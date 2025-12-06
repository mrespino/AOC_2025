#refactored as ruby this time
def try_cephalopod_math(lines)
  lines = lines.map(&:rstrip).reject { |line| line.strip.empty? }
  ops = lines[-1]
  nums = lines[0...-1]
  maxlen = ([ops.size] + nums.map(&:size)).max
  padded = nums.map { |row| row.ljust(maxlen) } + [ops.ljust(maxlen)]
  cols = padded.map(&:chars).transpose.map(&:join)
  problem_ranges = []
  in_problem = false
  start = nil
  cols.each_with_index do |col, i|
    if col.chars.all? { |c| c == ' ' }
      if in_problem
        problem_ranges << [start, i]
        in_problem = false
        start = nil
      end
    else
      unless in_problem
        in_problem = true
        start = i
      end
    end
  end
  problem_ranges << [start, cols.size] if in_problem
  answers = []
  problem_ranges.each do |start, end_idx|
    op = ops[start...end_idx].strip
    num_digits = Array.new(end_idx - start) { [] }
    (end_idx - 1).downto(start) do |col|
      nums.each_with_index do |row, row_idx|
        digit = row[col]
        num_digits[end_idx - 1 - col] << digit if digit != ' '
      end
    end
    numbers = num_digits.map { |digits| digits.empty? ? nil : digits.join.to_i }.compact
    if op == '+'
      answers << numbers.sum
    elsif op == '*'
      answers << numbers.inject(1) { |prod, n| prod * n }
    else
      raise "this broke again: #{op}"
    end
  end
  answers.sum
end

lines = File.readlines("input_06.txt").map(&:rstrip)
lines = lines.reject { |line| line.strip.empty? }
puts "answer: #{try_cephalopod_math(lines)}"
