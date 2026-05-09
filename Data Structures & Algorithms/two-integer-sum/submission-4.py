class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        l = 0
        r = len(nums) - 1

        while l < r:
            s = nums[l] + nums[r]
            if s == target:
                return [l, r]
            elif abs(s) < abs(target):
                l += 1
            else:
                r -= 1

        return [l, r]