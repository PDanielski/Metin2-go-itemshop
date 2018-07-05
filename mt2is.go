package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mt2is/pkg/account"
	"mt2is/pkg/wallet"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

const configDir = "conf/conf.json"

type Configuration struct {
	DatabaseDsn string
	TemplateDir string
}

func main() {
	conf, err := initConfig()
	if err != nil {
		panic(fmt.Errorf("Can't read configuration: %v", err))
	}

	db, err := initDB(conf.DatabaseDsn)
	if err != nil {
		panic(fmt.Errorf("Can't connect to database: %v", err))
	}

	currencies, err := initCurrencies()
	if err != nil {
		panic(fmt.Errorf("Can't read currencies: %v", err))
	}
	_ = currencies // TODO clean

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	sessStore := sessions.NewCookieStore([]byte("secret key"))
	repo := account.NewSQLRepository(db)
	checkAuth := SecurityMiddleware(sessStore, repo)

	catNodeProvider := &SQLNodeTreeProvider{db}
	catHandler, err := NewCategoryHandler(catNodeProvider)
	if err != nil {
		panic(err)
	}
	http.Handle("/category/", checkAuth(catHandler))
	loginHandler := &LoginHandler{repo}
	http.Handle("/login", loginHandler)
	http.ListenAndServe(":8080", nil)
}

func initConfig() (*Configuration, error) {
	jsonConfig, err := ioutil.ReadFile(configDir)
	if err != nil {
		return nil, err
	}
	conf := &Configuration{}
	err = json.Unmarshal(jsonConfig, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func initDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func initCurrencies() ([]*wallet.Currency, error) {
	fJSON, err := ioutil.ReadFile("conf/currencies.json")
	if err != nil {
		return nil, err
	}
	provider, err := NewJsonCurrencyProvider(fJSON)
	if err != nil {
		return nil, err
	}
	return provider.Provide()
}
