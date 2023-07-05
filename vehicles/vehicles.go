package vehicles

import "strings"

func IsTestVehicle(vrm string) bool {

	vrm = strings.ToUpper(strings.ReplaceAll(" ", "", vrm))

	switch vrm {

	case "ZP01TST":
	case "ZP02TST":
	case "ZP99TST":
		return true

	default:
		return false

	}
	return false
}
