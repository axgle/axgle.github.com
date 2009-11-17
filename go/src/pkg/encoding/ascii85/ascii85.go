<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/encoding/ascii85/ascii85.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/encoding/ascii85/ascii85.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Package ascii85 implements the ascii85 data encoding</span>
<a id="L6"></a><span class="comment">// as used in the btoa tool and Adobe&#39;s PostScript and PDF document formats.</span>
<a id="L7"></a>package ascii85

<a id="L9"></a>import (
    <a id="L10"></a>&#34;bytes&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;strconv&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">/*</span>
<a id="L17"></a><span class="comment"> * Encoder</span>
<a id="L18"></a><span class="comment"> */</span>

<a id="L20"></a><span class="comment">// Encode encodes src into at most MaxEncodedLen(len(src))</span>
<a id="L21"></a><span class="comment">// bytes of dst, returning the actual number of bytes written.</span>
<a id="L22"></a><span class="comment">//</span>
<a id="L23"></a><span class="comment">// The encoding handles 4-byte chunks, using a special encoding</span>
<a id="L24"></a><span class="comment">// for the last fragment, so Encode is not appropriate for use on</span>
<a id="L25"></a><span class="comment">// individual blocks of a large data stream.  Use NewEncoder() instead.</span>
<a id="L26"></a><span class="comment">//</span>
<a id="L27"></a><span class="comment">// Often, ascii85-encoded data is wrapped in &lt;~ and ~&gt; symbols.</span>
<a id="L28"></a><span class="comment">// Encode does not add these.</span>
<a id="L29"></a>func Encode(dst, src []byte) int {
    <a id="L30"></a>if len(src) == 0 {
        <a id="L31"></a>return 0
    <a id="L32"></a>}

    <a id="L34"></a>n := 0;
    <a id="L35"></a>for len(src) &gt; 0 {
        <a id="L36"></a>dst[0] = 0;
        <a id="L37"></a>dst[1] = 0;
        <a id="L38"></a>dst[2] = 0;
        <a id="L39"></a>dst[3] = 0;
        <a id="L40"></a>dst[4] = 0;

        <a id="L42"></a><span class="comment">// Unpack 4 bytes into uint32 to repack into base 85 5-byte.</span>
        <a id="L43"></a>var v uint32;
        <a id="L44"></a>switch len(src) {
        <a id="L45"></a>default:
            <a id="L46"></a>v |= uint32(src[3]);
            <a id="L47"></a>fallthrough;
        <a id="L48"></a>case 3:
            <a id="L49"></a>v |= uint32(src[2]) &lt;&lt; 8;
            <a id="L50"></a>fallthrough;
        <a id="L51"></a>case 2:
            <a id="L52"></a>v |= uint32(src[1]) &lt;&lt; 16;
            <a id="L53"></a>fallthrough;
        <a id="L54"></a>case 1:
            <a id="L55"></a>v |= uint32(src[0]) &lt;&lt; 24
        <a id="L56"></a>}

        <a id="L58"></a><span class="comment">// Special case: zero (!!!!!) shortens to z.</span>
        <a id="L59"></a>if v == 0 &amp;&amp; len(src) &gt;= 4 {
            <a id="L60"></a>dst[0] = &#39;z&#39;;
            <a id="L61"></a>dst = dst[1:len(dst)];
            <a id="L62"></a>n++;
            <a id="L63"></a>continue;
        <a id="L64"></a>}

        <a id="L66"></a><span class="comment">// Otherwise, 5 base 85 digits starting at !.</span>
        <a id="L67"></a>for i := 4; i &gt;= 0; i-- {
            <a id="L68"></a>dst[i] = &#39;!&#39; + byte(v%85);
            <a id="L69"></a>v /= 85;
        <a id="L70"></a>}

        <a id="L72"></a><span class="comment">// If src was short, discard the low destination bytes.</span>
        <a id="L73"></a>m := 5;
        <a id="L74"></a>if len(src) &lt; 4 {
            <a id="L75"></a>m -= 4 - len(src);
            <a id="L76"></a>src = nil;
        <a id="L77"></a>} else {
            <a id="L78"></a>src = src[4:len(src)]
        <a id="L79"></a>}
        <a id="L80"></a>dst = dst[m:len(dst)];
        <a id="L81"></a>n += m;
    <a id="L82"></a>}
    <a id="L83"></a>return n;
<a id="L84"></a>}

<a id="L86"></a><span class="comment">// MaxEncodedLen returns the maximum length of an encoding of n source bytes.</span>
<a id="L87"></a>func MaxEncodedLen(n int) int { return (n + 3) / 4 * 5 }

<a id="L89"></a><span class="comment">// NewEncoder returns a new ascii85 stream encoder.  Data written to</span>
<a id="L90"></a><span class="comment">// the returned writer will be encoded and then written to w.</span>
<a id="L91"></a><span class="comment">// Ascii85 encodings operate in 32-bit blocks; when finished</span>
<a id="L92"></a><span class="comment">// writing, the caller must Close the returned encoder to flush any</span>
<a id="L93"></a><span class="comment">// trailing partial block.</span>
<a id="L94"></a>func NewEncoder(w io.Writer) io.WriteCloser { return &amp;encoder{w: w} }

<a id="L96"></a>type encoder struct {
    <a id="L97"></a>err  os.Error;
    <a id="L98"></a>w    io.Writer;
    <a id="L99"></a>buf  [4]byte;    <span class="comment">// buffered data waiting to be encoded</span>
    <a id="L100"></a>nbuf int;        <span class="comment">// number of bytes in buf</span>
    <a id="L101"></a>out  [1024]byte; <span class="comment">// output buffer</span>
<a id="L102"></a>}

<a id="L104"></a>func (e *encoder) Write(p []byte) (n int, err os.Error) {
    <a id="L105"></a>if e.err != nil {
        <a id="L106"></a>return 0, e.err
    <a id="L107"></a>}

    <a id="L109"></a><span class="comment">// Leading fringe.</span>
    <a id="L110"></a>if e.nbuf &gt; 0 {
        <a id="L111"></a>var i int;
        <a id="L112"></a>for i = 0; i &lt; len(p) &amp;&amp; e.nbuf &lt; 4; i++ {
            <a id="L113"></a>e.buf[e.nbuf] = p[i];
            <a id="L114"></a>e.nbuf++;
        <a id="L115"></a>}
        <a id="L116"></a>n += i;
        <a id="L117"></a>p = p[i:len(p)];
        <a id="L118"></a>if e.nbuf &lt; 4 {
            <a id="L119"></a>return
        <a id="L120"></a>}
        <a id="L121"></a>nout := Encode(&amp;e.out, &amp;e.buf);
        <a id="L122"></a>if _, e.err = e.w.Write(e.out[0:nout]); e.err != nil {
            <a id="L123"></a>return n, e.err
        <a id="L124"></a>}
        <a id="L125"></a>e.nbuf = 0;
    <a id="L126"></a>}

    <a id="L128"></a><span class="comment">// Large interior chunks.</span>
    <a id="L129"></a>for len(p) &gt;= 4 {
        <a id="L130"></a>nn := len(e.out) / 5 * 4;
        <a id="L131"></a>if nn &gt; len(p) {
            <a id="L132"></a>nn = len(p)
        <a id="L133"></a>}
        <a id="L134"></a>nn -= nn % 4;
        <a id="L135"></a>if nn &gt; 0 {
            <a id="L136"></a>nout := Encode(&amp;e.out, p[0:nn]);
            <a id="L137"></a>if _, e.err = e.w.Write(e.out[0:nout]); e.err != nil {
                <a id="L138"></a>return n, e.err
            <a id="L139"></a>}
        <a id="L140"></a>}
        <a id="L141"></a>n += nn;
        <a id="L142"></a>p = p[nn:len(p)];
    <a id="L143"></a>}

    <a id="L145"></a><span class="comment">// Trailing fringe.</span>
    <a id="L146"></a>for i := 0; i &lt; len(p); i++ {
        <a id="L147"></a>e.buf[i] = p[i]
    <a id="L148"></a>}
    <a id="L149"></a>e.nbuf = len(p);
    <a id="L150"></a>n += len(p);
    <a id="L151"></a>return;
<a id="L152"></a>}

<a id="L154"></a><span class="comment">// Close flushes any pending output from the encoder.</span>
<a id="L155"></a><span class="comment">// It is an error to call Write after calling Close.</span>
<a id="L156"></a>func (e *encoder) Close() os.Error {
    <a id="L157"></a><span class="comment">// If there&#39;s anything left in the buffer, flush it out</span>
    <a id="L158"></a>if e.err == nil &amp;&amp; e.nbuf &gt; 0 {
        <a id="L159"></a>nout := Encode(&amp;e.out, e.buf[0:e.nbuf]);
        <a id="L160"></a>e.nbuf = 0;
        <a id="L161"></a>_, e.err = e.w.Write(e.out[0:nout]);
    <a id="L162"></a>}
    <a id="L163"></a>return e.err;
<a id="L164"></a>}

<a id="L166"></a><span class="comment">/*</span>
<a id="L167"></a><span class="comment"> * Decoder</span>
<a id="L168"></a><span class="comment"> */</span>

<a id="L170"></a>type CorruptInputError int64

<a id="L172"></a>func (e CorruptInputError) String() string {
    <a id="L173"></a>return &#34;illegal ascii85 data at input byte&#34; + strconv.Itoa64(int64(e))
<a id="L174"></a>}

<a id="L176"></a><span class="comment">// Decode decodes src into dst, returning both the number</span>
<a id="L177"></a><span class="comment">// of bytes written to dst and the number consumed from src.</span>
<a id="L178"></a><span class="comment">// If src contains invalid ascii85 data, Decode will return the</span>
<a id="L179"></a><span class="comment">// number of bytes successfully written and a CorruptInputError.</span>
<a id="L180"></a><span class="comment">// Decode ignores space and control characters in src.</span>
<a id="L181"></a><span class="comment">// Often, ascii85-encoded data is wrapped in &lt;~ and ~&gt; symbols.</span>
<a id="L182"></a><span class="comment">// Decode expects these to have been stripped by the caller.</span>
<a id="L183"></a><span class="comment">//</span>
<a id="L184"></a><span class="comment">// If flush is true, Decode assumes that src represents the</span>
<a id="L185"></a><span class="comment">// end of the input stream and processes it completely rather</span>
<a id="L186"></a><span class="comment">// than wait for the completion of another 32-bit block.</span>
<a id="L187"></a><span class="comment">//</span>
<a id="L188"></a><span class="comment">// NewDecoder wraps an io.Reader interface around Decode.</span>
<a id="L189"></a><span class="comment">//</span>
<a id="L190"></a>func Decode(dst, src []byte, flush bool) (ndst, nsrc int, err os.Error) {
    <a id="L191"></a>var v uint32;
    <a id="L192"></a>var nb int;
    <a id="L193"></a>for i, b := range src {
        <a id="L194"></a>if len(dst)-ndst &lt; 4 {
            <a id="L195"></a>return
        <a id="L196"></a>}
        <a id="L197"></a>switch {
        <a id="L198"></a>case b &lt;= &#39; &#39;:
            <a id="L199"></a>continue
        <a id="L200"></a>case b == &#39;z&#39; &amp;&amp; nb == 0:
            <a id="L201"></a>nb = 5;
            <a id="L202"></a>v = 0;
        <a id="L203"></a>case &#39;!&#39; &lt;= b &amp;&amp; b &lt;= &#39;u&#39;:
            <a id="L204"></a>v = v*85 + uint32(b-&#39;!&#39;);
            <a id="L205"></a>nb++;
        <a id="L206"></a>default:
            <a id="L207"></a>return 0, 0, CorruptInputError(i)
        <a id="L208"></a>}
        <a id="L209"></a>if nb == 5 {
            <a id="L210"></a>nsrc = i + 1;
            <a id="L211"></a>dst[ndst] = byte(v &gt;&gt; 24);
            <a id="L212"></a>dst[ndst+1] = byte(v &gt;&gt; 16);
            <a id="L213"></a>dst[ndst+2] = byte(v &gt;&gt; 8);
            <a id="L214"></a>dst[ndst+3] = byte(v);
            <a id="L215"></a>ndst += 4;
            <a id="L216"></a>nb = 0;
            <a id="L217"></a>v = 0;
        <a id="L218"></a>}
    <a id="L219"></a>}
    <a id="L220"></a>if flush {
        <a id="L221"></a>nsrc = len(src);
        <a id="L222"></a>if nb &gt; 0 {
            <a id="L223"></a><span class="comment">// The number of output bytes in the last fragment</span>
            <a id="L224"></a><span class="comment">// is the number of leftover input bytes - 1:</span>
            <a id="L225"></a><span class="comment">// the extra byte provides enough bits to cover</span>
            <a id="L226"></a><span class="comment">// the inefficiency of the encoding for the block.</span>
            <a id="L227"></a>if nb == 1 {
                <a id="L228"></a>return 0, 0, CorruptInputError(len(src))
            <a id="L229"></a>}
            <a id="L230"></a>for i := nb; i &lt; 5; i++ {
                <a id="L231"></a><span class="comment">// The short encoding truncated the output value.</span>
                <a id="L232"></a><span class="comment">// We have to assume the worst case values (digit 84)</span>
                <a id="L233"></a><span class="comment">// in order to ensure that the top bits are correct.</span>
                <a id="L234"></a>v = v*85 + 84
            <a id="L235"></a>}
            <a id="L236"></a>for i := 0; i &lt; nb-1; i++ {
                <a id="L237"></a>dst[ndst] = byte(v &gt;&gt; 24);
                <a id="L238"></a>v &lt;&lt;= 8;
                <a id="L239"></a>ndst++;
            <a id="L240"></a>}
        <a id="L241"></a>}
    <a id="L242"></a>}
    <a id="L243"></a>return;
<a id="L244"></a>}

<a id="L246"></a><span class="comment">// NewDecoder constructs a new ascii85 stream decoder.</span>
<a id="L247"></a>func NewDecoder(r io.Reader) io.Reader { return &amp;decoder{r: r} }

<a id="L249"></a>type decoder struct {
    <a id="L250"></a>err     os.Error;
    <a id="L251"></a>readErr os.Error;
    <a id="L252"></a>r       io.Reader;
    <a id="L253"></a>end     bool;       <span class="comment">// saw end of message</span>
    <a id="L254"></a>buf     [1024]byte; <span class="comment">// leftover input</span>
    <a id="L255"></a>nbuf    int;
    <a id="L256"></a>out     []byte; <span class="comment">// leftover decoded output</span>
    <a id="L257"></a>outbuf  [1024]byte;
<a id="L258"></a>}

<a id="L260"></a>func (d *decoder) Read(p []byte) (n int, err os.Error) {
    <a id="L261"></a>if len(p) == 0 {
        <a id="L262"></a>return 0, nil
    <a id="L263"></a>}
    <a id="L264"></a>if d.err != nil {
        <a id="L265"></a>return 0, d.err
    <a id="L266"></a>}

    <a id="L268"></a>for {
        <a id="L269"></a><span class="comment">// Copy leftover output from last decode.</span>
        <a id="L270"></a>if len(d.out) &gt; 0 {
            <a id="L271"></a>n = bytes.Copy(p, d.out);
            <a id="L272"></a>d.out = d.out[n:len(d.out)];
            <a id="L273"></a>return;
        <a id="L274"></a>}

        <a id="L276"></a><span class="comment">// Decode leftover input from last read.</span>
        <a id="L277"></a>var nn, nsrc, ndst int;
        <a id="L278"></a>if d.nbuf &gt; 0 {
            <a id="L279"></a>ndst, nsrc, d.err = Decode(&amp;d.outbuf, d.buf[0:d.nbuf], d.readErr != nil);
            <a id="L280"></a>if ndst &gt; 0 {
                <a id="L281"></a>d.out = d.outbuf[0:ndst];
                <a id="L282"></a>d.nbuf = bytes.Copy(&amp;d.buf, d.buf[nsrc:d.nbuf]);
                <a id="L283"></a>continue; <span class="comment">// copy out and return</span>
            <a id="L284"></a>}
        <a id="L285"></a>}

        <a id="L287"></a><span class="comment">// Out of input, out of decoded output.  Check errors.</span>
        <a id="L288"></a>if d.err != nil {
            <a id="L289"></a>return 0, d.err
        <a id="L290"></a>}
        <a id="L291"></a>if d.readErr != nil {
            <a id="L292"></a>d.err = d.readErr;
            <a id="L293"></a>return 0, d.err;
        <a id="L294"></a>}

        <a id="L296"></a><span class="comment">// Read more data.</span>
        <a id="L297"></a>nn, d.readErr = d.r.Read(d.buf[d.nbuf:len(d.buf)]);
        <a id="L298"></a>d.nbuf += nn;
    <a id="L299"></a>}
    <a id="L300"></a>panic(&#34;unreachable&#34;);
<a id="L301"></a>}
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
