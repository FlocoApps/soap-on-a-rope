# Soap on a Rope
A Go package designed to make your experiences with SOAP a little lesss unpleasant.

## Features
* SOAP Client support for TLS and basic auth
* Service oriented design

## Example
```
import (
	"fmt"
	"github.com/FlocoApps/soap-on-a-rope"
)

...

reqTemplate := rope.StdEnvelope(AvailabilityRequestTemplate)
	bodyString := fmt.Sprintf(reqTemplate, accessCode)
	url := PreProductionURL + AvailabilityService
	auth := &rope.BasicAuth{
		Username: "UserID",
		Password: "Password",
	}

	soap := &rope.SoapClient{
		UseTLS: true,
		Auth:   auth,
	}

	responseString, err := soap.SendRequest(bodyString, url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(responseString)
	}
```
