package gh

import (
	"encoding/base64"
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
	body := fmt.Sprintf(`{"message":%s,"content":%s}`, message, base64.StdEncoding.EncodeToString([]byte(content)))

	req, err := http.NewRequest("PUT", url, strings.NewReader(body))
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
