package account

import (
	"mt2is/pkg/test"
	"testing"
)

/*
 * Before running these tests make sure the account used as fixture exists in the test database
 */

var repo Repository

func init() {
	db, _ := test.NewDBConnection()
	repo = NewSQLRepository(db)
}

func TestByID(t *testing.T) {
	acc, ok := repo.ByID(ID(7))

	if ok != true {
		t.Error("ok must be true")
	}
	if acc == nil {
		t.Error("acc must be a pointer to a valid account")
	}

	if acc.ID() != ID(7) {
		t.Error("acc must have id 7")
	}

	if acc.Login() != "test" {
		t.Error("acc must have login 'test'")
	}

	if acc.Password() != "*B789F9621FFA4D4458099487777FD324AD8F02FC" {
		t.Error("acc must have as password '*B789F9621FFA4D4458099487777FD324AD8F02FC'")
	}
}

func TestByLogin(t *testing.T) {
	acc, ok := repo.ByLogin("test")

	if ok != true {
		t.Error("ok must be true")
	}
	if acc == nil {
		t.Error("acc must be a pointer to a valid account")
	}

	if acc.ID() != ID(7) {
		t.Error("acc must have id 7")
	}

	if acc.Login() != "test" {
		t.Error("acc must have login 'test'")
	}

	if acc.Password() != "*B789F9621FFA4D4458099487777FD324AD8F02FC" {
		t.Error("acc must have as password '*B789F9621FFA4D4458099487777FD324AD8F02FC'")
	}
}
