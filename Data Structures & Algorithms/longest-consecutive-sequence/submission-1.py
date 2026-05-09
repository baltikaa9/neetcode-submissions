class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        nums = set(nums)
        max_count = 0

        for n in nums:
            if n - 1 not in nums:
                count = 1
                while (n + count) in nums:
                    count += 1
                max_count = max(max_count, count)
        return max_count
        