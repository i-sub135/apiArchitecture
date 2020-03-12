package repo

import (
	db "../Db"
	m "../Model"
	"github.com/jinzhu/gorm"

	// mysql Driver
	jwt "../Middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var sql *gorm.DB

func init() {
	sql = db.Conenect()
}

// GetUsers -- get data GetUsers
func GetUsers() []m.TUser {
	var data []m.TUser
	sql.Table("t_users").Find(&data)
	return data
}

// GetJwtUser -- get jwt
func GetJwtUser() interface{} {
	var data m.TUser
	sql.Table("t_users").Where("id=?", 3).First(&data)

	j, errs := jwt.CreateJwt(data.ID, data.Email)
	if errs != nil {
		return nil
	}
	out := map[string]interface{}{
		"UserID": data.ID,
		"Name":   data.Fullname,
		"Email":  data.Email,
		"token":  j,
	}

	return out
}
