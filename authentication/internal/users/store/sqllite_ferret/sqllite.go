package sqllite_ferret

import (
	"context"
	"fmt"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/lib/persistence"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/users/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type debugLevel int

const (
	info    debugLevel = 1
	warning debugLevel = 2
	errors  debugLevel = 3
	crash   debugLevel = 4
)

type SqlLite struct {
	persistence persistence.SqlLite
	database    *mongo.Database
	users       *mongo.Collection
	debugger    []debugLevel
}

func NewSqlLite(persistence persistence.SqlLite) *SqlLite {

	s := &SqlLite{}

	s.debugger = []debugLevel{info, warning, errors, crash}

	s.persistence = persistence

	s.database = persistence.Dbm().Database("authentication")

	s.users = s.database.Collection("users")

	return s

}

func (s *SqlLite) debuggerPrintln(method string, state string, message string, level debugLevel) {

	for _, v := range s.debugger {
		if v == level {
			fmt.Printf("SqlLite.%s: %s: %s\n", method, state, message)
		}
	}

}

func (s *SqlLite) CheckIfEmailIsAlreadyInUse(ctx context.Context, email string) (err error, ok bool) {

	// Check if the email is already in use
	user := &model.User{}
	dbErr := s.users.FindOne(ctx, model.User{Email: email}).Decode(user)
	if dbErr != nil {
		s.debuggerPrintln("CheckIfEmailIsAlreadyInUse", "error", dbErr.Error(), errors)
		if dbErr.Error() == "mongo: no documents in result" {
			return nil, true
		} else {
			return dbErr, false
		}
	}

	// Return nil
	return nil, false
}

func (s *SqlLite) CreateUser(ctx context.Context, username string, password string) (err error, user model.User, ok bool) {

	// Create the user
	user = model.User{}
	user.Username = username
	user.BasicPasswordHash = password
	_, dbErr := s.users.InsertOne(ctx, user)
	if dbErr != nil {
		s.debuggerPrintln("CreateUser", "error", dbErr.Error(), errors)
		return dbErr, user, false
	}

	// Return nil
	return nil, user, true
}
