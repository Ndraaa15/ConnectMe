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

func (auth *AuthService) Register(ctx context.Context, req dto.SignUpRequest) (string, error) {
	repositoryClient := auth.repository.NewAuthRepositoryClient(false)

	hashedPassword, err := bcrypt.EncryptPassword(req.Password)
	if err != nil {
		return "", err
	}

	role, err := parseAccountRole(req.Role)
	if err != nil {
		return "", err
	}

	user := domain.User{
		ID:       uuid.New().String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     role,
		Phone:    req.Phone,
	}

	if user.Role == domain.RoleWorker {
		worker := domain.Worker{
			ID:   user.ID,
			Name: user.FullName,
		}

		err = repositoryClient.CreateWorker(ctx, &worker)
		if err != nil {
			return "", err
		}
	}

	err = repositoryClient.CreateUser(ctx, &user)
	if err != nil {
		return "", err

	}

	code := util.GenerateCode(4)
	err = auth.cache.Set(ctx, user.ID, code, 10*time.Minute)
	if err != nil {
		return "", err

	}

	auth.email.SetSubject("Verification Code")
	auth.email.SetReciever(user.Email)
	auth.email.SetSender("fuwafu212@gmail.com")
	err = auth.email.SetBodyHTML("verification_code.html", struct{ Code string }{Code: code})
	if err != nil {
		return "", err

	}

	err = auth.email.Send()
	if err != nil {
		return "", err

	}

	return user.ID, nil
}

func (auth *AuthService) Verify(ctx context.Context, req dto.VerifyAccountRequest) (string, error) {
	repositoryClient := auth.repository.NewAuthRepositoryClient(false)

	user, err := repositoryClient.GetUserByID(ctx, req.ID)
	if err != nil {
		return "", err
	}

	data, err := auth.cache.Get(ctx, user.ID)
	if err != nil {
		return "", err
	}

	if data != req.Code {
		return "", errx.New(fiber.StatusBadRequest, "invalid code", errors.New("invalid code not match"))
	} else {
		err = auth.cache.Delete(ctx, user.ID)
		if err != nil {
			return "", err
		}

		user.IsActive = true
	}

	err = repositoryClient.UpdateUser(ctx, user)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (auth *AuthService) Login(ctx context.Context, req dto.SignInRequest) (string, error) {
	repositoryClient := auth.repository.NewAuthRepositoryClient(false)

	user, err := repositoryClient.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", nil
	}

	if !user.IsActive {
		return "", errx.New(fiber.StatusForbidden, "user not verified", errors.New("user not verified, please verify first"))
	}

	if err := bcrypt.ComparePassword(user.Password, req.Password); err != nil {
		return "", errx.New(fiber.StatusBadRequest, "user email or password invalid", errors.New("user email or password invalid, please check again"))
	}

	token, err := auth.token.Encode(dto.NewPayload(user.ID, 72*time.Hour, user.Role.String()))
	if err != nil {
		return "", err
	}

	return token, nil
}

func parseAccountRole(role string) (domain.AccountRole, error) {
	switch role {
	case "user":
		return domain.RoleUser, nil
	case "worker":
		return domain.RoleWorker, nil
	default:
		return domain.RoleUnknown, errx.New(fiber.StatusBadRequest, "input role is unknown", errors.New("invalid role"))
	}
}
