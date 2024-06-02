package gorm_ex

import (
	"context"
	"github.com/DontBeProud/go-kits/error_ex"
	"gorm.io/gorm"
)

// ExecuteMethod sql执行方法
type ExecuteMethod func(_ctx context.Context, _db *gorm.DB) error

// BatchExecuteSqlCommands 批量执行sql请求
func BatchExecuteSqlCommands(ctx context.Context, eg *error_ex.ErrorGroupEx, db *gorm.DB, tasks []ExecuteMethod) error {
	// init execute method list
	taskNum := len(tasks)
	taskList := make([]func() error, taskNum)
	for index, task := range tasks {
		_task := task
		taskList[index] = func() error {
			return _task(ctx, db)
		}
	}
	return eg.BatchGo(taskList)
}
