<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/dwarf/buf.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/debug/dwarf/buf.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Buffered reading and decoding of DWARF data streams.</span>

<a id="L7"></a>package dwarf

<a id="L9"></a>import (
    <a id="L10"></a>&#34;encoding/binary&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;strconv&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">// Data buffer being decoded.</span>
<a id="L16"></a>type buf struct {
    <a id="L17"></a>dwarf    *Data;
    <a id="L18"></a>order    binary.ByteOrder;
    <a id="L19"></a>name     string;
    <a id="L20"></a>off      Offset;
    <a id="L21"></a>data     []byte;
    <a id="L22"></a>addrsize int;
    <a id="L23"></a>err      os.Error;
<a id="L24"></a>}

<a id="L26"></a>func makeBuf(d *Data, name string, off Offset, data []byte, addrsize int) buf {
    <a id="L27"></a>return buf{d, d.order, name, off, data, addrsize, nil}
<a id="L28"></a>}

<a id="L30"></a>func (b *buf) uint8() uint8 {
    <a id="L31"></a>if len(b.data) &lt; 1 {
        <a id="L32"></a>b.error(&#34;underflow&#34;);
        <a id="L33"></a>return 0;
    <a id="L34"></a>}
    <a id="L35"></a>val := b.data[0];
    <a id="L36"></a>b.data = b.data[1:len(b.data)];
    <a id="L37"></a>b.off++;
    <a id="L38"></a>return val;
<a id="L39"></a>}

<a id="L41"></a>func (b *buf) bytes(n int) []byte {
    <a id="L42"></a>if len(b.data) &lt; n {
        <a id="L43"></a>b.error(&#34;underflow&#34;);
        <a id="L44"></a>return nil;
    <a id="L45"></a>}
    <a id="L46"></a>data := b.data[0:n];
    <a id="L47"></a>b.data = b.data[n:len(b.data)];
    <a id="L48"></a>b.off += Offset(n);
    <a id="L49"></a>return data;
<a id="L50"></a>}

<a id="L52"></a>func (b *buf) skip(n int) { b.bytes(n) }

<a id="L54"></a>func (b *buf) string() string {
    <a id="L55"></a>for i := 0; i &lt; len(b.data); i++ {
        <a id="L56"></a>if b.data[i] == 0 {
            <a id="L57"></a>s := string(b.data[0:i]);
            <a id="L58"></a>b.data = b.data[i+1 : len(b.data)];
            <a id="L59"></a>b.off += Offset(i + 1);
            <a id="L60"></a>return s;
        <a id="L61"></a>}
    <a id="L62"></a>}
    <a id="L63"></a>b.error(&#34;underflow&#34;);
    <a id="L64"></a>return &#34;&#34;;
<a id="L65"></a>}

<a id="L67"></a>func (b *buf) uint16() uint16 {
    <a id="L68"></a>a := b.bytes(2);
    <a id="L69"></a>if a == nil {
        <a id="L70"></a>return 0
    <a id="L71"></a>}
    <a id="L72"></a>return b.order.Uint16(a);
<a id="L73"></a>}

<a id="L75"></a>func (b *buf) uint32() uint32 {
    <a id="L76"></a>a := b.bytes(4);
    <a id="L77"></a>if a == nil {
        <a id="L78"></a>return 0
    <a id="L79"></a>}
    <a id="L80"></a>return b.order.Uint32(a);
<a id="L81"></a>}

<a id="L83"></a>func (b *buf) uint64() uint64 {
    <a id="L84"></a>a := b.bytes(8);
    <a id="L85"></a>if a == nil {
        <a id="L86"></a>return 0
    <a id="L87"></a>}
    <a id="L88"></a>return b.order.Uint64(a);
<a id="L89"></a>}

<a id="L91"></a><span class="comment">// Read a varint, which is 7 bits per byte, little endian.</span>
<a id="L92"></a><span class="comment">// the 0x80 bit means read another byte.</span>
<a id="L93"></a>func (b *buf) varint() (c uint64, bits uint) {
    <a id="L94"></a>for i := 0; i &lt; len(b.data); i++ {
        <a id="L95"></a>byte := b.data[i];
        <a id="L96"></a>c |= uint64(byte&amp;0x7F) &lt;&lt; bits;
        <a id="L97"></a>bits += 7;
        <a id="L98"></a>if byte&amp;0x80 == 0 {
            <a id="L99"></a>b.off += Offset(i + 1);
            <a id="L100"></a>b.data = b.data[i+1 : len(b.data)];
            <a id="L101"></a>return c, bits;
        <a id="L102"></a>}
    <a id="L103"></a>}
    <a id="L104"></a>return 0, 0;
<a id="L105"></a>}

<a id="L107"></a><span class="comment">// Unsigned int is just a varint.</span>
<a id="L108"></a>func (b *buf) uint() uint64 {
    <a id="L109"></a>x, _ := b.varint();
    <a id="L110"></a>return x;
<a id="L111"></a>}

<a id="L113"></a><span class="comment">// Signed int is a sign-extended varint.</span>
<a id="L114"></a>func (b *buf) int() int64 {
    <a id="L115"></a>ux, bits := b.varint();
    <a id="L116"></a>x := int64(ux);
    <a id="L117"></a>if x&amp;(1&lt;&lt;(bits-1)) != 0 {
        <a id="L118"></a>x |= -1 &lt;&lt; bits
    <a id="L119"></a>}
    <a id="L120"></a>return x;
<a id="L121"></a>}

<a id="L123"></a><span class="comment">// Address-sized uint.</span>
<a id="L124"></a>func (b *buf) addr() uint64 {
    <a id="L125"></a>switch b.addrsize {
    <a id="L126"></a>case 1:
        <a id="L127"></a>return uint64(b.uint8())
    <a id="L128"></a>case 2:
        <a id="L129"></a>return uint64(b.uint16())
    <a id="L130"></a>case 4:
        <a id="L131"></a>return uint64(b.uint32())
    <a id="L132"></a>case 8:
        <a id="L133"></a>return uint64(b.uint64())
    <a id="L134"></a>}
    <a id="L135"></a>b.error(&#34;unknown address size&#34;);
    <a id="L136"></a>return 0;
<a id="L137"></a>}

<a id="L139"></a>func (b *buf) error(s string) {
    <a id="L140"></a>if b.err == nil {
        <a id="L141"></a>b.data = nil;
        <a id="L142"></a>b.err = DecodeError{b.name, b.off, s};
    <a id="L143"></a>}
<a id="L144"></a>}

<a id="L146"></a>type DecodeError struct {
    <a id="L147"></a>Name   string;
    <a id="L148"></a>Offset Offset;
    <a id="L149"></a>Error  string;
<a id="L150"></a>}

<a id="L152"></a>func (e DecodeError) String() string {
    <a id="L153"></a>return &#34;decoding dwarf section &#34; + e.Name + &#34; at offset 0x&#34; + strconv.Itob64(int64(e.Offset), 16) + &#34;: &#34; + e.Error
<a id="L154"></a>}
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
