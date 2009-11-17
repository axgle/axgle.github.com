<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/gosym/pclntab.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/debug/gosym/pclntab.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">/*</span>
<a id="L6"></a><span class="comment"> * Line tables</span>
<a id="L7"></a><span class="comment"> */</span>

<a id="L9"></a>package gosym

<a id="L11"></a>import &#34;encoding/binary&#34;

<a id="L13"></a>type LineTable struct {
    <a id="L14"></a>Data []byte;
    <a id="L15"></a>PC   uint64;
    <a id="L16"></a>Line int;
<a id="L17"></a>}

<a id="L19"></a><span class="comment">// TODO(rsc): Need to pull in quantum from architecture definition.</span>
<a id="L20"></a>const quantum = 1

<a id="L22"></a>func (t *LineTable) parse(targetPC uint64, targetLine int) (b []byte, pc uint64, line int) {
    <a id="L23"></a><span class="comment">// The PC/line table can be thought of as a sequence of</span>
    <a id="L24"></a><span class="comment">//  &lt;pc update&gt;* &lt;line update&gt;</span>
    <a id="L25"></a><span class="comment">// batches.  Each update batch results in a (pc, line) pair,</span>
    <a id="L26"></a><span class="comment">// where line applies to every PC from pc up to but not</span>
    <a id="L27"></a><span class="comment">// including the pc of the next pair.</span>
    <a id="L28"></a><span class="comment">//</span>
    <a id="L29"></a><span class="comment">// Here we process each update individually, which simplifies</span>
    <a id="L30"></a><span class="comment">// the code, but makes the corner cases more confusing.</span>
    <a id="L31"></a>b, pc, line = t.Data, t.PC, t.Line;
    <a id="L32"></a>for pc &lt;= targetPC &amp;&amp; line != targetLine &amp;&amp; len(b) &gt; 0 {
        <a id="L33"></a>code := b[0];
        <a id="L34"></a>b = b[1:len(b)];
        <a id="L35"></a>switch {
        <a id="L36"></a>case code == 0:
            <a id="L37"></a>if len(b) &lt; 4 {
                <a id="L38"></a>b = b[0:0];
                <a id="L39"></a>break;
            <a id="L40"></a>}
            <a id="L41"></a>val := binary.BigEndian.Uint32(b);
            <a id="L42"></a>b = b[4:len(b)];
            <a id="L43"></a>line += int(val);
        <a id="L44"></a>case code &lt;= 64:
            <a id="L45"></a>line += int(code)
        <a id="L46"></a>case code &lt;= 128:
            <a id="L47"></a>line -= int(code - 64)
        <a id="L48"></a>default:
            <a id="L49"></a>pc += quantum * uint64(code-128);
            <a id="L50"></a>continue;
        <a id="L51"></a>}
        <a id="L52"></a>pc += quantum;
    <a id="L53"></a>}
    <a id="L54"></a>return b, pc, line;
<a id="L55"></a>}

<a id="L57"></a>func (t *LineTable) slice(pc uint64) *LineTable {
    <a id="L58"></a>data, pc, line := t.parse(pc, -1);
    <a id="L59"></a>return &amp;LineTable{data, pc, line};
<a id="L60"></a>}

<a id="L62"></a>func (t *LineTable) PCToLine(pc uint64) int {
    <a id="L63"></a>_, _, line := t.parse(pc, -1);
    <a id="L64"></a>return line;
<a id="L65"></a>}

<a id="L67"></a>func (t *LineTable) LineToPC(line int, maxpc uint64) uint64 {
    <a id="L68"></a>_, pc, line1 := t.parse(maxpc, line);
    <a id="L69"></a>if line1 != line {
        <a id="L70"></a>return 0
    <a id="L71"></a>}
    <a id="L72"></a><span class="comment">// Subtract quantum from PC to account for post-line increment</span>
    <a id="L73"></a>return pc - quantum;
<a id="L74"></a>}

<a id="L76"></a><span class="comment">// NewLineTable returns a new PC/line table</span>
<a id="L77"></a><span class="comment">// corresponding to the encoded data.</span>
<a id="L78"></a><span class="comment">// Text must be the start address of the</span>
<a id="L79"></a><span class="comment">// corresponding text segment.</span>
<a id="L80"></a>func NewLineTable(data []byte, text uint64) *LineTable {
    <a id="L81"></a>return &amp;LineTable{data, text, 0}
<a id="L82"></a>}
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
