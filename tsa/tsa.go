package tsa

type Request interface {
	TSARequest(tsq []byte) ([]byte, error)
}
