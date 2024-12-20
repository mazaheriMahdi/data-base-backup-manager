package worker

import (
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

type Worker struct {
	queue     []Task
	isRunning bool
	Id        uuid.UUID
	popLock   sync.Mutex
	pushLock  sync.Mutex
}

func NewWorker() *Worker {
	return &Worker{
		queue:     make([]Task, 0),
		isRunning: false,
		Id:        uuid.New(),
		popLock:   sync.Mutex{},
		pushLock:  sync.Mutex{},
	}
}

func (w *Worker) Start() {
	w.isRunning = true
	go func() {
		for w.isRunning {
			if len(w.queue) == 0 {
				log.Println("Worker queue is empty.")
				time.Sleep(10 * time.Second)
				continue
			}
			task := w.queue[0]
			log.Printf("Executing task with id (%v)\n", task.Id)
			err := task.Action()
			if err != nil {
				log.Printf("Execution of task with id (%v) failed: %v\n", task.Id, err)
			} else {
				log.Printf("Execution of task with id (%v) finished.\n", task.Id)
			}
			w.pop()
		}
	}()
}
func (w *Worker) pop() {
	w.popLock.Lock()
	w.queue = w.queue[1:]
	w.popLock.Unlock()
}
func (w *Worker) AddTask(task Task) {
	w.pushLock.Lock()
	w.queue = append(w.queue, task)
	w.pushLock.Unlock()
}
