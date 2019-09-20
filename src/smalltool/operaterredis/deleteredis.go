/* 2019/05/05
 * AlvinFu
 * 主要用于删除Redis的key，可以进行模糊删除，也可以精确删除
 * 主要配合数据组件进行使用
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

//删除redis的key，范围删除
func main() {
	var host, prot, keys string
	fmt.Println("欢迎使用删除指定redis数据功能!")
	fmt.Println("请输入要删除的redis的ip地址")
	fmt.Scanln(&host)
	fmt.Println("请输入要删除的redis的端口号")
	fmt.Scanln(&prot)
	fmt.Println("请输入要删除的redis的keys，最好是唯一确定的")
	fmt.Scanln(&keys)
	fmt.Println("开始执行命令请等待")
	cmd := exec.Command("redis-cli", "-h", host, "-p", prot, "KEYS", "*"+keys+"*")
	tmp := MyCommand(cmd)
	for i := 0; i < len(tmp)-1; i++ {
		cmd := exec.Command("redis-cli", "-p", prot, "DEL", tmp[i])
		MyCommand(cmd)
	}
	fmt.Println("命令执行完毕,谢谢使用!")
}

func MyCommand(cmd *exec.Cmd) []string {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("pipe err:", err)
		return nil
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("start err:", err)
		return nil
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("read err:", err)
		return nil
	}
	fmt.Println("执行命令的返回结果", string(opBytes))
	tmp := strings.Split(string(opBytes), "\n")
	return tmp
}
