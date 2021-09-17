package pine

// writechan 管道满时立即落盘、或者有数据时，立即落盘
// 对不同时间的日志，启用不同的goroutine落盘
// goroutine的复用