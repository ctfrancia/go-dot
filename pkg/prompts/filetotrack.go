package prompts

import (
// "github.com/ctfrancia/go-dot/pkg/models"
)

// FileToTrack defines the
type FileToTrack struct {
	Manager string
	MPath   string
	MType   string
	RunCmd  string
}

// New returns a new instance of FileToTrack
func New() *FileToTrack {
	return &FileToTrack{}
}

// NameOfFileToTrack get's the file name to track eg: .zshrc, .bashrc, init.vim
func (ftt *FileToTrack) NameOfFileToTrack(fname string) error {
	return nil
}

// PathOfFileToTrack takes the path of the file
func (ftt *FileToTrack) PathOfFileToTrack(p string) error {
	return nil
}

// FileToTrackType takes the type of the file, Eg: vim, neovim, bash, etc.
// returns an error if that file type is already being tracked
func (ftt *FileToTrack) FileToTrackType(t string) error {
	return nil
}

// AddRunCmd adds to the
func (ftt *FileToTrack) AddRunCmd(cmd string) error {
	return nil
}
