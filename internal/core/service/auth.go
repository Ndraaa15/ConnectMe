package service

import (
	"context"
	"errors"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/bcrypt"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/util"
	"github.com/Ndraaa15/ConnectMe/internal/core/domain"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthService struct {
	repository port.AuthRepositoryItf
	cache      port.CacheItf
	token      port.TokenItf
	email      port.EmailItf
}

func NewAuthService(repository port.AuthRepositoryItf, cache port.CacheItf, token port.TokenItf, email port.EmailItf) *AuthService {
	return &AuthService{
		repository: repository,
		cache:      cache,
		token:      token,
		email:      email,
	}
}

func (auth *AuthService) Register(ctx context.Context, req dto.SignUpRequest) (uuid.UUID, error) {
	client := auth.repository.NewAuthRepositoryClient(false)

	hashedPassword, err := bcrypt.EncryptPassword(req.Password)
	if err != nil {
		return uuid.Nil, err
	}

	user := domain.User{
		ID:       uuid.New(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hashedPassword,
		Phone:    req.Phone,
	}

	err = client.CreateUser(ctx, &user)
	if err != nil {
		return uuid.Nil, err
	}

	code := util.GenerateCode(4)
	auth.cache.Set(ctx, user.ID.String(), code, 5*time.Minute)

	auth.email.SetSubject("Verification Code")
	auth.email.SetReciever(user.Email)
	auth.email.SetSender("fuwafu212@gmail.com")
	auth.email.SetBodyHTML("internal/adapter/pkg/template/verification_code.html", struct{ Code string }{Code: code})

	err = auth.email.Send()
	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

func (auth *AuthService) Verify(ctx context.Context, req dto.VerifyAccountRequest) (uuid.UUID, error) {
	client := auth.repository.NewAuthRepositoryClient(false)

	user, err := client.GetUserByID(ctx, req.ID)
	if err != nil {
		return uuid.Nil, err
	}

	data, err := auth.cache.Get(ctx, user.ID.String())
	if err != nil {
		return uuid.Nil, err
	}

	if data != req.Code {
		return uuid.Nil, errx.New(fiber.StatusBadRequest, "invalid code", errors.New("invalid code"))
	} else {
		err = auth.cache.Delete(ctx, user.ID.String())
		if err != nil {
			return uuid.Nil, err
		}

		user.IsActive = true
	}

	err = client.UpdateUser(ctx, user)
	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}
