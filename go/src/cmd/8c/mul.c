<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8c/mul.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/8c/mul.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/8c/mul.c
// http://code.google.com/p/inferno-os/source/browse/utils/8c/mul.c
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

typedef struct	Malg	Malg;
typedef struct	Mparam	Mparam;

struct	Malg
{
	char	vals[10];
};

struct	Mparam
{
	uint32	value;
	char	alg;
	char	neg;
	char	shift;
	char	arg;
	char	off;
};

static	Mparam	multab[32];
static	int	mulptr;

static	Malg	malgs[]	=
{
	{0, 100},
	{-1, 1, 100},
	{-9, -5, -3, 3, 5, 9, 100},
	{6, 10, 12, 18, 20, 24, 36, 40, 72, 100},
	{-8, -4, -2, 2, 4, 8, 100},
};

/*
 * return position of lowest 1
 */
int
lowbit(uint32 v)
{
	int s, i;
	uint32 m;

	s = 0;
	m = 0xFFFFFFFFUL;
	for(i = 16; i &gt; 0; i &gt;&gt;= 1) {
		m &gt;&gt;= i;
		if((v &amp; m) == 0) {
			v &gt;&gt;= i;
			s += i;
		}
	}
	return s;
}

void
genmuladd(Node *d, Node *s, int m, Node *a)
{
	Node nod;

	nod.op = OINDEX;
	nod.left = a;
	nod.right = s;
	nod.scale = m;
	nod.type = types[TIND];
	nod.xoffset = 0;
	xcom(&amp;nod);
	gopcode(OADDR, d-&gt;type, &amp;nod, d);
}

void
mulparam(uint32 m, Mparam *mp)
{
	int c, i, j, n, o, q, s;
	int bc, bi, bn, bo, bq, bs, bt;
	char *p;
	int32 u;
	uint32 t;

	bc = bq = 10;
	bi = bn = bo = bs = bt = 0;
	for(i = 0; i &lt; nelem(malgs); i++) {
		for(p = malgs[i].vals, j = 0; (o = p[j]) &lt; 100; j++)
		for(s = 0; s &lt; 2; s++) {
			c = 10;
			q = 10;
			u = m - o;
			if(u == 0)
				continue;
			if(s) {
				o = -o;
				if(o &gt; 0)
					continue;
				u = -u;
			}
			n = lowbit(u);
			t = (uint32)u &gt;&gt; n;
			switch(i) {
			case 0:
				if(t == 1) {
					c = s + 1;
					q = 0;
					break;
				}
				switch(t) {
				case 3:
				case 5:
				case 9:
					c = s + 1;
					if(n)
						c++;
					q = 0;
					break;
				}
				if(s)
					break;
				switch(t) {
				case 15:
				case 25:
				case 27:
				case 45:
				case 81:
					c = 2;
					if(n)
						c++;
					q = 1;
					break;
				}
				break;
			case 1:
				if(t == 1) {
					c = 3;
					q = 3;
					break;
				}
				switch(t) {
				case 3:
				case 5:
				case 9:
					c = 3;
					q = 2;
					break;
				}
				break;
			case 2:
				if(t == 1) {
					c = 3;
					q = 2;
					break;
				}
				break;
			case 3:
				if(s)
					break;
				if(t == 1) {
					c = 3;
					q = 1;
					break;
				}
				break;
			case 4:
				if(t == 1) {
					c = 3;
					q = 0;
					break;
				}
				break;
			}
			if(c &lt; bc || (c == bc &amp;&amp; q &gt; bq)) {
				bc = c;
				bi = i;
				bn = n;
				bo = o;
				bq = q;
				bs = s;
				bt = t;
			}
		}
	}
	mp-&gt;value = m;
	if(bc &lt;= 3) {
		mp-&gt;alg = bi;
		mp-&gt;shift = bn;
		mp-&gt;off = bo;
		mp-&gt;neg = bs;
		mp-&gt;arg = bt;
	}
	else
		mp-&gt;alg = -1;
}

