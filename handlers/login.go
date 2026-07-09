package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"login-system/database"

	"golang.org/x/crypto/bcrypt"
)

var loginTemplate = template.Must(template.ParseFiles("templates/login.html"))

// LoginPage - 显示登录页面
func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	loginTemplate.Execute(w, nil)
}

// Login - 处理登录请求
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// 从数据库查询用户
	var hashedPassword string
	err := database.DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hashedPassword)

	if err == sql.ErrNoRows {
		// 用户不存在
		data := struct{ Error string }{Error: "Invalid username or password"}
		loginTemplate.Execute(w, data)
		return
	} else if err != nil {
		log.Println("Database error:", err)
		data := struct{ Error string }{Error: "Server error. Please try again."}
		loginTemplate.Execute(w, data)
		return
	}

	// 比对密码
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// 密码错误
		data := struct{ Error string }{Error: "Invalid username or password"}
		loginTemplate.Execute(w, data)
		return
	}

	// ✅ 登录成功！设置 cookie 然后跳转到 dashboard
	http.SetCookie(w, &http.Cookie{
		Name:  "username",
		Value: username,
		Path:  "/",
	})

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
