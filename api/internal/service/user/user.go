package user

import (
	"fmt"

	"api/pkg/tvs"
)

type user struct {
	svc Service
}

func New(svc Service) User {
	return &user{
		svc: svc,
	}
}

func (u *user) Authorize(username, password string) (string, error) {
	token, err := u.svc.TVS().AuthorizeByPassword(&tvs.GrantRequest{
		GrantType: tvs.PasswordGrant,
		Username:  username,
		Password:  password,
		Scope:     "openid",
	})
	if err != nil {
		return "", fmt.Errorf("authorize by password: %w", err)
	}

	return token, nil
}
