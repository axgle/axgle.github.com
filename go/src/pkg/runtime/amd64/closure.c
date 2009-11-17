<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/amd64/closure.c</title>

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
  <h1 id="generatedHeader">Text file src/pkg/runtime/amd64/closure.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;runtime.h&#34;

#pragma textflag 7
// func closure(siz int32,
//	fn func(arg0, arg1, arg2 *ptr, callerpc uintptr, xxx) yyy,
//	arg0, arg1, arg2 *ptr) (func(xxx) yyy)
void
runtime·closure(int32 siz, byte *fn, byte *arg0)
{
	byte *p, *q, **ret;
	int32 i, n;
	int64 pcrel;

	if(siz &lt; 0 || siz%8 != 0)
		throw(&#34;bad closure size&#34;);

	ret = (byte**)((byte*)&amp;arg0 + siz);

	if(siz &gt; 100) {
		// TODO(rsc): implement stack growth preamble?
		throw(&#34;closure too big&#34;);
	}

	// compute size of new fn.
	// must match code laid out below.
	n = 7+10+3;	// SUBQ MOVQ MOVQ
	if(siz &lt;= 4*8)
		n += 2*siz/8;	// MOVSQ MOVSQ...
	else
		n += 7+3;	// MOVQ REP MOVSQ
	n += 12;	// CALL worst case; sometimes only 5
	n += 7+1;	// ADDQ RET

	// store args aligned after code, so gc can find them.
	n += siz;
	if(n%8)
		n += 8 - n%8;

	p = mal(n);
	*ret = p;
	q = p + n - siz;

	if(siz &gt; 0) {
		mcpy(q, (byte*)&amp;arg0, siz);

		// SUBQ $siz, SP
		*p++ = 0x48;
		*p++ = 0x81;
		*p++ = 0xec;
		*(uint32*)p = siz;
		p += 4;

		// MOVQ $q, SI
		*p++ = 0x48;
		*p++ = 0xbe;
		*(byte**)p = q;
		p += 8;

		// MOVQ SP, DI
		*p++ = 0x48;
		*p++ = 0x89;
		*p++ = 0xe7;

		if(siz &lt;= 4*8) {
			for(i=0; i&lt;siz; i+=8) {
				// MOVSQ
				*p++ = 0x48;
				*p++ = 0xa5;
			}
		} else {
			// MOVQ $(siz/8), CX  [32-bit immediate siz/8]
			*p++ = 0x48;
			*p++ = 0xc7;
			*p++ = 0xc1;
			*(uint32*)p = siz/8;
			p += 4;

			// REP; MOVSQ
			*p++ = 0xf3;
			*p++ = 0x48;
			*p++ = 0xa5;
		}
	}

	// call fn
	pcrel = fn - (p+5);
	if((int32)pcrel == pcrel) {
		// can use direct call with pc-relative offset
		// CALL fn
		*p++ = 0xe8;
		*(int32*)p = pcrel;
		p += 4;
	} else {
		// MOVQ $fn, CX  [64-bit immediate fn]
		*p++ = 0x48;
		*p++ = 0xb9;
		*(byte**)p = fn;
		p += 8;

		// CALL *CX
		*p++ = 0xff;
		*p++ = 0xd1;
	}

	// ADDQ $siz, SP
	*p++ = 0x48;
	*p++ = 0x81;
	*p++ = 0xc4;
	*(uint32*)p = siz;
	p += 4;

	// RET
	*p++ = 0xc3;

	if(p &gt; q)
		throw(&#34;bad math in sys.closure&#34;);
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
