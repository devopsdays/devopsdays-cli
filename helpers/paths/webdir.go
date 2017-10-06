package paths

import (
	"fmt"
	"os"
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
		return pwd
	} else {
		s := os.Getenv("DODPATH")
		s = strings.TrimSuffix(s, "/")
		s = strings.TrimSuffix(s, "\\")
		return s
	}
}
