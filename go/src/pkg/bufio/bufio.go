<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/bufio/bufio.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/bufio/bufio.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package implements buffered I/O.  It wraps an io.Reader or io.Writer</span>
<a id="L6"></a><span class="comment">// object, creating another object (Reader or Writer) that also implements</span>
<a id="L7"></a><span class="comment">// the interface but provides buffering and some help for textual I/O.</span>
<a id="L8"></a>package bufio

<a id="L10"></a>import (
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;strconv&#34;;
    <a id="L14"></a>&#34;utf8&#34;;
<a id="L15"></a>)


<a id="L18"></a>const (
    <a id="L19"></a>defaultBufSize = 4096;
<a id="L20"></a>)

<a id="L22"></a><span class="comment">// Errors introduced by this package.</span>
<a id="L23"></a>type Error struct {
    <a id="L24"></a>os.ErrorString;
<a id="L25"></a>}

<a id="L27"></a>var (
    <a id="L28"></a>ErrInvalidUnreadByte os.Error = &amp;Error{&#34;bufio: invalid use of UnreadByte&#34;};
    <a id="L29"></a>ErrBufferFull        os.Error = &amp;Error{&#34;bufio: buffer full&#34;};
    <a id="L30"></a>errInternal          os.Error = &amp;Error{&#34;bufio: internal error&#34;};
<a id="L31"></a>)

<a id="L33"></a><span class="comment">// BufSizeError is the error representing an invalid buffer size.</span>
<a id="L34"></a>type BufSizeError int

<a id="L36"></a>func (b BufSizeError) String() string {
    <a id="L37"></a>return &#34;bufio: bad buffer size &#34; + strconv.Itoa(int(b))
<a id="L38"></a>}

<a id="L40"></a>func copySlice(dst []byte, src []byte) {
    <a id="L41"></a>for i := 0; i &lt; len(dst); i++ {
        <a id="L42"></a>dst[i] = src[i]
    <a id="L43"></a>}
<a id="L44"></a>}


<a id="L47"></a><span class="comment">// Buffered input.</span>

<a id="L49"></a><span class="comment">// Reader implements buffering for an io.Reader object.</span>
<a id="L50"></a>type Reader struct {
    <a id="L51"></a>buf      []byte;
    <a id="L52"></a>rd       io.Reader;
    <a id="L53"></a>r, w     int;
    <a id="L54"></a>err      os.Error;
    <a id="L55"></a>lastbyte int;
<a id="L56"></a>}

<a id="L58"></a><span class="comment">// NewReaderSize creates a new Reader whose buffer has the specified size,</span>
<a id="L59"></a><span class="comment">// which must be greater than zero.  If the argument io.Reader is already a</span>
<a id="L60"></a><span class="comment">// Reader with large enough size, it returns the underlying Reader.</span>
<a id="L61"></a><span class="comment">// It returns the Reader and any error.</span>
<a id="L62"></a>func NewReaderSize(rd io.Reader, size int) (*Reader, os.Error) {
    <a id="L63"></a>if size &lt;= 0 {
        <a id="L64"></a>return nil, BufSizeError(size)
    <a id="L65"></a>}
    <a id="L66"></a><span class="comment">// Is it already a Reader?</span>
    <a id="L67"></a>b, ok := rd.(*Reader);
    <a id="L68"></a>if ok &amp;&amp; len(b.buf) &gt;= size {
        <a id="L69"></a>return b, nil
    <a id="L70"></a>}
    <a id="L71"></a>b = new(Reader);
    <a id="L72"></a>b.buf = make([]byte, size);
    <a id="L73"></a>b.rd = rd;
    <a id="L74"></a>b.lastbyte = -1;
    <a id="L75"></a>return b, nil;
<a id="L76"></a>}

<a id="L78"></a><span class="comment">// NewReader returns a new Reader whose buffer has the default size.</span>
<a id="L79"></a>func NewReader(rd io.Reader) *Reader {
    <a id="L80"></a>b, err := NewReaderSize(rd, defaultBufSize);
    <a id="L81"></a>if err != nil {
        <a id="L82"></a><span class="comment">// cannot happen - defaultBufSize is a valid size</span>
        <a id="L83"></a>panic(&#34;bufio: NewReader: &#34;, err.String())
    <a id="L84"></a>}
    <a id="L85"></a>return b;
<a id="L86"></a>}

<a id="L88"></a><span class="comment">// fill reads a new chunk into the buffer.</span>
<a id="L89"></a>func (b *Reader) fill() {
    <a id="L90"></a><span class="comment">// Slide existing data to beginning.</span>
    <a id="L91"></a>if b.w &gt; b.r {
        <a id="L92"></a>copySlice(b.buf[0:b.w-b.r], b.buf[b.r:b.w]);
        <a id="L93"></a>b.w -= b.r;
    <a id="L94"></a>} else {
        <a id="L95"></a>b.w = 0
    <a id="L96"></a>}
    <a id="L97"></a>b.r = 0;

    <a id="L99"></a><span class="comment">// Read new data.</span>
    <a id="L100"></a>n, e := b.rd.Read(b.buf[b.w:len(b.buf)]);
    <a id="L101"></a>b.w += n;
    <a id="L102"></a>if e != nil {
        <a id="L103"></a>b.err = e
    <a id="L104"></a>}
<a id="L105"></a>}

<a id="L107"></a><span class="comment">// Read reads data into p.</span>
<a id="L108"></a><span class="comment">// It returns the number of bytes read into p.</span>
<a id="L109"></a><span class="comment">// If nn &lt; len(p), also returns an error explaining</span>
<a id="L110"></a><span class="comment">// why the read is short.  At EOF, the count will be</span>
<a id="L111"></a><span class="comment">// zero and err will be os.EOF.</span>
<a id="L112"></a>func (b *Reader) Read(p []byte) (nn int, err os.Error) {
    <a id="L113"></a>nn = 0;
    <a id="L114"></a>for len(p) &gt; 0 {
        <a id="L115"></a>n := len(p);
        <a id="L116"></a>if b.w == b.r {
            <a id="L117"></a>if b.err != nil {
                <a id="L118"></a>return nn, b.err
            <a id="L119"></a>}
            <a id="L120"></a>if len(p) &gt;= len(b.buf) {
                <a id="L121"></a><span class="comment">// Large read, empty buffer.</span>
                <a id="L122"></a><span class="comment">// Read directly into p to avoid copy.</span>
                <a id="L123"></a>n, b.err = b.rd.Read(p);
                <a id="L124"></a>if n &gt; 0 {
                    <a id="L125"></a>b.lastbyte = int(p[n-1])
                <a id="L126"></a>}
                <a id="L127"></a>p = p[n:len(p)];
                <a id="L128"></a>nn += n;
                <a id="L129"></a>continue;
            <a id="L130"></a>}
            <a id="L131"></a>b.fill();
            <a id="L132"></a>continue;
        <a id="L133"></a>}
        <a id="L134"></a>if n &gt; b.w-b.r {
            <a id="L135"></a>n = b.w - b.r
        <a id="L136"></a>}
        <a id="L137"></a>copySlice(p[0:n], b.buf[b.r:b.r+n]);
        <a id="L138"></a>p = p[n:len(p)];
        <a id="L139"></a>b.r += n;
        <a id="L140"></a>b.lastbyte = int(b.buf[b.r-1]);
        <a id="L141"></a>nn += n;
    <a id="L142"></a>}
    <a id="L143"></a>return nn, nil;
<a id="L144"></a>}

<a id="L146"></a><span class="comment">// ReadByte reads and returns a single byte.</span>
<a id="L147"></a><span class="comment">// If no byte is available, returns an error.</span>
<a id="L148"></a>func (b *Reader) ReadByte() (c byte, err os.Error) {
    <a id="L149"></a>for b.w == b.r {
        <a id="L150"></a>if b.err != nil {
            <a id="L151"></a>return 0, b.err
        <a id="L152"></a>}
        <a id="L153"></a>b.fill();
    <a id="L154"></a>}
    <a id="L155"></a>c = b.buf[b.r];
    <a id="L156"></a>b.r++;
    <a id="L157"></a>b.lastbyte = int(c);
    <a id="L158"></a>return c, nil;
<a id="L159"></a>}

<a id="L161"></a><span class="comment">// UnreadByte unreads the last byte.  Only the most recently read byte can be unread.</span>
<a id="L162"></a>func (b *Reader) UnreadByte() os.Error {
    <a id="L163"></a>if b.r == b.w &amp;&amp; b.lastbyte &gt;= 0 {
        <a id="L164"></a>b.w = 1;
        <a id="L165"></a>b.r = 0;
        <a id="L166"></a>b.buf[0] = byte(b.lastbyte);
        <a id="L167"></a>b.lastbyte = -1;
        <a id="L168"></a>return nil;
    <a id="L169"></a>}
    <a id="L170"></a>if b.r &lt;= 0 {
        <a id="L171"></a>return ErrInvalidUnreadByte
    <a id="L172"></a>}
    <a id="L173"></a>b.r--;
    <a id="L174"></a>b.lastbyte = -1;
    <a id="L175"></a>return nil;
<a id="L176"></a>}

<a id="L178"></a><span class="comment">// ReadRune reads a single UTF-8 encoded Unicode character and returns the</span>
<a id="L179"></a><span class="comment">// rune and its size in bytes.</span>
<a id="L180"></a>func (b *Reader) ReadRune() (rune int, size int, err os.Error) {
    <a id="L181"></a>for b.r+utf8.UTFMax &gt; b.w &amp;&amp; !utf8.FullRune(b.buf[b.r:b.w]) &amp;&amp; b.err == nil {
        <a id="L182"></a>b.fill()
    <a id="L183"></a>}
    <a id="L184"></a>if b.r == b.w {
        <a id="L185"></a>return 0, 0, b.err
    <a id="L186"></a>}
    <a id="L187"></a>rune, size = int(b.buf[b.r]), 1;
    <a id="L188"></a>if rune &gt;= 0x80 {
        <a id="L189"></a>rune, size = utf8.DecodeRune(b.buf[b.r:b.w])
    <a id="L190"></a>}
    <a id="L191"></a>b.r += size;
    <a id="L192"></a>b.lastbyte = int(b.buf[b.r-1]);
    <a id="L193"></a>return rune, size, nil;
<a id="L194"></a>}

<a id="L196"></a><span class="comment">// Helper function: look for byte c in array p,</span>
<a id="L197"></a><span class="comment">// returning its index or -1.</span>
<a id="L198"></a>func findByte(p []byte, c byte) int {
    <a id="L199"></a>for i := 0; i &lt; len(p); i++ {
        <a id="L200"></a>if p[i] == c {
            <a id="L201"></a>return i
        <a id="L202"></a>}
    <a id="L203"></a>}
    <a id="L204"></a>return -1;
<a id="L205"></a>}

<a id="L207"></a><span class="comment">// Buffered returns the number of bytes that can be read from the current buffer.</span>
<a id="L208"></a>func (b *Reader) Buffered() int { return b.w - b.r }

<a id="L210"></a><span class="comment">// ReadSlice reads until the first occurrence of delim in the input,</span>
<a id="L211"></a><span class="comment">// returning a slice pointing at the bytes in the buffer.</span>
<a id="L212"></a><span class="comment">// The bytes stop being valid at the next read call.</span>
<a id="L213"></a><span class="comment">// If ReadSlice encounters an error before finding a delimiter,</span>
<a id="L214"></a><span class="comment">// it returns all the data in the buffer and the error itself (often os.EOF).</span>
<a id="L215"></a><span class="comment">// ReadSlice fails with error ErrBufferFull if the buffer fills without a delim.</span>
<a id="L216"></a><span class="comment">// Because the data returned from ReadSlice will be overwritten</span>
<a id="L217"></a><span class="comment">// by the next I/O operation, most clients should use</span>
<a id="L218"></a><span class="comment">// ReadBytes or ReadString instead.</span>
<a id="L219"></a><span class="comment">// ReadSlice returns err != nil if and only if line does not end in delim.</span>
<a id="L220"></a>func (b *Reader) ReadSlice(delim byte) (line []byte, err os.Error) {
    <a id="L221"></a><span class="comment">// Look in buffer.</span>
    <a id="L222"></a>if i := findByte(b.buf[b.r:b.w], delim); i &gt;= 0 {
        <a id="L223"></a>line1 := b.buf[b.r : b.r+i+1];
        <a id="L224"></a>b.r += i + 1;
        <a id="L225"></a>return line1, nil;
    <a id="L226"></a>}

    <a id="L228"></a><span class="comment">// Read more into buffer, until buffer fills or we find delim.</span>
    <a id="L229"></a>for {
        <a id="L230"></a>if b.err != nil {
            <a id="L231"></a>line := b.buf[b.r:b.w];
            <a id="L232"></a>b.r = b.w;
            <a id="L233"></a>return line, b.err;
        <a id="L234"></a>}

        <a id="L236"></a>n := b.Buffered();
        <a id="L237"></a>b.fill();

        <a id="L239"></a><span class="comment">// Search new part of buffer</span>
        <a id="L240"></a>if i := findByte(b.buf[n:b.w], delim); i &gt;= 0 {
            <a id="L241"></a>line := b.buf[0 : n+i+1];
            <a id="L242"></a>b.r = n + i + 1;
            <a id="L243"></a>return line, nil;
        <a id="L244"></a>}

        <a id="L246"></a><span class="comment">// Buffer is full?</span>
        <a id="L247"></a>if b.Buffered() &gt;= len(b.buf) {
            <a id="L248"></a>return nil, ErrBufferFull
        <a id="L249"></a>}
    <a id="L250"></a>}
    <a id="L251"></a>panic(&#34;not reached&#34;);
<a id="L252"></a>}

<a id="L254"></a><span class="comment">// ReadBytes reads until the first occurrence of delim in the input,</span>
<a id="L255"></a><span class="comment">// returning a string containing the data up to and including the delimiter.</span>
<a id="L256"></a><span class="comment">// If ReadBytes encounters an error before finding a delimiter,</span>
<a id="L257"></a><span class="comment">// it returns the data read before the error and the error itself (often os.EOF).</span>
<a id="L258"></a><span class="comment">// ReadBytes returns err != nil if and only if line does not end in delim.</span>
<a id="L259"></a>func (b *Reader) ReadBytes(delim byte) (line []byte, err os.Error) {
    <a id="L260"></a><span class="comment">// Use ReadSlice to look for array,</span>
    <a id="L261"></a><span class="comment">// accumulating full buffers.</span>
    <a id="L262"></a>var frag []byte;
    <a id="L263"></a>var full [][]byte;
    <a id="L264"></a>nfull := 0;
    <a id="L265"></a>err = nil;

    <a id="L267"></a>for {
        <a id="L268"></a>var e os.Error;
        <a id="L269"></a>frag, e = b.ReadSlice(delim);
        <a id="L270"></a>if e == nil { <span class="comment">// got final fragment</span>
            <a id="L271"></a>break
        <a id="L272"></a>}
        <a id="L273"></a>if e != ErrBufferFull { <span class="comment">// unexpected error</span>
            <a id="L274"></a>err = e;
            <a id="L275"></a>break;
        <a id="L276"></a>}

        <a id="L278"></a><span class="comment">// Read bytes out of buffer.</span>
        <a id="L279"></a>buf := make([]byte, b.Buffered());
        <a id="L280"></a>var n int;
        <a id="L281"></a>n, e = b.Read(buf);
        <a id="L282"></a>if e != nil {
            <a id="L283"></a>frag = buf[0:n];
            <a id="L284"></a>err = e;
            <a id="L285"></a>break;
        <a id="L286"></a>}
        <a id="L287"></a>if n != len(buf) {
            <a id="L288"></a>frag = buf[0:n];
            <a id="L289"></a>err = errInternal;
            <a id="L290"></a>break;
        <a id="L291"></a>}

        <a id="L293"></a><span class="comment">// Grow list if needed.</span>
        <a id="L294"></a>if full == nil {
            <a id="L295"></a>full = make([][]byte, 16)
        <a id="L296"></a>} else if nfull &gt;= len(full) {
            <a id="L297"></a>newfull := make([][]byte, len(full)*2);
            <a id="L298"></a>for i := 0; i &lt; len(full); i++ {
                <a id="L299"></a>newfull[i] = full[i]
            <a id="L300"></a>}
            <a id="L301"></a>full = newfull;
        <a id="L302"></a>}

        <a id="L304"></a><span class="comment">// Save buffer</span>
        <a id="L305"></a>full[nfull] = buf;
        <a id="L306"></a>nfull++;
    <a id="L307"></a>}

    <a id="L309"></a><span class="comment">// Allocate new buffer to hold the full pieces and the fragment.</span>
    <a id="L310"></a>n := 0;
    <a id="L311"></a>for i := 0; i &lt; nfull; i++ {
        <a id="L312"></a>n += len(full[i])
    <a id="L313"></a>}
    <a id="L314"></a>n += len(frag);

    <a id="L316"></a><span class="comment">// Copy full pieces and fragment in.</span>
    <a id="L317"></a>buf := make([]byte, n);
    <a id="L318"></a>n = 0;
    <a id="L319"></a>for i := 0; i &lt; nfull; i++ {
        <a id="L320"></a>copySlice(buf[n:n+len(full[i])], full[i]);
        <a id="L321"></a>n += len(full[i]);
    <a id="L322"></a>}
    <a id="L323"></a>copySlice(buf[n:n+len(frag)], frag);
    <a id="L324"></a>return buf, err;
<a id="L325"></a>}

<a id="L327"></a><span class="comment">// ReadString reads until the first occurrence of delim in the input,</span>
<a id="L328"></a><span class="comment">// returning a string containing the data up to and including the delimiter.</span>
<a id="L329"></a><span class="comment">// If ReadString encounters an error before finding a delimiter,</span>
<a id="L330"></a><span class="comment">// it returns the data read before the error and the error itself (often os.EOF).</span>
<a id="L331"></a><span class="comment">// ReadString returns err != nil if and only if line does not end in delim.</span>
<a id="L332"></a>func (b *Reader) ReadString(delim byte) (line string, err os.Error) {
    <a id="L333"></a>bytes, e := b.ReadBytes(delim);
    <a id="L334"></a>return string(bytes), e;
<a id="L335"></a>}


<a id="L338"></a><span class="comment">// buffered output</span>

<a id="L340"></a><span class="comment">// Writer implements buffering for an io.Writer object.</span>
<a id="L341"></a>type Writer struct {
    <a id="L342"></a>err os.Error;
    <a id="L343"></a>buf []byte;
    <a id="L344"></a>n   int;
    <a id="L345"></a>wr  io.Writer;
<a id="L346"></a>}

<a id="L348"></a><span class="comment">// NewWriterSize creates a new Writer whose buffer has the specified size,</span>
<a id="L349"></a><span class="comment">// which must be greater than zero. If the argument io.Writer is already a</span>
<a id="L350"></a><span class="comment">// Writer with large enough size, it returns the underlying Writer.</span>
<a id="L351"></a><span class="comment">// It returns the Writer and any error.</span>
<a id="L352"></a>func NewWriterSize(wr io.Writer, size int) (*Writer, os.Error) {
    <a id="L353"></a>if size &lt;= 0 {
        <a id="L354"></a>return nil, BufSizeError(size)
    <a id="L355"></a>}
    <a id="L356"></a><span class="comment">// Is it already a Writer?</span>
    <a id="L357"></a>b, ok := wr.(*Writer);
    <a id="L358"></a>if ok &amp;&amp; len(b.buf) &gt;= size {
        <a id="L359"></a>return b, nil
    <a id="L360"></a>}
    <a id="L361"></a>b = new(Writer);
    <a id="L362"></a>b.buf = make([]byte, size);
    <a id="L363"></a>b.wr = wr;
    <a id="L364"></a>return b, nil;
<a id="L365"></a>}

<a id="L367"></a><span class="comment">// NewWriter returns a new Writer whose buffer has the default size.</span>
<a id="L368"></a>func NewWriter(wr io.Writer) *Writer {
    <a id="L369"></a>b, err := NewWriterSize(wr, defaultBufSize);
    <a id="L370"></a>if err != nil {
        <a id="L371"></a><span class="comment">// cannot happen - defaultBufSize is valid size</span>
        <a id="L372"></a>panic(&#34;bufio: NewWriter: &#34;, err.String())
    <a id="L373"></a>}
    <a id="L374"></a>return b;
<a id="L375"></a>}

<a id="L377"></a><span class="comment">// Flush writes any buffered data to the underlying io.Writer.</span>
<a id="L378"></a>func (b *Writer) Flush() os.Error {
    <a id="L379"></a>if b.err != nil {
        <a id="L380"></a>return b.err
    <a id="L381"></a>}
    <a id="L382"></a>n, e := b.wr.Write(b.buf[0:b.n]);
    <a id="L383"></a>if n &lt; b.n &amp;&amp; e == nil {
        <a id="L384"></a>e = io.ErrShortWrite
    <a id="L385"></a>}
    <a id="L386"></a>if e != nil {
        <a id="L387"></a>if n &gt; 0 &amp;&amp; n &lt; b.n {
            <a id="L388"></a>copySlice(b.buf[0:b.n-n], b.buf[n:b.n])
        <a id="L389"></a>}
        <a id="L390"></a>b.n -= n;
        <a id="L391"></a>b.err = e;
        <a id="L392"></a>return e;
    <a id="L393"></a>}
    <a id="L394"></a>b.n = 0;
    <a id="L395"></a>return nil;
<a id="L396"></a>}

<a id="L398"></a><span class="comment">// Available returns how many bytes are unused in the buffer.</span>
<a id="L399"></a>func (b *Writer) Available() int { return len(b.buf) - b.n }

<a id="L401"></a><span class="comment">// Buffered returns the number of bytes that have been written into the current buffer.</span>
<a id="L402"></a>func (b *Writer) Buffered() int { return b.n }

<a id="L404"></a><span class="comment">// Write writes the contents of p into the buffer.</span>
<a id="L405"></a><span class="comment">// It returns the number of bytes written.</span>
<a id="L406"></a><span class="comment">// If nn &lt; len(p), also returns an error explaining</span>
<a id="L407"></a><span class="comment">// why the write is short.</span>
<a id="L408"></a>func (b *Writer) Write(p []byte) (nn int, err os.Error) {
    <a id="L409"></a>if b.err != nil {
        <a id="L410"></a>return 0, b.err
    <a id="L411"></a>}
    <a id="L412"></a>nn = 0;
    <a id="L413"></a>for len(p) &gt; 0 {
        <a id="L414"></a>n := b.Available();
        <a id="L415"></a>if n &lt;= 0 {
            <a id="L416"></a>if b.Flush(); b.err != nil {
                <a id="L417"></a>break
            <a id="L418"></a>}
            <a id="L419"></a>n = b.Available();
        <a id="L420"></a>}
        <a id="L421"></a>if b.Available() == 0 &amp;&amp; len(p) &gt;= len(b.buf) {
            <a id="L422"></a><span class="comment">// Large write, empty buffer.</span>
            <a id="L423"></a><span class="comment">// Write directly from p to avoid copy.</span>
            <a id="L424"></a>n, b.err = b.wr.Write(p);
            <a id="L425"></a>nn += n;
            <a id="L426"></a>p = p[n:len(p)];
            <a id="L427"></a>if b.err != nil {
                <a id="L428"></a>break
            <a id="L429"></a>}
            <a id="L430"></a>continue;
        <a id="L431"></a>}
        <a id="L432"></a>if n &gt; len(p) {
            <a id="L433"></a>n = len(p)
        <a id="L434"></a>}
        <a id="L435"></a>copySlice(b.buf[b.n:b.n+n], p[0:n]);
        <a id="L436"></a>b.n += n;
        <a id="L437"></a>nn += n;
        <a id="L438"></a>p = p[n:len(p)];
    <a id="L439"></a>}
    <a id="L440"></a>return nn, b.err;
<a id="L441"></a>}

<a id="L443"></a><span class="comment">// WriteByte writes a single byte.</span>
<a id="L444"></a>func (b *Writer) WriteByte(c byte) os.Error {
    <a id="L445"></a>if b.err != nil {
        <a id="L446"></a>return b.err
    <a id="L447"></a>}
    <a id="L448"></a>if b.Available() &lt;= 0 &amp;&amp; b.Flush() != nil {
        <a id="L449"></a>return b.err
    <a id="L450"></a>}
    <a id="L451"></a>b.buf[b.n] = c;
    <a id="L452"></a>b.n++;
    <a id="L453"></a>return nil;
<a id="L454"></a>}

<a id="L456"></a><span class="comment">// WriteString writes a string.</span>
<a id="L457"></a>func (b *Writer) WriteString(s string) os.Error {
    <a id="L458"></a>if b.err != nil {
        <a id="L459"></a>return b.err
    <a id="L460"></a>}
    <a id="L461"></a><span class="comment">// Common case, worth making fast.</span>
    <a id="L462"></a>if b.Available() &gt;= len(s) || len(b.buf) &gt;= len(s) &amp;&amp; b.Flush() == nil {
        <a id="L463"></a>for i := 0; i &lt; len(s); i++ { <span class="comment">// loop over bytes, not runes.</span>
            <a id="L464"></a>b.buf[b.n] = s[i];
            <a id="L465"></a>b.n++;
        <a id="L466"></a>}
        <a id="L467"></a>return nil;
    <a id="L468"></a>}
    <a id="L469"></a>for i := 0; i &lt; len(s); i++ { <span class="comment">// loop over bytes, not runes.</span>
        <a id="L470"></a>b.WriteByte(s[i])
    <a id="L471"></a>}
    <a id="L472"></a>return b.err;
<a id="L473"></a>}

<a id="L475"></a><span class="comment">// buffered input and output</span>

<a id="L477"></a><span class="comment">// ReadWriter stores pointers to a Reader and a Writer.</span>
<a id="L478"></a><span class="comment">// It implements io.ReadWriter.</span>
<a id="L479"></a>type ReadWriter struct {
    <a id="L480"></a>*Reader;
    <a id="L481"></a>*Writer;
<a id="L482"></a>}

<a id="L484"></a><span class="comment">// NewReadWriter allocates a new ReadWriter that dispatches to r and w.</span>
<a id="L485"></a>func NewReadWriter(r *Reader, w *Writer) *ReadWriter {
    <a id="L486"></a>return &amp;ReadWriter{r, w}
<a id="L487"></a>}
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
