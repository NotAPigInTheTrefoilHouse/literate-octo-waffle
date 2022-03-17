package sixinwuping //作为主函数里四星物品的包

import (
	"fmt"
	"math/rand"
	"time"
)

// 全局变量
var gaiup int

// 通用卡池
func tongyong(a int) {
	switch a {
	case 0:
		fmt.Println("西风剑 ****")
	case 1:
		fmt.Println("西风大剑 ****")
	case 2:
		fmt.Println("西风猎弓 ****")
	case 3:
		fmt.Println("西风秘典 ****")
	case 4:
		fmt.Println("西风长枪 ****")
	case 5:
		fmt.Println("祭礼剑 ****")
	case 6:
		fmt.Println("祭礼大剑 ****")
	case 7:
		fmt.Println("祭礼弓 ****")
	case 8:
		fmt.Println("祭礼残章 ****")
	case 9:
		fmt.Println("匣里灭尘 ****")
	case 10:
		fmt.Println("匣里龙吟 ****")
	case 11:
		fmt.Println("钟剑 ****")
	case 12:
		fmt.Println("绝弦 ****")
	case 13:
		fmt.Println("流浪乐章 ****")
	case 14:
		fmt.Println("笛剑 ****")
	case 15:
		fmt.Println("雨裁 ****")
	case 16:
		fmt.Println("弓藏 ****")
	case 17:
		fmt.Println("昭心 ****")
	case 18:
		fmt.Println("托马 ****")
	case 19:
		fmt.Println("烟绯 ****")
	case 20:
		fmt.Println("班尼特 ****")
	case 21:
		fmt.Println("辛焱 ****")
	case 22:
		fmt.Println("香菱 ****")
	case 23:
		fmt.Println("行秋 ****")
	case 24:
		fmt.Println("芭芭拉 ****")
	case 25:
		fmt.Println("早柚 ****")
	case 26:
		fmt.Println("砂糖 ****")
	case 27:
		fmt.Println("九条裟罗 ****")
	case 28:
		fmt.Println("雷泽 ****")
	case 29:
		fmt.Println("北斗 ****")
	case 30:
		fmt.Println("菲谢尔 ****")
	case 31:
		fmt.Println("罗莎莉亚 ****")
	case 32:
		fmt.Println("重云 ****")
	case 33:
		fmt.Println("迪奥娜 ****")
	case 34:
		fmt.Println("云堇 ****")
	case 35:
		fmt.Println("五郎 ****")
	case 36:
		fmt.Println("凝光 ****")
	case 37:
		fmt.Println("诺艾尔 ****")
	}
}

// up卡池
func up(a int) {
	switch a {
	case 0, 1:
		fmt.Println("辛焱 ****")
	case 2:
		fmt.Println("九条裟罗 ****")
	case 3:
		fmt.Println("班尼特 ****")
	}
}

func Sixin() {
	rand.Seed(time.Now().UnixNano()) //生成随机数种子
	var a int
	var b int = rand.Intn(2)
	if gaiup == 1 || b == 0 {
		a = rand.Intn(4)
		up(a)
		gaiup = 0
	} else if gaiup == 0 {
		a = rand.Intn(38)
		tongyong(a)
		gaiup = 1
		if a == 20 || a == 21 || a == 27 {
			gaiup = 0
		}
	}
}
