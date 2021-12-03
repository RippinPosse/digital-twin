package kerberos

import (
	"encoding/hex"
	"fmt"
	"github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/spnego"
	"net/http"
)

type Krb struct {
	address string

	client *client.Client
}

func New(address, username, keytabData, keytabConfig string) (*Krb, error) {
	b, err := hex.DecodeString(keytabData)
	if err != nil {
		return nil, fmt.Errorf("decode keytab string: %w", err)
	}

	kt := keytab.New()
	if err := kt.Unmarshal(b); err != nil {
		return nil, fmt.Errorf("unmarshal keytab data: %w", err)
	}

	c, err := config.NewFromString(keytabConfig)
	if err != nil {
		return nil, fmt.Errorf("create config: %w", err)
	}

	c.Realms[0].KDC = []string{address}

	cl := client.NewWithKeytab(username, "ARENA.RU", kt, c)

	krb := &Krb{
		address: address,
		client: cl,
	}

	return krb, nil
}

func (krb *Krb) Auth() error {
	if err := krb.client.Login(); err != nil {
		return fmt.Errorf("login: %w", err)
	}

	r, err := http.NewRequest("GET", krb.address, nil)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	err = spnego.SetSPNEGOHeader(krb.client, r, "HTTP/host.test.gokrb5")
	if err != nil {
		return fmt.Errorf("set SPNEGO header: %w", err)
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	fmt.Println(resp)

	return nil
}
