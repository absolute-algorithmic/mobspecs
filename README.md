```
go get github.com/absolute-algorithmic/MobSpecs@latest
```


```golang
package main

import (
	"fmt"

	"github.com/absolute-algorithmic/MobSpecs"
)

func main() {
	getSpec, err := mobspec.GetMobileSpec()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Product Build ID:", getSpec.BuildID)
	fmt.Println("Product Display ID:", getSpec.BuildDisplayID)
	fmt.Println("Product Name:", getSpec.ProductName)
	fmt.Println("Product Device:", getSpec.ProductDevice)
	fmt.Println("Product Board:", getSpec.ProductBoard)
	fmt.Println("Product Manufacturer:", getSpec.ProductManufacturer)
	fmt.Println("Product Brand:", getSpec.ProductBrand)
	fmt.Println("Product Model:", getSpec.ProductModel)
	fmt.Println("Bootloader:", getSpec.Bootloader)
	fmt.Println("Hardware:", getSpec.Hardware)
	fmt.Println("Build Type:", getSpec.BuildType)
	fmt.Println("Build Tags:", getSpec.BuildTags)
	fmt.Println("Build Fingerprint:", getSpec.BuildFingerprint)
	fmt.Println("Build User:", getSpec.BuildUser)
	fmt.Println("Build Host:", getSpec.BuildHost)
	fmt.Println("Build Version Incremental:", getSpec.BuildVersionIncremental)
	fmt.Println("Build Version Release:", getSpec.BuildVersionRelease)
	fmt.Println("Build Version SDK:", getSpec.BuildVersionSDK)
	fmt.Println("Build Version Codename:", getSpec.BuildVersionCodename)
	fmt.Println("Screen Height:", getSpec.ScreenHeight)
	fmt.Println("Screen Width:", getSpec.ScreenWidth)
	fmt.Println("Screen Density:", getSpec.ScreenDensity)
	fmt.Println("Screen X DPI:", getSpec.ScreenXDPI)
	fmt.Println("Screen Y DPI:", getSpec.ScreenYDPI)
	fmt.Println("Screen DPI:", getSpec.ScreenDPI)
	fmt.Println("Screen Scaled Density:", getSpec.ScreenScaledDensity)

}
```