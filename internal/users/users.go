package users

import (
	"encoding/json"
	"github.com/babenkoivan/todo/internal/db"
	"github.com/babenkoivan/todo/internal/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *User) Create() {
	now := time.Now().UTC()
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	res, _ := db.Connection.Exec("INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)", u.Username, string(hash), now)
	id, _ := res.LastInsertId()

	u.ID = id
	u.Password = string(hash)
	u.CreatedAt = now
}

func (u *User) Login() (string, error) {
	res := db.Connection.QueryRow("SELECT id, password FROM users WHERE username = ?", u.Username)

	var id int64
	var hash string

	if err := res.Scan(&id, &hash); err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.Password)); err != nil {
		return "", err
	}

	u.ID = id
	u.Password = hash

	return jwt.GenerateToken(u.Username), nil
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        int64
		Username  string
		CreatedAt time.Time
	}{
		ID:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	})
}

func Authenticate(tokenString string) (*User, error) {
	username, err := jwt.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	var user User
	res := db.Connection.QueryRow("SELECT * FROM users WHERE username = ?", username)
	err = res.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	return &user, err
}
