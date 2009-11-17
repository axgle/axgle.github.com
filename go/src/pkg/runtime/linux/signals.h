<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/linux/signals.h</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/linux/signals.h</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#define C SigCatch
#define I SigIgnore
#define R SigRestart

static SigTab sigtab[] = {
	/* 0 */	0, &#34;SIGNONE: no trap&#34;,
	/* 1 */	0, &#34;SIGHUP: terminal line hangup&#34;,
	/* 2 */	0, &#34;SIGINT: interrupt&#34;,
	/* 3 */	C, &#34;SIGQUIT: quit&#34;,
	/* 4 */	C, &#34;SIGILL: illegal instruction&#34;,
	/* 5 */	C, &#34;SIGTRAP: trace trap&#34;,
	/* 6 */	C, &#34;SIGABRT: abort&#34;,
	/* 7 */	C, &#34;SIGBUS: bus error&#34;,
	/* 8 */	C, &#34;SIGFPE: floating-point exception&#34;,
	/* 9 */	0, &#34;SIGKILL: kill&#34;,
	/* 10 */	0, &#34;SIGUSR1: user-defined signal 1&#34;,
	/* 11 */	C, &#34;SIGSEGV: segmentation violation&#34;,
	/* 12 */	0, &#34;SIGUSR2: user-defined signal 2&#34;,
	/* 13 */	I, &#34;SIGPIPE: write to broken pipe&#34;,
	/* 14 */	0, &#34;SIGALRM: alarm clock&#34;,
	/* 15 */	0, &#34;SIGTERM: termination&#34;,
	/* 16 */	0, &#34;SIGSTKFLT: stack fault&#34;,
	/* 17 */	I+R, &#34;SIGCHLD: child status has changed&#34;,
	/* 18 */	0, &#34;SIGCONT: continue&#34;,
	/* 19 */	0, &#34;SIGSTOP: stop, unblockable&#34;,
	/* 20 */	0, &#34;SIGTSTP: keyboard stop&#34;,
	/* 21 */	0, &#34;SIGTTIN: background read from tty&#34;,
	/* 22 */	0, &#34;SIGTTOU: background write to tty&#34;,
	/* 23 */	0, &#34;SIGURG: urgent condition on socket&#34;,
	/* 24 */	0, &#34;SIGXCPU: cpu limit exceeded&#34;,
	/* 25 */	0, &#34;SIGXFSZ: file size limit exceeded&#34;,
	/* 26 */	0, &#34;SIGVTALRM: virtual alarm clock&#34;,
	/* 27 */	0, &#34;SIGPROF: profiling alarm clock&#34;,
	/* 28 */	I+R, &#34;SIGWINCH: window size change&#34;,
	/* 29 */	0, &#34;SIGIO: i/o now possible&#34;,
	/* 30 */	0, &#34;SIGPWR: power failure restart&#34;,
	/* 31 */	C, &#34;SIGSYS: bad system call&#34;,
};
#undef C
#undef I
#undef R

#define	NSIG 32
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
