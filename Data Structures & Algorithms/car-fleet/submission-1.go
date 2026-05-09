type car struct {
	pos   int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]car, len(position))

	for i :=  range position {
		cars[i] = car{position[i], speed[i]}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})
	
	times := make([]float64, len(position))

	for i, car := range cars {
		times[i] = float64(target-car.pos) / float64(car.speed)
	}

	count := 1
	max := float64(target-cars[0].pos) / float64(cars[0].speed)

	for i := 1; i < len(cars); i++ {
		time := float64(target-cars[i].pos) / float64(cars[i].speed)
		if time > max {
			count++
			max = time
		}
	}

	return count
}
