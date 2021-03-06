// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "TEST REST API": Application Controllers
//
// Command:
// $ goagen
// --design=design
// --out=$(GOPATH)/src
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// TestServiceController is the controller interface for the TestService actions.
type TestServiceController interface {
	goa.Muxer
	LocalService(*LocalServiceTestServiceContext) error
	ServiceChain(*ServiceChainTestServiceContext) error
}

// MountTestServiceController "mounts" a TestService resource controller on the given service.
func MountTestServiceController(service *goa.Service, ctrl TestServiceController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLocalServiceTestServiceContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.LocalService(rctx)
	}
	service.Mux.Handle("GET", "/api/:svcLo/", ctrl.MuxHandler("local_service", h, nil))
	service.LogInfo("mount", "ctrl", "TestService", "action", "LocalService", "route", "GET /api/:svcLo/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewServiceChainTestServiceContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ServiceChain(rctx)
	}
	service.Mux.Handle("GET", "/api/:svcLo/*svcOther", ctrl.MuxHandler("service_chain", h, nil))
	service.LogInfo("mount", "ctrl", "TestService", "action", "ServiceChain", "route", "GET /api/:svcLo/*svcOther")
}
