package main

/* Import */
import (
	"bufio"
	"fmt"
	"os"
)

/*****************************
 *
 * Structures
 *
 *****************************/

/* storyNode */
type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

/*****************************
 *
 * Methods
 *
 *****************************/
/*playStory*/
func (node *storyNode) playStory() {

	/* Print */
	fmt.Println(node.text)

	/* Check for end of story*/
	if node.yesPath != nil && node.noPath != nil {
		/* Scan answer*/
		scanner := bufio.NewScanner(os.Stdin)
		/* Wait for an answer */
		for {
			/* Scan*/
			scanner.Scan()
			answer := scanner.Text()

			/* Evaluate answer */
			if (answer == "Y") || (answer == "Yes") || (answer == "yes") || (answer == "y") {
				node.yesPath.playStory()
				break
			} else if (answer == "N") || (answer == "No") || (answer == "no") || (answer == "n") {
				node.noPath.playStory()
				break
			} else {
				/**/
				fmt.Println("Please answer with Y or N")
			}
		}
	}
}

/**/
func (node *storyNode) printStory(depth int) {

	/**/
	for i := 0; i < depth*2; i++ {
		fmt.Print(" ")
	}

	/**/
	fmt.Println(node.text)
	if node.yesPath != nil {
		node.yesPath.printStory(depth + 1)
	}
	if node.noPath != nil {
		node.noPath.printStory(depth + 1)
	}
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

	/* Game start*/
	root := storyNode{"You are at the entrance of a dark cave.Do you want to enter?", nil, nil}
	/* Game end - WIN */
	winning := storyNode{"You won!", nil, nil}
	/* Gmae end - LOOSE */
	losing := storyNode{"You lost...", nil, nil}
	/* Assign paths */
	root.yesPath = &losing
	root.noPath = &winning
	/* Print story */
	//root.printStory(0)
	/* Play */
	root.playStory()
}
