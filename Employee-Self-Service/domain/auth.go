package domain

import (
	"database/sql"
	"employeeSelfService/config"
	"employeeSelfService/errs"
	"employeeSelfService/logger"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 24 hour
const ACCESS_TOKEN_DURATION = 30 * (24 * time.Hour)

type Auth struct {
	Email       string        `db:"email"`
	IdUser      sql.NullInt64 `db:"id_user"`
	IdEmploye   sql.NullInt64 `db:"id_employe"`
	UserRole    string        `db:"user_role"`
	Password    string        `db:"password"`
	NamaLengkap string        `db:"nama_lengkap"`
}

type AccessTokenClaims struct {
	Email       string        `json:"email"`
	IdUser      sql.NullInt64 `json:"id_user"`
	IdEmployee  sql.NullInt64 `json:"id_employe"`
	UserRole    string        `json:"user_role"`
	NamaLengkap string        `json:"nama_lengkap"`
	jwt.RegisteredClaims
}

type AuthToken struct {
	token *jwt.Token
}

func (a Auth) ClaimsAccessToken() AccessTokenClaims {
	if a.IdEmploye.Valid {
		return a.claimsForEmployee()
	} else {
		return a.claimsForAdmin()
	}
}

func (a Auth) claimsForEmployee() AccessTokenClaims {
	return AccessTokenClaims{
		Email:       a.Email,
		IdUser:      sql.NullInt64{Int64: a.IdUser.Int64, Valid: true},
		IdEmployee:  sql.NullInt64{Int64: a.IdEmploye.Int64, Valid: true},
		UserRole:    a.UserRole,
		NamaLengkap: a.NamaLengkap,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ACCESS_TOKEN_DURATION)),
		},
	}
}

func (a Auth) claimsForAdmin() AccessTokenClaims {
	return AccessTokenClaims{
		IdUser:   a.IdUser,
		UserRole: a.UserRole,
	}
}

func NewAuthToken(claims AccessTokenClaims) AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return AuthToken{token: token}
}

func (t AuthToken) NewAccessToken() (string, *errs.AppErr) {
	signedString, err := t.token.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		logger.Error("Failed while signing access token: " + err.Error())
		return "", errs.NewUnexpectedError("cannot generate access token")
	}
	return signedString, nil
}
