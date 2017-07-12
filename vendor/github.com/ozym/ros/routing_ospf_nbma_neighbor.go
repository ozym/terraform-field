package ros

import (
	"strconv"
)

func routingOspfNbmaNeighbors() Command {
	return Command{
		Path:    "/routing ospf nbma-neighbor",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingOspfNbmaNeighbors() ([]map[string]string, error) {
	return r.List(routingOspfNbmaNeighbors())
}

func routingOspfNbmaNeighbor(address string) Command {
	return Command{
		Path:    "/routing ospf nbma-neighbor",
		Command: "print",
		Filter: map[string]string{
			"address": address,
		},
		Detail: true,
	}
}

func (r Ros) RoutingOspfNbmaNeighbor(address string) (map[string]string, error) {
	return r.First(routingOspfNbmaNeighbor(address))
}

func setRoutingOspfNbmaNeighbor(address, key, value string) Command {
	return Command{
		Path:    "/routing ospf nbma-neighbor",
		Command: "set",
		Filter: map[string]string{
			"address": address,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingOspfNbmaNeighborComment(address, comment string) error {
	return r.Exec(setRoutingOspfNbmaNeighbor(address, "comment", comment))
}
func (r Ros) SetRoutingOspfNbmaNeighborPollInterval(address, interval string) error {
	return r.Exec(setRoutingOspfNbmaNeighbor(address, "poll-interval", interval))
}
func (r Ros) SetRoutingOspfNbmaNeighborPriority(address string, priority int) error {
	return r.Exec(setRoutingOspfNbmaNeighbor(address, "priority", strconv.Itoa(priority)))
}
