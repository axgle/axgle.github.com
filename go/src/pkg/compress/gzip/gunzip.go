<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/compress/gzip/gunzip.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/compress/gzip/gunzip.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The gzip package implements reading (and eventually writing) of</span>
<a id="L6"></a><span class="comment">// gzip format compressed files, as specified in RFC 1952.</span>
<a id="L7"></a>package gzip

<a id="L9"></a>import (
    <a id="L10"></a>&#34;bufio&#34;;
    <a id="L11"></a>&#34;compress/flate&#34;;
    <a id="L12"></a>&#34;hash&#34;;
    <a id="L13"></a>&#34;hash/crc32&#34;;
    <a id="L14"></a>&#34;io&#34;;
    <a id="L15"></a>&#34;os&#34;;
<a id="L16"></a>)

<a id="L18"></a>const (
    <a id="L19"></a>gzipID1     = 0x1f;
    <a id="L20"></a>gzipID2     = 0x8b;
    <a id="L21"></a>gzipDeflate = 8;
    <a id="L22"></a>flagText    = 1 &lt;&lt; 0;
    <a id="L23"></a>flagHdrCrc  = 1 &lt;&lt; 1;
    <a id="L24"></a>flagExtra   = 1 &lt;&lt; 2;
    <a id="L25"></a>flagName    = 1 &lt;&lt; 3;
    <a id="L26"></a>flagComment = 1 &lt;&lt; 4;
<a id="L27"></a>)

<a id="L29"></a>func makeReader(r io.Reader) flate.Reader {
    <a id="L30"></a>if rr, ok := r.(flate.Reader); ok {
        <a id="L31"></a>return rr
    <a id="L32"></a>}
    <a id="L33"></a>return bufio.NewReader(r);
<a id="L34"></a>}

