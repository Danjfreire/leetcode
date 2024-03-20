package main

func predictPartyVictory(senate string) string {
	queue := []rune{}

	for _, v := range senate {
		queue = append(queue, v)
	}

	direQueued := 0
	radiantQueued := 0

	for {
		if direQueued >= len(queue) {
			return "Dire"
		} else if radiantQueued >= len(queue) {
			return "Radiant"
		}

		current := queue[0]
		queue = queue[1:]

		if current == 'R' {
			if direQueued == 0 {
				radiantQueued++
				queue = append(queue, 'R')
			} else {
				direQueued--
			}
		} else if current == 'D' {
			if radiantQueued == 0 {
				direQueued++
				queue = append(queue, 'D')
			} else {
				radiantQueued--
			}
		}
	}
}