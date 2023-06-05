package system

import (
	"learn-gin/global"
	"learn-gin/model/base"
	"learn-gin/model/system"
	"learn-gin/model/system/query"
)

type UserRepository struct {
}

func (UserRepository) Create(user *system.User) uint {
	if err := global.DB.Create(&user).Error; err != nil {
		panic(err)
	}
	return user.ID
}

func (UserRepository) Update(user *system.User) uint {
	if err := global.DB.Model(&system.User{}).Where("id = ?", user.ID).Update("name", user.Username).Error; err != nil {
		panic(err)
	}
	return user.ID
}

func (UserRepository) Delete(id uint) uint {
	if err := global.DB.Delete(&system.User{}, id).Error; err != nil {
		panic(err)
	}
	return id
}

func (UserRepository) GetById(id uint) *system.User {
	var user *system.User
	if err := global.DB.Limit(1).Find(&user, id).Error; err != nil {
		panic(err)
	}
	return user
}

func (UserRepository) GetByName(username string) *system.User {
	var user *system.User
	if err := global.DB.Limit(1).Where("username = ?", username).First(&user).Error; err != nil {
		panic(err)
	}
	return user
}

func (UserRepository) List() []*system.User {
	var list []*system.User
	if err := global.DB.Find(&list).Error; err != nil {
		panic(err)
	}
	return list
}

func (UserRepository) ListByIds(ids []uint) []*system.User {
	var list []*system.User
	if err := global.DB.Find(&list, ids).Error; err != nil {
		panic(err)
	}
	return list
}

func (UserRepository) Page(query *query.UserQuery) *base.Page {
	var list []*system.User

	queryMap := make(map[string]interface{})

	if query.Name != "" {
		queryMap["name"] = query.Name
	}

	result := global.DB.Where(queryMap).Limit(int(query.Limit)).Offset(int(query.Offset)).Find(&list)
	if result.Error != nil {
		panic(result.Error)
	}

	return &base.Page{
		PageNo:   query.PageNo,
		PageSize: query.PageSize,
		Data:     list,
		Total:    uint(result.RowsAffected),
	}
}
