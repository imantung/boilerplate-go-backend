// wrapper for dependency injection using dig
package di

import "go.uber.org/dig"

var container = dig.New() // global DI Container

// Provide constructor function to DI Container
// The recommendation is to provide constructor in the same file with the definition. Here is a example:
//
//	var _ = di.Provide(YourContructor)
//	func YourContructor() {
//		// ...
//	}
func Provide(constructor interface{}, opts ...dig.ProvideOption) interface{} {
	err := container.Provide(constructor, opts...)
	if err != nil {
		panic(err)
	}
	return nil
}

// Invoke function using DI Container
func Invoke(fn interface{}, opts ...dig.InvokeOption) error {
	return container.Invoke(fn, opts...)
}
