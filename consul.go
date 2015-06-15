package mesosconsul

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "github.com/hashicorp/consul/api"
)

func (g *Client) GetConsulReg () {
  var err error

  g.services = map[string]*api.AgentService{}

  g.services, err = g.agent.Services()
  if err != nil {
    log.Fatal(err)
  }
}

func (g *Client) ConsulInit () {
  var err error
  consul_conf := &api.Config{
    Address:    g.Conf.ConsulUrl,
    Scheme:     "http",
    HttpClient: http.DefaultClient,
  }

  g.consul, err = api.NewClient(consul_conf)
  if err != nil {
    log.Fatal(err)
  }

  //g.catalog = g.consul.Catalog()
  g.agent = g.consul.Agent()
}

func (g *Client) ConsulRegIsRunning (reg string) bool{
  ret := false

  elems := len(g.running)
  for i := 0; i < elems; i += 1 {
    dock := g.running[i]
    if strings.Contains(reg, dock["docker_name"]) {
      ret = true
    }
    if strings.EqualFold(reg, "dockerhost"){
      ret = true
    }
    if strings.EqualFold(reg, "cassandra") {
      ret = true
    }
    if strings.EqualFold(reg,dock["mesos_ID"]) {
      ret = true
    }
    if ret == true {
      i = elems
    }
  }
  return ret
}

func (g *Client) ListConsulReg () {
  g.RefreshData("consul")

  fmt.Println("Service_ID\t\t\t\t\t\t\tService_name")
  fmt.Println("----------------------------------------------------------------------------------------------")
  for _, service := range g.services {
    fmt.Println(service.ID," - ",service.Service)
  }
  fmt.Println("----------------------------------------------------------------------------------------------")
  fmt.Println("REGISTERED SERVICES:", len(g.services))
  fmt.Println("----------------------------------------------------------------------------------------------")
  fmt.Println("")
}
