<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/encoding/git85/git.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/encoding/git85/git.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Package git85 implements the radix 85 data encoding</span>
<a id="L6"></a><span class="comment">// used in the Git version control system.</span>
<a id="L7"></a>package git85

<a id="L9"></a>import (
    <a id="L10"></a>&#34;bytes&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;strconv&#34;;
<a id="L14"></a>)

<a id="L16"></a>type CorruptInputError int64

<a id="L18"></a>func (e CorruptInputError) String() string {
    <a id="L19"></a>return &#34;illegal git85 data at input byte&#34; + strconv.Itoa64(int64(e))
<a id="L20"></a>}

<a id="L22"></a>const encode = &#34;0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&amp;()*+-;&lt;=&gt;?@^_`{|}~&#34;

<a id="L24"></a><span class="comment">// The decodings are 1+ the actual value, so that the</span>
<a id="L25"></a><span class="comment">// default zero value can be used to mean &#34;not valid&#34;.</span>
<a id="L26"></a>var decode = [256]uint8{
    <a id="L27"></a>&#39;0&#39;: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
    <a id="L28"></a>&#39;A&#39;: 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
    <a id="L29"></a>24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36,
    <a id="L30"></a>&#39;a&#39;: 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
    <a id="L31"></a>50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62,
    <a id="L32"></a>&#39;!&#39;: 63,
    <a id="L33"></a>&#39;#&#39;: 64, 65, 66, 67,
    <a id="L34"></a>&#39;(&#39;: 68, 69, 70, 71,
    <a id="L35"></a>&#39;-&#39;: 72,
    <a id="L36"></a>&#39;;&#39;: 73,
    <a id="L37"></a>&#39;&lt;&#39;: 74, 75, 76, 77,
    <a id="L38"></a>&#39;@&#39;: 78,
    <a id="L39"></a>&#39;^&#39;: 79, 80, 81,
    <a id="L40"></a>&#39;{&#39;: 82, 83, 84, 85,
<a id="L41"></a>}

