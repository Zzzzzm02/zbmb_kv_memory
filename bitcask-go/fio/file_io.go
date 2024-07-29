package fio

import "os"

// FileIO 标准系统文件 IO
type FileIO struct {
	fd *os.File // 系统文件描述符

}

// NewFileIOManager 初始化文件IO方法
func NewFileIOManager(fileName string) (*FileIO, error) {
	fd, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		DataFileParm,
	)
	// 最好去了解一下这个接口
	if err != nil {
		return nil, err // 如果有错误，可以去让上层调用者处理
	}
	// 没问题就返回
	return &FileIO{fd: fd}, nil
}

func (fio *FileIO) Read(b []byte, offset int64) (int, error) {
	return fio.fd.ReadAt(b, offset)
}

func (fio *FileIO) Write(b []byte) (int, error) {
	return fio.fd.Write(b)
}

func (fio *FileIO) Sync() error {
	return fio.fd.Sync()
}

func (fio *FileIO) Close() error {
	return fio.fd.Close()
}