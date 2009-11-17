<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/mparith2.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/mparith2.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include	&#34;go.h&#34;

//
// return the significant
// words of the argument
//
static int
mplen(Mpint *a)
{
	int i, n;
	long *a1;

	n = -1;
	a1 = &amp;a-&gt;a[0];
	for(i=0; i&lt;Mpprec; i++) {
		if(*a1++ != 0)
			n = i;
	}
	return n+1;
}

//
// left shift mpint by one
// ignores sign and overflow
//
static void
mplsh(Mpint *a)
{
	long *a1, x;
	int i, c;

	c = 0;
	a1 = &amp;a-&gt;a[0];
	for(i=0; i&lt;Mpprec; i++) {
		x = (*a1 &lt;&lt; 1) + c;
		c = 0;
		if(x &gt;= Mpbase) {
			x -= Mpbase;
			c = 1;
		}
		*a1++ = x;
	}
}

//
// left shift mpint by Mpscale
// ignores sign and overflow
//
static void
mplshw(Mpint *a)
{
	long *a1;
	int i;

	a1 = &amp;a-&gt;a[Mpprec-1];
	for(i=1; i&lt;Mpprec; i++) {
		a1[0] = a1[-1];
		a1--;
	}
	a1[0] = 0;
}

//
// right shift mpint by one
// ignores sign and overflow
//
static void
mprsh(Mpint *a)
{
	long *a1, x, lo;
	int i, c;

	c = 0;
	lo = a-&gt;a[0] &amp; 1;
	a1 = &amp;a-&gt;a[Mpprec];
	for(i=0; i&lt;Mpprec; i++) {
		x = *--a1;
		*a1 = (x + c) &gt;&gt; 1;
		c = 0;
		if(x &amp; 1)
			c = Mpbase;
	}
	if(a-&gt;neg &amp;&amp; lo != 0)
		mpaddcfix(a, -1);
}

//
// right shift mpint by Mpscale
// ignores sign and overflow
//
static void
mprshw(Mpint *a)
{
	long *a1, lo;
	int i;

	lo = a-&gt;a[0];
	a1 = &amp;a-&gt;a[0];
	for(i=1; i&lt;Mpprec; i++) {
		a1[0] = a1[1];
		a1++;
	}
	a1[0] = 0;
	if(a-&gt;neg &amp;&amp; lo != 0)
		mpaddcfix(a, -1);
}

