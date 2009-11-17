<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/encoding/base64/base64.go</title>

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
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/encoding/base64/base64.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Package base64 implements base64 encoding as specified by RFC 4648.</span>
<a id="L6"></a>package base64

<a id="L8"></a>import (
    <a id="L9"></a>&#34;bytes&#34;;
    <a id="L10"></a>&#34;io&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;strconv&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">/*</span>
<a id="L16"></a><span class="comment"> * Encodings</span>
<a id="L17"></a><span class="comment"> */</span>

<a id="L19"></a><span class="comment">// An Encoding is a radix 64 encoding/decoding scheme, defined by a</span>
<a id="L20"></a><span class="comment">// 64-character alphabet.  The most common encoding is the &#34;base64&#34;</span>
<a id="L21"></a><span class="comment">// encoding defined in RFC 4648 and used in MIME (RFC 2045) and PEM</span>
<a id="L22"></a><span class="comment">// (RFC 1421).  RFC 4648 also defines an alternate encoding, which is</span>
<a id="L23"></a><span class="comment">// the standard encoding with - and _ substituted for + and /.</span>
<a id="L24"></a>type Encoding struct {
    <a id="L25"></a>encode    string;
    <a id="L26"></a>decodeMap [256]byte;
<a id="L27"></a>}

<a id="L29"></a>const encodeStd = &#34;ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/&#34;
<a id="L30"></a>const encodeURL = &#34;ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_&#34;

<a id="L32"></a><span class="comment">// NewEncoding returns a new Encoding defined by the given alphabet,</span>
<a id="L33"></a><span class="comment">// which must be a 64-byte string.</span>
<a id="L34"></a>func NewEncoding(encoder string) *Encoding {
    <a id="L35"></a>e := new(Encoding);
    <a id="L36"></a>e.encode = encoder;
    <a id="L37"></a>for i := 0; i &lt; len(e.decodeMap); i++ {
        <a id="L38"></a>e.decodeMap[i] = 0xFF
    <a id="L39"></a>}
    <a id="L40"></a>for i := 0; i &lt; len(encoder); i++ {
        <a id="L41"></a>e.decodeMap[encoder[i]] = byte(i)
    <a id="L42"></a>}
    <a id="L43"></a>return e;
<a id="L44"></a>}

<a id="L46"></a><span class="comment">// StdEncoding is the standard base64 encoding, as defined in</span>
<a id="L47"></a><span class="comment">// RFC 4648.</span>
<a id="L48"></a>var StdEncoding = NewEncoding(encodeStd)

<a id="L50"></a><span class="comment">// URLEncoding is the alternate base64 encoding defined in RFC 4648.</span>
<a id="L51"></a><span class="comment">// It is typically used in URLs and file names.</span>
<a id="L52"></a>var URLEncoding = NewEncoding(encodeURL)

<a id="L54"></a><span class="comment">/*</span>
<a id="L55"></a><span class="comment"> * Encoder</span>
<a id="L56"></a><span class="comment"> */</span>

<a id="L58"></a><span class="comment">// Encode encodes src using the encoding enc, writing</span>
<a id="L59"></a><span class="comment">// EncodedLen(len(src)) bytes to dst.</span>
<a id="L60"></a><span class="comment">//</span>
<a id="L61"></a><span class="comment">// The encoding pads the output to a multiple of 4 bytes,</span>
<a id="L62"></a><span class="comment">// so Encode is not appropriate for use on individual blocks</span>
<a id="L63"></a><span class="comment">// of a large data stream.  Use NewEncoder() instead.</span>
<a id="L64"></a>func (enc *Encoding) Encode(dst, src []byte) {
    <a id="L65"></a>if len(src) == 0 {
        <a id="L66"></a>return
    <a id="L67"></a>}

    <a id="L69"></a>for len(src) &gt; 0 {
        <a id="L70"></a>dst[0] = 0;
        <a id="L71"></a>dst[1] = 0;
        <a id="L72"></a>dst[2] = 0;
        <a id="L73"></a>dst[3] = 0;

        <a id="L75"></a><span class="comment">// Unpack 4x 6-bit source blocks into a 4 byte</span>
        <a id="L76"></a><span class="comment">// destination quantum</span>
        <a id="L77"></a>switch len(src) {
        <a id="L78"></a>default:
            <a id="L79"></a>dst[3] |= src[2] &amp; 0x3F;
            <a id="L80"></a>dst[2] |= src[2] &gt;&gt; 6;
            <a id="L81"></a>fallthrough;
        <a id="L82"></a>case 2:
            <a id="L83"></a>dst[2] |= (src[1] &lt;&lt; 2) &amp; 0x3F;
            <a id="L84"></a>dst[1] |= src[1] &gt;&gt; 4;
            <a id="L85"></a>fallthrough;
        <a id="L86"></a>case 1:
            <a id="L87"></a>dst[1] |= (src[0] &lt;&lt; 4) &amp; 0x3F;
            <a id="L88"></a>dst[0] |= src[0] &gt;&gt; 2;
        <a id="L89"></a>}

        <a id="L91"></a><span class="comment">// Encode 6-bit blocks using the base64 alphabet</span>
        <a id="L92"></a>for j := 0; j &lt; 4; j++ {
            <a id="L93"></a>dst[j] = enc.encode[dst[j]]
        <a id="L94"></a>}

        <a id="L96"></a><span class="comment">// Pad the final quantum</span>
        <a id="L97"></a>if len(src) &lt; 3 {
            <a id="L98"></a>dst[3] = &#39;=&#39;;
            <a id="L99"></a>if len(src) &lt; 2 {
                <a id="L100"></a>dst[2] = &#39;=&#39;
            <a id="L101"></a>}
            <a id="L102"></a>break;
        <a id="L103"></a>}

        <a id="L105"></a>src = src[3:len(src)];
        <a id="L106"></a>dst = dst[4:len(dst)];
    <a id="L107"></a>}
<a id="L108"></a>}

<a id="L110"></a>type encoder struct {
    <a id="L111"></a>err  os.Error;
    <a id="L112"></a>enc  *Encoding;
    <a id="L113"></a>w    io.Writer;
    <a id="L114"></a>buf  [3]byte;    <span class="comment">// buffered data waiting to be encoded</span>
    <a id="L115"></a>nbuf int;        <span class="comment">// number of bytes in buf</span>
    <a id="L116"></a>out  [1024]byte; <span class="comment">// output buffer</span>
<a id="L117"></a>}

<a id="L119"></a>func (e *encoder) Write(p []byte) (n int, err os.Error) {
    <a id="L120"></a>if e.err != nil {
        <a id="L121"></a>return 0, e.err
    <a id="L122"></a>}

    <a id="L124"></a><span class="comment">// Leading fringe.</span>
    <a id="L125"></a>if e.nbuf &gt; 0 {
        <a id="L126"></a>var i int;
        <a id="L127"></a>for i = 0; i &lt; len(p) &amp;&amp; e.nbuf &lt; 3; i++ {
            <a id="L128"></a>e.buf[e.nbuf] = p[i];
            <a id="L129"></a>e.nbuf++;
        <a id="L130"></a>}
        <a id="L131"></a>n += i;
        <a id="L132"></a>p = p[i:len(p)];
        <a id="L133"></a>if e.nbuf &lt; 3 {
            <a id="L134"></a>return
        <a id="L135"></a>}
        <a id="L136"></a>e.enc.Encode(&amp;e.out, &amp;e.buf);
        <a id="L137"></a>if _, e.err = e.w.Write(e.out[0:4]); e.err != nil {
            <a id="L138"></a>return n, e.err
        <a id="L139"></a>}
        <a id="L140"></a>e.nbuf = 0;
    <a id="L141"></a>}

    <a id="L143"></a><span class="comment">// Large interior chunks.</span>
    <a id="L144"></a>for len(p) &gt;= 3 {
        <a id="L145"></a>nn := len(e.out) / 4 * 3;
        <a id="L146"></a>if nn &gt; len(p) {
            <a id="L147"></a>nn = len(p)
        <a id="L148"></a>}
        <a id="L149"></a>nn -= nn % 3;
        <a id="L150"></a>if nn &gt; 0 {
            <a id="L151"></a>e.enc.Encode(&amp;e.out, p[0:nn]);
            <a id="L152"></a>if _, e.err = e.w.Write(e.out[0 : nn/3*4]); e.err != nil {
                <a id="L153"></a>return n, e.err
            <a id="L154"></a>}
        <a id="L155"></a>}
        <a id="L156"></a>n += nn;
        <a id="L157"></a>p = p[nn:len(p)];
    <a id="L158"></a>}

    <a id="L160"></a><span class="comment">// Trailing fringe.</span>
    <a id="L161"></a>for i := 0; i &lt; len(p); i++ {
        <a id="L162"></a>e.buf[i] = p[i]
    <a id="L163"></a>}
    <a id="L164"></a>e.nbuf = len(p);
    <a id="L165"></a>n += len(p);
    <a id="L166"></a>return;
<a id="L167"></a>}

<a id="L169"></a><span class="comment">// Close flushes any pending output from the encoder.</span>
<a id="L170"></a><span class="comment">// It is an error to call Write after calling Close.</span>
<a id="L171"></a>func (e *encoder) Close() os.Error {
    <a id="L172"></a><span class="comment">// If there&#39;s anything left in the buffer, flush it out</span>
    <a id="L173"></a>if e.err == nil &amp;&amp; e.nbuf &gt; 0 {
        <a id="L174"></a>e.enc.Encode(&amp;e.out, e.buf[0:e.nbuf]);
        <a id="L175"></a>e.nbuf = 0;
        <a id="L176"></a>_, e.err = e.w.Write(e.out[0:4]);
    <a id="L177"></a>}
    <a id="L178"></a>return e.err;
<a id="L179"></a>}

<a id="L181"></a><span class="comment">// NewEncoder returns a new base64 stream encoder.  Data written to</span>
<a id="L182"></a><span class="comment">// the returned writer will be encoded using enc and then written to w.</span>
<a id="L183"></a><span class="comment">// Base64 encodings operate in 4-byte blocks; when finished</span>
<a id="L184"></a><span class="comment">// writing, the caller must Close the returned encoder to flush any</span>
<a id="L185"></a><span class="comment">// partially written blocks.</span>
<a id="L186"></a>func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser {
    <a id="L187"></a>return &amp;encoder{enc: enc, w: w}
<a id="L188"></a>}

<a id="L190"></a><span class="comment">// EncodedLen returns the length in bytes of the base64 encoding</span>
<a id="L191"></a><span class="comment">// of an input buffer of length n.</span>
<a id="L192"></a>func (enc *Encoding) EncodedLen(n int) int { return (n + 2) / 3 * 4 }

<a id="L194"></a><span class="comment">/*</span>
<a id="L195"></a><span class="comment"> * Decoder</span>
<a id="L196"></a><span class="comment"> */</span>

<a id="L198"></a>type CorruptInputError int64

<a id="L200"></a>func (e CorruptInputError) String() string {
    <a id="L201"></a>return &#34;illegal base64 data at input byte&#34; + strconv.Itoa64(int64(e))
<a id="L202"></a>}

<a id="L204"></a><span class="comment">// decode is like Decode but returns an additional &#39;end&#39; value, which</span>
<a id="L205"></a><span class="comment">// indicates if end-of-message padding was encountered and thus any</span>
<a id="L206"></a><span class="comment">// additional data is an error.  decode also assumes len(src)%4==0,</span>
<a id="L207"></a><span class="comment">// since it is meant for internal use.</span>
<a id="L208"></a>func (enc *Encoding) decode(dst, src []byte) (n int, end bool, err os.Error) {
    <a id="L209"></a>for i := 0; i &lt; len(src)/4 &amp;&amp; !end; i++ {
        <a id="L210"></a><span class="comment">// Decode quantum using the base64 alphabet</span>
        <a id="L211"></a>var dbuf [4]byte;
        <a id="L212"></a>dlen := 4;

    <a id="L214"></a>dbufloop:
        <a id="L215"></a>for j := 0; j &lt; 4; j++ {
            <a id="L216"></a>in := src[i*4+j];
            <a id="L217"></a>if in == &#39;=&#39; &amp;&amp; j &gt;= 2 &amp;&amp; i == len(src)/4-1 {
                <a id="L218"></a><span class="comment">// We&#39;ve reached the end and there&#39;s</span>
                <a id="L219"></a><span class="comment">// padding</span>
                <a id="L220"></a>if src[i*4+3] != &#39;=&#39; {
                    <a id="L221"></a>return n, false, CorruptInputError(i*4 + 2)
                <a id="L222"></a>}
                <a id="L223"></a>dlen = j;
                <a id="L224"></a>end = true;
                <a id="L225"></a>break dbufloop;
            <a id="L226"></a>}
            <a id="L227"></a>dbuf[j] = enc.decodeMap[in];
            <a id="L228"></a>if dbuf[j] == 0xFF {
                <a id="L229"></a>return n, false, CorruptInputError(i*4 + j)
            <a id="L230"></a>}
        <a id="L231"></a>}

        <a id="L233"></a><span class="comment">// Pack 4x 6-bit source blocks into 3 byte destination</span>
        <a id="L234"></a><span class="comment">// quantum</span>
        <a id="L235"></a>switch dlen {
        <a id="L236"></a>case 4:
            <a id="L237"></a>dst[i*3+2] = dbuf[2]&lt;&lt;6 | dbuf[3];
            <a id="L238"></a>fallthrough;
        <a id="L239"></a>case 3:
            <a id="L240"></a>dst[i*3+1] = dbuf[1]&lt;&lt;4 | dbuf[2]&gt;&gt;2;
            <a id="L241"></a>fallthrough;
        <a id="L242"></a>case 2:
            <a id="L243"></a>dst[i*3+0] = dbuf[0]&lt;&lt;2 | dbuf[1]&gt;&gt;4
        <a id="L244"></a>}
        <a id="L245"></a>n += dlen - 1;
    <a id="L246"></a>}

    <a id="L248"></a>return n, end, nil;
