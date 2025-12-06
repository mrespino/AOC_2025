lines = File.readlines("input_04.txt").map(&:strip).reject(&:empty?)
DIRECTIONS = [
  [-1, -1], [-1, 0], [-1, 1], [0, -1],
  [0, 1], [1, -1], [1, 0], [1, 1]
]
rows = lines.size
cols = rows > 0 ? lines[0].size : 0
result1 = 0

# Part One: count accessible rolls
rows.times do |r|
  cols.times do |c|
    next unless lines[r][c] == '@'
    adj = 0
  DIRECTIONS.each do |dr, dc|
      nr, nc = r + dr, c + dc
      if nr.between?(0, rows-1) && nc.between?(0, cols-1) && lines[nr][nc] == '@'
        adj += 1
        break if adj >= 4
      end
    end
    result1 += 1 if adj < 4
  end
end

def adj_count(grid, r, c)
  rows = grid.size
  cols = rows > 0 ? grid[0].size : 0
  cnt = 0
  DIRECTIONS.each do |dr, dc|
    nr, nc = r + dr, c + dc
    cnt += 1 if nr.between?(0, rows-1) && nc.between?(0, cols-1) && grid[nr][nc] == '@'
  end
  cnt
end

grid = lines.map(&:chars)
result2 = 0
loop do
  to_remove = []
  rows.times do |r|
    cols.times do |c|
      next unless grid[r][c] == '@'
      to_remove << [r, c] if adj_count(grid, r, c) < 4
    end
  end
  break if to_remove.empty?
  to_remove.each { |r, c| grid[r][c] = '.' }
  result2 += to_remove.size
end

puts "first: #{result1}"
puts "second: #{result2}"
