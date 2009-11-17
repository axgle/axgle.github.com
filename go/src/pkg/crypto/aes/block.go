<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/aes/block.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/aes/block.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This Go implementation is derived in part from the reference</span>
<a id="L6"></a><span class="comment">// ANSI C implementation, which carries the following notice:</span>
<a id="L7"></a><span class="comment">//</span>
<a id="L8"></a><span class="comment">//	rijndael-alg-fst.c</span>
<a id="L9"></a><span class="comment">//</span>
<a id="L10"></a><span class="comment">//	@version 3.0 (December 2000)</span>
<a id="L11"></a><span class="comment">//</span>
<a id="L12"></a><span class="comment">//	Optimised ANSI C code for the Rijndael cipher (now AES)</span>
<a id="L13"></a><span class="comment">//</span>
<a id="L14"></a><span class="comment">//	@author Vincent Rijmen &lt;vincent.rijmen@esat.kuleuven.ac.be&gt;</span>
<a id="L15"></a><span class="comment">//	@author Antoon Bosselaers &lt;antoon.bosselaers@esat.kuleuven.ac.be&gt;</span>
<a id="L16"></a><span class="comment">//	@author Paulo Barreto &lt;paulo.barreto@terra.com.br&gt;</span>
<a id="L17"></a><span class="comment">//</span>
<a id="L18"></a><span class="comment">//	This code is hereby placed in the public domain.</span>
<a id="L19"></a><span class="comment">//</span>
<a id="L20"></a><span class="comment">//	THIS SOFTWARE IS PROVIDED BY THE AUTHORS &#39;&#39;AS IS&#39;&#39; AND ANY EXPRESS</span>
<a id="L21"></a><span class="comment">//	OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED</span>
<a id="L22"></a><span class="comment">//	WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE</span>
<a id="L23"></a><span class="comment">//	ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHORS OR CONTRIBUTORS BE</span>
<a id="L24"></a><span class="comment">//	LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR</span>
<a id="L25"></a><span class="comment">//	CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF</span>
<a id="L26"></a><span class="comment">//	SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR</span>
<a id="L27"></a><span class="comment">//	BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,</span>
<a id="L28"></a><span class="comment">//	WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE</span>
<a id="L29"></a><span class="comment">//	OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE,</span>
<a id="L30"></a><span class="comment">//	EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.</span>
<a id="L31"></a><span class="comment">//</span>
<a id="L32"></a><span class="comment">// See FIPS 197 for specification, and see Daemen and Rijmen&#39;s Rijndael submission</span>
<a id="L33"></a><span class="comment">// for implementation details.</span>
<a id="L34"></a><span class="comment">//	http://www.csrc.nist.gov/publications/fips/fips197/fips-197.pdf</span>
<a id="L35"></a><span class="comment">//	http://csrc.nist.gov/archive/aes/rijndael/Rijndael-ammended.pdf</span>

<a id="L37"></a>package aes

