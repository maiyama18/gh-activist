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
	httpClient      *http.Client
	user            string
	authHeaderValue string
	repo            string
	file            string
}

func NewClient(user, token, repo, file string) *Client {
	return &Client{
		httpClient:      http.DefaultClient,
		user:            user,
		authHeaderValue: fmt.Sprintf("token %s", token),
		repo:            repo,
		file:            file,
	}
}

func (c *Client) Commit(message, content string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", c.user, c.repo, c.file)
	reqBody, err := newCommitRequest(message, content)
	if err != nil {
		return err
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Authorization", c.authHeaderValue)
	if err != nil {
		return err
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

type commitRequest struct {
	message string
	content string
}

func newCommitRequest(message, content string) (*commitRequest, error) {
	var errMsgs []string
	if message == "" {
		errMsgs = append(errMsgs, "message is empty")
	}
	if content == "" {
		errMsgs = append(errMsgs, "content is empty")
	}
	if len(errMsgs) != 0 {
		errMsg := strings.Join(errMsgs, ", ")
		return nil, errors.New(errMsg)
	}

	return &commitRequest{
		message: message,
		content: base64.StdEncoding.EncodeToString([]byte(content)),
	}, nil
}
