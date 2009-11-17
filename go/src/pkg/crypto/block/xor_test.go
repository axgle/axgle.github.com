<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/xor_test.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/xor_test.go</h1>

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

<a id="L15"></a><span class="comment">// Simple &#34;pseudo-random&#34; stream for testing.</span>
<a id="L16"></a>type incStream struct {
    <a id="L17"></a>buf []byte;
    <a id="L18"></a>n   byte;
<a id="L19"></a>}

<a id="L21"></a>func newIncStream(blockSize int) *incStream {
    <a id="L22"></a>x := new(incStream);
    <a id="L23"></a>x.buf = make([]byte, blockSize);
    <a id="L24"></a>return x;
<a id="L25"></a>}

<a id="L27"></a>func (x *incStream) Next() []byte {
    <a id="L28"></a>x.n++;
    <a id="L29"></a>for i := range x.buf {
        <a id="L30"></a>x.buf[i] = x.n;
        <a id="L31"></a>x.n++;
    <a id="L32"></a>}
    <a id="L33"></a>return x.buf;
<a id="L34"></a>}

<a id="L36"></a>func testXorWriter(t *testing.T, maxio int) {
    <a id="L37"></a>var plain, crypt [256]byte;
    <a id="L38"></a>for i := 0; i &lt; len(plain); i++ {
        <a id="L39"></a>plain[i] = byte(i)
    <a id="L40"></a>}
    <a id="L41"></a>b := new(bytes.Buffer);
    <a id="L42"></a>for block := 1; block &lt;= 64 &amp;&amp; block &lt;= maxio; block *= 2 {
        <a id="L43"></a><span class="comment">// compute encrypted version</span>
        <a id="L44"></a>n := byte(0);
        <a id="L45"></a>for i := 0; i &lt; len(crypt); i++ {
            <a id="L46"></a>if i%block == 0 {
                <a id="L47"></a>n++
            <a id="L48"></a>}
            <a id="L49"></a>crypt[i] = plain[i] ^ n;
            <a id="L50"></a>n++;
        <a id="L51"></a>}

        <a id="L53"></a>for frag := 0; frag &lt; 2; frag++ {
            <a id="L54"></a>test := fmt.Sprintf(&#34;block=%d frag=%d maxio=%d&#34;, block, frag, maxio);
            <a id="L55"></a>b.Reset();
            <a id="L56"></a>r := bytes.NewBuffer(&amp;plain);
            <a id="L57"></a>s := newIncStream(block);
            <a id="L58"></a>w := newXorWriter(s, b);

            <a id="L60"></a><span class="comment">// copy plain into w in increasingly large chunks: 1, 1, 2, 4, 8, ...</span>
            <a id="L61"></a><span class="comment">// if frag != 0, move the 1 to the end to cause fragmentation.</span>
            <a id="L62"></a>if frag == 0 {
                <a id="L63"></a>_, err := io.Copyn(w, r, 1);
                <a id="L64"></a>if err != nil {
                    <a id="L65"></a>t.Errorf(&#34;%s: first Copyn: %s&#34;, test, err);
                    <a id="L66"></a>continue;
                <a id="L67"></a>}
            <a id="L68"></a>}
            <a id="L69"></a>for n := 1; n &lt;= len(plain)/2; n *= 2 {
                <a id="L70"></a>_, err := io.Copyn(w, r, int64(n));
                <a id="L71"></a>if err != nil {
                    <a id="L72"></a>t.Errorf(&#34;%s: Copyn %d: %s&#34;, test, n, err)
                <a id="L73"></a>}
            <a id="L74"></a>}

            <a id="L76"></a><span class="comment">// check output</span>
            <a id="L77"></a>crypt := crypt[0 : len(crypt)-frag];
            <a id="L78"></a>data := b.Bytes();
            <a id="L79"></a>if len(data) != len(crypt) {
                <a id="L80"></a>t.Errorf(&#34;%s: want %d bytes, got %d&#34;, test, len(crypt), len(data));
                <a id="L81"></a>continue;
            <a id="L82"></a>}

            <a id="L84"></a>if string(data) != string(crypt) {
                <a id="L85"></a>t.Errorf(&#34;%s: want %x got %x&#34;, test, data, crypt)
            <a id="L86"></a>}
        <a id="L87"></a>}
    <a id="L88"></a>}
<a id="L89"></a>}


<a id="L92"></a>func TestXorWriter(t *testing.T) {
    <a id="L93"></a><span class="comment">// Do shorter I/O sizes first; they&#39;re easier to debug.</span>
    <a id="L94"></a>for n := 1; n &lt;= 256 &amp;&amp; !t.Failed(); n *= 2 {
        <a id="L95"></a>testXorWriter(t, n)
    <a id="L96"></a>}
<a id="L97"></a>}

<a id="L99"></a>func testXorReader(t *testing.T, maxio int) {
    <a id="L100"></a>var readers = []func(io.Reader) io.Reader{
        <a id="L101"></a>func(r io.Reader) io.Reader { return r },
        <a id="L102"></a>iotest.OneByteReader,
        <a id="L103"></a>iotest.HalfReader,
    <a id="L104"></a>};
    <a id="L105"></a>var plain, crypt [256]byte;
    <a id="L106"></a>for i := 0; i &lt; len(plain); i++ {
        <a id="L107"></a>plain[i] = byte(255 - i)
    <a id="L108"></a>}
    <a id="L109"></a>b := new(bytes.Buffer);
    <a id="L110"></a>for block := 1; block &lt;= 64 &amp;&amp; block &lt;= maxio; block *= 2 {
        <a id="L111"></a><span class="comment">// compute encrypted version</span>
        <a id="L112"></a>n := byte(0);
        <a id="L113"></a>for i := 0; i &lt; len(crypt); i++ {
            <a id="L114"></a>if i%block == 0 {
                <a id="L115"></a>n++
            <a id="L116"></a>}
            <a id="L117"></a>crypt[i] = plain[i] ^ n;
            <a id="L118"></a>n++;
        <a id="L119"></a>}

        <a id="L121"></a>for mode := 0; mode &lt; len(readers); mode++ {
            <a id="L122"></a>for frag := 0; frag &lt; 2; frag++ {
                <a id="L123"></a>test := fmt.Sprintf(&#34;block=%d mode=%d frag=%d maxio=%d&#34;, block, mode, frag, maxio);
                <a id="L124"></a>s := newIncStream(block);
                <a id="L125"></a>b.Reset();
                <a id="L126"></a>r := newXorReader(s, readers[mode](bytes.NewBuffer(crypt[0:maxio])));

                <a id="L128"></a><span class="comment">// read from crypt in increasingly large chunks: 1, 1, 2, 4, 8, ...</span>
                <a id="L129"></a><span class="comment">// if frag == 1, move the 1 to the end to cause fragmentation.</span>
                <a id="L130"></a>if frag == 0 {
                    <a id="L131"></a>_, err := io.Copyn(b, r, 1);
                    <a id="L132"></a>if err != nil {
                        <a id="L133"></a>t.Errorf(&#34;%s: first Copyn: %s&#34;, test, err);
                        <a id="L134"></a>continue;
                    <a id="L135"></a>}
                <a id="L136"></a>}
                <a id="L137"></a>for n := 1; n &lt;= maxio/2; n *= 2 {
                    <a id="L138"></a>_, err := io.Copyn(b, r, int64(n));
                    <a id="L139"></a>if err != nil {
                        <a id="L140"></a>t.Errorf(&#34;%s: Copyn %d: %s&#34;, test, n, err)
                    <a id="L141"></a>}
                <a id="L142"></a>}

                <a id="L144"></a><span class="comment">// check output</span>
                <a id="L145"></a>data := b.Bytes();
                <a id="L146"></a>crypt := crypt[0 : maxio-frag];
                <a id="L147"></a>plain := plain[0 : maxio-frag];
                <a id="L148"></a>if len(data) != len(plain) {
                    <a id="L149"></a>t.Errorf(&#34;%s: want %d bytes, got %d&#34;, test, len(plain), len(data));
                    <a id="L150"></a>continue;
                <a id="L151"></a>}

                <a id="L153"></a>if string(data) != string(plain) {
                    <a id="L154"></a>t.Errorf(&#34;%s: input=%x want %x got %x&#34;, test, crypt, plain, data)
                <a id="L155"></a>}
            <a id="L156"></a>}
        <a id="L157"></a>}
    <a id="L158"></a>}
<a id="L159"></a>}

<a id="L161"></a>func TestXorReader(t *testing.T) {
    <a id="L162"></a><span class="comment">// Do shorter I/O sizes first; they&#39;re easier to debug.</span>
    <a id="L163"></a>for n := 1; n &lt;= 256 &amp;&amp; !t.Failed(); n *= 2 {
        <a id="L164"></a>testXorReader(t, n)
    <a id="L165"></a>}
<a id="L166"></a>}

<a id="L168"></a><span class="comment">// TODO(rsc): Test handling of writes after write errors.</span>
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
