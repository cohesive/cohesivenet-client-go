package cohesivenetv1

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
	FirewallRules []FirewallRule `json:"response"`
}

type EndpointResponse struct {
	Endpoints interface{} `json:"response"`
}

type FirewallRule struct {
	ID   string `json:"id"`
	Rule string `json:"rule"`
}

type Endpoint struct {
	Name                    string `json:"name"`
	Description             string `json:"description"`
	Ipaddress               string `json:"ipaddress"`
	Secret                  string `json:"secret"`
	Pfs                     bool   `json:"pfs"`
	Ike_version             int    `json:"ike_version"`
	Nat_t_enabled           bool   `json:"nat_t_enabled"`
	Extra_config            string `json:"extra_config"`
	Vpn_type                string `json:"vpn_type"`
	Route_based_int_address string `json:"route_based_int_address"`
	Route_based_local       string `json:"route_based_local"`
	Route_based_remote      string `json:"route_based_remote"`
}

type NewEndpoint struct {
	Response struct {
		ID          int      `json:"id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Ipaddress   string   `json:"ipaddress"`
		Pfs         bool     `json:"pfs"`
		IkeVersion  int      `json:"ike_version"`
		NatTEnabled bool     `json:"nat_t_enabled"`
		ExtraConfig []string `json:"extra_config"`
		Tunnels     struct {
			Num3 struct {
				ID           int    `json:"id"`
				LocalSubnet  string `json:"local_subnet"`
				RemoteSubnet string `json:"remote_subnet"`
				EndpointID   int    `json:"endpoint_id"`
				Enabled      bool   `json:"enabled"`
				Description  string `json:"description"`
			} `json:"3"`
		} `json:"tunnels"`
		TrafficPairs struct {
		} `json:"traffic_pairs"`
		BgpPeers struct {
		} `json:"bgp_peers"`
		Type                 string `json:"type"`
		VpnType              string `json:"vpn_type"`
		RouteBasedIntAddress string `json:"route_based_int_address"`
		RouteBasedLocal      string `json:"route_based_local"`
		RouteBasedRemote     string `json:"route_based_remote"`
		Psk                  string `json:"psk"`
	} `json:"response"`
}

type Route struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Advertise   bool   `json:"advertise,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	Editable    bool   `json:"editable,omitempty"`
	Cidr        string `json:"cidr,omitempty"`
	Gateway     string `json:"gateway,omitempty"`
	Netmask     string `json:"netmask,omitempty"`
	Table       string `json:"table,omitempty"`
	Interface   string `json:"interface,omitempty"`
	Metric      int    `json:"metric,omitempty"`
	Tunnel      int    `json:"tunnel,omitempty"`
}

type RouteResponse struct {
	Routes []Route `json:"response"`
}
