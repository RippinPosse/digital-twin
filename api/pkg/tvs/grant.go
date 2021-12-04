package tvs

import "errors"

type GrantRequest struct {
	GrantType GrantType `json:"grant_type"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	Scope     string    `json:"scope"`

	// Required only for authorization_code grant type
	RedirectURL string `json:"redirect_url"`
}

type GrantType string

const (
	PasswordGrant = "password"
	CodeGrant     = "authorization_code"
)

func (r *GrantRequest) validate() error {
	var (
		errUnknownGrantType = errors.New("unknown grant type")
		errPasswordRequired = errors.New("password required")
		errRedirectRequired = errors.New("redirect URL required")
	)
	switch r.GrantType {
	case PasswordGrant:
		if r.Password == "" {
			return errPasswordRequired
		}
	case CodeGrant:
		if r.RedirectURL == "" {
			return errRedirectRequired
		}
	default:
		return errUnknownGrantType
	}

	return nil
}
