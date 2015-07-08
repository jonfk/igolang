package igolang

import ()

// https://ipython.org/ipython-doc/dev/development/kernels.html#connection-files
type ConnectionFile struct {
	ControlPort     int    `json:"control_port"`
	ShellPort       int    `json:"shell_port"`
	Transport       string `json:"transport"`
	SignatureScheme string `json:"signature_scheme"`
	StdinPort       int    `json:"stdin_port"`
	HbPort          int    `json:"hb_port"`
	Ip              string `json:"ip"`
	IopubPort       int    `json:"iopub_port"`
	Key             string `json:"key"`
}

type Message struct {
	Header       Header
	ParentHeader Header
}

type Header struct {
	MsgId    string `json:"msg_id"`
	Username string `json:"username"`
	Session  string `json:"session"`
	MsgType  string `json:"msg_type"`
}
