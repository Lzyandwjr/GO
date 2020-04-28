package main

import (
	"fmt"
	"time"
)

// time demo


func main(){
	Time := time.Now()
	fmt.Println(Time)
	year := 	Time.Year()
	mouth := Time.Month()
	day := Time.Day()
	hour := Time.Hour()
	minute := Time.Minute()
	second := Time.Second()
	fmt.Println(year,mouth,day,hour,minute,second)

	// 时间戳：from 1970.1.1 到现在的一个秒数
	timestamp1 := Time.Unix()
	timeStamp2 :=Time.UnixNano()
	fmt.Println(timestamp1,timeStamp2)
	// 将时间戳转换为具体的时间格式
	t := time.Unix(timestamp1+3600,0)
	fmt.Println(t)

	// 时间间隔
	//time.Sleep(5 * time.Second)
	fmt.Println("over")

	T2 := Time.Add(time.Hour)
	fmt.Println(T2)


	fmt.Println(T2.Sub(Time))

	// 定时器
	//for  temp :=range time.Tick(time.Second*2){
	//	fmt.Println(temp)
	//}
	// 时间格式化：Y-m-d H:M:S  -> Y m d H M S
	//       go中格式定义为 ： 2006 01 02 15 04 05
	ret1 := Time.Format("2006-01-02 15:04:05")
	fmt.Println(ret1)
	ret2 := Time.Format("2006-01-02 03:04:05 PM")
	fmt.Println(ret2)

	//解析字符串类型的时间
	timeStr := "2020/04/24 09:51:00"

	loc,err := time.LoadLocation("Asia/Shanghai")
	if err != nil{
		fmt.Println(err)
		return
	}

	//2.根据市区去解析一个字符串格式的时间

	timeObj ,err :=time.Parse("2006/01/02 15:04:05",timeStr)
	if err != nil{
		fmt.Printf("parse timeStr failed ,err:%v\n",err)
		return
	}
	fmt.Println(timeObj)
	timeObj2,err := time.ParseInLocation("2006/01/02 15:04:05",timeStr,loc)
	if err != nil{
		fmt.Printf("parse timeStr failed ,err:%v\n",err)
	}
	fmt.Println(timeObj2)
}

