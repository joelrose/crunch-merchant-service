package test_helper

import (
	"time"

	"firebase.google.com/go/auth"
	"github.com/joelrose/crunch-merchant-service/models"
)

var (
	MockTokenUID = "1234567890"
	MockToken    = auth.Token{
		UID: MockTokenUID,
	}
	MockUser = models.User{
		Id:           1,
		FirebaseId:   MockTokenUID,
		LanguageCode: "en",
		Firstname:    "John",
		Lastname:     "Doe",
		CreatedAt:    time.Now().Add(-1 * time.Hour),
	}
)
