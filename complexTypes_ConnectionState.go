package kurento

// State of the connection.
type ConnectionState string

// Implement fmt.Stringer interface
func (t ConnectionState) String() string {
	return string(t)
}

const (
	CONNECTIONSTATE_DISCONNECTED ConnectionState = "DISCONNECTED"
	CONNECTIONSTATE_CONNECTED    ConnectionState = "CONNECTED"
)