# proquint

An implementation of [A Proposal for Proquints: Identifiers that are Readable, Spellable, and Pronounceable](http://arxiv.org/html/0901.4016) in Go. This package converts integers to/from sets of pronounceable five-letter words, e.g. 791,594,501 becomes `fusov-simaj`. See original article for more details.

This implementation is based on [upspin.io/key/proquint](https://github.com/upspin/upspin/tree/master/key/proquint), with the antiquint subpackage providing output compatible with the output of [proquint-php](https://github.com/Fil/proquint-php). On top of that, this fork adds functions for handling 32-/64-bit integers and byte arrays, including the use of custom separators, and encodes proquints as strings instead of byte arrays. [Documentation can be found on pkg.go.dev.](https://pkg.go.dev/github.com/syrupyy/proquint)
