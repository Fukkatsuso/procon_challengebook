# 本と違う回答。正当性と実行速度要検証

n = gets.chomp.to_i
l = gets.chomp.split(" ").map(&:to_i).sort

ans = 0

sum = l.sum
while l.size > 1 do
  ans += sum
  sum -= l.pop
end

puts ans