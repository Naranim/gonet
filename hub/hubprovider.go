package hub

type HubProvider struct{}

func NewHubProvider() *HubProvider {
	return new(HubProvider)
}

func NewHub() *Hubber {
	return New()
}
