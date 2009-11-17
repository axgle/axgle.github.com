<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/archive/tar/common.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/archive/tar/common.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The tar package implements access to tar archives.</span>
<a id="L6"></a><span class="comment">// It aims to cover most of the variations, including those produced</span>
<a id="L7"></a><span class="comment">// by GNU and BSD tars.</span>
<a id="L8"></a><span class="comment">//</span>
<a id="L9"></a><span class="comment">// References:</span>
<a id="L10"></a><span class="comment">//   http://www.freebsd.org/cgi/man.cgi?query=tar&amp;sektion=5</span>
<a id="L11"></a><span class="comment">//   http://www.gnu.org/software/tar/manual/html_node/Standard.html</span>
<a id="L12"></a>package tar

<a id="L14"></a>const (
    <a id="L15"></a>blockSize = 512;

    <a id="L17"></a><span class="comment">// Types</span>
    <a id="L18"></a>TypeReg           = &#39;0&#39;;
    <a id="L19"></a>TypeRegA          = &#39;\x00&#39;;
    <a id="L20"></a>TypeLink          = &#39;1&#39;;
    <a id="L21"></a>TypeSymlink       = &#39;2&#39;;
    <a id="L22"></a>TypeChar          = &#39;3&#39;;
    <a id="L23"></a>TypeBlock         = &#39;4&#39;;
    <a id="L24"></a>TypeDir           = &#39;5&#39;;
    <a id="L25"></a>TypeFifo          = &#39;6&#39;;
    <a id="L26"></a>TypeCont          = &#39;7&#39;;
    <a id="L27"></a>TypeXHeader       = &#39;x&#39;;
    <a id="L28"></a>TypeXGlobalHeader = &#39;g&#39;;
<a id="L29"></a>)

<a id="L31"></a><span class="comment">// A Header represents a single header in a tar archive.</span>
<a id="L32"></a><span class="comment">// Some fields may not be populated.</span>
<a id="L33"></a>type Header struct {
    <a id="L34"></a>Name     string;
    <a id="L35"></a>Mode     int64;
    <a id="L36"></a>Uid      int64;
    <a id="L37"></a>Gid      int64;
    <a id="L38"></a>Size     int64;
    <a id="L39"></a>Mtime    int64;
    <a id="L40"></a>Typeflag byte;
    <a id="L41"></a>Linkname string;
    <a id="L42"></a>Uname    string;
    <a id="L43"></a>Gname    string;
    <a id="L44"></a>Devmajor int64;
    <a id="L45"></a>Devminor int64;
    <a id="L46"></a>Atime    int64;
    <a id="L47"></a>Ctime    int64;
<a id="L48"></a>}

<a id="L50"></a>var zeroBlock = make([]byte, blockSize)

<a id="L52"></a><span class="comment">// POSIX specifies a sum of the unsigned byte values, but the Sun tar uses signed byte values.</span>
<a id="L53"></a><span class="comment">// We compute and return both.</span>
<a id="L54"></a>func checksum(header []byte) (unsigned int64, signed int64) {
    <a id="L55"></a>for i := 0; i &lt; len(header); i++ {
        <a id="L56"></a>if i == 148 {
            <a id="L57"></a><span class="comment">// The chksum field (header[148:156]) is special: it should be treated as space bytes.</span>
            <a id="L58"></a>unsigned += &#39; &#39; * 8;
            <a id="L59"></a>signed += &#39; &#39; * 8;
            <a id="L60"></a>i += 7;
            <a id="L61"></a>continue;
        <a id="L62"></a>}
        <a id="L63"></a>unsigned += int64(header[i]);
        <a id="L64"></a>signed += int64(int8(header[i]));
    <a id="L65"></a>}
    <a id="L66"></a>return;
<a id="L67"></a>}

<a id="L69"></a>type slicer []byte

<a id="L71"></a>func (sp *slicer) next(n int) (b []byte) {
    <a id="L72"></a>s := *sp;
    <a id="L73"></a>b, *sp = s[0:n], s[n:len(s)];
    <a id="L74"></a>return;
<a id="L75"></a>}
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
