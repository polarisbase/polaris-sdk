package sqllite_ferret

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/lib/persistence"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type debugLevel int

const (
	info    debugLevel = 1
	warning debugLevel = 2
	errors_ debugLevel = 3
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

	s.debugger = []debugLevel{info, warning, errors_, crash}

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
	filter := bson.D{{"email", email}}
	res := s.users.FindOne(ctx, filter)
	if res != nil {
		if res.Err() != nil {
			if errors.Is(res.Err(), mongo.ErrNoDocuments) {
				return nil, true
			} else {
				return res.Err(), false
			}
		}
		return nil, false
	} else {
		return nil, true
	}
}

func (s *SqlLite) CreateUser(ctx context.Context, email string, passwordHash string, profileData map[string]string) (err error, user model.User, ok bool) {

	// Create the user
	user = model.User{}
	user.ID = uuid.New().String()
	user.Email = email

	user.BasicPasswordHash = passwordHash

	// Set the profile data
	for k, v := range profileData {
		switch k {
		case "first_name":
			user.FirstName = v
		case "last_name":
			user.LastName = v
		}
	}

	// Insert the user
	_, dbErr := s.users.InsertOne(ctx, user)
	if dbErr != nil {
		s.debuggerPrintln("CreateUser", "error", dbErr.Error(), errors_)
		return dbErr, user, false
	}

	// Return nil
	return nil, user, true
}

func (s *SqlLite) GetUsers(ctx context.Context, limit int, offset int) (err error, users []model.User, ok bool) {

	// Get the users
	cursor, dbErr := s.users.Find(ctx, bson.D{}, nil)
	if dbErr != nil {
		s.debuggerPrintln("GetUsers", "error", dbErr.Error(), errors_)
		return dbErr, users, false
	}

	// Iterate through the cursor
	for cursor.Next(ctx) {
		// Create the user
		user := model.User{}
		// Decode the user
		dbErr = cursor.Decode(&user)
		if dbErr != nil {
			s.debuggerPrintln("GetUsers", "error", dbErr.Error(), errors_)
			return dbErr, users, false
		}
		// Append the user
		users = append(users, user)
	}

	// Return nil
	return nil, users, true
}

func (s *SqlLite) FindUserByEmail(ctx context.Context, email string) (err error, user model.User, ok bool) {
	// Get the user
	filter := bson.D{{"email", email}}
	res := s.users.FindOne(ctx, filter)
	if res != nil {
		if res.Err() != nil {
			if errors.Is(res.Err(), mongo.ErrNoDocuments) {
				return nil, user, false
			} else {
				return res.Err(), user, false
			}
		}
		// Decode the user
		dbErr := res.Decode(&user)
		if dbErr != nil {
			s.debuggerPrintln("FindUserByEmail", "error", dbErr.Error(), errors_)
			return dbErr, user, false
		}
		return nil, user, true
	} else {
		return nil, user, false
	}
}
