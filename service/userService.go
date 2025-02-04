package service

import (
	"bytes"
	"fmt"
	"log"
	"phenikaa/infrastructure"
	"phenikaa/model"
	"phenikaa/utils"
	"phenikaa/utils/emailTemplate"
	"strings"
	"text/template"

	"github.com/golang/glog"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserService interface {
	CheckCredentials(username string, password string) (*model.UserResponse, error)
	GetByUsername(username string) (*model.UserResponse, error)
	CreateUser(newUser model.RegisterPayload) (*model.User, error)
	UpdateUserState(userId int32, state bool) error
	UpdateUserRole(listUser []model.UpdateUserStatePayload) error
	ResetPassword(username string) error
	ChangePassword(payload model.ChangePasswordPayload) error
	ForgotPassword(payload model.ForgotPasswordPayload) error
	CheckEmailExact(username string) error
	GetAllUsers() ([]*model.UserSystemResponse, error)
}

type userService struct {
	emailService EmailService
	db           *gorm.DB
}

func (s *userService) CheckCredentials(username string, password string) (*model.UserResponse, error) {
	var userResponse *model.UserResponse
	var user model.User

	if err := s.db.Debug().Where("username = ?", username).Preload("UserRoles.Role").First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	if !comparePassword(user.Password, password) {
		err := fmt.Errorf("password mismatch")
		glog.V(3).Infof(" - CheckCredentials() has err:", err)
		return nil, err
	}

	// ----------------------------------------------------------------
	var profile *model.Profile
	if err := s.db.Where("user_id = ?", user.ID).Find(&profile).Error; err != nil {
		return nil, err
	}

	userResponse = &model.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		FullName: profile.FirstName + " " + profile.LastName,
		Role:     user.UserRoles.Role.Type,
		Profile:  profile,
	}

	return userResponse, nil
}

func (s *userService) GetByUsername(username string) (*model.UserResponse, error) {
	var userResponse model.UserResponse
	var user model.User
	if err := s.db.Model(&model.User{}).Where("username = ?", username).
		Preload("UserRoles.Role").
		First(&user).Error; err != nil {
		return nil, err
	}
	var profile *model.Profile
	if err := s.db.Model(&model.Profile{}).Where("user_id = ?", user.ID).Find(&profile).Error; err != nil {
		return nil, err
	}

	userResponse.ID = user.ID
	userResponse.Username = user.Username
	userResponse.FullName = profile.FirstName + " " + profile.LastName
	userResponse.Role = user.UserRoles.Role.Type
	userResponse.Profile = profile

	return &userResponse, nil
}

