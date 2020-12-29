package notification

type Client interface {
	GetTargets() (*string, error)
}
