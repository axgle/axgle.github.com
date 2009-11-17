<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/amd64/asm.s</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/amd64/asm.s</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;amd64/asm.h&#34;

TEXT	_rt0_amd64(SB),7,$-8

	// copy arguments forward on an even stack
	MOVQ	0(SP), AX		// argc
	LEAQ	8(SP), BX		// argv
	SUBQ	$(4*8+7), SP		// 2args 2auto
	ANDQ	$~7, SP
	MOVQ	AX, 16(SP)
	MOVQ	BX, 24(SP)

	// set the per-goroutine and per-mach registers
	LEAQ	m0(SB), m
	LEAQ	g0(SB), g
	MOVQ	g, m_g0(m)		// m has pointer to its g0

	// create istack out of the given (operating system) stack
	LEAQ	(-8192+104)(SP), AX
	MOVQ	AX, g_stackguard(g)
	MOVQ	SP, g_stackbase(g)

	CLD				// convention is D is always left cleared
	CALL	check(SB)

	MOVL	16(SP), AX		// copy argc
	MOVL	AX, 0(SP)
	MOVQ	24(SP), AX		// copy argv
	MOVQ	AX, 8(SP)
	CALL	args(SB)
	CALL	osinit(SB)
	CALL	schedinit(SB)

	// create a new goroutine to start program
	PUSHQ	$mainstart(SB)		// entry
	PUSHQ	$0			// arg size
	CALL	runtime·newproc(SB)
	POPQ	AX
	POPQ	AX

	// start this M
	CALL	mstart(SB)

	CALL	notok(SB)		// never returns
	RET

TEXT mainstart(SB),7,$0
	CALL	main·init(SB)
	CALL	initdone(SB)
	CALL	main·main(SB)
	PUSHQ	$0
	CALL	exit(SB)
	POPQ	AX
	CALL	notok(SB)
	RET

TEXT	breakpoint(SB),7,$0
	BYTE	$0xcc
	RET

/*
 *  go-routine
 */

// uintptr gosave(Gobuf*)
// save state in Gobuf; setjmp
TEXT gosave(SB), 7, $0
	MOVQ	8(SP), AX		// gobuf
	LEAQ	8(SP), BX		// caller&#39;s SP
	MOVQ	BX, gobuf_sp(AX)
	MOVQ	0(SP), BX		// caller&#39;s PC
	MOVQ	BX, gobuf_pc(AX)
	MOVQ	g, gobuf_g(AX)
	MOVL	$0, AX			// return 0
	RET

// void gogo(Gobuf*, uintptr)
// restore state from Gobuf; longjmp
TEXT gogo(SB), 7, $0
	MOVQ	16(SP), AX		// return 2nd arg
	MOVQ	8(SP), BX		// gobuf
	MOVQ	gobuf_g(BX), g
	MOVQ	0(g), CX		// make sure g != nil
	MOVQ	gobuf_sp(BX), SP	// restore SP
	MOVQ	gobuf_pc(BX), BX
	JMP	BX

// void gogocall(Gobuf*, void (*fn)(void))
// restore state from Gobuf but then call fn.
// (call fn, returning to state in Gobuf)
TEXT gogocall(SB), 7, $0
	MOVQ	16(SP), AX		// fn
	MOVQ	8(SP), BX		// gobuf
	MOVQ	gobuf_g(BX), g
	MOVQ	0(g), CX		// make sure g != nil
	MOVQ	gobuf_sp(BX), SP	// restore SP
	MOVQ	gobuf_pc(BX), BX
	PUSHQ	BX
	JMP	AX
	POPQ	BX	// not reached

/*
 * support for morestack
 */

// Called during function prolog when more stack is needed.
TEXT runtime·morestack(SB),7,$0
	// Called from f.
	// Set m-&gt;morebuf to f&#39;s caller.
	MOVQ	8(SP), AX	// f&#39;s caller&#39;s PC
	MOVQ	AX, (m_morebuf+gobuf_pc)(m)
	LEAQ	16(SP), AX	// f&#39;s caller&#39;s SP
	MOVQ	AX, (m_morebuf+gobuf_sp)(m)
	MOVQ	AX, (m_morefp)(m)
	MOVQ	g, (m_morebuf+gobuf_g)(m)

	// Set m-&gt;morepc to f&#39;s PC.
	MOVQ	0(SP), AX
	MOVQ	AX, m_morepc(m)

	// Call newstack on m&#39;s scheduling stack.
	MOVQ	m_g0(m), g
	MOVQ	(m_sched+gobuf_sp)(m), SP
	CALL	newstack(SB)
	MOVQ	$0, 0x1003	// crash if newstack returns
	RET

// Called from reflection library.  Mimics morestack,
// reuses stack growth code to create a frame
// with the desired args running the desired function.
//
// func call(fn *byte, arg *byte, argsize uint32).
TEXT reflect·call(SB), 7, $0
	// Save our caller&#39;s state as the PC and SP to
	// restore when returning from f.
	MOVQ	0(SP), AX	// our caller&#39;s PC
	MOVQ	AX, (m_morebuf+gobuf_pc)(m)
	LEAQ	8(SP), AX	// our caller&#39;s SP
	MOVQ	AX, (m_morebuf+gobuf_sp)(m)
	MOVQ	g, (m_morebuf+gobuf_g)(m)

	// Set up morestack arguments to call f on a new stack.
	// We set f&#39;s frame size to zero, meaning
	// allocate a standard sized stack segment.
	// If it turns out that f needs a larger frame than this,
	// f&#39;s usual stack growth prolog will allocate
	// a new segment (and recopy the arguments).
	MOVQ	8(SP), AX	// fn
	MOVQ	16(SP), BX	// arg frame
	MOVL	24(SP), CX	// arg size

	MOVQ	AX, m_morepc(m)	// f&#39;s PC
	MOVQ	BX, m_morefp(m)	// argument frame pointer
	MOVL	CX, m_moreargs(m)	// f&#39;s argument size
	MOVL	$0, m_moreframe(m)	// f&#39;s frame size

	// Call newstack on m&#39;s scheduling stack.
	MOVQ	m_g0(m), g
	MOVQ	(m_sched+gobuf_sp)(m), SP
	CALL	newstack(SB)
	MOVQ	$0, 0x1103	// crash if newstack returns
	RET

