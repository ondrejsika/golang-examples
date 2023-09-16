package main

import (
	"fmt"
	"log"

	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

func main() {
	repoUrl := "https://helm.sikalabs.io"
	chartName := "hello-world"

	version, err := getLatestVersion(repoUrl, chartName)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(version)
}

func getLatestVersion(repoUrl, chartName string) (string, error) {
	r, err := repo.NewChartRepository(&repo.Entry{
		URL: repoUrl,
	}, getter.All(cli.New()))
	if err != nil {
		return "", err
	}
	// repo.CachePath = "/tmp/helm"

	indexFileName, err := r.DownloadIndexFile()
	if err != nil {
		return "", err
	}

	indexFile, err := repo.LoadIndexFile(indexFileName)
	if err != nil {
		return "", err
	}

	c, err := indexFile.Get(chartName, "")
	if err != nil {
		return "", err
	}

	return c.Version, nil
}
