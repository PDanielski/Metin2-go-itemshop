package main

import (
	"fmt"
	"html/template"
	"mt2is/pkg/account"
	"mt2is/pkg/auth"
	"net/http"
)

type LoginHandler struct {
	//store sessions.Store
	repo account.Repository
}

func (h *LoginHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		const html = `<div><form method="POST">
			<div>
				<label for="login">Login</label>
				<input type="text" name="login"/>
			</div>
			<div>
				<label for="psw">Password</label>
				<input type="password" name="psw"/>
			</div>
			<div>
				<input type="submit" value="Login now"/>
			</div>
		</form></div>`
		tpl := template.New("test")
		t, err := tpl.Parse(html)
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(writer, nil)
		return
	}

	request.ParseForm()
	loginpsw := auth.NewLoginPswAuthentication(h.repo)
	login := request.PostFormValue("login")
	psw := request.PostFormValue("psw")
	if !(login != "" && psw != "") {
		fmt.Fprintln(writer, "You must insert both login and psw")
		return
	}
	onFailure := auth.IfNotAuthenticated(func() {
		fmt.Fprintln(writer, "Invalid credentials")
	})
	onSuccess := auth.IfAuthenticated(func(acc *account.Account) {
		fmt.Fprintf(writer, "Welcome %s,\n your id is %d", acc.Login(), acc.ID())
	})
	loginpsw.Authenticate(login, psw, onSuccess, onFailure)
}
