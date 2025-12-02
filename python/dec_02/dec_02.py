with open("input_02.txt", "r") as input:
    data = input.read().strip()
result1 = 0
result2 = 0

# iterate through ranges
for line in data.split(","):
    begin, end = line.split("-")
    # iterate through range
    for i in range(int(begin), int(end) + 1):
        token = str(i)
        #if even
        if len(token) % 2 == 0:
            mid = len(token) // 2
            # if symetrical
            if token[:mid] == token[mid:]:
                result1 += i
        # check for repeated pattern
        for pattern_len in range(1, len(token) // 2 + 1):
          # if divisible  
            if len(token) % pattern_len == 0:
                # check if token is a repeated pattern
                if token == token[0:pattern_len] * (len(token) // pattern_len):
                    result2 += i
                    break

print(f"first: {result1}")
print(f"second: {result2}")