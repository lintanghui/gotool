// recover panic and write panic information to file

package gotool

import (
    "bytes"
    "fmt"
    "os"
    "runtime"
    "time"
)

var defFile string = "Crash.log"

func Crash() {
    if err := recover(); err != nil {
        fd, err := os.OpenFile(defFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
        if err != nil {
            fmt.Println("open file :%s fail ,err :%s", defFile, err.Error())
            return
        }
        defer fd.Close()

        fmt.Fprintf(fd, "%s panic : %v", time.Now().String(), err)
        stackInfo := string(stack(2))
        fmt.Fprintf(fd, "%s\n", stackInfo)
    }

}

// skip specified the depth to ingore
func stack(skip int) []byte {
    buf := new(bytes.Buffer)
    for i := skip; ; i++ {
        pc, file, line, ok := runtime.Caller(i)
        if !ok {
            break
        }
        fmt.Fprintf(buf, "\t%s:%d (0x%v)\n ", file, line, pc)
    }
    return buf.Bytes()
}
