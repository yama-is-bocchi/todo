package input_parser

import (
	"errors"
	"fmt"
	"time"

	"github.com/thamaji/date"
	"github.com/yama-is-bocchi/todo/database"
)

func ParseCreatedData(inputs []string) (database.CreatedData, error) {
	if len(inputs) != 3 {
		return database.CreatedData{}, errors.New("invalid todo data received")
	}
	time, err := time.Parse("2006-01-02", inputs[2])
	if err != nil {
		return database.CreatedData{}, fmt.Errorf("failed to perse time:%w", err)
	}
	return database.CreatedData{Title: inputs[0], Desc: inputs[1], Date: date.FromTime(time)}, nil
}
