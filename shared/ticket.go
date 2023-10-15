package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"time"
)

// Initialize default config
// This stores all of your app's sessions
var store = session.New(session.Config{
	KeyLookup: "cookie:sco-t",
})

type ticket struct {
	currentContext context.Context
	anonymous      bool
	UserID         string
	ExpiresIn      time.Duration
	IssuedAt       int64
	FistName       string
	LastName       string
	Email          string
}

type Ticket interface {
	GetContext() context.Context
	IsAuthenticated() bool
	GetUserID() string
	SaveTicketToSession(c *fiber.Ctx) error
	Print() string
	DeleteTicketFromSession(c *fiber.Ctx) error
}

func (t *ticket) GetContext() context.Context {
	if t.currentContext == nil {
		panic("No context set")
	}
	return t.currentContext
}

func (t *ticket) IsAuthenticated() bool {
	fmt.Printf("Is Authenticated: %s %v \n", t.UserID, t.anonymous)
	if t.anonymous || len(t.UserID) == 0 {
		return false
	}
	return true
}

func (t *ticket) GetUserID() string {
	return t.UserID
}

func (t *ticket) SaveTicketToSession(c *fiber.Ctx) error {
	// Get the session from the store
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Save: %s\n", t.Print())

	// Set session values
	sess.Set("ticket", t.ToByteArray())

	// Set expiration
	sess.SetExpiry(t.ExpiresIn)

	// Save session
	if err := sess.Save(); err != nil {
		panic(err)
	}

	return nil
}

func (t *ticket) DeleteTicketFromSession(c *fiber.Ctx) error {
	// Get the session from the store
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Delete: %s\n", t.Print())

	// Delete session values
	sess.Delete("ticket")

	// Save session
	if err := sess.Save(); err != nil {
		panic(err)
	}

	return nil
}

func (t *ticket) Print() string {
	return fmt.Sprintf("Ticket: %s %s %s %s %v %v \n", t.UserID, t.FistName, t.LastName, t.Email, t.ExpiresIn, t.anonymous)
}

func (t *ticket) setCurrentContext(c *fiber.Ctx) {
	t.currentContext = c.Context()
}

func (t *ticket) ToByteArray() []byte {
	val, err := json.Marshal(*t)
	if err != nil {
		panic(err)
	}
	return val
}

func fromByteArray(b []byte) Ticket {
	t := &ticket{}
	if err := json.Unmarshal(b, t); err != nil {
		panic(err)
	}
	return t
}

func NewTicket(userID string, firstName string, lastName string, email string, expiresIn time.Duration, isAnonymous bool) Ticket {
	t := &ticket{
		UserID:    userID,
		FistName:  firstName,
		LastName:  lastName,
		Email:     email,
		ExpiresIn: expiresIn,
		anonymous: isAnonymous,
	}
	return t
}

func GetTicket(c *fiber.Ctx) Ticket {
	// Get the session from the store
	sess, err := store.Get(c)
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	// Get the ticket from the session
	ticketOut := sess.Get("ticket")
	if ticketOut == nil {
		fmt.Println("No ticket found")
		anonymousTicket := NewTicket(
			"",
			"",
			"",
			"",
			time.Until(time.Now().Add(10*time.Minute)),
			true,
		).(*ticket)
		anonymousTicket.anonymous = true
		anonymousTicket.setCurrentContext(c)
		return anonymousTicket
	} else {
		fmt.Println("Ticket found")
		// Reconstruct the ticket from the session. From []byte to Ticket
		ticketReconstructed := fromByteArray(ticketOut.([]byte))
		ticketReconstructed.(*ticket).setCurrentContext(c)
		fmt.Printf("Get: %s\n", ticketReconstructed.Print())
		return ticketReconstructed
	}
}
