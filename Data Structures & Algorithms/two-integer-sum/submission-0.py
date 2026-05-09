class Solution:
        def twoSum(self, nums: list[int], target: int) -> list[int]:
        mapping = {}
        for i, val in enumerate(nums):
            diff = target - val
            if diff in mapping:
                return [mapping[diff], i]
            mapping[val] = i
        