package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/task"

	"github.com/robfig/cron/v3"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// Clean up DB scheduled tasks
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // The scheduled task method is set in the task file package
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "Regularly clear the contents of the database [log, blacklist]", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// Other scheduled tasks are set here. Please refer to the usage method above.

		//_, err := global.GVA_Timer.AddTaskByFunc("Timed task identifier", "corn expression", func() {
		// Specific execution content...
		//  ......
		//}, option...)
		//if err != nil {
		//	fmt.Println("add timer error:", err)
		//}
	}()
}
