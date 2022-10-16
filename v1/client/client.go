package client

type Config struct {
	host  string
	token string
}

type Client struct {
	cfg *Config
}

func New(cfg *Config) *Client {
	return &Client{
		cfg: cfg,
	}
}

func (c *Client) CommentOnPR(repo, comment string, issue int) error {
	return nil
}
