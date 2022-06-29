package kit

import (
	"strings"
	"time"

	"github.com/gookit/goutil/cliutil"
)

func getLogStr(params ...string) string {
	var stringBuilder strings.Builder
	stringBuilder.WriteString("[")
	stringBuilder.WriteString(time.Now().Format("15:04:05"))
	stringBuilder.WriteString("]")
	for _, val := range params {
		stringBuilder.WriteString(" ")
		stringBuilder.WriteString(val)
	}
	return stringBuilder.String()
}

func LogBlue(params ...string) {
	cliutil.Blueln(getLogStr(params...))
}
