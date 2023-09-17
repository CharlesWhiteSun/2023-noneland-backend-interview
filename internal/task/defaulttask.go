package task

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const qps int = 300 * 2

var taskDefaultOnce sync.Once
var task TaskDefault

type TaskDefault struct {
	UrgentTaskChan chan struct{}
	CommonTaskChan chan struct{}
}

func NewTask() *TaskDefault {
	taskDefaultOnce.Do(func() {
		task.UrgentTaskChan = make(chan struct{}, qps)
		task.CommonTaskChan = make(chan struct{}, qps)
	})
	return &task
}

// 寬鬆記數
// - 數值越高表示容忍次數越大，反之越低越不能容忍
const (
	CommonLooseCount int32 = 5
	UrgentLooseCount int32 = 0
)

// Deal 預設的任務處理器
// - 緊急任務收到則直接消化完畢
// - 普通任務累積到上限再分批消化完畢
func (t *TaskDefault) Deal() {
	var commonCnt int32 = 0
	var urgentCnt int32 = 0
	var commonUseCnt int = 3 // 每次處理指定次數的任務

	for {
		select {
		case <-t.CommonTaskChan:
			// 優先處理緊急任務
			if atomic.LoadInt32(&urgentCnt) > UrgentLooseCount {
				// do urgent api request
				fmt.Println(" [Common] 緊急任務處理中...")
				atomic.AddInt32(&urgentCnt, -1)
				fmt.Println(" [Common] 緊急任務處理完畢!")
			}

			// 普通任務累積到上限再分批消化完畢
			atomic.AddInt32(&commonCnt, 1)
			if atomic.LoadInt32(&commonCnt) > CommonLooseCount {
				for atomic.LoadInt32(&commonCnt) > CommonLooseCount {
					for i := commonUseCnt; i > 0; i-- {
						// do common api request
						fmt.Println(" [Common] 普通任務處理中...")
						atomic.AddInt32(&commonCnt, -1)
						fmt.Println(" [Common] 普通任務處理完畢!")
					}
				}
			}

		case <-t.UrgentTaskChan:
			fmt.Println(" [Urgent] 緊急任務處理中...")
			atomic.AddInt32(&urgentCnt, 1)
			// do urgent api request
			atomic.AddInt32(&urgentCnt, -1)
			fmt.Println(" [Urgent] 緊急任務處理完畢!")
		}
	}
}
