package cli

import (
	"errors"
	"gh-activist/gh"
	"log"
	"os"
	"time"
)

const (
	Success = iota
	CommitError
)

type Cli struct {
	ghClient *gh.Client
	logger   *log.Logger
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func New() (*Cli, error) {
	user := getEnv("GH_USER", "muiscript")
	token := getEnv("GH_TOKEN", "")
	if token == "" {
		return nil, errors.New("GH_TOKEN variable empty")
	}
	repo := getEnv("GH_REPO", "gh-activist-target")
	file := getEnv("GH_FILE", "commit.txt")
	ghClient := gh.NewClient(user, token, repo, file)

	logFile := getEnv("LOG_FILE", "")
	logWriter, err := os.Open(logFile)
	if err != nil {
		logWriter = os.Stdout
	}
	logger := log.New(logWriter, "[gh-activist]", log.LstdFlags)

	return &Cli{
		ghClient: ghClient,
		logger:   logger,
	}, nil
}

func (c *Cli) Run() int {
	now := time.Now().String()
	if err := c.ghClient.Commit(now, now); err != nil {
		return CommitError
	}

	c.logger.Printf("commited to repository: %s, file: %s\n", c.ghClient.Repo, c.ghClient.File)
	return Success
}
