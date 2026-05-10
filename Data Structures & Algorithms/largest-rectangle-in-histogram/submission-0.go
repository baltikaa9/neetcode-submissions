type Bar struct {
	index  int
	height int
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([]Bar, 0, len(heights))

	for i := range heights {
		bar := Bar{i, heights[i]}

		for len(stack) > 0 && bar.height < stack[len(stack)-1].height {
			area := stack[len(stack)-1].height * (i - stack[len(stack)-1].index)

			if area > maxArea {
				maxArea = area
			}

			bar.index = stack[len(stack)-1].index
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, bar)
	}

	if len(stack) > 0 {
		for len(stack) > 0 {
			area := stack[len(stack)-1].height * (len(heights) - stack[len(stack)-1].index)

			if area > maxArea {
				maxArea = area
			}

			stack = stack[:len(stack)-1]
		}
	}

	return maxArea
}
