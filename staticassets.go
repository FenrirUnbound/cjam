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

	"/": {
		isDir: true,
		local: "",
	},

	"/golang": {
		isDir: true,
		local: "golang",
	},
}
