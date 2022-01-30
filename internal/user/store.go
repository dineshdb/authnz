package user

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

type User struct {
	ID           int64  `json:"id,omitempty"`
	Email        string `json:"email,omitempty"`
	Name         string `json:"name,omitempty"`
	Location     string `json:"location,omitempty"`
	About        string `json:"about,omitempty"`
	PasswordHash string `json:"-"`
}

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type Repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// FIXME: Using serial number as an id can leak information about the
// user. Don't do this in production.
func (r *Repository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS users(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL DEFAULT '',
		location TEXT DEFAULT '',
		about TEXT DEFAULT '',
        passwd TEXT NOT NULL
    );
    `
	_, err := r.db.Exec(query)
	return err
}

func (r *Repository) Create(user *User, passwd string) (*User, error) {
	res, err := r.db.Exec("INSERT INTO users(email, name, location, about, passwd) values(?,?,?,?,?)", user.Email, user.Name, user.Location, user.About, passwd)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = id

	return user, nil
}

func (r *Repository) All() ([]User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Location, &user.About); err != nil {
			return nil, err
		}
		all = append(all, user)
	}
	return all, nil
}

func (r *Repository) GetByEmail(email string) (*User, error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE email = ?", email)

	var user User
	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Location, &user.About, &user.PasswordHash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetByID(id int64) (*User, error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	var user User
	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Location, &user.About, &user.PasswordHash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) Update(id int64, updated User) (*User, error) {
	if id == 0 {
		return nil, errors.New("invalid updated id")
	}
	res, err := r.db.Exec("UPDATE users SET name = ?, email = ?, location = ? WHERE id = ?", updated.Name, updated.Email, updated.Location, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}

func (r *Repository) Delete(id int64) error {
	res, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
