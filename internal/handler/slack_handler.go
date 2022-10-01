package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/youichiro/go-slack-my-unipos/internal/repositories"

	jsonpointer "github.com/mattn/go-jsonpointer"
)

type SlackHandler struct{}

func (s SlackHandler) Receive(c *gin.Context) {
	payloadJSON := c.Request.FormValue("payload")

	var payload interface{}
	err := json.Unmarshal([]byte(payloadJSON), &payload)
	if err != nil {
		c.IndentedJSON(500, gin.H{"message": err.Error()})
	}

	requestTypeObj, err := jsonpointer.Get(payload, "/type")
	if err != nil {
		c.IndentedJSON(500, gin.H{"message": err.Error()})
	}
	requestType := requestTypeObj.(string)

	var callbackIDObj interface{}
	switch requestType {
	case "shortcut":
		callbackIDObj, _ = jsonpointer.Get(payload, "/callback_id")
	case "view_submission":
		callbackIDObj, _ = jsonpointer.Get(payload, "/view/callback_id")
	}
	callbackID := callbackIDObj.(string)
	if len(callbackID) == 0 {
		c.IndentedJSON(400, gin.H{"message": "hoge"})
	}

	fmt.Println("callbackID")
	fmt.Println(callbackID)

	switch callbackID {
	case "unipos__post_card":
		err := SlackOpenCardForm(c, payload)
		if err != nil {
			c.IndentedJSON(500, gin.H{"message": err.Error()})
		}
		c.IndentedJSON(200, gin.H{"message": "ok"})
	}
}

func SlackOpenCardForm(c *gin.Context, payload interface{}) error {
	fmt.Println("called SlackOpenCardForm")

	token := os.Getenv("SLACK_TOKEN")
	fmt.Println("SLACK_TOKEN")
	fmt.Println(token)

	slackRepo := repositories.SlackRepository{
		Token:        os.Getenv("SLACK_TOKEN"),
		ViewsDirPath: "../configs/slack",
	}

	triggerID, _ := jsonpointer.Get(payload, "/trigger_id")

	fmt.Println("triggerID")
	fmt.Println(triggerID)

	_, err := slackRepo.OpenModal(triggerID.(string))
	if err != nil {
		c.Error(err)
	}

	return nil
}
