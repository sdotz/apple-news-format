package components

type AuthorStruct struct {
	TextStruct
}

func NewAuthor() *AuthorStruct {
	i := AuthorStruct{
		TextStruct: TextStruct{
			componentStruct: componentStruct{},
		},
	}
	i.SetRole("author")
	return &i
}
