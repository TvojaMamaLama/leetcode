class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s

        answer = ['' for _ in range(numRows)]

        arr_num = 0
        up_down = 0
        for char in s:
            if arr_num == numRows - 1:
                up_down = 1
            
            if arr_num == 0:
                up_down = 0
            
            answer[arr_num] += char

            if not up_down:
                arr_num+=1
            else:
                arr_num-=1

        return ''.join(answer)
            
         

s = Solution()
print(s.convert('AB', 1))