int
m0(int a)
{
	switch(a) {
	case -2:
	case 2:
		return 2;
	case -3:
	case 3:
		return 2;
	case -4:
	case 4:
		return 4;
	case -5:
	case 5:
		return 4;
	case 6:
		return 2;
	case -8:
	case 8:
		return 8;
	case -9:
	case 9:
		return 8;
	case 10:
		return 4;
	case 12:
		return 2;
	case 15:
		return 2;
	case 18:
		return 8;
	case 20:
		return 4;
	case 24:
		return 2;
	case 25:
		return 4;
	case 27:
		return 2;
	case 36:
		return 8;
	case 40:
		return 4;
	case 45:
		return 4;
	case 72:
		return 8;
	case 81:
		return 8;
	}
	diag(Z, &#34;bad m0&#34;);
	return 0;
}

int
m1(int a)
{
	switch(a) {
	case 15:
		return 4;
	case 25:
		return 4;
	case 27:
		return 8;
	case 45:
		return 8;
	case 81:
		return 8;
	}
	diag(Z, &#34;bad m1&#34;);
	return 0;
}

int
m2(int a)
{
	switch(a) {
	case 6:
		return 2;
	case 10:
		return 2;
	case 12:
		return 4;
	case 18:
		return 2;
	case 20:
		return 4;
	case 24:
		return 8;
	case 36:
		return 4;
	case 40:
		return 8;
	case 72:
		return 8;
	}
	diag(Z, &#34;bad m2&#34;);
	return 0;
}

void
shiftit(Type *t, Node *s, Node *d)
{
	int32 c;

	c = (int32)s-&gt;vconst &amp; 31;
	switch(c) {
	case 0:
		break;
	case 1:
		gopcode(OADD, t, d, d);
		break;
	default:
		gopcode(OASHL, t, s, d);
	}
}

static int
mulgen1(uint32 v, Node *n)
{
	int i, o;
	Mparam *p;
	Node nod, nods;

	for(i = 0; i &lt; nelem(multab); i++) {
		p = &amp;multab[i];
		if(p-&gt;value == v)
			goto found;
	}

	p = &amp;multab[mulptr];
	if(++mulptr == nelem(multab))
		mulptr = 0;

	mulparam(v, p);

found:
//	print(&#34;v=%.lx a=%d n=%d s=%d g=%d o=%d \n&#34;, p-&gt;value, p-&gt;alg, p-&gt;neg, p-&gt;shift, p-&gt;arg, p-&gt;off);
	if(p-&gt;alg &lt; 0)
		return 0;

	nods = *nodconst(p-&gt;shift);

	o = OADD;
	if(p-&gt;alg &gt; 0) {
		regalloc(&amp;nod, n, Z);
		if(p-&gt;off &lt; 0)
			o = OSUB;
	}

	switch(p-&gt;alg) {
	case 0:
		switch(p-&gt;arg) {
		case 1:
			shiftit(n-&gt;type, &amp;nods, n);
			break;
		case 15:
		case 25:
		case 27:
		case 45:
		case 81:
			genmuladd(n, n, m1(p-&gt;arg), n);
			/* fall thru */
		case 3:
		case 5:
		case 9:
			genmuladd(n, n, m0(p-&gt;arg), n);
			shiftit(n-&gt;type, &amp;nods, n);
			break;
		default:
			goto bad;
		}
		if(p-&gt;neg == 1)
			gins(ANEGL, Z, n);
		break;
	case 1:
		switch(p-&gt;arg) {
		case 1:
			gmove(n, &amp;nod);
			shiftit(n-&gt;type, &amp;nods, &amp;nod);
			break;
		case 3:
		case 5:
		case 9:
			genmuladd(&amp;nod, n, m0(p-&gt;arg), n);
			shiftit(n-&gt;type, &amp;nods, &amp;nod);
			break;
		default:
			goto bad;
		}
		if(p-&gt;neg)
			gopcode(o, n-&gt;type, &amp;nod, n);
		else {
			gopcode(o, n-&gt;type, n, &amp;nod);
			gmove(&amp;nod, n);
		}
		break;
	case 2:
		genmuladd(&amp;nod, n, m0(p-&gt;off), n);
		shiftit(n-&gt;type, &amp;nods, n);
		goto comop;
	case 3:
		genmuladd(&amp;nod, n, m0(p-&gt;off), n);
		shiftit(n-&gt;type, &amp;nods, n);
		genmuladd(n, &amp;nod, m2(p-&gt;off), n);
		break;
	case 4:
		genmuladd(&amp;nod, n, m0(p-&gt;off), nodconst(0));
		shiftit(n-&gt;type, &amp;nods, n);
		goto comop;
	default:
		diag(Z, &#34;bad mul alg&#34;);
		break;
	comop:
		if(p-&gt;neg) {
			gopcode(o, n-&gt;type, n, &amp;nod);
			gmove(&amp;nod, n);
		}
		else
			gopcode(o, n-&gt;type, &amp;nod, n);
	}

	if(p-&gt;alg &gt; 0)
		regfree(&amp;nod);

	return 1;

bad:
	diag(Z, &#34;mulgen botch&#34;);
	return 1;
}

void
mulgen(Type *t, Node *r, Node *n)
{
	if(!mulgen1(r-&gt;vconst, n))
		gopcode(OMUL, t, r, n);
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
