<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5l/l.h</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

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
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5l/l.h</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5l/l.h
// http://code.google.com/p/inferno-os/source/browse/utils/5l/l.h
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
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

#include	&lt;u.h&gt;
#include	&lt;libc.h&gt;
#include	&lt;bio.h&gt;
#include	&#34;../5l/5.out.h&#34;

enum
{
	PtrSize = 4
};

#ifndef	EXTERN
#define	EXTERN	extern
#endif

/* do not undefine this - code will be removed eventually */
#define	CALLEEBX

typedef	struct	Adr	Adr;
typedef	struct	Sym	Sym;
typedef	struct	Autom	Auto;
typedef	struct	Prog	Prog;
typedef	struct	Optab	Optab;
typedef	struct	Oprang	Oprang;
typedef	uchar	Opcross[32][2][32];
typedef	struct	Count	Count;
typedef	struct	Use	Use;

#define	P		((Prog*)0)
#define	S		((Sym*)0)
#define	U		((Use*)0)
#define	TNAME		(curtext&amp;&amp;curtext-&gt;from.sym?curtext-&gt;from.sym-&gt;name:noname)

struct	Adr
{
	union
	{
		int32	u0offset;
		char*	u0sval;
		Ieee*	u0ieee;
		char*	u0sbig;
	} u0;
	union
	{
		Auto*	u1autom;
		Sym*	u1sym;
	} u1;
	char	type;
	uchar	index; // not used on arm, required by ld/go.c
	char	reg;
	char	name;
	int32	offset2; // argsize
	char	class;
	Sym*	gotype;
};

#define	offset	u0.u0offset
#define	sval	u0.u0sval
#define	ieee	u0.u0ieee
#define	sbig	u0.u0sbig

#define	autom	u1.u1autom
#define	sym	u1.u1sym

struct	Prog
{
	Adr	from;
	Adr	to;
	union
	{
		int32	u0regused;
		Prog*	u0forwd;
	} u0;
	Prog*	cond;
	Prog*	link;
	Prog*	dlink;
	int32	pc;
	int32	line;
	uchar	mark;
	uchar	optab;
	uchar	as;
	uchar	scond;
	uchar	reg;
	uchar	align;
};
#define	regused	u0.u0regused
#define	forwd	u0.u0forwd

struct	Sym
{
	char	*name;
	short	type;
	short	version;
	short	become;
	short	frame;
	uchar	subtype;
	uchar	dupok;
	uchar	reachable;
	int32	value;
	int32	sig;
	uchar	used;
	uchar	thumb;	// thumb code
	uchar	foreign;	// called by arm if thumb, by thumb if arm
	uchar	fnptr;	// used as fn ptr
	Use*		use;
	Sym*	link;
	Prog*	text;
	Prog*	data;
	Sym*	gotype;
	char*	file;
	char*	dynldname;
	char*	dynldlib;
};

#define SIGNINTERN	(1729*325*1729)

struct	Autom
{
	Sym*	asym;
	Auto*	link;
	int32	aoffset;
	short	type;
	Sym*	gotype;
};
struct	Optab
{
	char	as;
	uchar	a1;
	char	a2;
	uchar	a3;
	uchar	type;
	char	size;
	char	param;
	char	flag;
};
struct	Oprang
{
	Optab*	start;
	Optab*	stop;
};
struct	Count
{
	int32	count;
	int32	outof;
};
struct	Use
{
	Prog*	p;	/* use */
	Prog*	ct;	/* curtext */
	Use*		link;
};

enum
{
	Sxxx,

	STEXT		= 1,
	SDATA,
	SBSS,
	SDATA1,
	SXREF,
	SLEAF,
	SFILE,
	SCONST,
	SSTRING,
	SUNDEF,
	SREMOVED,

	SIMPORT,
	SEXPORT,

	LFROM		= 1&lt;&lt;0,
	LTO		= 1&lt;&lt;1,
	LPOOL		= 1&lt;&lt;2,
	V4		= 1&lt;&lt;3,	/* arm v4 arch */

	C_NONE		= 0,
	C_REG,
	C_REGREG,
	C_SHIFT,
	C_FREG,
	C_PSR,
	C_FCR,

	C_RCON,		/* 0xff rotated */
	C_NCON,		/* ~RCON */
	C_SCON,		/* 0xffff */
	C_BCON,		/* thumb */
	C_LCON,
	C_FCON,
	C_GCON,		/* thumb */