//
// return the sign of (abs(a)-abs(b))
//
static int
mpcmp(Mpint *a, Mpint *b)
{
	long x, *a1, *b1;
	int i;

	if(a-&gt;ovf || b-&gt;ovf) {
		yyerror(&#34;ovf in cmp&#34;);
		return 0;
	}

	a1 = &amp;a-&gt;a[0] + Mpprec;
	b1 = &amp;b-&gt;a[0] + Mpprec;

	for(i=0; i&lt;Mpprec; i++) {
		x = *--a1 - *--b1;
		if(x &gt; 0)
			return +1;
		if(x &lt; 0)
			return -1;
	}
	return 0;
}

//
// negate a
// ignore sign and ovf
//
static void
mpneg(Mpint *a)
{
	long x, *a1;
	int i, c;

	a1 = &amp;a-&gt;a[0];
	c = 0;
	for(i=0; i&lt;Mpprec; i++) {
		x = -*a1 -c;
		c = 0;
		if(x &lt; 0) {
			x += Mpbase;
			c = 1;
		}
		*a1++ = x;
	}
}

// shift left by s (or right by -s)
void
mpshiftfix(Mpint *a, int s)
{
	if(s &gt;= 0) {
		while(s &gt;= Mpscale) {
			mplshw(a);
			s -= Mpscale;
		}
		while(s &gt; 0) {
			mplsh(a);
			s--;
		}
	} else {
		s = -s;
		while(s &gt;= Mpscale) {
			mprshw(a);
			s -= Mpscale;
		}
		while(s &gt; 0) {
			mprsh(a);
			s--;
		}
	}
}

/// implements fix arihmetic

void
mpaddfixfix(Mpint *a, Mpint *b)
{
	int i, c;
	long x, *a1, *b1;

	if(a-&gt;ovf || b-&gt;ovf) {
		yyerror(&#34;ovf in mpaddxx&#34;);
		a-&gt;ovf = 1;
		return;
	}

	c = 0;
	a1 = &amp;a-&gt;a[0];
	b1 = &amp;b-&gt;a[0];
	if(a-&gt;neg != b-&gt;neg)
		goto sub;

	// perform a+b
	for(i=0; i&lt;Mpprec; i++) {
		x = *a1 + *b1++ + c;
		c = 0;
		if(x &gt;= Mpbase) {
			x -= Mpbase;
			c = 1;
		}
		*a1++ = x;
	}
	a-&gt;ovf = c;
	if(a-&gt;ovf)
		yyerror(&#34;set ovf in mpaddxx&#34;);

	return;

sub:
	// perform a-b
	switch(mpcmp(a, b)) {
	case 0:
		mpmovecfix(a, 0);
		break;

	case 1:
		for(i=0; i&lt;Mpprec; i++) {
			x = *a1 - *b1++ - c;
			c = 0;
			if(x &lt; 0) {
				x += Mpbase;
				c = 1;
			}
			*a1++ = x;
		}
		break;

	case -1:
		a-&gt;neg ^= 1;
		for(i=0; i&lt;Mpprec; i++) {
			x = *b1++ - *a1 - c;
			c = 0;
			if(x &lt; 0) {
				x += Mpbase;
				c = 1;
			}
			*a1++ = x;
		}
		break;
	}
}

void
mpmulfixfix(Mpint *a, Mpint *b)
{

	int i, j, na, nb;
	long *a1, x;
	Mpint s, q;

	if(a-&gt;ovf || b-&gt;ovf) {
		yyerror(&#34;ovf in mpmulfixfix&#34;);
		a-&gt;ovf = 1;
		return;
	}

	// pick the smaller
	// to test for bits
	na = mplen(a);
	nb = mplen(b);
	if(na &gt; nb) {
		mpmovefixfix(&amp;s, a);
		a1 = &amp;b-&gt;a[0];
		na = nb;
	} else {
		mpmovefixfix(&amp;s, b);
		a1 = &amp;a-&gt;a[0];
	}
	s.neg = 0;

	mpmovecfix(&amp;q, 0);
	for(i=0; i&lt;na; i++) {
		x = *a1++;
		for(j=0; j&lt;Mpscale; j++) {
			if(x &amp; 1)
				mpaddfixfix(&amp;q, &amp;s);
			mplsh(&amp;s);
			x &gt;&gt;= 1;
		}
	}

	q.neg = a-&gt;neg ^ b-&gt;neg;
	mpmovefixfix(a, &amp;q);
	if(a-&gt;ovf)
		yyerror(&#34;set ovf in mpmulfixfix&#34;);
}

void
mpmulfract(Mpint *a, Mpint *b)
{

	int i, j;
	long *a1, x;
	Mpint s, q;

	if(a-&gt;ovf || b-&gt;ovf) {
		yyerror(&#34;ovf in mpmulflt&#34;);
		a-&gt;ovf = 1;
		return;
	}

	mpmovefixfix(&amp;s, b);
	a1 = &amp;a-&gt;a[Mpprec];
	s.neg = 0;
	mpmovecfix(&amp;q, 0);

	for(i=0; i&lt;Mpprec; i++) {
		x = *--a1;
		if(x == 0) {
			mprshw(&amp;s);
			continue;
		}
		for(j=0; j&lt;Mpscale; j++) {
			x &lt;&lt;= 1;
			if(x &amp; Mpbase)
				mpaddfixfix(&amp;q, &amp;s);
			mprsh(&amp;s);
		}
	}

	q.neg = a-&gt;neg ^ b-&gt;neg;
	mpmovefixfix(a, &amp;q);
	if(a-&gt;ovf)
		yyerror(&#34;set ovf in mpmulflt&#34;);
}

void
mporfixfix(Mpint *a, Mpint *b)
{
	int i;
	long x, *a1, *b1;

	if(a-&gt;ovf || b-&gt;ovf) {
		yyerror(&#34;ovf in mporfixfix&#34;);
		mpmovecfix(a, 0);
		a-&gt;ovf = 1;
		return;
	}
	if(a-&gt;neg) {
		a-&gt;neg = 0;
		mpneg(a);
	}
	if(b-&gt;neg)
		mpneg(b);

	a1 = &amp;a-&gt;a[0];
	b1 = &amp;b-&gt;a[0];
	for(i=0; i&lt;Mpprec; i++) {
		x = *a1 | *b1++;
		*a1++ = x;
	}

	if(b-&gt;neg)
		mpneg(b);
	if(x &amp; Mpsign) {
		a-&gt;neg = 1;
		mpneg(a);
	}
}

void
mpandfixfix(Mpint *a, Mpint *b)
{
	int i;
	long x, *a1, *b1;

	if(a-&gt;ovf || b-&gt;ovf) {
		yyerror(&#34;ovf in mpandfixfix&#34;);
		mpmovecfix(a, 0);
		a-&gt;ovf = 1;
		return;
	}
	if(a-&gt;neg) {
		a-&gt;neg = 0;
		mpneg(a);
	}
	if(b-&gt;neg)
		mpneg(b);

	a1 = &amp;a-&gt;a[0];
	b1 = &amp;b-&gt;a[0];
	for(i=0; i&lt;Mpprec; i++) {
		x = *a1 &amp; *b1++;
		*a1++ = x;
	}

	if(b-&gt;neg)
		mpneg(b);
	if(x &amp; Mpsign) {
		a-&gt;neg = 1;
		mpneg(a);
	}
}

void
mpandnotfixfix(Mpint *a, Mpint *b)
{
	int i;
	long x, *a1, *b1;

	if(a-&gt;ovf || b-&gt;ovf) {
		yyerror(&#34;ovf in mpandnotfixfix&#34;);
		mpmovecfix(a, 0);
		a-&gt;ovf = 1;
		return;
	}
	if(a-&gt;neg) {
		a-&gt;neg = 0;
		mpneg(a);
	}
	if(b-&gt;neg)
		mpneg(b);

	a1 = &amp;a-&gt;a[0];
	b1 = &amp;b-&gt;a[0];
	for(i=0; i&lt;Mpprec; i++) {
		x = *a1 &amp; ~*b1++;
		*a1++ = x;
	}

	if(b-&gt;neg)
		mpneg(b);
	if(x &amp; Mpsign) {
		a-&gt;neg = 1;
		mpneg(a);
	}
}

void
mpxorfixfix(Mpint *a, Mpint *b)
{
	int i;
	long x, *a1, *b1;

	if(a-&gt;ovf || b-&gt;ovf) {
		yyerror(&#34;ovf in mporfixfix&#34;);
		mpmovecfix(a, 0);
		a-&gt;ovf = 1;
		return;
	}
	if(a-&gt;neg) {
		a-&gt;neg = 0;
		mpneg(a);
	}
	if(b-&gt;neg)
		mpneg(b);

	a1 = &amp;a-&gt;a[0];
	b1 = &amp;b-&gt;a[0];
	for(i=0; i&lt;Mpprec; i++) {
		x = *a1 ^ *b1++;
		*a1++ = x;
	}

	if(b-&gt;neg)
		mpneg(b);
	if(x &amp; Mpsign) {
		a-&gt;neg = 1;
		mpneg(a);
	}
}

void
mplshfixfix(Mpint *a, Mpint *b)
{
	vlong s;

	if(a-&gt;ovf || b-&gt;ovf) {
		yyerror(&#34;ovf in mporfixfix&#34;);
		mpmovecfix(a, 0);
		a-&gt;ovf = 1;
		return;
	}
	s = mpgetfix(b);
	if(s &lt; 0 || s &gt;= Mpprec*Mpscale) {
		yyerror(&#34;stupid shift: %lld&#34;, s);
		mpmovecfix(a, 0);
		return;
	}

	mpshiftfix(a, s);
}

void
mprshfixfix(Mpint *a, Mpint *b)
{
	vlong s;

	if(a-&gt;ovf || b-&gt;ovf) {
		yyerror(&#34;ovf in mprshfixfix&#34;);
		mpmovecfix(a, 0);
		a-&gt;ovf = 1;
		return;
	}
	s = mpgetfix(b);
	if(s &lt; 0 || s &gt;= Mpprec*Mpscale) {
		yyerror(&#34;stupid shift: %lld&#34;, s);
		if(a-&gt;neg)
			mpmovecfix(a, -1);
		else
			mpmovecfix(a, 0);
		return;
	}

	mpshiftfix(a, -s);
}

void
mpnegfix(Mpint *a)
{
	a-&gt;neg ^= 1;
}

vlong
mpgetfix(Mpint *a)
{
	vlong v;

	if(a-&gt;ovf) {
		yyerror(&#34;ovf in mpgetfix&#34;);
		return 0;
	}

	v = (vlong)a-&gt;a[0];
	v |= (vlong)a-&gt;a[1] &lt;&lt; Mpscale;
	v |= (vlong)a-&gt;a[2] &lt;&lt; (Mpscale+Mpscale);
	if(a-&gt;neg)
		v = -v;
	return v;
}

void
mpmovecfix(Mpint *a, vlong c)
{
	int i;
	long *a1;
	vlong x;

	a-&gt;neg = 0;
	a-&gt;ovf = 0;

	x = c;
	if(x &lt; 0) {
		a-&gt;neg = 1;
		x = -x;
	}

	a1 = &amp;a-&gt;a[0];
	for(i=0; i&lt;Mpprec; i++) {
		*a1++ = x&amp;Mpmask;
		x &gt;&gt;= Mpscale;
	}
}

void
mpdivmodfixfix(Mpint *q, Mpint *r, Mpint *n, Mpint *d)
{
	int i, ns, ds;

	ns = n-&gt;neg;
	ds = d-&gt;neg;
	n-&gt;neg = 0;
	d-&gt;neg = 0;

	mpmovefixfix(r, n);
	mpmovecfix(q, 0);

	// shift denominator until it
	// is larger than numerator
	for(i=0; i&lt;Mpprec*Mpscale; i++) {
		if(mpcmp(d, r) &gt; 0)
			break;
		mplsh(d);
	}

	// if it never happens
	// denominator is probably zero
	if(i &gt;= Mpprec*Mpscale) {
		q-&gt;ovf = 1;
		r-&gt;ovf = 1;
		n-&gt;neg = ns;
		d-&gt;neg = ds;
		yyerror(&#34;set ovf in mpdivmodfixfix&#34;);
		return;
	}

	// shift denominator back creating
	// quotient a bit at a time
	// when done the remaining numerator
	// will be the remainder
	for(; i&gt;0; i--) {
		mplsh(q);
		mprsh(d);
		if(mpcmp(d, r) &lt;= 0) {
			mpaddcfix(q, 1);
			mpsubfixfix(r, d);
		}
	}

	n-&gt;neg = ns;
	d-&gt;neg = ds;
	r-&gt;neg = ns;
	q-&gt;neg = ns^ds;
}

void
mpdivfract(Mpint *a, Mpint *b)
{
	Mpint n, d;
	int i, j, neg;
	long *a1, x;

	mpmovefixfix(&amp;n, a);	// numerator
	mpmovefixfix(&amp;d, b);	// denominator
	a1 = &amp;a-&gt;a[Mpprec];	// quotient

	neg = n.neg ^ d.neg;
	n.neg = 0;
	d.neg = 0;

	for(i=0; i&lt;Mpprec; i++) {
		x = 0;
		for(j=0; j&lt;Mpscale; j++) {
			x &lt;&lt;= 1;
			if(mpcmp(&amp;d, &amp;n) &lt;= 0) {
				x |= 1;
				mpsubfixfix(&amp;n, &amp;d);
			}
			mprsh(&amp;d);
		}
		*--a1 = x;
	}
	a-&gt;neg = neg;
}

int
mptestfix(Mpint *a)
{
	Mpint b;
	int r;

	mpmovecfix(&amp;b, 0);
	r = mpcmp(a, &amp;b);
	if(a-&gt;neg) {
		if(r &gt; 0)
			return -1;
		if(r &lt; 0)
			return +1;
	}
	return r;
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
