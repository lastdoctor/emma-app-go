package pg

import (
	"github.com/google/uuid"
	"time"
)

// UserRepo is store for users
//go:generate mockery --dir . --name UserRepo --output ./mocks
type UserRepo interface {
	GetUser(context.Context, uuid.UUID) (*pg.DBUser, error)
	CreateUser(context.Context, *pg.DBUser) (*pg.DBUser, error)
	UpdateUser(context.Context, *pg.DBUser) (*pg.DBUser, error)
	DeleteUser(context.Context, uuid.UUID) error
}

// User is a JSON user
type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname" validate:"required"`
	Lastname  string    `json:"lastname" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// DBUser is a Postgres user
type DBUser struct {
	tableName struct{}  `pg:"users" gorm:"primaryKey"`
	ID        uuid.UUID `pg:"id,notnull,pk"`
	Firstname string    `pg:"firstname,notnull"`
	Lastname  string    `pg:"lastname,notnull"`
	CreatedAt time.Time `pg:"created_at,notnull"`
}

// ToDB converts User to DBUser
func (user *User) ToDB() *DBUser {
	return &DBUser{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		CreatedAt: user.CreatedAt,
	}
}

// TableName overrides default table name for gorm
func (DBUser) TableName() string {
	return "users"
}

// ToWeb converts DBUser to User
func (dbUser *DBUser) ToWeb() *User {
	return &User{
		ID:        dbUser.ID,
		Firstname: dbUser.Firstname,
		Lastname:  dbUser.Lastname,
		CreatedAt: dbUser.CreatedAt,
	}
}
