//go:build wireinject
// +build wireinject

package oauth

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	oauthHandler "online-course.mifwar.com/internal/oauth/delivery/http"
	oauthRepository "online-course.mifwar.com/internal/oauth/repository"
	oauthUseCase "online-course.mifwar.com/internal/oauth/usecase"
	userRepository "online-course.mifwar.com/internal/user/repository"
	userUseCase "online-course.mifwar.com/internal/user/usecase"
)

func InitializedService(db *gorm.DB) *oauthHandler.OauthHandler {
	wire.Build(
		oauthHandler.NewOauthHandler,
		oauthRepository.NewOauthClientRepository,
		oauthRepository.NewOauthAccessTokenRepository,
		oauthRepository.NewOauthRefreshTokenRepository,
		oauthUseCase.NewOauthUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
	)

	return &oauthHandler.OauthHandler{}
}
