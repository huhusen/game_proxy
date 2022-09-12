package task

import "sockets-proxy/gcld/cmd"

type FinishTask struct {
	cmd.Command
	Send2 struct {
		group string
		index string
		types string
	}
}

func NewFinishTask() *FinishTask {
	u := FinishTask{}
	u.Cmd = cmd.TaskFinishTask
	u.Zh = "【任务】完成任务"
	return &u
}
