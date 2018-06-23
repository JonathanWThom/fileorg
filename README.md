# Fileorg

The goal of this project is to be able to run a command, `fileorg`, in a directory, and
sort all of the files in it into sub directories based on their filetype/extension. Optionally,
if there are a lot of files (like a bunch of photos), it could organize them into further
subdirectories based on when they were created. It's a WIP!

### Usage

**This tool moves files. It could screw up your computer if executed on the
wrong directory. Make sure you know what you're doing.**

1. [Install Go](https://golang.org/doc/install)
2. `go get github.com/jonathanwthom/fileorg`
3. `go install github.com/jonathanwthom/fileorg`
4. Navigate to the directory you want to organize.
5. `fileorg`

### TODO

- Build out test coverage.
- Add confirmation message (e.g. "You are about to move these files, are you sure
you want to do that?").
- Create further subdirectories, sorted by date.
- Organize files by _type_ (like all images go together).
