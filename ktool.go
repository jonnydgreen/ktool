package main

import (
	"math/rand"
	"time"

	"github.com/projectjudge/ktool/pkg/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cmd.Execute()
}
