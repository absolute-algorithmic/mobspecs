package mobspec

import (
	"fmt"
	"math/rand"
	"time"
)

// Variables
var initialized bool

// MobileSpecs is a struct that represents a mobile specification
type MobileSpecs struct {
	BuildID                 string
	BuildDisplayID          string
	ProductName             string
	ProductDevice           string
	ProductBoard            string
	ProductManufacturer     string
	ProductBrand            string
	ProductModel            string
	Bootloader              string
	Hardware                string
	BuildType               string
	BuildTags               string
	BuildFingerprint        string
	BuildUser               string
	BuildHost               string
	BuildVersionIncremental string
	BuildVersionRelease     string
	BuildVersionSDK         string
	BuildVersionCodename    string
	ScreenHeight            string
	ScreenWidth             string
	ScreenDensity           string
	ScreenXDPI              string
	ScreenYDPI              string
	ScreenDPI               string
	ScreenScaledDensity     string
}

// InitializeSpecs initializes the mobile specs
func InitializeSpecs() error {
	rand.Seed(time.Now().UnixNano())

	// Shuffle the devices
	rand.Shuffle(len(MobileSpecsData), func(i, j int) {
		MobileSpecsData[i], MobileSpecsData[j] = MobileSpecsData[j], MobileSpecsData[i]
	})

	initialized = true
	return nil
}

// GetMobileSpecs returns the mobile specs
func GetMobileSpec() (MobileSpecs, error) {
	if !initialized {
		if err := InitializeSpecs(); err != nil {
			return MobileSpecs{}, err
		}
	}

	if len(MobileSpecsData) == 0 {
		return MobileSpecs{}, fmt.Errorf("no more mobile specs available")
	}

	spec := MobileSpecsData[len(MobileSpecsData)-1]
	MobileSpecsData = MobileSpecsData[0 : len(MobileSpecsData)-1]
	return spec, nil
}
