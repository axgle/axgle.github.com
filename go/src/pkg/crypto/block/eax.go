<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/eax.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/eax.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// EAX mode, not a NIST standard (yet).</span>
<a id="L6"></a><span class="comment">// EAX provides encryption and authentication.</span>
<a id="L7"></a><span class="comment">// EAX targets the same uses as NIST&#39;s CCM mode,</span>
<a id="L8"></a><span class="comment">// but EAX adds the ability to run in streaming mode.</span>

<a id="L10"></a><span class="comment">// See</span>
<a id="L11"></a><span class="comment">// http://csrc.nist.gov/groups/ST/toolkit/BCM/documents/proposedmodes/eax/eax-spec.pdf</span>
<a id="L12"></a><span class="comment">// http://www.cs.ucdavis.edu/~rogaway/papers/eax.pdf</span>
<a id="L13"></a><span class="comment">// What those papers call OMAC is now called CMAC.</span>

<a id="L15"></a>package block

<a id="L17"></a>import (
    <a id="L18"></a>&#34;fmt&#34;;
    <a id="L19"></a>&#34;hash&#34;;
    <a id="L20"></a>&#34;io&#34;;
    <a id="L21"></a>&#34;os&#34;;
<a id="L22"></a>)

<a id="L24"></a><span class="comment">// An EAXTagError is returned when the message has failed to authenticate,</span>
<a id="L25"></a><span class="comment">// because the tag at the end of the message stream (Read) does not match</span>
<a id="L26"></a><span class="comment">// the tag computed from the message itself (Computed).</span>
<a id="L27"></a>type EAXTagError struct {
    <a id="L28"></a>Read     []byte;
    <a id="L29"></a>Computed []byte;
<a id="L30"></a>}

