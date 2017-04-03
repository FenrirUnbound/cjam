package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/golang/main.go": {
		local:   "golang/main.go",
		size:    810,
		modtime: 1491176868,
		compressed: `
H4sIAAAAAAAA/5SSwY4TMQyGz8lTmKCVkmUY1GulniqtxIkVPXDY3UPUdYpFmowSz4CE+u7ImZnCCjjs
panH9m9/tgd//OZPCGdPSWs6D7kwWK1MiP5k5D2zPJQ/UB6Zohi5ym/lQulUjXZahzEd4TP6549pGNkG
ipj8GWGOcfDwNP+Dn1odc2JMXDvAUmC7g1m4l/Q7injNdlpRaEFvdpAoSrIafKKjxVKcVhet1dFXrKKy
tNMfhkhsZ8uutVwH5jEZp1VBHkuClvawgS1ETLZZ7v3mSV8Wli+FGD+N/A+aDnyq37HUK5WTzsKVJ9d+
X9Dzq0jUMwYsEPp9zBWt01qFXIA6mESz+HTCa2FJD2fu7wsljslOTj70renDTC7uwyD+YM3eV4S3N9MW
bqbHZDqgd5sOJtcqr8hyArahkCzxbqXe7kCOoV+EDZkOTAvp+QeL0ZYOAgsycE9JVu1jBP6K86hl9LmN
87+6WaTmmFV4XsBfyqJasI6RRVerJnPvyzK3xSUFDjlOaH8f5u0LNCfRf2769mWL3VrF6Yv+FQAA//8C
Qm5RKgMAAA==
`,
	},

	"/golang/solver.go": {
		local:   "golang/solver.go",
		size:    119,
		modtime: 1491176868,
		compressed: `
H4sIAAAAAAAA/ypITM5OTE9VyE3MzOPiyswtyC8qUdDQ5OJKK81LVgjOzylL1UhOLE4tVoiOLS4pysxL
14SzFKq5OItSi0tzShSsbOGi1UoZqTk5+Uo6Ckrl+UU5KUq1XCBlJaVFeQoQ1Vy1XIAAAAD//4sCsjB3
AAAA
`,
	},

	"/node6/main.js": {
		local:   "node6/main.js",
		size:    3820,
		modtime: 1491188266,
		compressed: `
H4sIAAAAAAAA/6xXX2/bNhB/16e4tQUkt7Kc7GmI4W5B0WEFtrVYB+whC1CaOtlsaVIjqThG4u8+HEn9
nbMERfVgi+Td8e74+x1Pz79bNNYs1kItUN2A0iUmSdpYBOuM4C5dJgnXyjqoLKzA4D+NMJillU1ny7hi
tbxBM1wtFmGOZJLFy5cJvIQPzFi0wMzmBiptwG0RhKobB0yVoBtHr5WQCIrt0Bak81PNDNsB3H10RqjN
1fUx6ANAbTRHawsae1GDrjEK7t6vPyN3R+ifMANcK8eEEmoDX/BgL0jrkeeZd/AZwDy6Sv6Re0/RDSE9
g/kwuKi8SKpGcSe0gprycmk2NqNQZnCXAMSMS7ahpNMMQDoX6QWk3o80j1PzMPzPvKapsG0vG8f9SgJw
XHbbGbSNdN1+3uAFpFE/qNC41ZLowLIb/B1vSSslqIA/n8JKwTH7flZU2rxlfJtlwuFuBqvXrfEKslZ3
FucgenDVLlzDCkhvGZcnm9HUEVBa9OZ8tq5I/rq3OFAZrEddCmTmnV4soMSKUfg3TDZok9aZIpz7ajy8
v48JL9yt867E5XjSq8mY5MNrVPAaHq9BcJkcO6L8gawcsKM2ei1xB5wRfSqjd8BAas6kB5TnyaUFBlaU
CFhVyF0OwsFeSAm1aRR6a5JZB1IosgvM6+aUOE0zlqzUBi0qd4J4xw67MNjaj3XlrY+8HPHxg9E7YXFA
yDgDbss86qhUWBDKaYrCb/gEfjFj2OGh7Qf8MsjKd5TJrPU5wCP6p3DfOpRl0ZkcDFLJGAC2sgUZ+llI
7OzkkDau+iHNIUNjciiZYwOVgHI0pofj4NTJvl9s0X1M4suEjWS1sLUULkv/VmknH8Qkqo3bwiqUkXfK
ZSR/dXadw/nZRJZvkX/5lWDQ4vMqwjRamcP59TJJJp76lGRZpz2DH1vvLlqcB8Kf5dGfV+ezuPfR/9Nv
D/C/jHABk0zZPRoLTvthiVZsFHM4ug9yYJxrU1LhDoJkhGtjqKoHE3Sj7Nhp6MIAvIN3p8E6baaOPHTt
RAF6LnuvkfFtBz4aPAL9h+G/36LyZx3J+IRnb4RzqE5kj4Ics4Bk8b1P6gC/MaqvIcQQpXRJRVPFjtVZ
FgY5CFXi7YgScZNPb5hFeP7izku8Oj9ewIu7oHX81COnpZ73fsK9uHXxWQsVmDHm4zdg4oQC/wPpN0xK
608hNkNd4p1uWxW/vBE3qDrEWHS+fv/CLLCyFKTBJHC9qyXeCnfwnZIVNGwPZb5mFksvw5xYCyncoYA/
IyuIpNRkYdAS1glOBV8x1xh/4qRKHmolDx1uTEc5D75YVT/6WB7qxGIQb/y9BEAuTC4rqpi0ub+0yGDE
5WMMeZglEeoRbE/kCUBMe+dHOKQxQziTMgScDSP7em7s2GGNbRiruOfI9je+OCZw9VYnXFosInBYAA7s
hdv60NeMf5lbd5AIjfW5jX4Mw5iNPLykG7gQ1v9P5O7vwR1q1NUoDYWjKrdarSBt856eDDJEMDLZhwuD
5q2n4G+MkBatFvAu8G10vFiGMushwI2oHQgLhqnTPZRvoIiGMUu6RDiDptYKbMOJ0jloA+d+ikxUTMjG
hJbsJMRPg9pbDJz3vo9xuWNCZaPPglgCbXvr+4+H4ffQbNxgTvsfG7rYWQLgDyTrod/PtRfe6vXJu8PG
1ra/Q3rVzGO5dYgSmJ2FZc4cfQyMqzNFpSUWaIw2WfqW/gAV141yaLBsW56xXM+A0Ubnw+IcMxAyuEz+
DQAA//+E1aJp7A4AAA==
`,
	},

	"/node6/solver.js": {
		local:   "node6/solver.js",
		size:    406,
		modtime: 1491188272,
		compressed: `
H4sIAAAAAAAA/1SQwW7yMBCE736KuRkiBA+A+PWjSn0BjlEOxtkkbh3b2t3QIpR3r0iLWuY6s9/urJ2E
IMrBq90bs6sqgwqnHC8EHQh9uFCCf3MjCudzpBFCes/8L47dCNxOyiH1dTM/Ei9OSADgyOyuyN2yIPUC
psIklDSkfsE/mP4+8kR9nZLXkNOM2rsYz86/N1h0TMjlbrmIh4XQQQenCEmJO+cJQVCYOmKmdiEz6cQJ
t1qvhZoZT6pbEs9h4TYG1c6MuZ0ibemzZFbBAau/9dY4/MPNAD4nUWhWF797H57esJUhdLpa740Bfi6o
f9Mb2IFizHYD+5E5trbZm9l8BQAA//+MSHaTlgEAAA==
`,
	},

	"/py27/main.py": {
		local:   "py27/main.py",
		size:    838,
		modtime: 1491188241,
		compressed: `
H4sIAAAAAAAA/3yST2/bMAzF7/oUnHqojTkKdi3gwzBg5923QVBrKhFgUwZFNx6KfvdBkt0/2dCTE+rx
xyc93nw6LomP94GOSI8w/5FzJKXCNEcWcHyaHSfc/6c4PiIrNaAHRjfYQPMijQ8jkpuwvVMAAA+RBEmg
h5+/S+ES5AxxRnpVgkvgq/x9izcZPAbC1LTXuNUk4TA3LfjIsEKg/bQOYpSFX2rV54WDoI2LZKeO0gU5
dXBl+V+HHdxebq9t5qmBBlw7yA7zfKRlQnaCO7t9lZcWUww0+ptLCDdPz3fw9PyLtPGRJydNwX3+UoFt
Wz1PLlCzgUoADP1LGOYrn5YJSX6Uk+2Vqsy4YbBuO2/0IegO9OFQYtIdDOjdMkqvS8HImotnHOdef9/u
DdGDnPPVcstH7FjZ9WXfwmvlQ/rW1KrCd3xK0O9jyicPygtQ53O8H3GyDy5hFr5ZvSwz5efOqilAvy2r
KZ/mHWKT/n8zCrEWW6WCB2uzc2uh70Fbm7OxVtdwalDqbwAAAP//9z/u4UYDAAA=
`,
	},

	"/py27/solver.py": {
		local:   "py27/solver.py",
		size:    107,
		modtime: 1491188241,
		compressed: `
H4sIAAAAAAAA/0pJTVMozs8pS9UoKMpPyknNjU9OLE4t1rTiUlBQUCjJL0nMUbBVQJHTK8gv0DDQBCso
Si0uzSkpRlfCBZUsKS3Kg6nhAgQAAP//sl947WsAAAA=
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/golang": {
		isDir: true,
		local: "golang",
	},

	"/node6": {
		isDir: true,
		local: "node6",
	},

	"/py27": {
		isDir: true,
		local: "py27",
	},
}
