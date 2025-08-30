package util

import "log"

func LogError(err error) {
	if err != nil {
		log.Fatal("cli_task_tracker: error ", err)
	}
}
