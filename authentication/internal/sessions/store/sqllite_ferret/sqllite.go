package sqllite_ferret

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/lib/persistence"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/sessions/model"
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
	sessions    *mongo.Collection
	debugger    []debugLevel
}

func NewSqlLite(persistence persistence.SqlLite) *SqlLite {

	s := &SqlLite{}

	s.debugger = []debugLevel{info, warning, errors_, crash}

	s.persistence = persistence

	s.database = persistence.Dbm().Database("session")

	s.sessions = s.database.Collection("sessions")

	return s

}

func (s *SqlLite) debuggerPrintln(method string, state string, message string, level debugLevel) {

	for _, v := range s.debugger {
		if v == level {
			fmt.Printf("SqlLite.%s: %s: %s\n", method, state, message)
		}
	}

}

func (s *SqlLite) CreateSession(ctx context.Context, session model.Session) (err error, sessionOut model.Session, ok bool) {

	// Set the session ID
	session.ID = uuid.New().String()

	// Create the session
	_, err = s.sessions.InsertOne(ctx, session)
	if err != nil {
		s.debuggerPrintln("CreateSession", "InsertOne", err.Error(), errors_)
		return err, session, false
	}

	return nil, session, true
}