<a id="L32"></a>func (e *EAXTagError) String() string {
    <a id="L33"></a>return fmt.Sprintf(&#34;crypto/block: EAX tag mismatch: read %x but computed %x&#34;, e.Read, e.Computed)
<a id="L34"></a>}

<a id="L36"></a>func setupEAX(c Cipher, iv, hdr []byte, tagBytes int) (ctrIV, tag []byte, cmac hash.Hash) {
    <a id="L37"></a>n := len(iv);
    <a id="L38"></a>if n != c.BlockSize() {
        <a id="L39"></a>panicln(&#34;crypto/block: EAX: iv length&#34;, n, &#34;!=&#34;, c.BlockSize())
    <a id="L40"></a>}
    <a id="L41"></a>buf := make([]byte, n); <span class="comment">// zeroed</span>

    <a id="L43"></a><span class="comment">// tag = CMAC(0 + iv) ^ CMAC(1 + hdr) ^ CMAC(2 + data)</span>
    <a id="L44"></a>cmac = NewCMAC(c);
    <a id="L45"></a>cmac.Write(buf); <span class="comment">// 0</span>
    <a id="L46"></a>cmac.Write(iv);
    <a id="L47"></a>sum := cmac.Sum();
    <a id="L48"></a>ctrIV = copy(sum);
    <a id="L49"></a>tag = copy(sum[0:tagBytes]);

    <a id="L51"></a>cmac.Reset();
    <a id="L52"></a>buf[n-1] = 1;
    <a id="L53"></a>cmac.Write(buf); <span class="comment">// 1</span>
    <a id="L54"></a>cmac.Write(hdr);
    <a id="L55"></a>sum = cmac.Sum();
    <a id="L56"></a>for i := 0; i &lt; tagBytes; i++ {
        <a id="L57"></a>tag[i] ^= sum[i]
    <a id="L58"></a>}

    <a id="L60"></a>cmac.Reset();
    <a id="L61"></a>buf[n-1] = 2; <span class="comment">// 2</span>
    <a id="L62"></a>cmac.Write(buf);

    <a id="L64"></a>return;
<a id="L65"></a>}

<a id="L67"></a>func finishEAX(tag []byte, cmac hash.Hash) {
    <a id="L68"></a><span class="comment">// Finish CMAC #2 and xor into tag.</span>
    <a id="L69"></a>sum := cmac.Sum();
    <a id="L70"></a>for i := range tag {
        <a id="L71"></a>tag[i] ^= sum[i]
    <a id="L72"></a>}
<a id="L73"></a>}

<a id="L75"></a><span class="comment">// Writer adapter.  Tees writes into both w and cmac.</span>
<a id="L76"></a><span class="comment">// Knows that cmac never returns write errors.</span>
<a id="L77"></a>type cmacWriter struct {
    <a id="L78"></a>w    io.Writer;
    <a id="L79"></a>cmac hash.Hash;
<a id="L80"></a>}

<a id="L82"></a>func (cw *cmacWriter) Write(p []byte) (n int, err os.Error) {
    <a id="L83"></a>n, err = cw.w.Write(p);
    <a id="L84"></a>cw.cmac.Write(p[0:n]);
    <a id="L85"></a>return;
<a id="L86"></a>}

<a id="L88"></a><span class="comment">// An eaxEncrypter implements the EAX encryption mode.</span>
<a id="L89"></a>type eaxEncrypter struct {
    <a id="L90"></a>ctr io.Writer;  <span class="comment">// CTR encrypter</span>
    <a id="L91"></a>cw  cmacWriter; <span class="comment">// CTR&#39;s output stream</span>
    <a id="L92"></a>tag []byte;
<a id="L93"></a>}

<a id="L95"></a><span class="comment">// NewEAXEncrypter creates and returns a new EAX encrypter</span>
<a id="L96"></a><span class="comment">// using the given cipher c, initialization vector iv, associated data hdr,</span>
<a id="L97"></a><span class="comment">// and tag length tagBytes.  The encrypter&#39;s Write method encrypts</span>
<a id="L98"></a><span class="comment">// the data it receives and writes that data to w.</span>
<a id="L99"></a><span class="comment">// The encrypter&#39;s Close method writes a final authenticating tag to w.</span>
<a id="L100"></a>func NewEAXEncrypter(c Cipher, iv []byte, hdr []byte, tagBytes int, w io.Writer) io.WriteCloser {
    <a id="L101"></a>x := new(eaxEncrypter);

    <a id="L103"></a><span class="comment">// Create new CTR instance writing to both</span>
    <a id="L104"></a><span class="comment">// w for encrypted output and cmac for digesting.</span>
    <a id="L105"></a>x.cw.w = w;
    <a id="L106"></a>var ctrIV []byte;
    <a id="L107"></a>ctrIV, x.tag, x.cw.cmac = setupEAX(c, iv, hdr, tagBytes);
    <a id="L108"></a>x.ctr = NewCTRWriter(c, ctrIV, &amp;x.cw);
    <a id="L109"></a>return x;
<a id="L110"></a>}

<a id="L112"></a>func (x *eaxEncrypter) Write(p []byte) (n int, err os.Error) {
    <a id="L113"></a>return x.ctr.Write(p)
<a id="L114"></a>}

<a id="L116"></a>func (x *eaxEncrypter) Close() os.Error {
    <a id="L117"></a>x.ctr = nil; <span class="comment">// crash if Write is called again</span>

    <a id="L119"></a><span class="comment">// Write tag.</span>
    <a id="L120"></a>finishEAX(x.tag, x.cw.cmac);
    <a id="L121"></a>n, err := x.cw.w.Write(x.tag);
    <a id="L122"></a>if n != len(x.tag) &amp;&amp; err == nil {
        <a id="L123"></a>err = io.ErrShortWrite
    <a id="L124"></a>}

    <a id="L126"></a>return err;
<a id="L127"></a>}

<a id="L129"></a><span class="comment">// Reader adapter.  Returns data read from r but hangs</span>
<a id="L130"></a><span class="comment">// on to the last len(tag) bytes for itself (returns EOF len(tag)</span>
<a id="L131"></a><span class="comment">// bytes early).  Also tees all data returned from Read into</span>
<a id="L132"></a><span class="comment">// the cmac digest.  The &#34;don&#39;t return the last t bytes&#34;</span>
<a id="L133"></a><span class="comment">// and the &#34;tee into digest&#34; functionality could be separated,</span>
<a id="L134"></a><span class="comment">// but the latter half is trivial.</span>
<a id="L135"></a>type cmacReader struct {
    <a id="L136"></a>r    io.Reader;
    <a id="L137"></a>cmac hash.Hash;
    <a id="L138"></a>tag  []byte;
    <a id="L139"></a>tmp  []byte;
<a id="L140"></a>}

<a id="L142"></a>func (cr *cmacReader) Read(p []byte) (n int, err os.Error) {
    <a id="L143"></a><span class="comment">// TODO(rsc): Maybe fall back to simpler code if</span>
    <a id="L144"></a><span class="comment">// we recognize the underlying r as a ByteBuffer</span>
    <a id="L145"></a><span class="comment">// or ByteReader.  Then we can just take the last piece</span>
    <a id="L146"></a><span class="comment">// off at the start.</span>

    <a id="L148"></a><span class="comment">// First, read a tag-sized chunk.</span>
    <a id="L149"></a><span class="comment">// It&#39;s probably not the tag (unless there&#39;s no data).</span>
    <a id="L150"></a>tag := cr.tag;
    <a id="L151"></a>if len(tag) &lt; cap(tag) {
        <a id="L152"></a>nt := len(tag);
        <a id="L153"></a>nn, err1 := io.ReadFull(cr.r, tag[nt:cap(tag)]);
        <a id="L154"></a>tag = tag[0 : nt+nn];
        <a id="L155"></a>cr.tag = tag;
        <a id="L156"></a>if err1 != nil {
            <a id="L157"></a>return 0, err1
        <a id="L158"></a>}
    <a id="L159"></a>}

    <a id="L161"></a>tagBytes := len(tag);
    <a id="L162"></a>if len(p) &gt; 4*tagBytes {
        <a id="L163"></a><span class="comment">// If p is big, try to read directly into p to avoid a copy.</span>
        <a id="L164"></a>n, err = cr.r.Read(p[tagBytes:len(p)]);
        <a id="L165"></a>if n == 0 {
            <a id="L166"></a>goto out
        <a id="L167"></a>}
        <a id="L168"></a><span class="comment">// copy old tag into p</span>
        <a id="L169"></a>for i := 0; i &lt; tagBytes; i++ {
            <a id="L170"></a>p[i] = tag[i]
        <a id="L171"></a>}
        <a id="L172"></a><span class="comment">// copy new tag out of p</span>
        <a id="L173"></a>for i := 0; i &lt; tagBytes; i++ {
            <a id="L174"></a>tag[i] = p[n+i]
        <a id="L175"></a>}
        <a id="L176"></a>goto out;
    <a id="L177"></a>}

    <a id="L179"></a><span class="comment">// Otherwise, read into p and then slide data</span>
    <a id="L180"></a>n, err = cr.r.Read(p);
    <a id="L181"></a>if n == 0 {
        <a id="L182"></a>goto out
    <a id="L183"></a>}

    <a id="L185"></a><span class="comment">// copy tag+p into p+tmp and then swap tmp, tag</span>
    <a id="L186"></a>tmp := cr.tmp;
    <a id="L187"></a>for i := n + tagBytes - 1; i &gt;= 0; i-- {
        <a id="L188"></a>var c byte;
        <a id="L189"></a>if i &lt; tagBytes {
            <a id="L190"></a>c = tag[i]
        <a id="L191"></a>} else {
            <a id="L192"></a>c = p[i-tagBytes]
        <a id="L193"></a>}
        <a id="L194"></a>if i &lt; n {
            <a id="L195"></a>p[i] = c
        <a id="L196"></a>} else {
            <a id="L197"></a>tmp[i] = c
        <a id="L198"></a>}
    <a id="L199"></a>}
    <a id="L200"></a>cr.tmp, cr.tag = tag, tmp;

<a id="L202"></a>out:
    <a id="L203"></a>cr.cmac.Write(p[0:n]);
    <a id="L204"></a>return;
<a id="L205"></a>}

<a id="L207"></a>type eaxDecrypter struct {
    <a id="L208"></a>ctr io.Reader;
    <a id="L209"></a>cr  cmacReader;
    <a id="L210"></a>tag []byte;
<a id="L211"></a>}

<a id="L213"></a><span class="comment">// NewEAXDecrypter creates and returns a new EAX decrypter</span>
<a id="L214"></a><span class="comment">// using the given cipher c, initialization vector iv, associated data hdr,</span>
<a id="L215"></a><span class="comment">// and tag length tagBytes.  The encrypter&#39;s Read method decrypts and</span>
<a id="L216"></a><span class="comment">// returns data read from r.  At r&#39;s EOF, the encrypter checks the final</span>
<a id="L217"></a><span class="comment">// authenticating tag and returns an EAXTagError if the tag is invalid.</span>
<a id="L218"></a><span class="comment">// In that case, the message should be discarded.</span>
<a id="L219"></a><span class="comment">// Note that the data stream returned from Read cannot be</span>
<a id="L220"></a><span class="comment">// assumed to be valid, authenticated data until Read returns</span>
<a id="L221"></a><span class="comment">// 0, nil to signal the end of the data.</span>
<a id="L222"></a>func NewEAXDecrypter(c Cipher, iv []byte, hdr []byte, tagBytes int, r io.Reader) io.Reader {
    <a id="L223"></a>x := new(eaxDecrypter);

    <a id="L225"></a>x.cr.r = r;
    <a id="L226"></a>x.cr.tag = make([]byte, 0, tagBytes);
    <a id="L227"></a>x.cr.tmp = make([]byte, 0, tagBytes);
    <a id="L228"></a>var ctrIV []byte;
    <a id="L229"></a>ctrIV, x.tag, x.cr.cmac = setupEAX(c, iv, hdr, tagBytes);
    <a id="L230"></a>x.ctr = NewCTRReader(c, ctrIV, &amp;x.cr);
    <a id="L231"></a>return x;
<a id="L232"></a>}

<a id="L234"></a>func (x *eaxDecrypter) checkTag() os.Error {
    <a id="L235"></a>x.ctr = nil; <span class="comment">// crash if Read is called again</span>

    <a id="L237"></a>finishEAX(x.tag, x.cr.cmac);
    <a id="L238"></a>if !same(x.tag, x.cr.tag) {
        <a id="L239"></a>e := new(EAXTagError);
        <a id="L240"></a>e.Computed = copy(x.tag);
        <a id="L241"></a>e.Read = copy(x.cr.tag);
        <a id="L242"></a>return e;
    <a id="L243"></a>}
    <a id="L244"></a>return nil;
<a id="L245"></a>}

<a id="L247"></a>func (x *eaxDecrypter) Read(p []byte) (n int, err os.Error) {
    <a id="L248"></a>n, err = x.ctr.Read(p);
    <a id="L249"></a>if n == 0 &amp;&amp; err == nil {
        <a id="L250"></a>err = x.checkTag()
    <a id="L251"></a>}
    <a id="L252"></a>return n, err;
<a id="L253"></a>}
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
