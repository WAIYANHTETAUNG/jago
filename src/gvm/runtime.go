package gvm

import "unsafe"

const (
	/* ----- CONSTANTS -----------*/
	NOP = 0		//(0X00)
	ACONST_NULL = 1		//(0X01)
	ICONST_M1 = 2		//(0X02)
	ICONST_0 = 3		//(0X03)
	ICONST_1 = 4		//(0X04)
	ICONST_2 = 05		//(0X05)
	ICONST_3 = 06		//(0X06)
	ICONST_4 = 07		//(0X07)
	ICONST_5 = 8		//(0X08)
	LCONST_0 = 9		//(0X09)
	LCONST_1 = 10		//(0X0A)
	FCONST_0 = 11		//(0X0B)
	FCONST_1 = 12		//(0X0C)
	FCONST_2 = 13		//(0X0D)
	DCONST_0 = 14		//(0X0E)
	DCONST_1 = 15		//(0X0F)
	BIPUSH = 16		//(0X10)
	SIPUSH = 17		//(0X11)
	LDC = 18		//(0X12)
	LDC_W = 19		//(0X13)
	LDC2_W = 20		//(0X14)

	/*--------LOADS ----------------*/
	ILOAD = 21		//(0X15)
	LLOAD = 22		//(0X16)
	FLOAD = 23		//(0X17)
	DLOAD = 24		//(0X18)
	ALOAD = 25		//(0X19)
	ILOAD_0 = 26		//(0X1A)
	ILOAD_1 = 27		//(0X1B)
	ILOAD_2 = 28		//(0X1C)
	ILOAD_3 = 29		//(0X1D)
	LLOAD_0 = 30		//(0X1E)
	LLOAD_1 = 31		//(0X1F)
	LLOAD_2 = 32		//(0X20)
	LLOAD_3 = 33		//(0X21)
	FLOAD_0 = 34		//(0X22)
	FLOAD_1 = 35		//(0X23)
	FLOAD_2 = 36		//(0X24)
	FLOAD_3 = 37		//(0X25)
	DLOAD_0 = 38		//(0X26)
	DLOAD_1 = 39		//(0X27)
	DLOAD_2 = 40		//(0X28)
	DLOAD_3 = 41		//(0X29)
	ALOAD_0 = 42		//(0X2A)
	ALOAD_1 = 43		//(0X2B)
	ALOAD_2 = 44		//(0X2C)
	ALOAD_3 = 45		//(0X2D)
	IALOAD = 46		//(0X2E)
	LALOAD = 47		//(0X2F)
	FALOAD = 48		//(0X30)
	DALOAD = 49		//(0X31)
	AALOAD = 50		//(0X32)
	BALOAD = 51		//(0X33)
	CALOAD = 52		//(0X34)
	SALOAD = 53		//(0X35)

	/*--------STORES ----------------*/
	ISTORE = 54		//(0X36)
	LSTORE = 55		//(0X37)
	FSTORE = 56		//(0X38)
	DSTORE = 57		//(0X39)
	ASTORE = 58		//(0X3A)
	ISTORE_0 = 59		//(0X3B)
	ISTORE_1 = 60		//(0X3C)
	ISTORE_2 = 61		//(0X3D)
	ISTORE_3 = 62		//(0X3E)
	LSTORE_0 = 63		//(0X3F)
	LSTORE_1 = 64		//(0X40)
	LSTORE_2 = 65		//(0X41)
	LSTORE_3 = 66		//(0X42)
	FSTORE_0 = 67		//(0X43)
	FSTORE_1 = 68		//(0X44)
	FSTORE_2 = 69		//(0X45)
	FSTORE_3 = 70		//(0X46)
	DSTORE_0 = 71		//(0X47)
	DSTORE_1 = 72		//(0X48)
	DSTORE_2 = 73		//(0X49)
	DSTORE_3 = 74		//(0X4A)
	ASTORE_0 = 75		//(0X4B)
	ASTORE_1 = 76		//(0X4C)
	ASTORE_2 = 77		//(0X4D)
	ASTORE_3 = 78		//(0X4E)
	IASTORE = 79		//(0X4F)
	LASTORE = 80		//(0X50)
	FASTORE = 81		//(0X51)
	DASTORE = 82		//(0X52)
	AASTORE = 83		//(0X53)
	BASTORE = 84		//(0X54)
	CASTORE = 85		//(0X55)
	SASTORE = 86		//(0X56)

	/*--------STACK---------------*/
	POP = 87		//(0X57)
	POP2 = 88		//(0X58)
	DUP = 89		//(0X59)
	DUP_X1 = 90		//(0X5A)
	DUP_X2 = 91		//(0X5B)
	DUP2 = 92		//(0X5C)
	DUP2_X1 = 93		//(0X5D)
	DUP2_X2 = 94		//(0X5E)
	SWAP = 95		//(0X5F)

	/*---------MATH -------------*/
	IADD = 96		//(0X60)
	LADD = 97		//(0X61)
	FADD = 98		//(0X62)
	DADD = 99		//(0X63)
	ISUB = 100		//(0X64)
	LSUB = 101		//(0X65)
	FSUB = 102		//(0X66)
	DSUB = 103		//(0X67)
	IMUL = 104		//(0X68)
	LMUL = 105		//(0X69)
	FMUL = 106		//(0X6A)
	DMUL = 107		//(0X6B)
	IDIV = 108		//(0X6C)
	LDIV = 109		//(0X6D)
	FDIV = 110		//(0X6E)
	DDIV = 111		//(0X6F)
	IREM = 112		//(0X70)
	LREM = 113		//(0X71)
	FREM = 114		//(0X72)
	DREM = 115		//(0X73)
	INEG = 116		//(0X74)
	LNEG = 117		//(0X75)
	FNEG = 118		//(0X76)
	DNEG = 119		//(0X77)
	ISHL = 120		//(0X78)
	LSHL = 121		//(0X79)
	ISHR = 122		//(0X7A)
	LSHR = 123		//(0X7B)
	IUSHR = 124		//(0X7C)
	LUSHR = 125		//(0X7D)
	IAND = 126		//(0X7E)
	LAND = 127		//(0X7F)
	IOR = 128		//(0X80)
	LOR = 129		//(0X81)
	IXOR = 130		//(0X82)
	LXOR = 131		//(0X83)
	IINC = 132		//(0X84)

	/*--------CONVERSIONS-----------*/
	I2L = 133		//(0X85)
	I2F = 134		//(0X86)
	I2D = 135		//(0X87)
	L2I = 136		//(0X88)
	L2F = 137		//(0X89)
	L2D = 138		//(0X8A)
	F2I = 139		//(0X8B)
	F2L = 140		//(0X8C)
	F2D = 141		//(0X8D)
	D2I = 142		//(0X8E)
	D2L = 143		//(0X8F)
	D2F = 144		//(0X90)
	I2B = 145		//(0X91)
	I2C = 146		//(0X92)
	I2S = 147		//(0X93)

	/*-----------COMPARASON -----------*/
	LCMP = 148		//(0X94)
	FCMPL = 149		//(0X95)
	FCMPG = 150		//(0X96)
	DCMPL = 151		//(0X97)
	DCMPG = 152		//(0X98)
	IFEQ = 153		//(0X99)
	IFNE = 154		//(0X9A)
	IFLT = 155		//(0X9B)
	IFGE = 156		//(0X9C)
	IFGT = 157		//(0X9D)
	IFLE = 158		//(0X9E)
	IF_ICMPEQ = 159		//(0X9F)
	IF_ICMPNE = 160		//(0XA0)
	IF_ICMPLT = 161		//(0XA1)
	IF_ICMPGE = 162		//(0XA2)
	IF_ICMPGT = 163		//(0XA3)
	IF_ICMPLE = 164		//(0XA4)
	IF_ACMPEQ = 165		//(0XA5)
	IF_ACMPNE = 166		//(0XA6)

	/*---------REFERENCES -------------*/
	GETSTATIC = 178		//(0XB2)
	PUTSTATIC = 179		//(0XB3)
	GETFIELD = 180		//(0XB4)
	PUTFIELD = 181		//(0XB5)
	INVOKEVIRTUAL = 182		//(0XB6)
	INVOKESPECIAL = 183		//(0XB7)
	INVOKESTATIC = 184		//(0XB8)
	INVOKEINTERFACE = 185		//(0XB9)
	INVOKEDYNAMIC = 186		//(0XBA)
	NEW = 187		//(0XBB)
	NEWARRAY = 188		//(0XBC)
	ANEWARRAY = 189		//(0XBD)
	ARRAYLENGTH = 190		//(0XBE)
	ATHROW = 191		//(0XBF)
	CHECKCAST = 192		//(0XC0)
	INSTANCEOF = 193		//(0XC1)
	MONITORENTER = 194		//(0XC2)
	MONITOREXIT = 195		//(0XC3)

	/*-------CONTROLS------------------*/
	GOTO = 167		//(0XA7)
	JSR = 168		//(0XA8)
	RET = 169		//(0XA9)
	TABLESWITCH = 170		//(0XAA)
	LOOKUPSWITCH = 171		//(0XAB)
	IRETURN = 172		//(0XAC)
	LRETURN = 173		//(0XAD)
	FRETURN = 174		//(0XAE)
	DRETURN = 175		//(0XAF)
	ARETURN = 176		//(0XB0)
	RETURN = 177		//(0XB1)

	/*--------EXTENDED-----------------*/
	WIDE = 196		//(0XC4)
	MULTIANEWARRAY = 197		//(0XC5)
	IFNULL = 198		//(0XC6)
	IFNONNULL = 199		//(0XC7)
	GOTO_W = 200		//(0XC8)
	JSR_W = 201		//(0XC9)

	/*----------RESERVED ---------------*/
	BREAKPOINT = 202		//(0XCA)
	IMPDEP1 = 254		//(0XFE)
	IMPDEP2 = 255		//(0XFF)
)


