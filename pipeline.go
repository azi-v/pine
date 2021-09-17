package pine

import "time"

// 原始数据
type origin struct {
	Level     LogLevel
	Msg       string
	Timestamp *time.Time
	Float     map[string]float64
	Int       map[string]int64
	Str       map[string]string
	Bool      map[string]bool
}

// 序列化数据
type Format struct {
	Msg    string            //输出的数据
	Fields map[string]string //日志字段
}

// 落盘数据
type CronData struct {
	Format
	Path string
}

const (
	// 默认数据管道长度
	defaultOriginLen = 1000
	defaultFormatLen = 500
	defaultCornLen   = 50

	OrignLen  = defaultOriginLen
	FormatLen = defaultFormatLen
	CornLen   = defaultCornLen
)

const (
	// 默认处理数据管道goroutine个数
	defaultOriginMax = 1
	defaultFormatMax = 2
	defaultCornMax   = 20

	OriginMax = defaultOriginMax
	FormatMax = defaultFormatMax
	CornMax   = defaultCornMax
)

var (
	// origin chan
	originchan = make(chan *origin, OrignLen)
)
