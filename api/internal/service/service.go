package service

import (
	"api/internal/service/user"
	"api/pkg/tvs"
)

type Service interface {
	User() user.User
}

type SVC struct {
	user user.User

	tvs tvs.Interactor
}

func New(tvs tvs.Interactor) Service {
	svc := &SVC{
		tvs: tvs,
	}

	svc.user = user.New(svc)

	return svc
}

func (svc *SVC) User() user.User {
	return svc.user
}

func (svc *SVC) TVS() tvs.Interactor {
	return svc.tvs
}
