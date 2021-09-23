package pine

import (
	"time"
)

const (
	// 默认数据管道长度
	defaultOriginLen = 1000
	defaultFormatLen = 500
	defaultCronLen   = 50

	OrignLen  = defaultOriginLen
	FormatLen = defaultFormatLen
	CronLen   = defaultCronLen
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

type OriginData struct {
	Level LogLevel
	Ts    time.Time
	Fields []Field
}

type ForamteData struct {
	Ts      time.Time
	Foramte []byte
}

type CronData struct {
	GoalFile string
	Foramte  string
}

var (
	originChan = make(chan *OriginData, OrignLen)
	formatChan = make(chan *ForamteData, FormatLen)
	cronChan = make(chan *CronData, CronLen)
)
