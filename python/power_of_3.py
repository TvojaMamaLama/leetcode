class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        if n == 1:
            return True

        power = 0
        i = 1
        while power < n:
            power = pow(3, i)
            if power == n:
                return True

            i += 1

        return False


print(Solution().isPowerOfThree(9))
