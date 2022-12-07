package main

import (
	"advent-of-code-22/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name           string
	parentDir      *Directory
	files          []File
	subdirectories []*Directory
}

func (dir *Directory) add_subdir(name string) {
	sub := Directory{name, dir, []File{}, []*Directory{}}
	dir.subdirectories = append(dir.subdirectories, &sub)
}

func (dir *Directory) add_file(name string, size int) {
	file := File{name, size}
	dir.files = append(dir.files, file)
}

func (dir Directory) total_size() int {
	var totalSize int
	for _, file := range dir.files {
		totalSize += file.size
	}
	for _, subdir := range dir.subdirectories {
		totalSize += subdir.total_size()
	}

	return totalSize
}

func (dir Directory) directories_smaller_than(n int) []Directory {
	var outputDirs []Directory
	if dir.total_size() <= n {
		outputDirs = append(outputDirs, dir)
	}
	for _, subdir := range dir.subdirectories {
		outputDirs = append(outputDirs, subdir.directories_smaller_than(n)...)
	}

	return outputDirs
}

func (dir Directory) directories_greater_than(n int) []Directory {
	var outputDirs []Directory
	if dir.total_size() >= n {
		outputDirs = append(outputDirs, dir)
	}
	for _, subdir := range dir.subdirectories {
		outputDirs = append(outputDirs, subdir.directories_greater_than(n)...)
	}

	return outputDirs
}

func createDirectoryTree(lines []string) Directory {
	rootDir := Directory{"/", nil, []File{}, []*Directory{}}
	currentDir := &rootDir
	for i := 0; i < len(lines); i++ {
		if string(lines[i][0]) == "$" {
			if lines[i][2:4] == "cd" {
				dirName := lines[i][5:]
				if dirName == ".." {
					// go up one
					currentDir = currentDir.parentDir
				} else {
					subdirs := currentDir.subdirectories
					for _, subdir := range subdirs {
						if subdir.name == dirName {
							currentDir = subdir
							break
						}
					}
					if currentDir.name != dirName {
						panic(dirName)
					}
				}
			}
		} else if lines[i][:3] == "dir" {
			dirName := lines[i][4:]
			currentDir.add_subdir(dirName)
		} else {
			re, _ := regexp.Compile(`([0-9]+) ([a-z\.]+)`)
			matches := re.FindAllStringSubmatch(lines[i], -1)
			sizeString, name := matches[0][1], matches[0][2]
			size, _ := strconv.Atoi(sizeString)
			currentDir.add_file(name, size)
		}
	}

	return rootDir
}

func A() {
	lines, _ := utils.ReadLines("input.txt")
	rootDir := createDirectoryTree(lines)

	outputDirs := rootDir.directories_smaller_than(100000)
	var totalSize int
	for _, dir := range outputDirs {
		totalSize += dir.total_size()
	}

	fmt.Println(totalSize)
}

func B() {
	lines, _ := utils.ReadLines("input.txt")
	rootDir := createDirectoryTree(lines)

	usedSpace := rootDir.total_size()
	unusedSpace := 70_000_000 - usedSpace
	extraSpaceNeeded := 30_000_000 - unusedSpace

	outputDirs := rootDir.directories_greater_than(extraSpaceNeeded)
	var dirSizes []int
	for _, dir := range outputDirs {
		dirSizes = append(dirSizes, dir.total_size())
	}
	sort.Ints(dirSizes)
	fmt.Println(dirSizes[2])

}

func main() {
	A()
	B()
}
