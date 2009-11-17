<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/386/asm.s</title>

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
  <h1 id="generatedHeader">Text file src/pkg/runtime/386/asm.s</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;386/asm.h&#34;

TEXT _rt0_386(SB),7,$0
	// copy arguments forward on an even stack
	MOVL	0(SP), AX		// argc
	LEAL	4(SP), BX		// argv
	SUBL	$128, SP		// plenty of scratch
	ANDL	$~15, SP
	MOVL	AX, 120(SP)		// save argc, argv away
	MOVL	BX, 124(SP)

	// if there is an initcgo, call it to let it
	// initialize and to set up GS.  if not,
	// we set up GS ourselves.
	MOVL	initcgo(SB), AX
	TESTL	AX, AX
	JZ	3(PC)
	CALL	AX
	JMP	ok

	// set up %gs
	CALL	ldt0setup(SB)

	// store through it, to make sure it works
	MOVL	$0x123, 0(GS)
	MOVL	tls0(SB), AX
	CMPL	AX, $0x123
	JEQ	ok
	MOVL	AX, 0	// abort
ok:
	// set up m and g &#34;registers&#34;
	LEAL	g0(SB), CX
	MOVL	CX, g
	LEAL	m0(SB), AX
	MOVL	AX, m

	// save m-&gt;g0 = g0
	MOVL	CX, m_g0(AX)

	// create istack out of the OS stack
	LEAL	(-8192+104)(SP), AX	// TODO: 104?
	MOVL	AX, g_stackguard(CX)
	MOVL	SP, g_stackbase(CX)
	CALL	emptyfunc(SB)	// fault if stack check is wrong

	// convention is D is always cleared
	CLD

	CALL	check(SB)

	// saved argc, argv
	MOVL	120(SP), AX
	MOVL	AX, 0(SP)
	MOVL	124(SP), AX
	MOVL	AX, 4(SP)
	CALL	args(SB)
	CALL	osinit(SB)
	CALL	schedinit(SB)

	// create a new goroutine to start program
	PUSHL	$mainstart(SB)	// entry
	PUSHL	$0	// arg size
	CALL	runtime·newproc(SB)
	POPL	AX
	POPL	AX

	// start this M
	CALL	mstart(SB)

	INT $3
	RET

TEXT mainstart(SB),7,$0
	CALL	main·init(SB)
	CALL	initdone(SB)
	CALL	main·main(SB)
	PUSHL	$0
	CALL	exit(SB)
	POPL	AX
	INT $3
	RET

TEXT	breakpoint(SB),7,$0
	INT $3
	RET

/*
 *  go-routine
 */

// uintptr gosave(Gobuf*)
// save state in Gobuf; setjmp
TEXT gosave(SB), 7, $0
	MOVL	4(SP), AX		// gobuf
	LEAL	4(SP), BX		// caller&#39;s SP
	MOVL	BX, gobuf_sp(AX)
	MOVL	0(SP), BX		// caller&#39;s PC
	MOVL	BX, gobuf_pc(AX)
	MOVL	g, BX
	MOVL	BX, gobuf_g(AX)
	MOVL	$0, AX			// return 0
	RET

// void gogo(Gobuf*, uintptr)
// restore state from Gobuf; longjmp
TEXT gogo(SB), 7, $0
	MOVL	8(SP), AX		// return 2nd arg
	MOVL	4(SP), BX		// gobuf
	MOVL	gobuf_g(BX), DX
	MOVL	0(DX), CX		// make sure g != nil
	MOVL	DX, g
	MOVL	gobuf_sp(BX), SP	// restore SP
	MOVL	gobuf_pc(BX), BX
	JMP	BX

// void gogocall(Gobuf*, void (*fn)(void))
// restore state from Gobuf but then call fn.
// (call fn, returning to state in Gobuf)
TEXT gogocall(SB), 7, $0
	MOVL	8(SP), AX		// fn
	MOVL	4(SP), BX		// gobuf
	MOVL	gobuf_g(BX), DX
	MOVL	DX, g
	MOVL	0(DX), CX		// make sure g != nil
	MOVL	gobuf_sp(BX), SP	// restore SP
	MOVL	gobuf_pc(BX), BX
	PUSHL	BX
	JMP	AX
	POPL	BX	// not reached

/*
 * support for morestack
 */

// Called during function prolog when more stack is needed.
TEXT runtime·morestack(SB),7,$0
	// Cannot grow scheduler stack (m-&gt;g0).
	MOVL	m, BX
	MOVL	m_g0(BX), SI
	CMPL	g, SI
	JNE	2(PC)
	INT	$3

	// frame size in DX
	// arg size in AX
	// Save in m.
	MOVL	DX, m_moreframe(BX)
	MOVL	AX, m_moreargs(BX)

	// Called from f.
	// Set m-&gt;morebuf to f&#39;s caller.
	MOVL	4(SP), DI	// f&#39;s caller&#39;s PC
	MOVL	DI, (m_morebuf+gobuf_pc)(BX)
	LEAL	8(SP), CX	// f&#39;s caller&#39;s SP
	MOVL	CX, (m_morebuf+gobuf_sp)(BX)
	MOVL	CX, (m_morefp)(BX)
	MOVL	g, SI
	MOVL	SI, (m_morebuf+gobuf_g)(BX)

	// Set m-&gt;morepc to f&#39;s PC.
	MOVL	0(SP), AX
	MOVL	AX, m_morepc(BX)

	// Call newstack on m&#39;s scheduling stack.
	MOVL	m_g0(BX), BP
	MOVL	BP, g
	MOVL	(m_sched+gobuf_sp)(BX), SP
	CALL	newstack(SB)
	MOVL	$0, 0x1003	// crash if newstack returns
	RET

