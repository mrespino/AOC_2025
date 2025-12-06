# Advent of Code 2025 Day 5 - Cafeteria Inventory

with open("input_05.txt", "r") as f:
    sections = f.read().strip().split("\n\n")

# Parse ranges
fresh_ranges = []
for line in sections[0].splitlines():
    start, end = map(int, line.split("-"))
    fresh_ranges.append((start, end))

# Parse IDs
available_ids = [int(line) for line in sections[1].splitlines() if line.strip()]

# Part I: Count fresh available IDs
result1 = 0
for id_ in available_ids:
    if any(start <= id_ <= end for start, end in fresh_ranges):
        result1 += 1


# Part II: ids in any range
def part_two(intervals):
    intervals = sorted(intervals)
    merged = []
    for start, end in intervals:
        if not merged or start > merged[-1][1] + 1:
            merged.append([start, end])
        else:
            merged[-1][1] = max(merged[-1][1], end)
    return merged

merged_intervals = part_two(fresh_ranges)
result2 = sum(end - start + 1 for start, end in merged_intervals)

print(f"First: {result1}")
print(f"Second: {result2}")
