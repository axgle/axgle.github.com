<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8l/l.h</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/8l/l.h</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/8l/l.h
// http://code.google.com/p/inferno-os/source/browse/utils/8l/l.h
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
#include	&#34;../8l/8.out.h&#34;

#ifndef	EXTERN
#define	EXTERN	extern
#endif

enum
{
	PtrSize = 4
};

#define	P		((Prog*)0)
#define	S		((Sym*)0)
#define	TNAME		(curtext?curtext-&gt;from.sym-&gt;name:noname)
#define	cput(c)\
	{ *cbp++ = c;\
	if(--cbc &lt;= 0)\
		cflush(); }

typedef	struct	Adr	Adr;
typedef	struct	Prog	Prog;
typedef	struct	Sym	Sym;
typedef	struct	Auto	Auto;
typedef	struct	Optab	Optab;

struct	Adr
{
	union
	{
		int32	u0offset;
		char	u0scon[8];
		Prog	*u0cond;	/* not used, but should be D_BRANCH */
		Ieee	u0ieee;
		char	*u0sbig;
	} u0;
	union
	{
		Auto*	u1autom;
		Sym*	u1sym;
	} u1;
	short	type;
	uchar	index;
	char	scale;
	int32	offset2;
};

#define	offset	u0.u0offset
#define	scon	u0.u0scon
#define	cond	u0.u0cond
#define	ieee	u0.u0ieee
#define	sbig	u0.u0sbig

#define	autom	u1.u1autom
#define	sym	u1.u1sym

struct	Prog
{
	Adr	from;
	Adr	to;
	Prog	*forwd;
	Prog*	link;
	Prog*	dlink;
	Prog*	pcond;	/* work on this */
	int32	pc;
	int32	line;
	short	as;
	char	width;		/* fake for DATA */
	char	ft;		/* oclass cache */
	char	tt;
	uchar	mark;	/* work on these */
	uchar	back;

};
struct	Auto
{
	Sym*	asym;
	Auto*	link;
	int32	aoffset;
	short	type;
	Sym*	gotype;
};
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
	int32	size;
	int32	sig;
	Sym*	link;
	Prog*	text;
	Prog*	data;
	Sym*	gotype;
	char*	file;
	char*	dynldname;
	char*	dynldlib;
};
struct	Optab
{
	short	as;
	uchar*	ytab;
	uchar	prefix;
	uchar	op[10];
};

enum
{
	Sxxx,

	STEXT,
	SDATA,
	SBSS,
	SDATA1,
	SXREF,
	SFILE,
	SCONST,
	SUNDEF,

	SIMPORT,
	SEXPORT,

	SMACHO,	/* pointer to mach-o imported symbol */

	NHASH		= 10007,
	NHUNK		= 100000,
	MINSIZ		= 4,
	STRINGSZ	= 200,
	MINLC		= 1,
	MAXIO		= 8192,
	MAXHIST		= 20,				/* limit of path elements for history symbols */

	Yxxx		= 0,
	Ynone,
	Yi0,
	Yi1,
	Yi8,
	Yi32,
	Yiauto,
	Yal,
	Ycl,
	Yax,
	Ycx,
	Yrb,
	Yrl,
	Yrf,
	Yf0,
	Yrx,
	Ymb,
	Yml,
	Ym,
	Ybr,
	Ycol,

	Ycs,	Yss,	Yds,	Yes,	Yfs,	Ygs,
	Ygdtr,	Yidtr,	Yldtr,	Ymsw,	Ytask,
	Ycr0,	Ycr1,	Ycr2,	Ycr3,	Ycr4,	Ycr5,	Ycr6,	Ycr7,
	Ydr0,	Ydr1,	Ydr2,	Ydr3,	Ydr4,	Ydr5,	Ydr6,	Ydr7,
	Ytr0,	Ytr1,	Ytr2,	Ytr3,	Ytr4,	Ytr5,	Ytr6,	Ytr7,
	Ymax,

	Zxxx		= 0,

	Zlit,
	Z_rp,
	Zbr,
	Zcall,
	Zcallcon,
	Zib_,
	Zib_rp,
	Zibo_m,
	Zil_,
	Zil_rp,
	Zilo_m,
	Zjmp,
	Zjmpcon,
	Zloop,
	Zm_o,
	Zm_r,
	Zaut_r,
	Zo_m,
	Zpseudo,
	Zr_m,
	Zrp_,
	Z_ib,
	Z_il,
	Zm_ibo,
	Zm_ilo,
	Zib_rr,
	Zil_rr,
	Zclr,
	Zbyte,
	Zmov,
	Zmax,

	Px		= 0,
	Pe		= 0x66,	/* operand escape */
	Pm		= 0x0f,	/* 2byte opcode escape */
	Pq		= 0xff,	/* both escape */
	Pb		= 0xfe,	/* byte operands */

	Roffset	= 22,		/* no. bits for offset in relocation address */
	Rindex	= 10,		/* no. bits for index in relocation address */
};

EXTERN union
{
	struct
	{
		char	obuf[MAXIO];			/* output buffer */
		uchar	ibuf[MAXIO];			/* input buffer */
	} u;
	char	dbuf[1];
} buf;

#define	cbuf	u.obuf
#define	xbuf	u.ibuf

