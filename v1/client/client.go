package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

type Interface interface {
	CommentOnPR(ctx context.Context, repo, comment string, issue int) error
}

type Client struct {
	client *http.Client
	host   string
	token  string
}

func New(client *http.Client, host, token string) *Client {
	c := &Client{
		client: client,
		host:   host,
		token:  token,
	}

	return c
}

func (c *Client) CommentOnPR(ctx context.Context, repo, comment string, issue int) error {
	params := url.Values{}
	params.Add("repo", repo)
	params.Add("comment", comment)
	params.Add("issue", fmt.Sprintf("%d", issue))

	uri := fmt.Sprintf("%s/pull_request/comment?%s", c.host, params.Encode())
	req, err := http.NewRequest(http.MethodPost, uri, http.NoBody)
	if err != nil {
		return err
	}

	req.Header.Add("X-Access-Token", c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	logrus.Infof("commented on PR %d", issue)

	return nil
}
