with open("input_03.txt", "r") as input:
    banks = [line.strip() for line in input]

result1 = 0
result2 = 0

# First
for bank in banks:
    max_joltage = 0
    
    for i in range(len(bank)):
        for j in range(i + 1, len(bank)):
            joltage = int(bank[i] + bank[j])
            max_joltage = max(max_joltage, joltage)
    
    result1 += max_joltage

# Second
for bank in banks:
    selected = []
    start_pos = 0
    needed = 12

    # Find the largest possible joltage
    for step in range(12):
        end_search = len(bank) - needed + 1
        best_pos = start_pos
        best_digit = bank[start_pos]
        
        # Find the largest digit in range
        for pos in range(start_pos, end_search):
            if bank[pos] > best_digit:
                best_digit = bank[pos]
                best_pos = pos
        
        selected.append(best_digit)
        start_pos = best_pos + 1
        needed -= 1
    
    joltage_str = ''.join(selected)
    joltage = int(joltage_str)
    result2 += joltage

print(f"first: {result1}")
print(f"second: {result2}")