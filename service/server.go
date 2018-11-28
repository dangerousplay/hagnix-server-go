package service

import (
	"github.com/InVisionApp/go-logger"
	"hagnix-server-go1/config"
	"hagnix-server-go1/routes/modelxml"
	"net"
	"strconv"
	"strings"
)

var logger = log.NewSimple()
var server = &ServerService{}

type ServerService struct{}

func (server *ServerService) GetServers() []modelxml.ServerItemXML {
	servers := config.GetConfig().ServerConfig.Servers

	var serverXML []modelxml.ServerItemXML

	for _, v := range servers {
		usage, err := getUsage(v.Address)

		if err != nil {
			logger.Warnf("Can't get usage of server: %f", usage)
			continue
		}

		serverXML = append(serverXML, modelxml.ServerItemXML{
			Name:      v.Name,
			AdminOnly: "false",
			Usage:     usage,
			DNS:       v.Address,
		})
	}

	return serverXML
}

func getUsage(address string) (float64, error) {
	hostPort := strings.Split(address, ":")

	var port = ":2050"

	if len(hostPort) > 0 {
		port = ""
	}

	conn, err := net.Dial("tcp", address+port)
	defer conn.Close()

	if err != nil {
		return 0, err
	}

	_, err = conn.Write([]byte{0x4d, 0x61, 0x64, 0x65, 0xff})

	if err != nil {
		return 0, err
	}

	var buffer []byte
	conn.Read(buffer)

	str := strings.Split(string(buffer), ":")

	number1, err := strconv.ParseFloat(str[1], 32)

	if err != nil {
		return 0, err
	}

	number2, err := strconv.ParseFloat(str[0], 32)

	if err != nil {
		return 0, err
	}

	return number1 / number2, err
}

func GetServerService() *ServerService {
	return server
}
