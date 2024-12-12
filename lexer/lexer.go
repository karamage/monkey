package lexer

type Lexer struct {
	input        string
	position     int  // 現在の文字の位置
	readPosition int  // これから読み込む位置
	ch           byte // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCIIコードでNULに対応する
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
