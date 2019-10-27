# x

n = gets.chomp.to_i
wv = []
n.times do
  wv << gets.chomp.split(" ").map(&:to_i)
end
w = gets.chomp.to_i

dp = []
(n+1).times do
  dp << [0] * (w + 1)
end

(0..n-1).each do |i|
  i = n - 1 - i
  (0..w).each do |j|
    if j < wv[i][0]
      dp[i][j] = dp[i+1][j]
    else
      dp[i][j] = [dp[i+1][j], dp[i+1][j-wv[i][0]] + wv[i][1]].max
    end
  end
end

puts dp[0][w]

# 4
# 2 3
# 1 2
# 3 4
# 2 2
# 5