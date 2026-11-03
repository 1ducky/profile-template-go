package assets

import (
	"fmt"
	"io"
	"net/http"
)

type Handler struct {
	AssetApp AssetApp
}

func NewHandler(assetsService AssetApp) Handler {
	return Handler{
		AssetApp: assetsService,
	}
}

func (h *Handler) ServeAssets(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(r.PathValue(KEY))
	reader, err := h.AssetApp.ServePublicAssets(r.Context(), r.PathValue(KEY))
	if err != nil {
		return err
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return err
	}

	w.Header().Set("Content-Type", http.DetectContentType(data))
	w.Write(data)

	return nil
}
