****# Notes

Instructions found at the following site

https://golang.cafe/blog/how-to-list-files-in-a-directory-in-go.html

### Option 1: ioutil.ReadDir
ioutil.ReadDir comes from the ioutil package in the Go standard library, you can check the documentation from the official Go Doc website

```func ReadDir(dirname string) ([]os.FileInfo, error)```

ReadDir reads the directory named by dirname and returns a list of directory entries sorted by filename.

---
### Option 2: filepath.Walk
filepath.Walk is another option you can use to list files in a directory structure, from the filepath Go package, it also allows us to recursively discover directories and files. The official documentation reads as follows

```func Walk(root string, walkFn WalkFunc) error```

Walk walks the file tree rooted at root, calling walkFn for each file or directory in the tree, including root. All errors that arise visiting files and directories are filtered by walkFn. The files are walked in lexical order, which makes the output deterministic but means that for very large directories Walk can be inefficient. Walk does not follow symbolic links.

---
### Option 3: os.File.Readdir
The last option is using just the file pointer coming from the os.File when reading a file from the file system. The package is the os package in the Go standard library.

```func (f *File) Readdir(n int) ([]FileInfo, error)```

Readdir reads the contents of the directory associated with file and returns a slice of up to n FileInfo values, as would be returned by Lstat, in directory order. Subsequent calls on the same file will yield further FileInfos.
If n > 0, Readdir returns at most n FileInfo structures. In this case, if Readdir returns an empty slice, it will return a non-nil error explaining why. At the end of a directory, the error is io.EOF.
If n <= 0, Readdir returns all the FileInfo from the directory in a single sli****

