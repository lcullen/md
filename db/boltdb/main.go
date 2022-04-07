package main

func main() {
	//db, err := bolt.Open("/Users/luoxiaowei/boltdb", os.FileMode(0666), nil)
	//if err != nil {
	//	log.Panic(err)
	//}
	//defer db.Close()
	//db.Begin(true)
	//db.Update(func(tx *bolt.Tx) error {
	//	tx.CreateBucketIfNotExists([]byte("learning boltdb"))
	//})

}

func getNum(s string) string {
	if len(s) <= 1 {
		return s
	}
	dmap := make(map[uint8][]bool, len(s))
	for i := range s {
		dmap[uint8(i)] = make([]bool, len(s))
		dmap[uint8(i)][i] = true
	}

	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if i-j == 1 {
					dmap[uint8(i)][j] = true
				} else {
					dmap[uint8(i)][j] = dmap[uint8(i)-1][j+1]
				}
			}
		}
	}
	l, r, maxLen := 0, 0, 1
	for i, arr := range dmap {
		for j, v := range arr {
			if v && j-int(i) > maxLen {
				l, r, maxLen = int(i), j, j-int(i)
			}
		}
	}
	return s[l:r]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return _gen(0, n)
}

func _gen(l, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}

	ret := []*TreeNode{}

	for i := l; i <= r; i++ {
		lnodes := _gen(l, i-1)
		rnodes := _gen(i+1, r)
		for _, lnode := range lnodes {
			for _, rnode := range rnodes {
				inode := &TreeNode{Val: i, Left: lnode, Right: rnode}
				ret = append(ret, inode)
			}
		}
	}
	return ret
}
