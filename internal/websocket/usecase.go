package websocket

type EventType struct {
	UpdateCustomerRanking
}
type Order struct {
	OrderID      string
	CustomerID   string
	Status       string
	BoosterID    string
	Username     string
	Password     string
	SummonerName string
	SummonerID   string
	Region       string
}
type UpdateCustomerRanking struct {
	CustomerID string
	Rank       string
	Division   string
	LP         string
}

type Event struct {
	Type    *EventType
	Payload interface{}
}
type UseCase interface {
	LPChangeEvent(event Event)
	OrderEvent(event Event)
}