type Thread struct {
	vmStack VMStack
}


type StackFrame struct {
	method *JavaMethod
	// if this frame is current frame, the pc is for the pc of this thread;
	// otherwise, it is a snapshot one since the last time
	pc uint32
	localVariables []uint64
	// operand stack
	operandStack []uint64
	operandStackSize uint
}

func NewStackFrame(method *JavaMethod) *StackFrame {
	stackFrame := &StackFrame{
		method: method,
		pc: 0,
		localVariables: make([]uint64, method.maxLocals),
		operandStack: make([]uint64, method.maxStack),
		operandStackSize: 0}
	return stackFrame
}

func (this *StackFrame) loadByteVar(index uint) java_byte  {
	return java_byte(this.localVariables[index])
}

func (this *StackFrame) storeByteVar(index uint, value java_byte)  {
	this.localVariables[index] = uint64(value) & uint64(0xFF)
}

func (this *StackFrame) loadCharVar(index uint) java_char  {
	return java_char(this.localVariables[index])
}

func (this *StackFrame) storeCharVar(index uint, value java_char)  {
	this.localVariables[index] = uint64(value) & uint64(0xFF)
}

func (this *StackFrame) loadShortVar(index uint) java_short  {
	return java_short(this.localVariables[index])
}

func (this *StackFrame) storeShortVar(index uint, value java_short)  {
	this.localVariables[index] = uint64(value) & uint64(0xFFFF)
}