<a id="L36"></a>var HeaderError os.Error = os.ErrorString(&#34;invalid gzip header&#34;)
<a id="L37"></a>var ChecksumError os.Error = os.ErrorString(&#34;gzip checksum error&#34;)

<a id="L39"></a><span class="comment">// An Inflater is an io.Reader that can be read to retrieve</span>
<a id="L40"></a><span class="comment">// uncompressed data from a gzip-format compressed file.</span>
<a id="L41"></a><span class="comment">// The gzip file stores a header giving metadata about the compressed file.</span>
<a id="L42"></a><span class="comment">// That header is exposed as the fields of the Inflater struct.</span>
<a id="L43"></a><span class="comment">//</span>
<a id="L44"></a><span class="comment">// In general, a gzip file can be a concatenation of gzip files,</span>
<a id="L45"></a><span class="comment">// each with its own header.  Reads from the Inflater</span>
<a id="L46"></a><span class="comment">// return the concatenation of the uncompressed data of each.</span>
<a id="L47"></a><span class="comment">// Only the first header is recorded in the Inflater fields.</span>
<a id="L48"></a><span class="comment">//</span>
<a id="L49"></a><span class="comment">// Gzip files store a length and checksum of the uncompressed data.</span>
<a id="L50"></a><span class="comment">// The Inflater will return a ChecksumError when Read</span>
<a id="L51"></a><span class="comment">// reaches the end of the uncompressed data if it does not</span>
<a id="L52"></a><span class="comment">// have the expected length or checksum.  Clients should treat data</span>
<a id="L53"></a><span class="comment">// returned by Read as tentative until they receive the successful</span>
<a id="L54"></a><span class="comment">// (zero length, nil error) Read marking the end of the data.</span>
<a id="L55"></a>type Inflater struct {
    <a id="L56"></a>Comment string; <span class="comment">// comment</span>
    <a id="L57"></a>Extra   []byte; <span class="comment">// &#34;extra data&#34;</span>
    <a id="L58"></a>Mtime   uint32; <span class="comment">// modification time (seconds since January 1, 1970)</span>
    <a id="L59"></a>Name    string; <span class="comment">// file name</span>
    <a id="L60"></a>OS      byte;   <span class="comment">// operating system type</span>

    <a id="L62"></a>r        flate.Reader;
    <a id="L63"></a>inflater io.ReadCloser;
    <a id="L64"></a>digest   hash.Hash32;
    <a id="L65"></a>size     uint32;
    <a id="L66"></a>flg      byte;
    <a id="L67"></a>buf      [512]byte;
    <a id="L68"></a>err      os.Error;
    <a id="L69"></a>eof      bool;
<a id="L70"></a>}

<a id="L72"></a><span class="comment">// NewInflater creates a new Inflater reading the given reader.</span>
<a id="L73"></a><span class="comment">// The implementation buffers input and may read more data than necessary from r.</span>
<a id="L74"></a><span class="comment">// It is the caller&#39;s responsibility to call Close on the Inflater when done.</span>
<a id="L75"></a>func NewInflater(r io.Reader) (*Inflater, os.Error) {
    <a id="L76"></a>z := new(Inflater);
    <a id="L77"></a>z.r = makeReader(r);
    <a id="L78"></a>z.digest = crc32.NewIEEE();
    <a id="L79"></a>if err := z.readHeader(true); err != nil {
        <a id="L80"></a>z.err = err;
        <a id="L81"></a>return nil, err;
    <a id="L82"></a>}
    <a id="L83"></a>return z, nil;
<a id="L84"></a>}

<a id="L86"></a><span class="comment">// GZIP (RFC 1952) is little-endian, unlike ZLIB (RFC 1950).</span>
<a id="L87"></a>func get4(p []byte) uint32 {
    <a id="L88"></a>return uint32(p[0]) | uint32(p[1])&lt;&lt;8 | uint32(p[2])&lt;&lt;16 | uint32(p[3])&lt;&lt;24
<a id="L89"></a>}

<a id="L91"></a>func (z *Inflater) readString() (string, os.Error) {
    <a id="L92"></a>var err os.Error;
    <a id="L93"></a>for i := 0; ; i++ {
        <a id="L94"></a>if i &gt;= len(z.buf) {
            <a id="L95"></a>return &#34;&#34;, HeaderError
        <a id="L96"></a>}
        <a id="L97"></a>z.buf[i], err = z.r.ReadByte();
        <a id="L98"></a>if err != nil {
            <a id="L99"></a>return &#34;&#34;, err
        <a id="L100"></a>}
        <a id="L101"></a>if z.buf[i] == 0 {
            <a id="L102"></a>return string(z.buf[0:i]), nil
        <a id="L103"></a>}
    <a id="L104"></a>}
    <a id="L105"></a>panic(&#34;not reached&#34;);
<a id="L106"></a>}

<a id="L108"></a>func (z *Inflater) read2() (uint32, os.Error) {
    <a id="L109"></a>_, err := z.r.Read(z.buf[0:2]);
    <a id="L110"></a>if err != nil {
        <a id="L111"></a>return 0, err
    <a id="L112"></a>}
    <a id="L113"></a>return uint32(z.buf[0]) | uint32(z.buf[1])&lt;&lt;8, nil;
<a id="L114"></a>}

<a id="L116"></a>func (z *Inflater) readHeader(save bool) os.Error {
    <a id="L117"></a>_, err := io.ReadFull(z.r, z.buf[0:10]);
    <a id="L118"></a>if err != nil {
        <a id="L119"></a>return err
    <a id="L120"></a>}
    <a id="L121"></a>if z.buf[0] != gzipID1 || z.buf[1] != gzipID2 || z.buf[2] != gzipDeflate {
        <a id="L122"></a>return HeaderError
    <a id="L123"></a>}
    <a id="L124"></a>z.flg = z.buf[3];
    <a id="L125"></a>if save {
        <a id="L126"></a>z.Mtime = get4(z.buf[4:8]);
        <a id="L127"></a><span class="comment">// z.buf[8] is xfl, ignored</span>
        <a id="L128"></a>z.OS = z.buf[9];
    <a id="L129"></a>}
    <a id="L130"></a>z.digest.Reset();
    <a id="L131"></a>z.digest.Write(z.buf[0:10]);

    <a id="L133"></a>if z.flg&amp;flagExtra != 0 {
        <a id="L134"></a>n, err := z.read2();
        <a id="L135"></a>if err != nil {
            <a id="L136"></a>return err
        <a id="L137"></a>}
        <a id="L138"></a>data := make([]byte, n);
        <a id="L139"></a>if _, err = io.ReadFull(z.r, data); err != nil {
            <a id="L140"></a>return err
        <a id="L141"></a>}
        <a id="L142"></a>if save {
            <a id="L143"></a>z.Extra = data
        <a id="L144"></a>}
    <a id="L145"></a>}

    <a id="L147"></a>var s string;
    <a id="L148"></a>if z.flg&amp;flagName != 0 {
        <a id="L149"></a>if s, err = z.readString(); err != nil {
            <a id="L150"></a>return err
        <a id="L151"></a>}
        <a id="L152"></a>if save {
            <a id="L153"></a>z.Name = s
        <a id="L154"></a>}
    <a id="L155"></a>}

    <a id="L157"></a>if z.flg&amp;flagComment != 0 {
        <a id="L158"></a>if s, err = z.readString(); err != nil {
            <a id="L159"></a>return err
        <a id="L160"></a>}
        <a id="L161"></a>if save {
            <a id="L162"></a>z.Comment = s
        <a id="L163"></a>}
    <a id="L164"></a>}

    <a id="L166"></a>if z.flg&amp;flagHdrCrc != 0 {
        <a id="L167"></a>n, err := z.read2();
        <a id="L168"></a>if err != nil {
            <a id="L169"></a>return err
        <a id="L170"></a>}
        <a id="L171"></a>sum := z.digest.Sum32() &amp; 0xFFFF;
        <a id="L172"></a>if n != sum {
            <a id="L173"></a>return HeaderError
        <a id="L174"></a>}
    <a id="L175"></a>}

    <a id="L177"></a>z.digest.Reset();
    <a id="L178"></a>z.inflater = flate.NewInflater(z.r);
    <a id="L179"></a>return nil;
<a id="L180"></a>}

<a id="L182"></a>func (z *Inflater) Read(p []byte) (n int, err os.Error) {
    <a id="L183"></a>if z.err != nil {
        <a id="L184"></a>return 0, z.err
    <a id="L185"></a>}
    <a id="L186"></a>if z.eof || len(p) == 0 {
        <a id="L187"></a>return 0, nil
    <a id="L188"></a>}

    <a id="L190"></a>n, err = z.inflater.Read(p);
    <a id="L191"></a>z.digest.Write(p[0:n]);
    <a id="L192"></a>z.size += uint32(n);
    <a id="L193"></a>if n != 0 || err != os.EOF {
        <a id="L194"></a>z.err = err;
        <a id="L195"></a>return;
    <a id="L196"></a>}

    <a id="L198"></a><span class="comment">// Finished file; check checksum + size.</span>
    <a id="L199"></a>if _, err := io.ReadFull(z.r, z.buf[0:8]); err != nil {
        <a id="L200"></a>z.err = err;
        <a id="L201"></a>return 0, err;
    <a id="L202"></a>}
    <a id="L203"></a>crc32, isize := get4(z.buf[0:4]), get4(z.buf[4:8]);
    <a id="L204"></a>sum := z.digest.Sum32();
    <a id="L205"></a>if sum != crc32 || isize != z.size {
        <a id="L206"></a>z.err = ChecksumError;
        <a id="L207"></a>return 0, z.err;
    <a id="L208"></a>}

    <a id="L210"></a><span class="comment">// File is ok; is there another?</span>
    <a id="L211"></a>if err = z.readHeader(false); err != nil {
        <a id="L212"></a>z.err = err;
        <a id="L213"></a>return;
    <a id="L214"></a>}

    <a id="L216"></a><span class="comment">// Yes.  Reset and read from it.</span>
    <a id="L217"></a>z.digest.Reset();
    <a id="L218"></a>z.size = 0;
    <a id="L219"></a>return z.Read(p);
<a id="L220"></a>}

<a id="L222"></a><span class="comment">// Calling Close does not close the wrapped io.Reader originally passed to NewInflater.</span>
<a id="L223"></a>func (z *Inflater) Close() os.Error { return z.inflater.Close() }
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
