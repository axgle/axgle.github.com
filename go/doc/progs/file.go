<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /doc/progs/file.go</title>

  <link rel="stylesheet" type="text/css" href="../style.css">
  <script type="text/javascript" src="../godocs.js"></script>

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
        <a href="../../index.html"><img src="../logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../go_tutorial.html">Tutorial</a></li>
    <li><a href="../effective_go.html">Effective Go</a></li>
    <li><a href="../go_faq.html">FAQ</a></li>
    <li><a href="../go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../go_spec.html">Language Specification</a></li>
    <li><a href="../go_mem.html">Memory Model</a></li>
    <li><a href="../go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../install.html">Install Go</a></li>
    <li><a href="../contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../src/index.html">Source files</a></li>

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
  <h1 id="generatedHeader">Source file /doc/progs/file.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package file

<a id="L7"></a>import (
    <a id="L8"></a>&#34;os&#34;;
    <a id="L9"></a>&#34;syscall&#34;;
<a id="L10"></a>)

<a id="L12"></a>type File struct {
    <a id="L13"></a>fd   int;    <span class="comment">// file descriptor number</span>
    <a id="L14"></a>name string; <span class="comment">// file name at Open time</span>
<a id="L15"></a>}

<a id="L17"></a>func newFile(fd int, name string) *File {
    <a id="L18"></a>if fd &lt; 0 {
        <a id="L19"></a>return nil
    <a id="L20"></a>}
    <a id="L21"></a>return &amp;File{fd, name};
<a id="L22"></a>}

<a id="L24"></a>var (
    <a id="L25"></a>Stdin  = newFile(0, &#34;/dev/stdin&#34;);
    <a id="L26"></a>Stdout = newFile(1, &#34;/dev/stdout&#34;);
    <a id="L27"></a>Stderr = newFile(2, &#34;/dev/stderr&#34;);
<a id="L28"></a>)

<a id="L30"></a>func Open(name string, mode int, perm int) (file *File, err os.Error) {
    <a id="L31"></a>r, e := syscall.Open(name, mode, perm);
    <a id="L32"></a>if e != 0 {
        <a id="L33"></a>err = os.Errno(e)
    <a id="L34"></a>}
    <a id="L35"></a>return newFile(r, name), err;
<a id="L36"></a>}

<a id="L38"></a>func (file *File) Close() os.Error {
    <a id="L39"></a>if file == nil {
        <a id="L40"></a>return os.EINVAL
    <a id="L41"></a>}
    <a id="L42"></a>e := syscall.Close(file.fd);
    <a id="L43"></a>file.fd = -1; <span class="comment">// so it can&#39;t be closed again</span>
    <a id="L44"></a>if e != 0 {
        <a id="L45"></a>return os.Errno(e)
    <a id="L46"></a>}
    <a id="L47"></a>return nil;
<a id="L48"></a>}

<a id="L50"></a>func (file *File) Read(b []byte) (ret int, err os.Error) {
    <a id="L51"></a>if file == nil {
        <a id="L52"></a>return -1, os.EINVAL
    <a id="L53"></a>}
    <a id="L54"></a>r, e := syscall.Read(file.fd, b);
    <a id="L55"></a>if e != 0 {
        <a id="L56"></a>err = os.Errno(e)
    <a id="L57"></a>}
    <a id="L58"></a>return int(r), err;
<a id="L59"></a>}

<a id="L61"></a>func (file *File) Write(b []byte) (ret int, err os.Error) {
    <a id="L62"></a>if file == nil {
        <a id="L63"></a>return -1, os.EINVAL
    <a id="L64"></a>}
    <a id="L65"></a>r, e := syscall.Write(file.fd, b);
    <a id="L66"></a>if e != 0 {
        <a id="L67"></a>err = os.Errno(e)
    <a id="L68"></a>}
    <a id="L69"></a>return int(r), err;
<a id="L70"></a>}

<a id="L72"></a>func (file *File) String() string { <a id="L73"></a>return file.name <a id="L74"></a>}
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
