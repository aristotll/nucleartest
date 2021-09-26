package main

import (
    "fmt"
    "sort"
)

// func threeSum(nums []int) [][]int {
//     r := make([][]int, 0)
//     //m := make(map[int]int)
//     if len(nums) <= 1 {
//         return r
//     }
 
//     p1, p2 := 0, 1
//     for p1 < len(nums) {
//         sum := nums[p1]
//         m := make(map[int]int)
//         rr := make([]int, 0)
//         for p2 < len(nums) {
//             //sum += nums[p2]
//             if _, ok := m[nums[p2]]; ok {
//                 continue
//             } 
//             sun += numsp[p2]
//             if sum < 0 {
//                 m[nums[p2]] = 0
//                 rr = append(rr, nums[p2])
//                 p2++
//             } else if sum == 0 {
                
//             } else {
                
//             }
//         } 
//     }
// }

func try(nums []int) [][]int {
    sort.Slice(nums, func(i, j int) bool {
        return i < j
    })

}

func main() {
    
}
