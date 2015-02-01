package tests

import (
	"testing"
	"github.com/bmizerany/assert"
	. "github.com/rapito/ph-go/printhouse"
)

var dummyOrder = &Order{}

func TestGetOrders(t *testing.T) {
	result := ph.GetOrders()
	assert.T(t, result == nil, "should be nil")
}

func TestGetOrder(t *testing.T) {
	result := ph.GetOrder("id")
	assert.T(t, result == nil, "should be nil")
}

func TestPostOrder(t *testing.T) {
	result := ph.PostOrder(dummyOrder)
	assert.T(t, result == nil, "should be nil")
}

func TestPutOrder(t *testing.T) {
	result, err := ph.PutOrder(dummyOrder)
	assert.T(t, result == nil, "should be nil")
	assert.T(t, err == nil, "should be nil")
}

func TestConfirmOrder(t *testing.T) {
	result := ph.ConfirmOrder(dummyOrder)
	assert.T(t, result == false, "should be false")
}

