package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

func getAnswer() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		answer := scanner.Text()

		if answer == "yes" || answer == "no" {
			return answer
		}

		fmt.Println("Wrong option! Answer yes or no.")
	}
}

func (node *storyNode) play() {
	fmt.Println(node.text)

	if node.yesPath == nil || node.noPath == nil {
		fmt.Println("End of the line.")
		return
	}

	answer := getAnswer()

	if answer == "yes" {
		node.yesPath.play()
	} else if answer == "no" {
		node.noPath.play()
	}
}

func (node *storyNode) printStory(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(("  "))
	}

	fmt.Print(node.text, "\n")

	if node.yesPath != nil {
		node.yesPath.printStory(depth + 1)
	}
	if node.noPath != nil {
		node.noPath.printStory(depth + 1)
	}
}

func main() {
	root := storyNode{"You are at the entrance of Akatsuki's hideout. Do you want to go in?", nil, nil}
	winningPath := storyNode{"You are alive!", nil, nil}
	losingPath := storyNode{"Pain killed you.", nil, nil}

	root.yesPath = &winningPath
	root.noPath = &losingPath

	root.printStory(0)
	fmt.Println()
	root.play()
}
