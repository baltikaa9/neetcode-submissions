class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if not nums: return 0
        nums = sorted(list(set(nums)))
        count = 1
        max_count = count

        for i in range(1, len(nums)):
            if nums[i] - nums[i-1] == 1:
                count += 1
                max_count = max(max_count, count)
            else:
                count = 1
        return max_count
        