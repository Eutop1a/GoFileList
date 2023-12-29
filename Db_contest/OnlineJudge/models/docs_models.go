package models

type _ResponseCode struct {
	Code int `json:"code"`
}
type _ResponseError struct {
	Error string `json:"error"`
}

type _ResponseMsg struct {
	Msg string `json:"msg"`
}

type _ResponseToken struct {
	Token string `json:"Token"`
}

type _ResponseReg struct {
	_ResponseToken
	EndTime string `json:"end_time"`
}

type _ResponseRegErr struct {
	_ResponseMsg
	Err string `json:"err"`
}

type _ResponseSort struct {
	//Data []User `json:"sort"`
}

type _ResponseAddScore struct {
	_ResponseMsg
	NewScore int `json:"new_score"`
}

type _ResponseUsername struct {
	Username string `json:"username"`
}

type _ResponseProblems struct {
	//Msg []Problems `json:"judge"`
}

type _ResponseAddProblems struct {
	Code string `json:"code"`
}

type _ResponseDataString struct {
	Data string `json:"data"`
}
type _ResponseGetData struct {
	Data []string `json:"data"`
}

type _ResponseQuestionDetail struct {
	//Data ProblemList `json:"data"`
}
