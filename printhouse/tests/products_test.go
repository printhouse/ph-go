package tests

import (
	"testing"
	"github.com/bmizerany/assert"

)

func TestGetProducts(t *testing.T) {
	result := ph.GetProducts("")
	assert.T(t, result == nil, "should be nil")
}

func TestGetProduct(t *testing.T) {
	result := ph.GetProduct("id")
	assert.T(t, result == nil, "should be nil")
}





