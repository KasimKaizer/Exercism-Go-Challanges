package spiralmatrix

func SpiralMatrix(size int) [][]int {
	num, max := 1, size*size
	output := make([][]int, size)
	for idx := range output {
		output[idx] = make([]int, size)
	}
	for i := 0; num <= max; i++ {
		iterMax := size - i - 1
		for j := i; j < iterMax; j++ { // walk right
			output[i][j] = num
			num++
		}
		for k := i; k < iterMax; k++ { // walk down
			output[k][iterMax] = num
			num++
		}
		for l := iterMax; l > i; l-- { // walk left
			output[iterMax][l] = num
			num++
		}
		for m := iterMax; m > i; m-- { // walk up
			output[m][i] = num
			num++
		}
		if size%2 != 0 && num == max { // check if odd, if so then add the last num
			half := size / 2
			output[half][half] = num
			num++
		}

	}
	return output
}
