class Solution:
    def rotate(self, matrix: list[list[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        lenght = len(matrix)
        for i in range(lenght // 2 + lenght % 2):
            for j in range(lenght // 2):
                temp = matrix[lenght - 1 - j][i]
                matrix[lenght - 1 - j][i] = matrix[lenght - 1 - i][lenght - j - 1]
                matrix[lenght - 1 - i][lenght - j - 1] = matrix[j][lenght - 1 - i]
                matrix[j][lenght - 1 - i] = matrix[i][j]
                matrix[i][j] = temp

        print(matrix)

s = Solution()
print(s.rotate([[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]))