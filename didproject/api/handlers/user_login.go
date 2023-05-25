package handlers

import (
    "database/sql"
    "net/http"
    "log"
    "strconv"
    "io/ioutil"

    "github.com/gin-gonic/gin"

    "didproject/api/models"
)

func Login(db *sql.DB) gin.HandlerFunc {
    log.SetOutput(ioutil.Discard) // 忽略日志输出

    return func(c *gin.Context) {
        var req models.RegisterRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        users := models.User{Username: req.Username, Password: req.Password}
        log.Println(users.Username)
        log.Println(users.Password)
        rows, err := db.Query("SELECT id, username, password FROM users where username=$1 AND password=$2", users.Username, users.Password)
        if err != nil {
            // 处理错误
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
            return
        }
        defer rows.Close()

        var id int
        var username, password string
        for rows.Next() {
            err = rows.Scan(&id, &username, &password)
            log.Println("id:", id, "username:", username, "password:", password)
            if username == users.Username && password == users.Password {
                // 登录验证通过，设置Cookie
                c.SetCookie("user_id", strconv.Itoa(id), 3600, "/", "localhost", false, true)
                c.SetCookie("username", username, 3600, "/", "localhost", false, true)
                c.JSON(http.StatusCreated, gin.H{"message": "Login successful", "user_id": id, "username": username})
                return
            }
        }

        // 登录验证失败
        c.JSON(http.StatusOK, gin.H{"message": "Invalid credentials"})
    }
}

func SomeHandler(c *gin.Context) {
    // 获取Cookie
    userID, err := c.Cookie("user_id")
    if err != nil {
        // 处理错误
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user ID"})
        return
    }

    username, err := c.Cookie("username")
    if err != nil {
        // 处理错误
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get username"})
        return
    }

    // 返回Cookie给前端
    c.JSON(http.StatusOK, gin.H{"user_id": userID, "username": username})
}
