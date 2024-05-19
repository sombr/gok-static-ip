package main;

import "log"
import netlink "github.com/vishvananda/netlink"
import "net"
import "os"

func main() {
    lnk, err := netlink.LinkByName(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    ipConfig := &netlink.Addr{IPNet: &net.IPNet{
        IP:   net.ParseIP(os.Args[2]),
        Mask: net.CIDRMask(30, 32),
    }}

    if err = netlink.AddrAdd(lnk, ipConfig); err != nil {
        log.Fatal(err)
    }
}
