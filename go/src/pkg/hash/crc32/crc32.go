<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/hash/crc32/crc32.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/hash/crc32/crc32.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package implements the 32-bit cyclic redundancy check, or CRC-32, checksum.</span>
<a id="L6"></a><span class="comment">// See http://en.wikipedia.org/wiki/Cyclic_redundancy_check for information.</span>
<a id="L7"></a>package crc32

<a id="L9"></a>import (
    <a id="L10"></a>&#34;hash&#34;;
    <a id="L11"></a>&#34;os&#34;;
<a id="L12"></a>)

<a id="L14"></a><span class="comment">// The size of a CRC-32 checksum in bytes.</span>
<a id="L15"></a>const Size = 4

<a id="L17"></a><span class="comment">// Predefined polynomials.</span>
<a id="L18"></a>const (
    <a id="L19"></a><span class="comment">// Far and away the most common CRC-32 polynomial.</span>
    <a id="L20"></a><span class="comment">// Used by ethernet (IEEE 802.3), v.42, fddi, gzip, zip, png, mpeg-2, ...</span>
    <a id="L21"></a>IEEE = 0xedb88320;

    <a id="L23"></a><span class="comment">// Castagnoli&#39;s polynomial, used in iSCSI.</span>
    <a id="L24"></a><span class="comment">// Has better error detection characteristics than IEEE.</span>
    <a id="L25"></a><span class="comment">// http://dx.doi.org/10.1109/26.231911</span>
    <a id="L26"></a>Castagnoli = 0x82f63b78;

    <a id="L28"></a><span class="comment">// Koopman&#39;s polynomial.</span>
    <a id="L29"></a><span class="comment">// Also has better error detection characteristics than IEEE.</span>
    <a id="L30"></a><span class="comment">// http://dx.doi.org/10.1109/DSN.2002.1028931</span>
    <a id="L31"></a>Koopman = 0xeb31d82e;
<a id="L32"></a>)

<a id="L34"></a><span class="comment">// Table is a 256-word table representing the polynomial for efficient processing.</span>
<a id="L35"></a>type Table [256]uint32

<a id="L37"></a><span class="comment">// MakeTable returns the Table constructed from the specified polynomial.</span>
<a id="L38"></a>func MakeTable(poly uint32) *Table {
    <a id="L39"></a>t := new(Table);
    <a id="L40"></a>for i := 0; i &lt; 256; i++ {
        <a id="L41"></a>crc := uint32(i);
        <a id="L42"></a>for j := 0; j &lt; 8; j++ {
            <a id="L43"></a>if crc&amp;1 == 1 {
                <a id="L44"></a>crc = (crc &gt;&gt; 1) ^ poly
            <a id="L45"></a>} else {
                <a id="L46"></a>crc &gt;&gt;= 1
            <a id="L47"></a>}
        <a id="L48"></a>}
        <a id="L49"></a>t[i] = crc;
    <a id="L50"></a>}
    <a id="L51"></a>return t;
<a id="L52"></a>}

<a id="L54"></a><span class="comment">// IEEETable is the table for the IEEE polynomial.</span>
<a id="L55"></a>var IEEETable = MakeTable(IEEE)

<a id="L57"></a><span class="comment">// digest represents the partial evaluation of a checksum.</span>
<a id="L58"></a>type digest struct {
    <a id="L59"></a>crc uint32;
    <a id="L60"></a>tab *Table;
<a id="L61"></a>}

<a id="L63"></a><span class="comment">// New creates a new hash.Hash32 computing the CRC-32 checksum</span>
<a id="L64"></a><span class="comment">// using the polynomial represented by the Table.</span>
<a id="L65"></a>func New(tab *Table) hash.Hash32 { return &amp;digest{0, tab} }

<a id="L67"></a><span class="comment">// NewIEEE creates a new hash.Hash32 computing the CRC-32 checksum</span>
<a id="L68"></a><span class="comment">// using the IEEE polynomial.</span>
<a id="L69"></a>func NewIEEE() hash.Hash32 { return New(IEEETable) }

<a id="L71"></a>func (d *digest) Size() int { return Size }

<a id="L73"></a>func (d *digest) Reset() { d.crc = 0 }

<a id="L75"></a>func update(crc uint32, tab *Table, p []byte) uint32 {
    <a id="L76"></a>crc = ^crc;
    <a id="L77"></a>for i := 0; i &lt; len(p); i++ {
        <a id="L78"></a>crc = tab[byte(crc)^p[i]] ^ (crc &gt;&gt; 8)
    <a id="L79"></a>}
    <a id="L80"></a>return ^crc;
<a id="L81"></a>}

<a id="L83"></a>func (d *digest) Write(p []byte) (n int, err os.Error) {
    <a id="L84"></a>d.crc = update(d.crc, d.tab, p);
    <a id="L85"></a>return len(p), nil;
<a id="L86"></a>}

<a id="L88"></a>func (d *digest) Sum32() uint32 { return d.crc }

<a id="L90"></a>func (d *digest) Sum() []byte {
    <a id="L91"></a>p := make([]byte, 4);
    <a id="L92"></a>s := d.Sum32();
    <a id="L93"></a>p[0] = byte(s &gt;&gt; 24);
    <a id="L94"></a>p[1] = byte(s &gt;&gt; 16);
    <a id="L95"></a>p[2] = byte(s &gt;&gt; 8);
    <a id="L96"></a>p[3] = byte(s);
    <a id="L97"></a>return p;
<a id="L98"></a>}

<a id="L100"></a><span class="comment">// Checksum returns the CRC-32 checksum of data</span>
<a id="L101"></a><span class="comment">// using the polynomial represented by the Table.</span>
<a id="L102"></a>func Checksum(data []byte, tab *Table) uint32 { return update(0, tab, data) }

<a id="L104"></a><span class="comment">// ChecksumIEEE returns the CRC-32 checksum of data</span>
<a id="L105"></a><span class="comment">// using the IEEE polynomial.</span>
<a id="L106"></a>func ChecksumIEEE(data []byte) uint32 { return update(0, IEEETable, data) }
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
