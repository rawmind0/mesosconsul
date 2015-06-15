package mesosconsul

import (
    "net"
    "fmt"
    "bytes"
    "strings"
)

type DnsReg struct {
  name  string
  port  uint16
}

func Dns(n string) []DnsReg{
  var dnslist []DnsReg
  var dns     DnsReg

  data := strings.Split(n, ".")

  buffer := bytes.NewBufferString("")

  for i := 1; i < len(data); i++ {
      buffer.WriteString(data[i])

      if i != len(data)-1 {
          buffer.WriteString(".")
      }
  }

  dnsdomain := buffer.String()
  dnsname := data[0]

  _, srv, err := net.LookupSRV(dnsname, "tcp",dnsdomain)

  if err != nil {
      // handle error
  }

  for i := range srv {
    dns.name = srv[i].Target
    dns.port = srv[i].Port
    dnslist = append(dnslist,dns)
    fmt.Printf("%s:%d\n", srv[i].Target, srv[i].Port)
  }

  return dnslist
}
