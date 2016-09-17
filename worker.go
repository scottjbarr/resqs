package resqs

import (
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Config struct {
}

type Worker struct {
	ID     string
	Config Config
}

type Handler interface {
	Perform() error
}

// NewWorker returns a Worker.
//
// The Worker will have an ID, but is has not yet regsitered.
func NewWorker(c Config) *Worker {
	return &Worker{
		ID:     uuid.NewV4().String(),
		Config: c,
	}
}

// Register adds this Worker to the list of active Workers.
func (s *Worker) Register() error {
	// add id to the worker registry

	// start a process to ping the worker registry

	// use etcd?

	return nil
}

func (s *Worker) Deregister() error {
	log.Printf("Deregistering Worker %v", s.ID)
	// stop the ping process

	// remove the id from the worker registry (this worker should just
	// disappear after timeout)

	return nil
}

func (s *Worker) Put(payload interface{}) (*Message, error) {
	return &Message{}, nil
}

// Get listens to the queue until a message is
func (s *Worker) Get(messages chan<- Message) error {
	s.Register()
	defer s.Deregister()

	for {
		time.Sleep(1 * time.Second)
		m := Message{
			ID: uuid.NewV4().String(),
		}
		messages <- m
	}

	return nil
}

func (s *Worker) Ack(id string) error {
	return nil
}

type Message struct {
	ID        string
	CreatedAt time.Time
}

func NewMessage() Message {
	return Message{
		ID:        uuid.NewV4().String(),
		CreatedAt: time.Now(),
	}
}
