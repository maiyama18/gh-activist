package gh

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	HTTPClient      *http.Client
	User            string
	AuthHeaderValue string
	Repo            string
	File            string
}

func NewClient(user, token, repo, file string) *Client {
	return &Client{
		HTTPClient:      http.DefaultClient,
		User:            user,
		AuthHeaderValue: fmt.Sprintf("token %s", token),
		Repo:            repo,
		File:            file,
	}
}

func (c *Client) Commit(message, content string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", c.User, c.Repo, c.File)
	reqBody, err := newCommitRequest(message, content)
	if err != nil {
		return err
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", reqBodyBytes)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Authorization", c.AuthHeaderValue)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

type commitRequest struct {
	Message string
	Content string
}

func newCommitRequest(message, content string) (*commitRequest, error) {
	var errMsgs []string
	if message == "" {
		errMsgs = append(errMsgs, "Message is empty")
	}
	if content == "" {
		errMsgs = append(errMsgs, "Content is empty")
	}
	if len(errMsgs) != 0 {
		errMsg := strings.Join(errMsgs, ", ")
		return nil, errors.New(errMsg)
	}

	return &commitRequest{
		Message: message,
		Content: base64.StdEncoding.EncodeToString([]byte(content)),
	}, nil
}
