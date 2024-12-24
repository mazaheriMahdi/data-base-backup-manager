package worker

import (
	"github.com/google/uuid"
	"log"
	"sync"
)

type Worker struct {
	queue           []Task
	isRunning       bool
	Id              uuid.UUID
	popLock         sync.Mutex
	pushLock        sync.Mutex
	starterTaskChan chan bool
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
		log.Printf("Wroker With Id %v started Successfully\n", w.Id)
		for w.isRunning {
			if len(w.queue) == 0 {
				log.Printf("Worker With Id %v is in idle mode\n", w.Id)
				<-w.starterTaskChan
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
	go func() {
		w.starterTaskChan <- true
	}()
}
