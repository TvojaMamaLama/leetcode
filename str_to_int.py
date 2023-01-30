class Solution:
    def myAtoi(self, s: str) -> int:
        if s[1].isalpha() or s[0].isalpha():
            return 0

        MIN, MAX = -2 ** 31, 2**31 - 1

        clean_str = [char for char in s if char.isdigit() or char == '-' or char == '.']
        try:
            clean_str = clean_str[:clean_str.index('.')]
        except:
            pass

        sign = -1 if clean_str[0] == '-' else 1
        if sign == -1:
            clean_str.remove('-')
        answer = 0
        for idx, s in enumerate(clean_str[::-1]):
            answer += int(s) * (10 ** idx)

        answer *= sign
        if answer > MAX:
            return MAX
        if answer < MIN:
            return MIN

        return answer

        

s = Solution()
print(s.myAtoi('42'))