// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package accounts

import (
	"context"
	"github.com/fernandodr19/library/pkg/domain/vos"
	"sync"
)

// AccountsMockUsecase is a mock implementation of Usecase.
//
// 	func TestSomethingThatUsesUsecase(t *testing.T) {
//
// 		// make and configure a mocked Usecase
// 		mockedUsecase := &AccountsMockUsecase{
// 			CreateAccountFunc: func(contextMoqParam context.Context, email vos.Email, password vos.Password) (vos.Tokens, error) {
// 				panic("mock out the CreateAccount method")
// 			},
// 		}
//
// 		// use mockedUsecase in code that requires Usecase
// 		// and then make assertions.
//
// 	}
type AccountsMockUsecase struct {
	// CreateAccountFunc mocks the CreateAccount method.
	CreateAccountFunc func(contextMoqParam context.Context, email vos.Email, password vos.Password) (vos.Tokens, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateAccount holds details about calls to the CreateAccount method.
		CreateAccount []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Email is the email argument value.
			Email vos.Email
			// Password is the password argument value.
			Password vos.Password
		}
	}
	lockCreateAccount sync.RWMutex
}

// CreateAccount calls CreateAccountFunc.
func (mock *AccountsMockUsecase) CreateAccount(contextMoqParam context.Context, email vos.Email, password vos.Password) (vos.Tokens, error) {
	callInfo := struct {
		ContextMoqParam context.Context
		Email           vos.Email
		Password        vos.Password
	}{
		ContextMoqParam: contextMoqParam,
		Email:           email,
		Password:        password,
	}
	mock.lockCreateAccount.Lock()
	mock.calls.CreateAccount = append(mock.calls.CreateAccount, callInfo)
	mock.lockCreateAccount.Unlock()
	if mock.CreateAccountFunc == nil {
		var (
			tokensOut vos.Tokens
			errOut    error
		)
		return tokensOut, errOut
	}
	return mock.CreateAccountFunc(contextMoqParam, email, password)
}

// CreateAccountCalls gets all the calls that were made to CreateAccount.
// Check the length with:
//     len(mockedUsecase.CreateAccountCalls())
func (mock *AccountsMockUsecase) CreateAccountCalls() []struct {
	ContextMoqParam context.Context
	Email           vos.Email
	Password        vos.Password
} {
	var calls []struct {
		ContextMoqParam context.Context
		Email           vos.Email
		Password        vos.Password
	}
	mock.lockCreateAccount.RLock()
	calls = mock.calls.CreateAccount
	mock.lockCreateAccount.RUnlock()
	return calls
}
