package database

import (
	"bufio"
	"kv-imdb/internal/common"
	"os"

	"go.uber.org/zap"
)

type Database struct {
	Compute Parser
	Stotage StorageEngine
	logger  *zap.Logger
}

type StorageEngine interface {
	Get(string) (string, bool)
	Set(string, string)
	Delete(string)
}

type Parser interface {
	Parse(string) (string, []string, error)
}

func NewDatabase(logger *zap.Logger, p Parser, s StorageEngine) *Database {

	return &Database{Compute: p,
		Stotage: s,
		logger:  logger,
	}
}

func (db *Database) Start() {
	reader := bufio.NewReader(os.Stdin)
	for {
		common.PrintCLI("input command:")

		line, err := reader.ReadString('\n')
		if err != nil {
			db.logger.Error("input error", zap.Error(err))
		}
		command, args, err := db.Compute.Parse(line)

		result := db.query(command, args)
		common.PrintCLI(result)

	}
}

func (db *Database) query(command string, args []string) string {
	switch command {
	case common.GetCommand:
		result, ok := db.Stotage.Get(args[1])
		if !ok {
			return "key not found"
		}
		return result
	case common.SetCommand:
		db.Stotage.Set(args[1], args[2])
		return "done"
	case common.DelCommand:
		db.Stotage.Delete(args[1])
		return "k-v pair has been deleted"
	}
	return "unexcepted command auery"
}
