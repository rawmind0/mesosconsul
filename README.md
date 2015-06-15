# mesosconsul

This package provides the `mesosconsul` package which attempts to provide 
programmatic and integrated access to consul, docker and mesos systems.

## Motivation

In a docker, mesos and consul (registrator) environment, due to mesos reubication service,
the consul catalog and mesos tasks, becomes unsync, leaving registered services that
are not running.

This package starts as an aim to resolve this issue. 

## Features

The package implemtens the folowing features:
- List own services regitered in consul.
- List own MESOS_TAKS_ID dockers running.
- List own services regitered and not running.
- Deregister not running services.

##TODO
- Implements running daemon mode.
- Implements full environment running mode.
- .....

## Usage

In examples folder, you can see main.go, an example to use the package.

The command has two modes of operation: 
- Agent mode (-a): To run in each mesos-slave. The command connects to the local consul agent and to the docker socket.
- Remote server mode (-s server): To run from remote node. The command connects to the server consul and docker agent.

The command implements two actions:
- List (-l): List registered, running and not running services.
- Sync (-sync): Deregister in consul, not running services.

environmet, domain, consulPort and dockerPort are used to form docker and consul urls as:
- dockerUrl tcp://server.environmet.domain:dockerPort
- consulUrl http://server.environmet.domain:consulPort

Usage: ./main -a | -s=server [-e=environment] [-d=domain] [-consulPort=port] [-dockerPort=port]) [-l] [-sync]"

```go
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
  flag.StringVar(&c.Environ, "e", "gbx0", "environment")
  flag.StringVar(&c.Domain, "d", "innotechapp.com", "domain name")
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
```
