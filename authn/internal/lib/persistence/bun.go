package persistence

type Bun interface {
	Close()
	Connect()
}