<a id="L43"></a><span class="comment">// Encode encodes src into EncodedLen(len(src))</span>
<a id="L44"></a><span class="comment">// bytes of dst.  As a convenience, it returns the number</span>
<a id="L45"></a><span class="comment">// of bytes written to dst, but this value is always EncodedLen(len(src)).</span>
<a id="L46"></a><span class="comment">// Encode implements the radix 85 encoding used in the</span>
<a id="L47"></a><span class="comment">// Git version control tool.</span>
<a id="L48"></a><span class="comment">//</span>
<a id="L49"></a><span class="comment">// The encoding splits src into chunks of at most 52 bytes</span>
<a id="L50"></a><span class="comment">// and encodes each chunk on its own line.</span>
<a id="L51"></a>func Encode(dst, src []byte) int {
    <a id="L52"></a>ndst := 0;
    <a id="L53"></a>for len(src) &gt; 0 {
        <a id="L54"></a>n := len(src);
        <a id="L55"></a>if n &gt; 52 {
            <a id="L56"></a>n = 52
        <a id="L57"></a>}
        <a id="L58"></a>if n &lt;= 27 {
            <a id="L59"></a>dst[ndst] = byte(&#39;A&#39; + n - 1)
        <a id="L60"></a>} else {
            <a id="L61"></a>dst[ndst] = byte(&#39;a&#39; + n - 26 - 1)
        <a id="L62"></a>}
        <a id="L63"></a>ndst++;
        <a id="L64"></a>for i := 0; i &lt; n; i += 4 {
            <a id="L65"></a>var v uint32;
            <a id="L66"></a>for j := 0; j &lt; 4 &amp;&amp; i+j &lt; n; j++ {
                <a id="L67"></a>v |= uint32(src[i+j]) &lt;&lt; uint(24-j*8)
            <a id="L68"></a>}
            <a id="L69"></a>for j := 4; j &gt;= 0; j-- {
                <a id="L70"></a>dst[ndst+j] = encode[v%85];
                <a id="L71"></a>v /= 85;
            <a id="L72"></a>}
            <a id="L73"></a>ndst += 5;
        <a id="L74"></a>}
        <a id="L75"></a>dst[ndst] = &#39;\n&#39;;
        <a id="L76"></a>ndst++;
        <a id="L77"></a>src = src[n:len(src)];
    <a id="L78"></a>}
    <a id="L79"></a>return ndst;
<a id="L80"></a>}

<a id="L82"></a><span class="comment">// EncodedLen returns the length of an encoding of n source bytes.</span>
<a id="L83"></a>func EncodedLen(n int) int {
    <a id="L84"></a>if n == 0 {
        <a id="L85"></a>return 0
    <a id="L86"></a>}
    <a id="L87"></a><span class="comment">// 5 bytes per 4 bytes of input, rounded up.</span>
    <a id="L88"></a><span class="comment">// 2 extra bytes for each line of 52 src bytes, rounded up.</span>
    <a id="L89"></a>return (n+3)/4*5 + (n+51)/52*2;
<a id="L90"></a>}

<a id="L92"></a>var newline = []byte{&#39;\n&#39;}

<a id="L94"></a><span class="comment">// Decode decodes src into at most MaxDecodedLen(len(src))</span>
<a id="L95"></a><span class="comment">// bytes, returning the actual number of bytes written to dst.</span>
<a id="L96"></a><span class="comment">//</span>
<a id="L97"></a><span class="comment">// If Decode encounters invalid input, it returns a CorruptInputError.</span>
<a id="L98"></a><span class="comment">//</span>
<a id="L99"></a>func Decode(dst, src []byte) (n int, err os.Error) {
    <a id="L100"></a>ndst := 0;
    <a id="L101"></a>nsrc := 0;
    <a id="L102"></a>for nsrc &lt; len(src) {
        <a id="L103"></a>var l int;
        <a id="L104"></a>switch ch := int(src[nsrc]); {
        <a id="L105"></a>case &#39;A&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;Z&#39;:
            <a id="L106"></a>l = ch - &#39;A&#39; + 1
        <a id="L107"></a>case &#39;a&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;z&#39;:
            <a id="L108"></a>l = ch - &#39;a&#39; + 26 + 1
        <a id="L109"></a>default:
            <a id="L110"></a>return ndst, CorruptInputError(nsrc)
        <a id="L111"></a>}
        <a id="L112"></a>if nsrc+1+l &gt; len(src) {
            <a id="L113"></a>return ndst, CorruptInputError(nsrc)
        <a id="L114"></a>}
        <a id="L115"></a>el := (l + 3) / 4 * 5; <span class="comment">// encoded len</span>
        <a id="L116"></a>if nsrc+1+el+1 &gt; len(src) || src[nsrc+1+el] != &#39;\n&#39; {
            <a id="L117"></a>return ndst, CorruptInputError(nsrc)
        <a id="L118"></a>}
        <a id="L119"></a>line := src[nsrc+1 : nsrc+1+el];
        <a id="L120"></a>for i := 0; i &lt; el; i += 5 {
            <a id="L121"></a>var v uint32;
            <a id="L122"></a>for j := 0; j &lt; 5; j++ {
                <a id="L123"></a>ch := decode[line[i+j]];
                <a id="L124"></a>if ch == 0 {
                    <a id="L125"></a>return ndst, CorruptInputError(nsrc + 1 + i + j)
                <a id="L126"></a>}
                <a id="L127"></a>v = v*85 + uint32(ch-1);
            <a id="L128"></a>}
            <a id="L129"></a>for j := 0; j &lt; 4; j++ {
                <a id="L130"></a>dst[ndst] = byte(v &gt;&gt; 24);
                <a id="L131"></a>v &lt;&lt;= 8;
                <a id="L132"></a>ndst++;
            <a id="L133"></a>}
        <a id="L134"></a>}
        <a id="L135"></a><span class="comment">// Last fragment may have run too far (but there was room in dst).</span>
        <a id="L136"></a><span class="comment">// Back up.</span>
        <a id="L137"></a>if l%4 != 0 {
            <a id="L138"></a>ndst -= 4 - l%4
        <a id="L139"></a>}
        <a id="L140"></a>nsrc += 1 + el + 1;
    <a id="L141"></a>}
    <a id="L142"></a>return ndst, nil;
<a id="L143"></a>}

<a id="L145"></a>func MaxDecodedLen(n int) int { return n / 5 * 4 }

<a id="L147"></a><span class="comment">// NewEncoder returns a new Git base85 stream encoder.  Data written to</span>
<a id="L148"></a><span class="comment">// the returned writer will be encoded and then written to w.</span>
<a id="L149"></a><span class="comment">// The Git encoding operates on 52-byte blocks; when finished</span>
<a id="L150"></a><span class="comment">// writing, the caller must Close the returned encoder to flush any</span>
<a id="L151"></a><span class="comment">// partially written blocks.</span>
<a id="L152"></a>func NewEncoder(w io.Writer) io.WriteCloser { return &amp;encoder{w: w} }

<a id="L154"></a>type encoder struct {
    <a id="L155"></a>w    io.Writer;
    <a id="L156"></a>err  os.Error;
    <a id="L157"></a>buf  [52]byte;
    <a id="L158"></a>nbuf int;
    <a id="L159"></a>out  [1024]byte;
    <a id="L160"></a>nout int;
<a id="L161"></a>}

<a id="L163"></a>func (e *encoder) Write(p []byte) (n int, err os.Error) {
    <a id="L164"></a>if e.err != nil {
        <a id="L165"></a>return 0, e.err
    <a id="L166"></a>}

    <a id="L168"></a><span class="comment">// Leading fringe.</span>
    <a id="L169"></a>if e.nbuf &gt; 0 {
        <a id="L170"></a>var i int;
        <a id="L171"></a>for i = 0; i &lt; len(p) &amp;&amp; e.nbuf &lt; 52; i++ {
            <a id="L172"></a>e.buf[e.nbuf] = p[i];
            <a id="L173"></a>e.nbuf++;
        <a id="L174"></a>}
        <a id="L175"></a>n += i;
        <a id="L176"></a>p = p[i:len(p)];
        <a id="L177"></a>if e.nbuf &lt; 52 {
            <a id="L178"></a>return
        <a id="L179"></a>}
        <a id="L180"></a>nout := Encode(&amp;e.out, &amp;e.buf);
        <a id="L181"></a>if _, e.err = e.w.Write(e.out[0:nout]); e.err != nil {
            <a id="L182"></a>return n, e.err
        <a id="L183"></a>}
        <a id="L184"></a>e.nbuf = 0;
    <a id="L185"></a>}

    <a id="L187"></a><span class="comment">// Large interior chunks.</span>
    <a id="L188"></a>for len(p) &gt;= 52 {
        <a id="L189"></a>nn := len(e.out) / (1 + 52/4*5 + 1) * 52;
        <a id="L190"></a>if nn &gt; len(p) {
            <a id="L191"></a>nn = len(p) / 52 * 52
        <a id="L192"></a>}
        <a id="L193"></a>if nn &gt; 0 {
            <a id="L194"></a>nout := Encode(&amp;e.out, p[0:nn]);
            <a id="L195"></a>if _, e.err = e.w.Write(e.out[0:nout]); e.err != nil {
                <a id="L196"></a>return n, e.err
            <a id="L197"></a>}
        <a id="L198"></a>}
        <a id="L199"></a>n += nn;
        <a id="L200"></a>p = p[nn:len(p)];
    <a id="L201"></a>}

    <a id="L203"></a><span class="comment">// Trailing fringe.</span>
    <a id="L204"></a>for i := 0; i &lt; len(p); i++ {
        <a id="L205"></a>e.buf[i] = p[i]
    <a id="L206"></a>}
    <a id="L207"></a>e.nbuf = len(p);
    <a id="L208"></a>n += len(p);
    <a id="L209"></a>return;
<a id="L210"></a>}

<a id="L212"></a>func (e *encoder) Close() os.Error {
    <a id="L213"></a><span class="comment">// If there&#39;s anything left in the buffer, flush it out</span>
    <a id="L214"></a>if e.err == nil &amp;&amp; e.nbuf &gt; 0 {
        <a id="L215"></a>nout := Encode(&amp;e.out, e.buf[0:e.nbuf]);
        <a id="L216"></a>e.nbuf = 0;
        <a id="L217"></a>_, e.err = e.w.Write(e.out[0:nout]);
    <a id="L218"></a>}
    <a id="L219"></a>return e.err;
<a id="L220"></a>}

<a id="L222"></a><span class="comment">// NewDecoder returns a new Git base85 stream decoder.</span>
<a id="L223"></a>func NewDecoder(r io.Reader) io.Reader { return &amp;decoder{r: r} }

<a id="L225"></a>type decoder struct {
    <a id="L226"></a>r       io.Reader;
    <a id="L227"></a>err     os.Error;
    <a id="L228"></a>readErr os.Error;
    <a id="L229"></a>buf     [1024]byte;
    <a id="L230"></a>nbuf    int;
    <a id="L231"></a>out     []byte;
    <a id="L232"></a>outbuf  [1024]byte;
    <a id="L233"></a>off     int64;
<a id="L234"></a>}

<a id="L236"></a>func (d *decoder) Read(p []byte) (n int, err os.Error) {
    <a id="L237"></a>if len(p) == 0 {
        <a id="L238"></a>return 0, nil
    <a id="L239"></a>}

    <a id="L241"></a>for {
        <a id="L242"></a><span class="comment">// Copy leftover output from last decode.</span>
        <a id="L243"></a>if len(d.out) &gt; 0 {
            <a id="L244"></a>n = bytes.Copy(p, d.out);
            <a id="L245"></a>d.out = d.out[n:len(d.out)];
            <a id="L246"></a>return;
        <a id="L247"></a>}

        <a id="L249"></a><span class="comment">// Out of decoded output.  Check errors.</span>
        <a id="L250"></a>if d.err != nil {
            <a id="L251"></a>return 0, d.err
        <a id="L252"></a>}
        <a id="L253"></a>if d.readErr != nil {
            <a id="L254"></a>d.err = d.readErr;
            <a id="L255"></a>return 0, d.err;
        <a id="L256"></a>}

        <a id="L258"></a><span class="comment">// Read and decode more input.</span>
        <a id="L259"></a>var nn int;
        <a id="L260"></a>nn, d.readErr = d.r.Read(d.buf[d.nbuf:len(d.buf)]);
        <a id="L261"></a>d.nbuf += nn;

        <a id="L263"></a><span class="comment">// Send complete lines to Decode.</span>
        <a id="L264"></a>nl := bytes.LastIndex(d.buf[0:d.nbuf], newline);
        <a id="L265"></a>if nl &lt; 0 {
            <a id="L266"></a>continue
        <a id="L267"></a>}
        <a id="L268"></a>nn, d.err = Decode(&amp;d.outbuf, d.buf[0:nl+1]);
        <a id="L269"></a>if e, ok := d.err.(CorruptInputError); ok {
            <a id="L270"></a>d.err = CorruptInputError(int64(e) + d.off)
        <a id="L271"></a>}
        <a id="L272"></a>d.out = d.outbuf[0:nn];
        <a id="L273"></a>d.nbuf = bytes.Copy(&amp;d.buf, d.buf[nl+1:d.nbuf]);
        <a id="L274"></a>d.off += int64(nl + 1);
    <a id="L275"></a>}
    <a id="L276"></a>panic(&#34;unreacahable&#34;);
<a id="L277"></a>}
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
