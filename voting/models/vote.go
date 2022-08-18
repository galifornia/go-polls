package models

type Vote struct {
	Id     string `json:"-"`
	PollId string `json:"poll_id"`
	Votes  []int  `json:"votes"`
}
