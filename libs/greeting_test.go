package libs_test

import (
	"testing"
	"wichy/libs"
)

func TestLogin(t *testing.T) {
	libs.GreetingText()
	if libs.Login("wichy", "password") == false {
		t.Errorf("password or username worng")
	}
}
