package usecases

import (
	"context"
	"os"
	"time"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/usecases"
)

type authenticationUsecase struct {
	repository   interfaces.IUserRepository
	hasher       interfaces.IHasher
	tokenManager interfaces.ITokenManager
}

func (pst authenticationUsecase) Perform(ctx context.Context, txn interface{}, dto dtos.AuthenticationDto) (dtos.AuthenticatedUserDto, error) {
	user, err := pst.repository.FindByEmail(ctx, txn, dto.Email)
	if err != nil {
		return dtos.AuthenticatedUserDto{}, err
	}

	if !pst.hasher.Verify(dto.Password, user.Password) {
		return dtos.AuthenticatedUserDto{}, errors.NewBadRequestError("Wrong password")
	}

	expireIn := time.Now().Add(time.Hour)
	accessToken, err := pst.tokenManager.GenerateToken(dtos.TokenDataDto{
		Id:       user.Id,
		Audience: os.Getenv("WebApi"),
		ExpireIn: expireIn,
	})
	if err != nil {
		return dtos.AuthenticatedUserDto{}, err
	}

	return dtos.AuthenticatedUserDto{
		AccessToken: accessToken,
		Kind:        os.Getenv("TOKEN_KIND"),
		ExpireIn:    expireIn,
	}, nil

}

func NewAuthenticationUseCase(repository interfaces.IUserRepository, hasher interfaces.IHasher, tokenManager interfaces.ITokenManager) usecases.IAuthenticationUseCase {
	return authenticationUsecase{
		repository:   repository,
		hasher:       hasher,
		tokenManager: tokenManager,
	}
}
