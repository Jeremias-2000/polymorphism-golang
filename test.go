package main

import "fmt"

type AbstractClient struct{}

type AbstractClientInterface interface {
	DoRequest()
}

func (a AbstractClient) DoRequest() {
	fmt.Println("AbstractClient is making a request.")
}

type Integration struct {
	ClientInterface AbstractClientInterface
}

type IntegrationInterface interface {
	DoSomething()
}

func (i Integration) DoSomething() {
	fmt.Println("Integration is doing something.")
}

type Handler struct {
	IntegrationInterface IntegrationInterface
}

func main() {
	handler := Handler{
		IntegrationInterface: Integration{
			ClientInterface: AbstractClient{},
		},
	}

	// Now you can call DoSomething through the Handler.
	handler.IntegrationInterface.DoSomething()

	// You can also access DoRequest through the AbstractClient.
	handler.IntegrationInterface.(Integration).ClientInterface.DoRequest()
}
