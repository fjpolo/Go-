package main

// Homework:
//
// Non Player Characters
//	- Talk
//	- Interact
//	- Fight
//	- Move aroung the graph
//
// Items that can be picked up or placed down
//
// Accept natural language as input (then parse)
//
//

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* Import */

/*****************************
 *
 * Structures
 *
 *****************************/

/* choices */
type choices struct {
	/**/
	cmd         string
	description string
	nextNode    *storyNode
	nextChoice  *choices
}

/* storyNode */
type storyNode struct {
	/**/
	text    string
	choices *choices
}

/* NPC */
type NPC struct {
	description string
	health      int
}

/*****************************
 *
 * Methods
 *
 *****************************/
/* addChoice*/
func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	/* Create choice*/
	choice := &choices{cmd, description, nextNode, nil}
	/* Check for NIL */
	if node.choices == nil {
		node.choices = choice
	} else {
		currentChoice := node.choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = choice
	}
}

/* render */
func (node *storyNode) render() {
	/* Print description */
	fmt.Println(node.text)
	/* Print choices */
	currentChoice := node.choices
	for currentChoice != nil {
		fmt.Println(currentChoice.cmd, ":", currentChoice.description)
		currentChoice = currentChoice.nextChoice
	}
}

/* executeCmd */
func (node *storyNode) executeCmd(cmd string) *storyNode {
	/*  */
	currentChoice := node.choices
	for currentChoice != nil {
		/* Match command */
		if strings.ToLower(currentChoice.cmd) == strings.ToLower(cmd) {
			return currentChoice.nextNode
		}
		currentChoice = currentChoice.nextChoice
	}
	/* Wrong command */
	fmt.Println("Wrong command!")
	return node
}

/* play */
var scanner *bufio.Scanner

func (node *storyNode) play() {
	/* Print on screen*/
	node.render()
	/* Wait for command */
	if node.choices != nil {
		/* Scan inputs */
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
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

	/* scanner */
	scanner = bufio.NewScanner(os.Stdin)
	/* Starting point*/
	start := storyNode{text: `
You are in a large chamber, deep underground.
You see three passages leading out. A north
passage leads into darkness. To the south, a
passage appears to head upwards. The eastern
passage appears flat and well traveled.
Choose wisely:`}

	/* Dark room */
	darkRoom := storyNode{text: "It is pitch black. You cannot see a thing."}

	/* darkRoomLit */
	darkRoomLit := storyNode{text: "The dark passage is now lit by your lantern. You can continue north or head back south."}

	/* Grue*/
	grue := storyNode{text: "While stumbling around in the darkness, you are eaten by a grue."}

	/* trap */
	trap := storyNode{text: "You head down the well traveled path when suddenly a trap door opens and you fall into a pit."}

	/* treasure */
	treasure := storyNode{text: "You arrive at a small chamber, filled with treasure!"}

	/* Start - Choices*/
	start.addChoice("N", "Go North.", &darkRoom)
	start.addChoice("S", "Go South.", &darkRoom)
	start.addChoice("E", "Go East.", &trap)
	//start.addChoice("W", "Go West", &darkRoom)

	/* Room options */
	darkRoom.addChoice("S", "Try to go back South.", &grue)
	darkRoom.addChoice("O", "Turn lantern on", &darkRoomLit)
	darkRoom.addChoice("N", "Go North.", &treasure)
	darkRoom.addChoice("S", "Go South.", &start)
	darkRoomLit.addChoice("N", "Go North.", &treasure)
	darkRoomLit.addChoice("S", "Try to go back South.", &grue)

	/* Play story */
	start.play()
	/* The end */
	fmt.Println()
	fmt.Println("The end.")
	/* Pause */
	var input string
	fmt.Scanln(&input)
}
