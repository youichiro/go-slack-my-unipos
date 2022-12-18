package repositories

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattn/go-jsonpointer"
	"github.com/youichiro/go-slack-my-unipos/internal/gateway"
	"github.com/youichiro/go-slack-my-unipos/internal/models"
)

type SlackRepository struct {
	Token        string
	ViewsDirPath string
}

type (
    FieldsSectionBlock struct {
        Type   string         `json:"type"`
        Fields []ContentBlock `json:"fields"`
    }
    TextSectionBlock struct {
        Type string       `json:"type"`
        Text ContentBlock `json:"text"`
    }
	ContentBlock struct {
        Type string `json:"type"`
        Text string `json:"text"`
    }
)

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

func (repo *SlackRepository) PostCardResult(userName string, channel string, card models.Card) ([]byte, error) {
	var err error
	var requestParams interface{}

	titleSection := TextSectionBlock{
        Type: "section",
        Text: ContentBlock{
            Type: "mrkdwn",
            Text: "@" + userName + " がuniposを送りました",
        },
    }
    contentList := []ContentBlock{
        {
            Type: "mrkdwn",
            Text: "ポイント: " + fmt.Sprint(card.Point),
        },
        {
            Type: "mrkdwn",
            Text: card.Message,
        },
    }
    contentSection := FieldsSectionBlock{
        Type:   "section",
        Fields: contentList,
    }

	requestJSON := `{"channel": "", "blocks": []}`
    err = json.Unmarshal([]byte(requestJSON), &requestParams)
    if err != nil {
        return nil, err
    }

    jsonpointer.Set(requestParams, "/channel", channel)
    jsonpointer.Set(requestParams, "/blocks", []interface{}{
        titleSection,
        contentSection,
    })

    // 作成した構造体をJSON文字列化する。
    requestParamsJSON, err := json.Marshal(requestParams)
    if err != nil {
        return nil, err
    }

    // 作成したJSON文字列パラメータとしてSlackメッセージ投稿APIに渡す。
    resp, err := gateway.SlackPostJSON(repo.Token, "chat.postMessage", string(requestParamsJSON))

    return resp, err
}