	C_RACON,
	C_SACON,	/* thumb */
	C_LACON,
	C_GACON,	/* thumb */

	C_RECON,
	C_LECON,

	C_SBRA,
	C_LBRA,
	C_GBRA,		/* thumb */

	C_HAUTO,	/* halfword insn offset (-0xff to 0xff) */
	C_FAUTO,	/* float insn offset (0 to 0x3fc, word aligned) */
	C_HFAUTO,	/* both H and F */
	C_SAUTO,	/* -0xfff to 0xfff */
	C_LAUTO,

	C_HEXT,
	C_FEXT,
	C_HFEXT,
	C_SEXT,
	C_LEXT,

	C_HOREG,
	C_FOREG,
	C_HFOREG,
	C_SOREG,
	C_ROREG,
	C_SROREG,	/* both S and R */
	C_LOREG,
	C_GOREG,		/* thumb */

	C_PC,
	C_SP,
	C_HREG,
	C_OFFPC,		/* thumb */

	C_ADDR,		/* relocatable address */

	C_GOK,

/* mark flags */
	FOLL		= 1&lt;&lt;0,
	LABEL		= 1&lt;&lt;1,
	LEAF		= 1&lt;&lt;2,

	BIG		= (1&lt;&lt;12)-4,
	STRINGSZ	= 200,
	NHASH		= 10007,
	NHUNK		= 100000,
	MINSIZ		= 64,
	NENT		= 100,
	MAXIO		= 8192,
	MAXHIST		= 20,	/* limit of path elements for history symbols */

	Roffset	= 22,		/* no. bits for offset in relocation address */
	Rindex	= 10,		/* no. bits for index in relocation address */
};

EXTERN union
{
	struct
	{
		uchar	obuf[MAXIO];			/* output buffer */
		uchar	ibuf[MAXIO];			/* input buffer */
	} u;
	char	dbuf[1];
} buf;

#define	cbuf	u.obuf
#define	xbuf	u.ibuf

#define	setarch(p)		if((p)-&gt;as==ATEXT) thumb=(p)-&gt;reg&amp;ALLTHUMBS
#define	setthumb(p)	if((p)-&gt;as==ATEXT) seenthumb|=(p)-&gt;reg&amp;ALLTHUMBS

#ifndef COFFCVT

EXTERN	int32	HEADR;			/* length of header */
EXTERN	int	HEADTYPE;		/* type of header */
EXTERN	int32	INITDAT;		/* data location */
EXTERN	int32	INITRND;		/* data round above text location */
EXTERN	int32	INITTEXT;		/* text location */
EXTERN	char*	INITENTRY;		/* entry point */
EXTERN	int32	autosize;
EXTERN	Biobuf	bso;
EXTERN	int32	bsssize;
EXTERN	int	cbc;
EXTERN	uchar*	cbp;
EXTERN	int	cout;
EXTERN	Auto*	curauto;
EXTERN	Auto*	curhist;
EXTERN	Prog*	curp;
EXTERN	Prog*	curtext;
EXTERN	Prog*	datap;
EXTERN	int32	datsize;
EXTERN	char	debug[128];
EXTERN	Prog*	edatap;
EXTERN	Prog*	etextp;
EXTERN	Prog*	firstp;
EXTERN	char*	noname;
EXTERN	int	xrefresolv;
EXTERN	Prog*	lastp;
EXTERN	int32	lcsize;
EXTERN	char	literal[32];
EXTERN	int	nerrors;
EXTERN	int32	instoffset;
EXTERN	Opcross	opcross[8];
EXTERN	Oprang	oprange[ALAST];
EXTERN	Oprang	thumboprange[ALAST];
EXTERN	char*	outfile;
EXTERN	int32	pc;
EXTERN	uchar	repop[ALAST];
EXTERN	uint32	stroffset;
EXTERN	int32	symsize;
EXTERN	Prog*	textp;
EXTERN	int32	textsize;
EXTERN	int	version;
EXTERN	char	xcmp[C_GOK+1][C_GOK+1];
EXTERN	Prog	zprg;
EXTERN	int	dtype;
EXTERN	int	armv4;
EXTERN	int	thumb;
EXTERN	int	seenthumb;
EXTERN	int	armsize;

EXTERN	int	doexp, dlm;
EXTERN	int	imports, nimports;
EXTERN	int	exports, nexports;
EXTERN	char*	EXPTAB;
EXTERN	Prog	undefp;

