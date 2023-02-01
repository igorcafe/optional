package optional_test

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
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
		panic(fmt.Sprintf("Expected to be equal: %s == %s", a, b))
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

	t.Run("invalid filter should return empty", func(t *testing.T) {
		opt := optional.Of("abc").Filter(func(a any) bool {
			return len(a.(string)) == 2
		})
		mustEqual(true, opt.IsEmpty())
	})
}

// test FlatMap

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

func TestMap(t *testing.T) {
	t.Run("mapping present optional", func(t *testing.T) {
		res := optional.Of(123).Map(func(a any) any {
			return a.(int) * 2
		}).Map(func(a any) any {
			return strconv.Itoa(a.(int))
		}).Get()

		mustEqual("246", res.(string))
	})

	t.Run("mapping empty optional", func(t *testing.T) {
		res := optional.OfNullable(nil).Map(func(a any) any {
			return a.(int) * 2
		}).Map(func(a any) any {
			return strconv.Itoa(a.(int))
		}).IsEmpty()

		mustEqual(true, res)
	})
}

func TestOrElse(t *testing.T) {
	mustEqual(123, optional.Of(123).OrElse("no!"))
	mustEqual("yes!", optional.OfNullable(nil).OrElse("yes!"))
}

func TestOrElseGet(t *testing.T) {
	mustEqual("present", optional.Of("present").OrElseGet(func() any {
		return "impossible!"
	}))

	mustEqual("empty", optional.OfNullable(nil).OrElseGet(func() any {
		return "empty"
	}))
}

func TestOrElseThrow(t *testing.T) {
	optional.Of("present").OrElseThrow()

	mustPanic(func() {
		optional.OfNullable(nil).OrElseThrow()
	})
}

func TestOrElseThrowErr(t *testing.T) {
	optional.Of("present").OrElseThrowErr("nothing happens")

	mustPanic(func() {
		optional.OfNullable(nil).OrElseThrowErr("error!")
	})
}

func TestToString(t *testing.T) {
	mustEqual("Optional[%!s(int=123)]", optional.Of(123).ToString())
	mustEqual("Optional.empty", optional.OfNullable(nil).ToString())
	mustEqual("Optional[abc]", optional.Of("abc").ToString())
}
