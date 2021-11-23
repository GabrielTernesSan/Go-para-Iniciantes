# A interface Writer

- A interface wirter do pacote io.
- Estrutura ``` type Writer interface {Write(p []byte) (n int, err error)}```. A função recebe uma fatia de bytes e retorna um inteiro e um erro.
  - pkg os: ```func (f *File) Write(b []byte) (n int, err error)```
  - pkg json: ```func NewEncoder(w io.Writer) *Encoder```
- "Println [...] writes to standard output."
  - func Println [...] return Fprintln(os.Stdout, a...)
  - func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
  - Stdout: NewFile(uintptr(syscall.Stdout), "/dev/stdout") (Google: Standard streams)
  - func NewFile(fd uintptr, name string) *File
  - func (f *File) Write(b []byte) (n int, err error)

### Preciso estudar mais isso