<a id="L39"></a><span class="comment">// Encrypt one block from src into dst, using the expanded key xk.</span>
<a id="L40"></a>func encryptBlock(xk []uint32, src, dst []byte) {
    <a id="L41"></a>var s0, s1, s2, s3, t0, t1, t2, t3 uint32;

    <a id="L43"></a>s0 = uint32(src[0])&lt;&lt;24 | uint32(src[1])&lt;&lt;16 | uint32(src[2])&lt;&lt;8 | uint32(src[3]);
    <a id="L44"></a>s1 = uint32(src[4])&lt;&lt;24 | uint32(src[5])&lt;&lt;16 | uint32(src[6])&lt;&lt;8 | uint32(src[7]);
    <a id="L45"></a>s2 = uint32(src[8])&lt;&lt;24 | uint32(src[9])&lt;&lt;16 | uint32(src[10])&lt;&lt;8 | uint32(src[11]);
    <a id="L46"></a>s3 = uint32(src[12])&lt;&lt;24 | uint32(src[13])&lt;&lt;16 | uint32(src[14])&lt;&lt;8 | uint32(src[15]);

    <a id="L48"></a><span class="comment">// First round just XORs input with key.</span>
    <a id="L49"></a>s0 ^= xk[0];
    <a id="L50"></a>s1 ^= xk[1];
    <a id="L51"></a>s2 ^= xk[2];
    <a id="L52"></a>s3 ^= xk[3];

    <a id="L54"></a><span class="comment">// Middle rounds shuffle using tables.</span>
    <a id="L55"></a><span class="comment">// Number of rounds is set by length of expanded key.</span>
    <a id="L56"></a>nr := len(xk)/4 - 2; <span class="comment">// - 2: one above, one more below</span>
    <a id="L57"></a>k := 4;
    <a id="L58"></a>for r := 0; r &lt; nr; r++ {
        <a id="L59"></a>t0 = xk[k+0] ^ te[0][s0&gt;&gt;24] ^ te[1][s1&gt;&gt;16&amp;0xff] ^ te[2][s2&gt;&gt;8&amp;0xff] ^ te[3][s3&amp;0xff];
        <a id="L60"></a>t1 = xk[k+1] ^ te[0][s1&gt;&gt;24] ^ te[1][s2&gt;&gt;16&amp;0xff] ^ te[2][s3&gt;&gt;8&amp;0xff] ^ te[3][s0&amp;0xff];
        <a id="L61"></a>t2 = xk[k+2] ^ te[0][s2&gt;&gt;24] ^ te[1][s3&gt;&gt;16&amp;0xff] ^ te[2][s0&gt;&gt;8&amp;0xff] ^ te[3][s1&amp;0xff];
        <a id="L62"></a>t3 = xk[k+3] ^ te[0][s3&gt;&gt;24] ^ te[1][s0&gt;&gt;16&amp;0xff] ^ te[2][s1&gt;&gt;8&amp;0xff] ^ te[3][s2&amp;0xff];
        <a id="L63"></a>k += 4;
        <a id="L64"></a>s0, s1, s2, s3 = t0, t1, t2, t3;
    <a id="L65"></a>}

    <a id="L67"></a><span class="comment">// Last round uses s-box directly and XORs to produce output.</span>
    <a id="L68"></a>s0 = uint32(sbox0[t0&gt;&gt;24])&lt;&lt;24 | uint32(sbox0[t1&gt;&gt;16&amp;0xff])&lt;&lt;16 | uint32(sbox0[t2&gt;&gt;8&amp;0xff])&lt;&lt;8 | uint32(sbox0[t3&amp;0xff]);
    <a id="L69"></a>s1 = uint32(sbox0[t1&gt;&gt;24])&lt;&lt;24 | uint32(sbox0[t2&gt;&gt;16&amp;0xff])&lt;&lt;16 | uint32(sbox0[t3&gt;&gt;8&amp;0xff])&lt;&lt;8 | uint32(sbox0[t0&amp;0xff]);
    <a id="L70"></a>s2 = uint32(sbox0[t2&gt;&gt;24])&lt;&lt;24 | uint32(sbox0[t3&gt;&gt;16&amp;0xff])&lt;&lt;16 | uint32(sbox0[t0&gt;&gt;8&amp;0xff])&lt;&lt;8 | uint32(sbox0[t1&amp;0xff]);
    <a id="L71"></a>s3 = uint32(sbox0[t3&gt;&gt;24])&lt;&lt;24 | uint32(sbox0[t0&gt;&gt;16&amp;0xff])&lt;&lt;16 | uint32(sbox0[t1&gt;&gt;8&amp;0xff])&lt;&lt;8 | uint32(sbox0[t2&amp;0xff]);

    <a id="L73"></a>s0 ^= xk[k+0];
    <a id="L74"></a>s1 ^= xk[k+1];
    <a id="L75"></a>s2 ^= xk[k+2];
    <a id="L76"></a>s3 ^= xk[k+3];

    <a id="L78"></a>dst[0], dst[1], dst[2], dst[3] = byte(s0&gt;&gt;24), byte(s0&gt;&gt;16), byte(s0&gt;&gt;8), byte(s0);
    <a id="L79"></a>dst[4], dst[5], dst[6], dst[7] = byte(s1&gt;&gt;24), byte(s1&gt;&gt;16), byte(s1&gt;&gt;8), byte(s1);
    <a id="L80"></a>dst[8], dst[9], dst[10], dst[11] = byte(s2&gt;&gt;24), byte(s2&gt;&gt;16), byte(s2&gt;&gt;8), byte(s2);
    <a id="L81"></a>dst[12], dst[13], dst[14], dst[15] = byte(s3&gt;&gt;24), byte(s3&gt;&gt;16), byte(s3&gt;&gt;8), byte(s3);
<a id="L82"></a>}

