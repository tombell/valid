package valid_test

import (
	"testing"

	"github.com/tombell/valid"
)

func TestValid(t *testing.T) {
	age := 21

	validator := valid.New()
	validator.Check("age", valid.Case{Cond: valid.Min(age, 18), Msg: "must be at least 18"})

	if got := validator.Valid(); !got {
		t.Errorf("Valid() = %v, expected %v", got, true)
	}
}

func TestError(t *testing.T) {
	username := "tombell"

	validator := valid.New()
	validator.Check("username",
		valid.Case{Cond: valid.NotEmpty(username), Msg: "must not be empty"},
		valid.Case{Cond: valid.MinLength(username, 10), Msg: "must be at least 10 characters"},
	)

	if got := validator.Valid(); got {
		t.Errorf("Valid() = %v, expected %v", got, true)
	}

	got := validator.Errors.Error()
	expected := "username: must be at least 10 characters"

	if got != expected {
		t.Errorf("Errors.Error() = %v, expected %v", got, expected)
	}
}
