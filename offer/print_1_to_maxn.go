package main

import (
    "fmt"
    "flag"
    "time"
)

var w = flag.Int("n", 1, "input n")

func find(n int) {
    start := time.Now()
    i, max, p := 1, 0, 1
    for i <= n {
        max += p * 9   
        i++
        p *= 10
    }
    fmt.Println(max)
    t := time.Since(start)
    fmt.Println(t)

   // for ii := 1; ii <= max; ii++ {
   //     fmt.Printf("%d ", ii)
   // }
}

//
func find1(n int) {
  start := time.Now()
  max := 1
  for ii := 0; ii < n; ii++ {
    max *= 10
  }
  fmt.Println(max)
  fmt.Println(time.Since(start))
}

// error
func find2(n int) {
    s := make([]byte, n)
    //flag := 0   // 进位
    i := n - 1

    for i < n && s[i] <= '9' {
        if s[i] == 0 {
            s[i] = '0' 
        }
        s[i] += 1
        fmt.Println(string(s))
        if s[i] == '9' {
            
            i++
        }
    }
    
}

// error
func find3(n int) {
    s := make([]byte, n)
	//flag := 0   // 进位
	i := n - 1
	for i >= 0 {
        if s[i] == 0 {
            s[i] = '0'
        }
		if s[i] == '9' && i+1 < n {
			if s[i+1] == 0 {
				s[i+1] = '1'
			}
			s[i+1] += 1
			i--
		} else {
			s[i] += 1
		}
        fmt.Println(string(s))
	}
}

func find4(n int) {
    s := make([]byte, n)
    memset(s)

    i := n - 1
    //temp := 0
    for i >= 0 {
        if s[i] < '9' {
            s[i] += 1
        } else if s[i] == '9' && i-1 != '9' {
            s[i-1] += 1
        } else {
            i--
        }
        fmt.Println(string(s))
    }
}

func memset(b []byte) {
    for i := 0; i < len(b); i++ {
        b[i] = '0'
    }
}

func main() {
    flag.Parse()
    //find(*w)
    //find1(*w)
    //find2(*w)
    //find3(*w)
    find4(*w)
}
