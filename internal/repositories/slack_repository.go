package repositories

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/mattn/go-jsonpointer"
	"github.com/youichiro/go-slack-my-unipos/internal/gateway"
)

type SlackRepository struct {
	Token        string
	ViewsDirPath string
}

func (repo *SlackRepository) OpenModal(triggerID string) ([]byte, error) {
	var requestParams, view interface{}

	viewPath := filepath.Join(repo.ViewsDirPath, "card_form.json")
	_, err := os.Stat(viewPath)
	if err != nil {
		return nil, err
	}

	requestJSON := `{"trigger_id": "", "view": {}}`
	err = json.Unmarshal([]byte(requestJSON), &requestParams)
	if err != nil {
		return nil, err
	}

	viewJSON, err := os.ReadFile(viewPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(viewJSON), &view)
	if err != nil {
		return nil, err
	}

	err = jsonpointer.Set(requestParams, "/view", view)
	if err != nil {
		return nil, err
	}

	err = jsonpointer.Set(requestParams, "/trigger_id", triggerID)
	if err != nil {
		return nil, err
	}

	requestParamsJSON, err := json.Marshal(requestParams)
	if err != nil {
		return nil, err
	}

	resp, err := gateway.SlackPostJSON(repo.Token, "views.open", string(requestParamsJSON))

	return resp, err
}
