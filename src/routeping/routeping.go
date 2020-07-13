package routeping

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"strings"

	. "../structs"

	ping "github.com/sparrc/go-ping"
)

//getRoutePoints trought terminal gets and parses traceroute
func GetRoutePoints(ip string) []string {
	var points []string
	cmd := exec.Command("traceroute", ip)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	lines := string(cmdOutput.Bytes())
	spl := strings.Split(lines, "  ")
	for _, txt := range spl {
		if strings.Contains(txt, "(") {
			txtSpl := strings.Split(txt, "(")[1]
			ip := strings.Replace(txtSpl, ")", "", -1)
			points = append(points, ip)
		}
	}
	return points
}

//pingAddress coroutine ping the endpoint and send Pinged struct back
func PingAddress(c chan Pinged, adress string, count int) {
	pinger, err := ping.NewPinger(adress)
	if err != nil {
		panic(err)
	}
	pinger.Count = count
	pinger.Run()
	stats := pinger.Statistics()
	avrg := floatString(stats.AvgRtt.Seconds())
	max := floatString(stats.MaxRtt.Seconds())
	min := floatString(stats.MinRtt.Seconds())
	c <- Pinged{pinger.IPAddr().IP.String(), avrg, max, min}
}

//floatString  converts float64 to string -- just for shorter code
func floatString(f float64) string {
	return strconv.FormatFloat(f, 'g', 6, 64)
}
