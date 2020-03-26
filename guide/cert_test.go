package guide

import (
	"testing"

	"github.com/p4gefau1t/trojan-go/common"
)

func TestCreate(t *testing.T) {
	caDir = "https://127.0.0.1:14000/dir"
	common.Must(CreateCert("localhost", "test@email.com"))
}

func TestRenew(t *testing.T) {
	caDir = "https://127.0.0.1:14000/dir"
	common.Must(RenewCert("localhost", "test@email.com"))
}