func (s *userService) CreateUser(newUser model.RegisterPayload) (*model.User, error) {
	var userInfo model.User
	user := model.User{
		Username: newUser.Username,
		Password: hashAndSalt(newUser.Password),
	}

	queryGetMaxId := "SELECT setval('users_id_seq', (SELECT MAX(id) FROM users)+1);"
	if err := s.db.Model(&model.User{}).Raw(queryGetMaxId).Error; err != nil {
		return nil, fmt.Errorf("set max id error: %v", err)
	}

	if err := s.db.Debug().Transaction(func(tx *gorm.DB) error {
		if err := s.db.Model(&user).Clauses(clause.Returning{}).
			Create(&user).Error; err != nil {
			return err
		}

		if err := s.db.Model(&model.UserRole{}).Create(&model.UserRole{
			UserID: user.ID,
			RoleID: infrastructure.GetClientRole(), // Default role is 1 (client)
			Active: true,
		}).Error; err != nil {
			return err
		}

		// take fisrt name form payload
		strArr := strings.Split(newUser.FullName, " ")
		firstName := strings.Join(strArr[0:len(strArr)-1], " ")

		var lastName string
		if len(strArr) > 1 {
			lastName = strings.Join(strArr[len(strArr)-1:], " ")
		} else {
			lastName = ""
		}

		if err := s.db.Model(&model.Profile{}).Create(&model.Profile{
			UserID:    user.ID,
			FirstName: firstName,
			LastName:  lastName,
			Email:     newUser.Email,
			Phone:     newUser.Phone,
		}).Error; err != nil {
			return err
		}

		if err := s.db.Model(&model.User{}).Where("id = ?", user.ID).Preload("UserRoles.Role").Preload("Profile").First(&userInfo).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	// Thông tin người dùng
	return &userInfo, nil
}

// Sử dụng mật khẩu mặc định phenikaa@123
func (s *userService) ResetPassword(username string) error {
	var user model.User
	if err := s.db.Model(&user).Where("username = ?", username).
		Update("password", hashAndSalt(model.DefaultPassword)).Error; err != nil {
		return err
	}
	return nil
}

func (s *userService) ChangePassword(payload model.ChangePasswordPayload) error {
	_, err := s.CheckCredentials(payload.Username, payload.OldPassword)
	if err != nil {
		return fmt.Errorf("Worng username or password: %v", err)
	}

	var user model.User
	if err := s.db.Model(&user).Where("username = ?", payload.Username).
		Update("password", hashAndSalt(payload.NewPassword)).Error; err != nil {
		return err
	}
	return nil
}

func (s *userService) ForgotPassword(payload model.ForgotPasswordPayload) error {
	var userForgotPassword model.UserForgotPassword
	if err := s.db.Model(&model.UserForgotPassword{}).Where("fogot_code = ?", payload.FogortCode).First(&userForgotPassword).Error; err != nil {
		return fmt.Errorf("code not exist: %v", err)
	}

	var user model.User
	if err := s.db.Model(&model.User{}).Where("id = ?", userForgotPassword.UserId).First(&user).Error; err != nil {
		return fmt.Errorf("user not exist: %v", err)
	}

	if err := s.db.Model(&model.User{}).Where("id = ?", userForgotPassword.UserId).
		Update("password", hashAndSalt(payload.NewPassword)).Error; err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateUserState(userId int32, state bool) error {
	if err := s.db.Model(&model.UserRole{}).Where("user_id = ?", userId).
		Update("active", state).Error; err != nil {
		return err
	}
	return nil
}

func (s *userService) UpdateUserRole(listUser []model.UpdateUserStatePayload) error {

	glog.V(3).Info("UpdateUserRole started")
	var roles []model.Role
	if err := s.db.Find(&roles).Error; err != nil {
		glog.V(3).Info("UpdateUserRole ERROR: %v", err)
		return err
	}

	glog.V(3).Info("UpdateUserRole started2")
	var listUserUpdate []model.UserRole
	for _, user := range listUser {
		roleId := int32(0)
		for _, role := range roles {
			if role.Code == user.Role {
				roleId = role.ID
				break
			}
		}
		if roleId == 0 {
			glog.V(3).Info("UpdateUserRole ERROR: User: %d - Role not exist: %s", user.ID, user.Role)
			continue
		}
		listUserUpdate = append(listUserUpdate, model.UserRole{
			UserID: user.ID,
			RoleID: roleId,
		})
	}

	// update roles
	glog.V(3).Info("UpdateUserRole started3")
	if err := s.db.Debug().Model(&model.UserRole{}).Save(listUserUpdate).Error; err != nil {
		glog.V(3).Info("UpdateUserRole ERROR: %v", err)
		return err
	}

	return nil
}

func hashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePassword(hashedPwd string, plainPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd)); err != nil {
		return false
	}
	return true
}

func (s *userService) CheckEmailExact(email string) error {
	var profile model.Profile
	emailExistErr := s.db.Model(&model.Profile{}).Where("email = ?", email).First(&profile).Error
	if emailExistErr != nil {
		return fmt.Errorf("Email not exist: %v", emailExistErr)
	}
	newCode, err := utils.GeneratePasswordKey(7)
	if err != nil {
		return fmt.Errorf("Generate code error: %v", err)
	}
	emailDataSend := model.EmailForgotPassword{
		FogortCode: newCode,
	}

	// Tạo template từ HTML
	emailHtmlTemplate := emailTemplate.ForgotPasswordTemplate
	tmpl, err := template.New("emailForgotTemplate").Parse(emailHtmlTemplate)
	if err != nil {
		fmt.Println("Error parsing HTML template:", err)
		return err
	}

	// Tạo buffer để lưu nội dung HTML đã render
	var tplBuffer bytes.Buffer
	err = tmpl.Execute(&tplBuffer, emailDataSend)
	if err != nil {
		return fmt.Errorf("Execute template error: %v", err)
	}

	// Chuyển đổi nội dung buffer thành chuỗi
	htmlBody := tplBuffer.String()
	errSendMail := s.emailService.SendEmail([]string{profile.Email}, "Xác Thực Quên Mật Khẩu Hệ Thống - Phenikaa Intern", htmlBody)
	if errSendMail != nil {
		return fmt.Errorf("Send email error: %v", errSendMail)
	}

	// Lưu code vào database
	newForgotCode := model.UserForgotPassword{
		UserId:    profile.UserID,
		FogotCode: newCode,
	}

	// Kiểm tra xem user đã có code chưa
	if err := s.db.Model(&model.UserForgotPassword{}).Where("user_id =?", profile.UserID).First(&model.UserForgotPassword{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := s.db.Model(&model.UserForgotPassword{}).Create(&newForgotCode).Error; err != nil {
				return fmt.Errorf("Create code error: %v", err)
			}
		} else {
			return fmt.Errorf("Check code error: %v", err)
		}
	} else {
		if err := s.db.Model(&model.UserForgotPassword{}).Where("user_id = ?", profile.UserID).Update("fogot_code", newCode).Error; err != nil {
			return fmt.Errorf("Update code error: %v", err)
		}
	}
	return nil
}

func (s *userService) GetAllUsers() ([]*model.UserSystemResponse, error) {
	var users []model.User
	if err := s.db.Preload("UserRoles.Role").Preload("Profile").Find(&users).Error; err != nil {
		return nil, fmt.Errorf("get all users error: %v", err)
	}

	var userResponses []*model.UserSystemResponse
	for _, User := range users {
		userSystem := &model.UserSystemResponse{
			ID:       User.ID,
			Username: User.Username,
			Role:     User.UserRoles.Role.Type,
			FullName: User.Profile.FirstName + " " + User.Profile.FirstName,
		}

		if !User.UserRoles.Active {
			userSystem.IsBaned = true
		} else {
			userSystem.IsBaned = false
		}

		userResponses = append(userResponses, userSystem)
	}
	return userResponses, nil
}

// Ở đấy hàm NewUserService() trả về một interface UserService
// sử dụng design pattern là Factory Method với mục đích tạo ra một đối tượng UserService mà không cần biết cụ thể nó là kiểu dữ liệu gì
func NewUserService() UserService {
	emailService := NewEmailService()
	return &userService{
		emailService: emailService,
		db:           infrastructure.GetDB(),
	}
}
