package main

import "fmt"

/*
	คุณได้รับรายชื่อผู้เข้าร่วมกิจกรรมและคะแนนที่พวกเขาได้รับจากการทำกิจกรรมต่างๆ ในรูปแบบของอาร์เรย์ โดยแต่ละตำแหน่งในอาร์เรย์แทนคะแนนของผู้เข้าร่วมคนหนึ่ง ให้คุณเขียนฟังก์ชันที่ทำการเรียงลำดับคะแนนจากน้อยไปมากโดยใช้ Bubble Sort และแสดงลำดับคะแนนที่เรียงแล้วออกมา

	ตัวอย่าง:
		อินพุต: arr := []int{45, 78, 12, 34, 89, 20}
		เอาต์พุต: Sorted array: [12, 20, 34, 45, 78, 89]
		ข้อกำหนดเพิ่มเติม:
			- คะแนนแต่ละคะแนนเป็นจำนวนเต็มบวก และคะแนนอาจซ้ำกันได้

	***การวิเคราะห์ประสิทธิภาพ***
		- Time Complexity: O(n²) ในกรณีที่แย่ที่สุดและกรณีเฉลี่ย เพราะต้องเปรียบเทียบและสลับตำแหน่งหลายครั้ง
		- Space Complexity: O(1) เนื่องจากไม่ใช้พื้นที่เพิ่มเติมสำหรับอาร์เรย์อื่น
*/

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ { // เปรียบเทียบและสลับตำแหน่งของค่าที่อยู่ติดกันในอาร์เรย์ ค่า j จะเริ่มจาก 0 และทำงานจนถึง n-1-i เนื่องจากค่าที่มากที่สุดจะถูกดันไปขวาสุดทุกครั้ง
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // ค่าที่ตำแหน่ง j มากกว่า j+1 จะทำการสลับตำแหน่ง (swap) โดยใช้ในการแลกเปลี่ยนค่า
			}
		}
	}
}

func main() {
	arr := []int{45, 78, 12, 34, 89, 20}
	bubbleSort(arr)
	fmt.Println("sorted array:", arr)
}
