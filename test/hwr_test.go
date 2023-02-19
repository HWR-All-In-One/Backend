package test

import (
	"os"
	"testing"

	"github.com/HWR-All-In-One/Backend/internal/pkg/hwr"
)

func TestValidateUser(t *testing.T) {
	email := os.Getenv("HWR_EMAIL")
	password := os.Getenv("HWR_PASSWORD")

	isValid, err := hwr.ValidateUser(email, password)

	if err != nil {
		t.Error(err)
	}

	if !isValid {
		t.Error("should be valid user but is not")
	}
}

func TestIsNotValidUser(t *testing.T) {
	email := os.Getenv("HWR_EMAIL") + "fail"
	password := os.Getenv("HWR_PASSWORD") + "fail"

	isValid, err := hwr.ValidateUser(email, password)

	if err != nil {
		t.Error(err)
	}

	if isValid {
		t.Error("should not be valid user but is")
	}

}
