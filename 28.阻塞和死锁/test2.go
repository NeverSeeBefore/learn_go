package main

import (
	"fmt"
	"strings"
)

// 编写一个流水线部件
// 他需要记住上一前面出现过的所有值，并且只有在值从未出现过的情况下才会将值流入下一部件

// 编写一个流水线部件
// 他接收字符串，并拆分成单词，然后向下一流水线一步一步发送这些单词

// 竞争选举
func main() {
	str := `
	Earth, the pale blue dot.

	We were hunters and foragers.
	The frontier was everywhere.
	We were bounded only by the earth, and the ocean, and the sky.

	The open road still softly calls.

	Our little terraqueous globe as the madhouse of those hundred thousand millions of worlds.
	We, who cannot even put our own planetary home in order,
	riven with rivalries and hatreds;
	are we to venture out into space?

	By the time we are ready to settle even the nearest other planetary systems,
	we will have changed.
	The simple passage of so many generations will have changed us.
	Necessity will have changed us.

	We are... an adaptable species.

	It will not be we who reach Alpha Centauri and the other nearby stars.
	It will be a species very like us,
	but with more of our strengths and fewer of our weaknesses,
	more confident, farseeing, capable and prudent.

	For all our failings,
	despite our limitations and fallibilities,
	we humans are capable of greatness.

	What new wonders undreamt of in our time
	will we have wrought in another generation? And another...?
	How far will our nomadic species have wandered by the end of the next century?
	And the next millennium...?
	Our remote descendants safely arrayed on many worlds through the Solar System
	and beyond,
	will be unified
	by their common heritage,
	by their regard for their home planet,
	and by the knowledge that, whatever other life may be, the only humans in all the Universe...
	come from Earth.

	They will gaze up and strain to find the blue dot in their skies.
	They will marvel at how vulnerable the repository of all our potential once was,
	how perilous our infancy,
	how humble our beginnings,
	how many rivers we had to cross
	before we found our way.`

	c0 := make(chan string)
	c1 := make(chan string)

	go splitGopher(str, c0)
	go filterGopher(c0, c1)

	printGopher(c1)
}

func splitGopher(s string, downstream chan string) {
	words := strings.Fields(s)
	for _, word := range words {
		downstream <- word
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	wordMap := map[string]bool{}

	for word := range upstream {
		if !wordMap[word] {
			wordMap[word] = true
			downstream <- word
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	count := 0
	for word := range upstream {
		fmt.Println(word)
		count += 1
	}
	fmt.Println("total:", count)
	return
}
