package main

import (
	"fmt"
)

func main() {
	var n, c int
	fmt.Scan(&n, &c)
	x := make([]int, n)
	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &v[i])
	}
	fmt.Printf("%d\n", findMaxCalorie(c, x, v))
}

func findMaxCalorie(c int, x, v []int) int {
	ls, ld := findLeftSafePoints(x, v)
	rs, rd := findRightSafePoints(c, x, v)
	lMax := calcLeftMaxCalorie(c, ls, rd)
	rMax := calcRightMaxCalorie(c, rs, ld)
	if rMax > lMax {
		return rMax
	}
	return lMax
}

func calcLeftMaxCalorie(c int, singlePoses, doublePoses []gain) (max int) {
	var curSingleIdx, curDoubleIdx, curCal int
	for i := 0; i < len(doublePoses); i++ {
		curCal += doublePoses[i].totalGain
	}
	max = curCal
	curDoubleIdx = len(doublePoses)
	for singleSafePontsNum := 1; singleSafePontsNum <= len(singlePoses); singleSafePontsNum++ {
		for i := curSingleIdx; i < singleSafePontsNum; i++ {
			curCal += singlePoses[i].totalGain
			curSingleIdx = i
		}
		curSinglePos := singlePoses[curSingleIdx].pos
		for i := curDoubleIdx - 1; i > -1 && doublePoses[i].pos <= curSinglePos; i-- {
			curCal -= doublePoses[i].totalGain
			curDoubleIdx = i
		}
		curSingleIdx++
		if curCal > max {
			max = curCal
		}
	}
	return
}

func calcRightMaxCalorie(c int, singlePoses, doublePoses []gain) (max int) {
	var curSingleIdx, curDoubleIdx, curCal int
	for i := 0; i < len(doublePoses); i++ {
		curCal += doublePoses[i].totalGain
	}
	max = curCal
	curDoubleIdx = len(doublePoses)
	for singleSafePontsNum := 1; singleSafePontsNum <= len(singlePoses); singleSafePontsNum++ {
		for i := curSingleIdx; i < singleSafePontsNum; i++ {
			curCal += singlePoses[i].totalGain
			curSingleIdx = i
		}
		curSinglePos := singlePoses[curSingleIdx].pos
		for i := curDoubleIdx - 1; i > -1 && doublePoses[i].pos >= curSinglePos; i-- {
			curCal -= doublePoses[i].totalGain
			curDoubleIdx = i
		}
		curSingleIdx++
		if curCal > max {
			max = curCal
		}
	}
	return
}

type gain struct {
	pos       int
	totalGain int
}

func findLeftSafePoints(x, v []int) (singles, doubles []gain) {
	sCur, dCur := 0, 0
	sCalorie, dCalorie := 0, 0
	for i := range x {
		if sGain := sCalorie + v[i] - (x[i] - sCur); sGain > 0 {
			sCur = x[i]
			singles = append(singles, gain{
				pos:       x[i],
				totalGain: sGain,
			})
			sCalorie = 0
		} else {
			sCalorie += v[i]
		}
		if dGain := dCalorie + v[i] - 2*(x[i]-dCur); dGain > 0 {
			dCur = x[i]
			doubles = append(doubles, gain{
				pos:       x[i],
				totalGain: dGain,
			})
			dCalorie = 0
		} else {
			dCalorie += v[i]
		}
	}
	return
}
func findRightSafePoints(c int, x, v []int) (singles, doubles []gain) {
	// TODO
	sCur, dCur := c, c
	sCalorie, dCalorie := 0, 0
	for i := len(x); i > 0; i-- {
		if sGain := sCalorie + v[i-1] - (sCur - x[i-1]); sGain > 0 {
			singles = append(singles, gain{
				pos:       x[i-1],
				totalGain: sGain,
			})
			sCur = x[i-1]
			sCalorie = 0
		} else {
			sCalorie += v[i-1]
		}
		if dGain := dCalorie + v[i-1] - 2*(dCur-x[i-1]); dGain > 0 {
			dCur = x[i-1]
			doubles = append(doubles, gain{
				pos:       x[i-1],
				totalGain: dGain,
			})
			dCalorie = 0
		} else {
			dCalorie += v[i-1]
		}
	}
	return
}
