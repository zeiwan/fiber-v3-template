package uberDig

import "go.uber.org/dig"

var container = dig.New()

func ProvideForDI(constructor any, opts ...dig.ProvideOption) error {
	return container.Provide(constructor, opts...)
}

func DI(function any, opts ...dig.InvokeOption) error {
	return container.Invoke(function, opts...)
}
