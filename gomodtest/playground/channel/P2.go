package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 初始数量设置为 50 个
var currentQuantity = 50

type order struct {
	// 需要生产的包子数
	needNum int
	// 订单编号
	id int
}

// 生产者
func producer(ch chan *order) {
	// fmt.Println("进入 生产者")
	for {
		infos := <- ch
		num := infos.needNum
		id := infos.id
		fmt.Printf("[%d]号顾客 需要 %d 个包子：\n", id, num)
		if currentQuantity < num {
			// 还差多少个包子
			needAdd := -(currentQuantity - num)
			fmt.Printf("【不足】包子库存不足，还差 %d 个，正在生产 \n", needAdd)
			currentQuantity += needAdd
			fmt.Printf("【不足】生产了 %d 个，当前库存 %d 个 \n",
				needAdd, currentQuantity)
			currentQuantity -= num
			fmt.Printf("【不足】消费了 %d 个，当前库存 %d 个 \n",
				num, currentQuantity)
		}else {
			currentQuantity -= num
			fmt.Printf("【充足】已扣除 %d 个包子，剩余 %d 个 \n",
				num, currentQuantity)
		}
	}
}

// 消费者
func consumer(ch chan *order, msg *order) {
	// fmt.Println("进入 消费者")
	ch <- msg
}

func main() {
	ch := make(chan *order)
	time.Sleep(time.Second * 3)
	go producer(ch)
	//for i := 0; i < 20; i++ {
	//	go consumer(ch, &order{
	//		needNum: i + rand.Intn(10),
	//		id:      i,
	//	})
	//}

	for {
		go consumer(ch, &order{
			needNum: rand.Intn(10),
			id:      rand.Intn(1000),
		})
	}
}
