<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/os/path.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/os/path.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package os


<a id="L8"></a><span class="comment">// MkdirAll creates a directory named path,</span>
<a id="L9"></a><span class="comment">// along with any necessary parents, and returns nil,</span>
<a id="L10"></a><span class="comment">// or else returns an error.</span>
<a id="L11"></a><span class="comment">// The permission bits perm are used for all</span>
<a id="L12"></a><span class="comment">// directories that MkdirAll creates.</span>
<a id="L13"></a><span class="comment">// If path is already a directory, MkdirAll does nothing</span>
<a id="L14"></a><span class="comment">// and returns nil.</span>
<a id="L15"></a>func MkdirAll(path string, perm int) Error {
    <a id="L16"></a><span class="comment">// If path exists, stop with success or error.</span>
    <a id="L17"></a>dir, err := Lstat(path);
    <a id="L18"></a>if err == nil {
        <a id="L19"></a>if dir.IsDirectory() {
            <a id="L20"></a>return nil
        <a id="L21"></a>}
        <a id="L22"></a>return &amp;PathError{&#34;mkdir&#34;, path, ENOTDIR};
    <a id="L23"></a>}

    <a id="L25"></a><span class="comment">// Doesn&#39;t already exist; make sure parent does.</span>
    <a id="L26"></a>i := len(path);
    <a id="L27"></a>for i &gt; 0 &amp;&amp; path[i-1] == &#39;/&#39; { <span class="comment">// Skip trailing slashes.</span>
        <a id="L28"></a>i--
    <a id="L29"></a>}

    <a id="L31"></a>j := i;
    <a id="L32"></a>for j &gt; 0 &amp;&amp; path[j-1] != &#39;/&#39; { <span class="comment">// Scan backward over element.</span>
        <a id="L33"></a>j--
    <a id="L34"></a>}

    <a id="L36"></a>if j &gt; 0 {
        <a id="L37"></a><span class="comment">// Create parent</span>
        <a id="L38"></a>err = MkdirAll(path[0:j-1], perm);
        <a id="L39"></a>if err != nil {
            <a id="L40"></a>return err
        <a id="L41"></a>}
    <a id="L42"></a>}

    <a id="L44"></a><span class="comment">// Now parent exists, try to create.</span>
    <a id="L45"></a>err = Mkdir(path, perm);
    <a id="L46"></a>if err != nil {
        <a id="L47"></a><span class="comment">// Handle arguments like &#34;foo/.&#34; by</span>
        <a id="L48"></a><span class="comment">// double-checking that directory doesn&#39;t exist.</span>
        <a id="L49"></a>dir, err1 := Lstat(path);
        <a id="L50"></a>if err1 == nil &amp;&amp; dir.IsDirectory() {
            <a id="L51"></a>return nil
        <a id="L52"></a>}
        <a id="L53"></a>return err;
    <a id="L54"></a>}
    <a id="L55"></a>return nil;
<a id="L56"></a>}

<a id="L58"></a><span class="comment">// RemoveAll removes path and any children it contains.</span>
<a id="L59"></a><span class="comment">// It removes everything it can but returns the first error</span>
<a id="L60"></a><span class="comment">// it encounters.  If the path does not exist, RemoveAll</span>
<a id="L61"></a><span class="comment">// returns nil (no error).</span>
<a id="L62"></a>func RemoveAll(path string) Error {
    <a id="L63"></a><span class="comment">// Simple case: if Remove works, we&#39;re done.</span>
    <a id="L64"></a>err := Remove(path);
    <a id="L65"></a>if err == nil {
        <a id="L66"></a>return nil
    <a id="L67"></a>}

    <a id="L69"></a><span class="comment">// Otherwise, is this a directory we need to recurse into?</span>
    <a id="L70"></a>dir, serr := Lstat(path);
    <a id="L71"></a>if serr != nil {
        <a id="L72"></a>if serr, ok := serr.(*PathError); ok &amp;&amp; serr.Error == ENOENT {
            <a id="L73"></a>return nil
        <a id="L74"></a>}
        <a id="L75"></a>return serr;
    <a id="L76"></a>}
    <a id="L77"></a>if !dir.IsDirectory() {
        <a id="L78"></a><span class="comment">// Not a directory; return the error from Remove.</span>
        <a id="L79"></a>return err
    <a id="L80"></a>}

    <a id="L82"></a><span class="comment">// Directory.</span>
    <a id="L83"></a>fd, err := Open(path, O_RDONLY, 0);
    <a id="L84"></a>if err != nil {
        <a id="L85"></a>return err
    <a id="L86"></a>}
    <a id="L87"></a>defer fd.Close();

    <a id="L89"></a><span class="comment">// Remove contents &amp; return first error.</span>
    <a id="L90"></a>err = nil;
    <a id="L91"></a>for {
        <a id="L92"></a>names, err1 := fd.Readdirnames(100);
        <a id="L93"></a>for _, name := range names {
            <a id="L94"></a>err1 := RemoveAll(path + &#34;/&#34; + name);
            <a id="L95"></a>if err == nil {
                <a id="L96"></a>err = err1
            <a id="L97"></a>}
        <a id="L98"></a>}
        <a id="L99"></a><span class="comment">// If Readdirnames returned an error, use it.</span>
        <a id="L100"></a>if err == nil {
            <a id="L101"></a>err = err1
        <a id="L102"></a>}
        <a id="L103"></a>if len(names) == 0 {
            <a id="L104"></a>break
        <a id="L105"></a>}
    <a id="L106"></a>}

    <a id="L108"></a><span class="comment">// Remove directory.</span>
    <a id="L109"></a>err1 := Remove(path);
    <a id="L110"></a>if err == nil {
        <a id="L111"></a>err = err1
    <a id="L112"></a>}
    <a id="L113"></a>return err;
<a id="L114"></a>}
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
