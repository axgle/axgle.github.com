<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/syscall/asm_linux_386.s</title>

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
  <h1 id="generatedHeader">Text file src/pkg/syscall/asm_linux_386.s</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//
// System calls for 386, Linux
//

// func Syscall(trap uintptr, a1, a2, a3 uintptr) (r1, r2, err uintptr);
// Trap # in AX, args in BX CX DX SI DI, return in AX

TEXT	syscall·Syscall(SB),7,$0
	CALL	runtime·entersyscall(SB)
	MOVL	4(SP), AX	// syscall entry
	MOVL	8(SP), BX
	MOVL	12(SP), CX
	MOVL	16(SP), DX
	MOVL	$0, SI
	MOVL	$0,  DI
	INT	$0x80
	CMPL	AX, $0xfffff001
	JLS	ok
	MOVL	$-1, 20(SP)	// r1
	MOVL	$0, 24(SP)	// r2
	NEGL	AX
	MOVL	AX, 28(SP)  // errno
	CALL	runtime·exitsyscall(SB)
	RET
ok:
	MOVL	AX, 20(SP)	// r1
	MOVL	DX, 24(SP)	// r2
	MOVL	$0, 28(SP)	// errno
	CALL	runtime·exitsyscall(SB)
	RET

// func Syscall6(trap uintptr, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr);
// Actually Syscall5 but the rest of the code expects it to be named Syscall6.
TEXT	syscall·Syscall6(SB),7,$0
	CALL	runtime·entersyscall(SB)
	MOVL	4(SP), AX	// syscall entry
	MOVL	8(SP), BX
	MOVL	12(SP), CX
	MOVL	16(SP), DX
	MOVL	20(SP), SI
	MOVL	24(SP), DI
	// 28(SP) is ignored
	INT	$0x80
	CMPL	AX, $0xfffff001
	JLS	ok6
	MOVL	$-1, 32(SP)	// r1
	MOVL	$0, 36(SP)	// r2
	NEGL	AX
	MOVL	AX, 40(SP)  // errno
	CALL	runtime·exitsyscall(SB)
	RET
ok6:
	MOVL	AX, 32(SP)	// r1
	MOVL	DX, 36(SP)	// r2
	MOVL	$0, 40(SP)	// errno
	CALL	runtime·exitsyscall(SB)
	RET

// func RawSyscall(trap uintptr, a1, a2, a3 uintptr) (r1, r2, err uintptr);
TEXT syscall·RawSyscall(SB),7,$0
	MOVL	4(SP), AX	// syscall entry
	MOVL	8(SP), BX
	MOVL	12(SP), CX
	MOVL	16(SP), DX
	MOVL	$0, SI
	MOVL	$0,  DI
	INT	$0x80
	CMPL	AX, $0xfffff001
	JLS	ok1
	MOVL	$-1, 20(SP)	// r1
	MOVL	$0, 24(SP)	// r2
	NEGL	AX
	MOVL	AX, 28(SP)  // errno
	CALL	runtime·exitsyscall(SB)
	RET
ok1:
	MOVL	AX, 20(SP)	// r1
	MOVL	DX, 24(SP)	// r2
	MOVL	$0, 28(SP)	// errno
	RET

#define SYS_SOCKETCALL 102	/* from zsysnum_linux_386.go */

// func socketcall(call int, a0, a1, a2, a3, a4, a5 uintptr) (n int, errno int)
// Kernel interface gets call sub-number and pointer to a0.
TEXT syscall·socketcall(SB),7,$0
	CALL	runtime·entersyscall(SB)
	MOVL	$SYS_SOCKETCALL, AX	// syscall entry
	MOVL	4(SP), BX	// socket call number
	LEAL		8(SP), CX	// pointer to call arguments
	MOVL	$0, DX
	MOVL	$0, SI
	MOVL	$0,  DI
	INT	$0x80
	CMPL	AX, $0xfffff001
	JLS	oksock
	MOVL	$-1, 32(SP)	// n
	NEGL	AX
	MOVL	AX, 36(SP)  // errno
	CALL	runtime·exitsyscall(SB)
	RET
oksock:
	MOVL	AX, 32(SP)	// n
	MOVL	$0, 36(SP)	// errno
	CALL	runtime·exitsyscall(SB)
	RET

#define SYS__LLSEEK 140	/* from zsysnum_linux_386.go */
// func Seek(fd int, offset int64, whence int) (newoffset int64, errno int)
// Implemented in assembly to avoid allocation when
// taking the address of the return value newoffset.
// Underlying system call is
//	llseek(int fd, int offhi, int offlo, int64 *result, int whence)
TEXT syscall·Seek(SB),7,$0
	CALL	runtime·entersyscall(SB)
	MOVL	$SYS__LLSEEK, AX	// syscall entry
	MOVL	4(SP), BX	// fd
	MOVL	12(SP), CX	// offset-high
	MOVL	8(SP), DX	// offset-low
	LEAL	20(SP), SI	// result pointer
	MOVL	16(SP),  DI	// whence
	INT	$0x80
	CMPL	AX, $0xfffff001
	JLS	okseek
	MOVL	$-1, 20(SP)	// newoffset low
	MOVL	$-1, 24(SP)	// newoffset high
	NEGL	AX
	MOVL	AX, 28(SP)  // errno
	CALL	runtime·exitsyscall(SB)
	RET
okseek:
	// system call filled in newoffset already
	MOVL	$0, 28(SP)	// errno
	CALL	runtime·exitsyscall(SB)
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
