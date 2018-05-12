
package grifts

import (
	"fmt"
	. "github.com/markbates/grift/grift"
	"github.com/nosnibor89/awww-muscle/tasks"
)

var _ = Desc("hello", "Say Hello!")
var _ = Add("hello", func(c *Context) error {
	fmt.Println("Hello World!")
	return nil
})

func init()  {
	tasks := tasks.GetTasks()

	for _, task := range tasks {
		var _ = Desc(task.Name, task.Description)
		var _ = Add(task.Name, task.Task)
	}
}