package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

var methods = map[string]string{
	"md5":    "",
	"sha1":   "",
	"sha224": "",
	"sha256": "",
	"sha384": "",
	"sha512": "",
}

func TestMain(m *testing.M) {
	if err := exec.Command("go", "build").Run(); err != nil {
		log.Fatal(err)
	}

	for m := range methods {
		out, err := exec.Command(fmt.Sprintf("%ssum", m), "hashcheck").Output()
		if err != nil {
			log.Fatal(err)
		}

		methods[m] = strings.Split(string(out), " ")[0]
	}

	ret := m.Run()

	if err := os.Remove("hashcheck"); err != nil {
		log.Fatal(err)
	}

	os.Exit(ret)
}

func TestHashcheck(t *testing.T) {
	for m := range methods {
		outbytes, err := exec.Command("./hashcheck", methods[m], "hashcheck").Output()
		if err != nil {
			t.Error(err)
		}

		out := string(outbytes)

		if !strings.Contains(out, strings.ToUpper(m)) {
			t.Errorf("hash %s was not recognized: %s", m, out)
		}

		if !strings.Contains(out, "OK") {
			t.Errorf("hash did not match: %s", out)
		}
	}
}
