package igolang

import (
	"fmt"
	"github.com/pebbe/zmq4"
)

func RunKernel(cinfo ConnectionFile) {
	heartBeat := zmq4.NewSocket(zmq4.REP)
	heartBeat.Bind(fmt.Sprintf("%s://%s:", cinfo.Transport, cinfo.Ip, cinfo.HbPort))
}
