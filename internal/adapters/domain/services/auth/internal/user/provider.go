package user

import (
	userCommon "github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/auth/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/pointmass/gorm_sqllite"
	"gorm.io/gorm"
)

type Provider struct {
	sqllite *gorm_sqllite.Database
	users   *gorm.DB
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
		Email:          email,
		PasswordHashed: password,
	}

	// check if user exists
	_, err := p.FindByEmail(email)
	if err == nil {
		return nil, userCommon.ErrUserAlreadyExists
	}

	res := p.users.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func NewProvider() *Provider {

	p := &Provider{}

	p.sqllite = gorm_sqllite.DatabaseDriver()

	users := p.sqllite.NewDatabase("users")

	p.users = gorm_sqllite.AsGorm(users)

	if err := p.users.AutoMigrate(&BaseUser{}); err != nil {
		panic(err)
	}

	return p
}
