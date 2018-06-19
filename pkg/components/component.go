package components


type Component interface {
	Role() string
	SetRole(string) Component
}

type componentStruct struct {
	Role string `json:"role"`
}
