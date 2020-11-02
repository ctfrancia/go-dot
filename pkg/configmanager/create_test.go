package pathmanager

import (
	cm "github.com/ctfrancia/go-dot/pkg/configmanager"
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

func TestCreateGodotDirPath(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Errorf("Error while getting the current user: %V", err)
	}

	tp := filepath.Join(usr.HomeDir, ".config", "godot", "testDir")

	err = cm.CreateGodotDirPath(tp)
	if err != nil {
		t.Errorf("creating directory error: %v", err)
	}

	if _, err = os.Stat(tp); os.IsNotExist(err) {
		t.Errorf("error path does not exist: %v", err)
	}
}
