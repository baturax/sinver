package main

import (
	"os"
	"os/user"
	"strings"

	"github.com/shirou/gopsutil/v4/host"
)

func getUser() string {
	userr, err := user.Current()
	if err != nil {
		return "Cant get user"
	}
	return userr.Username
}

func getDistro() string {
	a := getHost().Platform
	lets := strings.Split(a, "")
	return strings.ToUpper(lets[0])+ strings.Join(lets[1:], "")
}

func getKernel() string {
	return getHost().KernelVersion
}

func getHost() *host.InfoStat {
	h, err := host.Info()
	if err != nil {
		os.Exit(1)
	}

	return h
}
