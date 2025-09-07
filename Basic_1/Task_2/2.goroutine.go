package Task_2

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func coroutinePrint() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Printf("1. 奇数协程: %d\n", i)
				//ch1 <- i
			}
		}
	}()
	go func() {
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Printf(" 2. 偶数协程: %d\n", i)
			}
		}
	}()
}

// 题目 2
type Task func() error

type TaskResult struct {
	ID        string
	Error     error
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
}

type TaskScheduler struct {
	tasks   map[string]Task
	results map[string]TaskResult
	mutex   sync.RWMutex
}

func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:   make(map[string]Task),
		results: make(map[string]TaskResult),
	}
}

func (ts *TaskScheduler) AddTask(id string, task Task) {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	ts.tasks[id] = task
}

func (ts *TaskScheduler) ExecuteAll(ctx context.Context) map[string]TaskResult {
	ts.mutex.Lock()
	taskCount := len(ts.tasks)
	if taskCount == 0 {
		ts.mutex.Unlock()
		return ts.results
	}

	var wg sync.WaitGroup
	resultChan := make(chan TaskResult, taskCount)

	// 启动所有任务
	for id, task := range ts.tasks {
		wg.Add(1)
		go func(taskID string, taskFunc Task) {
			defer wg.Done()

			startTime := time.Now()
			err := taskFunc()
			endTime := time.Now()

			result := TaskResult{
				ID:        taskID,
				Error:     err,
				StartTime: startTime,
				EndTime:   endTime,
				Duration:  endTime.Sub(startTime),
			}

			// 尝试发送结果，但不要阻塞太久
			select {
			case resultChan <- result:
			case <-ctx.Done():
				// 上下文取消，但仍需发送结果
				result.Error = ctx.Err()
				resultChan <- result
			}
		}(id, task)
	}
	ts.mutex.Unlock()

	// 等待所有任务完成
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// 收集结果
	for {
		select {
		case result := <-resultChan:
			ts.mutex.Lock()
			ts.results[result.ID] = result
			ts.mutex.Unlock()
		case <-done:
			// 所有任务完成
			close(resultChan)
			return ts.results
		case <-ctx.Done():
			wg.Wait()
			close(resultChan)
			return ts.results
		}
	}
}

// GetResult 根据任务 ID 获取任务执行结果
// params：
//
//	id - 任务的唯一标识符
//
// return：
//
//	TaskResult - 任务执行结果
//	bool - 任务结果是否存在， true 表示存在，false 表示不存在
func (ts *TaskScheduler) GetResult(id string) (TaskResult, bool) {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()
	result, exists := ts.results[id]
	return result, exists
}

// GetAllResults 获取所有任务执行结果
func (ts *TaskScheduler) GetAllResults() map[string]TaskResult {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()

	results := make(map[string]TaskResult)
	for id, result := range ts.results {
		results[id] = result
	}
	return results
}
