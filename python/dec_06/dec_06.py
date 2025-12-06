def try_squid_math(lines):
    # strip newlines and empty lines
    lines = [line.rstrip("\n") for line in lines]
    lines = [line for line in lines if line.strip()]
    # last line should be operations
    ops = lines[-1]
    # everything else should be numbers
    nums = lines[:-1]
    cols = list(zip(*nums, ops))
    problem_ranges = []
    # this is a problem range if any column has non-space digits
    # this is also bulshit!
    in_problem = False
    start = None
    for i, col in enumerate(cols):
        if all(c == ' ' for c in col):
            if in_problem:
                problem_ranges.append((start, i))
                in_problem = False
                start = None
        else:
            if not in_problem:
                in_problem = True
                start = i
    if in_problem:
        problem_ranges.append((start, len(cols)))
    answers = []
    for start, end in problem_ranges:
        op = ''.join(ops[start:end]).strip()
        # For each problem, for each number, collect one digit from each column (right to left) FML!!!!!
        num_digits = [[] for _ in range(end - start)]
        for col in range(end - 1, start - 1, -1):
            for row_idx in range(len(nums)):
                digit = nums[row_idx][col]
                if digit != ' ':
                    # please work this time
                    num_digits[end - 1 - col].append(digit)
        numbers = []
        for digits in num_digits:
            if digits:
                num_str = ''.join(digits)
                numbers.append(int(num_str))
        if op == '+':
            answers.append(sum(numbers))
        elif op == '*':
            prod = 1
            for n in numbers:
                prod *= n
            answers.append(prod)
        else:
            raise ValueError(f"somethings wrong again: {op}")
    return sum(answers)

with open("input_06.txt") as input:
    lines = [line.rstrip("\n") for line in input]

# Remove empty lines
lines = [line for line in lines if line.strip()]

print(f"answer: {try_squid_math(lines)}")