package igolang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pebbe/zmq4"
)

type Kernel struct {
	Sockets Sockets
}

type Sockets struct {
	context       *zmq4.Context
	ShellSocket   *zmq4.Socket
	ControlSocket *zmq4.Socket
	StdinSocket   *zmq4.Socket
	IOPubSocket   *zmq4.Socket
	Key           []byte
}

func ReadConnectionFile(connFilePath string) ConnectionFile {
	connBytes, err := ioutil.ReadFile(connFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var connectionFile ConnectionFile
	err = json.Unmarshal(connBytes, &connectionFile)
	if err != nil {
		log.Println("Error unmarshalling connection_file")
		log.Fatal(err)
	}
	return connectionFile
}

func NewSockets(cinfo ConnectionFile) (Sockets, error) {
	context, err := zmq4.NewContext()
	if err != nil {
		return Sockets{}, err
	}
	s := Sockets{
		context: context,
		Key:     []byte(cinfo.Key)}

	address := func(port int) string { return fmt.Sprintf("%s://%s:%d", cinfo.Transport, cinfo.Ip, port) }

	// setup heartbeat
	heartBeat, err := context.NewSocket(zmq4.REP)
	if err != nil {
		log.Println("Error in setting up heart beat")
		log.Fatal(err)
	}
	err = heartBeat.Bind(address(cinfo.HbPort))
	if err != nil {
		log.Println("Error in setting up heart beat")
		log.Fatal(err)
	}

	// setup sockets
	s.ShellSocket, err = context.NewSocket(zmq4.ROUTER)
	if err != nil {
		return s, err
	}
	s.ControlSocket, err = context.NewSocket(zmq4.ROUTER)
	if err != nil {
		return s, err
	}
	s.StdinSocket, err = context.NewSocket(zmq4.ROUTER)
	if err != nil {
		return s, err
	}
	s.IOPubSocket, err = context.NewSocket(zmq4.ROUTER)
	if err != nil {
		return s, err
	}
	err = s.ShellSocket.Bind(address(cinfo.ShellPort))
	if err != nil {
		return s, err
	}
	err = s.ControlSocket.Bind(address(cinfo.ControlPort))
	if err != nil {
		return s, err
	}
	err = s.StdinSocket.Bind(address(cinfo.StdinPort))
	if err != nil {
		return s, err
	}
	err = s.IOPubSocket.Bind(address(cinfo.IOpubPort))
	if err != nil {
		return s, err
	}

	go func(heartBeat *zmq4.Socket) {
		err = zmq4.Proxy(heartBeat, heartBeat, nil)
		if err != nil {
			log.Fatal(err)
		}
	}(heartBeat)

	return s, nil
}

func NewKernel(connFilePath string) Kernel {
	cinfo := ReadConnectionFile(connFilePath)
	kernel := Kernel{}
	var err error
	kernel.Sockets, err = NewSockets(cinfo)
	if err != nil {
		log.Fatal(err)
	}

	return kernel
}

func (k Kernel) Run() {
	for {
	}
}
