class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = {}

        for s in strs:
            count = [0] * 26
            for char in s:
                count[ord(char) - ord('a')] += 1

            count = map(str, count)

            if (hash := ' '.join(count)) not in res:
                res[hash] = [s]
            else:
                res[hash].append(s)

        return list(res.values())