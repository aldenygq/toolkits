package cronjob

import (
	"context"
	"sync"

	cron "github.com/robfig/cron/v3"
)

type Scheduler struct {
	c      *cron.Cron
	mu     sync.Mutex
	tasks  map[string]cron.EntryID
	funcs  map[string]*Task
	status map[string]string
}

// 初始化
func NewScheduler() *Scheduler {
	return &Scheduler{
		c:      cron.New(cron.WithSeconds()),
		tasks:  make(map[string]cron.EntryID),
		funcs:  make(map[string]*Task),
		status: make(map[string]string),
	}
}

// 开始
func (s *Scheduler) Start() {
	s.c.Start()
}

// 结束
func (s *Scheduler) Stop() {
	s.c.Stop()
}

// 添加任务
func (s *Scheduler) AddTask(t *Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	id, err := s.c.AddFunc(t.Schedule, func() {
		ctx := context.Background()
		t.JobFunc(ctx)
	})
	if err != nil {
		return err
	}

	s.tasks[t.ID] = id
	s.funcs[t.ID] = t
	s.status[t.ID] = "running"
	return nil
}

// 立即执行
func (s *Scheduler) RunNow(taskID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if t, ok := s.funcs[taskID]; ok {
		go t.JobFunc(context.Background())
	}
}

// 暂停任务
func (s *Scheduler) PauseTask(taskID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if id, ok := s.tasks[taskID]; ok {
		s.c.Remove(id)
		s.status[taskID] = "paused"
	}
}

// 重试任务
func (s *Scheduler) ResumeTask(taskID string) error {
	if t, ok := s.funcs[taskID]; ok {
		return s.AddTask(t)
	}
	return nil
}

func (s *Scheduler) RetryTask(taskID string) {
	s.RunNow(taskID) // 简单实现，等同于立即执行
}

// 任务状态
func (s *Scheduler) Status(taskID string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.status[taskID]
}
