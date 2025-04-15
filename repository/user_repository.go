package repository

import (
	"context"
	"database/sql"
	"godesaapps/model"
	"time"
)

type UserRepository interface {
    CreateUser(ctx context.Context, tx *sql.Tx, user model.User) (model.User, error)
    FindByNik(ctx context.Context, tx *sql.Tx, nikadmin string) (model.User, error)
    FindByEmail(email string) (*model.User, error)
    FindByResetToken(token string) (*model.User, error)
    UpdateResetToken(email, token string, expiry time.Time) error
    UpdatePassword(email, hashedPass string) error
    FindRoleById(ctx context.Context, tx *sql.Tx, roleId string) (model.MstRole, error) 
    GetAllUsers(ctx context.Context) ([]model.User, error)
	DeleteUserByID(ctx context.Context, id string) error
    
}
