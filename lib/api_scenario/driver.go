package api_scenario

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wargasipil/tiktok_rpc/lib/driver_handler"
)

func UseDefaultDriver(t *testing.T) *driver_handler.DriverAccount {
	driver, err := driver_handler.NewDriverAccount("test@gmail.com", "", "", "")

	assert.Nil(t, err)

	return driver
}
