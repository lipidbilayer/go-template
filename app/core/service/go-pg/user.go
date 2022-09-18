package gopg

import (
	"car_pool/app/models"

	"github.com/go-pg/pg/v10/orm"
)

type DatabaseGoPG struct {
	DB orm.DB
}

func (d *DatabaseGoPG) GetUser(user *models.User) (*models.User, error) {
	err := d.DB.Model(user).Where("username = ?username").Where("password = crypt(?, password)", user.Password).First()
	return user, err
}
