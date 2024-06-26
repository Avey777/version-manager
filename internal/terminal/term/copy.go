//go:build darwin || freebsd || dragonfly || linux

package term

import (
	"io"
	"os"
	"syscall"

	"github.com/gvcgo/version-manager/internal/terminal/term/fdset"
)

func Copy(dst io.Writer, src *os.File) func() {
	r, w, _ := os.Pipe()

	go func() {
		copy(dst, src, r)
	}()

	return func() {
		w.Write([]byte("x"))
		r.Close()
		w.Close()
	}
}

func copy(dst io.Writer, src *os.File, finish *os.File) (written int64, err error) {
	fd := int(src.Fd())
	ffd := int(finish.Fd())
	maxfd := ffd + 1
	if fd > ffd {
		maxfd = fd + 1
	}
	rfds := &syscall.FdSet{}
	buf := make([]byte, 32*1024)

	for {
		fdset.FD_SET(rfds, fd)
		fdset.FD_SET(rfds, ffd)

		es := fdset.Select(maxfd, rfds, nil, nil, nil)
		if es != nil {
			if es == syscall.EINTR {
				continue
			}

			err = es
			break
		}

		if fdset.FD_ISSET(rfds, fd) {
			nr, er := src.Read(buf)
			if nr > 0 {
				nw, ew := dst.Write(buf[0:nr])
				if nw > 0 {
					written += int64(nw)
				}
				if ew != nil {
					err = ew
					break
				}
				if nr != nw {
					err = io.ErrShortWrite
					break
				}
			}
			if er == io.EOF {
				break
			}
			if er != nil {
				err = er
				break
			}
		}

		if fdset.FD_ISSET(rfds, ffd) {
			break
		}
	}

	return written, err
}
