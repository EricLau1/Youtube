package utils

type Alert struct {
	Message string
	Type string
}

func NewAlert(message, alert string) Alert {
	return Alert{message, alert}
}