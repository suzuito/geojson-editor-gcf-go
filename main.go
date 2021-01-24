package gcf

import (
	"context"
	"fmt"
	"time"

	"github.com/suzuito/geojson-editor-go/application"
	"github.com/suzuito/geojson-editor-go/entity"
	"github.com/suzuito/geojson-editor-go/usecase"
	"golang.org/x/xerrors"
)

func newUsecase(ctx context.Context) (usecase.Usecase, *application.Application, func(), error) {
	app, err := application.New(ctx)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("Cannot get application : %w", err)
	}
	u, closeFunc, err := application.NewUsecase(ctx, app)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("Cannot get usecase : %w", err)
	}
	return u, app, closeFunc, nil
}

type AuthEvent struct {
	UID         string `json:"uid"`
	DisplayName string `json:"displayName"`
	PhotoURL    string `json:"photoURL"`
}

func AfterUserSignUp(ctx context.Context, e AuthEvent) error {
	now := time.Now().Unix()
	fmt.Printf("User '%s' is created", e.UID)
	u, _, closeFunc, err := newUsecase(ctx)
	if err != nil {
		return err
	}
	defer closeFunc()
	user := entity.NewUser(
		entity.UserID(e.UID),
		e.DisplayName,
		e.PhotoURL,
	)
	if err := u.InitializeUser(ctx, user, now); err != nil {
		return xerrors.Errorf("Cannot init user '%s' : %w", user.UserID, err)
	}
	return nil
}
