// 93. Restore IP Addresses

// Given a string containing only digits, restore it by returning all possible valid IP address combinations.

// Example:
// Input: "25525511135"
// Output: ["255.255.11.135", "255.255.111.35"]

package leetcode

import "strings"

func restoreIpAddresses(s string) []string {
	bytes := make([]string, 0, 4)
	return dfs(nil, bytes, s)
}

func dfs(IPs, bytes []string, s string) []string {
	if len(bytes) == 4 {
		if len(s) == 0 {
			IPs = append(IPs, strings.Join(bytes, "."))
		}
		return IPs
	}

	if len(s) > 0 && s[0] == '0' {
		return dfs(IPs, append(bytes, "0"), s[1:])
	}

	var num int
	for i := 0; i < len(s); i++ {
		num = num*10 + int(s[i]-'0')
		if num > 255 {
			break
		}
		IPs = dfs(IPs, append(bytes, s[:i+1]), s[i+1:])
	}
	return IPs
}
