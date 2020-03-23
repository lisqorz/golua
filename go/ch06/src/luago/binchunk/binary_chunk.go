package binchunk

type binaryChunk struct {
	header
	sizeUpvalues byte
	mainFunc     *Prototype
}

/**/
const (
	LUA_SIGNATURE    = "\x1bLua"
	LUAC_VERSION     = 0x53
	LUAC_FORMAT      = 0
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSIZET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INFEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

/*文件头*/
type header struct {
	signature       [4]byte
	version         byte
	format          byte
	luacData        [6]byte
	cintSize        byte
	sizetSize       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}

const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x03
	TAG_INTEGER   = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x14
)

/*函数原型*/
type Prototype struct {
	Source          string        // 源文件名
	LineDefined     uint32        // 起止行号 用于记录在源文件中的行号，普通函数应该都大于0，如果是主函数，则起止行号都是0
	LastLineDefined uint32        // -
	NumParams       byte          // 固定参数个数
	IsVararg        byte          // 是否为变参
	MaxStackSize    byte          // 寄存器数量
	Code            []uint32      // 指令表
	Constants       []interface{} // 常量表
	Upvalues        []Upvalue     // 目前只需要知道占用2字节，第10章会介绍闭包和Upvalue
	Protos          []*Prototype  // 子函数 原型连
	LineInfo        []uint32      // 行号
	LocVars         []LocVar      // 局部变量
	UpvalueNames    []string      // Upvalue名列表
}

type Upvalue struct {
	Instack byte
	Idx     byte
}

type LocVar struct {
	VarName string
	StartPC uint32
	EndPC   uint32
}

func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()
	reader.readByte()
	return reader.readProto("")
}
