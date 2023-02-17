package model

// 存储在数据库中的Video结构
type VideoData struct {
	Id            int64  `json:"id,omitempty" grom:"autoIncrement"`
	Author        int64  `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty" gorm:"default:https://www.w3schools.com/html/movie.mp4"`
	CoverUrl      string `json:"cover_url,omitempty" gorm:"default:https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"`
	FavoriteCount int64  `json:"favorite_count,omitempty" gorm:"default:0"`
	CommentCount  int64  `json:"comment_count,omitempty" gorm:"default:0"`
	IsFavorite    bool   `json:"is_favorite,omitempty" gorm:"type:tinyint(1);default:0"`
	CreateTime    int64  `gorm:"column:create_at;"`
}

// 用于返回的Video结构
type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	CreateTime    int64  `gorm:"column:create_at;"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	//Title         string `json:"title,omitempty"`
}

// PublishListVideoStruct as just store author's id in database, need another struct to meet the interface of PublishList
type PublishListVideoStruct struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}
