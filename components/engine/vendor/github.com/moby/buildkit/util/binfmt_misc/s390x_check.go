// +build !s390x

package binfmt_misc

func s390xSupported() error {
	return check(Binarys390x)
}
