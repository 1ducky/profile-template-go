package registry

import (
	"net/http"
	"profile-template/internal/app/err"
	authApp "profile-template/internal/app/services/auth"
	profileApp "profile-template/internal/app/services/profile"
	"profile-template/internal/transport"
	"profile-template/internal/transport/middleware"
	"strings"
)

func NewRegisterRoutes(handlers RegisterHandler, uc RegisterUC, adapter *transport.Adapter) http.Handler {
	mux := http.NewServeMux()
	routeAuth(mux, handlers.AuthHandler, adapter)
	routeProfile(mux, handlers.ProfileHandler, adapter)

	return middleware.Chain(
		mux,
		middleware.CORSMiddleware,
		middleware.TrackerMiddleware(),
		middleware.NewAuthMiddleware(uc.auth),
	)
}

func routeAuth(mux *http.ServeMux, handler authApp.Handler, adapter *transport.Adapter) {
	mux.HandleFunc("POST /auth/register", adapter.Adapt(handler.Register, err.Auth))
	mux.HandleFunc("POST /auth/login", adapter.Adapt(handler.Login, err.Auth))
	mux.HandleFunc("GET /auth/session", adapter.Adapt(handler.GetSession, err.Auth))
}

func routeProfile(mux *http.ServeMux, handler profileApp.Handler, adapter *transport.Adapter) {
	mux.HandleFunc("GET /profile", adapter.Adapt(handler.GetOwnProfile, err.Profile))
	mux.HandleFunc("PATCH /profile", adapter.Adapt(handler.UpdateProfile, err.Profile))
	mux.HandleFunc("POST /profile/avatar", adapter.Adapt(handler.UploadAvatar, err.Profile))
	mux.HandleFunc(strings.ReplaceAll("GET /profile/{*}", "*", profileApp.DynamicIDPath), adapter.Adapt(handler.GetProfileByID, err.Profile))
}
