data = File.read("input_02.txt").strip
result1 = 0
result2 = 0

# iterate through ranges
data.split(",").each do |line|
  begin_s, end_s = line.split("-")
  (begin_s.to_i..end_s.to_i).each do |i|
    token = i.to_s
    # if even
    if token.length.even?
      mid = token.length / 2
      # if symmetrical
      if token[0...mid] == token[mid..-1]
        result1 += i
      end
    end
    # check for repeated pattern
    (1..(token.length / 2)).each do |pattern_len|
      if token.length % pattern_len == 0
        pattern = token[0, pattern_len]
        if pattern * (token.length / pattern_len) == token
          result2 += i
          break
        end
      end
    end
  end
end

puts "first: #{result1}"
puts "second: #{result2}"
