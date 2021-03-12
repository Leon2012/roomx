package common

func Success() *Resp {
	return &Resp{
		Code: 0,
		Msg:  "",
	}
}

func Error(code int32, msg string) *Resp {
	return &Resp{
		Code: code,
		Msg:  msg,
	}
}
