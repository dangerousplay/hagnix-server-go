package service

import (
	"github.com/deckarep/golang-set"
	"github.com/ivahaev/go-logger"
	"github.com/kataras/iris/core/errors"
	"hagnix-server-go1/config"
	"hagnix-server-go1/routes/modelxml"
	"net"
	"strconv"
	"strings"
)

var server = &ServerService{}
var servers = mapset.NewSet()

type ServerService struct{}

func (server *ServerService) GetServers() []modelxml.ServerItemXML {
	var xml []modelxml.ServerItemXML

	for v := range servers.Iter() {
		t, ok := v.(modelxml.ServerItemXML)

		if !ok {
			logger.Warnf("Invalid cast to ServerItemXML: %s", v)
			continue
		}

		xml = append(xml, t)
	}

	return xml
}

func getUsage(address string) (float64, error) {
	hostPort := strings.Split(address, ":")

	var port = ":2050"

	if len(hostPort) > 1 {
		port = ""
	}

	conn, err := net.Dial("tcp", address+port)

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	_, err = conn.Write([]byte{0x4d, 0x61, 0x64, 0x65, 0xff})

	if err != nil {
		return 0, err
	}

	buf := make([]byte, 256)

	readed, err := conn.Read(buf)

	if err != nil {
		return 0, err
	}

	decoded := strings.TrimSpace(string(buf[:readed]))

	str := strings.Split(decoded, ":")

	if len(str) < 2 {
		return 0, errors.New("invalid response from server: " + string(buf))
	}

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

func updateServers() {
	serversList := config.GetConfig().ServerConfig.Servers

	for _, v := range serversList {
		usage, err := getUsage(v.Address)

		if err != nil {
			logger.Warnf("Can't get usage of server: %s", v.Name)
			logger.Warn(err)
			continue
		}

		servers.Add(modelxml.ServerItemXML{
			Name:      v.Name,
			AdminOnly: "false",
			Usage:     usage,
			DNS:       v.Address,
		})
	}

}

func GetServerService() *ServerService {
	return server
}
