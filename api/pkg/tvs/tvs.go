package tvs

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Interactor interface {
	AuthorizeByPassword(req *GrantRequest) (string, error)
}

const (
	passwordGrantEndpoint = "/realms/master/protocol/openid-connect/token"
)

type tvs struct {
	address      string
	clientID     string
	clientSecret string
	client       http.Client
}

func New(address, clientID, clientSecret string) Interactor {
	c := http.Client{}

	tvs := &tvs{
		address:      address,
		clientID:     clientID,
		clientSecret: clientSecret,
		client:       c,
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

	data := url.Values{}
	data.Set("client_id", tvs.clientID)
	data.Set("client_secret", tvs.clientSecret)
	data.Set("grant_type", string(req.GrantType))
	data.Set("username", req.Username)
	data.Set("password", req.Password)
	data.Set("scope", req.Scope)
	encodedData := data.Encode()

	r, err := http.NewRequest("POST", tvs.address+passwordGrantEndpoint, strings.NewReader(encodedData))
	if err != nil {
		return "", fmt.Errorf("new request %w", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := tvs.client.Do(r)
	if err != nil {
		return "", fmt.Errorf("do request %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	authResp := &AuthResponse{}
	if err := json.NewDecoder(resp.Body).Decode(authResp); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	return authResp.AccessToken, nil
}
