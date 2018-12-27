package pgpool

import (
	"log"
	"os/exec"
	"strconv"
	"fmt"
)

func GetPrimaryNode(conn *PCPConn) Node {
	var node Node
	nodeCount := conn.pcpNodeCount()
	for i := 0; i < nodeCount; i++ {
		node = conn.pcpNodeInfo(i)
		if node.Role == "primary" { return node	}
	}
	return Node{}
}

func Sync(src Node, datadir string) {
	out, err := exec.Command("pg_basebackup",
		"--write-recovery-conf",
		"--wal-method=stream",
		"-D", datadir,
		"-h", src.Host,
		"-p", strconv.Itoa(src.Port)).CombinedOutput()

	if err != nil {
		fmt.Println(string(out))
		log.Fatal(err)
	}

	fmt.Println(string(out))
}


func InitDB(datadir string) {
	out, err := exec.Command("initdb",
		datadir).CombinedOutput()

	if err != nil {
		fmt.Println(string(out))
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
