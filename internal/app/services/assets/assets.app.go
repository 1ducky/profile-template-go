package assets

import (
	"context"
	"io"
	"path/filepath"
	"profile-template/internal/assets"

	"github.com/google/uuid"
)

type AssetApp struct {
	assets assets.UsecaseContract
}

func NewAssetsApp(assets assets.UsecaseContract) *AssetApp {
	return &AssetApp{
		assets: assets,
	}
}

func (a *AssetApp) ServePublicAssets(ctx context.Context, key string) (io.ReadCloser, error) {
	return a.assets.Read(ctx, key)
}

func (a *AssetApp) UploadPublicAssets(ctx context.Context, req UploadPublicAssetsReq) (AssetsUploadRes, error) {
	req.UserID = uuid.NewString()
	res, err := a.assets.Upload(ctx, req.Reader, assets.UploadPolicy{Category: assets.CategoryPost, MaxSize: 2 * assets.MB, UserID: req.UserID})
	res.FileKey = filepath.ToSlash(res.FileKey)
	return AssetsUploadRes{FileKey: res.FileKey, Metadata: res}, err
}
