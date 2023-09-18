package user

import (
	"github.com/google/uuid"
	userCommon "github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/pointmass/gorm_sqllite"
	"gorm.io/gorm"
)

type Provider struct {
	users *gorm.DB
}

func (p *Provider) FindByEmail(email string) (userCommon.User, error) {
	user := &BaseUser{}
	res := p.users.Where("email = ?", email).First(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (p *Provider) FindByID(id string) (userCommon.User, error) {
	user := &BaseUser{}
	res := p.users.Where("id = ?", id).First(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (p *Provider) FindByUsername(username string) (userCommon.User, error) {
	user := &BaseUser{}
	res := p.users.Where("username = ?", username).First(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (p *Provider) NewBasic(email string, password string) (userCommon.User, error) {
	user := &BaseUser{
		ID: uuid.New().String(),
	}

	user.SetEmail(email)

	user.SetPassword(password)

	// check if user exists
	_, err := p.FindByEmail(email)
	if err == nil {
		return nil, userCommon.PossibleErrors.UserAlreadyExists
	}

	res := p.users.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func NewProvider(unifiedGormDb *gorm.DB) *Provider {

	p := &Provider{}

	if unifiedGormDb != nil {
		p.users = unifiedGormDb
	} else {
		p.users = gorm_sqllite.AsGorm(
			gorm_sqllite.DatabaseDriver().NewDatabase("users"),
		)
	}

	if err := p.users.AutoMigrate(&BaseUser{}); err != nil {
		panic(err)
	}

	return p
}
