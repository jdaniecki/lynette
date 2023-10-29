package network

import (
	"log/slog"

	"github.com/vishvananda/netlink"
)

func EnsureBridge(name string) error {
	slog.Debug("Checking bridge...", "name", name)

	//  if bridge already exists exit early
	if _, err := netlink.LinkByName(name); err == nil {
		slog.Debug("Bridge already exists!")
		return nil
	}

	// create bridge (requires CAP_NET_ADMIN)
	slog.Debug("Creating bridge...", "name", name)
	la := netlink.NewLinkAttrs()
	la.Name = name

	bridge := &netlink.Bridge{LinkAttrs: la}
	err := netlink.LinkAdd(bridge)
	if err != nil {
		return err
	}
	return nil
}
