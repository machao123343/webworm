package base

import "logging"

//新建简易日志记录器
func NewLogger() logging.Logger {
	return logging.NewSimpleLogger()
}
