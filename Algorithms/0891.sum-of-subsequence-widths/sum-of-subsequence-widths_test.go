package problem0891

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{96, 87, 191, 197, 40, 101, 108, 35, 169, 50, 168, 182, 95, 80, 144, 43, 18, 60, 174, 13, 77, 173, 38, 46, 80, 117, 13, 19, 11, 6, 13, 118, 39, 80, 171, 36, 86, 156, 165, 190, 53, 49, 160, 192, 57, 42, 97, 35, 124, 200, 84, 70, 145, 180, 54, 141, 159, 42, 66, 66, 25, 95, 24, 136, 140, 159, 71, 131, 5, 140, 115, 76, 151, 137, 63, 47, 69, 164, 60, 172, 153, 183, 6, 70, 40, 168, 133, 45, 116, 188, 20, 52, 70, 156, 44, 27, 124, 59, 42, 172},
		136988321,
	},

	{
		[]int{2, 1, 3, 4},
		23,
	},

	{
		[]int{2, 1, 3},
		6,
	},

	// 可以有多个 testcase
}

func Test_sumSubseqWidths(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, sumSubseqWidths(tc.A), "输入:%v", tc)
	}
}

func Benchmark_sumSubseqWidths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sumSubseqWidths(tc.A)
		}
	}
}
