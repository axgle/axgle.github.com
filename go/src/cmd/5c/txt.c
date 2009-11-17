<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5c/txt.c</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5c/txt.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5c/txt.c
// http://code.google.com/p/inferno-os/source/browse/utils/5c/txt.c
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
ginit(void)
{
	Type *t;

	thechar = &#39;5&#39;;
	thestring = &#34;arm&#34;;
	exregoffset = REGEXT;
	exfregoffset = FREGEXT;
	listinit();
	nstring = 0;
	mnstring = 0;
	nrathole = 0;
	pc = 0;
	breakpc = -1;
	continpc = -1;
	cases = C;
	firstp = P;
	lastp = P;
	tfield = types[TLONG];

	zprog.link = P;
	zprog.as = AGOK;
	zprog.reg = NREG;
	zprog.from.type = D_NONE;
	zprog.from.name = D_NONE;
	zprog.from.reg = NREG;
	zprog.to = zprog.from;
	zprog.scond = 0xE;

	regnode.op = OREGISTER;
	regnode.class = CEXREG;
	regnode.reg = REGTMP;
	regnode.complex = 0;
	regnode.addable = 11;
	regnode.type = types[TLONG];

	constnode.op = OCONST;
	constnode.class = CXXX;
	constnode.complex = 0;
	constnode.addable = 20;
	constnode.type = types[TLONG];

	fconstnode.op = OCONST;
	fconstnode.class = CXXX;
	fconstnode.complex = 0;
	fconstnode.addable = 20;
	fconstnode.type = types[TDOUBLE];

	nodsafe = new(ONAME, Z, Z);
	nodsafe-&gt;sym = slookup(&#34;.safe&#34;);
	nodsafe-&gt;type = types[TINT];
	nodsafe-&gt;etype = types[TINT]-&gt;etype;
	nodsafe-&gt;class = CAUTO;
	complex(nodsafe);

	t = typ(TARRAY, types[TCHAR]);
	symrathole = slookup(&#34;.rathole&#34;);
	symrathole-&gt;class = CGLOBL;
	symrathole-&gt;type = t;

	nodrat = new(ONAME, Z, Z);
	nodrat-&gt;sym = symrathole;
	nodrat-&gt;type = types[TIND];
	nodrat-&gt;etype = TVOID;
	nodrat-&gt;class = CGLOBL;
	complex(nodrat);
	nodrat-&gt;type = t;

	nodret = new(ONAME, Z, Z);
	nodret-&gt;sym = slookup(&#34;.ret&#34;);
	nodret-&gt;type = types[TIND];
	nodret-&gt;etype = TIND;
	nodret-&gt;class = CPARAM;
	nodret = new(OIND, nodret, Z);
	complex(nodret);

	com64init();

	memset(reg, 0, sizeof(reg));
}

void
gclean(void)
{
	int i;
	Sym *s;

	for(i=0; i&lt;NREG; i++)
		if(reg[i])
			diag(Z, &#34;reg %d left allocated&#34;, i);
	for(i=NREG; i&lt;NREG+NFREG; i++)
		if(reg[i])
			diag(Z, &#34;freg %d left allocated&#34;, i-NREG);
	while(mnstring)
		outstring(&#34;&#34;, 1L);
	symstring-&gt;type-&gt;width = nstring;
	symrathole-&gt;type-&gt;width = nrathole;
	for(i=0; i&lt;NHASH; i++)
	for(s = hash[i]; s != S; s = s-&gt;link) {
		if(s-&gt;type == T)
			continue;
		if(s-&gt;type-&gt;width == 0)
			continue;
		if(s-&gt;class != CGLOBL &amp;&amp; s-&gt;class != CSTATIC)
			continue;
		if(s-&gt;type == types[TENUM])
			continue;
		gpseudo(AGLOBL, s, nodconst(s-&gt;type-&gt;width));
	}
	nextpc();
	p-&gt;as = AEND;
	outcode();
}

void
nextpc(void)
{

	p = alloc(sizeof(*p));
	*p = zprog;
	p-&gt;lineno = nearln;
	pc++;
	if(firstp == P) {
		firstp = p;
		lastp = p;
		return;
	}
	lastp-&gt;link = p;
	lastp = p;
}

void
gargs(Node *n, Node *tn1, Node *tn2)
{
	int32 regs;
	Node fnxargs[20], *fnxp;

	regs = cursafe;

	fnxp = fnxargs;
	garg1(n, tn1, tn2, 0, &amp;fnxp);	/* compile fns to temps */

	curarg = 0;
	fnxp = fnxargs;
	garg1(n, tn1, tn2, 1, &amp;fnxp);	/* compile normal args and temps */

	cursafe = regs;
}

void
garg1(Node *n, Node *tn1, Node *tn2, int f, Node **fnxp)
{
	Node nod;

	if(n == Z)
		return;
	if(n-&gt;op == OLIST) {
		garg1(n-&gt;left, tn1, tn2, f, fnxp);
		garg1(n-&gt;right, tn1, tn2, f, fnxp);
		return;
	}
	if(f == 0) {
		if(n-&gt;complex &gt;= FNX) {
			regsalloc(*fnxp, n);
			nod = znode;
			nod.op = OAS;
			nod.left = *fnxp;
			nod.right = n;
			nod.type = n-&gt;type;
			cgen(&amp;nod, Z);
			(*fnxp)++;
		}
		return;
	}
	if(typesuv[n-&gt;type-&gt;etype]) {
		regaalloc(tn2, n);
		if(n-&gt;complex &gt;= FNX) {
			sugen(*fnxp, tn2, n-&gt;type-&gt;width);
			(*fnxp)++;
		} else
			sugen(n, tn2, n-&gt;type-&gt;width);
		return;
	}
	if(REGARG &gt;= 0 &amp;&amp; curarg == 0 &amp;&amp; typechlp[n-&gt;type-&gt;etype]) {
		regaalloc1(tn1, n);
		if(n-&gt;complex &gt;= FNX) {
			cgen(*fnxp, tn1);
			(*fnxp)++;
		} else
			cgen(n, tn1);
		return;
	}
	regalloc(tn1, n, Z);
	if(n-&gt;complex &gt;= FNX) {
		cgen(*fnxp, tn1);
		(*fnxp)++;
	} else
		cgen(n, tn1);
	regaalloc(tn2, n);
	gopcode(OAS, tn1, Z, tn2);
	regfree(tn1);
}

Node*
nodconst(int32 v)
{
	constnode.vconst = v;
	return &amp;constnode;
}

Node*
nod32const(vlong v)
{
	constnode.vconst = v &amp; MASK(32);
	return &amp;constnode;
}

Node*
nodfconst(double d)
{
	fconstnode.fconst = d;
	return &amp;fconstnode;
}

void
nodreg(Node *n, Node *nn, int reg)
{
	*n = regnode;
	n-&gt;reg = reg;
	n-&gt;type = nn-&gt;type;
	n-&gt;lineno = nn-&gt;lineno;
}

void
regret(Node *n, Node *nn)
{
	int r;

	r = REGRET;
	if(typefd[nn-&gt;type-&gt;etype])
		r = FREGRET+NREG;
	nodreg(n, nn, r);
	reg[r]++;
}

int
tmpreg(void)
{
	int i;

	for(i=REGRET+1; i&lt;NREG; i++)
		if(reg[i] == 0)
			return i;
	diag(Z, &#34;out of fixed registers&#34;);
	return 0;
}

void
regalloc(Node *n, Node *tn, Node *o)
{
	int i, j;
	static int lasti;

	switch(tn-&gt;type-&gt;etype) {
	case TCHAR:
	case TUCHAR:
	case TSHORT:
	case TUSHORT:
	case TINT:
	case TUINT:
	case TLONG:
	case TULONG:
	case TIND:
		if(o != Z &amp;&amp; o-&gt;op == OREGISTER) {
			i = o-&gt;reg;
			if(i &gt;= 0 &amp;&amp; i &lt; NREG)
				goto out;
		}
		j = lasti + REGRET+1;
		for(i=REGRET+1; i&lt;NREG; i++) {
			if(j &gt;= NREG)
				j = REGRET+1;
			if(reg[j] == 0) {
				i = j;
				goto out;
			}
			j++;
		}
		diag(tn, &#34;out of fixed registers&#34;);
		goto err;

	case TFLOAT:
	case TDOUBLE:
	case TVLONG:
		if(o != Z &amp;&amp; o-&gt;op == OREGISTER) {
			i = o-&gt;reg;
			if(i &gt;= NREG &amp;&amp; i &lt; NREG+NFREG)
				goto out;
		}
		j = 0*2 + NREG;
		for(i=NREG; i&lt;NREG+NFREG; i++) {
			if(j &gt;= NREG+NFREG)
				j = NREG;
			if(reg[j] == 0) {
				i = j;
				goto out;
			}
			j++;
		}
		diag(tn, &#34;out of float registers&#34;);
		goto err;
	}
	diag(tn, &#34;unknown type in regalloc: %T&#34;, tn-&gt;type);
err:
	nodreg(n, tn, 0);
	return;
out:
	reg[i]++;
/* 	lasti++;	*** StrongARM does register forwarding */
	if(lasti &gt;= 5)
		lasti = 0;
	nodreg(n, tn, i);
}

void
regialloc(Node *n, Node *tn, Node *o)
{
	Node nod;

	nod = *tn;
	nod.type = types[TIND];
	regalloc(n, &amp;nod, o);
}

void
regfree(Node *n)
{
	int i;

	i = 0;
	if(n-&gt;op != OREGISTER &amp;&amp; n-&gt;op != OINDREG)
		goto err;
	i = n-&gt;reg;
	if(i &lt; 0 || i &gt;= sizeof(reg))
		goto err;
	if(reg[i] &lt;= 0)
		goto err;
	reg[i]--;
	return;
err:
	diag(n, &#34;error in regfree: %d&#34;, i);
}

void
regsalloc(Node *n, Node *nn)
{
	cursafe = align(cursafe, nn-&gt;type, Aaut3);
	maxargsafe = maxround(maxargsafe, cursafe+curarg);
	*n = *nodsafe;
	n-&gt;xoffset = -(stkoff + cursafe);
	n-&gt;type = nn-&gt;type;
	n-&gt;etype = nn-&gt;type-&gt;etype;
	n-&gt;lineno = nn-&gt;lineno;
}

void
regaalloc1(Node *n, Node *nn)
{
	nodreg(n, nn, REGARG);
	reg[REGARG]++;
	curarg = align(curarg, nn-&gt;type, Aarg1);
	curarg = align(curarg, nn-&gt;type, Aarg2);
	maxargsafe = maxround(maxargsafe, cursafe+curarg);
}

void
regaalloc(Node *n, Node *nn)
{
	curarg = align(curarg, nn-&gt;type, Aarg1);
	*n = *nn;
	n-&gt;op = OINDREG;
	n-&gt;reg = REGSP;
	n-&gt;xoffset = curarg + SZ_LONG;
	n-&gt;complex = 0;
	n-&gt;addable = 20;
	curarg = align(curarg, nn-&gt;type, Aarg2);
	maxargsafe = maxround(maxargsafe, cursafe+curarg);
}

void
regind(Node *n, Node *nn)
{

	if(n-&gt;op != OREGISTER) {
		diag(n, &#34;regind not OREGISTER&#34;);
		return;
	}
	n-&gt;op = OINDREG;
	n-&gt;type = nn-&gt;type;
}

void
raddr(Node *n, Prog *p)
{
	Adr a;

	naddr(n, &amp;a);
	if(R0ISZERO &amp;&amp; a.type == D_CONST &amp;&amp; a.offset == 0) {
		a.type = D_REG;
		a.reg = 0;
	}
	if(a.type != D_REG &amp;&amp; a.type != D_FREG) {
		if(n)
			diag(n, &#34;bad in raddr: %O&#34;, n-&gt;op);
		else
			diag(n, &#34;bad in raddr: &lt;null&gt;&#34;);
		p-&gt;reg = NREG;
	} else
		p-&gt;reg = a.reg;
}

void
naddr(Node *n, Adr *a)
{
	int32 v;

	a-&gt;type = D_NONE;
	if(n == Z)
		return;
	switch(n-&gt;op) {
	default:
	bad:
		diag(n, &#34;bad in naddr: %O&#34;, n-&gt;op);
		break;

	case OREGISTER:
		a-&gt;type = D_REG;
		a-&gt;sym = S;
		a-&gt;reg = n-&gt;reg;
		if(a-&gt;reg &gt;= NREG) {
			a-&gt;type = D_FREG;
			a-&gt;reg -= NREG;
		}
		break;

	case OIND:
		naddr(n-&gt;left, a);
		if(a-&gt;type == D_REG) {
			a-&gt;type = D_OREG;
			break;
		}
		if(a-&gt;type == D_CONST) {
			a-&gt;type = D_OREG;
			break;
		}
		goto bad;

	case OINDREG:
		a-&gt;type = D_OREG;
		a-&gt;sym = S;
		a-&gt;offset = n-&gt;xoffset;
		a-&gt;reg = n-&gt;reg;
		break;

	case ONAME:
		a-&gt;etype = n-&gt;etype;
		a-&gt;type = D_OREG;
		a-&gt;name = D_STATIC;
		a-&gt;sym = n-&gt;sym;
		a-&gt;offset = n-&gt;xoffset;
		if(n-&gt;class == CSTATIC)
			break;
		if(n-&gt;class == CEXTERN || n-&gt;class == CGLOBL) {
			a-&gt;name = D_EXTERN;
			break;
		}
		if(n-&gt;class == CAUTO) {
			a-&gt;name = D_AUTO;
			break;
		}
		if(n-&gt;class == CPARAM) {
			a-&gt;name = D_PARAM;
			break;
		}
		goto bad;

	case OCONST:
		a-&gt;sym = S;
		a-&gt;reg = NREG;
		if(typefd[n-&gt;type-&gt;etype]) {
			a-&gt;type = D_FCONST;
			a-&gt;dval = n-&gt;fconst;
		} else {
			a-&gt;type = D_CONST;
			a-&gt;offset = n-&gt;vconst;
		}
		break;

	case OADDR:
		naddr(n-&gt;left, a);
		if(a-&gt;type == D_OREG) {
			a-&gt;type = D_CONST;
			break;
		}
		goto bad;

	case OADD:
		if(n-&gt;left-&gt;op == OCONST) {
			naddr(n-&gt;left, a);
			v = a-&gt;offset;
			naddr(n-&gt;right, a);
		} else {
			naddr(n-&gt;right, a);
			v = a-&gt;offset;
			naddr(n-&gt;left, a);
		}
		a-&gt;offset += v;
		break;

	}
}

void
fop(int as, int f1, int f2, Node *t)
{
	Node nod1, nod2, nod3;

	nodreg(&amp;nod1, t, NREG+f1);
	nodreg(&amp;nod2, t, NREG+f2);
	regalloc(&amp;nod3, t, t);
	gopcode(as, &amp;nod1, &amp;nod2, &amp;nod3);
	gmove(&amp;nod3, t);
	regfree(&amp;nod3);
}

void
gmovm(Node *f, Node *t, int w)
{
	gins(AMOVM, f, t);
	p-&gt;scond |= C_UBIT;
	if(w)
		p-&gt;scond |= C_WBIT;
}

void
gmove(Node *f, Node *t)
{
	int ft, tt, a;
	Node nod;

	ft = f-&gt;type-&gt;etype;
	tt = t-&gt;type-&gt;etype;

	if(ft == TDOUBLE &amp;&amp; f-&gt;op == OCONST) {
	}
	if(ft == TFLOAT &amp;&amp; f-&gt;op == OCONST) {
	}

	/*
	 * a load --
	 * put it into a register then
	 * worry what to do with it.
	 */
	if(f-&gt;op == ONAME || f-&gt;op == OINDREG || f-&gt;op == OIND) {
		switch(ft) {
		default:
			a = AMOVW;
			break;
		case TFLOAT:
			a = AMOVF;
			break;
		case TDOUBLE:
			a = AMOVD;
			break;
		case TCHAR:
			a = AMOVB;
			break;
		case TUCHAR:
			a = AMOVBU;
			break;
		case TSHORT:
			a = AMOVH;
			break;
		case TUSHORT:
			a = AMOVHU;
			break;
		}
		if(typechlp[ft] &amp;&amp; typeilp[tt])
			regalloc(&amp;nod, t, t);
		else
			regalloc(&amp;nod, f, t);
		gins(a, f, &amp;nod);
		gmove(&amp;nod, t);
		regfree(&amp;nod);
		return;
	}

	/*
	 * a store --
	 * put it into a register then
	 * store it.
	 */
	if(t-&gt;op == ONAME || t-&gt;op == OINDREG || t-&gt;op == OIND) {
		switch(tt) {
		default:
			a = AMOVW;
			break;
		case TUCHAR:
			a = AMOVBU;
			break;
		case TCHAR:
			a = AMOVB;
			break;
		case TUSHORT:
			a = AMOVHU;
			break;
		case TSHORT:
			a = AMOVH;
			break;
		case TFLOAT:
			a = AMOVF;
			break;
		case TVLONG:
		case TDOUBLE:
			a = AMOVD;
			break;
		}
		if(ft == tt)
			regalloc(&amp;nod, t, f);
		else
			regalloc(&amp;nod, t, Z);
		gmove(f, &amp;nod);
		gins(a, &amp;nod, t);
		regfree(&amp;nod);
		return;
	}

	/*
	 * type x type cross table
	 */
	a = AGOK;
	switch(ft) {
	case TDOUBLE:
	case TVLONG:
	case TFLOAT:
		switch(tt) {
		case TDOUBLE:
		case TVLONG:
			a = AMOVD;
			if(ft == TFLOAT)
				a = AMOVFD;
			break;
		case TFLOAT:
			a = AMOVDF;
			if(ft == TFLOAT)
				a = AMOVF;
			break;
		case TINT:
		case TUINT:
		case TLONG:
		case TULONG:
		case TIND:
			a = AMOVDW;
			if(ft == TFLOAT)
				a = AMOVFW;
			break;
		case TSHORT:
		case TUSHORT:
		case TCHAR:
		case TUCHAR:
			a = AMOVDW;
			if(ft == TFLOAT)
				a = AMOVFW;
			break;
		}
		break;
	case TUINT:
	case TINT:
	case TULONG:
	case TLONG:
	case TIND:
		switch(tt) {
		case TDOUBLE:
		case TVLONG:
			gins(AMOVWD, f, t);
			if(ft == TULONG) {
			}
			return;
		case TFLOAT:
			gins(AMOVWF, f, t);
			if(ft == TULONG) {
			}
			return;
		case TINT:
		case TUINT:
		case TLONG:
		case TULONG:
		case TIND:
		case TSHORT:
		case TUSHORT:
		case TCHAR:
		case TUCHAR:
			a = AMOVW;
			break;
		}
		break;
	case TSHORT:
		switch(tt) {
		case TDOUBLE:
		case TVLONG:
			regalloc(&amp;nod, f, Z);
			gins(AMOVH, f, &amp;nod);
			gins(AMOVWD, &amp;nod, t);
			regfree(&amp;nod);
			return;
		case TFLOAT:
			regalloc(&amp;nod, f, Z);
			gins(AMOVH, f, &amp;nod);
			gins(AMOVWF, &amp;nod, t);
			regfree(&amp;nod);
			return;
		case TUINT:
		case TINT:
		case TULONG:
		case TLONG:
		case TIND:
			a = AMOVH;
			break;
		case TSHORT:
		case TUSHORT:
		case TCHAR:
		case TUCHAR:
			a = AMOVW;
			break;
		}
		break;
	case TUSHORT:
		switch(tt) {
		case TDOUBLE:
		case TVLONG:
			regalloc(&amp;nod, f, Z);
			gins(AMOVHU, f, &amp;nod);
			gins(AMOVWD, &amp;nod, t);
			regfree(&amp;nod);
			return;
		case TFLOAT:
			regalloc(&amp;nod, f, Z);
			gins(AMOVHU, f, &amp;nod);
			gins(AMOVWF, &amp;nod, t);
			regfree(&amp;nod);
			return;
		case TINT:
		case TUINT:
		case TLONG:
		case TULONG:
		case TIND:
			a = AMOVHU;
			break;
		case TSHORT:
		case TUSHORT:
		case TCHAR:
		case TUCHAR:
			a = AMOVW;
			break;
		}
		break;
	case TCHAR:
		switch(tt) {
		case TDOUBLE:
		case TVLONG:
			regalloc(&amp;nod, f, Z);
			gins(AMOVB, f, &amp;nod);
			gins(AMOVWD, &amp;nod, t);
			regfree(&amp;nod);
			return;
		case TFLOAT:
			regalloc(&amp;nod, f, Z);
			gins(AMOVB, f, &amp;nod);
			gins(AMOVWF, &amp;nod, t);
			regfree(&amp;nod);
			return;
		case TINT:
		case TUINT:
		case TLONG:
		case TULONG:
		case TIND:
		case TSHORT:
		case TUSHORT:
			a = AMOVB;
			break;
		case TCHAR:
		case TUCHAR:
			a = AMOVW;
			break;
		}
		break;
	case TUCHAR:
		switch(tt) {
		case TDOUBLE:
		case TVLONG:
			regalloc(&amp;nod, f, Z);
			gins(AMOVBU, f, &amp;nod);
			gins(AMOVWD, &amp;nod, t);
			regfree(&amp;nod);
			return;
		case TFLOAT:
			regalloc(&amp;nod, f, Z);
			gins(AMOVBU, f, &amp;nod);
			gins(AMOVWF, &amp;nod, t);
			regfree(&amp;nod);
			return;
		case TINT:
		case TUINT:
		case TLONG:
		case TULONG:
		case TIND:
		case TSHORT:
		case TUSHORT:
			a = AMOVBU;
			break;
		case TCHAR:
		case TUCHAR:
			a = AMOVW;
			break;
		}
		break;
	}
	if(a == AGOK)
		diag(Z, &#34;bad opcode in gmove %T -&gt; %T&#34;, f-&gt;type, t-&gt;type);
	if(a == AMOVW || a == AMOVF || a == AMOVD)
	if(samaddr(f, t))
		return;
	gins(a, f, t);
}

void
gmover(Node *f, Node *t)
{
	int ft, tt, a;

	ft = f-&gt;type-&gt;etype;
	tt = t-&gt;type-&gt;etype;
	a = AGOK;
	if(typechlp[ft] &amp;&amp; typechlp[tt] &amp;&amp; ewidth[ft] &gt;= ewidth[tt]){
		switch(tt){
		case TSHORT:
			a = AMOVH;
			break;
		case TUSHORT:
			a = AMOVHU;
			break;
		case TCHAR:
			a = AMOVB;
			break;
		case TUCHAR:
			a = AMOVBU;
			break;
		}
	}
	if(a == AGOK)
		gmove(f, t);
	else
		gins(a, f, t);
}

void
gins(int a, Node *f, Node *t)
{

	nextpc();
	p-&gt;as = a;
	if(f != Z)
		naddr(f, &amp;p-&gt;from);
	if(t != Z)
		naddr(t, &amp;p-&gt;to);
	if(debug[&#39;g&#39;])
		print(&#34;%P\n&#34;, p);
}

void
gopcode(int o, Node *f1, Node *f2, Node *t)
{
	int a, et;
	Adr ta;

	et = TLONG;
	if(f1 != Z &amp;&amp; f1-&gt;type != T)
		et = f1-&gt;type-&gt;etype;
	a = AGOK;
	switch(o) {
	case OAS:
		gmove(f1, t);
		return;

	case OASADD:
	case OADD:
		a = AADD;
		if(et == TFLOAT)
			a = AADDF;
		else
		if(et == TDOUBLE || et == TVLONG)
			a = AADDD;
		break;

	case OASSUB:
	case OSUB:
		if(f2 &amp;&amp; f2-&gt;op == OCONST) {
			Node *t = f1;
			f1 = f2;
			f2 = t;
			a = ARSB;
		} else
			a = ASUB;
		if(et == TFLOAT)
			a = ASUBF;
		else
		if(et == TDOUBLE || et == TVLONG)
			a = ASUBD;
		break;

	case OASOR:
	case OOR:
		a = AORR;
		break;

	case OASAND:
	case OAND:
		a = AAND;
		break;

	case OASXOR:
	case OXOR:
		a = AEOR;
		break;

	case OASLSHR:
	case OLSHR:
		a = ASRL;
		break;

	case OASASHR:
	case OASHR:
		a = ASRA;
		break;

	case OASASHL:
	case OASHL:
		a = ASLL;
		break;

	case OFUNC:
		a = ABL;
		break;

	case OASMUL:
	case OMUL:
		a = AMUL;
		if(et == TFLOAT)
			a = AMULF;
		else
		if(et == TDOUBLE || et == TVLONG)
			a = AMULD;
		break;

	case OASDIV:
	case ODIV:
		a = ADIV;
		if(et == TFLOAT)
			a = ADIVF;
		else
		if(et == TDOUBLE || et == TVLONG)
			a = ADIVD;
		break;

	case OASMOD:
	case OMOD:
		a = AMOD;
		break;

	case OASLMUL:
	case OLMUL:
		a = AMULU;
		break;

	case OASLMOD:
	case OLMOD:
		a = AMODU;
		break;

	case OASLDIV:
	case OLDIV:
		a = ADIVU;
		break;

	case OCASE:
	case OEQ:
	case ONE:
	case OLT:
	case OLE:
	case OGE:
	case OGT:
	case OLO:
	case OLS:
	case OHS:
	case OHI:
		a = ACMP;
		if(et == TFLOAT)
			a = ACMPF;
		else
		if(et == TDOUBLE || et == TVLONG)
			a = ACMPD;
		nextpc();
		p-&gt;as = a;
		naddr(f1, &amp;p-&gt;from);
		if(a == ACMP &amp;&amp; f1-&gt;op == OCONST &amp;&amp; p-&gt;from.offset &lt; 0) {
			p-&gt;as = ACMN;
			p-&gt;from.offset = -p-&gt;from.offset;
		}
		raddr(f2, p);
		switch(o) {
		case OEQ:
			a = ABEQ;
			break;
		case ONE:
			a = ABNE;
			break;
		case OLT:
			a = ABLT;
			break;
		case OLE:
			a = ABLE;
			break;
		case OGE:
			a = ABGE;
			break;
		case OGT:
			a = ABGT;
			break;
		case OLO:
			a = ABLO;
			break;
		case OLS:
			a = ABLS;
			break;
		case OHS:
			a = ABHS;
			break;
		case OHI:
			a = ABHI;
			break;
		case OCASE:
			nextpc();
			p-&gt;as = ACASE;
			p-&gt;scond = 0x9;
			naddr(f2, &amp;p-&gt;from);
			a = ABHI;
			break;
		}
		f1 = Z;
		f2 = Z;
		break;
	}
	if(a == AGOK)
		diag(Z, &#34;bad in gopcode %O&#34;, o);
	nextpc();
	p-&gt;as = a;
	if(f1 != Z)
		naddr(f1, &amp;p-&gt;from);
	if(f2 != Z) {
		naddr(f2, &amp;ta);
		p-&gt;reg = ta.reg;
	}
	if(t != Z)
		naddr(t, &amp;p-&gt;to);
	if(debug[&#39;g&#39;])
		print(&#34;%P\n&#34;, p);
}

int
samaddr(Node *f, Node *t)
{

	if(f-&gt;op != t-&gt;op)
		return 0;
	switch(f-&gt;op) {

	case OREGISTER:
		if(f-&gt;reg != t-&gt;reg)
			break;
		return 1;
	}
	return 0;
}

void
gbranch(int o)
{
	int a;

	a = AGOK;
	switch(o) {
	case ORETURN:
		a = ARET;
		break;
	case OGOTO:
		a = AB;
		break;
	}
	nextpc();
	if(a == AGOK) {
		diag(Z, &#34;bad in gbranch %O&#34;,  o);
		nextpc();
	}
	p-&gt;as = a;
}

void
patch(Prog *op, int32 pc)
{

	op-&gt;to.offset = pc;
	op-&gt;to.type = D_BRANCH;
}

void
gpseudo(int a, Sym *s, Node *n)
{

	nextpc();
	p-&gt;as = a;
	p-&gt;from.type = D_OREG;
	p-&gt;from.sym = s;
	p-&gt;from.name = D_EXTERN;
	if(a == ATEXT)
		p-&gt;reg = textflag;
	if(s-&gt;class == CSTATIC)
		p-&gt;from.name = D_STATIC;
	naddr(n, &amp;p-&gt;to);
	if(a == ADATA || a == AGLOBL)
		pc--;
}

int
sconst(Node *n)
{
	vlong vv;

	if(n-&gt;op == OCONST) {
		if(!typefd[n-&gt;type-&gt;etype]) {
			vv = n-&gt;vconst;
			if(vv &gt;= (vlong)(-32766) &amp;&amp; vv &lt; (vlong)32766)
				return 1;
			/*
			 * should be specialised for constant values which will
			 * fit in different instructionsl; for now, let 5l
			 * sort it out
			 */
			return 1;
		}
	}
	return 0;
}

int
sval(int32 v)
{
	int i;

	for(i=0; i&lt;16; i++) {
		if((v &amp; ~0xff) == 0)
			return 1;
		if((~v &amp; ~0xff) == 0)
			return 1;
		v = (v&lt;&lt;2) | ((uint32)v&gt;&gt;30);
	}
	return 0;
}

int32
exreg(Type *t)
{
	int32 o;

	if(typechlp[t-&gt;etype]) {
		if(exregoffset &lt;= REGEXT-4)
			return 0;
		o = exregoffset;
		exregoffset--;
		return o;
	}
	if(typefd[t-&gt;etype]) {
		if(exfregoffset &lt;= NFREG-1)
			return 0;
		o = exfregoffset + NREG;
		exfregoffset--;
		return o;
	}
	return 0;
}

schar	ewidth[NTYPE] =
{
	-1,		/* [TXXX] */
	SZ_CHAR,	/* [TCHAR] */
	SZ_CHAR,	/* [TUCHAR] */
	SZ_SHORT,	/* [TSHORT] */
	SZ_SHORT,	/* [TUSHORT] */
	SZ_INT,		/* [TINT] */
	SZ_INT,		/* [TUINT] */
	SZ_LONG,	/* [TLONG] */
	SZ_LONG,	/* [TULONG] */
	SZ_VLONG,	/* [TVLONG] */
	SZ_VLONG,	/* [TUVLONG] */
	SZ_FLOAT,	/* [TFLOAT] */
	SZ_DOUBLE,	/* [TDOUBLE] */
	SZ_IND,		/* [TIND] */
	0,		/* [TFUNC] */
	-1,		/* [TARRAY] */
	0,		/* [TVOID] */
	-1,		/* [TSTRUCT] */
	-1,		/* [TUNION] */
	SZ_INT,		/* [TENUM] */
};

int32	ncast[NTYPE] =
{
	0,				/* [TXXX] */
	BCHAR|BUCHAR,			/* [TCHAR] */
	BCHAR|BUCHAR,			/* [TUCHAR] */
	BSHORT|BUSHORT,			/* [TSHORT] */
	BSHORT|BUSHORT,			/* [TUSHORT] */
	BINT|BUINT|BLONG|BULONG|BIND,	/* [TINT] */
	BINT|BUINT|BLONG|BULONG|BIND,	/* [TUINT] */
	BINT|BUINT|BLONG|BULONG|BIND,	/* [TLONG] */
	BINT|BUINT|BLONG|BULONG|BIND,	/* [TULONG] */
	BVLONG|BUVLONG,			/* [TVLONG] */
	BVLONG|BUVLONG,			/* [TUVLONG] */
	BFLOAT,				/* [TFLOAT] */
	BDOUBLE,			/* [TDOUBLE] */
	BLONG|BULONG|BIND,		/* [TIND] */
	0,				/* [TFUNC] */
	0,				/* [TARRAY] */
	0,				/* [TVOID] */
	BSTRUCT,			/* [TSTRUCT] */
	BUNION,				/* [TUNION] */
	0,				/* [TENUM] */
};
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
