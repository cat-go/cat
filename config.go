package cat

import (
	"encoding/xml"
	"net"
	"os"
	"strings"
)

type Config struct {
	domain   string
	hostname string
	env      string
	ip       string
	ipHex    string

	httpServerPort      int
	httpServerAddresses []serverAddress

	serverAddress []serverAddress
}

type XMLConfig struct {
	Name    xml.Name         `xml:"config"`
	Servers XMLConfigServers `xml:"servers"`
}

type XMLConfigServers struct {
	Servers []XMLConfigServer `xml:"server"`
}

type XMLConfigServer struct {
	Host string `xml:"ip,attr"`
	Port int    `xml:"port,attr"`
}

var config = Config{
	domain:   defaultAppKey,
	hostname: defaultHostname,
	env:      defaultEnv,
	ip:       defaultIp,
	ipHex:    defaultIpHex,

	httpServerPort:      8080,
	httpServerAddresses: []serverAddress{},

	serverAddress: []serverAddress{},
}

func (config *Config) Init(opts *Options) (err error) {
	config.domain = opts.AppId

	defer func() {
		if err == nil {
			logger.Info("Cat has been initialized successfully with appkey: %s", config.domain)
		} else {
			logger.Error("Failed to initialize cat.")
		}
	}()

	// TODO load env.

	var ip net.IP
	if ip, err = getLocalhostIp(); err != nil {
		config.ip = defaultIp
		config.ipHex = defaultIpHex
		logger.Warning("Error while getting local ip, using default ip: %s", defaultIp)
	} else {
		config.ip = ip2String(ip)
		config.ipHex = ip2HexString(ip)
		logger.Info("Local ip has been configured to %s", config.ip)
	}

	if config.hostname, err = os.Hostname(); err != nil {
		config.hostname = defaultHostname
		logger.Warning("Error while getting hostname, using default hostname: %s", defaultHostname)
	} else {
		logger.Info("Hostname has been configured to %s", config.hostname)
	}

	//转为参数配置
	config.httpServerPort = opts.HttpPort
	var serversArray []string = strings.Split(opts.ServerAddr, ",")
	for _, s := range serversArray {
		config.httpServerAddresses = append(config.httpServerAddresses, serverAddress{
			host: s,
			port: opts.HttpPort,
		})
		config.serverAddress = append(config.serverAddress, serverAddress{
			host: s,
			port: opts.Port,
		})
	}

	return
}
