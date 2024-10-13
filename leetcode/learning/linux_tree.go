package learning

import (
	"fmt"
	"os"
)

// Tree is the function to print the tree structure of the current directory
func Tree(root string, depth int) error {
	fileInfo, err := os.Stat(root)

	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		if depth != 0 {
			printSpace(depth)
		}
		fmt.Printf("%s%s\n", "|--", fileInfo.Name())
	} else {
		if depth != 0 {
			printSpace(depth)
		}
		fmt.Printf("%s%s\n", "|--", fileInfo.Name())
		file, err := os.Open(root)
		if err != nil {
			return err
		}
		defer file.Close()

		names, err := file.Readdirnames(-1)
		if err != nil {
			return err
		}

		for _, name := range names {
			Tree(root+"/"+name, depth+1)
		}
	}

	return nil
}

func printSpace(depth int) {
	fmt.Print("|")
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
}
