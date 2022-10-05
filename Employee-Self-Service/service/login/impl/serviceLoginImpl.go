package serviceLoginImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"
	repositoryAuth "employeeSelfService/repository/auth"
	"employeeSelfService/request"
	"employeeSelfService/response"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type ServiceLoginImpl struct {
	repo repositoryAuth.RepositoryAuth
}

func NewLoginService(repo repositoryAuth.RepositoryAuth) ServiceLoginImpl {
	return ServiceLoginImpl{repo}
}

func (s *ServiceLoginImpl) Login(loginRequest *request.LoginRequest) (*response.ResponseLogin, *errs.AppErr) {
	// siapkan struct user dan error
	var err *errs.AppErr
	var auth *domain.Auth

	// ambil data user by username dari request
	// jika tidak ketemu, kembalikan error
	if auth, err = s.repo.FindByEmail(loginRequest.Email); err != nil {
		return nil, err
	}

	// cek apakah password sudah benar
	paswordSaltVerify := config.SECRET_KEY + loginRequest.Password
	if errorVerify := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(paswordSaltVerify)); errorVerify != nil {
		logger.Error(fmt.Sprintf("password from %s not verify", loginRequest.Email))
		return nil, errs.NewAuthenticationError("Your Login Failed! Invalid Credential")
	}

	// create claims token

	claims := auth.ClaimsAccessToken()
	authToken := domain.NewAuthToken(claims)

	if accessToken, appErr := authToken.NewAccessToken(); appErr != nil {
		return nil, appErr
	} else {
		return response.NewLoginSucess(accessToken), nil
	}
}
