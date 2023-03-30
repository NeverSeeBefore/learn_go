package main

import (
	"fmt"
	"strings"
)

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

	str = strings.ToLower(str)
	// str = strings.ReplaceAll(str, ",", "")
	// str = strings.ReplaceAll(str, ".", "")
	// str = strings.ReplaceAll(str, "?", "")
	// str = strings.ReplaceAll(str, ";", "")
	fields := strings.Fields(str)

	fmt.Println(fields)

	set := make(map[string]int, 0)
	for _, s := range fields {
		set[s]++
	}

	for k, v := range set {
		fmt.Printf("%v: %v\n", k, v)
	}
}
