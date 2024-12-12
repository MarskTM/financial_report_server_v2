package service

import (
	"bytes"
	"fmt"
	"log"
	"phenikaa/infrastructure"
	"phenikaa/model"
	"phenikaa/utils"
	"phenikaa/utils/emailTemplate"
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
	BanUser(username string) error
	ResetPassword(username string) error
	ChangePassword(payload model.ChangePasswordPayload) error
	ForgotPassword(payload model.ForgotPasswordPayload) error
	CheckEmailExact(username string) error
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

		if err := s.db.Model(&model.Profile{}).Create(&model.Profile{
			UserID:    user.ID,
			FirstName: newUser.FirstName,
			LastName:  newUser.LastName,
			Email:     newUser.Email,
			Phone:     newUser.Phone,
		}).Error; err != nil {
			return err
		}

		if err := s.db.Model(&model.User{}).Where("id = ?", user.ID).Preload("UserRoles.Role").First(&userInfo).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	// Thông tin người dùng
	userInfo.Password = "********"
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

func (s *userService) BanUser(username string) error {
	if err := s.db.Model(&model.User{}).Where("username = ?", username).
		Update("active", false).Error; err != nil {
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

// Ở đấy hàm NewUserService() trả về một interface UserService
// sử dụng design pattern là Factory Method với mục đích tạo ra một đối tượng UserService mà không cần biết cụ thể nó là kiểu dữ liệu gì
func NewUserService() UserService {
	emailService := NewEmailService()
	return &userService{
		emailService: emailService,
		db:           infrastructure.GetDB(),
	}
}
