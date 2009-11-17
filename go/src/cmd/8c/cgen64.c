<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8c/cgen64.c</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/8c/cgen64.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/8c/cgen64.c
// http://code.google.com/p/inferno-os/source/browse/utils/8c/cgen64.c
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

#include &#34;gc.h&#34;

void
zeroregm(Node *n)
{
	gins(AMOVL, nodconst(0), n);
}

/* do we need to load the address of a vlong? */
int
vaddr(Node *n, int a)
{
	switch(n-&gt;op) {
	case ONAME:
		if(a)
			return 1;
		return !(n-&gt;class == CEXTERN || n-&gt;class == CGLOBL || n-&gt;class == CSTATIC);

	case OCONST:
	case OREGISTER:
	case OINDREG:
		return 1;
	}
	return 0;
}

int32
hi64v(Node *n)
{
	if(align(0, types[TCHAR], Aarg1))	/* isbigendian */
		return (int32)(n-&gt;vconst) &amp; ~0L;
	else
		return (int32)((uvlong)n-&gt;vconst&gt;&gt;32) &amp; ~0L;
}

int32
lo64v(Node *n)
{
	if(align(0, types[TCHAR], Aarg1))	/* isbigendian */
		return (int32)((uvlong)n-&gt;vconst&gt;&gt;32) &amp; ~0L;
	else
		return (int32)(n-&gt;vconst) &amp; ~0L;
}

Node *
hi64(Node *n)
{
	return nodconst(hi64v(n));
}

Node *
lo64(Node *n)
{
	return nodconst(lo64v(n));
}

static Node *
anonreg(void)
{
	Node *n;

	n = new(OREGISTER, Z, Z);
	n-&gt;reg = D_NONE;
	n-&gt;type = types[TLONG];
	return n;
}

static Node *
regpair(Node *n, Node *t)
{
	Node *r;

	if(n != Z &amp;&amp; n-&gt;op == OREGPAIR)
		return n;
	r = new(OREGPAIR, anonreg(), anonreg());
	if(n != Z)
		r-&gt;type = n-&gt;type;
	else
		r-&gt;type = t-&gt;type;
	return r;
}

static void
evacaxdx(Node *r)
{
	Node nod1, nod2;

	if(r-&gt;reg == D_AX || r-&gt;reg == D_DX) {
		reg[D_AX]++;
		reg[D_DX]++;
		/*
		 * this is just an optim that should
		 * check for spill
		 */
		r-&gt;type = types[TULONG];
		regalloc(&amp;nod1, r, Z);
		nodreg(&amp;nod2, Z, r-&gt;reg);
		gins(AMOVL, &amp;nod2, &amp;nod1);
		regfree(r);
		r-&gt;reg = nod1.reg;
		reg[D_AX]--;
		reg[D_DX]--;
	}
}

/* lazy instantiation of register pair */
static int
instpair(Node *n, Node *l)
{
	int r;

	r = 0;
	if(n-&gt;left-&gt;reg == D_NONE) {
		if(l != Z) {
			n-&gt;left-&gt;reg = l-&gt;reg;
			r = 1;
		}
		else
			regalloc(n-&gt;left, n-&gt;left, Z);
	}
	if(n-&gt;right-&gt;reg == D_NONE)
		regalloc(n-&gt;right, n-&gt;right, Z);
	return r;
}

static void
zapreg(Node *n)
{
	if(n-&gt;reg != D_NONE) {
		regfree(n);
		n-&gt;reg = D_NONE;
	}
}

static void
freepair(Node *n)
{
	regfree(n-&gt;left);
	regfree(n-&gt;right);
}

/* n is not OREGPAIR, nn is */
void
loadpair(Node *n, Node *nn)
{
	Node nod;

	instpair(nn, Z);
	if(n-&gt;op == OCONST) {
		gins(AMOVL, lo64(n), nn-&gt;left);
		n-&gt;xoffset += SZ_LONG;
		gins(AMOVL, hi64(n), nn-&gt;right);
		n-&gt;xoffset -= SZ_LONG;
		return;
	}
	if(!vaddr(n, 0)) {
		/* steal the right register for the laddr */
		nod = regnode;
		nod.reg = nn-&gt;right-&gt;reg;
		lcgen(n, &amp;nod);
		n = &amp;nod;
		regind(n, n);
		n-&gt;xoffset = 0;
	}
	gins(AMOVL, n, nn-&gt;left);
	n-&gt;xoffset += SZ_LONG;
	gins(AMOVL, n, nn-&gt;right);
	n-&gt;xoffset -= SZ_LONG;
}

/* n is OREGPAIR, nn is not */
static void
storepair(Node *n, Node *nn, int f)
{
	Node nod;

	if(!vaddr(nn, 0)) {
		reglcgen(&amp;nod, nn, Z);
		nn = &amp;nod;
	}
	gins(AMOVL, n-&gt;left, nn);
	nn-&gt;xoffset += SZ_LONG;
	gins(AMOVL, n-&gt;right, nn);
	nn-&gt;xoffset -= SZ_LONG;
	if(nn == &amp;nod)
		regfree(&amp;nod);
	if(f)
		freepair(n);
}

enum
{
/* 4 only, see WW */
	WNONE	= 0,
	WCONST,
	WADDR,
	WHARD,
};

static int
whatof(Node *n, int a)
{
	if(n-&gt;op == OCONST)
		return WCONST;
	return !vaddr(n, a) ? WHARD : WADDR;
}

/* can upgrade an extern to addr for AND */
static int
reduxv(Node *n)
{
	return lo64v(n) == 0 || hi64v(n) == 0;
}

int
cond(int op)
{
	switch(op) {
	case OANDAND:
	case OOROR:
	case ONOT:
		return 1;

	case OEQ:
	case ONE:
	case OLE:
	case OLT:
	case OGE:
	case OGT:
	case OHI:
	case OHS:
	case OLO:
	case OLS:
		return 1;
	}
	return 0;
}

/*
 * for a func operand call it and then return
 * the safe node
 */
static Node *
vfunc(Node *n, Node *nn)
{
	Node *t;

	if(n-&gt;op != OFUNC)
		return n;
	t = new(0, Z, Z);
	if(nn == Z || nn == nodret)
		nn = n;
	regsalloc(t, nn);
	sugen(n, t, 8);
	return t;
}

/* try to steal a reg */
static int
getreg(Node **np, Node *t, int r)
{
	Node *n, *p;

	n = *np;
	if(n-&gt;reg == r) {
		p = new(0, Z, Z);
		regalloc(p, n, Z);
		gins(AMOVL, n, p);
		*t = *n;
		*np = p;
		return 1;
	}
	return 0;
}

static Node *
snarfreg(Node *n, Node *t, int r, Node *d, Node *c)
{
	if(n == Z || n-&gt;op != OREGPAIR || (!getreg(&amp;n-&gt;left, t, r) &amp;&amp; !getreg(&amp;n-&gt;right, t, r))) {
		if(nodreg(t, Z, r)) {
			regalloc(c, d, Z);
			gins(AMOVL, t, c);
			reg[r]++;
			return c;
		}
		reg[r]++;
	}
	return Z;
}

enum
{
	Vstart	= OEND,

	Vgo,
	Vamv,
	Vmv,
	Vzero,
	Vop,
	Vopx,
	Vins,
	Vins0,
	Vinsl,
	Vinsr,
	Vinsla,
	Vinsra,
	Vinsx,
	Vmul,
	Vshll,
	VT,
	VF,
	V_l_lo_f,
	V_l_hi_f,
	V_l_lo_t,
	V_l_hi_t,
	V_l_lo_u,
	V_l_hi_u,
	V_r_lo_f,
	V_r_hi_f,
	V_r_lo_t,
	V_r_hi_t,
	V_r_lo_u,
	V_r_hi_u,
	Vspazz,
	Vend,

	V_T0,
	V_T1,
	V_F0,
	V_F1,

	V_a0,
	V_a1,
	V_f0,
	V_f1,

	V_p0,
	V_p1,
	V_p2,
	V_p3,
	V_p4,

	V_s0,
	V_s1,
	V_s2,
	V_s3,
	V_s4,

	C00,
	C01,
	C31,
	C32,

	O_l_lo,
	O_l_hi,
	O_r_lo,
	O_r_hi,
	O_t_lo,
	O_t_hi,
	O_l,
	O_r,
	O_l_rp,
	O_r_rp,
	O_t_rp,
	O_r0,
	O_r1,
	O_Zop,

	O_a0,
	O_a1,

	V_C0,
	V_C1,

	V_S0,
	V_S1,

	VOPS	= 5,
	VLEN	= 5,
	VARGS	= 2,

	S00	= 0,
	Sc0,
	Sc1,
	Sc2,
	Sac3,
	Sac4,
	S10,

	SAgen	= 0,
	SAclo,
	SAc32,
	SAchi,
	SAdgen,
	SAdclo,
	SAdc32,
	SAdchi,

	B0c	= 0,
	Bca,
	Bac,

	T0i	= 0,
	Tii,

	Bop0	= 0,
	Bop1,
};

/*
 * _testv:
 * 	CMPL	lo,$0
 * 	JNE	true
 * 	CMPL	hi,$0
 * 	JNE	true
 * 	GOTO	false
 * false:
 * 	GOTO	code
 * true:
 * 	GOTO	patchme
 * code:
 */

