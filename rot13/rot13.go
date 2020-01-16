
package rot13

import "io"

type reader struct {
    r io.Reader
}

// NewReader returns a new io.Reader whose data will be encrypted via rot13
func NewReader(r io.Reader) io.Reader {
    return reader{r}
}


// c'est une methode
func (rr reader) Read(p []byte) (int, error) {
	 
	n, err := rr.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	
	
	return n, err 
}

func rot13(b byte) byte {
    var a, z byte
    switch {
    case 'a' <= b && b <= 'z':
        a, z = 'a', 'z'
    case 'A' <= b && b <= 'Z':
        a, z = 'A', 'Z'
    default:
        return b
    }
    return (b-a+13)%(z-a+1) + a
}



