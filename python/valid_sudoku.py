class Solution:
    def isValidSudoku(self, board: list[list[str]]) -> bool:
        for line in board:
            digits = [digit for digit in line if digit != '.']
            if len(digits) != len(set(digits)):
                return False

        for column in zip(*board):
            digits = [digit for digit in column if digit != '.']
            if len(digits) != len(set(digits)):
                return False

        for i in (0,3,6):
            for k in (0,3,6):
                digits = [board[x][y] for x in range(k, k+3) for y in range(i, i+3) if board[x][y] != '.']
                if len(digits) != len(set(digits)):
                    return False
        
        return True


s = Solution()
print(s.isValidSudoku([["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]))
