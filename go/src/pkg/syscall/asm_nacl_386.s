<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/syscall/asm_nacl_386.s</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/syscall/asm_nacl_386.s</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//
// System calls for 386, Native Client
//

#define SYSCALL(x)	$(0x10000+x * 32)

// func Syscall(trap uintptr, a1, a2, a3 uintptr) (r1, r2, err uintptr);
// Trap # in AX, args in BX CX DX SI DI, return in AX

TEXT	syscall·Syscall(SB),7,$20
	CALL	runtime·entersyscall(SB)
	MOVL	trap+0(FP), AX	// syscall entry
	MOVL	a1+4(FP), BX
	MOVL	a2+8(FP), CX
	MOVL	a3+12(FP), DX
	MOVL	$0, SI
	MOVL	$0,  DI

	MOVL	BX, 0(SP)
	MOVL	CX, 4(SP)
	MOVL	DX, 8(SP)
	MOVL	SI, 12(SP)
	MOVL	DI, 16(SP)

	// Call $(0x10000+32*AX)
	SHLL	$5, AX
	ADDL	$0x10000, AX
	CALL	AX

	CMPL	AX, $0xfffff001
	JLS	ok
	MOVL	$-1, r1+16(FP)
	MOVL	$0, r2+20(FP)
	NEGL	AX
	MOVL	AX, errno+24(FP)
	CALL	runtime·exitsyscall(SB)
	RET
ok:
	MOVL	AX, r1+16(FP)
	MOVL	DX, r2+20(FP)
	MOVL	$0, errno+24(FP)
	CALL	runtime·exitsyscall(SB)
	RET

// func Syscall6(trap uintptr, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr);
TEXT	syscall·Syscall6(SB),7,$24
	CALL	runtime·entersyscall(SB)
	MOVL	a1+4(FP), BX
	MOVL	a2+8(FP), CX
	MOVL	a3+12(FP), DX
	MOVL	a4+16(FP), SI
	MOVL	a5+20(FP), DI
	MOVL	a6+24(FP), AX

	MOVL	BX, 0(SP)
	MOVL	CX, 4(SP)
	MOVL	DX, 8(SP)
	MOVL	SI, 12(SP)
	MOVL	DI, 16(SP)
	MOVL	AX, 20(SP)

	// Call $(0x10000+32*trap)
	MOVL	trap+0(FP), AX	// syscall entry
	SHLL	$5, AX
	ADDL	$0x10000, AX
	CALL	AX

	CMPL	AX, $0xfffff001
	JLS	ok6
	MOVL	$-1, r1+28(FP)
	MOVL	$0, r2+32(FP)
	NEGL	AX
	MOVL	AX, errno+36(FP)
	CALL	runtime·exitsyscall(SB)
	RET
ok6:
	MOVL	AX, r1+28(FP)
	MOVL	DX, r2+32(FP)
	MOVL	$0, errno+36(FP)
	CALL	runtime·exitsyscall(SB)
	RET

// func RawSyscall(trap uintptr, a1, a2, a3 uintptr) (r1, r2, err uintptr);
TEXT syscall·RawSyscall(SB),7,$20
	MOVL	trap+0(FP), AX	// syscall entry
	MOVL	a1+4(FP), BX
	MOVL	a2+8(FP), CX
	MOVL	a3+12(FP), DX
	MOVL	$0, SI
	MOVL	$0,  DI

	MOVL	BX, 0(SP)
	MOVL	CX, 4(SP)
	MOVL	DX, 8(SP)
	MOVL	SI, 12(SP)
	MOVL	DI, 16(SP)

	// Call $(0x10000+32*AX)
	SHLL	$5, AX
	ADDL	$0x10000, AX
	CALL	AX

	CMPL	AX, $0xfffff001
	JLS	ok1
	MOVL	$-1, r1+16(FP)
	MOVL	$0, r2+20(FP)
	NEGL	AX
	MOVL	AX, errno+24(FP)
	RET
ok1:
	MOVL	AX, r1+16(FP)
	MOVL	DX, r2+20(FP)
	MOVL	$0, errno+24(FP)
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
