package webcomponents

import (
	"qwserve/init/initDB"
	"qwserve/models"
)

func Add(val *models.WebComponents) (uint, error) {
	result := initDB.DB.Create(val)
	if result.Error != nil {
		return 0, result.Error
	}
	return val.ID, nil
}

func Get(query interface{}, args ...interface{}) (*models.WebComponents, error) {
	obj := &models.WebComponents{}
	result := initDB.DB.Where("", args...).First(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

func Gets(query interface{}, args ...interface{}) ([]*models.WebComponents, error) {
	rows, err := initDB.DB.Model(&models.WebComponents{}).Where(query, args...).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	lists := make([]*models.WebComponents, 0)
	for rows.Next() {
		var user *models.WebComponents
		// ScanRows 扫描每一行进结构体
		initDB.DB.ScanRows(rows, user)

		// 对每一个 User 进行操作
		lists = append(lists, user)
	}

	return lists, nil
}

func Set(query interface{}, args ...interface{}) error {

	result := initDB.DB.Model(&models.WebComponents{}).Where(query, args).Update("name", "hello")
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SetS(query interface{}, args ...interface{}) error {
	//db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	//db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	result := initDB.DB.Model(&models.WebComponents{}).Where(query, args).Update("name", "hello")
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Del(query interface{}, args ...interface{}) error {

	result := initDB.DB.Where(query, args).Delete(&models.WebComponents{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//db.Where("email LIKE ?", "%jinzhu%").Delete(&Email{})
//var users = []User{{ID: 1}, {ID: 2}, {ID: 3}}
//db.Delete(&users)
