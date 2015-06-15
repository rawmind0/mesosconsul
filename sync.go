package mesosconsul

import (
    "fmt"
)

func (g *Client) GetNotRunning () {
  g.notrun = []string{}

  for _, service := range g.services {
    if ! g.ConsulRegIsRunning(service.ID) {
      g.notrun = append(g.notrun, service.ID)
    }
  }
}

func (g *Client) SyncNotRunning () {
  g.RefreshData("notrun")

  fmt.Println("----------------------------------------------------------------------------------------------")
  for _, notrun := range g.notrun {
    fmt.Println("Deregistering service",notrun)
    g.agent.ServiceDeregister(notrun)
  }
  fmt.Println("----------------------------------------------------------------------------------------------")
  fmt.Println("DEREGISTERED SERVICES:", len(g.notrun))
  fmt.Println("----------------------------------------------------------------------------------------------")
  fmt.Println("")
}

func (g *Client) ListNotRunning () {
  g.RefreshData("notrun")

  fmt.Println("Service_ID")
  fmt.Println("----------------------------------------------------------------------------------------------")
  for _, notrun := range g.notrun {
    fmt.Println(notrun)
  }
  fmt.Println("----------------------------------------------------------------------------------------------")
  fmt.Println("NOT RUNNING SERVICES:", len(g.notrun))
  fmt.Println("----------------------------------------------------------------------------------------------")
  fmt.Println("")
}
