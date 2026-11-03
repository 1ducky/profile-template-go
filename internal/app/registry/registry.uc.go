package registry

import (
	"profile-template/internal/auth"
	"profile-template/internal/profile"
)

type RegisterUC struct {
	auth    auth.UsecaseContract
	profile profile.UsecaseContract
}

func NewRegisterUC(repos RegisterRepo) RegisterUC {
	return buildRegisterUC(repos)
}

func buildRegisterUC(repos RegisterRepo) RegisterUC {
	return RegisterUC{
		auth:    auth.NewAuthService(repos.Cred, repos.Token, repos.Session),
		profile: profile.NewUsecase(repos.Profile),
	}
}
