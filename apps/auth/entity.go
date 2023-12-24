package auth

import (
	"github.com/RianIhsan/go-ecommerce/infra/response"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type Role string

const (
	ROLE_ADMIN Role = "admin"
	ROLE_USER  Role = "user"
)

type AuthEntity struct {
	Id        int       `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      Role      `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewFromRegisterRequestPayload(payload RegisterRequestPayload) (authEntity AuthEntity) {
	authEntity.Email = payload.Email
	authEntity.Password = payload.Password
	authEntity.Role = ROLE_USER
	authEntity.CreatedAt = time.Now()
	authEntity.UpdatedAt = time.Now()
	return
}

func (a AuthEntity) Validate() (err error) {
	if err = a.ValidateEmail(); err != nil {
		return
	}
	if err = a.ValidatePassword(); err != nil {
		return
	}
	return
}

func (a AuthEntity) ValidateEmail() (err error) {
	if a.Email == "" {
		return response.ErrEmailRequired
	}
	emails := strings.Split(a.Email, "@")
	if len(emails) != 2 {
		return response.ErrEmailInvalid
	}
	return
}

func (a AuthEntity) ValidatePassword() (err error) {
	if a.Password == "" {
		return response.ErrPasswordRequired
	}
	if len(a.Password) < 6 {
		return response.ErrPasswordInvalid
	}
	return
}

func (a AuthEntity) IsExists() bool {
	return a.Id != 0
}

func (a *AuthEntity) EncryptPassword(salt int) (err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	a.Password = string(encryptedPassword)
	return nil
}
