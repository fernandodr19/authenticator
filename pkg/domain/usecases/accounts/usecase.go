package accounts

import (
	"context"
	"time"

	"github.com/fernandodr19/library/pkg/domain/entities/accounts"
	"github.com/fernandodr19/library/pkg/domain/vos"
)

//go:generate moq -skip-ensure -stub -out mocks.gen.go . Usecase:AccountsMockUsecase

var _ Usecase = &AccountsUsecase{}

// Usecase of accounts
type Usecase interface {
	CreateAccount(context.Context, vos.Email, vos.Password) error
	Login(context.Context, vos.Email, vos.Password) (vos.Tokens, error)
}

// Repository of accounts
type Repository interface {
	GetAccountByEmail(context.Context, vos.Email) (accounts.Account, error)
	CreateAccount(context.Context, vos.Email, vos.HashedPassword) (vos.UserID, error)
}

// AccountsUsecase represents account's usecase
type AccountsUsecase struct {
	AccountsRepository Repository
	TokenGenerator     TokenGenerator
	Encrypter          Encrypter
}

// NewAccountsUsecase builds an account usecase
func NewAccountsUsecase(accRepo Repository, tokenGenerator TokenGenerator, encrypter Encrypter) *AccountsUsecase {
	return &AccountsUsecase{
		AccountsRepository: accRepo,
		TokenGenerator:     tokenGenerator,
		Encrypter:          encrypter,
	}
}

// TokenGenerator generates access & refresh tokens
type TokenGenerator interface {
	CreateTokens(acc accounts.Account, accessDuration time.Duration, refreshDuration time.Duration) (vos.Tokens, error)
}

// Encrypter encrypt passwords & validate them
type Encrypter interface {
	HashedPassword(password vos.Password) (vos.HashedPassword, error)
	PasswordMathces(password vos.Password, hashedPassword vos.HashedPassword) bool
}
