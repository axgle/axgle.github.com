<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/ecb.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/ecb.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Electronic codebook (ECB) mode.</span>
<a id="L6"></a><span class="comment">// ECB is a fancy name for ``encrypt and decrypt each block separately.&#39;&#39;</span>
<a id="L7"></a><span class="comment">// It&#39;s a pretty bad thing to do for any large amount of data (more than one block),</span>
<a id="L8"></a><span class="comment">// because the individual blocks can still be identified, duplicated, and reordered.</span>
<a id="L9"></a><span class="comment">// The ECB implementation exists mainly to provide buffering for</span>
<a id="L10"></a><span class="comment">// the other modes, which wrap it by providing modified Ciphers.</span>

<a id="L12"></a><span class="comment">// See NIST SP 800-38A, pp 9-10</span>

<a id="L14"></a>package block

<a id="L16"></a>import (
    <a id="L17"></a>&#34;io&#34;;
    <a id="L18"></a>&#34;os&#34;;
    <a id="L19"></a>&#34;strconv&#34;;
<a id="L20"></a>)

<a id="L22"></a>type ecbDecrypter struct {
    <a id="L23"></a>c         Cipher;
    <a id="L24"></a>r         io.Reader;
    <a id="L25"></a>blockSize int; <span class="comment">// block size</span>

    <a id="L27"></a><span class="comment">// Buffered data.</span>
    <a id="L28"></a><span class="comment">// The buffer buf is used as storage for both</span>
    <a id="L29"></a><span class="comment">// plain or crypt; at least one of those is nil at any given time.</span>
    <a id="L30"></a>buf   []byte;
    <a id="L31"></a>plain []byte; <span class="comment">// plain text waiting to be read</span>
    <a id="L32"></a>crypt []byte; <span class="comment">// ciphertext waiting to be decrypted</span>
<a id="L33"></a>}

<a id="L35"></a><span class="comment">// Read into x.crypt until it has a full block or EOF or an error happens.</span>
<a id="L36"></a>func (x *ecbDecrypter) fillCrypt() os.Error {
    <a id="L37"></a>var err os.Error;
    <a id="L38"></a>for len(x.crypt) &lt; x.blockSize {
        <a id="L39"></a>off := len(x.crypt);
        <a id="L40"></a>var m int;
        <a id="L41"></a>m, err = x.r.Read(x.crypt[off:x.blockSize]);
        <a id="L42"></a>x.crypt = x.crypt[0 : off+m];
        <a id="L43"></a>if m == 0 {
            <a id="L44"></a>break
        <a id="L45"></a>}

        <a id="L47"></a><span class="comment">// If an error happened but we got enough</span>
        <a id="L48"></a><span class="comment">// data to do some decryption, we can decrypt</span>
        <a id="L49"></a><span class="comment">// first and report the error (with some data) later.</span>
        <a id="L50"></a><span class="comment">// But if we don&#39;t have enough to decrypt,</span>
        <a id="L51"></a><span class="comment">// have to stop now.</span>
        <a id="L52"></a>if err != nil &amp;&amp; len(x.crypt) &lt; x.blockSize {
            <a id="L53"></a>break
        <a id="L54"></a>}
    <a id="L55"></a>}
    <a id="L56"></a>return err;
<a id="L57"></a>}

<a id="L59"></a><span class="comment">// Read from plain text buffer into p.</span>
<a id="L60"></a>func (x *ecbDecrypter) readPlain(p []byte) int {
    <a id="L61"></a>n := len(x.plain);
    <a id="L62"></a>if n &gt; len(p) {
        <a id="L63"></a>n = len(p)
    <a id="L64"></a>}
    <a id="L65"></a>for i := 0; i &lt; n; i++ {
        <a id="L66"></a>p[i] = x.plain[i]
    <a id="L67"></a>}
    <a id="L68"></a>if n &lt; len(x.plain) {
        <a id="L69"></a>x.plain = x.plain[n:len(x.plain)]
    <a id="L70"></a>} else {
        <a id="L71"></a>x.plain = nil
    <a id="L72"></a>}
    <a id="L73"></a>return n;
<a id="L74"></a>}

