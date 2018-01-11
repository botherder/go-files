package files

import (
    "os"
    "io"
    "errors"
    "hash"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "encoding/hex"
)

func hashTarget(target string, algorithm string, format string) (string, error) {
    var h hash.Hash

    switch algorithm {
    case "md5":
        h = md5.New()
    case "sha1":
        h = sha1.New()
    case "sha256":
        h = sha256.New()
    case "sha512":
        h = sha512.New()
    default:
        return "", errors.New("No valid algorithm specified")
    }

    switch format {
    case "file":
        file, err := os.Open(target)
    	if err != nil {
    		return "", errors.New("Unable to open specified file")
    	}
    	defer file.Close()

        _, err = io.Copy(h, file)
        if err != nil {
            return "", errors.New("Unable to hash file")
        }
    case "string":
        h.Write([]byte(target))
    }

    return hex.EncodeToString(h.Sum(nil)), nil
}

func HashFile(target string, algorithm string) (string, error) {
    return hashTarget(target, algorithm, "file")
}

func HashString(target string, algorithm string) (string, error) {
    return hashTarget(target, algorithm, "string")
}
