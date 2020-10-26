package models

// GoDotConfig is what is used inside of the program
type GoDotConfig struct {
	RepoURL      string
	RepoUsername string
	RepoToken    string
	Files        []File
}

// File represents the individual file looks like that is being tracked
type File struct {
	Manager string
	MPath   string
	MType   string
	RunCmd  string
}

// GoDotConfigJSON is the structure that it is being saved within the json file
type GoDotConfigJSON struct {
	Godot GdotConfig `json:"godot"`
}

// GdotConfig defines the structure of the config
type GdotConfig struct {
	RepoURL      string     `json:"repoURL"`
	RepoUsername string     `json:"repoUsername"`
	RepoKey      string     `json:"repoKey"`
	Files        []FileJSON `json:"files"`
}

// FileJSON defines how a File looks when interacting with the json
type FileJSON struct {
	Manager string `json:"manager"`
	MPath   string `json:"mpath"`
	MType   string `json:"mtype"`
	RunCmd  string `json:"runCmd"`
}
