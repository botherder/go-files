# go-files

Collection of functions to help with operation with and on files.

## Usage

Copy a file:

```go
err := files.Copy("/path/to/src", "/path/to/dir")
```

Hash a string:

```go
hash, err := files.HashString("test123", "md5")
```

Hash a file:

```go
hash, err := files.File("/path/to/file", "sha512")
fmt.Println(h)
```

## License

Released under MIT license.
