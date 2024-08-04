package models

import "gorm.io/gorm"

/* -------------------------------- Types ----------------------------------*/

type Post struct {
    gorm.Model
    Title    string `json:"title"`
    Body     string `json:"body"`
    AuthorID uint   `json:"author_id"`
    Author   User   `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
}


/* -------------------------------- DB Actions ----------------------------------*/

// Save post 
func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	err := db.Create(&p).Error
	if err != nil {
		return &Post{}, err
	}

	return p, nil
}

// Update post 
func (p *Post) UpdatePost(db *gorm.DB) (*Post, error) {
	err := db.Save(&p).Error
	if err != nil {
		return &Post{}, err
	}

	return p, nil
}

// Delete a post 
func (p *Post) DeletePost(db *gorm.DB) error {
	err := db.Delete(&p).Error
	
	return err
}

// Find post by ID 
func (p *Post) FindPostByID(db *gorm.DB, postID string) (*Post, error) {
	
	err := db.Where("id = ?", postID).First(&p).Error
	if err != nil {
		return &Post{}, err
	}
	
	return p, nil
}