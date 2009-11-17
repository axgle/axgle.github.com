<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5g/gg.h</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5g/gg.h</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &lt;u.h&gt;
#include &lt;libc.h&gt;

#include &#34;../gc/go.h&#34;
#include &#34;../5l/5.out.h&#34;

#ifndef	EXTERN
#define EXTERN	extern
#endif

typedef	struct	Addr	Addr;

struct	Addr
{
	int32	offset;
	int32	offset2;
	double	dval;
	Prog*	branch;
	char	sval[NSNAME];

	Sym*	sym;
	int	width;
	uchar	type;
	char	name;
	char	reg;
	uchar	etype;
};
#define	A	((Addr*)0)

struct	Prog
{
	short	as;			// opcode
	uint32	loc;		// pc offset in this func
	uint32	lineno;		// source line that generated this
	Addr	from;		// src address
	Addr	to;			// dst address
	Prog*	link;		// next instruction in this func
	char	reg;		// doubles as width in DATA op
	uchar	scond;
};

#define REGALLOC_R0 0
#define REGALLOC_RMAX REGEXT
#define REGALLOC_F0 (REGALLOC_RMAX+1)
#define REGALLOC_FMAX (REGALLOC_F0 + FREGEXT)

EXTERN	Biobuf*	bout;
EXTERN	int32	dynloc;
EXTERN	uchar	reg[REGALLOC_FMAX];
EXTERN	int32	pcloc;		// instruction counter
EXTERN	Strlit	emptystring;
extern	char*	anames[];
EXTERN	Hist*	hist;
EXTERN	Prog	zprog;
EXTERN	Node*	curfn;
EXTERN	Node*	newproc;
EXTERN	Node*	deferproc;
EXTERN	Node*	deferreturn;
EXTERN	Node*	throwindex;
EXTERN	Node*	throwreturn;
EXTERN	long	unmappedzero;
EXTERN	int	maxstksize;

/*
 * gen.c
 */
void	compile(Node*);
void	proglist(void);
void	gen(Node*);
Node*	lookdot(Node*, Node*, int);
void	cgen_as(Node*, Node*);
void	cgen_callmeth(Node*, int);
void	cgen_callinter(Node*, Node*, int);
void	cgen_proc(Node*, int);
void	cgen_callret(Node*, Node*);
void	cgen_dcl(Node*);
int	needconvert(Type*, Type*);
void	genconv(Type*, Type*);
void	allocparams(void);
void	checklabels();
void	ginscall(Node*, int);

/*
 * cgen
 */
void	agen(Node*, Node*);
void	igen(Node*, Node*, Node*);
void agenr(Node *n, Node *a, Node *res);
vlong	fieldoffset(Type*, Node*);
void	bgen(Node*, int, Prog*);
void	sgen(Node*, Node*, int32);
void	gmove(Node*, Node*);
Prog*	gins(int, Node*, Node*);
int	samaddr(Node*, Node*);
void	raddr(Node *n, Prog *p);
Prog*	gcmp(int, Node*, Node*);
Prog*	gshift(int as, Node *lhs, int32 stype, int32 sval, Node *rhs);
Prog *	gregshift(int as, Node *lhs, int32 stype, Node *reg, Node *rhs);
void	naddr(Node*, Addr*, int);
void	cgen_aret(Node*, Node*);
void	cgen_shift(int, Node*, Node*, Node*);

/*
 * cgen64.c
 */
void	cmp64(Node*, Node*, int, Prog*);
void	cgen64(Node*, Node*);

/*
 * gsubr.c
 */
void	clearp(Prog*);
void	proglist(void);
Prog*	gbranch(int, Type*);
Prog*	prog(int);
void	gaddoffset(Node*);
void	gconv(int, int);
int	conv2pt(Type*);
vlong	convvtox(vlong, int);
void	fnparam(Type*, int, int);
Prog*	gop(int, Node*, Node*, Node*);
void	setconst(Addr*, vlong);
void	setaddr(Addr*, Node*);
int	optoas(int, Type*);
void	ginit(void);
void	gclean(void);
void	regalloc(Node*, Type*, Node*);
void	regfree(Node*);
Node*	nodarg(Type*, int);
void	nodreg(Node*, Type*, int);
void	nodindreg(Node*, Type*, int);
void	buildtxt(void);
Plist*	newplist(void);
int	isfat(Type*);
int	dotaddable(Node*, Node*);
void	sudoclean(void);
int	sudoaddable(int, Node*, Addr*, int*);
void	afunclit(Addr*);
void	datagostring(Strlit*, Addr*);
void	split64(Node*, Node*, Node*);
void	splitclean(void);
Node*	ncon(uint32 i);

/*
 * obj.c
 */
void	datastring(char*, int, Addr*);

/*
 * list.c
 */
int	Aconv(Fmt*);
int	Cconv(Fmt*);
int	Dconv(Fmt*);
int	Mconv(Fmt*);
int	Pconv(Fmt*);
int	Rconv(Fmt*);
int	Yconv(Fmt*);
void	listinit(void);

void	zaddr(Biobuf*, Addr*, int);
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
