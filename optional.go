package optional

import "fmt"

type Optional struct {
	value any
}

// "static methods"

func Empty() Optional {
	return Optional{}
}

func Of(value any) Optional {
	if value == nil {
		panic("NullPointerException: null")
	}

	return OfNullable(value)
}

func OfNullable(value any) Optional {
	return Optional{
		value: value,
	}
}

// "object methods"

func (o Optional) Equals(value any) bool {
	return o.value == value
}

func (o Optional) Filter(f func(any) bool) Optional {
	if o.IsEmpty() || !f(o.value) {
		return Empty()
	}

	return Of(o.value)
}

func (o Optional) FlatMap(f func(any) any) Optional {
	panic("todo")
}

func (o Optional) Get() any {
	if o.IsEmpty() {
		panic("NoSuchElementException: No value present")
	}

	return o.value
}

func (o Optional) IfPresent(f func()) {
	if o.IsPresent() {
		f()
	}
}

func (o Optional) IsPresent() bool {
	return o.value != nil
}

func (o Optional) Map(f func(any) any) Optional {
	if o.IsEmpty() {
		return o
	}

	return OfNullable(f(o.value))
}

func (o Optional) IsEmpty() bool {
	return !o.IsPresent()
}

func (o Optional) OrElse(value any) any {
	if o.IsPresent() {
		return o.value
	}
	return value
}

func (o Optional) OrElseGet(f func() any) any {
	if o.IsPresent() {
		return o.value
	}
	return f()
}

func (o Optional) OrElseThrow() any {
	if o.IsPresent() {
		return o.value
	}

	panic("NullPointerException: null")
}

func (o Optional) OrElseThrowErr(err any) any {
	if o.IsPresent() {
		return o.value
	}

	panic(err)
}

func (o Optional) ToString() string {
	if o.IsEmpty() {
		return "Optional.empty"
	}

	return fmt.Sprintf("Optional[%s]", o.value)
}
