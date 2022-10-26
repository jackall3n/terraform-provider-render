package types

type OwnerItem struct {
	Owner  Owner  `json:"owner"`
	Cursor string `json:"cursor"`
}

type Owner struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
