package main

import (
	"fmt"
	"github.com/skynetservices/skydns/client"
)

var (
	domain      string
	environment string
	skydnsUrl   string
	secret      string
	appname     string

	skydns      Skydns
	running     = make(map[string]struct{})
	runningLock = sync.Mutex{}
)

func init() {
	flag.StringVar(&skydnsUrl, "skydns", "", "url to the skydns url")
	flag.StringVar(&secret, "secret", "", "skydns secret")
	flag.StringVar(&domain, "domain", "", "same domain passed to skydns")
	flag.StringVar(&environment, "environment", "dev", "environment name where service is running")
	flag.StringVar(&appname, "appname", "web", "service name/appname to load balance - ?.dev.docker")
	flag.Parse()
}

func validateSettings() {

	if skydnsUrl == "" {
		skydnsUrl = "http://" + os.Getenv("SKYDNS_PORT_8080_TCP_ADDR") + ":8080"
	}

	if domain == "" {
		fatal(fmt.Errorf("Must specify your skydns domain"))
	}

	if appname == "" {
		fatal(fmt.Errorf("Must specify the application name to be proxied"))
	}
}

func main() {

	// setup things
	// first a connection to SkyDNS
	if skydns, err = client.NewClient(skydnsUrl, secret, domain, "172.17.42.1:53"); err != nil {
		log.Logf(log.FATAL, "error connecting to skydns: %s", err)
		fatal(err)
	}

	// connect to skydns & get a list of the services that
	// match config - in a goroutine

	// wait until everything closes

}
