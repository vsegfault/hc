package accessory

import (
	"github.com/brutella/hc/model"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessoryIdentifyChanged(t *testing.T) {
	a := New(info)

	var identifyCalled = 0
	a.OnIdentify(func() {
		identifyCalled += 1
	})

	a.Info.Identify.SetValueFromRemote(true)
	// Identify is set to false immediately
	assert.False(t, a.Info.Identify.Identify())
	assert.Equal(t, identifyCalled, 1)
}

func TestAccessoryInfo(t *testing.T) {
	var accessoryInfo = model.Info{
		Name:         "My Accessory",
		SerialNumber: "0009",
		Manufacturer: "Matthias",
		Model:        "1A",
		Firmware:     "0.1",
		Hardware:     "1.0",
		Software:     "2.1",
	}

	a := New(accessoryInfo)
	assert.Equal(t, a.Name(), "My Accessory")
	assert.Equal(t, a.SerialNumber(), "0009")
	assert.Equal(t, a.Manufacturer(), "Matthias")
	assert.Equal(t, a.Model(), "1A")
	assert.Equal(t, a.Firmware(), "0.1")
	assert.Equal(t, a.Hardware(), "1.0")
	assert.Equal(t, a.Software(), "2.1")
}
