package pine

import "encoding/json"

type LogLevel string

const (
	// Level
	DebugLevel = LogLevel("Debug")
	InfoLevel  = LogLevel("Info")
	WarnLevel  = LogLevel("Warn")
	ErrorLevel = LogLevel("Error")
	FatalLevel = LogLevel("Fatal")
)

func formate() {
	for data := range originChan {

		fmap := make(map[string]string, len(data.Fields))
		for _, v := range data.Fields {
			res := v.Input()
			fmap[res.Key] = res.Val
		}
		bs,_:=json.Marshal(fmap)

		switch data.Level {
		case DebugLevel:
			formatChan <- &ForamteData{
				Ts:      data.Ts,
				Foramte: bs,
			}
		case InfoLevel:
		case ErrorLevel:
		case WarnLevel:
		case FatalLevel:
		default:
			// 打入黑洞

		}
	}

}
