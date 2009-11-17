<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/godoc/doc.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/godoc/doc.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">/*</span>

<a id="L7"></a><span class="comment">Godoc extracts and generates documentation for Go programs.</span>

<a id="L9"></a><span class="comment">It has two modes.</span>

<a id="L11"></a><span class="comment">Without the -http flag, it prints plain text documentation to standard output and exits.</span>

<a id="L13"></a><span class="comment">	godoc fmt</span>
<a id="L14"></a><span class="comment">	godoc fmt Printf</span>

<a id="L16"></a><span class="comment">With the -http flag, it runs as a web server and presents the documentation as a web page.</span>

<a id="L18"></a><span class="comment">	godoc -http=:6060</span>

<a id="L20"></a><span class="comment">Usage:</span>
<a id="L21"></a><span class="comment">	godoc [flag] package [name ...]</span>

<a id="L23"></a><span class="comment">The flags are:</span>
<a id="L24"></a><span class="comment">	-v</span>
<a id="L25"></a><span class="comment">		verbose mode</span>
<a id="L26"></a><span class="comment">	-tabwidth=4</span>
<a id="L27"></a><span class="comment">		width of tabs in units of spaces</span>
<a id="L28"></a><span class="comment">	-cmdroot=&#34;src/cmd&#34;</span>
<a id="L29"></a><span class="comment">		root command source directory (if unrooted, relative to -goroot)</span>
<a id="L30"></a><span class="comment">	-tmplroot=&#34;lib/godoc&#34;</span>
<a id="L31"></a><span class="comment">		root template directory (if unrooted, relative to -goroot)</span>
<a id="L32"></a><span class="comment">	-pkgroot=&#34;src/pkg&#34;</span>
<a id="L33"></a><span class="comment">		root package source directory (if unrooted, relative to -goroot)</span>
<a id="L34"></a><span class="comment">	-html</span>
<a id="L35"></a><span class="comment">		print HTML in command-line mode</span>
<a id="L36"></a><span class="comment">	-goroot=$GOROOT</span>
<a id="L37"></a><span class="comment">		Go root directory</span>
<a id="L38"></a><span class="comment">	-http=</span>
<a id="L39"></a><span class="comment">		HTTP service address (e.g., &#39;127.0.0.1:6060&#39; or just &#39;:6060&#39;)</span>
<a id="L40"></a><span class="comment">	-sync=&#34;command&#34;</span>
<a id="L41"></a><span class="comment">		if this and -sync_minutes are set, run the argument as a</span>
<a id="L42"></a><span class="comment">		command every sync_minutes; it is intended to update the</span>
<a id="L43"></a><span class="comment">		repository holding the source files.</span>
<a id="L44"></a><span class="comment">	-sync_minutes=0</span>
<a id="L45"></a><span class="comment">		sync interval in minutes; sync is disabled if &lt;= 0</span>

<a id="L47"></a><span class="comment">When godoc runs as a web server, it creates a search index from all .go files</span>
<a id="L48"></a><span class="comment">under $GOROOT (excluding files starting with .). The index is created at startup</span>
<a id="L49"></a><span class="comment">and is automatically updated every time the -sync command terminates with exit</span>
<a id="L50"></a><span class="comment">status 0, indicating that files have changed.</span>

<a id="L52"></a><span class="comment">If the sync exit status is 1, godoc assumes that it succeeded without errors</span>
<a id="L53"></a><span class="comment">but that no files changed; the index is not updated in this case.</span>

<a id="L55"></a><span class="comment">In all other cases, sync is assumed to have failed and godoc backs off running</span>
<a id="L56"></a><span class="comment">sync exponentially (up to 1 day). As soon as sync succeeds again (exit status 0</span>
<a id="L57"></a><span class="comment">or 1), the normal sync rhythm is re-established.</span>

<a id="L59"></a><span class="comment">*/</span>
<a id="L60"></a>package documentation
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
