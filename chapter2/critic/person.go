package critic

type Person struct {
	Name   string   `json:"name"`
	Movies []*Movie `json:"movies"`
}
