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

// Should create a new Printhouse Object.
func TestNew(t *testing.T) {

	assert.T(t, ph.clientID == clientID, "clientID should be the same")
	assert.T(t, ph.clientSecret == clientSecret, "clientSecret should be the same")
}

// Should create a new Printhouse Object.
func TestRequest(t *testing.T) {

	result, err := ph.Request("GET", "albums/%s", nil, "0sNOF9WDwhWunNAHPD3Baj");

	assert.T(t, result != nil, "Shouldnt be null")
	assert.T(t, err != nil, "Shouldnt be null")

}

// Should create a new Printhouse Object.
func TestGet(t *testing.T) {

	result, err := ph.Get("albums/%s", nil, "0sNOF9WDwhWunNAHPD3Baj")
	assert.T(t, result != nil, "Shouldnt be null")
	assert.T(t, err != nil, "Shouldnt be null")

}

// Should create a new Printhouse Object.
func TestGetEncodedKeys(t *testing.T) {

	result := ph.getEncodedKeys();
	assert.T(t, len(result) > 0, "shouldn't be null")

}

// Should create a new Printhouse Object.
func TestCreateTargetURL(t *testing.T) {

	result := ph.createTargetURL("products");
	assert.T(t, result == "https://api.printhouse.io/v1/products", "shouldn be same URL")

}

