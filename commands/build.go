package main

import (
	"go-test/util"
)

func main() {
	util.SystemCall("./client", "pnpm", "install")
	util.SystemCall("./client", "pnpm", "build")
	util.SystemCall("./", "go", "build")
}
