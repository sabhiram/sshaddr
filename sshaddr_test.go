package sshaddr

import (
	"testing"
)

func fatalIfError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Fatal error: %s\n", err.Error())
	}
}

func TestParse(t *testing.T) {
	for _, tc := range []struct {
		addr   string
		exUser string
		exPass string
		exHost string
		exPort int
		exDest string
	}{
		{
			addr:   "user:1@foobar.com:1234:/tmp/foobar",
			exUser: "user",
			exPass: "1",
			exHost: "foobar.com",
			exPort: 1234,
			exDest: "/tmp/foobar",
		},
		{
			addr:   "user@foobar.com:1234:/tmp/foobar",
			exUser: "user",
			exPass: "",
			exHost: "foobar.com",
			exPort: 1234,
			exDest: "/tmp/foobar",
		},
		{
			addr:   "user@foobar.com:/tmp/foobar",
			exUser: "user",
			exPass: "",
			exHost: "foobar.com",
			exPort: 22,
			exDest: "/tmp/foobar",
		},
	} {
		ssha, err := Parse(tc.addr)
		fatalIfError(t, err)

		if ssha.User() != tc.exUser {
			t.Errorf("Expected user=%s got=%s\n", tc.exUser, ssha.User())
		}
		if ssha.Pass() != tc.exPass {
			t.Errorf("Expected Pass=%s got=%s\n", tc.exPass, ssha.Pass())
		}
		if ssha.Host() != tc.exHost {
			t.Errorf("Expected Host=%s got=%s\n", tc.exHost, ssha.Host())
		}
		if ssha.Destination() != tc.exDest {
			t.Errorf("Expected Dest=%s got=%s\n", tc.exDest, ssha.Destination())
		}
		if ssha.Port() != tc.exPort {
			t.Errorf("Expected Port=%d got=%d\n", tc.exPort, ssha.Port())
		}
	}
}
