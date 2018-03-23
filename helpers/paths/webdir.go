package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetWebdir returns the value of the web directory
func GetWebdir() string {
	if os.Getenv("DODPATH") == "" {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if validateWebDir(pwd) != true {
			return pwd
		}
		fmt.Println("devopsdays-web directory is invalid")
		os.Exit(1)
	}
	s := os.Getenv("DODPATH")
	s = strings.TrimSuffix(s, "/")
	s = strings.TrimSuffix(s, "\\")
	if validateWebDir(s) != true {
		fmt.Println("devopsdays directory is invalid")
		os.Exit(1)
	}
	return s
}

func validateWebDir(webdir string) bool {

	themeFile := filepath.Join(webdir, "themes", "devopsdays-theme", "theme.toml")
	if _, err := os.Stat(themeFile); os.IsNotExist(err) {
		return false
	}
	return true
}