static uchar	testi[][VLEN] =
{
	{Vop, ONE, O_l_lo, C00},
	{V_s0, Vop, ONE, O_l_hi, C00},
	{V_s1, Vgo, V_s2, Vgo, V_s3},
	{VF, V_p0, V_p1, VT, V_p2},
	{Vgo, V_p3},
	{VT, V_p0, V_p1, VF, V_p2},
	{Vend},
};

/* shift left general case */
static uchar	shll00[][VLEN] =
{
	{Vop, OGE, O_r, C32},
	{V_s0, Vinsl, ASHLL, O_r, O_l_rp},
	{Vins, ASHLL, O_r, O_l_lo, Vgo},
	{V_p0, V_s0},
	{Vins, ASHLL, O_r, O_l_lo},
	{Vins, AMOVL, O_l_lo, O_l_hi},
	{Vzero, O_l_lo, V_p0, Vend},
};

/* shift left rp, const &lt; 32 */
static uchar	shllc0[][VLEN] =
{
	{Vinsl, ASHLL, O_r, O_l_rp},
	{Vshll, O_r, O_l_lo, Vend},
};

/* shift left rp, const == 32 */
static uchar	shllc1[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_l_hi},
	{Vzero, O_l_lo, Vend},
};

/* shift left rp, const &gt; 32 */
static uchar	shllc2[][VLEN] =
{
	{Vshll, O_r, O_l_lo},
	{Vins, AMOVL, O_l_lo, O_l_hi},
	{Vzero, O_l_lo, Vend},
};

/* shift left addr, const == 32 */
static uchar	shllac3[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_t_hi},
	{Vzero, O_t_lo, Vend},
};

/* shift left addr, const &gt; 32 */
static uchar	shllac4[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_t_hi},
	{Vshll, O_r, O_t_hi},
	{Vzero, O_t_lo, Vend},
};

/* shift left of constant */
static uchar	shll10[][VLEN] =
{
	{Vop, OGE, O_r, C32},
	{V_s0, Vins, AMOVL, O_l_lo, O_t_lo},
	{Vins, AMOVL, O_l_hi, O_t_hi},
	{Vinsl, ASHLL, O_r, O_t_rp},
	{Vins, ASHLL, O_r, O_t_lo, Vgo},
	{V_p0, V_s0},
	{Vins, AMOVL, O_l_lo, O_t_hi},
	{V_l_lo_t, Vins, ASHLL, O_r, O_t_hi},
	{Vzero, O_t_lo, V_p0, Vend},
};

static uchar	(*shlltab[])[VLEN] =
{
	shll00,
	shllc0,
	shllc1,
	shllc2,
	shllac3,
	shllac4,
	shll10,
};

/* shift right general case */
static uchar	shrl00[][VLEN] =
{
	{Vop, OGE, O_r, C32},
	{V_s0, Vinsr, ASHRL, O_r, O_l_rp},
	{Vins, O_a0, O_r, O_l_hi, Vgo},
	{V_p0, V_s0},
	{Vins, O_a0, O_r, O_l_hi},
	{Vins, AMOVL, O_l_hi, O_l_lo},
	{V_T1, Vzero, O_l_hi},
	{V_F1, Vins, ASARL, C31, O_l_hi},
	{V_p0, Vend},
};

/* shift right rp, const &lt; 32 */
static uchar	shrlc0[][VLEN] =
{
	{Vinsr, ASHRL, O_r, O_l_rp},
	{Vins, O_a0, O_r, O_l_hi, Vend},
};

/* shift right rp, const == 32 */
static uchar	shrlc1[][VLEN] =
{
	{Vins, AMOVL, O_l_hi, O_l_lo},
	{V_T1, Vzero, O_l_hi},
	{V_F1, Vins, ASARL, C31, O_l_hi},
	{Vend},
};

/* shift right rp, const &gt; 32 */
static uchar	shrlc2[][VLEN] =
{
	{Vins, O_a0, O_r, O_l_hi},
	{Vins, AMOVL, O_l_hi, O_l_lo},
	{V_T1, Vzero, O_l_hi},
	{V_F1, Vins, ASARL, C31, O_l_hi},
	{Vend},
};

/* shift right addr, const == 32 */
static uchar	shrlac3[][VLEN] =
{
	{Vins, AMOVL, O_l_hi, O_t_lo},
	{V_T1, Vzero, O_t_hi},
	{V_F1, Vins, AMOVL, O_t_lo, O_t_hi},
	{V_F1, Vins, ASARL, C31, O_t_hi},
	{Vend},
};

/* shift right addr, const &gt; 32 */
static uchar	shrlac4[][VLEN] =
{
	{Vins, AMOVL, O_l_hi, O_t_lo},
	{Vins, O_a0, O_r, O_t_lo},
	{V_T1, Vzero, O_t_hi},
	{V_F1, Vins, AMOVL, O_t_lo, O_t_hi},
	{V_F1, Vins, ASARL, C31, O_t_hi},
	{Vend},
};

/* shift right of constant */
static uchar	shrl10[][VLEN] =
{
	{Vop, OGE, O_r, C32},
	{V_s0, Vins, AMOVL, O_l_lo, O_t_lo},
	{Vins, AMOVL, O_l_hi, O_t_hi},
	{Vinsr, ASHRL, O_r, O_t_rp},
	{Vins, O_a0, O_r, O_t_hi, Vgo},
	{V_p0, V_s0},
	{Vins, AMOVL, O_l_hi, O_t_lo},
	{V_l_hi_t, Vins, O_a0, O_r, O_t_lo},
	{V_l_hi_u, V_S1},
	{V_T1, Vzero, O_t_hi, V_p0},
	{V_F1, Vins, AMOVL, O_t_lo, O_t_hi},
	{V_F1, Vins, ASARL, C31, O_t_hi},
	{Vend},
};

static uchar	(*shrltab[])[VLEN] =
{
	shrl00,
	shrlc0,
	shrlc1,
	shrlc2,
	shrlac3,
	shrlac4,
	shrl10,
};

/* shift asop left general case */
static uchar	asshllgen[][VLEN] =
{
	{V_a0, V_a1},
	{Vop, OGE, O_r, C32},
	{V_s0, Vins, AMOVL, O_l_lo, O_r0},
	{Vins, AMOVL, O_l_hi, O_r1},
	{Vinsla, ASHLL, O_r, O_r0},
	{Vins, ASHLL, O_r, O_r0},
	{Vins, AMOVL, O_r1, O_l_hi},
	{Vins, AMOVL, O_r0, O_l_lo, Vgo},
	{V_p0, V_s0},
	{Vins, AMOVL, O_l_lo, O_r0},
	{Vzero, O_l_lo},
	{Vins, ASHLL, O_r, O_r0},
	{Vins, AMOVL, O_r0, O_l_hi, V_p0},
	{V_f0, V_f1, Vend},
};

/* shift asop left, const &lt; 32 */
static uchar	asshllclo[][VLEN] =
{
	{V_a0, V_a1},
	{Vins, AMOVL, O_l_lo, O_r0},
	{Vins, AMOVL, O_l_hi, O_r1},
	{Vinsla, ASHLL, O_r, O_r0},
	{Vshll, O_r, O_r0},
	{Vins, AMOVL, O_r1, O_l_hi},
	{Vins, AMOVL, O_r0, O_l_lo},
	{V_f0, V_f1, Vend},
};

/* shift asop left, const == 32 */
static uchar	asshllc32[][VLEN] =
{
	{V_a0},
	{Vins, AMOVL, O_l_lo, O_r0},
	{Vzero, O_l_lo},
	{Vins, AMOVL, O_r0, O_l_hi},
	{V_f0, Vend},
};

/* shift asop left, const &gt; 32 */
static uchar	asshllchi[][VLEN] =
{
	{V_a0},
	{Vins, AMOVL, O_l_lo, O_r0},
	{Vzero, O_l_lo},
	{Vshll, O_r, O_r0},
	{Vins, AMOVL, O_r0, O_l_hi},
	{V_f0, Vend},
};

/* shift asop dest left general case */
static uchar	asdshllgen[][VLEN] =
{
	{Vop, OGE, O_r, C32},
	{V_s0, Vins, AMOVL, O_l_lo, O_t_lo},
	{Vins, AMOVL, O_l_hi, O_t_hi},
	{Vinsl, ASHLL, O_r, O_t_rp},
	{Vins, ASHLL, O_r, O_t_lo},
	{Vins, AMOVL, O_t_hi, O_l_hi},
	{Vins, AMOVL, O_t_lo, O_l_lo, Vgo},
	{V_p0, V_s0},
	{Vins, AMOVL, O_l_lo, O_t_hi},
	{Vzero, O_l_lo},
	{Vins, ASHLL, O_r, O_t_hi},
	{Vzero, O_t_lo},
	{Vins, AMOVL, O_t_hi, O_l_hi, V_p0},
	{Vend},
};

/* shift asop dest left, const &lt; 32 */
static uchar	asdshllclo[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_t_lo},
	{Vins, AMOVL, O_l_hi, O_t_hi},
	{Vinsl, ASHLL, O_r, O_t_rp},
	{Vshll, O_r, O_t_lo},
	{Vins, AMOVL, O_t_hi, O_l_hi},
	{Vins, AMOVL, O_t_lo, O_l_lo},
	{Vend},
};

