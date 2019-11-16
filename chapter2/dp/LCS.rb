n, m = gets.chomp.split(" ").map(&:to_i)
s = gets.chomp.split("")
t = gets.chomp.split("")

# (n*1) * (m+1) のテーブル
lcs = []
(n+1).times do
  lcs << [0] * (m + 1)
end

(1..n).each do |i|
  (1..m).each do |j|
    max = [lcs[i-1][j],
          lcs[i][j-1],
          s[i-1] == t[j-1] ? lcs[i-1][j-1] + 1 : lcs[i-1][j-1]].max
    lcs[i][j] = max
  end
end

puts lcs[n][m]


# 4 4
# abcd
# becd

# 3