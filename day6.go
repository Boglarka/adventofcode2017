package main

import (
  "fmt"
  "strings"
  "strconv"
)

var input = "0	5	10	0	11	14	13	4	11	8	8	7	1	4	12	11"

func task1() int {
  states, eq, found := [][]int{}, false, false
  numsS := strings.Split(input, "\t")
  nums := []int{}
  for _, i := range numsS {
    j, err := strconv.Atoi(i)
    if err != nil {
      panic(err)
    }
    nums = append(nums, j)
  }
  for found == false {
    pos, max := 0, 0
    for i := 0; i < len(nums); i++ {
      if nums[i] > max {
        max = nums[i]
        pos = i
      }
    }
    nums[pos] = 0
    i := pos + 1
    if i == len(nums) {
      i = 0
    }
    for j := 0;  j<max; j++ {
      nums[i]++
      i = (i+1) % len(nums)
    }
    c := make([]int, len(nums))
    copy(c, nums)
    states = append(states, c)
    if len(states) > 1 {
      for i := 0; i < len(states)-1; i++ {
        eq = true
        for k, _ := range states[i] {
          if states[i][k] != states[len(states)-1][k] {
            eq = false
            break
          }
        }
        if eq {
          break
        }
      }
    }
    if eq {
      break
    }
  }
  return len(states)
}

func task2() int {
  states, eq, found, which := [][]int{}, false, false, 0
  numsS := strings.Split(input, "\t")
  nums := []int{}
  for _, i := range numsS {
    j, err := strconv.Atoi(i)
    if err != nil {
      panic(err)
    }
    nums = append(nums, j)
  }
  for found == false {
    pos, max := 0, 0
    for i := 0; i < len(nums); i++ {
      if nums[i] > max {
        max = nums[i]
        pos = i
      }
    }
    nums[pos] = 0
    i := pos + 1
    if i == len(nums) {
      i = 0
    }
    for j := 0;  j<max; j++ {
      nums[i]++
      i = (i+1) % len(nums)
    }
    c := make([]int, len(nums))
    copy(c, nums)
    states = append(states, c)
    if len(states) > 1 {
      for i := 0; i < len(states)-1; i++ {
        eq = true
        for k, _ := range states[i] {
          if states[i][k] != states[len(states)-1][k] {
            eq = false
            break
          }
        }
        if eq {
          which = i
          break
        }
      }
    }
    if eq {
      break
    }
  }
  return len(states) - which - 1
}

func main() {
  //fmt.Println(task1())
  fmt.Println(task2())
}
