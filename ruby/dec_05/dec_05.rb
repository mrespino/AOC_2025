# Advent of Code 2025 Day 5 - Cafeteria Inventory (Ruby)

sections = File.read("input_05.txt").strip.split("\n\n")

# Parse ranges
fresh_ranges = sections[0].lines.map do |line|
  start, end_ = line.strip.split("-").map(&:to_i)
  [start, end_]
end

# Parse IDs
available_ids = sections[1].lines.map(&:strip).reject(&:empty?).map(&:to_i)

# Part I: Count fresh available IDs
result1 = available_ids.count do |id|
  fresh_ranges.any? { |start, end_| start <= id && id <= end_ }
end

# Part II: ids in any range
merged = fresh_ranges.sort_by(&:first)
merged_ranges = []
merged.each do |start, end_|
  if merged_ranges.empty? || start > merged_ranges[-1][1] + 1
    merged_ranges << [start, end_]
  else
    merged_ranges[-1][1] = [merged_ranges[-1][1], end_].max
  end
end
result2 = merged_ranges.sum { |start, end_| end_ - start + 1 }

puts "First: #{result1}"
puts "Second: #{result2}"
