package optional_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/igoracmelo/optional"
)

func mustPanic(f func()) {
	defer func() {
		if r := recover(); r == nil {
			panic("Didn't panic")
		}
	}()

	f()
}

func mustEqual(a any, b any) {
	if a != b {
		panic("Must be equal")
	}
}

func TestOf(t *testing.T) {
	t.Run("should create an optional", func(t *testing.T) {
		optional.Of(1)
		optional.Of("hello world")
		optional.Of(http.Client{})
	})

	t.Run("should throw NullPointerException", func(t *testing.T) {
		mustPanic(func() {
			optional.Of(nil)
		})
	})
}

func TestOfNullable(t *testing.T) {
	t.Run("should create an optional for any value", func(t *testing.T) {
		optional.OfNullable('a')
		optional.OfNullable(123i)
		optional.OfNullable(os.Stdin)
		optional.OfNullable(nil)
	})
}

func TestEquals(t *testing.T) {
	t.Run("should be equal", func(t *testing.T) {
		mustEqual(true, optional.Of('a').Equals('a'))
		mustEqual(true, optional.Of("string").Equals("string"))
	})

	t.Run("should not be equal", func(t *testing.T) {
		mustEqual(false, optional.Of('a').Equals("a"))
		mustEqual(false, optional.Of("string").Equals([]byte("string")))
	})
}

func TestFilter(t *testing.T) {
	t.Run("should pass filter", func(t *testing.T) {
		mustEqual(true, optional.Of('a').Filter(func(a any) bool {
			return true
		}).IsPresent())

		mustEqual(true, optional.Of("abc").Filter(func(a any) bool {
			return len(a.(string)) == 3
		}).IsPresent())
	})
}

func TestGet(t *testing.T) {
	t.Run("should get", func(t *testing.T) {
		mustEqual(1, optional.Of(1).Get())
		mustEqual("hello world", optional.Of("hello world").Get())
	})

	t.Run("should throw NoSuchElementException", func(t *testing.T) {
		mustPanic(func() {
			optional.OfNullable(nil).Get()
		})
	})
}

func TestIfPresent(t *testing.T) {
	t.Run("should call func", func(t *testing.T) {
		called := false

		optional.Of(1).IfPresent(func() {
			called = true
		})

		mustEqual(true, called)
	})

	t.Run("shouldn't call func", func(t *testing.T) {
		called := false

		optional.OfNullable(nil).IfPresent(func() {
			called = true
		})

		mustEqual(false, called)
	})
}
