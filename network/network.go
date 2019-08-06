package network

type Network struct {
	Name     string
	Filename string
}

func ListNames(networks []Network) []string {
	names := make([]string, 0, len(networks))
	for _, network := range networks {
		names = append(names, network.Name)
	}

	return names
}
