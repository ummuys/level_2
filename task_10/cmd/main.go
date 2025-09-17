package main

import (
	"fmt"
	"os"

	"github.com/ummuys/level_2/task_10/sortlib"
)

func main() {
	if err := sortlib.WBSort(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
