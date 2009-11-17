<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/utf8/utf8.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/utf8/utf8.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Functions and constants to support text encoded in UTF-8.</span>
<a id="L6"></a><span class="comment">// This package calls a Unicode character a rune for brevity.</span>
<a id="L7"></a>package utf8

<a id="L9"></a>import &#34;unicode&#34; <span class="comment">// only needed for a couple of constants</span>

<a id="L11"></a><span class="comment">// Numbers fundamental to the encoding.</span>
<a id="L12"></a>const (
    <a id="L13"></a>RuneError = unicode.ReplacementChar; <span class="comment">// the &#34;error&#34; Rune or &#34;replacement character&#34;.</span>
    <a id="L14"></a>RuneSelf  = 0x80;                    <span class="comment">// characters below Runeself are represented as themselves in a single byte.</span>
    <a id="L15"></a>UTFMax    = 4;                       <span class="comment">// maximum number of bytes of a UTF-8 encoded Unicode character.</span>
<a id="L16"></a>)

<a id="L18"></a>const (
    <a id="L19"></a>_T1 = 0x00; <span class="comment">// 0000 0000</span>
    <a id="L20"></a>_Tx = 0x80; <span class="comment">// 1000 0000</span>
    <a id="L21"></a>_T2 = 0xC0; <span class="comment">// 1100 0000</span>
    <a id="L22"></a>_T3 = 0xE0; <span class="comment">// 1110 0000</span>
    <a id="L23"></a>_T4 = 0xF0; <span class="comment">// 1111 0000</span>
    <a id="L24"></a>_T5 = 0xF8; <span class="comment">// 1111 1000</span>

    <a id="L26"></a>_Maskx = 0x3F; <span class="comment">// 0011 1111</span>
    <a id="L27"></a>_Mask2 = 0x1F; <span class="comment">// 0001 1111</span>
    <a id="L28"></a>_Mask3 = 0x0F; <span class="comment">// 0000 1111</span>
    <a id="L29"></a>_Mask4 = 0x07; <span class="comment">// 0000 0111</span>

    <a id="L31"></a>_Rune1Max = 1&lt;&lt;7 - 1;
    <a id="L32"></a>_Rune2Max = 1&lt;&lt;11 - 1;
    <a id="L33"></a>_Rune3Max = 1&lt;&lt;16 - 1;
    <a id="L34"></a>_Rune4Max = 1&lt;&lt;21 - 1;
<a id="L35"></a>)

