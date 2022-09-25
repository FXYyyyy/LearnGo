package main

import "fmt"

func main() {
	mat := [][]int{
		{1,2,3,4,5},
		{6,7,8,9,10},
		{11,12,13,14,15},
	}
	ret := findDiagonalOrder(mat)
	fmt.Println("ret = ", ret)
}

func findDiagonalOrder(mat [][]int) []int {
	const Up = 1
	const Down = 2
	ret := make([]int, 0)
	rows := len(mat)
	if rows == 0 {
		return ret
	}
	if rows == 1 {
		return mat[0]
	}
	columns := len(mat[0])
	if columns == 1 {
		for i:=0;i<rows;i++ {
			ret = append(ret, mat[i][0])
		}
		return ret
	}

	flag := Up
	i, j,count := 0, 0, 0
	for {
		count++
		if count == 15 {
			//break
		}
		//判断是不是最后一个数
		if i == rows - 1 && j == columns - 1 {
			ret = append(ret, mat[i][j])
			break
		}
		//判断是否出去了
		isOut := checkIsOut(i,j,rows,columns)
		if isOut {
			if flag == Up {
				i++
				temp := checkIsOut(i, j, rows, columns)
				if temp {
					j--
					i++
				}
				flag = Down
			} else {
				j++
				temp := checkIsOut(i,j,rows, columns)
				if temp {
					i--
					j++
				}
				flag = Up
			}
			continue
		}

		//将当前的数放入遍历的结果
		ret = append(ret, mat[i][j])

		//根据往上往下的标识挪坐标
		if flag == Up {
			i--
			j++
		} else {
			i++
			j--
		}
	}

	return ret
}

//判断是不是已经出去了
func checkIsOut(i,j,row, column int) bool {
	if i == -1 || j == -1 || i == row || j == column {
		return true
	}
	return false
}