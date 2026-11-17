import enum


class Direction(enum.Enum):

    UP = "up"
    LEFT = "left"
    DIAGONAL = "diagonal"


def lcs_length(x, y):
    results = [[-1 for _ in range(len(x) + 1)] for _ in range(len(y) + 1)]
    directions = [[-1 for _ in range(len(x) + 1)] for _ in range(len(y) + 1)]

    for i in range(len(x)):
        for j in range(len(y)):
            if i == 0 and j == 0:
                if x[i] == y[j]:
                    result = 1
                    directions[i][j] = Direction.DIAGONAL
                else:
                    result = 0
                    directions[i][j] = Direction.UP
            elif i == 0:
                if x[i] == y[j]:
                    result = 1
                    directions[i][j] = Direction.DIAGONAL
                else:
                    result = results[i][j - 1]
                    directions[i][j] = Direction.LEFT
            elif j == 0:
                if x[i] == y[j]:
                    result = 1
                    directions[i][j] = Direction.DIAGONAL
                else:
                    result = results[i - 1][j]
                    directions[i][j] = Direction.UP
            else:
                if x[i] == y[j]:
                    result = results[i - 1][j - 1] + 1
                    directions[i][j] = Direction.DIAGONAL
                else:
                    if results[i][j - 1] > results[i - 1][j]:
                        result = results[i][j - 1]
                        directions[i][j] = Direction.LEFT
                    else:
                        result = results[i - 1][j]
                        directions[i][j] = Direction.UP

            results[i][j] = result

    return results, directions


def print_optimal_strategy(x, i, j, directions):
    if i < 0 or j < 0:
        return
    else:
        if directions[i][j] == Direction.DIAGONAL:
            print_optimal_strategy(x, i - 1, j - 1, directions)
            print(x[i], end="")
        elif directions[i][j] == Direction.UP:
            print_optimal_strategy(x, i - 1, j, directions)
        else:
            print_optimal_strategy(x, i, j - 1, directions)


if __name__ == "__main__":
    x = "ABCBDAB"
    y = "BDCABA"

    results, directions = lcs_length(x, y)
    for i in range(len(x)):
        print(f"--- i = {i} ---")
        for j in range(len(y)):

            print(
                f"(i, j) = {i, j} / x = {x[:i + 1]}, y = {y[:j + 1]} / result = {results[i][j]} (",
                end="",
            )
            print_optimal_strategy(x, i, j, directions)
            print(")")