<a id="L37"></a>func decodeRuneInternal(p []byte) (rune, size int, short bool) {
    <a id="L38"></a>n := len(p);
    <a id="L39"></a>if n &lt; 1 {
        <a id="L40"></a>return RuneError, 0, true
    <a id="L41"></a>}
    <a id="L42"></a>c0 := p[0];

    <a id="L44"></a><span class="comment">// 1-byte, 7-bit sequence?</span>
    <a id="L45"></a>if c0 &lt; _Tx {
        <a id="L46"></a>return int(c0), 1, false
    <a id="L47"></a>}

    <a id="L49"></a><span class="comment">// unexpected continuation byte?</span>
    <a id="L50"></a>if c0 &lt; _T2 {
        <a id="L51"></a>return RuneError, 1, false
    <a id="L52"></a>}

    <a id="L54"></a><span class="comment">// need first continuation byte</span>
    <a id="L55"></a>if n &lt; 2 {
        <a id="L56"></a>return RuneError, 1, true
    <a id="L57"></a>}
    <a id="L58"></a>c1 := p[1];
    <a id="L59"></a>if c1 &lt; _Tx || _T2 &lt;= c1 {
        <a id="L60"></a>return RuneError, 1, false
    <a id="L61"></a>}

    <a id="L63"></a><span class="comment">// 2-byte, 11-bit sequence?</span>
    <a id="L64"></a>if c0 &lt; _T3 {
        <a id="L65"></a>rune = int(c0&amp;_Mask2)&lt;&lt;6 | int(c1&amp;_Maskx);
        <a id="L66"></a>if rune &lt;= _Rune1Max {
            <a id="L67"></a>return RuneError, 1, false
        <a id="L68"></a>}
        <a id="L69"></a>return rune, 2, false;
    <a id="L70"></a>}

    <a id="L72"></a><span class="comment">// need second continuation byte</span>
    <a id="L73"></a>if n &lt; 3 {
        <a id="L74"></a>return RuneError, 1, true
    <a id="L75"></a>}
    <a id="L76"></a>c2 := p[2];
    <a id="L77"></a>if c2 &lt; _Tx || _T2 &lt;= c2 {
        <a id="L78"></a>return RuneError, 1, false
    <a id="L79"></a>}

    <a id="L81"></a><span class="comment">// 3-byte, 16-bit sequence?</span>
    <a id="L82"></a>if c0 &lt; _T4 {
        <a id="L83"></a>rune = int(c0&amp;_Mask3)&lt;&lt;12 | int(c1&amp;_Maskx)&lt;&lt;6 | int(c2&amp;_Maskx);
        <a id="L84"></a>if rune &lt;= _Rune2Max {
            <a id="L85"></a>return RuneError, 1, false
        <a id="L86"></a>}
        <a id="L87"></a>return rune, 3, false;
    <a id="L88"></a>}

    <a id="L90"></a><span class="comment">// need third continuation byte</span>
    <a id="L91"></a>if n &lt; 4 {
        <a id="L92"></a>return RuneError, 1, true
    <a id="L93"></a>}
    <a id="L94"></a>c3 := p[3];
    <a id="L95"></a>if c3 &lt; _Tx || _T2 &lt;= c3 {
        <a id="L96"></a>return RuneError, 1, false
    <a id="L97"></a>}

    <a id="L99"></a><span class="comment">// 4-byte, 21-bit sequence?</span>
    <a id="L100"></a>if c0 &lt; _T5 {
        <a id="L101"></a>rune = int(c0&amp;_Mask4)&lt;&lt;18 | int(c1&amp;_Maskx)&lt;&lt;12 | int(c2&amp;_Maskx)&lt;&lt;6 | int(c3&amp;_Maskx);
        <a id="L102"></a>if rune &lt;= _Rune3Max {
            <a id="L103"></a>return RuneError, 1, false
        <a id="L104"></a>}
        <a id="L105"></a>return rune, 4, false;
    <a id="L106"></a>}

    <a id="L108"></a><span class="comment">// error</span>
    <a id="L109"></a>return RuneError, 1, false;
<a id="L110"></a>}

