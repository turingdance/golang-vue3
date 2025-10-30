package utils

import (
	"os"
)

type MsgHandler func(in string) error
type Ipc struct {
	path      string
	MsgChan   chan string
	pipeline  *os.File
	msgandler MsgHandler
}

func NewIpc(path ...string) (ipc *Ipc) {
	dstpath := "./ipc"
	if len(path) > 0 {
		dstpath = path[0]
	}
	ipc = &Ipc{
		path:    dstpath,
		MsgChan: make(chan string, 10),
	}
	return ipc
}

func (ipc *Ipc) Start() (err error) {
	ipc.pipeline, err = os.OpenFile(ipc.path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0)
	if err != nil {
		return
	}
	go ipc.read()
	return nil
}

func (ipc *Ipc) read() (err error) {
	for {
		msg := make([]byte, 1024)
		n, e := ipc.pipeline.Read(msg)
		if e != nil {
			return e
		}
		ipc.MsgChan <- string(msg[:n])
		if ipc.msgandler != nil {
			ipc.msgandler(string(msg[:n]))
		}
	}
}

func (ipc *Ipc) write(msg string) (err error) {
	_, err = ipc.pipeline.WriteString(msg)
	return
}

// 添加
func (ipc *Ipc) Notify(sig string) (err error) {
	return ipc.write(sig)
}

// 添加
func (ipc *Ipc) SetMsgHandler(fun MsgHandler) {
	ipc.msgandler = fun
}
