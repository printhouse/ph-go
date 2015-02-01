package tests

import (
	"testing"
	"github.com/bmizerany/assert"
	. "github.com/rapito/ph-go/printhouse"

)

var dummyShippingQuote = &ShippingQuote{}

func TestGetShippingQuote(t *testing.T) {
	result := ph.GetShippingQuote("id")
	assert.T(t, result == nil, "should be nil")
}

func TestPostShippingQuote(t *testing.T) {
	result := ph.PostShippingQuote(dummyShippingQuote)
	assert.T(t, result == nil, "should be nil")
}


