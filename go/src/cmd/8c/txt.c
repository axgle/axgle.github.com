<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8c/txt.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/8c/txt.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/8c/txt.c
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

#include &#34;gc.h&#34;

void
ginit(void)
{
	int i;
	Type *t;

	thechar = &#39;8&#39;;
	thestring = &#34;386&#34;;
	exregoffset = 0;
	exfregoffset = 0;
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
	zprog.from.type = D_NONE;
	zprog.from.index = D_NONE;
	zprog.from.scale = 0;
	zprog.to = zprog.from;

	regnode.op = OREGISTER;
	regnode.class = CEXREG;
	regnode.reg = REGTMP;
	regnode.complex = 0;
	regnode.addable = 11;
	regnode.type = types[TLONG];

	fregnode0 = regnode;
	fregnode0.reg = D_F0;
	fregnode0.type = types[TDOUBLE];

	fregnode1 = fregnode0;
	fregnode1.reg = D_F0+1;

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

	for(i=0; i&lt;nelem(reg); i++) {
		reg[i] = 1;
		if(i &gt;= D_AX &amp;&amp; i &lt;= D_DI &amp;&amp; i != D_SP)
			reg[i] = 0;
	}
}

void
gclean(void)
{
	int i;
	Sym *s;

	reg[D_SP]--;
	for(i=D_AX; i&lt;=D_DI; i++)
		if(reg[i])
			diag(Z, &#34;reg %R left allocated&#34;, i);
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

int nareg(void)
{
	int i, n;

	n = 0;
	for(i=D_AX; i&lt;=D_DI; i++)
		if(reg[i] == 0)
			n++;
	return n;
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
	if(typesu[n-&gt;type-&gt;etype] || typev[n-&gt;type-&gt;etype]) {
		regaalloc(tn2, n);
		if(n-&gt;complex &gt;= FNX) {
			sugen(*fnxp, tn2, n-&gt;type-&gt;width);
			(*fnxp)++;
		} else
			sugen(n, tn2, n-&gt;type-&gt;width);
		return;
	}
	if(REGARG &gt;= 0 &amp;&amp; curarg == 0 &amp;&amp; typeilp[n-&gt;type-&gt;etype]) {
		regaalloc1(tn1, n);
		if(n-&gt;complex &gt;= FNX) {
			cgen(*fnxp, tn1);
			(*fnxp)++;
		} else
			cgen(n, tn1);
		return;
	}
	if(vconst(n) == 0) {
		regaalloc(tn2, n);
		gmove(n, tn2);
		return;
	}
	regalloc(tn1, n, Z);
	if(n-&gt;complex &gt;= FNX) {
		cgen(*fnxp, tn1);
		(*fnxp)++;
	} else
		cgen(n, tn1);
	regaalloc(tn2, n);
	gmove(tn1, tn2);
	regfree(tn1);
}

Node*
nodconst(int32 v)
{
	constnode.vconst = v;
	return &amp;constnode;
}

Node*
nodfconst(double d)
{
	fconstnode.fconst = d;
	return &amp;fconstnode;
}

int
isreg(Node *n, int r)
{

	if(n-&gt;op == OREGISTER)
		if(n-&gt;reg == r)
			return 1;
	return 0;
}

int
nodreg(Node *n, Node *nn, int r)
{

	*n = regnode;
	n-&gt;reg = r;
	if(reg[r] == 0)
		return 0;
	if(nn != Z) {
		n-&gt;type = nn-&gt;type;
		n-&gt;lineno = nn-&gt;lineno;
		if(nn-&gt;op == OREGISTER)
		if(nn-&gt;reg == r)
			return 0;
	}
	return 1;
}

void
regret(Node *n, Node *nn)
{
	int r;

	r = REGRET;
	if(typefd[nn-&gt;type-&gt;etype])
		r = FREGRET;
	nodreg(n, nn, r);
	reg[r]++;
}

void
regalloc(Node *n, Node *tn, Node *o)
{
	int i;

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
			if(i &gt;= D_AX &amp;&amp; i &lt;= D_DI)
				goto out;
		}
		for(i=D_AX; i&lt;=D_DI; i++)
			if(reg[i] == 0)
				goto out;
		diag(tn, &#34;out of fixed registers&#34;);
		goto err;

	case TFLOAT:
	case TDOUBLE:
	case TVLONG:
		i = D_F0;
		goto out;
	}
	diag(tn, &#34;unknown type in regalloc: %T&#34;, tn-&gt;type);
err:
	i = 0;
out:
	if(i)
		reg[i]++;
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
	diag(n, &#34;error in regfree: %R&#34;, i);
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
	n-&gt;xoffset = curarg;
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
naddr(Node *n, Adr *a)
{
	int32 v;

	a-&gt;type = D_NONE;
	if(n == Z)
		return;
	switch(n-&gt;op) {
	default:
	bad:
		diag(n, &#34;bad in naddr: %O %D&#34;, n-&gt;op, a);
		break;

	case OREGISTER:
		a-&gt;type = n-&gt;reg;
		a-&gt;sym = S;
		break;

	case OEXREG:
		a-&gt;type = D_INDIR + D_GS;
		a-&gt;offset = n-&gt;reg - 1;
		break;

	case OIND:
		naddr(n-&gt;left, a);
		if(a-&gt;type &gt;= D_AX &amp;&amp; a-&gt;type &lt;= D_DI)
			a-&gt;type += D_INDIR;
		else
		if(a-&gt;type == D_CONST)
			a-&gt;type = D_NONE+D_INDIR;
		else
		if(a-&gt;type == D_ADDR) {
			a-&gt;type = a-&gt;index;
			a-&gt;index = D_NONE;
		} else
			goto bad;
		break;

	case OINDEX:
		a-&gt;type = idx.ptr;
		if(n-&gt;left-&gt;op == OADDR || n-&gt;left-&gt;op == OCONST)
			naddr(n-&gt;left, a);
		if(a-&gt;type &gt;= D_AX &amp;&amp; a-&gt;type &lt;= D_DI)
			a-&gt;type += D_INDIR;
		else
		if(a-&gt;type == D_CONST)
			a-&gt;type = D_NONE+D_INDIR;
		else
		if(a-&gt;type == D_ADDR) {
			a-&gt;type = a-&gt;index;
			a-&gt;index = D_NONE;
		} else
			goto bad;
		a-&gt;index = idx.reg;
		a-&gt;scale = n-&gt;scale;
		a-&gt;offset += n-&gt;xoffset;
		break;

	case OINDREG:
		a-&gt;type = n-&gt;reg+D_INDIR;
		a-&gt;sym = S;
		a-&gt;offset = n-&gt;xoffset;
		break;

	case ONAME:
		a-&gt;etype = n-&gt;etype;
		a-&gt;type = D_STATIC;
		a-&gt;sym = n-&gt;sym;
		a-&gt;offset = n-&gt;xoffset;
		if(n-&gt;class == CSTATIC)
			break;
		if(n-&gt;class == CEXTERN || n-&gt;class == CGLOBL) {
			a-&gt;type = D_EXTERN;
			break;
		}
		if(n-&gt;class == CAUTO) {
			a-&gt;type = D_AUTO;
			break;
		}
		if(n-&gt;class == CPARAM) {
			a-&gt;type = D_PARAM;
			break;
		}
		goto bad;

	case OCONST:
		if(typefd[n-&gt;type-&gt;etype]) {
			a-&gt;type = D_FCONST;
			a-&gt;dval = n-&gt;fconst;
			break;
		}
		a-&gt;sym = S;
		a-&gt;type = D_CONST;
		a-&gt;offset = n-&gt;vconst;
		break;

	case OADDR:
		naddr(n-&gt;left, a);
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
		goto bad;

	case OADD:
		if(n-&gt;right-&gt;op == OCONST) {
			v = n-&gt;right-&gt;vconst;
			naddr(n-&gt;left, a);
		} else
		if(n-&gt;left-&gt;op == OCONST) {
			v = n-&gt;left-&gt;vconst;
			naddr(n-&gt;right, a);
		} else
			goto bad;
		a-&gt;offset += v;
		break;

	}
}

#define	CASE(a,b)	((a&lt;&lt;8)|(b&lt;&lt;0))

void
gmove(Node *f, Node *t)
{
	int ft, tt, a;
	Node nod, nod1;
	Prog *p1;

	ft = f-&gt;type-&gt;etype;
	tt = t-&gt;type-&gt;etype;
	if(debug[&#39;M&#39;])
		print(&#34;gop: %O %O[%s],%O[%s]\n&#34;, OAS,
			f-&gt;op, tnames[ft], t-&gt;op, tnames[tt]);
	if(typefd[ft] &amp;&amp; f-&gt;op == OCONST) {
		if(f-&gt;fconst == 0)
			gins(AFLDZ, Z, Z);
		else
		if(f-&gt;fconst == 1)
			gins(AFLD1, Z, Z);
		else
			gins(AFMOVD, f, &amp;fregnode0);
		gmove(&amp;fregnode0, t);
		return;
	}
/*
 * load
 */
	if(f-&gt;op == ONAME || f-&gt;op == OINDREG ||
	   f-&gt;op == OIND || f-&gt;op == OINDEX)
	switch(ft) {
	case TCHAR:
		a = AMOVBLSX;
		goto ld;
	case TUCHAR:
		a = AMOVBLZX;
		goto ld;
	case TSHORT:
		if(typefd[tt]) {
			gins(AFMOVW, f, &amp;fregnode0);
			gmove(&amp;fregnode0, t);
			return;
		}
		a = AMOVWLSX;
		goto ld;
	case TUSHORT:
		a = AMOVWLZX;
		goto ld;
	case TINT:
	case TUINT:
	case TLONG:
	case TULONG:
	case TIND:
		if(typefd[tt]) {
			gins(AFMOVL, f, &amp;fregnode0);
			gmove(&amp;fregnode0, t);
			return;
		}
		a = AMOVL;

	ld:
		regalloc(&amp;nod, f, t);
		nod.type = types[TLONG];
		gins(a, f, &amp;nod);
		gmove(&amp;nod, t);
		regfree(&amp;nod);
		return;

	case TFLOAT:
		gins(AFMOVF, f, t);
		return;
	case TDOUBLE:
		gins(AFMOVD, f, t);
		return;
	case TVLONG:
		gins(AFMOVV, f, t);
		return;
	}

/*
 * store
 */
	if(t-&gt;op == ONAME || t-&gt;op == OINDREG ||
	   t-&gt;op == OIND || t-&gt;op == OINDEX)
	switch(tt) {
	case TCHAR:
	case TUCHAR:
		a = AMOVB;	goto st;
	case TSHORT:
	case TUSHORT:
		a = AMOVW;	goto st;
	case TINT:
	case TUINT:
	case TLONG:
	case TULONG:
	case TIND:
		a = AMOVL;	goto st;

	st:
		if(f-&gt;op == OCONST) {
			gins(a, f, t);
			return;
		}
		regalloc(&amp;nod, t, f);
		gmove(f, &amp;nod);
		gins(a, &amp;nod, t);
		regfree(&amp;nod);
		return;

	case TFLOAT:
		gins(AFMOVFP, f, t);
		return;
	case TDOUBLE:
		gins(AFMOVDP, f, t);
		return;
	case TVLONG:
		gins(AFMOVVP, f, t);
		return;
	}

/*
 * convert
 */
	switch(CASE(ft,tt)) {
	default:
/*
 * integer to integer
 ********
		a = AGOK;	break;

	case CASE(	TCHAR,	TCHAR):
	case CASE(	TUCHAR,	TCHAR):
	case CASE(	TSHORT,	TCHAR):
	case CASE(	TUSHORT,TCHAR):
	case CASE(	TINT,	TCHAR):
	case CASE(	TUINT,	TCHAR):
	case CASE(	TLONG,	TCHAR):
	case CASE(	TULONG,	TCHAR):
	case CASE(	TIND,	TCHAR):

	case CASE(	TCHAR,	TUCHAR):
	case CASE(	TUCHAR,	TUCHAR):
	case CASE(	TSHORT,	TUCHAR):
	case CASE(	TUSHORT,TUCHAR):
	case CASE(	TINT,	TUCHAR):
	case CASE(	TUINT,	TUCHAR):
	case CASE(	TLONG,	TUCHAR):
	case CASE(	TULONG,	TUCHAR):
	case CASE(	TIND,	TUCHAR):

	case CASE(	TSHORT,	TSHORT):
	case CASE(	TUSHORT,TSHORT):
	case CASE(	TINT,	TSHORT):
	case CASE(	TUINT,	TSHORT):
	case CASE(	TLONG,	TSHORT):
	case CASE(	TULONG,	TSHORT):
	case CASE(	TIND,	TSHORT):

	case CASE(	TSHORT,	TUSHORT):
	case CASE(	TUSHORT,TUSHORT):
	case CASE(	TINT,	TUSHORT):
	case CASE(	TUINT,	TUSHORT):
	case CASE(	TLONG,	TUSHORT):
	case CASE(	TULONG,	TUSHORT):
	case CASE(	TIND,	TUSHORT):

	case CASE(	TINT,	TINT):
	case CASE(	TUINT,	TINT):
	case CASE(	TLONG,	TINT):
	case CASE(	TULONG,	TINT):
	case CASE(	TIND,	TINT):

	case CASE(	TINT,	TUINT):
	case CASE(	TUINT,	TUINT):
	case CASE(	TLONG,	TUINT):
	case CASE(	TULONG,	TUINT):
	case CASE(	TIND,	TUINT):

	case CASE(	TINT,	TLONG):
	case CASE(	TUINT,	TLONG):
	case CASE(	TLONG,	TLONG):
	case CASE(	TULONG,	TLONG):
	case CASE(	TIND,	TLONG):

	case CASE(	TINT,	TULONG):
	case CASE(	TUINT,	TULONG):
	case CASE(	TLONG,	TULONG):
	case CASE(	TULONG,	TULONG):
	case CASE(	TIND,	TULONG):

	case CASE(	TINT,	TIND):
	case CASE(	TUINT,	TIND):
	case CASE(	TLONG,	TIND):
	case CASE(	TULONG,	TIND):
	case CASE(	TIND,	TIND):
 *****/
		a = AMOVL;
		break;

	case CASE(	TSHORT,	TINT):
	case CASE(	TSHORT,	TUINT):
	case CASE(	TSHORT,	TLONG):
	case CASE(	TSHORT,	TULONG):
	case CASE(	TSHORT,	TIND):
		a = AMOVWLSX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= 0xffff;
			if(f-&gt;vconst &amp; 0x8000)
				f-&gt;vconst |= 0xffff0000;
			a = AMOVL;
		}
		break;

	case CASE(	TUSHORT,TINT):
	case CASE(	TUSHORT,TUINT):
	case CASE(	TUSHORT,TLONG):
	case CASE(	TUSHORT,TULONG):
	case CASE(	TUSHORT,TIND):
		a = AMOVWLZX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= 0xffff;
			a = AMOVL;
		}
		break;

	case CASE(	TCHAR,	TSHORT):
	case CASE(	TCHAR,	TUSHORT):
	case CASE(	TCHAR,	TINT):
	case CASE(	TCHAR,	TUINT):
	case CASE(	TCHAR,	TLONG):
	case CASE(	TCHAR,	TULONG):
	case CASE(	TCHAR,	TIND):
		a = AMOVBLSX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= 0xff;
			if(f-&gt;vconst &amp; 0x80)
				f-&gt;vconst |= 0xffffff00;
			a = AMOVL;
		}
		break;

	case CASE(	TUCHAR,	TSHORT):
	case CASE(	TUCHAR,	TUSHORT):
	case CASE(	TUCHAR,	TINT):
	case CASE(	TUCHAR,	TUINT):
	case CASE(	TUCHAR,	TLONG):
	case CASE(	TUCHAR,	TULONG):
	case CASE(	TUCHAR,	TIND):
		a = AMOVBLZX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= 0xff;
			a = AMOVL;
		}
		break;

/*
 * float to fix
 */
	case CASE(	TFLOAT,	TCHAR):
	case CASE(	TFLOAT,	TUCHAR):
	case CASE(	TFLOAT,	TSHORT):
	case CASE(	TFLOAT,	TUSHORT):
	case CASE(	TFLOAT,	TINT):
	case CASE(	TFLOAT,	TUINT):
	case CASE(	TFLOAT,	TLONG):
	case CASE(	TFLOAT,	TULONG):
	case CASE(	TFLOAT,	TIND):

	case CASE(	TDOUBLE,TCHAR):
	case CASE(	TDOUBLE,TUCHAR):
	case CASE(	TDOUBLE,TSHORT):
	case CASE(	TDOUBLE,TUSHORT):
	case CASE(	TDOUBLE,TINT):
	case CASE(	TDOUBLE,TUINT):
	case CASE(	TDOUBLE,TLONG):
	case CASE(	TDOUBLE,TULONG):
	case CASE(	TDOUBLE,TIND):

	case CASE(	TVLONG,	TCHAR):
	case CASE(	TVLONG,	TUCHAR):
	case CASE(	TVLONG,	TSHORT):
	case CASE(	TVLONG,	TUSHORT):
	case CASE(	TVLONG,	TINT):
	case CASE(	TVLONG,	TUINT):
	case CASE(	TVLONG,	TLONG):
	case CASE(	TVLONG,	TULONG):
	case CASE(	TVLONG,	TIND):
		if(fproundflg) {
			regsalloc(&amp;nod, &amp;regnode);
			gins(AFMOVLP, f, &amp;nod);
			gmove(&amp;nod, t);
			return;
		}
		regsalloc(&amp;nod, &amp;regnode);
		regsalloc(&amp;nod1, &amp;regnode);
		gins(AFSTCW, Z, &amp;nod1);
		nod1.xoffset += 2;
		gins(AMOVW, nodconst(0xf7f), &amp;nod1);
		gins(AFLDCW, &amp;nod1, Z);
		gins(AFMOVLP, f, &amp;nod);
		nod1.xoffset -= 2;
		gins(AFLDCW, &amp;nod1, Z);
		gmove(&amp;nod, t);
		return;

/*
 * ulong to float
 */
	case CASE(	TULONG,	TDOUBLE):
	case CASE(	TULONG,	TVLONG):
	case CASE(	TULONG,	TFLOAT):
	case CASE(	TUINT,	TDOUBLE):
	case CASE(	TUINT,	TVLONG):
	case CASE(	TUINT,	TFLOAT):
		regalloc(&amp;nod, f, f);
		gmove(f, &amp;nod);
		regsalloc(&amp;nod1, &amp;regnode);
		gmove(&amp;nod, &amp;nod1);
		gins(AFMOVL, &amp;nod1, &amp;fregnode0);
		gins(ACMPL, &amp;nod, nodconst(0));
		gins(AJGE, Z, Z);
		p1 = p;
		gins(AFADDD, nodfconst(4294967296.), &amp;fregnode0);
		patch(p1, pc);
		regfree(&amp;nod);
		return;

/*
 * fix to float
 */
	case CASE(	TCHAR,	TFLOAT):
	case CASE(	TUCHAR,	TFLOAT):
	case CASE(	TSHORT,	TFLOAT):
	case CASE(	TUSHORT,TFLOAT):
	case CASE(	TINT,	TFLOAT):
	case CASE(	TLONG,	TFLOAT):
	case CASE(	TIND,	TFLOAT):

	case CASE(	TCHAR,	TDOUBLE):
	case CASE(	TUCHAR,	TDOUBLE):
	case CASE(	TSHORT,	TDOUBLE):
	case CASE(	TUSHORT,TDOUBLE):
	case CASE(	TINT,	TDOUBLE):
	case CASE(	TLONG,	TDOUBLE):
	case CASE(	TIND,	TDOUBLE):

	case CASE(	TCHAR,	TVLONG):
	case CASE(	TUCHAR,	TVLONG):
	case CASE(	TSHORT,	TVLONG):
	case CASE(	TUSHORT,TVLONG):
	case CASE(	TINT,	TVLONG):
	case CASE(	TLONG,	TVLONG):
	case CASE(	TIND,	TVLONG):
		regsalloc(&amp;nod, &amp;regnode);
		gmove(f, &amp;nod);
		gins(AFMOVL, &amp;nod, &amp;fregnode0);
		return;

/*
 * float to float
 */
	case CASE(	TFLOAT,	TFLOAT):
	case CASE(	TDOUBLE,TFLOAT):
	case CASE(	TVLONG,	TFLOAT):

	case CASE(	TFLOAT,	TDOUBLE):
	case CASE(	TDOUBLE,TDOUBLE):
	case CASE(	TVLONG,	TDOUBLE):

	case CASE(	TFLOAT,	TVLONG):
	case CASE(	TDOUBLE,TVLONG):
	case CASE(	TVLONG,	TVLONG):
		a = AFMOVD;	break;
	}
	if(a == AMOVL || a == AFMOVD)
	if(samaddr(f, t))
		return;
	gins(a, f, t);
}

void
doindex(Node *n)
{
	Node nod, nod1;
	int32 v;

if(debug[&#39;Y&#39;])
prtree(n, &#34;index&#34;);

if(n-&gt;left-&gt;complex &gt;= FNX)
print(&#34;botch in doindex\n&#34;);

	regalloc(&amp;nod, &amp;regnode, Z);
	v = constnode.vconst;
	cgen(n-&gt;right, &amp;nod);
	idx.ptr = D_NONE;
	if(n-&gt;left-&gt;op == OCONST)
		idx.ptr = D_CONST;
	else if(n-&gt;left-&gt;op == OREGISTER)
		idx.ptr = n-&gt;left-&gt;reg;
	else if(n-&gt;left-&gt;op != OADDR) {
		reg[D_BP]++;	// cant be used as a base
		regalloc(&amp;nod1, &amp;regnode, Z);
		cgen(n-&gt;left, &amp;nod1);
		idx.ptr = nod1.reg;
		regfree(&amp;nod1);
		reg[D_BP]--;
	}
	idx.reg = nod.reg;
	regfree(&amp;nod);
	constnode.vconst = v;
}

void
gins(int a, Node *f, Node *t)
{

	if(f != Z &amp;&amp; f-&gt;op == OINDEX)
		doindex(f);
	if(t != Z &amp;&amp; t-&gt;op == OINDEX)
		doindex(t);
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
fgopcode(int o, Node *f, Node *t, int pop, int rev)
{
	int a, et;
	Node nod;

	et = TLONG;
	if(f != Z &amp;&amp; f-&gt;type != T)
		et = f-&gt;type-&gt;etype;
	if(!typefd[et]) {
		diag(f, &#34;fop: integer %O&#34;, o);
		return;
	}
	if(debug[&#39;M&#39;]) {
		if(t != Z &amp;&amp; t-&gt;type != T)
			print(&#34;gop: %O %O-%s Z\n&#34;, o, f-&gt;op, tnames[et]);
		else
			print(&#34;gop: %O %O-%s %O-%s\n&#34;, o,
				f-&gt;op, tnames[et], t-&gt;op, tnames[t-&gt;type-&gt;etype]);
	}
	a = AGOK;
	switch(o) {

	case OASADD:
	case OADD:
		if(et == TFLOAT)
			a = AFADDF;
		else
		if(et == TDOUBLE || et == TVLONG) {
			a = AFADDD;
			if(pop)
				a = AFADDDP;
		}
		break;

	case OASSUB:
	case OSUB:
		if(et == TFLOAT) {
			a = AFSUBF;
			if(rev)
				a = AFSUBRF;
		} else
		if(et == TDOUBLE || et == TVLONG) {
			a = AFSUBD;
			if(pop)
				a = AFSUBDP;
			if(rev) {
				a = AFSUBRD;
				if(pop)
					a = AFSUBRDP;
			}
		}
		break;

	case OASMUL:
	case OMUL:
		if(et == TFLOAT)
			a = AFMULF;
		else
		if(et == TDOUBLE || et == TVLONG) {
			a = AFMULD;
			if(pop)
				a = AFMULDP;
		}
		break;

	case OASMOD:
	case OMOD:
	case OASDIV:
	case ODIV:
		if(et == TFLOAT) {
			a = AFDIVF;
			if(rev)
				a = AFDIVRF;
		} else
		if(et == TDOUBLE || et == TVLONG) {
			a = AFDIVD;
			if(pop)
				a = AFDIVDP;
			if(rev) {
				a = AFDIVRD;
				if(pop)
					a = AFDIVRDP;
			}
		}
		break;

	case OEQ:
	case ONE:
	case OLT:
	case OLE:
	case OGE:
	case OGT:
		pop += rev;
		if(et == TFLOAT) {
			a = AFCOMF;
			if(pop) {
				a = AFCOMFP;
				if(pop &gt; 1)
					a = AGOK;
			}
		} else
		if(et == TDOUBLE || et == TVLONG) {
			a = AFCOMF;
			if(pop) {
				a = AFCOMDP;
				if(pop &gt; 1)
					a = AFCOMDPP;
			}
		}
		gins(a, f, t);
		regalloc(&amp;nod, &amp;regnode, Z);
		if(nod.reg != D_AX) {
			regfree(&amp;nod);
			nod.reg = D_AX;
			gins(APUSHL, &amp;nod, Z);
			gins(AWAIT, Z, Z);
			gins(AFSTSW, Z, &amp;nod);
			gins(ASAHF, Z, Z);
			gins(APOPL, Z, &amp;nod);
		} else {
			gins(AWAIT, Z, Z);
			gins(AFSTSW, Z, &amp;nod);
			gins(ASAHF, Z, Z);
			regfree(&amp;nod);
		}
		switch(o) {
		case OEQ:	a = AJEQ; break;
		case ONE:	a = AJNE; break;
		case OLT:	a = AJCS; break;
		case OLE:	a = AJLS; break;
		case OGE:	a = AJCC; break;
		case OGT:	a = AJHI; break;
		}
		gins(a, Z, Z);
		return;
	}
	if(a == AGOK)
		diag(Z, &#34;bad in gopcode %O&#34;, o);
	gins(a, f, t);
}

void
gopcode(int o, Type *ty, Node *f, Node *t)
{
	int a, et;

	et = TLONG;
	if(ty != T)
		et = ty-&gt;etype;
	if(typefd[et] &amp;&amp; o != OADDR &amp;&amp; o != OFUNC) {
		diag(f, &#34;gop: float %O&#34;, o);
		return;
	}
	if(debug[&#39;M&#39;]) {
		if(f != Z &amp;&amp; f-&gt;type != T)
			print(&#34;gop: %O %O[%s],&#34;, o, f-&gt;op, tnames[et]);
		else
			print(&#34;gop: %O Z,&#34;, o);
		if(t != Z &amp;&amp; t-&gt;type != T)
			print(&#34;%O[%s]\n&#34;, t-&gt;op, tnames[t-&gt;type-&gt;etype]);
		else
			print(&#34;Z\n&#34;);
	}
	a = AGOK;
	switch(o) {
	case OCOM:
		a = ANOTL;
		if(et == TCHAR || et == TUCHAR)
			a = ANOTB;
		if(et == TSHORT || et == TUSHORT)
			a = ANOTW;
		break;

	case ONEG:
		a = ANEGL;
		if(et == TCHAR || et == TUCHAR)
			a = ANEGB;
		if(et == TSHORT || et == TUSHORT)
			a = ANEGW;
		break;

	case OADDR:
		a = ALEAL;
		break;

	case OASADD:
	case OADD:
		a = AADDL;
		if(et == TCHAR || et == TUCHAR)
			a = AADDB;
		if(et == TSHORT || et == TUSHORT)
			a = AADDW;
		break;

	case OASSUB:
	case OSUB:
		a = ASUBL;
		if(et == TCHAR || et == TUCHAR)
			a = ASUBB;
		if(et == TSHORT || et == TUSHORT)
			a = ASUBW;
		break;

	case OASOR:
	case OOR:
		a = AORL;
		if(et == TCHAR || et == TUCHAR)
			a = AORB;
		if(et == TSHORT || et == TUSHORT)
			a = AORW;
		break;

	case OASAND:
	case OAND:
		a = AANDL;
		if(et == TCHAR || et == TUCHAR)
			a = AANDB;
		if(et == TSHORT || et == TUSHORT)
			a = AANDW;
		break;

	case OASXOR:
	case OXOR:
		a = AXORL;
		if(et == TCHAR || et == TUCHAR)
			a = AXORB;
		if(et == TSHORT || et == TUSHORT)
			a = AXORW;
		break;

	case OASLSHR:
	case OLSHR:
		a = ASHRL;
		if(et == TCHAR || et == TUCHAR)
			a = ASHRB;
		if(et == TSHORT || et == TUSHORT)
			a = ASHRW;
		break;

	case OASASHR:
	case OASHR:
		a = ASARL;
		if(et == TCHAR || et == TUCHAR)
			a = ASARB;
		if(et == TSHORT || et == TUSHORT)
			a = ASARW;
		break;

	case OASASHL:
	case OASHL:
		a = ASALL;
		if(et == TCHAR || et == TUCHAR)
			a = ASALB;
		if(et == TSHORT || et == TUSHORT)
			a = ASALW;
		break;

	case OFUNC:
		a = ACALL;
		break;

	case OASMUL:
	case OMUL:
		if(f-&gt;op == OREGISTER &amp;&amp; t != Z &amp;&amp; isreg(t, D_AX) &amp;&amp; reg[D_DX] == 0)
			t = Z;
		a = AIMULL;
		break;

	case OASMOD:
	case OMOD:
	case OASDIV:
	case ODIV:
		a = AIDIVL;
		break;

	case OASLMUL:
	case OLMUL:
		a = AMULL;
		break;

	case OASLMOD:
	case OLMOD:
	case OASLDIV:
	case OLDIV:
		a = ADIVL;
		break;

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
		a = ACMPL;
		if(et == TCHAR || et == TUCHAR)
			a = ACMPB;
		if(et == TSHORT || et == TUSHORT)
			a = ACMPW;
		gins(a, f, t);
		switch(o) {
		case OEQ:	a = AJEQ; break;
		case ONE:	a = AJNE; break;
		case OLT:	a = AJLT; break;
		case OLE:	a = AJLE; break;
		case OGE:	a = AJGE; break;
		case OGT:	a = AJGT; break;
		case OLO:	a = AJCS; break;
		case OLS:	a = AJLS; break;
		case OHS:	a = AJCC; break;
		case OHI:	a = AJHI; break;
		}
		gins(a, Z, Z);
		return;
	}
	if(a == AGOK)
		diag(Z, &#34;bad in gopcode %O&#34;, o);
	gins(a, f, t);
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
		a = AJMP;
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
	p-&gt;from.type = D_EXTERN;
	p-&gt;from.sym = s;
	p-&gt;from.scale = textflag;
	textflag = 0;

	if(s-&gt;class == CSTATIC)
		p-&gt;from.type = D_STATIC;
	naddr(n, &amp;p-&gt;to);
	if(a == ADATA || a == AGLOBL)
		pc--;
}

int
sconst(Node *n)
{
	int32 v;

	if(n-&gt;op == OCONST &amp;&amp; !typefd[n-&gt;type-&gt;etype]) {
		v = n-&gt;vconst;
		if(v &gt;= -32766L &amp;&amp; v &lt; 32766L)
			return 1;
	}
	return 0;
}

int32
exreg(Type *t)
{
	int32 o;

	if(typechlp[t-&gt;etype]){
		if(exregoffset &gt;= 32)
			return 0;
		o = exregoffset;
		exregoffset += 4;
		return o+1;	// +1 to avoid 0 == failure; naddr case OEXREG will -1.
	}

	USED(t);
	return 0;
}

schar	ewidth[NTYPE] =
{
	-1,		/*[TXXX]*/
	SZ_CHAR,	/*[TCHAR]*/
	SZ_CHAR,	/*[TUCHAR]*/
	SZ_SHORT,	/*[TSHORT]*/
	SZ_SHORT,	/*[TUSHORT]*/
	SZ_INT,		/*[TINT]*/
	SZ_INT,		/*[TUINT]*/
	SZ_LONG,	/*[TLONG]*/
	SZ_LONG,	/*[TULONG]*/
	SZ_VLONG,	/*[TVLONG]*/
	SZ_VLONG,	/*[TUVLONG]*/
	SZ_FLOAT,	/*[TFLOAT]*/
	SZ_DOUBLE,	/*[TDOUBLE]*/
	SZ_IND,		/*[TIND]*/
	0,		/*[TFUNC]*/
	-1,		/*[TARRAY]*/
	0,		/*[TVOID]*/
	-1,		/*[TSTRUCT]*/
	-1,		/*[TUNION]*/
	SZ_INT,		/*[TENUM]*/
};
int32	ncast[NTYPE] =
{
	0,				/*[TXXX]*/
	BCHAR|BUCHAR,			/*[TCHAR]*/
	BCHAR|BUCHAR,			/*[TUCHAR]*/
	BSHORT|BUSHORT,			/*[TSHORT]*/
	BSHORT|BUSHORT,			/*[TUSHORT]*/
	BINT|BUINT|BLONG|BULONG|BIND,	/*[TINT]*/
	BINT|BUINT|BLONG|BULONG|BIND,	/*[TUINT]*/
	BINT|BUINT|BLONG|BULONG|BIND,	/*[TLONG]*/
	BINT|BUINT|BLONG|BULONG|BIND,	/*[TULONG]*/
	BVLONG|BUVLONG,			/*[TVLONG]*/
	BVLONG|BUVLONG,			/*[TUVLONG]*/
	BFLOAT,				/*[TFLOAT]*/
	BDOUBLE,			/*[TDOUBLE]*/
	BLONG|BULONG|BIND,		/*[TIND]*/
	0,				/*[TFUNC]*/
	0,				/*[TARRAY]*/
	0,				/*[TVOID]*/
	BSTRUCT,			/*[TSTRUCT]*/
	BUNION,				/*[TUNION]*/
	0,				/*[TENUM]*/
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
