package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"time"
)

var (
	length  int
	charset string
	num     int
	save string
)

const (
	// 适当重复，以平衡各类型字符的频率(字母 > 数字 > 特殊字符)
	NUmStr  = "0123456789012345678901234567890123456789"
	CharStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-#~.!^*$/&_+=-#~.!^*$/&_"
)

// 解析参数
func parseArgs() {
	flag.IntVar(&num, "n", 5, "-n 生成的个数")
	flag.IntVar(&length, "l", 8, "-l 生成密码的长度")
	flag.StringVar(&save, "o", "nil", `-o 将密码保存到到指定位置,
	如果不指定位置但使用了该参数, 请输入default, 将保存在默认位置: 【桌面/out.txt】
	必须对参数值添加单或双引号！`)
	flag.StringVar(&charset, "t", "mix",
		`-t 设置密码复杂度,
        num:只使用数字[0-9], 如: 637599124317,
        char:只使用英文字母[a-zA-Z], 如: SHXEXPHjoJsx,
        mix:使用数字[0-9]和字母[a-zA-Z], 如: r12srF27nfB3,
        advance:使用数字[0-9]、字母[a-zA-Z]以及特殊字符[+=-#~.!^*$/&_], 如: F[]QnK*a-2aE
	注意：
	1. 建议对参数值添加单或双引号,
	2. mix模式和advance模式下, 首字符必是字母`)
	flag.Parse()
}

func generatePasswd() string {
	// 初始化密码切片
	var passwd []byte = make([]byte, length)
	// 源字符串
	var sourceStr string

	if charset == "num" {
		// 数字模式
		sourceStr = NUmStr
	} else if charset == "char" {
		// 字母模式
		sourceStr = CharStr
	} else if charset == "mix" {
		// 混合模式
		sourceStr = fmt.Sprintf("%s%s", NUmStr, CharStr)
	} else if charset == "advance" {
		// 高级模式
		sourceStr = fmt.Sprintf("%s%s%s", NUmStr, CharStr, SpecStr)
	} else {
		sourceStr = CharStr
	}

	for i := 0; i < length; i++ {
		// 如果是mix或advance模式，则保证第一个位置为字母
		if i == 0 && (charset == "mix" || charset == "advance") {
			index := rand.Intn(len(CharStr))
			passwd[i] = CharStr[index]
		} else {
			index := rand.Intn(len(sourceStr))
			passwd[i] = sourceStr[index]
		}

	}
	return string(passwd)
}

func save2file(passes []string) {
	// 保存到文件
	myself, error := user.Current()
	if error != nil {
		panic(error)
	}
	homedir := myself.HomeDir
	if save == "default" {
		save = homedir + "/Desktop/" + "out.txt"
	}
	file, err := os.OpenFile(save, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
		os.Exit(3)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	for i := 0; i < len(passes); i++ {
		write.WriteString(passes[i])
		write.WriteString("\n")
	}
	write.Flush()
	fmt.Println("密码文件保存成功！位置:", save)
}

func main() {
	parseArgs()
	passes := []string{}
	for i := 0; i < num; i++ {
		// 随机种子
		time.Sleep(500 * time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		passwd := generatePasswd()
		fmt.Println(passwd)
		passes = append(passes, passwd)
	}
	if save != "nil" {
		save2file(passes)
	}
}
