package main

type TypeOneInterface interface {
	write()
}

type TypeTwoInterface interface {
	get()
}

type EmbeddedInterface interface {
	TypeOneInterface
	TypeTwoInterface
}

func testEmbeddedInterfaces(value EmbeddedInterface) {
	value.get()
	value.write()
}
