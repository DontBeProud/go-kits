package runtime_ex

import "runtime"

// GetCallerFuncName 获取调用者的函数名称
func GetCallerFuncName(callerSkip int) string {
	pc, _, _, _ := runtime.Caller(callerSkip + 1)
	return runtime.FuncForPC(pc).Name()
}
