package algs4

func ThreeSumCount(a []int64) int64 {
	n := len(a)
	cnt := int64(0)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if (a[i] + a[j] + a[k]) == 0 {
					cnt++
				}
			}
		}
	}

	return cnt
}
