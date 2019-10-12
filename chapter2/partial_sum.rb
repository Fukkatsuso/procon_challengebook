@n = gets.chomp.to_i
@a = gets.chomp.split(" ").map(&:to_i)
@k = gets.chomp.to_i

def dfs(i, sum)
  return sum == @k if i == @n
  return true if dfs(i + 1, sum)
  return true if dfs(i + 1, sum + @a[i])
  return false
end

if dfs(0, 0)
  puts "Yes"
else
  puts "No"
end