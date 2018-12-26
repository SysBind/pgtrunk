package pgpool

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Node struct {
	Host string
	Port int
	Status int
	Weight float32
	Role string
	ReplicationDelay int
}

func pcpNodeCount() int {	
	out, err := exec.Command("pcp_node_count", "-h", "127.0.0.1","--username", "root", "-w").Output()
	
	if err != nil {
		log.Fatal(err)
	}
	count, _ := strconv.Atoi(strings.Trim(string(out), " \n"))
	return count
}

func pcpNodeInfo(idx int) Node {
	out, err := exec.Command("pcp_node_info",
		strconv.Itoa(idx),
		"-h",
		"127.0.0.1",
		"--username",
		"root", "-w").Output()

	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(out), " ")

	host := parts[0]
	port, _ := strconv.Atoi(parts[1])
	status, _ := strconv.Atoi(parts[2])
	weight64, _ := strconv.ParseFloat(parts[3], 32)
	weight := float32(weight64)
	role := strings.Trim(parts[5], " ");
	replicationDelay, _ := strconv.Atoi(parts[6])
	return Node{host, port,	status, weight,	role, replicationDelay}
}
