package cohesivenet

type Idobject struct {
	Local_subnet  string `json:"local_subnet"`
	Remote_subnet string `json:"remote_subnet"`
	Endpointid    int    `json:"endpointid"`
	Endpoint_name string `json:"endpoint_name"`
	Active        bool   `json:"active"`
	Enabled       bool   `json:"enabled"`
	Descripton    int    `json:"description"`
	Connected     bool   `json:"connected"`
}

type Endpoints struct {
	Id Idobject `json:"1"`
}

type Intresponse struct {
	Res Endpoints `json:"response"`
}

type ContainerNetwork struct {
	Network string `json:"network"`
	Running bool   `json:"running"`
}

type ContainerSystem struct {
	Response ContainerNetwork `json:"response"`
}

type Config struct {
	PrivateIp        string `json:"private_ipaddress"`
	PublicIp         string `json:"public_ipaddress"`
	SubnetGateway    string `json:"subnet_gateway"`
	TopologyChecksum string `json:"topology_checksum"`
	Vns3Version      string `json:"vns3_version"`
	TopologyName     string `json:"topology_name"`
	NtpHosts         string `json:"ntp_hosts"`
	Licensed         bool   `json:"licensed"`
	Peered           bool   `json:"peered"`
	Asn              int    `json:"asn"`
	ManagerId        int    `json:"manager_id"`
	OverlayIpaddress string `json:"overlay_ipaddress"`
	SecurityToken    string `json:"security_token"`
}

type ConfigResponse struct {
	Config Config `json:"response"`
}

type FirewallResponse struct {
	Firewall [][]interface{} `json:"response"`
}

type RouteResponse struct {
	Response interface{} `json:"response"`
}

type EndpointResponse struct {
	Endpoints interface{} `json:"response"`
}
