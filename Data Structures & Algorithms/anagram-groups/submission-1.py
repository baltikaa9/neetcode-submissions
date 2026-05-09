class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = []
        anagrams = [{}] * len(strs)

        for i, s in enumerate(strs):
            anagrams[i] = {}
            for char in s:
                anagrams[i][char] = anagrams[i].get(char, 0) + 1

        for i in range(len(anagrams)):
            if anagrams[i] == {}:
                continue
            anagram = [strs[i]]
            for j in range(i + 1, len(anagrams)):
                if anagrams[i] == anagrams[j]:
                    anagram.append(strs[j])
                    anagrams[j] = {}
            res.append(anagram)
        return res