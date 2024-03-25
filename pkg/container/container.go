package container

import (
	"log"

	"go.uber.org/dig"
)

var Registry = dig.New()

func Inject(constructor interface{}, opts ...dig.ProvideOption) {
	if err := Registry.Provide(constructor, opts...); err != nil {
		log.Fatal(err)
	}
}
