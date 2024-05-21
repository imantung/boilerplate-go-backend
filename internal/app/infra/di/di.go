// wrapper for dependency injection using dig
package di

import "go.uber.org/dig"

var container = dig.New() // global DI Container

// provide constructor function to DI Container
func Provide(constructor interface{}, opts ...dig.ProvideOption) interface{} {
	err := container.Provide(constructor, opts...)
	if err != nil {
		panic(err)
	}
	return nil
}

// invoke function using DI Container
func Invoke(fn interface{}, opts ...dig.InvokeOption) error {
	return container.Invoke(fn, opts...)
}
