package article

import (
	"encoding/json"
	"net/http"
	"profile-template/internal/article"
	"profile-template/internal/contextapp"
	"profile-template/internal/transport"
	"profile-template/internal/utils"
)

type Handler struct {
	app ArticleApp
}

func NewHandler(app ArticleApp) *Handler {
	return &Handler{app: app}
}

// POST /articles
func (h *Handler) create(w http.ResponseWriter, r *http.Request) error {
	auth, ok := contextapp.GetAuth(r.Context())
	if !ok {
		return article.ErrUnauthenticated
	}
	var body article.CreateArticleBody
	if err := utils.DecodeJSON(w, r, &body, article.MAXSIZEREAD); err != nil {
		return article.ErrBadRequest
	}

	a, err := h.app.CreateArticle(r.Context(), article.CreateInput{
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

	return json.NewEncoder(w).Encode(transport.Response[article.ArticleEntity]{Status: transport.StatusSuccess, Data: a})
}

// PATCH /articles/{id}
func (h *Handler) update(w http.ResponseWriter, r *http.Request) error {
	auth, ok := contextapp.GetAuth(r.Context())
	if !ok {
		return article.ErrUnauthenticated
	}
	var body article.UpdateArticleBody
	if err := utils.DecodeJSON(w, r, &body, article.MAXSIZEREAD); err != nil {
		return article.ErrBadRequest
	}

	err := h.app.UpdateArticle(r.Context(), article.UpdateInput{
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
		return article.ErrUnauthenticated
	}
	err := h.app.DeleteArticle(r.Context(), article.DeleteRequest{
		AuthorID: &auth.ID,
		ID:       r.PathValue("id"),
	})
	if err != nil {
		return err
	}
	return nil
}

// GET /articles/{slug}
func (h *Handler) getBySlug(w http.ResponseWriter, r *http.Request) error {
	auth, ok := contextapp.GetAuth(r.Context())
	if !ok {
		return article.ErrUnauthenticated
	}
	// ViewerID boleh kosong (pengunjung anonim).
	a, err := h.app.GetBySlugArticle(r.Context(), auth.ID, r.PathValue("slug"))
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(transport.Response[article.ArticleResponse]{Status: transport.StatusSuccess, Data: a})
}

// GET /articles?q=&category=&author_id=&tag=go&tag=backend&page=1&page_size=10
func (h *Handler) list(w http.ResponseWriter, r *http.Request) error {
	auth, ok := contextapp.GetAuth(r.Context())
	if !ok {
		return article.ErrUnauthenticated
	}
	q := r.URL.Query()

	page, err := article.IntParam(q.Get("page"), "page")
	if err != nil {
		return err
	}
	pageSize, err := article.IntParam(q.Get("page_size"), "page_size")
	if err != nil {
		return err
	}

	out, err := h.app.GetListArticle(r.Context(), article.ListInput{
		ViewerID: auth.ID,
		AuthorID: q.Get("author_id"),
		Search:   q.Get("q"),
		Category: q.Get("category"),
		Tags:     q["tag"],
		Page:     page,
		PageSize: pageSize,
	})

	// writeJSON(w, http.StatusOK, listResponse{
	// 	Data: items,
	// 	Meta: listMeta{Total: out.Total, Page: out.Page, PageSize: out.PageSize},
	// })
	return json.NewEncoder(w).Encode(transport.Response[article.ListResponse]{Status: transport.StatusSuccess, Data: article.ListResponse{Articles: out.Articles, Meta: out.Meta}})
}
