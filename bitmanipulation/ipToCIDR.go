package bitmanipulation

import (
	"fmt"
	"strconv"
	"strings"
)

func ipToCIDR(ip string, n int) []string {
	// convert string ip to number value
	var x int64
	parts := strings.Split(ip, ".")
	for _, p := range parts {
		pv, _ := strconv.ParseInt(p, 10, 64)
		x = x * 256 + pv
	}
	// begin with start ip, calculate how many ips current ip/CIDR block can cover by doing x & -x
	// increment start ip, minus total n until n <= 0

	var res []string
	for ;n>0; {
		var count int64
		count = x & -x
		if count == 0 {
			count = int64(getCount(n))
		}
		// in case the count(mask) is too large. For example count is 7 but we only need 2 more ips
		for ;count > int64(n); {
			count /= 2
		}
		res = append(res, strinifyCIDR(x, count))
		n = n - int(count)
		x = x + count
	}
	return res
}

func getCount(x int) int {
	len := 0
	for ; x > 1; {
		x >>= 1
		len++
	}
	return 1 << len
}

func strinifyCIDR(num int64, count int64) string {
	var d, c, b, a int
	// Compute right-most part of ip
	d = int(num & 255)
	// throw away the right-most part of ip
	num >>= 8
	c = int(num & 255)
	num >>= 8
	b = int(num & 255)
	num >>= 8
	a = int(num & 255)

	// this while loop to know how many digits of count is in binary
	// for example, 00001000 here the len will be 4.
	len := 0
	for ;count >0; {
		count /= 2
		len++
	}
	// Think about 255.0.0.7 -> 11111111 00000000 00000000 00000111
	// x & -x of it is 00000001, the mask is 32. (which is 32 - (1 - 1))
	mask := 32 - (len - 1)
	return fmt.Sprintf("%d.%d.%d.%d/%d", a, b, c, d, mask)
}