package middlewares

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPIPath(t *testing.T) {
	path := "/v1/drivers/123/order/GK-12312/completion"
	expectedPath := "/v1/drivers/*/order/*/completion"
	newRelicCustomPath := GetAPIPath(path)
	assert.Equal(t, expectedPath, newRelicCustomPath)
}

func TestGetAPIPathIfPathIsEmpty(t *testing.T) {
	path := ""
	expectedPath := ""
	newRelicCustomPath := GetAPIPath(path)
	assert.Equal(t, expectedPath, newRelicCustomPath)
}

func TestGetAPIPathIfPathDoesNotContainMultipleDigit(t *testing.T) {
	path := "/v1/driver/status"
	expectedPath := "/v1/driver/status"
	newRelicCustomPath := GetAPIPath(path)
	assert.Equal(t, expectedPath, newRelicCustomPath)
}
