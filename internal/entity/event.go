package entity

import (
	"encoding/json"
	"io/fs"

	"github.com/Namchee/konfigured/internal/constant"
	"github.com/Namchee/konfigured/internal/utils"
)

type Branch struct {
	// actual branch name
	Ref string `json:"ref"`
}

type PullRequest struct {
	// branch head
	Head Branch `json:"head"`
}

// Event that triggers the action
type Event struct {
	// action name
	Action string `json:"action"`
	// pull request number
	Number int `json:"number"`
	// pull request 'object'
	PullRequest PullRequest `json:"pull_request"`
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
