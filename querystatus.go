package sqlow

// Status is  status after execution.
type Status string

var (
	// ADD is if added, this will be the status.
	ADD Status = "ADD"

	// PASS is if passed, the status will be this.
	PASS Status = "PASS"

	// UPDATE is will be the status if it has been updated.
	UPDATE Status = "UPDATE"

	// ERROR is will be in this status if the execution fails for some reason.
	ERROR Status = "ERROR"
)
