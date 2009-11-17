<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/float.c</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/float.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;runtime.h&#34;

static	uint64	uvnan		= 0x7FF0000000000001ULL;
static	uint64	uvinf		= 0x7FF0000000000000ULL;
static	uint64	uvneginf	= 0xFFF0000000000000ULL;

uint32
float32tobits(float32 f)
{
	// The obvious cast-and-pointer code is technically
	// not valid, and gcc miscompiles it.  Use a union instead.
	union {
		float32 f;
		uint32 i;
	} u;
	u.f = f;
	return u.i;
}

uint64
float64tobits(float64 f)
{
	// The obvious cast-and-pointer code is technically
	// not valid, and gcc miscompiles it.  Use a union instead.
	union {
		float64 f;
		uint64 i;
	} u;
	u.f = f;
	return u.i;
}

float64
float64frombits(uint64 i)
{
	// The obvious cast-and-pointer code is technically
	// not valid, and gcc miscompiles it.  Use a union instead.
	union {
		float64 f;
		uint64 i;
	} u;
	u.i = i;
	return u.f;
}

float32
float32frombits(uint32 i)
{
	// The obvious cast-and-pointer code is technically
	// not valid, and gcc miscompiles it.  Use a union instead.
	union {
		float32 f;
		uint32 i;
	} u;
	u.i = i;
	return u.f;
}

bool
isInf(float64 f, int32 sign)
{
	uint64 x;

	x = float64tobits(f);
	if(sign == 0)
		return x == uvinf || x == uvneginf;
	if(sign &gt; 0)
		return x == uvinf;
	return x == uvneginf;
}

float64
NaN(void)
{
	return float64frombits(uvnan);
}

bool
isNaN(float64 f)
{
	uint64 x;

	x = float64tobits(f);
	return ((uint32)(x&gt;&gt;52) &amp; 0x7FF) == 0x7FF &amp;&amp; !isInf(f, 0);
}

float64
Inf(int32 sign)
{
	if(sign &gt;= 0)
		return float64frombits(uvinf);
	else
		return float64frombits(uvneginf);
}

enum
{
	MASK	= 0x7ffL,
	SHIFT	= 64-11-1,
	BIAS	= 1022L,
};

float64
frexp(float64 d, int32 *ep)
{
	uint64 x;

	if(d == 0) {
		*ep = 0;
		return 0;
	}
	x = float64tobits(d);
	*ep = (int32)((x &gt;&gt; SHIFT) &amp; MASK) - BIAS;
	x &amp;= ~((uint64)MASK &lt;&lt; SHIFT);
	x |= (uint64)BIAS &lt;&lt; SHIFT;
	return float64frombits(x);
}

float64
ldexp(float64 d, int32 e)
{
	uint64 x;

	if(d == 0)
		return 0;
	x = float64tobits(d);
	e += (int32)(x &gt;&gt; SHIFT) &amp; MASK;
	if(e &lt;= 0)
		return 0;	/* underflow */
	if(e &gt;= MASK){		/* overflow */
		if(d &lt; 0)
			return Inf(-1);
		return Inf(1);
	}
	x &amp;= ~((uint64)MASK &lt;&lt; SHIFT);
	x |= (uint64)e &lt;&lt; SHIFT;
	return float64frombits(x);
}

float64
modf(float64 d, float64 *ip)
{
	float64 dd;
	uint64 x;
	int32 e;

	if(d &lt; 1) {
		if(d &lt; 0) {
			d = modf(-d, ip);
			*ip = -*ip;
			return -d;
		}
		*ip = 0;
		return d;
	}

	x = float64tobits(d);
	e = (int32)((x &gt;&gt; SHIFT) &amp; MASK) - BIAS;

	/*
	 * Keep the top 11+e bits; clear the rest.
	 */
	if(e &lt;= 64-11)
		x &amp;= ~(((uint64)1 &lt;&lt; (64LL-11LL-e))-1);
	dd = float64frombits(x);
	*ip = dd;
	return d - dd;
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
