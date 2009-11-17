<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/6c/txt.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/6c/txt.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/6c/txt.c
// http://code.google.com/p/inferno-os/source/browse/utils/6c/txt.c
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

	thechar = &#39;6&#39;;
	thestring = &#34;amd64&#34;;
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
	tfield = types[TINT];

	typeword = typechlvp;
	typecmplx = typesu;

	/* TO DO */
	memmove(typechlpv, typechlp, sizeof(typechlpv));
	typechlpv[TVLONG] = 1;
	typechlpv[TUVLONG] = 1;

	zprog.link = P;
	zprog.as = AGOK;
	zprog.from.type = D_NONE;
	zprog.from.index = D_NONE;
	zprog.from.scale = 0;
	zprog.to = zprog.from;

	lregnode.op = OREGISTER;
	lregnode.class = CEXREG;
	lregnode.reg = REGTMP;
	lregnode.complex = 0;
	lregnode.addable = 11;
	lregnode.type = types[TLONG];

	qregnode = lregnode;
	qregnode.type = types[TVLONG];

	constnode.op = OCONST;
	constnode.class = CXXX;
	constnode.complex = 0;
	constnode.addable = 20;
	constnode.type = types[TLONG];

	vconstnode = constnode;
	vconstnode.type = types[TVLONG];

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

	if(0)
		com64init();

	for(i=0; i&lt;nelem(reg); i++) {
		reg[i] = 1;
		if(i &gt;= D_AX &amp;&amp; i &lt;= D_R15 &amp;&amp; i != D_SP)
			reg[i] = 0;
		if(i &gt;= D_X0 &amp;&amp; i &lt;= D_X7)
			reg[i] = 0;
	}
}

