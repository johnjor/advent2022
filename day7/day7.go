package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func WalkTree(node *Node, sum uint64) uint64 {
	if node.IsDir {
		for name, child := range node.Links {
			if name == ".." {
				continue
			}
			sum = WalkTree(child, sum)
		}
	} else {
		sum += node.Size
	}
	return sum
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

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	root := NewDirNode("/")
	tree := &Tree{root, root}

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		parts := strings.Split(line, " ")

		if parts[0] == "$" {
			if parts[1] == "cd" {
				tree.ChangeDirectory(parts[2])
			} else if parts[1] == "ls" {
				continue
			} else {
				fmt.Printf("WTF? %s\n", line)
				continue
			}

		} else if parts[0] == "dir" {
			tree.Current.append(NewDirNode(parts[1]))
		} else {
			size, err := strconv.ParseUint(parts[0], 10, 64)
			if err != nil {
				panic(err)
			}
			tree.Current.append(NewFileNode(parts[1], size))
		}
	}

	//sum := WalkTree(root, 0)

	tree.ChangeDirectory("/")
	fmt.Println(tree.Current.Links)
	tree.ChangeDirectory("csmqbhjv")
	fmt.Println(tree.Current.Links)

	//fmt.Println(sum)
}
