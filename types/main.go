package types

type User struct {
	Username string
}

type Event struct {
	Name         string
	OwnerID      int32
	Participants []int32
}
