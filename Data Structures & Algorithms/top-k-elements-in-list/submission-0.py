class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = {}
        for num in nums:
            if num not in count:
                count[num] = 0
            count[num] += 1
        count = sorted(count, key=lambda x: count[x], reverse=True)
        return count[:k]
        