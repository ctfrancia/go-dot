package prompts

import (
	"github.com/ctfrancia/go-dot/pkg/models"
)

// New returns a new instance of File
func New() *models.File {
	return &models.File{
		Manager: "",
		MPath:   "",
		MType:   "",
		RunCmd:  "",
	}
}

// NameOfFile get's the file name to track eg: .zshrc, .bashrc, init.vim
func (m *models.File) NameOfFile(fname string) error {
	return nil
}

// PathOfFile takes the path of the file
func (m *models.File) PathOfFile(p string) error {
	return nil
}

// FileType takes the type of the file, Eg: vim, neovim, bash, etc.
// returns an error if that file type is already being tracked
func (m *models.File) FileType(t string) error {
	return nil
}

// AddRunCmd
func (m *models.File) AddRunCmd(cmd string) error {
}
