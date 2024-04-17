
Overview
--------

`domainify` is a command line utility for converting hostnames to "entity
domains", the organizational domain name associated with a hostname.

Entity domains are defined using the Profound Public Suffix List (PPSL),
a fork of the well-known [Mozilla Public Suffix List](https://publicsuffix.org)
used by browsers for restricting the scope of cookies across hostnames.
(The PPSL is bundled with the binary, so no external datasets are required.)


Usage
-----

domainify takes hostnames as arguments, and outputs entity domains to
standard output, one entity domain per line. A bad hostname or an error
mapping to an entity domain will result in a blank line being output for
that input.

Arguments can be givin either as command line arguments, or via standard
input, one hostname per line, with the `--stdin`  option.

For example:

```bash
$ domainify www.example.com
example.com

$ domainify www.profound.net bogus en.wikipedia.org
profound.net

wikipedia.org

$ echo -e "hosted.l.google.com\ncom\n31.pool85-61-184.dynamic.orange.es" |
  domainify --stdin
google.com

orange.es
```

Licence and Copyright
---------------------

Copyright (c) 2024 by [Profound Networks](https://profound.net).

`domainify` is licensed under the MIT License. See the file `LICENCE.md` for
details.

