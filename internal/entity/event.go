package entity

import (
	"encoding/json"
	"io/fs"

	"github.com/Namchee/setel/internal/constant"
	"github.com/Namchee/setel/internal/utils"
)

// Event that triggers the action
type Event struct {
	// action name
	Action string `json:"action"`
	// pull request number
	Number int `json:"number"`
}

// ReadEvent reads and parse event meta definition
func ReadEvent(fsys fs.FS) (*Event, error) {
	file, err := fsys.Open(
		utils.ReadEnvString("GITHUB_EVENT_PATH")[1:],
	)

	if err != nil {
		return nil, constant.ErrEventFileRead
	}

	var event Event

	if err := json.NewDecoder(file).Decode(&event); err != nil {
		return nil, constant.ErrEventFileParse
	}

	return &event, nil
}
