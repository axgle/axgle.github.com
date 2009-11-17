<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/big/arith_amd64.s</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/big/arith_amd64.s</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file provides fast assembly versions for the elementary
// arithmetic operations on vectors implemented in arith.go.

TEXT big·useAsm(SB),7,$0
	MOVB $1, 8(SP)  // assembly routines enabled
	RET


// TODO(gri) - experiment with unrolled loops for faster execution

// func addVV_s(z, x, y *Word, n int) (c Word)
TEXT big·addVV_s(SB),7,$0
	MOVQ z+0(FP), R10
	MOVQ x+8(FP), R8
	MOVQ y+16(FP), R9
	MOVL n+24(FP), R11
	MOVQ $0, BX         // i = 0
	MOVQ $0, DX         // c = 0
	JMP E1

L1:	MOVQ (R8)(BX*8), AX
	RCRQ $1, DX
	ADCQ (R9)(BX*8), AX
	RCLQ $1, DX
	MOVQ AX, (R10)(BX*8)
	ADDL $1, BX			// i++

E1:	CMPQ BX, R11		// i &lt; n
	JL L1

	MOVQ DX, c+32(FP)
	RET


// func subVV_s(z, x, y *Word, n int) (c Word)
// (same as addVV_s except for SBBQ instead of ADCQ and label names)
TEXT big·subVV_s(SB),7,$0
	MOVQ z+0(FP), R10
	MOVQ x+8(FP), R8
	MOVQ y+16(FP), R9
	MOVL n+24(FP), R11
	MOVQ $0, BX         // i = 0
	MOVQ $0, DX         // c = 0
	JMP E2

L2:	MOVQ (R8)(BX*8), AX
	RCRQ $1, DX
	SBBQ (R9)(BX*8), AX
	RCLQ $1, DX
	MOVQ AX, (R10)(BX*8)
	ADDL $1, BX         // i++

E2:	CMPQ BX, R11        // i &lt; n
	JL L2

	MOVQ DX, c+32(FP)
	RET


// func addVW_s(z, x *Word, y Word, n int) (c Word)
TEXT big·addVW_s(SB),7,$0
	MOVQ z+0(FP), R10
	MOVQ x+8(FP), R8
	MOVQ y+16(FP), AX   // c = y
	MOVL n+24(FP), R11
	MOVQ $0, BX         // i = 0
	JMP E3

L3:	ADDQ (R8)(BX*8), AX
	MOVQ AX, (R10)(BX*8)
	RCLQ $1, AX
	ANDQ $1, AX
	ADDL $1, BX         // i++

E3:	CMPQ BX, R11        // i &lt; n
	JL L3

	MOVQ AX, c+32(FP)
	RET


// func subVW_s(z, x *Word, y Word, n int) (c Word)
TEXT big·subVW_s(SB),7,$0
	MOVQ z+0(FP), R10
	MOVQ x+8(FP), R8
	MOVQ y+16(FP), AX   // c = y
	MOVL n+24(FP), R11
	MOVQ $0, BX         // i = 0
	JMP E4

L4:	MOVQ (R8)(BX*8), DX	// TODO(gri) is there a reverse SUBQ?
	SUBQ AX, DX
	MOVQ DX, (R10)(BX*8)
	RCLQ $1, AX
	ANDQ $1, AX
	ADDL $1, BX          // i++

E4:	CMPQ BX, R11         // i &lt; n
	JL L4

	MOVQ AX, c+32(FP)
	RET


// func mulAddVWW_s(z, x *Word, y, r Word, n int) (c Word)
TEXT big·mulAddVWW_s(SB),7,$0
	MOVQ z+0(FP), R10
	MOVQ x+8(FP), R8
	MOVQ y+16(FP), R9
	MOVQ r+24(FP), CX   // c = r
	MOVL n+32(FP), R11
	MOVQ $0, BX         // i = 0
	JMP E5

L5:	MOVQ (R8)(BX*8), AX
	MULQ R9
	ADDQ CX, AX
	ADCQ $0, DX
	MOVQ AX, (R10)(BX*8)
	MOVQ DX, CX
	ADDL $1, BX         // i++

E5:	CMPQ BX, R11        // i &lt; n
	JL L5

	MOVQ CX, c+40(FP)
	RET


// func addMulVVW_s(z, x *Word, y Word, n int) (c Word)
TEXT big·addMulVVW_s(SB),7,$0
	MOVQ z+0(FP), R10
	MOVQ x+8(FP), R8
	MOVQ y+16(FP), R9
	MOVL n+24(FP), R11
	MOVQ $0, BX         // i = 0
	MOVQ $0, CX         // c = 0
	JMP E6

L6:	MOVQ (R8)(BX*8), AX
	MULQ R9
	ADDQ (R10)(BX*8), AX
	ADCQ $0, DX
	ADDQ CX, AX
	ADCQ $0, DX
	MOVQ AX, (R10)(BX*8)
	MOVQ DX, CX
	ADDL $1, BX         // i++

E6:	CMPQ BX, R11        // i &lt; n
	JL L6

	MOVQ CX, c+32(FP)
	RET


// divWVW_s(z* Word, xn Word, x *Word, y Word, n int) (r Word)
TEXT big·divWVW_s(SB),7,$0
	MOVQ z+0(FP), R10
	MOVQ xn+8(FP), DX   // r = xn
	MOVQ x+16(FP), R8
	MOVQ y+24(FP), R9
	MOVL n+32(FP), BX   // i = n
	JMP E7

L7:	MOVQ (R8)(BX*8), AX
	DIVQ R9
	MOVQ AX, (R10)(BX*8)

E7:	SUBL $1, BX         // i--
	JGE L7              // i &gt;= 0

	MOVQ DX, r+40(FP)
	RET
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
