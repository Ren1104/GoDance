package config

type Comment struct {
	ID        uint `gorm:"primaryKey"` //主键
	CreatedAt int
	VideoID   uint   `gorm:"not null,index"` //非空，创建索引
	UserID    uint   `gorm:"not null"`       //非空
	Content   string `gorm:"not null"`       //非空
}

// 添加评论
func CommentAdd(c Comment) error {

	err := DB.Table("comments").Create(&c).Error
	if err != nil {
		return err
	}
	return nil

}

// 删除评论
func CommentDelete(id int) error {

	var comment Comment
	err := DB.Table("comments").Delete(&comment, id).Error
	if err != nil {
		return err
	}
	return nil
}
