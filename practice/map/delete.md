"".main STEXT size=48 args=0x0 locals=0x8
	0x0000 00000 (delete.go:7)	TEXT	"".main(SB), $8-0
	0x0000 00000 (delete.go:7)	MOVQ	(TLS), CX
	0x0009 00009 (delete.go:7)	CMPQ	SP, 16(CX)
	0x000d 00013 (delete.go:7)	JLS	41
	0x000f 00015 (delete.go:7)	SUBQ	$8, SP
	0x0013 00019 (delete.go:7)	MOVQ	BP, (SP)
	0x0017 00023 (delete.go:7)	LEAQ	(SP), BP
	0x001b 00027 (delete.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (delete.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (delete.go:7)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (delete.go:8)	PCDATA	$2, $0
	0x001b 00027 (delete.go:8)	PCDATA	$0, $0
	0x001b 00027 (delete.go:8)	CALL	"".demo2(SB)
	0x0020 00032 (delete.go:9)	MOVQ	(SP), BP
	0x0024 00036 (delete.go:9)	ADDQ	$8, SP
	0x0028 00040 (delete.go:9)	RET
	0x0029 00041 (delete.go:9)	NOP
	0x0029 00041 (delete.go:7)	PCDATA	$0, $-1
	0x0029 00041 (delete.go:7)	PCDATA	$2, $-1
	0x0029 00041 (delete.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x002e 00046 (delete.go:7)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 1a 48  dH..%....H;a.v.H
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 e8 00 00 00 00  ...H.,$H.,$.....
	0x0020 48 8b 2c 24 48 83 c4 08 c3 e8 00 00 00 00 eb d0  H.,$H...........
	rel 5+4 t=16 TLS+0
	rel 28+4 t=8 "".demo2+0
	rel 42+4 t=8 runtime.morestack_noctxt+0
"".demo2 STEXT size=684 args=0x0 locals=0x100
	0x0000 00000 (delete.go:43)	TEXT	"".demo2(SB), $256-0
	0x0000 00000 (delete.go:43)	MOVQ	(TLS), CX
	0x0009 00009 (delete.go:43)	LEAQ	-128(SP), AX
	0x000e 00014 (delete.go:43)	CMPQ	AX, 16(CX)
	0x0012 00018 (delete.go:43)	JLS	674
	0x0018 00024 (delete.go:43)	SUBQ	$256, SP
	0x001f 00031 (delete.go:43)	MOVQ	BP, 248(SP)
	0x0027 00039 (delete.go:43)	LEAQ	248(SP), BP
	0x002f 00047 (delete.go:43)	FUNCDATA	$0, gclocals·f5be5308b59e045b7c5b33ee8908cfb7(SB)
	0x002f 00047 (delete.go:43)	FUNCDATA	$1, gclocals·0970b3a362ebabf95df0bb0ec675cbde(SB)
	0x002f 00047 (delete.go:43)	FUNCDATA	$3, gclocals·0879caa565d3c697358d27e15966481d(SB)
	0x002f 00047 (delete.go:44)	PCDATA	$2, $0
	0x002f 00047 (delete.go:44)	PCDATA	$0, $0
	0x002f 00047 (delete.go:44)	CALL	runtime.makemap_small(SB)
	0x0034 00052 (delete.go:44)	PCDATA	$2, $1
	0x0034 00052 (delete.go:44)	MOVQ	(SP), AX
	0x0038 00056 (delete.go:44)	PCDATA	$0, $1
	0x0038 00056 (delete.go:44)	MOVQ	AX, "".dm+80(SP)
	0x003d 00061 (delete.go:44)	MOVL	$1, CX
	0x0042 00066 (delete.go:45)	JMP	120
	0x0044 00068 (delete.go:45)	MOVQ	CX, "".i+64(SP)
	0x0049 00073 (delete.go:46)	PCDATA	$2, $2
	0x0049 00073 (delete.go:46)	LEAQ	type.map[int]int(SB), DX
	0x0050 00080 (delete.go:46)	PCDATA	$2, $1
	0x0050 00080 (delete.go:46)	MOVQ	DX, (SP)
	0x0054 00084 (delete.go:46)	PCDATA	$2, $0
	0x0054 00084 (delete.go:46)	MOVQ	AX, 8(SP)
	0x0059 00089 (delete.go:46)	MOVQ	CX, 16(SP)
	0x005e 00094 (delete.go:46)	CALL	runtime.mapassign_fast64(SB)
	0x0063 00099 (delete.go:46)	PCDATA	$2, $1
	0x0063 00099 (delete.go:46)	MOVQ	24(SP), AX
	0x0068 00104 (delete.go:46)	MOVQ	"".i+64(SP), CX
	0x006d 00109 (delete.go:46)	PCDATA	$2, $0
	0x006d 00109 (delete.go:46)	MOVQ	CX, (AX)
	0x0070 00112 (delete.go:45)	INCQ	CX
	0x0073 00115 (delete.go:46)	PCDATA	$2, $1
	0x0073 00115 (delete.go:46)	MOVQ	"".dm+80(SP), AX
	0x0078 00120 (delete.go:45)	CMPQ	CX, $11
	0x007c 00124 (delete.go:45)	JLT	68
	0x007e 00126 (delete.go:51)	PCDATA	$2, $3
	0x007e 00126 (delete.go:51)	PCDATA	$0, $2
	0x007e 00126 (delete.go:51)	LEAQ	""..autotmp_6+152(SP), DI
	0x0086 00134 (delete.go:51)	XORPS	X0, X0
	0x0089 00137 (delete.go:51)	PCDATA	$2, $1
	0x0089 00137 (delete.go:51)	LEAQ	-32(DI), DI
	0x008d 00141 (delete.go:51)	DUFFZERO	$273
	0x00a0 00160 (delete.go:51)	PCDATA	$2, $4
	0x00a0 00160 (delete.go:51)	LEAQ	type.map[int]int(SB), CX
	0x00a7 00167 (delete.go:51)	PCDATA	$2, $1
	0x00a7 00167 (delete.go:51)	MOVQ	CX, (SP)
	0x00ab 00171 (delete.go:51)	PCDATA	$2, $0
	0x00ab 00171 (delete.go:51)	MOVQ	AX, 8(SP)
	0x00b0 00176 (delete.go:51)	PCDATA	$2, $5
	0x00b0 00176 (delete.go:51)	LEAQ	""..autotmp_6+152(SP), DX
	0x00b8 00184 (delete.go:51)	PCDATA	$2, $0
	0x00b8 00184 (delete.go:51)	MOVQ	DX, 16(SP)
	0x00bd 00189 (delete.go:51)	CALL	runtime.mapiterinit(SB)
	0x00c2 00194 (delete.go:51)	PCDATA	$2, $1
	0x00c2 00194 (delete.go:51)	MOVQ	"".dm+80(SP), AX
	0x00c7 00199 (delete.go:51)	XORL	CX, CX
	0x00c9 00201 (delete.go:51)	JMP	387
	0x00ce 00206 (delete.go:53)	MOVQ	BX, "".i+72(SP)
	0x00d3 00211 (delete.go:54)	PCDATA	$2, $4
	0x00d3 00211 (delete.go:54)	LEAQ	type.map[int]int(SB), CX
	0x00da 00218 (delete.go:54)	PCDATA	$2, $1
	0x00da 00218 (delete.go:54)	MOVQ	CX, (SP)
	0x00de 00222 (delete.go:54)	PCDATA	$2, $0
	0x00de 00222 (delete.go:54)	MOVQ	AX, 8(SP)
	0x00e3 00227 (delete.go:54)	MOVQ	BX, 16(SP)
	0x00e8 00232 (delete.go:54)	CALL	runtime.mapassign_fast64(SB)
	0x00ed 00237 (delete.go:54)	PCDATA	$2, $1
	0x00ed 00237 (delete.go:54)	MOVQ	24(SP), AX
	0x00f2 00242 (delete.go:54)	MOVQ	"".i+72(SP), CX
	0x00f7 00247 (delete.go:54)	PCDATA	$2, $0
	0x00f7 00247 (delete.go:54)	MOVQ	CX, (AX)
	0x00fa 00250 (delete.go:53)	LEAQ	1(CX), BX
	0x00fe 00254 (delete.go:54)	PCDATA	$2, $1
	0x00fe 00254 (delete.go:54)	MOVQ	"".dm+80(SP), AX
	0x0103 00259 (delete.go:57)	MOVQ	"".num+48(SP), CX
	0x0108 00264 (delete.go:58)	MOVQ	"".k+56(SP), DX
	0x010d 00269 (delete.go:53)	CMPQ	BX, $21
	0x0111 00273 (delete.go:53)	JLT	206
	0x0113 00275 (delete.go:58)	PCDATA	$2, $0
	0x0113 00275 (delete.go:58)	PCDATA	$0, $3
	0x0113 00275 (delete.go:58)	XORPS	X0, X0
	0x0116 00278 (delete.go:58)	MOVUPS	X0, ""..autotmp_7+104(SP)
	0x011b 00283 (delete.go:58)	PCDATA	$2, $1
	0x011b 00283 (delete.go:58)	LEAQ	type.int(SB), AX
	0x0122 00290 (delete.go:58)	PCDATA	$2, $0
	0x0122 00290 (delete.go:58)	MOVQ	AX, (SP)
	0x0126 00294 (delete.go:58)	MOVQ	DX, 8(SP)
	0x012b 00299 (delete.go:58)	CALL	runtime.convT2E64(SB)
	0x0130 00304 (delete.go:58)	PCDATA	$2, $1
	0x0130 00304 (delete.go:58)	MOVQ	24(SP), AX
	0x0135 00309 (delete.go:58)	MOVQ	16(SP), CX
	0x013a 00314 (delete.go:58)	MOVQ	CX, ""..autotmp_7+104(SP)
	0x013f 00319 (delete.go:58)	PCDATA	$2, $0
	0x013f 00319 (delete.go:58)	MOVQ	AX, ""..autotmp_7+112(SP)
	0x0144 00324 (delete.go:58)	PCDATA	$2, $1
	0x0144 00324 (delete.go:58)	LEAQ	""..autotmp_7+104(SP), AX
	0x0149 00329 (delete.go:58)	PCDATA	$2, $0
	0x0149 00329 (delete.go:58)	MOVQ	AX, (SP)
	0x014d 00333 (delete.go:58)	MOVQ	$1, 8(SP)
	0x0156 00342 (delete.go:58)	MOVQ	$1, 16(SP)
	0x015f 00351 (delete.go:58)	CALL	fmt.Println(SB)
	0x0164 00356 (delete.go:51)	PCDATA	$2, $1
	0x0164 00356 (delete.go:51)	PCDATA	$0, $2
	0x0164 00356 (delete.go:51)	LEAQ	""..autotmp_6+152(SP), AX
	0x016c 00364 (delete.go:51)	PCDATA	$2, $0
	0x016c 00364 (delete.go:51)	MOVQ	AX, (SP)
	0x0170 00368 (delete.go:51)	CALL	runtime.mapiternext(SB)
	0x0175 00373 (delete.go:57)	MOVQ	"".num+48(SP), AX
	0x017a 00378 (delete.go:57)	LEAQ	1(AX), CX
	0x017e 00382 (delete.go:51)	PCDATA	$2, $1
	0x017e 00382 (delete.go:51)	MOVQ	"".dm+80(SP), AX
	0x0183 00387 (delete.go:51)	PCDATA	$2, $2
	0x0183 00387 (delete.go:51)	MOVQ	""..autotmp_6+152(SP), DX
	0x018b 00395 (delete.go:51)	TESTQ	DX, DX
	0x018e 00398 (delete.go:51)	JEQ	433
	0x0190 00400 (delete.go:57)	MOVQ	CX, "".num+48(SP)
	0x0195 00405 (delete.go:51)	PCDATA	$2, $1
	0x0195 00405 (delete.go:51)	MOVQ	(DX), DX
	0x0198 00408 (delete.go:52)	CMPQ	DX, $5
	0x019c 00412 (delete.go:52)	JNE	275
	0x01a2 00418 (delete.go:51)	MOVQ	DX, "".k+56(SP)
	0x01a7 00423 (delete.go:51)	MOVL	$12, BX
	0x01ac 00428 (delete.go:53)	JMP	269
	0x01b1 00433 (delete.go:60)	PCDATA	$2, $0
	0x01b1 00433 (delete.go:60)	PCDATA	$0, $4
	0x01b1 00433 (delete.go:60)	XORPS	X0, X0
	0x01b4 00436 (delete.go:60)	MOVUPS	X0, ""..autotmp_9+88(SP)
	0x01b9 00441 (delete.go:60)	PCDATA	$2, $1
	0x01b9 00441 (delete.go:60)	LEAQ	type.int(SB), AX
	0x01c0 00448 (delete.go:60)	PCDATA	$2, $0
	0x01c0 00448 (delete.go:60)	MOVQ	AX, (SP)
	0x01c4 00452 (delete.go:60)	MOVQ	CX, 8(SP)
	0x01c9 00457 (delete.go:60)	CALL	runtime.convT2E64(SB)
	0x01ce 00462 (delete.go:60)	PCDATA	$2, $1
	0x01ce 00462 (delete.go:60)	MOVQ	24(SP), AX
	0x01d3 00467 (delete.go:60)	MOVQ	16(SP), CX
	0x01d8 00472 (delete.go:60)	MOVQ	CX, ""..autotmp_9+88(SP)
	0x01dd 00477 (delete.go:60)	PCDATA	$2, $0
	0x01dd 00477 (delete.go:60)	MOVQ	AX, ""..autotmp_9+96(SP)
	0x01e2 00482 (delete.go:60)	PCDATA	$2, $1
	0x01e2 00482 (delete.go:60)	LEAQ	""..autotmp_9+88(SP), AX
	0x01e7 00487 (delete.go:60)	PCDATA	$2, $0
	0x01e7 00487 (delete.go:60)	MOVQ	AX, (SP)
	0x01eb 00491 (delete.go:60)	MOVQ	$1, 8(SP)
	0x01f4 00500 (delete.go:60)	MOVQ	$1, 16(SP)
	0x01fd 00509 (delete.go:60)	CALL	fmt.Println(SB)
	0x0202 00514 (delete.go:62)	PCDATA	$2, $1
	0x0202 00514 (delete.go:62)	PCDATA	$0, $1
	0x0202 00514 (delete.go:62)	MOVQ	"".dm+80(SP), AX
	0x0207 00519 (delete.go:62)	TESTQ	AX, AX
	0x020a 00522 (delete.go:62)	JEQ	667
	0x0210 00528 (delete.go:62)	PCDATA	$2, $0
	0x0210 00528 (delete.go:62)	MOVQ	(AX), CX
	0x0213 00531 (delete.go:62)	PCDATA	$0, $5
	0x0213 00531 (delete.go:62)	XORPS	X0, X0
	0x0216 00534 (delete.go:62)	MOVUPS	X0, ""..autotmp_11+120(SP)
	0x021b 00539 (delete.go:62)	MOVUPS	X0, ""..autotmp_11+136(SP)
	0x0223 00547 (delete.go:62)	PCDATA	$2, $1
	0x0223 00547 (delete.go:62)	LEAQ	type.int(SB), AX
	0x022a 00554 (delete.go:62)	PCDATA	$2, $0
	0x022a 00554 (delete.go:62)	MOVQ	AX, (SP)
	0x022e 00558 (delete.go:62)	MOVQ	CX, 8(SP)
	0x0233 00563 (delete.go:62)	CALL	runtime.convT2E64(SB)
	0x0238 00568 (delete.go:62)	MOVQ	16(SP), AX
	0x023d 00573 (delete.go:62)	PCDATA	$2, $6
	0x023d 00573 (delete.go:62)	MOVQ	24(SP), CX
	0x0242 00578 (delete.go:62)	MOVQ	AX, ""..autotmp_11+120(SP)
	0x0247 00583 (delete.go:62)	PCDATA	$2, $0
	0x0247 00583 (delete.go:62)	MOVQ	CX, ""..autotmp_11+128(SP)
	0x024f 00591 (delete.go:62)	PCDATA	$2, $1
	0x024f 00591 (delete.go:62)	LEAQ	type.map[int]int(SB), AX
	0x0256 00598 (delete.go:62)	PCDATA	$2, $0
	0x0256 00598 (delete.go:62)	MOVQ	AX, ""..autotmp_11+136(SP)
	0x025e 00606 (delete.go:62)	PCDATA	$2, $1
	0x025e 00606 (delete.go:62)	PCDATA	$0, $6
	0x025e 00606 (delete.go:62)	MOVQ	"".dm+80(SP), AX
	0x0263 00611 (delete.go:62)	PCDATA	$2, $0
	0x0263 00611 (delete.go:62)	MOVQ	AX, ""..autotmp_11+144(SP)
	0x026b 00619 (delete.go:62)	PCDATA	$2, $1
	0x026b 00619 (delete.go:62)	LEAQ	""..autotmp_11+120(SP), AX
	0x0270 00624 (delete.go:62)	PCDATA	$2, $0
	0x0270 00624 (delete.go:62)	MOVQ	AX, (SP)
	0x0274 00628 (delete.go:62)	MOVQ	$2, 8(SP)
	0x027d 00637 (delete.go:62)	MOVQ	$2, 16(SP)
	0x0286 00646 (delete.go:62)	CALL	fmt.Println(SB)
	0x028b 00651 (delete.go:63)	PCDATA	$0, $0
	0x028b 00651 (delete.go:63)	MOVQ	248(SP), BP
	0x0293 00659 (delete.go:63)	ADDQ	$256, SP
	0x029a 00666 (delete.go:63)	RET
	0x029b 00667 (delete.go:63)	PCDATA	$0, $1
	0x029b 00667 (delete.go:63)	XORL	CX, CX
	0x029d 00669 (delete.go:62)	JMP	531
	0x02a2 00674 (delete.go:62)	NOP
	0x02a2 00674 (delete.go:43)	PCDATA	$0, $-1
	0x02a2 00674 (delete.go:43)	PCDATA	$2, $-1
	0x02a2 00674 (delete.go:43)	CALL	runtime.morestack_noctxt(SB)
	0x02a7 00679 (delete.go:43)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 8d 44 24 80 48 3b  dH..%....H.D$.H;
	0x0010 41 10 0f 86 8a 02 00 00 48 81 ec 00 01 00 00 48  A.......H......H
	0x0020 89 ac 24 f8 00 00 00 48 8d ac 24 f8 00 00 00 e8  ..$....H..$.....
	0x0030 00 00 00 00 48 8b 04 24 48 89 44 24 50 b9 01 00  ....H..$H.D$P...
	0x0040 00 00 eb 34 48 89 4c 24 40 48 8d 15 00 00 00 00  ...4H.L$@H......
	0x0050 48 89 14 24 48 89 44 24 08 48 89 4c 24 10 e8 00  H..$H.D$.H.L$...
	0x0060 00 00 00 48 8b 44 24 18 48 8b 4c 24 40 48 89 08  ...H.D$.H.L$@H..
	0x0070 48 ff c1 48 8b 44 24 50 48 83 f9 0b 7c c6 48 8d  H..H.D$PH...|.H.
	0x0080 bc 24 98 00 00 00 0f 57 c0 48 8d 7f e0 48 89 6c  .$.....W.H...H.l
	0x0090 24 f0 48 8d 6c 24 f0 e8 00 00 00 00 48 8b 6d 00  $.H.l$......H.m.
	0x00a0 48 8d 0d 00 00 00 00 48 89 0c 24 48 89 44 24 08  H......H..$H.D$.
	0x00b0 48 8d 94 24 98 00 00 00 48 89 54 24 10 e8 00 00  H..$....H.T$....
	0x00c0 00 00 48 8b 44 24 50 31 c9 e9 b5 00 00 00 48 89  ..H.D$P1......H.
	0x00d0 5c 24 48 48 8d 0d 00 00 00 00 48 89 0c 24 48 89  \$HH......H..$H.
	0x00e0 44 24 08 48 89 5c 24 10 e8 00 00 00 00 48 8b 44  D$.H.\$......H.D
	0x00f0 24 18 48 8b 4c 24 48 48 89 08 48 8d 59 01 48 8b  $.H.L$HH..H.Y.H.
	0x0100 44 24 50 48 8b 4c 24 30 48 8b 54 24 38 48 83 fb  D$PH.L$0H.T$8H..
	0x0110 15 7c bb 0f 57 c0 0f 11 44 24 68 48 8d 05 00 00  .|..W...D$hH....
	0x0120 00 00 48 89 04 24 48 89 54 24 08 e8 00 00 00 00  ..H..$H.T$......
	0x0130 48 8b 44 24 18 48 8b 4c 24 10 48 89 4c 24 68 48  H.D$.H.L$.H.L$hH
	0x0140 89 44 24 70 48 8d 44 24 68 48 89 04 24 48 c7 44  .D$pH.D$hH..$H.D
	0x0150 24 08 01 00 00 00 48 c7 44 24 10 01 00 00 00 e8  $.....H.D$......
	0x0160 00 00 00 00 48 8d 84 24 98 00 00 00 48 89 04 24  ....H..$....H..$
	0x0170 e8 00 00 00 00 48 8b 44 24 30 48 8d 48 01 48 8b  .....H.D$0H.H.H.
	0x0180 44 24 50 48 8b 94 24 98 00 00 00 48 85 d2 74 21  D$PH..$....H..t!
	0x0190 48 89 4c 24 30 48 8b 12 48 83 fa 05 0f 85 71 ff  H.L$0H..H.....q.
	0x01a0 ff ff 48 89 54 24 38 bb 0c 00 00 00 e9 5c ff ff  ..H.T$8......\..
	0x01b0 ff 0f 57 c0 0f 11 44 24 58 48 8d 05 00 00 00 00  ..W...D$XH......
	0x01c0 48 89 04 24 48 89 4c 24 08 e8 00 00 00 00 48 8b  H..$H.L$......H.
	0x01d0 44 24 18 48 8b 4c 24 10 48 89 4c 24 58 48 89 44  D$.H.L$.H.L$XH.D
	0x01e0 24 60 48 8d 44 24 58 48 89 04 24 48 c7 44 24 08  $`H.D$XH..$H.D$.
	0x01f0 01 00 00 00 48 c7 44 24 10 01 00 00 00 e8 00 00  ....H.D$........
	0x0200 00 00 48 8b 44 24 50 48 85 c0 0f 84 8b 00 00 00  ..H.D$PH........
	0x0210 48 8b 08 0f 57 c0 0f 11 44 24 78 0f 11 84 24 88  H...W...D$x...$.
	0x0220 00 00 00 48 8d 05 00 00 00 00 48 89 04 24 48 89  ...H......H..$H.
	0x0230 4c 24 08 e8 00 00 00 00 48 8b 44 24 10 48 8b 4c  L$......H.D$.H.L
	0x0240 24 18 48 89 44 24 78 48 89 8c 24 80 00 00 00 48  $.H.D$xH..$....H
	0x0250 8d 05 00 00 00 00 48 89 84 24 88 00 00 00 48 8b  ......H..$....H.
	0x0260 44 24 50 48 89 84 24 90 00 00 00 48 8d 44 24 78  D$PH..$....H.D$x
	0x0270 48 89 04 24 48 c7 44 24 08 02 00 00 00 48 c7 44  H..$H.D$.....H.D
	0x0280 24 10 02 00 00 00 e8 00 00 00 00 48 8b ac 24 f8  $..........H..$.
	0x0290 00 00 00 48 81 c4 00 01 00 00 c3 31 c9 e9 71 ff  ...H.......1..q.
	0x02a0 ff ff e8 00 00 00 00 e9 54 fd ff ff              ........T...
	rel 5+4 t=16 TLS+0
	rel 48+4 t=8 runtime.makemap_small+0
	rel 76+4 t=15 type.map[int]int+0
	rel 95+4 t=8 runtime.mapassign_fast64+0
	rel 152+4 t=8 runtime.duffzero+273
	rel 163+4 t=15 type.map[int]int+0
	rel 190+4 t=8 runtime.mapiterinit+0
	rel 214+4 t=15 type.map[int]int+0
	rel 233+4 t=8 runtime.mapassign_fast64+0
	rel 286+4 t=15 type.int+0
	rel 300+4 t=8 runtime.convT2E64+0
	rel 352+4 t=8 fmt.Println+0
	rel 369+4 t=8 runtime.mapiternext+0
	rel 444+4 t=15 type.int+0
	rel 458+4 t=8 runtime.convT2E64+0
	rel 510+4 t=8 fmt.Println+0
	rel 550+4 t=15 type.int+0
	rel 564+4 t=8 runtime.convT2E64+0
	rel 594+4 t=15 type.map[int]int+0
	rel 647+4 t=8 fmt.Println+0
	rel 675+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=92 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	85
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	PCDATA	$2, $0
	0x001b 00027 (<autogenerated>:1)	PCDATA	$0, $0
	0x001b 00027 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (<autogenerated>:1)	CMPB	AL, $1
	0x0025 00037 (<autogenerated>:1)	JLS	48
	0x0027 00039 (<autogenerated>:1)	PCDATA	$2, $-2
	0x0027 00039 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0027 00039 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002b 00043 (<autogenerated>:1)	ADDQ	$8, SP
	0x002f 00047 (<autogenerated>:1)	RET
	0x0030 00048 (<autogenerated>:1)	JNE	57
	0x0032 00050 (<autogenerated>:1)	PCDATA	$2, $0
	0x0032 00050 (<autogenerated>:1)	PCDATA	$0, $0
	0x0032 00050 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x0037 00055 (<autogenerated>:1)	UNDEF
	0x0039 00057 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x0040 00064 (<autogenerated>:1)	CALL	fmt.init(SB)
	0x0045 00069 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x004c 00076 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0050 00080 (<autogenerated>:1)	ADDQ	$8, SP
	0x0054 00084 (<autogenerated>:1)	RET
	0x0055 00085 (<autogenerated>:1)	NOP
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0055 00085 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0055 00085 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x005a 00090 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 46 48  dH..%....H;a.vFH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 80 f8 01 76 09 48 8b 2c 24 48 83 c4 08 c3  .....v.H.,$H....
	0x0030 75 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 01  u...............
	0x0040 e8 00 00 00 00 c6 05 00 00 00 00 02 48 8b 2c 24  ............H.,$
	0x0050 48 83 c4 08 c3 e8 00 00 00 00 eb a4              H...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 51+4 t=8 runtime.throwinit+0
	rel 59+4 t=15 "".initdone·+-1
	rel 65+4 t=8 fmt.init+0
	rel 71+4 t=15 "".initdone·+-1
	rel 86+4 t=8 runtime.morestack_noctxt+0
type..hash.[2]interface {} STEXT dupok size=110 args=0x18 locals=0x28
	0x0000 00000 (<autogenerated>:1)	TEXT	type..hash.[2]interface {}(SB), DUPOK, $40-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	103
	0x000f 00015 (<autogenerated>:1)	SUBQ	$40, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, 32(SP)
	0x0018 00024 (<autogenerated>:1)	LEAQ	32(SP), BP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$3, gclocals·ee104e299ed2e4539b82c61c5a4b843d(SB)
	0x001d 00029 (<autogenerated>:1)	PCDATA	$2, $0
	0x001d 00029 (<autogenerated>:1)	PCDATA	$0, $0
	0x001d 00029 (<autogenerated>:1)	XORL	AX, AX
	0x001f 00031 (<autogenerated>:1)	MOVQ	"".h+56(SP), CX
	0x0024 00036 (<autogenerated>:1)	JMP	82
	0x0026 00038 (<autogenerated>:1)	MOVQ	AX, "".i+24(SP)
	0x002b 00043 (<autogenerated>:1)	SHLQ	$4, AX
	0x002f 00047 (<autogenerated>:1)	PCDATA	$2, $1
	0x002f 00047 (<autogenerated>:1)	MOVQ	"".p+48(SP), BX
	0x0034 00052 (<autogenerated>:1)	PCDATA	$2, $2
	0x0034 00052 (<autogenerated>:1)	ADDQ	BX, AX
	0x0037 00055 (<autogenerated>:1)	PCDATA	$2, $0
	0x0037 00055 (<autogenerated>:1)	MOVQ	AX, (SP)
	0x003b 00059 (<autogenerated>:1)	MOVQ	CX, 8(SP)
	0x0040 00064 (<autogenerated>:1)	CALL	runtime.nilinterhash(SB)
	0x0045 00069 (<autogenerated>:1)	MOVQ	16(SP), CX
	0x004a 00074 (<autogenerated>:1)	MOVQ	"".i+24(SP), AX
	0x004f 00079 (<autogenerated>:1)	INCQ	AX
	0x0052 00082 (<autogenerated>:1)	CMPQ	AX, $2
	0x0056 00086 (<autogenerated>:1)	JLT	38
	0x0058 00088 (<autogenerated>:1)	PCDATA	$0, $1
	0x0058 00088 (<autogenerated>:1)	MOVQ	CX, "".~r2+64(SP)
	0x005d 00093 (<autogenerated>:1)	MOVQ	32(SP), BP
	0x0062 00098 (<autogenerated>:1)	ADDQ	$40, SP
	0x0066 00102 (<autogenerated>:1)	RET
	0x0067 00103 (<autogenerated>:1)	NOP
	0x0067 00103 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0067 00103 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0067 00103 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x006c 00108 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 58 48  dH..%....H;a.vXH
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 31 c0 48  ..(H.l$ H.l$ 1.H
	0x0020 8b 4c 24 38 eb 2c 48 89 44 24 18 48 c1 e0 04 48  .L$8.,H.D$.H...H
	0x0030 8b 5c 24 30 48 01 d8 48 89 04 24 48 89 4c 24 08  .\$0H..H..$H.L$.
	0x0040 e8 00 00 00 00 48 8b 4c 24 10 48 8b 44 24 18 48  .....H.L$.H.D$.H
	0x0050 ff c0 48 83 f8 02 7c ce 48 89 4c 24 40 48 8b 6c  ..H...|.H.L$@H.l
	0x0060 24 20 48 83 c4 28 c3 e8 00 00 00 00 eb 92        $ H..(........
	rel 5+4 t=16 TLS+0
	rel 65+4 t=8 runtime.nilinterhash+0
	rel 104+4 t=8 runtime.morestack_noctxt+0
type..eq.[2]interface {} STEXT dupok size=182 args=0x18 locals=0x30
	0x0000 00000 (<autogenerated>:1)	TEXT	type..eq.[2]interface {}(SB), DUPOK, $48-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	172
	0x0013 00019 (<autogenerated>:1)	SUBQ	$48, SP
	0x0017 00023 (<autogenerated>:1)	MOVQ	BP, 40(SP)
	0x001c 00028 (<autogenerated>:1)	LEAQ	40(SP), BP
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$0, gclocals·dc9b0298814590ca3ffc3a889546fc8b(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$3, gclocals·a1bdf42ea3370bf425f59e11a41daee2(SB)
	0x0021 00033 (<autogenerated>:1)	PCDATA	$2, $1
	0x0021 00033 (<autogenerated>:1)	PCDATA	$0, $0
	0x0021 00033 (<autogenerated>:1)	MOVQ	"".p+56(SP), AX
	0x0026 00038 (<autogenerated>:1)	PCDATA	$2, $2
	0x0026 00038 (<autogenerated>:1)	MOVQ	"".q+64(SP), CX
	0x002b 00043 (<autogenerated>:1)	XORL	DX, DX
	0x002d 00045 (<autogenerated>:1)	JMP	72
	0x002f 00047 (<autogenerated>:1)	PCDATA	$2, $0
	0x002f 00047 (<autogenerated>:1)	MOVQ	""..autotmp_8+32(SP), BX
	0x0034 00052 (<autogenerated>:1)	LEAQ	1(BX), DX
	0x0038 00056 (<autogenerated>:1)	PCDATA	$2, $3
	0x0038 00056 (<autogenerated>:1)	MOVQ	"".p+56(SP), BX
	0x003d 00061 (<autogenerated>:1)	PCDATA	$2, $4
	0x003d 00061 (<autogenerated>:1)	MOVQ	"".q+64(SP), SI
	0x0042 00066 (<autogenerated>:1)	PCDATA	$2, $5
	0x0042 00066 (<autogenerated>:1)	MOVQ	BX, AX
	0x0045 00069 (<autogenerated>:1)	PCDATA	$2, $2
	0x0045 00069 (<autogenerated>:1)	MOVQ	SI, CX
	0x0048 00072 (<autogenerated>:1)	CMPQ	DX, $2
	0x004c 00076 (<autogenerated>:1)	JGE	157
	0x004e 00078 (<autogenerated>:1)	MOVQ	DX, BX
	0x0051 00081 (<autogenerated>:1)	SHLQ	$4, DX
	0x0055 00085 (<autogenerated>:1)	PCDATA	$2, $6
	0x0055 00085 (<autogenerated>:1)	MOVQ	8(DX)(AX*1), SI
	0x005a 00090 (<autogenerated>:1)	PCDATA	$2, $7
	0x005a 00090 (<autogenerated>:1)	MOVQ	(DX)(AX*1), DI
	0x005e 00094 (<autogenerated>:1)	PCDATA	$2, $8
	0x005e 00094 (<autogenerated>:1)	MOVQ	8(DX)(CX*1), R8
	0x0063 00099 (<autogenerated>:1)	PCDATA	$2, $9
	0x0063 00099 (<autogenerated>:1)	MOVQ	(DX)(CX*1), DX
	0x0067 00103 (<autogenerated>:1)	CMPQ	DI, DX
	0x006a 00106 (<autogenerated>:1)	JNE	142
	0x006c 00108 (<autogenerated>:1)	MOVQ	BX, ""..autotmp_8+32(SP)
	0x0071 00113 (<autogenerated>:1)	MOVQ	DI, (SP)
	0x0075 00117 (<autogenerated>:1)	PCDATA	$2, $10
	0x0075 00117 (<autogenerated>:1)	MOVQ	SI, 8(SP)
	0x007a 00122 (<autogenerated>:1)	PCDATA	$2, $0
	0x007a 00122 (<autogenerated>:1)	MOVQ	R8, 16(SP)
	0x007f 00127 (<autogenerated>:1)	CALL	runtime.efaceeq(SB)
	0x0084 00132 (<autogenerated>:1)	PCDATA	$2, $1
	0x0084 00132 (<autogenerated>:1)	LEAQ	24(SP), AX
	0x0089 00137 (<autogenerated>:1)	PCDATA	$2, $0
	0x0089 00137 (<autogenerated>:1)	CMPB	(AX), $0
	0x008c 00140 (<autogenerated>:1)	JNE	47
	0x008e 00142 (<autogenerated>:1)	PCDATA	$0, $1
	0x008e 00142 (<autogenerated>:1)	MOVB	$0, "".~r2+72(SP)
	0x0093 00147 (<autogenerated>:1)	MOVQ	40(SP), BP
	0x0098 00152 (<autogenerated>:1)	ADDQ	$48, SP
	0x009c 00156 (<autogenerated>:1)	RET
	0x009d 00157 (<autogenerated>:1)	MOVB	$1, "".~r2+72(SP)
	0x00a2 00162 (<autogenerated>:1)	MOVQ	40(SP), BP
	0x00a7 00167 (<autogenerated>:1)	ADDQ	$48, SP
	0x00ab 00171 (<autogenerated>:1)	RET
	0x00ac 00172 (<autogenerated>:1)	NOP
	0x00ac 00172 (<autogenerated>:1)	PCDATA	$0, $-1
	0x00ac 00172 (<autogenerated>:1)	PCDATA	$2, $-1
	0x00ac 00172 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x00b1 00177 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 99  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 30 48 89 6c 24 28 48 8d 6c 24  ...H..0H.l$(H.l$
	0x0020 28 48 8b 44 24 38 48 8b 4c 24 40 31 d2 eb 19 48  (H.D$8H.L$@1...H
	0x0030 8b 5c 24 20 48 8d 53 01 48 8b 5c 24 38 48 8b 74  .\$ H.S.H.\$8H.t
	0x0040 24 40 48 89 d8 48 89 f1 48 83 fa 02 7d 4f 48 89  $@H..H..H...}OH.
	0x0050 d3 48 c1 e2 04 48 8b 74 02 08 48 8b 3c 02 4c 8b  .H...H.t..H.<.L.
	0x0060 44 0a 08 48 8b 14 0a 48 39 d7 75 22 48 89 5c 24  D..H...H9.u"H.\$
	0x0070 20 48 89 3c 24 48 89 74 24 08 4c 89 44 24 10 e8   H.<$H.t$.L.D$..
	0x0080 00 00 00 00 48 8d 44 24 18 80 38 00 75 a1 c6 44  ....H.D$..8.u..D
	0x0090 24 48 00 48 8b 6c 24 28 48 83 c4 30 c3 c6 44 24  $H.H.l$(H..0..D$
	0x00a0 48 01 48 8b 6c 24 28 48 83 c4 30 c3 e8 00 00 00  H.H.l$(H..0.....
	0x00b0 00 e9 4a ff ff ff                                ..J...
	rel 5+4 t=16 TLS+0
	rel 128+4 t=8 runtime.efaceeq+0
	rel 173+4 t=8 runtime.morestack_noctxt+0
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=33
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+48
	rel 27+4 t=29 gofile../mnt/hgfs/study/golang/src/readruntime/practice/map/delete.go+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 15 00                             .......
go.loc."".demo2 SDWARFLOC size=360
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 38 00 00 00 00 00 00 00 63 00 00 00 00 00 00 00  8.......c.......
	0x0020 01 00 50 63 00 00 00 00 00 00 00 ac 02 00 00 00  ..Pc............
	0x0030 00 00 00 03 00 91 c8 7e 00 00 00 00 00 00 00 00  .......~........
	0x0040 00 00 00 00 00 00 00 00 ff ff ff ff ff ff ff ff  ................
	0x0050 00 00 00 00 00 00 00 00 da 00 00 00 00 00 00 00  ................
	0x0060 8b 01 00 00 00 00 00 00 03 00 91 a8 7e 8b 01 00  ............~...
	0x0070 00 00 00 00 00 ce 01 00 00 00 00 00 00 01 00 52  ...............R
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x00a0 49 00 00 00 00 00 00 00 7c 00 00 00 00 00 00 00  I.......|.......
	0x00b0 03 00 91 b8 7e 7c 00 00 00 00 00 00 00 a7 00 00  ....~|..........
	0x00c0 00 00 00 00 00 01 00 52 00 00 00 00 00 00 00 00  .......R........
	0x00d0 00 00 00 00 00 00 00 00 ff ff ff ff ff ff ff ff  ................
	0x00e0 00 00 00 00 00 00 00 00 ed 00 00 00 00 00 00 00  ................
	0x00f0 30 01 00 00 00 00 00 00 03 00 91 b0 7e 98 01 00  0...........~...
	0x0100 00 00 00 00 00 ac 02 00 00 00 00 00 00 01 00 51  ...............Q
	0x0110 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0120 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0130 d3 00 00 00 00 00 00 00 11 01 00 00 00 00 00 00  ................
	0x0140 03 00 91 c0 7e 11 01 00 00 00 00 00 00 ac 02 00  ....~...........
	0x0150 00 00 00 00 00 01 00 53 00 00 00 00 00 00 00 00  .......S........
	0x0160 00 00 00 00 00 00 00 00                          ........
	rel 8+8 t=1 "".demo2+0
	rel 80+8 t=1 "".demo2+0
	rel 152+8 t=1 "".demo2+0
	rel 224+8 t=1 "".demo2+0
	rel 296+8 t=1 "".demo2+0
go.info."".demo2 SDWARFINFO size=139
	0x0000 02 22 22 2e 64 65 6d 6f 32 00 00 00 00 00 00 00  ."".demo2.......
	0x0010 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00  ................
	0x0020 01 0a 64 6d 00 2c 00 00 00 00 00 00 00 00 0a 6e  ..dm.,.........n
	0x0030 75 6d 00 32 00 00 00 00 00 00 00 00 14 00 00 00  um.2............
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 0a 69 00  ..............i.
	0x0050 2d 00 00 00 00 00 00 00 00 00 14 00 00 00 00 00  -...............
	0x0060 00 00 00 00 00 00 00 00 00 00 00 0a 6b 00 33 00  ............k.3.
	0x0070 00 00 00 00 00 00 00 13 00 00 00 00 0a 69 00 35  .............i.5
	0x0080 00 00 00 00 00 00 00 00 00 00 00                 ...........
	rel 10+8 t=1 "".demo2+0
	rel 18+8 t=1 "".demo2+684
	rel 28+4 t=29 gofile../mnt/hgfs/study/golang/src/readruntime/practice/map/delete.go+0
	rel 38+4 t=28 go.info.map[int]int+0
	rel 42+4 t=28 go.loc."".demo2+0
	rel 52+4 t=28 go.info.int+0
	rel 56+4 t=28 go.loc."".demo2+72
	rel 61+8 t=1 "".demo2+66
	rel 69+8 t=1 "".demo2+126
	rel 81+4 t=28 go.info.int+0
	rel 85+4 t=28 go.loc."".demo2+144
	rel 91+8 t=1 "".demo2+126
	rel 99+8 t=1 "".demo2+433
	rel 111+4 t=28 go.info.int+0
	rel 115+4 t=28 go.loc."".demo2+216
	rel 120+4 t=28 go.range."".demo2+0
	rel 128+4 t=28 go.info.int+0
	rel 132+4 t=28 go.loc."".demo2+288
go.range."".demo2 SDWARFRANGE size=80
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 ce 00 00 00 00 00 00 00 03 01 00 00 00 00 00 00  ................
	0x0020 0d 01 00 00 00 00 00 00 13 01 00 00 00 00 00 00  ................
	0x0030 ac 01 00 00 00 00 00 00 b1 01 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 8+8 t=1 "".demo2+0
go.isstmt."".demo2 SDWARFMISC size=0
	0x0000 04 18 04 17 03 05 01 0e 02 02 01 05 02 07 01 20  ............... 
	0x0010 02 03 01 05 02 04 01 02 02 08 01 4d 02 07 01 20  ...........M... 
	0x0020 02 04 01 0f 02 04 01 02 02 03 01 4e 02 08 01 09  ...........N....
	0x0030 02 05 01 09 02 08 01 0d 02 04 01 15 02 03 01 4e  ...............N
	0x0040 02 05 01 0c 02 03 01 75 02 10 01 07 02 0a 00     .......u.......
go.loc."".init SDWARFLOC size=0
go.info."".init SDWARFINFO size=33
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+92
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
go.isstmt."".init SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 07 01 05 02 09 01 07 02 09 01 15  ................
	0x0010 02 07 00                                         ...
"".initdone· SNOPTRBSS size=1
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 00 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
go.loc.type..hash.[2]interface {} SDWARFLOC dupok size=173
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 2b 00 00 00 00 00 00 00 56 00 00 00 00 00 00 00  +.......V.......
	0x0020 02 00 91 68 56 00 00 00 00 00 00 00 6e 00 00 00  ...hV.......n...
	0x0030 00 00 00 00 01 00 50 00 00 00 00 00 00 00 00 00  ......P.........
	0x0040 00 00 00 00 00 00 00 ff ff ff ff ff ff ff ff 00  ................
	0x0050 00 00 00 00 00 00 00 1f 00 00 00 00 00 00 00 6e  ...............n
	0x0060 00 00 00 00 00 00 00 01 00 9c 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 ff ff ff ff ff ff  ................
	0x0080 ff ff 00 00 00 00 00 00 00 00 4a 00 00 00 00 00  ..........J.....
	0x0090 00 00 6e 00 00 00 00 00 00 00 01 00 52 00 00 00  ..n.........R...
	0x00a0 00 00 00 00 00 00 00 00 00 00 00 00 00           .............
	rel 8+8 t=1 type..hash.[2]interface {}+0
	rel 79+8 t=1 type..hash.[2]interface {}+0
	rel 130+8 t=1 type..hash.[2]interface {}+0
go.info.type..hash.[2]interface {} SDWARFINFO dupok size=102
	0x0000 02 74 79 70 65 2e 2e 68 61 73 68 2e 5b 32 5d 69  .type..hash.[2]i
	0x0010 6e 74 65 72 66 61 63 65 20 7b 7d 00 00 00 00 00  nterface {}.....
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00  ................
	0x0030 00 00 01 0a 69 00 01 00 00 00 00 00 00 00 00 0f  ....i...........
	0x0040 70 00 00 01 00 00 00 00 00 00 00 00 0f 68 00 00  p............h..
	0x0050 01 00 00 00 00 00 00 00 00 0e 7e 72 32 00 01 01  ..........~r2...
	0x0060 00 00 00 00 00 00                                ......
	rel 28+8 t=1 type..hash.[2]interface {}+0
	rel 36+8 t=1 type..hash.[2]interface {}+110
	rel 46+4 t=29 gofile..<autogenerated>+0
	rel 55+4 t=28 go.info.int+0
	rel 59+4 t=28 go.loc.type..hash.[2]interface {}+0
	rel 68+4 t=28 go.info.*[2]interface {}+0
	rel 72+4 t=28 go.loc.type..hash.[2]interface {}+71
	rel 81+4 t=28 go.info.uintptr+0
	rel 85+4 t=28 go.loc.type..hash.[2]interface {}+122
	rel 96+4 t=28 go.info.uintptr+0
go.range.type..hash.[2]interface {} SDWARFRANGE dupok size=0
go.isstmt.type..hash.[2]interface {} SDWARFMISC dupok size=0
	0x0000 04 0f 04 0e 03 02 01 33 02 04 01 02 02 05 01 0a  .......3........
	0x0010 02 07 00                                         ...
go.loc.type..eq.[2]interface {} SDWARFLOC dupok size=154
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 4c 00 00 00 00 00 00 00 55 00 00 00 00 00 00 00  L.......U.......
	0x0020 01 00 51 00 00 00 00 00 00 00 00 00 00 00 00 00  ..Q.............
	0x0030 00 00 00 ff ff ff ff ff ff ff ff 00 00 00 00 00  ................
	0x0040 00 00 00 26 00 00 00 00 00 00 00 b6 00 00 00 00  ...&............
	0x0050 00 00 00 01 00 9c 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 ff ff ff ff ff ff ff ff 00 00  ................
	0x0070 00 00 00 00 00 00 26 00 00 00 00 00 00 00 b6 00  ......&.........
	0x0080 00 00 00 00 00 00 02 00 91 08 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 00 00                    ..........
	rel 8+8 t=1 type..eq.[2]interface {}+0
	rel 59+8 t=1 type..eq.[2]interface {}+0
	rel 110+8 t=1 type..eq.[2]interface {}+0
go.info.type..eq.[2]interface {} SDWARFINFO dupok size=100
	0x0000 02 74 79 70 65 2e 2e 65 71 2e 5b 32 5d 69 6e 74  .type..eq.[2]int
	0x0010 65 72 66 61 63 65 20 7b 7d 00 00 00 00 00 00 00  erface {}.......
	0x0020 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00  ................
	0x0030 01 0a 69 00 01 00 00 00 00 00 00 00 00 0f 70 00  ..i...........p.
	0x0040 00 01 00 00 00 00 00 00 00 00 0f 71 00 00 01 00  ...........q....
	0x0050 00 00 00 00 00 00 00 0e 7e 72 32 00 01 01 00 00  ........~r2.....
	0x0060 00 00 00 00                                      ....
	rel 26+8 t=1 type..eq.[2]interface {}+0
	rel 34+8 t=1 type..eq.[2]interface {}+182
	rel 44+4 t=29 gofile..<autogenerated>+0
	rel 53+4 t=28 go.info.int+0
	rel 57+4 t=28 go.loc.type..eq.[2]interface {}+0
	rel 66+4 t=28 go.info.*[2]interface {}+0
	rel 70+4 t=28 go.loc.type..eq.[2]interface {}+51
	rel 79+4 t=28 go.info.*[2]interface {}+0
	rel 83+4 t=28 go.loc.type..eq.[2]interface {}+102
	rel 94+4 t=28 go.info.bool+0
go.range.type..eq.[2]interface {} SDWARFRANGE dupok size=0
go.isstmt.type..eq.[2]interface {} SDWARFMISC dupok size=0
	0x0000 04 13 04 0e 03 05 01 22 02 04 01 42 02 05 01 0a  ......."...B....
	0x0010 02 05 01 0a 02 0a 00                             .......
type..hashfunc.[2]interface {} SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..hash.[2]interface {}+0
type..eqfunc.[2]interface {} SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq.[2]interface {}+0
type..alg.[2]interface {} SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc.[2]interface {}+0
	rel 8+8 t=1 type..eqfunc.[2]interface {}+0
type..namedata.*[2]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 32 5d 69 6e 74 65 72 66 61 63 65  ...*[2]interface
	0x0010 20 7b 7d                                          {}
type.*[2]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 be 73 2d 71 00 08 08 36 00 00 00 00 00 00 00 00  .s-q...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[2]interface {}-+0
	rel 48+8 t=1 type.[2]interface {}+0
runtime.gcbits.0a SRODATA dupok size=1
	0x0000 0a                                               .
type.[2]interface {} SRODATA dupok size=72
	0x0000 20 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00   ....... .......
	0x0010 2c 59 a4 f1 02 08 08 11 00 00 00 00 00 00 00 00  ,Y..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..alg.[2]interface {}+0
	rel 32+8 t=1 runtime.gcbits.0a+0
	rel 40+4 t=5 type..namedata.*[2]interface {}-+0
	rel 44+4 t=6 type.*[2]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..namedata.*[]uint8- SRODATA dupok size=11
	0x0000 00 00 08 2a 5b 5d 75 69 6e 74 38                 ...*[]uint8
type.*[]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a5 8e d0 69 00 08 08 36 00 00 00 00 00 00 00 00  ...i...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 48+8 t=1 type.[]uint8+0
type.[]uint8 SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 df 7e 2e 38 02 08 08 17 00 00 00 00 00 00 00 00  .~.8............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 44+4 t=6 type.*[]uint8+0
	rel 48+8 t=1 type.uint8+0
type..namedata.*[8]uint8- SRODATA dupok size=12
	0x0000 00 00 09 2a 5b 38 5d 75 69 6e 74 38              ...*[8]uint8
type.*[8]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a9 89 a5 7a 00 08 08 36 00 00 00 00 00 00 00 00  ...z...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[8]uint8-+0
	rel 48+8 t=1 type.[8]uint8+0
runtime.gcbits. SRODATA dupok size=0
type.[8]uint8 SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 3e f9 30 b4 02 01 01 91 00 00 00 00 00 00 00 00  >.0.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 08 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[8]uint8-+0
	rel 44+4 t=6 type.*[8]uint8+0
	rel 48+8 t=1 type.uint8+0
	rel 56+8 t=1 type.[]uint8+0
type..namedata.*[]int- SRODATA dupok size=9
	0x0000 00 00 06 2a 5b 5d 69 6e 74                       ...*[]int
type.*[]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 1b 31 52 88 00 08 08 36 00 00 00 00 00 00 00 00  .1R....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int-+0
	rel 48+8 t=1 type.[]int+0
type.[]int SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8e 66 f9 1b 02 08 08 17 00 00 00 00 00 00 00 00  .f..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int-+0
	rel 44+4 t=6 type.*[]int+0
	rel 48+8 t=1 type.int+0
type..namedata.*[8]int- SRODATA dupok size=10
	0x0000 00 00 07 2a 5b 38 5d 69 6e 74                    ...*[8]int
type.*[8]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 3f a8 3b 00 08 08 36 00 00 00 00 00 00 00 00  .?.;...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[8]int-+0
	rel 48+8 t=1 type.noalg.[8]int+0
type.noalg.[8]int SRODATA dupok size=72
	0x0000 40 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  @...............
	0x0010 96 99 d5 05 02 08 08 91 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 08 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[8]int-+0
	rel 44+4 t=6 type.*[8]int+0
	rel 48+8 t=1 type.int+0
	rel 56+8 t=1 type.[]int+0
type..namedata.*map.bucket[int]int- SRODATA dupok size=22
	0x0000 00 00 13 2a 6d 61 70 2e 62 75 63 6b 65 74 5b 69  ...*map.bucket[i
	0x0010 6e 74 5d 69 6e 74                                nt]int
type.*map.bucket[int]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 2a ff 7c 44 00 08 08 36 00 00 00 00 00 00 00 00  *.|D...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map.bucket[int]int-+0
	rel 48+8 t=1 type.noalg.map.bucket[int]int+0
type..importpath.. SRODATA dupok size=3
	0x0000 00 00 00                                         ...
type..namedata.topbits- SRODATA dupok size=10
	0x0000 00 00 07 74 6f 70 62 69 74 73                    ...topbits
type..namedata.keys- SRODATA dupok size=7
	0x0000 00 00 04 6b 65 79 73                             ...keys
type..namedata.values- SRODATA dupok size=9
	0x0000 00 00 06 76 61 6c 75 65 73                       ...values
type..namedata.overflow- SRODATA dupok size=11
	0x0000 00 00 08 6f 76 65 72 66 6c 6f 77                 ...overflow
type.noalg.map.bucket[int]int SRODATA dupok size=176
	0x0000 90 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 b3 28 ed bb 02 08 08 99 00 00 00 00 00 00 00 00  .(..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 04 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 90 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00a0 00 00 00 00 00 00 00 00 10 01 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*map.bucket[int]int-+0
	rel 44+4 t=6 type.*map.bucket[int]int+0
	rel 48+8 t=1 type..importpath..+0
	rel 56+8 t=1 type.noalg.map.bucket[int]int+80
	rel 80+8 t=1 type..namedata.topbits-+0
	rel 88+8 t=1 type.[8]uint8+0
	rel 104+8 t=1 type..namedata.keys-+0
	rel 112+8 t=1 type.noalg.[8]int+0
	rel 128+8 t=1 type..namedata.values-+0
	rel 136+8 t=1 type.noalg.[8]int+0
	rel 152+8 t=1 type..namedata.overflow-+0
	rel 160+8 t=1 type.uintptr+0
type..namedata.*map[int]int- SRODATA dupok size=15
	0x0000 00 00 0c 2a 6d 61 70 5b 69 6e 74 5d 69 6e 74     ...*map[int]int
type.*map[int]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ab c4 20 42 00 08 08 36 00 00 00 00 00 00 00 00  .. B...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map[int]int-+0
	rel 48+8 t=1 type.map[int]int+0
type.map[int]int SRODATA dupok size=80
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 50 1b 58 23 02 08 08 35 00 00 00 00 00 00 00 00  P.X#...5........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 08 00 08 00 90 00 01 00  ................
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map[int]int-+0
	rel 44+4 t=6 type.*map[int]int+0
	rel 48+8 t=1 type.int+0
	rel 56+8 t=1 type.int+0
	rel 64+8 t=1 type.noalg.map.bucket[int]int+0
runtime.gcbits.2c SRODATA dupok size=1
	0x0000 2c                                               ,
type..namedata.*map.hdr[int]int- SRODATA dupok size=19
	0x0000 00 00 10 2a 6d 61 70 2e 68 64 72 5b 69 6e 74 5d  ...*map.hdr[int]
	0x0010 69 6e 74                                         int
type..namedata.count- SRODATA dupok size=8
	0x0000 00 00 05 63 6f 75 6e 74                          ...count
type..namedata.flags- SRODATA dupok size=8
	0x0000 00 00 05 66 6c 61 67 73                          ...flags
type..namedata.B. SRODATA dupok size=4
	0x0000 01 00 01 42                                      ...B
type..namedata.noverflow- SRODATA dupok size=12
	0x0000 00 00 09 6e 6f 76 65 72 66 6c 6f 77              ...noverflow
type..namedata.hash0- SRODATA dupok size=8
	0x0000 00 00 05 68 61 73 68 30                          ...hash0
type..namedata.buckets- SRODATA dupok size=10
	0x0000 00 00 07 62 75 63 6b 65 74 73                    ...buckets
type..namedata.oldbuckets- SRODATA dupok size=13
	0x0000 00 00 0a 6f 6c 64 62 75 63 6b 65 74 73           ...oldbuckets
type..namedata.nevacuate- SRODATA dupok size=12
	0x0000 00 00 09 6e 65 76 61 63 75 61 74 65              ...nevacuate
type..namedata.extra- SRODATA dupok size=8
	0x0000 00 00 05 65 78 74 72 61                          ...extra
type.noalg.map.hdr[int]int SRODATA dupok size=296
	0x0000 30 00 00 00 00 00 00 00 30 00 00 00 00 00 00 00  0.......0.......
	0x0010 3c 29 b0 35 02 08 08 19 00 00 00 00 00 00 00 00  <).5............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 09 00 00 00 00 00 00 00 09 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 12 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00a0 00 00 00 00 00 00 00 00 14 00 00 00 00 00 00 00  ................
	0x00b0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00c0 18 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00d0 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	0x00e0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00f0 30 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  0...............
	0x0100 00 00 00 00 00 00 00 00 40 00 00 00 00 00 00 00  ........@.......
	0x0110 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0120 50 00 00 00 00 00 00 00                          P.......
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.2c+0
	rel 40+4 t=5 type..namedata.*map.hdr[int]int-+0
	rel 44+4 t=6 type.*map.hdr[int]int+0
	rel 48+8 t=1 type..importpath..+0
	rel 56+8 t=1 type.noalg.map.hdr[int]int+80
	rel 80+8 t=1 type..namedata.count-+0
	rel 88+8 t=1 type.int+0
	rel 104+8 t=1 type..namedata.flags-+0
	rel 112+8 t=1 type.uint8+0
	rel 128+8 t=1 type..namedata.B.+0
	rel 136+8 t=1 type.uint8+0
	rel 152+8 t=1 type..namedata.noverflow-+0
	rel 160+8 t=1 type.uint16+0
	rel 176+8 t=1 type..namedata.hash0-+0
	rel 184+8 t=1 type.uint32+0
	rel 200+8 t=1 type..namedata.buckets-+0
	rel 208+8 t=1 type.*map.bucket[int]int+0
	rel 224+8 t=1 type..namedata.oldbuckets-+0
	rel 232+8 t=1 type.*map.bucket[int]int+0
	rel 248+8 t=1 type..namedata.nevacuate-+0
	rel 256+8 t=1 type.uintptr+0
	rel 272+8 t=1 type..namedata.extra-+0
	rel 280+8 t=1 type.unsafe.Pointer+0
type.*map.hdr[int]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8b 92 74 97 00 08 08 36 00 00 00 00 00 00 00 00  ..t....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map.hdr[int]int-+0
	rel 48+8 t=1 type.noalg.map.hdr[int]int+0
type..namedata.*map.iter[int]int- SRODATA dupok size=20
	0x0000 00 00 11 2a 6d 61 70 2e 69 74 65 72 5b 69 6e 74  ...*map.iter[int
	0x0010 5d 69 6e 74                                      ]int
type.*map.iter[int]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 39 ec 8d d2 00 08 08 36 00 00 00 00 00 00 00 00  9......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map.iter[int]int-+0
	rel 48+8 t=1 type.noalg.map.iter[int]int+0
runtime.gcbits.ff SRODATA dupok size=1
	0x0000 ff                                               .
type..namedata.key- SRODATA dupok size=6
	0x0000 00 00 03 6b 65 79                                ...key
type..namedata.val- SRODATA dupok size=6
	0x0000 00 00 03 76 61 6c                                ...val
type..namedata.t- SRODATA dupok size=4
	0x0000 00 00 01 74                                      ...t
type..namedata.h- SRODATA dupok size=4
	0x0000 00 00 01 68                                      ...h
type..namedata.bptr- SRODATA dupok size=7
	0x0000 00 00 04 62 70 74 72                             ...bptr
type..namedata.oldoverflow- SRODATA dupok size=14
	0x0000 00 00 0b 6f 6c 64 6f 76 65 72 66 6c 6f 77        ...oldoverflow
type..namedata.startBucket- SRODATA dupok size=14
	0x0000 00 00 0b 73 74 61 72 74 42 75 63 6b 65 74        ...startBucket
type..namedata.offset- SRODATA dupok size=9
	0x0000 00 00 06 6f 66 66 73 65 74                       ...offset
type..namedata.wrapped- SRODATA dupok size=10
	0x0000 00 00 07 77 72 61 70 70 65 64                    ...wrapped
type..namedata.i- SRODATA dupok size=4
	0x0000 00 00 01 69                                      ...i
type..namedata.bucket- SRODATA dupok size=9
	0x0000 00 00 06 62 75 63 6b 65 74                       ...bucket
type..namedata.checkBucket- SRODATA dupok size=14
	0x0000 00 00 0b 63 68 65 63 6b 42 75 63 6b 65 74        ...checkBucket
type.noalg.map.iter[int]int SRODATA dupok size=440
	0x0000 60 00 00 00 00 00 00 00 40 00 00 00 00 00 00 00  `.......@.......
	0x0010 85 17 e7 d0 02 08 08 19 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 0f 00 00 00 00 00 00 00 0f 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 20 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ...............
	0x00a0 00 00 00 00 00 00 00 00 30 00 00 00 00 00 00 00  ........0.......
	0x00b0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00c0 40 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  @...............
	0x00d0 00 00 00 00 00 00 00 00 50 00 00 00 00 00 00 00  ........P.......
	0x00e0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00f0 60 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  `...............
	0x0100 00 00 00 00 00 00 00 00 70 00 00 00 00 00 00 00  ........p.......
	0x0110 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0120 80 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0130 00 00 00 00 00 00 00 00 90 00 00 00 00 00 00 00  ................
	0x0140 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0150 92 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0160 00 00 00 00 00 00 00 00 94 00 00 00 00 00 00 00  ................
	0x0170 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0180 96 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0190 00 00 00 00 00 00 00 00 a0 00 00 00 00 00 00 00  ................
	0x01a0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x01b0 b0 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.ff+0
	rel 40+4 t=5 type..namedata.*map.iter[int]int-+0
	rel 44+4 t=6 type.*map.iter[int]int+0
	rel 48+8 t=1 type..importpath..+0
	rel 56+8 t=1 type.noalg.map.iter[int]int+80
	rel 80+8 t=1 type..namedata.key-+0
	rel 88+8 t=1 type.*int+0
	rel 104+8 t=1 type..namedata.val-+0
	rel 112+8 t=1 type.*int+0
	rel 128+8 t=1 type..namedata.t-+0
	rel 136+8 t=1 type.unsafe.Pointer+0
	rel 152+8 t=1 type..namedata.h-+0
	rel 160+8 t=1 type.*map.hdr[int]int+0
	rel 176+8 t=1 type..namedata.buckets-+0
	rel 184+8 t=1 type.*map.bucket[int]int+0
	rel 200+8 t=1 type..namedata.bptr-+0
	rel 208+8 t=1 type.*map.bucket[int]int+0
	rel 224+8 t=1 type..namedata.overflow-+0
	rel 232+8 t=1 type.unsafe.Pointer+0
	rel 248+8 t=1 type..namedata.oldoverflow-+0
	rel 256+8 t=1 type.unsafe.Pointer+0
	rel 272+8 t=1 type..namedata.startBucket-+0
	rel 280+8 t=1 type.uintptr+0
	rel 296+8 t=1 type..namedata.offset-+0
	rel 304+8 t=1 type.uint8+0
	rel 320+8 t=1 type..namedata.wrapped-+0
	rel 328+8 t=1 type.bool+0
	rel 344+8 t=1 type..namedata.B.+0
	rel 352+8 t=1 type.uint8+0
	rel 368+8 t=1 type..namedata.i-+0
	rel 376+8 t=1 type.uint8+0
	rel 392+8 t=1 type..namedata.bucket-+0
	rel 400+8 t=1 type.uintptr+0
	rel 416+8 t=1 type..namedata.checkBucket-+0
	rel 424+8 t=1 type.uintptr+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·f5be5308b59e045b7c5b33ee8908cfb7 SRODATA dupok size=8
	0x0000 07 00 00 00 00 00 00 00                          ........
gclocals·0970b3a362ebabf95df0bb0ec675cbde SRODATA dupok size=29
	0x0000 07 00 00 00 15 00 00 00 00 00 00 01 00 00 01 fe  ................
	0x0010 01 11 fe 01 05 00 00 41 01 00 40 01 00           .......A..@..
gclocals·0879caa565d3c697358d27e15966481d SRODATA dupok size=15
	0x0000 07 00 00 00 07 00 00 00 00 01 05 41 03 04 02     ...........A...
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·ee104e299ed2e4539b82c61c5a4b843d SRODATA dupok size=11
	0x0000 03 00 00 00 04 00 00 00 00 08 01                 ...........
gclocals·dc9b0298814590ca3ffc3a889546fc8b SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 03 00                    ..........
gclocals·a1bdf42ea3370bf425f59e11a41daee2 SRODATA dupok size=19
	0x0000 0b 00 00 00 08 00 00 00 00 01 03 08 28 21 23 22  ............(!#"
	0x0010 a2 a0 80                                         ...
