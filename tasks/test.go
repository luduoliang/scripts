package tasks

import (
	"fmt"
	"log"
	"scripts/db"
	"scripts/library"
	"time"
)

type Test struct {
	SleepTime time.Duration
}

//初始化脚本
func init() {
	//脚本执行时间间隔（秒）
	task := Test{SleepTime:2}
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
func (t *Test)Execute() {
	//查询多条数据
	business := []Business{}
	db.DB.Select(&business, "SELECT id FROM business WHERE id < ?", 10)
	log.Println(business)

	//查询单条数据
	business_info := Business{}
	db.DB.Get(&business_info, "SELECT id FROM business WHERE id < ? ORDER BY id DESC", 10)
	log.Println(business_info)

	//联表查询
	rows, _ :=db.DB.Query("SELECT a.description as description, b.phone as phone FROM user_wallet_log a LEFT JOIN user b ON b.id=a.user_id WHERE a.id < ?", 10)
	var list []map[string]interface{}
	for rows.Next() {
		var description, phone string
		rows.Scan(&description, &phone)
		list = append(list, map[string]interface{}{"description":description, "phone":phone})
	}
	fmt.Println(list)


	//事务处理
	if tx, err := db.DB.Begin();err == nil{
		_, err = tx.Exec(
			"UPDATE business SET expise_time = ? WHERE id < ?",
			library.CurrentTime(),
			10,
		)
		if err != nil {
			tx.Rollback()
		}
		_, err = tx.Exec(
			"INSERT INTO business (name) VALUES(?)",
			"test",
		)
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}

}

type Business struct {
	Id int64
}