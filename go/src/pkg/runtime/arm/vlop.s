<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/arm/vlop.s</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/arm/vlop.s</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno&#39;s libkern/vlop-arm.s
// http://code.google.com/p/inferno-os/source/browse/libkern/vlop-arm.s
//
//         Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//         Revisions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com).  All rights reserved.
//         Portions Copyright 2009 The Go Authors. All rights reserved.
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

#define UMULL(Rs,Rm,Rhi,Rlo,S)  WORD	 $((14&lt;&lt;28)|(4&lt;&lt;21)|(S&lt;&lt;20)|(Rhi&lt;&lt;16)|(Rlo&lt;&lt;12)|(Rs&lt;&lt;8)|(9&lt;&lt;4)|Rm)
#define UMLAL(Rs,Rm,Rhi,Rlo,S)  WORD	 $((14&lt;&lt;28)|(5&lt;&lt;21)|(S&lt;&lt;20)|(Rhi&lt;&lt;16)|(Rlo&lt;&lt;12)|(Rs&lt;&lt;8)|(9&lt;&lt;4)|Rm)
#define MUL(Rs,Rm,Rd,S) WORD	 $((14&lt;&lt;28)|(0&lt;&lt;21)|(S&lt;&lt;20)|(Rd&lt;&lt;16)|(Rs&lt;&lt;8)|(9&lt;&lt;4)|Rm)
arg=0

/* replaced use of R10 by R11 because the former can be the data segment base register */

TEXT	_mulv(SB), $0
	MOVW	0(FP), R0
	MOVW	8(FP), R2		/* l0 */
	MOVW	4(FP), R3	  /* h0 */
	MOVW	16(FP), R4	  /* l1 */
	MOVW	12(FP), R5	  /* h1 */
	UMULL(4, 2, 7, 6, 0)
	MUL(3, 4, 8, 0)
	ADD	R8, R7
	MUL(2, 5, 8, 0)
	ADD	R8, R7
	MOVW	R6, 4(R(arg))
	MOVW	R7, 0(R(arg))
	RET


Q	= 0
N	= 1
D	= 2
CC	= 3
TMP	= 11

TEXT	save&lt;&gt;(SB), 7, $0
	MOVW	R(Q), 0(FP)
	MOVW	R(N), 4(FP)
	MOVW	R(D), 8(FP)
	MOVW	R(CC), 12(FP)

	MOVW	R(TMP), R(Q)		/* numerator */
	MOVW	20(FP), R(D)		/* denominator */
	CMP	$0, R(D)
	BNE	s1
	SWI		 0
/*	  MOVW	-1(R(D)), R(TMP)	/* divide by zero fault */
s1:	 RET

TEXT	rest&lt;&gt;(SB), 7, $0
	MOVW	0(FP), R(Q)
	MOVW	4(FP), R(N)
	MOVW	8(FP), R(D)
	MOVW	12(FP), R(CC)
/*
 * return to caller
 * of rest&lt;&gt;
 */
	MOVW	0(R13), R14
	ADD	$20, R13
	B	(R14)

TEXT	div&lt;&gt;(SB), 7, $0
	MOVW	$32, R(CC)
/*
 * skip zeros 8-at-a-time
 */
e1:
	AND.S	$(0xff&lt;&lt;24),R(Q), R(N)
	BNE	e2
	SLL	$8, R(Q)
	SUB.S	$8, R(CC)
	BNE	e1
	RET
e2:
	MOVW	$0, R(N)

loop:
/*
 * shift R(N||Q) left one
 */
	SLL	$1, R(N)
	CMP	$0, R(Q)
	ORR.LT  $1, R(N)
	SLL	$1, R(Q)

/*
 * compare numerator to denominator
 * if less, subtract and set quotent bit
 */
	CMP	R(D), R(N)
	ORR.HS  $1, R(Q)
	SUB.HS  R(D), R(N)
	SUB.S	$1, R(CC)
	BNE	loop
	RET

TEXT	_div(SB), 7, $16
	BL	save&lt;&gt;(SB)
	CMP	$0, R(Q)
	BGE	d1
	RSB	$0, R(Q), R(Q)
	CMP	$0, R(D)
	BGE	d2
	RSB	$0, R(D), R(D)
d0:
	BL	div&lt;&gt;(SB)			/* none/both neg */
	MOVW	R(Q), R(TMP)
	B	out
d1:
	CMP	$0, R(D)
	BGE	d0
	RSB	$0, R(D), R(D)
d2:
	BL	div&lt;&gt;(SB)			/* one neg */
	RSB	$0, R(Q), R(TMP)
	B	out

TEXT	_mod(SB), 7, $16
	BL	save&lt;&gt;(SB)
	CMP	$0, R(D)
	RSB.LT	$0, R(D), R(D)
	CMP	$0, R(Q)
	BGE	m1
	RSB	$0, R(Q), R(Q)
	BL	div&lt;&gt;(SB)			/* neg numerator */
	RSB	$0, R(N), R(TMP)
	B	out
m1:
	BL	div&lt;&gt;(SB)			/* pos numerator */
	MOVW	R(N), R(TMP)
	B	out

TEXT	_divu(SB), 7, $16
	BL	save&lt;&gt;(SB)
	BL	div&lt;&gt;(SB)
	MOVW	R(Q), R(TMP)
	B	out

TEXT	_modu(SB), 7, $16
	BL	save&lt;&gt;(SB)
	BL	div&lt;&gt;(SB)
	MOVW	R(N), R(TMP)
	B	out

out:
	BL	rest&lt;&gt;(SB)
	B	out
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
