package newRelic

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"log"
	"os"
)

type NewRelic struct {
	App newrelic.Application
}

func InitializeNewRelic() *NewRelic {
	goLicenseKey := os.Getenv("NEW_RELIC_GO_LICENSE_KEY")
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("badge-service"),
		newrelic.ConfigLicense(goLicenseKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	return &NewRelic{
		App: *app,
	}
}
