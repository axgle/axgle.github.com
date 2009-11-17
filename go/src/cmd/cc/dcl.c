<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cc/dcl.c</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/cc/dcl.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/cc/dcl.c
// http://code.google.com/p/inferno-os/source/browse/utils/cc/dcl.c
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

#include &#34;cc.h&#34;

Node*
dodecl(void (*f)(int,Type*,Sym*), int c, Type *t, Node *n)
{
	Sym *s;
	Node *n1;
	int32 v;

	nearln = lineno;
	lastfield = 0;

loop:
	if(n != Z)
	switch(n-&gt;op) {
	default:
		diag(n, &#34;unknown declarator: %O&#34;, n-&gt;op);
		break;

	case OARRAY:
		t = typ(TARRAY, t);
		t-&gt;width = 0;
		n1 = n-&gt;right;
		n = n-&gt;left;
		if(n1 != Z) {
			complex(n1);
			v = -1;
			if(n1-&gt;op == OCONST)
				v = n1-&gt;vconst;
			if(v &lt;= 0) {
				diag(n, &#34;array size must be a positive constant&#34;);
				v = 1;
			}
			t-&gt;width = v * t-&gt;link-&gt;width;
		}
		goto loop;

	case OIND:
		t = typ(TIND, t);
		t-&gt;garb = n-&gt;garb;
		n = n-&gt;left;
		goto loop;

	case OFUNC:
		t = typ(TFUNC, t);
		t-&gt;down = fnproto(n);
		n = n-&gt;left;
		goto loop;

	case OBIT:
		n1 = n-&gt;right;
		complex(n1);
		lastfield = -1;
		if(n1-&gt;op == OCONST)
			lastfield = n1-&gt;vconst;
		if(lastfield &lt; 0) {
			diag(n, &#34;field width must be non-negative constant&#34;);
			lastfield = 1;
		}
		if(lastfield == 0) {
			lastbit = 0;
			firstbit = 1;
			if(n-&gt;left != Z) {
				diag(n, &#34;zero width named field&#34;);
				lastfield = 1;
			}
		}
		if(!typei[t-&gt;etype]) {
			diag(n, &#34;field type must be int-like&#34;);
			t = types[TINT];
			lastfield = 1;
		}
		if(lastfield &gt; tfield-&gt;width*8) {
			diag(n, &#34;field width larger than field unit&#34;);
			lastfield = 1;
		}
		lastbit += lastfield;
		if(lastbit &gt; tfield-&gt;width*8) {
			lastbit = lastfield;
			firstbit = 1;
		}
		n = n-&gt;left;
		goto loop;

	case ONAME:
		if(f == NODECL)
			break;
		s = n-&gt;sym;
		(*f)(c, t, s);
		if(s-&gt;class == CLOCAL)
			s = mkstatic(s);
		firstbit = 0;
		n-&gt;sym = s;
		n-&gt;type = s-&gt;type;
		n-&gt;xoffset = s-&gt;offset;
		n-&gt;class = s-&gt;class;
		n-&gt;etype = TVOID;
		if(n-&gt;type != T)
			n-&gt;etype = n-&gt;type-&gt;etype;
		if(debug[&#39;d&#39;])
			dbgdecl(s);
		acidvar(s);
		s-&gt;varlineno = lineno;
		break;
	}
	lastdcl = t;
	return n;
}

Sym*
mkstatic(Sym *s)
{
	Sym *s1;

	if(s-&gt;class != CLOCAL)
		return s;
	snprint(symb, NSYMB, &#34;%s$%d&#34;, s-&gt;name, s-&gt;block);
	s1 = lookup();
	if(s1-&gt;class != CSTATIC) {
		s1-&gt;type = s-&gt;type;
		s1-&gt;offset = s-&gt;offset;
		s1-&gt;block = s-&gt;block;
		s1-&gt;class = CSTATIC;
	}
	return s1;
}

/*
 * make a copy of a typedef
 * the problem is to split out incomplete
 * arrays so that it is in the variable
 * rather than the typedef.
 */
Type*
tcopy(Type *t)
{
	Type *tl, *tx;
	int et;

	if(t == T)
		return t;
	et = t-&gt;etype;
	if(typesu[et])
		return t;
	tl = tcopy(t-&gt;link);
	if(tl != t-&gt;link ||
	  (et == TARRAY &amp;&amp; t-&gt;width == 0)) {
		tx = copytyp(t);
		tx-&gt;link = tl;
		return tx;
	}
	return t;
}

Node*
doinit(Sym *s, Type *t, int32 o, Node *a)
{
	Node *n;

	if(t == T)
		return Z;
	if(s-&gt;class == CEXTERN) {
		s-&gt;class = CGLOBL;
		if(debug[&#39;d&#39;])
			dbgdecl(s);
	}
	if(debug[&#39;i&#39;]) {
		print(&#34;t = %T; o = %ld; n = %s\n&#34;, t, o, s-&gt;name);
		prtree(a, &#34;doinit value&#34;);
	}


	n = initlist;
	if(a-&gt;op == OINIT)
		a = a-&gt;left;
	initlist = a;

	a = init1(s, t, o, 0);
	if(initlist != Z)
		diag(initlist, &#34;more initializers than structure: %s&#34;,
			s-&gt;name);
	initlist = n;

	return a;
}

/*
 * get next major operator,
 * dont advance initlist.
 */
Node*
peekinit(void)
{
	Node *a;

	a = initlist;

loop:
	if(a == Z)
		return a;
	if(a-&gt;op == OLIST) {
		a = a-&gt;left;
		goto loop;
	}
	return a;
}

/*
 * consume and return next element on
 * initlist. expand strings.
 */
Node*
nextinit(void)
{
	Node *a, *b, *n;

	a = initlist;
	n = Z;

	if(a == Z)
		return a;
	if(a-&gt;op == OLIST) {
		n = a-&gt;right;
		a = a-&gt;left;
	}
	if(a-&gt;op == OUSED) {
		a = a-&gt;left;
		b = new(OCONST, Z, Z);
		b-&gt;type = a-&gt;type-&gt;link;
		if(a-&gt;op == OSTRING) {
			b-&gt;vconst = convvtox(*a-&gt;cstring, TCHAR);
			a-&gt;cstring++;
		}
		if(a-&gt;op == OLSTRING) {
			b-&gt;vconst = convvtox(*a-&gt;rstring, TUSHORT);
			a-&gt;rstring++;
		}
		a-&gt;type-&gt;width -= b-&gt;type-&gt;width;
		if(a-&gt;type-&gt;width &lt;= 0)
			initlist = n;
		return b;
	}
	initlist = n;
	return a;
}

int
isstruct(Node *a, Type *t)
{
	Node *n;

	switch(a-&gt;op) {
	case ODOTDOT:
		n = a-&gt;left;
		if(n &amp;&amp; n-&gt;type &amp;&amp; sametype(n-&gt;type, t))
			return 1;
	case OSTRING:
	case OLSTRING:
	case OCONST:
	case OINIT:
	case OELEM:
		return 0;
	}

	n = new(ODOTDOT, Z, Z);
	*n = *a;

	/*
	 * ODOTDOT is a flag for tcom
	 * a second tcom will not be performed
	 */
	a-&gt;op = ODOTDOT;
	a-&gt;left = n;
	a-&gt;right = Z;

	if(tcom(n))
		return 0;

	if(sametype(n-&gt;type, t))
		return 1;
	return 0;
}

Node*
init1(Sym *s, Type *t, int32 o, int exflag)
{
	Node *a, *l, *r, nod;
	Type *t1;
	int32 e, w, so, mw;

	a = peekinit();
	if(a == Z)
		return Z;

	if(debug[&#39;i&#39;]) {
		print(&#34;t = %T; o = %ld; n = %s\n&#34;, t, o, s-&gt;name);
		prtree(a, &#34;init1 value&#34;);
	}

	if(exflag &amp;&amp; a-&gt;op == OINIT)
		return doinit(s, t, o, nextinit());

	switch(t-&gt;etype) {
	default:
		diag(Z, &#34;unknown type in initialization: %T to: %s&#34;, t, s-&gt;name);
		return Z;

	case TCHAR:
	case TUCHAR:
	case TINT:
	case TUINT:
	case TSHORT:
	case TUSHORT:
	case TLONG:
	case TULONG:
	case TVLONG:
	case TUVLONG:
	case TFLOAT:
	case TDOUBLE:
	case TIND:
	single:
		if(a-&gt;op == OARRAY || a-&gt;op == OELEM)
			return Z;

		a = nextinit();
		if(a == Z)
			return Z;

		if(t-&gt;nbits)
			diag(Z, &#34;cannot initialize bitfields&#34;);
		if(s-&gt;class == CAUTO) {
			l = new(ONAME, Z, Z);
			l-&gt;sym = s;
			l-&gt;type = t;
			l-&gt;etype = TVOID;
			if(s-&gt;type)
				l-&gt;etype = s-&gt;type-&gt;etype;
			l-&gt;xoffset = s-&gt;offset + o;
			l-&gt;class = s-&gt;class;

			l = new(OASI, l, a);
			return l;
		}

		complex(a);
		if(a-&gt;type == T)
			return Z;

		if(a-&gt;op == OCONST) {
			if(vconst(a) &amp;&amp; t-&gt;etype == TIND &amp;&amp; a-&gt;type &amp;&amp; a-&gt;type-&gt;etype != TIND){
				diag(a, &#34;initialize pointer to an integer: %s&#34;, s-&gt;name);
				return Z;
			}
			if(!sametype(a-&gt;type, t)) {
				/* hoop jumping to save malloc */
				if(nodcast == Z)
					nodcast = new(OCAST, Z, Z);
				nod = *nodcast;
				nod.left = a;
				nod.type = t;
				nod.lineno = a-&gt;lineno;
				complex(&amp;nod);
				if(nod.type)
					*a = nod;
			}
			if(a-&gt;op != OCONST) {
				diag(a, &#34;initializer is not a constant: %s&#34;,
					s-&gt;name);
				return Z;
			}
			if(vconst(a) == 0)
				return Z;
			goto gext;
		}
		if(t-&gt;etype == TIND) {
			while(a-&gt;op == OCAST) {
				warn(a, &#34;CAST in initialization ignored&#34;);
				a = a-&gt;left;
			}
			if(!sametype(t, a-&gt;type)) {
				diag(a, &#34;initialization of incompatible pointers: %s\n%T and %T&#34;,
					s-&gt;name, t, a-&gt;type);
			}
			if(a-&gt;op == OADDR)
				a = a-&gt;left;
			goto gext;
		}

		while(a-&gt;op == OCAST)
			a = a-&gt;left;
		if(a-&gt;op == OADDR) {
			warn(a, &#34;initialize pointer to an integer: %s&#34;, s-&gt;name);
			a = a-&gt;left;
			goto gext;
		}
		diag(a, &#34;initializer is not a constant: %s&#34;, s-&gt;name);
		return Z;

	gext:
		gextern(s, a, o, t-&gt;width);

		return Z;

	case TARRAY:
		w = t-&gt;link-&gt;width;
		if(a-&gt;op == OSTRING || a-&gt;op == OLSTRING)
		if(typei[t-&gt;link-&gt;etype]) {
			/*
			 * get rid of null if sizes match exactly
			 */
			a = nextinit();
			mw = t-&gt;width/w;
			so = a-&gt;type-&gt;width/a-&gt;type-&gt;link-&gt;width;
			if(mw &amp;&amp; so &gt; mw) {
				if(so != mw+1)
					diag(a, &#34;string initialization larger than array&#34;);
				a-&gt;type-&gt;width -= a-&gt;type-&gt;link-&gt;width;
			}

			/*
			 * arrange strings to be expanded
			 * inside OINIT braces.
			 */
			a = new(OUSED, a, Z);
			return doinit(s, t, o, a);
		}

		mw = -w;
		l = Z;
		for(e=0;;) {
			/*
			 * peek ahead for element initializer
			 */
			a = peekinit();
			if(a == Z)
				break;
			if(a-&gt;op == OELEM &amp;&amp; t-&gt;link-&gt;etype != TSTRUCT)
				break;
			if(a-&gt;op == OARRAY) {
				if(e &amp;&amp; exflag)
					break;
				a = nextinit();
				r = a-&gt;left;
				complex(r);
				if(r-&gt;op != OCONST) {
					diag(r, &#34;initializer subscript must be constant&#34;);
					return Z;
				}
				e = r-&gt;vconst;
				if(t-&gt;width != 0)
					if(e &lt; 0 || e*w &gt;= t-&gt;width) {
						diag(a, &#34;initialization index out of range: %ld&#34;, e);
						continue;
					}
			}

			so = e*w;
			if(so &gt; mw)
				mw = so;
			if(t-&gt;width != 0)
				if(mw &gt;= t-&gt;width)
					break;
			r = init1(s, t-&gt;link, o+so, 1);
			l = newlist(l, r);
			e++;
		}
		if(t-&gt;width == 0)
			t-&gt;width = mw+w;
		return l;

	case TUNION:
	case TSTRUCT:
		/*
		 * peek ahead to find type of rhs.
		 * if its a structure, then treat
		 * this element as a variable
		 * rather than an aggregate.
		 */
		if(isstruct(a, t))
			goto single;

		if(t-&gt;width &lt;= 0) {
			diag(Z, &#34;incomplete structure: %s&#34;, s-&gt;name);
			return Z;
		}
		l = Z;

	again:
		for(t1 = t-&gt;link; t1 != T; t1 = t1-&gt;down) {
			if(a-&gt;op == OARRAY &amp;&amp; t1-&gt;etype != TARRAY)
				break;
			if(a-&gt;op == OELEM) {
				if(t1-&gt;sym != a-&gt;sym)
					continue;
				nextinit();
			}
			r = init1(s, t1, o+t1-&gt;offset, 1);
			l = newlist(l, r);
			a = peekinit();
			if(a == Z)
				break;
			if(a-&gt;op == OELEM)
				goto again;
		}
		if(a &amp;&amp; a-&gt;op == OELEM)
			diag(a, &#34;structure element not found %F&#34;, a);
		return l;
	}
}

Node*
newlist(Node *l, Node *r)
{
	if(r == Z)
		return l;
	if(l == Z)
		return r;
	return new(OLIST, l, r);
}

void
suallign(Type *t)
{
	Type *l;
	int32 o, w;

	o = 0;
	switch(t-&gt;etype) {

	case TSTRUCT:
		t-&gt;offset = 0;
		w = 0;
		for(l = t-&gt;link; l != T; l = l-&gt;down) {
			if(l-&gt;nbits) {
				if(l-&gt;shift &lt;= 0) {
					l-&gt;shift = -l-&gt;shift;
					w = xround(w, tfield-&gt;width);
					o = w;
					w += tfield-&gt;width;
				}
				l-&gt;offset = o;
			} else {
				if(l-&gt;width &lt;= 0)
				if(l-&gt;down != T)
					if(l-&gt;sym)
						diag(Z, &#34;incomplete structure element: %s&#34;,
							l-&gt;sym-&gt;name);
					else
						diag(Z, &#34;incomplete structure element&#34;);
				w = align(w, l, Ael1);
				l-&gt;offset = w;
				w = align(w, l, Ael2);
			}
		}
		w = align(w, t, Asu2);
		t-&gt;width = w;
		acidtype(t);
		pickletype(t);
		return;

	case TUNION:
		t-&gt;offset = 0;
		w = 0;
		for(l = t-&gt;link; l != T; l = l-&gt;down) {
			if(l-&gt;width &lt;= 0)
				if(l-&gt;sym)
					diag(Z, &#34;incomplete union element: %s&#34;,
						l-&gt;sym-&gt;name);
				else
					diag(Z, &#34;incomplete union element&#34;);
			l-&gt;offset = 0;
			l-&gt;shift = 0;
			o = align(align(0, l, Ael1), l, Ael2);
			if(o &gt; w)
				w = o;
		}
		w = align(w, t, Asu2);
		t-&gt;width = w;
		acidtype(t);
		pickletype(t);
		return;

	default:
		diag(Z, &#34;unknown type in suallign: %T&#34;, t);
		break;
	}
}

int32
xround(int32 v, int w)
{
	int r;

	if(w &lt;= 0 || w &gt; 8) {
		diag(Z, &#34;rounding by %d&#34;, w);
		w = 1;
	}
	r = v%w;
	if(r)
		v += w-r;
	return v;
}

Type*
ofnproto(Node *n)
{
	Type *tl, *tr, *t;

	if(n == Z)
		return T;
	switch(n-&gt;op) {
	case OLIST:
		tl = ofnproto(n-&gt;left);
		tr = ofnproto(n-&gt;right);
		if(tl == T)
			return tr;
		tl-&gt;down = tr;
		return tl;

	case ONAME:
		t = copytyp(n-&gt;sym-&gt;type);
		t-&gt;down = T;
		return t;
	}
	return T;
}

#define	ANSIPROTO	1
#define	OLDPROTO	2

void
argmark(Node *n, int pass)
{
	Type *t;

	autoffset = align(0, thisfn-&gt;link, Aarg0);
	stkoff = 0;
	for(; n-&gt;left != Z; n = n-&gt;left) {
		if(n-&gt;op != OFUNC || n-&gt;left-&gt;op != ONAME)
			continue;
		walkparam(n-&gt;right, pass);
		if(pass != 0 &amp;&amp; anyproto(n-&gt;right) == OLDPROTO) {
			t = typ(TFUNC, n-&gt;left-&gt;sym-&gt;type-&gt;link);
			t-&gt;down = typ(TOLD, T);
			t-&gt;down-&gt;down = ofnproto(n-&gt;right);
			tmerge(t, n-&gt;left-&gt;sym);
			n-&gt;left-&gt;sym-&gt;type = t;
		}
		break;
	}
	autoffset = 0;
	stkoff = 0;
}

void
walkparam(Node *n, int pass)
{
	Sym *s;
	Node *n1;

	if(n != Z &amp;&amp; n-&gt;op == OPROTO &amp;&amp; n-&gt;left == Z &amp;&amp; n-&gt;type == types[TVOID])
		return;

loop:
	if(n == Z)
		return;
	switch(n-&gt;op) {
	default:
		diag(n, &#34;argument not a name/prototype: %O&#34;, n-&gt;op);
		break;

	case OLIST:
		walkparam(n-&gt;left, pass);
		n = n-&gt;right;
		goto loop;

	case OPROTO:
		for(n1 = n; n1 != Z; n1=n1-&gt;left)
			if(n1-&gt;op == ONAME) {
				if(pass == 0) {
					s = n1-&gt;sym;
					push1(s);
					s-&gt;offset = -1;
					break;
				}
				dodecl(pdecl, CPARAM, n-&gt;type, n-&gt;left);
				break;
			}
		if(n1)
			break;
		if(pass == 0) {
			/*
			 * extension:
			 *	allow no name in argument declaration
			diag(Z, &#34;no name in argument declaration&#34;);
			 */
			break;
		}
		dodecl(NODECL, CPARAM, n-&gt;type, n-&gt;left);
		pdecl(CPARAM, lastdcl, S);
		break;

	case ODOTDOT:
		break;
	
	case ONAME:
		s = n-&gt;sym;
		if(pass == 0) {
			push1(s);
			s-&gt;offset = -1;
			break;
		}
		if(s-&gt;offset != -1) {
			if(autoffset == 0) {
				firstarg = s;
				firstargtype = s-&gt;type;
			}
			autoffset = align(autoffset, s-&gt;type, Aarg1);
			s-&gt;offset = autoffset;
			autoffset = align(autoffset, s-&gt;type, Aarg2);
		} else
			dodecl(pdecl, CXXX, types[TINT], n);
		break;
	}
}

void
markdcl(void)
{
	Decl *d;

	blockno++;
	d = push();
	d-&gt;val = DMARK;
	d-&gt;offset = autoffset;
	d-&gt;block = autobn;
	autobn = blockno;
}

Node*
revertdcl(void)
{
	Decl *d;
	Sym *s;
	Node *n, *n1;

	n = Z;
	for(;;) {
		d = dclstack;
		if(d == D) {
			diag(Z, &#34;pop off dcl stack&#34;);
			break;
		}
		dclstack = d-&gt;link;
		s = d-&gt;sym;
		switch(d-&gt;val) {
		case DMARK:
			autoffset = d-&gt;offset;
			autobn = d-&gt;block;
			return n;

		case DAUTO:
			if(debug[&#39;d&#39;])
				print(&#34;revert1 \&#34;%s\&#34;\n&#34;, s-&gt;name);
			if(s-&gt;aused == 0) {
				nearln = s-&gt;varlineno;
				if(s-&gt;class == CAUTO)
					warn(Z, &#34;auto declared and not used: %s&#34;, s-&gt;name);
				if(s-&gt;class == CPARAM)
					warn(Z, &#34;param declared and not used: %s&#34;, s-&gt;name);
			}
			if(s-&gt;type &amp;&amp; (s-&gt;type-&gt;garb &amp; GVOLATILE)) {
				n1 = new(ONAME, Z, Z);
				n1-&gt;sym = s;
				n1-&gt;type = s-&gt;type;
				n1-&gt;etype = TVOID;
				if(n1-&gt;type != T)
					n1-&gt;etype = n1-&gt;type-&gt;etype;
				n1-&gt;xoffset = s-&gt;offset;
				n1-&gt;class = s-&gt;class;

				n1 = new(OADDR, n1, Z);
				n1 = new(OUSED, n1, Z);
				if(n == Z)
					n = n1;
				else
					n = new(OLIST, n1, n);
			}
			s-&gt;type = d-&gt;type;
			s-&gt;class = d-&gt;class;
			s-&gt;offset = d-&gt;offset;
			s-&gt;block = d-&gt;block;
			s-&gt;varlineno = d-&gt;varlineno;
			s-&gt;aused = d-&gt;aused;
			break;

		case DSUE:
			if(debug[&#39;d&#39;])
				print(&#34;revert2 \&#34;%s\&#34;\n&#34;, s-&gt;name);
			s-&gt;suetag = d-&gt;type;
			s-&gt;sueblock = d-&gt;block;
			break;

		case DLABEL:
			if(debug[&#39;d&#39;])
				print(&#34;revert3 \&#34;%s\&#34;\n&#34;, s-&gt;name);
			if(s-&gt;label &amp;&amp; s-&gt;label-&gt;addable == 0)
				warn(s-&gt;label, &#34;label declared and not used \&#34;%s\&#34;&#34;, s-&gt;name);
			s-&gt;label = Z;
			break;
		}
	}
	return n;
}

Type*
fnproto(Node *n)
{
	int r;

	r = anyproto(n-&gt;right);
	if(r == 0 || (r &amp; OLDPROTO)) {
		if(r &amp; ANSIPROTO)
			diag(n, &#34;mixed ansi/old function declaration: %F&#34;, n-&gt;left);
		return T;
	}
	return fnproto1(n-&gt;right);
}

int
anyproto(Node *n)
{
	int r;

	r = 0;

loop:
	if(n == Z)
		return r;
	switch(n-&gt;op) {
	case OLIST:
		r |= anyproto(n-&gt;left);
		n = n-&gt;right;
		goto loop;

	case ODOTDOT:
	case OPROTO:
		return r | ANSIPROTO;
	}
	return r | OLDPROTO;
}

Type*
fnproto1(Node *n)
{
	Type *t;

	if(n == Z)
		return T;
	switch(n-&gt;op) {
	case OLIST:
		t = fnproto1(n-&gt;left);
		if(t != T)
			t-&gt;down = fnproto1(n-&gt;right);
		return t;

	case OPROTO:
		lastdcl = T;
		dodecl(NODECL, CXXX, n-&gt;type, n-&gt;left);
		t = typ(TXXX, T);
		if(lastdcl != T)
			*t = *paramconv(lastdcl, 1);
		return t;

	case ONAME:
		diag(n, &#34;incomplete argument prototype&#34;);
		return typ(TINT, T);

	case ODOTDOT:
		return typ(TDOT, T);
	}
	diag(n, &#34;unknown op in fnproto&#34;);
	return T;
}

void
dbgdecl(Sym *s)
{
	print(&#34;decl \&#34;%s\&#34;: C=%s [B=%d:O=%ld] T=%T\n&#34;,
		s-&gt;name, cnames[s-&gt;class], s-&gt;block, s-&gt;offset, s-&gt;type);
}

Decl*
push(void)
{
	Decl *d;

	d = alloc(sizeof(*d));
	d-&gt;link = dclstack;
	dclstack = d;
	return d;
}

Decl*
push1(Sym *s)
{
	Decl *d;

	d = push();
	d-&gt;sym = s;
	d-&gt;val = DAUTO;
	d-&gt;type = s-&gt;type;
	d-&gt;class = s-&gt;class;
	d-&gt;offset = s-&gt;offset;
	d-&gt;block = s-&gt;block;
	d-&gt;varlineno = s-&gt;varlineno;
	d-&gt;aused = s-&gt;aused;
	return d;
}

int
sametype(Type *t1, Type *t2)
{

	if(t1 == t2)
		return 1;
	return rsametype(t1, t2, 5, 1);
}

int
rsametype(Type *t1, Type *t2, int n, int f)
{
	int et;

	n--;
	for(;;) {
		if(t1 == t2)
			return 1;
		if(t1 == T || t2 == T)
			return 0;
		if(n &lt;= 0)
			return 1;
		et = t1-&gt;etype;
		if(et != t2-&gt;etype)
			return 0;
		if(et == TFUNC) {
			if(!rsametype(t1-&gt;link, t2-&gt;link, n, 0))
				return 0;
			t1 = t1-&gt;down;
			t2 = t2-&gt;down;
			while(t1 != T &amp;&amp; t2 != T) {
				if(t1-&gt;etype == TOLD) {
					t1 = t1-&gt;down;
					continue;
				}
				if(t2-&gt;etype == TOLD) {
					t2 = t2-&gt;down;
					continue;
				}
				while(t1 != T || t2 != T) {
					if(!rsametype(t1, t2, n, 0))
						return 0;
					t1 = t1-&gt;down;
					t2 = t2-&gt;down;
				}
				break;
			}
			return 1;
		}
		if(et == TARRAY)
			if(t1-&gt;width != t2-&gt;width &amp;&amp; t1-&gt;width != 0 &amp;&amp; t2-&gt;width != 0)
				return 0;
		if(typesu[et]) {
			if(t1-&gt;link == T)
				snap(t1);
			if(t2-&gt;link == T)
				snap(t2);
			t1 = t1-&gt;link;
			t2 = t2-&gt;link;
			for(;;) {
				if(t1 == t2)
					return 1;
				if(!rsametype(t1, t2, n, 0))
					return 0;
				t1 = t1-&gt;down;
				t2 = t2-&gt;down;
			}
		}
		t1 = t1-&gt;link;
		t2 = t2-&gt;link;
		if((f || !debug[&#39;V&#39;]) &amp;&amp; et == TIND) {
			if(t1 != T &amp;&amp; t1-&gt;etype == TVOID)
				return 1;
			if(t2 != T &amp;&amp; t2-&gt;etype == TVOID)
				return 1;
		}
	}
}

typedef struct Typetab Typetab;

struct Typetab{
	int n;
	Type **a;
};

static int
sigind(Type *t, Typetab *tt)
{
	int n;
	Type **a, **na, **p, **e;

	n = tt-&gt;n;
	a = tt-&gt;a;
	e = a+n;
	/* linear search seems ok */
	for(p = a ; p &lt; e; p++)
		if(sametype(*p, t))
			return p-a;
	if((n&amp;15) == 0){
		na = malloc((n+16)*sizeof(Type*));
		memmove(na, a, n*sizeof(Type*));
		free(a);
		a = tt-&gt;a = na;
	}
	a[tt-&gt;n++] = t;
	return -1;
}

static uint32
signat(Type *t, Typetab *tt)
{
	int i;
	Type *t1;
	int32 s;

	s = 0;
	for(; t; t=t-&gt;link) {
		s = s*thash1 + thash[t-&gt;etype];
		if(t-&gt;garb&amp;GINCOMPLETE)
			return s;
		switch(t-&gt;etype) {
		default:
			return s;
		case TARRAY:
			s = s*thash2 + 0;	/* was t-&gt;width */
			break;
		case TFUNC:
			for(t1=t-&gt;down; t1; t1=t1-&gt;down)
				s = s*thash3 + signat(t1, tt);
			break;
		case TSTRUCT:
		case TUNION:
			if((i = sigind(t, tt)) &gt;= 0){
				s = s*thash2 + i;
				return s;
			}
			for(t1=t-&gt;link; t1; t1=t1-&gt;down)
				s = s*thash3 + signat(t1, tt);
			return s;
		case TIND:
			break;
		}
	}
	return s;
}

uint32
signature(Type *t)
{
	uint32 s;
	Typetab tt;

	tt.n = 0;
	tt.a = nil;
	s = signat(t, &amp;tt);
	free(tt.a);
	return s;
}

uint32
sign(Sym *s)
{
	uint32 v;
	Type *t;

	if(s-&gt;sig == SIGINTERN)
		return SIGNINTERN;
	if((t = s-&gt;type) == T)
		return 0;
	v = signature(t);
	if(v == 0)
		v = SIGNINTERN;
	return v;
}

void
snap(Type *t)
{
	if(typesu[t-&gt;etype])
	if(t-&gt;link == T &amp;&amp; t-&gt;tag &amp;&amp; t-&gt;tag-&gt;suetag) {
		t-&gt;link = t-&gt;tag-&gt;suetag-&gt;link;
		t-&gt;width = t-&gt;tag-&gt;suetag-&gt;width;
	}
}

Type*
dotag(Sym *s, int et, int bn)
{
	Decl *d;

	if(bn != 0 &amp;&amp; bn != s-&gt;sueblock) {
		d = push();
		d-&gt;sym = s;
		d-&gt;val = DSUE;
		d-&gt;type = s-&gt;suetag;
		d-&gt;block = s-&gt;sueblock;
		s-&gt;suetag = T;
	}
	if(s-&gt;suetag == T) {
		s-&gt;suetag = typ(et, T);
		s-&gt;sueblock = autobn;
	}
	if(s-&gt;suetag-&gt;etype != et)
		diag(Z, &#34;tag used for more than one type: %s&#34;,
			s-&gt;name);
	if(s-&gt;suetag-&gt;tag == S)
		s-&gt;suetag-&gt;tag = s;
	return s-&gt;suetag;
}

Node*
dcllabel(Sym *s, int f)
{
	Decl *d, d1;
	Node *n;

	n = s-&gt;label;
	if(n != Z) {
		if(f) {
			if(n-&gt;complex)
				diag(Z, &#34;label reused: %s&#34;, s-&gt;name);
			n-&gt;complex = 1;	// declared
		} else
			n-&gt;addable = 1;	// used
		return n;
	}

	d = push();
	d-&gt;sym = s;
	d-&gt;val = DLABEL;
	dclstack = d-&gt;link;

	d1 = *firstdcl;
	*firstdcl = *d;
	*d = d1;

	firstdcl-&gt;link = d;
	firstdcl = d;

	n = new(OXXX, Z, Z);
	n-&gt;sym = s;
	n-&gt;complex = f;
	n-&gt;addable = !f;
	s-&gt;label = n;

	if(debug[&#39;d&#39;])
		dbgdecl(s);
	return n;
}

Type*
paramconv(Type *t, int f)
{

	switch(t-&gt;etype) {
	case TUNION:
	case TSTRUCT:
		if(t-&gt;width &lt;= 0)
			diag(Z, &#34;incomplete structure: %s&#34;, t-&gt;tag-&gt;name);
		break;

	case TARRAY:
		t = typ(TIND, t-&gt;link);
		t-&gt;width = types[TIND]-&gt;width;
		break;

	case TFUNC:
		t = typ(TIND, t);
		t-&gt;width = types[TIND]-&gt;width;
		break;

	case TFLOAT:
		if(!f)
			t = types[TDOUBLE];
		break;

	case TCHAR:
	case TSHORT:
		if(!f)
			t = types[TINT];
		break;

	case TUCHAR:
	case TUSHORT:
		if(!f)
			t = types[TUINT];
		break;
	}
	return t;
}

void
adecl(int c, Type *t, Sym *s)
{

	if(c == CSTATIC)
		c = CLOCAL;
	if(t-&gt;etype == TFUNC) {
		if(c == CXXX)
			c = CEXTERN;
		if(c == CLOCAL)
			c = CSTATIC;
		if(c == CAUTO || c == CEXREG)
			diag(Z, &#34;function cannot be %s %s&#34;, cnames[c], s-&gt;name);
	}
	if(c == CXXX)
		c = CAUTO;
	if(s) {
		if(s-&gt;class == CSTATIC)
			if(c == CEXTERN || c == CGLOBL) {
				warn(Z, &#34;just say static: %s&#34;, s-&gt;name);
				c = CSTATIC;
			}
		if(s-&gt;class == CAUTO || s-&gt;class == CPARAM || s-&gt;class == CLOCAL)
		if(s-&gt;block == autobn)
			diag(Z, &#34;auto redeclaration of: %s&#34;, s-&gt;name);
		if(c != CPARAM)
			push1(s);
		s-&gt;block = autobn;
		s-&gt;offset = 0;
		s-&gt;type = t;
		s-&gt;class = c;
		s-&gt;aused = 0;
	}
	switch(c) {
	case CAUTO:
		autoffset = align(autoffset, t, Aaut3);
		stkoff = maxround(stkoff, autoffset);
		s-&gt;offset = -autoffset;
		break;

	case CPARAM:
		if(autoffset == 0) {
			firstarg = s;
			firstargtype = t;
		}
		autoffset = align(autoffset, t, Aarg1);
		if(s)
			s-&gt;offset = autoffset;
		autoffset = align(autoffset, t, Aarg2);
		break;
	}
}

void
pdecl(int c, Type *t, Sym *s)
{
	if(s &amp;&amp; s-&gt;offset != -1) {
		diag(Z, &#34;not a parameter: %s&#34;, s-&gt;name);
		return;
	}
	t = paramconv(t, c==CPARAM);
	if(c == CXXX)
		c = CPARAM;
	if(c != CPARAM) {
		diag(Z, &#34;parameter cannot have class: %s&#34;, s-&gt;name);
		c = CPARAM;
	}
	adecl(c, t, s);
}

void
xdecl(int c, Type *t, Sym *s)
{
	int32 o;

	o = 0;
	switch(c) {
	case CEXREG:
		o = exreg(t);
		if(o == 0)
			c = CEXTERN;
		if(s-&gt;class == CGLOBL)
			c = CGLOBL;
		break;

	case CEXTERN:
		if(s-&gt;class == CGLOBL)
			c = CGLOBL;
		break;

	case CXXX:
		c = CGLOBL;
		if(s-&gt;class == CEXTERN)
			s-&gt;class = CGLOBL;
		break;

	case CAUTO:
		diag(Z, &#34;overspecified class: %s %s %s&#34;, s-&gt;name, cnames[c], cnames[s-&gt;class]);
		c = CEXTERN;
		break;

	case CTYPESTR:
		if(!typesuv[t-&gt;etype]) {
			diag(Z, &#34;typestr must be struct/union: %s&#34;, s-&gt;name);
			break;
		}
		dclfunct(t, s);
		break;
	}

	if(s-&gt;class == CSTATIC)
		if(c == CEXTERN || c == CGLOBL) {
			warn(Z, &#34;overspecified class: %s %s %s&#34;, s-&gt;name, cnames[c], cnames[s-&gt;class]);
			c = CSTATIC;
		}
	if(s-&gt;type != T)
		if(s-&gt;class != c || !sametype(t, s-&gt;type) || t-&gt;etype == TENUM) {
			diag(Z, &#34;external redeclaration of: %s&#34;, s-&gt;name);
			Bprint(&amp;diagbuf, &#34;	%s %T %L\n&#34;, cnames[c], t, nearln);
			Bprint(&amp;diagbuf, &#34;	%s %T %L\n&#34;, cnames[s-&gt;class], s-&gt;type, s-&gt;varlineno);
		}
	tmerge(t, s);
	s-&gt;type = t;
	s-&gt;class = c;
	s-&gt;block = 0;
	s-&gt;offset = o;
}

void
tmerge(Type *t1, Sym *s)
{
	Type *ta, *tb, *t2;

	t2 = s-&gt;type;
/*print(&#34;merge	%T; %T\n&#34;, t1, t2);/**/
	for(;;) {
		if(t1 == T || t2 == T || t1 == t2)
			break;
		if(t1-&gt;etype != t2-&gt;etype)
			break;
		switch(t1-&gt;etype) {
		case TFUNC:
			ta = t1-&gt;down;
			tb = t2-&gt;down;
			if(ta == T) {
				t1-&gt;down = tb;
				break;
			}
			if(tb == T)
				break;
			while(ta != T &amp;&amp; tb != T) {
				if(ta == tb)
					break;
				/* ignore old-style flag */
				if(ta-&gt;etype == TOLD) {
					ta = ta-&gt;down;
					continue;
				}
				if(tb-&gt;etype == TOLD) {
					tb = tb-&gt;down;
					continue;
				}
				/* checking terminated by ... */
				if(ta-&gt;etype == TDOT &amp;&amp; tb-&gt;etype == TDOT) {
					ta = T;
					tb = T;
					break;
				}
				if(!sametype(ta, tb))
					break;
				ta = ta-&gt;down;
				tb = tb-&gt;down;
			}
			if(ta != tb)
				diag(Z, &#34;function inconsistently declared: %s&#34;, s-&gt;name);

			/* take new-style over old-style */
			ta = t1-&gt;down;
			tb = t2-&gt;down;
			if(ta != T &amp;&amp; ta-&gt;etype == TOLD)
				if(tb != T &amp;&amp; tb-&gt;etype != TOLD)
					t1-&gt;down = tb;
			break;

		case TARRAY:
			/* should we check array size change? */
			if(t2-&gt;width &gt; t1-&gt;width)
				t1-&gt;width = t2-&gt;width;
			break;

		case TUNION:
		case TSTRUCT:
			return;
		}
		t1 = t1-&gt;link;
		t2 = t2-&gt;link;
	}
}

void
edecl(int c, Type *t, Sym *s)
{
	Type *t1;

	if(s == S) {
		if(!typesu[t-&gt;etype])
			diag(Z, &#34;unnamed structure element must be struct/union&#34;);
		if(c != CXXX)
			diag(Z, &#34;unnamed structure element cannot have class&#34;);
	} else
		if(c != CXXX)
			diag(Z, &#34;structure element cannot have class: %s&#34;, s-&gt;name);
	t1 = t;
	t = copytyp(t1);
	t-&gt;sym = s;
	t-&gt;down = T;
	if(lastfield) {
		t-&gt;shift = lastbit - lastfield;
		t-&gt;nbits = lastfield;
		if(firstbit)
			t-&gt;shift = -t-&gt;shift;
		if(typeu[t-&gt;etype])
			t-&gt;etype = tufield-&gt;etype;
		else
			t-&gt;etype = tfield-&gt;etype;
	}
	if(strf == T)
		strf = t;
	else
		strl-&gt;down = t;
	strl = t;
}

/*
 * this routine is very suspect.
 * ansi requires the enum type to
 * be represented as an &#39;int&#39;
 * this means that 0x81234567
 * would be illegal. this routine
 * makes signed and unsigned go
 * to unsigned.
 */
Type*
maxtype(Type *t1, Type *t2)
{

	if(t1 == T)
		return t2;
	if(t2 == T)
		return t1;
	if(t1-&gt;etype &gt; t2-&gt;etype)
		return t1;
	return t2;
}

void
doenum(Sym *s, Node *n)
{

	if(n) {
		complex(n);
		if(n-&gt;op != OCONST) {
			diag(n, &#34;enum not a constant: %s&#34;, s-&gt;name);
			return;
		}
		en.cenum = n-&gt;type;
		en.tenum = maxtype(en.cenum, en.tenum);

		if(!typefd[en.cenum-&gt;etype])
			en.lastenum = n-&gt;vconst;
		else
			en.floatenum = n-&gt;fconst;
	}
	if(dclstack)
		push1(s);
	xdecl(CXXX, types[TENUM], s);

	if(en.cenum == T) {
		en.tenum = types[TINT];
		en.cenum = types[TINT];
		en.lastenum = 0;
	}
	s-&gt;tenum = en.cenum;

	if(!typefd[s-&gt;tenum-&gt;etype]) {
		s-&gt;vconst = convvtox(en.lastenum, s-&gt;tenum-&gt;etype);
		en.lastenum++;
	} else {
		s-&gt;fconst = en.floatenum;
		en.floatenum++;
	}

	if(debug[&#39;d&#39;])
		dbgdecl(s);
	acidvar(s);
}

void
symadjust(Sym *s, Node *n, int32 del)
{

	switch(n-&gt;op) {
	default:
		if(n-&gt;left)
			symadjust(s, n-&gt;left, del);
		if(n-&gt;right)
			symadjust(s, n-&gt;right, del);
		return;

	case ONAME:
		if(n-&gt;sym == s)
			n-&gt;xoffset -= del;
		return;

	case OCONST:
	case OSTRING:
	case OLSTRING:
	case OINDREG:
	case OREGISTER:
		return;
	}
}

Node*
contig(Sym *s, Node *n, int32 v)
{
	Node *p, *r, *q, *m;
	int32 w;
	Type *zt;

	if(debug[&#39;i&#39;]) {
		print(&#34;contig v = %ld; s = %s\n&#34;, v, s-&gt;name);
		prtree(n, &#34;doinit value&#34;);
	}

	if(n == Z)
		goto no;
	w = s-&gt;type-&gt;width;

	/*
	 * nightmare: an automatic array whose size
	 * increases when it is initialized
	 */
	if(v != w) {
		if(v != 0)
			diag(n, &#34;automatic adjustable array: %s&#34;, s-&gt;name);
		v = s-&gt;offset;
		autoffset = align(autoffset, s-&gt;type, Aaut3);
		s-&gt;offset = -autoffset;
		stkoff = maxround(stkoff, autoffset);
		symadjust(s, n, v - s-&gt;offset);
	}
	if(w &lt;= ewidth[TIND])
		goto no;
	if(n-&gt;op == OAS)
		diag(Z, &#34;oops in contig&#34;);
/*ZZZ this appears incorrect
need to check if the list completely covers the data.
if not, bail
 */
	if(n-&gt;op == OLIST)
		goto no;
	if(n-&gt;op == OASI)
		if(n-&gt;left-&gt;type)
		if(n-&gt;left-&gt;type-&gt;width == w)
			goto no;
	while(w &amp; (ewidth[TIND]-1))
		w++;
/*
 * insert the following code, where long becomes vlong if pointers are fat
 *
	*(long**)&amp;X = (long*)((char*)X + sizeof(X));
	do {
		*(long**)&amp;X -= 1;
		**(long**)&amp;X = 0;
	} while(*(long**)&amp;X);
 */

	for(q=n; q-&gt;op != ONAME; q=q-&gt;left)
		;

	zt = ewidth[TIND] &gt; ewidth[TLONG]? types[TVLONG]: types[TLONG];

	p = new(ONAME, Z, Z);
	*p = *q;
	p-&gt;type = typ(TIND, zt);
	p-&gt;xoffset = s-&gt;offset;

	r = new(ONAME, Z, Z);
	*r = *p;
	r = new(OPOSTDEC, r, Z);

	q = new(ONAME, Z, Z);
	*q = *p;
	q = new(OIND, q, Z);

	m = new(OCONST, Z, Z);
	m-&gt;vconst = 0;
	m-&gt;type = zt;

	q = new(OAS, q, m);

	r = new(OLIST, r, q);

	q = new(ONAME, Z, Z);
	*q = *p;
	r = new(ODWHILE, q, r);

	q = new(ONAME, Z, Z);
	*q = *p;
	q-&gt;type = q-&gt;type-&gt;link;
	q-&gt;xoffset += w;
	q = new(OADDR, q, 0);

	q = new(OASI, p, q);
	r = new(OLIST, q, r);

	n = new(OLIST, r, n);

no:
	return n;
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
