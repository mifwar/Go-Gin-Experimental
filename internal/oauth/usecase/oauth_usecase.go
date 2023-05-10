package oauth

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	adminUseCase "online-course.mifwar.com/internal/admin/usecase"
	dto "online-course.mifwar.com/internal/oauth/dto"
	entity "online-course.mifwar.com/internal/oauth/entity"
	repository "online-course.mifwar.com/internal/oauth/repository"
	userUseCase "online-course.mifwar.com/internal/user/usecase"
	"online-course.mifwar.com/pkg/utils"
)

type OauthUseCase interface {
	Login(loginRequestBody dto.LoginRequestBody) (*dto.LoginResponse, error)
	Refresh(refreshTokenRequestBody dto.RefreshTokenRequestBody) (*dto.LoginResponse, error)
}

type OauthUseCaseImpl struct {
	oauthClientRepository       repository.OauthClientRepository
	oauthAccessTokenRepository  repository.OauthAccessTokenRepository
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository
	userUseCase                 userUseCase.UserUseCase
	adminUseCase                adminUseCase.AdminUseCase
}

// Login implements OauthUseCase
func (usecase *OauthUseCaseImpl) Login(loginRequestBody dto.LoginRequestBody) (*dto.LoginResponse, error) {
	oauthClient, err := usecase.oauthClientRepository.FindByClientIdAndClientSecret(loginRequestBody.ClientId, loginRequestBody.ClientSecret)

	if err != nil {
		return nil, err
	}

	var user dto.UserResponse

	//login using admin
	if oauthClient.Name == "web-admin" {
		dataAdmin, err := usecase.adminUseCase.FindByEmail(loginRequestBody.Email)

		if err != nil {
			return nil, errors.New("email or password is invalid")
		}

		user.ID = dataAdmin.ID
		user.Email = dataAdmin.Email
		user.Password = dataAdmin.Password
	} else {
		// Login menggunakan user
		dataUser, err := usecase.userUseCase.FindByEmail(loginRequestBody.Email)

		if err != nil {
			return nil, errors.New("username or password is invalid")
		}

		user.ID = dataUser.ID
		user.Name = dataUser.Name
		user.Email = dataUser.Email
		user.Password = dataUser.Password
	}

	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequestBody.Password))

	if err != nil {
		return nil, errors.New("username or password is invalid")
	}

	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	claims := &dto.ClaimsResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	if oauthClient.Name != "web-admin" {
		claims.IsAdmin = false
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	dataOauthAccessToken := entity.OauthAccessToken{
		OauthClientID: &oauthClient.ID,
		UserID:        user.ID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt: sql.NullTime{
			Time: expirationTime,
		},
	}

	oauthAccessToken, err := usecase.oauthAccessTokenRepository.Create(dataOauthAccessToken)

	if err != nil {
		return nil, err
	}

	dataOauthRefreshToken := entity.OauthRefreshToken{
		OauthAccessTokenID: &oauthAccessToken.ID,
		UserID:             user.ID,
		Token:              utils.GenerateRandomString(128),
		ExpiredAt: sql.NullTime{
			Time: time.Now().Add(24 * 366 * time.Hour),
		},
	}

	oauthRefreshToken, err := usecase.oauthRefreshTokenRepository.Create(dataOauthRefreshToken)

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  tokenString,
		RefreshToken: oauthRefreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil

}

// Refresh implements OauthUseCase
func (*OauthUseCaseImpl) Refresh(refreshTokenRequestBody dto.RefreshTokenRequestBody) (*dto.LoginResponse, error) {
	panic("unimplemented")
}

func NewOauthUseCase(
	oauthClientRepository repository.OauthClientRepository,
	oauthAccessTokenRepository repository.OauthAccessTokenRepository,
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository,
	userUseCase userUseCase.UserUseCase,
	adminUseCase adminUseCase.AdminUseCase,
) OauthUseCase {
	return &OauthUseCaseImpl{
		oauthClientRepository,
		oauthAccessTokenRepository,
		oauthRefreshTokenRepository,
		userUseCase,
		adminUseCase,
	}

}
