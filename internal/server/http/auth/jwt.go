package auth

import (
	"context"
	"errors"

	"github.com/krivenkov/pkg/auth"
	"go.uber.org/zap"
)

type JWT struct {
	authCli *auth.Client
	logger  *zap.Logger
}

func NewJWT(logger *zap.Logger, authCli *auth.Client) *JWT {
	return &JWT{
		logger:  logger,
		authCli: authCli,
	}
}

func (j *JWT) Handle(tokenStr string) (interface{}, error) {
	session, err := j.authCli.SessionFromToken(tokenStr)
	if err != nil {
		if errors.Is(err, auth.ErrNoUserFound) {
			return nil, ErrInvalidGrand{
				Description: "access token: " + err.Error(),
			}
		}

		j.logger.Error("error parse jwt token", zap.String("token", tokenStr), zap.Error(err))

		return nil, errors.New("internal error")
	}

	user, err := j.authCli.ExtractUserByUsername(context.Background(), session.PreferredUsername)
	if err != nil {
		return nil, ErrInvalidGrand{
			Description: "access token: " + err.Error(),
		}
	}

	return user.ID, nil
}