func (this *StackFrame) loadIntVar(index uint) java_int  {
	return java_int(this.localVariables[index])
}

func (this *StackFrame) storeIntVar(index uint, value java_int)  {
	this.localVariables[index] = uint64(value) & uint64(0xFFFFFFFF)
}

func (this *StackFrame) loadLongVar(index uint) java_long  {
	return java_long(this.localVariables[index])
}

func (this *StackFrame) storeLongVar(index uint, value java_long)  {
	this.localVariables[index] = uint64(value) & uint64(0xFFFFFFFFFFFFFFFF)
}

func (this *StackFrame) loadFloatVar(index uint) java_float  {
	return java_float(this.localVariables[index])
}

func (this *StackFrame) storeFloatVar(index uint, value java_float)  {
	this.localVariables[index] = uint64(value) & uint64(0xFFFFFFFF)
}

func (this *StackFrame) loadDoubleVar(index uint) java_double  {
	return java_double(this.localVariables[index])
}

func (this *StackFrame) storeDoubleVar(index uint, value java_double)  {
	this.localVariables[index] = uint64(value) & uint64(0xFFFFFFFFFFFFFFFF)
}

func (this *StackFrame) loadBooleanVar(index uint) java_boolean  {
	return java_boolean(this.localVariables[index])
}

