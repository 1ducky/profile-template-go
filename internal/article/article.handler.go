package article

import (
	"encoding/json"
	"net/http"
	"profile-template/internal/contextapp"
	"profile-template/internal/transport"
	"profile-template/internal/utils"
)

type Handler struct {
	uc ArticleUsecaseContract
}

func NewHandler(uc ArticleUsecaseContract) *Handler {
	return &Handler{uc: uc}
}

// POST /articles
func (h *Handler) create(w http.ResponseWriter, r *http.Request) error {
	auth, ok := contextapp.GetAuth(r.Context())
	if !ok {
		return ErrUnauthenticated
	}
	var body CreateArticleBody
	if err := utils.DecodeJSON(w, r, &body, MAXSIZEREAD); err != nil {
		return ErrBadRequest
	}

	a, err := h.uc.Create(r.Context(), CreateInput{
		AuthorID:  auth.ID,
		Title:     body.Title,
		Excerpt:   body.Excerpt,
		Category:  body.Category,
		Badges:    body.Badges,
		Tags:      body.Tags,
		Cover:     body.Cover,
		Content:   body.Content,
		PublishAt: body.PublishedAt,
	})
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(transport.Response[ArticleEntity]{Status: transport.StatusSuccess, Data: *a})
}

// PATCH /articles/{id}
func (h *Handler) update(w http.ResponseWriter, r *http.Request) error {
	auth, ok := contextapp.GetAuth(r.Context())
	if !ok {
		return ErrUnauthenticated
	}
	var body UpdateArticleBody
	if err := utils.DecodeJSON(w, r, &body, MAXSIZEREAD); err != nil {
		return ErrBadRequest
	}

	err := h.uc.Update(r.Context(), UpdateInput{
		AuthorID:  auth.ID,
		ID:        r.PathValue("id"),
		Slug:      body.Slug,
		Title:     body.Title,
		Excerpt:   body.Excerpt,
		Category:  body.Category,
		Badges:    body.Badges,
		Tags:      body.Tags,
		Cover:     body.Cover,
		Content:   body.Content,
		PublishAt: body.PublishedAt,
	})
	if err != nil {
		return err
	}

	return nil
}

// DELETE /articles/{id}
func (h *Handler) delete(w http.ResponseWriter, r *http.Request) error {
	auth, ok := contextapp.GetAuth(r.Context())
	if !ok {
		return ErrUnauthenticated
	}
	err := h.uc.Delete(r.Context(), auth.ID, r.PathValue("id"))
	if err != nil {
		return err
	}
	return nil
}

// GET /articles/{slug}
func (h *Handler) getBySlug(w http.ResponseWriter, r *http.Request) error {
	auth, ok := contextapp.GetAuth(r.Context())
	if !ok {
		return ErrUnauthenticated
	}
	// ViewerID boleh kosong (pengunjung anonim).
	a, err := h.uc.GetBySlug(r.Context(), r.PathValue("slug"), auth.ID)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(transport.Response[ArticleResponse]{Status: transport.StatusSuccess, Data: toResponse(a)})
}

// GET /articles?q=&category=&author_id=&tag=go&tag=backend&page=1&page_size=10
func (h *Handler) list(w http.ResponseWriter, r *http.Request) error {
	auth, ok := contextapp.GetAuth(r.Context())
	if !ok {
		return ErrUnauthenticated
	}
	q := r.URL.Query()

	page, err := intParam(q.Get("page"), "page")
	if err != nil {
		return err
	}
	pageSize, err := intParam(q.Get("page_size"), "page_size")
	if err != nil {
		return err
	}

	out, err := h.uc.List(r.Context(), ListInput{
		ViewerID: auth.ID,
		AuthorID: q.Get("author_id"),
		Search:   q.Get("q"),
		Category: q.Get("category"),
		Tags:     q["tag"],
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return err
	}

	items := make([]ArticleResponse, 0, len(out.Items))
	for i := range out.Items {
		items = append(items, toResponse(&out.Items[i]))
	}
	// writeJSON(w, http.StatusOK, listResponse{
	// 	Data: items,
	// 	Meta: listMeta{Total: out.Total, Page: out.Page, PageSize: out.PageSize},
	// })
	return json.NewEncoder(w).Encode(transport.Response[ListResponse]{Status: transport.StatusSuccess, Data: ListResponse{Data: items, Meta: ListMeta{}}})
}
