package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Name     string
	Size     uint64
	IsDir    bool
	Children []*Node
	Parent   *Node
}

type Tree struct {
	Root    *Node
	Current *Node
}

func (parent *Node) append(children ...*Node) {
	for _, child := range children {
		parent.Children = append(parent.Children, child)
		child.Parent = parent
	}
}

func (tree *Tree) GoToRoot() {
	tree.Current = tree.Root
}

func (tree *Tree) GoToParent() {
	tree.Current = tree.Current.Parent
}

func WalkTree(node *Node, sum uint64) uint64 {
	if node.IsDir {
		for _, child := range node.Children {
			sum = WalkTree(child, sum)
		}
	} else {
		sum += node.Size
	}
	return sum
}

func LoadTestTree(root *Node) *Node {
	a := &Node{Name: "foobar.txt", Size: 10}

	b := &Node{Name: "images", IsDir: true}
	c := &Node{Name: "img.jpg", Size: 10}
	d := &Node{Name: "img2.jpg", Size: 10}
	b.append(c, d)

	e := &Node{Name: "css", IsDir: true}
	f := &Node{Name: "style.css", Size: 11}
	g := &Node{Name: "style2.css", Size: 11}
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
	root := &Node{Name: "/", IsDir: true}
	current := root

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		parts := strings.Split(line, " ")

		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == "/" {
					current = root
				} else if parts[2] == ".." {
					if current == root {
						fmt.Println("Warning: cd .. at root")
						continue
					} else {
						current = current.Parent
					}
				} else {
					// TODO
				}
			} else if parts[1] == "ls" {
				continue
			} else {
				fmt.Printf("WTF? %s\n", line)
				continue
			}

		} else if parts[0] == "dir" {
			dirName := strings.SplitN(line, " ", 2)
			current.append(&Node{Name: dirName[1], IsDir: true})
		} else {
			splits := strings.SplitN(line, " ", 2)
			name := splits[1]
			size, _ := strconv.ParseUint(splits[0], 10, 64)
			current.append(&Node{Name: name, Size: size})
		}

	}

	sum := WalkTree(root, 0)
	fmt.Println(sum)
}
