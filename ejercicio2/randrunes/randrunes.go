package randrunes

import (
	"fmt"
	"math/rand"
)

func GenNum(max_n int) int {
    return rand.Intn(max_n)
}

func RandRunes(input_size int, input_chain string) string {
	output := make([]byte, input_size)

	fmt.Println("Input Size = %v", len(input_chain))

	for i := 0; i < input_size; i++ {
		output[i] = input_chain[GenNum(len(input_chain))]
	}

	return string(output)
}
