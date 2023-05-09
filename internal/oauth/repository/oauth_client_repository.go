package oauth

import (
	"gorm.io/gorm"
	oauth "online-course.mifwar.com/internal/oauth/entity"
)

type OauthClientRepository interface {
	FindByClientIdAndClientSecret(clientId, clientSecret string) (*oauth.OauthClient, error)
}

type OauthClientRepositoryImpl struct {
	db *gorm.DB
}

// FindByClientIdAndClientSecret implements OauthClientRepository
func (repository *OauthClientRepositoryImpl) FindByClientIdAndClientSecret(clientId string, clientSecret string) (*oauth.OauthClient, error) {
	var oauthClient oauth.OauthClient

	if err := repository.db.Where("client_id = ?", clientId).Where("client_secret = ?", clientSecret).First(&oauthClient).Error; err != nil {
		return nil, err
	}

	return &oauthClient, nil
}

func NewOauthClientRepository(db *gorm.DB) OauthClientRepository {
	return &OauthClientRepositoryImpl{db}
}
