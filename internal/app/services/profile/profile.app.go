package profile

import (
	"context"
	"profile-template/internal/app/uow"
	"profile-template/internal/assets"
	"profile-template/internal/contextapp"
	"profile-template/internal/profile"
)

type TXUsecase struct {
	Profile profile.UsecaseContract
	Asset   assets.UsecaseContract
}

type ProfileApp struct {
	uow  uow.UnitOfWork[TXUsecase]
	prof profile.UsecaseContract
}

func NewProfileApp(prof profile.UsecaseContract, uow uow.UnitOfWork[TXUsecase]) *ProfileApp {
	return &ProfileApp{prof: prof, uow: uow}
}

func (p *ProfileApp) GetOwnProfile(ctx context.Context) (ProfileResource, error) {
	identity, ok := contextapp.GetAuth(ctx)
	if !ok {
		return ProfileResource{}, profile.ErrNotFound
	}
	pro, err := p.prof.GetUserByID(ctx, identity.ID)
	if err != nil {
		return ProfileResource{}, err
	}
	return ProfileResource{Username: pro.Username, Tag: pro.Tag, Avatar: pro.AvatarURL, Bio: pro.Bio}, nil
}

func (p *ProfileApp) GetProfileByID(ctx context.Context, id string) (ProfileResource, error) {
	pro, err := p.prof.GetUserByID(ctx, id)
	if err != nil {
		return ProfileResource{}, err
	}
	return ProfileResource{Username: pro.Username, Tag: pro.Tag, Avatar: pro.AvatarURL, Bio: pro.Bio}, nil
}

func (p *ProfileApp) UpdateProfile(ctx context.Context, req UpdateProfileRequest) error {
	user, ok := contextapp.GetAuth(ctx)
	if !ok {
		return profile.ErrNotFound
	}
	err := p.prof.UpdateProfile(ctx, profile.UpdateRequest{ID: user.ID, Username: req.Username, Tag: req.Tag, Bio: req.Bio})
	if err != nil {
		return err
	}
	return nil
}
