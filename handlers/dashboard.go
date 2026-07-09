package handlers

import (
	"html/template"
	"net/http"
)

var dashboardTemplate = template.Must(template.ParseFiles("templates/dashboard.html"))

func Dashboard(w http.ResponseWriter, r *http.Request) {
	// 检查 cookie
	cookie, err := r.Cookie("username")
	if err != nil {
		// 没有 cookie，跳回登录页
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := struct{ Username string }{Username: cookie.Value}
	dashboardTemplate.Execute(w, data)
}
