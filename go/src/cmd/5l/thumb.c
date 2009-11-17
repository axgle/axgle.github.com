<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5l/thumb.c</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5l/thumb.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5l/thumb.c
// http://code.google.com/p/inferno-os/source/browse/utils/5l/thumb.c
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

#include &#34;l.h&#34;

static int32 thumboprr(int);
static int32 thumboprrr(int, int);
static int32 thumbopirr(int , int);
static int32 thumbopri(int);
static int32 thumbophh(int);
static int32 thumbopbra(int);
static int32 thumbopmv(int, int);
static void lowreg(Prog *, int);
static void mult(Prog *, int, int);
static void numr(Prog *, int, int, int);
static void regis(Prog *, int, int, int);
static void dis(int, int);

// build a constant using neg, add and shift - only worth it if &lt; 6 bytes */
static int
immbuildcon(int c, Prog *p)
{
	int n = 0;

	USED(p);
	if(c &gt;= 0 &amp;&amp; c &lt;= 255)
		return 0;			// mv
	if(c &gt;= -255 &amp;&amp; c &lt; 0)	// mv, neg
		return 1;
	if(c &gt;= 256 &amp;&amp; c &lt;= 510)	// mv, add
		return 1;
	if(c &lt; 0)
		return 0;
	while(!(c &amp; 1)){
		n++;
		c &gt;&gt;= 1;
	}
	if(c &gt;= 0 &amp;&amp; c &lt;= 255)	// mv, lsl
		return 1;
	return 0;
}

