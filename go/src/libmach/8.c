<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/libmach/8.c</title>

  <link rel="stylesheet" type="text/css" href="../../doc/style.css">
  <script type="text/javascript" src="../../doc/godocs.js"></script>

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
        <a href="../../index.html"><img src="../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../doc/install.html">Install Go</a></li>
    <li><a href="../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../pkg/index.html">Package documentation</a></li>
    <li><a href="../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/libmach/8.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno libmach/8.c
// http://code.google.com/p/inferno-os/source/browse/utils/libmach/8.c
//
// 	Copyright © 1994-1999 Lucent Technologies Inc.
// 	Power PC support Copyright © 1995-2004 C H Forsyth (forsyth@terzarima.net).
// 	Portions Copyright © 1997-1999 Vita Nuova Limited.
// 	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com).
// 	Revisions Copyright © 2000-2004 Lucent Technologies Inc. and others.
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

/*
 * 386 definition
 */
#include &lt;u.h&gt;
#include &lt;bio.h&gt;
#include &lt;ureg_x86.h&gt;
#include &lt;mach.h&gt;

#define	REGOFF(x)	(uintptr)(&amp;((struct Ureg *) 0)-&gt;x)

#define PC		REGOFF(pc)
#define SP		REGOFF(sp)
#define	AX		REGOFF(ax)

#define	REGSIZE		sizeof(struct Ureg)
#define FP_CTL(x)	(REGSIZE+4*(x))
#define FP_REG(x)	(FP_CTL(7)+10*(x))
#define	FPREGSIZE	(7*4+8*10)

Reglist i386reglist[] = {
	{&#34;DI&#34;,		REGOFF(di),	RINT, &#39;X&#39;},
	{&#34;SI&#34;,		REGOFF(si),	RINT, &#39;X&#39;},
	{&#34;BP&#34;,		REGOFF(bp),	RINT, &#39;X&#39;},
	{&#34;BX&#34;,		REGOFF(bx),	RINT, &#39;X&#39;},
	{&#34;DX&#34;,		REGOFF(dx),	RINT, &#39;X&#39;},
	{&#34;CX&#34;,		REGOFF(cx),	RINT, &#39;X&#39;},
	{&#34;AX&#34;,		REGOFF(ax),	RINT, &#39;X&#39;},
	{&#34;GS&#34;,		REGOFF(gs),	RINT, &#39;X&#39;},
	{&#34;FS&#34;,		REGOFF(fs),	RINT, &#39;X&#39;},
	{&#34;ES&#34;,		REGOFF(es),	RINT, &#39;X&#39;},
	{&#34;DS&#34;,		REGOFF(ds),	RINT, &#39;X&#39;},
	{&#34;TRAP&#34;,	REGOFF(trap), 	RINT, &#39;X&#39;},
	{&#34;ECODE&#34;,	REGOFF(ecode),	RINT, &#39;X&#39;},
	{&#34;PC&#34;,		PC,		RINT, &#39;X&#39;},
	{&#34;CS&#34;,		REGOFF(cs),	RINT, &#39;X&#39;},
	{&#34;EFLAGS&#34;,	REGOFF(flags),	RINT, &#39;X&#39;},
	{&#34;SP&#34;,		SP,		RINT, &#39;X&#39;},
	{&#34;SS&#34;,		REGOFF(ss),	RINT, &#39;X&#39;},

	{&#34;E0&#34;,		FP_CTL(0),	RFLT, &#39;X&#39;},
	{&#34;E1&#34;,		FP_CTL(1),	RFLT, &#39;X&#39;},
	{&#34;E2&#34;,		FP_CTL(2),	RFLT, &#39;X&#39;},
	{&#34;E3&#34;,		FP_CTL(3),	RFLT, &#39;X&#39;},
	{&#34;E4&#34;,		FP_CTL(4),	RFLT, &#39;X&#39;},
	{&#34;E5&#34;,		FP_CTL(5),	RFLT, &#39;X&#39;},
	{&#34;E6&#34;,		FP_CTL(6),	RFLT, &#39;X&#39;},
	{&#34;F0&#34;,		FP_REG(0),	RFLT, &#39;3&#39;},
	{&#34;F1&#34;,		FP_REG(1),	RFLT, &#39;3&#39;},
	{&#34;F2&#34;,		FP_REG(2),	RFLT, &#39;3&#39;},
	{&#34;F3&#34;,		FP_REG(3),	RFLT, &#39;3&#39;},
	{&#34;F4&#34;,		FP_REG(4),	RFLT, &#39;3&#39;},
	{&#34;F5&#34;,		FP_REG(5),	RFLT, &#39;3&#39;},
	{&#34;F6&#34;,		FP_REG(6),	RFLT, &#39;3&#39;},
	{&#34;F7&#34;,		FP_REG(7),	RFLT, &#39;3&#39;},
	{  0 }
};

Mach mi386 =
{
	&#34;386&#34;,
	MI386,		/* machine type */
	i386reglist,	/* register list */
	REGSIZE,	/* size of registers in bytes */
	FPREGSIZE,	/* size of fp registers in bytes */
	&#34;PC&#34;,		/* name of PC */
	&#34;SP&#34;,		/* name of SP */
	0,		/* link register */
	&#34;setSB&#34;,	/* static base register name (bogus anyways) */
	0,		/* static base register value */
	0x1000,		/* page size */
	0x80100000ULL,	/* kernel base */
	0xF0000000ULL,	/* kernel text mask */
	0xFFFFFFFFULL,	/* user stack top */
	1,		/* quantization of pc */
	4,		/* szaddr */
	4,		/* szreg */
	4,		/* szfloat */
	8,		/* szdouble */
};
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
