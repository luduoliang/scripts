package tasks

import (
	"fmt"
	"time"
)

type Good struct {
	SleepTime time.Duration
}

//初始化脚本
func init() {
	//脚本执行时间间隔（秒）
	task := Good{SleepTime:3}
	//taskChan := make(chan bool)
	//每个脚本对应一个主线程，永远不会结束，因为里面是个死循环
	go func(){
		for {
			/*//子线程运行一遍后，执行for下一次循环。备注：子线程这快可以不用
			go func(){
				task.Execute()
				defer func(){
					taskChan <- true
				}()
			}()
			//等待子线程结束
			<-taskChan*/

			task.Execute()
			time.Sleep(task.SleepTime*time.Second)
		}
	}()
}

//入口函数,每隔一定时间会执行一次
func (t *Good)Execute() {
	fmt.Println("good...")
}