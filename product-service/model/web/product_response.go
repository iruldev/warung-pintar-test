package web

type ProductResponse struct {
	Id		int		`json:"id"`
	Name 	string 	`json:"name"`
	Price	int 	`json:"price"`
	Stock	int 	`json:"stock"`
	Weight	int 	`json:"weight"`
}
