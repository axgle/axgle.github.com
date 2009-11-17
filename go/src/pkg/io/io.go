<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/io/io.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/io/io.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package provides basic interfaces to I/O primitives.</span>
<a id="L6"></a><span class="comment">// Its primary job is to wrap existing implementations of such primitives,</span>
<a id="L7"></a><span class="comment">// such as those in package os, into shared public interfaces that</span>
<a id="L8"></a><span class="comment">// abstract the functionality.</span>
<a id="L9"></a><span class="comment">// It also provides buffering primitives and some other basic operations.</span>
<a id="L10"></a>package io

<a id="L12"></a>import (
    <a id="L13"></a>&#34;os&#34;;
    <a id="L14"></a>&#34;strings&#34;;
<a id="L15"></a>)

<a id="L17"></a><span class="comment">// Error represents an unexpected I/O behavior.</span>
<a id="L18"></a>type Error struct {
    <a id="L19"></a>os.ErrorString;
<a id="L20"></a>}

<a id="L22"></a><span class="comment">// ErrShortWrite means that a write accepted fewer bytes than requested</span>
<a id="L23"></a><span class="comment">// but failed to return an explicit error.</span>
<a id="L24"></a>var ErrShortWrite os.Error = &amp;Error{&#34;short write&#34;}

<a id="L26"></a><span class="comment">// ErrUnexpectedEOF means that os.EOF was encountered in the</span>
<a id="L27"></a><span class="comment">// middle of reading a fixed-size block or data structure.</span>
<a id="L28"></a>var ErrUnexpectedEOF os.Error = &amp;Error{&#34;unexpected EOF&#34;}

<a id="L30"></a><span class="comment">// Reader is the interface that wraps the basic Read method.</span>
<a id="L31"></a><span class="comment">//</span>
<a id="L32"></a><span class="comment">// Read reads up to len(p) bytes into p.  It returns the number of bytes</span>
<a id="L33"></a><span class="comment">// read (0 &lt;= n &lt;= len(p)) and any error encountered.</span>
<a id="L34"></a><span class="comment">// Even if Read returns n &lt; len(p),</span>
<a id="L35"></a><span class="comment">// it may use all of p as scratch space during the call.</span>
<a id="L36"></a><span class="comment">// If some data is available but not len(p) bytes, Read conventionally</span>
<a id="L37"></a><span class="comment">// returns what is available rather than block waiting for more.</span>
<a id="L38"></a><span class="comment">//</span>
<a id="L39"></a><span class="comment">// At the end of the input stream, Read returns 0, os.EOF.</span>
<a id="L40"></a><span class="comment">// Read may return a non-zero number of bytes with a non-nil err.</span>
<a id="L41"></a><span class="comment">// In particular, a Read that exhausts the input may return n &gt; 0, os.EOF.</span>
<a id="L42"></a>type Reader interface {
    <a id="L43"></a>Read(p []byte) (n int, err os.Error);
<a id="L44"></a>}

<a id="L46"></a><span class="comment">// Writer is the interface that wraps the basic Write method.</span>
<a id="L47"></a><span class="comment">//</span>
<a id="L48"></a><span class="comment">// Write writes len(p) bytes from p to the underlying data stream.</span>
<a id="L49"></a><span class="comment">// It returns the number of bytes written from p (0 &lt;= n &lt;= len(p))</span>
<a id="L50"></a><span class="comment">// and any error encountered that caused the write to stop early.</span>
<a id="L51"></a><span class="comment">// Write must return a non-nil error if it returns n &lt; len(p).</span>
<a id="L52"></a>type Writer interface {
    <a id="L53"></a>Write(p []byte) (n int, err os.Error);
<a id="L54"></a>}

<a id="L56"></a><span class="comment">// Closer is the interface that wraps the basic Close method.</span>
<a id="L57"></a>type Closer interface {
    <a id="L58"></a>Close() os.Error;
<a id="L59"></a>}

<a id="L61"></a><span class="comment">// Seeker is the interface that wraps the basic Seek method.</span>
<a id="L62"></a><span class="comment">//</span>
<a id="L63"></a><span class="comment">// Seek sets the offset for the next Read or Write to offset,</span>
<a id="L64"></a><span class="comment">// interpreted according to whence: 0 means relative to the origin of</span>
<a id="L65"></a><span class="comment">// the file, 1 means relative to the current offset, and 2 means</span>
<a id="L66"></a><span class="comment">// relative to the end.  Seek returns the new offset and an Error, if</span>
<a id="L67"></a><span class="comment">// any.</span>
<a id="L68"></a>type Seeker interface {
    <a id="L69"></a>Seek(offset int64, whence int) (ret int64, err os.Error);
<a id="L70"></a>}

<a id="L72"></a><span class="comment">// ReadWrite is the interface that groups the basic Read and Write methods.</span>
<a id="L73"></a>type ReadWriter interface {
    <a id="L74"></a>Reader;
    <a id="L75"></a>Writer;
<a id="L76"></a>}

<a id="L78"></a><span class="comment">// ReadCloser is the interface that groups the basic Read and Close methods.</span>
<a id="L79"></a>type ReadCloser interface {
    <a id="L80"></a>Reader;
    <a id="L81"></a>Closer;
<a id="L82"></a>}

<a id="L84"></a><span class="comment">// WriteCloser is the interface that groups the basic Write and Close methods.</span>
<a id="L85"></a>type WriteCloser interface {
    <a id="L86"></a>Writer;
    <a id="L87"></a>Closer;
<a id="L88"></a>}

<a id="L90"></a><span class="comment">// ReadWriteCloser is the interface that groups the basic Read, Write and Close methods.</span>
<a id="L91"></a>type ReadWriteCloser interface {
    <a id="L92"></a>Reader;
    <a id="L93"></a>Writer;
    <a id="L94"></a>Closer;
<a id="L95"></a>}

<a id="L97"></a><span class="comment">// ReadSeeker is the interface that groups the basic Read and Seek methods.</span>
<a id="L98"></a>type ReadSeeker interface {
    <a id="L99"></a>Reader;
    <a id="L100"></a>Seeker;
<a id="L101"></a>}

<a id="L103"></a><span class="comment">// WriteSeeker is the interface that groups the basic Write and Seek methods.</span>
<a id="L104"></a>type WriteSeeker interface {
    <a id="L105"></a>Writer;
    <a id="L106"></a>Seeker;
<a id="L107"></a>}

<a id="L109"></a><span class="comment">// ReadWriteSeeker is the interface that groups the basic Read, Write and Seek methods.</span>
<a id="L110"></a>type ReadWriteSeeker interface {
    <a id="L111"></a>Reader;
    <a id="L112"></a>Writer;
    <a id="L113"></a>Seeker;
<a id="L114"></a>}

<a id="L116"></a><span class="comment">// ReaderAt is the interface that wraps the basic ReadAt method.</span>
<a id="L117"></a><span class="comment">//</span>
<a id="L118"></a><span class="comment">// ReadAt reads len(p) bytes into p starting at offset off in the</span>
<a id="L119"></a><span class="comment">// underlying data stream.  It returns the number of bytes</span>
<a id="L120"></a><span class="comment">// read (0 &lt;= n &lt;= len(p)) and any error encountered.</span>
<a id="L121"></a><span class="comment">//</span>
<a id="L122"></a><span class="comment">// Even if ReadAt returns n &lt; len(p),</span>
<a id="L123"></a><span class="comment">// it may use all of p as scratch space during the call.</span>
<a id="L124"></a><span class="comment">// If some data is available but not len(p) bytes, ReadAt blocks</span>
<a id="L125"></a><span class="comment">// until either all the data is available or an error occurs.</span>
<a id="L126"></a><span class="comment">//</span>
<a id="L127"></a><span class="comment">// At the end of the input stream, ReadAt returns 0, os.EOF.</span>
<a id="L128"></a><span class="comment">// ReadAt may return a non-zero number of bytes with a non-nil err.</span>
<a id="L129"></a><span class="comment">// In particular, a ReadAt that exhausts the input may return n &gt; 0, os.EOF.</span>
<a id="L130"></a>type ReaderAt interface {
    <a id="L131"></a>ReadAt(p []byte, off int64) (n int, err os.Error);
<a id="L132"></a>}

<a id="L134"></a><span class="comment">// WriterAt is the interface that wraps the basic WriteAt method.</span>
<a id="L135"></a><span class="comment">//</span>
<a id="L136"></a><span class="comment">// WriteAt writes len(p) bytes from p to the underlying data stream</span>
<a id="L137"></a><span class="comment">// at offset off.  It returns the number of bytes written from p (0 &lt;= n &lt;= len(p))</span>
<a id="L138"></a><span class="comment">// and any error encountered that caused the write to stop early.</span>
<a id="L139"></a><span class="comment">// WriteAt must return a non-nil error if it returns n &lt; len(p).</span>
<a id="L140"></a>type WriterAt interface {
    <a id="L141"></a>WriteAt(p []byte, off int64) (n int, err os.Error);
<a id="L142"></a>}

<a id="L144"></a><span class="comment">// WriteString writes the contents of the string s to w, which accepts an array of bytes.</span>
<a id="L145"></a>func WriteString(w Writer, s string) (n int, err os.Error) {
    <a id="L146"></a>return w.Write(strings.Bytes(s))
<a id="L147"></a>}

<a id="L149"></a><span class="comment">// ReadAtLeast reads from r into buf until it has read at least min bytes.</span>
<a id="L150"></a><span class="comment">// It returns the number of bytes copied and an error if fewer bytes were read.</span>
<a id="L151"></a><span class="comment">// The error is os.EOF only if no bytes were read.</span>
<a id="L152"></a><span class="comment">// If an EOF happens after reading fewer than min bytes,</span>
<a id="L153"></a><span class="comment">// ReadAtLeast returns ErrUnexpectedEOF.</span>
<a id="L154"></a>func ReadAtLeast(r Reader, buf []byte, min int) (n int, err os.Error) {
    <a id="L155"></a>n = 0;
    <a id="L156"></a>for n &lt; min {
        <a id="L157"></a>nn, e := r.Read(buf[n:len(buf)]);
        <a id="L158"></a>if nn &gt; 0 {
            <a id="L159"></a>n += nn
        <a id="L160"></a>}
        <a id="L161"></a>if e != nil {
            <a id="L162"></a>if e == os.EOF &amp;&amp; n &gt; 0 {
                <a id="L163"></a>e = ErrUnexpectedEOF
            <a id="L164"></a>}
            <a id="L165"></a>return n, e;
        <a id="L166"></a>}
    <a id="L167"></a>}
    <a id="L168"></a>return n, nil;
<a id="L169"></a>}

<a id="L171"></a><span class="comment">// ReadFull reads exactly len(buf) bytes from r into buf.</span>
<a id="L172"></a><span class="comment">// It returns the number of bytes copied and an error if fewer bytes were read.</span>
<a id="L173"></a><span class="comment">// The error is os.EOF only if no bytes were read.</span>
<a id="L174"></a><span class="comment">// If an EOF happens after reading some but not all the bytes,</span>
<a id="L175"></a><span class="comment">// ReadFull returns ErrUnexpectedEOF.</span>
<a id="L176"></a>func ReadFull(r Reader, buf []byte) (n int, err os.Error) {
    <a id="L177"></a>return ReadAtLeast(r, buf, len(buf))
<a id="L178"></a>}

<a id="L180"></a><span class="comment">// Copyn copies n bytes (or until an error) from src to dst.</span>
<a id="L181"></a><span class="comment">// It returns the number of bytes copied and the error, if any.</span>
<a id="L182"></a>func Copyn(dst Writer, src Reader, n int64) (written int64, err os.Error) {
    <a id="L183"></a>buf := make([]byte, 32*1024);
    <a id="L184"></a>for written &lt; n {
        <a id="L185"></a>l := len(buf);
        <a id="L186"></a>if d := n - written; d &lt; int64(l) {
            <a id="L187"></a>l = int(d)
        <a id="L188"></a>}
        <a id="L189"></a>nr, er := src.Read(buf[0:l]);
        <a id="L190"></a>if nr &gt; 0 {
            <a id="L191"></a>nw, ew := dst.Write(buf[0:nr]);
            <a id="L192"></a>if nw &gt; 0 {
                <a id="L193"></a>written += int64(nw)
            <a id="L194"></a>}
            <a id="L195"></a>if ew != nil {
                <a id="L196"></a>err = ew;
                <a id="L197"></a>break;
            <a id="L198"></a>}
            <a id="L199"></a>if nr != nw {
                <a id="L200"></a>err = ErrShortWrite;
                <a id="L201"></a>break;
            <a id="L202"></a>}
        <a id="L203"></a>}
        <a id="L204"></a>if er != nil {
            <a id="L205"></a>err = er;
            <a id="L206"></a>break;
        <a id="L207"></a>}
    <a id="L208"></a>}
    <a id="L209"></a>return written, err;
<a id="L210"></a>}

<a id="L212"></a><span class="comment">// Copy copies from src to dst until either EOF is reached</span>
<a id="L213"></a><span class="comment">// on src or an error occurs.  It returns the number of bytes</span>
<a id="L214"></a><span class="comment">// copied and the error, if any.</span>
<a id="L215"></a>func Copy(dst Writer, src Reader) (written int64, err os.Error) {
    <a id="L216"></a>buf := make([]byte, 32*1024);
    <a id="L217"></a>for {
        <a id="L218"></a>nr, er := src.Read(buf);
        <a id="L219"></a>if nr &gt; 0 {
            <a id="L220"></a>nw, ew := dst.Write(buf[0:nr]);
            <a id="L221"></a>if nw &gt; 0 {
                <a id="L222"></a>written += int64(nw)
            <a id="L223"></a>}
            <a id="L224"></a>if ew != nil {
                <a id="L225"></a>err = ew;
                <a id="L226"></a>break;
            <a id="L227"></a>}
            <a id="L228"></a>if nr != nw {
                <a id="L229"></a>err = ErrShortWrite;
                <a id="L230"></a>break;
            <a id="L231"></a>}
        <a id="L232"></a>}
        <a id="L233"></a>if er == os.EOF {
            <a id="L234"></a>break
        <a id="L235"></a>}
        <a id="L236"></a>if er != nil {
            <a id="L237"></a>err = er;
            <a id="L238"></a>break;
        <a id="L239"></a>}
    <a id="L240"></a>}
    <a id="L241"></a>return written, err;
<a id="L242"></a>}

<a id="L244"></a><span class="comment">// LimitReader returns a Reader that reads from r</span>
<a id="L245"></a><span class="comment">// but stops with os.EOF after n bytes.</span>
<a id="L246"></a>func LimitReader(r Reader, n int64) Reader { return &amp;limitedReader{r, n} }

<a id="L248"></a>type limitedReader struct {
    <a id="L249"></a>r   Reader;
    <a id="L250"></a>n   int64;
<a id="L251"></a>}

<a id="L253"></a>func (l *limitedReader) Read(p []byte) (n int, err os.Error) {
    <a id="L254"></a>if l.n &lt;= 0 {
        <a id="L255"></a>return 0, os.EOF
    <a id="L256"></a>}
    <a id="L257"></a>if int64(len(p)) &gt; l.n {
        <a id="L258"></a>p = p[0:l.n]
    <a id="L259"></a>}
    <a id="L260"></a>n, err = l.r.Read(p);
    <a id="L261"></a>l.n -= int64(n);
    <a id="L262"></a>return;
<a id="L263"></a>}

<a id="L265"></a><span class="comment">// NewSectionReader returns a SectionReader that reads from r</span>
<a id="L266"></a><span class="comment">// starting at offset off and stops with os.EOF after n bytes.</span>
<a id="L267"></a>func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader {
    <a id="L268"></a>return &amp;SectionReader{r, off, off, off + n}
<a id="L269"></a>}

<a id="L271"></a><span class="comment">// SectionReader implements Read, Seek, and ReadAt on a section</span>
<a id="L272"></a><span class="comment">// of an underlying ReaderAt.</span>
<a id="L273"></a>type SectionReader struct {
    <a id="L274"></a>r     ReaderAt;
    <a id="L275"></a>base  int64;
    <a id="L276"></a>off   int64;
    <a id="L277"></a>limit int64;
<a id="L278"></a>}

<a id="L280"></a>func (s *SectionReader) Read(p []byte) (n int, err os.Error) {
    <a id="L281"></a>if s.off &gt;= s.limit {
        <a id="L282"></a>return 0, os.EOF
    <a id="L283"></a>}
    <a id="L284"></a>if max := s.limit - s.off; int64(len(p)) &gt; max {
        <a id="L285"></a>p = p[0:max]
    <a id="L286"></a>}
    <a id="L287"></a>n, err = s.r.ReadAt(p, s.off);
    <a id="L288"></a>s.off += int64(n);
    <a id="L289"></a>return;
<a id="L290"></a>}

<a id="L292"></a>func (s *SectionReader) Seek(offset int64, whence int) (ret int64, err os.Error) {
    <a id="L293"></a>switch whence {
    <a id="L294"></a>default:
        <a id="L295"></a>return 0, os.EINVAL
    <a id="L296"></a>case 0:
        <a id="L297"></a>offset += s.base
    <a id="L298"></a>case 1:
        <a id="L299"></a>offset += s.off
    <a id="L300"></a>case 2:
        <a id="L301"></a>offset += s.limit
    <a id="L302"></a>}
    <a id="L303"></a>if offset &lt; s.off || offset &gt; s.limit {
        <a id="L304"></a>return 0, os.EINVAL
    <a id="L305"></a>}
    <a id="L306"></a>s.off = offset;
    <a id="L307"></a>return offset - s.base, nil;
<a id="L308"></a>}

<a id="L310"></a>func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err os.Error) {
    <a id="L311"></a>if off &lt; 0 || off &gt;= s.limit-s.base {
        <a id="L312"></a>return 0, os.EOF
    <a id="L313"></a>}
    <a id="L314"></a>off += s.base;
    <a id="L315"></a>if max := s.limit - off; int64(len(p)) &gt; max {
        <a id="L316"></a>p = p[0:max]
    <a id="L317"></a>}
    <a id="L318"></a>return s.r.ReadAt(p, off);
<a id="L319"></a>}

<a id="L321"></a><span class="comment">// Size returns the size of the section in bytes.</span>
<a id="L322"></a>func (s *SectionReader) Size() int64 { return s.limit - s.base }
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
