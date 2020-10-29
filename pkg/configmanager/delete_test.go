package pathmanager

import (
	"fmt"
	cm "github.com/ctfrancia/go-dot/pkg/configmanager"
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

var (
	fileInfo *os.FileInfo
	e        error
)

func TestDeleteConfig(t *testing.T) {
	err := cm.DeleteConfig()
	if err != nil {
		t.Errorf("error when deleting the folder: %s", err)
	}

	usr, err := user.Current()
	if err != nil {
		t.Errorf("error when getting user current: %s", err)
	}

	fp := filepath.Join(usr.HomeDir, ".config", "godot", "config.json")
	fmt.Println("the path of the config file", fp)

	if _, e := os.Stat(fp); !os.IsNotExist(e) {
		t.Errorf("error path still exists: %v", e)
	}

	fmt.Printf("val after the check of file existing: %v \n", fileInfo)
}
