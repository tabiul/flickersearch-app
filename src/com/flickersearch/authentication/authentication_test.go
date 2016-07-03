package authentication

import (
	"com/flickersearch/test"
	"testing"
)

func TestAddUser(t *testing.T) {
	var err error
	err = AddUser("test1", "abc")
	test.AssertNotError(t, err)
}

func TestCheckValidUser(t *testing.T) {
	var err error
	err = AddUser("test3", "abc")
	test.AssertNotError(t, err)

	err = CheckValidUser("test2", "abc")
	test.AssertError(t, err)

	err = CheckValidUser("test3", "abcd")
	test.AssertError(t, err)

	err = CheckValidUser("test3", "abc")
	test.AssertNotError(t, err)

}
