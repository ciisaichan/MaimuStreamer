package global

import (
	"io"
	"net/http"
	"os"
)

type FileDownloader struct {
	Reader *Reader
	HttpResponse *http.Response
	Url string
	FileName string
	Downloading bool
	Error error
}

type Reader struct {
	IoReader io.Reader
	Total int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error){
	n, err = r.IoReader.Read(p)

	r.Current += int64(n)
	//fmt.Printf("\r已下载: %.2f MB", float64(r.Current)/1024/1024)

	return n,err
}

func (downloader *FileDownloader) Start() {
	go downloader.downloadProgress()
}

func (downloader *FileDownloader) Stop() {
	downloader.HttpResponse.Body.Close()
}

func (downloader *FileDownloader) downloadProgress() {
	//var err error = nil
	client := http.Client{}
	request, err := http.NewRequest("GET", downloader.Url, nil)
	if err != nil {
		downloader.finishProgress(err)
		return
	}

	request.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36 Edg/97.0.1072.62")

	downloader.HttpResponse, err = client.Do(request)
	if err != nil {
		downloader.finishProgress(err)
		return
	}
	defer downloader.HttpResponse.Body.Close()
	/*
	downloader.HttpResponse, err = http.Get(downloader.Url)
	defer func() {_ = downloader.HttpResponse.Body.Close()}()
	if err != nil {
		downloader.finishProgress(err)
		return
	}
	 */

	f, err2 := os.Create(downloader.FileName)
	if err2 != nil {
		downloader.finishProgress(err2)
		return
	}
	defer func() {_ = f.Close()}()

	downloader.Downloading = true

	downloader.Reader = &Reader{
		IoReader: downloader.HttpResponse.Body,
		Total: downloader.HttpResponse.ContentLength,
	}

	_, err = io.Copy(f, downloader.Reader)
	if err != nil {
		downloader.finishProgress(err)
		return
	}
	downloader.finishProgress(nil)
}

func (downloader *FileDownloader) finishProgress(err error) {
	downloader.Error = err
	downloader.Downloading = false
}