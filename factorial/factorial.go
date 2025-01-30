package main

import (
	"fmt"
)

/*
	***การวิเคราะห์ประสิทธิภาพ***
		- การทำงานในเวลา O(n)
*/

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {

	// ในเกมสุ่มจับลูกบอลจากกล่องที่มีลูกบอลทั้งหมด 8 ลูก ลูกบอลแต่ละลูกมีหมายเลขกำกับตั้งแต่ 1 ถึง 8 ถ้าสุ่มจับลูกบอลออกมา 3 ลูกโดยไม่สนใจลำดับ คุณต้องเขียนโปรแกรมเพื่อคำนวณจำนวนวิธีที่สามารถสุ่มจับลูกบอล 3 ลูกได้
	n, r := 8, 3
	fmt.Printf("Number of ways to choose %d balls from %d: %d\n", r, n, combination(n, r))

	// มีสมาชิกทั้งหมด 12 คนในชมรม และต้องการเลือกคณะกรรมการ 5 คนเพื่อทำหน้าที่บริหารชมรม คุณต้องเขียนโปรแกรมเพื่อคำนวณจำนวนวิธีที่สามารถเลือกคณะกรรมการชุดนี้ได้
	n, r = 12, 5
	fmt.Printf("Number of ways to choose %d balls from %d: %d\n", r, n, combination(n, r))

	// ในห้องประชุมมีที่นั่งทั้งหมด 10 ที่ และต้องการจัดที่นั่งสำหรับสมาชิก 10 คนให้ไม่ซ้ำกัน คุณต้องเขียนโปรแกรมเพื่อคำนวณจำนวนวิธีการจัดที่นั่งทั้งหมด
	n = 10
	fmt.Printf("Number of ways to arrange %d members: %d\n", n, factorial(n))

}
