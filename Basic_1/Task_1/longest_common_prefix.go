package main

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(prefix, s string) string {
	length := min(len(prefix), len(s))
	index := 0
	for index < length && prefix[index] == s[index] {
		index++
	}
	return prefix[:index]
}
