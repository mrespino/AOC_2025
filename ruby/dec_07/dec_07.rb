def first
  lines = File.readlines("input_07.txt").map(&:rstrip)
  rows = lines.size
  cols = lines[0].size

  # Find S (start column)
  start_col = lines[0].index('S')
  raise "No start S found" unless start_col

  queue = [[0, start_col]]
  split_count = 0
  visited = {}

  until queue.empty?
    row, col = queue.shift
    next if row >= rows
    next if visited[[row, col]]
    visited[[row, col]] = true
    cell = lines[row][col]
    if cell == '.' || cell == 'S'
      queue << [row + 1, col]
    elsif cell == '^'
      split_count += 1
      queue << [row + 1, col - 1] if col > 0
      queue << [row + 1, col + 1] if col < cols - 1
    end
  end

  puts "first: #{split_count}"
end

def second
  lines = File.readlines("input_07.txt").map(&:rstrip)
  rows = lines.size
  cols = lines[0].size

  # Find S (start column)
  start_col = lines[0].index('S')
  raise "No start S found" unless start_col

  ways = Array.new(rows + 1) { Array.new(cols, 0) }
  ways[0][start_col] = 1

  rows.times do |row|
    cols.times do |col|
      w = ways[row][col]
      next if w == 0
      cell = lines[row][col]
      if cell == '.' || cell == 'S'
        ways[row + 1][col] += w
      elsif cell == '^'
        ways[row + 1][col - 1] += w if col > 0
        ways[row + 1][col + 1] += w if col < cols - 1
      end
    end
  end

  timeline_count = ways[rows].sum
  puts "second: #{timeline_count}"
end

first
second
