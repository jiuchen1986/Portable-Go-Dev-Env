package main

import (
	"app"
	"github.com/goadesign/goa"
)

// TestServiceController implements the Test Service resource.
type TestServiceController struct {
	*goa.Controller
}

// NewTestServiceController creates a Test Service controller.
func NewTestServiceController(service *goa.Service) *TestServiceController {
	return &TestServiceController{Controller: service.NewController("TestServiceController")}
}

// LocalService runs the local service action.
func (c *TestServiceController) LocalService(ctx *app.LocalServiceTestServiceContext) error {
	// TestServiceController_LocalService: start_implement

	// Put your logic here

	// TestServiceController_LocalService: end_implement
	return nil
}

// ServiceChain runs the service chain action.
func (c *TestServiceController) ServiceChain(ctx *app.ServiceChainTestServiceContext) error {
	// TestServiceController_ServiceChain: start_implement

	// Put your logic here

	// TestServiceController_ServiceChain: end_implement
	return nil
}
