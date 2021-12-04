package user

import "api/pkg/tvs"

type Service interface {
	TVS() tvs.Interactor
}

type User interface {
	Authorize(username, password string) (string, error)
}
