**entrop** is a password generator (command line tool), that converts a sequence of words with some separator into
a password in some charset, using some good algorithms:
- https://en.wikipedia.org/wiki/PBKDF2 ("-a pbs5" or "-a pbs2")
- https://en.wikipedia.org/wiki/Argon2 ("-a ar")
- and some others (see alg.go)

_Usage of entrop_:
```
  -V uint
        alg settings ('defaults') version, default = 0
  -a string
        algorithm with optional params: e.g. ar:3:32768 or rh:125000 (default "pbs5")
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

In inline mode (-i), separator can be specified with -s option, and words must be specified after flags:
```
entrop -i -l 15 -s "+++" hello world 12345
# separator=+++; words=hello,world,12345
```
In a default non-inline mode, separator and words will be asked in hidden inputs (as in `read -s`).
Inline mode is not recommended, it is insecure. Use inline mode only if you can disable commands history in your shell.


## GUI

*entrop* also is available for usage online (for that puprose it was compiled to wasm):

https://dlepex.github.io/7w/index.html
