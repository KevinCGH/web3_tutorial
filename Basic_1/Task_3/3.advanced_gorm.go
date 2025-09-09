package Task_3

import (
	"gorm.io/gorm"
)

// Model

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	PostCount int    `gorm:"default:0" json:"post_count"`
	Posts     []Post `gorm:"foreignKey:UserID" json:"posts"`
}

type Post struct {
	gorm.Model
	ID            uint      `gorm:"primaryKey" json:"id"`
	Title         string    `gorm:"not null" json:"title"`
	Content       string    `gorm:"type:text" json:"content"`
	UserID        uint      `gorm:"not null" json:"user_id"`
	User          User      `gorm:"foreignKey:UserID" json:"user"`
	Comments      []Comment `gorm:"foreignKey:PostID" json:"comments"`
	CommentCount  int       `gorm:"default:0" json:"comment_count"`
	CommentStatus string    `gorm:"default:无评论" json:"comment_status"`
}

type Comment struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey" json:"id"`
	Content string `gorm:"type:text;not null" json:"content"`
	PostID  uint   `gorm:"not null" json:"post_id"`
	Post    Post   `gorm:"foreignKey:PostID" json:"post"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
}

// BeforeCreate 钩子函数：在文章创建时自动更新用户的文章数量统计字段
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	if err := tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// AfterDelete 钩子函数：在文章删除时自动更新用户的文章数量统计字段
func (p *Post) AfterDelete(tx *gorm.DB) error {
	if err := tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) AfterCreate(tx *gorm.DB) error {
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}
	values := map[string]interface{}{
		"comment_count": count,
	}

	// 如果评论数为 0，则更新文章的评论状态
	if count > 0 {
		values["comment_status"] = "有评论"
	}
	if err := tx.Model(&Post{}).Where("id = ?", c.PostID).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

// AfterDelete 钩子函数：在评论删除时检查文章的评论数
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}
	values := map[string]interface{}{
		"comment_count": count,
	}
	// 如果评论数为 0，则更新文章的评论状态
	if count == 0 {
		values["comment_status"] = "无评论"
	}
	if err := tx.Model(&Post{}).Where("id = ?", c.PostID).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

// Service

type BlogService struct {
	DB *gorm.DB
}

func (s *BlogService) GetUserPostsAndComments(userID uint) ([]Post, error) {
	var posts []Post
	err := s.DB.Preload("Comments").Preload("Comments.User").Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

func (s *BlogService) GetMostCommentedPost() (*Post, error) {
	var post Post
	err := s.DB.Preload("User").Preload("Comments").Order("comment_count DESC").Find(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, err
}
