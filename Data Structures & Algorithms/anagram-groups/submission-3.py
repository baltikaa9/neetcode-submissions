class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = {}

        for s in strs:
            count = [0] * 26
            for char in s:
                count[ord(char) - ord('a')] += 1

            hash = 0
            for i, c in enumerate(count):
                hash += c * 2 ** (len(count) - i - 1)

            if hash not in res:
                res[hash] = [s]
            else:
                res[hash].append(s)
        return list(res.values())