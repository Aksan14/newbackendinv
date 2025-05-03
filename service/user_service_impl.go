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

func NewUserServiceImpl(userRepository repository.UserRepository, db *sql.DB) UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}


func (s *userServiceImpl) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.UserRepository.GetAllUsers(ctx)
}

func (service *userServiceImpl) DeleteUser(ctx context.Context, id string) error {
	return service.UserRepository.DeleteUserByID(ctx, id)
}

func (s *userServiceImpl) CreateUser(ctx context.Context, req dto.CreateUserRequest) dto.UserResponse {
	tx, err := s.DB.BeginTx(ctx, nil)
	util.SentPanicIfError(err)
	defer util.CommitOrRollBack(tx)

	hashedPass, err := hashPassword(req.Pass)
	util.SentPanicIfError(err)

	user := model.User{
		Id:          uuid.New().String(),
		Email:       req.Email,
		Nikadmin:    req.Nikadmin,
		Password:    hashedPass,
		NamaLengkap: req.NamaLengkap,
		RoleID:      req.Role_id,
	}

	newUser, err := s.UserRepository.CreateUser(ctx, tx, user)
	util.SentPanicIfError(err)

	role, err := s.UserRepository.FindRoleById(ctx, tx, user.RoleID)
	util.SentPanicIfError(err)

	return convertToResponseDTO(newUser, role)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func verifyPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
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

//authh
type Claims struct {
	Nikadmin    string `json:"nikadmin"`
	Email       string `json:"email"`
	RoleId      string `json:"role_id"`
	NamaLengkap string `json:"namalengkap"`
	jwt.StandardClaims
}

func (s *userServiceImpl) GenerateJWT(email, nikadmin, namalengkap, role string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email:       email,
		Nikadmin:    nikadmin,
		RoleId:      role,
		NamaLengkap: namalengkap,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "go-auth-invdes",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (s *userServiceImpl) LoginUser(ctx context.Context, req dto.LoginUserRequest) (string, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}
	defer util.CommitOrRollBack(tx)

	user, err := s.UserRepository.FindByNik(ctx, tx, req.Nikadmin)
	if err != nil || !verifyPassword(user.Password, req.Pass) {
		return "", errors.New("invalid nikadmin or password")
	}

	return s.GenerateJWT(user.Email, user.Nikadmin, user.NamaLengkap, user.RoleID)
}


func (s *userServiceImpl) GetUserInfoByNikAdmin(ctx context.Context, nikadmin string) (dto.UserResponse, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return dto.UserResponse{}, err
	}
	defer tx.Rollback()

	user, err := s.UserRepository.FindByNik(ctx, tx, nikadmin)
	if err != nil {
		return dto.UserResponse{}, err
	}

	role, err := s.UserRepository.FindRoleById(ctx, tx, user.RoleID)
	if err != nil {
		return dto.UserResponse{}, err
	}

	tx.Commit()
	return convertToResponseDTO(user, role), nil
}

func (s *userServiceImpl) GetRoleByUserId(ctx context.Context, roleID string) (dto.RoleResponse, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return dto.RoleResponse{}, err
	}
	defer tx.Rollback()

	role, err := s.UserRepository.FindRoleById(ctx, tx, roleID)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	tx.Commit()
	return dto.RoleResponse{
		IdRole:   role.IdRole,
		RoleName: role.RoleName,
	}, nil
}

func (s *userServiceImpl) FindByNIK(ctx context.Context, nik string) (*dto.UserResponse, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	user, err := s.UserRepository.FindByNik(ctx, tx, nik)
	if err != nil {
		return nil, err
	}

	role, err := s.UserRepository.FindRoleById(ctx, tx, user.RoleID)
	if err != nil {
		return nil, err
	}

	tx.Commit()
	response := convertToResponseDTO(user, role)
	return &response, nil
}

func (s *userServiceImpl) ForgotPassword(req dto.ForgotPasswordRequest) error {
	user, err := s.UserRepository.FindByEmail(req.Email)
	if err != nil {
		return errors.New("email tidak ditemukan")
	}

	token := generateToken(32)
	expiry := time.Now().Add(15 * time.Minute)

	if err := s.UserRepository.UpdateResetToken(user.Email, token, expiry); err != nil {
		return fmt.Errorf("gagal menyimpan token reset: %w", err)
	}

	resetURL := fmt.Sprintf("http://localhost:5800/authentication/reset-password?token=%s", token)
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

func (s *userServiceImpl) ResetPassword(req dto.ResetPasswordRequest) error {
	if req.Token == "" {
		return errors.New("token tidak ditemukan")
	}

	user, err := s.UserRepository.FindByResetToken(req.Token)
	if err != nil {
		return errors.New("token tidak valid atau sudah kadaluarsa")
	}

	hashedPass, err := hashPassword(req.Password)
	if err != nil {
		return errors.New("gagal mengenkripsi password")
	}

	return s.UserRepository.UpdatePassword(user.Email, hashedPass)
}

func generateToken(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)
}
