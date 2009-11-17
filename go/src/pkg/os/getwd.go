<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/os/getwd.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/os/getwd.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package os

<a id="L7"></a>import (
    <a id="L8"></a>&#34;syscall&#34;;
<a id="L9"></a>)

<a id="L11"></a><span class="comment">// Getwd returns a rooted path name corresponding to the</span>
<a id="L12"></a><span class="comment">// current directory.  If the current directory can be</span>
<a id="L13"></a><span class="comment">// reached via multiple paths (due to symbolic links),</span>
<a id="L14"></a><span class="comment">// Getwd may return any one of them.</span>
<a id="L15"></a>func Getwd() (string, Error) {
    <a id="L16"></a><span class="comment">// If the operating system provides a Getwd call, use it.</span>
    <a id="L17"></a>if syscall.ImplementsGetwd {
        <a id="L18"></a>s, e := syscall.Getwd();
        <a id="L19"></a>return s, NewSyscallError(&#34;getwd&#34;, e);
    <a id="L20"></a>}

    <a id="L22"></a><span class="comment">// Otherwise, we&#39;re trying to find our way back to &#34;.&#34;.</span>
    <a id="L23"></a>dot, err := Stat(&#34;.&#34;);
    <a id="L24"></a>if err != nil {
        <a id="L25"></a>return &#34;&#34;, err
    <a id="L26"></a>}

    <a id="L28"></a><span class="comment">// Clumsy but widespread kludge:</span>
    <a id="L29"></a><span class="comment">// if $PWD is set and matches &#34;.&#34;, use it.</span>
    <a id="L30"></a>pwd := Getenv(&#34;PWD&#34;);
    <a id="L31"></a>if len(pwd) &gt; 0 &amp;&amp; pwd[0] == &#39;/&#39; {
        <a id="L32"></a>d, err := Stat(pwd);
        <a id="L33"></a>if err == nil &amp;&amp; d.Dev == dot.Dev &amp;&amp; d.Ino == dot.Ino {
            <a id="L34"></a>return pwd, nil
        <a id="L35"></a>}
    <a id="L36"></a>}

    <a id="L38"></a><span class="comment">// Root is a special case because it has no parent</span>
    <a id="L39"></a><span class="comment">// and ends in a slash.</span>
    <a id="L40"></a>root, err := Stat(&#34;/&#34;);
    <a id="L41"></a>if err != nil {
        <a id="L42"></a><span class="comment">// Can&#39;t stat root - no hope.</span>
        <a id="L43"></a>return &#34;&#34;, err
    <a id="L44"></a>}
    <a id="L45"></a>if root.Dev == dot.Dev &amp;&amp; root.Ino == dot.Ino {
        <a id="L46"></a>return &#34;/&#34;, nil
    <a id="L47"></a>}

    <a id="L49"></a><span class="comment">// General algorithm: find name in parent</span>
    <a id="L50"></a><span class="comment">// and then find name of parent.  Each iteration</span>
    <a id="L51"></a><span class="comment">// adds /name to the beginning of pwd.</span>
    <a id="L52"></a>pwd = &#34;&#34;;
    <a id="L53"></a>for parent := &#34;..&#34;; ; parent = &#34;../&#34; + parent {
        <a id="L54"></a>if len(parent) &gt;= 1024 { <span class="comment">// Sanity check</span>
            <a id="L55"></a>return &#34;&#34;, ENAMETOOLONG
        <a id="L56"></a>}
        <a id="L57"></a>fd, err := Open(parent, O_RDONLY, 0);
        <a id="L58"></a>if err != nil {
            <a id="L59"></a>return &#34;&#34;, err
        <a id="L60"></a>}

        <a id="L62"></a>for {
            <a id="L63"></a>names, err := fd.Readdirnames(100);
            <a id="L64"></a>if err != nil {
                <a id="L65"></a>fd.Close();
                <a id="L66"></a>return &#34;&#34;, err;
            <a id="L67"></a>}
            <a id="L68"></a>for _, name := range names {
                <a id="L69"></a>d, _ := Lstat(parent + &#34;/&#34; + name);
                <a id="L70"></a>if d.Dev == dot.Dev &amp;&amp; d.Ino == dot.Ino {
                    <a id="L71"></a>pwd = &#34;/&#34; + name + pwd;
                    <a id="L72"></a>goto Found;
                <a id="L73"></a>}
            <a id="L74"></a>}
        <a id="L75"></a>}
        <a id="L76"></a>fd.Close();
        <a id="L77"></a>return &#34;&#34;, ENOENT;

    <a id="L79"></a>Found:
        <a id="L80"></a>pd, err := fd.Stat();
        <a id="L81"></a>if err != nil {
            <a id="L82"></a>return &#34;&#34;, err
        <a id="L83"></a>}
        <a id="L84"></a>fd.Close();
        <a id="L85"></a>if pd.Dev == root.Dev &amp;&amp; pd.Ino == root.Ino {
            <a id="L86"></a>break
        <a id="L87"></a>}
        <a id="L88"></a><span class="comment">// Set up for next round.</span>
        <a id="L89"></a>dot = pd;
    <a id="L90"></a>}
    <a id="L91"></a>return pwd, nil;
<a id="L92"></a>}
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
