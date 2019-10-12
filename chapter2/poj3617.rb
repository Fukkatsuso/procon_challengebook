# x

n = gets.chomp.to_i
s = gets.chomp.split("")
t = ""

head = 0
tail = n - 1

while head <= tail do
  left = false
  (0..tail-head).each do |i|
    if s[head + i] < s[tail - i]
      left = true
      break
    elsif s[head + i] > s[head - i]
      left = false
      break
    end
  end
  if left
    t << s[head]
    head += 1
  else
    t << s[tail]
    tail -= 1
  end
end

puts t

# 6
# ACDBCB