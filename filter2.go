package go_math

import (
	"context"
	"runtime"
	"sync"

	"golang.org/x/sync/semaphore"
	"golang.org/x/xerrors"
)

// 卷积核滤波full模式
// xss为原矩阵
// tss为原矩阵补0
// yss为m*n矩阵 m行n列
// hss为i*j矩阵 i行j列
func filter2full[T ~int | ~float32 | ~float64](hss, xss [][]T) ([][]T, error) {
	if len(hss)%2 == 0 {
		return nil, xerrors.New("卷积核行数需要是奇数")
	} else if len(hss[0])%2 == 0 {
		return nil, xerrors.New("卷积核列数需要是奇数")
	}
	sem := semaphore.NewWeighted(int64(runtime.NumCPU() * 1024))
	wg := new(sync.WaitGroup)
	tss := make([][]T, len(xss)+(len(hss)-1)*2)
	for m := range tss {
		sem.Acquire(context.Background(), 1)
		wg.Add(1)
		go func(m int) {
			defer wg.Done()
			defer sem.Release(1)
			tss[m] = make([]T, len(xss[0])+(len(hss[0])-1)*2)
			for n := range tss[m] {
				if m-(len(hss)-1) >= 0 && m-(len(hss)-1) < len(xss) &&
					n-(len(hss[0])-1) >= 0 && n-(len(hss[0])-1) < len(xss[0]) {
					tss[m][n] = xss[m-(len(hss)-1)][n-(len(hss[0])-1)]
				}
			}
		}(m)
	}
	wg.Wait()
	yss := make([][]T, len(xss)+(len(hss)-1))
	for m := range yss {
		sem.Acquire(context.Background(), 1)
		wg.Add(1)
		go func(m int) {
			defer wg.Done()
			defer sem.Release(1)
			yss[m] = make([]T, len(xss[0])+(len(hss[0])-1))
			for n := range yss[m] {
				for i := range hss {
					for j := range hss[i] {
						yss[m][n] += tss[m+i][n+j] * hss[i][j]
					}
				}
			}
		}(m)
	}
	wg.Wait()
	return yss, nil
}

// 卷积核滤波same模式
// xss为原矩阵
// tss为原矩阵补0
// yss为m*n矩阵 m行n列
// hss为i*j矩阵 i行j列
func filter2same[T ~int | ~float32 | ~float64](hss, xss [][]T) ([][]T, error) {
	if len(hss)%2 == 0 {
		return nil, xerrors.New("卷积核行数需要是奇数")
	} else if len(hss[0])%2 == 0 {
		return nil, xerrors.New("卷积核列数需要是奇数")
	}
	sem := semaphore.NewWeighted(int64(runtime.NumCPU() * 1024))
	wg := new(sync.WaitGroup)
	tss := make([][]T, len(xss)+(len(hss)-1))
	for m := range tss {
		sem.Acquire(context.Background(), 1)
		wg.Add(1)
		go func(m int) {
			defer wg.Done()
			defer sem.Release(1)
			tss[m] = make([]T, len(xss[0])+(len(hss[0])-1))
			for n := range tss[m] {
				if m-(len(hss)-1)/2 >= 0 && m-(len(hss)-1)/2 < len(xss) &&
					n-(len(hss[0])-1)/2 >= 0 && n-(len(hss[0])-1)/2 < len(xss[0]) {
					tss[m][n] = xss[m-(len(hss)-1)/2][n-(len(hss[0])-1)/2]
				}
			}
		}(m)
	}
	wg.Wait()
	yss := make([][]T, len(xss))
	for m := range yss {
		sem.Acquire(context.Background(), 1)
		wg.Add(1)
		go func(m int) {
			defer wg.Done()
			defer sem.Release(1)
			yss[m] = make([]T, len(xss[0]))
			for n := range yss[m] {
				for i := range hss {
					for j := range hss[i] {
						yss[m][n] += tss[m+i][n+j] * hss[i][j]
					}
				}
			}
		}(m)
	}
	wg.Wait()
	return yss, nil
}

// 卷积核滤波valid模式
// xss为原矩阵
// yss为m*n矩阵 m行n列
// hss为i*j矩阵 i行j列
func filter2valid[T ~int | ~float32 | ~float64](hss, xss [][]T) ([][]T, error) {
	if len(hss)%2 == 0 {
		return nil, xerrors.New("卷积核行数需要是奇数")
	} else if len(hss[0])%2 == 0 {
		return nil, xerrors.New("卷积核列数需要是奇数")
	} else if len(hss) > len(xss) {
		return nil, xerrors.New("卷积核行数需要小于矩阵的行数")
	} else if len(hss[0]) > len(xss[0]) {
		return nil, xerrors.New("卷积核列数需要小于矩阵的列数")
	}
	sem := semaphore.NewWeighted(int64(runtime.NumCPU() * 1024))
	wg := new(sync.WaitGroup)
	yss := make([][]T, len(xss)-(len(hss)-1))
	for m := range yss {
		sem.Acquire(context.Background(), 1)
		wg.Add(1)
		go func(m int) {
			defer wg.Done()
			defer sem.Release(1)
			yss[m] = make([]T, len(xss[0])-(len(hss[0])-1))
			for n := range yss[m] {
				for i := range hss {
					for j := range hss[i] {
						yss[m][n] += xss[m+i][n+j] * hss[i][j]
					}
				}
			}
		}(m)
	}
	wg.Wait()
	return yss, nil
}
