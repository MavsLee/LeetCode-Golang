package graph

type UF struct {
	root  []int // stores parent node
	rank  []int // stores node length including itself
	count int
}

func NewUF(n int) *UF {
	root := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
		rank[i] = 1
	}
	return &UF{
		root:  root,
		rank:  rank,
		count: n,
	}
}

// path compress
func (u *UF) Find(n int) int {
	temp := n
	for u.root[temp] != temp {
		temp = u.root[temp]
	}
	for n != temp {
		c := u.root[n]
		u.root[n] = temp
		n = c
	}
	return temp
}

// union by rank, try to keep tree balanced
func (u *UF) Connect(n1 int, n2 int) {
	p1 := u.Find(n1)
	p2 := u.Find(n2)
	if p1 != p2 {
		u.root[p2] = p1
		u.count--
	}
}

func (u *UF) IsConnected(n1 int, n2 int) bool {
	return u.Find(n1) == u.Find(n2)
}

func (u *UF) UnionByRank(n1 int, n2 int) {
	p1 := u.Find(n1)
	p2 := u.Find(n2)
	if p1 != p2 {
		if u.rank[p1] > u.rank[p2] {
			u.root[p2] = p1 // connect p2 to p1, since p1 has more child
		} else if u.rank[p1] < u.rank[p2] {
			u.root[p1] = p2
		} else {
			u.root[p2] = p1
			u.rank[p1]++
		}
		u.count--
	}
}

func (u *UF) GetConnectedComponents() int {
	return u.count
}
