from typing import List

# Read input
with open("input_04.txt", "r") as f:
    lines = [line.rstrip("\n") for line in f if line.strip() != ""]

# Directions for 8-neighbors
dirs = [
    (-1, -1), (-1, 0), (-1, 1), (0, -1), 
     (0, 1),(1, -1),  (1, 0),  (1, 1),
]

# Part One: count accessible rolls
grid_lines = lines
rows = len(grid_lines)
cols = len(grid_lines[0]) if rows > 0 else 0
result1 = 0
for r in range(rows):
    for c in range(cols):
        if grid_lines[r][c] != '@':
            continue
        adj = 0
        for dr, dc in dirs:
            nr, nc = r + dr, c + dc
            if 0 <= nr < rows and 0 <= nc < cols and grid_lines[nr][nc] == '@':
                adj += 1
                if adj >= 4:
                    break
        if adj < 4:
            result1 += 1

# count adjacent '@'
def adj_count(grid: List[List[str]], r: int, c: int) -> int:
    rows = len(grid)
    cols = len(grid[0]) if rows > 0 else 0
    cnt = 0
    for dr, dc in dirs:
        nr, nc = r + dr, c + dc
        if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == '@':
            cnt += 1
    return cnt

# Part Two: iterative until no accessible rolls remain
grid = [list(line) for line in lines]
result2 = 0
while True:
    to_remove: List[tuple] = []
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] != '@':
                continue
            if adj_count(grid, r, c) < 4:
                to_remove.append((r, c))
    if not to_remove:
        break
    for r, c in to_remove:
        grid[r][c] = '.'
    result2 += len(to_remove)

print(f"first: {result1}")
print(f"second: {result2}")
