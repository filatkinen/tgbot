package lorry

import "fmt"

type Lorry struct {
	Model string `json:"model"`
	ID    uint64 `json:"id"`
}

func (l Lorry) String() string {
	return fmt.Sprintf("ID -%d, Модель-%s", l.Model)
}
