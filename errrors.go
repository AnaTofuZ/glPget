package glPget

// error wrap系のmethod管理

// 純粋なerrorを取ってくるインターフェイス
// http://deeeet.com/writing/2016/04/25/go-pkg-errors/ を参照
type causer interface {
	Cause() error
}

type exiter interface {
	ExitCode() int
}

func (glp *glPget) ErrTrap(err error) (int, error) {
	for e := err; e != nil; {
		switch e.(type) {
		case exiter: // errorの返り値をよしなに変更
			return e.(exiter).ExitCode(), e
		case causer:
			e = e.(causer).Cause()
		default:
			return 1, e
		}
	}
	return 0, nil
}
