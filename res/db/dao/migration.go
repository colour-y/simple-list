package dao

import "simplelist/res/db/model"

func migration() {
	err := _db.Set("gorm:simple_list", "charset=utf8mb4").AutoMigrate(&model.User{}, &model.Task{})

	if err != nil {
		return
	}

}
