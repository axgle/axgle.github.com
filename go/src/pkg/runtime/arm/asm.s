<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/arm/asm.s</title>

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
  <h1 id="generatedHeader">Text file src/pkg/runtime/arm/asm.s</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;arm/asm.h&#34;

// using frame size $-4 means do not save LR on stack.
TEXT _rt0_arm(SB),7,$-4
	MOVW $setR12(SB), R12

	// copy arguments forward on an even stack
	// use R13 instead of SP to avoid linker rewriting the offsets
	MOVW	0(R13), R0		// argc
	MOVW	$4(R13), R1		// argv
	SUB	$128, R13		// plenty of scratch
	AND	$~7, R13
	MOVW	R0, 120(R13)		// save argc, argv away
	MOVW	R1, 124(R13)

	// set up m and g registers
	// g is R10, m is R9
	MOVW	$g0(SB), g
	MOVW	$m0(SB), m

	// save m-&gt;g0 = g0
	MOVW	g, m_g0(m)

	// create istack out of the OS stack
	MOVW	$(-8192+104)(R13), R0
	MOVW	R0, g_stackguard(g)	// (w 104b guard)
	MOVW	R13, g_stackbase(g)
	BL	emptyfunc(SB)	// fault if stack check is wrong

	BL	check(SB)

	// saved argc, argv
	MOVW	120(R13), R0
	MOVW	R0, 4(R13)
	MOVW	124(R13), R1
	MOVW	R1, 8(R13)
	BL	args(SB)
	BL	osinit(SB)
	BL	schedinit(SB)

	// create a new goroutine to start program
	MOVW	$mainstart(SB), R0
	MOVW.W	R0, -4(R13)
	MOVW	$8, R0
	MOVW.W	R0, -4(R13)
	MOVW	$0, R0
	MOVW.W	R0, -4(R13)	// push $0 as guard
	BL	runtime·newproc(SB)
	MOVW	$12(R13), R13	// pop args and LR

	// start this M
	BL	mstart(SB)

	MOVW	$1234, R0
	MOVW	$1000, R1
	MOVW	R0, (R1)	// fail hard
	B	_dep_dummy(SB)	// Never reached


TEXT mainstart(SB),7,$4
	BL	main·init(SB)
	BL	initdone(SB)
	BL	main·main(SB)
	MOVW	$0, R0
	MOVW	R0, 4(SP)
	BL	exit(SB)
	MOVW	$1234, R0
	MOVW	$1001, R1
	MOVW	R0, (R1)	// fail hard
	RET

// TODO(kaib): remove these once linker works properly
// pull in dummy dependencies
TEXT _dep_dummy(SB),7,$0
	BL	_div(SB)
	BL	_divu(SB)
	BL	_mod(SB)
	BL	_modu(SB)
	BL	_modu(SB)

TEXT	breakpoint(SB),7,$0
	BL	abort(SB)
//	BYTE $0xcc
//	RET

/*
 *  go-routine
 */

// uintptr gosave(Gobuf*)
// save state in Gobuf; setjmp
TEXT gosave(SB), 7, $-4
	MOVW	0(FP), R0
	MOVW	SP, gobuf_sp(R0)
	MOVW	LR, gobuf_pc(R0)
	MOVW	g, gobuf_g(R0)
	MOVW	$0, R0			// return 0
	RET

// void gogo(Gobuf*, uintptr)
// restore state from Gobuf; longjmp
TEXT	gogo(SB), 7, $-4
	MOVW	0(FP), R1			// gobuf
	MOVW	4(FP), R0		// return 2nd arg
	MOVW	gobuf_g(R1), g
	MOVW	0(g), R2		// make sure g != nil
	MOVW	gobuf_sp(R1), SP	// restore SP
	MOVW	gobuf_pc(R1), PC

// void gogocall(Gobuf*, void (*fn)(void))
// restore state from Gobuf but then call fn.
// (call fn, returning to state in Gobuf)
// using frame size $-4 means do not save LR on stack.
TEXT gogocall(SB), 7, $-4
	MOVW	0(FP), R0
	MOVW	4(FP), R1		// fn
	MOVW	gobuf_g(R0), g
	MOVW	0(g), R2		// make sure g != nil
	MOVW	gobuf_sp(R0), SP	// restore SP
	MOVW	gobuf_pc(R0), LR
	MOVW	R1, PC

/*
 * support for morestack
 */

