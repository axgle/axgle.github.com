<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/mparith1.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/mparith1.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include	&#34;go.h&#34;

/// uses arithmetic

int
mpcmpfixflt(Mpint *a, Mpflt *b)
{
	char buf[500];
	Mpflt c;

	snprint(buf, sizeof(buf), &#34;%B&#34;, a);
	mpatoflt(&amp;c, buf);
	return mpcmpfltflt(&amp;c, b);
}

int
mpcmpfltfix(Mpflt *a, Mpint *b)
{
	char buf[500];
	Mpflt c;

	snprint(buf, sizeof(buf), &#34;%B&#34;, b);
	mpatoflt(&amp;c, buf);
	return mpcmpfltflt(a, &amp;c);
}

int
mpcmpfixfix(Mpint *a, Mpint *b)
{
	Mpint c;

	mpmovefixfix(&amp;c, a);
	mpsubfixfix(&amp;c, b);
	return mptestfix(&amp;c);
}

int
mpcmpfixc(Mpint *b, vlong c)
{
	Mpint a;

	mpmovecfix(&amp;a, c);
	return mpcmpfixfix(&amp;a, b);
}

int
mpcmpfltflt(Mpflt *a, Mpflt *b)
{
	Mpflt c;

	mpmovefltflt(&amp;c, a);
	mpsubfltflt(&amp;c, b);
	return mptestflt(&amp;c);
}

int
mpcmpfltc(Mpflt *b, double c)
{
	Mpflt a;

	mpmovecflt(&amp;a, c);
	return mpcmpfltflt(&amp;a, b);
}

void
mpsubfixfix(Mpint *a, Mpint *b)
{
	mpnegfix(a);
	mpaddfixfix(a, b);
	mpnegfix(a);
}

void
mpsubfltflt(Mpflt *a, Mpflt *b)
{
	mpnegflt(a);
	mpaddfltflt(a, b);
	mpnegflt(a);
}

void
mpaddcfix(Mpint *a, vlong c)
{
	Mpint b;

	mpmovecfix(&amp;b, c);
	mpaddfixfix(a, &amp;b);
}

void
mpaddcflt(Mpflt *a, double c)
{
	Mpflt b;

	mpmovecflt(&amp;b, c);
	mpaddfltflt(a, &amp;b);
}

void
mpmulcfix(Mpint *a, vlong c)
{
	Mpint b;

	mpmovecfix(&amp;b, c);
	mpmulfixfix(a, &amp;b);
}

void
mpmulcflt(Mpflt *a, double c)
{
	Mpflt b;

	mpmovecflt(&amp;b, c);
	mpmulfltflt(a, &amp;b);
}

void
mpdivfixfix(Mpint *a, Mpint *b)
{
	Mpint q, r;

	mpdivmodfixfix(&amp;q, &amp;r, a, b);
	mpmovefixfix(a, &amp;q);
}

void
mpmodfixfix(Mpint *a, Mpint *b)
{
	Mpint q, r;

	mpdivmodfixfix(&amp;q, &amp;r, a, b);
	mpmovefixfix(a, &amp;r);
}

void
mpcomfix(Mpint *a)
{
	Mpint b;

	mpmovecfix(&amp;b, 1);
	mpnegfix(a);
	mpsubfixfix(a, &amp;b);
}

void
mpmovefixflt(Mpflt *a, Mpint *b)
{
	a-&gt;val = *b;
	a-&gt;exp = 0;
	mpnorm(a);
}

// convert (truncate) b to a.
// return -1 (but still convert) if b was non-integer.
int
mpmovefltfix(Mpint *a, Mpflt *b)
{
	Mpflt f;
	*a = b-&gt;val;
	mpshiftfix(a, b-&gt;exp);
	if(b-&gt;exp &lt; 0) {
		f.val = *a;
		f.exp = 0;
		mpnorm(&amp;f);
		if(mpcmpfltflt(b, &amp;f) != 0)
			return -1;
	}
	return 0;
}

void
mpmovefixfix(Mpint *a, Mpint *b)
{
	*a = *b;
}

void
mpmovefltflt(Mpflt *a, Mpflt *b)
{
	*a = *b;
}

static	double	tab[] = { 1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7 };
static void
mppow10flt(Mpflt *a, int p)
{
	if(p &lt; nelem(tab)) {
		mpmovecflt(a, tab[p]);
		return;
	}
	mppow10flt(a, p&gt;&gt;1);
	mpmulfltflt(a, a);
	if(p &amp; 1)
		mpmulcflt(a, 10);
}

//
// floating point input
// required syntax is [+-]d*[.]d*[e[+-]d*]
//
void
mpatoflt(Mpflt *a, char *as)
{
	Mpflt b;
	int dp, c, f, ef, ex, eb, zer;
	char *s;

	s = as;
	dp = 0;		/* digits after decimal point */
	f = 0;		/* sign */
	ex = 0;		/* exponent */
	eb = 0;		/* binary point */
	zer = 1;	/* zero */

	mpmovecflt(a, 0.0);
	for(;;) {
		switch(c = *s++) {
		default:
			goto bad;

		case &#39;-&#39;:
			f = 1;

		case &#39; &#39;:
		case  &#39;\t&#39;:
		case  &#39;+&#39;:
			continue;

		case &#39;.&#39;:
			dp = 1;
			continue;

		case &#39;1&#39;:
		case &#39;2&#39;:
		case &#39;3&#39;:
		case &#39;4&#39;:
		case &#39;5&#39;:
		case &#39;6&#39;:
		case &#39;7&#39;:
		case &#39;8&#39;:
		case &#39;9&#39;:
			zer = 0;

		case &#39;0&#39;:
			mpmulcflt(a, 10);
			mpaddcflt(a, c-&#39;0&#39;);
			if(dp)
				dp++;
			continue;

		case &#39;P&#39;:
		case &#39;p&#39;:
			eb = 1;

		case &#39;E&#39;:
		case &#39;e&#39;:
			ex = 0;
			ef = 0;
			for(;;) {
				c = *s++;
				if(c == &#39;+&#39; || c == &#39; &#39; || c == &#39;\t&#39;)
					continue;
				if(c == &#39;-&#39;) {
					ef = 1;
					continue;
				}
				if(c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;9&#39;) {
					ex = ex*10 + (c-&#39;0&#39;);
					continue;
				}
				break;
			}
			if(ef)
				ex = -ex;

		case 0:
			break;
		}
		break;
	}

	if(eb) {
		if(dp)
			goto bad;
		a-&gt;exp += ex;
		goto out;
	}

	if(dp)
		dp--;
	if(mpcmpfltc(a, 0.0) != 0) {
		if(ex &gt;= dp) {
			mppow10flt(&amp;b, ex-dp);
			mpmulfltflt(a, &amp;b);
		} else {
			mppow10flt(&amp;b, dp-ex);
			mpdivfltflt(a, &amp;b);
		}
	}

out:
	if(f)
		mpnegflt(a);
	return;

bad:
	yyerror(&#34;set ovf in mpatof&#34;);
	mpmovecflt(a, 0.0);
}

//
// fixed point input
// required syntax is [+-][0[x]]d*
//
void
mpatofix(Mpint *a, char *as)
{
	int c, f;
	char *s;

	s = as;
	f = 0;
	mpmovecfix(a, 0);

	c = *s++;
	switch(c) {
	case &#39;-&#39;:
		f = 1;

	case &#39;+&#39;:
		c = *s++;
		if(c != &#39;0&#39;)
			break;

	case &#39;0&#39;:
		goto oct;
	}

	while(c) {
		if(c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;9&#39;) {
			mpmulcfix(a, 10);
			mpaddcfix(a, c-&#39;0&#39;);
			c = *s++;
			continue;
		}
		goto bad;
	}
	goto out;

oct:
	c = *s++;
	if(c == &#39;x&#39; || c == &#39;X&#39;)
		goto hex;
	while(c) {
		if(c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;7&#39;) {
			mpmulcfix(a, 8);
			mpaddcfix(a, c-&#39;0&#39;);
			c = *s++;
			continue;
		}
		goto bad;
	}
	goto out;

hex:
	c = *s++;
	while(c) {
		if(c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;9&#39;) {
			mpmulcfix(a, 16);
			mpaddcfix(a, c-&#39;0&#39;);
			c = *s++;
			continue;
		}
		if(c &gt;= &#39;a&#39; &amp;&amp; c &lt;= &#39;f&#39;) {
			mpmulcfix(a, 16);
			mpaddcfix(a, c+10-&#39;a&#39;);
			c = *s++;
			continue;
		}
		if(c &gt;= &#39;A&#39; &amp;&amp; c &lt;= &#39;F&#39;) {
			mpmulcfix(a, 16);
			mpaddcfix(a, c+10-&#39;A&#39;);
			c = *s++;
			continue;
		}
		goto bad;
	}

out:
	if(f)
		mpnegfix(a);
	return;

bad:
	yyerror(&#34;set ovf in mpatov: %s&#34;, as);
	mpmovecfix(a, 0);
}

int
Bconv(Fmt *fp)
{
	char buf[500], *p;
	Mpint *xval, q, r, ten;
	int f;

	xval = va_arg(fp-&gt;args, Mpint*);
	mpmovefixfix(&amp;q, xval);
	f = 0;
	if(mptestfix(&amp;q) &lt; 0) {
		f = 1;
		mpnegfix(&amp;q);
	}
	mpmovecfix(&amp;ten, 10);

	p = &amp;buf[sizeof(buf)];
	*--p = 0;
	for(;;) {
		mpdivmodfixfix(&amp;q, &amp;r, &amp;q, &amp;ten);
		*--p = mpgetfix(&amp;r) + &#39;0&#39;;
		if(mptestfix(&amp;q) &lt;= 0)
			break;
	}
	if(f)
		*--p = &#39;-&#39;;
	return fmtstrcpy(fp, p);
}

int
Fconv(Fmt *fp)
{
	char buf[500];
	Mpflt *fvp, fv;
	double d;

	fvp = va_arg(fp-&gt;args, Mpflt*);
	if(fp-&gt;flags &amp; FmtSharp) {
		// alternate form - decimal for error messages.
		// for well in range, convert to double and use print&#39;s %g
		if(-900 &lt; fvp-&gt;exp &amp;&amp; fvp-&gt;exp &lt; 900) {
			d = mpgetflt(fvp);
			return fmtprint(fp, &#34;%g&#34;, d);
		}
		// TODO(rsc): for well out of range, print
		// an approximation like 1.234e1000
	}

	if(sigfig(fvp) == 0) {
		snprint(buf, sizeof(buf), &#34;0p+0&#34;);
		goto out;
	}
	fv = *fvp;

	while(fv.val.a[0] == 0) {
		mpshiftfix(&amp;fv.val, -Mpscale);
		fv.exp += Mpscale;
	}
	while((fv.val.a[0]&amp;1) == 0) {
		mpshiftfix(&amp;fv.val, -1);
		fv.exp += 1;
	}

	if(fv.exp &gt;= 0) {
		snprint(buf, sizeof(buf), &#34;%Bp+%d&#34;, &amp;fv.val, fv.exp);
		goto out;
	}
	snprint(buf, sizeof(buf), &#34;%Bp-%d&#34;, &amp;fv.val, -fv.exp);

out:
	return fmtstrcpy(fp, buf);
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
