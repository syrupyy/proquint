# proquint
An implementation of [A Proposal for Proquints: Identifiers that are Readable, Spellable, and Pronounceable](http://arxiv.org/html/0901.4016) in Go. This package converts integers to/from sets of pronounceable five-letter words, e.g. 96,874,031 becomes `fusov-simaj`. See original article for more details.

This implementation is based on [upspin.io/key/proquint](https://github.com/upspin/upspin/tree/master/key/proquint) and is subject to the Apache License 2.0. This fork adds functions for handling 32-bit and 64-bit integers, including the use of custom separators, and uses strings instead of byte arrays. [Documentation can be found on pkg.go.dev.](https://pkg.go.dev/github.com/syrupyy/proquint)
