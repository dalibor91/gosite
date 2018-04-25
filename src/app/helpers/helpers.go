package helpers

import (
	"time"
	"fmt"
	"strings"
)

func GetTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetDate() string {
	return time.Now().Format("2006-01-02")
}

func Log(msgs ... string) {
	lngth := len(msgs)
	if lngth > 1 {
		lastIdx := len(msgs)-1
		last := strings.ToLower(msgs[lastIdx])
		excludeLast := (last == "info" || last == "debug" || last == "warning" || last == "error" )

		for i := range(msgs) {
			if excludeLast && (i != lastIdx) {
				fmt.Println(fmt.Sprintf("[ %s ] - %s - %s", GetTimestamp(), strings.ToUpper(last), msgs[i]))
			}
		}
	} else { Log(msgs[0], "info") }
}