func (this *StackFrame) storeBooleanVar(index uint, value java_boolean)  {
	this.localVariables[index] = uint64(value) & uint64(0x1)
}

func (this *StackFrame) loadReferenceVar(index uint) java_reference  {
	return java_reference(unsafe.Pointer(uintptr(this.localVariables[index])))
}

func (this *StackFrame) storeReferenceVar(index uint, value java_reference)  {
	this.localVariables[index] = uint64(unsafe.Alignof((value)))
}

func (this *StackFrame) loadArrayVar(index uint) java_array  {
	return java_array(unsafe.Pointer(uintptr(this.localVariables[index])))
}

func (this *StackFrame) storeArrayVar(index uint, value java_array)  {
	this.localVariables[index] = uint64(unsafe.Alignof((value)))
}

func (this *StackFrame) passParameters(callee *StackFrame)  {
	method := callee.method
	for i := 0; i < len(method.parameterDescriptors); i++ {
		switch rune(method.parameterDescriptors[i][0]) {
		case 'B':callee.storeByteVar(uint(i), this.popByte())
		case 'C':callee.storeCharVar(uint(i), this.popChar())
		case 'D':callee.storeDoubleVar(uint(i), this.popDouble())
		case 'F':callee.storeFloatVar(uint(i), this.popFloat())
		case 'I':callee.storeIntVar(uint(i), this.popInt())
		case 'J':callee.storeLongVar(uint(i), this.popLong())
		case 'S':callee.storeShortVar(uint(i), this.popShort())
		case 'Z':callee.storeBooleanVar(uint(i), this.popBoolean())
		case 'L', '[':callee.storeReferenceVar(uint(i), this.popReference())
		}
	}
}

func (this *StackFrame) passReturn(caller *StackFrame)  {
	method := this.method
	switch rune(method.returnDescriptor[0]) {
	case 'B':caller.pushByte(this.popByte())
	case 'C':caller.pushChar(this.popChar())
	case 'D':caller.pushDouble(this.popDouble())
	case 'F':caller.pushFloat(this.popFloat())
	case 'I':caller.pushInt(this.popInt())
	case 'J':caller.pushLong(this.popLong())
	case 'S':caller.pushShort(this.popShort())
	case 'Z':caller.pushBoolean(this.popBoolean())
	case 'L', '[':caller.pushReference(this.popReference())
	}
}

func (this *StackFrame) pushByte(jbyte java_byte)  {
	this.operandStack[this.operandStackSize] = uint64(jbyte) & uint64(0xFF)
	this.operandStackSize++
}

func (this *StackFrame) popByte() java_byte {
	jbyte := java_byte(this.operandStack[this.operandStackSize-1])
	this.operandStackSize--
	return jbyte
}

