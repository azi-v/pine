package pine

import (
	"context"
	"strconv"
	"time"
)

type Fields interface {
	Input() map[string]string
}

type DefaultFields struct {
	Ts        int64
	ServiceID string
	Service   string
	LogID     string
}

func (d DefaultFields) Input() map[string]string {
	return map[string]string{
		"ts":         strconv.FormatInt(d.Ts, 10),
		"service_id": d.ServiceID,
		"service":    d.Service,
		"logid":      d.LogID,
	}
}

func input(ctx context.Context, level LogLevel, msg string, f Fields,
	strs map[string]string, ints map[string]int64,
	floats map[string]float64, bools map[string]bool) {

	if strs == nil {
		strs = make(map[string]string)
	}

	now := time.Now()
	strs["timestamp"] = strconv.FormatUint(uint64(now.Unix()), 10)

	fm := f.Input()
	for k, v := range fm {
		strs[k] = v
	}

	originchan <- &origin{
		Level: level,
		Msg:   msg,
		Str:   strs,
		Float: floats,
		Int:   ints,
		Bool:  bools,
	}

}

func Debug(ctx context.Context, msg string, f Fields,
	strs map[string]string, ints map[string]int64,
	floats map[string]float64, bools map[string]bool) {
	input(ctx, DebugLevel, msg, f, strs, ints, floats, bools)
}

func Info(ctx context.Context, msg string, f Fields,
	strs map[string]string, ints map[string]int64,
	floats map[string]float64, bools map[string]bool) {
	input(ctx, InfoLevel, msg, f, strs, ints, floats, bools)

}

func Warn(ctx context.Context, msg string, f Fields,
	strs map[string]string, ints map[string]int64,
	floats map[string]float64, bools map[string]bool) {
	input(ctx, WarnLevel, msg, f, strs, ints, floats, bools)

}

func Error(ctx context.Context, msg string, f Fields,
	strs map[string]string, ints map[string]int64,
	floats map[string]float64, bools map[string]bool) {
	input(ctx, ErrorLevel, msg, f, strs, ints, floats, bools)

}

func Fatal(ctx context.Context, msg string, f Fields,
	strs map[string]string, ints map[string]int64,
	floats map[string]float64, bools map[string]bool) {
	// TODO: context处理，在日志顺利落盘以后，panic
	input(ctx, FatalLevel, msg, f, strs, ints, floats, bools)

}
