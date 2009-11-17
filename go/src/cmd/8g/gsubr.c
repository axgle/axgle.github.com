<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8g/gsubr.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/8g/gsubr.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Derived from Inferno utils/8c/txt.c
// http://code.google.com/p/inferno-os/source/browse/utils/8c/txt.c
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

#include &#34;gg.h&#34;

// TODO(rsc): Can make this bigger if we move
// the text segment up higher in 8l for all GOOS.
uint32 unmappedzero = 4096;

#define	CASE(a,b)	(((a)&lt;&lt;16)|((b)&lt;&lt;0))

void
clearp(Prog *p)
{
	p-&gt;as = AEND;
	p-&gt;from.type = D_NONE;
	p-&gt;from.index = D_NONE;
	p-&gt;to.type = D_NONE;
	p-&gt;to.index = D_NONE;
	p-&gt;loc = pcloc;
	pcloc++;
}

/*
 * generate and return proc with p-&gt;as = as,
 * linked into program.  pc is next instruction.
 */
Prog*
prog(int as)
{
	Prog *p;

	p = pc;
	pc = mal(sizeof(*pc));

	clearp(pc);

	if(lineno == 0) {
		if(debug[&#39;K&#39;])
			warn(&#34;prog: line 0&#34;);
	}

	p-&gt;as = as;
	p-&gt;lineno = lineno;
	p-&gt;link = pc;
	return p;
}

/*
 * generate a branch.
 * t is ignored.
 */
Prog*
gbranch(int as, Type *t)
{
	Prog *p;

	p = prog(as);
	p-&gt;to.type = D_BRANCH;
	p-&gt;to.branch = P;
	return p;
}

/*
 * patch previous branch to jump to to.
 */
void
patch(Prog *p, Prog *to)
{
	if(p-&gt;to.type != D_BRANCH)
		fatal(&#34;patch: not a branch&#34;);
	p-&gt;to.branch = to;
	p-&gt;to.offset = to-&gt;loc;
}

/*
 * start a new Prog list.
 */
Plist*
newplist(void)
{
	Plist *pl;

	pl = mal(sizeof(*pl));
	if(plist == nil)
		plist = pl;
	else
		plast-&gt;link = pl;
	plast = pl;

	pc = mal(sizeof(*pc));
	clearp(pc);
	pl-&gt;firstpc = pc;

	return pl;
}

void
gused(Node *n)
{
	gins(ANOP, n, N);	// used
}

Prog*
gjmp(Prog *to)
{
	Prog *p;

	p = gbranch(AJMP, T);
	if(to != P)
		patch(p, to);
	return p;
}

void
ggloblnod(Node *nam, int32 width)
{
	Prog *p;

	p = gins(AGLOBL, nam, N);
	p-&gt;lineno = nam-&gt;lineno;
	p-&gt;to.sym = S;
	p-&gt;to.type = D_CONST;
	p-&gt;to.offset = width;
}

void
ggloblsym(Sym *s, int32 width, int dupok)
{
	Prog *p;

	p = gins(AGLOBL, N, N);
	p-&gt;from.type = D_EXTERN;
	p-&gt;from.index = D_NONE;
	p-&gt;from.sym = s;
	p-&gt;to.type = D_CONST;
	p-&gt;to.index = D_NONE;
	p-&gt;to.offset = width;
	if(dupok)
		p-&gt;from.scale = DUPOK;
}

int
isfat(Type *t)
{
	if(t != T)
	switch(t-&gt;etype) {
	case TSTRUCT:
	case TARRAY:
	case TSTRING:
	case TINTER:	// maybe remove later
	case TDDD:	// maybe remove later
		return 1;
	}
	return 0;
}

/*
 * naddr of func generates code for address of func.
 * if using opcode that can take address implicitly,
 * call afunclit to fix up the argument.
 */
void
afunclit(Addr *a)
{
	if(a-&gt;type == D_ADDR &amp;&amp; a-&gt;index == D_EXTERN) {
		a-&gt;type = D_EXTERN;
		a-&gt;index = D_NONE;
	}
}

/*
 * return Axxx for Oxxx on type t.
 */
int
optoas(int op, Type *t)
{
	int a;

	if(t == T)
		fatal(&#34;optoas: t is nil&#34;);

	a = AGOK;
	switch(CASE(op, simtype[t-&gt;etype])) {
	default:
		fatal(&#34;optoas: no entry %O-%T&#34;, op, t);
		break;

	case CASE(OADDR, TPTR32):
		a = ALEAL;
		break;

	case CASE(OEQ, TBOOL):
	case CASE(OEQ, TINT8):
	case CASE(OEQ, TUINT8):
	case CASE(OEQ, TINT16):
	case CASE(OEQ, TUINT16):
	case CASE(OEQ, TINT32):
	case CASE(OEQ, TUINT32):
	case CASE(OEQ, TINT64):
	case CASE(OEQ, TUINT64):
	case CASE(OEQ, TPTR32):
	case CASE(OEQ, TPTR64):
	case CASE(OEQ, TFLOAT32):
	case CASE(OEQ, TFLOAT64):
		a = AJEQ;
		break;

	case CASE(ONE, TBOOL):
	case CASE(ONE, TINT8):
	case CASE(ONE, TUINT8):
	case CASE(ONE, TINT16):
	case CASE(ONE, TUINT16):
	case CASE(ONE, TINT32):
	case CASE(ONE, TUINT32):
	case CASE(ONE, TINT64):
	case CASE(ONE, TUINT64):
	case CASE(ONE, TPTR32):
	case CASE(ONE, TPTR64):
	case CASE(ONE, TFLOAT32):
	case CASE(ONE, TFLOAT64):
		a = AJNE;
		break;

	case CASE(OLT, TINT8):
	case CASE(OLT, TINT16):
	case CASE(OLT, TINT32):
	case CASE(OLT, TINT64):
		a = AJLT;
		break;

	case CASE(OLT, TUINT8):
	case CASE(OLT, TUINT16):
	case CASE(OLT, TUINT32):
	case CASE(OLT, TUINT64):
	case CASE(OGT, TFLOAT32):
	case CASE(OGT, TFLOAT64):
		a = AJCS;
		break;

	case CASE(OLE, TINT8):
	case CASE(OLE, TINT16):
	case CASE(OLE, TINT32):
	case CASE(OLE, TINT64):
		a = AJLE;
		break;

	case CASE(OLE, TUINT8):
	case CASE(OLE, TUINT16):
	case CASE(OLE, TUINT32):
	case CASE(OLE, TUINT64):
	case CASE(OGE, TFLOAT32):
	case CASE(OGE, TFLOAT64):
		a = AJLS;
		break;

	case CASE(OGT, TINT8):
	case CASE(OGT, TINT16):
	case CASE(OGT, TINT32):
	case CASE(OGT, TINT64):
		a = AJGT;
		break;

	case CASE(OGT, TUINT8):
	case CASE(OGT, TUINT16):
	case CASE(OGT, TUINT32):
	case CASE(OGT, TUINT64):
	case CASE(OLT, TFLOAT32):
	case CASE(OLT, TFLOAT64):
		a = AJHI;
		break;

	case CASE(OGE, TINT8):
	case CASE(OGE, TINT16):
	case CASE(OGE, TINT32):
	case CASE(OGE, TINT64):
		a = AJGE;
		break;

	case CASE(OGE, TUINT8):
	case CASE(OGE, TUINT16):
	case CASE(OGE, TUINT32):
	case CASE(OGE, TUINT64):
	case CASE(OLE, TFLOAT32):
	case CASE(OLE, TFLOAT64):
		a = AJCC;
		break;

	case CASE(OCMP, TBOOL):
	case CASE(OCMP, TINT8):
	case CASE(OCMP, TUINT8):
		a = ACMPB;
		break;

	case CASE(OCMP, TINT16):
	case CASE(OCMP, TUINT16):
		a = ACMPW;
		break;

	case CASE(OCMP, TINT32):
	case CASE(OCMP, TUINT32):
	case CASE(OCMP, TPTR32):
		a = ACMPL;
		break;

	case CASE(OAS, TBOOL):
	case CASE(OAS, TINT8):
	case CASE(OAS, TUINT8):
		a = AMOVB;
		break;

	case CASE(OAS, TINT16):
	case CASE(OAS, TUINT16):
		a = AMOVW;
		break;

	case CASE(OAS, TINT32):
	case CASE(OAS, TUINT32):
	case CASE(OAS, TPTR32):
		a = AMOVL;
		break;

	case CASE(OADD, TINT8):
	case CASE(OADD, TUINT8):
		a = AADDB;
		break;

	case CASE(OADD, TINT16):
	case CASE(OADD, TUINT16):
		a = AADDW;
		break;

	case CASE(OADD, TINT32):
	case CASE(OADD, TUINT32):
	case CASE(OADD, TPTR32):
		a = AADDL;
		break;

	case CASE(OSUB, TINT8):
	case CASE(OSUB, TUINT8):
		a = ASUBB;
		break;

	case CASE(OSUB, TINT16):
	case CASE(OSUB, TUINT16):
		a = ASUBW;
		break;

	case CASE(OSUB, TINT32):
	case CASE(OSUB, TUINT32):
	case CASE(OSUB, TPTR32):
		a = ASUBL;
		break;

	case CASE(OINC, TINT8):
	case CASE(OINC, TUINT8):
		a = AINCB;
		break;

	case CASE(OINC, TINT16):
	case CASE(OINC, TUINT16):
		a = AINCW;
		break;

	case CASE(OINC, TINT32):
	case CASE(OINC, TUINT32):
	case CASE(OINC, TPTR32):
		a = AINCL;
		break;

	case CASE(ODEC, TINT8):
	case CASE(ODEC, TUINT8):
		a = ADECB;
		break;

	case CASE(ODEC, TINT16):
	case CASE(ODEC, TUINT16):
		a = ADECW;
		break;

	case CASE(ODEC, TINT32):
	case CASE(ODEC, TUINT32):
	case CASE(ODEC, TPTR32):
		a = ADECL;
		break;

	case CASE(OCOM, TINT8):
	case CASE(OCOM, TUINT8):
		a = ANOTB;
		break;

	case CASE(OCOM, TINT16):
	case CASE(OCOM, TUINT16):
		a = ANOTW;
		break;

	case CASE(OCOM, TINT32):
	case CASE(OCOM, TUINT32):
	case CASE(OCOM, TPTR32):
		a = ANOTL;
		break;

	case CASE(OMINUS, TINT8):
	case CASE(OMINUS, TUINT8):
		a = ANEGB;
		break;

	case CASE(OMINUS, TINT16):
	case CASE(OMINUS, TUINT16):
		a = ANEGW;
		break;

	case CASE(OMINUS, TINT32):
	case CASE(OMINUS, TUINT32):
	case CASE(OMINUS, TPTR32):
		a = ANEGL;
		break;

	case CASE(OAND, TINT8):
	case CASE(OAND, TUINT8):
		a = AANDB;
		break;

	case CASE(OAND, TINT16):
	case CASE(OAND, TUINT16):
		a = AANDW;
		break;

	case CASE(OAND, TINT32):
	case CASE(OAND, TUINT32):
	case CASE(OAND, TPTR32):
		a = AANDL;
		break;

	case CASE(OOR, TINT8):
	case CASE(OOR, TUINT8):
		a = AORB;
		break;

	case CASE(OOR, TINT16):
	case CASE(OOR, TUINT16):
		a = AORW;
		break;

	case CASE(OOR, TINT32):
	case CASE(OOR, TUINT32):
	case CASE(OOR, TPTR32):
		a = AORL;
		break;

	case CASE(OXOR, TINT8):
	case CASE(OXOR, TUINT8):
		a = AXORB;
		break;

	case CASE(OXOR, TINT16):
	case CASE(OXOR, TUINT16):
		a = AXORW;
		break;

	case CASE(OXOR, TINT32):
	case CASE(OXOR, TUINT32):
	case CASE(OXOR, TPTR32):
		a = AXORL;
		break;

	case CASE(OLSH, TINT8):
	case CASE(OLSH, TUINT8):
		a = ASHLB;
		break;

	case CASE(OLSH, TINT16):
	case CASE(OLSH, TUINT16):
		a = ASHLW;
		break;

	case CASE(OLSH, TINT32):
	case CASE(OLSH, TUINT32):
	case CASE(OLSH, TPTR32):
		a = ASHLL;
		break;

	case CASE(ORSH, TUINT8):
		a = ASHRB;
		break;

	case CASE(ORSH, TUINT16):
		a = ASHRW;
		break;

	case CASE(ORSH, TUINT32):
	case CASE(ORSH, TPTR32):
		a = ASHRL;
		break;

	case CASE(ORSH, TINT8):
		a = ASARB;
		break;

	case CASE(ORSH, TINT16):
		a = ASARW;
		break;

	case CASE(ORSH, TINT32):
		a = ASARL;
		break;

	case CASE(OMUL, TINT8):
	case CASE(OMUL, TUINT8):
		a = AIMULB;
		break;

	case CASE(OMUL, TINT16):
	case CASE(OMUL, TUINT16):
		a = AIMULW;
		break;

	case CASE(OMUL, TINT32):
	case CASE(OMUL, TUINT32):
	case CASE(OMUL, TPTR32):
		a = AIMULL;
		break;

	case CASE(ODIV, TINT8):
	case CASE(OMOD, TINT8):
		a = AIDIVB;
		break;

	case CASE(ODIV, TUINT8):
	case CASE(OMOD, TUINT8):
		a = ADIVB;
		break;

	case CASE(ODIV, TINT16):
	case CASE(OMOD, TINT16):
		a = AIDIVW;
		break;

	case CASE(ODIV, TUINT16):
	case CASE(OMOD, TUINT16):
		a = ADIVW;
		break;

	case CASE(ODIV, TINT32):
	case CASE(OMOD, TINT32):
		a = AIDIVL;
		break;

	case CASE(ODIV, TUINT32):
	case CASE(ODIV, TPTR32):
	case CASE(OMOD, TUINT32):
	case CASE(OMOD, TPTR32):
		a = ADIVL;
		break;

	case CASE(OEXTEND, TINT16):
		a = ACWD;
		break;

	case CASE(OEXTEND, TINT32):
		a = ACDQ;
		break;
	}
	return a;
}

#define FCASE(a, b, c)  (((a)&lt;&lt;16)|((b)&lt;&lt;8)|(c))
int
foptoas(int op, Type *t, int flg)
{
	int et;

	et = simtype[t-&gt;etype];

	// If we need Fpop, it means we&#39;re working on
	// two different floating-point registers, not memory.
	// There the instruction only has a float64 form.
	if(flg &amp; Fpop)
		et = TFLOAT64;

	// clear Frev if unneeded
	switch(op) {
	case OADD:
	case OMUL:
		flg &amp;= ~Frev;
		break;
	}

	switch(FCASE(op, et, flg)) {
	case FCASE(OADD, TFLOAT32, 0):
		return AFADDF;
	case FCASE(OADD, TFLOAT64, 0):
		return AFADDD;
	case FCASE(OADD, TFLOAT64, Fpop):
		return AFADDDP;

	case FCASE(OSUB, TFLOAT32, 0):
		return AFSUBF;
	case FCASE(OSUB, TFLOAT32, Frev):
		return AFSUBRF;

	case FCASE(OSUB, TFLOAT64, 0):
		return AFSUBD;
	case FCASE(OSUB, TFLOAT64, Frev):
		return AFSUBRD;
	case FCASE(OSUB, TFLOAT64, Fpop):
		return AFSUBDP;
	case FCASE(OSUB, TFLOAT64, Fpop|Frev):
		return AFSUBRDP;

	case FCASE(OMUL, TFLOAT32, 0):
		return AFMULF;
	case FCASE(OMUL, TFLOAT64, 0):
		return AFMULD;
	case FCASE(OMUL, TFLOAT64, Fpop):
		return AFMULDP;

	case FCASE(ODIV, TFLOAT32, 0):
		return AFDIVF;
	case FCASE(ODIV, TFLOAT32, Frev):
		return AFDIVRF;

	case FCASE(ODIV, TFLOAT64, 0):
		return AFDIVD;
	case FCASE(ODIV, TFLOAT64, Frev):
		return AFDIVRD;
	case FCASE(ODIV, TFLOAT64, Fpop):
		return AFDIVDP;
	case FCASE(ODIV, TFLOAT64, Fpop|Frev):
		return AFDIVRDP;

	case FCASE(OCMP, TFLOAT32, 0):
		return AFCOMF;
	case FCASE(OCMP, TFLOAT32, Fpop):
		return AFCOMFP;
	case FCASE(OCMP, TFLOAT64, 0):
		return AFCOMD;
	case FCASE(OCMP, TFLOAT64, Fpop):
		return AFCOMDP;
	case FCASE(OCMP, TFLOAT64, Fpop2):
		return AFCOMDPP;
	}

	fatal(&#34;foptoas %O %T %#x&#34;, op, t, flg);
	return 0;
}

static	int	resvd[] =
{
//	D_DI,	// for movstring
//	D_SI,	// for movstring

	D_AX,	// for divide
	D_CX,	// for shift
	D_DX,	// for divide
	D_SP,	// for stack

	D_BL,	// because D_BX can be allocated
	D_BH,
};

void
ginit(void)
{
	int i;

	for(i=0; i&lt;nelem(reg); i++)
		reg[i] = 1;
	for(i=D_AL; i&lt;=D_DI; i++)
		reg[i] = 0;

	for(i=0; i&lt;nelem(resvd); i++)
		reg[resvd[i]]++;
}

ulong regpc[D_NONE];

void
gclean(void)
{
	int i;

	for(i=0; i&lt;nelem(resvd); i++)
		reg[resvd[i]]--;

	for(i=D_AL; i&lt;=D_DI; i++)
		if(reg[i])
			yyerror(&#34;reg %R left allocated at %lux&#34;, i, regpc[i]);
}

/*
 * allocate register of type t, leave in n.
 * if o != N, o is desired fixed register.
 * caller must regfree(n).
 */
void
regalloc(Node *n, Type *t, Node *o)
{
	int i, et;

	if(t == T)
		fatal(&#34;regalloc: t nil&#34;);
	et = simtype[t-&gt;etype];

	switch(et) {
	case TINT8:
	case TUINT8:
	case TINT16:
	case TUINT16:
	case TINT32:
	case TUINT32:
	case TINT64:
	case TUINT64:
	case TPTR32:
	case TPTR64:
	case TBOOL:
		if(o != N &amp;&amp; o-&gt;op == OREGISTER) {
			i = o-&gt;val.u.reg;
			if(i &gt;= D_AX &amp;&amp; i &lt;= D_DI)
				goto out;
		}
		for(i=D_AX; i&lt;=D_DI; i++)
			if(reg[i] == 0)
				goto out;

		fprint(2, &#34;registers allocated at\n&#34;);
		for(i=D_AX; i&lt;=D_DI; i++)
			fprint(2, &#34;\t%R\t%#lux\n&#34;, i, regpc[i]);
		yyerror(&#34;out of fixed registers&#34;);
		goto err;

	case TFLOAT32:
	case TFLOAT64:
		i = D_F0;
		goto out;
	}
	yyerror(&#34;regalloc: unknown type %T&#34;, t);
	i = 0;

err:
	nodreg(n, t, 0);
	return;

out:
	if(reg[i] == 0) {
		regpc[i] = (ulong)__builtin_return_address(0);
		if(i == D_AX || i == D_CX || i == D_DX || i == D_SP) {
			dump(&#34;regalloc-o&#34;, o);
			fatal(&#34;regalloc %R&#34;, i);
		}
	}
	reg[i]++;
	nodreg(n, t, i);
}

void
regfree(Node *n)
{
	int i;

	if(n-&gt;op != OREGISTER &amp;&amp; n-&gt;op != OINDREG)
		fatal(&#34;regfree: not a register&#34;);
	i = n-&gt;val.u.reg;
	if(i &lt; 0 || i &gt;= sizeof(reg))
		fatal(&#34;regfree: reg out of range&#34;);
	if(reg[i] &lt;= 0)
		fatal(&#34;regfree: reg not allocated&#34;);
	reg[i]--;
	if(reg[i] == 0 &amp;&amp; (i == D_AX || i == D_CX || i == D_DX || i == D_SP))
		fatal(&#34;regfree %R&#34;, i);
}

void
tempalloc(Node *n, Type *t)
{
	int w;

	dowidth(t);

	memset(n, 0, sizeof(*n));
	n-&gt;op = ONAME;
	n-&gt;sym = S;
	n-&gt;type = t;
	n-&gt;etype = t-&gt;etype;
	n-&gt;class = PAUTO;
	n-&gt;addable = 1;
	n-&gt;ullman = 1;
	n-&gt;noescape = 1;
	n-&gt;ostk = stksize;

	w = t-&gt;width;
	stksize += w;
	stksize = rnd(stksize, w);
	n-&gt;xoffset = -stksize;
//print(&#34;tempalloc %d -&gt; %d from %p\n&#34;, n-&gt;ostk, n-&gt;xoffset, __builtin_return_address(0));
	if(stksize &gt; maxstksize)
		maxstksize = stksize;
}

void
tempfree(Node *n)
{
//print(&#34;tempfree %d\n&#34;, n-&gt;xoffset);
	if(n-&gt;xoffset != -stksize)
		fatal(&#34;tempfree %lld %d&#34;, -n-&gt;xoffset, stksize);
	stksize = n-&gt;ostk;
}

/*
 * initialize n to be register r of type t.
 */
void
nodreg(Node *n, Type *t, int r)
{
	if(t == T)
		fatal(&#34;nodreg: t nil&#34;);

	memset(n, 0, sizeof(*n));
	n-&gt;op = OREGISTER;
	n-&gt;addable = 1;
	ullmancalc(n);
	n-&gt;val.u.reg = r;
	n-&gt;type = t;
}

/*
 * initialize n to be indirect of register r; n is type t.
 */
void
nodindreg(Node *n, Type *t, int r)
{
	nodreg(n, t, r);
	n-&gt;op = OINDREG;
}

Node*
nodarg(Type *t, int fp)
{
	Node *n;
	Type *first;
	Iter savet;

	// entire argument struct, not just one arg
	switch(t-&gt;etype) {
	default:
		fatal(&#34;nodarg %T&#34;, t);

	case TSTRUCT:
		if(!t-&gt;funarg)
			fatal(&#34;nodarg: TSTRUCT but not funarg&#34;);
		n = nod(ONAME, N, N);
		n-&gt;sym = lookup(&#34;.args&#34;);
		n-&gt;type = t;
		first = structfirst(&amp;savet, &amp;t);
		if(first == nil)
			fatal(&#34;nodarg: bad struct&#34;);
		if(first-&gt;width == BADWIDTH)
			fatal(&#34;nodarg: offset not computed for %T&#34;, t);
		n-&gt;xoffset = first-&gt;width;
		n-&gt;addable = 1;
		break;

	case TFIELD:
		n = nod(ONAME, N, N);
		n-&gt;type = t-&gt;type;
		n-&gt;sym = t-&gt;sym;
		if(t-&gt;width == BADWIDTH)
			fatal(&#34;nodarg: offset not computed for %T&#34;, t);
		n-&gt;xoffset = t-&gt;width;
		n-&gt;addable = 1;
		break;
	}

	switch(fp) {
	default:
		fatal(&#34;nodarg %T %d&#34;, t, fp);

	case 0:		// output arg
		n-&gt;op = OINDREG;
		n-&gt;val.u.reg = D_SP;
		break;

	case 1:		// input arg
		n-&gt;class = PPARAM;
		break;
	}

	n-&gt;typecheck = 1;
	return n;
}

/*
 * generate
 *	as $c, reg
 */
void
gconreg(int as, vlong c, int reg)
{
	Node n1, n2;

	nodconst(&amp;n1, types[TINT64], c);
	nodreg(&amp;n2, types[TINT64], reg);
	gins(as, &amp;n1, &amp;n2);
}

/*
 * swap node contents
 */
void
nswap(Node *a, Node *b)
{
	Node t;

	t = *a;
	*a = *b;
	*b = t;
}

/*
 * return constant i node.
 * overwritten by next call, but useful in calls to gins.
 */
Node*
ncon(uint32 i)
{
	static Node n;

	if(n.type == T)
		nodconst(&amp;n, types[TUINT32], 0);
	mpmovecfix(n.val.u.xval, i);
	return &amp;n;
}

/*
 * Is this node a memory operand?
 */
int
ismem(Node *n)
{
	switch(n-&gt;op) {
	case OLEN:
	case OCAP:
	case OINDREG:
	case ONAME:
	case OPARAM:
		return 1;
	}
	return 0;
}

Node sclean[10];
int nsclean;

/*
 * n is a 64-bit value.  fill in lo and hi to refer to its 32-bit halves.
 */
void
split64(Node *n, Node *lo, Node *hi)
{
	Node n1;
	int64 i;

	if(!is64(n-&gt;type))
		fatal(&#34;split64 %T&#34;, n-&gt;type);

	sclean[nsclean].op = OEMPTY;
	if(nsclean &gt;= nelem(sclean))
		fatal(&#34;split64 clean&#34;);
	nsclean++;
	switch(n-&gt;op) {
	default:
		if(!dotaddable(n, &amp;n1)) {
			igen(n, &amp;n1, N);
			sclean[nsclean-1] = n1;
		}
		n = &amp;n1;
		goto common;
	case ONAME:
		if(n-&gt;class == PPARAMREF) {
			cgen(n-&gt;heapaddr, &amp;n1);
			sclean[nsclean-1] = n1;
			// fall through.
			n = &amp;n1;
		}
		goto common;
	case OINDREG:
	common:
		*lo = *n;
		*hi = *n;
		lo-&gt;type = types[TUINT32];
		if(n-&gt;type-&gt;etype == TINT64)
			hi-&gt;type = types[TINT32];
		else
			hi-&gt;type = types[TUINT32];
		hi-&gt;xoffset += 4;
		break;

	case OLITERAL:
		convconst(&amp;n1, n-&gt;type, &amp;n-&gt;val);
		i = mpgetfix(n1.val.u.xval);
		nodconst(lo, types[TUINT32], (uint32)i);
		i &gt;&gt;= 32;
		if(n-&gt;type-&gt;etype == TINT64)
			nodconst(hi, types[TINT32], (int32)i);
		else
			nodconst(hi, types[TUINT32], (uint32)i);
		break;
	}
}

void
splitclean(void)
{
	if(nsclean &lt;= 0)
		fatal(&#34;splitclean&#34;);
	nsclean--;
	if(sclean[nsclean].op != OEMPTY)
		regfree(&amp;sclean[nsclean]);
}

/*
 * set up nodes representing fp constants
 */
Node zerof;
Node two64f;
Node two63f;

void
bignodes(void)
{
	static int did;

	if(did)
		return;
	did = 1;

	two64f = *ncon(0);
	two64f.type = types[TFLOAT64];
	two64f.val.ctype = CTFLT;
	two64f.val.u.fval = mal(sizeof *two64f.val.u.fval);
	mpmovecflt(two64f.val.u.fval, 18446744073709551616.);

	two63f = two64f;
	two63f.val.u.fval = mal(sizeof *two63f.val.u.fval);
	mpmovecflt(two63f.val.u.fval, 9223372036854775808.);

	zerof = two64f;
	zerof.val.u.fval = mal(sizeof *zerof.val.u.fval);
	mpmovecflt(zerof.val.u.fval, 0);
}

void
gmove(Node *f, Node *t)
{
	int a, ft, tt;
	Type *cvt;
	Node r1, r2, t1, t2, flo, fhi, tlo, thi, con, f0, f1, ax, dx, cx;
	Prog *p1, *p2, *p3;

	if(debug[&#39;M&#39;])
		print(&#34;gmove %N -&gt; %N\n&#34;, f, t);

	ft = simsimtype(f-&gt;type);
	tt = simsimtype(t-&gt;type);
	cvt = t-&gt;type;

	// cannot have two integer memory operands;
	// except 64-bit, which always copies via registers anyway.
	if(isint[ft] &amp;&amp; isint[tt] &amp;&amp; !is64(f-&gt;type) &amp;&amp; !is64(t-&gt;type) &amp;&amp; ismem(f) &amp;&amp; ismem(t))
		goto hard;

	// convert constant to desired type
	if(f-&gt;op == OLITERAL) {
		if(tt == TFLOAT32)
			convconst(&amp;con, types[TFLOAT64], &amp;f-&gt;val);
		else
			convconst(&amp;con, t-&gt;type, &amp;f-&gt;val);
		f = &amp;con;
		ft = simsimtype(con.type);

		// some constants can&#39;t move directly to memory.
		if(ismem(t)) {
			// float constants come from memory.
			if(isfloat[tt])
				goto hard;
		}
	}

	// value -&gt; value copy, only one memory operand.
	// figure out the instruction to use.
	// break out of switch for one-instruction gins.
	// goto rdst for &#34;destination must be register&#34;.
	// goto hard for &#34;convert to cvt type first&#34;.
	// otherwise handle and return.

	switch(CASE(ft, tt)) {
	default:
		goto fatal;

	/*
	 * integer copy and truncate
	 */
	case CASE(TINT8, TINT8):	// same size
	case CASE(TINT8, TUINT8):
	case CASE(TUINT8, TINT8):
	case CASE(TUINT8, TUINT8):
	case CASE(TINT16, TINT8):	// truncate
	case CASE(TUINT16, TINT8):
	case CASE(TINT32, TINT8):
	case CASE(TUINT32, TINT8):
	case CASE(TINT16, TUINT8):
	case CASE(TUINT16, TUINT8):
	case CASE(TINT32, TUINT8):
	case CASE(TUINT32, TUINT8):
		a = AMOVB;
		break;

	case CASE(TINT64, TINT8):	// truncate low word
	case CASE(TUINT64, TINT8):
	case CASE(TINT64, TUINT8):
	case CASE(TUINT64, TUINT8):
		split64(f, &amp;flo, &amp;fhi);
		nodreg(&amp;r1, t-&gt;type, D_AX);
		gins(AMOVB, &amp;flo, &amp;r1);
		gins(AMOVB, &amp;r1, t);
		splitclean();
		return;

	case CASE(TINT16, TINT16):	// same size
	case CASE(TINT16, TUINT16):
	case CASE(TUINT16, TINT16):
	case CASE(TUINT16, TUINT16):
	case CASE(TINT32, TINT16):	// truncate
	case CASE(TUINT32, TINT16):
	case CASE(TINT32, TUINT16):
	case CASE(TUINT32, TUINT16):
		a = AMOVW;
		break;

	case CASE(TINT64, TINT16):	// truncate low word
	case CASE(TUINT64, TINT16):
	case CASE(TINT64, TUINT16):
	case CASE(TUINT64, TUINT16):
		split64(f, &amp;flo, &amp;fhi);
		nodreg(&amp;r1, t-&gt;type, D_AX);
		gins(AMOVW, &amp;flo, &amp;r1);
		gins(AMOVW, &amp;r1, t);
		splitclean();
		return;

	case CASE(TINT32, TINT32):	// same size
	case CASE(TINT32, TUINT32):
	case CASE(TUINT32, TINT32):
	case CASE(TUINT32, TUINT32):
		a = AMOVL;
		break;

	case CASE(TINT64, TINT32):	// truncate
	case CASE(TUINT64, TINT32):
	case CASE(TINT64, TUINT32):
	case CASE(TUINT64, TUINT32):
		split64(f, &amp;flo, &amp;fhi);
		nodreg(&amp;r1, t-&gt;type, D_AX);
		gins(AMOVL, &amp;flo, &amp;r1);
		gins(AMOVL, &amp;r1, t);
		splitclean();
		return;

	case CASE(TINT64, TINT64):	// same size
	case CASE(TINT64, TUINT64):
	case CASE(TUINT64, TINT64):
	case CASE(TUINT64, TUINT64):
		split64(f, &amp;flo, &amp;fhi);
		split64(t, &amp;tlo, &amp;thi);
		if(f-&gt;op == OLITERAL) {
			gins(AMOVL, &amp;flo, &amp;tlo);
			gins(AMOVL, &amp;fhi, &amp;thi);
		} else {
			nodreg(&amp;r1, t-&gt;type, D_AX);
			nodreg(&amp;r2, t-&gt;type, D_DX);
			gins(AMOVL, &amp;flo, &amp;r1);
			gins(AMOVL, &amp;fhi, &amp;r2);
			gins(AMOVL, &amp;r1, &amp;tlo);
			gins(AMOVL, &amp;r2, &amp;thi);
		}
		splitclean();
		splitclean();
		return;

	/*
	 * integer up-conversions
	 */
	case CASE(TINT8, TINT16):	// sign extend int8
	case CASE(TINT8, TUINT16):
		a = AMOVBWSX;
		goto rdst;
	case CASE(TINT8, TINT32):
	case CASE(TINT8, TUINT32):
		a = AMOVBLSX;
		goto rdst;
	case CASE(TINT8, TINT64):	// convert via int32
	case CASE(TINT8, TUINT64):
		cvt = types[TINT32];
		goto hard;

	case CASE(TUINT8, TINT16):	// zero extend uint8
	case CASE(TUINT8, TUINT16):
		a = AMOVBWZX;
		goto rdst;
	case CASE(TUINT8, TINT32):
	case CASE(TUINT8, TUINT32):
		a = AMOVBLZX;
		goto rdst;
	case CASE(TUINT8, TINT64):	// convert via uint32
	case CASE(TUINT8, TUINT64):
		cvt = types[TUINT32];
		goto hard;

	case CASE(TINT16, TINT32):	// sign extend int16
	case CASE(TINT16, TUINT32):
		a = AMOVWLSX;
		goto rdst;
	case CASE(TINT16, TINT64):	// convert via int32
	case CASE(TINT16, TUINT64):
		cvt = types[TINT32];
		goto hard;

	case CASE(TUINT16, TINT32):	// zero extend uint16
	case CASE(TUINT16, TUINT32):
		a = AMOVWLZX;
		goto rdst;
	case CASE(TUINT16, TINT64):	// convert via uint32
	case CASE(TUINT16, TUINT64):
		cvt = types[TUINT32];
		goto hard;

	case CASE(TINT32, TINT64):	// sign extend int32
	case CASE(TINT32, TUINT64):
		split64(t, &amp;tlo, &amp;thi);
		nodreg(&amp;flo, tlo.type, D_AX);
		nodreg(&amp;fhi, thi.type, D_DX);
		gmove(f, &amp;flo);
		gins(ACDQ, N, N);
		gins(AMOVL, &amp;flo, &amp;tlo);
		gins(AMOVL, &amp;fhi, &amp;thi);
		splitclean();
		return;

	case CASE(TUINT32, TINT64):	// zero extend uint32
	case CASE(TUINT32, TUINT64):
		split64(t, &amp;tlo, &amp;thi);
		gmove(f, &amp;tlo);
		gins(AMOVL, ncon(0), &amp;thi);
		splitclean();
		return;

	/*
	* float to integer
	*/
	case CASE(TFLOAT32, TINT16):
	case CASE(TFLOAT32, TINT32):
	case CASE(TFLOAT32, TINT64):
	case CASE(TFLOAT64, TINT16):
	case CASE(TFLOAT64, TINT32):
	case CASE(TFLOAT64, TINT64):
		if(t-&gt;op == OREGISTER)
			goto hardmem;
		nodreg(&amp;r1, types[ft], D_F0);
		if(ft == TFLOAT32 &amp;&amp; f-&gt;op != OREGISTER)
			gins(AFMOVF, f, &amp;r1);
		else
			gins(AFMOVD, f, &amp;r1);

		// set round to zero mode during conversion
		tempalloc(&amp;t1, types[TUINT16]);
		tempalloc(&amp;t2, types[TUINT16]);
		gins(AFSTCW, N, &amp;t1);
		gins(AMOVW, ncon(0xf7f), &amp;t2);
		gins(AFLDCW, &amp;t2, N);
		if(tt == TINT16)
			gins(AFMOVWP, &amp;r1, t);
		else if(tt == TINT32)
			gins(AFMOVLP, &amp;r1, t);
		else
			gins(AFMOVVP, &amp;r1, t);
		gins(AFLDCW, &amp;t1, N);
		tempfree(&amp;t2);
		tempfree(&amp;t1);
		return;

	case CASE(TFLOAT32, TINT8):
	case CASE(TFLOAT32, TUINT16):
	case CASE(TFLOAT32, TUINT8):
	case CASE(TFLOAT64, TINT8):
	case CASE(TFLOAT64, TUINT16):
	case CASE(TFLOAT64, TUINT8):
		// convert via int32.
		tempalloc(&amp;t1, types[TINT32]);
		gmove(f, &amp;t1);
		switch(tt) {
		default:
			fatal(&#34;gmove %T&#34;, t);
		case TINT8:
			gins(ACMPL, &amp;t1, ncon(-0x80));
			p1 = gbranch(optoas(OLT, types[TINT32]), T);
			gins(ACMPL, &amp;t1, ncon(0x7f));
			p2 = gbranch(optoas(OGT, types[TINT32]), T);
			p3 = gbranch(AJMP, T);
			patch(p1, pc);
			patch(p2, pc);
			gmove(ncon(-0x80), &amp;t1);
			patch(p3, pc);
			gmove(&amp;t1, t);
			break;
		case TUINT8:
			gins(ATESTL, ncon(0xffffff00), &amp;t1);
			p1 = gbranch(AJEQ, T);
			gins(AMOVB, ncon(0), &amp;t1);
			patch(p1, pc);
			gmove(&amp;t1, t);
			break;
		case TUINT16:
			gins(ATESTL, ncon(0xffff0000), &amp;t1);
			p1 = gbranch(AJEQ, T);
			gins(AMOVW, ncon(0), &amp;t1);
			patch(p1, pc);
			gmove(&amp;t1, t);
			break;
		}
		tempfree(&amp;t1);
		return;

	case CASE(TFLOAT32, TUINT32):
	case CASE(TFLOAT64, TUINT32):
		// convert via int64.
		tempalloc(&amp;t1, types[TINT64]);
		gmove(f, &amp;t1);
		split64(&amp;t1, &amp;tlo, &amp;thi);
		gins(ACMPL, &amp;thi, ncon(0));
		p1 = gbranch(AJEQ, T);
		gins(AMOVL, ncon(0), &amp;tlo);
		patch(p1, pc);
		gmove(&amp;tlo, t);
		splitclean();
		tempfree(&amp;t1);
		return;

	case CASE(TFLOAT32, TUINT64):
	case CASE(TFLOAT64, TUINT64):
		bignodes();
		nodreg(&amp;f0, types[ft], D_F0);
		nodreg(&amp;f1, types[ft], D_F0 + 1);
		nodreg(&amp;ax, types[TUINT16], D_AX);

		gmove(f, &amp;f0);

		// if 0 &gt; v { answer = 0 }
		gmove(&amp;zerof, &amp;f0);
		gins(AFUCOMP, &amp;f0, &amp;f1);
		gins(AFSTSW, N, &amp;ax);
		gins(ASAHF, N, N);
		p1 = gbranch(optoas(OGT, types[tt]), T);
		// if 1&lt;&lt;64 &lt;= v { answer = 0 too }
		gmove(&amp;two64f, &amp;f0);
		gins(AFUCOMP, &amp;f0, &amp;f1);
		gins(AFSTSW, N, &amp;ax);
		gins(ASAHF, N, N);
		p2 = gbranch(optoas(OGT, types[tt]), T);
		patch(p1, pc);
		gins(AFMOVVP, &amp;f0, t);	// don&#39;t care about t, but will pop the stack
		split64(t, &amp;tlo, &amp;thi);
		gins(AMOVL, ncon(0), &amp;tlo);
		gins(AMOVL, ncon(0), &amp;thi);
		splitclean();
		p1 = gbranch(AJMP, T);
		patch(p2, pc);

		// in range; algorithm is:
		//	if small enough, use native float64 -&gt; int64 conversion.
		//	otherwise, subtract 2^63, convert, and add it back.

		// set round to zero mode during conversion
		tempalloc(&amp;t1, types[TUINT16]);
		tempalloc(&amp;t2, types[TUINT16]);
		gins(AFSTCW, N, &amp;t1);
		gins(AMOVW, ncon(0xf7f), &amp;t2);
		gins(AFLDCW, &amp;t2, N);
		tempfree(&amp;t2);

		// actual work
		gmove(&amp;two63f, &amp;f0);
		gins(AFUCOMP, &amp;f0, &amp;f1);
		gins(AFSTSW, N, &amp;ax);
		gins(ASAHF, N, N);
		p2 = gbranch(optoas(OLE, types[tt]), T);
		gins(AFMOVVP, &amp;f0, t);
		p3 = gbranch(AJMP, T);
		patch(p2, pc);
		gmove(&amp;two63f, &amp;f0);
		gins(AFSUBDP, &amp;f0, &amp;f1);
		gins(AFMOVVP, &amp;f0, t);
		split64(t, &amp;tlo, &amp;thi);
		gins(AXORL, ncon(0x80000000), &amp;thi);	// + 2^63
		patch(p3, pc);
		patch(p1, pc);
		splitclean();

		// restore rounding mode
		gins(AFLDCW, &amp;t1, N);
		tempfree(&amp;t1);
		return;

	/*
	 * integer to float
	 */
	case CASE(TINT16, TFLOAT32):
	case CASE(TINT16, TFLOAT64):
	case CASE(TINT32, TFLOAT32):
	case CASE(TINT32, TFLOAT64):
	case CASE(TINT64, TFLOAT32):
	case CASE(TINT64, TFLOAT64):
		if(t-&gt;op != OREGISTER)
			goto hard;
		if(f-&gt;op == OREGISTER) {
			cvt = f-&gt;type;
			goto hardmem;
		}
		switch(ft) {
		case TINT16:
			a = AFMOVW;
			break;
		case TINT32:
			a = AFMOVL;
			break;
		default:
			a = AFMOVV;
			break;
		}
		break;

	case CASE(TINT8, TFLOAT32):
	case CASE(TINT8, TFLOAT64):
	case CASE(TUINT16, TFLOAT32):
	case CASE(TUINT16, TFLOAT64):
	case CASE(TUINT8, TFLOAT32):
	case CASE(TUINT8, TFLOAT64):
		// convert via int32 memory
		cvt = types[TINT32];
		goto hardmem;

	case CASE(TUINT32, TFLOAT32):
	case CASE(TUINT32, TFLOAT64):
		// convert via int64 memory
		cvt = types[TINT64];
		goto hardmem;

	case CASE(TUINT64, TFLOAT32):
	case CASE(TUINT64, TFLOAT64):
		// algorithm is:
		//	if small enough, use native int64 -&gt; uint64 conversion.
		//	otherwise, halve (rounding to odd?), convert, and double.
		nodreg(&amp;ax, types[TUINT32], D_AX);
		nodreg(&amp;dx, types[TUINT32], D_DX);
		nodreg(&amp;cx, types[TUINT32], D_CX);
		tempalloc(&amp;t1, f-&gt;type);
		split64(&amp;t1, &amp;tlo, &amp;thi);
		gmove(f, &amp;t1);
		gins(ACMPL, &amp;thi, ncon(0));
		p1 = gbranch(AJLT, T);
		// native
		t1.type = types[TINT64];
		gmove(&amp;t1, t);
		p2 = gbranch(AJMP, T);
		// simulated
		patch(p1, pc);
		gmove(&amp;tlo, &amp;ax);
		gmove(&amp;thi, &amp;dx);
		p1 = gins(ASHRL, ncon(1), &amp;ax);
		p1-&gt;from.index = D_DX;	// double-width shift DX -&gt; AX
		p1-&gt;from.scale = 0;
		gins(ASETCC, N, &amp;cx);
		gins(AORB, &amp;cx, &amp;ax);
		gins(ASHRL, ncon(1), &amp;dx);
		gmove(&amp;dx, &amp;thi);
		gmove(&amp;ax, &amp;tlo);
		nodreg(&amp;r1, types[tt], D_F0);
		nodreg(&amp;r2, types[tt], D_F0 + 1);
		gmove(&amp;t1, &amp;r1);	// t1.type is TINT64 now, set above
		gins(AFMOVD, &amp;r1, &amp;r1);
		gins(AFADDDP, &amp;r1, &amp;r2);
		gmove(&amp;r1, t);
		patch(p2, pc);
		splitclean();
		tempfree(&amp;t1);
		return;

	/*
	 * float to float
	 */
	case CASE(TFLOAT32, TFLOAT32):
	case CASE(TFLOAT64, TFLOAT64):
		// The way the code generator uses floating-point
		// registers, a move from F0 to F0 is intended as a no-op.
		// On the x86, it&#39;s not: it pushes a second copy of F0
		// on the floating point stack.  So toss it away here.
		// Also, F0 is the *only* register we ever evaluate
		// into, so we should only see register/register as F0/F0.
		if(f-&gt;op == OREGISTER &amp;&amp; t-&gt;op == OREGISTER) {
			if(f-&gt;val.u.reg != D_F0 || t-&gt;val.u.reg != D_F0)
				goto fatal;
			return;
		}
		if(ismem(f) &amp;&amp; ismem(t))
			goto hard;
		a = AFMOVF;
		if(ft == TFLOAT64)
			a = AFMOVD;
		if(ismem(t)) {
			if(f-&gt;op != OREGISTER || f-&gt;val.u.reg != D_F0)
				fatal(&#34;gmove %N&#34;, f);
			a = AFMOVFP;
			if(ft == TFLOAT64)
				a = AFMOVDP;
		}
		break;

	case CASE(TFLOAT32, TFLOAT64):
		if(f-&gt;op == OREGISTER &amp;&amp; t-&gt;op == OREGISTER) {
			if(f-&gt;val.u.reg != D_F0 || t-&gt;val.u.reg != D_F0)
				goto fatal;
			return;
		}
		if(f-&gt;op == OREGISTER)
			gins(AFMOVDP, f, t);
		else
			gins(AFMOVF, f, t);
		return;

	case CASE(TFLOAT64, TFLOAT32):
		if(f-&gt;op == OREGISTER &amp;&amp; t-&gt;op == OREGISTER) {
			tempalloc(&amp;r1, types[TFLOAT32]);
			gins(AFMOVFP, f, &amp;r1);
			gins(AFMOVF, &amp;r1, t);
			tempfree(&amp;r1);
			return;
		}
		if(f-&gt;op == OREGISTER)
			gins(AFMOVFP, f, t);
		else
			gins(AFMOVD, f, t);
		return;
	}

	gins(a, f, t);
	return;

rdst:
	// requires register destination
	regalloc(&amp;r1, t-&gt;type, t);
	gins(a, f, &amp;r1);
	gmove(&amp;r1, t);
	regfree(&amp;r1);
	return;

hard:
	// requires register intermediate
	regalloc(&amp;r1, cvt, t);
	gmove(f, &amp;r1);
	gmove(&amp;r1, t);
	regfree(&amp;r1);
	return;

hardmem:
	// requires memory intermediate
	tempalloc(&amp;r1, cvt);
	gmove(f, &amp;r1);
	gmove(&amp;r1, t);
	tempfree(&amp;r1);
	return;

fatal:
	// should not happen
	fatal(&#34;gmove %N -&gt; %N&#34;, f, t);
}

int
samaddr(Node *f, Node *t)
{

	if(f-&gt;op != t-&gt;op)
		return 0;

	switch(f-&gt;op) {
	case OREGISTER:
		if(f-&gt;val.u.reg != t-&gt;val.u.reg)
			break;
		return 1;
	}
	return 0;
}
/*
 * generate one instruction:
 *	as f, t
 */
Prog*
gins(int as, Node *f, Node *t)
{
	Prog *p;
	Addr af, at;

	if(as == AFMOVF &amp;&amp; f &amp;&amp; f-&gt;op == OREGISTER &amp;&amp; t &amp;&amp; t-&gt;op == OREGISTER)
		fatal(&#34;gins MOVF reg, reg&#34;);

	switch(as) {
	case AMOVB:
	case AMOVW:
	case AMOVL:
		if(f != N &amp;&amp; t != N &amp;&amp; samaddr(f, t))
			return nil;
	}

	memset(&amp;af, 0, sizeof af);
	memset(&amp;at, 0, sizeof at);
	if(f != N)
		naddr(f, &amp;af, 1);
	if(t != N)
		naddr(t, &amp;at, 1);
	p = prog(as);
	if(f != N)
		p-&gt;from = af;
	if(t != N)
		p-&gt;to = at;
	if(debug[&#39;g&#39;])
		print(&#34;%P\n&#34;, p);
	return p;
}

static void
checkoffset(Addr *a, int canemitcode)
{
	Prog *p;

	if(a-&gt;offset &lt; unmappedzero)
		return;
	if(!canemitcode)
		fatal(&#34;checkoffset %#llx, cannot emit code&#34;, a-&gt;offset);

	// cannot rely on unmapped nil page at 0 to catch
	// reference with large offset.  instead, emit explicit
	// test of 0(reg).
	p = gins(ATESTB, nodintconst(0), N);
	p-&gt;to = *a;
	p-&gt;to.offset = 0;
}

/*
 * generate code to compute n;
 * make a refer to result.
 */
void
naddr(Node *n, Addr *a, int canemitcode)
{
	a-&gt;scale = 0;
	a-&gt;index = D_NONE;
	a-&gt;type = D_NONE;
	a-&gt;gotype = S;
	if(n == N)
		return;

	switch(n-&gt;op) {
	default:
		fatal(&#34;naddr: bad %O %D&#34;, n-&gt;op, a);
		break;

	case OREGISTER:
		a-&gt;type = n-&gt;val.u.reg;
		a-&gt;sym = S;
		break;

	case OINDREG:
		a-&gt;type = n-&gt;val.u.reg+D_INDIR;
		a-&gt;sym = n-&gt;sym;
		a-&gt;offset = n-&gt;xoffset;
		break;

	case OPARAM:
		// n-&gt;left is PHEAP ONAME for stack parameter.
		// compute address of actual parameter on stack.
		a-&gt;etype = n-&gt;left-&gt;type-&gt;etype;
		a-&gt;width = n-&gt;left-&gt;type-&gt;width;
		a-&gt;offset = n-&gt;xoffset;
		a-&gt;sym = n-&gt;left-&gt;sym;
		a-&gt;type = D_PARAM;
		break;

	case ONAME:
		a-&gt;etype = 0;
		a-&gt;width = 0;
		if(n-&gt;type != T) {
			a-&gt;etype = simtype[n-&gt;type-&gt;etype];
			a-&gt;width = n-&gt;type-&gt;width;
			a-&gt;gotype = ngotype(n);
		}
		a-&gt;offset = n-&gt;xoffset;
		a-&gt;sym = n-&gt;sym;
		if(a-&gt;sym == S)
			a-&gt;sym = lookup(&#34;.noname&#34;);
		if(n-&gt;method) {
			if(n-&gt;type != T)
			if(n-&gt;type-&gt;sym != S)
			if(n-&gt;type-&gt;sym-&gt;package != nil)
				a-&gt;sym = pkglookup(a-&gt;sym-&gt;name, n-&gt;type-&gt;sym-&gt;package);
		}

		switch(n-&gt;class) {
		default:
			fatal(&#34;naddr: ONAME class %S %d\n&#34;, n-&gt;sym, n-&gt;class);
		case PEXTERN:
			a-&gt;type = D_EXTERN;
			break;
		case PAUTO:
			a-&gt;type = D_AUTO;
			break;
		case PPARAM:
		case PPARAMOUT:
			a-&gt;type = D_PARAM;
			break;
		case PFUNC:
			a-&gt;index = D_EXTERN;
			a-&gt;type = D_ADDR;
			break;
		}
		break;

	case OLITERAL:
		switch(n-&gt;val.ctype) {
		default:
			fatal(&#34;naddr: const %lT&#34;, n-&gt;type);
			break;
		case CTFLT:
			a-&gt;type = D_FCONST;
			a-&gt;dval = mpgetflt(n-&gt;val.u.fval);
			break;
		case CTINT:
			a-&gt;sym = S;
			a-&gt;type = D_CONST;
			a-&gt;offset = mpgetfix(n-&gt;val.u.xval);
			break;
		case CTSTR:
			datagostring(n-&gt;val.u.sval, a);
			break;
		case CTBOOL:
			a-&gt;sym = S;
			a-&gt;type = D_CONST;
			a-&gt;offset = n-&gt;val.u.bval;
			break;
		case CTNIL:
			a-&gt;sym = S;
			a-&gt;type = D_CONST;
			a-&gt;offset = 0;
			break;
		}
		break;

	case OADDR:
		naddr(n-&gt;left, a, canemitcode);
		if(a-&gt;type &gt;= D_INDIR) {
			a-&gt;type -= D_INDIR;
			break;
		}
		if(a-&gt;type == D_EXTERN || a-&gt;type == D_STATIC ||
		   a-&gt;type == D_AUTO || a-&gt;type == D_PARAM)
			if(a-&gt;index == D_NONE) {
				a-&gt;index = a-&gt;type;
				a-&gt;type = D_ADDR;
				break;
			}
		fatal(&#34;naddr: OADDR\n&#34;);

	case OLEN:
		// len of string or slice
		naddr(n-&gt;left, a, canemitcode);
		a-&gt;offset += Array_nel;
		if(a-&gt;offset &gt;= unmappedzero &amp;&amp; a-&gt;offset-Array_nel &lt; unmappedzero)
			checkoffset(a, canemitcode);
		break;

	case OCAP:
		// cap of string or slice
		naddr(n-&gt;left, a, canemitcode);
		a-&gt;offset += Array_cap;
		if(a-&gt;offset &gt;= unmappedzero &amp;&amp; a-&gt;offset-Array_nel &lt; unmappedzero)
			checkoffset(a, canemitcode);
		break;

//	case OADD:
//		if(n-&gt;right-&gt;op == OLITERAL) {
//			v = n-&gt;right-&gt;vconst;
//			naddr(n-&gt;left, a, canemitcode);
//		} else
//		if(n-&gt;left-&gt;op == OLITERAL) {
//			v = n-&gt;left-&gt;vconst;
//			naddr(n-&gt;right, a, canemitcode);
//		} else
//			goto bad;
//		a-&gt;offset += v;
//		break;

	}
}

int
dotaddable(Node *n, Node *n1)
{
	int o, oary[10];
	Node *nn;

	if(n-&gt;op != ODOT)
		return 0;

	o = dotoffset(n, oary, &amp;nn);
	if(nn != N &amp;&amp; nn-&gt;addable &amp;&amp; o == 1 &amp;&amp; oary[0] &gt;= 0) {
		*n1 = *nn;
		n1-&gt;type = n-&gt;type;
		n1-&gt;xoffset += oary[0];
		return 1;
	}
	return 0;
}

void
sudoclean(void)
{
}

int
sudoaddable(int as, Node *n, Addr *a)
{
	return 0;
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
