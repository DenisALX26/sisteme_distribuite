package main

import "fmt"

func main() {
	heights := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}

	names := make([]string, 0, len(heights))
	for name := range heights {
		names = append(names, name)
	}

	for i := range names {
		for j := i + 1; j < len(names); j++ {
			if heights[names[i]] < heights[names[j]] ||
				(heights[names[i]] == heights[names[j]] && names[i] > names[j]) {
				names[i], names[j] = names[j], names[i]
			}
		}
	}

	for _, name := range names {
		fmt.Printf("%s: %d\n", name, heights[name])
	}
}