<a id="L249"></a>}

<a id="L251"></a><span class="comment">// Decode decodes src using the encoding enc.  It writes at most</span>
<a id="L252"></a><span class="comment">// DecodedLen(len(src)) bytes to dst and returns the number of bytes</span>
<a id="L253"></a><span class="comment">// written.  If src contains invalid base64 data, it will return the</span>
<a id="L254"></a><span class="comment">// number of bytes successfully written and CorruptInputError.</span>
<a id="L255"></a>func (enc *Encoding) Decode(dst, src []byte) (n int, err os.Error) {
    <a id="L256"></a>if len(src)%4 != 0 {
        <a id="L257"></a>return 0, CorruptInputError(len(src) / 4 * 4)
    <a id="L258"></a>}

    <a id="L260"></a>n, _, err = enc.decode(dst, src);
    <a id="L261"></a>return;
<a id="L262"></a>}

<a id="L264"></a>type decoder struct {
    <a id="L265"></a>err    os.Error;
    <a id="L266"></a>enc    *Encoding;
    <a id="L267"></a>r      io.Reader;
    <a id="L268"></a>end    bool;       <span class="comment">// saw end of message</span>
    <a id="L269"></a>buf    [1024]byte; <span class="comment">// leftover input</span>
    <a id="L270"></a>nbuf   int;
    <a id="L271"></a>out    []byte; <span class="comment">// leftover decoded output</span>
    <a id="L272"></a>outbuf [1024 / 4 * 3]byte;
<a id="L273"></a>}

