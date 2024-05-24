package applications

type InternalAvailableApplicationProvider struct {
}

func (this *InternalAvailableApplicationProvider) GetMapIndex() map[string][]Application {
	return available
}

func (this *InternalAvailableApplicationProvider) List() []Application {
	return FlattenMap(this.GetMapIndex())
}
