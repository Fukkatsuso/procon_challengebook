# x

n = gets.chomp.to_i
s = gets.chomp.split(" ").map(&:to_i)
t = gets.chomp.split(" ").map(&:to_i)

jobs = []
(0..n-1).each do |i|
  jobs << [s[i], t[i]]
end
# 早く終わる順にソート
jobs.sort_by {|x| x[1]}

last = 0
ans = 0

jobs.each do |job|
  # job開始時刻、終了時刻
  js, je = job
  if last < js
    ans += 1
    last = je
  end
end

puts ans

# 5
# 1 2 4 6 8
# 3 5 7 9 10