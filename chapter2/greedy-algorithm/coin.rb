c = gets.chomp.split(" ").map(&:to_i)
a = gets.chomp.to_i

n = 6
price = [1, 5, 10, 50, 100, 500]
ans = 0 # 使うコインの総数
sum = 0 # 現時点で選んだコインの総額
(0..n - 1).each do |i|
  break if sum == a

  i = (n - 1) - i
  if a - sum >= price[i] && c[i] > 0
    ans += 1
    sum += price[i]
    c[i] -= 1
  end
end

puts ans

# 3 2 1 3 0 2
# 620