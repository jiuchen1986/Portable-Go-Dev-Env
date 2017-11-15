// The executor for requests' processing using a new goroutine in the handlers

package handler

import (
    "fmt"
    "context"
    
    "github.com/goadesign/goa"
)

type executor struct {
    c