package main

import (
	"os"

	"github.com/haojunyu/fcbeat/cmd"

	_ "github.com/haojunyu/fcbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
