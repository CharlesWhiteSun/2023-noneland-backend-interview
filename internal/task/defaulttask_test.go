package task_test

import (
	"nonelandBackendInterview/internal/task"
	"testing"
	"time"
)

func TestDefaultTask_Output_example(t *testing.T) {
	if testing.Short() {
		t.Skip("do not need run on short ver.")
	}
	tasks := task.NewTask()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			tasks.CommonTaskChan <- struct{}{}
		}
	}()

	go func() {
		for {
			time.Sleep(250 * time.Millisecond)
			tasks.UrgentTaskChan <- struct{}{}
		}
	}()

	taskInit := task.NewTaskInitiator()
	taskInit.SetHandler(tasks)
	taskInit.Handler.Deal()
}