// Called during function prolog when more stack is needed.
// R1 frame size
// R2 arg size
// R3 prolog&#39;s LR
// NB. we do not save R0 because the we&#39;ve forced 5c to pass all arguments
// on the stack.
// using frame size $-4 means do not save LR on stack.
TEXT runtime·morestack(SB),7,$-4
	// Cannot grow scheduler stack (m-&gt;g0).
	MOVW	m_g0(m), R4
	CMP	g, R4
	BNE	2(PC)
	BL	abort(SB)

	// Save in m.
	MOVW	R1, m_moreframe(m)
	MOVW	R2, m_moreargs(m)

	// Called from f.
	// Set m-&gt;morebuf to f&#39;s caller.
	MOVW	R3, (m_morebuf+gobuf_pc)(m) // f&#39;s caller&#39;s PC
	MOVW	SP, (m_morebuf+gobuf_sp)(m) // f&#39;s caller&#39;s SP
	MOVW	SP, m_morefp(m) // f&#39;s caller&#39;s SP
	MOVW	g, (m_morebuf+gobuf_g)(m)

	// Set m-&gt;morepc to f&#39;s PC.
	MOVW	LR, m_morepc(m)

	// Call newstack on m&#39;s scheduling stack.
	MOVW	m_g0(m), g
	MOVW	(m_sched+gobuf_sp)(m), SP
	B	newstack(SB)

// Called from reflection library.  Mimics morestack,
// reuses stack growth code to create a frame
// with the desired args running the desired function.
//
// func call(fn *byte, arg *byte, argsize uint32).
TEXT reflect·call(SB), 7, $-4
	// Save our caller&#39;s state as the PC and SP to
	// restore when returning from f.
	MOVW	LR, (m_morebuf+gobuf_pc)(m)	// our caller&#39;s PC
	MOVW	SP, (m_morebuf+gobuf_sp)(m)	// our caller&#39;s SP
	MOVW	g,  (m_morebuf+gobuf_g)(m)

	// Set up morestack arguments to call f on a new stack.
	// We set f&#39;s frame size to zero, meaning
	// allocate a standard sized stack segment.
	// If it turns out that f needs a larger frame than this,
	// f&#39;s usual stack growth prolog will allocate
	// a new segment (and recopy the arguments).
	MOVW	4(SP), R0	// fn
	MOVW	8(SP), R1	// arg frame
	MOVW	12(SP), R2	// arg size

	MOVW	R0, m_morepc(m)	// f&#39;s PC
	MOVW	R1, m_morefp(m)	// argument frame pointer
	MOVW	R2, m_moreargs(m)	// f&#39;s argument size
	MOVW	$0, R3
	MOVW	R3, m_moreframe(m)	// f&#39;s frame size

	// Call newstack on m&#39;s scheduling stack.
	MOVW	m_g0(m), g
	MOVW	(m_sched+gobuf_sp)(m), SP
	B	newstack(SB)

// Return point when leaving stack.
// using frame size $-4 means do not save LR on stack.
TEXT runtime·lessstack(SB), 7, $-4
	// Save return value in m-&gt;cret
	MOVW	R0, m_cret(m)

	// Call oldstack on m&#39;s scheduling stack.
	MOVW	m_g0(m), g
	MOVW	(m_sched+gobuf_sp)(m), SP
	B	oldstack(SB)

// void jmpdefer(fn, sp);
// called from deferreturn.
// 1. grab stored LR for caller
// 2. sub 4 bytes to get back to BL deferreturn
// 3. B to fn
TEXT jmpdefer(SB), 7, $0
	MOVW	0(SP), LR
	MOVW	$-4(LR), LR	// BL deferreturn
	MOVW	4(SP), R0		// fn
	MOVW	8(SP), R1
	MOVW	$-4(R1), SP	// correct for sp pointing to arg0, past stored lr
	B		(R0)

TEXT	runtime·memclr(SB),7,$20
	MOVW	0(FP), R0
	MOVW	$0, R1		// c = 0
	MOVW	R1, -16(SP)
	MOVW	4(FP), R1	// n
	MOVW	R1, -12(SP)
	MOVW	m, -8(SP)	// Save m and g
	MOVW	g, -4(SP)
	BL	memset(SB)
	MOVW	-8(SP), m	// Restore m and g, memset clobbers them
	MOVW	-4(SP), g
	RET

TEXT	runtime·getcallerpc+0(SB),7,$-4
	MOVW	0(SP), R0
	RET

TEXT	runtime·setcallerpc+0(SB),7,$-4
	MOVW	x+4(FP), R0
	MOVW	R0, 0(SP)
	RET

// runcgo(void(*fn)(void*), void *arg)
// Just call fn(arg), but first align the stack
// appropriately for the gcc ABI.
// TODO(kaib): figure out the arm-gcc ABI
TEXT	runcgo(SB),7,$16
	BL	abort(SB)
//	MOVL	fn+0(FP), AX
//	MOVL	arg+4(FP), BX
//	MOVL	SP, CX
//	ANDL	$~15, SP	// alignment for gcc ABI
//	MOVL	CX, 4(SP)
//	MOVL	BX, 0(SP)
//	CALL	AX
//	MOVL	4(SP), SP
//	RET

TEXT emptyfunc(SB),0,$0
	RET

TEXT abort(SB),7,$0
	MOVW	$0, R0
	MOVW	(R0), R1

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
