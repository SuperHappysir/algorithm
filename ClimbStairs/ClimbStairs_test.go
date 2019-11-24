package ClimbStairs

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			`假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
            每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
            注意：给定 n 是一个正整数。`,
			args{
				1000,
			},
			9079565065540428013,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
