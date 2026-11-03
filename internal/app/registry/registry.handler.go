package registry

import (
	assetsApp "profile-template/internal/app/services/assets"
	authApp "profile-template/internal/app/services/auth"
	profileApp "profile-template/internal/app/services/profile"
)

type RegisterHandler struct {
	AuthHandler    authApp.Handler
	ProfileHandler profileApp.Handler
	AssetsHandler  assetsApp.Handler
}

func NewRegisterHandler(apps RegisterApp) RegisterHandler {
	return RegisterHandler{
		AuthHandler:    authApp.NewHandler(apps.authService),
		ProfileHandler: profileApp.NewHandler(apps.profileService),
		AssetsHandler:  assetsApp.NewHandler(apps.assetsService),
	}
}
