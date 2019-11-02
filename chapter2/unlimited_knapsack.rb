n = gets.chomp.to_i
wv = []
n.times do 
  wv << gets.chomp.split(" ").map(&:to_i)
end
W = gets.chomp.to_i

# (n+1) * (W+1)
dp = []
(n+1).times do
  dp << [0] * (W + 1)
end

(1..n).each do |i|
  (1..W).each do |j|
    if wv[i-1][0] > j # 選べない場合
      dp[i][j] = dp[i-1][j]
    else
      dp[i][j] = [dp[i-1][j], dp[i][j-wv[i-1][0]] + wv[i-1][1]].max
    end
  end
end

puts dp[n][W]


# 3
# 3 4
# 4 5
# 2 3
# 7

# 10