package tests

import (
	"testing"
	"github.com/bmizerany/assert"
	. "github.com/rapito/ph-go/printhouse"

)

var dummyOrder = &Order{}

func TestGet(t *testing.T) {
	result, err := ph.Get("products/%s", nil, "0001")
	assert.T(t, result != nil, "Shouldnt be null")
	assert.T(t, err != nil, "Shouldnt be null")
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

