package scheduler

import (
	"backupManager/worker"
	"github.com/google/uuid"
	"time"
)

type ScheduledTask struct {
	Id        uuid.UUID
	Name      string
	Action    *worker.Task
	NextRun   time.Time
	RunOffset time.Duration
	timer     *time.Timer
}
