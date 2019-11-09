package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStore() {
	for page != nil {
		fmt.Println(page.text)
		scanner.Scan()
		page = page.nextPage
	}
}

func (page *storyPage) addPageToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage
	}

	page.nextPage = &storyPage{text, nil}
}

func (page *storyPage) addPageAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

// TODO: delete

func main() {
	book := storyPage{"It was a dark and stormy night.", nil}
	book.addPageToEnd("You are alone and need to find the secret scroll before Akatsuki does.")
	book.addPageToEnd("You see Kakuzu ahead, holding a dead shinobi.")

	book.addPageAfter("Testing")
	book.playStore()
}
