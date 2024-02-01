package cohesivenetv1

import "time"

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

type Tunnel struct {
	Remote_Subnet  string `json:"remote_subnet,omitempty"`
	Local_Subnet   string `json:"local_subnet,omitempty"`
	Ping_Ipaddress string `json:"ping_ipaddress,omitempty"`
	Ping_Interval  int    `json:"ping_interval,omitempty"`
	Ping_Interface string `json:"ping_interface,omitempty"`
	Enabled        bool   `json:"enabled,omitempty"`
	Description    string `json:"description,omitempty"`
}

type TrafficPair struct {
	ID             int    `json:"id,omitempty"`
	Remote_Subnet  string `json:"remote_subnet,omitempty"`
	Local_Subnet   string `json:"local_subnet,omitempty"`
	Ping_Ipaddress string `json:"ping_ipaddress,omitempty"`
	Ping_Interval  int    `json:"ping_interval,omitempty"`
	Ping_Interface string `json:"ping_interface,omitempty"`
	Enabled        bool   `json:"enabled,omitempty"`
	Description    string `json:"description,omitempty"`
}

type NewTunnelResponse struct {
	NewTunnels []NewTunnel `json:"response"`
}
type NewTunnel struct {
	ID            int    `json:"id,omitempty"`
	LocalSubnet   string `json:"local_subnet,omitempty"`
	RemoteSubnet  string `json:"remote_subnet,omitempty"`
	EndpointID    int    `json:"endpoint_id,omitempty"`
	Enabled       bool   `json:"enabled,omitempty"`
	Description   string `json:"description,omitempty"`
	PingIpaddress string `json:"ping_ipaddress,omitempty"`
	PingInterface string `json:"ping_interface,omitempty"`
	PingInterval  int    `json:"ping_interval,omitempty"`
}

type Route struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Advertise   bool   `json:"advertise"`
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

type EbgpPeer struct {
	ID                          int    `json:"id,omitempty"`
	Ipaddress                   string `json:"ipaddress,omitempty"`
	Asn                         int    `json:"asn,omitempty"`
	LocalAsnAlias               int    `json:"local_asn_alias,omitempty"`
	AccessList                  string `json:"access_list,omitempty"`
	AddNetworkDistanceHops      int    `json:"add_network_distance_hops,omitempty"`
	BgpPassword                 string `json:"bgp_password,omitempty"`
	AddNetworkDistance          bool   `json:"add_network_distance,omitempty"`
	AddNetworkDistanceDirection string `json:"add_network_distance_direction,omitempty"`
}

type CreateImageResponse struct {
	NewImage struct {
		Status      string `json:"status,omitempty"`
		Import_uuid string `json:"import_uuid,omitempty"`
	} `json:"response,omitempty"`
}

type PluginImage struct {
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`
	Buildurl    string `json:"buildurl,omitempty"`
	Localbuild  string `json:"localbuild,omitempty"`
	Localimage  string `json:"localimage,omitempty"`
	Imagefile   string `json:"imagefile,omitempty"`
	Buildfile   string `json:"buildfile,omitempty"`
	Description string `json:"description,omitempty"`
	Import_uuid string `json:"import_uuid,omitempty"`
}

type ImageResponse struct {
	Images []Image `json:"response,omitempty"`
}

type Image struct {
	ID        string `json:"id,omitempty"`
	ImageName string `json:"image_name,omitempty"`
	TagName   string `json:"tag_name,omitempty"`
	Status    string `json:"status,omitempty"`
	//StatusMsg string `json:"status_msg,omitempty"`
	StatusMsg interface{} `json:"status_msg,omitempty"`
	ImportID  string      `json:"import_id,omitempty"`
	//Created   time.Time `json:"created,omitempty"`
	Created     string `json:"created,omitempty"`
	Description string `json:"description,omitempty"`
	//Comment     string `json:"comment,omitempty"`
	ImportUUID      string      `json:"import_uuid,omitempty"`
	Comment         interface{} `json:"comment,omitempty"`
	ContainerConfig interface{} `json:"container_config,omitempty"`
}

type CreatePluginInstance struct {
	UUID        string `json:"uuid,omitempty"`
	ImageUUID   string `json:"image_uuid,omitempty"`
	Name        string `json:"name,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Description string `json:"description,omitempty"`
	Command     string `json:"command,omitempty"`
	Environment string `json:"environment,omitempty"`
}

