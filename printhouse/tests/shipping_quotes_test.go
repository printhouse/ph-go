package tests

import (
	"testing"
	"github.com/bmizerany/assert"
	. "github.com/rapito/ph-go/printhouse"

)

var dummyShippingQuote = &ShippingQuote{}


func TestPostShippingQuote(t *testing.T) {
	result := ph.PostShippingQuote(dummyShippingQuote)
	assert.T(t, result == nil, "should be nil")
}

func TestGetShippingQuotse(t *testing.T) {
	result := ph.GetShippingQuotes()
	assert.T(t, result == nil, "should be nil")
}

func TestGetShippingQuote(t *testing.T) {
	result := ph.GetShippingQuote("id")
	assert.T(t, result == nil, "should be nil")
}


