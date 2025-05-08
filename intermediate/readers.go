/*
reader - the interface for reading streams of data and populating the given byte slice.
*/
package intermediate

import "io"

// Exercise: Readers
// Goal: Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

// Exercise: rot13Reader
// Goal: A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
//
// For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).
//
// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.
//
// The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(p []byte) (int, error) {
	n, err := rr.r.Read(p)

	for i := 0; i < n; i++ {
		b := p[i]
		if 'A' <= b && b <= 'Z' {
			p[i] = 'A' + (b-'A'+13)%26
		} else if 'a' <= b && b <= 'z' {
			p[i] = 'a' + (b-'a'+13)%26
		}
	}

	return n, err
}

// Exercise: CountingReader
// Goal: Wrap any io.Reader so that you can measure exactly how many bytes have passed through it.

type CountingReader struct {
	R     io.Reader
	Count int64
}

func (cr *CountingReader) Read(p []byte) (int, error) {
	n, err := cr.R.Read(p)

	cr.Count += int64(n)

	return n, err
}

// Exercise: LimitReader
// Goal: Write a LimitReader that wraps an io.Reader but stops after N bytes, returning io.EOF once the limit is reached.

type LimitReader struct {
	R io.Reader
	N int64
}

func (lr *LimitReader) Read(p []byte) (int, error) {
	if lr.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > lr.N {
		p = p[:lr.N]
	}

	n, err := lr.R.Read(p)
	lr.N -= int64(n)

	if lr.N <= 0 {
		err = io.EOF
	}

	return n, err
}
