package user

import (
	"errors"
	"healthy/internal/database/models"
	"healthy/internal/database/repositories"
	"healthy/internal/pkg/defined"
	"healthy/internal/pkg/enc"

	"github.com/rs/zerolog/log"
)

func (u *UserHandler) SignUp(req *SignUpReq) error {
	hashPassword, err := enc.Create(req.Password)
	if err != nil {
		log.Error().Err(err).Msg("could not hash password")
		return err
	}
	err = repositories.AddUser(u.db, &models.User{
		Username: req.Username,
		Gender:   req.Gender,
		Birthday: req.Birthday,
		Password: hashPassword,
	})
	if err != nil {
		if errors.Is(err, defined.ErrUsernameExists) {
			return errors.New("username already exists")
		}
		return err
	}
	return nil
}

func (h *UserHandler) SignIn(req *SignInReq) (*User, error) {
	user, err := repositories.GetUserByUsername(h.db, req.Username)
	if err != nil {
		return nil, err
	}
	matched, err := enc.Check(req.Password, user.Password)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, errors.New("password not matched")
	}

	return &User{
		Username: user.Username,
		Gender:   user.Gender,
	}, nil
}

type SignInReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type SignUpReq struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}
