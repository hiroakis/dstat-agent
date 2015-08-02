package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type dstatLabel struct {
	section string
	keys    []string
}

var dstatLabels = []dstatLabel{
	dstatLabel{
		section: "cpu",
		keys:    []string{"usr", "sys", "idl", "wai", "hiq", "siq"},
	},
	dstatLabel{
		section: "disk",
		keys:    []string{"read", "writ"},
	},
	dstatLabel{
		section: "load",
		keys:    []string{"1m", "5m", "15m"},
	},
	dstatLabel{
		section: "memory",
		keys:    []string{"used", "buff", "cach", "free"},
	},
	dstatLabel{
		section: "net",
		keys:    []string{"recv", "send"},
	},
	dstatLabel{
		section: "procs",
		keys:    []string{"run", "blk", "new"},
	},
	dstatLabel{
		section: "swap",
		keys:    []string{"used", "free"},
	},
	dstatLabel{
		section: "system",
		keys:    []string{"init", "csw"},
	},
}

var dstatResult string

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, dstatResult)
}

func main() {
	host := flag.String("host", "0.0.0.0", "Listen host")
	port := flag.String("port", "8888", "Listen port")
	flag.Parse()

	listenAddr := fmt.Sprintf("%s:%s", *host, *port)

	cmd := exec.Command("dstat", "-cdlmnpsy", "--noheaders", "--nocolor", "--noupdate")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		cmd.Start()
		scanner := bufio.NewScanner(stdout)
		results := make(map[string]map[string]string)
		for scanner.Scan() {
			text := scanner.Text()
			if strings.Index(text, "---") != -1 || strings.Index(text, "usr sys") != -1 {
				// skip headers
				continue
			}

			// divide into each section "cpu(usr sys..), disk(read writ), ..."
			sections := strings.Split(text, "|")

			for i, j := range sections {

				// cpu(usr sys..) => []string{usr, sys...}
				// disk(read, writ) => []string{read, writ}
				sectionValues := strings.Fields(j)

				sectionResult := make(map[string]string)
				for p, q := range sectionValues {
					sectionResult[dstatLabels[i].keys[p]] = q
				}
				results[dstatLabels[i].section] = sectionResult
			}

			jsonStr, err := json.MarshalIndent(results, "", "\t")
			if err != nil {
				fmt.Println(err)
			}
			dstatResult = fmt.Sprintf("%s", jsonStr)
		}
	}()

	http.HandleFunc("/", handler)
	http.ListenAndServe(listenAddr, nil)
}
