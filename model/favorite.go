package model

type FavoriteData struct {
	Token   string `json:"column:token"`
	VideoId string `json:"column:video_id"`
	//user_id  int64  `gorm:"column:user_id"`	//because FavoriteAction has no req para named user_id
}

type Favorite struct {
}
