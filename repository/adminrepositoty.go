package repository

import (
	"database/sql"
	"godesaapps/model"
)

type AdminRepository interface {
	InsertAdmin(admin model.Admin) error
	FindPegawaiById(id int) (model.Admin, error)
	RoleExists(roleId string) (bool, error)
}

type adminRepositoryImpl struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepositoryImpl{db}
}

func (r *adminRepositoryImpl) FindPegawaiById(id int) (model.Admin, error) {
	var admin model.Admin
	query := "SELECT id, email, namalengkap, nip FROM pegawai WHERE id = ?"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&admin.Id, &admin.Email, &admin.NikAdmin, &admin.NamaLengkap)
	return admin, err
}

func (r *adminRepositoryImpl) InsertAdmin(admin model.Admin) error {
	_, err := r.db.Exec(`
		INSERT INTO admin (id, email, nikadmin, namalengkap, role_id, pass)
		VALUES (?, ?, ?, ?, ?, ?)`,
		admin.Id, admin.Email, admin.NikAdmin, admin.NamaLengkap, admin.RoleId, admin.Pass)
	return err
}

func (r *adminRepositoryImpl) RoleExists(roleId string) (bool, error) {
	var count int
	query := "SELECT COUNT(1) FROM role_admin WHERE id = ?"
	err := r.db.QueryRow(query, roleId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

