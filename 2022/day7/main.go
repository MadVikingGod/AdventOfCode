package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var testInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func main() {
	root := parseDirs(testInput)
	fmt.Println(root.findMaxSize(100000))
	unused := 70000000 - root.size()
	target := 30000000 - unused
	fmt.Println(root.findMinDir(target))

	root = parseDirs(input)
	fmt.Println(root.findMaxSize(100000))
	unused = 70000000 - root.size()
	target = 30000000 - unused
	fmt.Println(root.findMinDir(target))
	// 34474241 is too high
}

type dir struct {
	parent *dir
	dirs   map[string]*dir
	files  map[string]int
}

func (d *dir) size() int {
	s := 0
	for _, f := range d.files {
		s += f
	}
	for _, d := range d.dirs {
		s += d.size()
	}
	return s
}

func (d *dir) findMaxSize(max int) int {
	size := 0
	if d.size() <= max {
		size = d.size()
	}
	for _, d := range d.dirs {
		size += d.findMaxSize(max)
	}
	return size
}

func (d *dir) findMinDir(min int) (int, error) {
	size := d.size()
	if size < min {
		return 0, fmt.Errorf("Size too small")
	}
	for _, d := range d.dirs {
		s, err := d.findMinDir(min)
		if err == nil && s < size && s > min {
			size = s
		}
	}
	return size, nil
}

func parseDirs(input string) *dir {
	root := &dir{
		dirs:  make(map[string]*dir),
		files: make(map[string]int),
	}
	root.parent = root
	cwd := root

	for _, line := range strings.Split(input, "\n") {
		switch {
		case line == "" || line == "$ ls":
			continue
		case line == "$ cd /":
			cwd = root
		case line == "$ cd ..":
			cwd = cwd.parent
		case strings.HasPrefix(line, "$ cd "):
			name := line[5:]
			if d, ok := cwd.dirs[name]; ok {
				cwd = d
			} else {
				panic("no such dir: " + name)
			}
		case strings.HasPrefix(line, "dir "):
			name := line[4:]
			if _, ok := cwd.dirs[name]; ok {
				panic("dir already exists: " + name)
			}
			cwd.dirs[name] = &dir{
				parent: cwd,
				dirs:   make(map[string]*dir),
				files:  make(map[string]int),
			}
		default:
			var size int
			var name string
			if _, err := fmt.Sscanf(line, "%d %s", &size, &name); err == nil {
				if _, ok := cwd.dirs[name]; ok {
					panic("file already exists: " + name)
				}
				cwd.files[name] = size
			}
		}
	}
	return root
}
