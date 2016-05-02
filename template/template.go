package template

import (
	"bytes"
	"fmt"
	"net"
	goexec "os/exec"
	"regexp"
	"strconv"
	"strings"
	gotemplate "text/template"

	"github.com/mattn/go-shellwords"
)

func Merge(template, left, right string, m map[string]interface{}) (string, error) {
	template = trim(template, left, right)
	t := gotemplate.New(template)
	t = t.Delims(left, right)
	funcMap := map[string]interface{}{
		"toNetwork":    toNetwork,
		"toPrefixLen":  toPrefixLen,
		"toSubnetMask": toSubnetMask,
		"exec":         exec,
	}
	t = t.Funcs(funcMap)
	t, err := t.Parse(template)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	t.Execute(&out, m)
	return fmt.Sprintf("%s", out.String()), nil
}

func trim(template, left, right string) string {
	match := []string{
		left + `/\*(\n|.)*?\*/` + right + `\n`,         // comment
		left + `\s*\$\w*\s*:\=\s*.*\s*` + right + `\n`, // variable
		left + `\s*range.*` + right + `\n`,             // range
		left + `\s*end\s*` + right + `\n`,              // end
	}
	for _, m := range match {
		r := regexp.MustCompile(m)
		submatches := r.FindAllStringSubmatch(template, -1)
		for _, s := range submatches {
			template = strings.Replace(template, s[0], strings.TrimRight(s[0], "\n"), 1)
		}
	}
	return template
}

func toNetwork(ip, mask string) string {
	var cidr string

	switch {
	case len(mask) > 2:
		if m := net.ParseIP(mask); m == nil {
			return "Invalid mask"
		}
		cidr = ip + "/" + toPrefixLen(mask)
	default:
		if m, err := strconv.Atoi(mask); err != nil || m < 0 || 32 < m {
			return "Invalid prefix length"
		}
		cidr = ip + "/" + mask
	}

	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "Invalid IP address"
	}

	return fmt.Sprintf("%s", ipnet.IP)
}

func toPrefixLen(mask string) string {
	m := net.ParseIP(mask).To4()
	if m == nil {
		return "Invalid mask"
	}
	prefixLen, _ := net.IPMask(m).Size()
	return fmt.Sprintf("%d", prefixLen)
}

func toSubnetMask(length string) string {
	l, err := strconv.Atoi(length)
	if err != nil || l < 0 || 32 < l {
		return "Invalid prefix length"
	}
	// :p
	maskList := []string{
		"0.0.0.0", "128.0.0.0", "192.0.0.0", "224.0.0.0", "240.0.0.0", "248.0.0.0", "252.0.0.0", "254.0.0.0", "255.0.0.0",
		"255.128.0.0", "255.192.0.0", "255.224.0.0", "255.240.0.0", "255.248.0.0", "255.252.0.0", "255.254.0.0", "255.255.0.0",
		"255.255.128.0", "255.255.192.0", "255.255.224.0", "255.255.240.0", "255.255.248.0", "255.255.252.0", "255.255.254.0", "255.255.255.0",
		"255.255.255.128", "255.255.255.192", "255.255.255.224", "255.255.255.240", "255.255.255.248", "255.255.255.252", "255.255.255.254", "255.255.255.255",
	}
	return maskList[l]
}

func exec(cmdstr string) string {
	c, err := shellwords.Parse(cmdstr)
	if err != nil {
		return "Invalid command"
	}
	switch len(c) {
	case 0:
		return "Invalid command"
	case 1:
		out, _ := goexec.Command(c[0]).Output()
		return fmt.Sprintf("%s", out)
	default:
		out, _ := goexec.Command(c[0], c[1:]...).Output()
		return fmt.Sprintf("%s", out)
	}
}
