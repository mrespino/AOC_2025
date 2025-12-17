

lines = File.readlines("input_07.txt").map(&:rstrip)
rows = lines.size
cols = lines[0].size

# Find S (start column)
start_col = lines[0].index('S')
raise "No start S found" unless start_col






# DP table: ways[row][col] = number of ways to reach (row, col)
ways = Array.new(rows+1) { Array.new(cols, 0) }
ways[0][start_col] = 1

(0...rows).each do |row|
  (0...cols).each do |col|
    w = ways[row][col]
    next if w == 0
    cell = lines[row][col]
    if cell == '.' || cell == 'S'
      ways[row+1][col] += w
    elsif cell == '^'
      ways[row+1][col-1] += w if col > 0
      ways[row+1][col+1] += w if col < cols-1
    end
  end
end

# Sum all ways to exit the grid from the last row
timeline_count = ways[rows].sum
puts "answer2: #{timeline_count}"
