# go-files

Collection of functions to help with operation with and on files.

## Usage

Copy a file:

```go
err := files.Copy("/path/to/src", "/path/to/dir")
```

Hash a file:

```go
hash, err := files.HashFile("/path/to/file", "sha512")
```

Hash a string:

```go
hash, err := files.HashString("test123", "md5")
```

On Windows, expand paths:

```go
path := files.ExpandWindows("%ProgramFiles%\\Application\\file.exe")
```

## License

Released under MIT license.