type CreatePluginInstanceResponse struct {
	Instance struct {
		UUID             string `json:"uuid,omitempty"`
		ContainerStarted bool   `json:"container_started,omitempty"`
		IPAddress        string `json:"ip_address,omitempty"`
		Status           string `json:"status,omitempty"`
	} `json:"response,omitempty"`
}

type GetPluginInstances struct {
	Instances struct {
		Containers []struct {
			ID      string        `json:"Id,omitempty"`
			Created time.Time     `json:"Created,omitempty"`
			Path    string        `json:"Path,omitempty"`
			Args    []interface{} `json:"Args,omitempty"`
			State   struct {
				Status     string    `json:"Status,omitempty"`
				Running    bool      `json:"Running,omitempty"`
				Paused     bool      `json:"Paused,omitempty"`
				Restarting bool      `json:"Restarting,omitempty"`
				OOMKilled  bool      `json:"OOMKilled,omitempty"`
				Dead       bool      `json:"Dead,omitempty"`
				Pid        int       `json:"Pid,omitempty"`
				ExitCode   int       `json:"ExitCode,omitempty"`
				Error      string    `json:"Error,omitempty"`
				StartedAt  time.Time `json:"StartedAt,omitempty"`
				FinishedAt time.Time `json:"FinishedAt,omitempty"`
			} `json:"State,omitempty"`
			Image           string      `json:"Image,omitempty"`
			ResolvConfPath  string      `json:"ResolvConfPath,omitempty"`
			HostnamePath    string      `json:"HostnamePath,omitempty"`
			HostsPath       string      `json:"HostsPath,omitempty"`
			LogPath         string      `json:"LogPath,omitempty"`
			Name            string      `json:"Name,omitempty"`
			RestartCount    int         `json:"RestartCount,omitempty"`
			Driver          string      `json:"Driver,omitempty"`
			Platform        string      `json:"Platform,omitempty"`
			MountLabel      string      `json:"MountLabel,omitempty"`
			ProcessLabel    string      `json:"ProcessLabel,omitempty"`
			AppArmorProfile string      `json:"AppArmorProfile,omitempty"`
			ExecIDs         interface{} `json:"ExecIDs,omitempty"`
			HostConfig      struct {
				Binds           interface{} `json:"Binds,omitempty"`
				ContainerIDFile string      `json:"ContainerIDFile,omitempty"`
				LogConfig       struct {
					Type   string `json:"Type,omitempty"`
					Config struct {
					} `json:"Config,omitempty"`
				} `json:"LogConfig,omitempty"`
				NetworkMode  string `json:"NetworkMode,omitempty"`
				PortBindings struct {
				} `json:"PortBindings,omitempty"`
				RestartPolicy struct {
					Name              string `json:"Name,omitempty"`
					MaximumRetryCount int    `json:"MaximumRetryCount,omitempty"`
				} `json:"RestartPolicy,omitempty"`
				AutoRemove           bool          `json:"AutoRemove,omitempty"`
				VolumeDriver         string        `json:"VolumeDriver,omitempty"`
				VolumesFrom          interface{}   `json:"VolumesFrom,omitempty"`
				CapAdd               interface{}   `json:"CapAdd,omitempty"`
				CapDrop              interface{}   `json:"CapDrop,omitempty"`
				CgroupnsMode         string        `json:"CgroupnsMode,omitempty"`
				DNS                  []interface{} `json:"Dns,omitempty"`
				DNSOptions           []interface{} `json:"DnsOptions,omitempty"`
				DNSSearch            []interface{} `json:"DnsSearch,omitempty"`
				ExtraHosts           interface{}   `json:"ExtraHosts,omitempty"`
				GroupAdd             interface{}   `json:"GroupAdd,omitempty"`
				IpcMode              string        `json:"IpcMode,omitempty"`
				Cgroup               string        `json:"Cgroup,omitempty"`
				Links                interface{}   `json:"Links,omitempty"`
				OomScoreAdj          int           `json:"OomScoreAdj,omitempty"`
				PidMode              string        `json:"PidMode,omitempty"`
				Privileged           bool          `json:"Privileged,omitempty"`
				PublishAllPorts      bool          `json:"PublishAllPorts,omitempty"`
				ReadonlyRootfs       bool          `json:"ReadonlyRootfs,omitempty"`
				SecurityOpt          interface{}   `json:"SecurityOpt,omitempty"`
				UTSMode              string        `json:"UTSMode,omitempty"`
				UsernsMode           string        `json:"UsernsMode,omitempty"`
				ShmSize              int           `json:"ShmSize,omitempty"`
				Runtime              string        `json:"Runtime,omitempty"`
				ConsoleSize          []int         `json:"ConsoleSize,omitempty"`
				Isolation            string        `json:"Isolation,omitempty"`
				CPUShares            int           `json:"CpuShares,omitempty"`
				Memory               int           `json:"Memory,omitempty"`
				NanoCpus             int           `json:"NanoCpus,omitempty"`
				CgroupParent         string        `json:"CgroupParent,omitempty"`
				BlkioWeight          int           `json:"BlkioWeight,omitempty"`
				BlkioWeightDevice    []interface{} `json:"BlkioWeightDevice,omitempty"`
				BlkioDeviceReadBps   interface{}   `json:"BlkioDeviceReadBps,omitempty"`
				BlkioDeviceWriteBps  interface{}   `json:"BlkioDeviceWriteBps,omitempty"`
				BlkioDeviceReadIOps  interface{}   `json:"BlkioDeviceReadIOps,omitempty"`
				BlkioDeviceWriteIOps interface{}   `json:"BlkioDeviceWriteIOps,omitempty"`
				CPUPeriod            int           `json:"CpuPeriod,omitempty"`
				CPUQuota             int           `json:"CpuQuota,omitempty"`
				CPURealtimePeriod    int           `json:"CpuRealtimePeriod,omitempty"`
				CPURealtimeRuntime   int           `json:"CpuRealtimeRuntime,omitempty"`
				CpusetCpus           string        `json:"CpusetCpus,omitempty"`
				CpusetMems           string        `json:"CpusetMems,omitempty"`
				Devices              []interface{} `json:"Devices,omitempty"`
				DeviceCgroupRules    interface{}   `json:"DeviceCgroupRules,omitempty"`
				DeviceRequests       interface{}   `json:"DeviceRequests,omitempty"`
				KernelMemory         int           `json:"KernelMemory,omitempty"`
				KernelMemoryTCP      int           `json:"KernelMemoryTCP,omitempty"`
				MemoryReservation    int           `json:"MemoryReservation,omitempty"`
				MemorySwap           int           `json:"MemorySwap,omitempty"`
				MemorySwappiness     interface{}   `json:"MemorySwappiness,omitempty"`
				OomKillDisable       bool          `json:"OomKillDisable,omitempty"`
				PidsLimit            interface{}   `json:"PidsLimit,omitempty"`
				Ulimits              interface{}   `json:"Ulimits,omitempty"`
				CPUCount             int           `json:"CpuCount,omitempty"`
				CPUPercent           int           `json:"CpuPercent,omitempty"`
				IOMaximumIOps        int           `json:"IOMaximumIOps,omitempty"`
				IOMaximumBandwidth   int           `json:"IOMaximumBandwidth,omitempty"`
				MaskedPaths          []string      `json:"MaskedPaths,omitempty"`
				ReadonlyPaths        []string      `json:"ReadonlyPaths,omitempty"`
			} `json:"HostConfig,omitempty"`
			GraphDriver struct {
				Data struct {
					LowerDir  string `json:"LowerDir,omitempty"`
					MergedDir string `json:"MergedDir,omitempty"`
					UpperDir  string `json:"UpperDir,omitempty"`
					WorkDir   string `json:"WorkDir,omitempty"`
				} `json:"Data,omitempty"`
				Name string `json:"Name,omitempty"`
			} `json:"GraphDriver,omitempty"`
			Mounts []interface{} `json:"Mounts,omitempty"`
			Config struct {
				Hostname     string      `json:"Hostname,omitempty"`
				Domainname   string      `json:"Domainname,omitempty"`
				User         string      `json:"User,omitempty"`
				AttachStdin  bool        `json:"AttachStdin,omitempty"`
				AttachStdout bool        `json:"AttachStdout,omitempty"`
				AttachStderr bool        `json:"AttachStderr,omitempty"`
				Tty          bool        `json:"Tty,omitempty"`
				OpenStdin    bool        `json:"OpenStdin,omitempty"`
				StdinOnce    bool        `json:"StdinOnce,omitempty"`
				Env          []string    `json:"Env,omitempty"`
				Cmd          []string    `json:"Cmd,omitempty"`
				Image        string      `json:"Image,omitempty"`
				Volumes      interface{} `json:"Volumes,omitempty"`
				WorkingDir   string      `json:"WorkingDir,omitempty"`
				Entrypoint   interface{} `json:"Entrypoint,omitempty"`
				OnBuild      interface{} `json:"OnBuild,omitempty"`
				Labels       struct {
				} `json:"Labels,omitempty"`
			} `json:"Config,omitempty"`
			NetworkSettings struct {
				Bridge                 string `json:"Bridge,omitempty"`
				SandboxID              string `json:"SandboxID,omitempty"`
				HairpinMode            bool   `json:"HairpinMode,omitempty"`
				LinkLocalIPv6Address   string `json:"LinkLocalIPv6Address,omitempty"`
				LinkLocalIPv6PrefixLen int    `json:"LinkLocalIPv6PrefixLen,omitempty"`
				Ports                  struct {
				} `json:"Ports,omitempty"`
				SandboxKey             string      `json:"SandboxKey,omitempty"`
				SecondaryIPAddresses   interface{} `json:"SecondaryIPAddresses,omitempty"`
				SecondaryIPv6Addresses interface{} `json:"SecondaryIPv6Addresses,omitempty"`
				EndpointID             string      `json:"EndpointID,omitempty"`
				Gateway                string      `json:"Gateway,omitempty"`
				GlobalIPv6Address      string      `json:"GlobalIPv6Address,omitempty"`
				GlobalIPv6PrefixLen    int         `json:"GlobalIPv6PrefixLen,omitempty"`
				IPAddress              string      `json:"IPAddress,omitempty"`
				IPPrefixLen            int         `json:"IPPrefixLen,omitempty"`
				IPv6Gateway            string      `json:"IPv6Gateway,omitempty"`
				MacAddress             string      `json:"MacAddress,omitempty"`
				Networks               struct {
					CohesiveNet struct {
						IPAMConfig struct {
							IPv4Address string `json:"IPv4Address,omitempty"`
						} `json:"IPAMConfig,omitempty"`
						Links               interface{} `json:"Links,omitempty"`
						Aliases             []string    `json:"Aliases,omitempty"`
						NetworkID           string      `json:"NetworkID,omitempty"`
						EndpointID          string      `json:"EndpointID,omitempty"`
						Gateway             string      `json:"Gateway,omitempty"`
						IPAddress           string      `json:"IPAddress,omitempty"`
						IPPrefixLen         int         `json:"IPPrefixLen,omitempty"`
						IPv6Gateway         string      `json:"IPv6Gateway,omitempty"`
						GlobalIPv6Address   string      `json:"GlobalIPv6Address,omitempty"`
						GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen,omitempty"`
						MacAddress          string      `json:"MacAddress,omitempty"`
						DriverOpts          interface{} `json:"DriverOpts,omitempty"`
					} `json:"cohesive.net,omitempty"`
				} `json:"Networks,omitempty"`
			} `json:"NetworkSettings,omitempty"`
		} `json:"containers,omitempty"`
	} `json:"response,omitempty"`
}

