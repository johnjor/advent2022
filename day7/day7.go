package main

import (
	"fmt"
	"os"
)

type Node struct {
	Name  string
	Size  uint64
	IsDir bool
	Links map[string]*Node
}

type Tree struct {
	Root    *Node
	Current *Node
}

func NewFileNode(name string, size uint64) *Node {
	return &Node{Name: name, Size: size, Links: make(map[string]*Node)}
}

func NewDirNode(name string) *Node {
	return &Node{Name: name, IsDir: true, Links: make(map[string]*Node)}
}

func (parent *Node) append(children ...*Node) {
	for _, child := range children {
		parent.Links[child.Name] = child
		child.Links[".."] = parent
	}
}

func (tree *Tree) ChangeDirectory(name string) {
	if name == "/" {
		tree.Current = tree.Root
		return
	}

	node, present := tree.Current.Links[name]
	if present {
		tree.Current = node
	} else {
		panic(fmt.Sprintf("Directory %s not found in %s", name, tree.Current.Name))
	}

}

func WalkTree(node *Node, acc uint64) (uint64, uint64) {
	var dirSum uint64 = 0
	if node.IsDir {
		for name, child := range node.Links {
			if name == ".." {
				continue
			}
			ret, xacc := WalkTree(child, acc)
			dirSum += ret
			acc = xacc
		}
		node.Size = dirSum
		if dirSum <= 100000 {
			acc += dirSum
		}
	} else {
		dirSum += node.Size
	}
	return dirSum, acc
}

func WalkTreePart2(node *Node, threshold uint64) {
	if node.IsDir {
		for name, child := range node.Links {
			if name == ".." {
				continue
			}
			WalkTreePart2(child, threshold)
		}
		if node.Size > threshold {
			fmt.Println(node.Size)
		}
	}
}

func LoadTestTree(root *Node) *Node {
	a := NewFileNode("foobar.txt", 10)

	b := NewDirNode("images")
	c := NewFileNode("img.jpg", 10)
	d := NewFileNode("img2.jpg", 10)
	b.append(c, d)

	e := NewDirNode("css")
	f := NewFileNode("style.css", 11)
	g := NewFileNode("style2.css", 11)
	e.append(f, g)

	root.append(a, b, e)
	return root
}

func LoadTestTree2(root *Node) *Node {
	// - / (dir)
	//  - a (dir)
	//    - e (dir)
	//      - i (file, size=584)
	//    - f (file, size=29116)
	//    - g (file, size=2557)
	//    - h.lst (file, size=62596)
	//  - b.txt (file, size=14848514)
	//  - c.dat (file, size=8504156)
	//  - d (dir)
	//    - j (file, size=4060174)
	//    - d.log (file, size=8033020)
	//    - d.ext (file, size=5626152)
	//    - k (file, size=7214296)
	a := NewDirNode("a")
	e := NewDirNode("e")
	i := NewFileNode("i", 584)
	e.append(i)

	f := NewFileNode("f", 29116)
	g := NewFileNode("g", 2557)
	h := NewFileNode("h.lst", 62596)
	a.append(e, f, g, h)

	b := NewFileNode("b.txt", 14848514)
	c := NewFileNode("c.dat", 8504156)
	root.append(a, b, c)

	d := NewDirNode("d")
	j := NewFileNode("j", 4060174)
	dlog := NewFileNode("d.log", 8033020)
	dext := NewFileNode("d.ext", 5626152)
	k := NewFileNode("k", 7214296)
	d.append(j, dlog, dext, k)
	root.append(d)
	return root
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//scanner := bufio.NewScanner(file)
	root := NewDirNode("/")
	//tree := &Tree{root, root}

	//for scanner.Scan() {
	//	line := strings.TrimSuffix(scanner.Text(), "\n")
	//	parts := strings.Split(line, " ")
	//
	//	if parts[0] == "$" {
	//		if parts[1] == "cd" {
	//			tree.ChangeDirectory(parts[2])
	//		} else if parts[1] == "ls" {
	//			continue
	//		} else {
	//			fmt.Printf("WTF? %s\n", line)
	//			continue
	//		}
	//
	//	} else if parts[0] == "dir" {
	//		tree.Current.append(NewDirNode(parts[1]))
	//	} else {
	//		size, err := strconv.ParseUint(parts[0], 10, 64)
	//		if err != nil {
	//			panic(err)
	//		}
	//		tree.Current.append(NewFileNode(parts[1], size))
	//	}
	//}

	// Part 1
	//_, acc := WalkTree(root, 0)
	//fmt.Println(acc)

	// Part 2
	//var totalSize uint64 = 70_000_000
	//var requiredFree uint64 = 30_000_000
	//usedSpace, _ := WalkTree(root, 0)
	//spaceNeeded := usedSpace - requiredFree
	//
	//fmt.Printf("Space needed: %d\n", spaceNeeded)
	//WalkTreePart2(root, spaceNeeded)

	LoadTestTree2(root)
	var totalSpace uint64 = 70_000_000
	usedSpace, _ := WalkTree(root, 0)
	freeSpace := totalSpace - usedSpace
	threshold := 30_000_000 - freeSpace
	fmt.Printf("Used Space: %d, Free Space: %d, Needed Space: %d\n", usedSpace, freeSpace, threshold)

	WalkTreePart2(root, threshold)

}
