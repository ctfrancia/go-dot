package pathmanager

func GodotDefault() GdotConfig {
	return GdotConfig{
		RepoURL:      "na",
		RepoUsername: "na",
		Paths:        []Path{},
	}
}

type GdotConfig struct {
	RepoURL      string
	RepoUsername string
	Paths        []Path
}

type Path struct {
	Mtype   string
	Manager string
	MPath   string
}
