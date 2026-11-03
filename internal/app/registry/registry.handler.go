package registry

import (
	authApp "profile-template/internal/app/services/auth"
	profileApp "profile-template/internal/app/services/profile"
)

type RegisterHandler struct {
	AuthHandler    authApp.Handler
	ProfileHandler profileApp.Handler
}

func NewRegisterHandler(apps RegisterApp) RegisterHandler {
	return RegisterHandler{
		AuthHandler:    authApp.NewHandler(apps.authService),
		ProfileHandler: profileApp.NewHandler(apps.profileService),
	}
}
