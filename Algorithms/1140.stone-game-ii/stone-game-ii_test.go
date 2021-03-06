package problem1140

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	piles []int
	ans   int
}{

	{
		[]int{5819, 9551, 3626, 8100, 6991, 4067, 581, 3914, 895, 9859, 3463, 4463, 851, 1993, 6596, 408, 2950, 5818, 1433, 6552, 8416, 837, 7084, 5066, 1514, 6417, 9411, 9331, 5321, 7705, 1376, 6956, 6964, 2371, 5858, 9570, 6367, 9973, 7921, 2004, 8642, 8935, 861, 3857, 7807, 5708, 5020, 4558, 9641, 2286, 7931, 9637, 7542, 5899, 3814, 491, 6356, 9458, 9074, 8037, 7722, 5403, 7363, 8774, 9165, 3799, 7304, 2596, 2319, 5555, 3382, 8311, 6396, 7246, 2193, 7019, 3019, 4814, 6450, 1934, 9388, 4501, 909, 215, 1656, 3799, 6611, 8907, 739, 2678, 1342, 8707, 4648, 4223, 5271, 5970, 9702, 9413, 6121, 3915},
		276186,
	},

	{
		[]int{8270, 7145, 575, 5156, 5126, 2905, 8793, 7817, 5532, 5726, 7071, 7730, 5200, 5369, 5763, 7148, 8287, 9449, 7567, 4850, 1385, 2135, 1737, 9511, 8065, 7063, 8023, 7729, 7084, 8407},
		98008,
	},

	{
		[]int{2, 7, 9, 4, 4},
		10,
	},

	// 可以有多个 testcase
}

func Test_stoneGameII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, stoneGameII(tc.piles), "输入:%v", tc)
	}
}

func Benchmark_stoneGameII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			stoneGameII(tc.piles)
		}
	}
}
