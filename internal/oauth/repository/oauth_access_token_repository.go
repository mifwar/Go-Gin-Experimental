package oauth

import (
	"gorm.io/gorm"
	entity "online-course.mifwar.com/internal/oauth/entity"
)

type OauthAccessTokenRepository interface {
	Create(oauthAccessToken entity.OauthAccessToken) (*entity.OauthAccessToken, error)
	Delete(id int) error
}

type OauthAccessTokenRepositoryImpl struct {
	db *gorm.DB
}

// Create implements OauthAccessTokenRepository
func (repository *OauthAccessTokenRepositoryImpl) Create(oauthAccessToken entity.OauthAccessToken) (*entity.OauthAccessToken, error) {
	if err := repository.db.Create(&oauthAccessToken).Error; err != nil {

		return nil, err
	}

	return &oauthAccessToken, nil
}

// Delete implements OauthAccessTokenRepository
func (repository *OauthAccessTokenRepositoryImpl) Delete(id int) error {
	var oauthAccessToken entity.OauthAccessToken

	if err := repository.db.Delete(&oauthAccessToken, id).Error; err != nil {
		return err
	}

	return nil
}

func NewOauthAccessTokenRepository(db *gorm.DB) OauthAccessTokenRepository {
	return &OauthAccessTokenRepositoryImpl{db}
}
