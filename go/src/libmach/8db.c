<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/libmach/8db.c</title>

  <link rel="stylesheet" type="text/css" href="../../doc/style.css">
  <script type="text/javascript" src="../../doc/godocs.js"></script>

</head>

<body>

  <script>
    // Catch 'enter' key down events and trigger the search form submission.
    function codesearchKeyDown(event) {
      if (event.which == 13) {
        var form = document.getElementById('codesearch');
        var query = document.getElementById('codesearchQuery');
        form.q.value = "lang:go package:go.googlecode.com " + query.value;
        document.getElementById('codesearch').submit();
}      return true;
}
    // Capture the submission event and construct the query parameter.
    function codeSearchSubmit() {
      var query = document.getElementById('codesearchQuery');
      var form = document.getElementById('codesearch');
      form.q.value = "lang:go package:go.googlecode.com " + query.value;
      return true;
}  </script>

<div id="topnav">
  <table summary="">
    <tr>
      <td id="headerImage">
        <a href="../../index.html"><img src="../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
      </td>
      <td>
        <div id="headerDocSetTitle">The Go Programming Language</div>
      </td>
      <td>
        <!-- <table>
          <tr>
            <td>
              <! The input box is outside of the form because we want to add
              a couple of restricts to the query before submitting. If we just
              add the restricts to the text box before submitting, then they
              appear in the box when the user presses 'back'. Thus we use a
              hidden field in the form. However, there's no way to stop the
              non-hidden text box from also submitting a value unless we move
              it outside of the form
              <input type="search" id="codesearchQuery" value="" size="30" onkeydown="return codesearchKeyDown(event);"/>
              <form method="GET" action="http://www.google.com/codesearch" id="codesearch" class="search" onsubmit="return codeSearchSubmit();" style="display:inline;">
                <input type="hidden" name="q" value=""/>
                <input type="submit" value="Code search" />
                <span style="color: red">(TODO: remove for now?)</span>
              </form>
            </td>
          </tr>
          <tr>
            <td>
              <span style="color: gray;">(e.g. &ldquo;pem&rdquo; or &ldquo;xml&rdquo;)</span>
            </td>
          </tr>
        </table> -->
      </td>
    </tr>
  </table>
</div>

<div id="linkList">
  <ul>
    <li class="navhead"><a href="../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../doc/install.html">Install Go</a></li>
    <li><a href="../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../pkg/index.html">Package documentation</a></li>
    <li><a href="../index.html">Source files</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Help</li>
    <li>#go-nuts on irc.freenode.net</li>
    <li><a href="http://groups.google.com/group/golang-nuts">Go Nuts mailing list</a></li>
    <li><a href="http://code.google.com/p/go/issues/list">Issue tracker</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Go code search</li>
    <form method="GET" action="http://golang.org/search" class="search">
    <input type="search" name="q" value="" size="25" style="width:80%; max-width:200px" />
    <input type="submit" value="Go" />
    </form>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Last update</li>
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/libmach/8db.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno libmach/8db.c
// http://code.google.com/p/inferno-os/source/browse/utils/libmach/8db.c
//
// 	Copyright © 1994-1999 Lucent Technologies Inc.
// 	Power PC support Copyright © 1995-2004 C H Forsyth (forsyth@terzarima.net).
// 	Portions Copyright © 1997-1999 Vita Nuova Limited.
// 	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com).
// 	Revisions Copyright © 2000-2004 Lucent Technologies Inc. and others.
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the &#34;Software&#34;), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

#include &lt;u.h&gt;
#include &lt;libc.h&gt;
#include &lt;bio.h&gt;
#include &lt;mach.h&gt;
#define Ureg UregAmd64
#include &lt;ureg_amd64.h&gt;
#undef Ureg
#define Ureg Ureg386
#include &lt;ureg_x86.h&gt;
#undef Ureg

typedef struct UregAmd64 UregAmd64;
typedef struct Ureg386 Ureg386;

/*
 * i386-specific debugger interface
 * also amd64 extensions
 */

static	char	*i386excep(Map*, Rgetter);

static	int	i386trace(Map*, uvlong, uvlong, uvlong, Tracer);
static	uvlong	i386frame(Map*, uvlong, uvlong, uvlong, uvlong);
static	int	i386foll(Map*, uvlong, Rgetter, uvlong*);
static	int	i386inst(Map*, uvlong, char, char*, int);
static	int	i386das(Map*, uvlong, char*, int);
static	int	i386instlen(Map*, uvlong);

static	char	STARTSYM[] =	&#34;_main&#34;;
static	char	GOSTARTSYM[] =	&#34;sys·goexit&#34;;
static	char	PROFSYM[] =	&#34;_mainp&#34;;
static	char	FRAMENAME[] =	&#34;.frame&#34;;
static	char	LESSSTACK[] = &#34;sys·lessstack&#34;;
static	char	MORESTACK[] = &#34;sys·morestack&#34;;
static char *excname[] =
{
[0]	&#34;divide error&#34;,
[1]	&#34;debug exception&#34;,
[4]	&#34;overflow&#34;,
[5]	&#34;bounds check&#34;,
[6]	&#34;invalid opcode&#34;,
[7]	&#34;math coprocessor emulation&#34;,
[8]	&#34;double fault&#34;,
[9]	&#34;math coprocessor overrun&#34;,
[10]	&#34;invalid TSS&#34;,
[11]	&#34;segment not present&#34;,
[12]	&#34;stack exception&#34;,
[13]	&#34;general protection violation&#34;,
[14]	&#34;page fault&#34;,
[16]	&#34;math coprocessor error&#34;,
[17]	&#34;alignment check&#34;,
[18]	&#34;machine check&#34;,
[19]	&#34;floating-point exception&#34;,
[24]	&#34;clock&#34;,
[25]	&#34;keyboard&#34;,
[27]	&#34;modem status&#34;,
[28]	&#34;serial line status&#34;,
[30]	&#34;floppy disk&#34;,
[36]	&#34;mouse&#34;,
[37]	&#34;math coprocessor&#34;,
[38]	&#34;hard disk&#34;,
[64]	&#34;system call&#34;,
};

Machdata i386mach =
{
	{0xCC, 0, 0, 0},	/* break point: INT 3 */
	1,			/* break point size */

	leswab,			/* convert short to local byte order */
	leswal,			/* convert int32 to local byte order */
	leswav,			/* convert vlong to local byte order */
	i386trace,		/* C traceback */
	i386frame,		/* frame finder */
	i386excep,		/* print exception */
	0,			/* breakpoint fixup */
	leieeesftos,		/* single precision float printer */
	leieeedftos,		/* double precision float printer */
	i386foll,		/* following addresses */
	i386inst,		/* print instruction */
	i386das,		/* dissembler */
	i386instlen,		/* instruction size calculation */
};

