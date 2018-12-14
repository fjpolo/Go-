package main

/* Import */
import (
	"fmt"
)

/*****************************
 *
 * Structures
 *
 *****************************/

/* Story page structure */
type storyPage struct {
	text     string
	nextPage *storyPage
}

/*****************************
 *
 * Methods
 *
 *****************************/

/* playStory */
func (page *storyPage) playStory() {

	/* Tail call elimination loop */
	for page != nil {
		/* Print current page */
		fmt.Println(page.text)
		/* Go to next page*/
		page = page.nextPage
		//playStory(page.nextPage)   // If playStory is a function
		//page.nextPage.playStory() // if playStory is a method

	}
}

/* addToEnd */
func (page *storyPage) addToEnd(text string) {

	/* Seek last page */
	for page.nextPage != nil {
		/* Next page */
		page = page.nextPage
	}
	/* We are at the end page. Assign next page. */
	page.nextPage = &storyPage{text, nil}
}

/* addToEnd */
func (page *storyPage) addAfter(text string) {

	/* new page */
	newPage := &storyPage{text, page.nextPage}
	/* link previous page's nextpage to this new page */
	page.nextPage = newPage

}

/*****************************
 *
 * Functions
 *
 *****************************/

/*****************************
 *
 *          MAIN
 *
 *****************************/
func main() {

	/* Initialisation */
	page1 := storyPage{"You are stannding in an open field west of a white house.", nil}
	page1.addToEnd("You climb into the attic, it is pitch black, you can't see a thing!")
	page1.addToEnd("You are eaten by a Grue.")

	/* Test addAfter*/
	page1.addAfter("Testing addAfter")

	/* Play story */
	page1.playStory()
}
