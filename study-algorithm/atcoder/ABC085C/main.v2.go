package main

import (
	"fmt"
)

type combi struct{ yukichi, higuchi, natsume int }

func main() {
	var n, y int
	fmt.Scanf("%d %d", &n, &y)
	y /= 1000
	c := search(n, y)
	fmt.Printf("%d %d %d", c.yukichi, c.higuchi, c.natsume)
	return
}

func search(num, y int) combi {
	state := combi{}
	state.natsume = num

	d := (y - price(state)) / 9
	state.yukichi += d
	state.natsume -= d

	if state.natsume < 0 || state.higuchi < 0 || state.yukichi < 0 {
		return combi{-1, -1, -1}
	}

	d = (y - price(state)) / 4
	state.higuchi += d
	state.natsume -= d

	if state.natsume < 0 || state.higuchi < 0 || state.yukichi < 0 {
		return combi{-1, -1, -1}
	}

	d = (y - price(state)) / 3
	state.yukichi -= d
	state.higuchi += 3 * d
	state.natsume -= 2 * d

	if state.natsume < 0 || state.higuchi < 0 || state.yukichi < 0 {
		return combi{-1, -1, -1}
	}

	d = (y - price(state)) / 2
	state.yukichi -= 2 * d
	state.higuchi += 5 * d
	state.natsume -= 3 * d

	if state.natsume < 0 || state.higuchi < 0 || state.yukichi < 0 {
		return combi{-1, -1, -1}
	}

	d = y - price(state)
	state.yukichi -= 3 * d
	state.higuchi += 7 * d
	state.natsume -= 4 * d

	if state.natsume < 0 || state.higuchi < 0 || state.yukichi < 0 {
		return combi{-1, -1, -1}
	}

	return state
}

func adjust(c combi, y int) combi {
	if p := price(c); p > y {
	} else if p < y {
	} else {
	}
	return combi{-1, -1, -1}
}

func price(c combi) int {
	return c.yukichi*10 + c.higuchi*5 + c.natsume
}
