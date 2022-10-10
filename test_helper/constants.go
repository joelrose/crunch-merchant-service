package test_helper

import (
	"time"

	"firebase.google.com/go/auth"
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/models"
)

var (
	MockTokenUID = "1234567890"
	MockToken    = auth.Token{
		UID: MockTokenUID,
	}
	MockUser = models.User{
		Id:           uuid.New(),
		FirebaseId:   MockTokenUID,
		LanguageCode: "en",
		Firstname:    "John",
		Lastname:     "Doe",
		CreatedAt:    time.Now().UTC().Add(-1 * time.Hour),
	}
)
