package dm01

import (
    "fmt"
)

type FileDmObject struct {
    Name string
    Size float64
}

type FileDmNotFound struct {
	Full string
	Spec string
}

func (e *FileDmNotFound) Error() string {
	return fmt.Sprintf("yaml: %s: %q not found", e.Full, e.Spec)
}

type FileTypeDmMismatch struct {
	Full     string
	Spec     string
	Token    string
    Name     string
    Size     float64
	Expected string
}

func (e *FileTypeDmMismatch) ShowInfo() string {
	return fmt.Sprintf("yaml: %s: type mismatch: %q is %T, want %s (at %q)",
		e.Full, e.Spec, e.Name, e.Expected, e.Token)
}

func (a FileTypeDmMismatch) ShowToken() {
    fmt.Println(a.Token)
}
func (a *FileTypeDmMismatch) FunPtrA1() {
    fmt.Println("FunPtrA1")
}
func (a *FileTypeDmMismatch) FunPtrA2() {
    fmt.Println("FunPtrA2")
    fmt.Println(a.Token)
}

