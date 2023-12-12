package main

import "log"

type IntegrationResponse struct {
	status string
	body   string
}

type RealIntegration struct{}

type MockIntegration struct{}

type Integration interface {
	MakeIntegration(method, url, body string) *IntegrationResponse
}

func (r RealIntegration) MakeIntegration(method, url , body string) *IntegrationResponse {
	log.Printf("Requisicao Real: method=%s,url=%s,body=%s",method,url,body)
	return &IntegrationResponse{
		status: "OK",
		body:   "",
	}
}

func (m MockIntegration) MakeIntegration(method, url, body string) *IntegrationResponse {
	log.Printf("Requisicao Mockada: method=%s,url=%s,body=%s \n",method,url,body)
	return nil
}

type Controller struct {
	Integration Integration
}

type InterfaceController interface {
	Request()
}

func (c Controller) Request() {
	c.Integration.MakeIntegration("POST","login","{'username':'joaozinho','pass':'123'}")
}

func main() {
	controller:= Controller {
		Integration: RealIntegration{},
	}

	controller.Request()

}