<a id="L112"></a>func decodeRuneInStringInternal(s string) (rune, size int, short bool) {
    <a id="L113"></a>n := len(s);
    <a id="L114"></a>if n &lt; 1 {
        <a id="L115"></a>return RuneError, 0, true
    <a id="L116"></a>}
    <a id="L117"></a>c0 := s[0];

    <a id="L119"></a><span class="comment">// 1-byte, 7-bit sequence?</span>
    <a id="L120"></a>if c0 &lt; _Tx {
        <a id="L121"></a>return int(c0), 1, false
    <a id="L122"></a>}

    <a id="L124"></a><span class="comment">// unexpected continuation byte?</span>
    <a id="L125"></a>if c0 &lt; _T2 {
        <a id="L126"></a>return RuneError, 1, false
    <a id="L127"></a>}

    <a id="L129"></a><span class="comment">// need first continuation byte</span>
    <a id="L130"></a>if n &lt; 2 {
        <a id="L131"></a>return RuneError, 1, true
    <a id="L132"></a>}
    <a id="L133"></a>c1 := s[1];
    <a id="L134"></a>if c1 &lt; _Tx || _T2 &lt;= c1 {
        <a id="L135"></a>return RuneError, 1, false
    <a id="L136"></a>}

    <a id="L138"></a><span class="comment">// 2-byte, 11-bit sequence?</span>
    <a id="L139"></a>if c0 &lt; _T3 {
        <a id="L140"></a>rune = int(c0&amp;_Mask2)&lt;&lt;6 | int(c1&amp;_Maskx);
        <a id="L141"></a>if rune &lt;= _Rune1Max {
            <a id="L142"></a>return RuneError, 1, false
        <a id="L143"></a>}
        <a id="L144"></a>return rune, 2, false;
    <a id="L145"></a>}

    <a id="L147"></a><span class="comment">// need second continuation byte</span>
    <a id="L148"></a>if n &lt; 3 {
        <a id="L149"></a>return RuneError, 1, true
    <a id="L150"></a>}
    <a id="L151"></a>c2 := s[2];
    <a id="L152"></a>if c2 &lt; _Tx || _T2 &lt;= c2 {
        <a id="L153"></a>return RuneError, 1, false
    <a id="L154"></a>}

    <a id="L156"></a><span class="comment">// 3-byte, 16-bit sequence?</span>
    <a id="L157"></a>if c0 &lt; _T4 {
        <a id="L158"></a>rune = int(c0&amp;_Mask3)&lt;&lt;12 | int(c1&amp;_Maskx)&lt;&lt;6 | int(c2&amp;_Maskx);
        <a id="L159"></a>if rune &lt;= _Rune2Max {
            <a id="L160"></a>return RuneError, 1, false
        <a id="L161"></a>}
        <a id="L162"></a>return rune, 3, false;
    <a id="L163"></a>}

    <a id="L165"></a><span class="comment">// need third continuation byte</span>
    <a id="L166"></a>if n &lt; 4 {
        <a id="L167"></a>return RuneError, 1, true
    <a id="L168"></a>}
    <a id="L169"></a>c3 := s[3];
    <a id="L170"></a>if c3 &lt; _Tx || _T2 &lt;= c3 {
        <a id="L171"></a>return RuneError, 1, false
    <a id="L172"></a>}

    <a id="L174"></a><span class="comment">// 4-byte, 21-bit sequence?</span>
    <a id="L175"></a>if c0 &lt; _T5 {
        <a id="L176"></a>rune = int(c0&amp;_Mask4)&lt;&lt;18 | int(c1&amp;_Maskx)&lt;&lt;12 | int(c2&amp;_Maskx)&lt;&lt;6 | int(c3&amp;_Maskx);
        <a id="L177"></a>if rune &lt;= _Rune3Max {
            <a id="L178"></a>return RuneError, 1, false
        <a id="L179"></a>}
        <a id="L180"></a>return rune, 4, false;
    <a id="L181"></a>}

    <a id="L183"></a><span class="comment">// error</span>
    <a id="L184"></a>return RuneError, 1, false;
<a id="L185"></a>}

<a id="L187"></a><span class="comment">// FullRune reports whether the bytes in p begin with a full UTF-8 encoding of a rune.</span>
<a id="L188"></a><span class="comment">// An invalid encoding is considered a full Rune since it will convert as a width-1 error rune.</span>
<a id="L189"></a>func FullRune(p []byte) bool {
    <a id="L190"></a>_, _, short := decodeRuneInternal(p);
    <a id="L191"></a>return !short;
<a id="L192"></a>}

<a id="L194"></a><span class="comment">// FullRuneInString is like FullRune but its input is a string.</span>
<a id="L195"></a>func FullRuneInString(s string) bool {
    <a id="L196"></a>_, _, short := decodeRuneInStringInternal(s);
    <a id="L197"></a>return !short;
<a id="L198"></a>}

<a id="L200"></a><span class="comment">// DecodeRune unpacks the first UTF-8 encoding in p and returns the rune and its width in bytes.</span>
<a id="L201"></a>func DecodeRune(p []byte) (rune, size int) {
    <a id="L202"></a>rune, size, _ = decodeRuneInternal(p);
    <a id="L203"></a>return;
<a id="L204"></a>}

<a id="L206"></a><span class="comment">// DecodeRuneInString is like DecodeRune but its input is a string.</span>
<a id="L207"></a>func DecodeRuneInString(s string) (rune, size int) {
    <a id="L208"></a>rune, size, _ = decodeRuneInStringInternal(s);
    <a id="L209"></a>return;
<a id="L210"></a>}

<a id="L212"></a><span class="comment">// RuneLen returns the number of bytes required to encode the rune.</span>
<a id="L213"></a>func RuneLen(rune int) int {
    <a id="L214"></a>switch {
    <a id="L215"></a>case rune &lt;= _Rune1Max:
        <a id="L216"></a>return 1
    <a id="L217"></a>case rune &lt;= _Rune2Max:
        <a id="L218"></a>return 2
    <a id="L219"></a>case rune &lt;= _Rune3Max:
        <a id="L220"></a>return 3
    <a id="L221"></a>case rune &lt;= _Rune4Max:
        <a id="L222"></a>return 4
    <a id="L223"></a>}
    <a id="L224"></a>return -1;
<a id="L225"></a>}

