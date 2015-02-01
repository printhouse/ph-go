package printhouse

// Import Testing frameworks needed
import (
	"testing"
	"github.com/bmizerany/assert"
)

// Create out api variables for easy access
const (
	clientID     = "your-clientID-here"
	clientSecret = "your-client-secret-here"
)

var ph = New(clientID, clientSecret)

func TestNew(t *testing.T) {

	assert.T(t, ph.clientID == clientID, "clientID should be the same")
	assert.T(t, ph.clientSecret == clientSecret, "clientSecret should be the same")
}

func TestRequest(t *testing.T) {
	result, err := ph.Request("GET", "products/%s", nil, "0001");

	assert.T(t, result != nil, "Shouldnt be null")
	assert.T(t, err != nil, "Shouldnt be null")
}

func TestGet(t *testing.T) {

	result, err := ph.Get("products/%s", nil, "0001")
	assert.T(t, result != nil, "Shouldnt be null")
	assert.T(t, err != nil, "Shouldnt be null")

}

func TestGetEncodedKeys(t *testing.T) {

	result := ph.getEncodedKeys();
	assert.T(t, len(result) > 0, "shouldn't be null")
}

func TestCreateTargetURL(t *testing.T) {

	result := ph.createTargetURL("products");
	assert.T(t, result == "https://api.printhouse.io/v1/products", "shouldn be same URL")

}

