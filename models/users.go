package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// ErrNotFound is returned when a resource cannot be found in the database.
	ErrNotFound = errors.New("models: resource not found")
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	return &UserService{db}, nil
}

// Close closes the UserService database connection.
func (us *UserService) Close() error {
	return us.db.Close()
}

// DestructiveReset drops the User table and rebuilds it.
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

// ByID will look up a user with the provided id.
//
// If the user is found, a nil error is returned.
// If the user is not found, an ErrNotFound error will be returned.
// If there was a different error, it will be returned.
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	err := us.db.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Create creates the provided user and will backfill data (e.g. ID, CreatedAt, UpdatedAt, DeletedAt).
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}
