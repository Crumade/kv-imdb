package compute

import (
	"errors"
	"kv-imdb/internal/common"
	"strings"

	"go.uber.org/zap"
)

type Compute struct {
	logger *zap.Logger
}

func NewCompute(logger *zap.Logger) *Compute {
	return &Compute{logger: logger}
}

func parseCommand(s string) (string, error) {

	switch strings.ToUpper(s) {
	case common.GetCommand:
		return common.GetCommand, nil
	case common.SetCommand:
		return common.SetCommand, nil
	case common.DelCommand:
		return common.DelCommand, nil
	default:
		return "", errors.New("unknown command")
	}
}

func (c *Compute) Parse(rq string) (string, []string, error) {
	args := strings.Fields(rq)
	if len(args) < 2 {
		return "unavailable input", nil, nil

	}
	command, err := parseCommand(args[0])
	if err != nil {
		c.logger.Error("parsing error", zap.Error(err))
	}

	return command, args, nil
}
