package pine

import (
	"time"
)

func Debug(msg string, fields ...Field) {
	fields = append(fields, &Str{Key: "msg", Val: msg})
	originChan <- &OriginData{
		Level: DebugLevel,
		Ts:    time.Now(),
		Fields: fields,
	}
}
