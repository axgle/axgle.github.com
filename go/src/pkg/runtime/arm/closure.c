<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/arm/closure.c</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

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
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

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
  <h1 id="generatedHeader">Text file src/pkg/runtime/arm/closure.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;runtime.h&#34;

/*
	There are two bits of magic:
	- The signature of the compiler generated function uses two stack frames
	as arguments (callerpc separates these frames)
	- size determines how many arguments runtime.closure actually has
	starting at arg0.

	Example closure with 3 captured variables:
	func closure(siz int32,
	fn func(arg0, arg1, arg2 *ptr, callerpc uintptr, xxx) yyy,
		arg0, arg1, arg2 *ptr) (func(xxx) yyy)

	Code generated:
	src R0
	dst R1
	end R3
	tmp R4
	frame = siz+4

//skip loop for 0 size closures
		MOVW.W	R14,-frame(R13)

		MOVW	$vars(PC), R0
		MOVW	$4(SP), R1
		MOVW	$siz(R0), R3
loop:		MOVW.P	4(R0), R4
		MOVW.P	R4, 4(R1)
		CMP		R0, R3
		BNE		loop

		MOVW	8(PC), R0
		BL		(R0)			// 2 words
		MOVW.P	frame(R13),R15
fptr:		WORD	*fn
vars:		WORD	arg0
		WORD	arg1
		WORD	arg2
*/

extern void cacheflush(byte* start, byte* end);

#pragma textflag 7
void
runtime·closure(int32 siz, byte *fn, byte *arg0)
{
	byte *p, *q, **ret;
	uint32 *pc;
	int32 n;

	if(siz &lt; 0 || siz%4 != 0)
		throw(&#34;bad closure size&#34;);

	ret = (byte**)((byte*)&amp;arg0 + siz);

	if(siz &gt; 100) {
		// TODO(kaib): implement stack growth preamble?
		throw(&#34;closure too big&#34;);
	}

	// size of new fn.
	// must match code laid out below.
	if (siz &gt; 0)
		n = 6 * 4 + 7 * 4;
	else
		n = 6 * 4;

	// store args aligned after code, so gc can find them.
	n += siz;

	p = mal(n);
	*ret = p;
	q = p + n - siz;

	pc = (uint32*)p;

	//	MOVW.W	R14,-frame(R13)
	*pc++ = 0xe52de000 | (siz + 4);

	if(siz &gt; 0) {
		mcpy(q, (byte*)&amp;arg0, siz);

		//	MOVW	$vars(PC), R0
		*pc = 0xe28f0000 | (int32)(q - (byte*)pc - 8);
		pc++;

		//	MOVW	$4(SP), R1
		*pc++ = 0xe28d1004;

		//	MOVW	$siz(R0), R3
		*pc++ = 0xe2803000 | siz;

		//	MOVW.P	4(R0), R4
		*pc++ = 0xe4904004;
		//	MOVW.P	R4, 4(R1)
		*pc++ = 0xe4814004;
		//	CMP		R0, R3
		*pc++ = 0xe1530000;
		//	BNE		loop
		*pc++ = 0x1afffffb;
	}

	//	MOVW	fptr(PC), R0
	*pc = 0xe59f0008 | (int32)((q - 4) -(byte*) pc - 8);
	pc++;

	//	BL		(R0)
	*pc++ = 0xe28fe000;
	*pc++ = 0xe280f000;

	//	MOVW.P	frame(R13),R15
	*pc++ = 0xe49df000 | (siz + 4);

	//	WORD	*fn
	*pc++ = (uint32)fn;

	p = (byte*)pc;

	if(p &gt; q)
		throw(&#34;bad math in sys.closure&#34;);

	cacheflush(*ret, q+siz);
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