#define	UP	(&amp;undefp)

extern	char*	anames[];
extern	Optab	optab[];
extern	Optab	thumboptab[];

void	addpool(Prog*, Adr*);
EXTERN	Prog*	blitrl;
EXTERN	Prog*	elitrl;

void	initdiv(void);
EXTERN	Prog*	prog_div;
EXTERN	Prog*	prog_divu;
EXTERN	Prog*	prog_mod;
EXTERN	Prog*	prog_modu;

#pragma	varargck	type	&#34;A&#34;	int
#pragma	varargck	type	&#34;C&#34;	int
#pragma	varargck	type	&#34;D&#34;	Adr*
#pragma	varargck	type	&#34;N&#34;	Adr*
#pragma	varargck	type	&#34;P&#34;	Prog*
#pragma	varargck	type	&#34;S&#34;	char*

int	Aconv(Fmt*);
int	Cconv(Fmt*);
int	Dconv(Fmt*);
int	Nconv(Fmt*);
int	Oconv(Fmt*);
int	Pconv(Fmt*);
int	Sconv(Fmt*);
int	aclass(Adr*);
int	thumbaclass(Adr*, Prog*);
void	addhist(int32, int);
Prog*	appendp(Prog*);
void	asmb(void);
void	asmdyn(void);
void	asmlc(void);
void	asmthumbmap(void);
void	asmout(Prog*, Optab*);
void	thumbasmout(Prog*, Optab*);
void	asmsym(void);
int32	atolwhex(char*);
Prog*	brloop(Prog*);
void	buildop(void);
void	thumbbuildop(void);
void	buildrep(int, int);
void	cflush(void);
void	ckoff(Sym*, int32);
int	chipfloat(Ieee*);
int	cmp(int, int);
int	compound(Prog*);
double	cputime(void);
void	datblk(int32, int32, int);
void	diag(char*, ...);
void	divsig(void);
void	dodata(void);
void	doprof1(void);
void	doprof2(void);
void	dynreloc(Sym*, int32, int);
int32	entryvalue(void);
void	exchange(Prog*);
void	export(void);
void	follow(void);
void	hputl(int);
void	import(void);
int	isnop(Prog*);
void	listinit(void);
Sym*	lookup(char*, int);
void	cput(int);
void	hput(int32);
void	lput(int32);
void	lputl(int32);
void	mkfwd(void);
void*	mysbrk(uint32);
void	names(void);
Prog*	newdata(Sym *s, int o, int w, int t);
void	nocache(Prog*);
int	ocmp(const void*, const void*);
int32	opirr(int);
Optab*	oplook(Prog*);
int32	oprrr(int, int);
int32	olr(int32, int, int, int);
int32	olhr(int32, int, int, int);
int32	olrr(int, int, int, int);
int32	olhrr(int, int, int, int);
int32	osr(int, int, int32, int, int);
int32	oshr(int, int32, int, int);
int32	ofsr(int, int, int32, int, int, Prog*);
int32	osrr(int, int, int, int);
int32	oshrr(int, int, int, int);
int32	omvl(Prog*, Adr*, int);
void	patch(void);
void	prasm(Prog*);
void	prepend(Prog*, Prog*);
Prog*	prg(void);
int	pseudo(Prog*);
void	putsymb(char*, int, int32, int);
int32	regoff(Adr*);
int	relinv(int);
int32	rnd(int32, int32);
void	span(void);
void	strnput(char*, int);
void	undef(void);
void	wput(int32);
void    wputl(ushort w);
void	xdefine(char*, int, int32);
void	xfol(Prog*);
void	noops(void);
int32	immrot(uint32);
int32	immaddr(int32);
int32	opbra(int, int);
int	brextra(Prog*);
int	isbranch(Prog*);
int	fnpinc(Sym *);
int	fninc(Sym *);
void	thumbcount(void);
void reachable(void);
void fnptrs(void);

uint32	linuxheadr(void);
void	linuxphdr(int type, int flags, vlong foff,
	vlong vaddr, vlong paddr,
	vlong filesize, vlong memsize, vlong align);
void	linuxshdr(char *name, uint32 type, vlong flags, vlong addr, vlong off,
	vlong size, uint32 link, uint32 info, vlong align, vlong entsize);
int	linuxstrtable(void);

/*
 *	go.c
 */
void	deadcode(void);
char*	gotypefor(char *name);
void	ldpkg(Biobuf *f, int64 len, char *filename);

#endif
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
