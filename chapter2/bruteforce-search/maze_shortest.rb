# x

n, m = gets.chomp.split(" ").map(&:to_i)
maze = []
n.times do
  maze << gets.chomp.split("")
end

# スタート地点
sx = 0
sy = 0
# ゴール地点
gx = 0
gy = 0
# queue
que = []
inf = n * m
d = [] # 各点までの最短距離
# 移動方向のベクトル
dx = [1, 0, -1, 0]
dy = [0, 1, 0, -1]

# スタート地点を探してキューに入れる
# ゴール地点も探す
(0..n-1).each do |i|
  (0..m-1).each do |j|
    if maze[i][j] == "S"
      sx = i
      sy = j
    elsif maze[i][j] == "G"
      gx = i
      gy = j
    end
  end
  d << [inf] * m
end
que << [sx, sy]
d[sx][sy] = 0

while que.size > 0 do
  p = que.shift
  break if p[0] == gx && p[1] == gy # 到着
  
  (0..3).each do |i|
    nx = p[0] + dx[i]
    ny = p[1] + dy[i]

    if (0 <= nx && nx < n) && (0 <= ny && ny < m) && maze[nx][ny] != "#" && d[nx][ny] == inf
      que << [nx, ny]
      d[nx][ny] = d[p[0]][p[1]] + 1
    end
  end
end

puts d[gx][gy]

# 10 10
# #S######.#
# ......#..#
# .#.##.##.#
# .#........
# ##.##.####
# ....#....#
# .#######.#
# ....#.....
# .####.###.
# ....#...G#