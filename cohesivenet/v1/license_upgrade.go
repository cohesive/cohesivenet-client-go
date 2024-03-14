package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const licenseEndpoint = "/license"
const licenseUpgradeEndpoint = "/license/upgrade/"
const licenseParamsEndpoint = "/license/parameters/"

func (c *Client) GetControllerLicense() (ControllerLicense, error) {

	licenseRequest, errNewRequest := http.NewRequest("GET", fmt.Sprintf("%s%s", c.HostURL, licenseEndpoint), nil)
	if errNewRequest != nil {
		return ControllerLicense{}, fmt.Errorf("error creating request: %w", errNewRequest)
	}

	licenseRequestBody, errDoRequest := c.doRequest(licenseRequest)
	if errDoRequest != nil {
		return ControllerLicense{}, fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if licenseRequestBody == nil {
		return ControllerLicense{}, fmt.Errorf("empty response body")
	}

	controllerLicense := ControllerLicense{}
	errUnmarshal := json.Unmarshal([]byte(licenseRequestBody), &controllerLicense)
	if errUnmarshal != nil {
		return ControllerLicense{}, fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}

	log.Printf("Read controller license: %+v", controllerLicense)

	return controllerLicense, nil
}

func (c *Client) CreateLicenseUpgrade(licenseUpgradeKeyPath string, clientpackIps string, managerIps string) error {

	licenseUpgradeKey, errReadfile := os.ReadFile(licenseUpgradeKeyPath)
	if errReadfile != nil {
		return fmt.Errorf("error reading upgarde file: %w", errReadfile)
	}
	licenseRequestResponse, errNewRequest := http.NewRequest("PUT", fmt.Sprintf("%s%s", c.HostURL, licenseUpgradeEndpoint), strings.NewReader(string(licenseUpgradeKey)))
	if errNewRequest != nil {
		return fmt.Errorf("error creating request: %w", errNewRequest)
	}

	licenseRequestBody, errDoRequest := c.doRequestWithTextHeader(licenseRequestResponse)
	if errDoRequest != nil {
		return fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if licenseRequestBody == nil {
		return fmt.Errorf("empty license request body")
	}

	newLicenseUpgrade := LicenseUpgrade{}
	errUnmarshal := json.Unmarshal([]byte(licenseRequestBody), &newLicenseUpgrade)
	if errUnmarshal != nil {
		return fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}

	licenseParams := LicenseParams{
		Clients:  clientpackIps,
		Managers: managerIps,
	}

	requestBody, errMarshal := json.Marshal(licenseParams)
	if errMarshal != nil {
		return fmt.Errorf("error marshaling license parameters: %w", errMarshal)
	}

	paramsRequestResponse, errNewRequest := http.NewRequest("PUT", fmt.Sprintf("%s%s", c.HostURL, licenseParamsEndpoint), strings.NewReader(string(requestBody)))
	if errNewRequest != nil {
		return fmt.Errorf("error creating request: %w", errNewRequest)
	}

	paramsRequestBody, errDoRequest := c.doRequest(paramsRequestResponse)
	if errDoRequest != nil {
		return fmt.Errorf("error doing request: %w", errDoRequest)
	}

	log.Println("License upgrade response: ", string(paramsRequestBody))
	if paramsRequestBody == nil {
		return fmt.Errorf("empty params request body")
	}

	log.Printf("Created new upgrade: %+v", newLicenseUpgrade)

	return nil
}

func (c *Client) UpdateLicenseUpgrade(licenseUpgradeKeyPath string, clientpackIps string, managerIps string) error {

	licenseUpgradeKey, errReadfile := os.ReadFile(licenseUpgradeKeyPath)
	if errReadfile != nil {
		return fmt.Errorf("error reading upgarde file: %w", errReadfile)
	}
	licenseRequestResponse, errNewRequest := http.NewRequest("PUT", fmt.Sprintf("%s%s", c.HostURL, licenseUpgradeEndpoint), strings.NewReader(string(licenseUpgradeKey)))
	if errNewRequest != nil {
		return fmt.Errorf("error creating request: %w", errNewRequest)
	}

	licenseRequestBody, errDoRequest := c.doRequestWithTextHeader(licenseRequestResponse)
	if errDoRequest != nil {
		return fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if licenseRequestBody == nil {
		return fmt.Errorf("empty license request body")
	}

	newLicenseUpgrade := LicenseUpgrade{}
	errUnmarshal := json.Unmarshal([]byte(licenseRequestBody), &newLicenseUpgrade)
	if errUnmarshal != nil {
		return fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}

	licenseParams := LicenseParams{
		Clients:  clientpackIps,
		Managers: managerIps,
	}

	requestBody, errMarshal := json.Marshal(licenseParams)
	if errMarshal != nil {
		return fmt.Errorf("error marshaling license parameters: %w", errMarshal)
	}

	paramsRequestResponse, errNewRequest := http.NewRequest("PUT", fmt.Sprintf("%s%s", c.HostURL, licenseParamsEndpoint), strings.NewReader(string(requestBody)))
	if errNewRequest != nil {
		return fmt.Errorf("error creating request: %w", errNewRequest)
	}

	paramsRequestBody, errDoRequest := c.doRequest(paramsRequestResponse)
	if errDoRequest != nil {
		return fmt.Errorf("error doing request: %w", errDoRequest)
	}

	log.Println("License upgrade response: ", string(paramsRequestBody))
	if paramsRequestBody == nil {
		return fmt.Errorf("empty params request body")
	}

	log.Printf("Created new upgrade: %+v", newLicenseUpgrade)

	return nil
}
