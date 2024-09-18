package learning

import (
	"fmt"
	"os"
	"path/filepath"
)

// Tree is the function to print the tree structure of the current directory
func Tree(fp string) {
	filepath.Walk(fp, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if path == fp {
			fmt.Println(path)
			return nil
		}
		fmt.Println(path)
		return err
	})
}
