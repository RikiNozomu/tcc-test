// filepath: c:\Users\RIKI\Desktop\work\tcc-test\api\adapters\db\sqlite\user.go
package sqlite

import (
	"gorm.io/gorm"

	models "tcc-test/api/core/models"
)

type SQLiteUserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(adapter *SQLiteAdapter) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: adapter.Get()}
}

// Create creates a new user
func (r *SQLiteUserRepository) Create(user models.UserCreate) (*models.User, error) {
	userCreated := models.User{Username: user.Username, Password: user.Password}
	err := r.db.Create(&userCreated).Error
	if err != nil {
		return nil, err
	}
	return &userCreated, nil
}

// GetOne retrieves a user by key and value
func (r *SQLiteUserRepository) GetOne(id string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *SQLiteUserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("Username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
