package pgpool

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type pcpServer struct {
	hostname string
	port int	
}

type PCPConn struct {
	servers []pcpServer
	username string
	password string
}

type Node struct {
	Host string
	Port int
	Status int
	Weight float32
	Role string
	ReplicationDelay int
}

// PCPConnection generate PCPConn object from first server details
// Should then fetch other servers automatically
func PCPConnection(hostname string, port int, username string, password string) *PCPConn {
	return &PCPConn{[]pcpServer{{hostname, port}}, username, password}
}

// pcpCommand 
func (c *PCPConn) pcpCommand(command string) string {	
	out, err := exec.Command(command, 
		"-h", c.servers[0].hostname,
		"-p", strconv.Itoa(c.servers[0].port),
		"--username", c.username, "-w").CombinedOutput()
	
	if err != nil {
		log.Println(string(out))
		log.Fatal(err)
	}

	return strings.Trim(string(out), " \n")
}

// pcpNodeCommand
func (c *PCPConn) pcpNodeCommand(command string, node int) string {	
	out, err := exec.Command(command,  strconv.Itoa(node),
		"-h", c.servers[0].hostname,
		"-p", strconv.Itoa(c.servers[0].port),
		"--username", c.username, "-w").CombinedOutput()
	
	if err != nil {
		log.Println(string(out))
		log.Fatal(err)
	}

	return strings.Trim(string(out), " \n")
}

// pcpNodeCount
func (c *PCPConn) pcpNodeCount() int {	
	count, err := strconv.Atoi(c.pcpCommand("pcp_node_count"))

	if err != nil {
		log.Fatal(err)
	}

	return count
}

func (c *PCPConn) pcpNodeInfo(idx int) Node {
	out := c.pcpNodeCommand("pcp_node_info", idx)

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
