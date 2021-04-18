**entrop** is a password generator, that converts a sequence of words with some separator into
a password in some charset, using some good algorithms:
- https://en.wikipedia.org/wiki/PBKDF2 ("-a pbs5" or "-a pbs2")
- https://en.wikipedia.org/wiki/Argon2 ("-a ar")
- and some others (see alg.go)

_Usage of entrop_:
```
  -V uint
        alg settings ('defaults') version, default = 0
  -a string
        algorithm with optional params: e.g. ar:3:32768 or rsha:11111111 (default "pbs5")
  -c string
        charset, see charsets.go (default "alnum")
  -i    inline mode, i.e. no hidden inputs for words & separator
  -l uint
        password length (default 25)
  -ncw
        no word numbering/counting
  -nq
        no quality check/retry
  -s string
        separator (default " ")
  -v    verbose mode
```