void
gclean(void)
{
	int i;
	Sym *s;

	reg[D_SP]--;
	for(i=D_AX; i&lt;=D_R15; i++)
		if(reg[i])
			diag(Z, &#34;reg %R left allocated&#34;, i);
	for(i=D_X0; i&lt;=D_X7; i++)
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

int
nareg(void)
{
	int i, n;

	n = 0;
	for(i=D_AX; i&lt;=D_R15; i++)
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
	if(typesu[n-&gt;type-&gt;etype]) {
		regaalloc(tn2, n);
		if(n-&gt;complex &gt;= FNX) {
			sugen(*fnxp, tn2, n-&gt;type-&gt;width);
			(*fnxp)++;
		} else
			sugen(n, tn2, n-&gt;type-&gt;width);
		return;
	}
	if(REGARG &gt;= 0 &amp;&amp; curarg == 0 &amp;&amp; typechlpv[n-&gt;type-&gt;etype]) {
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
nodgconst(vlong v, Type *t)
{
	if(!typev[t-&gt;etype])
		return nodconst((int32)v);
	vconstnode.vconst = v;
	return &amp;vconstnode;
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
	int et;

	*n = qregnode;
	n-&gt;reg = r;
	if(nn != Z){
		et = nn-&gt;type-&gt;etype;
		if(!typefd[et] &amp;&amp; nn-&gt;type-&gt;width &lt;= SZ_LONG &amp;&amp; 0)
			n-&gt;type = typeu[et]? types[TUINT]: types[TINT];
		else
			n-&gt;type = nn-&gt;type;
//print(&#34;nodreg %s [%s]\n&#34;, tnames[et], tnames[n-&gt;type-&gt;etype]);
		n-&gt;lineno = nn-&gt;lineno;
	}
	if(reg[r] == 0)
		return 0;
	if(nn != Z) {
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
	case TVLONG:
	case TUVLONG:
	case TIND:
		if(o != Z &amp;&amp; o-&gt;op == OREGISTER) {
			i = o-&gt;reg;
			if(i &gt;= D_AX &amp;&amp; i &lt;= D_R15)
				goto out;
		}
		for(i=D_AX; i&lt;=D_R15; i++)
			if(reg[i] == 0)
				goto out;
		diag(tn, &#34;out of fixed registers&#34;);
		goto err;

	case TFLOAT:
	case TDOUBLE:
		if(o != Z &amp;&amp; o-&gt;op == OREGISTER) {
			i = o-&gt;reg;
			if(i &gt;= D_X0 &amp;&amp; i &lt;= D_X7)
				goto out;
		}
		for(i=D_X0; i&lt;=D_X7; i++)
			if(reg[i] == 0)
				goto out;
		diag(tn, &#34;out of float registers&#34;);
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
	if(REGARG &lt; 0)
		diag(n, &#34;regaalloc1 and REGARG&lt;0&#34;);
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


	case OIND:
		naddr(n-&gt;left, a);
		if(a-&gt;type &gt;= D_AX &amp;&amp; a-&gt;type &lt;= D_R15)
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
		if(a-&gt;type &gt;= D_AX &amp;&amp; a-&gt;type &lt;= D_R15)
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
		if(typev[n-&gt;type-&gt;etype] || n-&gt;type-&gt;etype == TIND)
			a-&gt;offset = n-&gt;vconst;
		else
			a-&gt;offset = convvtox(n-&gt;vconst, typeu[n-&gt;type-&gt;etype]? TULONG: TLONG);
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

void
gcmp(int op, Node *n, vlong val)
{
	Node *cn, nod;

	cn = nodgconst(val, n-&gt;type);
	if(!immconst(cn)){
		regalloc(&amp;nod, n, Z);
		gmove(cn, &amp;nod);
		gopcode(op, n-&gt;type, n, &amp;nod);
		regfree(&amp;nod);
	}else
		gopcode(op, n-&gt;type, n, cn);
}

#define	CASE(a,b)	((a&lt;&lt;8)|(b&lt;&lt;0))

void
gmove(Node *f, Node *t)
{
	int ft, tt, t64, a;
	Node nod, nod1, nod2, nod3;
	Prog *p1, *p2;

	ft = f-&gt;type-&gt;etype;
	tt = t-&gt;type-&gt;etype;
	t64 = tt == TVLONG || tt == TUVLONG || tt == TIND;
	if(debug[&#39;M&#39;])
		print(&#34;gop: %O %O[%s],%O[%s]\n&#34;, OAS,
			f-&gt;op, tnames[ft], t-&gt;op, tnames[tt]);
	if(typefd[ft] &amp;&amp; f-&gt;op == OCONST) {
		/* TO DO: pick up special constants, possibly preloaded */
		if(f-&gt;fconst == 0.0){
			regalloc(&amp;nod, t, t);
			gins(AXORPD, &amp;nod, &amp;nod);
			gmove(&amp;nod, t);
			regfree(&amp;nod);
			return;
		}
	}
/*
 * load
 */
	if(ft == TVLONG || ft == TUVLONG)
	if(f-&gt;op == OCONST)
	if(f-&gt;vconst &gt; 0x7fffffffLL || f-&gt;vconst &lt; -0x7fffffffLL)
	if(t-&gt;op != OREGISTER) {
		regalloc(&amp;nod, f, Z);
		gmove(f, &amp;nod);
		gmove(&amp;nod, t);
		regfree(&amp;nod);
		return;
	}

	if(f-&gt;op == ONAME || f-&gt;op == OINDREG ||
	   f-&gt;op == OIND || f-&gt;op == OINDEX)
	switch(ft) {
	case TCHAR:
		a = AMOVBLSX;
		if(t64)
			a = AMOVBQSX;
		goto ld;
	case TUCHAR:
		a = AMOVBLZX;
		if(t64)
			a = AMOVBQZX;
		goto ld;
	case TSHORT:
		a = AMOVWLSX;
		if(t64)
			a = AMOVWQSX;
		goto ld;
	case TUSHORT:
		a = AMOVWLZX;
		if(t64)
			a = AMOVWQZX;
		goto ld;
	case TINT:
	case TLONG:
		if(typefd[tt]) {
			regalloc(&amp;nod, t, t);
			if(tt == TDOUBLE)
				a = ACVTSL2SD;
			else
				a = ACVTSL2SS;
			gins(a, f, &amp;nod);
			gmove(&amp;nod, t);
			regfree(&amp;nod);
			return;
		}
		a = AMOVL;
		if(t64)
			a = AMOVLQSX;
		goto ld;
	case TUINT:
	case TULONG:
		a = AMOVL;
		if(t64)
			a = AMOVLQZX;	/* could probably use plain MOVL */
		goto ld;
	case TVLONG:
		if(typefd[tt]) {
			regalloc(&amp;nod, t, t);
			if(tt == TDOUBLE)
				a = ACVTSQ2SD;
			else
				a = ACVTSQ2SS;
			gins(a, f, &amp;nod);
			gmove(&amp;nod, t);
			regfree(&amp;nod);
			return;
		}
	case TUVLONG:
		a = AMOVQ;
		goto ld;
	case TIND:
		a = AMOVQ;

	ld:
		regalloc(&amp;nod, f, t);
		nod.type = t64? types[TVLONG]: types[TINT];
		gins(a, f, &amp;nod);
		gmove(&amp;nod, t);
		regfree(&amp;nod);
		return;

	case TFLOAT:
		a = AMOVSS;
		goto fld;
	case TDOUBLE:
		a = AMOVSD;
	fld:
		regalloc(&amp;nod, f, t);
		if(tt != TDOUBLE &amp;&amp; tt != TFLOAT){	/* TO DO: why is this here */
			prtree(f, &#34;odd tree&#34;);
			nod.type = t64? types[TVLONG]: types[TINT];
		}
		gins(a, f, &amp;nod);
		gmove(&amp;nod, t);
		regfree(&amp;nod);
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
		a = AMOVL;	goto st;
	case TVLONG:
	case TUVLONG:
	case TIND:
		a = AMOVQ;	goto st;

	st:
		if(f-&gt;op == OCONST) {
			gins(a, f, t);
			return;
		}
	fst:
		regalloc(&amp;nod, t, f);
		gmove(f, &amp;nod);
		gins(a, &amp;nod, t);
		regfree(&amp;nod);
		return;

	case TFLOAT:
		a = AMOVSS;
		goto fst;
	case TDOUBLE:
		a = AMOVSD;
		goto fst;
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

	case CASE(	TUINT,	TIND):
	case CASE(	TVLONG,	TUINT):
	case CASE(	TVLONG,	TULONG):
	case CASE(	TUVLONG, TUINT):
	case CASE(	TUVLONG, TULONG):
 *****/
		a = AMOVL;
		break;

	case CASE(	TVLONG,	TCHAR):
	case	CASE(	TVLONG,	TSHORT):
	case CASE(	TVLONG,	TINT):
	case CASE(	TVLONG,	TLONG):
	case CASE(	TUVLONG, TCHAR):
	case	CASE(	TUVLONG, TSHORT):
	case CASE(	TUVLONG, TINT):
	case CASE(	TUVLONG, TLONG):
	case CASE(	TINT,	TVLONG):
	case CASE(	TINT,	TUVLONG):
	case CASE(	TLONG,	TVLONG):
	case CASE(	TINT,	TIND):
	case CASE(	TLONG,	TIND):
		a = AMOVLQSX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= (uvlong)0xffffffffU;
			if(f-&gt;vconst &amp; 0x80000000)
				f-&gt;vconst |= (vlong)0xffffffff &lt;&lt; 32;
			a = AMOVQ;
		}
		break;

	case CASE(	TUINT,	TIND):
	case CASE(	TUINT,	TVLONG):
	case CASE(	TUINT,	TUVLONG):
	case CASE(	TULONG,	TVLONG):
	case CASE(	TULONG,	TUVLONG):
	case CASE(	TULONG,	TIND):
		a = AMOVL;	/* same effect as AMOVLQZX */
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= (uvlong)0xffffffffU;
			a = AMOVQ;
		}
		break;

	case CASE(	TIND,	TVLONG):
	case CASE(	TVLONG,	TVLONG):
	case CASE(	TUVLONG,	TVLONG):
	case CASE(	TVLONG,	TUVLONG):
	case CASE(	TUVLONG,	TUVLONG):
	case CASE(	TIND,	TUVLONG):
	case CASE(	TVLONG,	TIND):
	case CASE(	TUVLONG,	TIND):
	case CASE(	TIND,	TIND):
		a = AMOVQ;
		break;

	case CASE(	TSHORT,	TINT):
	case CASE(	TSHORT,	TUINT):
	case CASE(	TSHORT,	TLONG):
	case CASE(	TSHORT,	TULONG):
		a = AMOVWLSX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= 0xffff;
			if(f-&gt;vconst &amp; 0x8000)
				f-&gt;vconst |= 0xffff0000;
			a = AMOVL;
		}
		break;

	case CASE(	TSHORT,	TVLONG):
	case CASE(	TSHORT,	TUVLONG):
	case CASE(	TSHORT,	TIND):
		a = AMOVWQSX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= 0xffff;
			if(f-&gt;vconst &amp; 0x8000){
				f-&gt;vconst |= 0xffff0000;
				f-&gt;vconst |= (vlong)~0 &lt;&lt; 32;
			}
			a = AMOVL;
		}
		break;

	case CASE(	TUSHORT,TINT):
	case CASE(	TUSHORT,TUINT):
	case CASE(	TUSHORT,TLONG):
	case CASE(	TUSHORT,TULONG):
		a = AMOVWLZX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= 0xffff;
			a = AMOVL;
		}
		break;

	case CASE(	TUSHORT,TVLONG):
	case CASE(	TUSHORT,TUVLONG):
	case CASE(	TUSHORT,TIND):
		a = AMOVWQZX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= 0xffff;
			a = AMOVL;	/* MOVL also zero-extends to 64 bits */
		}
		break;

	case CASE(	TCHAR,	TSHORT):
	case CASE(	TCHAR,	TUSHORT):
	case CASE(	TCHAR,	TINT):
	case CASE(	TCHAR,	TUINT):
	case CASE(	TCHAR,	TLONG):
	case CASE(	TCHAR,	TULONG):
		a = AMOVBLSX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= 0xff;
			if(f-&gt;vconst &amp; 0x80)
				f-&gt;vconst |= 0xffffff00;
			a = AMOVL;
		}
		break;

	case CASE(	TCHAR,	TVLONG):
	case CASE(	TCHAR,	TUVLONG):
	case CASE(	TCHAR,	TIND):
		a = AMOVBQSX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= 0xff;
			if(f-&gt;vconst &amp; 0x80){
				f-&gt;vconst |= 0xffffff00;
				f-&gt;vconst |= (vlong)~0 &lt;&lt; 32;
			}
			a = AMOVQ;
		}
		break;

	case CASE(	TUCHAR,	TSHORT):
	case CASE(	TUCHAR,	TUSHORT):
	case CASE(	TUCHAR,	TINT):
	case CASE(	TUCHAR,	TUINT):
	case CASE(	TUCHAR,	TLONG):
	case CASE(	TUCHAR,	TULONG):
		a = AMOVBLZX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= 0xff;
			a = AMOVL;
		}
		break;

	case CASE(	TUCHAR,	TVLONG):
	case CASE(	TUCHAR,	TUVLONG):
	case CASE(	TUCHAR,	TIND):
		a = AMOVBQZX;
		if(f-&gt;op == OCONST) {
			f-&gt;vconst &amp;= 0xff;
			a = AMOVL;	/* zero-extends to 64-bits */
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
	case CASE(	TFLOAT,	TVLONG):
	case CASE(	TFLOAT,	TUVLONG):
	case CASE(	TFLOAT,	TIND):

	case CASE(	TDOUBLE,TCHAR):
	case CASE(	TDOUBLE,TUCHAR):
	case CASE(	TDOUBLE,TSHORT):
	case CASE(	TDOUBLE,TUSHORT):
	case CASE(	TDOUBLE,TINT):
	case CASE(	TDOUBLE,TUINT):
	case CASE(	TDOUBLE,TLONG):
	case CASE(	TDOUBLE,TULONG):
	case CASE(	TDOUBLE,TVLONG):
	case CASE(	TDOUBLE,TUVLONG):
	case CASE(	TDOUBLE,TIND):
		regalloc(&amp;nod, t, Z);
		if(ewidth[tt] == SZ_VLONG || typeu[tt] &amp;&amp; ewidth[tt] == SZ_INT){
			if(ft == TFLOAT)
				a = ACVTTSS2SQ;
			else
				a = ACVTTSD2SQ;
		}else{
			if(ft == TFLOAT)
				a = ACVTTSS2SL;
			else
				a = ACVTTSD2SL;
		}
		gins(a, f, &amp;nod);
		gmove(&amp;nod, t);
		regfree(&amp;nod);
		return;

/*
 * uvlong to float
 */
	case CASE(	TUVLONG,	TDOUBLE):
	case CASE(	TUVLONG,	TFLOAT):
		a = ACVTSQ2SS;
		if(tt == TDOUBLE)
			a = ACVTSQ2SD;
		regalloc(&amp;nod, f, f);
		gmove(f, &amp;nod);
		regalloc(&amp;nod1, t, t);
		gins(ACMPQ, &amp;nod, nodconst(0));
		gins(AJLT, Z, Z);
		p1 = p;
		gins(a, &amp;nod, &amp;nod1);
		gins(AJMP, Z, Z);
		p2 = p;
		patch(p1, pc);
		regalloc(&amp;nod2, f, Z);
		regalloc(&amp;nod3, f, Z);
		gmove(&amp;nod, &amp;nod2);
		gins(ASHRQ, nodconst(1), &amp;nod2);
		gmove(&amp;nod, &amp;nod3);
		gins(AANDL, nodconst(1), &amp;nod3);
		gins(AORQ, &amp;nod3, &amp;nod2);
		gins(a, &amp;nod2, &amp;nod1);
		gins(tt == TDOUBLE? AADDSD: AADDSS, &amp;nod1, &amp;nod1);
		regfree(&amp;nod2);
		regfree(&amp;nod3);
		patch(p2, pc);
		regfree(&amp;nod);
		regfree(&amp;nod1);
		return;

	case CASE(	TULONG,	TDOUBLE):
	case CASE(	TUINT,	TDOUBLE):
	case CASE(	TULONG,	TFLOAT):
	case CASE(	TUINT,	TFLOAT):
		a = ACVTSQ2SS;
		if(tt == TDOUBLE)
			a = ACVTSQ2SD;
		regalloc(&amp;nod, f, f);
		gins(AMOVLQZX, f, &amp;nod);
		regalloc(&amp;nod1, t, t);
		gins(a, &amp;nod, &amp;nod1);
		gmove(&amp;nod1, t);
		regfree(&amp;nod);
		regfree(&amp;nod1);
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
	case	CASE(	TVLONG,	TFLOAT):
	case CASE(	TIND,	TFLOAT):

	case CASE(	TCHAR,	TDOUBLE):
	case CASE(	TUCHAR,	TDOUBLE):
	case CASE(	TSHORT,	TDOUBLE):
	case CASE(	TUSHORT,TDOUBLE):
	case CASE(	TINT,	TDOUBLE):
	case CASE(	TLONG,	TDOUBLE):
	case CASE(	TVLONG,	TDOUBLE):
	case CASE(	TIND,	TDOUBLE):
		regalloc(&amp;nod, t, t);
		if(ewidth[ft] == SZ_VLONG){
			if(tt == TFLOAT)
				a = ACVTSQ2SS;
			else
				a = ACVTSQ2SD;
		}else{
			if(tt == TFLOAT)
				a = ACVTSL2SS;
			else
				a = ACVTSL2SD;
		}
		gins(a, f, &amp;nod);
		gmove(&amp;nod, t);
		regfree(&amp;nod);
		return;

/*
 * float to float
 */
	case CASE(	TFLOAT,	TFLOAT):
		a = AMOVSS;
		break;
	case CASE(	TDOUBLE,TFLOAT):
		a = ACVTSD2SS;
		break;
	case CASE(	TFLOAT,	TDOUBLE):
		a = ACVTSS2SD;
		break;
	case CASE(	TDOUBLE,TDOUBLE):
		a = AMOVSD;
		break;
	}
	if(a == AMOVQ || a == AMOVSD || a == AMOVSS || a == AMOVL &amp;&amp; ewidth[ft] == ewidth[tt])	/* TO DO: check AMOVL */
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

	regalloc(&amp;nod, &amp;qregnode, Z);
	v = constnode.vconst;
	cgen(n-&gt;right, &amp;nod);
	idx.ptr = D_NONE;
	if(n-&gt;left-&gt;op == OCONST)
		idx.ptr = D_CONST;
	else if(n-&gt;left-&gt;op == OREGISTER)
		idx.ptr = n-&gt;left-&gt;reg;
	else if(n-&gt;left-&gt;op != OADDR) {
		reg[D_BP]++;	// cant be used as a base
		regalloc(&amp;nod1, &amp;qregnode, Z);
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
gopcode(int o, Type *ty, Node *f, Node *t)
{
	int a, et;

	et = TLONG;
	if(ty != T)
		et = ty-&gt;etype;
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
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = ANOTQ;
		break;

	case ONEG:
		a = ANEGL;
		if(et == TCHAR || et == TUCHAR)
			a = ANEGB;
		if(et == TSHORT || et == TUSHORT)
			a = ANEGW;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = ANEGQ;
		break;

	case OADDR:
		a = ALEAQ;
		break;

	case OASADD:
	case OADD:
		a = AADDL;
		if(et == TCHAR || et == TUCHAR)
			a = AADDB;
		if(et == TSHORT || et == TUSHORT)
			a = AADDW;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = AADDQ;
		if(et == TFLOAT)
			a = AADDSS;
		if(et == TDOUBLE)
			a = AADDSD;
		break;

	case OASSUB:
	case OSUB:
		a = ASUBL;
		if(et == TCHAR || et == TUCHAR)
			a = ASUBB;
		if(et == TSHORT || et == TUSHORT)
			a = ASUBW;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = ASUBQ;
		if(et == TFLOAT)
			a = ASUBSS;
		if(et == TDOUBLE)
			a = ASUBSD;
		break;

	case OASOR:
	case OOR:
		a = AORL;
		if(et == TCHAR || et == TUCHAR)
			a = AORB;
		if(et == TSHORT || et == TUSHORT)
			a = AORW;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = AORQ;
		break;

	case OASAND:
	case OAND:
		a = AANDL;
		if(et == TCHAR || et == TUCHAR)
			a = AANDB;
		if(et == TSHORT || et == TUSHORT)
			a = AANDW;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = AANDQ;
		break;

	case OASXOR:
	case OXOR:
		a = AXORL;
		if(et == TCHAR || et == TUCHAR)
			a = AXORB;
		if(et == TSHORT || et == TUSHORT)
			a = AXORW;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = AXORQ;
		break;

	case OASLSHR:
	case OLSHR:
		a = ASHRL;
		if(et == TCHAR || et == TUCHAR)
			a = ASHRB;
		if(et == TSHORT || et == TUSHORT)
			a = ASHRW;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = ASHRQ;
		break;

	case OASASHR:
	case OASHR:
		a = ASARL;
		if(et == TCHAR || et == TUCHAR)
			a = ASARB;
		if(et == TSHORT || et == TUSHORT)
			a = ASARW;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = ASARQ;
		break;

	case OASASHL:
	case OASHL:
		a = ASALL;
		if(et == TCHAR || et == TUCHAR)
			a = ASALB;
		if(et == TSHORT || et == TUSHORT)
			a = ASALW;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = ASALQ;
		break;

	case OFUNC:
		a = ACALL;
		break;

	case OASMUL:
	case OMUL:
		if(f-&gt;op == OREGISTER &amp;&amp; t != Z &amp;&amp; isreg(t, D_AX) &amp;&amp; reg[D_DX] == 0)
			t = Z;
		a = AIMULL;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = AIMULQ;
		if(et == TFLOAT)
			a = AMULSS;
		if(et == TDOUBLE)
			a = AMULSD;
		break;

	case OASMOD:
	case OMOD:
	case OASDIV:
	case ODIV:
		a = AIDIVL;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = AIDIVQ;
		if(et == TFLOAT)
			a = ADIVSS;
		if(et == TDOUBLE)
			a = ADIVSD;
		break;

	case OASLMUL:
	case OLMUL:
		a = AMULL;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = AMULQ;
		break;

	case OASLMOD:
	case OLMOD:
	case OASLDIV:
	case OLDIV:
		a = ADIVL;
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = ADIVQ;
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
		if(et == TVLONG || et == TUVLONG || et == TIND)
			a = ACMPQ;
		if(et == TFLOAT)
			a = AUCOMISS;
		if(et == TDOUBLE)
			a = AUCOMISD;
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
	return f-&gt;op == OREGISTER &amp;&amp; t-&gt;op == OREGISTER &amp;&amp; f-&gt;reg == t-&gt;reg;
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

	if(typechlpv[t-&gt;etype]) {
		if(exregoffset &lt;= REGEXT-4)
			return 0;
		o = exregoffset;
		exregoffset--;
		return o;
	}
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
	BINT|BUINT|BLONG|BULONG,	/*[TINT]*/
	BINT|BUINT|BLONG|BULONG,	/*[TUINT]*/
	BINT|BUINT|BLONG|BULONG,	/*[TLONG]*/
	BINT|BUINT|BLONG|BULONG,	/*[TULONG]*/
	BVLONG|BUVLONG|BIND,			/*[TVLONG]*/
	BVLONG|BUVLONG|BIND,			/*[TUVLONG]*/
	BFLOAT,				/*[TFLOAT]*/
	BDOUBLE,			/*[TDOUBLE]*/
	BVLONG|BUVLONG|BIND,		/*[TIND]*/
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
