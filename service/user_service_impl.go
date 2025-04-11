package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"godesaapps/dto"
	"godesaapps/model"
	"godesaapps/repository"
	"godesaapps/util"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key")

type userServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func (service *userServiceImpl) GetUserInfoByNikAdmin(ctx context.Context, nikadmin string) (dto.UserResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return dto.UserResponse{}, err
	}
	defer tx.Rollback()

	user, err := service.UserRepository.FindByNik(ctx, tx, nikadmin)
	if err != nil {
		return dto.UserResponse{}, err
	}

	role, err := service.UserRepository.FindRoleById(ctx, tx, user.RoleID)
	if err != nil {
		return dto.UserResponse{}, err
	}

	if err := tx.Commit(); err != nil {
		return dto.UserResponse{}, err
	}

	roleResponse := dto.RoleResponse{
		IdRole:   role.IdRole,
		RoleName: role.RoleName,
	}

	return dto.UserResponse{
		Id:          user.Id,
		Email:       user.Email,
		Nikadmin:    user.Nikadmin,
		NamaLengkap: user.NamaLengkap,
		Role:        roleResponse,
	}, nil
}

func (service *userServiceImpl) GetRoleByUserId(ctx context.Context, roleID string) (dto.RoleResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return dto.RoleResponse{}, err
	}
	defer tx.Rollback()

	role, err := service.UserRepository.FindRoleById(ctx, tx, roleID)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	if err := tx.Commit(); err != nil {
		return dto.RoleResponse{}, err
	}

	return dto.RoleResponse{
		IdRole:   role.IdRole,
		RoleName: role.RoleName,
	}, nil
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *sql.DB) UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func verifyPassword(storedHash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)) == nil
}

func (service *userServiceImpl) CreateUser(ctx context.Context, userRequest dto.CreateUserRequest) dto.UserResponse {
	tx, err := service.DB.Begin()
	util.SentPanicIfError(err)
	defer util.CommitOrRollBack(tx)

	hashedPass, err := hashPassword(userRequest.Pass)
	util.SentPanicIfError(err)

	user := model.User{
		Id:          uuid.New().String(),
		Email:       userRequest.Email,
		Nikadmin:    userRequest.Nikadmin,
		Password:    hashedPass,
		NamaLengkap: userRequest.NamaLengkap,
		RoleID:      userRequest.Role_id,
	}

	createUser, errSave := service.UserRepository.CreateUser(ctx, tx, user)
	util.SentPanicIfError(errSave)

	role, err := service.UserRepository.FindRoleById(ctx, tx, user.RoleID)
	util.SentPanicIfError(err)

	return convertToResponseDTO(createUser, role)
}

func convertToResponseDTO(user model.User, role model.MstRole) dto.UserResponse {
	return dto.UserResponse{
		Id:          user.Id,
		Email:       user.Email,
		Nikadmin:    user.Nikadmin,
		NamaLengkap: user.NamaLengkap,
		Role: dto.RoleResponse{
			IdRole:   role.IdRole,
			RoleName: role.RoleName,
		},
	}
}

type Claims struct {
	Nikadmin 	string `json:"nikadmin"`
	Email    	string `json:"email"`
	RoleId   	string `json:"role_id"`
	NamaLengkap string `json:"namalengkap"`
	jwt.StandardClaims
}

func (service *userServiceImpl) GenerateJWT(email, nikadmin,namalengkap, roleadmin string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Email:    email,
		Nikadmin: nikadmin,
		RoleId: roleadmin,
		NamaLengkap: namalengkap,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "go-auth-invdes",
		},
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func (service *userServiceImpl) LoginUser(ctx context.Context, loginRequest dto.LoginUserRequest) (string, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return "", fmt.Errorf("failed to start transaction: %v", err)
	}
	defer util.CommitOrRollBack(tx)

	user, err := service.UserRepository.FindByNik(ctx, tx, loginRequest.Nikadmin)
	if err != nil {
		return "", fmt.Errorf("invalid nikadmin")
	}

	if !verifyPassword(user.Password, loginRequest.Pass) {
		return "", fmt.Errorf("invalid password")
	}

	token, err := service.GenerateJWT(user.Email, user.Nikadmin, user.RoleID, user.NamaLengkap)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

func (service *userServiceImpl) FindByNIK(ctx context.Context, nik string) (*dto.UserResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	user, err := service.UserRepository.FindByNik(ctx, tx, nik)
	if err != nil {
		return nil, err
	}

	role, err := service.UserRepository.FindRoleById(ctx, tx, user.RoleID)
	if err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		Id:          user.Id,
		Nikadmin:    user.Nikadmin,
		Email:       user.Email,
		NamaLengkap: user.NamaLengkap,
		Role: dto.RoleResponse{
			IdRole:   role.IdRole,
			RoleName: role.RoleName,
		},
	}

	tx.Commit()
	return &response, nil
}

func (service *userServiceImpl) ForgotPassword(request dto.ForgotPasswordRequest) error {
	user, err := service.UserRepository.FindByEmail(request.Email)
	if err != nil {
		return errors.New("email tidak ditemukan")
	}

	token := generateToken(32)
	expiry := time.Now().Add(15 * time.Minute)

	err = service.UserRepository.UpdateResetToken(user.Email, token, expiry)
	if err != nil {
		return fmt.Errorf("gagal menyimpan token reset: %w", err)
	}

	resetURL := fmt.Sprintf("http://localhost:3000/authentication/reset-password?token=%s", token)

	emailBody := fmt.Sprintf(`
        <html>
        <body>
            <p>Halo <strong>%s</strong>,</p>
            <p>Klik tombol di bawah ini untuk mengatur ulang password Anda:</p>
            <a href="%s">Reset Password</a>
            <p>Link ini akan kedaluwarsa dalam 15 menit.</p>
        </body>
        </html>
    `, user.Email, resetURL)

	return util.SendEmail(user.Email, "Reset Password", emailBody)
}

func (service *userServiceImpl) ResetPassword(request dto.ResetPasswordRequest) error {
	if request.Token == "" {
		return errors.New("token tidak ditemukan")
	}

	user, err := service.UserRepository.FindByResetToken(request.Token)
	if err != nil {
		return errors.New("token tidak valid atau sudah kadaluarsa")
	}

	hashedPass, err := hashPassword(request.Password)
	if err != nil {
		return errors.New("gagal mengenkripsi password")
	}

	err = service.UserRepository.UpdatePassword(user.Email, hashedPass)
	if err != nil {
		return errors.New("gagal mengubah password")
	}

	return nil
}

func generateToken(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)
}
