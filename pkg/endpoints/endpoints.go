package endpoints

import "fmt"

const (
	V20 = "https://api.scanii.com/v2.0"
	V21 = "https://api.scanii.com/v2.1"

	// US
	V20_US1 = "https://api-us1.scanii.com/v2.0"
	V21_US1 = "https://api-us1.scanii.com/v2.1"

	// EU
	V20_EU1 = "https://api-eu1.scanii.com/v2.0"
	V20_EU2 = "https://api-eu2.scanii.com/v2.0"
	V21_EU1 = "https://api-eu1.scanii.com/v2.1"
	V21_EU2 = "https://api-eu2.scanii.com/v2.1"

	// AP
	V20_AP1 = "https://api-ap1.scanii.com/v2.0"
	V20_AP2 = "https://api-ap2.scanii.com/v2.0"
	V21_AP1 = "https://api-ap1.scanii.com/v2.1"
	V21_AP2 = "https://api-ap2.scanii.com/v2.1"
	LATEST  = V21
)

func Resolve(endpoint, path string) string {
	return fmt.Sprintf("%s/%s", endpoint, path)
}
