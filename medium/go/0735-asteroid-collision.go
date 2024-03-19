package main

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for i := 0; i < len(asteroids); i++ {
		if len(stack) == 0 || asteroids[i] > 0  {
			stack = append(stack, asteroids[i])
			continue
		}

		for  {
			lastAsteroid := stack[len(stack)-1]
			if lastAsteroid < 0 {
				stack = append(stack, asteroids[i])
				break
			} else if lastAsteroid == -asteroids[i] {
				stack = stack[:len(stack)-1]
				break
			} else if lastAsteroid > -asteroids[i] {
				break
			} else {
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					stack = append(stack, asteroids[i])
					break
				}
			}
		}
	}

	return stack
}