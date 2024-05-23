package error_ex

import (
	"errors"
	"fmt"
	"github.com/DontBeProud/go-kits/runtime_ex"
)

// ErrorEx 封装error
type ErrorEx interface {
	error
}

// NewErrorEx errors.New
func NewErrorEx(format string, params ...interface{}) ErrorEx {
	return errors.New(fmt.Sprintf(format, params...))
}

// NewErrorExWithFuncNamePrefix 创建将当前所在函数名称作为前缀的错误
// extraCallerSkip: 需要额外跳过的调用者层级。默认填0，若函数内存在层级式调用，则需要填入对应的skip value
func NewErrorExWithFuncNamePrefix(callerSkip int, format string, params ...interface{}) ErrorEx {
	return NewErrorExWithPrefix(fmt.Sprintf("%s: ", runtime_ex.GetCallerFuncName(callerSkip+1)), format, params...)
}

// NewErrorExWithPrefix 创建附带前缀的错误
func NewErrorExWithPrefix(prefix string, format string, params ...interface{}) ErrorEx {
	return errors.New(prefix + fmt.Sprintf(format, params...))
}

// NewErrorExWithSuffix 创建附带错误的后缀
func NewErrorExWithSuffix(suffix string, format string, params ...interface{}) ErrorEx {
	return errors.New(fmt.Sprintf(format, params...) + suffix)
}

// SetFuncNameAsErrorPrefix 将所在函数的名称设置为错误前缀
// callerSkip: 需要额外跳过的调用者层级。默认填0，若函数内存在层级式调用，则需要填入对应的skip value
func SetFuncNameAsErrorPrefix(callerSkip int, err ErrorEx) ErrorEx {
	return SetErrorPrefix(fmt.Sprintf("%s: ", runtime_ex.GetCallerFuncName(callerSkip+1)), err)
}

// SetErrorPrefix 设置error前缀(便于链式调用场景中调整错误格式)
func SetErrorPrefix(prefix string, err ErrorEx) ErrorEx {
	if err == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("%s%s", prefix, err.Error()))
}

// SetErrorSuffix 设置错误后缀(便于链式调用场景中调整错误格式)
func SetErrorSuffix(suffix string, err ErrorEx) ErrorEx {
	if err == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("%s%s", err.Error(), suffix))
}
