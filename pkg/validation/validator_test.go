package validation_test

import (
	"testing"

	"github.com/payfazz/golib/pkg/validation"
)

func TestEmailValidation(t *testing.T) {
	if validation.IsEmail("") {
		t.Error("expected value:'false' as result of validation.IsEmail()")
	}
	if validation.IsEmail("plain") {
		t.Error("expected value:'false' as result of validation.IsEmail()")
	}
	if validation.IsEmail("some space@mail.com") {
		t.Error("expected value:'false' as result of validation.IsEmail()")
	}
	if validation.IsEmail("@gmail.com") {
		t.Error("expected value:'false' as result of validation.IsEmail()")
	}

	if !validation.IsEmail("someone@example") {
		t.Error("expected value:'true' as result of validation.IsEmail()")
	}
	if !validation.IsEmail("someone@example.go") {
		t.Error("expected value:'true' as result of validation.IsEmail()")
	}
	if !validation.IsEmail("someone@example.go.lang") {
		t.Error("expected value:'true' as result of validation.IsEmail()")
	}
}

func TestPhoneSanitization(t *testing.T) {
	expected := "+62812345678910"
	if validation.SanitizePhone("0812-3456-7891-0") != expected {
		t.Errorf("expected value:'%s' as result of validation.SanitizePhone()", expected)
	}
	if validation.SanitizePhone("0812345678910") != expected {
		t.Errorf("expected value:'%s' as result of validation.SanitizePhone()", expected)
	}
	if validation.SanitizePhone("62812345678910") != expected {
		t.Errorf("expected value:'%s' as result of validation.SanitizePhone()", expected)
	}
	if validation.SanitizePhone("+62812345678910") != expected {
		t.Errorf("expected value:'%s' as result of validation.SanitizePhone()", expected)
	}
	if validation.SanitizePhone("(62)812-3456-7891-0") != expected {
		t.Errorf("expected value:'%s' as result of validation.SanitizePhone()", expected)
	}
}

func TestValidPassword(t *testing.T) {
	password := "Ultraman123"

	valid := validation.IsValidPassword(password)
	if !valid {
		t.Error("expected true")
	}
}
