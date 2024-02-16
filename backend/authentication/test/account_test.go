package test

import (
	"mtrain-main/models"
	"mtrain-main/usecases"
	"net/http"
	"testing"
)

type fakeStore struct{}

func (fakeStore) Find(map[string]interface{}) ([]models.Account, error) {
	return []models.Account{
		{
			Email:       "example@gmail.com",
			Password:    "admin123",
			FirstName:   "John",
			LastName:    "Doe",
			AccountID:   "ACC_001",
			AccountType: "default",
		},
		{
			Email:       "example2@gmail.com",
			Password:    "admin123",
			FirstName:   "Alice",
			LastName:    "Wonder",
			AccountID:   "ACC_002",
			AccountType: "default",
		},
	}, nil
}

func (fakeStore) Save(models.Account) error {
	return nil
}

type fakeStoreHasNoData struct{}

func (fakeStoreHasNoData) Find(map[string]interface{}) ([]models.Account, error) {
	return []models.Account{}, nil
}

func (fakeStoreHasNoData) Save(models.Account) error {
	return nil
}

func (*fakeStoreHasNoData) Bind(interface{}) error {
	return nil
}

type fakeContext struct {
	code    int
	respone interface{}
}

func (c *fakeContext) JSON(code int, v interface{}) {
	c.code = code
	c.respone = v.(map[string]string)
}

func (*fakeContext) Bind(interface{}) error {
	return nil
}

func TestCreateAccountSuccessful(t *testing.T) {
	handler := &usecases.AccountHandler{
		Store: fakeStoreHasNoData{},
	}
	c := &fakeContext{}
	handler.CreateAccount(c)

	want := http.StatusCreated

	if want != c.code {
		t.Errorf("%d is expected but got %d", want, c.code)
	}
}

func TestCreateAccountEmailAlreadyTaken(t *testing.T) {
	handler := &usecases.AccountHandler{
		Store: fakeStore{},
	}
	c := &fakeContext{}
	handler.CreateAccount(c)

	want := http.StatusBadRequest

	if want != c.code {
		t.Errorf("%d is expected but got %d", want, c.code)
	}
}
