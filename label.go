package trello

type Label struct {
	client  *Client
	Id      string  `json:"id"`
	IdBoard string  `json:"idBoard"`
	Name    string  `json:"name"`
	Color   string  `json:"color"`
	Uses    int 	`json:"uses"`
}
