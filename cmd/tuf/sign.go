package main

import (
	"github.com/flynn/go-tuf"
	"github.com/flynn/go-tuf/Godeps/_workspace/src/github.com/flynn/go-docopt"
)

func init() {
	register("sign", cmdSign, `
usage: tuf sign <manifest>

Sign a manifest.
`)
}

func cmdSign(args *docopt.Args, repo *tuf.Repo) error {
	return repo.Sign(args.String["<manifest>"])
}
