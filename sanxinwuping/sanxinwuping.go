package sanxinwuping

import (
	"fmt"
	"math/rand"
	"time"
)

func Sanxin() {
	rand.Seed(time.Now().UnixNano()) //生成随机数种子
	var a int = rand.Intn(22)
	switch a {
	case 0:
		fmt.Println("冷刃")
	case 1:
		fmt.Println("黎明神剑")
	case 2:
		fmt.Println("吃虎鱼刀")
	case 3:
		fmt.Println("飞天御剑")
	case 4:
		fmt.Println("铁影阔剑")
	case 5:
		fmt.Println("沐浴龙血的剑")
	case 6:
		fmt.Println("白铁大剑")
	case 7:
		fmt.Println("以理服人")
	case 8:
		fmt.Println("飞天大御剑")
	case 9:
		fmt.Println("白缨枪")
	case 10:
		fmt.Println("钺矛")
	case 11:
		fmt.Println("黑缨枪")
	case 12:
		fmt.Println("魔导绪论")
	case 13:
		fmt.Println("讨龙英杰谭")
	case 14:
		fmt.Println("异世界行记")
	case 15:
		fmt.Println("翡玉法球")
	case 16:
		fmt.Println("甲级宝珏")
	case 17:
		fmt.Println("鸦羽弓")
	case 18:
		fmt.Println("神射手之誓")
	case 19:
		fmt.Println("反曲弓")
	case 20:
		fmt.Println("弹弓")
	case 21:
		fmt.Println("信使")
	}

}
