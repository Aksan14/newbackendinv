package repository

import (
	"context"
	"database/sql"
	"errors"
	"godesaapps/model"
	"godesaapps/util"
	"time"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (repository *userRepositoryImpl) CreateUser(ctx context.Context, tx *sql.Tx, user model.User) (model.User, error) {
	query := `INSERT INTO admin(id, email, nikadmin, pass) VALUES(?, ?, ?, ?)`

	_, err := tx.ExecContext(ctx, query, user.Id, user.Email, user.Nikadmin, user.Password)
	util.SentPanicIfError(err)

	return user, nil
}

func (repository *userRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, idUser string) (model.User, error) {
	query := `SELECT id, email, nikadmin, pass FROM admin WHERE id = ?`

	rows, err := tx.QueryContext(ctx, query, idUser)
	util.SentPanicIfError(err)

	defer rows.Close()
	users := model.User{}
	if rows.Next() {
		err := rows.Scan(&users.Id, &users.Email, &users.Nikadmin, &users.Password)
		util.SentPanicIfError(err)
		return users, err
	} else {
		return users, errors.New("id user not found")
	}
}

func (repository *userRepositoryImpl) FindByNik(ctx context.Context, tx *sql.Tx, nikadmin string) (model.User, error) {
	query := `SELECT id, email, nikadmin, pass FROM admin WHERE nikadmin = ?`

	rows, err := tx.QueryContext(ctx, query, nikadmin)
	util.SentPanicIfError(err)

	defer rows.Close()
	users := model.User{}
	if rows.Next() {
		err := rows.Scan(&users.Id, &users.Nikadmin, &users.Email, &users.Password)
		util.SentPanicIfError(err)
		return users, err
	} else {
		return users, errors.New("id user not found")
	}
}

func (repository *userRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := repository.DB.QueryRow("SELECT id, email FROM admin WHERE email = ?", email).
		Scan(&user.Id, &user.Email)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}
	return &user, nil
}

func (repository *userRepositoryImpl) ReadUser(ctx context.Context, tx *sql.Tx) []model.User {
	query := `SELECT id, nikadmin, email FROM admin`

	rows, err := tx.QueryContext(ctx, query)
	util.SentPanicIfError(err)
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Nikadmin, &user.Email)
		util.SentPanicIfError(err)
		users = append(users, user)
	}

	return users
}

func (r *userRepositoryImpl) UpdateResetToken(email, token string, expiry time.Time) error {
    expiryTime := expiry.Format("2006-01-02 15:05:05")
    query := `UPDATE admin SET reset_token = ?, reset_expiry = ? WHERE email = ?`
    _, err := r.DB.Exec(query, token, expiryTime, email)
    return err
}


func (repository *userRepositoryImpl) FindByResetToken(token string) (*model.User, error) {
	var user model.User
	err := repository.DB.QueryRow("SELECT id, email FROM admin WHERE reset_token = ? AND reset_expiry > UNIX_TIMESTAMP()", token).
		Scan(&user.Id, &user.Email)
	if err != nil {
		return nil, errors.New("token tidak valid atau kadaluarsa")
	}
	return &user, nil
}

func (repository *userRepositoryImpl) UpdatePassword(email, newPassword string) error {
	_, err := repository.DB.Exec("UPDATE admin SET pass = ?, reset_token = NULL, reset_expiry = NULL WHERE email = ?", newPassword, email)
	return err
}


type wargaRepositoryImpl struct {
	db *sql.DB
}

func NewWargaRepository(db *sql.DB) WargaRepository {
	return &wargaRepositoryImpl{db: db}
}

func (r *wargaRepositoryImpl) InsertWarga(warga model.Warga) error {
	query := `INSERT INTO warga (nik, nama_lengkap, alamat, jenis_surat, keterangan, file_upload, no_hp, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`
	_, err := r.db.Exec(query, warga.NIK, warga.NamaLengkap, warga.Alamat, warga.JenisSurat, warga.Keterangan, warga.FileUpload, warga.NoHP)
	return err
}