<a id="L84"></a><span class="comment">// Decrypt one block from src into dst, using the expanded key xk.</span>
<a id="L85"></a>func decryptBlock(xk []uint32, src, dst []byte) {
    <a id="L86"></a>var s0, s1, s2, s3, t0, t1, t2, t3 uint32;

    <a id="L88"></a>s0 = uint32(src[0])&lt;&lt;24 | uint32(src[1])&lt;&lt;16 | uint32(src[2])&lt;&lt;8 | uint32(src[3]);
    <a id="L89"></a>s1 = uint32(src[4])&lt;&lt;24 | uint32(src[5])&lt;&lt;16 | uint32(src[6])&lt;&lt;8 | uint32(src[7]);
    <a id="L90"></a>s2 = uint32(src[8])&lt;&lt;24 | uint32(src[9])&lt;&lt;16 | uint32(src[10])&lt;&lt;8 | uint32(src[11]);
    <a id="L91"></a>s3 = uint32(src[12])&lt;&lt;24 | uint32(src[13])&lt;&lt;16 | uint32(src[14])&lt;&lt;8 | uint32(src[15]);

    <a id="L93"></a><span class="comment">// First round just XORs input with key.</span>
    <a id="L94"></a>s0 ^= xk[0];
    <a id="L95"></a>s1 ^= xk[1];
    <a id="L96"></a>s2 ^= xk[2];
    <a id="L97"></a>s3 ^= xk[3];

    <a id="L99"></a><span class="comment">// Middle rounds shuffle using tables.</span>
    <a id="L100"></a><span class="comment">// Number of rounds is set by length of expanded key.</span>
    <a id="L101"></a>nr := len(xk)/4 - 2; <span class="comment">// - 2: one above, one more below</span>
    <a id="L102"></a>k := 4;
    <a id="L103"></a>for r := 0; r &lt; nr; r++ {
        <a id="L104"></a>t0 = xk[k+0] ^ td[0][s0&gt;&gt;24] ^ td[1][s3&gt;&gt;16&amp;0xff] ^ td[2][s2&gt;&gt;8&amp;0xff] ^ td[3][s1&amp;0xff];
        <a id="L105"></a>t1 = xk[k+1] ^ td[0][s1&gt;&gt;24] ^ td[1][s0&gt;&gt;16&amp;0xff] ^ td[2][s3&gt;&gt;8&amp;0xff] ^ td[3][s2&amp;0xff];
        <a id="L106"></a>t2 = xk[k+2] ^ td[0][s2&gt;&gt;24] ^ td[1][s1&gt;&gt;16&amp;0xff] ^ td[2][s0&gt;&gt;8&amp;0xff] ^ td[3][s3&amp;0xff];
        <a id="L107"></a>t3 = xk[k+3] ^ td[0][s3&gt;&gt;24] ^ td[1][s2&gt;&gt;16&amp;0xff] ^ td[2][s1&gt;&gt;8&amp;0xff] ^ td[3][s0&amp;0xff];
        <a id="L108"></a>k += 4;
        <a id="L109"></a>s0, s1, s2, s3 = t0, t1, t2, t3;
    <a id="L110"></a>}

    <a id="L112"></a><span class="comment">// Last round uses s-box directly and XORs to produce output.</span>
    <a id="L113"></a>s0 = uint32(sbox1[t0&gt;&gt;24])&lt;&lt;24 | uint32(sbox1[t3&gt;&gt;16&amp;0xff])&lt;&lt;16 | uint32(sbox1[t2&gt;&gt;8&amp;0xff])&lt;&lt;8 | uint32(sbox1[t1&amp;0xff]);
    <a id="L114"></a>s1 = uint32(sbox1[t1&gt;&gt;24])&lt;&lt;24 | uint32(sbox1[t0&gt;&gt;16&amp;0xff])&lt;&lt;16 | uint32(sbox1[t3&gt;&gt;8&amp;0xff])&lt;&lt;8 | uint32(sbox1[t2&amp;0xff]);
    <a id="L115"></a>s2 = uint32(sbox1[t2&gt;&gt;24])&lt;&lt;24 | uint32(sbox1[t1&gt;&gt;16&amp;0xff])&lt;&lt;16 | uint32(sbox1[t0&gt;&gt;8&amp;0xff])&lt;&lt;8 | uint32(sbox1[t3&amp;0xff]);
    <a id="L116"></a>s3 = uint32(sbox1[t3&gt;&gt;24])&lt;&lt;24 | uint32(sbox1[t2&gt;&gt;16&amp;0xff])&lt;&lt;16 | uint32(sbox1[t1&gt;&gt;8&amp;0xff])&lt;&lt;8 | uint32(sbox1[t0&amp;0xff]);

    <a id="L118"></a>s0 ^= xk[k+0];
    <a id="L119"></a>s1 ^= xk[k+1];
    <a id="L120"></a>s2 ^= xk[k+2];
    <a id="L121"></a>s3 ^= xk[k+3];

    <a id="L123"></a>dst[0], dst[1], dst[2], dst[3] = byte(s0&gt;&gt;24), byte(s0&gt;&gt;16), byte(s0&gt;&gt;8), byte(s0);
    <a id="L124"></a>dst[4], dst[5], dst[6], dst[7] = byte(s1&gt;&gt;24), byte(s1&gt;&gt;16), byte(s1&gt;&gt;8), byte(s1);
    <a id="L125"></a>dst[8], dst[9], dst[10], dst[11] = byte(s2&gt;&gt;24), byte(s2&gt;&gt;16), byte(s2&gt;&gt;8), byte(s2);
    <a id="L126"></a>dst[12], dst[13], dst[14], dst[15] = byte(s3&gt;&gt;24), byte(s3&gt;&gt;16), byte(s3&gt;&gt;8), byte(s3);
<a id="L127"></a>}

