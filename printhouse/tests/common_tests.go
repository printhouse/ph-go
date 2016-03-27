package tests

import (
	. "github.com/rapito/ph-go/printhouse"
)

// Create out api variables for easy access
const (
	clientID     = "your-clientID-here"
	clientSecret = "your-client-secret-here"
)

var ph = New(clientID, clientSecret)

