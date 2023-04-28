package main

type NetworkConfig struct {
	Network Network `yaml:"network"`
}

type Network struct {
	Ethernets map[string]Ethernet `yaml:"ethernets"`
	Renderer  string              `yaml:"renderer"`
	Version   int                 `yaml:"version"`
}

type Ethernet struct {
	Addresses   []string `yaml:"addresses"`
	Dhcp4       string   `yaml:"dhcp4"`
	Nameservers struct {
		Addresses []string `yaml:"addresses"`
	} `yaml:"nameservers"`
	Optional string `yaml:"optional"`
}

type address struct {
	IFName  string `json:"ifname" binding:"required"`
	Ipv4    string `json:"ipv4" binding:"required,ip4_addr"`
	Netmask int    `json:"netmask" binding:"required,lte=32"`
}
