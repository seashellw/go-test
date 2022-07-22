package lib

import (
	"strings"

	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/timex"
)

func getLogStr(params ...string) string {
	var stringBuilder strings.Builder
	stringBuilder.WriteString("[")
	stringBuilder.WriteString(timex.Now().DateFormat("H:I:S"))
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

func LogRed(params ...string) {
	cliutil.Redln(getLogStr(params...))
}
