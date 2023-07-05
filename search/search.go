package search

import (
	"github.com/transfer360/sys360/vehicle_search"
	"github.com/transfer360/zatpark_testing/zatpark_errors"
	"strings"
)

func Search(sr vehicle_search.SearchRequest) (result vehicle_search.SearchResult, err error) {

	sr.VRM = strings.ToUpper(strings.ReplaceAll(sr.VRM, " ", ""))

	switch sr.VRM {

	case "ZP01TST":
		result.IsLeaseVehicle = true
		result.Response = vehicle_search.LeaseCompany{
			CompanyName:    "Acme Hire Company",
			AddressLine1:   "Regent House",
			AddressLine2:   "Hove Street",
			AddressLine3:   "",
			City:           "Hove",
			PostCode:       "BN3 2DW",
			LeaseCompanyId: 1,
		}
		break

	case "ZP02TST":
		result.IsLeaseVehicle = true
		result.Response = vehicle_search.LeaseCompany{
			CompanyName:    "Lease Fake Company",
			AddressLine1:   "PO Box 5419",
			AddressLine2:   "",
			AddressLine3:   "",
			City:           "Hove",
			PostCode:       "BN52 9AN",
			LeaseCompanyId: 2,
		}
		break

	case "ZP99TST":
		result.IsLeaseVehicle = false
		break

	default:
		return result, zatpark_errors.ErrNotATestVehicle
	}

	return result, nil
}