<a id="L129"></a><span class="comment">// Apply sbox0 to each byte in w.</span>
<a id="L130"></a>func subw(w uint32) uint32 {
    <a id="L131"></a>return uint32(sbox0[w&gt;&gt;24])&lt;&lt;24 |
        <a id="L132"></a>uint32(sbox0[w&gt;&gt;16&amp;0xff])&lt;&lt;16 |
        <a id="L133"></a>uint32(sbox0[w&gt;&gt;8&amp;0xff])&lt;&lt;8 |
        <a id="L134"></a>uint32(sbox0[w&amp;0xff])
<a id="L135"></a>}

<a id="L137"></a><span class="comment">// Rotate</span>
<a id="L138"></a>func rotw(w uint32) uint32 { return w&lt;&lt;8 | w&gt;&gt;24 }

<a id="L140"></a><span class="comment">// Key expansion algorithm.  See FIPS-197, Figure 11.</span>
<a id="L141"></a><span class="comment">// Their rcon[i] is our powx[i-1] &lt;&lt; 24.</span>
<a id="L142"></a>func expandKey(key []byte, enc, dec []uint32) {
    <a id="L143"></a><span class="comment">// Encryption key setup.</span>
    <a id="L144"></a>var i int;
    <a id="L145"></a>nk := len(key) / 4;
    <a id="L146"></a>for i = 0; i &lt; nk; i++ {
        <a id="L147"></a>enc[i] = uint32(key[4*i])&lt;&lt;24 | uint32(key[4*i+1])&lt;&lt;16 | uint32(key[4*i+2])&lt;&lt;8 | uint32(key[4*i+3])
    <a id="L148"></a>}
    <a id="L149"></a>for ; i &lt; len(enc); i++ {
        <a id="L150"></a>t := enc[i-1];
        <a id="L151"></a>if i%nk == 0 {
            <a id="L152"></a>t = subw(rotw(t)) ^ (uint32(powx[i/nk-1]) &lt;&lt; 24)
        <a id="L153"></a>} else if nk &gt; 6 &amp;&amp; i%nk == 4 {
            <a id="L154"></a>t = subw(t)
        <a id="L155"></a>}
        <a id="L156"></a>enc[i] = enc[i-nk] ^ t;
    <a id="L157"></a>}

    <a id="L159"></a><span class="comment">// Derive decryption key from encryption key.</span>
    <a id="L160"></a><span class="comment">// Reverse the 4-word round key sets from enc to produce dec.</span>
    <a id="L161"></a><span class="comment">// All sets but the first and last get the MixColumn transform applied.</span>
    <a id="L162"></a>if dec == nil {
        <a id="L163"></a>return
    <a id="L164"></a>}
    <a id="L165"></a>n := len(enc);
    <a id="L166"></a>for i := 0; i &lt; n; i += 4 {
        <a id="L167"></a>ei := n - i - 4;
        <a id="L168"></a>for j := 0; j &lt; 4; j++ {
            <a id="L169"></a>x := enc[ei+j];
            <a id="L170"></a>if i &gt; 0 &amp;&amp; i+4 &lt; n {
                <a id="L171"></a>x = td[0][sbox0[x&gt;&gt;24]] ^ td[1][sbox0[x&gt;&gt;16&amp;0xff]] ^ td[2][sbox0[x&gt;&gt;8&amp;0xff]] ^ td[3][sbox0[x&amp;0xff]]
            <a id="L172"></a>}
            <a id="L173"></a>dec[i+j] = x;
        <a id="L174"></a>}
    <a id="L175"></a>}
<a id="L176"></a>}
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
