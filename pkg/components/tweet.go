package components

type TweetStruct struct {
	componentStruct
	URL string `json:"URL"`
}

func NewTweet() *TweetStruct {
	i := TweetStruct{}
	i.SetRole("tweet")
	return &i
}
