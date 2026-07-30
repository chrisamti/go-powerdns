package powerdns_test

import (
	"context"
	"fmt"
	"log"

	"github.com/joeig/go-powerdns/v3"
)

func ExampleCryptokeysService_Create() {
	pdns := powerdns.New("http://localhost:8080", "localhost", powerdns.WithAPIKey("apipw"))
	ctx := context.Background()

	domain := "cryptokey-example.com."

	_ = pdns.Zones.Delete(ctx, domain)

	if _, err := pdns.Zones.AddNative(ctx, domain, false, "", false, "", "", true, []string{"localhost."}); err != nil {
		log.Fatalf("%v", err)
	}

	cryptokey, err := pdns.Cryptokeys.Create(ctx, domain, powerdns.Cryptokey{
		KeyType:   powerdns.String("ksk"),
		Active:    powerdns.Bool(true),
		Published: powerdns.Bool(true),
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("active=%t published=%t\n",
		powerdns.BoolValue(cryptokey.Active),
		powerdns.BoolValue(cryptokey.Published))
	// Output: active=true published=true
}

func ExampleCryptokeysService_Change() {
	pdns := powerdns.New("http://localhost:8080", "localhost", powerdns.WithAPIKey("apipw"))
	ctx := context.Background()

	domain := "cryptokey-example.com."

	cryptokeys, err := pdns.Cryptokeys.List(ctx, domain)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if err := pdns.Cryptokeys.Change(ctx, domain, powerdns.Uint64Value(cryptokeys[0].ID), powerdns.Cryptokey{
		Active: powerdns.Bool(false),
	}); err != nil {
		log.Fatalf("%v", err)
	}
}

func ExampleCryptokeysService_List() {
	pdns := powerdns.New("http://localhost:8080", "localhost", powerdns.WithAPIKey("apipw"))
	ctx := context.Background()

	if _, err := pdns.Cryptokeys.List(ctx, "cryptokey-example.com."); err != nil {
		log.Fatalf("%v", err)
	}
}

func ExampleCryptokeysService_Get() {
	pdns := powerdns.New("http://localhost:8080", "localhost", powerdns.WithAPIKey("apipw"))
	ctx := context.Background()

	if _, err := pdns.Cryptokeys.Get(ctx, "cryptokey-example.com.", 1); err != nil {
		log.Fatalf("%v", err)
	}
}

func ExampleCryptokeysService_Delete() {
	pdns := powerdns.New("http://localhost:8080", "localhost", powerdns.WithAPIKey("apipw"))
	ctx := context.Background()

	if err := pdns.Cryptokeys.Delete(ctx, "cryptokey-example.com.", 1); err != nil {
		log.Fatalf("%v", err)
	}
}
