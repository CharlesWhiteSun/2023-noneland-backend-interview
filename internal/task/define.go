package task

// ITaskHandler 各調度類型任務需要實作的介面
type ITaskHandler interface {
	Deal()
}

var taskInit TaskInitiator

type TaskInitiator struct {
	Handler ITaskHandler
}

// NewTaskInitiator 返回主要任務註冊的初始物件
func NewTaskInitiator() *TaskInitiator {
	return &taskInit
}

func (t *TaskInitiator) SetHandler(it ITaskHandler) {
	t.Handler = it
}
