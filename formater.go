package pine

type LogLevel string

const (
	// Level
	DebugLevel = LogLevel("Debug")
	InfoLevel  = LogLevel("Info")
	WarnLevel  = LogLevel("Warn")
	ErrorLevel = LogLevel("Error")
	FatalLevel = LogLevel("Fatal")
)

func Formate(ch <-chan *origin) {
	// data := <-ch
	// for k,v:=range  ints{
	// 	strs[k]=strconv.FormatInt(v, 10)
	// }

	// for k,v := range floats {
	// 	strs[k]=strconv.FormatFloat(v, 'f', -1, 64)
	// }

	// for k,v := range bools {
	// 	strs[k]=strconv.FormatBool(v)
	// }
}
