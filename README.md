# Fileorg

A command line utility to organize files. Run `fileorg` on a messy directory,
and all files will be moved into subdirectories based on their file extension.

### Usage

**This tool moves files. It could screw up your computer if executed on the
wrong directory. Make sure you know what you're doing.**

_With Go_
1. [Install Go](https://golang.org/doc/install)
2. `go get github.com/jonathanwthom/fileorg`
3. `go install github.com/jonathanwthom/fileorg`
4. Navigate to the directory you want to organize.
5. `fileorg`

_Without Go_
1. `git clone https://github.com/JonathanWThom/fileorg.git`
2. `cd fileorg`
3. `sudo cp fileorg /usr/local/bin`
4. `cd ~/path/to/your/directory`
5. `fileorg`

If you want to organize files even futher, passing the `-date` option will
will create subdirectories based on when each file was last updated. For example:

```
$ ls
$ old.pdf also_old.pdf new.pdf
$ fileorg -date
$ ls
$ pdf
$ ls pdf
$ July_2014 March_2018
$ ls pdf/July_2014
$ old.pdf also_old.pdf
```

### TODO

- Build out test coverage.
- Add confirmation message (e.g. "You are about to move these files, are you sure
you want to do that?").
- Organize files by _type_ (like all images go together).
