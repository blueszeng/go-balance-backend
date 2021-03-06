package validate

import (
  "reflect"
  "strings"
  "regexp"
  "strconv"

  . "github.com/tksasha/balance/rest/resource"
  . "github.com/tksasha/balance/rest/validate/presence_of"
  . "github.com/tksasha/balance/rest/validate/format_of_email"
  . "github.com/tksasha/balance/rest/validate/length_of"
  . "github.com/tksasha/balance/rest/validate/greater_than"
  . "github.com/tksasha/balance/rest/validate/uniqueness_of"
)

func Validate(resource Resource) {
  reflected := reflect.ValueOf(resource).Elem()

  for idx := 0; idx < reflected.NumField(); idx++ {
    tags := strings.Split(reflected.Type().Field(idx).Tag.Get("valid"), ",")

    for _, tag := range tags {
      attribute := reflected.Type().Field(idx).Name

      switch tag {
      case "present":
        ValidatePresenceOf(resource, attribute)
      case "email":
        ValidateFormatOfEmail(resource, attribute)
      case "unique":
        ValidateUniquenessOf(resource, attribute)
      }

      lengthAtLeast(resource, attribute, tag)

      lengthAtMost(resource, attribute, tag)

      lengthExactly(resource, attribute, tag)

      greaterThan(resource, attribute, tag)
    }
  }
}

func lengthAtLeast(resource Resource, attribute string, tag string) {
  mask := regexp.MustCompile(`length_at_least\((\d+)\)`)

  if mask.MatchString(tag) {
    var length int

    var err error

    if length, err = strconv.Atoi(mask.FindStringSubmatch(tag)[1]); err != nil {
      length = 0
    }

    ValidateLengthOf(resource, attribute, "at least", length)
  }
}

func lengthAtMost(resource Resource, attribute string, tag string) {
  mask := regexp.MustCompile(`length_at_most\((\d+)\)`)

  if mask.MatchString(tag) {
    var length int

    var err error

    if length, err = strconv.Atoi(mask.FindStringSubmatch(tag)[1]); err != nil {
      length = 0
    }

    ValidateLengthOf(resource, attribute, "at most", length)
  }
}

func lengthExactly(resource Resource, attribute string, tag string) {
  mask := regexp.MustCompile(`length_exactly\((\d+)\)`)

  if mask.MatchString(tag) {
    var length int

    var err error

    if length, err = strconv.Atoi(mask.FindStringSubmatch(tag)[1]); err != nil {
      length = 0
    }

    ValidateLengthOf(resource, attribute, "exactly", length)
  }
}

func greaterThan(resource Resource, attribute string, tag string) {
  value := reflect.ValueOf(resource).Elem()

  value = reflect.Indirect(value).FieldByName(attribute)

  switch value.Type().Name() {
  case "int":
    greaterThanForInt(resource, attribute, tag)
  case "float64":
    greaterThanForFloat(resource, attribute, tag)
  }
}

func greaterThanForInt(resource Resource, attribute string, tag string) {
  mask := regexp.MustCompile(`greater_than\((\d+)\)`)

  if mask.MatchString(tag) {
    var value int

    var err error

    if value, err = strconv.Atoi(mask.FindStringSubmatch(tag)[1]); err != nil {
      value = 0
    }

    ValidateGreaterThan(resource, attribute, value)
  }
}

func greaterThanForFloat(resource Resource, attribute string, tag string) {
  mask := regexp.MustCompile(`greater_than\((\d+|\d+\.\d+)\)`)

  if mask.MatchString(tag) {
    var value float64

    var err error

    if value, err = strconv.ParseFloat(mask.FindStringSubmatch(tag)[1], 64); err != nil {
      value = 0.0
    }

    ValidateGreaterThan(resource, attribute, value)
  }
}
