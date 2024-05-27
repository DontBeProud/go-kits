package cron_ex

import (
	"github.com/DontBeProud/go-kits/error_ex"
	"github.com/robfig/cron"
	"time"
)

// StandardCronMessage 标准的计划任务执行信息
type StandardCronMessage struct {
	CronServiceName string
	StartTime       time.Time
	EndTime         time.Time
	Error           error
}

// NewStandardCronMessage 新建标准的计划任务执行信息
func NewStandardCronMessage(serviceName string, startTime time.Time, endTime time.Time, err error) *StandardCronMessage {
	return &StandardCronMessage{
		CronServiceName: serviceName,
		StartTime:       startTime,
		EndTime:         endTime,
		Error:           err,
	}
}

// AddStandardCronTask 添加标准的计划任务
func AddStandardCronTask(cronObj *cron.Cron, serviceName string, cronSpec string, fn func() error,
	cronMsgChan chan *StandardCronMessage) error {

	errPrefix := "AddStandardCronTask: "
	if cronObj == nil {
		return error_ex.NewErrorExWithPrefix(errPrefix, "计划任务对象为空")
	}

	if fn == nil {
		return error_ex.NewErrorExWithPrefix(errPrefix, "执行方法为空")
	}

	if cronMsgChan == nil {
		return error_ex.NewErrorExWithPrefix(errPrefix, "接收计划任务执行消息的管道为空")
	}

	return error_ex.SetErrorPrefix(errPrefix, cronObj.AddFunc(cronSpec, func() {
		start := time.Now()
		cronMsgChan <- NewStandardCronMessage(serviceName, start, time.Now(), fn())
	}))
}