/* shift asop dest left, const == 32 */
static uchar	asdshllc32[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_t_hi},
	{Vzero, O_t_lo},
	{Vins, AMOVL, O_t_hi, O_l_hi},
	{Vins, AMOVL, O_t_lo, O_l_lo},
	{Vend},
};

/* shift asop dest, const &gt; 32 */
static uchar	asdshllchi[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_t_hi},
	{Vzero, O_t_lo},
	{Vshll, O_r, O_t_hi},
	{Vins, AMOVL, O_t_lo, O_l_lo},
	{Vins, AMOVL, O_t_hi, O_l_hi},
	{Vend},
};

static uchar	(*asshlltab[])[VLEN] =
{
	asshllgen,
	asshllclo,
	asshllc32,
	asshllchi,
	asdshllgen,
	asdshllclo,
	asdshllc32,
	asdshllchi,
};

/* shift asop right general case */
static uchar	asshrlgen[][VLEN] =
{
	{V_a0, V_a1},
	{Vop, OGE, O_r, C32},
	{V_s0, Vins, AMOVL, O_l_lo, O_r0},
	{Vins, AMOVL, O_l_hi, O_r1},
	{Vinsra, ASHRL, O_r, O_r0},
	{Vinsx, Bop0, O_r, O_r1},
	{Vins, AMOVL, O_r0, O_l_lo},
	{Vins, AMOVL, O_r1, O_l_hi, Vgo},
	{V_p0, V_s0},
	{Vins, AMOVL, O_l_hi, O_r0},
	{Vinsx, Bop0, O_r, O_r0},
	{V_T1, Vzero, O_l_hi},
	{Vins, AMOVL, O_r0, O_l_lo},
	{V_F1, Vins, ASARL, C31, O_r0},
	{V_F1, Vins, AMOVL, O_r0, O_l_hi},
	{V_p0, V_f0, V_f1, Vend},
};

/* shift asop right, const &lt; 32 */
static uchar	asshrlclo[][VLEN] =
{
	{V_a0, V_a1},
	{Vins, AMOVL, O_l_lo, O_r0},
	{Vins, AMOVL, O_l_hi, O_r1},
	{Vinsra, ASHRL, O_r, O_r0},
	{Vinsx, Bop0, O_r, O_r1},
	{Vins, AMOVL, O_r0, O_l_lo},
	{Vins, AMOVL, O_r1, O_l_hi},
	{V_f0, V_f1, Vend},
};

/* shift asop right, const == 32 */
static uchar	asshrlc32[][VLEN] =
{
	{V_a0},
	{Vins, AMOVL, O_l_hi, O_r0},
	{V_T1, Vzero, O_l_hi},
	{Vins, AMOVL, O_r0, O_l_lo},
	{V_F1, Vins, ASARL, C31, O_r0},
	{V_F1, Vins, AMOVL, O_r0, O_l_hi},
	{V_f0, Vend},
};

/* shift asop right, const &gt; 32 */
static uchar	asshrlchi[][VLEN] =
{
	{V_a0},
	{Vins, AMOVL, O_l_hi, O_r0},
	{V_T1, Vzero, O_l_hi},
	{Vinsx, Bop0, O_r, O_r0},
	{Vins, AMOVL, O_r0, O_l_lo},
	{V_F1, Vins, ASARL, C31, O_r0},
	{V_F1, Vins, AMOVL, O_r0, O_l_hi},
	{V_f0, Vend},
};

/* shift asop dest right general case */
static uchar	asdshrlgen[][VLEN] =
{
	{Vop, OGE, O_r, C32},
	{V_s0, Vins, AMOVL, O_l_lo, O_t_lo},
	{Vins, AMOVL, O_l_hi, O_t_hi},
	{Vinsr, ASHRL, O_r, O_t_rp},
	{Vinsx, Bop0, O_r, O_t_hi},
	{Vins, AMOVL, O_t_lo, O_l_lo},
	{Vins, AMOVL, O_t_hi, O_l_hi, Vgo},
	{V_p0, V_s0},
	{Vins, AMOVL, O_l_hi, O_t_lo},
	{V_T1, Vzero, O_t_hi},
	{Vinsx, Bop0, O_r, O_t_lo},
	{V_F1, Vins, AMOVL, O_t_lo, O_t_hi},
	{V_F1, Vins, ASARL, C31, O_t_hi},
	{Vins, AMOVL, O_t_hi, O_l_hi, V_p0},
	{Vend},
};

/* shift asop dest right, const &lt; 32 */
static uchar	asdshrlclo[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_t_lo},
	{Vins, AMOVL, O_l_hi, O_t_hi},
	{Vinsr, ASHRL, O_r, O_t_rp},
	{Vinsx, Bop0, O_r, O_t_hi},
	{Vins, AMOVL, O_t_lo, O_l_lo},
	{Vins, AMOVL, O_t_hi, O_l_hi},
	{Vend},
};

/* shift asop dest right, const == 32 */
static uchar	asdshrlc32[][VLEN] =
{
	{Vins, AMOVL, O_l_hi, O_t_lo},
	{V_T1, Vzero, O_t_hi},
	{V_F1, Vins, AMOVL, O_t_lo, O_t_hi},
	{V_F1, Vins, ASARL, C31, O_t_hi},
	{Vins, AMOVL, O_t_lo, O_l_lo},
	{Vins, AMOVL, O_t_hi, O_l_hi},
	{Vend},
};

/* shift asop dest, const &gt; 32 */
static uchar	asdshrlchi[][VLEN] =
{
	{Vins, AMOVL, O_l_hi, O_t_lo},
	{V_T1, Vzero, O_t_hi},
	{Vinsx, Bop0, O_r, O_t_lo},
	{V_T1, Vins, AMOVL, O_t_hi, O_l_hi},
	{V_T1, Vins, AMOVL, O_t_lo, O_l_lo},
	{V_F1, Vins, AMOVL, O_t_lo, O_t_hi},
	{V_F1, Vins, ASARL, C31, O_t_hi},
	{V_F1, Vins, AMOVL, O_t_lo, O_l_lo},
	{V_F1, Vins, AMOVL, O_t_hi, O_l_hi},
	{Vend},
};

static uchar	(*asshrltab[])[VLEN] =
{
	asshrlgen,
	asshrlclo,
	asshrlc32,
	asshrlchi,
	asdshrlgen,
	asdshrlclo,
	asdshrlc32,
	asdshrlchi,
};

static uchar	shrlargs[]	= { ASHRL, 1 };
static uchar	sarlargs[]	= { ASARL, 0 };

/* ++ -- */
static uchar	incdec[][VLEN] =
{
	{Vinsx, Bop0, C01, O_l_lo},
	{Vinsx, Bop1, C00, O_l_hi, Vend},
};

/* ++ -- *p */
static uchar	incdecpre[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_t_lo},
	{Vins, AMOVL, O_l_hi, O_t_hi},
	{Vinsx, Bop0, C01, O_t_lo},
	{Vinsx, Bop1, C00, O_t_hi},
	{Vins, AMOVL, O_t_lo, O_l_lo},
	{Vins, AMOVL, O_t_hi, O_l_hi, Vend},
};

/* *p ++ -- */
static uchar	incdecpost[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_t_lo},
	{Vins, AMOVL, O_l_hi, O_t_hi},
	{Vinsx, Bop0, C01, O_l_lo},
	{Vinsx, Bop1, C00, O_l_hi, Vend},
};

/* binop rp, rp */
static uchar	binop00[][VLEN] =
{
	{Vinsx, Bop0, O_r_lo, O_l_lo},
	{Vinsx, Bop1, O_r_hi, O_l_hi, Vend},
	{Vend},
};

/* binop rp, addr */
static uchar	binoptmp[][VLEN] =
{
	{V_a0, Vins, AMOVL, O_r_lo, O_r0},
	{Vinsx, Bop0, O_r0, O_l_lo},
	{Vins, AMOVL, O_r_hi, O_r0},
	{Vinsx, Bop1, O_r0, O_l_hi},
	{V_f0, Vend},
};

/* binop t = *a op *b */
static uchar	binop11[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_t_lo},
	{Vinsx, Bop0, O_r_lo, O_t_lo},
	{Vins, AMOVL, O_l_hi, O_t_hi},
	{Vinsx, Bop1, O_r_hi, O_t_hi, Vend},
};

/* binop t = rp +- c */
static uchar	add0c[][VLEN] =
{
	{V_r_lo_t, Vinsx, Bop0, O_r_lo, O_l_lo},
	{V_r_lo_f, Vamv, Bop0, Bop1},
	{Vinsx, Bop1, O_r_hi, O_l_hi},
	{Vend},
};

/* binop t = rp &amp; c */
static uchar	and0c[][VLEN] =
{
	{V_r_lo_t, Vinsx, Bop0, O_r_lo, O_l_lo},
	{V_r_lo_f, Vins, AMOVL, C00, O_l_lo},
	{V_r_hi_t, Vinsx, Bop1, O_r_hi, O_l_hi},
	{V_r_hi_f, Vins, AMOVL, C00, O_l_hi},
	{Vend},
};

/* binop t = rp | c */
static uchar	or0c[][VLEN] =
{
	{V_r_lo_t, Vinsx, Bop0, O_r_lo, O_l_lo},
	{V_r_hi_t, Vinsx, Bop1, O_r_hi, O_l_hi},
	{Vend},
};

