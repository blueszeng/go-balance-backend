package models_test

import (
  "testing"

  "github.com/stretchr/testify/assert"

  . "github.com/tksasha/balance/config/test"
  . "github.com/tksasha/balance/models"
  . "github.com/tksasha/balance/rest/validate"
)

func TestCashErrors(t *testing.T) {
  cash := new(Cash)

  cash.Build(Values)

  assert.NotPanics(t, func() {
    cash.Errors().Add("foo", "is invalid")
  })
}

func TestCashBuild(t *testing.T) {
  cash := new(Cash)

  Values.Set("cash[name]", "In the Purse")

  Values.Set("cash[sum]", "16.45")

  cash.Build(Values)

  assert.Equal(t, "In the Purse", cash.Name)

  assert.Equal(t, 16.45, cash.Sum)
}

func TestCashValidatePresenceOfName(t *testing.T) {
  cash := new(Cash)

  Values.Set("cash[name]", "")

  cash.Build(Values)

  Validate(cash)

  assert.False(t, cash.IsValid())

  assert.Equal(t, []string{"can't be blank"}, cash.Errors().Get("Name"))
}
