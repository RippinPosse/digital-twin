package tvs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Interactor interface {
	AuthorizeByPassword(req *GrantRequest) (string, error)
}

const (
	passwordGrantEndpoint = "/realms/master/protocol/openid-connect/token"
)

type tvs struct {
	address string
	client  http.Client
}

func New(address string) Interactor {
	c := http.Client{}

	tvs := &tvs{
		address: address,
		client:  c,
	}

	return tvs
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func (tvs *tvs) AuthorizeByPassword(req *GrantRequest) (string, error) {
	if err := req.validate(); err != nil {
		return "", fmt.Errorf("invalid request: %w", err)
	}

	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	buf := bytes.NewBuffer(body)
	r, err := tvs.client.Post(tvs.address + passwordGrantEndpoint, "application/json", buf)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}

	resp := &AuthResponse{}
	if err := json.NewDecoder(r.Body).Decode(resp); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	return resp.AccessToken, nil
}
