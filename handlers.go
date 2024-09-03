package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
    var post Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := DB.Create(&post).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
    var posts []Post
    if err := DB.Preload("Comments").Find(&posts).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var post Post
    if err := DB.Preload("Comments").First(&post, id).Error; err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var post Post
    if err := DB.First(&post, id).Error; err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }

    var updatedPost Post
    if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    post.Title = updatedPost.Title
    post.Content = updatedPost.Content
    post.ImageURL = updatedPost.ImageURL

    if err := DB.Save(&post).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var post Post
    if err := DB.First(&post, id).Error; err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }

    if err := DB.Delete(&post).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var post Post
    if err := DB.First(&post, id).Error; err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }

    post.Likes++

    if err := DB.Save(&post).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(post)
}

func AddComment(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var post Post
    if err := DB.First(&post, id).Error; err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }

    var comment Comment
    if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    comment.PostID = post.ID

    if err := DB.Create(&comment).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(comment)
}
