<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/godefs/doc.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/godefs/doc.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">/*</span>

<a id="L7"></a><span class="comment">Godefs is a bootstrapping tool for porting the Go runtime to new systems.</span>
<a id="L8"></a><span class="comment">It translates C type declarations into C or Go type declarations</span>
<a id="L9"></a><span class="comment">with the same memory layout.</span>

<a id="L11"></a><span class="comment">Usage: godefs [-g package] [-c cc] [-f cc-arg]... [defs.c ...]</span>

<a id="L13"></a><span class="comment">Godefs takes as input a host-compilable C file that includes</span>
<a id="L14"></a><span class="comment">standard system headers.  From that input file, it generates</span>
<a id="L15"></a><span class="comment">a standalone (no #includes) C or Go file containing equivalent</span>
<a id="L16"></a><span class="comment">definitions.</span>

<a id="L18"></a><span class="comment">The input to godefs is a C input file that can be compiled by</span>
<a id="L19"></a><span class="comment">the host system&#39;s standard C compiler (typically gcc).</span>
<a id="L20"></a><span class="comment">This file is expected to define new types and enumerated constants</span>
<a id="L21"></a><span class="comment">whose names begin with $ (a legal identifier character in gcc).</span>
<a id="L22"></a><span class="comment">Godefs compile the given input file with the host compiler and</span>
<a id="L23"></a><span class="comment">then parses the debug info embedded in the assembly output.</span>
<a id="L24"></a><span class="comment">This is far easier than reading system headers on most machines.</span>

<a id="L26"></a><span class="comment">The output from godefs is either C output intended for the</span>
<a id="L27"></a><span class="comment">Plan 9 C compiler tool chain (6c, 8c, or 5c) or Go output.</span>

<a id="L29"></a><span class="comment">The options are:</span>

<a id="L31"></a><span class="comment">	-g package</span>
<a id="L32"></a><span class="comment">		generate Go output using the given package name.</span>
<a id="L33"></a><span class="comment">		In the Go output, struct fields have leading xx_ prefixes</span>
<a id="L34"></a><span class="comment">		removed and the first character capitalized (exported).</span>

<a id="L36"></a><span class="comment">	-c cc</span>
<a id="L37"></a><span class="comment">		set the name of the host system&#39;s C compiler (default &#34;gcc&#34;)</span>

<a id="L39"></a><span class="comment">	-f cc-arg</span>
<a id="L40"></a><span class="comment">		add cc-arg to the command line when invoking the system C compiler</span>
<a id="L41"></a><span class="comment">		(for example, -f -m64 to invoke gcc -m64).</span>
<a id="L42"></a><span class="comment">		Repeating this option adds multiple flags to the command line.</span>

<a id="L44"></a><span class="comment">For example, if this is x.c:</span>

<a id="L46"></a><span class="comment">	#include &lt;sys/stat.h&gt;</span>

<a id="L48"></a><span class="comment">	typedef struct timespec $Timespec;</span>
<a id="L49"></a><span class="comment">	enum {</span>
<a id="L50"></a><span class="comment">		$S_IFMT = S_IFMT,</span>
<a id="L51"></a><span class="comment">		$S_IFIFO = S_IFIFO,</span>
<a id="L52"></a><span class="comment">		$S_IFCHR = S_IFCHR,</span>
<a id="L53"></a><span class="comment">	};</span>

<a id="L55"></a><span class="comment">then &#34;godefs x.c&#34; generates:</span>

<a id="L57"></a><span class="comment">	// godefs x.c</span>
<a id="L58"></a><span class="comment">	// MACHINE GENERATED - DO NOT EDIT.</span>

<a id="L60"></a><span class="comment">	// Constants</span>
<a id="L61"></a><span class="comment">	enum {</span>
<a id="L62"></a><span class="comment">		S_IFMT = 0xf000,</span>
<a id="L63"></a><span class="comment">		S_IFIFO = 0x1000,</span>
<a id="L64"></a><span class="comment">		S_IFCHR = 0x2000,</span>
<a id="L65"></a><span class="comment">	};</span>

<a id="L67"></a><span class="comment">	// Types</span>
<a id="L68"></a><span class="comment">	#pragma pack on</span>

<a id="L70"></a><span class="comment">	typedef struct Timespec Timespec;</span>
<a id="L71"></a><span class="comment">	struct Timespec {</span>
<a id="L72"></a><span class="comment">		int64 tv_sec;</span>
<a id="L73"></a><span class="comment">		int64 tv_nsec;</span>
<a id="L74"></a><span class="comment">	};</span>
<a id="L75"></a><span class="comment">	#pragma pack off</span>

<a id="L77"></a><span class="comment">and &#34;godefs -g MyPackage x.c&#34; generates:</span>

<a id="L79"></a><span class="comment">	// godefs -g MyPackage x.c</span>
<a id="L80"></a><span class="comment">	// MACHINE GENERATED - DO NOT EDIT.</span>

<a id="L82"></a><span class="comment">	package MyPackage</span>

<a id="L84"></a><span class="comment">	// Constants</span>
<a id="L85"></a><span class="comment">	const (</span>
<a id="L86"></a><span class="comment">		S_IFMT = 0xf000;</span>
<a id="L87"></a><span class="comment">		S_IFIFO = 0x1000;</span>
<a id="L88"></a><span class="comment">		S_IFCHR = 0x2000;</span>
<a id="L89"></a><span class="comment">	)</span>

<a id="L91"></a><span class="comment">	// Types</span>

<a id="L93"></a><span class="comment">	type Timespec struct {</span>
<a id="L94"></a><span class="comment">		Sec int64;</span>
<a id="L95"></a><span class="comment">		Nsec int64;</span>
<a id="L96"></a><span class="comment">	}</span>

<a id="L98"></a><span class="comment">*/</span>
<a id="L99"></a>package documentation
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
