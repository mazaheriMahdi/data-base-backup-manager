package scheduler

import (
	"backupManager/worker"
	"log"
	"reflect"
	"time"
)

type Scheduler struct {
	tasks   []*ScheduledTask
	workers []*worker.Worker
	timers  []*time.Timer
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks:   make([]*ScheduledTask, 0),
		workers: make([]*worker.Worker, 0),
	}
}

func (s *Scheduler) AddTask(task *ScheduledTask) {
	s.tasks = append(s.tasks, task)
	s.runTask(task)
}
func (s *Scheduler) AddWorker(worker *worker.Worker) {
	s.workers = append(s.workers, worker)
}

func (s *Scheduler) Run() {
	for _, element := range s.tasks {
		go func() {
			s.runTask(element)
		}()
	}
}

func (s *Scheduler) runTask(task *ScheduledTask) {
	log.Printf("%v: Adding task with Name (%v) to worker queue", reflect.TypeOf(s).Elem(), task.Name)
	s.getWorker().AddTask(*task.Action)
	log.Printf("%v: Task with Name (%v) added to worker queue", reflect.TypeOf(s).Elem(), task.Name)
	task.NextRun = time.Now().Add(task.RunOffset)
	s.scheduleNextRun(task)
	log.Printf("%v: Task with Name (%v) scheduled for next run", reflect.TypeOf(s).Elem(), task.Name)
}

func (s *Scheduler) StopTask(task *ScheduledTask) {
	task.timer.Stop()
}

func (s *Scheduler) StopAll(task *ScheduledTask) {
	for _, element := range s.tasks {
		element.timer.Stop()
	}
}

func (s *Scheduler) scheduleNextRun(task *ScheduledTask) {
	timer := time.AfterFunc(task.RunOffset, func() {
		s.runTask(task)
	})
	task.timer = timer
}

func (s *Scheduler) getWorker() *worker.Worker {
	// In the future, it's going to choose the best worker
	return s.workers[0]
}
