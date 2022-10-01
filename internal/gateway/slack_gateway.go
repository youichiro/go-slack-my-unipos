package gateway

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SlackPostJSON(token string, command string, paramJSON string) ([]byte, error) {
	fmt.Println("called SlackPostJSON")
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://slack.com/api/"+command, strings.NewReader(paramJSON))
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	fmt.Println("response.StatusCode")
	fmt.Println(response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("body")
	fmt.Println(string(body))
	return body, nil
}
