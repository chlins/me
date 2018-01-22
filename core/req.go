package core

// Request connect param
type Request struct {
	// p-producer c-consumer
	Role  string `json:"role"`
	Topic string `json:"topic"`
}
