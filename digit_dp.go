//https://leetcode.com/problems/total-waviness-of-numbers-in-range-ii/description/?envType=daily-question&envId=2026-06-05

import "strconv"

type pair struct {
	first  int64
	second int64
}

var dp [16][11][11][2][2]pair

func reset() {
	for i := 0; i < 16; i++ {
		for a := 0; a < 11; a++ {
			for b := 0; b < 11; b++ {
				for tg := 0; tg < 2; tg++ {
					for st := 0; st < 2; st++ {
						dp[i][a][b][tg][st] = pair{-1, -1}
					}
				}
			}
		}
	}
}

func totalWaviness(n1 int64, n2 int64) int64 {
	if n2 <= 100 {
		return 0
	}
	reset()
	var f func(ind, a, b, tg, st int, t string) pair
	f = func(ind, a, b, tg, st int, t string) pair {
		n := len(t)
		if ind >= n {
			return pair{int64(0), int64(1)}
		}
		if dp[ind][a+1][b+1][tg][st].first != -1 {
			return dp[ind][a+1][b+1][tg][st]
		}
		rel := int64(0)
		cnt := int64(0)
		lim := 9
		if tg == 1 {
			lim = int(byte(t[ind]) - '0')
		}
		for i := 0; i <= lim; i++ {
			ntg := 0
			yes := int64(0)
			if tg == 1 && i == lim {
				ntg = 1
			}
			x := pair{0, 0}
			if st == 0 {
				if i == 0 {
					x = f(ind+1, a, b, ntg, st, t)
				} else {
					x = f(ind+1, b, i, ntg, 1, t)
				}
			} else {
				x = f(ind+1, b, i, ntg, st, t)
				if a != -1 {
					if (a < b && b > i) || (a > b && b < i) {
						yes = int64(1)
					}
				}
			}
			rel += yes*x.second + x.first
			cnt += x.second
		}
		dp[ind][a+1][b+1][tg][st] = pair{rel, cnt}
		return pair{rel, cnt}
	}
	t1 := strconv.FormatInt(n1-1, 10)
	l := f(0, -1, -1, 1, 0, t1)
	reset()
	t2 := strconv.FormatInt(n2, 10)
	r := f(0, -1, -1, 1, 0, t2)
	return r.first - l.first
}