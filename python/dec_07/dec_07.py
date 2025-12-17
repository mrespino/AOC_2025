
def first():
    from collections import deque
    with open("input_07.txt") as f:
        lines = [line.rstrip() for line in f]
    rows = len(lines)
    cols = len(lines[0])

    # Find S (start column)
    try:
        start_col = lines[0].index('S')
    except ValueError:
        raise Exception("No start S found")

    # Each beam is (row, col)
    queue = deque()
    queue.append((0, start_col))
    split_count = 0

    visited = set()
    while queue:
        row, col = queue.popleft()
        if row >= rows:
            continue
        if (row, col) in visited:
            continue
        visited.add((row, col))
        cell = lines[row][col]
        if cell == '.' or cell == 'S':
            queue.append((row + 1, col))
        elif cell == '^':
            split_count += 1
            if col > 0:
                queue.append((row + 1, col - 1))
            if col < cols - 1:
                queue.append((row + 1, col + 1))

    print(f"first: {split_count}")

# Placeholder for part 2 if needed in the future
def second():
    with open("input_07.txt") as f:
        lines = [line.rstrip() for line in f]
    rows = len(lines)
    cols = len(lines[0])

    # Find S (start column)
    try:
        start_col = lines[0].index('S')
    except ValueError:
        raise Exception("No start S found")

    # DP table: ways[row][col] = number of ways to reach (row, col)
    ways = [[0] * cols for _ in range(rows + 1)]
    ways[0][start_col] = 1

    for row in range(rows):
        for col in range(cols):
            w = ways[row][col]
            if w == 0:
                continue
            cell = lines[row][col]
            if cell == '.' or cell == 'S':
                ways[row + 1][col] += w
            elif cell == '^':
                if col > 0:
                    ways[row + 1][col - 1] += w
                if col < cols - 1:
                    ways[row + 1][col + 1] += w

    timeline_count = sum(ways[rows])
    print(f"second: {timeline_count}")

first()
second()