func (this *StackFrame) pushShort(jshort java_short) {
	this.operandStack[this.operandStackSize] = uint64(jshort) & uint64(0xFFFF)
	this.operandStackSize++
}

func (this *StackFrame) popShort() java_short {
	jshort := java_short(this.operandStack[this.operandStackSize-1])
	this.operandStackSize--
	return jshort
}

func (this *StackFrame) pushChar(jchar java_char)  {
	this.operandStack[this.operandStackSize] = uint64(jchar) & uint64(0xFFFF)
	this.operandStackSize++
}

func (this *StackFrame) popChar() java_char {
	jchar := java_char(this.operandStack[this.operandStackSize-1])
	this.operandStackSize--
	return jchar
}

func (this *StackFrame) pushInt(jint java_int)  {
	this.operandStack[this.operandStackSize] = uint64(jint) & uint64(0xFFFFFFFF)
	this.operandStackSize++
}

func (this *StackFrame) popInt() java_int {
	jint := java_int(this.operandStack[this.operandStackSize-1])
	this.operandStackSize--
	return jint
}

func (this *StackFrame) pushLong(jlong java_long)  {
	this.operandStack[this.operandStackSize] = uint64(jlong) & uint64(0xFFFFFFFFFFFFFFFF)
	this.operandStackSize++
}

func (this *StackFrame) popLong() java_long {
	jlong := java_long(this.operandStack[this.operandStackSize-1])
	this.operandStackSize--
	return jlong
}

func (this *StackFrame) pushFloat(jfloat java_float)  {
	this.operandStack[this.operandStackSize] = uint64(jfloat) & uint64(0xFFFFFFFF)
	this.operandStackSize++
}

func (this *StackFrame) popFloat() java_float {
	jfloat := java_float(this.operandStack[this.operandStackSize-1])
	this.operandStackSize--
	return jfloat
}

func (this *StackFrame) pushDouble(jdouble java_double)  {
	this.operandStack[this.operandStackSize] = uint64(jdouble) & uint64(0xFFFFFFFF)
	this.operandStackSize++
}

func (this *StackFrame) popDouble() java_double {
	jdouble := java_double(this.operandStack[this.operandStackSize-1])
	this.operandStackSize--
	return jdouble
}

func (this *StackFrame) pushBoolean(jboolean java_boolean)  {
	this.operandStack[this.operandStackSize] = uint64(jboolean) & uint64(0x1)
	this.operandStackSize++
}

func (this *StackFrame) popBoolean() java_boolean {
	jboolean := java_boolean(this.operandStack[this.operandStackSize-1])
	this.operandStackSize--
	return jboolean
}

func (this *StackFrame) pushReference(jreference java_reference)  {
	this.operandStack[this.operandStackSize] = uint64(unsafe.Alignof(jreference))
	this.operandStackSize++
}

func (this *StackFrame) popReference() java_reference {
	jreference := java_reference(unsafe.Pointer(uintptr(this.operandStack[this.operandStackSize-1])))
	this.operandStackSize--
	return jreference
}

func (this *StackFrame) pushArray(jreference java_array)  {
	this.operandStack[this.operandStackSize] = uint64(unsafe.Alignof(jreference))
	this.operandStackSize++
}

func (this *StackFrame) popArray() java_array {
	jreference := java_array(unsafe.Pointer(uintptr(this.operandStack[this.operandStackSize-1])))
	this.operandStackSize--
	return jreference
}

type VMStack struct {
	stackFrames []*StackFrame
	size uint
	capacity uint
}

func (this *VMStack) push(stackFrame *StackFrame)  {
	this.stackFrames[this.size] = stackFrame
	this.size++
}

func (this *VMStack) pop()  {
	if this.size == 0 {
		return
	}
	this.stackFrames[this.size-1] = nil
	this.size--
}

func (this *VMStack) peek() *StackFrame  {
	return this.stackFrames[this.size-1]
}