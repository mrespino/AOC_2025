banks = File.readlines("input_03.txt").map(&:strip)
result1 = 0
result2 = 0

# First
banks.each do |bank|
  max_joltage = 0
  bank.length.times do |i|
    (i+1...bank.length).each do |j|
      joltage = (bank[i] + bank[j]).to_i
      max_joltage = [max_joltage, joltage].max
    end
  end
  result1 += max_joltage
end

# Second
banks.each do |bank|
  selected = []
  start_pos = 0
  needed = 12
  12.times do
    end_search = bank.length - needed + 1
    best_pos = start_pos
    best_digit = bank[start_pos]
    (start_pos...end_search).each do |pos|
      if bank[pos] > best_digit
        best_digit = bank[pos]
        best_pos = pos
      end
    end
    selected << best_digit
    start_pos = best_pos + 1
    needed -= 1
  end
  joltage_str = selected.join
  joltage = joltage_str.to_i
  result2 += joltage
end

puts "first: #{result1}"
puts "second: #{result2}"
