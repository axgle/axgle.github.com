<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/ecb_test.go</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/ecb_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package block

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;fmt&#34;;
    <a id="L10"></a>&#34;io&#34;;
    <a id="L11"></a>&#34;testing&#34;;
    <a id="L12"></a>&#34;testing/iotest&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">// Simple Cipher for testing: adds an incrementing amount</span>
<a id="L16"></a><span class="comment">// to each byte in each</span>
<a id="L17"></a>type IncCipher struct {
    <a id="L18"></a>blockSize  int;
    <a id="L19"></a>delta      byte;
    <a id="L20"></a>encrypting bool;
<a id="L21"></a>}

<a id="L23"></a>func (c *IncCipher) BlockSize() int { return c.blockSize }

<a id="L25"></a>func (c *IncCipher) Encrypt(src, dst []byte) {
    <a id="L26"></a>if !c.encrypting {
        <a id="L27"></a>panicln(&#34;encrypt: not encrypting&#34;)
    <a id="L28"></a>}
    <a id="L29"></a>if len(src) != c.blockSize || len(dst) != c.blockSize {
        <a id="L30"></a>panicln(&#34;encrypt: wrong block size&#34;, c.blockSize, len(src), len(dst))
    <a id="L31"></a>}
    <a id="L32"></a>c.delta++;
    <a id="L33"></a>for i, b := range src {
        <a id="L34"></a>dst[i] = b + c.delta
    <a id="L35"></a>}
<a id="L36"></a>}

<a id="L38"></a>func (c *IncCipher) Decrypt(src, dst []byte) {
    <a id="L39"></a>if c.encrypting {
        <a id="L40"></a>panicln(&#34;decrypt: not decrypting&#34;)
    <a id="L41"></a>}
    <a id="L42"></a>if len(src) != c.blockSize || len(dst) != c.blockSize {
        <a id="L43"></a>panicln(&#34;decrypt: wrong block size&#34;, c.blockSize, len(src), len(dst))
    <a id="L44"></a>}
    <a id="L45"></a>c.delta--;
    <a id="L46"></a>for i, b := range src {
        <a id="L47"></a>dst[i] = b + c.delta
    <a id="L48"></a>}
<a id="L49"></a>}

<a id="L51"></a>func TestECBEncrypter(t *testing.T) {
    <a id="L52"></a>var plain, crypt [256]byte;
    <a id="L53"></a>for i := 0; i &lt; len(plain); i++ {
        <a id="L54"></a>plain[i] = byte(i)
    <a id="L55"></a>}
    <a id="L56"></a>b := new(bytes.Buffer);
    <a id="L57"></a>for block := 1; block &lt;= 64; block *= 2 {
        <a id="L58"></a><span class="comment">// compute encrypted version</span>
        <a id="L59"></a>delta := byte(0);
        <a id="L60"></a>for i := 0; i &lt; len(crypt); i++ {
            <a id="L61"></a>if i%block == 0 {
                <a id="L62"></a>delta++
            <a id="L63"></a>}
            <a id="L64"></a>crypt[i] = plain[i] + delta;
        <a id="L65"></a>}

        <a id="L67"></a>for frag := 0; frag &lt; 2; frag++ {
            <a id="L68"></a>c := &amp;IncCipher{block, 0, true};
            <a id="L69"></a>b.Reset();
            <a id="L70"></a>r := bytes.NewBuffer(&amp;plain);
            <a id="L71"></a>w := NewECBEncrypter(c, b);

            <a id="L73"></a><span class="comment">// copy plain into w in increasingly large chunks: 1, 1, 2, 4, 8, ...</span>
            <a id="L74"></a><span class="comment">// if frag != 0, move the 1 to the end to cause fragmentation.</span>
            <a id="L75"></a>if frag == 0 {
                <a id="L76"></a>_, err := io.Copyn(w, r, 1);
                <a id="L77"></a>if err != nil {
                    <a id="L78"></a>t.Errorf(&#34;block=%d frag=0: first Copyn: %s&#34;, block, err);
                    <a id="L79"></a>continue;
                <a id="L80"></a>}
            <a id="L81"></a>}
            <a id="L82"></a>for n := 1; n &lt;= len(plain)/2; n *= 2 {
                <a id="L83"></a>_, err := io.Copyn(w, r, int64(n));
                <a id="L84"></a>if err != nil {
                    <a id="L85"></a>t.Errorf(&#34;block=%d frag=%d: Copyn %d: %s&#34;, block, frag, n, err)
                <a id="L86"></a>}
            <a id="L87"></a>}
            <a id="L88"></a>if frag != 0 {
                <a id="L89"></a>_, err := io.Copyn(w, r, 1);
                <a id="L90"></a>if err != nil {
                    <a id="L91"></a>t.Errorf(&#34;block=%d frag=1: last Copyn: %s&#34;, block, err);
                    <a id="L92"></a>continue;
                <a id="L93"></a>}
            <a id="L94"></a>}

            <a id="L96"></a><span class="comment">// check output</span>
            <a id="L97"></a>data := b.Bytes();
            <a id="L98"></a>if len(data) != len(crypt) {
                <a id="L99"></a>t.Errorf(&#34;block=%d frag=%d: want %d bytes, got %d&#34;, block, frag, len(crypt), len(data));
                <a id="L100"></a>continue;
            <a id="L101"></a>}

            <a id="L103"></a>if string(data) != string(&amp;crypt) {
                <a id="L104"></a>t.Errorf(&#34;block=%d frag=%d: want %x got %x&#34;, block, frag, data, crypt)
            <a id="L105"></a>}
        <a id="L106"></a>}
    <a id="L107"></a>}
<a id="L108"></a>}

<a id="L110"></a>func testECBDecrypter(t *testing.T, maxio int) {
    <a id="L111"></a>var readers = []func(io.Reader) io.Reader{
        <a id="L112"></a>func(r io.Reader) io.Reader { return r },
        <a id="L113"></a>iotest.OneByteReader,
        <a id="L114"></a>iotest.HalfReader,
    <a id="L115"></a>};
    <a id="L116"></a>var plain, crypt [256]byte;
    <a id="L117"></a>for i := 0; i &lt; len(plain); i++ {
        <a id="L118"></a>plain[i] = byte(255 - i)
    <a id="L119"></a>}
    <a id="L120"></a>b := new(bytes.Buffer);
    <a id="L121"></a>for block := 1; block &lt;= 64 &amp;&amp; block &lt;= maxio; block *= 2 {
        <a id="L122"></a><span class="comment">// compute encrypted version</span>
        <a id="L123"></a>delta := byte(0);
        <a id="L124"></a>for i := 0; i &lt; len(crypt); i++ {
            <a id="L125"></a>if i%block == 0 {
                <a id="L126"></a>delta++
            <a id="L127"></a>}
            <a id="L128"></a>crypt[i] = plain[i] + delta;
        <a id="L129"></a>}

        <a id="L131"></a>for mode := 0; mode &lt; len(readers); mode++ {
            <a id="L132"></a>for frag := 0; frag &lt; 2; frag++ {
                <a id="L133"></a>test := fmt.Sprintf(&#34;block=%d mode=%d frag=%d maxio=%d&#34;, block, mode, frag, maxio);
                <a id="L134"></a>c := &amp;IncCipher{block, 0, false};
                <a id="L135"></a>b.Reset();
                <a id="L136"></a>r := NewECBDecrypter(c, readers[mode](bytes.NewBuffer(crypt[0:maxio])));

                <a id="L138"></a><span class="comment">// read from crypt in increasingly large chunks: 1, 1, 2, 4, 8, ...</span>
                <a id="L139"></a><span class="comment">// if frag == 1, move the 1 to the end to cause fragmentation.</span>
                <a id="L140"></a>if frag == 0 {
                    <a id="L141"></a>_, err := io.Copyn(b, r, 1);
                    <a id="L142"></a>if err != nil {
                        <a id="L143"></a>t.Errorf(&#34;%s: first Copyn: %s&#34;, test, err);
                        <a id="L144"></a>continue;
                    <a id="L145"></a>}
                <a id="L146"></a>}
                <a id="L147"></a>for n := 1; n &lt;= maxio/2; n *= 2 {
                    <a id="L148"></a>_, err := io.Copyn(b, r, int64(n));
                    <a id="L149"></a>if err != nil {
                        <a id="L150"></a>t.Errorf(&#34;%s: Copyn %d: %s&#34;, test, n, err)
                    <a id="L151"></a>}
                <a id="L152"></a>}
                <a id="L153"></a>if frag != 0 {
                    <a id="L154"></a>_, err := io.Copyn(b, r, 1);
                    <a id="L155"></a>if err != nil {
                        <a id="L156"></a>t.Errorf(&#34;%s: last Copyn: %s&#34;, test, err);
                        <a id="L157"></a>continue;
                    <a id="L158"></a>}
                <a id="L159"></a>}

                <a id="L161"></a><span class="comment">// check output</span>
                <a id="L162"></a>data := b.Bytes();
                <a id="L163"></a>if len(data) != maxio {
                    <a id="L164"></a>t.Errorf(&#34;%s: want %d bytes, got %d&#34;, test, maxio, len(data));
                    <a id="L165"></a>continue;
                <a id="L166"></a>}

                <a id="L168"></a>if string(data) != string(plain[0:maxio]) {
                    <a id="L169"></a>t.Errorf(&#34;%s: input=%x want %x got %x&#34;, test, crypt[0:maxio], plain[0:maxio], data)
                <a id="L170"></a>}
            <a id="L171"></a>}
        <a id="L172"></a>}
    <a id="L173"></a>}
<a id="L174"></a>}

<a id="L176"></a>func TestECBDecrypter(t *testing.T) {
    <a id="L177"></a><span class="comment">// Do shorter I/O sizes first; they&#39;re easier to debug.</span>
    <a id="L178"></a>for n := 1; n &lt;= 256 &amp;&amp; !t.Failed(); n *= 2 {
        <a id="L179"></a>testECBDecrypter(t, n)
    <a id="L180"></a>}
<a id="L181"></a>}
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
