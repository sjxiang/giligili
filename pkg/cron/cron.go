package cron

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/robfig/cron"
)

// Cron 定时器单例
var Cron *cron.Cron



type handler func() error


// Run 运行
func Run(job handler) {
	from := time.Now().UnixNano()

	err := job()

	to := time.Now().UnixNano()

	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil {
		fmt.Printf("%s error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		fmt.Printf("%s success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}

}

// CronJob 定时任务
func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	} 

	// 在凌晨执行
	Cron.AddFunc("0 0 0 * * *", func() { Run(RestartDailyRank) })
	Cron.Start()

	fmt.Println("定时任务开启 ...... ")
}	