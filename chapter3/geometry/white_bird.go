package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
	g              = 9.8
	eps            = 1e-6
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N, V := getInt(), getInt()
	X, Y := getInt(), getInt()
	L, B, R, T := make([]int, N), make([]int, N), make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		L[i], B[i], R[i], T[i] = getInt(), getInt(), getInt(), getInt()
	}

	// 点(qx,qy)を通るように打ち出したとき，豚に卵をぶつけられるか判定
	canHit := func(qx, qy float64) bool {
		// 初速のx方向成分をvx，y方向成分をvyとし，(qx,qy)を通る時刻をtとしたときの連立方程式
		// vx^2 + vy^2 = V^2
		// vx*t = qx
		// vy*t - 1/2 g t^2 = qy
		// を解く．解は2通りある（点(qx,qy)を鳥が上昇中に通るか，下降中に通るか）
		a, b, c := g*g/4, g*qy-float64(V*V), qx*qx+qy*qy
		D := b*b - 4*a*c // 判別式

		if -eps < D && D < 0 {
			D = 0
		}
		if D < 0 {
			return false
		}

		for d := float64(-1); d <= 1; d += 2 { // 連立方程式の2つの解を試すループ
			t2 := (-b + d*math.Sqrt(D)) / (2 * a)
			if t2 <= 0 {
				continue
			}
			t := math.Sqrt(t2)
			vx, vy := qx/t, (qy+g*t*t/2)/t

			// 豚より上を通過できるか
			yt := pointY(vy, float64(X)/vx)
			if yt < float64(Y)-eps {
				continue
			}

			ok := true
			for i := 0; i < N; i++ {
				if L[i] >= X {
					continue
				}
				// 豚の真上まできたときに，間に障害物がないか
				if R[i] == X && Y <= T[i] && float64(B[i]) <= yt {
					ok = false
				}
				// 途中で障害物にぶつからないか
				yL := cmp(float64(B[i]), float64(T[i]), pointY(vy, float64(L[i])/vx)) // 障害物iの左端の辺
				yR := cmp(float64(B[i]), float64(T[i]), pointY(vy, float64(R[i])/vx)) // 障害物iの右端の辺
				xH := cmp(float64(L[i]), float64(R[i]), vx*(vy/g))                    // 鳥の最高地点
				yH := cmp(float64(B[i]), float64(T[i]), pointY(vy, vy/g))
				if xH == 0 && yH >= 0 && yL < 0 { // 鳥の最高地点の下に障害物iがあり，障害物iの左辺より下を通過
					ok = false
				}
				if yL*yR <= 0 { // 障害物iの左右の辺どちらかに衝突してしまう
					ok = false
				}
			}
			if ok {
				return true
			}
		}

		return false
	}

	// 豚より右にある障害物を縮めておく
	for i := 0; i < N; i++ {
		if R[i] > X {
			R[i] = X
		}
	}

	ok := canHit(float64(X), float64(Y)) // 直接ぶつける場合
	for i := 0; i < N; i++ {
		ok = ok ||
			canHit(float64(L[i]), float64(T[i])) || // 左上の角を通る場合
			canHit(float64(R[i]), float64(T[i])) // 右上の角を通る場合
	}
	if ok {
		puts("Yes")
	} else {
		puts("No")
	}
}

// 鉛直上向きに初速vyで打ち出した際の，t秒後の位置
func pointY(vy, t float64) float64 {
	return vy*t - g*t*t/2
}

// xのlbとubに対する相対的な位置
func cmp(lb, ub, x float64) int {
	if x < lb+eps {
		return -1
	}
	if x > ub-eps {
		return 1
	}
	return 0
}
