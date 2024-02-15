package common

type BusinessCode int
type Resp struct {
	Code BusinessCode `json:"code"`
	Data any          `json:"data"`
	Msg  string       `json:"msg"`
}

func (r *Resp) Success(data any) *Resp {
	r.Code = 200
	r.Msg = "success"
	r.Data = data
	return r
}

func (r *Resp) Fail(code BusinessCode, msg string) *Resp {
	r.Code = code
	r.Msg = msg
	return r
}
