package models

// GoDotConfig is what is used inside of the program
type GoDotConfig struct {
	RepoURL   string
	RepoToken string
	Files     []File
}

// File represents the individual file looks like that is being tracked
type File struct {
	Manager string
	MPath   string
	MType   string
	RunCmd  string
}

// GoDotConfigJSON is the structure that it is being saved within the yaml file
type GoDotConfigJSON struct {
	Godot GdotConfig `yaml:"godot"`
}

// GdotConfig defines the structure of the config
type GdotConfig struct {
	RepoURL string     `yaml:"repoURL"`
	RepoKey string     `yaml:"repoKey"`
	Files   []FileJSON `yaml:"files"`
}

// FileJSON defines how a File looks when interacting with the yaml
type FileJSON struct {
	Manager string `yaml:"manager"`
	MPath   string `yaml:"mpath"`
	MType   string `yaml:"mtype"`
	RunCmd  string `yaml:"runCmd"`
}
