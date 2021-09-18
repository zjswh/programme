package types

type ProgrammeStruct struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	MssUrl    string `json:"mssUrl"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	ChannelId int    `json:"channelId"`
	Uin       int    `json:"uin"`
	Type      int    `json:"type"`
	Week      int    `json:"week"`
	STaskId   int    `json:"sTaskId"`
	ETaskId   int    `json:"eTaskId"`
}

type Request struct {
}

type Response struct {
}
