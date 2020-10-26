package models

// GodotDefault returns config structure
func GodotDefault() GdotConfig {
	return GdotConfig{
		RepoURL:      "",
		RepoUsername: "",
		RepoKey:      "",
		Paths:        []Path{},
	}
}

// GDotC is the structure that it is being saved within the json file
type GDotC struct {
	Godot GdotConfig `json:"godot"`
}

// GdotConfig defines the structure of the config
type GdotConfig struct {
	RepoURL      string `json:"repoURL"`
	RepoUsername string `json:"repoUsername"`
	RepoKey      string `json:"repoKey"`
	Paths        []Path `json:"paths"`
}

// Path defines how a Path looks
type Path struct {
	Manager string `json:"manager"`
	MPath   string `json:"mpath"`
	MType   string `json:"mtype"`
}
