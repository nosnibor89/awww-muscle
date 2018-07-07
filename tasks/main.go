package tasks

import "github.com/markbates/grift/grift"


type TaskConfig struct {
	Name string
	Description string
	Namespace string
	Task func(c *grift.Context) error
}

type TaskConfigs []TaskConfig

func GetTasks() []TaskConfig {
	tasks := TaskConfigs{
		TaskConfig{
			Name: "migrate",
			Description: "Migrating DB...",
			Task: migrateDB,
		},
	}

	return tasks
}