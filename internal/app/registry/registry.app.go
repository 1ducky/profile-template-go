package registry

import (
	authApp "profile-template/internal/app/services/auth"
	profileApp "profile-template/internal/app/services/profile"
)

type RegisterApp struct {
	authService    authApp.AuthApp
	profileService profileApp.ProfileApp
}

func NewRegisterApp(ucs RegisterUC, uows RegisterUoW) RegisterApp {
	return RegisterApp{
		authService:    *authApp.NewAuthApp(ucs.auth, ucs.profile, uows.RegisterCase),
		profileService: *profileApp.NewProfileApp(ucs.profile, uows.ProfileCase),
	}
}
