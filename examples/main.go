package main

import (
    "log"
    "fmt"
    "os"
    "bytes"
    "flag"
    "mesosConsul"
)

func CheckArgs () mesosconsul.Config {
  var err error
  var c mesosconsul.Config

  flag.BoolVar(&c.Agent, "a", false, "Agent execution")
  flag.StringVar(&c.Server, "s", "", "Server")
  flag.StringVar(&c.Environ, "e", "env1", "environment")
  flag.StringVar(&c.Domain, "d", "domain1", "domain name")
  flag.StringVar(&c.ConsulPort, "consul-port", "8500", "Consul port")
  flag.StringVar(&c.DockerPort, "docker-port", "2375", "Docker port")
  flag.BoolVar(&c.List, "l", false, "List services")
  flag.BoolVar(&c.Sync, "sync", false, "Sync services")

  flag.Parse()

  if ! c.List && ! c.Sync {
    c.List = true
  }

  if ! c.Agent {
    if c.Server == "" {
      fmt.Println("Usage: ", os.Args[0],"-a | -s=<server> [-e=<environment>] [-d=<domain>] [-consulPort=<port>] [-dockerPort=<port>]) [-l] [-sync]")
      os.Exit(1)
    } else {

      buffer2 := bytes.NewBufferString("tcp://")
      buffer2.WriteString(c.Server)
      buffer2.WriteString(".")
      buffer2.WriteString(c.Environ)
      buffer2.WriteString(".")
      buffer2.WriteString(c.Domain)
      buffer2.WriteString(":")
      buffer2.WriteString(c.DockerPort)
      c.DockerUrl = buffer2.String()
    }
  } else {

    c.Server, err = os.Hostname()
    if err != nil {
      log.Fatal(err)
    }

    c.DockerUrl = "unix:///var/run/docker.sock"

  }

  buffer := bytes.NewBufferString("")
  buffer.WriteString(c.Server)
  buffer.WriteString(".")
  buffer.WriteString(c.Environ)
  buffer.WriteString(".")
  buffer.WriteString(c.Domain)
  buffer.WriteString(":")
  buffer.WriteString(c.ConsulPort)
  c.ConsulUrl = buffer.String()

  return c

}

func Run (g mesosconsul.Client) {

  if g.Conf.List {
    g.ListConsulReg()
    g.ListDockerRun()
    g.ListNotRunning()
  }

  if g.Conf.Sync {
    g.SyncNotRunning()
  }
}

func main() {

  cli := mesosconsul.Client{}

  cli.PutConf(CheckArgs())
  cli.Init()

  Run(cli)

}
