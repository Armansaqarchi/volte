package service

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
	"volte/backend/utils/test"
)

func NewFakeAuthService(t *testing.T) AuthService {
	return AuthService{
		mongoClient: test.NewFakeMongoClient(),
	}
}

func TestRegister(t *testing.T) {
	ctx, recorder, _ := newTestGinContextWithSession()

	test.CreateNewFakeMongoServer(t)
	req := createFakePostRequest(t, map[string]any{
		"username": "example_username",
		"password": "exmpale_password",
	}, "/auth/signup")
	ctx.Request = req

	authService := NewFakeAuthService(t)
	authService.Register(ctx)

	assert.Equal(t, recorder.Code, http.StatusOK)
}

func TestLogin(t *testing.T) {
	ctx, recorder, _ := newTestGinContextWithSession()

	test.CreateNewFakeMongoServer(t)
	req := createFakePostRequest(t, map[string]any{
		"username": "example_username",
		"password": "exmpale_password",
	}, "/auth/signup")
	ctx.Request = req

	authService := NewFakeAuthService(t)
	authService.Register(ctx)
	ctx, recorder, _ = newTestGinContextWithSession()
	req = createFakePostRequest(t, map[string]any{
		"username": "example_username",
		"password": "exmpale_password",
	}, "/auth/login")
	ctx.Request = req
	authService.Login(ctx)

	assert.Equal(t, recorder.Code, http.StatusOK)
}