// Called from reflection library.  Mimics morestack,
// reuses stack growth code to create a frame
// with the desired args running the desired function.
//
// func call(fn *byte, arg *byte, argsize uint32).
TEXT reflect·call(SB), 7, $0
	MOVL	m, BX

	// Save our caller&#39;s state as the PC and SP to
	// restore when returning from f.
	MOVL	0(SP), AX	// our caller&#39;s PC
	MOVL	AX, (m_morebuf+gobuf_pc)(BX)
	LEAL	4(SP), AX	// our caller&#39;s SP
	MOVL	AX, (m_morebuf+gobuf_sp)(BX)
	MOVL	g, AX
	MOVL	AX, (m_morebuf+gobuf_g)(BX)

	// Set up morestack arguments to call f on a new stack.
	// We set f&#39;s frame size to zero, meaning
	// allocate a standard sized stack segment.
	// If it turns out that f needs a larger frame than this,
	// f&#39;s usual stack growth prolog will allocate
	// a new segment (and recopy the arguments).
	MOVL	4(SP), AX	// fn
	MOVL	8(SP), DX	// arg frame
	MOVL	12(SP), CX	// arg size

	MOVL	AX, m_morepc(BX)	// f&#39;s PC
	MOVL	DX, m_morefp(BX)	// argument frame pointer
	MOVL	CX, m_moreargs(BX)	// f&#39;s argument size
	MOVL	$0, m_moreframe(BX)	// f&#39;s frame size

	// Call newstack on m&#39;s scheduling stack.
	MOVL	m_g0(BX), BP
	MOVL	BP, g
	MOVL	(m_sched+gobuf_sp)(BX), SP
	CALL	newstack(SB)
	MOVL	$0, 0x1103	// crash if newstack returns
	RET


// Return point when leaving stack.
TEXT runtime·lessstack(SB), 7, $0
	// Save return value in m-&gt;cret
	MOVL	m, BX
	MOVL	AX, m_cret(BX)

	// Call oldstack on m&#39;s scheduling stack.
	MOVL	m_g0(BX), DX
	MOVL	DX, g
	MOVL	(m_sched+gobuf_sp)(BX), SP
	CALL	oldstack(SB)
	MOVL	$0, 0x1004	// crash if oldstack returns
	RET


// bool cas(int32 *val, int32 old, int32 new)
// Atomically:
//	if(*val == old){
//		*val = new;
//		return 1;
//	}else
//		return 0;
TEXT cas(SB), 7, $0
	MOVL	4(SP), BX
	MOVL	8(SP), AX
	MOVL	12(SP), CX
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
	MOVL	4(SP), AX	// fn
	MOVL	8(SP), BX	// caller sp
	LEAL	-4(BX), SP	// caller sp after CALL
	SUBL	$5, (SP)	// return to CALL again
	JMP	AX	// but first run the deferred function

TEXT	runtime·memclr(SB),7,$0
	MOVL	4(SP), DI		// arg 1 addr
	MOVL	8(SP), CX		// arg 2 count
	ADDL	$3, CX
	SHRL	$2, CX
	MOVL	$0, AX
	CLD
	REP
	STOSL
	RET

TEXT	runtime·getcallerpc+0(SB),7,$0
	MOVL	x+0(FP),AX		// addr of first arg
	MOVL	-4(AX),AX		// get calling pc
	RET

TEXT	runtime·setcallerpc+0(SB),7,$0
	MOVL	x+0(FP),AX		// addr of first arg
	MOVL	x+4(FP), BX
	MOVL	BX, -4(AX)		// set calling pc
	RET

TEXT ldt0setup(SB),7,$16
	// set up ldt 7 to point at tls0
	// ldt 1 would be fine on Linux, but on OS X, 7 is as low as we can go.
	// the entry number is just a hint.  setldt will set up GS with what it used.
	MOVL	$7, 0(SP)
	LEAL	tls0(SB), AX
	MOVL	AX, 4(SP)
	MOVL	$32, 8(SP)	// sizeof(tls array)
	CALL	setldt(SB)
	RET

TEXT emptyfunc(SB),0,$0
	RET

TEXT	abort(SB),7,$0
	INT $0x3

// runcgo(void(*fn)(void*), void *arg)
// Call fn(arg) on the scheduler stack,
// aligned appropriately for the gcc ABI.
TEXT	runcgo(SB),7,$16
	MOVL	fn+0(FP), AX
	MOVL	arg+4(FP), BX
	MOVL	SP, CX

	// Figure out if we need to switch to m-&gt;g0 stack.
	MOVL	m, DX
	MOVL	m_g0(DX), SI
	CMPL	g, SI
	JEQ	2(PC)
	MOVL	(m_sched+gobuf_sp)(DX), SP

	// Now on a scheduling stack (a pthread-created stack).
	SUBL	$16, SP
	ANDL	$~15, SP	// alignment for gcc ABI
	MOVL	CX, 4(SP)
	MOVL	BX, 0(SP)
	CALL	AX
	MOVL	4(SP), SP
	RET


GLOBL m0(SB), $1024
GLOBL g0(SB), $1024
GLOBL tls0(SB), $32
GLOBL initcgo(SB), $4

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
