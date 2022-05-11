package go_math

import "golang.org/x/xerrors"

// 卷积核滤波full模式
func filter2full[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](hss, xss [][]T) ([][]T, error) {
	return nil, nil
}

// 卷积核滤波same模式
func filter2same[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](hss, xss [][]T) ([][]T, error) {
	if len(hss)%2 == 0 {
		return nil, xerrors.New("卷积核行数需要是奇数")
	} else if len(hss[0])%2 == 0 {
		return nil, xerrors.New("卷积核列数需要是奇数")
	}
	yss := make([][]T, len(xss))
	return yss, nil
}

// 卷积核滤波valid模式
func filter2valid[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](hss, xss [][]T) ([][]T, error) {
	return nil, nil
}