type InstanceResponse struct {
	Instances []Instance `json:"response,omitempty"`
}

type Instance struct {
	ID        string `json:"Id,omitempty"`
	Image     string `json:"Image,omitempty"`
	Hostname  string `json:"Hostame,omitempty"`
	IPAddress string `json:"IPAddress,omitempty"`
	Path      string `json:"Path,omitempty"`
}

type HttpsCert struct {
	Cert string `json:"cert,omitempty"`
	Key  string `json:"key,omitempty"`
}

type HttpsCertResponse struct {
	Response struct {
		Status string `json:"status"`
		UUID   string `json:"uuid"`
	} `json:"response"`
}

type UpdateTunnelResponse struct {
	Response struct {
		ID           int    `json:"id,omitempty"`
		LocalSubnet  string `json:"local_subnet,omitempty"`
		RemoteSubnet string `json:"remote_subnet,omitempty"`
		EndpointID   int    `json:"endpoint_id,omitempty"`
		Enabled      bool   `json:"enabled,omitempty"`
		Description  string `json:"description,omitempty"`
		Bounce       bool   `json:"bounce,omitempty"`
	} `json:"response,omitempty"`
}

type CreateTrafficPairResponse struct {
	Response struct {
		ID              int       `json:"id,omitempty"`
		RemoteSubnet    string    `json:"remote_subnet,omitempty"`
		LocalSubnet     string    `json:"local_subnet,omitempty"`
		PingIpaddress   string    `json:"ping_ipaddress,omitempty"`
		PingInterval    int       `json:"ping_interval,omitempty"`
		PingInterface   string    `json:"ping_interface,omitempty"`
		Enabled         bool      `json:"enabled,omitempty"`
		Description     string    `json:"description,omitempty"`
		IpsecEndpointID int       `json:"ipsec_endpoint_id,omitempty"`
		CreatedAt       time.Time `json:"created_at,omitempty"`
		UpdatedAt       time.Time `json:"updated_at,omitempty"`
	} `json:"response,omitempty"`
}

