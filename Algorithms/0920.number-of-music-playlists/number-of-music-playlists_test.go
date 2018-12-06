package problem0920

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	L   int
	K   int
	ans int
}{

	{
		3,
		3,
		1,
		6,
	},

	{
		2,
		3,
		0,
		6,
	},

	{
		2,
		3,
		1,
		2,
	},

	// 可以有多个 testcase
}

func Test_numMusicPlaylists(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numMusicPlaylists(tc.N, tc.L, tc.K), "输入:%v", tc)
	}
}

func Benchmark_numMusicPlaylists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numMusicPlaylists(tc.N, tc.L, tc.K)
		}
	}
}