<a id="L275"></a>func (d *decoder) Read(p []byte) (n int, err os.Error) {
    <a id="L276"></a>if d.err != nil {
        <a id="L277"></a>return 0, d.err
    <a id="L278"></a>}

    <a id="L280"></a><span class="comment">// Use leftover decoded output from last read.</span>
    <a id="L281"></a>if len(d.out) &gt; 0 {
        <a id="L282"></a>n = bytes.Copy(p, d.out);
        <a id="L283"></a>d.out = d.out[n:len(d.out)];
        <a id="L284"></a>return n, nil;
    <a id="L285"></a>}

    <a id="L287"></a><span class="comment">// Read a chunk.</span>
    <a id="L288"></a>nn := len(p) / 3 * 4;
    <a id="L289"></a>if nn &lt; 4 {
        <a id="L290"></a>nn = 4
    <a id="L291"></a>}
    <a id="L292"></a>if nn &gt; len(d.buf) {
        <a id="L293"></a>nn = len(d.buf)
    <a id="L294"></a>}
    <a id="L295"></a>nn, d.err = io.ReadAtLeast(d.r, d.buf[d.nbuf:nn], 4-d.nbuf);
    <a id="L296"></a>d.nbuf += nn;
    <a id="L297"></a>if d.nbuf &lt; 4 {
        <a id="L298"></a>return 0, d.err
    <a id="L299"></a>}

    <a id="L301"></a><span class="comment">// Decode chunk into p, or d.out and then p if p is too small.</span>
    <a id="L302"></a>nr := d.nbuf / 4 * 4;
    <a id="L303"></a>nw := d.nbuf / 4 * 3;
    <a id="L304"></a>if nw &gt; len(p) {
        <a id="L305"></a>nw, d.end, d.err = d.enc.decode(&amp;d.outbuf, d.buf[0:nr]);
        <a id="L306"></a>d.out = d.outbuf[0:nw];
        <a id="L307"></a>n = bytes.Copy(p, d.out);
        <a id="L308"></a>d.out = d.out[n:len(d.out)];
    <a id="L309"></a>} else {
        <a id="L310"></a>n, d.end, d.err = d.enc.decode(p, d.buf[0:nr])
    <a id="L311"></a>}
    <a id="L312"></a>d.nbuf -= nr;
    <a id="L313"></a>for i := 0; i &lt; d.nbuf; i++ {
        <a id="L314"></a>d.buf[i] = d.buf[i+nr]
    <a id="L315"></a>}

    <a id="L317"></a>if d.err == nil {
        <a id="L318"></a>d.err = err
    <a id="L319"></a>}
    <a id="L320"></a>return n, d.err;
<a id="L321"></a>}

<a id="L323"></a><span class="comment">// NewDecoder constructs a new base64 stream decoder.</span>
<a id="L324"></a>func NewDecoder(enc *Encoding, r io.Reader) io.Reader {
    <a id="L325"></a>return &amp;decoder{enc: enc, r: r}
<a id="L326"></a>}

<a id="L328"></a><span class="comment">// DecodeLen returns the maximum length in bytes of the decoded data</span>
<a id="L329"></a><span class="comment">// corresponding to n bytes of base64-encoded data.</span>
<a id="L330"></a>func (enc *Encoding) DecodedLen(n int) int { return n / 4 * 3 }
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
