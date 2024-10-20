package common

import (
	"fmt"
	"time"
)

const (
	GetCommand = "GET"
	SetCommand = "SET"
	DelCommand = "DEL"
)

func PrintCLI(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), " ", msg)
}
