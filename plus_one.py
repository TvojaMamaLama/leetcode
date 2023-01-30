class Solution:
    def plusOne(self, digits: list[int]) -> list[int]:
        s = ''.join(map(str, digits))
        s = int(s)
        s += 1
        s = str(s)
        answer = []
        for d in s:
            answer.append(int(d))

        return answer
