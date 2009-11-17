<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/mparith3.c</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/gc/mparith3.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include	&#34;go.h&#34;

/*
 * returns the leading non-zero
 * word of the number
 */
int
sigfig(Mpflt *a)
{
	int i;

	for(i=Mpprec-1; i&gt;=0; i--)
		if(a-&gt;val.a[i] != 0)
			break;
//print(&#34;sigfig %d %d\n&#34;, i-z+1, z);
	return i+1;
}

/*
 * shifts the leading non-zero
 * word of the number to Mpnorm
 */
void
mpnorm(Mpflt *a)
{
	int s;

	s = sigfig(a);
	if(s == 0) {
		// zero
		a-&gt;exp = 0;
		a-&gt;val.neg = 0;
		return;
	}
	s = (Mpnorm-s) * Mpscale;
	mpshiftfix(&amp;a-&gt;val, s);
	a-&gt;exp -= s;
}

/// implements float arihmetic

void
mpaddfltflt(Mpflt *a, Mpflt *b)
{
	int sa, sb, s;
	Mpflt c;

	if(Mpdebug)
		print(&#34;\n%F + %F&#34;, a, b);

	sa = sigfig(a);
	sb = sigfig(b);

	if(sa == 0) {
		if(sb == 0) {
			// zero
			a-&gt;exp = 0;
			a-&gt;val.neg = 0;
			return;
		}
		mpmovefltflt(a, b);
		goto out;
	}
	if(sb == 0)
		goto out;

	s = a-&gt;exp - b-&gt;exp;
	if(s &gt; 0) {
		// a is larger, shift b right
		mpmovefltflt(&amp;c, b);
		mpshiftfix(&amp;c.val, -s);
		mpaddfixfix(&amp;a-&gt;val, &amp;c.val);
		goto out;
	}
	if(s &lt; 0) {
		// b is larger, shift a right
		mpshiftfix(&amp;a-&gt;val, s);
		a-&gt;exp -= s;
		mpaddfixfix(&amp;a-&gt;val, &amp;b-&gt;val);
		goto out;
	}
	mpaddfixfix(&amp;a-&gt;val, &amp;b-&gt;val);

out:
	mpnorm(a);
	if(Mpdebug)
		print(&#34; = %F\n\n&#34;, a);
}

void
mpmulfltflt(Mpflt *a, Mpflt *b)
{
	int sa, sb;

	if(Mpdebug)
		print(&#34;%F\n * %F\n&#34;, a, b);

	sa = sigfig(a);
	sb = sigfig(b);

	if(sa == 0 || sb == 0) {
		// zero
		a-&gt;exp = 0;
		a-&gt;val.neg = 0;
		return;
	}

	mpmulfract(&amp;a-&gt;val, &amp;b-&gt;val);
	a-&gt;exp = (a-&gt;exp + b-&gt;exp) + Mpscale*Mpprec - 1;

	mpnorm(a);
	if(Mpdebug)
		print(&#34; = %F\n\n&#34;, a);
}

void
mpdivfltflt(Mpflt *a, Mpflt *b)
{
	int sa, sb;
	Mpflt c;

	if(Mpdebug)
		print(&#34;%F\n / %F\n&#34;, a, b);

	sa = sigfig(a);
	sb = sigfig(b);

	if(sb == 0) {
		// zero and ovfl
		a-&gt;exp = 0;
		a-&gt;val.neg = 0;
		a-&gt;val.ovf = 1;
		yyerror(&#34;mpdivfltflt divide by zero&#34;);
		return;
	}
	if(sa == 0) {
		// zero
		a-&gt;exp = 0;
		a-&gt;val.neg = 0;
		return;
	}

	// adjust b to top
	mpmovefltflt(&amp;c, b);
	mpshiftfix(&amp;c.val, Mpscale);

	// divide
	mpdivfract(&amp;a-&gt;val, &amp;c.val);
	a-&gt;exp = (a-&gt;exp-c.exp) - Mpscale*(Mpprec-1) + 1;

	mpnorm(a);
	if(Mpdebug)
		print(&#34; = %F\n\n&#34;, a);
}

double
mpgetflt(Mpflt *a)
{
	int s, i;
	uvlong v, vm;
	double f;

	if(a-&gt;val.ovf)
		yyerror(&#34;mpgetflt ovf&#34;);

	s = sigfig(a);
	if(s == 0)
		return 0;

	if(s != Mpnorm) {
		yyerror(&#34;mpgetflt norm&#34;);
		mpnorm(a);
	}

	while((a-&gt;val.a[Mpnorm-1] &amp; Mpsign) == 0) {
		mpshiftfix(&amp;a-&gt;val, 1);
		a-&gt;exp -= 1;
	}

	// the magic numbers (64, 63, 53, 10) are
	// IEEE specific. this should be done machine
	// independently or in the 6g half of the compiler

	// pick up the mantissa in a uvlong
	s = 53;
	v = 0;
	for(i=Mpnorm-1; s&gt;=Mpscale; i--) {
		v = (v&lt;&lt;Mpscale) | a-&gt;val.a[i];
		s -= Mpscale;
	}
	vm = v;
	if(s &gt; 0)
		vm = (vm&lt;&lt;s) | (a-&gt;val.a[i]&gt;&gt;(Mpscale-s));

	// continue with 64 more bits
	s += 64;
	for(; s&gt;=Mpscale; i--) {
		v = (v&lt;&lt;Mpscale) | a-&gt;val.a[i];
		s -= Mpscale;
	}
	if(s &gt; 0)
		v = (v&lt;&lt;s) | (a-&gt;val.a[i]&gt;&gt;(Mpscale-s));

//print(&#34;vm=%.16llux v=%.16llux\n&#34;, vm, v);
	// round toward even
	if(v != (1ULL&lt;&lt;63) || (vm&amp;1ULL) != 0)
		vm += v&gt;&gt;63;

	f = (double)(vm);
	f = ldexp(f, Mpnorm*Mpscale + a-&gt;exp - 53);

	if(a-&gt;val.neg)
		f = -f;
	return f;
}

void
mpmovecflt(Mpflt *a, double c)
{
	int i;
	double f;
	long l;

	if(Mpdebug)
		print(&#34;\nconst %g&#34;, c);
	mpmovecfix(&amp;a-&gt;val, 0);
	a-&gt;exp = 0;
	if(c == 0)
		goto out;
	if(c &lt; 0) {
		a-&gt;val.neg = 1;
		c = -c;
	}

	f = frexp(c, &amp;i);
	a-&gt;exp = i;

	for(i=0; i&lt;10; i++) {
		f = f*Mpbase;
		l = floor(f);
		f = f - l;
		a-&gt;exp -= Mpscale;
		a-&gt;val.a[0] = l;
		if(f == 0)
			break;
		mpshiftfix(&amp;a-&gt;val, Mpscale);
	}

out:
	mpnorm(a);
	if(Mpdebug)
		print(&#34; = %F\n&#34;, a);
}

void
mpnegflt(Mpflt *a)
{
	a-&gt;val.neg ^= 1;
}

int
mptestflt(Mpflt *a)
{
	int s;

	if(Mpdebug)
		print(&#34;\n%F?&#34;, a);
	s = sigfig(a);
	if(s != 0) {
		s = +1;
		if(a-&gt;val.neg)
			s = -1;
	}
	if(Mpdebug)
		print(&#34; = %d\n&#34;, s);
	return s;
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
