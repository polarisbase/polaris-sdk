package session

import (
	userCommon "github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/pointmass/gorm_sqllite"
	"gorm.io/gorm"
)

type Provider struct {
	sessions *gorm.DB
}

func (p *Provider) NewSession(user userCommon.User) (userCommon.Session, error) {
	session := NewBasic(user.GetID())
	res := p.sessions.Create(session)
	if res.Error != nil {
		return nil, res.Error
	}
	return session, nil
}

func (p *Provider) NewSessionFromToken(token string) (userCommon.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Provider) NewSessionFromRefreshToken(refreshToken string) (userCommon.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Provider) FindByID(id string) (userCommon.Session, error) {
	session := &BaseSession{}
	res := p.sessions.Where("id = ?", id).First(session)
	if res.Error != nil {
		return nil, res.Error
	}
	return session, nil
}

func NewProvider(unifiedGormDb *gorm.DB) *Provider {

	p := &Provider{}

	if unifiedGormDb != nil {
		p.sessions = unifiedGormDb
	} else {
		p.sessions = gorm_sqllite.AsGorm(
			gorm_sqllite.DatabaseDriver().NewDatabase("sessions"),
		)
	}

	if err := p.sessions.AutoMigrate(&BaseSession{}); err != nil {
		panic(err)
	}

	return p

}
