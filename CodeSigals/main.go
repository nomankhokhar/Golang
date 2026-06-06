package main

import (
	"fmt"
	"strings"
)

type Directory struct {
	Subdirectories map[string]*Directory
	Files          map[string]string
}

type FileSystem struct {
	root *Directory
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		root: &Directory{
			Subdirectories: make(map[string]*Directory),
			Files:          make(map[string]string),
		},
	}
}

// traverse goes to parent folder
// Example:
// /dirA/file.txt
// returns:
// parent = dirA
// name = file.txt
func (fs *FileSystem) traverse(path string) (*Directory, string) {

	// Remove "/" from start and end
	path = strings.Trim(path, "/") // this will remove leading and trailing slashes like "/dirA/dirB/" -> "dirA/dirB"
	// fmt.Println("path:", path)

	parts := strings.Split(path, "/") // this will split the path into parts like "dirA/dirB/file.txt" -> ["dirA", "dirB", "file.txt"]
	// fmt.Println("parts:", parts)

	current := fs.root // this will start from the root directory like "/"
	// fmt.Println("current:", current)

	// Move through folders
	for i := 0; i < len(parts)-1; i++ {

		next, exists := current.Subdirectories[parts[i]]
		// fmt.Println("next:", next, "exists:", exists, "current.Subdirectories:", current.Subdirectories)
		if !exists {
			return nil, ""
		}

		current = next
	}
	// fmt.Println("current return:", *current, "name return:", parts[len(parts)-1])
	return current, parts[len(parts)-1]
}

// AddDirectory adds new folder
func (fs *FileSystem) AddDirectory(path string) bool {

	parent, dirName := fs.traverse(path) // this will return the parent directory and the name of the new directory to be added like "/dirA/dirB" -> parent = dirA, dirName = dirB
	// Parent folder not found
	if parent == nil {
		return false
	}

	// Folder already exists
	if _, exists := parent.Subdirectories[dirName]; exists {
		return false
	}

	// Create folder
	parent.Subdirectories[dirName] = &Directory{
		Subdirectories: make(map[string]*Directory),
		Files:          make(map[string]string),
	}

	return true
}

// AddFile adds file in folder
func (fs *FileSystem) AddFile(path, content string) bool {
	// this will return the parent directory and the name of the new file to be added like "/dirA/dirB/file.txt" -> parent = dirB, fileName = file.txt

	parent, fileName := fs.traverse(path)
	// Parent folder not found
	if parent == nil {
		return false
	}

	// File already exists
	if _, exists := parent.Files[fileName]; exists {
		return false
	}

	// Save content
	parent.Files[fileName] = content

	return true
}

// GetContent gets file content
func (fs *FileSystem) GetContent(path string) (string, bool) {
	parent, fileName := fs.traverse(path)

	// Parent folder not found
	if parent == nil {
		return "", false
	}

	content, exists := parent.Files[fileName]

	if !exists {
		return "", false
	}

	return content, true
}

// DeleteItem deletes file or folder
func (fs *FileSystem) DeleteItem(path string) bool {

	parent, itemName := fs.traverse(path)

	// Parent folder not found
	if parent == nil {
		return false
	}

	// Path can be file or folder, we need to check both
	// Delete file
	if _, exists := parent.Files[itemName]; exists {
		delete(parent.Files, itemName)
		return true
	}

	// Delete folder
	if _, exists := parent.Subdirectories[itemName]; exists {
		delete(parent.Subdirectories, itemName)
		return true
	}

	return false
}

// printDir recursively prints a directory and its contents
func printDir(dir *Directory, name string, indent string) {
	fmt.Printf("%s[%s]/\n", indent, name)
	for subName, subDir := range dir.Subdirectories {
		printDir(subDir, subName, indent+"  ")
	}
	for fileName, content := range dir.Files {
		fmt.Printf("%s  %s = %q\n", indent, fileName, content)
	}
}

// PrintTree prints the full filesystem tree
func (fs *FileSystem) PrintTree() {
	fmt.Println("--- FileSystem Tree ---")
	printDir(fs.root, "/", "")
	fmt.Println("----------------------")
}

func main() {
	fs := NewFileSystem()

	fs.AddDirectory("/dirA")
	fs.PrintTree()

	fs.AddDirectory("/dirA/dirB")
	fs.PrintTree()

	fs.AddDirectory("/dirA/dirB/dirC")
	fs.PrintTree()

	fs.AddFile("/dirA/dirB/file.txt", "Hello, World!")
	fs.PrintTree()

	content, _ := fs.GetContent("/dirA/dirB/file.txt")
	fmt.Println("GetContent:", content)

	fs.DeleteItem("/dirA/dirB/file.txt")
	fs.PrintTree()

	_, exists := fs.GetContent("/dirA/dirB/file.txt")
	fmt.Println("Exists after delete:", exists)
}
