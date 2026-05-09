class Solution:
    def isValidSudoku(self, board: list[list[str]]) -> bool:
        boxes = [[] for _ in range(9)]

        for i in range(9):
            row = []
            for num in board[i]:
                if num == '.':
                    continue
                if num in row:
                    return False
                row.append(num)

        for i in range(9):
            col = []
            for row in board:
                if row[i] == '.':
                    continue
                if row[i] in col:
                    return False
                col.append(row[i])

        for i in range(9):
            for j in range(9):
                if board[i][j] == '.':
                    continue
                box_i = (i // 3 + j // 3) + (i // 3 * 2)
                boxes[box_i].append(board[i][j])

        for box in boxes:
            if len(box) != len(set(box)):
                return False

        return True