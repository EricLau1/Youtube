package routes

import (
	"net/http"
	"go-webapp/utils"
	"go-webapp/sessions"
	"go-webapp/auth"
	"go-webapp/models"
	"fmt"
)

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
  _, isAuth := sessions.IsLogged(r)
  if isAuth {
    http.Redirect(w, r, "/admin", 302)
    return
  }
	message, alert := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "login.html", struct{
		Alert utils.Alert
	}{
		Alert: utils.NewAlert(message, alert),
	})
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	user, err := auth.Singin(email, password)
	checkErrAuthenticate(err, w, r, user)
}

func checkErrAuthenticate(err error, w http.ResponseWriter, r *http.Request, user models.User) {
	session, _ := sessions.Store.Get(r, "session")
	if err != nil {
		switch(err) {
			case auth.ErrEmptyFields,
				 auth.ErrEmailNotFound,
				 models.ErrInvalidEmail,
				 auth.ErrInvalidPassword:
				 session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
				 session.Values["ALERT"] = "danger"
				 session.Save(r, w)
				 http.Redirect(w, r, "/login", 302)
			default:
				utils.InternalServerError(w)
		}
		return
	}
	session.Values["USERID"] = user.Id
	session.Save(r, w)
	http.Redirect(w, r, "/admin", 302)	
}

func logoutGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	delete(session.Values, "USERID")
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}
