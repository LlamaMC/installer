package main

import (
	"fmt"
	"installer/json"
	"installer/util"

	"github.com/charmbracelet/huh"
)

var (
	selectedVersion       string
	selectedConfiguration string
)

func main() {
	versionsObject := json.ParseJsonFromUrl("https://raw.githubusercontent.com/LlamaMC/manifest/refs/heads/master/versions.json")
	manifestVersions := versionsObject["versions"].([]any)
	var parsedVersions []util.Version
	for _, version := range manifestVersions {
		parsedVersions = append(parsedVersions, util.Version{
			Id:    version.(map[string]any)["id"].(string),
			Build: int(version.(map[string]any)["build"].(float64)),
			Name:  version.(map[string]any)["name"].(string)})
	}
	util.InitializeVersions(parsedVersions)
	configurationObject := json.ParseJsonFromUrl("https://raw.githubusercontent.com/LlamaMC/manifest/refs/heads/master/configurations.json")
	manifestConfigurations := configurationObject["configurations"].([]any)
	var parsedConfigurations []util.Configuration
	for _, configuration := range manifestConfigurations {
		parsedConfigurations = append(parsedConfigurations, util.Configuration{
			Id:   configuration.(map[string]any)["id"].(string),
			Name: configuration.(map[string]any)["name"].(string),
			Path: configuration.(map[string]any)["path"].(string)})
	}
	util.InitializeConfigurations(parsedConfigurations)
	var versionOptions []huh.Option[string]
	var configurationOptions []huh.Option[string]
	for _, version := range util.Versions {
		versionOptions = append(versionOptions, huh.NewOption(version.Name, version.Id))
	}
	for _, configuration := range util.Configurations {
		configurationOptions = append(configurationOptions, huh.NewOption(configuration.Name, configuration.Id))
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose the Minecraft version").
				Options(versionOptions...).
				Value(&selectedVersion),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose the preconfigured configuration").
				Options(configurationOptions...).
				Value(&selectedConfiguration),
		),
	)
	err := form.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println(selectedVersion, selectedConfiguration)
}
