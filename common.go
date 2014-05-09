package citadel

type (
	Host struct {
		Name        string  `json:"name,omitempty"`
		IPAddress   string  `json:"ip_address,omitempty"`
		Cpus        int     `json:"cpus,omitempty"`
		TotalMemory float64 `json:"total_memory,omitempty"`
		Disks       []*Disk `json:"disks,omitempty"`
	}

	Disk struct {
		Path      string  `json:"path,omitempty"`
		Used      float64 `json:"used,omitempty"`
		Free      float64 `json:"free,omitempty"`
		Files     float64 `json:"files,omitempty"`
		Available float64 `json:"available,omitempty"`
		Total     float64 `json:"total,omitempty"`
	}

	ContainerGroup struct {
		Name      string  `json:"name,omitempty"`
		Instances int     `json:"instances,omitempty"`
		Cpu       float64 `json:"cpu,omitempty"`
		Memory    float64 `json:"memory,omitempty"`
		Status    string  `json:"status,omitempty"`
	}

	Container struct {
		Host   string  `json:"host,omitempty"`
		Name   string  `json:"name,omitempty"`
		Image  string  `json:"image,omitempty"`
		Cpu    float64 `json:"cpu,omitempty"`
		Memory float64 `json:"memory,omitempty"`
	}
)
