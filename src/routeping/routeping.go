package routeping

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "../structs"
)

//GetRoutePoints trought terminal gets and parses traceroute
func GetRoutePoints(ip string) []string {
	var points []string
	fmt.Println("- Fetching ips on route")
	cmd := exec.Command("traceroute", ip)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	lines := string(cmdOutput.Bytes())
	spl := strings.Split(lines, "\n")
	for _, txt := range spl {
		if strings.Contains(txt, "(") &&
			!strings.Contains(txt, "traceroute") {
			txtSpl := strings.Split(txt, "(")[1]
			ip := strings.Replace(txtSpl, ")", "", -1)
			points = append(points, ip)
		}
	}
	fmt.Println(fmt.Sprintf("- Found %d ips on route", len(points)))
	return points
}

//PingAddress coroutine ping the endpoint and send Pinged struct back
func PingAddress(c chan Pinged, adress string, count int) {
	cmd := exec.Command("ping", adress, "-c 1")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	lines := string(cmdOutput.Bytes())
	values := parserPingResponse(lines)
	if err != nil {
		c <- Pinged{
			Address: adress,
			Average: values[1],
			Max:     values[2],
			Min:     values[0],
			Message: parsePingError(err, lines),
		}
	}

	c <- Pinged{
		Address: adress,
		Average: values[1],
		Max:     values[2],
		Min:     values[0],
		Message: "Ok",
	}
}

func parsePingError(err error, response string) string {
	if err.Error() == "exit status 2" {
		spl := strings.Split(response, "\n")
		for _, txt := range spl {
			if strings.Contains(txt, "100.0%") {
				return "100% packet loss"
			}
		}
	}
	return response
}

func parserPingResponse(response string) []string {
	spl := strings.Split(response, "\n")
	values := []string{"0", "0", "0", "0"}
	for _, txt := range spl {
		if strings.Contains(txt, "rtt") ||
			strings.Contains(txt, "round-trip") {
			txtSpl := strings.Split(txt, "=")[1]
			values = strings.Split(strings.Trim(txtSpl, " "), "/")
		}
	}
	return values
}
