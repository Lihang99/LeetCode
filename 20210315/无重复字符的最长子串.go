/*
滑动窗口
一个map记录是否存在
右指针先走 如果是未重复的就记录进map 计算长度
*/

func lengthOfLongestSubstring(s string) int {
	m := map[byte]bool{}
	res := 0
	p1,p2 := 0,0
	for p2<len(s){
		for m[s[p2]]{
			m[s[p1]]=false
			p1++
		}
		m[s[p2]]=true
		res = max(res,(p2-p1+1))
		p2++
	}
	return res
}

func max(a,b int) int{
	if a>b {
		return a
	} else{
		return b
	}
}