package scheduler

import (
	"backupManager/worker"
	"log"
	"reflect"
	"time"
)

type Scheduler struct {
	Tasks   []*ScheduledTask
	Workers []*worker.Worker
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Tasks:   make([]*ScheduledTask, 0),
		Workers: make([]*worker.Worker, 0),
	}
}

func (s *Scheduler) AddTask(task *ScheduledTask) {
	s.Tasks = append(s.Tasks, task)
}
func (s *Scheduler) AddWorker(worker *worker.Worker) {
	s.Workers = append(s.Workers, worker)
}

func (s *Scheduler) Run() {
	for _, element := range s.Tasks {
		go func() {
			s.runTask(element)
		}()
	}
}
func (s *Scheduler) ScheduleNextRun(task *ScheduledTask) {
	time.AfterFunc(task.RunOffset, func() {
		s.runTask(task)
	})
}

func (s *Scheduler) runTask(task *ScheduledTask) {
	log.Printf("%v: Adding task with Name (%v) to worker queue", reflect.TypeOf(s).Elem(), task.Name)
	s.getWorker().AddTask(task.Action)
	log.Printf("%v: Task with Name (%v) added to worker queue", reflect.TypeOf(s).Elem(), task.Name)
	task.NextRun = time.Now().Add(task.RunOffset)
	s.ScheduleNextRun(task)
	log.Printf("%v: Task with Name (%v) scheduled for next run", reflect.TypeOf(s).Elem(), task.Name)
}
func (s *Scheduler) getWorker() *worker.Worker {
	return s.Workers[0]
}
