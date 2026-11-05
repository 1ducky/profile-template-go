package article

import (
	"context"
	"profile-template/internal/app/uow"
	"profile-template/internal/article"
	"profile-template/internal/assets"

	"github.com/google/uuid"
)

type TXUsecase struct {
	Article article.ArticleUsecaseContract
	Asset   assets.UsecaseContract
}

type ArticleApp struct {
	uow  uow.UnitOfWork[TXUsecase]
	artc article.ArticleUsecaseContract
}

func NewArticleApp(artc article.ArticleUsecaseContract, uow uow.UnitOfWork[TXUsecase]) *ArticleApp {
	return &ArticleApp{artc: artc, uow: uow}
}

func (a *ArticleApp) CreateArticle(ctx context.Context, req article.CreateInput) error {
	if req.AuthorID == "" {
		return article.ErrUnauthenticated
	}
	filekeys := article.ExtractFileKeys(req.Content)
	filekeys = append(filekeys, req.Cover)
	err := a.uow.Do(ctx, func(ctx context.Context, r TXUsecase) error {
		id := uuid.NewString()

	})

}
