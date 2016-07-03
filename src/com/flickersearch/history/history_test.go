package history

import (
	"com/flickersearch/test"
	"testing"
)

func TestHistory(t *testing.T) {
	AddHistory("abc", "snow1")
	AddHistory("abc", "snow2")
	AddHistory("abc", "snow3")
	AddHistory("abc", "snow4")
	AddHistory("abc", "snow5")
	AddHistory("abc", "snow6")
	AddHistory("abc", "snow7")
	AddHistory("abc", "snow8")
	AddHistory("abc", "snow9")
	AddHistory("abc", "snow10")

	bytes, err := GetHistory("abc")
	test.AssertNotError(t, err)
	test.AssertEquals(t, "[\"snow1\",\"snow2\",\"snow3\",\"snow4\",\"snow5\",\"snow6\",\"snow7\",\"snow8\",\"snow9\",\"snow10\"]", string(bytes))
	AddHistory("abc", "snow11")
	bytes, err = GetHistory("abc")
	test.AssertNotError(t, err)
	test.AssertEquals(t, "[\"snow2\",\"snow3\",\"snow4\",\"snow5\",\"snow6\",\"snow7\",\"snow8\",\"snow9\",\"snow10\",\"snow11\"]", string(bytes))
	AddHistory("abc", "snow2")
	bytes, err = GetHistory("abc")
	test.AssertNotError(t, err)
	test.AssertEquals(t, "[\"snow2\",\"snow3\",\"snow4\",\"snow5\",\"snow6\",\"snow7\",\"snow8\",\"snow9\",\"snow10\",\"snow11\"]", string(bytes))
	bytes, err = GetHistory("cde")
	test.AssertNotError(t, err)
	test.AssertEquals(t, 0, len(bytes))
}
