package response

import (
	"github.com/transfer360/zatpark_testing/zatpark_errors"
	"strings"
)

type Data struct {
	SearchReference    string    `json:"sref"`
	YourReference      string    `json:"your_reference"`
	AgreementNo        string    `json:"agreement_no"`
	HireFrom           string    `json:"hire_from"`
	HireTo             string    `json:"hire_end"`
	ContactInformation []Contact `json:"contact_information"`
}

type Contact struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	CompanyName  string `json:"company_name,omitempty"`
	Telephone    string `json:"telephone"`
	Email        string `json:"email"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	AddressLine3 string `json:"address_line3"`
	City         string `json:"city,omitempty"`
	Postcode     string `json:"postcode"`
	Country      string `json:"country"`
	Priority     bool   `json:"priority"`
}

func LeaseResponse(sref string, reference string, vrm string) (Data, error) {
	vrm = strings.ToUpper(strings.ReplaceAll(vrm, " ", ""))

	if vrm == "ZP01TST" {
		return testResultOne(sref, reference), nil
	} else if vrm == "ZP02TST" {
		return testResultTwo(sref, reference), nil
	} else {
		return Data{}, zatpark_errors.ErrNotATestVehicle
	}

}

func testResultOne(sref string, reference string) Data {

	testResponse := Data{
		SearchReference: sref,
		YourReference:   reference,
		AgreementNo:     "89000001000",
		HireFrom:        "2023-07-01",
		HireTo:          "2025-07-31",
	}
	testResponse.ContactInformation = make([]Contact, 0)

	cnt := Contact{
		FirstName:    "Michelle",
		LastName:     "Donovan",
		CompanyName:  "",
		Telephone:    "01483799246",
		Email:        "michelled@gmail.com",
		AddressLine1: "77 Connaught Crescent",
		AddressLine2: "Brookwood",
		AddressLine3: "",
		City:         "Woking",
		Postcode:     "GU240AW",
		Country:      "United Kingdom",
		Priority:     true,
	}
	testResponse.ContactInformation = append(testResponse.ContactInformation, cnt)

	return testResponse
}

// -------------------------------------------------------------------
func testResultTwo(sref string, reference string) Data {

	testResponse := Data{
		SearchReference: sref,
		YourReference:   reference,
		AgreementNo:     "8900000000",
		HireFrom:        "2023-05-01",
		HireTo:          "2025-08-31",
	}
	testResponse.ContactInformation = make([]Contact, 0)

	cnt := Contact{
		FirstName:    "Jason",
		LastName:     "Williams",
		CompanyName:  "",
		Telephone:    "012730010002",
		Email:        "jwilliams@outlook.com",
		AddressLine1: "Regent House",
		AddressLine2: "Hove Street",
		AddressLine3: "",
		City:         "Hove",
		Postcode:     "BN32DW",
		Country:      "United Kingdom",
		Priority:     true,
	}
	testResponse.ContactInformation = append(testResponse.ContactInformation, cnt)

	return testResponse
}

// -------------------------------------------------------------------
