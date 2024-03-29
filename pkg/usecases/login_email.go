package sparkleuc

import (
	"context"
	"errors"
	"github.com/octofoxio/sparkle"
	entitiesv1 "github.com/octofoxio/sparkle/pkg/entities/v1"
	svcsv1 "github.com/octofoxio/sparkle/pkg/svcs/v1"
)

func (l *LoginUseCase) LoginWithEmail(ctx context.Context, in *svcsv1.LoginInputWithEmail) (*entitiesv1.SessionRecord, error) {
	user, err := l.user.FindOne(ctx, &entitiesv1.User{
		Email: in.Email,
	})

	if err != nil && err == sparkle.ErrNotFound {
		return nil, errors.New("invalid user credential")
	} else if err != nil {
		return nil, err
	}

	if !user.ValidatePassword(in.PlainPassword.GetData()) {
		return nil, errors.New("invalid user credential")
	}

	s, err := l.CreateSession(ctx, user.ID.GetData())
	if err != nil {
		return nil, err
	}

	return s, nil

}
