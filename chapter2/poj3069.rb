n = gets.chomp.to_i
r = gets.chomp.to_i
x = gets.chomp.split(" ").map(&:to_i).sort

ans = 0

left = 0
right = 0
while right < n do
  # 印を付ける点を探す
  while (right < n) && (x[right] - x[left] <= r) do
    right += 1
  end

  # 印を付ける点
  right -= 1
  left = right
  ans += 1
  
  # 印の範囲に収まらなくなる点までrightを増やす
  while (right < n) && (x[right] - x[left] <= r) do
    right += 1
  end
  left = right
end

puts ans