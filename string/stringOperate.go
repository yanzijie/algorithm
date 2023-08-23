package string

/*
字符串模式匹配:
	主串 : "ABCdkjeGoMVP"
	模式串 : "Go"
	判断模式串是否存在于主串中
朴素模式匹配:
	暴力匹配, 拿到模式串的长度, 在主串中从第一个字符开始匹配, 如果匹配失败，主串向右滑动一位再进行匹配
*/

// NaivePatternMatching 朴素模式匹配
// 参数: text-主串，pattern-模式串
// 返回值: 成功返回位置下标，失败返回-1
func NaivePatternMatching(text, pattern string) int {
	pLen := len(pattern)
	tLen := len(text)

	for i := 0; i <= (tLen - pLen); i++ {
		j := 0
		for j < pLen {
			if text[i+j] != pattern[j] {
				break
			}
			j++
		}

		if j == pLen {
			return i
		}
	}

	return -1
}
