package validations_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/validations"
)

func TestValidateEmailWithEmptyValue(t *testing.T) {
  m := new(M)

  ValidateEmail(m)

  assert.Contains(t, m.Errors().Get("Email"), "is invalid")
}

func TestValidateEmailWithInvalidValue(t *testing.T) {
  m := &M{ Email: "invalid" }

  ValidateEmail(m)

  assert.Contains(t, m.Errors().Get("Email"), "is invalid")
}

func TestValidateEmailWithValidValue(t *testing.T) {
  m := &M{ Email: "one@digits.com" }

  ValidateEmail(m)

  assert.NotContains(t, m.Errors().Get("Email"), "is invalid")
}
