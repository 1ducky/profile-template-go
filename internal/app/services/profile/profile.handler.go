package profile

import (
	"encoding/json"
	"log"
	"net/http"
	"profile-template/internal/profile"
	"profile-template/internal/transport"
)

type Handler struct {
	app ProfileApp
}

func NewHandler(app ProfileApp) Handler {
	return Handler{app: app}
}

func (h *Handler) GetOwnProfile(w http.ResponseWriter, r *http.Request) error {
	log.Print("GetOwnProfile")
	res, err := h.app.GetOwnProfile(r.Context())
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(transport.Response[ProfileResource]{Status: transport.StatusSuccess, Data: res})
}

func (h *Handler) GetProfileByID(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue(DynamicIDPath)
	if id == "" {
		return profile.ErrNotFound
	}
	res, err := h.app.GetProfileByID(r.Context(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(transport.Response[ProfileResource]{Status: transport.StatusSuccess, Data: res})
}

func (h Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) error {
	var payload UpdateProfileRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	err = h.app.UpdateProfile(r.Context(), payload)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(transport.Response[map[string]string]{Status: transport.StatusSuccess, Data: map[string]string{"message": "Profile updated successfully"}})
}
