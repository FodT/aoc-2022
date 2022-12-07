package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type file struct {
	Size int
	Name string
	Path string
	Dir  bool
}

type dir struct {
	Name      string
	TotalSize int
}

func main() {
	files := map[string]file{}

	scanner := bufio.NewScanner(strings.NewReader(Input))
	var path []string
	for scanner.Scan() {

		line := scanner.Text()
		if line[0] == '$' {
			//fmt.Printf("cmd: %s\n", line)
			if line[2:4] == "cd" {
				dir := line[5:]
				if dir != ".." {
					//fmt.Printf("dir: %s\n", dir)
					path = append(path, dir)
					files[strings.Join(path, "/")] = file{Name: dir, Path: strings.Join(path, "/"), Size: 0, Dir: true}
				} else {
					path = path[:len(path)-1]
				}
			}
		} else if line[0] != 'd' { // assume file info output
			var f file
			fmt.Sscanf(line, "%d %s", &f.Size, &f.Name)
			f.Path = strings.Join(path, "/")

			files[f.Path+"/"+f.Name] = f
			//fmt.Println(f)
		}
	}

	type size struct {
		DirName   string
		TotalSize int
	}
	var sizes []size

	for k, v := range files {
		dirSize := 0
		for k_, v_ := range files {
			if strings.HasPrefix(k_, k) {
				dirSize += v_.Size
			}
		}
		if v.Dir {
			fmt.Printf("dir %s size: %d\n", k, dirSize)
			sizes = append(sizes, size{k, dirSize})
		}
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i].TotalSize < sizes[j].TotalSize
	})

	totalUsableSpace := 70000000
	totalNeeded := 30000000
	totalDiskUsed := sizes[len(sizes)-1].TotalSize
	unused := totalUsableSpace - totalDiskUsed
	needToFree := totalNeeded - unused
	fmt.Printf("totalUsableSpace: %d, totalDiskUsed: %d, unused: %d, needToFree: %d\n",
		totalUsableSpace, totalDiskUsed, unused, needToFree)
	total := 0
	foundDelete := false
	for _, s := range sizes {
		if s.TotalSize <= 100000 {
			total += s.TotalSize
		}
		if s.TotalSize >= needToFree && !foundDelete {
			foundDelete = true
			fmt.Printf("pt2 total: %d\n", s.TotalSize)
		}
	}
	fmt.Printf("pt1 total: %d\n", total)
}
