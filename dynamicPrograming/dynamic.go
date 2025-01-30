package main

import "fmt"

func main() {
	n := 10
	memo := make(map[int]int)
	result := fibonacci(n, memo)
	fmt.Printf("Number of ways to reach step %d is %d\n", n, result)

	s1 := "The quick brown fox jumps over the lazy dog"
	s2 := "The brown fox jumps over the quick dog"
	result = longestCommonSubsequence(s1, s2)
	fmt.Printf("Length of LCS is %d\n", result)
}

/*
	Dynamic fibonacci
	จำลองสำหรับการเติบโตของการออมในแต่ละเดือน แต่ละเดือนจำนวนเงินที่ออมเพิ่มขึ้นเป็นผลรวมของสองเดือนก่อนหน้า ผู้ใช้งานสามารถระบุจำนวนเดือนที่ต้องการคำนวณ

	***การวิเคราะห์ประสิทธิภาพ***
		- การทำงานในเวลา O(n)

	Memoization: ใช้ recursion + การเก็บผลลัพธ์ที่คำนวณแล้วใน map เพื่อหลีกเลี่ยงการคำนวณซ้ำ
	ข้อดี:
		- ประหยัดหน่วยความจำมากกว่าในบางกรณี เพราะเก็บเฉพาะค่าที่คำนวณแล้วใน map
		- ใช้งานสะดวกเมื่อมีการคำนวณค่า Fibonacci หลายค่าในลำดับที่ไม่ต่อเนื่อง
	ข้อเสีย:
		- มี overhead จากการเรียกฟังก์ชันแบบ recursive หลายครั้ง
		- ประสิทธิภาพอาจลดลงถ้าความลึกของ recursive สูง
*/

func fibonacci(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if _, ok := memo[n]; !ok {
		memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
	}

	return memo[n]
}

/*
	Dynamic Longest Common Subsequence (LCS)
	คุณได้รับข้อความสองชุดที่เป็นประโยค และต้องการหาความยาวของ Longest Common Subsequence (LCS) ของคำในทั้งสองข้อความ โดยไม่ต้องคำนึงถึงลำดับของคำที่ตรงกัน แต่ต้องรักษาคำเหล่านั้นในลำดับที่ปรากฏในทั้งสองข้อความ

	***การวิเคราะห์ประสิทธิภาพ***
		- การทำงานในเวลา O(m*n)
*/

func longestCommonSubsequence(s1, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1) //เก็บข้อมูลรอง LCS
	for i := range dp {
		dp[i] = make([]int, n+1) //เก็บข้อมูล LCS
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] { // คำที่ตำแหน่งนี้เหมือนกันในทั้งสองข้อความ
				dp[i][j] = dp[i-1][j-1] + 1 // LCS จะเพิ่มขึ้น 1 จากค่าก่อนหน้า
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) // ตัวอักษรไม่เท่ากัน เลือกค่าที่มากที่สุดระหว่างการข้ามตัวอักษรใน s1 หรือข้ามตัวอักษรใน s2
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
