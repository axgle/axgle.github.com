<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cc/sub.c</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/cc/sub.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/cc/sub.c
// http://code.google.com/p/inferno-os/source/browse/utils/cc/sub.c
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

#include	&#34;cc.h&#34;

Node*
new(int t, Node *l, Node *r)
{
	Node *n;

	n = alloc(sizeof(*n));
	n-&gt;op = t;
	n-&gt;left = l;
	n-&gt;right = r;
	if(l &amp;&amp; t != OGOTO)
		n-&gt;lineno = l-&gt;lineno;
	else if(r)
		n-&gt;lineno = r-&gt;lineno;
	else
		n-&gt;lineno = lineno;
	newflag = 1;
	return n;
}

Node*
new1(int o, Node *l, Node *r)
{
	Node *n;

	n = new(o, l, r);
	n-&gt;lineno = nearln;
	return n;
}

void
prtree(Node *n, char *s)
{

	print(&#34; == %s ==\n&#34;, s);
	prtree1(n, 0, 0);
	print(&#34;\n&#34;);
}

void
prtree1(Node *n, int d, int f)
{
	int i;

	if(f)
	for(i=0; i&lt;d; i++)
		print(&#34;   &#34;);
	if(n == Z) {
		print(&#34;Z\n&#34;);
		return;
	}
	if(n-&gt;op == OLIST) {
		prtree1(n-&gt;left, d, 0);
		prtree1(n-&gt;right, d, 1);
		return;
	}
	d++;
	print(&#34;%O&#34;, n-&gt;op);
	i = 3;
	switch(n-&gt;op)
	{
	case ONAME:
		print(&#34; \&#34;%F\&#34;&#34;, n);
		print(&#34; %ld&#34;, n-&gt;xoffset);
		i = 0;
		break;

	case OINDREG:
		print(&#34; %ld(R%d)&#34;, n-&gt;xoffset, n-&gt;reg);
		i = 0;
		break;

	case OREGISTER:
		if(n-&gt;xoffset)
			print(&#34; %ld+R%d&#34;, n-&gt;xoffset, n-&gt;reg);
		else
			print(&#34; R%d&#34;, n-&gt;reg);
		i = 0;
		break;

	case OSTRING:
		print(&#34; \&#34;%s\&#34;&#34;, n-&gt;cstring);
		i = 0;
		break;

	case OLSTRING:
		print(&#34; \&#34;%S\&#34;&#34;, n-&gt;rstring);
		i = 0;
		break;

	case ODOT:
	case OELEM:
		print(&#34; \&#34;%F\&#34;&#34;, n);
		break;

	case OCONST:
		if(typefd[n-&gt;type-&gt;etype])
			print(&#34; \&#34;%.8e\&#34;&#34;, n-&gt;fconst);
		else
			print(&#34; \&#34;%lld\&#34;&#34;, n-&gt;vconst);
		i = 0;
		break;
	}
	if(n-&gt;addable != 0)
		print(&#34; &lt;%d&gt;&#34;, n-&gt;addable);
	if(n-&gt;type != T)
		print(&#34; %T&#34;, n-&gt;type);
	if(n-&gt;complex != 0)
		print(&#34; (%d)&#34;, n-&gt;complex);
	print(&#34; %L\n&#34;, n-&gt;lineno);
	if(i &amp; 2)
		prtree1(n-&gt;left, d, 1);
	if(i &amp; 1)
		prtree1(n-&gt;right, d, 1);
}

Type*
typ(int et, Type *d)
{
	Type *t;

	t = alloc(sizeof(*t));
	t-&gt;etype = et;
	t-&gt;link = d;
	t-&gt;down = T;
	t-&gt;sym = S;
	t-&gt;width = ewidth[et];
	t-&gt;offset = 0;
	t-&gt;shift = 0;
	t-&gt;nbits = 0;
	t-&gt;garb = 0;
	return t;
}

Type*
copytyp(Type *t)
{
	Type *nt;

	nt = typ(TXXX, T);
	*nt = *t;
	return nt;
}

Type*
garbt(Type *t, int32 b)
{
	Type *t1;

	if(b &amp; BGARB) {
		t1 = copytyp(t);
		t1-&gt;garb = simpleg(b);
		return t1;
	}
	return t;
}

int
simpleg(int32 b)
{

	b &amp;= BGARB;
	switch(b) {
	case BCONSTNT:
		return GCONSTNT;
	case BVOLATILE:
		return GVOLATILE;
	case BVOLATILE|BCONSTNT:
		return GCONSTNT|GVOLATILE;
	}
	return GXXX;
}

int
simplec(int32 b)
{

	b &amp;= BCLASS;
	switch(b) {
	case 0:
	case BREGISTER:
		return CXXX;
	case BAUTO:
	case BAUTO|BREGISTER:
		return CAUTO;
	case BEXTERN:
		return CEXTERN;
	case BEXTERN|BREGISTER:
		return CEXREG;
	case BSTATIC:
		return CSTATIC;
	case BTYPEDEF:
		return CTYPEDEF;
	case BTYPESTR:
		return CTYPESTR;
	}
	diag(Z, &#34;illegal combination of classes %Q&#34;, b);
	return CXXX;
}

Type*
simplet(int32 b)
{

	b &amp;= ~BCLASS &amp; ~BGARB;
	switch(b) {
	case BCHAR:
	case BCHAR|BSIGNED:
		return types[TCHAR];

	case BCHAR|BUNSIGNED:
		return types[TUCHAR];

	case BSHORT:
	case BSHORT|BINT:
	case BSHORT|BSIGNED:
	case BSHORT|BINT|BSIGNED:
		return types[TSHORT];

	case BUNSIGNED|BSHORT:
	case BUNSIGNED|BSHORT|BINT:
		return types[TUSHORT];

	case 0:
	case BINT:
	case BINT|BSIGNED:
	case BSIGNED:
		return types[TINT];

	case BUNSIGNED:
	case BUNSIGNED|BINT:
		return types[TUINT];

	case BLONG:
	case BLONG|BINT:
	case BLONG|BSIGNED:
	case BLONG|BINT|BSIGNED:
		return types[TLONG];

	case BUNSIGNED|BLONG:
	case BUNSIGNED|BLONG|BINT:
		return types[TULONG];

	case BVLONG|BLONG:
	case BVLONG|BLONG|BINT:
	case BVLONG|BLONG|BSIGNED:
	case BVLONG|BLONG|BINT|BSIGNED:
		return types[TVLONG];

	case BVLONG|BLONG|BUNSIGNED:
	case BVLONG|BLONG|BINT|BUNSIGNED:
		return types[TUVLONG];

	case BFLOAT:
		return types[TFLOAT];

	case BDOUBLE:
	case BDOUBLE|BLONG:
	case BFLOAT|BLONG:
		return types[TDOUBLE];

	case BVOID:
		return types[TVOID];
	}

	diag(Z, &#34;illegal combination of types %Q&#34;, b);
	return types[TINT];
}

int
stcompat(Node *n, Type *t1, Type *t2, int32 ttab[])
{
	int i;
	uint32 b;

	i = 0;
	if(t2 != T)
		i = t2-&gt;etype;
	b = 1L &lt;&lt; i;
	i = 0;
	if(t1 != T)
		i = t1-&gt;etype;
	if(b &amp; ttab[i]) {
		if(ttab == tasign)
			if(b == BSTRUCT || b == BUNION)
				if(!sametype(t1, t2))
					return 1;
		if(n-&gt;op != OCAST)
		 	if(b == BIND &amp;&amp; i == TIND)
				if(!sametype(t1, t2))
					return 1;
		return 0;
	}
	return 1;
}

int
tcompat(Node *n, Type *t1, Type *t2, int32 ttab[])
{

	if(stcompat(n, t1, t2, ttab)) {
		if(t1 == T)
			diag(n, &#34;incompatible type: \&#34;%T\&#34; for op \&#34;%O\&#34;&#34;,
				t2, n-&gt;op);
		else
			diag(n, &#34;incompatible types: \&#34;%T\&#34; and \&#34;%T\&#34; for op \&#34;%O\&#34;&#34;,
				t1, t2, n-&gt;op);
		return 1;
	}
	return 0;
}

void
makedot(Node *n, Type *t, int32 o)
{
	Node *n1, *n2;

	if(t-&gt;nbits) {
		n1 = new(OXXX, Z, Z);
		*n1 = *n;
		n-&gt;op = OBIT;
		n-&gt;left = n1;
		n-&gt;right = Z;
		n-&gt;type = t;
		n-&gt;addable = n1-&gt;left-&gt;addable;
		n = n1;
	}
	n-&gt;addable = n-&gt;left-&gt;addable;
	if(n-&gt;addable == 0) {
		n1 = new1(OCONST, Z, Z);
		n1-&gt;vconst = o;
		n1-&gt;type = types[TLONG];
		n-&gt;right = n1;
		n-&gt;type = t;
		return;
	}
	n-&gt;left-&gt;type = t;
	if(o == 0) {
		*n = *n-&gt;left;
		return;
	}
	n-&gt;type = t;
	n1 = new1(OCONST, Z, Z);
	n1-&gt;vconst = o;
	t = typ(TIND, t);
	t-&gt;width = types[TIND]-&gt;width;
	n1-&gt;type = t;

	n2 = new1(OADDR, n-&gt;left, Z);
	n2-&gt;type = t;

	n1 = new1(OADD, n1, n2);
	n1-&gt;type = t;

	n-&gt;op = OIND;
	n-&gt;left = n1;
	n-&gt;right = Z;
}

Type*
dotsearch(Sym *s, Type *t, Node *n, int32 *off)
{
	Type *t1, *xt, *rt;

	xt = T;

	/*
	 * look it up by name
	 */
	for(t1 = t; t1 != T; t1 = t1-&gt;down)
		if(t1-&gt;sym == s) {
			if(xt != T)
				goto ambig;
			xt = t1;
		}

	/*
	 * look it up by type
	 */
	if(s-&gt;class == CTYPEDEF || s-&gt;class == CTYPESTR)
		for(t1 = t; t1 != T; t1 = t1-&gt;down)
			if(t1-&gt;sym == S &amp;&amp; typesu[t1-&gt;etype])
				if(sametype(s-&gt;type, t1)) {
					if(xt != T)
						goto ambig;
					xt = t1;
				}
	if(xt != T) {
		*off = xt-&gt;offset;
		return xt;
	}

	/*
	 * look it up in unnamed substructures
	 */
	for(t1 = t; t1 != T; t1 = t1-&gt;down)
		if(t1-&gt;sym == S &amp;&amp; typesu[t1-&gt;etype]){
			rt = dotsearch(s, t1-&gt;link, n, off);
			if(rt != T) {
				if(xt != T)
					goto ambig;
				xt = rt;
				*off += t1-&gt;offset;
			}
		}
	return xt;

ambig:
	diag(n, &#34;ambiguous structure element: %s&#34;, s-&gt;name);
	return xt;
}

int32
dotoffset(Type *st, Type *lt, Node *n)
{
	Type *t;
	Sym *g;
	int32 o, o1;

	o = -1;
	/*
	 * first try matching at the top level
	 * for matching tag names
	 */
	g = st-&gt;tag;
	if(g != S)
		for(t=lt-&gt;link; t!=T; t=t-&gt;down)
			if(t-&gt;sym == S)
				if(g == t-&gt;tag) {
					if(o &gt;= 0)
						goto ambig;
					o = t-&gt;offset;
				}
	if(o &gt;= 0)
		return o;

	/*
	 * second try matching at the top level
	 * for similar types
	 */
	for(t=lt-&gt;link; t!=T; t=t-&gt;down)
		if(t-&gt;sym == S)
			if(sametype(st, t)) {
				if(o &gt;= 0)
					goto ambig;
				o = t-&gt;offset;
			}
	if(o &gt;= 0)
		return o;

	/*
	 * last try matching sub-levels
	 */
	for(t=lt-&gt;link; t!=T; t=t-&gt;down)
		if(t-&gt;sym == S)
		if(typesu[t-&gt;etype]) {
			o1 = dotoffset(st, t, n);
			if(o1 &gt;= 0) {
				if(o &gt;= 0)
					goto ambig;
				o = o1 + t-&gt;offset;
			}
		}
	return o;

ambig:
	diag(n, &#34;ambiguous unnamed structure element&#34;);
	return o;
}

/*
 * look into tree for floating point constant expressions
 */
int
allfloat(Node *n, int flag)
{

	if(n != Z) {
		if(n-&gt;type-&gt;etype != TDOUBLE)
			return 1;
		switch(n-&gt;op) {
		case OCONST:
			if(flag)
				n-&gt;type = types[TFLOAT];
			return 1;
		case OADD:	/* no need to get more exotic than this */
		case OSUB:
		case OMUL:
		case ODIV:
			if(!allfloat(n-&gt;right, flag))
				break;
		case OCAST:
			if(!allfloat(n-&gt;left, flag))
				break;
			if(flag)
				n-&gt;type = types[TFLOAT];
			return 1;
		}
	}
	return 0;
}

void
constas(Node *n, Type *il, Type *ir)
{
	Type *l, *r;

	l = il;
	r = ir;

	if(l == T)
		return;
	if(l-&gt;garb &amp; GCONSTNT) {
		warn(n, &#34;assignment to a constant type (%T)&#34;, il);
		return;
	}
	if(r == T)
		return;
	for(;;) {
		if(l-&gt;etype != TIND || r-&gt;etype != TIND)
			break;
		l = l-&gt;link;
		r = r-&gt;link;
		if(l == T || r == T)
			break;
		if(r-&gt;garb &amp; GCONSTNT)
			if(!(l-&gt;garb &amp; GCONSTNT)) {
				warn(n, &#34;assignment of a constant pointer type (%T)&#34;, ir);
				break;
			}
	}
}

void
typeext1(Type *st, Node *l)
{
	if(st-&gt;etype == TFLOAT &amp;&amp; allfloat(l, 0))
		allfloat(l, 1);
}

void
typeext(Type *st, Node *l)
{
	Type *lt;
	Node *n1, *n2;
	int32 o;

	lt = l-&gt;type;
	if(lt == T)
		return;
	if(st-&gt;etype == TIND &amp;&amp; vconst(l) == 0) {
		l-&gt;type = st;
		l-&gt;vconst = 0;
		return;
	}
	typeext1(st, l);

	/*
	 * extension of C
	 * if assign of struct containing unnamed sub-struct
	 * to type of sub-struct, insert the DOT.
	 * if assign of *struct containing unnamed substruct
	 * to type of *sub-struct, insert the add-offset
	 */
	if(typesu[st-&gt;etype] &amp;&amp; typesu[lt-&gt;etype]) {
		o = dotoffset(st, lt, l);
		if(o &gt;= 0) {
			n1 = new1(OXXX, Z, Z);
			*n1 = *l;
			l-&gt;op = ODOT;
			l-&gt;left = n1;
			l-&gt;right = Z;
			makedot(l, st, o);
		}
		return;
	}
	if(st-&gt;etype == TIND &amp;&amp; typesu[st-&gt;link-&gt;etype])
	if(lt-&gt;etype == TIND &amp;&amp; typesu[lt-&gt;link-&gt;etype]) {
		o = dotoffset(st-&gt;link, lt-&gt;link, l);
		if(o &gt;= 0) {
			l-&gt;type = st;
			if(o == 0)
				return;
			n1 = new1(OXXX, Z, Z);
			*n1 = *l;
			n2 = new1(OCONST, Z, Z);
			n2-&gt;vconst = o;
			n2-&gt;type = st;
			l-&gt;op = OADD;
			l-&gt;left = n1;
			l-&gt;right = n2;
		}
		return;
	}
}

/*
 * a cast that generates no code
 * (same size move)
 */
int
nocast(Type *t1, Type *t2)
{
	int i, b;

	if(t1-&gt;nbits)
		return 0;
	i = 0;
	if(t2 != T)
		i = t2-&gt;etype;
	b = 1&lt;&lt;i;
	i = 0;
	if(t1 != T)
		i = t1-&gt;etype;
	if(b &amp; ncast[i])
		return 1;
	return 0;
}

/*
 * a cast that has a noop semantic
 * (small to large, convert)
 */
int
nilcast(Type *t1, Type *t2)
{
	int et1, et2;

	if(t1 == T)
		return 0;
	if(t1-&gt;nbits)
		return 0;
	if(t2 == T)
		return 0;
	et1 = t1-&gt;etype;
	et2 = t2-&gt;etype;
	if(et1 == et2)
		return 1;
	if(typefd[et1] &amp;&amp; typefd[et2]) {
		if(ewidth[et1] &lt; ewidth[et2])
			return 1;
		return 0;
	}
	if(typechlp[et1] &amp;&amp; typechlp[et2]) {
		if(ewidth[et1] &lt; ewidth[et2])
			return 1;
		return 0;
	}
	return 0;
}

/*
 * &#34;the usual arithmetic conversions are performed&#34;
 */
void
arith(Node *n, int f)
{
	Type *t1, *t2;
	int i, j, k;
	Node *n1;
	int32 w;

	t1 = n-&gt;left-&gt;type;
	if(n-&gt;right == Z)
		t2 = t1;
	else
		t2 = n-&gt;right-&gt;type;
	i = TXXX;
	if(t1 != T)
		i = t1-&gt;etype;
	j = TXXX;
	if(t2 != T)
		j = t2-&gt;etype;
	k = tab[i][j];
	if(k == TIND) {
		if(i == TIND)
			n-&gt;type = t1;
		else
		if(j == TIND)
			n-&gt;type = t2;
	} else {
		/* convert up to at least int */
		if(f == 1)
		while(k &lt; TINT)
			k += 2;
		n-&gt;type = types[k];
	}
	if(n-&gt;op == OSUB)
	if(i == TIND &amp;&amp; j == TIND) {
		w = n-&gt;right-&gt;type-&gt;link-&gt;width;
		if(w &lt; 1 || n-&gt;left-&gt;type-&gt;link == T || n-&gt;left-&gt;type-&gt;link-&gt;width &lt; 1)
			goto bad;
		n-&gt;type = types[ewidth[TIND] &lt;= ewidth[TLONG]? TLONG: TVLONG];
		if(0 &amp;&amp; ewidth[TIND] &gt; ewidth[TLONG]){
			n1 = new1(OXXX, Z, Z);
			*n1 = *n;
			n-&gt;op = OCAST;
			n-&gt;left = n1;
			n-&gt;right = Z;
			n-&gt;type = types[TLONG];
		}
		if(w &gt; 1) {
			n1 = new1(OXXX, Z, Z);
			*n1 = *n;
			n-&gt;op = ODIV;
			n-&gt;left = n1;
			n1 = new1(OCONST, Z, Z);
			n1-&gt;vconst = w;
			n1-&gt;type = n-&gt;type;
			n-&gt;right = n1;
			w = vlog(n1);
			if(w &gt;= 0) {
				n-&gt;op = OASHR;
				n1-&gt;vconst = w;
			}
		}
		return;
	}
	if(!sametype(n-&gt;type, n-&gt;left-&gt;type)) {
		n-&gt;left = new1(OCAST, n-&gt;left, Z);
		n-&gt;left-&gt;type = n-&gt;type;
		if(n-&gt;type-&gt;etype == TIND) {
			w = n-&gt;type-&gt;link-&gt;width;
			if(w &lt; 1) {
				snap(n-&gt;type-&gt;link);
				w = n-&gt;type-&gt;link-&gt;width;
				if(w &lt; 1)
					goto bad;
			}
			if(w &gt; 1) {
				n1 = new1(OCONST, Z, Z);
				n1-&gt;vconst = w;
				n1-&gt;type = n-&gt;type;
				n-&gt;left = new1(OMUL, n-&gt;left, n1);
				n-&gt;left-&gt;type = n-&gt;type;
			}
		}
	}
	if(n-&gt;right != Z)
	if(!sametype(n-&gt;type, n-&gt;right-&gt;type)) {
		n-&gt;right = new1(OCAST, n-&gt;right, Z);
		n-&gt;right-&gt;type = n-&gt;type;
		if(n-&gt;type-&gt;etype == TIND) {
			w = n-&gt;type-&gt;link-&gt;width;
			if(w &lt; 1) {
				snap(n-&gt;type-&gt;link);
				w = n-&gt;type-&gt;link-&gt;width;
				if(w &lt; 1)
					goto bad;
			}
			if(w != 1) {
				n1 = new1(OCONST, Z, Z);
				n1-&gt;vconst = w;
				n1-&gt;type = n-&gt;type;
				n-&gt;right = new1(OMUL, n-&gt;right, n1);
				n-&gt;right-&gt;type = n-&gt;type;
			}
		}
	}
	return;
bad:
	diag(n, &#34;pointer addition not fully declared: %T&#34;, n-&gt;type-&gt;link);
}

/*
 * try to rewrite shift &amp; mask
 */
void
simplifyshift(Node *n)
{
	uint32 c3;
	int o, s1, s2, c1, c2;

	if(!typechlp[n-&gt;type-&gt;etype])
		return;
	switch(n-&gt;op) {
	default:
		return;
	case OASHL:
		s1 = 0;
		break;
	case OLSHR:
		s1 = 1;
		break;
	case OASHR:
		s1 = 2;
		break;
	}
	if(n-&gt;right-&gt;op != OCONST)
		return;
	if(n-&gt;left-&gt;op != OAND)
		return;
	if(n-&gt;left-&gt;right-&gt;op != OCONST)
		return;
	switch(n-&gt;left-&gt;left-&gt;op) {
	default:
		return;
	case OASHL:
		s2 = 0;
		break;
	case OLSHR:
		s2 = 1;
		break;
	case OASHR:
		s2 = 2;
		break;
	}
	if(n-&gt;left-&gt;left-&gt;right-&gt;op != OCONST)
		return;

	c1 = n-&gt;right-&gt;vconst;
	c2 = n-&gt;left-&gt;left-&gt;right-&gt;vconst;
	c3 = n-&gt;left-&gt;right-&gt;vconst;

/*
	if(debug[&#39;h&#39;])
		print(&#34;%.3o %ld %ld %d #%.lux\n&#34;,
			(s1&lt;&lt;3)|s2, c1, c2, topbit(c3), c3);
*/

	o = n-&gt;op;
	switch((s1&lt;&lt;3)|s2) {
	case 000:	/* (((e &lt;&lt;u c2) &amp; c3) &lt;&lt;u c1) */
		c3 &gt;&gt;= c2;
		c1 += c2;
		if(c1 &gt;= 32)
			break;
		goto rewrite1;

	case 002:	/* (((e &gt;&gt;s c2) &amp; c3) &lt;&lt;u c1) */
		if(topbit(c3) &gt;= (32-c2))
			break;
	case 001:	/* (((e &gt;&gt;u c2) &amp; c3) &lt;&lt;u c1) */
		if(c1 &gt; c2) {
			c3 &lt;&lt;= c2;
			c1 -= c2;
			o = OASHL;
			goto rewrite1;
		}
		c3 &lt;&lt;= c1;
		if(c1 == c2)
			goto rewrite0;
		c1 = c2-c1;
		o = OLSHR;
		goto rewrite2;

	case 022:	/* (((e &gt;&gt;s c2) &amp; c3) &gt;&gt;s c1) */
		if(c2 &lt;= 0)
			break;
	case 012:	/* (((e &gt;&gt;s c2) &amp; c3) &gt;&gt;u c1) */
		if(topbit(c3) &gt;= (32-c2))
			break;
		goto s11;
	case 021:	/* (((e &gt;&gt;u c2) &amp; c3) &gt;&gt;s c1) */
		if(topbit(c3) &gt;= 31 &amp;&amp; c2 &lt;= 0)
			break;
		goto s11;
	case 011:	/* (((e &gt;&gt;u c2) &amp; c3) &gt;&gt;u c1) */
	s11:
		c3 &lt;&lt;= c2;
		c1 += c2;
		if(c1 &gt;= 32)
			break;
		o = OLSHR;
		goto rewrite1;

	case 020:	/* (((e &lt;&lt;u c2) &amp; c3) &gt;&gt;s c1) */
		if(topbit(c3) &gt;= 31)
			break;
	case 010:	/* (((e &lt;&lt;u c2) &amp; c3) &gt;&gt;u c1) */
		c3 &gt;&gt;= c1;
		if(c1 == c2)
			goto rewrite0;
		if(c1 &gt; c2) {
			c1 -= c2;
			goto rewrite2;
		}
		c1 = c2 - c1;
		o = OASHL;
		goto rewrite2;
	}
	return;

rewrite0:	/* get rid of both shifts */
if(debug[&#39;&lt;&#39;])prtree(n, &#34;rewrite0&#34;);
	*n = *n-&gt;left;
	n-&gt;left = n-&gt;left-&gt;left;
	n-&gt;right-&gt;vconst = c3;
	return;
rewrite1:	/* get rid of lower shift */
if(debug[&#39;&lt;&#39;])prtree(n, &#34;rewrite1&#34;);
	n-&gt;left-&gt;left = n-&gt;left-&gt;left-&gt;left;
	n-&gt;left-&gt;right-&gt;vconst = c3;
	n-&gt;right-&gt;vconst = c1;
	n-&gt;op = o;
	return;
rewrite2:	/* get rid of upper shift */
if(debug[&#39;&lt;&#39;])prtree(n, &#34;rewrite2&#34;);
	*n = *n-&gt;left;
	n-&gt;right-&gt;vconst = c3;
	n-&gt;left-&gt;right-&gt;vconst = c1;
	n-&gt;left-&gt;op = o;
}

int
side(Node *n)
{

loop:
	if(n != Z)
	switch(n-&gt;op) {
	case OCAST:
	case ONOT:
	case OADDR:
	case OIND:
		n = n-&gt;left;
		goto loop;

	case OCOND:
		if(side(n-&gt;left))
			break;
		n = n-&gt;right;

	case OEQ:
	case ONE:
	case OLT:
	case OGE:
	case OGT:
	case OLE:
	case OADD:
	case OSUB:
	case OMUL:
	case OLMUL:
	case ODIV:
	case OLDIV:
	case OLSHR:
	case OASHL:
	case OASHR:
	case OAND:
	case OOR:
	case OXOR:
	case OMOD:
	case OLMOD:
	case OANDAND:
	case OOROR:
	case OCOMMA:
	case ODOT:
		if(side(n-&gt;left))
			break;
		n = n-&gt;right;
		goto loop;

	case OSIGN:
	case OSIZE:
	case OCONST:
	case OSTRING:
	case OLSTRING:
	case ONAME:
		return 0;
	}
	return 1;
}

int
vconst(Node *n)
{
	int i;

	if(n == Z)
		goto no;
	if(n-&gt;op != OCONST)
		goto no;
	if(n-&gt;type == T)
		goto no;
	switch(n-&gt;type-&gt;etype)
	{
	case TFLOAT:
	case TDOUBLE:
		i = 100;
		if(n-&gt;fconst &gt; i || n-&gt;fconst &lt; -i)
			goto no;
		i = n-&gt;fconst;
		if(i != n-&gt;fconst)
			goto no;
		return i;

	case TVLONG:
	case TUVLONG:
		i = n-&gt;vconst;
		if(i != n-&gt;vconst)
			goto no;
		return i;

	case TCHAR:
	case TUCHAR:
	case TSHORT:
	case TUSHORT:
	case TINT:
	case TUINT:
	case TLONG:
	case TULONG:
	case TIND:
		i = n-&gt;vconst;
		if(i != n-&gt;vconst)
			goto no;
		return i;
	}
no:
	return -159;	/* first uninteresting constant */
}

/*
 * return log(n) if n is a power of 2 constant
 */
int
xlog2(uvlong v)
{
	int s, i;
	uvlong m;

	s = 0;
	m = MASK(8*sizeof(uvlong));
	for(i=32; i; i&gt;&gt;=1) {
		m &gt;&gt;= i;
		if(!(v &amp; m)) {
			v &gt;&gt;= i;
			s += i;
		}
	}
	if(v == 1)
		return s;
	return -1;
}

int
vlog(Node *n)
{
	if(n-&gt;op != OCONST)
		goto bad;
	if(typefd[n-&gt;type-&gt;etype])
		goto bad;

	return xlog2(n-&gt;vconst);

bad:
	return -1;
}

int
topbit(uint32 v)
{
	int i;

	for(i = -1; v; i++)
		v &gt;&gt;= 1;
	return i;
}

/*
 * try to cast a constant down
 * rather than cast a variable up
 * example:
 *	if(c == &#39;a&#39;)
 */
void
relcon(Node *l, Node *r)
{
	vlong v;

	if(l-&gt;op != OCONST)
		return;
	if(r-&gt;op != OCAST)
		return;
	if(!nilcast(r-&gt;left-&gt;type, r-&gt;type))
		return;
	switch(r-&gt;type-&gt;etype) {
	default:
		return;
	case TCHAR:
	case TUCHAR:
	case TSHORT:
	case TUSHORT:
		v = convvtox(l-&gt;vconst, r-&gt;type-&gt;etype);
		if(v != l-&gt;vconst)
			return;
		break;
	}
	l-&gt;type = r-&gt;left-&gt;type;
	*r = *r-&gt;left;
}

int
relindex(int o)
{

	switch(o) {
	default:
		diag(Z, &#34;bad in relindex: %O&#34;, o);
	case OEQ: return 0;
	case ONE: return 1;
	case OLE: return 2;
	case OLS: return 3;
	case OLT: return 4;
	case OLO: return 5;
	case OGE: return 6;
	case OHS: return 7;
	case OGT: return 8;
	case OHI: return 9;
	}
}

Node*
invert(Node *n)
{
	Node *i;

	if(n == Z || n-&gt;op != OLIST)
		return n;
	i = n;
	for(n = n-&gt;left; n != Z; n = n-&gt;left) {
		if(n-&gt;op != OLIST)
			break;
		i-&gt;left = n-&gt;right;
		n-&gt;right = i;
		i = n;
	}
	i-&gt;left = n;
	return i;
}

int
bitno(int32 b)
{
	int i;

	for(i=0; i&lt;32; i++)
		if(b &amp; (1L&lt;&lt;i))
			return i;
	diag(Z, &#34;bad in bitno&#34;);
	return 0;
}

int32
typebitor(int32 a, int32 b)
{
	int32 c;

	c = a | b;
	if(a &amp; b)
		if((a &amp; b) == BLONG)
			c |= BVLONG;		/* long long =&gt; vlong */
		else
			warn(Z, &#34;once is enough: %Q&#34;, a &amp; b);
	return c;
}

void
diag(Node *n, char *fmt, ...)
{
	char buf[STRINGSZ];
	va_list arg;

	va_start(arg, fmt);
	vseprint(buf, buf+sizeof(buf), fmt, arg);
	va_end(arg);
	Bprint(&amp;diagbuf, &#34;%L %s\n&#34;, (n==Z)? nearln: n-&gt;lineno, buf);

	if(debug[&#39;X&#39;]){
		Bflush(&amp;diagbuf);
		abort();
	}
	if(n != Z)
	if(debug[&#39;v&#39;])
		prtree(n, &#34;diagnostic&#34;);

	nerrors++;
	if(nerrors &gt; 10) {
		Bprint(&amp;diagbuf, &#34;too many errors\n&#34;);
		errorexit();
	}
}

void
warn(Node *n, char *fmt, ...)
{
	char buf[STRINGSZ];
	va_list arg;

	if(debug[&#39;w&#39;]) {
		Bprint(&amp;diagbuf, &#34;warning: &#34;);
		va_start(arg, fmt);
		vseprint(buf, buf+sizeof(buf), fmt, arg);
		va_end(arg);
		Bprint(&amp;diagbuf, &#34;%L %s\n&#34;, (n==Z)? nearln: n-&gt;lineno, buf);

		if(n != Z)
		if(debug[&#39;v&#39;])
			prtree(n, &#34;warning&#34;);
	}
}

void
yyerror(char *fmt, ...)
{
	char buf[STRINGSZ];
	va_list arg;

	/*
	 * hack to intercept message from yaccpar
	 */
	if(strcmp(fmt, &#34;syntax error&#34;) == 0) {
		yyerror(&#34;syntax error, last name: %s&#34;, symb);
		return;
	}
	va_start(arg, fmt);
	vseprint(buf, buf+sizeof(buf), fmt, arg);
	va_end(arg);
	Bprint(&amp;diagbuf, &#34;%L %s\n&#34;, lineno, buf);
	nerrors++;
	if(nerrors &gt; 10) {
		Bprint(&amp;diagbuf, &#34;too many errors\n&#34;);
		errorexit();
	}
}

void
fatal(Node *n, char *fmt, ...)
{
	char buf[STRINGSZ];
	va_list arg;

	va_start(arg, fmt);
	vseprint(buf, buf+sizeof(buf), fmt, arg);
	va_end(arg);
	Bprint(&amp;diagbuf, &#34;%L %s\n&#34;, (n==Z)? nearln: n-&gt;lineno, buf);

	if(debug[&#39;X&#39;]){
		Bflush(&amp;diagbuf);
		abort();
	}
	if(n != Z)
	if(debug[&#39;v&#39;])
		prtree(n, &#34;diagnostic&#34;);

	nerrors++;
	errorexit();
}

uint32	thash1	= 0x2edab8c9;
uint32	thash2	= 0x1dc74fb8;
uint32	thash3	= 0x1f241331;
uint32	thash[NALLTYPES];
Init	thashinit[] =
{
	TXXX,		0x17527bbd,	0,
	TCHAR,		0x5cedd32b,	0,
	TUCHAR,		0x552c4454,	0,
	TSHORT,		0x63040b4b,	0,
	TUSHORT,	0x32a45878,	0,
	TINT,		0x4151d5bd,	0,
	TUINT,		0x5ae707d6,	0,
	TLONG,		0x5ef20f47,	0,
	TULONG,		0x36d8eb8f,	0,
	TVLONG,		0x6e5e9590,	0,
	TUVLONG,	0x75910105,	0,
	TFLOAT,		0x25fd7af1,	0,
	TDOUBLE,	0x7c40a1b2,	0,
	TIND,		0x1b832357,	0,
	TFUNC,		0x6babc9cb,	0,
	TARRAY,		0x7c50986d,	0,
	TVOID,		0x44112eff,	0,
	TSTRUCT,	0x7c2da3bf,	0,
	TUNION,		0x3eb25e98,	0,
	TENUM,		0x44b54f61,	0,
	TFILE,		0x19242ac3,	0,
	TOLD,		0x22b15988,	0,
	TDOT,		0x0204f6b3,	0,
	-1,		0,		0,
};

char*	bnames[NALIGN];
Init	bnamesinit[] =
{
	Axxx,	0,	&#34;Axxx&#34;,
	Ael1,	0,	&#34;el1&#34;,
	Ael2,	0,	&#34;el2&#34;,
	Asu2,	0,	&#34;su2&#34;,
	Aarg0,	0,	&#34;arg0&#34;,
	Aarg1,	0,	&#34;arg1&#34;,
	Aarg2,	0,	&#34;arg2&#34;,
	Aaut3,	0,	&#34;aut3&#34;,
	-1,	0,	0,
};

char*	tnames[NALLTYPES];
Init	tnamesinit[] =
{
	TXXX,		0,	&#34;TXXX&#34;,
	TCHAR,		0,	&#34;CHAR&#34;,
	TUCHAR,		0,	&#34;UCHAR&#34;,
	TSHORT,		0,	&#34;SHORT&#34;,
	TUSHORT,	0,	&#34;USHORT&#34;,
	TINT,		0,	&#34;INT&#34;,
	TUINT,		0,	&#34;UINT&#34;,
	TLONG,		0,	&#34;LONG&#34;,
	TULONG,		0,	&#34;ULONG&#34;,
	TVLONG,		0,	&#34;VLONG&#34;,
	TUVLONG,	0,	&#34;UVLONG&#34;,
	TFLOAT,		0,	&#34;FLOAT&#34;,
	TDOUBLE,	0,	&#34;DOUBLE&#34;,
	TIND,		0,	&#34;IND&#34;,
	TFUNC,		0,	&#34;FUNC&#34;,
	TARRAY,		0,	&#34;ARRAY&#34;,
	TVOID,		0,	&#34;VOID&#34;,
	TSTRUCT,	0,	&#34;STRUCT&#34;,
	TUNION,		0,	&#34;UNION&#34;,
	TENUM,		0,	&#34;ENUM&#34;,
	TFILE,		0,	&#34;FILE&#34;,
	TOLD,		0,	&#34;OLD&#34;,
	TDOT,		0,	&#34;DOT&#34;,
	-1,		0,	0,
};

char*	gnames[NGTYPES];
Init	gnamesinit[] =
{
	GXXX,			0,	&#34;GXXX&#34;,
	GCONSTNT,		0,	&#34;CONST&#34;,
	GVOLATILE,		0,	&#34;VOLATILE&#34;,
	GVOLATILE|GCONSTNT,	0,	&#34;CONST-VOLATILE&#34;,
	-1,			0,	0,
};

char*	qnames[NALLTYPES];
Init	qnamesinit[] =
{
	TXXX,		0,	&#34;TXXX&#34;,
	TCHAR,		0,	&#34;CHAR&#34;,
	TUCHAR,		0,	&#34;UCHAR&#34;,
	TSHORT,		0,	&#34;SHORT&#34;,
	TUSHORT,	0,	&#34;USHORT&#34;,
	TINT,		0,	&#34;INT&#34;,
	TUINT,		0,	&#34;UINT&#34;,
	TLONG,		0,	&#34;LONG&#34;,
	TULONG,		0,	&#34;ULONG&#34;,
	TVLONG,		0,	&#34;VLONG&#34;,
	TUVLONG,	0,	&#34;UVLONG&#34;,
	TFLOAT,		0,	&#34;FLOAT&#34;,
	TDOUBLE,	0,	&#34;DOUBLE&#34;,
	TIND,		0,	&#34;IND&#34;,
	TFUNC,		0,	&#34;FUNC&#34;,
	TARRAY,		0,	&#34;ARRAY&#34;,
	TVOID,		0,	&#34;VOID&#34;,
	TSTRUCT,	0,	&#34;STRUCT&#34;,
	TUNION,		0,	&#34;UNION&#34;,
	TENUM,		0,	&#34;ENUM&#34;,

	TAUTO,		0,	&#34;AUTO&#34;,
	TEXTERN,	0,	&#34;EXTERN&#34;,
	TSTATIC,	0,	&#34;STATIC&#34;,
	TTYPEDEF,	0,	&#34;TYPEDEF&#34;,
	TTYPESTR,	0,	&#34;TYPESTR&#34;,
	TREGISTER,	0,	&#34;REGISTER&#34;,
	TCONSTNT,	0,	&#34;CONSTNT&#34;,
	TVOLATILE,	0,	&#34;VOLATILE&#34;,
	TUNSIGNED,	0,	&#34;UNSIGNED&#34;,
	TSIGNED,	0,	&#34;SIGNED&#34;,
	TDOT,		0,	&#34;DOT&#34;,
	TFILE,		0,	&#34;FILE&#34;,
	TOLD,		0,	&#34;OLD&#34;,
	-1,		0,	0,
};
char*	cnames[NCTYPES];
Init	cnamesinit[] =
{
	CXXX,		0,	&#34;CXXX&#34;,
	CAUTO,		0,	&#34;AUTO&#34;,
	CEXTERN,	0,	&#34;EXTERN&#34;,
	CGLOBL,		0,	&#34;GLOBL&#34;,
	CSTATIC,	0,	&#34;STATIC&#34;,
	CLOCAL,		0,	&#34;LOCAL&#34;,
	CTYPEDEF,	0,	&#34;TYPEDEF&#34;,
	CTYPESTR,	0,	&#34;TYPESTR&#34;,
	CPARAM,		0,	&#34;PARAM&#34;,
	CSELEM,		0,	&#34;SELEM&#34;,
	CLABEL,		0,	&#34;LABEL&#34;,
	CEXREG,		0,	&#34;EXREG&#34;,
	-1,		0,	0,
};

char*	onames[OEND+1];
Init	onamesinit[] =
{
	OXXX,		0,	&#34;OXXX&#34;,
	OADD,		0,	&#34;ADD&#34;,
	OADDR,		0,	&#34;ADDR&#34;,
	OAND,		0,	&#34;AND&#34;,
	OANDAND,	0,	&#34;ANDAND&#34;,
	OARRAY,		0,	&#34;ARRAY&#34;,
	OAS,		0,	&#34;AS&#34;,
	OASI,		0,	&#34;ASI&#34;,
	OASADD,		0,	&#34;ASADD&#34;,
	OASAND,		0,	&#34;ASAND&#34;,
	OASASHL,	0,	&#34;ASASHL&#34;,
	OASASHR,	0,	&#34;ASASHR&#34;,
	OASDIV,		0,	&#34;ASDIV&#34;,
	OASHL,		0,	&#34;ASHL&#34;,
	OASHR,		0,	&#34;ASHR&#34;,
	OASLDIV,	0,	&#34;ASLDIV&#34;,
	OASLMOD,	0,	&#34;ASLMOD&#34;,
	OASLMUL,	0,	&#34;ASLMUL&#34;,
	OASLSHR,	0,	&#34;ASLSHR&#34;,
	OASMOD,		0,	&#34;ASMOD&#34;,
	OASMUL,		0,	&#34;ASMUL&#34;,
	OASOR,		0,	&#34;ASOR&#34;,
	OASSUB,		0,	&#34;ASSUB&#34;,
	OASXOR,		0,	&#34;ASXOR&#34;,
	OBIT,		0,	&#34;BIT&#34;,
	OBREAK,		0,	&#34;BREAK&#34;,
	OCASE,		0,	&#34;CASE&#34;,
	OCAST,		0,	&#34;CAST&#34;,
	OCOMMA,		0,	&#34;COMMA&#34;,
	OCOND,		0,	&#34;COND&#34;,
	OCONST,		0,	&#34;CONST&#34;,
	OCONTINUE,	0,	&#34;CONTINUE&#34;,
	ODIV,		0,	&#34;DIV&#34;,
	ODOT,		0,	&#34;DOT&#34;,
	ODOTDOT,	0,	&#34;DOTDOT&#34;,
	ODWHILE,	0,	&#34;DWHILE&#34;,
	OENUM,		0,	&#34;ENUM&#34;,
	OEQ,		0,	&#34;EQ&#34;,
	OEXREG,	0,	&#34;EXREG&#34;,
	OFOR,		0,	&#34;FOR&#34;,
	OFUNC,		0,	&#34;FUNC&#34;,
	OGE,		0,	&#34;GE&#34;,
	OGOTO,		0,	&#34;GOTO&#34;,
	OGT,		0,	&#34;GT&#34;,
	OHI,		0,	&#34;HI&#34;,
	OHS,		0,	&#34;HS&#34;,
	OIF,		0,	&#34;IF&#34;,
	OIND,		0,	&#34;IND&#34;,
	OINDREG,	0,	&#34;INDREG&#34;,
	OINIT,		0,	&#34;INIT&#34;,
	OLABEL,		0,	&#34;LABEL&#34;,
	OLDIV,		0,	&#34;LDIV&#34;,
	OLE,		0,	&#34;LE&#34;,
	OLIST,		0,	&#34;LIST&#34;,
	OLMOD,		0,	&#34;LMOD&#34;,
	OLMUL,		0,	&#34;LMUL&#34;,
	OLO,		0,	&#34;LO&#34;,
	OLS,		0,	&#34;LS&#34;,
	OLSHR,		0,	&#34;LSHR&#34;,
	OLT,		0,	&#34;LT&#34;,
	OMOD,		0,	&#34;MOD&#34;,
	OMUL,		0,	&#34;MUL&#34;,
	ONAME,		0,	&#34;NAME&#34;,
	ONE,		0,	&#34;NE&#34;,
	ONOT,		0,	&#34;NOT&#34;,
	OOR,		0,	&#34;OR&#34;,
	OOROR,		0,	&#34;OROR&#34;,
	OPOSTDEC,	0,	&#34;POSTDEC&#34;,
	OPOSTINC,	0,	&#34;POSTINC&#34;,
	OPREDEC,	0,	&#34;PREDEC&#34;,
	OPREINC,	0,	&#34;PREINC&#34;,
	OPROTO,		0,	&#34;PROTO&#34;,
	OREGISTER,	0,	&#34;REGISTER&#34;,
	ORETURN,	0,	&#34;RETURN&#34;,
	OSET,		0,	&#34;SET&#34;,
	OSIGN,		0,	&#34;SIGN&#34;,
	OSIZE,		0,	&#34;SIZE&#34;,
	OSTRING,	0,	&#34;STRING&#34;,
	OLSTRING,	0,	&#34;LSTRING&#34;,
	OSTRUCT,	0,	&#34;STRUCT&#34;,
	OSUB,		0,	&#34;SUB&#34;,
	OSWITCH,	0,	&#34;SWITCH&#34;,
	OUNION,		0,	&#34;UNION&#34;,
	OUSED,		0,	&#34;USED&#34;,
	OWHILE,		0,	&#34;WHILE&#34;,
	OXOR,		0,	&#34;XOR&#34;,
	OPOS,		0,	&#34;POS&#34;,
	ONEG,		0,	&#34;NEG&#34;,
	OCOM,		0,	&#34;COM&#34;,
	OELEM,		0,	&#34;ELEM&#34;,
	OTST,		0,	&#34;TST&#34;,
	OINDEX,		0,	&#34;INDEX&#34;,
	OFAS,		0,	&#34;FAS&#34;,
	OREGPAIR,	0,	&#34;REGPAIR&#34;,
	OEND,		0,	&#34;END&#34;,
	-1,		0,	0,
};

/*	OEQ, ONE, OLE, OLS, OLT, OLO, OGE, OHS, OGT, OHI */
uchar	comrel[12] =
{
	ONE, OEQ, OGT, OHI, OGE, OHS, OLT, OLO, OLE, OLS,
};
uchar	invrel[12] =
{
	OEQ, ONE, OGE, OHS, OGT, OHI, OLE, OLS, OLT, OLO,
};
uchar	logrel[12] =
{
	OEQ, ONE, OLS, OLS, OLO, OLO, OHS, OHS, OHI, OHI,
};

uchar	typei[NTYPE];
int	typeiinit[] =
{
	TCHAR, TUCHAR, TSHORT, TUSHORT, TINT, TUINT, TLONG, TULONG, TVLONG, TUVLONG, -1,
};
uchar	typeu[NTYPE];
int	typeuinit[] =
{
	TUCHAR, TUSHORT, TUINT, TULONG, TUVLONG, TIND, -1,
};

uchar	typesuv[NTYPE];
int	typesuvinit[] =
{
	TVLONG, TUVLONG, TSTRUCT, TUNION, -1,
};

uchar	typeilp[NTYPE];
int	typeilpinit[] =
{
	TINT, TUINT, TLONG, TULONG, TIND, -1
};

uchar	typechl[NTYPE];
uchar	typechlv[NTYPE];
uchar	typechlvp[NTYPE];
int	typechlinit[] =
{
	TCHAR, TUCHAR, TSHORT, TUSHORT, TINT, TUINT, TLONG, TULONG, -1,
};

uchar	typechlp[NTYPE];
int	typechlpinit[] =
{
	TCHAR, TUCHAR, TSHORT, TUSHORT, TINT, TUINT, TLONG, TULONG, TIND, -1,
};

uchar	typechlpfd[NTYPE];
int	typechlpfdinit[] =
{
	TCHAR, TUCHAR, TSHORT, TUSHORT, TINT, TUINT, TLONG, TULONG, TFLOAT, TDOUBLE, TIND, -1,
};

uchar	typec[NTYPE];
int	typecinit[] =
{
	TCHAR, TUCHAR, -1
};

uchar	typeh[NTYPE];
int	typehinit[] =
{
	TSHORT, TUSHORT, -1,
};

uchar	typeil[NTYPE];
int	typeilinit[] =
{
	TINT, TUINT, TLONG, TULONG, -1,
};

uchar	typev[NTYPE];
int	typevinit[] =
{
	TVLONG,	TUVLONG, -1,
};

uchar	typefd[NTYPE];
int	typefdinit[] =
{
	TFLOAT, TDOUBLE, -1,
};

uchar	typeaf[NTYPE];
int	typeafinit[] =
{
	TFUNC, TARRAY, -1,
};

uchar	typesu[NTYPE];
int	typesuinit[] =
{
	TSTRUCT, TUNION, -1,
};

int32	tasign[NTYPE];
Init	tasigninit[] =
{
	TCHAR,		BNUMBER,	0,
	TUCHAR,		BNUMBER,	0,
	TSHORT,		BNUMBER,	0,
	TUSHORT,	BNUMBER,	0,
	TINT,		BNUMBER,	0,
	TUINT,		BNUMBER,	0,
	TLONG,		BNUMBER,	0,
	TULONG,		BNUMBER,	0,
	TVLONG,		BNUMBER,	0,
	TUVLONG,	BNUMBER,	0,
	TFLOAT,		BNUMBER,	0,
	TDOUBLE,	BNUMBER,	0,
	TIND,		BIND,		0,
	TSTRUCT,	BSTRUCT,	0,
	TUNION,		BUNION,		0,
	-1,		0,		0,
};

int32	tasadd[NTYPE];
Init	tasaddinit[] =
{
	TCHAR,		BNUMBER,	0,
	TUCHAR,		BNUMBER,	0,
	TSHORT,		BNUMBER,	0,
	TUSHORT,	BNUMBER,	0,
	TINT,		BNUMBER,	0,
	TUINT,		BNUMBER,	0,
	TLONG,		BNUMBER,	0,
	TULONG,		BNUMBER,	0,
	TVLONG,		BNUMBER,	0,
	TUVLONG,	BNUMBER,	0,
	TFLOAT,		BNUMBER,	0,
	TDOUBLE,	BNUMBER,	0,
	TIND,		BINTEGER,	0,
	-1,		0,		0,
};

int32	tcast[NTYPE];
Init	tcastinit[] =
{
	TCHAR,		BNUMBER|BIND|BVOID,	0,
	TUCHAR,		BNUMBER|BIND|BVOID,	0,
	TSHORT,		BNUMBER|BIND|BVOID,	0,
	TUSHORT,	BNUMBER|BIND|BVOID,	0,
	TINT,		BNUMBER|BIND|BVOID,	0,
	TUINT,		BNUMBER|BIND|BVOID,	0,
	TLONG,		BNUMBER|BIND|BVOID,	0,
	TULONG,		BNUMBER|BIND|BVOID,	0,
	TVLONG,		BNUMBER|BIND|BVOID,	0,
	TUVLONG,	BNUMBER|BIND|BVOID,	0,
	TFLOAT,		BNUMBER|BVOID,		0,
	TDOUBLE,	BNUMBER|BVOID,		0,
	TIND,		BINTEGER|BIND|BVOID,	0,
	TVOID,		BVOID,			0,
	TSTRUCT,	BSTRUCT|BVOID,		0,
	TUNION,		BUNION|BVOID,		0,
	-1,		0,			0,
};

int32	tadd[NTYPE];
Init	taddinit[] =
{
	TCHAR,		BNUMBER|BIND,	0,
	TUCHAR,		BNUMBER|BIND,	0,
	TSHORT,		BNUMBER|BIND,	0,
	TUSHORT,	BNUMBER|BIND,	0,
	TINT,		BNUMBER|BIND,	0,
	TUINT,		BNUMBER|BIND,	0,
	TLONG,		BNUMBER|BIND,	0,
	TULONG,		BNUMBER|BIND,	0,
	TVLONG,		BNUMBER|BIND,	0,
	TUVLONG,	BNUMBER|BIND,	0,
	TFLOAT,		BNUMBER,	0,
	TDOUBLE,	BNUMBER,	0,
	TIND,		BINTEGER,	0,
	-1,		0,		0,
};

int32	tsub[NTYPE];
Init	tsubinit[] =
{
	TCHAR,		BNUMBER,	0,
	TUCHAR,		BNUMBER,	0,
	TSHORT,		BNUMBER,	0,
	TUSHORT,	BNUMBER,	0,
	TINT,		BNUMBER,	0,
	TUINT,		BNUMBER,	0,
	TLONG,		BNUMBER,	0,
	TULONG,		BNUMBER,	0,
	TVLONG,		BNUMBER,	0,
	TUVLONG,	BNUMBER,	0,
	TFLOAT,		BNUMBER,	0,
	TDOUBLE,	BNUMBER,	0,
	TIND,		BINTEGER|BIND,	0,
	-1,		0,		0,
};

int32	tmul[NTYPE];
Init	tmulinit[] =
{
	TCHAR,		BNUMBER,	0,
	TUCHAR,		BNUMBER,	0,
	TSHORT,		BNUMBER,	0,
	TUSHORT,	BNUMBER,	0,
	TINT,		BNUMBER,	0,
	TUINT,		BNUMBER,	0,
	TLONG,		BNUMBER,	0,
	TULONG,		BNUMBER,	0,
	TVLONG,		BNUMBER,	0,
	TUVLONG,	BNUMBER,	0,
	TFLOAT,		BNUMBER,	0,
	TDOUBLE,	BNUMBER,	0,
	-1,		0,		0,
};

int32	tand[NTYPE];
Init	tandinit[] =
{
	TCHAR,		BINTEGER,	0,
	TUCHAR,		BINTEGER,	0,
	TSHORT,		BINTEGER,	0,
	TUSHORT,	BINTEGER,	0,
	TINT,		BNUMBER,	0,
	TUINT,		BNUMBER,	0,
	TLONG,		BINTEGER,	0,
	TULONG,		BINTEGER,	0,
	TVLONG,		BINTEGER,	0,
	TUVLONG,	BINTEGER,	0,
	-1,		0,		0,
};

int32	trel[NTYPE];
Init	trelinit[] =
{
	TCHAR,		BNUMBER,	0,
	TUCHAR,		BNUMBER,	0,
	TSHORT,		BNUMBER,	0,
	TUSHORT,	BNUMBER,	0,
	TINT,		BNUMBER,	0,
	TUINT,		BNUMBER,	0,
	TLONG,		BNUMBER,	0,
	TULONG,		BNUMBER,	0,
	TVLONG,		BNUMBER,	0,
	TUVLONG,	BNUMBER,	0,
	TFLOAT,		BNUMBER,	0,
	TDOUBLE,	BNUMBER,	0,
	TIND,		BIND,		0,
	-1,		0,		0,
};

int32	tfunct[1] =
{
	BFUNC,
};

int32	tindir[1] =
{
	BIND,
};

int32	tdot[1] =
{
	BSTRUCT|BUNION,
};

int32	tnot[1] =
{
	BNUMBER|BIND,
};

int32	targ[1] =
{
	BNUMBER|BIND|BSTRUCT|BUNION,
};

uchar	tab[NTYPE][NTYPE] =
{
/*TXXX*/	{ 0,
		},

/*TCHAR*/	{ 0,	TCHAR, TUCHAR, TSHORT, TUSHORT, TINT, TUINT, TLONG,
			TULONG, TVLONG, TUVLONG, TFLOAT, TDOUBLE, TIND,
		},
/*TUCHAR*/	{ 0,	TUCHAR, TUCHAR, TUSHORT, TUSHORT, TUINT, TUINT, TULONG,
			TULONG, TUVLONG, TUVLONG, TFLOAT, TDOUBLE, TIND,
		},
/*TSHORT*/	{ 0,	TSHORT, TUSHORT, TSHORT, TUSHORT, TINT, TUINT, TLONG,
			TULONG, TVLONG, TUVLONG, TFLOAT, TDOUBLE, TIND,
		},
/*TUSHORT*/	{ 0,	TUSHORT, TUSHORT, TUSHORT, TUSHORT, TUINT, TUINT, TULONG,
			TULONG, TUVLONG, TUVLONG, TFLOAT, TDOUBLE, TIND,
		},
/*TINT*/	{ 0,	TINT, TUINT, TINT, TUINT, TINT, TUINT, TLONG,
			TULONG, TVLONG, TUVLONG, TFLOAT, TDOUBLE, TIND,
		},
/*TUINT*/	{ 0,	TUINT, TUINT, TUINT, TUINT, TUINT, TUINT, TULONG,
			TULONG, TUVLONG, TUVLONG, TFLOAT, TDOUBLE, TIND,
		},
/*TLONG*/	{ 0,	TLONG, TULONG, TLONG, TULONG, TLONG, TULONG, TLONG,
			TULONG, TVLONG, TUVLONG, TFLOAT, TDOUBLE, TIND,
		},
/*TULONG*/	{ 0,	TULONG, TULONG, TULONG, TULONG, TULONG, TULONG, TULONG,
			TULONG, TUVLONG, TUVLONG, TFLOAT, TDOUBLE, TIND,
		},
/*TVLONG*/	{ 0,	TVLONG, TUVLONG, TVLONG, TUVLONG, TVLONG, TUVLONG, TVLONG,
			TUVLONG, TVLONG, TUVLONG, TFLOAT, TDOUBLE, TIND,
		},
/*TUVLONG*/	{ 0,	TUVLONG, TUVLONG, TUVLONG, TUVLONG, TUVLONG, TUVLONG, TUVLONG,
			TUVLONG, TUVLONG, TUVLONG, TFLOAT, TDOUBLE, TIND,
		},
/*TFLOAT*/	{ 0,	TFLOAT, TFLOAT, TFLOAT, TFLOAT, TFLOAT, TFLOAT, TFLOAT,
			TFLOAT, TFLOAT, TFLOAT, TFLOAT, TDOUBLE, TIND,
		},
/*TDOUBLE*/	{ 0,	TDOUBLE, TDOUBLE, TDOUBLE, TDOUBLE, TDOUBLE, TDOUBLE, TDOUBLE,
			TDOUBLE, TDOUBLE, TDOUBLE, TFLOAT, TDOUBLE, TIND,
		},
/*TIND*/	{ 0,	TIND, TIND, TIND, TIND, TIND, TIND, TIND,
			 TIND, TIND, TIND, TIND, TIND, TIND,
		},
};

void
urk(char *name, int max, int i)
{
	if(i &gt;= max) {
		fprint(2, &#34;bad tinit: %s %d&gt;=%d\n&#34;, name, i, max);
		exits(&#34;init&#34;);
	}
}

void
tinit(void)
{
	int *ip;
	Init *p;

	for(p=thashinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;thash&#34;, nelem(thash), p-&gt;code);
		thash[p-&gt;code] = p-&gt;value;
	}
	for(p=bnamesinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;bnames&#34;, nelem(bnames), p-&gt;code);
		bnames[p-&gt;code] = p-&gt;s;
	}
	for(p=tnamesinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;tnames&#34;, nelem(tnames), p-&gt;code);
		tnames[p-&gt;code] = p-&gt;s;
	}
	for(p=gnamesinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;gnames&#34;, nelem(gnames), p-&gt;code);
		gnames[p-&gt;code] = p-&gt;s;
	}
	for(p=qnamesinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;qnames&#34;, nelem(qnames), p-&gt;code);
		qnames[p-&gt;code] = p-&gt;s;
	}
	for(p=cnamesinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;cnames&#34;, nelem(cnames), p-&gt;code);
		cnames[p-&gt;code] = p-&gt;s;
	}
	for(p=onamesinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;onames&#34;, nelem(onames), p-&gt;code);
		onames[p-&gt;code] = p-&gt;s;
	}
	for(ip=typeiinit; *ip&gt;=0; ip++) {
		urk(&#34;typei&#34;, nelem(typei), *ip);
		typei[*ip] = 1;
	}
	for(ip=typeuinit; *ip&gt;=0; ip++) {
		urk(&#34;typeu&#34;, nelem(typeu), *ip);
		typeu[*ip] = 1;
	}
	for(ip=typesuvinit; *ip&gt;=0; ip++) {
		urk(&#34;typesuv&#34;, nelem(typesuv), *ip);
		typesuv[*ip] = 1;
	}
	for(ip=typeilpinit; *ip&gt;=0; ip++) {
		urk(&#34;typeilp&#34;, nelem(typeilp), *ip);
		typeilp[*ip] = 1;
	}
	for(ip=typechlinit; *ip&gt;=0; ip++) {
		urk(&#34;typechl&#34;, nelem(typechl), *ip);
		typechl[*ip] = 1;
		typechlv[*ip] = 1;
		typechlvp[*ip] = 1;
	}
	for(ip=typechlpinit; *ip&gt;=0; ip++) {
		urk(&#34;typechlp&#34;, nelem(typechlp), *ip);
		typechlp[*ip] = 1;
		typechlvp[*ip] = 1;
	}
	for(ip=typechlpfdinit; *ip&gt;=0; ip++) {
		urk(&#34;typechlpfd&#34;, nelem(typechlpfd), *ip);
		typechlpfd[*ip] = 1;
	}
	for(ip=typecinit; *ip&gt;=0; ip++) {
		urk(&#34;typec&#34;, nelem(typec), *ip);
		typec[*ip] = 1;
	}
	for(ip=typehinit; *ip&gt;=0; ip++) {
		urk(&#34;typeh&#34;, nelem(typeh), *ip);
		typeh[*ip] = 1;
	}
	for(ip=typeilinit; *ip&gt;=0; ip++) {
		urk(&#34;typeil&#34;, nelem(typeil), *ip);
		typeil[*ip] = 1;
	}
	for(ip=typevinit; *ip&gt;=0; ip++) {
		urk(&#34;typev&#34;, nelem(typev), *ip);
		typev[*ip] = 1;
		typechlv[*ip] = 1;
		typechlvp[*ip] = 1;
	}
	for(ip=typefdinit; *ip&gt;=0; ip++) {
		urk(&#34;typefd&#34;, nelem(typefd), *ip);
		typefd[*ip] = 1;
	}
	for(ip=typeafinit; *ip&gt;=0; ip++) {
		urk(&#34;typeaf&#34;, nelem(typeaf), *ip);
		typeaf[*ip] = 1;
	}
	for(ip=typesuinit; *ip &gt;= 0; ip++) {
		urk(&#34;typesu&#34;, nelem(typesu), *ip);
		typesu[*ip] = 1;
	}
	for(p=tasigninit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;tasign&#34;, nelem(tasign), p-&gt;code);
		tasign[p-&gt;code] = p-&gt;value;
	}
	for(p=tasaddinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;tasadd&#34;, nelem(tasadd), p-&gt;code);
		tasadd[p-&gt;code] = p-&gt;value;
	}
	for(p=tcastinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;tcast&#34;, nelem(tcast), p-&gt;code);
		tcast[p-&gt;code] = p-&gt;value;
	}
	for(p=taddinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;tadd&#34;, nelem(tadd), p-&gt;code);
		tadd[p-&gt;code] = p-&gt;value;
	}
	for(p=tsubinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;tsub&#34;, nelem(tsub), p-&gt;code);
		tsub[p-&gt;code] = p-&gt;value;
	}
	for(p=tmulinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;tmul&#34;, nelem(tmul), p-&gt;code);
		tmul[p-&gt;code] = p-&gt;value;
	}
	for(p=tandinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;tand&#34;, nelem(tand), p-&gt;code);
		tand[p-&gt;code] = p-&gt;value;
	}
	for(p=trelinit; p-&gt;code &gt;= 0; p++) {
		urk(&#34;trel&#34;, nelem(trel), p-&gt;code);
		trel[p-&gt;code] = p-&gt;value;
	}
	
	/* 32-bit defaults */
	typeword = typechlp;
	typecmplx = typesuv;
}

/*
 * return 1 if it is impossible to jump into the middle of n.
 */
static int
deadhead(Node *n, int caseok)
{
loop:
	if(n == Z)
		return 1;
	switch(n-&gt;op) {
	case OLIST:
		if(!deadhead(n-&gt;left, caseok))
			return 0;
	rloop:
		n = n-&gt;right;
		goto loop;

	case ORETURN:
		break;

	case OLABEL:
		return 0;

	case OGOTO:
		break;

	case OCASE:
		if(!caseok)
			return 0;
		goto rloop;

	case OSWITCH:
		return deadhead(n-&gt;right, 1);

	case OWHILE:
	case ODWHILE:
		goto rloop;

	case OFOR:
		goto rloop;

	case OCONTINUE:
		break;

	case OBREAK:
		break;

	case OIF:
		return deadhead(n-&gt;right-&gt;left, caseok) &amp;&amp; deadhead(n-&gt;right-&gt;right, caseok);

	case OSET:
	case OUSED:
		break;
	}
	return 1;
}

int
deadheads(Node *c)
{
	return deadhead(c-&gt;left, 0) &amp;&amp; deadhead(c-&gt;right, 0);
}

int
mixedasop(Type *l, Type *r)
{
	return !typefd[l-&gt;etype] &amp;&amp; typefd[r-&gt;etype];
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
