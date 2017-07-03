package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// Client for working with polls.
type Client struct {
	Endpoint string
}

// CreateInput is the input for creating a poll.
type CreateInput struct {
	Options []string `json:"options"`
}

// CreateOutput is the output for creating a poll.
type CreateOutput struct {
	ID string `json:"id"`
}

// Create a new poll.
func (c *Client) Create(in *CreateInput) (*CreateOutput, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling input")
	}

	res, err := http.Post(c.Endpoint+"/poll", "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}
	defer res.Body.Close()

	// TODO: status and retries

	out := new(CreateOutput)
	if err := json.NewDecoder(res.Body).Decode(out); err != nil {
		return nil, errors.Wrap(err, "unmarshaling output")
	}

	return out, nil
}
