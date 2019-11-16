# x

@n, @m = gets.chomp.split(" ").map(&:to_i)
@lake = []
@n.times do
  @lake << gets.chomp.split("")
end

def dfs(x, y)
  @lake[x][y] = "."

  (-1..1).each do |dx|
    (-1..1).each do |dy|
      nx = x + dx
      ny = y + dy
      if (0 <= nx && nx < @n) && (0 <= ny && ny < @m) && @lake[nx][ny] == "W"
        dfs(nx, ny)
      end
    end
  end
end

ans = 0

(0..@n-1).each do |i|
  (0..@m-1).each do |j|
    if @lake[i][j] == "W"
      dfs(i, j)
      ans += 1
    end
  end
end

puts ans

# 10 12
# W........WW.
# .WWW.....WWW
# ....WW...WW.
# .........WW.
# .........W..
# ..W......W..
# .W.W.....WW.
# W.W.W.....W.
# .W.W......W.
# ..W.......W.