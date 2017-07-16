func lengthOfLongestSubstring(s string) int {
    start, end, maxLen := 0, 0, 0
    m := map[byte]int{}

    for ; end < len(s); end++ {
        id, ok := m[s[end]]
        if ok && id >= start {
            start = id + 1
        }
        m[s[end]] = end
        curLen := end - start + 1
        if curLen > maxLen {
            maxLen = curLen
        }
    }

    return maxLen
}
