package repositories

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/models"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(c echo.Context, user models.User) (err error)
	GetUserByID(c echo.Context, userID uint) (user *models.User, err error)
	GetUserByUsername(c echo.Context, username string) (user *models.User, err error)
	UpdateUser(c echo.Context, user models.User) (err error)
	DeleteUser(c echo.Context, user models.User) (err error)
}

type UserRepositoryImpl struct {
	db	*gorm.DB
}

func NewUserRepository(ioc di.Container) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: ioc.Get(constants.POSTGRES).(*gorm.DB),
	}
}

func (r *UserRepositoryImpl) CreateUser(c echo.Context, user models.User) (err error) {
	err = r.db.Create(&user).WithContext(c.Request().Context()).Error
	return
}

func (r *UserRepositoryImpl) GetUserByID(c echo.Context, userID uint) (user *models.User, err error) {
	err = r.db.Where("id = ?", userID).Find(&user).Limit(1).WithContext(c.Request().Context()).Error
	if user.ID == 0 {
		return nil, nil
	}
	return
}

func (r *UserRepositoryImpl) GetUserByUsername(c echo.Context, username string) (user *models.User, err error) {
	err = r.db.Where("username = ?", username).Find(&user).Limit(1).WithContext(c.Request().Context()).Error
	if user.ID == 0 {
		return nil, nil
	}
	return
}

func (r *UserRepositoryImpl) UpdateUser(c echo.Context, user models.User) (err error) {
	err = r.db.Save(&user).WithContext(c.Request().Context()).Error
	return
}

func (r *UserRepositoryImpl) DeleteUser(c echo.Context, user models.User) (err error) {
	err = r.db.Delete(&user).WithContext(c.Request().Context()).Error
	return
}
