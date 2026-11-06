package article

import (
	"context"
	"log"
	"path/filepath"
	assetsApp "profile-template/internal/app/services/assets"
	"profile-template/internal/app/uow"
	"profile-template/internal/article"
	"profile-template/internal/assets"
	"profile-template/internal/utils"
	"strings"
	"time"
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

func (a *ArticleApp) CreateArticle(ctx context.Context, req article.CreateInput) (article.ArticleEntity, error) {
	if req.AuthorID == "" {
		return article.ArticleEntity{}, article.ErrUnauthenticated
	}
	filekeys := article.ExtractFileKeys(req.Content)
	filekeys = append(filekeys, req.Cover)

	filekeys = utils.MapField(filekeys, func(key string) (string, bool) {
		if key == "" {
			return "", false
		}
		return a.clearToValidFileKey(key), true
	})
	var articleRes article.ArticleEntity
	err := a.uow.Do(ctx, func(ctx context.Context, r TXUsecase) error {
		cratedAt := time.Now()
		res, err := r.Article.Create(ctx, article.CreateInput{
			AuthorID:  req.AuthorID,
			Title:     req.Title,
			URL:       req.URL,
			Content:   req.Content,
			Cover:     req.Cover,
			Excerpt:   req.Excerpt,
			Category:  req.Category,
			Badges:    req.Badges,
			Tags:      req.Tags,
			PublishAt: &cratedAt,
		})
		if err != nil {
			return err
		}
		articleRes = *res
		err = r.Asset.LinkingAssetByFileKeys(ctx, filekeys, res.ID)
		if err != nil {
			return err
		}
		return nil
	})

	return articleRes, err

}

func (a *ArticleApp) UpdateArticle(ctx context.Context, req article.UpdateInput) error {
	if req.AuthorID == "" {
		return article.ErrUnauthenticated
	}
	newFileKeys := article.ExtractFileKeys(*req.Content)
	newFileKeys = append(newFileKeys, *req.Cover)
	newFileKeys = utils.MapField(newFileKeys, func(key string) (string, bool) {
		if key == "" {
			return "", false
		}
		return a.clearToValidFileKey(key), true
	})
	log.Print(newFileKeys)
	return a.uow.Do(ctx, func(ctx context.Context, r TXUsecase) error {
		lastArticle, err := r.Article.GetByID(ctx, req.ID)
		if err != nil {
			return article.ErrNotFound
		}

		err = r.Article.Update(ctx, req)
		if err != nil {
			return err
		}
		err = r.Asset.ReleaseByParentID(ctx, []string{lastArticle.ID})
		if err != nil {
			return err
		}
		err = r.Asset.LinkingAssetByFileKeys(ctx, newFileKeys, lastArticle.ID)
		if err != nil {
			return err
		}
		return nil
	})

}

func (a *ArticleApp) DeleteArticle(ctx context.Context, req article.DeleteRequest) error {
	if req.AuthorID == nil {
		return article.ErrUnauthenticated
	}
	return a.uow.Do(ctx, func(ctx context.Context, r TXUsecase) error {
		err := r.Article.Delete(ctx, *req.AuthorID, req.ID)
		if err != nil {
			return err
		}
		err = r.Asset.ReleaseByParentID(ctx, []string{req.ID})
		if err != nil {
			return err
		}
		return nil
	})
}

func (a *ArticleApp) GetBySlugArticle(ctx context.Context, userID string, slug string) (article.ArticleResponse, error) {
	log.Println("app:", slug)
	res, err := a.artc.GetBySlug(ctx, slug, userID)
	if err != nil {
		return article.ArticleResponse{}, err
	}
	return article.ToResponse(res), nil
}

func (a *ArticleApp) GetListArticle(ctx context.Context, req article.ListInput) (article.ListResponse, error) {
	res, err := a.artc.List(ctx, req)
	if err == article.ErrNotFound {
		return article.ListResponse{
			Articles: []article.ArticleResponse{},
			Meta: article.ListMeta{
				Total:    0,
				Page:     req.Page,
				PageSize: req.PageSize,
			},
		}, err
	} else if err != nil {
		return article.ListResponse{}, err
	}
	responses := make([]article.ArticleResponse, len(res.Items))
	for i, item := range res.Items {
		responses[i] = article.ToResponse(&item)
	}
	Metadata := article.ListMeta{
		Total:    res.Total,
		Page:     res.Page,
		PageSize: res.PageSize,
	}
	return article.ListResponse{
		Articles: responses,
		Meta:     Metadata,
	}, nil
}

func (a *ArticleApp) clearToValidFileKey(filekey string) string {
	return filepath.FromSlash(strings.TrimPrefix(filekey, assetsApp.PREFIX))
}
