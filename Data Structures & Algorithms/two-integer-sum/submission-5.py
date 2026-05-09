class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        indices = {}

        for i, n in enumerate(nums):
            indices[n] = i

        for i in range(len(nums)):
            diff = target - nums[i]
            if diff in indices:
                return [i, indices[diff]]