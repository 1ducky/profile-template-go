package assets

import (
	"context"
	"io"
	"profile-template/internal/assets"
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
	res, err := a.assets.Upload(ctx, req.Reader, assets.UploadPolicy{Category: assets.CategoryPost, MaxSize: 2 * assets.MB, UserID: req.UserID})
	return AssetsUploadRes{FileKey: res.FileKey, Metadata: res}, err
}
