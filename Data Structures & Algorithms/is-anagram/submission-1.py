class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        count_s = [0 for _ in range(26)]
        count_t = count_s.copy()

        for letter in s:
            count_s[ord(letter) - ord('a')] += 1

        for letter in t:
            count_t[ord(letter) - ord('a')] += 1

        return count_s == count_t