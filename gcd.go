package algs4

func Gcd(p, q uint64) uint64 {
	if q == 0 {
		return p
	}
	r := p % q
	return Gcd(q, r)
}
