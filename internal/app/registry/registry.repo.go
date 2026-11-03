package registry

import (
	"profile-template/internal/auth"
	"profile-template/internal/auth/credential"
	"profile-template/internal/auth/jwt"
	"profile-template/internal/auth/session"
	"profile-template/internal/config"
	"profile-template/internal/db"
	"profile-template/internal/profile"
)

type RegisterRepo struct {
	Cred    auth.CredentialRepositoryContract
	Token   auth.TokenContract
	Session auth.SessionRepositoryContract
	Profile profile.RepoContract
}

func NewRegisterRepo(persistent db.DBTX, conf config.AppConfig) RegisterRepo {
	return RegisterRepo{
		Cred:    credential.NewMySQLRepository(persistent),
		Token:   jwt.NewJWT(conf.AuthConfig),
		Session: session.NewMySQLRepository(persistent),
		Profile: profile.NewMySQLRepository(persistent),
	}
}
