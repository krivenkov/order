package auth

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// const prefix = "Bearer "

type JWT struct {
	// tokenizer *tokenizer.Tokenizer
	logger *zap.Logger
}

func NewJWT(logger *zap.Logger) *JWT {
	return &JWT{
		logger: logger,
	}
}

func (j *JWT) Handle(tokenStr string) (interface{}, error) {
	// tokenStr = strings.TrimPrefix(tokenStr, prefix)

	// TODO: add check
	// tok, err := j.tokenizer.Parse(tokenizer.TypeAccess, tokenStr)
	// if err != nil {
	// 	if brerr.As[tokenizer.ValidationError](err) {
	// 		return nil, ErrInvalidGrand{
	// 			Description: "access token: " + err.Error(),
	// 		}
	// 	}
	//
	// 	j.logger.Error("error parse jwt token", zap.String("token", tokenStr), zap.Error(err))
	//
	// 	return nil, errors.New("internal error")
	// }

	return uuid.New().String(), nil
}
