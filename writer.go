package pine

import (
	"bufio"
	"os"
)

func write() {
	for data := range cronChan {
		fileHandle, err := os.OpenFile(data.GoalFile, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer fileHandle.Close()
		buf := bufio.NewWriterSize(fileHandle, 4096)
		buf.WriteString(data.Foramte + "\n")
		err = buf.Flush()
		if err != nil {
			panic(err)
		}
	}

}