type Alert struct {
	Response struct {
		ID               int           `json:"id,omitempty"`
		Name             string        `json:"name,omitempty"`
		URL              string        `json:"url,omitempty"`
		Enabled          bool          `json:"enabled,omitempty"`
		WebhookID        int           `json:"webhook_id,omitempty"`
		CreatedAt        string        `json:"created_at,omitempty"`
		UpdatedAt        string        `json:"updated_at,omitempty"`
		Events           []interface{} `json:"events,omitempty"`
		CustomProperties []interface{} `json:"custom_properties,omitempty"`
	} `json:"response,omitempty"`
}

type NewAlert struct {
	ID               int           `json:"id,omitempty"`
	Name             string        `json:"name,omitempty"`
	Url              string        `json:"url,omitempty"`
	Enabled          bool          `json:"enabled,omitempty"`
	WebhookId        int           `json:"webhook_id,omitempty"`
	Events           []interface{} `json:"events,omitempty"`
	CustomProperties interface{}   `json:"custom_properties,omitempty"`
}
type Webhook struct {
	Response struct {
		ID               int       `json:"id,omitempty"`
		Name             string    `json:"name,omitempty"`
		ValidateCert     bool      `json:"validate_cert,omitempty"`
		CreatedAt        time.Time `json:"created_at,omitempty"`
		UpdatedAt        time.Time `json:"updated_at,omitempty"`
		System           bool      `json:"system,omitempty"`
		URL              string    `json:"url,omitempty"`
		Body             string    `json:"body,omitempty"`
		CustomProperties []struct {
			Name        string `json:"name,omitempty"`
			Value       string `json:"value,omitempty"`
			Description string `json:"description,omitempty"`
		} `json:"custom_properties,omitempty"`
		Headers    interface{}   `json:"headers,omitempty"`
		Parameters interface{}   `json:"parameters,omitempty"`
		Events     []interface{} `json:"events,omitempty"`
	} `json:"response,omitempty"`
}

type NewWebhook struct {
	Name             string        `json:"name,omitempty"`
	ValidateCert     bool          `json:"validate_cert,omitempty"`
	URL              string        `json:"url,omitempty"`
	Body             string        `json:"body,omitempty"`
	Headers          interface{}   `json:"headers,omitempty"`
	Parameters       interface{}   `json:"parameters,omitempty"`
	Events           []interface{} `json:"events,omitempty"`
	CustomProperties interface{}   `json:"custom_properties,omitempty"`
}
