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