/* binop t = c - rp */
static uchar	sub10[][VLEN] =
{
	{V_a0, Vins, AMOVL, O_l_lo, O_r0},
	{Vinsx, Bop0, O_r_lo, O_r0},
	{Vins, AMOVL, O_l_hi, O_r_lo},
	{Vinsx, Bop1, O_r_hi, O_r_lo},
	{Vspazz, V_f0, Vend},
};

/* binop t = c + *b */
static uchar	addca[][VLEN] =
{
	{Vins, AMOVL, O_r_lo, O_t_lo},
	{V_l_lo_t, Vinsx, Bop0, O_l_lo, O_t_lo},
	{V_l_lo_f, Vamv, Bop0, Bop1},
	{Vins, AMOVL, O_r_hi, O_t_hi},
	{Vinsx, Bop1, O_l_hi, O_t_hi},
	{Vend},
};

/* binop t = c &amp; *b */
static uchar	andca[][VLEN] =
{
	{V_l_lo_t, Vins, AMOVL, O_r_lo, O_t_lo},
	{V_l_lo_t, Vinsx, Bop0, O_l_lo, O_t_lo},
	{V_l_lo_f, Vzero, O_t_lo},
	{V_l_hi_t, Vins, AMOVL, O_r_hi, O_t_hi},
	{V_l_hi_t, Vinsx, Bop1, O_l_hi, O_t_hi},
	{V_l_hi_f, Vzero, O_t_hi},
	{Vend},
};

/* binop t = c | *b */
static uchar	orca[][VLEN] =
{
	{Vins, AMOVL, O_r_lo, O_t_lo},
	{V_l_lo_t, Vinsx, Bop0, O_l_lo, O_t_lo},
	{Vins, AMOVL, O_r_hi, O_t_hi},
	{V_l_hi_t, Vinsx, Bop1, O_l_hi, O_t_hi},
	{Vend},
};

/* binop t = c - *b */
static uchar	subca[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_t_lo},
	{Vins, AMOVL, O_l_hi, O_t_hi},
	{Vinsx, Bop0, O_r_lo, O_t_lo},
	{Vinsx, Bop1, O_r_hi, O_t_hi},
	{Vend},
};

/* binop t = *a +- c */
static uchar	addac[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_t_lo},
	{V_r_lo_t, Vinsx, Bop0, O_r_lo, O_t_lo},
	{V_r_lo_f, Vamv, Bop0, Bop1},
	{Vins, AMOVL, O_l_hi, O_t_hi},
	{Vinsx, Bop1, O_r_hi, O_t_hi},
	{Vend},
};

/* binop t = *a | c */
static uchar	orac[][VLEN] =
{
	{Vins, AMOVL, O_l_lo, O_t_lo},
	{V_r_lo_t, Vinsx, Bop0, O_r_lo, O_t_lo},
	{Vins, AMOVL, O_l_hi, O_t_hi},
	{V_r_hi_t, Vinsx, Bop1, O_r_hi, O_t_hi},
	{Vend},
};

/* binop t = *a &amp; c */
static uchar	andac[][VLEN] =
{
	{V_r_lo_t, Vins, AMOVL, O_l_lo, O_t_lo},
	{V_r_lo_t, Vinsx, Bop0, O_r_lo, O_t_lo},
	{V_r_lo_f, Vzero, O_t_lo},
	{V_r_hi_t, Vins, AMOVL, O_l_hi, O_t_hi},
	{V_r_hi_t, Vinsx, Bop0, O_r_hi, O_t_hi},
	{V_r_hi_f, Vzero, O_t_hi},
	{Vend},
};

static uchar	ADDargs[]	= { AADDL, AADCL };
static uchar	ANDargs[]	= { AANDL, AANDL };
static uchar	ORargs[]	= { AORL, AORL };
static uchar	SUBargs[]	= { ASUBL, ASBBL };
static uchar	XORargs[]	= { AXORL, AXORL };

static uchar	(*ADDtab[])[VLEN] =
{
	add0c, addca, addac,
};

static uchar	(*ANDtab[])[VLEN] =
{
	and0c, andca, andac,
};

static uchar	(*ORtab[])[VLEN] =
{
	or0c, orca, orac,
};

static uchar	(*SUBtab[])[VLEN] =
{
	add0c, subca, addac,
};

/* mul of const32 */
static uchar	mulc32[][VLEN] =
{
	{V_a0, Vop, ONE, O_l_hi, C00},
	{V_s0, Vins, AMOVL, O_r_lo, O_r0},
	{Vins, AMULL, O_r0, O_Zop},
	{Vgo, V_p0, V_s0},
	{Vins, AMOVL, O_l_hi, O_r0},
	{Vmul, O_r_lo, O_r0},
	{Vins, AMOVL, O_r_lo, O_l_hi},
	{Vins, AMULL, O_l_hi, O_Zop},
	{Vins, AADDL, O_r0, O_l_hi},
	{V_f0, V_p0, Vend},
};

/* mul of const64 */
static uchar	mulc64[][VLEN] =
{
	{V_a0, Vins, AMOVL, O_r_hi, O_r0},
	{Vop, OOR, O_l_hi, O_r0},
	{Vop, ONE, O_r0, C00},
	{V_s0, Vins, AMOVL, O_r_lo, O_r0},
	{Vins, AMULL, O_r0, O_Zop},
	{Vgo, V_p0, V_s0},
	{Vmul, O_r_lo, O_l_hi},
	{Vins, AMOVL, O_l_lo, O_r0},
	{Vmul, O_r_hi, O_r0},
	{Vins, AADDL, O_l_hi, O_r0},
	{Vins, AMOVL, O_r_lo, O_l_hi},
	{Vins, AMULL, O_l_hi, O_Zop},
	{Vins, AADDL, O_r0, O_l_hi},
	{V_f0, V_p0, Vend},
};

/* mul general */
static uchar	mull[][VLEN] =
{
	{V_a0, Vins, AMOVL, O_r_hi, O_r0},
	{Vop, OOR, O_l_hi, O_r0},
	{Vop, ONE, O_r0, C00},
	{V_s0, Vins, AMOVL, O_r_lo, O_r0},
	{Vins, AMULL, O_r0, O_Zop},
	{Vgo, V_p0, V_s0},
	{Vins, AIMULL, O_r_lo, O_l_hi},
	{Vins, AMOVL, O_l_lo, O_r0},
	{Vins, AIMULL, O_r_hi, O_r0},
	{Vins, AADDL, O_l_hi, O_r0},
	{Vins, AMOVL, O_r_lo, O_l_hi},
	{Vins, AMULL, O_l_hi, O_Zop},
	{Vins, AADDL, O_r0, O_l_hi},
	{V_f0, V_p0, Vend},
};

/* cast rp l to rp t */
static uchar	castrp[][VLEN] =
{
	{Vmv, O_l, O_t_lo},
	{VT, Vins, AMOVL, O_t_lo, O_t_hi},
	{VT, Vins, ASARL, C31, O_t_hi},
	{VF, Vzero, O_t_hi},
	{Vend},
};

/* cast rp l to addr t */
static uchar	castrpa[][VLEN] =
{
	{VT, V_a0, Vmv, O_l, O_r0},
	{VT, Vins, AMOVL, O_r0, O_t_lo},
	{VT, Vins, ASARL, C31, O_r0},
	{VT, Vins, AMOVL, O_r0, O_t_hi},
	{VT, V_f0},
	{VF, Vmv, O_l, O_t_lo},
	{VF, Vzero, O_t_hi},
	{Vend},
};

static uchar	netab0i[][VLEN] =
{
	{Vop, ONE, O_l_lo, O_r_lo},
	{V_s0, Vop, ONE, O_l_hi, O_r_hi},
	{V_s1, Vgo, V_s2, Vgo, V_s3},
	{VF, V_p0, V_p1, VT, V_p2},
	{Vgo, V_p3},
	{VT, V_p0, V_p1, VF, V_p2},
	{Vend},
};

static uchar	netabii[][VLEN] =
{
	{V_a0, Vins, AMOVL, O_l_lo, O_r0},
	{Vop, ONE, O_r0, O_r_lo},
	{V_s0, Vins, AMOVL, O_l_hi, O_r0},
	{Vop, ONE, O_r0, O_r_hi},
	{V_s1, Vgo, V_s2, Vgo, V_s3},
	{VF, V_p0, V_p1, VT, V_p2},
	{Vgo, V_p3},
	{VT, V_p0, V_p1, VF, V_p2},
	{V_f0, Vend},
};

static uchar	cmptab0i[][VLEN] =
{
	{Vopx, Bop0, O_l_hi, O_r_hi},
	{V_s0, Vins0, AJNE},
	{V_s1, Vopx, Bop1, O_l_lo, O_r_lo},
	{V_s2, Vgo, V_s3, Vgo, V_s4},
	{VT, V_p1, V_p3},
	{VF, V_p0, V_p2},
	{Vgo, V_p4},
	{VT, V_p0, V_p2},
	{VF, V_p1, V_p3},
	{Vend},
};

static uchar	cmptabii[][VLEN] =
{
	{V_a0, Vins, AMOVL, O_l_hi, O_r0},
	{Vopx, Bop0, O_r0, O_r_hi},
	{V_s0, Vins0, AJNE},
	{V_s1, Vins, AMOVL, O_l_lo, O_r0},
	{Vopx, Bop1, O_r0, O_r_lo},
	{V_s2, Vgo, V_s3, Vgo, V_s4},
	{VT, V_p1, V_p3},
	{VF, V_p0, V_p2},
	{Vgo, V_p4},
	{VT, V_p0, V_p2},
	{VF, V_p1, V_p3},
	{V_f0, Vend},
};

