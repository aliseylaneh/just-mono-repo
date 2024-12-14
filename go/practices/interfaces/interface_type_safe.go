package main

type LogicProvider struct {
}

func (lp LogicProvider) end() {
	// TODO implement me
	panic("implement me")
}

func (lp LogicProvider) process() {
	//	TODO implement me
	panic("implement me")
}

type Logic interface {
	process()
	end()
}

type Client struct {
	logic Logic
}

func main() {
	client := Client{
		logic: LogicProvider{},
	}
	client.logic.process()
}
