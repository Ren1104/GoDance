package service

import (
	"GoDance/model"
	"errors"
	"fmt"
	"sync/atomic"
)

var UserIdSequence = int64(1)

func FeedResponseUser(user model.User) model.User {
	return model.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
		TotalFav:      user.TotalFav,
		FavCount:      user.FavCount,
	}
}

func GetUserInfoVO(userId int64) (model.User, error) {
	db := model.Db
	user := model.UserData{}
	var userVO model.User
	if err := db.Table("user_data").Where("user_id=  ?", userId).Find(&user).Error; err != nil {
		return userVO, err
	}
	fmt.Println(user)
	userVO.FollowCount = user.FollowCount
	userVO.Id = user.Id
	userVO.Name = user.Name
	userVO.FollowerCount = user.FollowerCount
	userVO.IsFollow = user.IsFollow
	return userVO, nil
}

func Login(username string, password string) (model.UserData, error) {
	var user model.UserData
	model.Db.Where("name=? AND password=?", username, password).Find(&user)
	fmt.Println(user)
	if user.Id == 0 {
		return model.UserData{}, errors.New("login failed,please check your username or password")
	} else {
		fmt.Println("success")
		return user, nil
	}
}

func Register(username string, password string) (model.UserData, error) {
	var user model.UserData
	model.Db.Find(&user, "name=?", username)
	if user.Id != 0 {
		return model.UserData{}, errors.New("user already exist")
	} else {
		atomic.AddInt64(&UserIdSequence, 1)
		user := model.UserData{
			Id:            UserIdSequence,
			Name:          username,
			Password:      password,
			FollowCount:   0,
			FollowerCount: 0,
			TotalFav:      0,
			FavCount:      0,
			IsFollow:      false,
			Token:         username + password,
		}
		if model.Db.Create(&user).RowsAffected == 1 {
			return user, nil
		} else {
			return model.UserData{}, errors.New("register failed")
		}
	}
}
func FindUserByToken(token string) (model.User, error) {
	var user model.UserData
	fmt.Println(token)
	model.Db.Table("user_data").Find(&user, "token=?", token)
	fmt.Println(user)
	if user.Id == 0 {
		return model.User{}, errors.New("user doesn't exist")
	} else {
		ret_user := model.User{
			Id:            user.Id,
			Name:          user.Name,
			Password:      user.Password,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowCount,
			TotalFav:      user.TotalFav,
			FavCount:      user.FavCount,
			IsFollow:      user.IsFollow,
			Token:         user.Token,
		}
		fmt.Println(ret_user)
		return ret_user, nil
	}
}

func FindUserById(id int64) (model.User, error) {
	var user model.UserData
	model.Db.Table("user_data").Find(&user, "user_id=?", id)
	if user.Id == 0 {
		return model.User{}, errors.New("user doesn't exist")
	} else {
		ret_user := model.User{
			Id:            user.Id,
			Name:          user.Name,
			Password:      user.Password,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowCount,
			TotalFav:      user.TotalFav,
			FavCount:      user.FavCount,
			IsFollow:      user.IsFollow,
			Token:         user.Token,
		}
		fmt.Println(ret_user)
		return ret_user, nil
	}
}