static uchar	(*NEtab[])[VLEN] =
{
	netab0i, netabii,
};

static uchar	(*cmptab[])[VLEN] =
{
	cmptab0i, cmptabii,
};

static uchar	GEargs[]	= { OGT, OHS };
static uchar	GTargs[]	= { OGT, OHI };
static uchar	HIargs[]	= { OHI, OHI };
static uchar	HSargs[]	= { OHI, OHS };

/* Big Generator */
static void
biggen(Node *l, Node *r, Node *t, int true, uchar code[][VLEN], uchar *a)
{
	int i, j, g, oc, op, lo, ro, to, xo, *xp;
	Type *lt;
	Prog *pr[VOPS];
	Node *ot, *tl, *tr, tmps[2];
	uchar *c, (*cp)[VLEN], args[VARGS];

	if(a != nil)
		memmove(args, a, VARGS);
//print(&#34;biggen %d %d %d\n&#34;, args[0], args[1], args[2]);
//if(l) prtree(l, &#34;l&#34;);
//if(r) prtree(r, &#34;r&#34;);
//if(t) prtree(t, &#34;t&#34;);
	lo = ro = to = 0;
	cp = code;

	for (;;) {
		c = *cp++;
		g = 1;
		i = 0;
//print(&#34;code %d %d %d %d %d\n&#34;, c[0], c[1], c[2], c[3], c[4]);
		for(;;) {
			switch(op = c[i]) {
			case Vgo:
				if(g)
					gbranch(OGOTO);
				i++;
				break;

			case Vamv:
				i += 3;
				if(i &gt; VLEN) {
					diag(l, &#34;bad Vop&#34;);
					return;
				}
				if(g)
					args[c[i - 1]] = args[c[i - 2]];
				break;

			case Vzero:
				i += 2;
				if(i &gt; VLEN) {
					diag(l, &#34;bad Vop&#34;);
					return;
				}
				j = i - 1;
				goto op;

			case Vspazz:	// nasty hack to save a reg in SUB
//print(&#34;spazz\n&#34;);
				if(g) {
//print(&#34;hi %R lo %R t %R\n&#34;, r-&gt;right-&gt;reg, r-&gt;left-&gt;reg, tmps[0].reg);
					ot = r-&gt;right;
					r-&gt;right = r-&gt;left;
					tl = new(0, Z, Z);
					*tl = tmps[0];
					r-&gt;left = tl;
					tmps[0] = *ot;
//print(&#34;hi %R lo %R t %R\n&#34;, r-&gt;right-&gt;reg, r-&gt;left-&gt;reg, tmps[0].reg);
				}
				i++;
				break;

			case Vmv:
			case Vmul:
			case Vshll:
				i += 3;
				if(i &gt; VLEN) {
					diag(l, &#34;bad Vop&#34;);
					return;
				}
				j = i - 2;
				goto op;

			case Vins0:
				i += 2;
				if(i &gt; VLEN) {
					diag(l, &#34;bad Vop&#34;);
					return;
				}
				gins(c[i - 1], Z, Z);
				break;

			case Vop:
			case Vopx:
			case Vins:
			case Vinsl:
			case Vinsr:
			case Vinsla:
			case Vinsra:
			case Vinsx:
				i += 4;
				if(i &gt; VLEN) {
					diag(l, &#34;bad Vop&#34;);
					return;
				}
				j = i - 2;
				goto op;

			op:
				if(!g)
					break;
				tl = Z;
				tr = Z;
				for(; j &lt; i; j++) {
					switch(c[j]) {
					case C00:
						ot = nodconst(0);
						break;
					case C01:
						ot = nodconst(1);
						break;
					case C31:
						ot = nodconst(31);
						break;
					case C32:
						ot = nodconst(32);
						break;

					case O_l:
					case O_l_lo:
						ot = l; xp = &amp;lo; xo = 0;
						goto op0;
					case O_l_hi:
						ot = l; xp = &amp;lo; xo = SZ_LONG;
						goto op0;
					case O_r:
					case O_r_lo:
						ot = r; xp = &amp;ro; xo = 0;
						goto op0;
					case O_r_hi:
						ot = r; xp = &amp;ro; xo = SZ_LONG;
						goto op0;
					case O_t_lo:
						ot = t; xp = &amp;to; xo = 0;
						goto op0;
					case O_t_hi:
						ot = t; xp = &amp;to; xo = SZ_LONG;
						goto op0;
					case O_l_rp:
						ot = l;
						break;
					case O_r_rp:
						ot = r;
						break;
					case O_t_rp:
						ot = t;
						break;
					case O_r0:
					case O_r1:
						ot = &amp;tmps[c[j] - O_r0];
						break;
					case O_Zop:
						ot = Z;
						break;

					op0:
						switch(ot-&gt;op) {
						case OCONST:
							if(xo)
								ot = hi64(ot);
							else
								ot = lo64(ot);
							break;
						case OREGPAIR:
							if(xo)
								ot = ot-&gt;right;
							else
								ot = ot-&gt;left;
							break;
						case OREGISTER:
							break;
						default:
							if(xo != *xp) {
								ot-&gt;xoffset += xo - *xp;
								*xp = xo;
							}
						}
						break;
					
					default:
						diag(l, &#34;bad V_lop&#34;);
						return;
					}
					if(tl == nil)
						tl = ot;
					else
						tr = ot;
				}
				if(op == Vzero) {
					zeroregm(tl);
					break;
				}
				oc = c[i - 3];
				if(op == Vinsx || op == Vopx) {
//print(&#34;%d -&gt; %d\n&#34;, oc, args[oc]);
					oc = args[oc];
				}
				else {
					switch(oc) {
					case O_a0:
					case O_a1:
						oc = args[oc - O_a0];
						break;
					}
				}
				switch(op) {
				case Vmul:
					mulgen(tr-&gt;type, tl, tr);
					break;
				case Vmv:
					gmove(tl, tr);
					break;
				case Vshll:
					shiftit(tr-&gt;type, tl, tr);
					break;
				case Vop:
				case Vopx:
					gopcode(oc, types[TULONG], tl, tr);
					break;
				case Vins:
				case Vinsx:
					gins(oc, tl, tr);
					break;
				case Vinsl:
					gins(oc, tl, tr-&gt;right);
					p-&gt;from.index = tr-&gt;left-&gt;reg;
					break;
				case Vinsr:
					gins(oc, tl, tr-&gt;left);
					p-&gt;from.index = tr-&gt;right-&gt;reg;
					break;
				case Vinsla:
					gins(oc, tl, tr + 1);
					p-&gt;from.index = tr-&gt;reg;
					break;
				case Vinsra:
					gins(oc, tl, tr);
					p-&gt;from.index = (tr + 1)-&gt;reg;
					break;
				}
				break;

			case VT:
				g = true;
				i++;
				break;
			case VF:
				g = !true;
				i++;
				break;

			case V_T0: case V_T1:
				g = args[op - V_T0];
				i++;
				break;

			case V_F0: case V_F1:
				g = !args[op - V_F0];
				i++;
				break;

			case V_C0: case V_C1:
				if(g)
					args[op - V_C0] = 0;
				i++;
				break;

			case V_S0: case V_S1:
				if(g)
					args[op - V_S0] = 1;
				i++;
				break;

			case V_l_lo_f:
				g = lo64v(l) == 0;
				i++;
				break;
			case V_l_hi_f:
				g = hi64v(l) == 0;
				i++;
				break;
			case V_l_lo_t:
				g = lo64v(l) != 0;
				i++;
				break;
			case V_l_hi_t:
				g = hi64v(l) != 0;
				i++;
				break;
			case V_l_lo_u:
				g = lo64v(l) &gt;= 0;
				i++;
				break;
			case V_l_hi_u:
				g = hi64v(l) &gt;= 0;
				i++;
				break;
			case V_r_lo_f:
				g = lo64v(r) == 0;
				i++;
				break;
			case V_r_hi_f:
				g = hi64v(r) == 0;
				i++;
				break;
			case V_r_lo_t:
				g = lo64v(r) != 0;
				i++;
				break;
			case V_r_hi_t:
				g = hi64v(r) != 0;
				i++;
				break;
			case V_r_lo_u:
				g = lo64v(r) &gt;= 0;
				i++;
				break;
			case V_r_hi_u:
				g = hi64v(r) &gt;= 0;
				i++;
				break;

			case Vend:
				goto out;

			case V_a0: case V_a1:
				if(g) {
					lt = l-&gt;type;
					l-&gt;type = types[TULONG];
					regalloc(&amp;tmps[op - V_a0], l, Z);
					l-&gt;type = lt;
				}
				i++;
				break;

			case V_f0: case V_f1:
				if(g)
					regfree(&amp;tmps[op - V_f0]);
				i++;
				break;

			case V_p0: case V_p1: case V_p2: case V_p3: case V_p4:
				if(g)
					patch(pr[op - V_p0], pc);
				i++;
				break;

			case V_s0: case V_s1: case V_s2: case V_s3: case V_s4:
				if(g)
					pr[op - V_s0] = p;
				i++;
				break;

			default:
				diag(l, &#34;bad biggen: %d&#34;, op);
				return;
			}
			if(i == VLEN || c[i] == 0)
				break;
		}
	}
out:
	if(lo)
		l-&gt;xoffset -= lo;
	if(ro)
		r-&gt;xoffset -= ro;
	if(to)
		t-&gt;xoffset -= to;
}

int
cgen64(Node *n, Node *nn)
{
	Type *dt;
	uchar *args, (*cp)[VLEN], (**optab)[VLEN];
	int li, ri, lri, dr, si, m, op, sh, cmp, true;
	Node *c, *d, *l, *r, *t, *s, nod1, nod2, nod3, nod4, nod5;

	if(debug[&#39;g&#39;]) {
		prtree(nn, &#34;cgen64 lhs&#34;);
		prtree(n, &#34;cgen64&#34;);
		print(&#34;AX = %d\n&#34;, reg[D_AX]);
	}
	cmp = 0;
	sh = 0;

	switch(n-&gt;op) {
	case ONEG:
		d = regpair(nn, n);
		sugen(n-&gt;left, d, 8);
		gins(ANOTL, Z, d-&gt;right);
		gins(ANEGL, Z, d-&gt;left);
		gins(ASBBL, nodconst(-1), d-&gt;right);
		break;

	case OCOM:
		if(!vaddr(n-&gt;left, 0) || !vaddr(nn, 0))
			d = regpair(nn, n);
		else
			return 0;
		sugen(n-&gt;left, d, 8);
		gins(ANOTL, Z, d-&gt;left);
		gins(ANOTL, Z, d-&gt;right);
		break;

	case OADD:
		optab = ADDtab;
		args = ADDargs;
		goto twoop;
	case OAND:
		optab = ANDtab;
		args = ANDargs;
		goto twoop;
	case OOR:
		optab = ORtab;
		args = ORargs;
		goto twoop;
	case OSUB:
		optab = SUBtab;
		args = SUBargs;
		goto twoop;
	case OXOR:
		optab = ORtab;
		args = XORargs;
		goto twoop;
	case OASHL:
		sh = 1;
		args = nil;
		optab = shlltab;
		goto twoop;
	case OLSHR:
		sh = 1;
		args = shrlargs;
		optab = shrltab;
		goto twoop;
	case OASHR:
		sh = 1;
		args = sarlargs;
		optab = shrltab;
		goto twoop;
	case OEQ:
		cmp = 1;
		args = nil;
		optab = nil;
		goto twoop;
	case ONE:
		cmp = 1;
		args = nil;
		optab = nil;
		goto twoop;
	case OLE:
		cmp = 1;
		args = nil;
		optab = nil;
		goto twoop;
	case OLT:
		cmp = 1;
		args = nil;
		optab = nil;
		goto twoop;
	case OGE:
		cmp = 1;
		args = nil;
		optab = nil;
		goto twoop;
	case OGT:
		cmp = 1;
		args = nil;
		optab = nil;
		goto twoop;
	case OHI:
		cmp = 1;
		args = nil;
		optab = nil;
		goto twoop;
	case OHS:
		cmp = 1;
		args = nil;
		optab = nil;
		goto twoop;
	case OLO:
		cmp = 1;
		args = nil;
		optab = nil;
		goto twoop;
	case OLS:
		cmp = 1;
		args = nil;
		optab = nil;
		goto twoop;

twoop:
		dr = nn != Z &amp;&amp; nn-&gt;op == OREGPAIR;
		l = vfunc(n-&gt;left, nn);
		if(sh)
			r = n-&gt;right;
		else
			r = vfunc(n-&gt;right, nn);

		li = l-&gt;op == ONAME || l-&gt;op == OINDREG || l-&gt;op == OCONST;
		ri = r-&gt;op == ONAME || r-&gt;op == OINDREG || r-&gt;op == OCONST;

#define	IMM(l, r)	((l) | ((r) &lt;&lt; 1))

		lri = IMM(li, ri);

		/* find out what is so easy about some operands */
		if(li)
			li = whatof(l, sh | cmp);
		if(ri)
			ri = whatof(r, cmp);

		if(sh)
			goto shift;

		if(cmp)
			goto cmp;

		/* evaluate hard subexps, stealing nn if possible. */
		switch(lri) {
		case IMM(0, 0):
		bin00:
			if(l-&gt;complex &gt; r-&gt;complex) {
				if(dr)
					t = nn;
				else
					t = regpair(Z, n);
				sugen(l, t, 8);
				l = t;
				t = regpair(Z, n);
				sugen(r, t, 8);
				r = t;
			}
			else {
				t = regpair(Z, n);
				sugen(r, t, 8);
				r = t;
				if(dr)
					t = nn;
				else
					t = regpair(Z, n);
				sugen(l, t, 8);
				l = t;
			}
			break;
		case IMM(0, 1):
			if(dr)
				t = nn;
			else
				t = regpair(Z, n);
			sugen(l, t, 8);
			l = t;
			break;
		case IMM(1, 0):
			if(n-&gt;op == OSUB &amp;&amp; l-&gt;op == OCONST &amp;&amp; hi64v(l) == 0) {
				lri = IMM(0, 0);
				goto bin00;
			}
			if(dr)
				t = nn;
			else
				t = regpair(Z, n);
			sugen(r, t, 8);
			r = t;
			break;
		case IMM(1, 1):
			break;
		}

#define	WW(l, r)	((l) | ((r) &lt;&lt; 2))
		d = Z;
		dt = nn-&gt;type;
		nn-&gt;type = types[TLONG];

		switch(lri) {
		case IMM(0, 0):
			biggen(l, r, Z, 0, binop00, args);
			break;
		case IMM(0, 1):
			switch(ri) {
			case WNONE:
				diag(r, &#34;bad whatof\n&#34;);
				break;
			case WCONST:
				biggen(l, r, Z, 0, optab[B0c], args);
				break;
			case WHARD:
				reglcgen(&amp;nod2, r, Z);
				r = &amp;nod2;
				/* fall thru */
			case WADDR:
				biggen(l, r, Z, 0, binoptmp, args);
				if(ri == WHARD)
					regfree(r);
				break;
			}
			break;
		case IMM(1, 0):
			if(n-&gt;op == OSUB) {
				switch(li) {
				case WNONE:
					diag(l, &#34;bad whatof\n&#34;);
					break;
				case WHARD:
					reglcgen(&amp;nod2, l, Z);
					l = &amp;nod2;
					/* fall thru */
				case WADDR:
				case WCONST:
					biggen(l, r, Z, 0, sub10, args);
					break;
				}
				if(li == WHARD)
					regfree(l);
			}
			else {
				switch(li) {
				case WNONE:
					diag(l, &#34;bad whatof\n&#34;);
					break;
				case WCONST:
					biggen(r, l, Z, 0, optab[B0c], args);
					break;
				case WHARD:
					reglcgen(&amp;nod2, l, Z);
					l = &amp;nod2;
					/* fall thru */
				case WADDR:
					biggen(r, l, Z, 0, binoptmp, args);
					if(li == WHARD)
						regfree(l);
					break;
				}
			}
			break;
		case IMM(1, 1):
			switch(WW(li, ri)) {
			case WW(WCONST, WHARD):
				if(r-&gt;op == ONAME &amp;&amp; n-&gt;op == OAND &amp;&amp; reduxv(l))
					ri = WADDR;
				break;
			case WW(WHARD, WCONST):
				if(l-&gt;op == ONAME &amp;&amp; n-&gt;op == OAND &amp;&amp; reduxv(r))
					li = WADDR;
				break;
			}
			if(li == WHARD) {
				reglcgen(&amp;nod3, l, Z);
				l = &amp;nod3;
			}
			if(ri == WHARD) {
				reglcgen(&amp;nod2, r, Z);
				r = &amp;nod2;
			}
			d = regpair(nn, n);
			instpair(d, Z);
			switch(WW(li, ri)) {
			case WW(WCONST, WADDR):
			case WW(WCONST, WHARD):
				biggen(l, r, d, 0, optab[Bca], args);
				break;

			case WW(WADDR, WCONST):
			case WW(WHARD, WCONST):
				biggen(l, r, d, 0, optab[Bac], args);
				break;

			case WW(WADDR, WADDR):
			case WW(WADDR, WHARD):
			case WW(WHARD, WADDR):
			case WW(WHARD, WHARD):
				biggen(l, r, d, 0, binop11, args);
				break;

			default:
				diag(r, &#34;bad whatof pair %d %d\n&#34;, li, ri);
				break;
			}
			if(li == WHARD)
				regfree(l);
			if(ri == WHARD)
				regfree(r);
			break;
		}

		nn-&gt;type = dt;

		if(d != Z)
			goto finished;

		switch(lri) {
		case IMM(0, 0):
			freepair(r);
			/* fall thru */;
		case IMM(0, 1):
			if(!dr)
				storepair(l, nn, 1);
			break;
		case IMM(1, 0):
			if(!dr)
				storepair(r, nn, 1);
			break;
		case IMM(1, 1):
			break;
		}
		return 1;

	shift:
		c = Z;

		/* evaluate hard subexps, stealing nn if possible. */
		/* must also secure CX.  not as many optims as binop. */
		switch(lri) {
		case IMM(0, 0):
		imm00:
			if(l-&gt;complex + 1 &gt; r-&gt;complex) {
				if(dr)
					t = nn;
				else
					t = regpair(Z, l);
				sugen(l, t, 8);
				l = t;
				t = &amp;nod1;
				c = snarfreg(l, t, D_CX, r, &amp;nod2);
				cgen(r, t);
				r = t;
			}
			else {
				t = &amp;nod1;
				c = snarfreg(nn, t, D_CX, r, &amp;nod2);
				cgen(r, t);
				r = t;
				if(dr)
					t = nn;
				else
					t = regpair(Z, l);
				sugen(l, t, 8);
				l = t;
			}
			break;
		case IMM(0, 1):
		imm01:
			if(ri != WCONST) {
				lri = IMM(0, 0);
				goto imm00;
			}
			if(dr)
				t = nn;
			else
				t = regpair(Z, n);
			sugen(l, t, 8);
			l = t;
			break;
		case IMM(1, 0):
		imm10:
			if(li != WCONST) {
				lri = IMM(0, 0);
				goto imm00;
			}
			t = &amp;nod1;
			c = snarfreg(nn, t, D_CX, r, &amp;nod2);
			cgen(r, t);
			r = t;
			break;
		case IMM(1, 1):
			if(ri != WCONST) {
				lri = IMM(1, 0);
				goto imm10;
			}
			if(li == WHARD) {
				lri = IMM(0, 1);
				goto imm01;
			}
			break;
		}

		d = Z;

		switch(lri) {
		case IMM(0, 0):
			biggen(l, r, Z, 0, optab[S00], args);
			break;
		case IMM(0, 1):
			switch(ri) {
			case WNONE:
			case WADDR:
			case WHARD:
				diag(r, &#34;bad whatof\n&#34;);
				break;
			case WCONST:
				m = r-&gt;vconst &amp; 63;
				s = nodconst(m);
				if(m &lt; 32)
					cp = optab[Sc0];
				else if(m == 32)
					cp = optab[Sc1];
				else
					cp = optab[Sc2];
				biggen(l, s, Z, 0, cp, args);
				break;
			}
			break;
		case IMM(1, 0):
			/* left is const */
			d = regpair(nn, n);
			instpair(d, Z);
			biggen(l, r, d, 0, optab[S10], args);
			regfree(r);
			break;
		case IMM(1, 1):
			d = regpair(nn, n);
			instpair(d, Z);
			switch(WW(li, ri)) {
			case WW(WADDR, WCONST):
				m = r-&gt;vconst &amp; 63;
				s = nodconst(m);
				if(m &lt; 32) {
					loadpair(l, d);
					l = d;
					cp = optab[Sc0];
				}
				else if(m == 32)
					cp = optab[Sac3];
				else
					cp = optab[Sac4];
				biggen(l, s, d, 0, cp, args);
				break;

			default:
				diag(r, &#34;bad whatof pair %d %d\n&#34;, li, ri);
				break;
			}
			break;
		}

		if(c != Z) {
			gins(AMOVL, c, r);
			regfree(c);
		}

		if(d != Z)
			goto finished;

		switch(lri) {
		case IMM(0, 0):
			regfree(r);
			/* fall thru */
		case IMM(0, 1):
			if(!dr)
				storepair(l, nn, 1);
			break;
		case IMM(1, 0):
			regfree(r);
			break;
		case IMM(1, 1):
			break;
		}
		return 1;

	cmp:
		op = n-&gt;op;
		/* evaluate hard subexps */
		switch(lri) {
		case IMM(0, 0):
			if(l-&gt;complex &gt; r-&gt;complex) {
				t = regpair(Z, l);
				sugen(l, t, 8);
				l = t;
				t = regpair(Z, r);
				sugen(r, t, 8);
				r = t;
			}
			else {
				t = regpair(Z, r);
				sugen(r, t, 8);
				r = t;
				t = regpair(Z, l);
				sugen(l, t, 8);
				l = t;
			}
			break;
		case IMM(1, 0):
			t = r;
			r = l;
			l = t;
			ri = li;
			op = invrel[relindex(op)];
			/* fall thru */
		case IMM(0, 1):
			t = regpair(Z, l);
			sugen(l, t, 8);
			l = t;
			break;
		case IMM(1, 1):
			break;
		}

		true = 1;
		optab = cmptab;
		switch(op) {
		case OEQ:
			optab = NEtab;
			true = 0;
			break;
		case ONE:
			optab = NEtab;
			break;
		case OLE:
			args = GTargs;
			true = 0;
			break;
		case OGT:
			args = GTargs;
			break;
		case OLS:
			args = HIargs;
			true = 0;
			break;
		case OHI:
			args = HIargs;
			break;
		case OLT:
			args = GEargs;
			true = 0;
			break;
		case OGE:
			args = GEargs;
			break;
		case OLO:
			args = HSargs;
			true = 0;
			break;
		case OHS:
			args = HSargs;
			break;
		default:
			diag(n, &#34;bad cmp\n&#34;);
			SET(optab);
		}

		switch(lri) {
		case IMM(0, 0):
			biggen(l, r, Z, true, optab[T0i], args);
			break;
		case IMM(0, 1):
		case IMM(1, 0):
			switch(ri) {
			case WNONE:
				diag(l, &#34;bad whatof\n&#34;);
				break;
			case WCONST:
				biggen(l, r, Z, true, optab[T0i], args);
				break;
			case WHARD:
				reglcgen(&amp;nod2, r, Z);
				r = &amp;nod2;
				/* fall thru */
			case WADDR:
				biggen(l, r, Z, true, optab[T0i], args);
				if(ri == WHARD)
					regfree(r);
				break;
			}
			break;
		case IMM(1, 1):
			if(li == WHARD) {
				reglcgen(&amp;nod3, l, Z);
				l = &amp;nod3;
			}
			if(ri == WHARD) {
				reglcgen(&amp;nod2, r, Z);
				r = &amp;nod2;
			}
			biggen(l, r, Z, true, optab[Tii], args);
			if(li == WHARD)
				regfree(l);
			if(ri == WHARD)
				regfree(r);
			break;
		}

		switch(lri) {
		case IMM(0, 0):
			freepair(r);
			/* fall thru */;
		case IMM(0, 1):
		case IMM(1, 0):
			freepair(l);
			break;
		case IMM(1, 1):
			break;
		}
		return 1;

	case OASMUL:
	case OASLMUL:
		m = 0;
		goto mulop;

	case OMUL:
	case OLMUL:
		m = 1;
		goto mulop;

	mulop:
		dr = nn != Z &amp;&amp; nn-&gt;op == OREGPAIR;
		l = vfunc(n-&gt;left, nn);
		r = vfunc(n-&gt;right, nn);
		if(r-&gt;op != OCONST) {
			if(l-&gt;complex &gt; r-&gt;complex) {
				if(m) {
					t = l;
					l = r;
					r = t;
				}
				else if(!vaddr(l, 1)) {
					reglcgen(&amp;nod5, l, Z);
					l = &amp;nod5;
					evacaxdx(l);
				}
			}
			t = regpair(Z, n);
			sugen(r, t, 8);
			r = t;
			evacaxdx(r-&gt;left);
			evacaxdx(r-&gt;right);
			if(l-&gt;complex &lt;= r-&gt;complex &amp;&amp; !m &amp;&amp; !vaddr(l, 1)) {
				reglcgen(&amp;nod5, l, Z);
				l = &amp;nod5;
				evacaxdx(l);
			}
		}
		if(dr)
			t = nn;
		else
			t = regpair(Z, n);
		c = Z;
		d = Z;
		if(!nodreg(&amp;nod1, t-&gt;left, D_AX)) {
			if(t-&gt;left-&gt;reg != D_AX){
				t-&gt;left-&gt;reg = D_AX;
				reg[D_AX]++;
			}else if(reg[D_AX] == 0)
				fatal(Z, &#34;vlong mul AX botch&#34;);
		}
		if(!nodreg(&amp;nod2, t-&gt;right, D_DX)) {
			if(t-&gt;right-&gt;reg != D_DX){
				t-&gt;right-&gt;reg = D_DX;
				reg[D_DX]++;
			}else if(reg[D_DX] == 0)
				fatal(Z, &#34;vlong mul DX botch&#34;);
		}
		if(m)
			sugen(l, t, 8);
		else
			loadpair(l, t);
		if(t-&gt;left-&gt;reg != D_AX) {
			c = &amp;nod3;
			regsalloc(c, t-&gt;left);
			gmove(&amp;nod1, c);
			gmove(t-&gt;left, &amp;nod1);
			zapreg(t-&gt;left);
		}
		if(t-&gt;right-&gt;reg != D_DX) {
			d = &amp;nod4;
			regsalloc(d, t-&gt;right);
			gmove(&amp;nod2, d);
			gmove(t-&gt;right, &amp;nod2);
			zapreg(t-&gt;right);
		}
		if(c != Z || d != Z) {
			s = regpair(Z, n);
			s-&gt;left = &amp;nod1;
			s-&gt;right = &amp;nod2;
		}
		else
			s = t;
		if(r-&gt;op == OCONST) {
			if(hi64v(r) == 0)
				biggen(s, r, Z, 0, mulc32, nil);
			else
				biggen(s, r, Z, 0, mulc64, nil);
		}
		else
			biggen(s, r, Z, 0, mull, nil);
		instpair(t, Z);
		if(c != Z) {
			gmove(&amp;nod1, t-&gt;left);
			gmove(&amp;nod3, &amp;nod1);
		}
		if(d != Z) {
			gmove(&amp;nod2, t-&gt;right);
			gmove(&amp;nod4, &amp;nod2);
		}
		if(r-&gt;op == OREGPAIR)
			freepair(r);
		if(!m)
			storepair(t, l, 0);
		if(l == &amp;nod5)
			regfree(l);
		if(!dr) {
			if(nn != Z)
				storepair(t, nn, 1);
			else
				freepair(t);
		}
		return 1;

	case OASADD:
		args = ADDargs;
		goto vasop;
	case OASAND:
		args = ANDargs;
		goto vasop;
	case OASOR:
		args = ORargs;
		goto vasop;
	case OASSUB:
		args = SUBargs;
		goto vasop;
	case OASXOR:
		args = XORargs;
		goto vasop;

	vasop:
		l = n-&gt;left;
		r = n-&gt;right;
		dr = nn != Z &amp;&amp; nn-&gt;op == OREGPAIR;
		m = 0;
		if(l-&gt;complex &gt; r-&gt;complex) {
			if(!vaddr(l, 1)) {
				reglcgen(&amp;nod1, l, Z);
				l = &amp;nod1;
			}
			if(!vaddr(r, 1) || nn != Z || r-&gt;op == OCONST) {
				if(dr)
					t = nn;
				else
					t = regpair(Z, r);
				sugen(r, t, 8);
				r = t;
				m = 1;
			}
		}
		else {
			if(!vaddr(r, 1) || nn != Z || r-&gt;op == OCONST) {
				if(dr)
					t = nn;
				else
					t = regpair(Z, r);
				sugen(r, t, 8);
				r = t;
				m = 1;
			}
			if(!vaddr(l, 1)) {
				reglcgen(&amp;nod1, l, Z);
				l = &amp;nod1;
			}
		}
		if(nn != Z) {
			if(n-&gt;op == OASSUB)
				biggen(l, r, Z, 0, sub10, args);
			else
				biggen(r, l, Z, 0, binoptmp, args);
			storepair(r, l, 0);
		}
		else {
			if(m)
				biggen(l, r, Z, 0, binop00, args);
			else
				biggen(l, r, Z, 0, binoptmp, args);
		}
		if(l == &amp;nod1)
			regfree(&amp;nod1);
		if(m) {
			if(nn == Z)
				freepair(r);
			else if(!dr)
				storepair(r, nn, 1);
		}
		return 1;

	case OASASHL:
		args = nil;
		optab = asshlltab;
		goto assh;
	case OASLSHR:
		args = shrlargs;
		optab = asshrltab;
		goto assh;
	case OASASHR:
		args = sarlargs;
		optab = asshrltab;
		goto assh;

	assh:
		c = Z;
		l = n-&gt;left;
		r = n-&gt;right;
		if(r-&gt;op == OCONST) {
			m = r-&gt;vconst &amp; 63;
			if(m &lt; 32)
				m = SAclo;
			else if(m == 32)
				m = SAc32;
			else
				m = SAchi;
		}
		else
			m = SAgen;
		if(l-&gt;complex &gt; r-&gt;complex) {
			if(!vaddr(l, 0)) {
				reglcgen(&amp;nod1, l, Z);
				l = &amp;nod1;
			}
			if(m == SAgen) {
				t = &amp;nod2;
				if(l-&gt;reg == D_CX) {
					regalloc(t, r, Z);
					gmove(l, t);
					l-&gt;reg = t-&gt;reg;
					t-&gt;reg = D_CX;
				}
				else
					c = snarfreg(nn, t, D_CX, r, &amp;nod3);
				cgen(r, t);
				r = t;
			}
		}
		else {
			if(m == SAgen) {
				t = &amp;nod2;
				c = snarfreg(nn, t, D_CX, r, &amp;nod3);
				cgen(r, t);
				r = t;
			}
			if(!vaddr(l, 0)) {
				reglcgen(&amp;nod1, l, Z);
				l = &amp;nod1;
			}
		}

		if(nn != Z) {
			m += SAdgen - SAgen;
			d = regpair(nn, n);
			instpair(d, Z);
			biggen(l, r, d, 0, optab[m], args);
			if(l == &amp;nod1) {
				regfree(&amp;nod1);
				l = Z;
			}
			if(r == &amp;nod2 &amp;&amp; c == Z) {
				regfree(&amp;nod2);
				r = Z;
			}
			if(d != nn)
				storepair(d, nn, 1);
		}
		else
			biggen(l, r, Z, 0, optab[m], args);

		if(c != Z) {
			gins(AMOVL, c, r);
			regfree(c);
		}
		if(l == &amp;nod1)
			regfree(&amp;nod1);
		if(r == &amp;nod2)
			regfree(&amp;nod2);
		return 1;

	case OPOSTINC:
		args = ADDargs;
		cp = incdecpost;
		goto vinc;
	case OPOSTDEC:
		args = SUBargs;
		cp = incdecpost;
		goto vinc;
	case OPREINC:
		args = ADDargs;
		cp = incdecpre;
		goto vinc;
	case OPREDEC:
		args = SUBargs;
		cp = incdecpre;
		goto vinc;

	vinc:
		l = n-&gt;left;
		if(!vaddr(l, 1)) {
			reglcgen(&amp;nod1, l, Z);
			l = &amp;nod1;
		}
		
		if(nn != Z) {
			d = regpair(nn, n);
			instpair(d, Z);
			biggen(l, Z, d, 0, cp, args);
			if(l == &amp;nod1) {
				regfree(&amp;nod1);
				l = Z;
			}
			if(d != nn)
				storepair(d, nn, 1);
		}
		else
			biggen(l, Z, Z, 0, incdec, args);

		if(l == &amp;nod1)
			regfree(&amp;nod1);
		return 1;

	case OCAST:
		l = n-&gt;left;
		if(typev[l-&gt;type-&gt;etype]) {
			if(!vaddr(l, 1)) {
				if(l-&gt;complex + 1 &gt; nn-&gt;complex) {
					d = regpair(Z, l);
					sugen(l, d, 8);
					if(!vaddr(nn, 1)) {
						reglcgen(&amp;nod1, nn, Z);
						r = &amp;nod1;
					}
					else
						r = nn;
				}
				else {
					if(!vaddr(nn, 1)) {
						reglcgen(&amp;nod1, nn, Z);
						r = &amp;nod1;
					}
					else
						r = nn;
					d = regpair(Z, l);
					sugen(l, d, 8);
				}
//				d-&gt;left-&gt;type = r-&gt;type;
				d-&gt;left-&gt;type = types[TLONG];
				gmove(d-&gt;left, r);
				freepair(d);
			}
			else {
				if(nn-&gt;op != OREGISTER &amp;&amp; !vaddr(nn, 1)) {
					reglcgen(&amp;nod1, nn, Z);
					r = &amp;nod1;
				}
				else
					r = nn;
//				l-&gt;type = r-&gt;type;
				l-&gt;type = types[TLONG];
				gmove(l, r);
			}
			if(r != nn)
				regfree(r);
		}
		else {
			if(typeu[l-&gt;type-&gt;etype] || cond(l-&gt;op))
				si = TUNSIGNED;
			else
				si = TSIGNED;
			regalloc(&amp;nod1, l, Z);
			cgen(l, &amp;nod1);
			if(nn-&gt;op == OREGPAIR) {
				m = instpair(nn, &amp;nod1);
				biggen(&amp;nod1, Z, nn, si == TSIGNED, castrp, nil);
			}
			else {
				m = 0;
				if(!vaddr(nn, si != TSIGNED)) {
					dt = nn-&gt;type;
					nn-&gt;type = types[TLONG];
					reglcgen(&amp;nod2, nn, Z);
					nn-&gt;type = dt;
					nn = &amp;nod2;
				}
				dt = nn-&gt;type;
				nn-&gt;type = types[TLONG];
				biggen(&amp;nod1, Z, nn, si == TSIGNED, castrpa, nil);
				nn-&gt;type = dt;
				if(nn == &amp;nod2)
					regfree(&amp;nod2);
			}
			if(!m)
				regfree(&amp;nod1);
		}
		return 1;

	default:
		if(n-&gt;op == OREGPAIR) {
			storepair(n, nn, 1);
			return 1;
		}
		if(nn-&gt;op == OREGPAIR) {
			loadpair(n, nn);
			return 1;
		}
		return 0;
	}
finished:
	if(d != nn)
		storepair(d, nn, 1);
	return 1;
}

void
testv(Node *n, int true)
{
	Type *t;
	Node *nn, nod;

	switch(n-&gt;op) {
	case OINDREG:
	case ONAME:
		biggen(n, Z, Z, true, testi, nil);
		break;

	default:
		n = vfunc(n, n);
		if(n-&gt;addable &gt;= INDEXED) {
			t = n-&gt;type;
			n-&gt;type = types[TLONG];
			reglcgen(&amp;nod, n, Z);
			n-&gt;type = t;
			n = &amp;nod;
			biggen(n, Z, Z, true, testi, nil);
			if(n == &amp;nod)
				regfree(n);
		}
		else {
			nn = regpair(Z, n);
			sugen(n, nn, 8);
			biggen(nn, Z, Z, true, testi, nil);
			freepair(nn);
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