// Return point when leaving stack.
TEXT runtime·lessstack(SB), 7, $0
	// Save return value in m-&gt;cret
	MOVQ	AX, m_cret(m)

	// Call oldstack on m&#39;s scheduling stack.
	MOVQ	m_g0(m), g
	MOVQ	(m_sched+gobuf_sp)(m), SP
	CALL	oldstack(SB)
	MOVQ	$0, 0x1004	// crash if oldstack returns
	RET

// morestack trampolines
TEXT	runtime·morestack00+0(SB),7,$0
	MOVQ	$0, AX
	MOVQ	AX, m_moreframe(m)
	MOVQ	$runtime·morestack+0(SB), AX
	JMP	AX

TEXT	runtime·morestack01+0(SB),7,$0
	SHLQ	$32, AX
	MOVQ	AX, m_moreframe(m)
	MOVQ	$runtime·morestack+0(SB), AX
	JMP	AX

TEXT	runtime·morestack10+0(SB),7,$0
	MOVLQZX	AX, AX
	MOVQ	AX, m_moreframe(m)
	MOVQ	$runtime·morestack+0(SB), AX
	JMP	AX

TEXT	runtime·morestack11+0(SB),7,$0
	MOVQ	AX, m_moreframe(m)
	MOVQ	$runtime·morestack+0(SB), AX
	JMP	AX

// subcases of morestack01
// with const of 8,16,...48
TEXT	runtime·morestack8(SB),7,$0
	PUSHQ	$1
	MOVQ	$runtime·morestackx(SB), AX
	JMP	AX

TEXT	runtime·morestack16(SB),7,$0
	PUSHQ	$2
	MOVQ	$runtime·morestackx(SB), AX
	JMP	AX

TEXT	runtime·morestack24(SB),7,$0
	PUSHQ	$3
	MOVQ	$runtime·morestackx(SB), AX
	JMP	AX

TEXT	runtime·morestack32(SB),7,$0
	PUSHQ	$4
	MOVQ	$runtime·morestackx(SB), AX
	JMP	AX

TEXT	runtime·morestack40(SB),7,$0
	PUSHQ	$5
	MOVQ	$runtime·morestackx(SB), AX
	JMP	AX

TEXT	runtime·morestack48(SB),7,$0
	PUSHQ	$6
	MOVQ	$runtime·morestackx(SB), AX
	JMP	AX

TEXT	runtime·morestackx(SB),7,$0
	POPQ	AX
	SHLQ	$35, AX
	MOVQ	AX, m_moreframe(m)
	MOVQ	$runtime·morestack(SB), AX
	JMP	AX

// bool cas(int32 *val, int32 old, int32 new)
// Atomically:
//	if(*val == old){
//		*val = new;
//		return 1;
//	} else
//		return 0;
TEXT cas(SB), 7, $0
	MOVQ	8(SP), BX
	MOVL	16(SP), AX
	MOVL	20(SP), CX
	LOCK
	CMPXCHGL	CX, 0(BX)
	JZ 3(PC)
	MOVL	$0, AX
	RET
	MOVL	$1, AX
	RET

// void jmpdefer(fn, sp);
// called from deferreturn.
// 1. pop the caller
// 2. sub 5 bytes from the callers return
// 3. jmp to the argument
TEXT jmpdefer(SB), 7, $0
	MOVQ	8(SP), AX	// fn
	MOVQ	16(SP), BX	// caller sp
	LEAQ	-8(BX), SP	// caller sp after CALL
	SUBQ	$5, (SP)	// return to CALL again
	JMP	AX	// but first run the deferred function

// runcgo(void(*fn)(void*), void *arg)
// Call fn(arg) on the scheduler stack,
// aligned appropriately for the gcc ABI.
// Save g and m across the call,
// since the foreign code might reuse them.
TEXT runcgo(SB),7,$32
	// Save old registers.
	MOVQ	fn+0(FP),AX
	MOVQ	arg+8(FP),DI	// DI = first argument in AMD64 ABI
	MOVQ	SP, CX

	// Figure out if we need to switch to m-&gt;g0 stack.
	MOVQ	m_g0(m), R8
	CMPQ	R8, g
	JEQ	2(PC)
	MOVQ	(m_sched+gobuf_sp)(m), SP

	// Now on a scheduling stack (a pthread-created stack).
	SUBQ	$32, SP
	ANDQ	$~15, SP	// alignment for gcc ABI
	MOVQ	g, 24(SP)	// save old g, m, SP
	MOVQ	m, 16(SP)
	MOVQ	CX, 8(SP)
	CALL	AX

	// Restore registers, stack pointer.
	MOVQ	16(SP), m
	MOVQ	24(SP), g
	MOVQ	8(SP), SP
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
