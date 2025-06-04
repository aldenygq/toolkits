package cronjob

import "context"

type Task struct {
	ID       string
	Name     string
	Schedule string
	JobFunc  func(ctx context.Context)
	Retry    int
}
(mye