<a id="L227"></a><span class="comment">// EncodeRune writes into p (which must be large enough) the UTF-8 encoding of the rune.</span>
<a id="L228"></a><span class="comment">// It returns the number of bytes written.</span>
<a id="L229"></a>func EncodeRune(rune int, p []byte) int {
    <a id="L230"></a>if rune &lt;= _Rune1Max {
        <a id="L231"></a>p[0] = byte(rune);
        <a id="L232"></a>return 1;
    <a id="L233"></a>}

    <a id="L235"></a>if rune &lt;= _Rune2Max {
        <a id="L236"></a>p[0] = _T2 | byte(rune&gt;&gt;6);
        <a id="L237"></a>p[1] = _Tx | byte(rune)&amp;_Maskx;
        <a id="L238"></a>return 2;
    <a id="L239"></a>}

    <a id="L241"></a>if rune &gt; unicode.MaxRune {
        <a id="L242"></a>rune = RuneError
    <a id="L243"></a>}

    <a id="L245"></a>if rune &lt;= _Rune3Max {
        <a id="L246"></a>p[0] = _T3 | byte(rune&gt;&gt;12);
        <a id="L247"></a>p[1] = _Tx | byte(rune&gt;&gt;6)&amp;_Maskx;
        <a id="L248"></a>p[2] = _Tx | byte(rune)&amp;_Maskx;
        <a id="L249"></a>return 3;
    <a id="L250"></a>}

    <a id="L252"></a>p[0] = _T4 | byte(rune&gt;&gt;18);
    <a id="L253"></a>p[1] = _Tx | byte(rune&gt;&gt;12)&amp;_Maskx;
    <a id="L254"></a>p[2] = _Tx | byte(rune&gt;&gt;6)&amp;_Maskx;
    <a id="L255"></a>p[3] = _Tx | byte(rune)&amp;_Maskx;
    <a id="L256"></a>return 4;
<a id="L257"></a>}

<a id="L259"></a><span class="comment">// RuneCount returns the number of runes in p.  Erroneous and short</span>
<a id="L260"></a><span class="comment">// encodings are treated as single runes of width 1 byte.</span>
<a id="L261"></a>func RuneCount(p []byte) int {
    <a id="L262"></a>i := 0;
    <a id="L263"></a>var n int;
    <a id="L264"></a>for n = 0; i &lt; len(p); n++ {
        <a id="L265"></a>if p[i] &lt; RuneSelf {
            <a id="L266"></a>i++
        <a id="L267"></a>} else {
            <a id="L268"></a>_, size := DecodeRune(p[i:len(p)]);
            <a id="L269"></a>i += size;
        <a id="L270"></a>}
    <a id="L271"></a>}
    <a id="L272"></a>return n;
<a id="L273"></a>}

<a id="L275"></a><span class="comment">// RuneCountInString is like RuneCount but its input is a string.</span>
<a id="L276"></a>func RuneCountInString(s string) int {
    <a id="L277"></a>ei := len(s);
    <a id="L278"></a>i := 0;
    <a id="L279"></a>var n int;
    <a id="L280"></a>for n = 0; i &lt; ei; n++ {
        <a id="L281"></a>if s[i] &lt; RuneSelf {
            <a id="L282"></a>i++
        <a id="L283"></a>} else {
            <a id="L284"></a>_, size, _ := decodeRuneInStringInternal(s[i:ei]);
            <a id="L285"></a>i += size;
        <a id="L286"></a>}
    <a id="L287"></a>}
    <a id="L288"></a>return n;
<a id="L289"></a>}

<a id="L291"></a><span class="comment">// RuneStart reports whether the byte could be the first byte of</span>
<a id="L292"></a><span class="comment">// an encoded rune.  Second and subsequent bytes always have the top</span>
<a id="L293"></a><span class="comment">// two bits set to 10.</span>
<a id="L294"></a>func RuneStart(b byte) bool { return b&amp;0xC0 != 0x80 }
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
