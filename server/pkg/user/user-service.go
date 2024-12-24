package user

import (
	"database/sql"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user *types.UserPayLoad) error {
	_, err := s.db.Exec("INSERT INTO users (email, password, role, name, surname, mobile) VALUES (?, ?, ?, ?, ?, ?)", user.Email, user.Password, user.Role, user.Name, user.Surname, user.Mobile)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT id, created_at, deleted_at, email, password, role, name, surname, mobile FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	user := new(types.User)
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.CreatedAt, &user.DeletedAt, &user.Email, &user.Password, &user.Role, &user.Name, &user.Surname, &user.Mobile)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == 0 {
		return nil, nil
	}
	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT id, created_at, deleted_at, email, password, role, name, surname, mobile FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	user := new(types.User)
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.CreatedAt, &user.DeletedAt, &user.Email, &user.Password, &user.Role, &user.Name, &user.Surname, &user.Mobile)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == 0 {
		return nil, nil
	}
	return user, nil
}

func (s *Store) UpdateUser(user *types.User) error {
	_, err := s.db.Exec("UPDATE users SET email = ?, password = ?, role = ?, name = ?, surname = ?, mobile = ? WHERE id = ?", user.Email, user.Password, user.Role, user.Name, user.Surname, user.Mobile, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUsers() ([]types.User, error) {
	rows, err := s.db.Query("SELECT id, created_at, deleted_at, email, password, role, name, surname, mobile FROM users")
	if err != nil {
		return nil, err
	}
	users := make([]types.User, 0)
	for rows.Next() {
		user, err := ScanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return users, nil
}

func (s *Store) DeleteUser(id int) error {
	_, err := s.db.Exec("UPGRADE users SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func ScanRowsIntoUser(row *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := row.Scan(
		&user.ID, 
		&user.CreatedAt, 
		&user.DeletedAt, 
		&user.Email, 
		&user.Password, 
		&user.Role, 
		&user.Name, 
		&user.Surname, 
		&user.Mobile)
	if err != nil {
		return nil, err
	}
	return user, nil
}
