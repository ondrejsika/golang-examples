package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Unleash/unleash-client-go/v3"
)

type metricsInterface struct{}

func init() {
	gitlabDomain := os.Getenv("GITLAB_DOMAIN")
	projectId := os.Getenv("PROJECT_ID")
	instanceId := os.Getenv("INSTANCE_ID")
	appName := os.Getenv("APP_NAME")

	unleash.Initialize(
		unleash.WithUrl("https://"+gitlabDomain+"/api/v4/feature_flags/unleash/"+projectId),
		unleash.WithInstanceId(instanceId),
		unleash.WithAppName(appName),
		unleash.WithListener(&metricsInterface{}),
	)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if unleash.IsEnabled("czech") {
			io.WriteString(w, "Ahoj Svete!\n")
		} else if unleash.IsEnabled("german") {
			io.WriteString(w, "Hallo Welt!\n")
		} else {
			io.WriteString(w, "Hello World!\n")
		}
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
