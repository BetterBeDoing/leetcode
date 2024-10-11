package learning

import (
	"fmt"
	"os"
	"path/filepath"
)

// Tree is the function to print the tree structure of the current directory
func Tree(root string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if path == root {
			fmt.Println(path)
			return nil
		}
		fmt.Println(path)
		return err
	})
}