static char*
i386excep(Map *map, Rgetter rget)
{
	uint32 c;
	uvlong pc;
	static char buf[16];

	c = (*rget)(map, &#34;TRAP&#34;);
	if(c &gt; 64 || excname[c] == 0) {
		if (c == 3) {
			pc = (*rget)(map, &#34;PC&#34;);
			if (get1(map, pc, (uchar*)buf, machdata-&gt;bpsize) &gt; 0)
			if (memcmp(buf, machdata-&gt;bpinst, machdata-&gt;bpsize) == 0)
				return &#34;breakpoint&#34;;
		}
		snprint(buf, sizeof(buf), &#34;exception %ld&#34;, c);
		return buf;
	} else
		return excname[c];
}

static int
i386trace(Map *map, uvlong pc, uvlong sp, uvlong link, Tracer trace)
{
	int i;
	uvlong osp, pc1;
	Symbol s, f, s1;
	extern Mach mamd64;
	int isamd64;
	uvlong g, m, lessstack, morestack, stktop;

	isamd64 = (mach == &amp;mamd64);

	// ../pkg/runtime/runtime.h
	// G is
	//	byte* stackguard
	//	byte* stackbase (= Stktop*)
	//	Defer* defer
	//	Gobuf sched
	// TODO(rsc): Need some way to get at the g for other threads.
	// Probably need to pass it into the trace function.
	g = 0;
	if(isamd64)
		geta(map, offsetof(struct UregAmd64, r15), &amp;g);
	else {
		// TODO(rsc): How to fetch g on 386?
	}
	stktop = 0;
	if(g != 0)
		geta(map, g+1*mach-&gt;szaddr, &amp;stktop);

	lessstack = 0;
	if(lookup(0, LESSSTACK, &amp;s))
		lessstack = s.value;
	morestack = 0;
	if(lookup(0, MORESTACK, &amp;s))
		morestack = s.value;

	USED(link);
	osp = 0;
	i = 0;

	for(;;) {
		if(!findsym(pc, CTEXT, &amp;s)) {
			// check for closure return sequence
			uchar buf[8], *p;
			if(get1(map, pc, buf, 8) &lt; 0)
				break;
			// ADDQ $xxx, SP; RET
			p = buf;
			if(mach == &amp;mamd64) {
				if(p[0] != 0x48)
					break;
				p++;
			}
			if(p[0] != 0x81 || p[1] != 0xc4 || p[6] != 0xc3)
				break;
			sp += p[2] | (p[3]&lt;&lt;8) | (p[4]&lt;&lt;16) | (p[5]&lt;&lt;24);
			if(geta(map, sp, &amp;pc) &lt; 0)
				break;
			sp += mach-&gt;szaddr;
			continue;
		}

		if (osp == sp)
			break;
		osp = sp;

		if(strcmp(STARTSYM, s.name) == 0 ||
		   strcmp(GOSTARTSYM, s.name) == 0 ||
		   strcmp(PROFSYM, s.name) == 0)
			break;

		if(s.value == morestack) {
			// In the middle of morestack.
			// Caller is m-&gt;morepc.
			// Caller&#39;s caller is in m-&gt;morearg.
			// TODO(rsc): 386
			geta(map, offsetof(struct UregAmd64, r14), &amp;m);

			pc = 0;
			sp = 0;
			pc1 = 0;
			s1 = s;
			memset(&amp;s, 0, sizeof s);
			geta(map, m+1*mach-&gt;szaddr, &amp;pc1);	// m-&gt;morepc
			geta(map, m+2*mach-&gt;szaddr, &amp;sp);	// m-&gt;morebuf.sp
			geta(map, m+3*mach-&gt;szaddr, &amp;pc);	// m-&gt;morebuf.pc
			findsym(pc1, CTEXT, &amp;s);
			(*trace)(map, pc1, sp-mach-&gt;szaddr, &amp;s1);	// morestack symbol; caller&#39;s PC/SP

			// caller&#39;s caller
			s1 = s;
			findsym(pc, CTEXT, &amp;s);
			(*trace)(map, pc, sp, &amp;s1);		// morestack&#39;s caller; caller&#39;s caller&#39;s PC/SP
			continue;
		}

		if(pc == lessstack) {
			// ../pkg/runtime/runtime.h
			// Stktop is
			//	byte* stackguard
			//	byte* stackbase
			//	Gobuf gobuf
			//		byte* sp;
			//		byte* pc;
			//		G*	g;
			if(!isamd64)
				fprint(2, &#34;warning: cannot unwind stack split on 386\n&#34;);
			if(stktop == 0)
				break;
			pc = 0;
			sp = 0;
			geta(map, stktop+2*mach-&gt;szaddr, &amp;sp);
			geta(map, stktop+3*mach-&gt;szaddr, &amp;pc);
			geta(map, stktop+1*mach-&gt;szaddr, &amp;stktop);
			(*trace)(map, pc, sp, &amp;s1);
			continue;
		}

		s1 = s;
		pc1 = 0;
		if(pc != s.value) {	/* not at first instruction */
			if(findlocal(&amp;s, FRAMENAME, &amp;f) == 0)
				break;
			geta(map, sp, &amp;pc1);
			sp += f.value-mach-&gt;szaddr;
		}
		if(geta(map, sp, &amp;pc) &lt; 0)
			break;

		// If PC is not valid, assume we caught the function
		// before it moved the stack pointer down or perhaps
		// after it moved the stack pointer back up.
		// Try the PC we&#39;d have gotten without the stack
		// pointer adjustment above (pc != s.value).
		// This only matters for the first frame, and it is only
		// a heuristic, but it does help.
		if(!findsym(pc, CTEXT, &amp;s) || strcmp(s.name, &#34;etext&#34;) == 0)
			pc = pc1;

		if(pc == 0)
			break;

		if(pc != lessstack)
			(*trace)(map, pc, sp, &amp;s1);
		sp += mach-&gt;szaddr;

		if(++i &gt; 1000)
			break;
	}
	return i;
}

static uvlong
i386frame(Map *map, uvlong addr, uvlong pc, uvlong sp, uvlong link)
{
	Symbol s, f;

	USED(link);
	while (findsym(pc, CTEXT, &amp;s)) {
		if(strcmp(STARTSYM, s.name) == 0 || strcmp(PROFSYM, s.name) == 0)
			break;

		if(pc != s.value) {	/* not first instruction */
			if(findlocal(&amp;s, FRAMENAME, &amp;f) == 0)
				break;
			sp += f.value-mach-&gt;szaddr;
		}

		if (s.value == addr)
			return sp;

		if (geta(map, sp, &amp;pc) &lt; 0)
			break;
		sp += mach-&gt;szaddr;
	}
	return 0;
}

	/* I386/486 - Disassembler and related functions */

/*
 *  an instruction
 */
typedef struct Instr Instr;
struct	Instr
{
	uchar	mem[1+1+1+1+2+1+1+4+4];		/* raw instruction */
	uvlong	addr;		/* address of start of instruction */
	int	n;		/* number of bytes in instruction */
	char	*prefix;	/* instr prefix */
	char	*segment;	/* segment override */
	uchar	jumptype;	/* set to the operand type for jump/ret/call */
	uchar	amd64;
	uchar	rex;		/* REX prefix (or zero) */
	char	osize;		/* &#39;W&#39; or &#39;L&#39; (or &#39;Q&#39; on amd64) */
	char	asize;		/* address size &#39;W&#39; or &#39;L&#39; (or &#39;Q&#39; or amd64) */
	uchar	mod;		/* bits 6-7 of mod r/m field */
	uchar	reg;		/* bits 3-5 of mod r/m field */
	char	ss;		/* bits 6-7 of SIB */
	char	index;		/* bits 3-5 of SIB */
	char	base;		/* bits 0-2 of SIB */
	char	rip;		/* RIP-relative in amd64 mode */
	uchar	opre;		/* f2/f3 could introduce media */
	short	seg;		/* segment of far address */
	uint32	disp;		/* displacement */
	uint32 	imm;		/* immediate */
	uint32 	imm2;		/* second immediate operand */
	uvlong	imm64;		/* big immediate */
	char	*curr;		/* fill level in output buffer */
	char	*end;		/* end of output buffer */
	char	*err;		/* error message */
};

	/* 386 register (ha!) set */
enum{
	AX=0,
	CX,
	DX,
	BX,
	SP,
	BP,
	SI,
	DI,

	/* amd64 */
	R8,
	R9,
	R10,
	R11,
	R12,
	R13,
	R14,
	R15
};

	/* amd64 rex extension byte */
enum{
	REXW		= 1&lt;&lt;3,	/* =1, 64-bit operand size */
	REXR		= 1&lt;&lt;2,	/* extend modrm reg */
	REXX		= 1&lt;&lt;1,	/* extend sib index */
	REXB		= 1&lt;&lt;0	/* extend modrm r/m, sib base, or opcode reg */
};

	/* Operand Format codes */
/*
%A	-	address size register modifier (!asize -&gt; &#39;E&#39;)
%C	-	Control register CR0/CR1/CR2
%D	-	Debug register DR0/DR1/DR2/DR3/DR6/DR7
%I	-	second immediate operand
%O	-	Operand size register modifier (!osize -&gt; &#39;E&#39;)
%T	-	Test register TR6/TR7
%S	-	size code (&#39;W&#39; or &#39;L&#39;)
%W	-	Weird opcode: OSIZE == &#39;W&#39; =&gt; &#34;CBW&#34;; else =&gt; &#34;CWDE&#34;
%d	-	displacement 16-32 bits
%e	-	effective address - Mod R/M value
%f	-	floating point register F0-F7 - from Mod R/M register
%g	-	segment register
%i	-	immediate operand 8-32 bits
%p	-	PC-relative - signed displacement in immediate field
%r	-	Reg from Mod R/M
%w	-	Weird opcode: OSIZE == &#39;W&#39; =&gt; &#34;CWD&#34;; else =&gt; &#34;CDQ&#34;
*/

typedef struct Optable Optable;
struct Optable
{
	char	operand[2];
	void	*proto;		/* actually either (char*) or (Optable*) */
};
	/* Operand decoding codes */
enum {
	Ib = 1,			/* 8-bit immediate - (no sign extension)*/
	Ibs,			/* 8-bit immediate (sign extended) */
	Jbs,			/* 8-bit sign-extended immediate in jump or call */
	Iw,			/* 16-bit immediate -&gt; imm */
	Iw2,			/* 16-bit immediate -&gt; imm2 */
	Iwd,			/* Operand-sized immediate (no sign extension)*/
	Iwdq,			/* Operand-sized immediate, possibly 64 bits */
	Awd,			/* Address offset */
	Iwds,			/* Operand-sized immediate (sign extended) */
	RM,			/* Word or int32 R/M field with register (/r) */
	RMB,			/* Byte R/M field with register (/r) */
	RMOP,			/* Word or int32 R/M field with op code (/digit) */
	RMOPB,			/* Byte R/M field with op code (/digit) */
	RMR,			/* R/M register only (mod = 11) */
	RMM,			/* R/M memory only (mod = 0/1/2) */
	R0,			/* Base reg of Mod R/M is literal 0x00 */
	R1,			/* Base reg of Mod R/M is literal 0x01 */
	FRMOP,			/* Floating point R/M field with opcode */
	FRMEX,			/* Extended floating point R/M field with opcode */
	JUMP,			/* Jump or Call flag - no operand */
	RET,			/* Return flag - no operand */
	OA,			/* literal 0x0a byte */
	PTR,			/* Seg:Displacement addr (ptr16:16 or ptr16:32) */
	AUX,			/* Multi-byte op code - Auxiliary table */
	AUXMM,			/* multi-byte op code - auxiliary table chosen by prefix */
	PRE,			/* Instr Prefix */
	OPRE,			/* Instr Prefix or media op extension */
	SEG,			/* Segment Prefix */
	OPOVER,			/* Operand size override */
	ADDOVER,		/* Address size override */
};

static Optable optab0F00[8]=
{
[0x00]	0,0,		&#34;MOVW	LDT,%e&#34;,
[0x01]	0,0,		&#34;MOVW	TR,%e&#34;,
[0x02]	0,0,		&#34;MOVW	%e,LDT&#34;,
[0x03]	0,0,		&#34;MOVW	%e,TR&#34;,
[0x04]	0,0,		&#34;VERR	%e&#34;,
[0x05]	0,0,		&#34;VERW	%e&#34;,
};

static Optable optab0F01[8]=
{
[0x00]	0,0,		&#34;MOVL	GDTR,%e&#34;,
[0x01]	0,0,		&#34;MOVL	IDTR,%e&#34;,
[0x02]	0,0,		&#34;MOVL	%e,GDTR&#34;,
[0x03]	0,0,		&#34;MOVL	%e,IDTR&#34;,
[0x04]	0,0,		&#34;MOVW	MSW,%e&#34;,	/* word */
[0x06]	0,0,		&#34;MOVW	%e,MSW&#34;,	/* word */
[0x07]	0,0,		&#34;INVLPG	%e&#34;,		/* or SWAPGS */
};

static Optable optab0F01F8[1]=
{
[0x00]	0,0,		&#34;SWAPGS&#34;,
};

/* 0F71 */
/* 0F72 */
/* 0F73 */

static Optable optab0FAE[8]=
{
[0x00]	0,0,		&#34;FXSAVE	%e&#34;,
[0x01]	0,0,		&#34;FXRSTOR	%e&#34;,
[0x02]	0,0,		&#34;LDMXCSR	%e&#34;,
[0x03]	0,0,		&#34;STMXCSR	%e&#34;,
[0x05]	0,0,		&#34;LFENCE&#34;,
[0x06]	0,0,		&#34;MFENCE&#34;,
[0x07]	0,0,		&#34;SFENCE&#34;,
};

/* 0F18 */
/* 0F0D */

static Optable optab0FBA[8]=
{
[0x04]	Ib,0,		&#34;BT%S	%i,%e&#34;,
[0x05]	Ib,0,		&#34;BTS%S	%i,%e&#34;,
[0x06]	Ib,0,		&#34;BTR%S	%i,%e&#34;,
[0x07]	Ib,0,		&#34;BTC%S	%i,%e&#34;,
};

static Optable optab0F0F[256]=
{
[0x0c]	0,0,		&#34;PI2FW	%m,%M&#34;,
[0x0d]	0,0,		&#34;PI2L	%m,%M&#34;,
[0x1c]	0,0,		&#34;PF2IW	%m,%M&#34;,
[0x1d]	0,0,		&#34;PF2IL	%m,%M&#34;,
[0x8a]	0,0,		&#34;PFNACC	%m,%M&#34;,
[0x8e]	0,0,		&#34;PFPNACC	%m,%M&#34;,
[0x90]	0,0,		&#34;PFCMPGE	%m,%M&#34;,
[0x94]	0,0,		&#34;PFMIN	%m,%M&#34;,
[0x96]	0,0,		&#34;PFRCP	%m,%M&#34;,
[0x97]	0,0,		&#34;PFRSQRT	%m,%M&#34;,
[0x9a]	0,0,		&#34;PFSUB	%m,%M&#34;,
[0x9e]	0,0,		&#34;PFADD	%m,%M&#34;,
[0xa0]	0,0,		&#34;PFCMPGT	%m,%M&#34;,
[0xa4]	0,0,		&#34;PFMAX	%m,%M&#34;,
[0xa6]	0,0,		&#34;PFRCPIT1	%m,%M&#34;,
[0xa7]	0,0,		&#34;PFRSQIT1	%m,%M&#34;,
[0xaa]	0,0,		&#34;PFSUBR	%m,%M&#34;,
[0xae]	0,0,		&#34;PFACC	%m,%M&#34;,
[0xb0]	0,0,		&#34;PFCMPEQ	%m,%M&#34;,
[0xb4]	0,0,		&#34;PFMUL	%m,%M&#34;,
[0xb6]	0,0,		&#34;PFRCPI2T	%m,%M&#34;,
[0xb7]	0,0,		&#34;PMULHRW	%m,%M&#34;,
[0xbb]	0,0,		&#34;PSWAPL	%m,%M&#34;,
};

static Optable optab0FC7[8]=
{
[0x01]	0,0,		&#34;CMPXCHG8B	%e&#34;,
};

static Optable optab660F71[8]=
{
[0x02]	Ib,0,		&#34;PSRLW	%i,%X&#34;,
[0x04]	Ib,0,		&#34;PSRAW	%i,%X&#34;,
[0x06]	Ib,0,		&#34;PSLLW	%i,%X&#34;,
};

static Optable optab660F72[8]=
{
[0x02]	Ib,0,		&#34;PSRLL	%i,%X&#34;,
[0x04]	Ib,0,		&#34;PSRAL	%i,%X&#34;,
[0x06]	Ib,0,		&#34;PSLLL	%i,%X&#34;,
};

static Optable optab660F73[8]=
{
[0x02]	Ib,0,		&#34;PSRLQ	%i,%X&#34;,
[0x03]	Ib,0,		&#34;PSRLO	%i,%X&#34;,
[0x06]	Ib,0,		&#34;PSLLQ	%i,%X&#34;,
[0x07]	Ib,0,		&#34;PSLLO	%i,%X&#34;,
};

static Optable optab660F[256]=
{
[0x2B]	RM,0,		&#34;MOVNTPD	%x,%e&#34;,
[0x2E]	RM,0,		&#34;UCOMISD	%x,%X&#34;,
[0x2F]	RM,0,		&#34;COMISD	%x,%X&#34;,
[0x5A]	RM,0,		&#34;CVTPD2PS	%x,%X&#34;,
[0x5B]	RM,0,		&#34;CVTPS2PL	%x,%X&#34;,
[0x6A]	RM,0,		&#34;PUNPCKHLQ %x,%X&#34;,
[0x6B]	RM,0,		&#34;PACKSSLW %x,%X&#34;,
[0x6C]	RM,0,		&#34;PUNPCKLQDQ %x,%X&#34;,
[0x6D]	RM,0,		&#34;PUNPCKHQDQ %x,%X&#34;,
[0x6E]	RM,0,		&#34;MOV%S	%e,%X&#34;,
[0x6F]	RM,0,		&#34;MOVO	%x,%X&#34;,		/* MOVDQA */
[0x70]	RM,Ib,		&#34;PSHUFL	%i,%x,%X&#34;,
[0x71]	RMOP,0,		optab660F71,
[0x72]	RMOP,0,		optab660F72,
[0x73]	RMOP,0,		optab660F73,
[0x7E]	RM,0,		&#34;MOV%S	%X,%e&#34;,
[0x7F]	RM,0,		&#34;MOVO	%X,%x&#34;,
[0xC4]	RM,Ib,		&#34;PINSRW	%i,%e,%X&#34;,
[0xC5]	RMR,Ib,		&#34;PEXTRW	%i,%X,%e&#34;,
[0xD4]	RM,0,		&#34;PADDQ	%x,%X&#34;,
[0xD5]	RM,0,		&#34;PMULLW	%x,%X&#34;,
[0xD6]	RM,0,		&#34;MOVQ	%X,%x&#34;,
[0xE6]	RM,0,		&#34;CVTTPD2PL	%x,%X&#34;,
[0xE7]	RM,0,		&#34;MOVNTO	%X,%e&#34;,
[0xF7]	RM,0,		&#34;MASKMOVOU	%x,%X&#34;,
};

static Optable optabF20F[256]=
{
[0x10]	RM,0,		&#34;MOVSD	%x,%X&#34;,
[0x11]	RM,0,		&#34;MOVSD	%X,%x&#34;,
[0x2A]	RM,0,		&#34;CVTS%S2SD	%e,%X&#34;,
[0x2C]	RM,0,		&#34;CVTTSD2S%S	%x,%r&#34;,
[0x2D]	RM,0,		&#34;CVTSD2S%S	%x,%r&#34;,
[0x5A]	RM,0,		&#34;CVTSD2SS	%x,%X&#34;,
[0x6F]	RM,0,		&#34;MOVOU	%x,%X&#34;,
[0x70]	RM,Ib,		&#34;PSHUFLW	%i,%x,%X&#34;,
[0x7F]	RM,0,		&#34;MOVOU	%X,%x&#34;,
[0xD6]	RM,0,		&#34;MOVQOZX	%M,%X&#34;,
[0xE6]	RM,0,		&#34;CVTPD2PL	%x,%X&#34;,
};

static Optable optabF30F[256]=
{
[0x10]	RM,0,		&#34;MOVSS	%x,%X&#34;,
[0x11]	RM,0,		&#34;MOVSS	%X,%x&#34;,
[0x2A]	RM,0,		&#34;CVTS%S2SS	%e,%X&#34;,
[0x2C]	RM,0,		&#34;CVTTSS2S%S	%x,%r&#34;,
[0x2D]	RM,0,		&#34;CVTSS2S%S	%x,%r&#34;,
[0x5A]	RM,0,		&#34;CVTSS2SD	%x,%X&#34;,
[0x5B]	RM,0,		&#34;CVTTPS2PL	%x,%X&#34;,
[0x6F]	RM,0,		&#34;MOVOU	%x,%X&#34;,
[0x70]	RM,Ib,		&#34;PSHUFHW	%i,%x,%X&#34;,
[0x7E]	RM,0,		&#34;MOVQOZX	%x,%X&#34;,
[0x7F]	RM,0,		&#34;MOVOU	%X,%x&#34;,
[0xD6]	RM,0,		&#34;MOVQOZX	%m*,%X&#34;,
[0xE6]	RM,0,		&#34;CVTPL2PD	%x,%X&#34;,
};

static Optable optab0F[256]=
{
[0x00]	RMOP,0,		optab0F00,
[0x01]	RMOP,0,		optab0F01,
[0x02]	RM,0,		&#34;LAR	%e,%r&#34;,
[0x03]	RM,0,		&#34;LSL	%e,%r&#34;,
[0x05]	0,0,		&#34;SYSCALL&#34;,
[0x06]	0,0,		&#34;CLTS&#34;,
[0x07]	0,0,		&#34;SYSRET&#34;,
[0x08]	0,0,		&#34;INVD&#34;,
[0x09]	0,0,		&#34;WBINVD&#34;,
[0x0B]	0,0,		&#34;UD2&#34;,
[0x0F]	RM,AUX,		optab0F0F,		/* 3DNow! */
[0x10]	RM,0,		&#34;MOVU%s	%x,%X&#34;,
[0x11]	RM,0,		&#34;MOVU%s	%X,%x&#34;,
[0x12]	RM,0,		&#34;MOV[H]L%s	%x,%X&#34;,	/* TO DO: H if source is XMM */
[0x13]	RM,0,		&#34;MOVL%s	%X,%e&#34;,
[0x14]	RM,0,		&#34;UNPCKL%s	%x,%X&#34;,
[0x15]	RM,0,		&#34;UNPCKH%s	%x,%X&#34;,
[0x16]	RM,0,		&#34;MOV[L]H%s	%x,%X&#34;,	/* TO DO: L if source is XMM */
[0x17]	RM,0,		&#34;MOVH%s	%X,%x&#34;,
[0x20]	RMR,0,		&#34;MOVL	%C,%e&#34;,
[0x21]	RMR,0,		&#34;MOVL	%D,%e&#34;,
[0x22]	RMR,0,		&#34;MOVL	%e,%C&#34;,
[0x23]	RMR,0,		&#34;MOVL	%e,%D&#34;,
[0x24]	RMR,0,		&#34;MOVL	%T,%e&#34;,
[0x26]	RMR,0,		&#34;MOVL	%e,%T&#34;,
[0x28]	RM,0,		&#34;MOVA%s	%x,%X&#34;,
[0x29]	RM,0,		&#34;MOVA%s	%X,%x&#34;,
[0x2A]	RM,0,		&#34;CVTPL2%s	%m*,%X&#34;,
[0x2B]	RM,0,		&#34;MOVNT%s	%X,%e&#34;,
[0x2C]	RM,0,		&#34;CVTT%s2PL	%x,%M&#34;,
[0x2D]	RM,0,		&#34;CVT%s2PL	%x,%M&#34;,
[0x2E]	RM,0,		&#34;UCOMISS	%x,%X&#34;,
[0x2F]	RM,0,		&#34;COMISS	%x,%X&#34;,
[0x30]	0,0,		&#34;WRMSR&#34;,
[0x31]	0,0,		&#34;RDTSC&#34;,
[0x32]	0,0,		&#34;RDMSR&#34;,
[0x33]	0,0,		&#34;RDPMC&#34;,
[0x42]	RM,0,		&#34;CMOVC	%e,%r&#34;,		/* CF */
[0x43]	RM,0,		&#34;CMOVNC	%e,%r&#34;,		/* ¬ CF */
[0x44]	RM,0,		&#34;CMOVZ	%e,%r&#34;,		/* ZF */
[0x45]	RM,0,		&#34;CMOVNZ	%e,%r&#34;,		/* ¬ ZF */
[0x46]	RM,0,		&#34;CMOVBE	%e,%r&#34;,		/* CF ∨ ZF */
[0x47]	RM,0,		&#34;CMOVA	%e,%r&#34;,		/* ¬CF ∧ ¬ZF */
[0x48]	RM,0,		&#34;CMOVS	%e,%r&#34;,		/* SF */
[0x49]	RM,0,		&#34;CMOVNS	%e,%r&#34;,		/* ¬ SF */
[0x4A]	RM,0,		&#34;CMOVP	%e,%r&#34;,		/* PF */
[0x4B]	RM,0,		&#34;CMOVNP	%e,%r&#34;,		/* ¬ PF */
[0x4C]	RM,0,		&#34;CMOVLT	%e,%r&#34;,		/* LT ≡ OF ≠ SF */
[0x4D]	RM,0,		&#34;CMOVGE	%e,%r&#34;,		/* GE ≡ ZF ∨ SF */
[0x4E]	RM,0,		&#34;CMOVLE	%e,%r&#34;,		/* LE ≡ ZF ∨ LT */
[0x4F]	RM,0,		&#34;CMOVGT	%e,%r&#34;,		/* GT ≡ ¬ZF ∧ GE */
[0x50]	RM,0,		&#34;MOVMSK%s	%X,%r&#34;,	/* TO DO: check */
[0x51]	RM,0,		&#34;SQRT%s	%x,%X&#34;,
[0x52]	RM,0,		&#34;RSQRT%s	%x,%X&#34;,
[0x53]	RM,0,		&#34;RCP%s	%x,%X&#34;,
[0x54]	RM,0,		&#34;AND%s	%x,%X&#34;,
[0x55]	RM,0,		&#34;ANDN%s	%x,%X&#34;,
[0x56]	RM,0,		&#34;OR%s	%x,%X&#34;,		/* TO DO: S/D */
[0x57]	RM,0,		&#34;XOR%s	%x,%X&#34;,		/* S/D */
[0x58]	RM,0,		&#34;ADD%s	%x,%X&#34;,		/* S/P S/D */
[0x59]	RM,0,		&#34;MUL%s	%x,%X&#34;,
[0x5A]	RM,0,		&#34;CVTPS2PD	%x,%X&#34;,
[0x5B]	RM,0,		&#34;CVTPL2PS	%x,%X&#34;,
[0x5C]	RM,0,		&#34;SUB%s	%x,%X&#34;,
[0x5D]	RM,0,		&#34;MIN%s	%x,%X&#34;,
[0x5E]	RM,0,		&#34;DIV%s	%x,%X&#34;,		/* TO DO: S/P S/D */
[0x5F]	RM,0,		&#34;MAX%s	%x,%X&#34;,
[0x60]	RM,0,		&#34;PUNPCKLBW %m,%M&#34;,
[0x61]	RM,0,		&#34;PUNPCKLWL %m,%M&#34;,
[0x62]	RM,0,		&#34;PUNPCKLLQ %m,%M&#34;,
[0x63]	RM,0,		&#34;PACKSSWB %m,%M&#34;,
[0x64]	RM,0,		&#34;PCMPGTB %m,%M&#34;,
[0x65]	RM,0,		&#34;PCMPGTW %m,%M&#34;,
[0x66]	RM,0,		&#34;PCMPGTL %m,%M&#34;,
[0x67]	RM,0,		&#34;PACKUSWB %m,%M&#34;,
[0x68]	RM,0,		&#34;PUNPCKHBW %m,%M&#34;,
[0x69]	RM,0,		&#34;PUNPCKHWL %m,%M&#34;,
[0x6A]	RM,0,		&#34;PUNPCKHLQ %m,%M&#34;,
[0x6B]	RM,0,		&#34;PACKSSLW %m,%M&#34;,
[0x6E]	RM,0,		&#34;MOV%S %e,%M&#34;,
[0x6F]	RM,0,		&#34;MOVQ %m,%M&#34;,
[0x70]	RM,Ib,		&#34;PSHUFW	%i,%m,%M&#34;,
[0x74]	RM,0,		&#34;PCMPEQB %m,%M&#34;,
[0x75]	RM,0,		&#34;PCMPEQW %m,%M&#34;,
[0x76]	RM,0,		&#34;PCMPEQL %m,%M&#34;,
[0x7E]	RM,0,		&#34;MOV%S %M,%e&#34;,
[0x7F]	RM,0,		&#34;MOVQ %M,%m&#34;,
[0xAE]	RMOP,0,		optab0FAE,
[0xAA]	0,0,		&#34;RSM&#34;,
[0xB0]	RM,0,		&#34;CMPXCHGB	%r,%e&#34;,
[0xB1]	RM,0,		&#34;CMPXCHG%S	%r,%e&#34;,
[0xC0]	RMB,0,		&#34;XADDB	%r,%e&#34;,
[0xC1]	RM,0,		&#34;XADD%S	%r,%e&#34;,
[0xC2]	RM,Ib,		&#34;CMP%s	%i,%x,%X&#34;,
[0xC3]	RM,0,		&#34;MOVNTI%S	%r,%e&#34;,
[0xC6]	RM,Ib,		&#34;SHUF%s	%i,%x,%X&#34;,
[0xC8]	0,0,		&#34;BSWAP	AX&#34;,
[0xC9]	0,0,		&#34;BSWAP	CX&#34;,
[0xCA]	0,0,		&#34;BSWAP	DX&#34;,
[0xCB]	0,0,		&#34;BSWAP	BX&#34;,
[0xCC]	0,0,		&#34;BSWAP	SP&#34;,
[0xCD]	0,0,		&#34;BSWAP	BP&#34;,
[0xCE]	0,0,		&#34;BSWAP	SI&#34;,
[0xCF]	0,0,		&#34;BSWAP	DI&#34;,
[0xD1]	RM,0,		&#34;PSRLW %m,%M&#34;,
[0xD2]	RM,0,		&#34;PSRLL %m,%M&#34;,
[0xD3]	RM,0,		&#34;PSRLQ %m,%M&#34;,
[0xD5]	RM,0,		&#34;PMULLW %m,%M&#34;,
[0xD6]	RM,0,		&#34;MOVQOZX	%m*,%X&#34;,
[0xD7]	RM,0,		&#34;PMOVMSKB %m,%r&#34;,
[0xD8]	RM,0,		&#34;PSUBUSB %m,%M&#34;,
[0xD9]	RM,0,		&#34;PSUBUSW %m,%M&#34;,
[0xDA]	RM,0,		&#34;PMINUB %m,%M&#34;,
[0xDB]	RM,0,		&#34;PAND %m,%M&#34;,
[0xDC]	RM,0,		&#34;PADDUSB %m,%M&#34;,
[0xDD]	RM,0,		&#34;PADDUSW %m,%M&#34;,
[0xDE]	RM,0,		&#34;PMAXUB %m,%M&#34;,
[0xDF]	RM,0,		&#34;PANDN %m,%M&#34;,
[0xE0]	RM,0,		&#34;PAVGB %m,%M&#34;,
[0xE1]	RM,0,		&#34;PSRAW %m,%M&#34;,
[0xE2]	RM,0,		&#34;PSRAL %m,%M&#34;,
[0xE3]	RM,0,		&#34;PAVGW %m,%M&#34;,
[0xE4]	RM,0,		&#34;PMULHUW %m,%M&#34;,
[0xE5]	RM,0,		&#34;PMULHW %m,%M&#34;,
[0xE7]	RM,0,		&#34;MOVNTQ	%M,%e&#34;,
[0xE8]	RM,0,		&#34;PSUBSB %m,%M&#34;,
[0xE9]	RM,0,		&#34;PSUBSW %m,%M&#34;,
[0xEA]	RM,0,		&#34;PMINSW %m,%M&#34;,
[0xEB]	RM,0,		&#34;POR %m,%M&#34;,
[0xEC]	RM,0,		&#34;PADDSB %m,%M&#34;,
[0xED]	RM,0,		&#34;PADDSW %m,%M&#34;,
[0xEE]	RM,0,		&#34;PMAXSW %m,%M&#34;,
[0xEF]	RM,0,		&#34;PXOR %m,%M&#34;,
[0xF1]	RM,0,		&#34;PSLLW %m,%M&#34;,
[0xF2]	RM,0,		&#34;PSLLL %m,%M&#34;,
[0xF3]	RM,0,		&#34;PSLLQ %m,%M&#34;,
[0xF4]	RM,0,		&#34;PMULULQ	%m,%M&#34;,
[0xF5]	RM,0,		&#34;PMADDWL %m,%M&#34;,
[0xF6]	RM,0,		&#34;PSADBW %m,%M&#34;,
[0xF7]	RMR,0,		&#34;MASKMOVQ	%m,%M&#34;,
[0xF8]	RM,0,		&#34;PSUBB %m,%M&#34;,
[0xF9]	RM,0,		&#34;PSUBW %m,%M&#34;,
[0xFA]	RM,0,		&#34;PSUBL %m,%M&#34;,
[0xFC]	RM,0,		&#34;PADDB %m,%M&#34;,
[0xFD]	RM,0,		&#34;PADDW %m,%M&#34;,
[0xFE]	RM,0,		&#34;PADDL %m,%M&#34;,

[0x80]	Iwds,0,		&#34;JOS	%p&#34;,
[0x81]	Iwds,0,		&#34;JOC	%p&#34;,
[0x82]	Iwds,0,		&#34;JCS	%p&#34;,
[0x83]	Iwds,0,		&#34;JCC	%p&#34;,
[0x84]	Iwds,0,		&#34;JEQ	%p&#34;,
[0x85]	Iwds,0,		&#34;JNE	%p&#34;,
[0x86]	Iwds,0,		&#34;JLS	%p&#34;,
[0x87]	Iwds,0,		&#34;JHI	%p&#34;,
[0x88]	Iwds,0,		&#34;JMI	%p&#34;,
[0x89]	Iwds,0,		&#34;JPL	%p&#34;,
[0x8a]	Iwds,0,		&#34;JPS	%p&#34;,
[0x8b]	Iwds,0,		&#34;JPC	%p&#34;,
[0x8c]	Iwds,0,		&#34;JLT	%p&#34;,
[0x8d]	Iwds,0,		&#34;JGE	%p&#34;,
[0x8e]	Iwds,0,		&#34;JLE	%p&#34;,
[0x8f]	Iwds,0,		&#34;JGT	%p&#34;,
[0x90]	RMB,0,		&#34;SETOS	%e&#34;,
[0x91]	RMB,0,		&#34;SETOC	%e&#34;,
[0x92]	RMB,0,		&#34;SETCS	%e&#34;,
[0x93]	RMB,0,		&#34;SETCC	%e&#34;,
[0x94]	RMB,0,		&#34;SETEQ	%e&#34;,
[0x95]	RMB,0,		&#34;SETNE	%e&#34;,
[0x96]	RMB,0,		&#34;SETLS	%e&#34;,
[0x97]	RMB,0,		&#34;SETHI	%e&#34;,
[0x98]	RMB,0,		&#34;SETMI	%e&#34;,
[0x99]	RMB,0,		&#34;SETPL	%e&#34;,
[0x9a]	RMB,0,		&#34;SETPS	%e&#34;,
[0x9b]	RMB,0,		&#34;SETPC	%e&#34;,
[0x9c]	RMB,0,		&#34;SETLT	%e&#34;,
[0x9d]	RMB,0,		&#34;SETGE	%e&#34;,
[0x9e]	RMB,0,		&#34;SETLE	%e&#34;,
[0x9f]	RMB,0,		&#34;SETGT	%e&#34;,
[0xa0]	0,0,		&#34;PUSHL	FS&#34;,
[0xa1]	0,0,		&#34;POPL	FS&#34;,
[0xa2]	0,0,		&#34;CPUID&#34;,
[0xa3]	RM,0,		&#34;BT%S	%r,%e&#34;,
[0xa4]	RM,Ib,		&#34;SHLD%S	%r,%i,%e&#34;,
[0xa5]	RM,0,		&#34;SHLD%S	%r,CL,%e&#34;,
[0xa8]	0,0,		&#34;PUSHL	GS&#34;,
[0xa9]	0,0,		&#34;POPL	GS&#34;,
[0xab]	RM,0,		&#34;BTS%S	%r,%e&#34;,
[0xac]	RM,Ib,		&#34;SHRD%S	%r,%i,%e&#34;,
[0xad]	RM,0,		&#34;SHRD%S	%r,CL,%e&#34;,
[0xaf]	RM,0,		&#34;IMUL%S	%e,%r&#34;,
[0xb2]	RMM,0,		&#34;LSS	%e,%r&#34;,
[0xb3]	RM,0,		&#34;BTR%S	%r,%e&#34;,
[0xb4]	RMM,0,		&#34;LFS	%e,%r&#34;,
[0xb5]	RMM,0,		&#34;LGS	%e,%r&#34;,
[0xb6]	RMB,0,		&#34;MOVBZX	%e,%R&#34;,
[0xb7]	RM,0,		&#34;MOVWZX	%e,%R&#34;,
[0xba]	RMOP,0,		optab0FBA,
[0xbb]	RM,0,		&#34;BTC%S	%e,%r&#34;,
[0xbc]	RM,0,		&#34;BSF%S	%e,%r&#34;,
[0xbd]	RM,0,		&#34;BSR%S	%e,%r&#34;,
[0xbe]	RMB,0,		&#34;MOVBSX	%e,%R&#34;,
[0xbf]	RM,0,		&#34;MOVWSX	%e,%R&#34;,
[0xc7]	RMOP,0,		optab0FC7,
};

static Optable optab80[8]=
{
[0x00]	Ib,0,		&#34;ADDB	%i,%e&#34;,
[0x01]	Ib,0,		&#34;ORB	%i,%e&#34;,
[0x02]	Ib,0,		&#34;ADCB	%i,%e&#34;,
[0x03]	Ib,0,		&#34;SBBB	%i,%e&#34;,
[0x04]	Ib,0,		&#34;ANDB	%i,%e&#34;,
[0x05]	Ib,0,		&#34;SUBB	%i,%e&#34;,
[0x06]	Ib,0,		&#34;XORB	%i,%e&#34;,
[0x07]	Ib,0,		&#34;CMPB	%e,%i&#34;,
};

static Optable optab81[8]=
{
[0x00]	Iwd,0,		&#34;ADD%S	%i,%e&#34;,
[0x01]	Iwd,0,		&#34;OR%S	%i,%e&#34;,
[0x02]	Iwd,0,		&#34;ADC%S	%i,%e&#34;,
[0x03]	Iwd,0,		&#34;SBB%S	%i,%e&#34;,
[0x04]	Iwd,0,		&#34;AND%S	%i,%e&#34;,
[0x05]	Iwd,0,		&#34;SUB%S	%i,%e&#34;,
[0x06]	Iwd,0,		&#34;XOR%S	%i,%e&#34;,
[0x07]	Iwd,0,		&#34;CMP%S	%e,%i&#34;,
};

static Optable optab83[8]=
{
[0x00]	Ibs,0,		&#34;ADD%S	%i,%e&#34;,
[0x01]	Ibs,0,		&#34;OR%S	%i,%e&#34;,
[0x02]	Ibs,0,		&#34;ADC%S	%i,%e&#34;,
[0x03]	Ibs,0,		&#34;SBB%S	%i,%e&#34;,
[0x04]	Ibs,0,		&#34;AND%S	%i,%e&#34;,
[0x05]	Ibs,0,		&#34;SUB%S	%i,%e&#34;,
[0x06]	Ibs,0,		&#34;XOR%S	%i,%e&#34;,
[0x07]	Ibs,0,		&#34;CMP%S	%e,%i&#34;,
};

static Optable optabC0[8] =
{
[0x00]	Ib,0,		&#34;ROLB	%i,%e&#34;,
[0x01]	Ib,0,		&#34;RORB	%i,%e&#34;,
[0x02]	Ib,0,		&#34;RCLB	%i,%e&#34;,
[0x03]	Ib,0,		&#34;RCRB	%i,%e&#34;,
[0x04]	Ib,0,		&#34;SHLB	%i,%e&#34;,
[0x05]	Ib,0,		&#34;SHRB	%i,%e&#34;,
[0x07]	Ib,0,		&#34;SARB	%i,%e&#34;,
};

static Optable optabC1[8] =
{
[0x00]	Ib,0,		&#34;ROL%S	%i,%e&#34;,
[0x01]	Ib,0,		&#34;ROR%S	%i,%e&#34;,
[0x02]	Ib,0,		&#34;RCL%S	%i,%e&#34;,
[0x03]	Ib,0,		&#34;RCR%S	%i,%e&#34;,
[0x04]	Ib,0,		&#34;SHL%S	%i,%e&#34;,
[0x05]	Ib,0,		&#34;SHR%S	%i,%e&#34;,
[0x07]	Ib,0,		&#34;SAR%S	%i,%e&#34;,
};

static Optable optabD0[8] =
{
[0x00]	0,0,		&#34;ROLB	%e&#34;,
[0x01]	0,0,		&#34;RORB	%e&#34;,
[0x02]	0,0,		&#34;RCLB	%e&#34;,
[0x03]	0,0,		&#34;RCRB	%e&#34;,
[0x04]	0,0,		&#34;SHLB	%e&#34;,
[0x05]	0,0,		&#34;SHRB	%e&#34;,
[0x07]	0,0,		&#34;SARB	%e&#34;,
};

static Optable optabD1[8] =
{
[0x00]	0,0,		&#34;ROL%S	%e&#34;,
[0x01]	0,0,		&#34;ROR%S	%e&#34;,
[0x02]	0,0,		&#34;RCL%S	%e&#34;,
[0x03]	0,0,		&#34;RCR%S	%e&#34;,
[0x04]	0,0,		&#34;SHL%S	%e&#34;,
[0x05]	0,0,		&#34;SHR%S	%e&#34;,
[0x07]	0,0,		&#34;SAR%S	%e&#34;,
};

static Optable optabD2[8] =
{
[0x00]	0,0,		&#34;ROLB	CL,%e&#34;,
[0x01]	0,0,		&#34;RORB	CL,%e&#34;,
[0x02]	0,0,		&#34;RCLB	CL,%e&#34;,
[0x03]	0,0,		&#34;RCRB	CL,%e&#34;,
[0x04]	0,0,		&#34;SHLB	CL,%e&#34;,
[0x05]	0,0,		&#34;SHRB	CL,%e&#34;,
[0x07]	0,0,		&#34;SARB	CL,%e&#34;,
};

static Optable optabD3[8] =
{
[0x00]	0,0,		&#34;ROL%S	CL,%e&#34;,
[0x01]	0,0,		&#34;ROR%S	CL,%e&#34;,
[0x02]	0,0,		&#34;RCL%S	CL,%e&#34;,
[0x03]	0,0,		&#34;RCR%S	CL,%e&#34;,
[0x04]	0,0,		&#34;SHL%S	CL,%e&#34;,
[0x05]	0,0,		&#34;SHR%S	CL,%e&#34;,
[0x07]	0,0,		&#34;SAR%S	CL,%e&#34;,
};

static Optable optabD8[8+8] =
{
[0x00]	0,0,		&#34;FADDF	%e,F0&#34;,
[0x01]	0,0,		&#34;FMULF	%e,F0&#34;,
[0x02]	0,0,		&#34;FCOMF	%e,F0&#34;,
[0x03]	0,0,		&#34;FCOMFP	%e,F0&#34;,
[0x04]	0,0,		&#34;FSUBF	%e,F0&#34;,
[0x05]	0,0,		&#34;FSUBRF	%e,F0&#34;,
[0x06]	0,0,		&#34;FDIVF	%e,F0&#34;,
[0x07]	0,0,		&#34;FDIVRF	%e,F0&#34;,
[0x08]	0,0,		&#34;FADDD	%f,F0&#34;,
[0x09]	0,0,		&#34;FMULD	%f,F0&#34;,
[0x0a]	0,0,		&#34;FCOMD	%f,F0&#34;,
[0x0b]	0,0,		&#34;FCOMPD	%f,F0&#34;,
[0x0c]	0,0,		&#34;FSUBD	%f,F0&#34;,
[0x0d]	0,0,		&#34;FSUBRD	%f,F0&#34;,
[0x0e]	0,0,		&#34;FDIVD	%f,F0&#34;,
[0x0f]	0,0,		&#34;FDIVRD	%f,F0&#34;,
};
/*
 *	optabD9 and optabDB use the following encoding:
 *	if (0 &lt;= modrm &lt;= 2) instruction = optabDx[modrm&amp;0x07];
 *	else instruction = optabDx[(modrm&amp;0x3f)+8];
 *
 *	the instructions for MOD == 3, follow the 8 instructions
 *	for the other MOD values stored at the front of the table.
 */
static Optable optabD9[64+8] =
{
[0x00]	0,0,		&#34;FMOVF	%e,F0&#34;,
[0x02]	0,0,		&#34;FMOVF	F0,%e&#34;,
[0x03]	0,0,		&#34;FMOVFP	F0,%e&#34;,
[0x04]	0,0,		&#34;FLDENV%S %e&#34;,
[0x05]	0,0,		&#34;FLDCW	%e&#34;,
[0x06]	0,0,		&#34;FSTENV%S %e&#34;,
[0x07]	0,0,		&#34;FSTCW	%e&#34;,
[0x08]	0,0,		&#34;FMOVD	F0,F0&#34;,		/* Mod R/M = 11xx xxxx*/
[0x09]	0,0,		&#34;FMOVD	F1,F0&#34;,
[0x0a]	0,0,		&#34;FMOVD	F2,F0&#34;,
[0x0b]	0,0,		&#34;FMOVD	F3,F0&#34;,
[0x0c]	0,0,		&#34;FMOVD	F4,F0&#34;,
[0x0d]	0,0,		&#34;FMOVD	F5,F0&#34;,
[0x0e]	0,0,		&#34;FMOVD	F6,F0&#34;,
[0x0f]	0,0,		&#34;FMOVD	F7,F0&#34;,
[0x10]	0,0,		&#34;FXCHD	F0,F0&#34;,
[0x11]	0,0,		&#34;FXCHD	F1,F0&#34;,
[0x12]	0,0,		&#34;FXCHD	F2,F0&#34;,
[0x13]	0,0,		&#34;FXCHD	F3,F0&#34;,
[0x14]	0,0,		&#34;FXCHD	F4,F0&#34;,
[0x15]	0,0,		&#34;FXCHD	F5,F0&#34;,
[0x16]	0,0,		&#34;FXCHD	F6,F0&#34;,
[0x17]	0,0,		&#34;FXCHD	F7,F0&#34;,
[0x18]	0,0,		&#34;FNOP&#34;,
[0x28]	0,0,		&#34;FCHS&#34;,
[0x29]	0,0,		&#34;FABS&#34;,
[0x2c]	0,0,		&#34;FTST&#34;,
[0x2d]	0,0,		&#34;FXAM&#34;,
[0x30]	0,0,		&#34;FLD1&#34;,
[0x31]	0,0,		&#34;FLDL2T&#34;,
[0x32]	0,0,		&#34;FLDL2E&#34;,
[0x33]	0,0,		&#34;FLDPI&#34;,
[0x34]	0,0,		&#34;FLDLG2&#34;,
[0x35]	0,0,		&#34;FLDLN2&#34;,
[0x36]	0,0,		&#34;FLDZ&#34;,
[0x38]	0,0,		&#34;F2XM1&#34;,
[0x39]	0,0,		&#34;FYL2X&#34;,
[0x3a]	0,0,		&#34;FPTAN&#34;,
[0x3b]	0,0,		&#34;FPATAN&#34;,
[0x3c]	0,0,		&#34;FXTRACT&#34;,
[0x3d]	0,0,		&#34;FPREM1&#34;,
[0x3e]	0,0,		&#34;FDECSTP&#34;,
[0x3f]	0,0,		&#34;FNCSTP&#34;,
[0x40]	0,0,		&#34;FPREM&#34;,
[0x41]	0,0,		&#34;FYL2XP1&#34;,
[0x42]	0,0,		&#34;FSQRT&#34;,
[0x43]	0,0,		&#34;FSINCOS&#34;,
[0x44]	0,0,		&#34;FRNDINT&#34;,
[0x45]	0,0,		&#34;FSCALE&#34;,
[0x46]	0,0,		&#34;FSIN&#34;,
[0x47]	0,0,		&#34;FCOS&#34;,
};

static Optable optabDA[8+8] =
{
[0x00]	0,0,		&#34;FADDL	%e,F0&#34;,
[0x01]	0,0,		&#34;FMULL	%e,F0&#34;,
[0x02]	0,0,		&#34;FCOML	%e,F0&#34;,
[0x03]	0,0,		&#34;FCOMLP	%e,F0&#34;,
[0x04]	0,0,		&#34;FSUBL	%e,F0&#34;,
[0x05]	0,0,		&#34;FSUBRL	%e,F0&#34;,
[0x06]	0,0,		&#34;FDIVL	%e,F0&#34;,
[0x07]	0,0,		&#34;FDIVRL	%e,F0&#34;,
[0x0d]	R1,0,		&#34;FUCOMPP&#34;,
};

static Optable optabDB[8+64] =
{
[0x00]	0,0,		&#34;FMOVL	%e,F0&#34;,
[0x02]	0,0,		&#34;FMOVL	F0,%e&#34;,
[0x03]	0,0,		&#34;FMOVLP	F0,%e&#34;,
[0x05]	0,0,		&#34;FMOVX	%e,F0&#34;,
[0x07]	0,0,		&#34;FMOVXP	F0,%e&#34;,
[0x2a]	0,0,		&#34;FCLEX&#34;,
[0x2b]	0,0,		&#34;FINIT&#34;,
};

static Optable optabDC[8+8] =
{
[0x00]	0,0,		&#34;FADDD	%e,F0&#34;,
[0x01]	0,0,		&#34;FMULD	%e,F0&#34;,
[0x02]	0,0,		&#34;FCOMD	%e,F0&#34;,
[0x03]	0,0,		&#34;FCOMDP	%e,F0&#34;,
[0x04]	0,0,		&#34;FSUBD	%e,F0&#34;,
[0x05]	0,0,		&#34;FSUBRD	%e,F0&#34;,
[0x06]	0,0,		&#34;FDIVD	%e,F0&#34;,
[0x07]	0,0,		&#34;FDIVRD	%e,F0&#34;,
[0x08]	0,0,		&#34;FADDD	F0,%f&#34;,
[0x09]	0,0,		&#34;FMULD	F0,%f&#34;,
[0x0c]	0,0,		&#34;FSUBRD	F0,%f&#34;,
[0x0d]	0,0,		&#34;FSUBD	F0,%f&#34;,
[0x0e]	0,0,		&#34;FDIVRD	F0,%f&#34;,
[0x0f]	0,0,		&#34;FDIVD	F0,%f&#34;,
};

static Optable optabDD[8+8] =
{
[0x00]	0,0,		&#34;FMOVD	%e,F0&#34;,
[0x02]	0,0,		&#34;FMOVD	F0,%e&#34;,
[0x03]	0,0,		&#34;FMOVDP	F0,%e&#34;,
[0x04]	0,0,		&#34;FRSTOR%S %e&#34;,
[0x06]	0,0,		&#34;FSAVE%S %e&#34;,
[0x07]	0,0,		&#34;FSTSW	%e&#34;,
[0x08]	0,0,		&#34;FFREED	%f&#34;,
[0x0a]	0,0,		&#34;FMOVD	%f,F0&#34;,
[0x0b]	0,0,		&#34;FMOVDP	%f,F0&#34;,
[0x0c]	0,0,		&#34;FUCOMD	%f,F0&#34;,
[0x0d]	0,0,		&#34;FUCOMDP %f,F0&#34;,
};

static Optable optabDE[8+8] =
{
[0x00]	0,0,		&#34;FADDW	%e,F0&#34;,
[0x01]	0,0,		&#34;FMULW	%e,F0&#34;,
[0x02]	0,0,		&#34;FCOMW	%e,F0&#34;,
[0x03]	0,0,		&#34;FCOMWP	%e,F0&#34;,
[0x04]	0,0,		&#34;FSUBW	%e,F0&#34;,
[0x05]	0,0,		&#34;FSUBRW	%e,F0&#34;,
[0x06]	0,0,		&#34;FDIVW	%e,F0&#34;,
[0x07]	0,0,		&#34;FDIVRW	%e,F0&#34;,
[0x08]	0,0,		&#34;FADDDP	F0,%f&#34;,
[0x09]	0,0,		&#34;FMULDP	F0,%f&#34;,
[0x0b]	R1,0,		&#34;FCOMPDP&#34;,
[0x0c]	0,0,		&#34;FSUBRDP F0,%f&#34;,
[0x0d]	0,0,		&#34;FSUBDP	F0,%f&#34;,
[0x0e]	0,0,		&#34;FDIVRDP F0,%f&#34;,
[0x0f]	0,0,		&#34;FDIVDP	F0,%f&#34;,
};

static Optable optabDF[8+8] =
{
[0x00]	0,0,		&#34;FMOVW	%e,F0&#34;,
[0x02]	0,0,		&#34;FMOVW	F0,%e&#34;,
[0x03]	0,0,		&#34;FMOVWP	F0,%e&#34;,
[0x04]	0,0,		&#34;FBLD	%e&#34;,
[0x05]	0,0,		&#34;FMOVL	%e,F0&#34;,
[0x06]	0,0,		&#34;FBSTP	%e&#34;,
[0x07]	0,0,		&#34;FMOVLP	F0,%e&#34;,
[0x0c]	R0,0,		&#34;FSTSW	%OAX&#34;,
};

static Optable optabF6[8] =
{
[0x00]	Ib,0,		&#34;TESTB	%i,%e&#34;,
[0x02]	0,0,		&#34;NOTB	%e&#34;,
[0x03]	0,0,		&#34;NEGB	%e&#34;,
[0x04]	0,0,		&#34;MULB	AL,%e&#34;,
[0x05]	0,0,		&#34;IMULB	AL,%e&#34;,
[0x06]	0,0,		&#34;DIVB	AL,%e&#34;,
[0x07]	0,0,		&#34;IDIVB	AL,%e&#34;,
};

static Optable optabF7[8] =
{
[0x00]	Iwd,0,		&#34;TEST%S	%i,%e&#34;,
[0x02]	0,0,		&#34;NOT%S	%e&#34;,
[0x03]	0,0,		&#34;NEG%S	%e&#34;,
[0x04]	0,0,		&#34;MUL%S	%OAX,%e&#34;,
[0x05]	0,0,		&#34;IMUL%S	%OAX,%e&#34;,
[0x06]	0,0,		&#34;DIV%S	%OAX,%e&#34;,
[0x07]	0,0,		&#34;IDIV%S	%OAX,%e&#34;,
};

static Optable optabFE[8] =
{
[0x00]	0,0,		&#34;INCB	%e&#34;,
[0x01]	0,0,		&#34;DECB	%e&#34;,
};

static Optable optabFF[8] =
{
[0x00]	0,0,		&#34;INC%S	%e&#34;,
[0x01]	0,0,		&#34;DEC%S	%e&#34;,
[0x02]	JUMP,0,		&#34;CALL*	%e&#34;,
[0x03]	JUMP,0,		&#34;CALLF*	%e&#34;,
[0x04]	JUMP,0,		&#34;JMP*	%e&#34;,
[0x05]	JUMP,0,		&#34;JMPF*	%e&#34;,
[0x06]	0,0,		&#34;PUSHL	%e&#34;,
};

static Optable optable[256+1] =
{
[0x00]	RMB,0,		&#34;ADDB	%r,%e&#34;,
[0x01]	RM,0,		&#34;ADD%S	%r,%e&#34;,
[0x02]	RMB,0,		&#34;ADDB	%e,%r&#34;,
[0x03]	RM,0,		&#34;ADD%S	%e,%r&#34;,
[0x04]	Ib,0,		&#34;ADDB	%i,AL&#34;,
[0x05]	Iwd,0,		&#34;ADD%S	%i,%OAX&#34;,
[0x06]	0,0,		&#34;PUSHL	ES&#34;,
[0x07]	0,0,		&#34;POPL	ES&#34;,
[0x08]	RMB,0,		&#34;ORB	%r,%e&#34;,
[0x09]	RM,0,		&#34;OR%S	%r,%e&#34;,
[0x0a]	RMB,0,		&#34;ORB	%e,%r&#34;,
[0x0b]	RM,0,		&#34;OR%S	%e,%r&#34;,
[0x0c]	Ib,0,		&#34;ORB	%i,AL&#34;,
[0x0d]	Iwd,0,		&#34;OR%S	%i,%OAX&#34;,
[0x0e]	0,0,		&#34;PUSHL	CS&#34;,
[0x0f]	AUXMM,0,	optab0F,
[0x10]	RMB,0,		&#34;ADCB	%r,%e&#34;,
[0x11]	RM,0,		&#34;ADC%S	%r,%e&#34;,
[0x12]	RMB,0,		&#34;ADCB	%e,%r&#34;,
[0x13]	RM,0,		&#34;ADC%S	%e,%r&#34;,
[0x14]	Ib,0,		&#34;ADCB	%i,AL&#34;,
[0x15]	Iwd,0,		&#34;ADC%S	%i,%OAX&#34;,
[0x16]	0,0,		&#34;PUSHL	SS&#34;,
[0x17]	0,0,		&#34;POPL	SS&#34;,
[0x18]	RMB,0,		&#34;SBBB	%r,%e&#34;,
[0x19]	RM,0,		&#34;SBB%S	%r,%e&#34;,
[0x1a]	RMB,0,		&#34;SBBB	%e,%r&#34;,
[0x1b]	RM,0,		&#34;SBB%S	%e,%r&#34;,
[0x1c]	Ib,0,		&#34;SBBB	%i,AL&#34;,
[0x1d]	Iwd,0,		&#34;SBB%S	%i,%OAX&#34;,
[0x1e]	0,0,		&#34;PUSHL	DS&#34;,
[0x1f]	0,0,		&#34;POPL	DS&#34;,
[0x20]	RMB,0,		&#34;ANDB	%r,%e&#34;,
[0x21]	RM,0,		&#34;AND%S	%r,%e&#34;,
[0x22]	RMB,0,		&#34;ANDB	%e,%r&#34;,
[0x23]	RM,0,		&#34;AND%S	%e,%r&#34;,
[0x24]	Ib,0,		&#34;ANDB	%i,AL&#34;,
[0x25]	Iwd,0,		&#34;AND%S	%i,%OAX&#34;,
[0x26]	SEG,0,		&#34;ES:&#34;,
[0x27]	0,0,		&#34;DAA&#34;,
[0x28]	RMB,0,		&#34;SUBB	%r,%e&#34;,
[0x29]	RM,0,		&#34;SUB%S	%r,%e&#34;,
[0x2a]	RMB,0,		&#34;SUBB	%e,%r&#34;,
[0x2b]	RM,0,		&#34;SUB%S	%e,%r&#34;,
[0x2c]	Ib,0,		&#34;SUBB	%i,AL&#34;,
[0x2d]	Iwd,0,		&#34;SUB%S	%i,%OAX&#34;,
[0x2e]	SEG,0,		&#34;CS:&#34;,
[0x2f]	0,0,		&#34;DAS&#34;,
[0x30]	RMB,0,		&#34;XORB	%r,%e&#34;,
[0x31]	RM,0,		&#34;XOR%S	%r,%e&#34;,
[0x32]	RMB,0,		&#34;XORB	%e,%r&#34;,
[0x33]	RM,0,		&#34;XOR%S	%e,%r&#34;,
[0x34]	Ib,0,		&#34;XORB	%i,AL&#34;,
[0x35]	Iwd,0,		&#34;XOR%S	%i,%OAX&#34;,
[0x36]	SEG,0,		&#34;SS:&#34;,
[0x37]	0,0,		&#34;AAA&#34;,
[0x38]	RMB,0,		&#34;CMPB	%r,%e&#34;,
[0x39]	RM,0,		&#34;CMP%S	%r,%e&#34;,
[0x3a]	RMB,0,		&#34;CMPB	%e,%r&#34;,
[0x3b]	RM,0,		&#34;CMP%S	%e,%r&#34;,
[0x3c]	Ib,0,		&#34;CMPB	%i,AL&#34;,
[0x3d]	Iwd,0,		&#34;CMP%S	%i,%OAX&#34;,
[0x3e]	SEG,0,		&#34;DS:&#34;,
[0x3f]	0,0,		&#34;AAS&#34;,
[0x40]	0,0,		&#34;INC%S	%OAX&#34;,
[0x41]	0,0,		&#34;INC%S	%OCX&#34;,
[0x42]	0,0,		&#34;INC%S	%ODX&#34;,
[0x43]	0,0,		&#34;INC%S	%OBX&#34;,
[0x44]	0,0,		&#34;INC%S	%OSP&#34;,
[0x45]	0,0,		&#34;INC%S	%OBP&#34;,
[0x46]	0,0,		&#34;INC%S	%OSI&#34;,
[0x47]	0,0,		&#34;INC%S	%ODI&#34;,
[0x48]	0,0,		&#34;DEC%S	%OAX&#34;,
[0x49]	0,0,		&#34;DEC%S	%OCX&#34;,
[0x4a]	0,0,		&#34;DEC%S	%ODX&#34;,
[0x4b]	0,0,		&#34;DEC%S	%OBX&#34;,
[0x4c]	0,0,		&#34;DEC%S	%OSP&#34;,
[0x4d]	0,0,		&#34;DEC%S	%OBP&#34;,
[0x4e]	0,0,		&#34;DEC%S	%OSI&#34;,
[0x4f]	0,0,		&#34;DEC%S	%ODI&#34;,
[0x50]	0,0,		&#34;PUSH%S	%OAX&#34;,
[0x51]	0,0,		&#34;PUSH%S	%OCX&#34;,
[0x52]	0,0,		&#34;PUSH%S	%ODX&#34;,
[0x53]	0,0,		&#34;PUSH%S	%OBX&#34;,
[0x54]	0,0,		&#34;PUSH%S	%OSP&#34;,
[0x55]	0,0,		&#34;PUSH%S	%OBP&#34;,
[0x56]	0,0,		&#34;PUSH%S	%OSI&#34;,
[0x57]	0,0,		&#34;PUSH%S	%ODI&#34;,
[0x58]	0,0,		&#34;POP%S	%OAX&#34;,
[0x59]	0,0,		&#34;POP%S	%OCX&#34;,
[0x5a]	0,0,		&#34;POP%S	%ODX&#34;,
[0x5b]	0,0,		&#34;POP%S	%OBX&#34;,
[0x5c]	0,0,		&#34;POP%S	%OSP&#34;,
[0x5d]	0,0,		&#34;POP%S	%OBP&#34;,
[0x5e]	0,0,		&#34;POP%S	%OSI&#34;,
[0x5f]	0,0,		&#34;POP%S	%ODI&#34;,
[0x60]	0,0,		&#34;PUSHA%S&#34;,
[0x61]	0,0,		&#34;POPA%S&#34;,
[0x62]	RMM,0,		&#34;BOUND	%e,%r&#34;,
[0x63]	RM,0,		&#34;ARPL	%r,%e&#34;,
[0x64]	SEG,0,		&#34;FS:&#34;,
[0x65]	SEG,0,		&#34;GS:&#34;,
[0x66]	OPOVER,0,	&#34;&#34;,
[0x67]	ADDOVER,0,	&#34;&#34;,
[0x68]	Iwd,0,		&#34;PUSH%S	%i&#34;,
[0x69]	RM,Iwd,		&#34;IMUL%S	%e,%i,%r&#34;,
[0x6a]	Ib,0,		&#34;PUSH%S	%i&#34;,
[0x6b]	RM,Ibs,		&#34;IMUL%S	%e,%i,%r&#34;,
[0x6c]	0,0,		&#34;INSB	DX,(%ODI)&#34;,
[0x6d]	0,0,		&#34;INS%S	DX,(%ODI)&#34;,
[0x6e]	0,0,		&#34;OUTSB	(%ASI),DX&#34;,
[0x6f]	0,0,		&#34;OUTS%S	(%ASI),DX&#34;,
[0x70]	Jbs,0,		&#34;JOS	%p&#34;,
[0x71]	Jbs,0,		&#34;JOC	%p&#34;,
[0x72]	Jbs,0,		&#34;JCS	%p&#34;,
[0x73]	Jbs,0,		&#34;JCC	%p&#34;,
[0x74]	Jbs,0,		&#34;JEQ	%p&#34;,
[0x75]	Jbs,0,		&#34;JNE	%p&#34;,
[0x76]	Jbs,0,		&#34;JLS	%p&#34;,
[0x77]	Jbs,0,		&#34;JHI	%p&#34;,
[0x78]	Jbs,0,		&#34;JMI	%p&#34;,
[0x79]	Jbs,0,		&#34;JPL	%p&#34;,
[0x7a]	Jbs,0,		&#34;JPS	%p&#34;,
[0x7b]	Jbs,0,		&#34;JPC	%p&#34;,
[0x7c]	Jbs,0,		&#34;JLT	%p&#34;,
[0x7d]	Jbs,0,		&#34;JGE	%p&#34;,
[0x7e]	Jbs,0,		&#34;JLE	%p&#34;,
[0x7f]	Jbs,0,		&#34;JGT	%p&#34;,
[0x80]	RMOPB,0,	optab80,
[0x81]	RMOP,0,		optab81,
[0x83]	RMOP,0,		optab83,
[0x84]	RMB,0,		&#34;TESTB	%r,%e&#34;,
[0x85]	RM,0,		&#34;TEST%S	%r,%e&#34;,
[0x86]	RMB,0,		&#34;XCHGB	%r,%e&#34;,
[0x87]	RM,0,		&#34;XCHG%S	%r,%e&#34;,
[0x88]	RMB,0,		&#34;MOVB	%r,%e&#34;,
[0x89]	RM,0,		&#34;MOV%S	%r,%e&#34;,
[0x8a]	RMB,0,		&#34;MOVB	%e,%r&#34;,
[0x8b]	RM,0,		&#34;MOV%S	%e,%r&#34;,
[0x8c]	RM,0,		&#34;MOVW	%g,%e&#34;,
[0x8d]	RM,0,		&#34;LEA%S	%e,%r&#34;,
[0x8e]	RM,0,		&#34;MOVW	%e,%g&#34;,
[0x8f]	RM,0,		&#34;POP%S	%e&#34;,
[0x90]	0,0,		&#34;NOP&#34;,
[0x91]	0,0,		&#34;XCHG	%OCX,%OAX&#34;,
[0x92]	0,0,		&#34;XCHG	%ODX,%OAX&#34;,
[0x93]	0,0,		&#34;XCHG	%OBX,%OAX&#34;,
[0x94]	0,0,		&#34;XCHG	%OSP,%OAX&#34;,
[0x95]	0,0,		&#34;XCHG	%OBP,%OAX&#34;,
[0x96]	0,0,		&#34;XCHG	%OSI,%OAX&#34;,
[0x97]	0,0,		&#34;XCHG	%ODI,%OAX&#34;,
[0x98]	0,0,		&#34;%W&#34;,			/* miserable CBW or CWDE */
[0x99]	0,0,		&#34;%w&#34;,			/* idiotic CWD or CDQ */
[0x9a]	PTR,0,		&#34;CALL%S	%d&#34;,
[0x9b]	0,0,		&#34;WAIT&#34;,
[0x9c]	0,0,		&#34;PUSHF&#34;,
[0x9d]	0,0,		&#34;POPF&#34;,
[0x9e]	0,0,		&#34;SAHF&#34;,
[0x9f]	0,0,		&#34;LAHF&#34;,
[0xa0]	Awd,0,		&#34;MOVB	%i,AL&#34;,
[0xa1]	Awd,0,		&#34;MOV%S	%i,%OAX&#34;,
[0xa2]	Awd,0,		&#34;MOVB	AL,%i&#34;,
[0xa3]	Awd,0,		&#34;MOV%S	%OAX,%i&#34;,
[0xa4]	0,0,		&#34;MOVSB	(%ASI),(%ADI)&#34;,
[0xa5]	0,0,		&#34;MOVS%S	(%ASI),(%ADI)&#34;,
[0xa6]	0,0,		&#34;CMPSB	(%ASI),(%ADI)&#34;,
[0xa7]	0,0,		&#34;CMPS%S	(%ASI),(%ADI)&#34;,
[0xa8]	Ib,0,		&#34;TESTB	%i,AL&#34;,
[0xa9]	Iwd,0,		&#34;TEST%S	%i,%OAX&#34;,
[0xaa]	0,0,		&#34;STOSB	AL,(%ADI)&#34;,
[0xab]	0,0,		&#34;STOS%S	%OAX,(%ADI)&#34;,
[0xac]	0,0,		&#34;LODSB	(%ASI),AL&#34;,
[0xad]	0,0,		&#34;LODS%S	(%ASI),%OAX&#34;,
[0xae]	0,0,		&#34;SCASB	(%ADI),AL&#34;,
[0xaf]	0,0,		&#34;SCAS%S	(%ADI),%OAX&#34;,
[0xb0]	Ib,0,		&#34;MOVB	%i,AL&#34;,
[0xb1]	Ib,0,		&#34;MOVB	%i,CL&#34;,
[0xb2]	Ib,0,		&#34;MOVB	%i,DL&#34;,
[0xb3]	Ib,0,		&#34;MOVB	%i,BL&#34;,
[0xb4]	Ib,0,		&#34;MOVB	%i,AH&#34;,
[0xb5]	Ib,0,		&#34;MOVB	%i,CH&#34;,
[0xb6]	Ib,0,		&#34;MOVB	%i,DH&#34;,
[0xb7]	Ib,0,		&#34;MOVB	%i,BH&#34;,
[0xb8]	Iwdq,0,		&#34;MOV%S	%i,%OAX&#34;,
[0xb9]	Iwdq,0,		&#34;MOV%S	%i,%OCX&#34;,
[0xba]	Iwdq,0,		&#34;MOV%S	%i,%ODX&#34;,
[0xbb]	Iwdq,0,		&#34;MOV%S	%i,%OBX&#34;,
[0xbc]	Iwdq,0,		&#34;MOV%S	%i,%OSP&#34;,
[0xbd]	Iwdq,0,		&#34;MOV%S	%i,%OBP&#34;,
[0xbe]	Iwdq,0,		&#34;MOV%S	%i,%OSI&#34;,
[0xbf]	Iwdq,0,		&#34;MOV%S	%i,%ODI&#34;,
[0xc0]	RMOPB,0,	optabC0,
[0xc1]	RMOP,0,		optabC1,
[0xc2]	Iw,0,		&#34;RET	%i&#34;,
[0xc3]	RET,0,		&#34;RET&#34;,
[0xc4]	RM,0,		&#34;LES	%e,%r&#34;,
[0xc5]	RM,0,		&#34;LDS	%e,%r&#34;,
[0xc6]	RMB,Ib,		&#34;MOVB	%i,%e&#34;,
[0xc7]	RM,Iwd,		&#34;MOV%S	%i,%e&#34;,
[0xc8]	Iw2,Ib,		&#34;ENTER	%i,%I&#34;,		/* loony ENTER */
[0xc9]	RET,0,		&#34;LEAVE&#34;,		/* bizarre LEAVE */
[0xca]	Iw,0,		&#34;RETF	%i&#34;,
[0xcb]	RET,0,		&#34;RETF&#34;,
[0xcc]	0,0,		&#34;INT	3&#34;,
[0xcd]	Ib,0,		&#34;INTB	%i&#34;,
[0xce]	0,0,		&#34;INTO&#34;,
[0xcf]	0,0,		&#34;IRET&#34;,
[0xd0]	RMOPB,0,	optabD0,
[0xd1]	RMOP,0,		optabD1,
[0xd2]	RMOPB,0,	optabD2,
[0xd3]	RMOP,0,		optabD3,
[0xd4]	OA,0,		&#34;AAM&#34;,
[0xd5]	OA,0,		&#34;AAD&#34;,
[0xd7]	0,0,		&#34;XLAT&#34;,
[0xd8]	FRMOP,0,	optabD8,
[0xd9]	FRMEX,0,	optabD9,
[0xda]	FRMOP,0,	optabDA,
[0xdb]	FRMEX,0,	optabDB,
[0xdc]	FRMOP,0,	optabDC,
[0xdd]	FRMOP,0,	optabDD,
[0xde]	FRMOP,0,	optabDE,
[0xdf]	FRMOP,0,	optabDF,
[0xe0]	Jbs,0,		&#34;LOOPNE	%p&#34;,
[0xe1]	Jbs,0,		&#34;LOOPE	%p&#34;,
[0xe2]	Jbs,0,		&#34;LOOP	%p&#34;,
[0xe3]	Jbs,0,		&#34;JCXZ	%p&#34;,
[0xe4]	Ib,0,		&#34;INB	%i,AL&#34;,
[0xe5]	Ib,0,		&#34;IN%S	%i,%OAX&#34;,
[0xe6]	Ib,0,		&#34;OUTB	AL,%i&#34;,
[0xe7]	Ib,0,		&#34;OUT%S	%OAX,%i&#34;,
[0xe8]	Iwds,0,		&#34;CALL	%p&#34;,
[0xe9]	Iwds,0,		&#34;JMP	%p&#34;,
[0xea]	PTR,0,		&#34;JMP	%d&#34;,
[0xeb]	Jbs,0,		&#34;JMP	%p&#34;,
[0xec]	0,0,		&#34;INB	DX,AL&#34;,
[0xed]	0,0,		&#34;IN%S	DX,%OAX&#34;,
[0xee]	0,0,		&#34;OUTB	AL,DX&#34;,
[0xef]	0,0,		&#34;OUT%S	%OAX,DX&#34;,
[0xf0]	PRE,0,		&#34;LOCK&#34;,
[0xf2]	OPRE,0,		&#34;REPNE&#34;,
[0xf3]	OPRE,0,		&#34;REP&#34;,
[0xf4]	0,0,		&#34;HLT&#34;,
[0xf5]	0,0,		&#34;CMC&#34;,
[0xf6]	RMOPB,0,	optabF6,
[0xf7]	RMOP,0,		optabF7,
[0xf8]	0,0,		&#34;CLC&#34;,
[0xf9]	0,0,		&#34;STC&#34;,
[0xfa]	0,0,		&#34;CLI&#34;,
[0xfb]	0,0,		&#34;STI&#34;,
[0xfc]	0,0,		&#34;CLD&#34;,
[0xfd]	0,0,		&#34;STD&#34;,
[0xfe]	RMOPB,0,	optabFE,
[0xff]	RMOP,0,		optabFF,
[0x100]	RM,0,		&#34;MOVLQSX	%r,%e&#34;,
};

/*
 *  get a byte of the instruction
 */
static int
igetc(Map *map, Instr *ip, uchar *c)
{
	if(ip-&gt;n+1 &gt; sizeof(ip-&gt;mem)){
		werrstr(&#34;instruction too long&#34;);
		return -1;
	}
	if (get1(map, ip-&gt;addr+ip-&gt;n, c, 1) &lt; 0) {
		werrstr(&#34;can&#39;t read instruction: %r&#34;);
		return -1;
	}
	ip-&gt;mem[ip-&gt;n++] = *c;
	return 1;
}

/*
 *  get two bytes of the instruction
 */
static int
igets(Map *map, Instr *ip, ushort *sp)
{
	uchar c;
	ushort s;

	if (igetc(map, ip, &amp;c) &lt; 0)
		return -1;
	s = c;
	if (igetc(map, ip, &amp;c) &lt; 0)
		return -1;
	s |= (c&lt;&lt;8);
	*sp = s;
	return 1;
}

/*
 *  get 4 bytes of the instruction
 */
static int
igetl(Map *map, Instr *ip, uint32 *lp)
{
	ushort s;
	int32	l;

	if (igets(map, ip, &amp;s) &lt; 0)
		return -1;
	l = s;
	if (igets(map, ip, &amp;s) &lt; 0)
		return -1;
	l |= (s&lt;&lt;16);
	*lp = l;
	return 1;
}

/*
 *  get 8 bytes of the instruction
 *
static int
igetq(Map *map, Instr *ip, vlong *qp)
{
	uint32	l;
	uvlong q;

	if (igetl(map, ip, &amp;l) &lt; 0)
		return -1;
	q = l;
	if (igetl(map, ip, &amp;l) &lt; 0)
		return -1;
	q |= ((uvlong)l&lt;&lt;32);
	*qp = q;
	return 1;
}
 */

static int
getdisp(Map *map, Instr *ip, int mod, int rm, int code, int pcrel)
{
	uchar c;
	ushort s;

	if (mod &gt; 2)
		return 1;
	if (mod == 1) {
		if (igetc(map, ip, &amp;c) &lt; 0)
			return -1;
		if (c&amp;0x80)
			ip-&gt;disp = c|0xffffff00;
		else
			ip-&gt;disp = c&amp;0xff;
	} else if (mod == 2 || rm == code) {
		if (ip-&gt;asize == &#39;E&#39;) {
			if (igetl(map, ip, &amp;ip-&gt;disp) &lt; 0)
				return -1;
			if (mod == 0)
				ip-&gt;rip = pcrel;
		} else {
			if (igets(map, ip, &amp;s) &lt; 0)
				return -1;
			if (s&amp;0x8000)
				ip-&gt;disp = s|0xffff0000;
			else
				ip-&gt;disp = s;
		}
		if (mod == 0)
			ip-&gt;base = -1;
	}
	return 1;
}

static int
modrm(Map *map, Instr *ip, uchar c)
{
	uchar rm, mod;

	mod = (c&gt;&gt;6)&amp;3;
	rm = c&amp;7;
	ip-&gt;mod = mod;
	ip-&gt;base = rm;
	ip-&gt;reg = (c&gt;&gt;3)&amp;7;
	ip-&gt;rip = 0;
	if (mod == 3)			/* register */
		return 1;
	if (ip-&gt;asize == 0) {		/* 16-bit mode */
		switch(rm) {
		case 0:
			ip-&gt;base = BX; ip-&gt;index = SI;
			break;
		case 1:
			ip-&gt;base = BX; ip-&gt;index = DI;
			break;
		case 2:
			ip-&gt;base = BP; ip-&gt;index = SI;
			break;
		case 3:
			ip-&gt;base = BP; ip-&gt;index = DI;
			break;
		case 4:
			ip-&gt;base = SI;
			break;
		case 5:
			ip-&gt;base = DI;
			break;
		case 6:
			ip-&gt;base = BP;
			break;
		case 7:
			ip-&gt;base = BX;
			break;
		default:
			break;
		}
		return getdisp(map, ip, mod, rm, 6, 0);
	}
	if (rm == 4) {	/* scummy sib byte */
		if (igetc(map, ip, &amp;c) &lt; 0)
			return -1;
		ip-&gt;ss = (c&gt;&gt;6)&amp;0x03;
		ip-&gt;index = (c&gt;&gt;3)&amp;0x07;
		if (ip-&gt;index == 4)
			ip-&gt;index = -1;
		ip-&gt;base = c&amp;0x07;
		return getdisp(map, ip, mod, ip-&gt;base, 5, 0);
	}
	return getdisp(map, ip, mod, rm, 5, ip-&gt;amd64);
}

static Optable *
mkinstr(Map *map, Instr *ip, uvlong pc)
{
	int i, n, norex;
	uchar c;
	ushort s;
	Optable *op, *obase;
	char buf[128];

	memset(ip, 0, sizeof(*ip));
	norex = 1;
	ip-&gt;base = -1;
	ip-&gt;index = -1;
	if(asstype == AI8086)
		ip-&gt;osize = &#39;W&#39;;
	else {
		ip-&gt;osize = &#39;L&#39;;
		ip-&gt;asize = &#39;E&#39;;
		ip-&gt;amd64 = asstype != AI386;
		norex = 0;
	}
	ip-&gt;addr = pc;
	if (igetc(map, ip, &amp;c) &lt; 0)
		return 0;
	obase = optable;
newop:
	if(ip-&gt;amd64 &amp;&amp; !norex){
		if(c &gt;= 0x40 &amp;&amp; c &lt;= 0x4f) {
			ip-&gt;rex = c;
			if(igetc(map, ip, &amp;c) &lt; 0)
				return 0;
		}
		if(c == 0x63){
			op = &amp;obase[0x100];	/* MOVLQSX */
			goto hack;
		}
	}
	op = &amp;obase[c];
hack:
	if (op-&gt;proto == 0) {
badop:
		n = snprint(buf, sizeof(buf), &#34;opcode: ??&#34;);
		for (i = 0; i &lt; ip-&gt;n &amp;&amp; n &lt; sizeof(buf)-3; i++, n+=2)
			_hexify(buf+n, ip-&gt;mem[i], 1);
		strcpy(buf+n, &#34;??&#34;);
		werrstr(buf);
		return 0;
	}
	for(i = 0; i &lt; 2 &amp;&amp; op-&gt;operand[i]; i++) {
		switch(op-&gt;operand[i]) {
		case Ib:	/* 8-bit immediate - (no sign extension)*/
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			ip-&gt;imm = c&amp;0xff;
			ip-&gt;imm64 = ip-&gt;imm;
			break;
		case Jbs:	/* 8-bit jump immediate (sign extended) */
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if (c&amp;0x80)
				ip-&gt;imm = c|0xffffff00;
			else
				ip-&gt;imm = c&amp;0xff;
			ip-&gt;imm64 = (int32)ip-&gt;imm;
			ip-&gt;jumptype = Jbs;
			break;
		case Ibs:	/* 8-bit immediate (sign extended) */
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if (c&amp;0x80)
				if (ip-&gt;osize == &#39;L&#39;)
					ip-&gt;imm = c|0xffffff00;
				else
					ip-&gt;imm = c|0xff00;
			else
				ip-&gt;imm = c&amp;0xff;
			ip-&gt;imm64 = (int32)ip-&gt;imm;
			break;
		case Iw:	/* 16-bit immediate -&gt; imm */
			if (igets(map, ip, &amp;s) &lt; 0)
				return 0;
			ip-&gt;imm = s&amp;0xffff;
			ip-&gt;imm64 = ip-&gt;imm;
			ip-&gt;jumptype = Iw;
			break;
		case Iw2:	/* 16-bit immediate -&gt; in imm2*/
			if (igets(map, ip, &amp;s) &lt; 0)
				return 0;
			ip-&gt;imm2 = s&amp;0xffff;
			break;
		case Iwd:	/* Operand-sized immediate (no sign extension unless 64 bits)*/
			if (ip-&gt;osize == &#39;L&#39;) {
				if (igetl(map, ip, &amp;ip-&gt;imm) &lt; 0)
					return 0;
				ip-&gt;imm64 = ip-&gt;imm;
				if(ip-&gt;rex&amp;REXW &amp;&amp; (ip-&gt;imm &amp; (1&lt;&lt;31)) != 0)
					ip-&gt;imm64 |= (vlong)~0 &lt;&lt; 32;
			} else {
				if (igets(map, ip, &amp;s)&lt; 0)
					return 0;
				ip-&gt;imm = s&amp;0xffff;
				ip-&gt;imm64 = ip-&gt;imm;
			}
			break;
		case Iwdq:	/* Operand-sized immediate, possibly big */
			if (ip-&gt;osize == &#39;L&#39;) {
				if (igetl(map, ip, &amp;ip-&gt;imm) &lt; 0)
					return 0;
				ip-&gt;imm64 = ip-&gt;imm;
				if (ip-&gt;rex &amp; REXW) {
					uint32 l;
					if (igetl(map, ip, &amp;l) &lt; 0)
						return 0;
					ip-&gt;imm64 |= (uvlong)l &lt;&lt; 32;
				}
			} else {
				if (igets(map, ip, &amp;s)&lt; 0)
					return 0;
				ip-&gt;imm = s&amp;0xffff;
			}
			break;
		case Awd:	/* Address-sized immediate (no sign extension)*/
			if (ip-&gt;asize == &#39;E&#39;) {
				if (igetl(map, ip, &amp;ip-&gt;imm) &lt; 0)
					return 0;
				/* TO DO: REX */
			} else {
				if (igets(map, ip, &amp;s)&lt; 0)
					return 0;
				ip-&gt;imm = s&amp;0xffff;
			}
			break;
		case Iwds:	/* Operand-sized immediate (sign extended) */
			if (ip-&gt;osize == &#39;L&#39;) {
				if (igetl(map, ip, &amp;ip-&gt;imm) &lt; 0)
					return 0;
			} else {
				if (igets(map, ip, &amp;s)&lt; 0)
					return 0;
				if (s&amp;0x8000)
					ip-&gt;imm = s|0xffff0000;
				else
					ip-&gt;imm = s&amp;0xffff;
			}
			ip-&gt;jumptype = Iwds;
			break;
		case OA:	/* literal 0x0a byte */
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if (c != 0x0a)
				goto badop;
			break;
		case R0:	/* base register must be R0 */
			if (ip-&gt;base != 0)
				goto badop;
			break;
		case R1:	/* base register must be R1 */
			if (ip-&gt;base != 1)
				goto badop;
			break;
		case RMB:	/* R/M field with byte register (/r)*/
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if (modrm(map, ip, c) &lt; 0)
				return 0;
			ip-&gt;osize = &#39;B&#39;;
			break;
		case RM:	/* R/M field with register (/r) */
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if (modrm(map, ip, c) &lt; 0)
				return 0;
			break;
		case RMOPB:	/* R/M field with op code (/digit) */
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if (modrm(map, ip, c) &lt; 0)
				return 0;
			c = ip-&gt;reg;		/* secondary op code */
			obase = (Optable*)op-&gt;proto;
			ip-&gt;osize = &#39;B&#39;;
			goto newop;
		case RMOP:	/* R/M field with op code (/digit) */
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if (modrm(map, ip, c) &lt; 0)
				return 0;
			obase = (Optable*)op-&gt;proto;
			if(ip-&gt;amd64 &amp;&amp; obase == optab0F01 &amp;&amp; c == 0xF8)
				return optab0F01F8;
			c = ip-&gt;reg;
			goto newop;
		case FRMOP:	/* FP R/M field with op code (/digit) */
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if (modrm(map, ip, c) &lt; 0)
				return 0;
			if ((c&amp;0xc0) == 0xc0)
				c = ip-&gt;reg+8;		/* 16 entry table */
			else
				c = ip-&gt;reg;
			obase = (Optable*)op-&gt;proto;
			goto newop;
		case FRMEX:	/* Extended FP R/M field with op code (/digit) */
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if (modrm(map, ip, c) &lt; 0)
				return 0;
			if ((c&amp;0xc0) == 0xc0)
				c = (c&amp;0x3f)+8;		/* 64-entry table */
			else
				c = ip-&gt;reg;
			obase = (Optable*)op-&gt;proto;
			goto newop;
		case RMR:	/* R/M register only (mod = 11) */
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if ((c&amp;0xc0) != 0xc0) {
				werrstr(&#34;invalid R/M register: %x&#34;, c);
				return 0;
			}
			if (modrm(map, ip, c) &lt; 0)
				return 0;
			break;
		case RMM:	/* R/M register only (mod = 11) */
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if ((c&amp;0xc0) == 0xc0) {
				werrstr(&#34;invalid R/M memory mode: %x&#34;, c);
				return 0;
			}
			if (modrm(map, ip, c) &lt; 0)
				return 0;
			break;
		case PTR:	/* Seg:Displacement addr (ptr16:16 or ptr16:32) */
			if (ip-&gt;osize == &#39;L&#39;) {
				if (igetl(map, ip, &amp;ip-&gt;disp) &lt; 0)
					return 0;
			} else {
				if (igets(map, ip, &amp;s)&lt; 0)
					return 0;
				ip-&gt;disp = s&amp;0xffff;
			}
			if (igets(map, ip, (ushort*)&amp;ip-&gt;seg) &lt; 0)
				return 0;
			ip-&gt;jumptype = PTR;
			break;
		case AUXMM:	/* Multi-byte op code; prefix determines table selection */
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			obase = (Optable*)op-&gt;proto;
			switch (ip-&gt;opre) {
			case 0x66:	op = optab660F; break;
			case 0xF2:	op = optabF20F; break;
			case 0xF3:	op = optabF30F; break;
			default:	op = nil; break;
			}
			if(op != nil &amp;&amp; op[c].proto != nil)
				obase = op;
			norex = 1;	/* no more rex prefixes */
			/* otherwise the optab entry captures it */
			goto newop;
		case AUX:	/* Multi-byte op code - Auxiliary table */
			obase = (Optable*)op-&gt;proto;
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			goto newop;
		case OPRE:	/* Instr Prefix or media op */
			ip-&gt;opre = c;
			/* fall through */
		case PRE:	/* Instr Prefix */
			ip-&gt;prefix = (char*)op-&gt;proto;
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if (ip-&gt;opre &amp;&amp; c == 0x0F)
				ip-&gt;prefix = 0;
			goto newop;
		case SEG:	/* Segment Prefix */
			ip-&gt;segment = (char*)op-&gt;proto;
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			goto newop;
		case OPOVER:	/* Operand size override */
			ip-&gt;opre = c;
			ip-&gt;osize = &#39;W&#39;;
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			if (c == 0x0F)
				ip-&gt;osize = &#39;L&#39;;
			else if (ip-&gt;amd64 &amp;&amp; (c&amp;0xF0) == 0x40)
				ip-&gt;osize = &#39;Q&#39;;
			goto newop;
		case ADDOVER:	/* Address size override */
			ip-&gt;asize = 0;
			if (igetc(map, ip, &amp;c) &lt; 0)
				return 0;
			goto newop;
		case JUMP:	/* mark instruction as JUMP or RET */
		case RET:
			ip-&gt;jumptype = op-&gt;operand[i];
			break;
		default:
			werrstr(&#34;bad operand type %d&#34;, op-&gt;operand[i]);
			return 0;
		}
	}
	return op;
}

#pragma	varargck	argpos	bprint		2

static void
bprint(Instr *ip, char *fmt, ...)
{
	va_list arg;

	va_start(arg, fmt);
	ip-&gt;curr = vseprint(ip-&gt;curr, ip-&gt;end, fmt, arg);
	va_end(arg);
}

/*
 *  if we want to call 16 bit regs AX,BX,CX,...
 *  and 32 bit regs EAX,EBX,ECX,... then
 *  change the defs of ANAME and ONAME to:
 *  #define	ANAME(ip)	((ip-&gt;asize == &#39;E&#39; ? &#34;E&#34; : &#34;&#34;)
 *  #define	ONAME(ip)	((ip)-&gt;osize == &#39;L&#39; ? &#34;E&#34; : &#34;&#34;)
 */
#define	ANAME(ip)	&#34;&#34;
#define	ONAME(ip)	&#34;&#34;

static char *reg[] =  {
[AX]	&#34;AX&#34;,
[CX]	&#34;CX&#34;,
[DX]	&#34;DX&#34;,
[BX]	&#34;BX&#34;,
[SP]	&#34;SP&#34;,
[BP]	&#34;BP&#34;,
[SI]	&#34;SI&#34;,
[DI]	&#34;DI&#34;,

	/* amd64 */
[R8]	&#34;R8&#34;,
[R9]	&#34;R9&#34;,
[R10]	&#34;R10&#34;,
[R11]	&#34;R11&#34;,
[R12]	&#34;R12&#34;,
[R13]	&#34;R13&#34;,
[R14]	&#34;R14&#34;,
[R15]	&#34;R15&#34;,
};

static char *breg[] = { &#34;AL&#34;, &#34;CL&#34;, &#34;DL&#34;, &#34;BL&#34;, &#34;AH&#34;, &#34;CH&#34;, &#34;DH&#34;, &#34;BH&#34; };
static char *breg64[] = { &#34;AL&#34;, &#34;CL&#34;, &#34;DL&#34;, &#34;BL&#34;, &#34;SPB&#34;, &#34;BPB&#34;, &#34;SIB&#34;, &#34;DIB&#34;,
	&#34;R8B&#34;, &#34;R9B&#34;, &#34;R10B&#34;, &#34;R11B&#34;, &#34;R12B&#34;, &#34;R13B&#34;, &#34;R14B&#34;, &#34;R15B&#34; };
static char *sreg[] = { &#34;ES&#34;, &#34;CS&#34;, &#34;SS&#34;, &#34;DS&#34;, &#34;FS&#34;, &#34;GS&#34; };

static void
plocal(Instr *ip)
{
	int ret;
	int32 offset;
	Symbol s;
	char *reg;

	offset = ip-&gt;disp;
	if (!findsym(ip-&gt;addr, CTEXT, &amp;s) || !findlocal(&amp;s, FRAMENAME, &amp;s)) {
		bprint(ip, &#34;%lux(SP)&#34;, offset);
		return;
	}

	if (s.value &gt; ip-&gt;disp) {
		ret = getauto(&amp;s, s.value-ip-&gt;disp-mach-&gt;szaddr, CAUTO, &amp;s);
		reg = &#34;(SP)&#34;;
	} else {
		offset -= s.value;
		ret = getauto(&amp;s, offset, CPARAM, &amp;s);
		reg = &#34;(FP)&#34;;
	}
	if (ret)
		bprint(ip, &#34;%s+&#34;, s.name);
	else
		offset = ip-&gt;disp;
	bprint(ip, &#34;%lux%s&#34;, offset, reg);
}

static int
isjmp(Instr *ip)
{
	switch(ip-&gt;jumptype){
	case Iwds:
	case Jbs:
	case JUMP:
		return 1;
	default:
		return 0;
	}
}

/*
 * This is too smart for its own good, but it really is nice
 * to have accurate translations when debugging, and it
 * helps us identify which code is different in binaries that
 * are changed on sources.
 */
static int
issymref(Instr *ip, Symbol *s, int32 w, int32 val)
{
	Symbol next, tmp;
	int32 isstring, size;

	if (isjmp(ip))
		return 1;
	if (s-&gt;class==CTEXT &amp;&amp; w==0)
		return 1;
	if (s-&gt;class==CDATA) {
		/* use first bss symbol (or &#34;end&#34;) rather than edata */
		if (s-&gt;name[0]==&#39;e&#39; &amp;&amp; strcmp(s-&gt;name, &#34;edata&#34;) == 0){
			if((s -&gt;index &gt;= 0 &amp;&amp; globalsym(&amp;tmp, s-&gt;index+1) &amp;&amp; tmp.value==s-&gt;value)
			|| (s-&gt;index &gt; 0 &amp;&amp; globalsym(&amp;tmp, s-&gt;index-1) &amp;&amp; tmp.value==s-&gt;value))
				*s = tmp;
		}
		if (w == 0)
			return 1;
		for (next=*s; next.value==s-&gt;value; next=tmp)
			if (!globalsym(&amp;tmp, next.index+1))
				break;
		size = next.value - s-&gt;value;
		if (w &gt;= size)
			return 0;
		if (w &gt; size-w)
			w = size-w;
		/* huge distances are usually wrong except in .string */
		isstring = (s-&gt;name[0]==&#39;.&#39; &amp;&amp; strcmp(s-&gt;name, &#34;.string&#34;) == 0);
		if (w &gt; 8192 &amp;&amp; !isstring)
			return 0;
		/* medium distances are tricky - look for constants */
		/* near powers of two */
		if ((val&amp;(val-1)) == 0 || (val&amp;(val+1)) == 0)
			return 0;
		return 1;
	}
	return 0;
}

static void
immediate(Instr *ip, vlong val)
{
	Symbol s;
	int32 w;

	if (findsym(val, CANY, &amp;s)) {		/* TO DO */
		w = val - s.value;
		if (w &lt; 0)
			w = -w;
		if (issymref(ip, &amp;s, w, val)) {
			if (w)
				bprint(ip, &#34;%s+%#lux(SB)&#34;, s.name, w);
			else
				bprint(ip, &#34;%s(SB)&#34;, s.name);
			return;
		}
/*
		if (s.class==CDATA &amp;&amp; globalsym(&amp;s, s.index+1)) {
			w = s.value - val;
			if (w &lt; 0)
				w = -w;
			if (w &lt; 4096) {
				bprint(ip, &#34;%s-%#lux(SB)&#34;, s.name, w);
				return;
			}
		}
*/
	}
	if((ip-&gt;rex &amp; REXW) == 0)
		bprint(ip, &#34;%lux&#34;, (long)val);
	else
		bprint(ip, &#34;%llux&#34;, val);
}

static void
pea(Instr *ip)
{
	if (ip-&gt;mod == 3) {
		if (ip-&gt;osize == &#39;B&#39;)
			bprint(ip, (ip-&gt;rex &amp; REXB? breg64: breg)[(uchar)ip-&gt;base]);
		else if(ip-&gt;rex &amp; REXB)
			bprint(ip, &#34;%s%s&#34;, ANAME(ip), reg[ip-&gt;base+8]);
		else
			bprint(ip, &#34;%s%s&#34;, ANAME(ip), reg[(uchar)ip-&gt;base]);
		return;
	}
	if (ip-&gt;segment)
		bprint(ip, ip-&gt;segment);
	if (ip-&gt;asize == &#39;E&#39; &amp;&amp; ip-&gt;base == SP)
		plocal(ip);
	else {
		if (ip-&gt;base &lt; 0)
			immediate(ip, ip-&gt;disp);
		else {
			bprint(ip, &#34;%lux&#34;, ip-&gt;disp);
			if(ip-&gt;rip)
				bprint(ip, &#34;(RIP)&#34;);
			bprint(ip,&#34;(%s%s)&#34;, ANAME(ip), reg[ip-&gt;rex&amp;REXB? ip-&gt;base+8: ip-&gt;base]);
		}
	}
	if (ip-&gt;index &gt;= 0)
		bprint(ip,&#34;(%s%s*%d)&#34;, ANAME(ip), reg[ip-&gt;rex&amp;REXX? ip-&gt;index+8: ip-&gt;index], 1&lt;&lt;ip-&gt;ss);
}

static void
prinstr(Instr *ip, char *fmt)
{
	vlong v;

	if (ip-&gt;prefix)
		bprint(ip, &#34;%s &#34;, ip-&gt;prefix);
	for (; *fmt &amp;&amp; ip-&gt;curr &lt; ip-&gt;end; fmt++) {
		if (*fmt != &#39;%&#39;){
			*ip-&gt;curr++ = *fmt;
			continue;
		}
		switch(*++fmt){
		case &#39;%&#39;:
			*ip-&gt;curr++ = &#39;%&#39;;
			break;
		case &#39;A&#39;:
			bprint(ip, &#34;%s&#34;, ANAME(ip));
			break;
		case &#39;C&#39;:
			bprint(ip, &#34;CR%d&#34;, ip-&gt;reg);
			break;
		case &#39;D&#39;:
			if (ip-&gt;reg &lt; 4 || ip-&gt;reg == 6 || ip-&gt;reg == 7)
				bprint(ip, &#34;DR%d&#34;,ip-&gt;reg);
			else
				bprint(ip, &#34;???&#34;);
			break;
		case &#39;I&#39;:
			bprint(ip, &#34;$&#34;);
			immediate(ip, ip-&gt;imm2);
			break;
		case &#39;O&#39;:
			bprint(ip,&#34;%s&#34;, ONAME(ip));
			break;
		case &#39;i&#39;:
			bprint(ip, &#34;$&#34;);
			v = ip-&gt;imm;
			if(ip-&gt;rex &amp; REXW)
				v = ip-&gt;imm64;
			immediate(ip, v);
			break;
		case &#39;R&#39;:
			bprint(ip, &#34;%s%s&#34;, ONAME(ip), reg[ip-&gt;rex&amp;REXR? ip-&gt;reg+8: ip-&gt;reg]);
			break;
		case &#39;S&#39;:
			if(ip-&gt;osize == &#39;Q&#39; || ip-&gt;osize == &#39;L&#39; &amp;&amp; ip-&gt;rex &amp; REXW)
				bprint(ip, &#34;Q&#34;);
			else
				bprint(ip, &#34;%c&#34;, ip-&gt;osize);
			break;
		case &#39;s&#39;:
			if(ip-&gt;opre == 0 || ip-&gt;opre == 0x66)
				bprint(ip, &#34;P&#34;);
			else
				bprint(ip, &#34;S&#34;);
			if(ip-&gt;opre == 0xf2 || ip-&gt;opre == 0x66)
				bprint(ip, &#34;D&#34;);
			else
				bprint(ip, &#34;S&#34;);
			break;
		case &#39;T&#39;:
			if (ip-&gt;reg == 6 || ip-&gt;reg == 7)
				bprint(ip, &#34;TR%d&#34;,ip-&gt;reg);
			else
				bprint(ip, &#34;???&#34;);
			break;
		case &#39;W&#39;:
			if (ip-&gt;osize == &#39;Q&#39; || ip-&gt;osize == &#39;L&#39; &amp;&amp; ip-&gt;rex &amp; REXW)
				bprint(ip, &#34;CDQE&#34;);
			else if (ip-&gt;osize == &#39;L&#39;)
				bprint(ip,&#34;CWDE&#34;);
			else
				bprint(ip, &#34;CBW&#34;);
			break;
		case &#39;d&#39;:
			bprint(ip,&#34;%ux:%lux&#34;,ip-&gt;seg,ip-&gt;disp);
			break;
		case &#39;m&#39;:
			if (ip-&gt;mod == 3 &amp;&amp; ip-&gt;osize != &#39;B&#39;) {
				if(fmt[1] != &#39;*&#39;){
					if(ip-&gt;opre != 0) {
						bprint(ip, &#34;X%d&#34;, ip-&gt;rex&amp;REXB? ip-&gt;base+8: ip-&gt;base);
						break;
					}
				} else
					fmt++;
				bprint(ip, &#34;M%d&#34;, ip-&gt;base);
				break;
			}
			pea(ip);
			break;
		case &#39;e&#39;:
			pea(ip);
			break;
		case &#39;f&#39;:
			bprint(ip, &#34;F%d&#34;, ip-&gt;base);
			break;
		case &#39;g&#39;:
			if (ip-&gt;reg &lt; 6)
				bprint(ip,&#34;%s&#34;,sreg[ip-&gt;reg]);
			else
				bprint(ip,&#34;???&#34;);
			break;
		case &#39;p&#39;:
			/*
			 * signed immediate in the uint32 ip-&gt;imm.
			 */
			v = (int32)ip-&gt;imm;
			immediate(ip, v+ip-&gt;addr+ip-&gt;n);
			break;
		case &#39;r&#39;:
			if (ip-&gt;osize == &#39;B&#39;)
				bprint(ip,&#34;%s&#34;, (ip-&gt;rex? breg64: breg)[ip-&gt;rex&amp;REXR? ip-&gt;reg+8: ip-&gt;reg]);
			else
				bprint(ip, reg[ip-&gt;rex&amp;REXR? ip-&gt;reg+8: ip-&gt;reg]);
			break;
		case &#39;w&#39;:
			if (ip-&gt;osize == &#39;Q&#39; || ip-&gt;rex &amp; REXW)
				bprint(ip, &#34;CQO&#34;);
			else if (ip-&gt;osize == &#39;L&#39;)
				bprint(ip,&#34;CDQ&#34;);
			else
				bprint(ip, &#34;CWD&#34;);
			break;
		case &#39;M&#39;:
			if(ip-&gt;opre != 0)
				bprint(ip, &#34;X%d&#34;, ip-&gt;rex&amp;REXR? ip-&gt;reg+8: ip-&gt;reg);
			else
				bprint(ip, &#34;M%d&#34;, ip-&gt;reg);
			break;
		case &#39;x&#39;:
			if (ip-&gt;mod == 3 &amp;&amp; ip-&gt;osize != &#39;B&#39;) {
				bprint(ip, &#34;X%d&#34;, ip-&gt;rex&amp;REXB? ip-&gt;base+8: ip-&gt;base);
				break;
			}
			pea(ip);
			break;
		case &#39;X&#39;:
			bprint(ip, &#34;X%d&#34;, ip-&gt;rex&amp;REXR? ip-&gt;reg+8: ip-&gt;reg);
			break;
		default:
			bprint(ip, &#34;%%%c&#34;, *fmt);
			break;
		}
	}
	*ip-&gt;curr = 0;		/* there&#39;s always room for 1 byte */
}

static int
i386inst(Map *map, uvlong pc, char modifier, char *buf, int n)
{
	Instr instr;
	Optable *op;

	USED(modifier);
	op = mkinstr(map, &amp;instr, pc);
	if (op == 0) {
		errstr(buf, n);
		return -1;
	}
	instr.curr = buf;
	instr.end = buf+n-1;
	prinstr(&amp;instr, op-&gt;proto);
	return instr.n;
}

static int
i386das(Map *map, uvlong pc, char *buf, int n)
{
	Instr instr;
	int i;

	if (mkinstr(map, &amp;instr, pc) == 0) {
		errstr(buf, n);
		return -1;
	}
	for(i = 0; i &lt; instr.n &amp;&amp; n &gt; 2; i++) {
		_hexify(buf, instr.mem[i], 1);
		buf += 2;
		n -= 2;
	}
	*buf = 0;
	return instr.n;
}

static int
i386instlen(Map *map, uvlong pc)
{
	Instr i;

	if (mkinstr(map, &amp;i, pc))
		return i.n;
	return -1;
}

static int
i386foll(Map *map, uvlong pc, Rgetter rget, uvlong *foll)
{
	Instr i;
	Optable *op;
	ushort s;
	uvlong l, addr;
	vlong v;
	int n;

	op = mkinstr(map, &amp;i, pc);
	if (!op)
		return -1;

	n = 0;

	switch(i.jumptype) {
	case RET:		/* RETURN or LEAVE */
	case Iw:		/* RETURN */
		if (strcmp(op-&gt;proto, &#34;LEAVE&#34;) == 0) {
			if (geta(map, (*rget)(map, &#34;BP&#34;), &amp;l) &lt; 0)
				return -1;
		} else if (geta(map, (*rget)(map, mach-&gt;sp), &amp;l) &lt; 0)
			return -1;
		foll[0] = l;
		return 1;
	case Iwds:		/* pc relative JUMP or CALL*/
	case Jbs:		/* pc relative JUMP or CALL */
		v = (int32)i.imm;
		foll[0] = pc+v+i.n;
		n = 1;
		break;
	case PTR:		/* seg:displacement JUMP or CALL */
		foll[0] = (i.seg&lt;&lt;4)+i.disp;
		return 1;
	case JUMP:		/* JUMP or CALL EA */

		if(i.mod == 3) {
			foll[0] = (*rget)(map, reg[i.rex&amp;REXB? i.base+8: i.base]);
			return 1;
		}
			/* calculate the effective address */
		addr = i.disp;
		if (i.base &gt;= 0) {
			if (geta(map, (*rget)(map, reg[i.rex&amp;REXB? i.base+8: i.base]), &amp;l) &lt; 0)
				return -1;
			addr += l;
		}
		if (i.index &gt;= 0) {
			if (geta(map, (*rget)(map, reg[i.rex&amp;REXX? i.index+8: i.index]), &amp;l) &lt; 0)
				return -1;
			addr += l*(1&lt;&lt;i.ss);
		}
			/* now retrieve a seg:disp value at that address */
		if (get2(map, addr, &amp;s) &lt; 0)			/* seg */
			return -1;
		foll[0] = s&lt;&lt;4;
		addr += 2;
		if (i.asize == &#39;L&#39;) {
			if (geta(map, addr, &amp;l) &lt; 0)		/* disp32 */
				return -1;
			foll[0] += l;
		} else {					/* disp16 */
			if (get2(map, addr, &amp;s) &lt; 0)
				return -1;
			foll[0] += s;
		}
		return 1;
	default:
		break;
	}
	if (strncmp(op-&gt;proto,&#34;JMP&#34;, 3) == 0 || strncmp(op-&gt;proto,&#34;CALL&#34;, 4) == 0)
		return 1;
	foll[n++] = pc+i.n;
	return n;
}
</pre>

</div>

<div id="footer">
<p>Except as noted, this content is
   licensed under <a href="http://creativecommons.org/licenses/by/3.0/">
   Creative Commons Attribution 3.0</a>.
</div>

<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-11222381-2");
pageTracker._trackPageview();
</script>
</body>
</html>
<!-- generated at Thu Nov 12 15:42:51 PST 2009 -->
