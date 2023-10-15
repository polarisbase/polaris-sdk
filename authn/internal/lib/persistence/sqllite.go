package persistence

import (
	"context"
	"github.com/FerretDB/FerretDB/ferretdb"
	"go.mongodb.org/mongo-driver/mongo"
)

type SqlLite interface {
	Close()
	Connect()
	Dbm() *mongo.Client
	GetFerretDB() *ferretdb.FerretDB
	GetContext() context.Context
}
