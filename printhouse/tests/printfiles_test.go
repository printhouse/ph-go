package tests

import (
	"testing"
	"github.com/bmizerany/assert"
	. "github.com/rapito/ph-go/printhouse"

)

var dummyPrintFile = &PrintFile{}

func TestGetPrintFiles(t *testing.T) {
	result := ph.GetPrintFiles()
	assert.T(t, result == nil, "should be nil")
}

func TestGetPrintFile(t *testing.T) {
	result := ph.GetPrintFile("id")
	assert.T(t, result == nil, "should be nil")
}

func TestPostPrintFile(t *testing.T) {
	result := ph.PostPrintFile(dummyPrintFile)
	assert.T(t, result == nil, "should be nil")
}



