package mesosconsul

import (
    "fmt"
    "log"
    "strings"
    "github.com/fsouza/go-dockerclient"
)

func (g *Client) GetDockerRun () {
  g.running = []map[string]string{}

  list, err := g.docker.ListContainers(docker.ListContainersOptions{All: false})
  if err != nil {
    log.Fatal(err)
  }
  g.running = g.GetInfoDockers(list)
}

func (g *Client) GetInfoDockers (r []docker.APIContainers) []map[string]string {
  var data []map[string]string

  for _, dock := range r {
    d, err := g.docker.InspectContainer(dock.ID)
    if err != nil {
      log.Fatal(err)
    }
    m := make(map[string]string)
    m["docker_ID"] = dock.ID
    m["docker_name"] = d.Name[1:]
    for _, c := range d.Config.Env {
      v := strings.Split(c, "=")
      if strings.EqualFold(v[0],"MESOS_TASK_ID") {
        m["mesos_ID"] = v[1]
      }
    }

    data = append(data,m)
  }
  return data
}

func (g *Client) DockerInit () {
  var err error

  g.docker, err = docker.NewClient(g.Conf.DockerUrl)
  if err != nil {
    log.Fatal(err)
  }
}

func (g *Client) ListDockerRun () {
  g.RefreshData("docker")

  fmt.Println("Docker_id\t\t\t\t\t\t\t\tDocker_name\t\t\t\t\tMesos_ID")
  fmt.Println("----------------------------------------------------------------------------------------------")
  for _, dock := range g.running {
    fmt.Println(dock["docker_ID"]," - ",dock["docker_name"]," - ",dock["mesos_ID"])
  }
  fmt.Println("----------------------------------------------------------------------------------------------")
  fmt.Println("RUNNING SERVICES:", len(g.running))
  fmt.Println("----------------------------------------------------------------------------------------------")
  fmt.Println("")
}

func (g *Client) saveDockerRun () {
  //var keyvalue *api.KV


}
