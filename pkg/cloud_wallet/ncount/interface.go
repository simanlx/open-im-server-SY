package ncount

type NCounter interface {
	NewAccountURL() string
}

type counter struct {
}
