package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"phenikaa/infrastructure"
	"phenikaa/model"
	"phenikaa/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type AccessService interface {
	CreateToken(userId uint, role string) (*model.TokenDetail, error)
	CreateAuth(userID int, tokenDetail *model.TokenDetail) error
	DeleteAuth(uuid string) (int64, error)
	ExtractTokenMetadata(r *http.Request) (*model.AccessDetail, error)
}

type accessService struct {
	userService UserService
	db          *gorm.DB
}

var (
	infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	errLog  = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func (s *accessService) CreateToken(userId uint, role string) (*model.TokenDetail, error) {
	var err error

	// Create token details
	tokenDetail := &model.TokenDetail{}

	var profile model.Profile
	if userId != 0 {
		if err = s.db.Model(&model.Profile{}).Where("user_id = ?", userId).Preload("User").Find(&profile).Error; err != nil {
			return nil, err
		}
		tokenDetail.Username = profile.User.Username
	}

	tokenDetail.AtExpires = time.Now().Add(time.Hour * time.Duration(infrastructure.GetExtendAccessHour())).Unix()
	tokenDetail.AccessUUID = utils.PatternGet(userId) + uuid.NewV4().String()
	tokenDetail.RtExpires = time.Now().Add(time.Hour * time.Duration(infrastructure.GetExtendRefreshHour())).Unix()
	tokenDetail.RefreshUUID = utils.PatternGet(userId) + uuid.NewV4().String()

	// Create Access Token
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = tokenDetail.AccessUUID
	atClaims["username"] = tokenDetail.Username
	atClaims["user_id"] = userId
	atClaims["role"] = role
	atClaims["exp"] = tokenDetail.AtExpires

	_, tokenDetail.AccessToken, err = infrastructure.GetEncodeAuth().Encode(atClaims)
	if err != nil {
		return nil, err
	}

	// Create Resfresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tokenDetail.RefreshUUID
	rtClaims["username"] = tokenDetail.Username
	rtClaims["user_id"] = userId
	rtClaims["role"] = role
	rtClaims["exp"] = tokenDetail.RtExpires
	_, tokenDetail.RefreshToken, err = infrastructure.GetEncodeAuth().Encode(rtClaims)
	if err != nil {
		return nil, err
	}

	return tokenDetail, nil
}

func (s *accessService) CreateAuth(userID int, tokenDetail *model.TokenDetail) error {
	// converting Unix to UTC(to Time Object)
	accessToken := time.Unix(tokenDetail.AtExpires, 0)
	refreshToken := time.Unix(tokenDetail.RtExpires, 0)
	now := time.Now()

	if errAccess := infrastructure.
		GetRedisClient().
		Set(tokenDetail.AccessUUID, strconv.Itoa(userID), accessToken.Sub(now)).
		Err(); errAccess != nil {
		return errAccess
	}

	if errRefresh := infrastructure.
		GetRedisClient().
		Set(tokenDetail.RefreshUUID, strconv.Itoa(userID), refreshToken.Sub(now)).
		Err(); errRefresh != nil {
		return errRefresh
	}

	return nil
}

func (c *accessService) DeleteAuth(uuid string) (int64, error) {
	deleted, err := infrastructure.GetRedisClient().Del(uuid).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}

func (s *accessService) ExtractTokenMetadata(r *http.Request) (*model.AccessDetail, error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return nil, err
	}

	accessUUID, ok := claims["access_uuid"].(string)
	if !ok {
		errLog.Println("can't parse access uuid from token")
		return nil, errors.New("can't parse access uuid from token")
	}

	userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		errLog.Println(err)
		return nil, err
	}

	return &model.AccessDetail{
		AccessUUID: accessUUID,
		UserID:     int(userID),
	}, nil
}

func NewAccessService() AccessService {
	return &accessService{
		userService: NewUserService(),
		db:          infrastructure.GetDB(),
	}
}
