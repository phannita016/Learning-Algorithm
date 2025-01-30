package main

import (
	"fmt"
	"sort"
)

/*
	ปัญหาการหาเหรียญที่น้อยที่สุดเพื่อทอนเงิน โดยที่เรามีเหรียญทอนหลายประเภทและต้องการทอนเงินให้ได้จำนวนเหรียญน้อยที่สุด
	***การวิเคราะห์ประสิทธิภาพ***
		- การทำงานในเวลา O(n)
*/

func findMincount(coins []int, amount int) []int {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	result := make([]int, 0) // ตัวแปรเก็บเหรียญที่ใช้ในการทอน
	for _, coin := range coins {
		for amount >= coin { // จำนวนเงินที่ยังต้องทอนมีมากกว่าหรือเท่ากับมูลค่าของเหรียญ
			result = append(result, coin) // เพิ่มเหรียญเข้าไปใน result
			amount -= coin                // ลดจำนวนเงินที่ต้องทอนลง จนกว่าจะไม่สามารถใช้เหรียญนี้ได้อีก
		}
	}

	// ถ้า amount ยังไม่เป็น 0 แสดงว่าไม่สามารถทอนเงินให้หมดได้ด้วยเหรียญที่มีอยู่
	if amount != 0 {
		return []int{}
	}

	return result
}

/*
	คุณมีกระเป๋า (knapsack) ที่สามารถบรรจุของได้จำกัด และแต่ละของมีน้ำหนักและมูลค่า คุณต้องการที่จะใส่ของลงไปในกระเป๋าเพื่อให้มูลค่ารวมมากที่สุดโดยที่น้ำหนักรวมไม่เกินขีดจำกัด
	***การวิเคราะห์ประสิทธิภาพ***
		- การทำงานในเวลา O(n)
*/

type Item struct {
	value  int
	weight int
	index  int
}

func knapsack(capacity int, items []Item) int {
	sort.Slice(items, func(i, j int) bool {
		return items[i].value/items[i].weight > items[j].value/items[j].weight
	}) // เรียงลำดับ item ตามมูลค่าต่อน้ำหนัก

	totalValue := 0 // เก็บมูลค่ารวมของสิ่งที่ใส่ในกระเป๋า
	for _, item := range items {
		if capacity >= item.weight { // ความจุกระเป๋าที่เหลือ >= น้ำหนักของ item
			totalValue += item.value // เพิ่มมูลค่าของitem เข้าไปใน totalValue
			capacity -= item.weight  // ลดความจุกระเป๋าตามน้ำนหักของที่ใส่
		} else { // กรณีของมีน้ำหนักมากเกินไป ไม่สามารถใส่ทั้งหมดได้ จะใส่ของแค่บางส่วนที่สามารถใส่ได้
			totalValue += item.value * capacity / item.weight
			break
		}
	}

	return totalValue
}

/*
เลือกของที่มีมูลค่าต่อน้ำหนัก (value per weight) สูงที่สุดก่อน แล้วค่อยๆ เพิ่มของที่เหลือไป สำหรับของที่มีมูลค่าสูงแต่มีน้ำหนักมาก, อาจเลือกนำบางส่วนของมันไปบรรจุในกระเป๋า (Fractional)
*/
type Activity struct {
	start int
	end   int
}

func activitySelection(activities []Activity) []Activity {
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].end < activities[j].end
	}) // เรียงกิจกรรมตามเวลาที่จบ จากน้อยไปมาก

	selected := []Activity{activities[0]} // เลือกกิจกรรมแรกที่จบเร็วที่สุด
	lastSelected := activities[0]         // ให้เท่ากับกิจกรรมแรก เพื่อใช้ในการเช็ตกิจกรรมถัดไปว่าเริ่มจากกิจกรรมนี้หรือไม่

	for i := 1; i < len(activities); i++ { // วนลูปกิจกรรมที่เหลือ
		if activities[i].start >= lastSelected.end { // กิจกรรมจะเริ่มได้หลังจากกิจกรรมที่เลือกไปแล้วจบ
			selected = append(selected, activities[i]) // ถ้ากิจกรรมนี้สามารถดำเนินได้ ก็จะถูกเพิ่มเข้าไปใน seleted
			lastSelected = activities[i]               // lastSelected จะอัปเดตให้เป็นกิจกรรมลาสุดที่เลือก
		}
	}

	return selected
}

func main() {
	coins := []int{1, 5, 10, 25}
	amount := 63
	result := findMincount(coins, amount)
	fmt.Println("The minimum number of coins is: ", result)

	//
	items := []Item{
		{weight: 10, value: 60},
		{weight: 20, value: 100},
		{weight: 30, value: 120},
	}
	capacity := 50
	maxValue := knapsack(capacity, items)
	fmt.Printf("Maximum value in Knapsack = %d\n", maxValue)

	//
	activities := []Activity{
		{start: 1, end: 4},
		{start: 3, end: 5},
		{start: 0, end: 6},
		{start: 5, end: 7},
		{start: 3, end: 9},
		{start: 5, end: 9},
		{start: 6, end: 10},
		{start: 8, end: 11},
		{start: 8, end: 12},
	}
	selected := activitySelection(activities)
	fmt.Println("Selected activities:")
	for _, activity := range selected {
		fmt.Printf("Start: %d, End: %d\n", activity.start, activity.end)
	}
}
