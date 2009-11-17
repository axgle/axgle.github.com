<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/arm/traceback.c</title>

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
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/arm/traceback.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;runtime.h&#34;

// TODO(rsc): Move this into portable code, with calls to a
// machine-dependent isclosure() function.

void
traceback(byte *pc0, byte *sp, G *g)
{
// 	Stktop *stk;
// 	uintptr pc;
// 	int32 i, n;
// 	Func *f;
// 	byte *p;

// 	pc = (uintptr)pc0;

// 	// If the PC is zero, it&#39;s likely a nil function call.
// 	// Start in the caller&#39;s frame.
// 	if(pc == 0) {
// 		pc = *(uintptr*)sp;
// 		sp += sizeof(uintptr);
// 	}

// 	stk = (Stktop*)g-&gt;stackbase;
// 	for(n=0; n&lt;100; n++) {
// 		while(pc == (uintptr)retfromnewstack) {
// 			// pop to earlier stack block
// 			sp = stk-&gt;oldsp;
// 			stk = (Stktop*)stk-&gt;oldbase;
// 			pc = *(uintptr*)(sp+sizeof(uintptr));
// 			sp += 2*sizeof(uintptr);	// two irrelevant calls on stack: morestack plus its call
// 		}
// 		f = findfunc(pc);
// 		if(f == nil) {
// 			// dangerous, but poke around to see if it is a closure
// 			p = (byte*)pc;
// 			// ADDL $xxx, SP; RET
// 			if(p[0] == 0x81 &amp;&amp; p[1] == 0xc4 &amp;&amp; p[6] == 0xc3) {
// 				sp += *(uint32*)(p+2) + 8;
// 				pc = *(uintptr*)(sp - 8);
// 				if(pc &lt;= 0x1000)
// 					return;
// 				continue;
// 			}
// 			printf(&#34;%p unknown pc\n&#34;, pc);
// 			return;
// 		}
// 		if(f-&gt;frame &lt; sizeof(uintptr))	// assembly funcs say 0 but lie
// 			sp += sizeof(uintptr);
// 		else
// 			sp += f-&gt;frame;

// 		// print this frame
// 		//	main+0xf /home/rsc/go/src/runtime/x.go:23
// 		//		main(0x1, 0x2, 0x3)
// 		printf(&#34;%S&#34;, f-&gt;name);
// 		if(pc &gt; f-&gt;entry)
// 			printf(&#34;+%p&#34;, (uintptr)(pc - f-&gt;entry));
// 		printf(&#34; %S:%d\n&#34;, f-&gt;src, funcline(f, pc-1));	// -1 to get to CALL instr.
// 		printf(&#34;\t%S(&#34;, f-&gt;name);
// 		for(i = 0; i &lt; f-&gt;args; i++) {
// 			if(i != 0)
// 				prints(&#34;, &#34;);
// 			runtime·printhex(((uint32*)sp)[i]);
// 			if(i &gt;= 4) {
// 				prints(&#34;, ...&#34;);
// 				break;
// 			}
// 		}
// 		prints(&#34;)\n&#34;);

// 		pc = *(uintptr*)(sp-sizeof(uintptr));
// 		if(pc &lt;= 0x1000)
// 			return;
// 	}
// 	prints(&#34;...\n&#34;);
}

// func caller(n int) (pc uintptr, file string, line int, ok bool)
void
runtime·Caller(int32 n, uintptr retpc, String retfile, int32 retline, bool retbool)
{
// 	uintptr pc;
// 	byte *sp;
// 	byte *p;
// 	Stktop *stk;
// 	Func *f;

// 	// our caller&#39;s pc, sp.
// 	sp = (byte*)&amp;n;
// 	pc = *((uintptr*)sp - 1);
// 	if((f = findfunc(pc)) == nil) {
// 	error:
// 		retpc = 0;
// 		retline = 0;
// 		retfile = emptystring;
// 		retbool = false;
// 		FLUSH(&amp;retpc);
// 		FLUSH(&amp;retfile);
// 		FLUSH(&amp;retline);
// 		FLUSH(&amp;retbool);
// 		return;
// 	}

// 	// now unwind n levels
// 	stk = (Stktop*)g-&gt;stackbase;
// 	while(n-- &gt; 0) {
// 		while(pc == (uintptr)retfromnewstack) {
// 			sp = stk-&gt;oldsp;
// 			stk = (Stktop*)stk-&gt;oldbase;
// 			pc = *((uintptr*)sp + 1);
// 			sp += 2*sizeof(uintptr);
// 		}

// 		if(f-&gt;frame &lt; sizeof(uintptr))	// assembly functions lie
// 			sp += sizeof(uintptr);
// 		else
// 			sp += f-&gt;frame;

// 	loop:
// 		pc = *((uintptr*)sp - 1);
// 		if(pc &lt;= 0x1000 || (f = findfunc(pc)) == nil) {
// 			// dangerous, but let&#39;s try this.
// 			// see if it is a closure.
// 			p = (byte*)pc;
// 			// ADDL $xxx, SP; RET
// 			if(p[0] == 0x81 &amp;&amp; p[1] == 0xc4 &amp;&amp; p[6] == 0xc3) {
// 				sp += *(uint32*)(p+2) + sizeof(uintptr);
// 				goto loop;
// 			}
// 			goto error;
// 		}
// 	}

// 	retpc = pc;
// 	retfile = f-&gt;src;
// 	retline = funcline(f, pc-1);
// 	retbool = true;
// 	FLUSH(&amp;retpc);
// 	FLUSH(&amp;retfile);
// 	FLUSH(&amp;retline);
// 	FLUSH(&amp;retbool);
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
