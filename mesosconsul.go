package mesosconsul

import (
    "fmt"
    "github.com/fsouza/go-dockerclient"
    "github.com/hashicorp/consul/api"
)

type Config struct {
  Agent   bool
  Server  string
  Environ string
  Domain  string
  ConsulPort  string
  ConsulUrl   string
  DockerPort  string
  DockerUrl   string
  List  bool
  Sync  bool
}

type Client struct {
  host    string
  Conf    Config
  docker  *docker.Client
  consul  *api.Client
  catalog *api.Catalog
  agent *api.Agent
  running []map[string]string
  registered *api.CatalogNode
  services map[string]*api.AgentService
  notrun []string

}

type node struct {
  Node    string
  Address string
}

func (g *Client) Init () {
  g.ConsulInit()
  g.DockerInit()

}

func (g *Client) RefreshData (s string) {

  switch s {
    case "docker" :
      g.GetDockerRun()
      fmt.Println("Refreshing docker data")
    case "consul" :
      g.GetConsulReg()
      fmt.Println("Refreshing consul data")
    case "notrun" :
      g.RefreshData("docker")
      g.RefreshData("consul")
      g.GetNotRunning()
      fmt.Println("Refreshing merge data")
    case "all" :
      g.RefreshData("notrun")
  }
}

func (g *Client) PutConf (c Config) {

  g.Conf.Agent = c.Agent
  g.Conf.Server = c.Server
  g.Conf.Environ = c.Environ
  g.Conf.Domain = c.Domain
  g.Conf.ConsulPort = c.ConsulPort
  g.Conf.DockerPort = c.DockerPort
  g.Conf.List = c.List
  g.Conf.Sync = c.Sync
  g.Conf.DockerUrl = c.DockerUrl
  g.Conf.ConsulUrl = c.ConsulUrl
  g.host = c.Server

  if ! g.Conf.Agent {
    fmt.Println("Running in remote server mode")
  } else {
    fmt.Println("Running in agent mode")
  }

  fmt.Println("Conecting to:")
  fmt.Println("- Consul server", g.Conf.ConsulUrl)
  fmt.Println("- Docker server", g.Conf.DockerUrl)
}
