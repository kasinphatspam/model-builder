package usecases

import (
	"log"
	"mtrain-main/models"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type storer interface {
	Save(models.Account) error
	Find(map[string]interface{}) ([]models.Account, error)
	// FindOne(map[string]interface{}) (models.Account, error)
}

type Context interface {
	JSON(int, interface{})
	Bind(interface{}) error
}

type AccountHandler struct {
	Store storer
}

func NewAccountHandler(store storer) *AccountHandler {
	return &AccountHandler{Store: store}
}

func (h *AccountHandler) generateAccountID(c Context, prefix string) (string, error) {
	timestamp := time.Now().Unix()
	timestampStr := strconv.FormatInt(timestamp, 10)
	accountID := prefix + "_" + timestampStr

	// Check duplicate Account ID
	accounts, err := h.Store.Find(map[string]interface{}{
		"accountid": accountID,
	})
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if len(accounts) != 0 {
		return h.generateAccountID(c, prefix)
	}

	return accountID, nil
}

func (h *AccountHandler) generateHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(hashedPassword), nil
}

func (h *AccountHandler) CreateAccount(c Context) {
	// Bind http request body
	var account models.Account
	if err := c.Bind(&account); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	// Check duplicate email address
	accounts, _ := h.Store.Find(map[string]interface{}{
		"email": account.Email,
	})
	if len(accounts) != 0 {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "This email address is already in use. Please try another email.",
		})
		return
	}

	// Generate account id
	accountID, err := h.generateAccountID(c, "ACC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	account.AccountID = accountID

	// Generate a hash of the password
	hashedPassword, err := h.generateHashedPassword(account.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	account.Password = string(hashedPassword)

	// Save into database
	h.Store.Save(account)

	c.JSON(http.StatusCreated, map[string]string{
		"email": account.Email,
	})
}

func (h *AccountHandler) EditAccount(c Context) {

}

func (h *AccountHandler) DeleteAccount(c Context) {

}
