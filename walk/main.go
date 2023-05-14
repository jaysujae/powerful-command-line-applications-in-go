package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	ext     string
	size    int64
	list    bool
	del     bool
	wLog    io.Writer
	archive string
}

func main() {
	root := flag.String("root", ".", "root")
	list := flag.Bool("list", false, "list file")
	ext := flag.String("ext", "", "File extension to filter")
	size := flag.Int64("size", 0, "minimum file size")
	del := flag.Bool("del", false, "Delete files")
	logFile := flag.String("log", "", "Log deletes to this file")
	archive := flag.String("archive", "", "Archive directory")
	flag.Parse()

	var (
		f   = os.Stdout
		err error
	)
	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	c := config{
		ext:     *ext,
		size:    *size,
		list:    *list,
		del:     *del,
		wLog:    f,
		archive: *archive,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		panic(err)
	}
}

func run(root string, out io.Writer, cfg config) error {

	delLogger := log.New(cfg.wLog, "Deleted file: ", log.LstdFlags)

	return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filterOut(path, cfg.ext, cfg.size, info) {
			return nil
		}
		if cfg.list {
			return listFile(path, out)
		}
		if cfg.del {
			return delFile(path, delLogger)
		}

		if cfg.archive != "" {
			if err := archiveFile(cfg.archive, root, path); err != nil {
				return err
			}
		}

		return listFile(path, out)
	})
}
