package common

type Database interface {
	TryGetUnderlyingDatabaseConnection() interface{}
}

type DatabaseDriver interface {
	NewDatabase(name string) Database
}
