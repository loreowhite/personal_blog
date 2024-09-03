package main

import "gorm.io/gorm"

type Post struct {
    gorm.Model
    Title    string    `json:"title"`
    Content  string    `json:"content"`
    ImageURL string    `json:"image_url"`
    Likes    int       `json:"likes"`
    Comments []Comment `json:"comments" gorm:"foreignKey:PostID"`
}

type Comment struct {
    gorm.Model
    PostID  uint   `json:"post_id"`
    Content string `json:"content"`
    Author  string `json:"author"`
}
