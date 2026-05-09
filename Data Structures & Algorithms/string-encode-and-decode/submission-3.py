class Solution:

    def encode(self, strs: List[str]) -> str:
        res = ''
        for str in strs:
            res += f'{len(str)}#{str}'
        return res

    def decode(self, s: str) -> List[str]:
        res = []
        i = 0
        length = ''
        while i < len(s):
            if s[i] != '#':
                length += s[i]
                i += 1
            else:
                res.append(s[i+1:i+int(length)+1])
                i += int(length) + 1
                length = ''
        return res
