// Inspiration from https://www.reddit.com/r/golang/comments/mtvm96/install_helm_chart_using_go_programmatically/, thank you

package main

import (
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
)

func main() {
	dryRun := false
	chartPath := "./hello-world-0.5.0.tgz"
	namespace := "default"
	releaseName := "hello-world"

	settings := cli.New()

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), namespace,
		os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Fatalln(err)
	}

	vals := map[string]interface{}{
		"TEXT":     "Hello World from Helm Chart installed using HELM Go Client!",
		"replicas": 2,
	}

	chart, err := loader.Load(chartPath)
	if err != nil {
		log.Fatalln(err)
	}

	instal := action.NewInstall(actionConfig)
	instal.Namespace = namespace
	instal.ReleaseName = releaseName
	instal.DryRun = dryRun

	rel, err := instal.Run(chart, vals)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Installed Chart from path: %s in namespace: %s\n", rel.Name, rel.Namespace)
}
