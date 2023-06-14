package service

import "go-rest-user/model"

type UserAttribute model.UserAttribute

func (userAttribute *UserAttribute) Save() (*UserAttribute, error) {
	err := model.Database.Create(&userAttribute).Error

	if err != nil {
		return &UserAttribute{}, err
	}

	return userAttribute, nil
}

func FindUserById(id uint) (UserAttribute, error) {

	var user UserAttribute

	err := model.Database.Where("id=?", id).Find(&user).Error

	if err != nil {
		return UserAttribute{}, nil
	}

	return user, nil

}

func Update(data *UserAttribute) (err error) {
	model.Database.Save(&data)
	return nil
}

func Delete(id uint) (UserAttribute, error) {
	var user UserAttribute

	if err := model.Database.Where("id=?", id).Delete(&user); err != nil {
		return UserAttribute{}, nil
	}

	return user, nil
}
