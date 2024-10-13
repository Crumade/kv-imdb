package compute

import (
	"bufio"
	"errors"
	"fmt"
	"kv-imdb/internal/storage"
	"log"
	"os"
	"strings"
	"time"
)

const (
	GetCommand = "GET"
	SetCommand = "SET"
	DelCommand = "DEL"
)

type Database interface {
	Get(string) (string, error)
	Set(string, string)
	Delete(string)
}

func Read() {
	stor := make(map[string]string)
	var db Database
	db = &storage.Database{Data: stor}
	for {
		print("input command:")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		str := strings.Fields(line)
		if len(str) < 2 {
			print("unavailable input")
			continue
		}
		command, err := parseCommand(str[0])
		if err != nil {
			log.Fatal(err)
		}

		switch command {
		case GetCommand:
			result, err := db.Get(str[1])
			if err != nil {
				print(err.Error())
			}
			print(result)
		case SetCommand:
			db.Set(str[1], str[2])
		case DelCommand:
			db.Delete(str[1])
			print("k-v pair has been deleted")
		}

		print(command)
	}
}

func parseCommand(s string) (string, error) {

	switch strings.ToUpper(s) {
	case GetCommand:
		return GetCommand, nil
	case SetCommand:
		return SetCommand, nil
	case DelCommand:
		return DelCommand, nil
	default:
		return "", errors.New("unknown command")
	}
}

func print(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), " ", msg)
}
