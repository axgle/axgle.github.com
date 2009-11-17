<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/bytes/buffer.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/bytes/buffer.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package bytes

<a id="L7"></a><span class="comment">// Simple byte buffer for marshaling data.</span>

<a id="L9"></a>import (
    <a id="L10"></a>&#34;os&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">// Copy from string to byte array at offset doff.  Assume there&#39;s room.</span>
<a id="L14"></a>func copyString(dst []byte, doff int, str string) {
    <a id="L15"></a>for soff := 0; soff &lt; len(str); soff++ {
        <a id="L16"></a>dst[doff] = str[soff];
        <a id="L17"></a>doff++;
    <a id="L18"></a>}
<a id="L19"></a>}

<a id="L21"></a><span class="comment">// Copy from bytes to byte array at offset doff.  Assume there&#39;s room.</span>
<a id="L22"></a>func copyBytes(dst []byte, doff int, src []byte) {
    <a id="L23"></a>for soff := 0; soff &lt; len(src); soff++ {
        <a id="L24"></a>dst[doff] = src[soff];
        <a id="L25"></a>doff++;
    <a id="L26"></a>}
<a id="L27"></a>}

<a id="L29"></a><span class="comment">// A Buffer is a variable-sized buffer of bytes</span>
<a id="L30"></a><span class="comment">// with Read and Write methods.</span>
<a id="L31"></a><span class="comment">// The zero value for Buffer is an empty buffer ready to use.</span>
<a id="L32"></a>type Buffer struct {
    <a id="L33"></a>buf     []byte; <span class="comment">// contents are the bytes buf[off : len(buf)]</span>
    <a id="L34"></a>off     int;    <span class="comment">// read at &amp;buf[off], write at &amp;buf[len(buf)]</span>
    <a id="L35"></a>oneByte []byte; <span class="comment">// avoid allocation of slice on each WriteByte</span>
<a id="L36"></a>}

<a id="L38"></a><span class="comment">// Bytes returns the contents of the unread portion of the buffer;</span>
<a id="L39"></a><span class="comment">// len(b.Bytes()) == b.Len().</span>
<a id="L40"></a>func (b *Buffer) Bytes() []byte { return b.buf[b.off:len(b.buf)] }

<a id="L42"></a><span class="comment">// String returns the contents of the unread portion of the buffer</span>
<a id="L43"></a><span class="comment">// as a string.  If the Buffer is a nil pointer, it returns &#34;&lt;nil&gt;&#34;.</span>
<a id="L44"></a>func (b *Buffer) String() string {
    <a id="L45"></a>if b == nil {
        <a id="L46"></a><span class="comment">// Special case, useful in debugging.</span>
        <a id="L47"></a>return &#34;&lt;nil&gt;&#34;
    <a id="L48"></a>}
    <a id="L49"></a>return string(b.buf[b.off:len(b.buf)]);
<a id="L50"></a>}

<a id="L52"></a><span class="comment">// Len returns the number of bytes of the unread portion of the buffer;</span>
<a id="L53"></a><span class="comment">// b.Len() == len(b.Bytes()).</span>
<a id="L54"></a>func (b *Buffer) Len() int { return len(b.buf) - b.off }

<a id="L56"></a><span class="comment">// Truncate discards all but the first n unread bytes from the buffer.</span>
<a id="L57"></a><span class="comment">// It is an error to call b.Truncate(n) with n &gt; b.Len().</span>
<a id="L58"></a>func (b *Buffer) Truncate(n int) {
    <a id="L59"></a>if n == 0 {
        <a id="L60"></a><span class="comment">// Reuse buffer space.</span>
        <a id="L61"></a>b.off = 0
    <a id="L62"></a>}
    <a id="L63"></a>b.buf = b.buf[0 : b.off+n];
<a id="L64"></a>}

<a id="L66"></a><span class="comment">// Reset resets the buffer so it has no content.</span>
<a id="L67"></a><span class="comment">// b.Reset() is the same as b.Truncate(0).</span>
<a id="L68"></a>func (b *Buffer) Reset() { b.Truncate(0) }

<a id="L70"></a><span class="comment">// Write appends the contents of p to the buffer.  The return</span>
<a id="L71"></a><span class="comment">// value n is the length of p; err is always nil.</span>
<a id="L72"></a>func (b *Buffer) Write(p []byte) (n int, err os.Error) {
    <a id="L73"></a>m := b.Len();
    <a id="L74"></a>n = len(p);

    <a id="L76"></a>if len(b.buf)+n &gt; cap(b.buf) {
        <a id="L77"></a><span class="comment">// not enough space at end</span>
        <a id="L78"></a>buf := b.buf;
        <a id="L79"></a>if m+n &gt; cap(b.buf) {
            <a id="L80"></a><span class="comment">// not enough space anywhere</span>
            <a id="L81"></a>buf = make([]byte, 2*cap(b.buf)+n)
        <a id="L82"></a>}
        <a id="L83"></a>copyBytes(buf, 0, b.buf[b.off:b.off+m]);
        <a id="L84"></a>b.buf = buf;
        <a id="L85"></a>b.off = 0;
    <a id="L86"></a>}

    <a id="L88"></a>b.buf = b.buf[0 : b.off+m+n];
    <a id="L89"></a>copyBytes(b.buf, b.off+m, p);
    <a id="L90"></a>return n, nil;
<a id="L91"></a>}

<a id="L93"></a><span class="comment">// WriteString appends the contents of s to the buffer.  The return</span>
<a id="L94"></a><span class="comment">// value n is the length of s; err is always nil.</span>
<a id="L95"></a>func (b *Buffer) WriteString(s string) (n int, err os.Error) {
    <a id="L96"></a>m := b.Len();
    <a id="L97"></a>n = len(s);

    <a id="L99"></a>if len(b.buf)+n &gt; cap(b.buf) {
        <a id="L100"></a><span class="comment">// not enough space at end</span>
        <a id="L101"></a>buf := b.buf;
        <a id="L102"></a>if m+n &gt; cap(b.buf) {
            <a id="L103"></a><span class="comment">// not enough space anywhere</span>
            <a id="L104"></a>buf = make([]byte, 2*cap(b.buf)+n)
        <a id="L105"></a>}
        <a id="L106"></a>copyBytes(buf, 0, b.buf[b.off:b.off+m]);
        <a id="L107"></a>b.buf = buf;
        <a id="L108"></a>b.off = 0;
    <a id="L109"></a>}

    <a id="L111"></a>b.buf = b.buf[0 : b.off+m+n];
    <a id="L112"></a>copyString(b.buf, b.off+m, s);
    <a id="L113"></a>return n, nil;
<a id="L114"></a>}

<a id="L116"></a><span class="comment">// WriteByte appends the byte c to the buffer.</span>
<a id="L117"></a><span class="comment">// The returned error is always nil, but is included</span>
<a id="L118"></a><span class="comment">// to match bufio.Writer&#39;s WriteByte.</span>
<a id="L119"></a>func (b *Buffer) WriteByte(c byte) os.Error {
    <a id="L120"></a>if b.oneByte == nil {
        <a id="L121"></a><span class="comment">// Only happens once per Buffer, and then we have a slice.</span>
        <a id="L122"></a>b.oneByte = make([]byte, 1)
    <a id="L123"></a>}
    <a id="L124"></a>b.oneByte[0] = c;
    <a id="L125"></a>b.Write(b.oneByte);
    <a id="L126"></a>return nil;
<a id="L127"></a>}

<a id="L129"></a><span class="comment">// Read reads the next len(p) bytes from the buffer or until the buffer</span>
<a id="L130"></a><span class="comment">// is drained.  The return value n is the number of bytes read.  If the</span>
<a id="L131"></a><span class="comment">// buffer has no data to return, err is os.EOF even if len(p) is zero;</span>
<a id="L132"></a><span class="comment">// otherwise it is nil.</span>
<a id="L133"></a>func (b *Buffer) Read(p []byte) (n int, err os.Error) {
    <a id="L134"></a>if b.off &gt;= len(b.buf) {
        <a id="L135"></a>return 0, os.EOF
    <a id="L136"></a>}
    <a id="L137"></a>m := b.Len();
    <a id="L138"></a>n = len(p);

    <a id="L140"></a>if n &gt; m {
        <a id="L141"></a><span class="comment">// more bytes requested than available</span>
        <a id="L142"></a>n = m
    <a id="L143"></a>}

    <a id="L145"></a>copyBytes(p, 0, b.buf[b.off:b.off+n]);
    <a id="L146"></a>b.off += n;
    <a id="L147"></a>return n, err;
<a id="L148"></a>}

<a id="L150"></a><span class="comment">// ReadByte reads and returns the next byte from the buffer.</span>
<a id="L151"></a><span class="comment">// If no byte is available, it returns error os.EOF.</span>
<a id="L152"></a>func (b *Buffer) ReadByte() (c byte, err os.Error) {
    <a id="L153"></a>if b.off &gt;= len(b.buf) {
        <a id="L154"></a>return 0, os.EOF
    <a id="L155"></a>}
    <a id="L156"></a>c = b.buf[b.off];
    <a id="L157"></a>b.off++;
    <a id="L158"></a>return c, nil;
<a id="L159"></a>}

<a id="L161"></a><span class="comment">// NewBuffer creates and initializes a new Buffer</span>
<a id="L162"></a><span class="comment">// using buf as its initial contents.</span>
<a id="L163"></a>func NewBuffer(buf []byte) *Buffer { return &amp;Buffer{buf: buf} }

<a id="L165"></a><span class="comment">// NewBufferString creates and initializes a new Buffer</span>
<a id="L166"></a><span class="comment">// using string s as its initial contents.</span>
<a id="L167"></a>func NewBufferString(s string) *Buffer {
    <a id="L168"></a>buf := make([]byte, len(s));
    <a id="L169"></a>copyString(buf, 0, s);
    <a id="L170"></a>return &amp;Buffer{buf: buf};
<a id="L171"></a>}
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
