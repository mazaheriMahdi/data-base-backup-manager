package worker

import "github.com/google/uuid"

type Task struct {
	Id     uuid.UUID
	Action func() error
}
