package mod

import (
	"fmt"
	"os"
)

// Root contains the absolute path of the module root. Its value is empty if
// used outside of a module.
var Root string

func init() {
	/*
		stdout, _ := exec.Command("go", "env", "GOMOD").Output()

		var ok bool
		Root, ok = strings.CutSuffix(strings.TrimSpace(string(stdout)), "/go.mod")
		if !ok {
			Root = ""
		}
		fmt.Println("Root: ", Root)
	*/

	// download solc to app's workdir 
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	Root = wd
	fmt.Println("current work directory: ", wd)
}
