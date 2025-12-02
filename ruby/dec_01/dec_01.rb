def first
  lines = File.readlines("input_01.txt").map(&:strip)

  result = 0
  dial = 50

  lines.each do |line|
    direction = line[0]
    number = line[1..].to_i

    dial += direction == "R" ? number : -number
    dial %= 100

    result += 1 if dial == 0
  end

  puts "first: #{result}"
end

def second
  lines = File.readlines("input_01.txt").map(&:strip)

  result = 0
  dial = 50

  lines.each do |line|
    direction = line[0]
    number = line[1..].to_i

    step = direction == "R" ? 1 : -1

    number.times do
      dial = (dial + step) % 100
      result += 1 if dial == 0
    end
  end

  puts "second: #{result}"
end

first
second