#pragma	varargck	type	&#34;A&#34;	uint
#pragma	varargck	type	&#34;D&#34;	Adr*
#pragma	varargck	type	&#34;P&#34;	Prog*
#pragma	varargck	type	&#34;R&#34;	int
#pragma	varargck	type	&#34;S&#34;	char*

EXTERN	int32	HEADR;
EXTERN	int32	HEADTYPE;
EXTERN	int32	INITDAT;
EXTERN	int32	INITRND;
EXTERN	int32	INITTEXT;
EXTERN	char*	INITENTRY;		/* entry point */
EXTERN	Biobuf	bso;
EXTERN	int32	bsssize;
EXTERN	int32	casepc;
EXTERN	int	cbc;
EXTERN	char*	cbp;
EXTERN	char*	pcstr;
EXTERN	Auto*	curauto;
EXTERN	Auto*	curhist;
EXTERN	Prog*	curp;
EXTERN	Prog*	curtext;
EXTERN	Prog*	datap;
EXTERN	Prog*	edatap;
EXTERN	int32	datsize;
EXTERN	int32	dynptrsize;
EXTERN	char	debug[128];
EXTERN	char	literal[32];
EXTERN	Prog*	etextp;
EXTERN	Prog*	firstp;
EXTERN	int	xrefresolv;
EXTERN	uchar	ycover[Ymax*Ymax];
EXTERN	uchar*	andptr;
EXTERN	uchar	and[100];
EXTERN	char	reg[D_NONE];
EXTERN	Prog*	lastp;
EXTERN	int32	lcsize;
EXTERN	int	maxop;
EXTERN	int	nerrors;
EXTERN	char*	noname;
EXTERN	int32	pc;
EXTERN	int32	spsize;
EXTERN	Sym*	symlist;
EXTERN	int32	symsize;
EXTERN	Prog*	textp;
EXTERN	int32	textsize;
EXTERN	int	version;
EXTERN	Prog	zprg;
EXTERN	int	dtype;
EXTERN	int	tlsoffset;
EXTERN	Sym*	adrgotype;	// type symbol on last Adr read
EXTERN	Sym*	fromgotype;	// type symbol on last p-&gt;from read

EXTERN	Adr*	reloca;
EXTERN	int	doexp, dlm;
EXTERN	int	imports, nimports;
EXTERN	int	exports, nexports;
EXTERN	char*	EXPTAB;
EXTERN	Prog	undefp;

#define	UP	(&amp;undefp)

extern	Optab	optab[];
extern	char*	anames[];

int	Aconv(Fmt*);
int	Dconv(Fmt*);
int	Pconv(Fmt*);
int	Rconv(Fmt*);
int	Sconv(Fmt*);
void	addhist(int32, int);
Prog*	appendp(Prog*);
void	asmb(void);
void	asmdyn(void);
void	asmins(Prog*);
void	asmlc(void);
void	asmsp(void);
void	asmsym(void);
int32	atolwhex(char*);
Prog*	brchain(Prog*);
Prog*	brloop(Prog*);
void	cflush(void);
void	ckoff(Sym*, int32);
Prog*	copyp(Prog*);
double	cputime(void);
void	datblk(int32, int32);
void	diag(char*, ...);
void	dodata(void);
void	doelf(void);
void	doinit(void);
void	doprof1(void);
void	doprof2(void);
void	dostkoff(void);
void	dynreloc(Sym*, uint32, int);
int32	entryvalue(void);
void	export(void);
void	follow(void);
void	import(void);
void	listinit(void);
Sym*	lookup(char*, int);
void	lput(int32);
void	lputl(int32);
void	vputl(uvlong);
void	strnput(char*, int);
void	main(int, char*[]);
void	mkfwd(void);
void*	mal(uint32);
Prog*	newdata(Sym*, int, int, int);
Prog*	newtext(Prog*, Sym*);
int	opsize(Prog*);
void	patch(void);
Prog*	prg(void);
int	relinv(int);
int32	reuse(Prog*, Sym*);
int32	rnd(int32, int32);
void	s8put(char*);
void	span(void);
void	undef(void);
int32	vaddr(Adr*);
int32	symaddr(Sym*);
void	wput(ushort);
void	wputl(ushort);
void	xdefine(char*, int, int32);
void	xfol(Prog*);
void	zaddr(Biobuf*, Adr*, Sym*[]);
uint32	machheadr(void);
vlong		addaddr(Sym *s, Sym *t);
vlong		addsize(Sym *s, Sym *t);
vlong		addstring(Sym *s, char *str);
vlong		adduint16(Sym *s, uint16 v);
vlong		adduint32(Sym *s, uint32 v);
vlong		adduint64(Sym *s, uint64 v);
vlong		adduint8(Sym *s, uint8 v);
vlong		adduintxx(Sym *s, uint64 v, int wid);

/*
 *	go.c
 */
void	deadcode(void);

/* Native is little-endian */
#define	LPUT(a)	lputl(a)
#define	WPUT(a)	wputl(a)
#define	VPUT(a)	vputl(a)

#pragma	varargck	type	&#34;D&#34;	Adr*
#pragma	varargck	type	&#34;P&#34;	Prog*
#pragma	varargck	type	&#34;R&#34;	int
#pragma	varargck	type	&#34;A&#34;	int
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
