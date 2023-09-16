package session

import (
	"github.com/google/uuid"
	userCommon "github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/pointmass/gorm_sqllite"
	"gorm.io/gorm"
)

type Provider struct {
	sqllite  *gorm_sqllite.Database
	sessions *gorm.DB
}

func (p *Provider) NewSession(user userCommon.User) (userCommon.Session, error) {
	session := &BaseSession{
		ID:     uuid.New().String(),
		UserID: user.GetID(),
	}
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

func (p *Provider) FindByToken(token string) (userCommon.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Provider) FindByRefreshToken(refreshToken string) (userCommon.Session, error) {
	//TODO implement me
	panic("implement me")
}

func NewProvider() *Provider {

	p := &Provider{}

	p.sqllite = gorm_sqllite.DatabaseDriver()

	sessions := p.sqllite.NewDatabase("sessions")

	p.sessions = gorm_sqllite.AsGorm(sessions)

	if err := p.sessions.AutoMigrate(&BaseSession{}); err != nil {
		panic(err)
	}

	return p

}
