package models

type Account struct {
	AccountID   string `json:"accountID" xml:"accountID" form:"accountID"`
	AccountType string `json:"accountType" xml:"accountType" form:"accountType"`
	Email       string `json:"email" xml:"email" form:"email"`
	Password    string `json:"password" xml:"password" form:"password"`
	FirstName   string `json:"firstName" xml:"firstName" form:"firstName"`
	LastName    string `json:"lastName" xml:"lastName" form:"lastName"`
	PhoneNumber string `json:"phoneNumber" xml:"phoneNumber" form:"phoneNumber"`
}
