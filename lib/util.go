package lib

import "github.com/rs/xid"

func IdGet() string {
	guid := xid.New()
	return guid.String()
}
