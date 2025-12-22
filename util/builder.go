package util

var Versions []Version
var Configurations []Configuration

type Version struct {
	Id    string `json:"id"`
	Build int    `json:"build"`
	Name  string `json:"name"`
}

type Configuration struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func InitializeVersions(versions []Version) {
	Versions = versions
}

func InitializeConfigurations(configurations []Configuration) {
	Configurations = configurations
}
