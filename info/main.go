package info

type Info struct {
	App     string `json:"app"`
	Version string `json:"version"`
	Tag     string `json:"tag"`
	Build   string `json:"build"`
}

var App = Info{
	App: "" +
		"service-scaffold",
}
