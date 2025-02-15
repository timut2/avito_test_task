package handlers_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/timut2/avito_test_task/internal/handlers"
)

func TestValidatePasswordLenError(t *testing.T) {
	in := "asdf"

	err := handlers.ValidatePassword(in)
	require.NotNil(t, err)
	require.EqualError(t, err, handlers.ErrPasswordTooShort.Error())

}

func TestValidatePasswordNumError(t *testing.T) {
	in2 := "asdfasdkfjla"
	err2 := handlers.ValidatePassword(in2)
	require.NotNil(t, err2)
	require.EqualError(t, err2, handlers.ErrPasswordNoDigit.Error())

}

func TestValidatePasswordSpecSymbolError(t *testing.T) {

	in3 := "asdfasdkfjla1"
	err3 := handlers.ValidatePassword(in3)
	require.NotNil(t, err3)
	require.EqualError(t, err3, handlers.ErrPasswordNoSpecialChar.Error())
}

func TestValidatePasswordNoError(t *testing.T) {
	in4 := "adsfjasd1321!"
	err4 := handlers.ValidatePassword(in4)
	require.Nil(t, err4)
}
