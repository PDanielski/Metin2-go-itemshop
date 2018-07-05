package account

import (
	"time"
)

//ID is an abstraction layer on top of the actual type of the identifier.
//If the identifier changes type, it will be easier to refactor
type ID int

//Account holds the representation of an account
type Account struct {
	id           ID
	login        string
	password     Password
	socialID     string
	email        string
	creationTime time.Time
	gold         int
	warpoints    int
	biscuits     int
}

//New returns a new instance of an Account. If the account id is unknown during instantation, provide 0 instead.
func New(id ID, login string, password Password, socialID string, email string, gold, warpoints, biscuits int) *Account {
	return &Account{id: id, login: login, password: password, socialID: socialID, email: email, gold: gold, warpoints: warpoints, biscuits: biscuits}
}

//IsIDValid returns true if the account has a valid id, which is any int != 0
func (acc *Account) IsIDValid() bool {
	return acc.id != ID(0)
}

//ID is the unique identifier of the account
func (acc *Account) ID() ID {
	return acc.id
}

//Login is the username used for signing in the game and in the website
func (acc *Account) Login() string {
	return acc.login
}

//Password is the password used for signing in the game and website
func (acc *Account) Password() Password {
	return acc.password
}

//SocialID is 7-sized numeric string code used by the users for confirming the cancellation of their in game players
func (acc *Account) SocialID() string {
	return acc.socialID
}

//Email is a email. Is it really? I don't know. I see only a string.
func (acc *Account) Email() string {
	return acc.email
}

//CreationTime returns the time on which the account was initially created
func (acc *Account) CreationTime() time.Time {
	return acc.creationTime
}

//Gold returns the balance of the gold currency
func (acc *Account) Gold() int {
	return acc.gold
}

//Warpoints returns the balance of the warpoint currency
func (acc *Account) Warpoints() int {
	return acc.warpoints
}

//Biscuits returns the balance of the biscuit currency
func (acc *Account) Biscuits() int {
	return acc.biscuits
}
