<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8c/div.c</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/8c/div.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/8c/div.c
// http://code.google.com/p/inferno-os/source/browse/utils/8c/div.c
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

/*
 * Based on: Granlund, T.; Montgomery, P.L.
 * &#34;Division by Invariant Integers using Multiplication&#34;.
 * SIGPLAN Notices, Vol. 29, June 1994, page 61.
 */

#define	TN(n)	((uvlong)1 &lt;&lt; (n))
#define	T31	TN(31)
#define	T32	TN(32)

int
multiplier(uint32 d, int p, uvlong *mp)
{
	int l;
	uvlong mlo, mhi, tlo, thi;

	l = topbit(d - 1) + 1;
	mlo = (((TN(l) - d) &lt;&lt; 32) / d) + T32;
	if(l + p == 64)
		mhi = (((TN(l) + 1 - d) &lt;&lt; 32) / d) + T32;
	else
		mhi = (TN(32 + l) + TN(32 + l - p)) / d;
	/*assert(mlo &lt; mhi);*/
	while(l &gt; 0) {
		tlo = mlo &gt;&gt; 1;
		thi = mhi &gt;&gt; 1;
		if(tlo == thi)
			break;
		mlo = tlo;
		mhi = thi;
		l--;
	}
	*mp = mhi;
	return l;
}

int
sdiv(uint32 d, uint32 *mp, int *sp)
{
	int s;
	uvlong m;

	s = multiplier(d, 32 - 1, &amp;m);
	*mp = m;
	*sp = s;
	if(m &gt;= T31)
		return 1;
	else
		return 0;
}

int
udiv(uint32 d, uint32 *mp, int *sp, int *pp)
{
	int p, s;
	uvlong m;

	s = multiplier(d, 32, &amp;m);
	p = 0;
	if(m &gt;= T32) {
		while((d &amp; 1) == 0) {
			d &gt;&gt;= 1;
			p++;
		}
		s = multiplier(d, 32 - p, &amp;m);
	}
	*mp = m;
	*pp = p;
	if(m &gt;= T32) {
		/*assert(p == 0);*/
		*sp = s - 1;
		return 1;
	}
	else {
		*sp = s;
		return 0;
	}
}

void
sdivgen(Node *l, Node *r, Node *ax, Node *dx)
{
	int a, s;
	uint32 m;
	vlong c;

	c = r-&gt;vconst;
	if(c &lt; 0)
		c = -c;
	a = sdiv(c, &amp;m, &amp;s);
//print(&#34;a=%d i=%ld s=%d m=%lux\n&#34;, a, (int32)r-&gt;vconst, s, m);
	gins(AMOVL, nodconst(m), ax);
	gins(AIMULL, l, Z);
	gins(AMOVL, l, ax);
	if(a)
		gins(AADDL, ax, dx);
	gins(ASHRL, nodconst(31), ax);
	gins(ASARL, nodconst(s), dx);
	gins(AADDL, ax, dx);
	if(r-&gt;vconst &lt; 0)
		gins(ANEGL, Z, dx);
}

void
udivgen(Node *l, Node *r, Node *ax, Node *dx)
{
	int a, s, t;
	uint32 m;
	Node nod;

	a = udiv(r-&gt;vconst, &amp;m, &amp;s, &amp;t);
//print(&#34;a=%ud i=%ld p=%d s=%d m=%lux\n&#34;, a, (int32)r-&gt;vconst, t, s, m);
	if(t != 0) {
		gins(AMOVL, l, ax);
		gins(ASHRL, nodconst(t), ax);
		gins(AMOVL, nodconst(m), dx);
		gins(AMULL, dx, Z);
	}
	else if(a) {
		if(l-&gt;op != OREGISTER) {
			regalloc(&amp;nod, l, Z);
			gins(AMOVL, l, &amp;nod);
			l = &amp;nod;
		}
		gins(AMOVL, nodconst(m), ax);
		gins(AMULL, l, Z);
		gins(AADDL, l, dx);
		gins(ARCRL, nodconst(1), dx);
		if(l == &amp;nod)
			regfree(l);
	}
	else {
		gins(AMOVL, nodconst(m), ax);
		gins(AMULL, l, Z);
	}
	if(s != 0)
		gins(ASHRL, nodconst(s), dx);
}

void
sext(Node *d, Node *s, Node *l)
{
	if(s-&gt;reg == D_AX &amp;&amp; !nodreg(d, Z, D_DX)) {
		reg[D_DX]++;
		gins(ACDQ, Z, Z);
	}
	else {
		regalloc(d, l, Z);
		gins(AMOVL, s, d);
		gins(ASARL, nodconst(31), d);
	}
}

void
sdiv2(int32 c, int v, Node *l, Node *n)
{
	Node nod;

	if(v &gt; 0) {
		if(v &gt; 1) {
			sext(&amp;nod, n, l);
			gins(AANDL, nodconst((1 &lt;&lt; v) - 1), &amp;nod);
			gins(AADDL, &amp;nod, n);
			regfree(&amp;nod);
		}
		else {
			gins(ACMPL, n, nodconst(0x80000000));
			gins(ASBBL, nodconst(-1), n);
		}
		gins(ASARL, nodconst(v), n);
	}
	if(c &lt; 0)
		gins(ANEGL, Z, n);
}

void
smod2(int32 c, int v, Node *l, Node *n)
{
	Node nod;

	if(c == 1) {
		zeroregm(n);
		return;
	}

	sext(&amp;nod, n, l);
	if(v == 0) {
		zeroregm(n);
		gins(AXORL, &amp;nod, n);
		gins(ASUBL, &amp;nod, n);
	}
	else if(v &gt; 1) {
		gins(AANDL, nodconst((1 &lt;&lt; v) - 1), &amp;nod);
		gins(AADDL, &amp;nod, n);
		gins(AANDL, nodconst((1 &lt;&lt; v) - 1), n);
		gins(ASUBL, &amp;nod, n);
	}
	else {
		gins(AANDL, nodconst(1), n);
		gins(AXORL, &amp;nod, n);
		gins(ASUBL, &amp;nod, n);
	}
	regfree(&amp;nod);
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
