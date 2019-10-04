package alg

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_LongestNoDuplicateSubString(t *testing.T) {
	l, p := longestNoDuplicateSubString("aabccakdac")
	t.Log(l)
	t.Log(p)
}

//最长无重复子串: 采用滑动窗口的方式记录最大长度&&最大长度的起始地点
func longestNoDuplicateSubString(str string) (length, wndStartPos int) {
	if len(str) <= 1 {
		return len(str), len(str)
	}
	//用于记录字符对应出现的下标
	wnd := make(map[uint8]int)
	wnd[str[0]] = 0

	for i := 1; i < len(str); i++ {
		//lasPos 上一个字符最近出现的位置，wndStartPos 表示 窗口开始的位置
		if lastPos, ok := wnd[str[i]]; ok && lastPos > wndStartPos {
			wndStartPos = lastPos + 1
		} else {
			length = i + 1 - wndStartPos
		}
		wnd[str[i]] = i
	}
	return
}

func Test_longestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("bb"))
}

/*
最长回文子串: dp[j][i] = true 表示 从j ~ i 是回文串 那么j+1 ~ i-1 也是回文串
得到状态转移方程: i-j <= 1 表示对称或者i==j dp[j][i] = dp[j+1][i-1]
				i-j > 1 && str[j] = str[i] dp
*/
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	maxLen := 1 //字符串长度大于0 保证maxLen >= 1
	longestStr := s[:1]
	dp := make(map[int]map[int]bool, len(s))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			dp[i] = map[int]bool{j: false}
		}
	}
	for i := 0; i < len(s); i++ {
		for j := i; j >= 0; j-- {
			//trick 防止 oom
			if s[j] == s[i] && (i-j <= 1 || dp[j+1][i-1]) {
				dp[j][i] = true
				if maxLen <= i-j+1 {
					maxLen = i - j + 1
					longestStr = s[j : i+1]
				}
			}
		}
	}
	return longestStr
}

//todo test
func restoreIpAddresses(s string) []string {
	result := []string{}
	if len(s) < 4 || len(s) > 12 {
		return result
	}
	_restoreIpAddresses(s, "", 0, 4, result)
	return result
}

func isNum(str string) bool {
	if v, err := strconv.Atoi(str); err == nil && v <= 255 {
		return true
	}
	return false
}

func _restoreIpAddresses(s, ip string, pos, n int, result []string) {
	//n 表示left .
	if n == 0 {
		if pos == len(s) {
			result = append(result, ip)
		}
		return
	}

	for i := pos; i < len(s); i++ {
		if isNum(s[pos:i]) {
			ip = fmt.Sprintf("%s.%s", ip, s[pos:i])
			_restoreIpAddresses(s, ip, i, n-1, result)
		}
		return
	}
}

//76. 最小覆盖子串
func minWindow(s string, t string) string {
	tMap := make(map[int32]bool, len(t))
	for _, v := range t {
		tMap[v] = true
	}

	for i := 0; i < len(s); i++ {
	}

	return ""
}

//10 正则表达式匹配问题
func isMatch(s, p string) bool {
	dp := make(map[int]map[int]bool, len(s))
	dp[0] = map[int]bool{0: true}
	for i := 1; i <= len(s); i++ {
		tmp := make(map[int]bool, len(p))
		dp[i] = tmp
	}

	for j := 1; j <= len(p); j++ {
		if string(p[j-1]) == "*" && j > 1 {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j < len(p); j++ {
			if string(p[j-1]) == "." || string(p[j-1]) == string(s[j-1]) {
				dp[i][j] = dp[i-1][j-1]
			}

			if string(p[j-1]) == "*" {
				if j > 1 && string(p[j-2]) != "." && string(p[j-2]) != string(s[j-1]) {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2] || dp[i][j-1] || dp[i-1][j]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

//单词拆分 II [refer](https://blog.csdn.net/qq_17550379/article/details/85847803)
func Solution(s string, list []string) []string {
}
