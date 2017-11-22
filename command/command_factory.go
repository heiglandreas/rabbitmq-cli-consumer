package command

import (
	"os/exec"
    "bytes"
)

type CommandFactory struct {
	Cmd  string
	Args []string
}

func (me CommandFactory) Create(body string) *exec.Cmd {
    execution := exec.Command(me.Cmd, me.Args...)

    var b bytes.Buffer
    b.Write([]byte(body))
    execution.Stdin = &b

    return execution

//	return exec.Command(me.Cmd, append(me.Args, body)...)
}
