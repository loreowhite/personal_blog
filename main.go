package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func main() {
    InitDB()
    router := NewRouter()
    fmt.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    if err := DB.AutoMigrate(&Post{}, &Comment{}); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/admin/posts", CreatePost).Methods("POST")
    router.HandleFunc("/admin/posts/{id}", UpdatePost).Methods("PUT")
    router.HandleFunc("/admin/posts/{id}", DeletePost).Methods("DELETE")
    router.HandleFunc("/posts", GetPosts).Methods("GET")
    router.HandleFunc("/posts/{id}", GetPost).Methods("GET")
    router.HandleFunc("/posts/{id}/like", LikePost).Methods("POST")
    router.HandleFunc("/posts/{id}/comments", AddComment).Methods("POST")

    return router
}