// positive 5 bit offset from register - O(R)
// positive 8 bit offset from register - mov O, R then [R, R]
// otherwise O goes in literal pool - mov O1(PC), R then [R, R]
static int
immoreg(int off, Prog *p)
{
	int v = 1;
	int as = p-&gt;as;

	if(off &lt; 0)
		return C_GOREG;
	if(as == AMOVW)
		v = 4;
	else if(as == AMOVH || as == AMOVHU)
		v = 2;
	else if(as == AMOVB || as == AMOVBU)
		v = 1;
	else
		diag(&#34;bad op in immoreg&#34;);
	if(off/v &lt;= 31)
		return C_SOREG;
	if(off &lt;= 255)
		return C_LOREG;
	return C_GOREG;
}

// positive 8 bit - mov O, R then 0(R)
// otherwise O goes in literal pool - mov O1(PC), R then 0(R)
static int
immacon(int off, Prog *p, int t1, int t2)
{
	USED(p);
	if(off &lt; 0)
		return t2;
	if(off &lt;= 255)
		return t1;
	return t2;
}

// unsigned 8 bit in words
static int
immauto(int off, Prog *p)
{
	if(p-&gt;as != AMOVW)
		diag(&#34;bad op in immauto&#34;);
	mult(p, off, 4);
	if(off &gt;= 0 &amp;&amp; off &lt;= 1020)
		return C_SAUTO;
	return C_LAUTO;
}

static int
immsmall(int off, Prog *p, int t1, int t2, int t3)
{
	USED(p);
	if(off &gt;= 0 &amp;&amp; off &lt;= 7)
		return t1;
	if(off &gt;= 0 &amp;&amp; off &lt;= 255)
		return t2;
	return t3;
}

static int
immcon(int off, Prog *p)
{
	int as = p-&gt;as;

	if(as == ASLL || as == ASRL || as == ASRA)
		return C_SCON;
	if(p-&gt;to.type == D_REG &amp;&amp; p-&gt;to.reg == REGSP){
		if(as == AADD || as == ASUB){
			if(off &gt;= 0 &amp;&amp; off &lt;= 508)
				return C_SCON;
			if(as == ASUB){
				p-&gt;as = AADD;
				p-&gt;from.offset = -p-&gt;from.offset;
			}
			return C_LCON;
		}
		diag(&#34;unknown type in immcon&#34;);
	}
	if(as == AADD || as == ASUB){
		if(p-&gt;reg != NREG)
			return immsmall(off, p, C_SCON, C_LCON, C_GCON);
		return immacon(off, p, C_SCON, C_LCON);
	}
	if(as == AMOVW &amp;&amp; p-&gt;from.type == D_CONST &amp;&amp; p-&gt;to.type == D_REG &amp;&amp; immbuildcon(off, p))
		return C_BCON;
	if(as == ACMP &amp;&amp; p-&gt;from.type == D_CONST &amp;&amp; immbuildcon(off, p))
		return C_BCON;
	if(as == ACMP || as == AMOVW)
		return immacon(off, p, C_SCON, C_LCON);
	return C_LCON;
}

int
thumbaclass(Adr *a, Prog *p)
{
	Sym *s;
	int t;

	switch(a-&gt;type) {
	case D_NONE:
		return C_NONE;
	case D_REG:
		if(a-&gt;reg == REGSP)
			return C_SP;
		if(a-&gt;reg == REGPC)
			return C_PC;
		if(a-&gt;reg &gt;= 8)
			return C_HREG;
		return C_REG;
	case D_SHIFT:
		diag(&#34;D_SHIFT in thumbaclass&#34;);
		return C_SHIFT;
	case D_FREG:
		diag(&#34;D_FREG in thumbaclass&#34;);
		return C_FREG;
	case D_FPCR:
		diag(&#34;D_FPCR in thumbaclass&#34;);
		return C_FCR;
	case D_OREG:
		switch(a-&gt;name) {
		case D_EXTERN:
		case D_STATIC:
			if(a-&gt;sym == 0 || a-&gt;sym-&gt;name == 0) {
				print(&#34;null sym external\n&#34;);
				print(&#34;%D\n&#34;, a);
				return C_GOK;
			}
			t = a-&gt;sym-&gt;type;
			if(t == 0 || t == SXREF) {
				diag(&#34;undefined external: %s in %s\n&#34;,
					a-&gt;sym-&gt;name, TNAME);
				a-&gt;sym-&gt;type = SDATA;
			}
			instoffset = a-&gt;sym-&gt;value + a-&gt;offset + INITDAT;
			return C_LEXT;	/* INITDAT unknown at this stage */
			// return immacon(instoffset, p, C_SEXT, C_LEXT);
		case D_AUTO:
			instoffset = autosize + a-&gt;offset;
			return immauto(instoffset, p);
		case D_PARAM:
			instoffset = autosize + a-&gt;offset + 4L;
// print(&#34;D_PARAM %s %d+%d+%d = %d\n&#34;, a-&gt;sym != S ? a-&gt;sym-&gt;name : &#34;noname&#34;, autosize, a-&gt;offset, 4, autosize+a-&gt;offset+4);
			return immauto(instoffset, p);
		case D_NONE:
			instoffset = a-&gt;offset;
			if(a-&gt;reg == REGSP)
				return immauto(instoffset, p);
			else
				return immoreg(instoffset, p);
		}
		return C_GOK;
	case D_PSR:
		diag(&#34;D_PSR in thumbaclass&#34;);
		return C_PSR;
	case D_OCONST:
		switch(a-&gt;name) {
		case D_EXTERN:
		case D_STATIC:
			s = a-&gt;sym;
			t = s-&gt;type;
			if(t == 0 || t == SXREF) {
				diag(&#34;undefined external: %s in %s\n&#34;,
					s-&gt;name, TNAME);
				s-&gt;type = SDATA;
			}
			instoffset = s-&gt;value + a-&gt;offset + INITDAT;
			if(s-&gt;type == STEXT || s-&gt;type == SLEAF){
				instoffset = s-&gt;value + a-&gt;offset;
#ifdef CALLEEBX
				instoffset += fnpinc(s);
#else
				if(s-&gt;thumb)
					instoffset++;	// T bit
#endif
				return C_LCON;
			}
			return C_LCON;	/* INITDAT unknown at this stage */
			// return immcon(instoffset, p);
		}
		return C_GOK;
	case D_FCONST:
		diag(&#34;D_FCONST in thumaclass&#34;);
		return C_FCON;
	case D_CONST:
		switch(a-&gt;name) {
		case D_NONE:
			instoffset = a-&gt;offset;
			if(a-&gt;reg != NREG)
				goto aconsize;
			return immcon(instoffset, p);
		case D_EXTERN:
		case D_STATIC:
			s = a-&gt;sym;
			if(s == S)
				break;
			t = s-&gt;type;
			switch(t) {
			case 0:
			case SXREF:
				diag(&#34;undefined external: %s in %s\n&#34;,
					s-&gt;name, TNAME);
				s-&gt;type = SDATA;
				break;
			case SCONST:
			case STEXT:
			case SLEAF:
				instoffset = s-&gt;value + a-&gt;offset;
#ifdef CALLEEBX
				instoffset += fnpinc(s);
#else
				if(s-&gt;thumb)
					instoffset++;	// T bit
#endif
				return C_LCON;
			}
			instoffset = s-&gt;value + a-&gt;offset + INITDAT;
			return C_LCON;	/* INITDAT unknown at this stage */
			// return immcon(instoffset, p);
		case D_AUTO:
			instoffset = autosize + a-&gt;offset;
			goto aconsize;
		case D_PARAM:
			instoffset = autosize + a-&gt;offset + 4L;
		aconsize:
			if(p-&gt;from.reg == REGSP || p-&gt;from.reg == NREG)
				return instoffset &gt;= 0 &amp;&amp; instoffset &lt; 1024 ? C_SACON : C_GACON;
			else if(p-&gt;from.reg == p-&gt;to.reg)
				return immacon(instoffset, p, C_SACON, C_GACON);
			return immsmall(instoffset, p, C_SACON, C_LACON, C_GACON);
		}
		return C_GOK;
	case D_BRANCH: {
		int v, va;

		p-&gt;align = 0;
		v = -4;
		va = 0;
		if(p-&gt;cond != P){
			v = (p-&gt;cond-&gt;pc - p-&gt;pc) - 4;
			va = p-&gt;cond-&gt;pc;
		}
		instoffset = v;
		if(p-&gt;as == AB){
			if(v &gt;= -2048 &amp;&amp; v &lt;= 2046)
				return C_SBRA;
			p-&gt;align = 4;
			instoffset = va;
			return C_LBRA;
		}
		if(p-&gt;as == ABL){
#ifdef CALLEEBX
			int e;

			if((e = fninc(p-&gt;to.sym))) {
				v += e;
				va += e;
				instoffset += e;
			}
#endif
			if(v &gt;= -4194304 &amp;&amp; v &lt;= 4194302)
				return C_SBRA;
			p-&gt;align = 2;
			instoffset = va;
			return C_LBRA;
		}
		if(p-&gt;as == ABX){
			v = va;
			if(v &gt;= 0 &amp;&amp; v &lt;= 255)
				return C_SBRA;
			p-&gt;align = 2;
			instoffset = va;
			return C_LBRA;
		}
		if(v &gt;= -256 &amp;&amp; v &lt;= 254)
			return C_SBRA;
		if(v &gt;= -(2048-2) &amp;&amp; v &lt;= (2046+2))
			return C_LBRA;
		p-&gt;align = 2;
		instoffset = va;
		return C_GBRA;
	}
	}
	return C_GOK;
}

// as a1 a2 a3 type size param lit vers
Optab thumboptab[] =
{
	{ ATEXT,		C_LEXT,		C_NONE,		C_LCON,		0,	0,	0 },
	{ ATEXT,		C_LEXT,		C_REG,		C_LCON,		0,	0,	0 },
	{ AMVN,		C_REG,		C_NONE,		C_REG,		1,	2,	0 },
	{ ASRL,		C_REG,		C_NONE,		C_REG,		1,	2,	0 },
	{ ACMP,		C_REG,		C_REG,		C_NONE,		1,	2,	0 },
	{ ACMN,		C_REG,		C_REG,		C_NONE,		1,	2,	0 },
	{ AADD,		C_REG,		C_REG,		C_REG,		2,	2,	0 },
	{ AADD,		C_REG,		C_NONE,		C_REG,		2,	2,	0 },
	{ AADD,		C_SCON,		C_REG,		C_REG,		3,	2,	0 },
	{ AADD,		C_LCON,		C_REG,		C_REG,		49,	4,	0 },
	{ AADD,		C_GCON,		C_REG,		C_REG,		36,	4,	0,	LFROM },
	// { AADD,		C_LCON,		C_NONE,		C_REG,		3,	2,	0,	LFROM },
	{ ASRL,		C_SCON,		C_REG,		C_REG,		4,	2,	0 },
	{ ASRL,		C_SCON,		C_NONE,		C_REG,		4,	2,	0 },
	{ AADD,		C_SCON,		C_NONE,		C_REG,		5,	2,	0 },
	{ AADD,		C_LCON,		C_NONE,		C_REG,		37,	4,	0,	LFROM },
	{ ACMP,		C_SCON,		C_REG,		C_NONE,		5,	2,	0 },
	{ ACMP,		C_BCON,		C_REG,		C_NONE,		48,	6,	0 },
	{ ACMP,		C_LCON,		C_REG,		C_NONE,		39,	4,	0,	LFROM },
	{ AMOVW,		C_SCON,		C_NONE,		C_REG,		5,	2,	0 },
	{ AMOVW,		C_BCON,		C_NONE,		C_REG,		47,	4,	0 },
	{ AMOVW,		C_LCON,		C_NONE,		C_REG,		38,	2,	0,	LFROM },
	// { AADD,		C_LCON,		C_PC,		C_REG,		6,	2,	0,	LFROM },
	// { AADD,		C_LCON,		C_SP,		C_REG,		6,	2,	0,	LFROM },
	{ AADD,		C_SCON,		C_NONE,		C_SP,		7,	2,	0 },
	{ AADD,		C_LCON,		C_NONE,		C_SP,		40,	4,	0,	LFROM },
	{ AADD,		C_REG,		C_NONE,		C_HREG,		8,	2,	0 },
	{ AADD,		C_HREG,		C_NONE,		C_REG,		8,	2,	0 },
	{ AADD,		C_HREG,		C_NONE,		C_HREG,		8,	2,	0 },
	{ AMOVW,		C_REG,		C_NONE,		C_HREG,		8,	2,	0 },
	{ AMOVW,		C_HREG,		C_NONE,		C_REG,		8,	2,	0 },
	{ AMOVW,		C_HREG,		C_NONE,		C_HREG,		8,	2,	0 },
	{ ACMP,		C_REG,		C_HREG,		C_NONE,		8,	2,	0 },
	{ ACMP,		C_HREG,		C_REG,		C_NONE,		8,	2,	0 },
	{ ACMP,		C_HREG,		C_HREG,		C_NONE,		8,	2,	0 },
	{ AB,			C_NONE,		C_NONE,		C_SBRA,		9,	2,	0,	LPOOL },
	{ ABEQ,		C_NONE,		C_NONE,		C_SBRA,		10,	2,	0 },
	{ ABL,		C_NONE,		C_NONE,		C_SBRA,		11,	4,	0 },
	{ ABX,		C_NONE,		C_NONE,		C_SBRA,		12,	10,	0 },
	{ AB,			C_NONE,		C_NONE,		C_LBRA,		41,	8,	0,	LPOOL },
	{ ABEQ,		C_NONE,		C_NONE,		C_LBRA,		46,	4,	0 },
	{ ABL,		C_NONE,		C_NONE,		C_LBRA,		43,	14,	0 },
	{ ABX,		C_NONE,		C_NONE,		C_LBRA,		44,	14,	0 },
	{ ABEQ,		C_NONE,		C_NONE,		C_GBRA,		42,  10, 	0 },
	// { AB,		C_NONE,		C_NONE,		C_SOREG,		13,	0,	0 },
	// { ABL,		C_NONE,		C_NONE,		C_SOREG,		14,	0,	0 },
	{ ABL,		C_NONE,		C_NONE,		C_REG,		51,	4,	0 },
	{ ABX,		C_NONE,		C_NONE,		C_REG,		15,	8,	0 },
	{ ABX,		C_NONE,		C_NONE,		C_HREG,		15,	8,	0 },
	{ ABXRET,		C_NONE,		C_NONE,		C_REG,		45,	2,	0 },
	{ ABXRET,		C_NONE,		C_NONE,		C_HREG,		45,	2,	0 },
	{ ASWI,		C_NONE,		C_NONE,		C_LCON,		16,	2,	0 },
	{ AWORD,		C_NONE,		C_NONE,		C_LCON,		17,	4,	0 },
	{ AWORD,		C_NONE,		C_NONE,		C_GCON,		17,	4,	0 },
	{ AWORD,		C_NONE,		C_NONE,		C_LEXT,		17,	4, 	0 },
	{ ADWORD,	C_LCON,		C_NONE,		C_LCON,		50,	8,	0 },
	{ AMOVW,		C_SAUTO,		C_NONE,		C_REG,		18,	2,	REGSP },
	{ AMOVW,		C_LAUTO,		C_NONE,		C_REG,		33,	6,	0,	LFROM  },
	// { AMOVW,		C_OFFPC,		C_NONE,		C_REG,		18,	2,	REGPC,	LFROM  },
	{ AMOVW,		C_SEXT,		C_NONE,		C_REG,		30,	4,	0 },
	{ AMOVW,		C_SOREG,		C_NONE,		C_REG,		19,	2,	0 },
	{ AMOVHU,	C_SEXT,		C_NONE,		C_REG,		30,	4,	0 },
	{ AMOVHU,	C_SOREG,		C_NONE,		C_REG,		19,	2,	0 },
	{ AMOVBU,	C_SEXT,		C_NONE,		C_REG,		30,	4,	0 },
	{ AMOVBU,	C_SOREG,		C_NONE,		C_REG,		19,	2,	0 },
	{ AMOVW,		C_REG,		C_NONE,		C_SAUTO,		20,	2,	0 },
	{ AMOVW,		C_REG,		C_NONE,		C_LAUTO,		34,	6,	0,	LTO },
	{ AMOVW,		C_REG,		C_NONE,		C_SEXT,		31,	4,	0 },
	{ AMOVW,		C_REG,		C_NONE,		C_SOREG,		21,	2,	0 },
	{ AMOVH,		C_REG,		C_NONE,		C_SEXT,		31,	4,	0 },
	{ AMOVH,		C_REG,		C_NONE,		C_SOREG,		21,	2,	0 },
	{ AMOVB,		C_REG,		C_NONE,		C_SEXT,		31,	4,	0 },
	{ AMOVB,		C_REG,		C_NONE,		C_SOREG,		21,	2,	0 },
	{ AMOVHU,	C_REG,		C_NONE,		C_SEXT,		31,	4,	0 },
	{ AMOVHU,	C_REG,		C_NONE,		C_SOREG,		21,	2,	0 },
	{ AMOVBU,	C_REG,		C_NONE,		C_SEXT,		31,	4,	0 },
	{ AMOVBU,	C_REG,		C_NONE,		C_SOREG,		21,	2,	0 },
	{ AMOVW,		C_REG,		C_NONE,		C_REG,		22,	2,	0 },
	{ AMOVB,		C_REG,		C_NONE,		C_REG,		23,	4,	0 },
	{ AMOVH,		C_REG,		C_NONE,		C_REG,		23,	4,	0 },
	{ AMOVBU,	C_REG,		C_NONE,		C_REG,		23,	4,	0 },
	{ AMOVHU,	C_REG,		C_NONE,		C_REG,		23,	4,	0 },
	{ AMOVH,		C_SEXT,		C_NONE,		C_REG,		32,	6,	0 },
	{ AMOVH,		C_SOREG,		C_NONE,		C_REG,		24,	4,	0 },
	{ AMOVB,		C_SEXT,		C_NONE,		C_REG,		32,	6,	0 },
	{ AMOVB,		C_SOREG,		C_NONE,		C_REG,		24,	4,	0 },
	{ AMOVW,		C_SACON,	C_NONE,		C_REG,		25,	2,	0 },
	{ AMOVW,		C_LACON,	C_NONE,		C_REG,		35,	4,	0 },
	{ AMOVW,		C_GACON,	C_NONE,		C_REG,		35,	4,	0,	LFROM },
	{ AMOVM,		C_LCON,		C_NONE,		C_REG,		26,	2,	0 },
	{ AMOVM,		C_REG,		C_NONE,		C_LCON,		27,	2,	0 },
	{ AMOVW,		C_LOREG,		C_NONE,		C_REG,		28,	4,	0 },
	{ AMOVH,		C_LOREG,		C_NONE,		C_REG,		28,	4,	0 },
	{ AMOVB,		C_LOREG,		C_NONE,		C_REG,		28,	4,	0 },
	{ AMOVHU,	C_LOREG,		C_NONE,		C_REG,		28,	4,	0 },
	{ AMOVBU,	C_LOREG,		C_NONE,		C_REG,		28,	4,	0 },
	{ AMOVW,		C_REG,		C_NONE,		C_LOREG,		29,	4,	0 },
	{ AMOVH,		C_REG,		C_NONE,		C_LOREG,		29,	4,	0 },
	{ AMOVB,		C_REG,		C_NONE,		C_LOREG,		29,	4,	0 },
	{ AMOVHU,	C_REG,		C_NONE,		C_LOREG,		29,	4,	0 },
	{ AMOVBU,	C_REG,		C_NONE,		C_LOREG,		29,	4,	0 },
	{ AMOVW,		C_GOREG,		C_NONE,		C_REG,		28,	4,	0,	LFROM },
	{ AMOVH,		C_GOREG,		C_NONE,		C_REG,		28,	4,	0,	LFROM },
	{ AMOVB,		C_GOREG,		C_NONE,		C_REG,		28,	4,	0,	LFROM },
	{ AMOVHU,	C_GOREG,		C_NONE,		C_REG,		28,	4,	0,	LFROM },
	{ AMOVBU,	C_GOREG,		C_NONE,		C_REG,		28,	4,	0,	LFROM },
	{ AMOVW,		C_REG,		C_NONE,		C_GOREG,		29,	4,	0,	LTO },
	{ AMOVH,		C_REG,		C_NONE,		C_GOREG,		29,	4,	0,	LTO },
	{ AMOVB,		C_REG,		C_NONE,		C_GOREG,		29,	4,	0,	LTO },
	{ AMOVHU,	C_REG,		C_NONE,		C_GOREG,		29,	4,	0,	LTO },
	{ AMOVBU,	C_REG,		C_NONE,		C_GOREG,		29,	4,	0,	LTO },
	{ AMOVW,		C_LEXT,		C_NONE,		C_REG,		30,	4,	0,	LFROM },
	{ AMOVH,		C_LEXT,		C_NONE,		C_REG,		32,	6,	0,	LFROM },
	{ AMOVB,		C_LEXT,		C_NONE,		C_REG,		32,	6,	0,	LFROM },
	{ AMOVHU,	C_LEXT,		C_NONE,		C_REG,		30,	4,	0,	LFROM },
	{ AMOVBU,	C_LEXT,		C_NONE,		C_REG,		30,	4,	0,	LFROM },
	{ AMOVW,		C_REG,		C_NONE,		C_LEXT,		31,	4,	0,	LTO },
	{ AMOVH,		C_REG,		C_NONE,		C_LEXT,		31,	4,	0,	LTO },
	{ AMOVB,		C_REG,		C_NONE,		C_LEXT,		31,	4,	0,	LTO },
	{ AMOVHU,	C_REG,		C_NONE,		C_LEXT,		31,	4,	0,	LTO },
	{ AMOVBU,	C_REG,		C_NONE,		C_LEXT,		31,	4,	0,	LTO },

	{ AXXX,		C_NONE,		C_NONE,		C_NONE,		0,	2,	0 },
};

#define OPCNTSZ	52
int opcount[OPCNTSZ];

// is this too pessimistic ?
int
brextra(Prog *p)
{
	int c;

	// +2 is for padding
	if(p-&gt;as == ATEXT)
		return 0-0+2;
	if(!isbranch(p))
		diag(&#34;bad op in brextra()&#34;);
	c = thumbaclass(&amp;p-&gt;to, p);
	switch(p-&gt;as){
		case AB:
			if(c != C_SBRA)
				return 0;
			return 8-2+2;
		case ABL:
			if(c != C_SBRA)
				return 0;
			return 14-4+2;
		case ABX:
			if(c == C_REG || c == C_HREG)
				return 0;
#ifdef CALLEEBX
			diag(&#34;ABX $I in brextra&#34;);
#endif
			if(c != C_SBRA)
				return 0;
			return 14-10+2;
		default:
			if(c == C_GBRA)
				return 0;
			if(c == C_LBRA)
				return 10-4+2;
			return 10-2+2;
	}
	return 0;
}

#define high(r)	((r)&gt;=8)

static int32
mv(Prog *p, int r, int off)
{
	int v, o;
	if(p != nil &amp;&amp; p-&gt;cond != nil){	// in literal pool
		v = p-&gt;cond-&gt;pc - p-&gt;pc - 4;
		if(p-&gt;cond-&gt;pc &amp; 3)
			diag(&#34;mv: bad literal pool alignment&#34;);
		if(v &amp; 3)
			v += 2;	// ensure M(4) offset
		mult(p, v, 4);
		off = v/4;
		numr(p, off, 0, 255);
		o = 0x9&lt;&lt;11;
	}
	else{
		numr(p, off, 0, 255);
		o = 0x4&lt;&lt;11;
	}
	o |= (r&lt;&lt;8) | off;
	return o;
}

static void
mvcon(Prog *p, int r, int c, int32 *o1, int32 *o2)
{
	int op = 0, n = 0;

	if(c &gt;= 0 &amp;&amp; c &lt;= 255)
		diag(&#34;bad c in mvcon&#34;);
	if(c &gt;= -255 &amp;&amp; c &lt; 0)	// mv, neg
		c = -c;
	else if(c &gt;= 256 &amp;&amp; c &lt;= 510){	// mv, add
		n = rand()%(511-c) + (c-255);
		c -= n;
		// n = c-255;
		// c = 255;
		op = AADD;
	}
	else{
		if(c &lt; 0)
			diag(&#34;-ve in mvcon&#34;);
		while(!(c &amp; 1)){
			n++;
			c &gt;&gt;= 1;
		}
		if(c &gt;= 0 &amp;&amp; c &lt;= 255)	// mv, lsl
			op = ASLL;
		else
			diag(&#34;bad shift in mvcon&#34;);
	}
	*o1 = mv(p, r, c);
	switch(op){
		case 0:
			*o2 = (1&lt;&lt;14) | (9&lt;&lt;6) | (r&lt;&lt;3) | r;
			break;
		case AADD:
			*o2 = (6&lt;&lt;11) | (r&lt;&lt;8) | n;
			break;
		case ASLL:
			*o2 = (n&lt;&lt;6) | (r&lt;&lt;3) | r;
			break;
	}
}

static int32
mvlh(int rs, int rd)
{
	int o = 0x46&lt;&lt;8;

	if(high(rs)){
		rs -= 8;
		o |= 1&lt;&lt;6;
	}
	if(high(rd)){
		rd -= 8;
		o |= 1&lt;&lt;7;
	}
	o |= (rs&lt;&lt;3) | rd;
	return o;
}

void
thumbbuildop()
{
	int i, n, r;
	Optab *optab = thumboptab;
	Oprang *oprange = thumboprange;

	for(n=0; optab[n].as != AXXX; n++)
		;
	qsort(optab, n, sizeof(optab[0]), ocmp);
	for(i=0; i&lt;n; i++) {
		r = optab[i].as;
		oprange[r].start = optab+i;
		while(optab[i].as == r)
			i++;
		oprange[r].stop = optab+i;
		i--;

		switch(r)
		{
		default:
			break;
		case ABEQ:
			oprange[ABNE] = oprange[r];
			oprange[ABCS] = oprange[r];
			oprange[ABHS] = oprange[r];
			oprange[ABCC] = oprange[r];
			oprange[ABLO] = oprange[r];
			oprange[ABMI] = oprange[r];
			oprange[ABPL] = oprange[r];
			oprange[ABVS] = oprange[r];
			oprange[ABVC] = oprange[r];
			oprange[ABHI] = oprange[r];
			oprange[ABLS] = oprange[r];
			oprange[ABGE] = oprange[r];
			oprange[ABLT] = oprange[r];
			oprange[ABGT] = oprange[r];
			oprange[ABLE] = oprange[r];
			break;
		case AMVN:
			oprange[AADC] = oprange[r];
			oprange[ASBC] = oprange[r];
			oprange[AMUL] = oprange[r];
			oprange[AAND] = oprange[r];
			oprange[AEOR] = oprange[r];
			oprange[AORR] = oprange[r];
			oprange[ABIC] = oprange[r];
			oprange[AMULU] = oprange[r];
			break;
		case ACMN:
			oprange[ATST] = oprange[r];
			break;
		case ASRL:
			oprange[ASRA] = oprange[r];
			oprange[ASLL] = oprange[r];
			break;
		case AADD:
			oprange[ASUB] = oprange[r];
			break;
		}
	}
}

void
thumbasmout(Prog *p, Optab *o)
{
	int32 o1, o2, o3, o4, o5, o6, o7, v;
	int r, rf, rt;

	rf = p-&gt;from.reg;
	rt = p-&gt;to.reg;
	r = p-&gt;reg;
	o1 = o2 = o3 = o4 = o5 = o6 = o7 = 0;
if(debug[&#39;P&#39;]) print(&#34;%ulx: %P	type %d %d\n&#34;, (uint32)(p-&gt;pc), p, o-&gt;type, p-&gt;align);
	opcount[o-&gt;type] += o-&gt;size;
	switch(o-&gt;type) {
	default:
		diag(&#34;unknown asm %d&#34;, o-&gt;type);
		prasm(p);
		break;
	case 0:		/* pseudo ops */
if(debug[&#39;G&#39;]) print(&#34;%ulx: %s: thumb\n&#34;, (uint32)(p-&gt;pc), p-&gt;from.sym-&gt;name);
		break;
	case 1:		/* op R, -, R or op R, R, - */
		o1 = thumboprr(p-&gt;as);
		if(rt == NREG)
			rt = r;
		lowreg(p, rf);
		lowreg(p, rt);
		o1 |= (0x10&lt;&lt;10) | (rf&lt;&lt;3) | rt;
		break;
	case 2:		/* add/sub R, R, R or add/sub R, -, R */
		o1 = p-&gt;as == AADD ? 0x0&lt;&lt;9 : 0x1&lt;&lt;9;
		if(r == NREG)
			r = rt;
		lowreg(p, rf);
		lowreg(p, r);
		lowreg(p, rt);
		o1 |= (0x6&lt;&lt;10) | (rf&lt;&lt;6) | (r&lt;&lt;3) | rt;
		break;
	case 3:		/* add/sub $I, R, R or add/sub $I, -, R */
		thumbaclass(&amp;p-&gt;from, p);
		o1 = p-&gt;as == AADD ? 0x0&lt;&lt;9 : 0x1&lt;&lt;9;
		if(r == NREG)
			r = rt;
		numr(p, instoffset, 0, 7);
		lowreg(p, r);
		lowreg(p, rt);
		o1 |= (0x7&lt;&lt;10) | (instoffset&lt;&lt;6) | (r&lt;&lt;3) | rt;
		break;
	case 4:		/* shift $I, R, R or shift $I, -, R */
		thumbaclass(&amp;p-&gt;from, p);
		if(instoffset &lt; 0)
			diag(&#34;negative shift in thumbasmout&#34;);
		instoffset %= 32;
		o1 = thumbopri(p-&gt;as);
		if(r == NREG)
			r = rt;
		numr(p, instoffset, 0, 31);
		lowreg(p, r);
		lowreg(p, rt);
		o1 |= (0x0&lt;&lt;13) | (instoffset&lt;&lt;6) | (r&lt;&lt;3) | rt;
		break;
	case 5:		/* add/sub/mov $I, -, R or cmp $I, R, - */
		thumbaclass(&amp;p-&gt;from, p);
		o1 = thumbopri(p-&gt;as);
		if(rt == NREG)
			rt = r;
		numr(p, instoffset, 0, 255);
		lowreg(p, rt);
		o1 |= (0x1&lt;&lt;13) | (rt&lt;&lt;8) | instoffset;
		break;
	case 6:		/* add $I, PC/SP, R */
		if(p-&gt;as == ASUB)
			diag(&#34;subtract in add $I, PC/SP, R&#34;);
		thumbaclass(&amp;p-&gt;from, p);
		o1 = r == REGSP ? 0x1&lt;&lt;11 : 0x0&lt;&lt;11;
		numr(p, instoffset, 0, 255);
		regis(p, r, REGSP, REGPC);
		lowreg(p, rt);
		o1 |= (0xa&lt;&lt;12) | (rt&lt;&lt;8) | instoffset;
		break;
	case 7:		/* add, sub $I, SP */
		thumbaclass(&amp;p-&gt;from, p);
		o1 = p-&gt;as == AADD ? 0x0&lt;&lt;7 : 0x1&lt;&lt;7;
		numr(p, instoffset, 0, 508);
		mult(p, instoffset, 4);
		regis(p, rt, REGSP, REGSP);
		o1 |= (0xb0&lt;&lt;8) | (instoffset&gt;&gt;2);
		break;
	case 8:		/* add/mov/cmp R, R where at least 1 reg is high */
		o1 = 0;
		if(rt == NREG)
			rt = r;
		if(high(rf)){
			o1 |= 1&lt;&lt;6;
			rf -= 8;
		}
		if(high(rt)){
			o1 |= 2&lt;&lt;6;
			rt -= 8;
		}
		if(o1 == 0)
			diag(&#34;no high register(%P)&#34;, p);
		o1 |= thumbophh(p-&gt;as);
		o1 |= (0x11&lt;&lt;10) | (rf&lt;&lt;3) | rt;
		break;
	case 9:		/* B	$I */
		thumbaclass(&amp;p-&gt;to, p);
		numr(p, instoffset, -2048, 2046);
		o1 = (0x1c&lt;&lt;11) | ((instoffset&gt;&gt;1)&amp;0x7ff);
		break;
	case 10:		/* Bcc $I */
		thumbaclass(&amp;p-&gt;to, p);
		numr(p, instoffset, -256, 254);
		o1 = thumbopbra(p-&gt;as);
		o1 |= (0xd&lt;&lt;12) | ((instoffset&gt;&gt;1)&amp;0xff);
		break;
	case 11:		/* BL $I */
		thumbaclass(&amp;p-&gt;to, p);
		numr(p, instoffset, -4194304, 4194302);
		o1 = (0x1e&lt;&lt;11) | ((instoffset&gt;&gt;12)&amp;0x7ff);
		o2 = (0x1f&lt;&lt;11) | ((instoffset&gt;&gt;1)&amp;0x7ff);
		break;
	case 12:		/* BX $I */
#ifdef CALLEEBX
		diag(&#34;BX $I case&#34;);
#endif
		thumbaclass(&amp;p-&gt;to, p);
		if(p-&gt;to.sym-&gt;thumb)
			instoffset  |= 1;	// T bit
		o1 = mvlh(REGPC, REGTMPT);
		o2 = (0x6&lt;&lt;11) | (REGTMPT&lt;&lt;8) | 7;	// add 7, RTMP	(T bit + PC offset)
		o3 = mvlh(REGTMPT, REGLINK);
		o4 = mv(nil, REGTMPT, instoffset);
		o5 = (0x11c&lt;&lt;6) | (REGTMPT&lt;&lt;3);
		// o1 = mv(nil, REGTMPT, v);
		// o2 = (0x11b&lt;&lt;6) | (REGPC&lt;&lt;3) | REGLINK;
		// o3 = (0x11c&lt;&lt;6) | (REGTMPT&lt;&lt;3);
		break;
	case 13:		/* B O(R)  */
		diag(&#34;B O(R)&#34;);
		break;
	case 14:		/* BL O(R) */
		diag(&#34;BL O(R)&#34;);
		break;
	case 15:		/* BX R */
		o1 = mvlh(REGPC, REGTMPT);
		o2 = (0x6&lt;&lt;11) | (REGTMPT&lt;&lt;8) | 5;	// add 5, RTMP (T bit + PC offset)
		o3 = mvlh(REGTMPT, REGLINK);
		o4 = 0;
		if(high(rt)){
			rt -= 8;
			o4 |= 1&lt;&lt;6;
		}
		o4 |= (0x8e&lt;&lt;7) | (rt&lt;&lt;3);
		// o1 = (0x11c&lt;&lt;6) | (rt&lt;&lt;3);
		break;
	case 16:		/* SWI $I */
		thumbaclass(&amp;p-&gt;to, p);
		numr(p, instoffset, 0, 255);
		o1 = (0xdf&lt;&lt;8) | instoffset;
		break;
	case 17:		/* AWORD */
		thumbaclass(&amp;p-&gt;to, p);
		o1 = instoffset&amp;0xffff;
		o2 = (instoffset&gt;&gt;16)&amp;0xffff;
		break;
	case 18:		/* AMOVW O(SP), R and AMOVW O(PC), R */
		thumbaclass(&amp;p-&gt;from, p);
		rf = o-&gt;param;
		o1 = rf == REGSP ? 0x13&lt;&lt;11 : 0x9&lt;&lt;11;
		regis(p, rf, REGSP, REGPC);
		lowreg(p, rt);
		mult(p, instoffset, 4);
		numr(p, instoffset/4, 0, 255);
		o1 |= (rt&lt;&lt;8) | (instoffset/4);
		break;
	case 19:		/* AMOVW... O(R), R */
		thumbaclass(&amp;p-&gt;from, p);
		o1 = thumbopmv(p-&gt;as, 1);
		v = 4;
		if(p-&gt;as == AMOVHU)
			v = 2;
		else if(p-&gt;as == AMOVBU)
			v = 1;
		mult(p, instoffset, v);
		lowreg(p, rf);
		lowreg(p, rt);
		numr(p, instoffset/v, 0, 31);
		o1 |= ((instoffset/v)&lt;&lt;6) | (rf&lt;&lt;3) | rt;
		break;
	case 20:		/* AMOVW R, O(SP) */
		thumbaclass(&amp;p-&gt;to, p);
		o1 = 0x12&lt;&lt;11;
		if(rt != NREG) regis(p, rt, REGSP, REGSP);
		lowreg(p, rf);
		mult(p, instoffset, 4);
		numr(p, instoffset/4, 0, 255);
		o1 |= (rf&lt;&lt;8) | (instoffset/4);
		break;
	case 21:		/* AMOVW... R, O(R) */
		thumbaclass(&amp;p-&gt;to, p);
		o1 = thumbopmv(p-&gt;as, 0);
		v = 4;
		if(p-&gt;as == AMOVHU || p-&gt;as == AMOVH)
			v = 2;
		else if(p-&gt;as == AMOVBU || p-&gt;as == AMOVB)
			v = 1;
		lowreg(p, rf);
		lowreg(p, rt);
		mult(p, instoffset, v);
		numr(p, instoffset/v, 0, 31);
		o1 |= ((instoffset/v)&lt;&lt;6) | (rt&lt;&lt;3) | rf;
		break;
	case 22:		/* AMOVW R, R -&gt; ASLL $0, R, R */
		o1 = thumbopri(ASLL);
		lowreg(p, rf);
		lowreg(p, rt);
		o1 |= (0x0&lt;&lt;13) | (rf&lt;&lt;3) | rt;
		break;
	case 23:		/* AMOVB/AMOVH/AMOVBU/AMOVHU R, R */
		o1 = thumbopri(ASLL);
		o2 = p-&gt;as == AMOVB || p-&gt;as == AMOVH ? thumbopri(ASRA) : thumbopri(ASRL);
		v = p-&gt;as == AMOVB || p-&gt;as == AMOVBU ? 24 : 16;
		lowreg(p, rf);
		lowreg(p, rt);
		o1 |= (0x0&lt;&lt;13) | (v&lt;&lt;6) | (rf&lt;&lt;3) | rt;
		o2 |= (0x0&lt;&lt;13) | (v&lt;&lt;6) | (rt&lt;&lt;3) | rt;
		break;
	case 24:	/* AMOVH/AMOVB O(R), R -&gt; AMOVH/AMOVB [R, R], R */
		thumbaclass(&amp;p-&gt;from, p);
		lowreg(p, rf);
		lowreg(p, rt);
		if(rf == rt)
			r = REGTMPT;
		else
			r = rt;
		if(p-&gt;as == AMOVB)
			numr(p, instoffset, 0, 31);
		else{
			mult(p, instoffset, 2);
			numr(p, instoffset, 0, 62);
		}
		o1 = mv(p, r, instoffset);
		o2 = p-&gt;as == AMOVH ? 0x2f&lt;&lt;9 : 0x2b&lt;&lt;9;
		o2 |= (r&lt;&lt;6) | (rf&lt;&lt;3) | rt;
		break;
	case 25:	/* MOVW $sacon, R */
		thumbaclass(&amp;p-&gt;from, p);
// print(&#34;25: %d %d %d %d\n&#34;, instoffset, rf, r, rt);
		if(rf == NREG)
			rf = REGSP;
		lowreg(p, rt);
		if(rf == REGSP){
			mult(p, instoffset, 4);
			numr(p, instoffset&gt;&gt;2, 0, 255);
			o1 = (0x15&lt;&lt;11) | (rt&lt;&lt;8) | (instoffset&gt;&gt;2);	// add $O, SP, R
		}
		else if(rf == rt){
			numr(p, instoffset, 0, 255);
			o1 = (0x6&lt;&lt;11) | (rt&lt;&lt;8) | instoffset;		// add $O, R
		}
		else{
			lowreg(p, rf);
			numr(p, instoffset, 0, 7);
			o1 = (0xe&lt;&lt;9) | (instoffset&lt;&lt;6) | (rf&lt;&lt;3) | rt;	// add $O, Rs, Rd
		}
		break;
	case 26:	/* AMOVM $c, oreg -&gt; stmia */
		lowreg(p, rt);
		numr(p, p-&gt;from.offset, -256, 255);
		o1 = (0x18&lt;&lt;11) | (rt&lt;&lt;8) | (p-&gt;from.offset&amp;0xff);
		break;
	case 27:	/* AMOVM oreg, $c -&gt;ldmia */
		lowreg(p, rf);
		numr(p, p-&gt;to.offset, -256, 256);
		o1 = (0x19&lt;&lt;11) | (rf&lt;&lt;8) | (p-&gt;to.offset&amp;0xff);
		break;
	case 28:	/* AMOV* O(R), R -&gt; AMOV* [R, R], R 	(offset large)	*/
		thumbaclass(&amp;p-&gt;from, p);
		lowreg(p, rf);
		lowreg(p, rt);
		if(rf == rt)
			r = REGTMPT;
		else
			r = rt;
		o1 = mv(p, r, instoffset);
		o2 = thumboprrr(p-&gt;as, 1);
		o2 |= (r&lt;&lt;6) | (rf&lt;&lt;3) | rt;
		break;
	case 29:	/* AMOV* R, O(R) -&gt; AMOV* R, [R, R]	(offset large)	*/
		thumbaclass(&amp;p-&gt;to, p);
		lowreg(p, rf);
		lowreg(p, rt);
		if(rt == REGTMPT){	// used as tmp reg
			if(instoffset &gt;= 0 &amp;&amp; instoffset &lt;= 255){
				o1 = (1&lt;&lt;13) | (2&lt;&lt;11) | (rt&lt;&lt;8) | instoffset;	// add $O, R7
				o2 = thumbopirr(p-&gt;as, 0);
				o2 |= (0&lt;&lt;6) | (rt&lt;&lt;3) | rf;					// mov* R, 0(R)
			}
			else
				diag(&#34;big offset - case 29&#34;);
		}
		else{
			o1 = mv(p, REGTMPT, instoffset);
			o2 = thumboprrr(p-&gt;as, 0);
			o2 |= (REGTMPT&lt;&lt;6) | (rt&lt;&lt;3) | rf;
		}
		break;
	case 30:		/* AMOVW... *addr, R */
		thumbaclass(&amp;p-&gt;from, p);
		o1 = mv(p, rt, instoffset);		// MOV addr, rtmp
		o2 = thumbopmv(p-&gt;as, 1);
		lowreg(p, rt);
		o2 |= (rt&lt;&lt;3) | rt;			// MOV* 0(rtmp), R
		break;
	case 31:		/* AMOVW... R, *addr */
		thumbaclass(&amp;p-&gt;to, p);
		o1 = mv(p, REGTMPT, instoffset);
		o2 = thumbopmv(p-&gt;as, 0);
		lowreg(p, rf);
		o2 |= (REGTMPT&lt;&lt;3) | rf;
		break;
	case 32:	/* AMOVH/AMOVB *addr, R -&gt; AMOVH/AMOVB [R, R], R */
		thumbaclass(&amp;p-&gt;from, p);
		o1 = mv(p, rt, instoffset);
		lowreg(p, rt);
		o2 = mv(nil, REGTMPT, 0);
		o3 = p-&gt;as == AMOVH ? 0x2f&lt;&lt;9 : 0x2b&lt;&lt;9;
		o3 |= (REGTMPT&lt;&lt;6) | (rt&lt;&lt;3) | rt;
		break;
	case 33:	/* AMOVW O(SP), R	(O large) */
		thumbaclass(&amp;p-&gt;from, p);
		lowreg(p, rt);
		o1 = mv(p, rt, instoffset);
		o2 = (0x111&lt;&lt;6) | (REGSP-8)&lt;&lt;3 | rt;	// add SP, rt
		o3 = thumbopmv(p-&gt;as, 1);
		o3 |= (rt&lt;&lt;3) | rt;
		break;
	case 34:	/* AMOVW R, O(SP)	(O large) */
		thumbaclass(&amp;p-&gt;to, p);
		lowreg(p, rf);
		o1 = mv(p, REGTMPT, instoffset);
		o2 = (0x111&lt;&lt;6) | (REGSP-8)&lt;&lt;3 | REGTMPT;	// add SP, REGTMP
		o3 = thumbopmv(p-&gt;as, 0);
		o3 |= (REGTMPT&lt;&lt;3) | rf;
		break;
	case 35:	/* AMOVW $lacon, R */
		thumbaclass(&amp;p-&gt;from, p);
		lowreg(p, rt);
		if(rf == NREG)
			rf = REGSP;
		if(rf == rt)
			rf = r = REGTMPT;
		else
			r = rt;
// print(&#34;35: io=%d rf=%d rt=%d\n&#34;, instoffset, rf, rt);
		o1 = mv(p, r, instoffset);		// mov O, Rd
		if(high(rf))
			o2 = (0x44&lt;&lt;8) | (0x1&lt;&lt;6) | ((rf-8)&lt;&lt;3) | rt;	// add Rs, Rd
		else
			o2 = (0x6&lt;&lt;10) | (rf&lt;&lt;6) | (rt&lt;&lt;3) | rt;		// add Rs, Rd
		break;
	case 36:	/* AADD/ASUB $i, r, r when $i too big */
		thumbaclass(&amp;p-&gt;from, p);
		lowreg(p, r);
		lowreg(p, rt);
		o1 = mv(p, REGTMPT, instoffset);
		o2 = p-&gt;as == AADD ? 0xc&lt;&lt;9 : 0xd&lt;&lt;9;
		o2 |= (REGTMPT&lt;&lt;6) | (r&lt;&lt;3) | rt;
		break;
	case 37:	/* AADD/ASUB $i, r when $i too big */
		thumbaclass(&amp;p-&gt;from, p);
		lowreg(p, rt);
		o1 = mv(p, REGTMPT, instoffset);
		o2 = p-&gt;as == AADD ? 0xc&lt;&lt;9 : 0xd&lt;&lt;9;
		o2 |= (REGTMPT&lt;&lt;6) | (rt&lt;&lt;3) | rt;
		break;
	case 38:	/* AMOVW $i, r when $i too big */
		thumbaclass(&amp;p-&gt;from, p);
		lowreg(p, rt);
		o1 = mv(p, rt, instoffset);
		break;
	case 39:	/* ACMP $i, r when $i too big */
		thumbaclass(&amp;p-&gt;from, p);
		lowreg(p, r);
		o1 = mv(p, REGTMPT, instoffset);
		o2 = (0x10a&lt;&lt;6) | (REGTMPT&lt;&lt;3) | r;
		break;
	case 40:		/* add, sub $I, SP when $I large*/
		thumbaclass(&amp;p-&gt;from, p);
		if(p-&gt;as == ASUB)
			instoffset = -instoffset;
		o1 = mv(p, REGTMPT, instoffset);
		o2 = (0x112&lt;&lt;6) | (REGTMPT&lt;&lt;3) | (REGSP-8);
		regis(p, rt, REGSP, REGSP);
		break;
	case	41:		/* BL LBRA */
		thumbaclass(&amp;p-&gt;to, p);
		o1 = (0x9&lt;&lt;11) | (REGTMPT&lt;&lt;8);	// mov 0(pc), r7
		o2 = mvlh(REGTMPT, REGPC);		// mov r7, pc
		o3 = instoffset&amp;0xffff;			// $lab
		o4 = (instoffset&gt;&gt;16)&amp;0xffff;
		break;
	case 42:		/* Bcc GBRA */
		thumbaclass(&amp;p-&gt;to, p);
		o1 = (0xd&lt;&lt;12) | thumbopbra(relinv(p-&gt;as)) | (6&gt;&gt;1);		// bccnot
		// ab lbra
		o2 = (0x9&lt;&lt;11) | (REGTMPT&lt;&lt;8);	// mov 0(pc), r7
		o3 = mvlh(REGTMPT, REGPC);		// mov r7, pc
		o4 = instoffset&amp;0xffff;			// $lab
		o5 = (instoffset&gt;&gt;16)&amp;0xffff;
		break;
	case 43:		/* BL LBRA */
		thumbaclass(&amp;p-&gt;to, p);
		o1 = mvlh(REGPC, REGTMPT);						// mov pc, r7
		o2 = (0x6&lt;&lt;11) | (REGTMPT&lt;&lt;8) | 10;				// add 10, r7
		o3 = mvlh(REGTMPT, REGLINK);					// mov r7, lr
		o4 = (0x9&lt;&lt;11) | (REGTMPT&lt;&lt;8);					// mov o(pc), r7
		o5 = mvlh(REGTMPT, REGPC);						// mov r7, pc
		o6 = instoffset&amp;0xffff;							// $lab
		o7 = (instoffset&gt;&gt;16)&amp;0xffff;
		break;
	case 44:		/* BX LBRA */
#ifdef CALLEEBX
		diag(&#34;BX LBRA case&#34;);
#endif
		thumbaclass(&amp;p-&gt;to, p);
		if(p-&gt;to.sym-&gt;thumb)
			instoffset  |= 1;	// T bit
		o1 = mvlh(REGPC, REGTMPT);						// mov pc, r7
		o2 = (0x6&lt;&lt;11) | (REGTMPT&lt;&lt;8) | 11;				// add 11, r7
		o3 = mvlh(REGTMPT, REGLINK);					// mov r7, lr
		o4 = (0x9&lt;&lt;11) | (REGTMPT&lt;&lt;8);					// mov o(pc), r7
		o5 = (0x11c&lt;&lt;6) | (REGTMPT&lt;&lt;3);					// bx r7
		o6 = instoffset&amp;0xffff;							// $lab
		o7 = (instoffset&gt;&gt;16)&amp;0xffff;
		break;
	case 45:	/* BX R when returning from fn */
		o1 = 0;
		if(high(rt)){
			rt -= 8;
			o1 |= 1&lt;&lt;6;
		}
		o1 |= (0x8e&lt;&lt;7) | (rt&lt;&lt;3);
		break;
	case 46:		/* Bcc LBRA */
		thumbaclass(&amp;p-&gt;to, p);
		o1 = (0xd&lt;&lt;12) | thumbopbra(relinv(p-&gt;as)) | (0&gt;&gt;1);		// bccnot
		// ab lbra
		instoffset -= 2;
		numr(p, instoffset, -2048, 2046);
		o2 = (0x1c&lt;&lt;11) | ((instoffset&gt;&gt;1)&amp;0x7ff);
		break;
	case 47:	/* mov $i, R where $i can be built */
		thumbaclass(&amp;p-&gt;from, p);
		mvcon(p, rt, instoffset, &amp;o1, &amp;o2);
		break;
	case 48: /* ACMP $i, r when $i built up */
		thumbaclass(&amp;p-&gt;from, p);
		lowreg(p, r);
		mvcon(p, REGTMPT, instoffset, &amp;o1, &amp;o2);
		o3 = (0x10a&lt;&lt;6) | (REGTMPT&lt;&lt;3) | r;
		break;
	case 49:	/* AADD $i, r, r when $i is between 0 and 255 - could merge with case 36 */
		thumbaclass(&amp;p-&gt;from, p);
		lowreg(p, r);
		lowreg(p, rt);
		numr(p, instoffset, 0, 255);
		o1 = mv(p, REGTMPT, instoffset);
		o2 = p-&gt;as == AADD ? 0xc&lt;&lt;9 : 0xd&lt;&lt;9;
		o2 |= (REGTMPT&lt;&lt;6) | (r&lt;&lt;3) | rt;
		break;
	case 50:		/* ADWORD */
		thumbaclass(&amp;p-&gt;from, p);
		o1 = instoffset&amp;0xffff;
		o2 = (instoffset&gt;&gt;16)&amp;0xffff;
		thumbaclass(&amp;p-&gt;to, p);
		o3 = instoffset&amp;0xffff;
		o4 = (instoffset&gt;&gt;16)&amp;0xffff;
		break;
	case 51:	/* BL r */
		o1 = mvlh(REGPC, REGLINK);	// mov pc, lr
		o2 = mvlh(rt, REGPC);		// mov r, pc
		break;
	}

	v = p-&gt;pc;
	switch(o-&gt;size) {
	default:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34; %.8lux:\t\t%P\n&#34;, v, p);
		break;
	case 2:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34; %.8lux: %.8lux\t%P\n&#34;, v, o1, p);
		hputl(o1);
		break;
	case 4:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34; %.8lux: %.8lux %.8lux\t%P\n&#34;, v, o1, o2, p);
		hputl(o1);
		hputl(o2);
		break;
	case 6:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34;%.8lux: %.8lux %.8lux %.8lux\t%P\n&#34;, v, o1, o2, o3, p);
		hputl(o1);
		hputl(o2);
		hputl(o3);
		break;
	case 8:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34;%.8lux: %.8lux %.8lux %.8lux %.8lux\t%P\n&#34;, v, o1, o2, o3, o4, p);
		hputl(o1);
		hputl(o2);
		hputl(o3);
		hputl(o4);
		break;
	case 10:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34;%.8lux: %.8lux %.8lux %.8lux %.8lux %.8lux\t%P\n&#34;, v, o1, o2, o3, o4, o5, p);
		hputl(o1);
		hputl(o2);
		hputl(o3);
		hputl(o4);
		hputl(o5);
		break;
	case 12:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34;%.8lux: %.8lux %.8lux %.8lux %.8lux %.8lux %.8lux\t%P\n&#34;, v, o1, o2, o3, o4, o5, o6, p);
		hputl(o1);
		hputl(o2);
		hputl(o3);
		hputl(o4);
		hputl(o5);
		hputl(o6);
		break;
	case 14:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34;%.8lux: %.8lux %.8lux %.8lux %.8lux %.8lux %.8lux %.8lux\t%P\n&#34;, v, o1, o2, o3, o4, o5, o6, o7, p);
		hputl(o1);
		hputl(o2);
		hputl(o3);
		hputl(o4);
		hputl(o5);
		hputl(o6);
		hputl(o7);
		break;
	}
	if(debug[&#39;G&#39;]){
		if(o-&gt;type == 17){
			print(&#34;%lx:	word %ld\n&#34;, p-&gt;pc, (o2&lt;&lt;16)+o1);
			return;
		}
		if(o-&gt;type == 50){
			print(&#34;%lx:	word %ld\n&#34;, p-&gt;pc, (o2&lt;&lt;16)+o1);
			print(&#34;%lx:	word %ld\n&#34;, p-&gt;pc, (o4&lt;&lt;16)+o3);
			return;
		}
		if(o-&gt;size &gt; 0) dis(o1, p-&gt;pc);
		if(o-&gt;size &gt; 2) dis(o2, p-&gt;pc+2);
		if(o-&gt;size &gt; 4) dis(o3, p-&gt;pc+4);
		if(o-&gt;size &gt; 6) dis(o4, p-&gt;pc+6);
		if(o-&gt;size &gt; 8) dis(o5, p-&gt;pc+8);
		if(o-&gt;size &gt; 10) dis(o6, p-&gt;pc+10);
		if(o-&gt;size &gt; 12) dis(o7, p-&gt;pc+12);
		// if(o-&gt;size &gt; 14) dis(o8, p-&gt;pc+14);
	}
}

static int32
thumboprr(int a)
{
	switch(a) {
	case AMVN:	return 0xf&lt;&lt;6;
	case ACMP:	return 0xa&lt;&lt;6;
	case ACMN:	return 0xb&lt;&lt;6;
	case ATST:	return 0x8&lt;&lt;6;
	case AADC:	return 0x5&lt;&lt;6;
	case ASBC:	return 0x6&lt;&lt;6;
	case AMUL:
	case AMULU:	return 0xd&lt;&lt;6;
	case AAND:	return 0x0&lt;&lt;6;
	case AEOR:	return 0x1&lt;&lt;6;
	case AORR:	return 0xc&lt;&lt;6;
	case ABIC:	return 0xe&lt;&lt;6;
	case ASRL:	return 0x3&lt;&lt;6;
	case ASRA:	return 0x4&lt;&lt;6;
	case ASLL:	return 0x2&lt;&lt;6;
	}
	diag(&#34;bad thumbop oprr %d&#34;, a);
	prasm(curp);
	return 0;
}

static int32
thumbopirr(int a, int ld)
{
	if(ld)
		diag(&#34;load in thumbopirr&#34;);
	switch(a){
		case AMOVW:	return 0xc&lt;&lt;11;
		case AMOVH:
		case AMOVHU:	return 0x10&lt;&lt;11;
		case AMOVB:
		case AMOVBU:	return 0xe&lt;&lt;11;
	}
	return 0;
}

static int32
thumboprrr(int a, int ld)
{
	if(ld){
		switch(a){
		case AMOVW:	return 0x2c&lt;&lt;9;
		case AMOVH:	return 0x2f&lt;&lt;9;
		case AMOVB:	return 0x2b&lt;&lt;9;
		case AMOVHU:	return 0x2d&lt;&lt;9;
		case AMOVBU:	return 0x2e&lt;&lt;9;
		}
	}
	else{
		switch(a){
		case AMOVW:	return 0x28&lt;&lt;9;
		case AMOVHU:
		case AMOVH:	return 0x29&lt;&lt;9;
		case AMOVBU:
		case AMOVB:	return 0x2a&lt;&lt;9;
		}
	}
	diag(&#34;bad thumbop oprrr %d&#34;, a);
	prasm(curp);
	return 0;
}

static int32
thumbopri(int a)
{
	switch(a) {
	case ASRL:	return 0x1&lt;&lt;11;
	case ASRA:	return 0x2&lt;&lt;11;
	case ASLL:	return 0x0&lt;&lt;11;
	case AADD:	return 0x2&lt;&lt;11;
	case ASUB:	return 0x3&lt;&lt;11;
	case AMOVW:	return 0x0&lt;&lt;11;
	case ACMP:	return 0x1&lt;&lt;11;
	}
	diag(&#34;bad thumbop opri %d&#34;, a);
	prasm(curp);
	return 0;
}

static int32
thumbophh(int a)
{
	switch(a) {
	case AADD:	return 0x0&lt;&lt;8;
	case AMOVW:	return 0x2&lt;&lt;8;
	case ACMP:	return 0x1&lt;&lt;8;
	}
	diag(&#34;bad thumbop ophh %d&#34;, a);
	prasm(curp);
	return 0;
}

static int32
thumbopbra(int a)
{
	switch(a) {
	case ABEQ:	return 0x0&lt;&lt;8;
	case ABNE:	return 0x1&lt;&lt;8;
	case ABCS:	return 0x2&lt;&lt;8;
	case ABHS:	return 0x2&lt;&lt;8;
	case ABCC:	return 0x3&lt;&lt;8;
	case ABLO:	return 0x3&lt;&lt;8;
	case ABMI:	return 0x4&lt;&lt;8;
	case ABPL:	return 0x5&lt;&lt;8;
	case ABVS:	return 0x6&lt;&lt;8;
	case ABVC:	return 0x7&lt;&lt;8;
	case ABHI:	return 0x8&lt;&lt;8;
	case ABLS:	return 0x9&lt;&lt;8;
	case ABGE:	return 0xa&lt;&lt;8;
	case ABLT:	return 0xb&lt;&lt;8;
	case ABGT:	return 0xc&lt;&lt;8;
	case ABLE:	return 0xd&lt;&lt;8;
	}
	diag(&#34;bad thumbop opbra %d&#34;, a);
	prasm(curp);
	return 0;
}

static int32
thumbopmv(int a, int ld)
{
	switch(a) {
	case AMOVW: 	return (ld ? 0xd : 0xc)&lt;&lt;11;
	case AMOVH:
	case AMOVHU:	return (ld ? 0x11: 0x10)&lt;&lt;11;
	case AMOVB:
	case AMOVBU:	return (ld ? 0xf : 0xe)&lt;&lt;11;
	}
	diag(&#34;bad thumbop opmv %d&#34;, a);
	prasm(curp);
	return 0;
}

static void
lowreg(Prog *p, int r)
{
	if(high(r))
		diag(&#34;high reg [%P]&#34;, p);
}

static void
mult(Prog *p, int n, int m)
{
	if(m*(n/m) != n)
		diag(&#34;%d not M(%d) [%P]&#34;, n, m, p);
}

static void
numr(Prog *p, int n, int min, int max)
{
	if(n &lt; min || n &gt; max)
		diag(&#34;%d not in %d-%d [%P]&#34;, n, min, max, p);
}

static void
regis(Prog *p, int r, int r1, int r2)
{
	if(r != r1 &amp;&amp; r != r2)
		diag(&#34;reg %d not %d or %d [%P]&#34;, r, r1, r2, p);
}

void
hputl(int n)
{
	cbp[1] = n&gt;&gt;8;
	cbp[0] = n;
	cbp += 2;
	cbc -= 2;
	if(cbc &lt;= 0)
		cflush();
}

void
thumbcount()
{
	int i, c = 0, t = 0;

	for (i = 0; i &lt; OPCNTSZ; i++)
		t += opcount[i];
	if(t == 0)
		return;
	for (i = 0; i &lt; OPCNTSZ; i++){
		c += opcount[i];
		print(&#34;%d:	%d %d %d%%\n&#34;, i, opcount[i], c, (opcount[i]*100+t/2)/t);
	}
}

char *op1[] = { &#34;lsl&#34;, &#34;lsr&#34;, &#34;asr&#34; };
char *op2[] = { &#34;add&#34;, &#34;sub&#34; };
char *op3[] = { &#34;movw&#34;, &#34;cmp&#34;, &#34;add&#34;, &#34;sub&#34; };
char *op4[] = { &#34;and&#34;, &#34;eor&#34;, &#34;lsl&#34;, &#34;lsr&#34;, &#34;asr&#34;, &#34;adc&#34;, &#34;sbc&#34;, &#34;ror&#34;,
		        &#34;tst&#34;, &#34;neg&#34;, &#34;cmp&#34;, &#34;cmpn&#34;, &#34;or&#34;, &#34;mul&#34;, &#34;bitc&#34;, &#34;movn&#34; };
char *op5[] = { &#34;add&#34;, &#34;cmp&#34;, &#34;movw&#34;, &#34;bx&#34; };
char *op6[] = { &#34;smovw&#34;, &#34;smovh&#34;, &#34;smovb&#34;, &#34;lmovb&#34;, &#34;lmovw&#34;, &#34;lmovhu&#34;, &#34;lmovbu&#34;, &#34;lmovh&#34; };
char *op7[] = { &#34;smovw&#34;, &#34;lmovw&#34;, &#34;smovb&#34;, &#34;lmovbu&#34; };
char *op8[] = { &#34;smovh&#34;, &#34;lmovhu&#34; };
char *op9[] = { &#34;smovw&#34;, &#34;lmovw&#34; };
char *op10[] = { &#34;push&#34;, &#34;pop&#34; };
char *op11[] = { &#34;stmia&#34;, &#34;ldmia&#34; };

char *cond[] = { &#34;eq&#34;, &#34;ne&#34;, &#34;hs&#34;, &#34;lo&#34;, &#34;mi&#34;, &#34;pl&#34;, &#34;vs&#34;, &#34;vc&#34;,
			 &#34;hi&#34;, &#34;ls&#34;, &#34;ge&#34;, &#34;lt&#34;, &#34;gt&#34;, &#34;le&#34;, &#34;al&#34;, &#34;nv&#34; };

#define B(h, l)		bits(i, h, l)
#define IMM(h, l)	B(h, l)
#define REG(h, l)	reg(B(h, l))
#define LHREG(h, l, lh)	lhreg(B(h, l), B(lh, lh))
#define COND(h, l)	cond[B(h, l)]
#define OP1(h, l)	op1[B(h, l)]
#define OP2(h, l)	op2[B(h, l)]
#define OP3(h, l)	op3[B(h, l)]
#define OP4(h, l)	op4[B(h, l)]
#define OP5(h, l)	op5[B(h, l)]
#define OP6(h, l)	op6[B(h, l)]
#define OP7(h, l)	op7[B(h, l)]
#define OP8(h, l)	op8[B(h, l)]
#define OP9(h, l)	op9[B(h, l)]
#define OP10(h, l)	op10[B(h, l)]
#define OP11(h, l)	op11[B(h, l)]
#define SBZ(h, l)	if(IMM(h, l) != 0) diag(&#34;%x: %x bits %d,%d not zero&#34;, pc, i, h, l)
#define SNBZ(h, l)	if(IMM(h, l) == 0) diag(&#34;%x: %x bits %d,%d zero&#34;, pc, i, h, l)
#define SBO(h, l)	if(IMM(h, l) != 1) diag(&#34;%x: %x bits %d,%d not one&#34;, pc, i, h, l)

static int
bits(int i, int h, int l)
{
	if(h &lt; l)
		diag(&#34;h &lt; l in bits&#34;);
	return (i&amp;(((1&lt;&lt;(h-l+1))-1)&lt;&lt;l))&gt;&gt;l;
}

static char *
reg(int r)
{
	static char s[4][4];
	static int i = 0;

	if(r &lt; 0 || r &gt; 7)
		diag(&#34;register %d out of range&#34;, r);
	i++;
	if(i == 4)
		i = 0;
	sprint(s[i], &#34;r%d&#34;, r);
	return s[i];
}

static char *regnames[] = { &#34;sp&#34;, &#34;lr&#34;, &#34;pc&#34; };

static char *
lhreg(int r, int lh)
{
	static char s[4][4];
	static int i = 0;

	if(lh == 0)
		return reg(r);
	if(r &lt; 0 || r &gt; 7)
		diag(&#34;high register %d out of range&#34;, r);
	i++;
	if(i == 4)
		i = 0;
	if(r &gt;= 5)
		sprint(s[i], &#34;%s&#34;, regnames[r-5]);
	else
		sprint(s[i], &#34;r%d&#34;, r+8);
	return s[i];
}

static void
illegal(int i, int pc)
{
	diag(&#34;%x: %x illegal instruction&#34;, pc, i);
}

static void
dis(int i, int pc)
{
	static int lasto;
	int o, l;
	char *op;

	print(&#34;%x: %x:	&#34;, pc, i);
	if(i&amp;0xffff0000)
		illegal(i, pc);
	o = B(15, 13);
	switch(o){
	case 0:
		o = B(12, 11);
		switch(o){
			case 0:
			case 1:
			case 2:
				print(&#34;%s	%d, %s, %s\n&#34;, OP1(12, 11), IMM(10, 6), REG(5, 3), REG(2, 0));
				return;
			case 3:
				if(B(10, 10) == 0)
					print(&#34;%s	%s, %s, %s\n&#34;, OP2(9, 9), REG(8, 6), REG(5, 3), REG(2, 0));
				else
					print(&#34;%s	%d, %s, %s\n&#34;, OP2(9, 9), IMM(8, 6), REG(5, 3), REG(2, 0));
				return;
		}
	case 1:
		print(&#34;%s	%d, %s\n&#34;, OP3(12, 11), IMM(7, 0), REG(10, 8));
		return;
	case 2:
		o = B(12, 10);
		if(o == 0){
			print(&#34;%s	%s, %s\n&#34;, OP4(9, 6), REG(5, 3), REG(2, 0));
			return;
		}
		if(o == 1){
			o = B(9, 8);
			if(o == 3){
				SBZ(7, 7);
				SBZ(2, 0);
				print(&#34;%s	%s\n&#34;, OP5(9, 8), LHREG(5, 3, 6));
				return;
			}
			SNBZ(7, 6);
			print(&#34;%s	%s, %s\n&#34;, OP5(9, 8), LHREG(5, 3, 6), LHREG(2, 0, 7));
			return;
		}
		if(o == 2 || o == 3){
			print(&#34;movw	%d(pc)[%x], %s\n&#34;, 4*IMM(7, 0), 4*IMM(7, 0)+pc+4, REG(10, 8));
			return;
		}
		op = OP6(11, 9);
		if(*op == &#39;l&#39;)
			print(&#34;%s	[%s, %s], %s\n&#34;, op+1, REG(8, 6), REG(5, 3), REG(2, 0));
		else
			print(&#34;%s	%s, [%s, %s]\n&#34;, op+1, REG(2, 0), REG(8, 6), REG(5, 3));
		return;
	case 3:
		op = OP7(12, 11);
		if(B(12, 11) == 0 || B(12,11) == 1)
			l = 4;
		else
			l = 1;
		if(*op == &#39;l&#39;)
			print(&#34;%s	%d(%s), %s\n&#34;, op+1, l*IMM(10, 6), REG(5, 3), REG(2, 0));
		else
			print(&#34;%s	%s, %d(%s)\n&#34;, op+1, REG(2, 0), l*IMM(10, 6), REG(5, 3));
		return;
	case 4:
		if(B(12, 12) == 0){
			op = OP8(11, 11);
			if(*op == &#39;l&#39;)
				print(&#34;%s	%d(%s), %s\n&#34;, op+1, 2*IMM(10, 6), REG(5, 3), REG(2, 0));
			else
				print(&#34;%s	%s, %d(%s)\n&#34;, op+1, REG(2, 0), 2*IMM(10, 6), REG(5, 3));
			return;
		}
		op = OP9(11, 11);
		if(*op == &#39;l&#39;)
			print(&#34;%s	%d(sp), %s\n&#34;, op+1, 4*IMM(7, 0), REG(10, 8));
		else
			print(&#34;%s	%s, %d(sp)\n&#34;, op+1, REG(10, 8), 4*IMM(7, 0));
		return;
	case 5:
		if(B(12, 12) == 0){
			if(B(11, 11) == 0)
				print(&#34;add	%d, pc, %s\n&#34;, 4*IMM(7, 0), REG(10, 8));
			else
				print(&#34;add	%d, sp, %s\n&#34;, 4*IMM(7, 0), REG(10, 8));
			return;
		}
		if(B(11, 8) == 0){
			print(&#34;%s	%d, sp\n&#34;, OP2(7, 7), 4*IMM(6, 0));
			return;
		}
		SBO(10, 10);
		SBZ(9, 9);
		if(B(8, 8) == 0)
			print(&#34;%s	sp, %d\n&#34;, OP10(11, 11), IMM(7, 0));
		else
			print(&#34;%s	sp, %d|15\n&#34;, OP10(11, 11), IMM(7, 0));
		return;
	case 6:
		if(B(12, 12) == 0){
			print(&#34;%s	%s, %d\n&#34;, OP11(11, 11), REG(10, 8), IMM(7, 0));
			return;
		}
		if(B(11, 8) == 0xf){
			print(&#34;swi	%d\n&#34;, IMM(7, 0));
			return;
		}
		o = IMM(7, 0);
		if(o&amp;0x80)
			o |= 0xffffff00;
		o = pc+4+(o&lt;&lt;1);
		print(&#34;b%s	%x\n&#34;, COND(11, 8), o);
		return;
	case 7:
		o = B(12, 11);
		switch(o){
			case 0:
				o = IMM(10, 0);
				if(o&amp;0x400)
					o |= 0xfffff800;
				o = pc+4+(o&lt;&lt;1);
				print(&#34;b	%x\n&#34;, o);
				return;
			case 1:
				illegal(i, pc);
				return;
			case 2:
				lasto = IMM(10, 0);
				print(&#34;bl\n&#34;);
				return;
			case 3:
				if(lasto&amp;0x400)
					lasto |= 0xfffff800;
				o = IMM(10, 0);
				o = (pc-2)+4+(o&lt;&lt;1)+(lasto&lt;&lt;12);
				print(&#34;bl %x\n&#34;, o);
				return;
		}
	}
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
