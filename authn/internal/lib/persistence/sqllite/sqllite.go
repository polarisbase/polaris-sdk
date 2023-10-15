package sqllite

import (
	"context"
	"github.com/FerretDB/FerretDB/ferretdb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type SqlLite struct {
	inMem  bool
	path   string
	f      *ferretdb.FerretDB
	ctx    context.Context
	cancel context.CancelFunc
	client *mongo.Client
}

func New(inMem bool) *SqlLite {

	s := &SqlLite{
		inMem: inMem,
		path:  "db",
	}

	return s

}

func (s *SqlLite) Close() {
	s.disconnectClient()
	s.cancel()
}

func (s *SqlLite) Connect() {

	pathToUse := "file:" + s.path
	if s.inMem {
		pathToUse = "file::memory:?cache=shared"
	} else {
		// Check and ensure that the directory exists for the database 'db'
		// If it does not exist, create it
		os.MkdirAll(s.path, os.ModePerm)
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())

	f, err := ferretdb.New(&ferretdb.Config{
		Listener: ferretdb.ListenerConfig{
			TCP: "127.0.0.1:17027",
		},
		Handler: "sqlite",
		//PostgreSQLURL: "postgres://127.0.0.1:5432/ferretdb",
		SQLiteURL: pathToUse,
	})
	if err != nil {
		panic(err)
		log.Fatal(err)
	}

	s.f = f

	go func() {
		log.Print("Starting FerretDB")
		if err := s.f.Run(s.ctx); err != nil {
			log.Fatal(err)
		}
		log.Print("FerretDB stopped")
	}()

	s.connectClient()

}

func (s *SqlLite) connectClient() {

	uri := s.f.MongoDBURI()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	s.client = client

}

func (s *SqlLite) disconnectClient() {

	err := s.client.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}

}

func (s *SqlLite) Dbm() *mongo.Client {
	return s.client
}

func (s *SqlLite) GetFerretDB() *ferretdb.FerretDB {
	return s.f
}

func (s *SqlLite) GetContext() context.Context {
	return s.ctx
}
