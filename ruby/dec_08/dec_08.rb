class UnionFind
	def initialize(n)
		@parent = Array.new(n) { |i| i }
		@size = Array.new(n, 1)
	end
	def find(x)
		@parent[x] = find(@parent[x]) if @parent[x] != x
		@parent[x]
	end
	def union(x, y)
		xr, yr = find(x), find(y)
		return if xr == yr
		if @size[xr] < @size[yr]
			@parent[xr] = yr
			@size[yr] += @size[xr]
		else
			@parent[yr] = xr
			@size[xr] += @size[yr]
		end
	end
	def sizes
		roots = @parent.map { |x| find(x) }
		counts = Hash.new(0)
		roots.each { |r| counts[r] += 1 }
		counts.values
	end
end

def first
	boxes = File.readlines("input_08.txt").map { |line| line.strip.split(",").map(&:to_i) }
	n = boxes.size
	pairs = []
	(0...n).each do |i|
		(i+1...n).each do |j|
			a, b = boxes[i], boxes[j]
			dist = Math.sqrt((a[0]-b[0])**2 + (a[1]-b[1])**2 + (a[2]-b[2])**2)
			pairs << [dist, i, j]
		end
	end
	pairs.sort_by!(&:first)
	uf = UnionFind.new(n)
	(0...1000).each do |idx|
		_, i, j = pairs[idx]
		uf.union(i, j)
	end
	sizes = uf.sizes.sort.reverse
	sizes << 1 while sizes.size < 3
	result = sizes[0] * sizes[1] * sizes[2]
	puts "first: #{result} (sizes: #{sizes[0]}, #{sizes[1]}, #{sizes[2]})"
end

def second
	boxes = File.readlines("input_08.txt").map { |line| line.strip.split(",").map(&:to_i) }
	n = boxes.size
	pairs = []
	(0...n).each do |i|
		(i+1...n).each do |j|
			a, b = boxes[i], boxes[j]
			dist = Math.sqrt((a[0]-b[0])**2 + (a[1]-b[1])**2 + (a[2]-b[2])**2)
			pairs << [dist, i, j]
		end
	end
	pairs.sort_by!(&:first)
	uf = UnionFind.new(n)
	last_i = last_j = nil
	pairs.each do |_, i, j|
		if uf.find(i) != uf.find(j)
			uf.union(i, j)
			last_i, last_j = i, j
			if (0...n).all? { |x| uf.find(x) == uf.find(0) }
				break
			end
		end
	end
	if last_i && last_j
		result = boxes[last_i][0] * boxes[last_j][0]
		puts "second: #{result} (X: #{boxes[last_i][0]}, #{boxes[last_j][0]})"
	else
		puts "second: not found"
	end
end

first
second
