package kata

import "math"

/*
At most:
0 participated -> 0 handshakes
1 participated -> 0 handshakes
2 participated -> 1 handshakes
3 participated -> 3 handshakes
4 participated -> 6 handshakes
5 participated -> 10 handshakes
...
n participated -> n*(n-1)/2 handshakes

2h + 1/4 = n^2 - n + 1/4
2h + 1/4 = (n - 1/2)^2
n = sqrt(2h + 1/4) + 1/2
*/
func GetParticipants(h int) (n int) {
	if h <= 0 {
		return 0
	}
	n = int(math.Ceil(math.Sqrt(2*float64(h)+1.0/4) + 1.0/2))
	return
}