<a id="L76"></a>type ecbFragmentError int

<a id="L78"></a>func (n ecbFragmentError) String() string {
    <a id="L79"></a>return &#34;crypto/block: &#34; + strconv.Itoa(int(n)) + &#34;-byte fragment at EOF&#34;
<a id="L80"></a>}

<a id="L82"></a>func (x *ecbDecrypter) Read(p []byte) (n int, err os.Error) {
    <a id="L83"></a>if len(p) == 0 {
        <a id="L84"></a>return
    <a id="L85"></a>}

    <a id="L87"></a><span class="comment">// If there&#39;s no plaintext waiting and p is not big enough</span>
    <a id="L88"></a><span class="comment">// to hold a whole cipher block, we&#39;ll have to work in the</span>
    <a id="L89"></a><span class="comment">// cipher text buffer.  Set it to non-nil so that the</span>
    <a id="L90"></a><span class="comment">// code below will fill it.</span>
    <a id="L91"></a>if x.plain == nil &amp;&amp; len(p) &lt; x.blockSize &amp;&amp; x.crypt == nil {
        <a id="L92"></a>x.crypt = x.buf[0:0]
    <a id="L93"></a>}

    <a id="L95"></a><span class="comment">// If there is a leftover cipher text buffer,</span>
    <a id="L96"></a><span class="comment">// try to accumulate a full block.</span>
    <a id="L97"></a>if x.crypt != nil {
        <a id="L98"></a>err = x.fillCrypt();
        <a id="L99"></a>if err != nil || len(x.crypt) == 0 {
            <a id="L100"></a>return
        <a id="L101"></a>}
        <a id="L102"></a>x.c.Decrypt(x.crypt, x.crypt);
        <a id="L103"></a>x.plain = x.crypt;
        <a id="L104"></a>x.crypt = nil;
    <a id="L105"></a>}

    <a id="L107"></a><span class="comment">// If there is a leftover plain text buffer, read from it.</span>
    <a id="L108"></a>if x.plain != nil {
        <a id="L109"></a>n = x.readPlain(p);
        <a id="L110"></a>return;
    <a id="L111"></a>}

    <a id="L113"></a><span class="comment">// Read and decrypt directly in caller&#39;s buffer.</span>
    <a id="L114"></a>n, err = io.ReadAtLeast(x.r, p, x.blockSize);
    <a id="L115"></a>if err == os.EOF &amp;&amp; n &gt; 0 {
        <a id="L116"></a><span class="comment">// EOF is only okay on block boundary</span>
        <a id="L117"></a>err = os.ErrorString(&#34;block fragment at EOF during decryption&#34;);
        <a id="L118"></a>return;
    <a id="L119"></a>}
    <a id="L120"></a>var i int;
    <a id="L121"></a>for i = 0; i+x.blockSize &lt;= n; i += x.blockSize {
        <a id="L122"></a>a := p[i : i+x.blockSize];
        <a id="L123"></a>x.c.Decrypt(a, a);
    <a id="L124"></a>}

    <a id="L126"></a><span class="comment">// There might be an encrypted fringe remaining.</span>
    <a id="L127"></a><span class="comment">// Save it for next time.</span>
    <a id="L128"></a>if i &lt; n {
        <a id="L129"></a>p = p[i:n];
        <a id="L130"></a>for j, v := range p {
            <a id="L131"></a>x.buf[j] = v
        <a id="L132"></a>}
        <a id="L133"></a>x.crypt = x.buf[0:len(p)];
        <a id="L134"></a>n = i;
    <a id="L135"></a>}

    <a id="L137"></a>return;
<a id="L138"></a>}

<a id="L140"></a><span class="comment">// NewECBDecrypter returns a reader that reads data from r and decrypts it using c.</span>
<a id="L141"></a><span class="comment">// It decrypts by calling c.Decrypt on each block in sequence;</span>
<a id="L142"></a><span class="comment">// this mode is known as electronic codebook mode, or ECB.</span>
<a id="L143"></a><span class="comment">// The returned Reader does not buffer or read ahead except</span>
<a id="L144"></a><span class="comment">// as required by the cipher&#39;s block size.</span>
<a id="L145"></a>func NewECBDecrypter(c Cipher, r io.Reader) io.Reader {
    <a id="L146"></a>x := new(ecbDecrypter);
    <a id="L147"></a>x.c = c;
    <a id="L148"></a>x.r = r;
    <a id="L149"></a>x.blockSize = c.BlockSize();
    <a id="L150"></a>x.buf = make([]byte, x.blockSize);
    <a id="L151"></a>return x;
<a id="L152"></a>}

<a id="L154"></a>type ecbEncrypter struct {
    <a id="L155"></a>c         Cipher;
    <a id="L156"></a>w         io.Writer;
    <a id="L157"></a>blockSize int;

    <a id="L159"></a><span class="comment">// Buffered data.</span>
    <a id="L160"></a><span class="comment">// The buffer buf is used as storage for both</span>
    <a id="L161"></a><span class="comment">// plain or crypt.  If both are non-nil, plain</span>
    <a id="L162"></a><span class="comment">// follows crypt in buf.</span>
    <a id="L163"></a>buf   []byte;
    <a id="L164"></a>plain []byte; <span class="comment">// plain text waiting to be encrypted</span>
    <a id="L165"></a>crypt []byte; <span class="comment">// encrypted text waiting to be written</span>
<a id="L166"></a>}

<a id="L168"></a><span class="comment">// Flush the x.crypt buffer to x.w.</span>
<a id="L169"></a>func (x *ecbEncrypter) flushCrypt() os.Error {
    <a id="L170"></a>if len(x.crypt) == 0 {
        <a id="L171"></a>return nil
    <a id="L172"></a>}
    <a id="L173"></a>n, err := x.w.Write(x.crypt);
    <a id="L174"></a>if n &lt; len(x.crypt) {
        <a id="L175"></a>x.crypt = x.crypt[n:len(x.crypt)];
        <a id="L176"></a>if err == nil {
            <a id="L177"></a>err = io.ErrShortWrite
        <a id="L178"></a>}
    <a id="L179"></a>}
    <a id="L180"></a>if err != nil {
        <a id="L181"></a>return err
    <a id="L182"></a>}
    <a id="L183"></a>x.crypt = nil;
    <a id="L184"></a>return nil;
<a id="L185"></a>}

<a id="L187"></a><span class="comment">// Slide x.plain down to the beginning of x.buf.</span>
<a id="L188"></a><span class="comment">// Plain is known to have less than one block of data,</span>
<a id="L189"></a><span class="comment">// so this is cheap enough.</span>
<a id="L190"></a>func (x *ecbEncrypter) slidePlain() {
    <a id="L191"></a>if len(x.plain) == 0 {
        <a id="L192"></a>x.plain = x.buf[0:0]
    <a id="L193"></a>} else if cap(x.plain) &lt; cap(x.buf) {
        <a id="L194"></a><span class="comment">// plain and buf share same data,</span>
        <a id="L195"></a><span class="comment">// but buf is before plain, so forward loop is correct</span>
        <a id="L196"></a>for i := 0; i &lt; len(x.plain); i++ {
            <a id="L197"></a>x.buf[i] = x.plain[i]
        <a id="L198"></a>}
        <a id="L199"></a>x.plain = x.buf[0:len(x.plain)];
    <a id="L200"></a>}
<a id="L201"></a>}

<a id="L203"></a><span class="comment">// Fill x.plain from the data in p.</span>
<a id="L204"></a><span class="comment">// Return the number of bytes copied.</span>
<a id="L205"></a>func (x *ecbEncrypter) fillPlain(p []byte) int {
    <a id="L206"></a>off := len(x.plain);
    <a id="L207"></a>n := len(p);
    <a id="L208"></a>if max := cap(x.plain) - off; n &gt; max {
        <a id="L209"></a>n = max
    <a id="L210"></a>}
    <a id="L211"></a>x.plain = x.plain[0 : off+n];
    <a id="L212"></a>for i := 0; i &lt; n; i++ {
        <a id="L213"></a>x.plain[off+i] = p[i]
    <a id="L214"></a>}
    <a id="L215"></a>return n;
<a id="L216"></a>}

<a id="L218"></a><span class="comment">// Encrypt x.plain; record encrypted range as x.crypt.</span>
<a id="L219"></a>func (x *ecbEncrypter) encrypt() {
    <a id="L220"></a>var i int;
    <a id="L221"></a>n := len(x.plain);
    <a id="L222"></a>for i = 0; i+x.blockSize &lt;= n; i += x.blockSize {
        <a id="L223"></a>a := x.plain[i : i+x.blockSize];
        <a id="L224"></a>x.c.Encrypt(a, a);
    <a id="L225"></a>}
    <a id="L226"></a>x.crypt = x.plain[0:i];
    <a id="L227"></a>x.plain = x.plain[i:n];
<a id="L228"></a>}

<a id="L230"></a>func (x *ecbEncrypter) Write(p []byte) (n int, err os.Error) {
    <a id="L231"></a>for {
        <a id="L232"></a><span class="comment">// If there is data waiting to be written, write it.</span>
        <a id="L233"></a><span class="comment">// This can happen on the first iteration</span>
        <a id="L234"></a><span class="comment">// if a write failed in an earlier call.</span>
        <a id="L235"></a>if err = x.flushCrypt(); err != nil {
            <a id="L236"></a>return
        <a id="L237"></a>}

        <a id="L239"></a><span class="comment">// Now that encrypted data is gone (flush ran),</span>
        <a id="L240"></a><span class="comment">// perhaps we need to slide the plaintext down.</span>
        <a id="L241"></a>x.slidePlain();

        <a id="L243"></a><span class="comment">// Fill plaintext buffer from p.</span>
        <a id="L244"></a>m := x.fillPlain(p);
        <a id="L245"></a>if m == 0 {
            <a id="L246"></a>break
        <a id="L247"></a>}
        <a id="L248"></a>n += m;
        <a id="L249"></a>p = p[m:len(p)];

        <a id="L251"></a><span class="comment">// Encrypt, adjusting crypt and plain.</span>
        <a id="L252"></a>x.encrypt();

        <a id="L254"></a><span class="comment">// Write x.crypt.</span>
        <a id="L255"></a>if err = x.flushCrypt(); err != nil {
            <a id="L256"></a>break
        <a id="L257"></a>}
    <a id="L258"></a>}
    <a id="L259"></a>return;
<a id="L260"></a>}

<a id="L262"></a><span class="comment">// NewECBEncrypter returns a writer that encrypts data using c and writes it to w.</span>
<a id="L263"></a><span class="comment">// It encrypts by calling c.Encrypt on each block in sequence;</span>
<a id="L264"></a><span class="comment">// this mode is known as electronic codebook mode, or ECB.</span>
<a id="L265"></a><span class="comment">// The returned Writer does no buffering except as required</span>
<a id="L266"></a><span class="comment">// by the cipher&#39;s block size, so there is no need for a Flush method.</span>
<a id="L267"></a>func NewECBEncrypter(c Cipher, w io.Writer) io.Writer {
    <a id="L268"></a>x := new(ecbEncrypter);
    <a id="L269"></a>x.c = c;
    <a id="L270"></a>x.w = w;
    <a id="L271"></a>x.blockSize = c.BlockSize();

    <a id="L273"></a><span class="comment">// Create a buffer that is an integral number of blocks.</span>
    <a id="L274"></a>x.buf = make([]byte, 8192/x.blockSize*x.blockSize);
    <a id="L275"></a>return x;
<a id="L276"></a>}
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
