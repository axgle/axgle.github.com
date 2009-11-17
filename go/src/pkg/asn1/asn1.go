<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/asn1/asn1.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/asn1/asn1.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The asn1 package implements parsing of DER-encoded ASN.1 data structures,</span>
<a id="L6"></a><span class="comment">// as defined in ITU-T Rec X.690.</span>
<a id="L7"></a><span class="comment">//</span>
<a id="L8"></a><span class="comment">// See also ``A Layman&#39;s Guide to a Subset of ASN.1, BER, and DER,&#39;&#39;</span>
<a id="L9"></a><span class="comment">// http://luca.ntop.org/Teaching/Appunti/asn1.html.</span>
<a id="L10"></a>package asn1

<a id="L12"></a><span class="comment">// ASN.1 is a syntax for specifying abstract objects and BER, DER, PER, XER etc</span>
<a id="L13"></a><span class="comment">// are different encoding formats for those objects. Here, we&#39;ll be dealing</span>
<a id="L14"></a><span class="comment">// with DER, the Distinguished Encoding Rules. DER is used in X.509 because</span>
<a id="L15"></a><span class="comment">// it&#39;s fast to parse and, unlike BER, has a unique encoding for every object.</span>
<a id="L16"></a><span class="comment">// When calculating hashes over objects, it&#39;s important that the resulting</span>
<a id="L17"></a><span class="comment">// bytes be the same at both ends and DER removes this margin of error.</span>
<a id="L18"></a><span class="comment">//</span>
<a id="L19"></a><span class="comment">// ASN.1 is very complex and this package doesn&#39;t attempt to implement</span>
<a id="L20"></a><span class="comment">// everything by any means.</span>

<a id="L22"></a>import (
    <a id="L23"></a>&#34;fmt&#34;;
    <a id="L24"></a>&#34;os&#34;;
    <a id="L25"></a>&#34;reflect&#34;;
    <a id="L26"></a>&#34;strconv&#34;;
    <a id="L27"></a>&#34;strings&#34;;
    <a id="L28"></a>&#34;time&#34;;
<a id="L29"></a>)

<a id="L31"></a><span class="comment">// A StructuralError suggests that the ASN.1 data is valid, but the Go type</span>
<a id="L32"></a><span class="comment">// which is receiving it doesn&#39;t match.</span>
<a id="L33"></a>type StructuralError struct {
    <a id="L34"></a>Msg string;
<a id="L35"></a>}

<a id="L37"></a>func (e StructuralError) String() string { return &#34;ASN.1 structure error: &#34; + e.Msg }

<a id="L39"></a><span class="comment">// A SyntaxError suggests that the ASN.1 data is invalid.</span>
<a id="L40"></a>type SyntaxError struct {
    <a id="L41"></a>Msg string;
<a id="L42"></a>}

<a id="L44"></a>func (e SyntaxError) String() string { return &#34;ASN.1 syntax error: &#34; + e.Msg }

<a id="L46"></a><span class="comment">// We start by dealing with each of the primitive types in turn.</span>

<a id="L48"></a><span class="comment">// BOOLEAN</span>

<a id="L50"></a>func parseBool(bytes []byte) (ret bool, err os.Error) {
    <a id="L51"></a>if len(bytes) != 1 {
        <a id="L52"></a>err = SyntaxError{&#34;invalid boolean&#34;};
        <a id="L53"></a>return;
    <a id="L54"></a>}

    <a id="L56"></a>return bytes[0] != 0, nil;
<a id="L57"></a>}

<a id="L59"></a><span class="comment">// INTEGER</span>

<a id="L61"></a><span class="comment">// parseInt64 treats the given bytes as a big-endian, signed integer and</span>
<a id="L62"></a><span class="comment">// returns the result.</span>
<a id="L63"></a>func parseInt64(bytes []byte) (ret int64, err os.Error) {
    <a id="L64"></a>if len(bytes) &gt; 8 {
        <a id="L65"></a><span class="comment">// We&#39;ll overflow an int64 in this case.</span>
        <a id="L66"></a>err = StructuralError{&#34;integer too large&#34;};
        <a id="L67"></a>return;
    <a id="L68"></a>}
    <a id="L69"></a>for bytesRead := 0; bytesRead &lt; len(bytes); bytesRead++ {
        <a id="L70"></a>ret &lt;&lt;= 8;
        <a id="L71"></a>ret |= int64(bytes[bytesRead]);
    <a id="L72"></a>}

    <a id="L74"></a><span class="comment">// Shift up and down in order to sign extend the result.</span>
    <a id="L75"></a>ret &lt;&lt;= 64 - uint8(len(bytes))*8;
    <a id="L76"></a>ret &gt;&gt;= 64 - uint8(len(bytes))*8;
    <a id="L77"></a>return;
<a id="L78"></a>}

<a id="L80"></a><span class="comment">// parseInt treats the given bytes as a big-endian, signed integer and returns</span>
<a id="L81"></a><span class="comment">// the result.</span>
<a id="L82"></a>func parseInt(bytes []byte) (int, os.Error) {
    <a id="L83"></a>ret64, err := parseInt64(bytes);
    <a id="L84"></a>if err != nil {
        <a id="L85"></a>return 0, err
    <a id="L86"></a>}
    <a id="L87"></a>if ret64 != int64(int(ret64)) {
        <a id="L88"></a>return 0, StructuralError{&#34;integer too large&#34;}
    <a id="L89"></a>}
    <a id="L90"></a>return int(ret64), nil;
<a id="L91"></a>}

<a id="L93"></a><span class="comment">// BIT STRING</span>

<a id="L95"></a><span class="comment">// BitString is the structure to use when you want an ASN.1 BIT STRING type. A</span>
<a id="L96"></a><span class="comment">// bit string is padded up to the nearest byte in memory and the number of</span>
<a id="L97"></a><span class="comment">// valid bits is recorded. Padding bits will be zero.</span>
<a id="L98"></a>type BitString struct {
    <a id="L99"></a>Bytes     []byte; <span class="comment">// bits packed into bytes.</span>
    <a id="L100"></a>BitLength int;    <span class="comment">// length in bits.</span>
<a id="L101"></a>}

<a id="L103"></a><span class="comment">// At returns the bit at the given index. If the index is out of range it</span>
<a id="L104"></a><span class="comment">// returns false.</span>
<a id="L105"></a>func (b BitString) At(i int) int {
    <a id="L106"></a>if i &lt; 0 || i &gt;= b.BitLength {
        <a id="L107"></a>return 0
    <a id="L108"></a>}
    <a id="L109"></a>x := i / 8;
    <a id="L110"></a>y := 7 - uint(i%8);
    <a id="L111"></a>return int(b.Bytes[x]&gt;&gt;y) &amp; 1;
<a id="L112"></a>}

<a id="L114"></a><span class="comment">// parseBitString parses an ASN.1 bit string from the given byte array and returns it.</span>
<a id="L115"></a>func parseBitString(bytes []byte) (ret BitString, err os.Error) {
    <a id="L116"></a>if len(bytes) == 0 {
        <a id="L117"></a>err = SyntaxError{&#34;zero length BIT STRING&#34;};
        <a id="L118"></a>return;
    <a id="L119"></a>}
    <a id="L120"></a>paddingBits := int(bytes[0]);
    <a id="L121"></a>if paddingBits &gt; 7 ||
        <a id="L122"></a>len(bytes) == 1 &amp;&amp; paddingBits &gt; 0 ||
        <a id="L123"></a>bytes[len(bytes)-1]&amp;((1&lt;&lt;bytes[0])-1) != 0 {
        <a id="L124"></a>err = SyntaxError{&#34;invalid padding bits in BIT STRING&#34;};
        <a id="L125"></a>return;
    <a id="L126"></a>}
    <a id="L127"></a>ret.BitLength = (len(bytes)-1)*8 - paddingBits;
    <a id="L128"></a>ret.Bytes = bytes[1:len(bytes)];
    <a id="L129"></a>return;
<a id="L130"></a>}

<a id="L132"></a><span class="comment">// OBJECT IDENTIFIER</span>

<a id="L134"></a><span class="comment">// An ObjectIdentifier represents an ASN.1 OBJECT IDENTIFIER.</span>
<a id="L135"></a>type ObjectIdentifier []int

<a id="L137"></a><span class="comment">// parseObjectIdentifier parses an OBJECT IDENTIFER from the given bytes and</span>
<a id="L138"></a><span class="comment">// returns it. An object identifer is a sequence of variable length integers</span>
<a id="L139"></a><span class="comment">// that are assigned in a hierarachy.</span>
<a id="L140"></a>func parseObjectIdentifier(bytes []byte) (s []int, err os.Error) {
    <a id="L141"></a>if len(bytes) == 0 {
        <a id="L142"></a>err = SyntaxError{&#34;zero length OBJECT IDENTIFIER&#34;};
        <a id="L143"></a>return;
    <a id="L144"></a>}

    <a id="L146"></a><span class="comment">// In the worst case, we get two elements from the first byte (which is</span>
    <a id="L147"></a><span class="comment">// encoded differently) and then every varint is a single byte long.</span>
    <a id="L148"></a>s = make([]int, len(bytes)+1);

    <a id="L150"></a><span class="comment">// The first byte is 40*value1 + value2:</span>
    <a id="L151"></a>s[0] = int(bytes[0]) / 40;
    <a id="L152"></a>s[1] = int(bytes[0]) % 40;
    <a id="L153"></a>i := 2;
    <a id="L154"></a>for offset := 1; offset &lt; len(bytes); i++ {
        <a id="L155"></a>var v int;
        <a id="L156"></a>v, offset, err = parseBase128Int(bytes, offset);
        <a id="L157"></a>if err != nil {
            <a id="L158"></a>return
        <a id="L159"></a>}
        <a id="L160"></a>s[i] = v;
    <a id="L161"></a>}
    <a id="L162"></a>s = s[0:i];
    <a id="L163"></a>return;
<a id="L164"></a>}

<a id="L166"></a><span class="comment">// parseBase128Int parses a base-128 encoded int from the given offset in the</span>
<a id="L167"></a><span class="comment">// given byte array. It returns the value and the new offset.</span>
<a id="L168"></a>func parseBase128Int(bytes []byte, initOffset int) (ret, offset int, err os.Error) {
    <a id="L169"></a>offset = initOffset;
    <a id="L170"></a>for shifted := 0; offset &lt; len(bytes); shifted++ {
        <a id="L171"></a>if shifted &gt; 4 {
            <a id="L172"></a>err = StructuralError{&#34;base 128 integer too large&#34;};
            <a id="L173"></a>return;
        <a id="L174"></a>}
        <a id="L175"></a>ret &lt;&lt;= 7;
        <a id="L176"></a>b := bytes[offset];
        <a id="L177"></a>ret |= int(b &amp; 0x7f);
        <a id="L178"></a>offset++;
        <a id="L179"></a>if b&amp;0x80 == 0 {
            <a id="L180"></a>return
        <a id="L181"></a>}
    <a id="L182"></a>}
    <a id="L183"></a>err = SyntaxError{&#34;truncated base 128 integer&#34;};
    <a id="L184"></a>return;
<a id="L185"></a>}

<a id="L187"></a><span class="comment">// UTCTime</span>

<a id="L189"></a>func isDigit(b byte) bool { return &#39;0&#39; &lt;= b &amp;&amp; b &lt;= &#39;9&#39; }

<a id="L191"></a><span class="comment">// twoDigits returns the value of two, base 10 digits.</span>
<a id="L192"></a>func twoDigits(bytes []byte, max int) (int, bool) {
    <a id="L193"></a>for i := 0; i &lt; 2; i++ {
        <a id="L194"></a>if !isDigit(bytes[i]) {
            <a id="L195"></a>return 0, false
        <a id="L196"></a>}
    <a id="L197"></a>}
    <a id="L198"></a>value := (int(bytes[0])-&#39;0&#39;)*10 + int(bytes[1]-&#39;0&#39;);
    <a id="L199"></a>if value &gt; max {
        <a id="L200"></a>return 0, false
    <a id="L201"></a>}
    <a id="L202"></a>return value, true;
<a id="L203"></a>}

<a id="L205"></a><span class="comment">// parseUTCTime parses the UTCTime from the given byte array and returns the</span>
<a id="L206"></a><span class="comment">// resulting time.</span>
<a id="L207"></a>func parseUTCTime(bytes []byte) (ret time.Time, err os.Error) {
    <a id="L208"></a><span class="comment">// A UTCTime can take the following formats:</span>
    <a id="L209"></a><span class="comment">//</span>
    <a id="L210"></a><span class="comment">//             1111111</span>
    <a id="L211"></a><span class="comment">//   01234567890123456</span>
    <a id="L212"></a><span class="comment">//</span>
    <a id="L213"></a><span class="comment">//   YYMMDDhhmmZ</span>
    <a id="L214"></a><span class="comment">//   YYMMDDhhmm+hhmm</span>
    <a id="L215"></a><span class="comment">//   YYMMDDhhmm-hhmm</span>
    <a id="L216"></a><span class="comment">//   YYMMDDhhmmssZ</span>
    <a id="L217"></a><span class="comment">//   YYMMDDhhmmss+hhmm</span>
    <a id="L218"></a><span class="comment">//   YYMMDDhhmmss-hhmm</span>
    <a id="L219"></a>if len(bytes) &lt; 11 {
        <a id="L220"></a>err = SyntaxError{&#34;UTCTime too short&#34;};
        <a id="L221"></a>return;
    <a id="L222"></a>}
    <a id="L223"></a>var ok1, ok2, ok3, ok4, ok5 bool;
    <a id="L224"></a>year, ok1 := twoDigits(bytes[0:2], 99);
    <a id="L225"></a><span class="comment">// RFC 5280, section 5.1.2.4 says that years 2050 or later use another date</span>
    <a id="L226"></a><span class="comment">// scheme.</span>
    <a id="L227"></a>if year &gt; 50 {
        <a id="L228"></a>ret.Year = 1900 + int64(year)
    <a id="L229"></a>} else {
        <a id="L230"></a>ret.Year = 2000 + int64(year)
    <a id="L231"></a>}
    <a id="L232"></a>ret.Month, ok2 = twoDigits(bytes[2:4], 12);
    <a id="L233"></a>ret.Day, ok3 = twoDigits(bytes[4:6], 31);
    <a id="L234"></a>ret.Hour, ok4 = twoDigits(bytes[6:8], 23);
    <a id="L235"></a>ret.Minute, ok5 = twoDigits(bytes[8:10], 59);
    <a id="L236"></a>if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 {
        <a id="L237"></a>goto Error
    <a id="L238"></a>}
    <a id="L239"></a>bytes = bytes[10:len(bytes)];
    <a id="L240"></a>switch bytes[0] {
    <a id="L241"></a>case &#39;0&#39;, &#39;1&#39;, &#39;2&#39;, &#39;3&#39;, &#39;4&#39;, &#39;5&#39;, &#39;6&#39;:
        <a id="L242"></a>if len(bytes) &lt; 3 {
            <a id="L243"></a>goto Error
        <a id="L244"></a>}
        <a id="L245"></a>ret.Second, ok1 = twoDigits(bytes[0:2], 60); <span class="comment">// 60, not 59, because of leap seconds.</span>
        <a id="L246"></a>if !ok1 {
            <a id="L247"></a>goto Error
        <a id="L248"></a>}
        <a id="L249"></a>bytes = bytes[2:len(bytes)];
    <a id="L250"></a>}
    <a id="L251"></a>if len(bytes) == 0 {
        <a id="L252"></a>goto Error
    <a id="L253"></a>}
    <a id="L254"></a>switch bytes[0] {
    <a id="L255"></a>case &#39;Z&#39;:
        <a id="L256"></a>if len(bytes) != 1 {
            <a id="L257"></a>goto Error
        <a id="L258"></a>}
        <a id="L259"></a>return;
    <a id="L260"></a>case &#39;-&#39;, &#39;+&#39;:
        <a id="L261"></a>if len(bytes) != 5 {
            <a id="L262"></a>goto Error
        <a id="L263"></a>}
        <a id="L264"></a>hours, ok1 := twoDigits(bytes[1:3], 12);
        <a id="L265"></a>minutes, ok2 := twoDigits(bytes[3:5], 59);
        <a id="L266"></a>if !ok1 || !ok2 {
            <a id="L267"></a>goto Error
        <a id="L268"></a>}
        <a id="L269"></a>sign := 1;
        <a id="L270"></a>if bytes[0] == &#39;-&#39; {
            <a id="L271"></a>sign = -1
        <a id="L272"></a>}
        <a id="L273"></a>ret.ZoneOffset = sign * (60 * (hours*60 + minutes));
    <a id="L274"></a>default:
        <a id="L275"></a>goto Error
    <a id="L276"></a>}
    <a id="L277"></a>return;

<a id="L279"></a>Error:
    <a id="L280"></a>err = SyntaxError{&#34;invalid UTCTime&#34;};
    <a id="L281"></a>return;
<a id="L282"></a>}

<a id="L284"></a><span class="comment">// PrintableString</span>

<a id="L286"></a><span class="comment">// parsePrintableString parses a ASN.1 PrintableString from the given byte</span>
<a id="L287"></a><span class="comment">// array and returns it.</span>
<a id="L288"></a>func parsePrintableString(bytes []byte) (ret string, err os.Error) {
    <a id="L289"></a>for _, b := range bytes {
        <a id="L290"></a>if !isPrintable(b) {
            <a id="L291"></a>err = SyntaxError{&#34;PrintableString contains invalid character&#34;};
            <a id="L292"></a>return;
        <a id="L293"></a>}
    <a id="L294"></a>}
    <a id="L295"></a>ret = string(bytes);
    <a id="L296"></a>return;
<a id="L297"></a>}

<a id="L299"></a><span class="comment">// isPrintable returns true iff the given b is in the ASN.1 PrintableString set.</span>
<a id="L300"></a>func isPrintable(b byte) bool {
    <a id="L301"></a>return &#39;a&#39; &lt;= b &amp;&amp; b &lt;= &#39;z&#39; ||
        <a id="L302"></a>&#39;A&#39; &lt;= b &amp;&amp; b &lt;= &#39;Z&#39; ||
        <a id="L303"></a>&#39;0&#39; &lt;= b &amp;&amp; b &lt;= &#39;9&#39; ||
        <a id="L304"></a>&#39;\&#39;&#39; &lt;= b &amp;&amp; b &lt;= &#39;)&#39; ||
        <a id="L305"></a>&#39;+&#39; &lt;= b &amp;&amp; b &lt;= &#39;/&#39; ||
        <a id="L306"></a>b == &#39; &#39; ||
        <a id="L307"></a>b == &#39;:&#39; ||
        <a id="L308"></a>b == &#39;=&#39; ||
        <a id="L309"></a>b == &#39;?&#39;
<a id="L310"></a>}

<a id="L312"></a><span class="comment">// IA5String</span>

<a id="L314"></a><span class="comment">// parseIA5String parses a ASN.1 IA5String (ASCII string) from the given</span>
<a id="L315"></a><span class="comment">// byte array and returns it.</span>
<a id="L316"></a>func parseIA5String(bytes []byte) (ret string, err os.Error) {
    <a id="L317"></a>for _, b := range bytes {
        <a id="L318"></a>if b &gt;= 0x80 {
            <a id="L319"></a>err = SyntaxError{&#34;IA5String contains invalid character&#34;};
            <a id="L320"></a>return;
        <a id="L321"></a>}
    <a id="L322"></a>}
    <a id="L323"></a>ret = string(bytes);
    <a id="L324"></a>return;
<a id="L325"></a>}

<a id="L327"></a><span class="comment">// A RawValue represents an undecoded ASN.1 object.</span>
<a id="L328"></a>type RawValue struct {
    <a id="L329"></a>Class, Tag int;
    <a id="L330"></a>IsCompound bool;
    <a id="L331"></a>Bytes      []byte;
<a id="L332"></a>}

<a id="L334"></a><span class="comment">// Tagging</span>

<a id="L336"></a><span class="comment">// ASN.1 objects have metadata preceeding them:</span>
<a id="L337"></a><span class="comment">//   the tag: the type of the object</span>
<a id="L338"></a><span class="comment">//   a flag denoting if this object is compound or not</span>
<a id="L339"></a><span class="comment">//   the class type: the namespace of the tag</span>
<a id="L340"></a><span class="comment">//   the length of the object, in bytes</span>

<a id="L342"></a><span class="comment">// Here are some standard tags and classes</span>

<a id="L344"></a>const (
    <a id="L345"></a>tagBoolean         = 1;
    <a id="L346"></a>tagInteger         = 2;
    <a id="L347"></a>tagBitString       = 3;
    <a id="L348"></a>tagOctetString     = 4;
    <a id="L349"></a>tagOID             = 6;
    <a id="L350"></a>tagSequence        = 16;
    <a id="L351"></a>tagSet             = 17;
    <a id="L352"></a>tagPrintableString = 19;
    <a id="L353"></a>tagIA5String       = 22;
    <a id="L354"></a>tagUTCTime         = 23;
<a id="L355"></a>)

<a id="L357"></a>const (
    <a id="L358"></a>classUniversal       = 0;
    <a id="L359"></a>classApplication     = 1;
    <a id="L360"></a>classContextSpecific = 2;
    <a id="L361"></a>classPrivate         = 3;
<a id="L362"></a>)

<a id="L364"></a>type tagAndLength struct {
    <a id="L365"></a>class, tag, length int;
    <a id="L366"></a>isCompound         bool;
<a id="L367"></a>}

<a id="L369"></a><span class="comment">// parseTagAndLength parses an ASN.1 tag and length pair from the given offset</span>
<a id="L370"></a><span class="comment">// into a byte array. It returns the parsed data and the new offset. SET and</span>
<a id="L371"></a><span class="comment">// SET OF (tag 17) are mapped to SEQUENCE and SEQUENCE OF (tag 16) since we</span>
<a id="L372"></a><span class="comment">// don&#39;t distinguish between ordered and unordered objects in this code.</span>
<a id="L373"></a>func parseTagAndLength(bytes []byte, initOffset int) (ret tagAndLength, offset int, err os.Error) {
    <a id="L374"></a>offset = initOffset;
    <a id="L375"></a>b := bytes[offset];
    <a id="L376"></a>offset++;
    <a id="L377"></a>ret.class = int(b &gt;&gt; 6);
    <a id="L378"></a>ret.isCompound = b&amp;0x20 == 0x20;
    <a id="L379"></a>ret.tag = int(b &amp; 0x1f);

    <a id="L381"></a><span class="comment">// If the bottom five bits are set, then the tag number is actually base 128</span>
    <a id="L382"></a><span class="comment">// encoded afterwards</span>
    <a id="L383"></a>if ret.tag == 0x1f {
        <a id="L384"></a>ret.tag, offset, err = parseBase128Int(bytes, offset);
        <a id="L385"></a>if err != nil {
            <a id="L386"></a>return
        <a id="L387"></a>}
    <a id="L388"></a>}
    <a id="L389"></a>if offset &gt;= len(bytes) {
        <a id="L390"></a>err = SyntaxError{&#34;truncated tag or length&#34;};
        <a id="L391"></a>return;
    <a id="L392"></a>}
    <a id="L393"></a>b = bytes[offset];
    <a id="L394"></a>offset++;
    <a id="L395"></a>if b&amp;0x80 == 0 {
        <a id="L396"></a><span class="comment">// The length is encoded in the bottom 7 bits.</span>
        <a id="L397"></a>ret.length = int(b &amp; 0x7f)
    <a id="L398"></a>} else {
        <a id="L399"></a><span class="comment">// Bottom 7 bits give the number of length bytes to follow.</span>
        <a id="L400"></a>numBytes := int(b &amp; 0x7f);
        <a id="L401"></a><span class="comment">// We risk overflowing a signed 32-bit number if we accept more than 3 bytes.</span>
        <a id="L402"></a>if numBytes &gt; 3 {
            <a id="L403"></a>err = StructuralError{&#34;length too large&#34;};
            <a id="L404"></a>return;
        <a id="L405"></a>}
        <a id="L406"></a>if numBytes == 0 {
            <a id="L407"></a>err = SyntaxError{&#34;indefinite length found (not DER)&#34;};
            <a id="L408"></a>return;
        <a id="L409"></a>}
        <a id="L410"></a>ret.length = 0;
        <a id="L411"></a>for i := 0; i &lt; numBytes; i++ {
            <a id="L412"></a>if offset &gt;= len(bytes) {
                <a id="L413"></a>err = SyntaxError{&#34;truncated tag or length&#34;};
                <a id="L414"></a>return;
            <a id="L415"></a>}
            <a id="L416"></a>b = bytes[offset];
            <a id="L417"></a>offset++;
            <a id="L418"></a>ret.length &lt;&lt;= 8;
            <a id="L419"></a>ret.length |= int(b);
        <a id="L420"></a>}
    <a id="L421"></a>}

    <a id="L423"></a><span class="comment">// We magically map SET and SET OF to SEQUENCE and SEQUENCE OF</span>
    <a id="L424"></a><span class="comment">// because we treat everything as ordered.</span>
    <a id="L425"></a>if ret.tag == tagSet {
        <a id="L426"></a>ret.tag = tagSequence
    <a id="L427"></a>}
    <a id="L428"></a>return;
<a id="L429"></a>}

<a id="L431"></a><span class="comment">// ASN.1 has IMPLICIT and EXPLICIT tags, which can be translated as &#34;instead</span>
<a id="L432"></a><span class="comment">// of&#34; and &#34;in addition to&#34;. When not specified, every primitive type has a</span>
<a id="L433"></a><span class="comment">// default tag in the UNIVERSAL class.</span>
<a id="L434"></a><span class="comment">//</span>
<a id="L435"></a><span class="comment">// For example: a BIT STRING is tagged [UNIVERSAL 3] by default (although ASN.1</span>
<a id="L436"></a><span class="comment">// doesn&#39;t actually have a UNIVERSAL keyword). However, by saying [IMPLICIT</span>
<a id="L437"></a><span class="comment">// CONTEXT-SPECIFIC 42], that means that the tag is replaced by another.</span>
<a id="L438"></a><span class="comment">//</span>
<a id="L439"></a><span class="comment">// On the other hand, if it said [EXPLICIT CONTEXT-SPECIFIC 10], then an</span>
<a id="L440"></a><span class="comment">// /additional/ tag would wrap the default tag. This explicit tag will have the</span>
<a id="L441"></a><span class="comment">// compound flag set.</span>
<a id="L442"></a><span class="comment">//</span>
<a id="L443"></a><span class="comment">// (This is used in order to remove ambiguity with optional elements.)</span>
<a id="L444"></a><span class="comment">//</span>
<a id="L445"></a><span class="comment">// You can layer EXPLICIT and IMPLICIT tags to an arbitrary depth, however we</span>
<a id="L446"></a><span class="comment">// don&#39;t support that here. We support a single layer of EXPLICIT or IMPLICIT</span>
<a id="L447"></a><span class="comment">// tagging with tag strings on the fields of a structure.</span>

<a id="L449"></a><span class="comment">// fieldParameters is the parsed representation of tag string from a structure field.</span>
<a id="L450"></a>type fieldParameters struct {
    <a id="L451"></a>optional     bool;   <span class="comment">// true iff the field is OPTIONAL</span>
    <a id="L452"></a>explicit     bool;   <span class="comment">// true iff and EXPLICIT tag is in use.</span>
    <a id="L453"></a>defaultValue *int64; <span class="comment">// a default value for INTEGER typed fields (maybe nil).</span>
    <a id="L454"></a>tag          *int;   <span class="comment">// the EXPLICIT or IMPLICIT tag (maybe nil).</span>

    <a id="L456"></a><span class="comment">// Invariants:</span>
    <a id="L457"></a><span class="comment">//   if explicit is set, tag is non-nil.</span>
<a id="L458"></a>}

<a id="L460"></a><span class="comment">// Given a tag string with the format specified in the package comment,</span>
<a id="L461"></a><span class="comment">// parseFieldParameters will parse it into a fieldParameters structure,</span>
<a id="L462"></a><span class="comment">// ignoring unknown parts of the string.</span>
<a id="L463"></a>func parseFieldParameters(str string) (ret fieldParameters) {
    <a id="L464"></a>for _, part := range strings.Split(str, &#34;,&#34;, 0) {
        <a id="L465"></a>switch {
        <a id="L466"></a>case part == &#34;optional&#34;:
            <a id="L467"></a>ret.optional = true
        <a id="L468"></a>case part == &#34;explicit&#34;:
            <a id="L469"></a>ret.explicit = true;
            <a id="L470"></a>if ret.tag == nil {
                <a id="L471"></a>ret.tag = new(int);
                <a id="L472"></a>*ret.tag = 0;
            <a id="L473"></a>}
        <a id="L474"></a>case strings.HasPrefix(part, &#34;default:&#34;):
            <a id="L475"></a>i, err := strconv.Atoi64(part[8:len(part)]);
            <a id="L476"></a>if err == nil {
                <a id="L477"></a>ret.defaultValue = new(int64);
                <a id="L478"></a>*ret.defaultValue = i;
            <a id="L479"></a>}
        <a id="L480"></a>case strings.HasPrefix(part, &#34;tag:&#34;):
            <a id="L481"></a>i, err := strconv.Atoi(part[4:len(part)]);
            <a id="L482"></a>if err == nil {
                <a id="L483"></a>ret.tag = new(int);
                <a id="L484"></a>*ret.tag = i;
            <a id="L485"></a>}
        <a id="L486"></a>}
    <a id="L487"></a>}
    <a id="L488"></a>return;
<a id="L489"></a>}

<a id="L491"></a><span class="comment">// Given a reflected Go type, getUniversalType returns the default tag number</span>
<a id="L492"></a><span class="comment">// and expected compound flag.</span>
<a id="L493"></a>func getUniversalType(t reflect.Type) (tagNumber int, isCompound, ok bool) {
    <a id="L494"></a>switch t {
    <a id="L495"></a>case objectIdentifierType:
        <a id="L496"></a>return tagOID, false, true
    <a id="L497"></a>case bitStringType:
        <a id="L498"></a>return tagBitString, false, true
    <a id="L499"></a>case timeType:
        <a id="L500"></a>return tagUTCTime, false, true
    <a id="L501"></a>}
    <a id="L502"></a>switch i := t.(type) {
    <a id="L503"></a>case *reflect.BoolType:
        <a id="L504"></a>return tagBoolean, false, true
    <a id="L505"></a>case *reflect.IntType:
        <a id="L506"></a>return tagInteger, false, true
    <a id="L507"></a>case *reflect.Int64Type:
        <a id="L508"></a>return tagInteger, false, true
    <a id="L509"></a>case *reflect.StructType:
        <a id="L510"></a>return tagSequence, true, true
    <a id="L511"></a>case *reflect.SliceType:
        <a id="L512"></a>if _, ok := t.(*reflect.SliceType).Elem().(*reflect.Uint8Type); ok {
            <a id="L513"></a>return tagOctetString, false, true
        <a id="L514"></a>}
        <a id="L515"></a>return tagSequence, true, true;
    <a id="L516"></a>case *reflect.StringType:
        <a id="L517"></a>return tagPrintableString, false, true
    <a id="L518"></a>}
    <a id="L519"></a>return 0, false, false;
<a id="L520"></a>}

<a id="L522"></a><span class="comment">// parseSequenceOf is used for SEQUENCE OF and SET OF values. It tries to parse</span>
<a id="L523"></a><span class="comment">// a number of ASN.1 values from the given byte array and returns them as a</span>
<a id="L524"></a><span class="comment">// slice of Go values of the given type.</span>
<a id="L525"></a>func parseSequenceOf(bytes []byte, sliceType *reflect.SliceType, elemType reflect.Type) (ret *reflect.SliceValue, err os.Error) {
    <a id="L526"></a>expectedTag, compoundType, ok := getUniversalType(elemType);
    <a id="L527"></a>if !ok {
        <a id="L528"></a>err = StructuralError{&#34;unknown Go type for slice&#34;};
        <a id="L529"></a>return;
    <a id="L530"></a>}

    <a id="L532"></a><span class="comment">// First we iterate over the input and count the number of elements,</span>
    <a id="L533"></a><span class="comment">// checking that the types are correct in each case.</span>
    <a id="L534"></a>numElements := 0;
    <a id="L535"></a>for offset := 0; offset &lt; len(bytes); {
        <a id="L536"></a>var t tagAndLength;
        <a id="L537"></a>t, offset, err = parseTagAndLength(bytes, offset);
        <a id="L538"></a>if err != nil {
            <a id="L539"></a>return
        <a id="L540"></a>}
        <a id="L541"></a>if t.class != classUniversal || t.isCompound != compoundType || t.tag != expectedTag {
            <a id="L542"></a>err = StructuralError{&#34;sequence tag mismatch&#34;};
            <a id="L543"></a>return;
        <a id="L544"></a>}
        <a id="L545"></a>if invalidLength(offset, t.length, len(bytes)) {
            <a id="L546"></a>err = SyntaxError{&#34;truncated sequence&#34;};
            <a id="L547"></a>return;
        <a id="L548"></a>}
        <a id="L549"></a>offset += t.length;
        <a id="L550"></a>numElements++;
    <a id="L551"></a>}
    <a id="L552"></a>ret = reflect.MakeSlice(sliceType, numElements, numElements);
    <a id="L553"></a>params := fieldParameters{};
    <a id="L554"></a>offset := 0;
    <a id="L555"></a>for i := 0; i &lt; numElements; i++ {
        <a id="L556"></a>offset, err = parseField(ret.Elem(i), bytes, offset, params);
        <a id="L557"></a>if err != nil {
            <a id="L558"></a>return
        <a id="L559"></a>}
    <a id="L560"></a>}
    <a id="L561"></a>return;
<a id="L562"></a>}

<a id="L564"></a>var (
    <a id="L565"></a>bitStringType        = reflect.Typeof(BitString{});
    <a id="L566"></a>objectIdentifierType = reflect.Typeof(ObjectIdentifier{});
    <a id="L567"></a>timeType             = reflect.Typeof(time.Time{});
    <a id="L568"></a>rawValueType         = reflect.Typeof(RawValue{});
<a id="L569"></a>)

<a id="L571"></a><span class="comment">// invalidLength returns true iff offset + length &gt; sliceLength, or if the</span>
<a id="L572"></a><span class="comment">// addition would overflow.</span>
<a id="L573"></a>func invalidLength(offset, length, sliceLength int) bool {
    <a id="L574"></a>return offset+length &lt; offset || offset+length &gt; sliceLength
<a id="L575"></a>}

<a id="L577"></a><span class="comment">// parseField is the main parsing function. Given a byte array and an offset</span>
<a id="L578"></a><span class="comment">// into the array, it will try to parse a suitable ASN.1 value out and store it</span>
<a id="L579"></a><span class="comment">// in the given Value.</span>
<a id="L580"></a>func parseField(v reflect.Value, bytes []byte, initOffset int, params fieldParameters) (offset int, err os.Error) {
    <a id="L581"></a>offset = initOffset;
    <a id="L582"></a>fieldType := v.Type();

    <a id="L584"></a><span class="comment">// If we have run out of data, it may be that there are optional elements at the end.</span>
    <a id="L585"></a>if offset == len(bytes) {
        <a id="L586"></a>if !setDefaultValue(v, params) {
            <a id="L587"></a>err = SyntaxError{&#34;sequence truncated&#34;}
        <a id="L588"></a>}
        <a id="L589"></a>return;
    <a id="L590"></a>}

    <a id="L592"></a><span class="comment">// Deal with raw values.</span>
    <a id="L593"></a>if fieldType == rawValueType {
        <a id="L594"></a>var t tagAndLength;
        <a id="L595"></a>t, offset, err = parseTagAndLength(bytes, offset);
        <a id="L596"></a>if err != nil {
            <a id="L597"></a>return
        <a id="L598"></a>}
        <a id="L599"></a>if invalidLength(offset, t.length, len(bytes)) {
            <a id="L600"></a>err = SyntaxError{&#34;data truncated&#34;};
            <a id="L601"></a>return;
        <a id="L602"></a>}
        <a id="L603"></a>result := RawValue{t.class, t.tag, t.isCompound, bytes[offset : offset+t.length]};
        <a id="L604"></a>offset += t.length;
        <a id="L605"></a>v.(*reflect.StructValue).Set(reflect.NewValue(result).(*reflect.StructValue));
        <a id="L606"></a>return;
    <a id="L607"></a>}

    <a id="L609"></a><span class="comment">// Deal with the ANY type.</span>
    <a id="L610"></a>if ifaceType, ok := fieldType.(*reflect.InterfaceType); ok &amp;&amp; ifaceType.NumMethod() == 0 {
        <a id="L611"></a>ifaceValue := v.(*reflect.InterfaceValue);
        <a id="L612"></a>var t tagAndLength;
        <a id="L613"></a>t, offset, err = parseTagAndLength(bytes, offset);
        <a id="L614"></a>if err != nil {
            <a id="L615"></a>return
        <a id="L616"></a>}
        <a id="L617"></a>if invalidLength(offset, t.length, len(bytes)) {
            <a id="L618"></a>err = SyntaxError{&#34;data truncated&#34;};
            <a id="L619"></a>return;
        <a id="L620"></a>}
        <a id="L621"></a>var result interface{}
        <a id="L622"></a>if !t.isCompound &amp;&amp; t.class == classUniversal {
            <a id="L623"></a>innerBytes := bytes[offset : offset+t.length];
            <a id="L624"></a>switch t.tag {
            <a id="L625"></a>case tagPrintableString:
                <a id="L626"></a>result, err = parsePrintableString(innerBytes)
            <a id="L627"></a>case tagIA5String:
                <a id="L628"></a>result, err = parseIA5String(innerBytes)
            <a id="L629"></a>case tagInteger:
                <a id="L630"></a>result, err = parseInt64(innerBytes)
            <a id="L631"></a>case tagBitString:
                <a id="L632"></a>result, err = parseBitString(innerBytes)
            <a id="L633"></a>case tagOID:
                <a id="L634"></a>result, err = parseObjectIdentifier(innerBytes)
            <a id="L635"></a>case tagUTCTime:
                <a id="L636"></a>result, err = parseUTCTime(innerBytes)
            <a id="L637"></a>case tagOctetString:
                <a id="L638"></a>result = innerBytes
            <a id="L639"></a>default:
                <a id="L640"></a><span class="comment">// If we don&#39;t know how to handle the type, we just leave Value as nil.</span>
            <a id="L641"></a>}
        <a id="L642"></a>}
        <a id="L643"></a>offset += t.length;
        <a id="L644"></a>if err != nil {
            <a id="L645"></a>return
        <a id="L646"></a>}
        <a id="L647"></a>if result != nil {
            <a id="L648"></a>ifaceValue.Set(reflect.NewValue(result))
        <a id="L649"></a>}
        <a id="L650"></a>return;
    <a id="L651"></a>}
    <a id="L652"></a>universalTag, compoundType, ok1 := getUniversalType(fieldType);
    <a id="L653"></a>if !ok1 {
        <a id="L654"></a>err = StructuralError{fmt.Sprintf(&#34;unknown Go type: %v&#34;, fieldType)};
        <a id="L655"></a>return;
    <a id="L656"></a>}

    <a id="L658"></a>t, offset, err := parseTagAndLength(bytes, offset);
    <a id="L659"></a>if err != nil {
        <a id="L660"></a>return
    <a id="L661"></a>}
    <a id="L662"></a>if params.explicit {
        <a id="L663"></a>if t.class == classContextSpecific &amp;&amp; t.tag == *params.tag &amp;&amp; t.isCompound {
            <a id="L664"></a>t, offset, err = parseTagAndLength(bytes, offset);
            <a id="L665"></a>if err != nil {
                <a id="L666"></a>return
            <a id="L667"></a>}
        <a id="L668"></a>} else {
            <a id="L669"></a><span class="comment">// The tags didn&#39;t match, it might be an optional element.</span>
            <a id="L670"></a>ok := setDefaultValue(v, params);
            <a id="L671"></a>if ok {
                <a id="L672"></a>offset = initOffset
            <a id="L673"></a>} else {
                <a id="L674"></a>err = StructuralError{&#34;explicitly tagged member didn&#39;t match&#34;}
            <a id="L675"></a>}
            <a id="L676"></a>return;
        <a id="L677"></a>}
    <a id="L678"></a>}

    <a id="L680"></a><span class="comment">// Special case for strings: PrintableString and IA5String both map to</span>
    <a id="L681"></a><span class="comment">// the Go type string. getUniversalType returns the tag for</span>
    <a id="L682"></a><span class="comment">// PrintableString when it sees a string so, if we see an IA5String on</span>
    <a id="L683"></a><span class="comment">// the wire, we change the universal type to match.</span>
    <a id="L684"></a>if universalTag == tagPrintableString &amp;&amp; t.tag == tagIA5String {
        <a id="L685"></a>universalTag = tagIA5String
    <a id="L686"></a>}

    <a id="L688"></a>expectedClass := classUniversal;
    <a id="L689"></a>expectedTag := universalTag;

    <a id="L691"></a>if !params.explicit &amp;&amp; params.tag != nil {
        <a id="L692"></a>expectedClass = classContextSpecific;
        <a id="L693"></a>expectedTag = *params.tag;
    <a id="L694"></a>}

    <a id="L696"></a><span class="comment">// We have unwrapped any explicit tagging at this point.</span>
    <a id="L697"></a>if t.class != expectedClass || t.tag != expectedTag || t.isCompound != compoundType {
        <a id="L698"></a><span class="comment">// Tags don&#39;t match. Again, it could be an optional element.</span>
        <a id="L699"></a>ok := setDefaultValue(v, params);
        <a id="L700"></a>if ok {
            <a id="L701"></a>offset = initOffset
        <a id="L702"></a>} else {
            <a id="L703"></a>err = StructuralError{fmt.Sprintf(&#34;tags don&#39;t match (%d vs %+v) %+v %s %#v&#34;, expectedTag, t, params, fieldType.Name(), bytes[offset:len(bytes)])}
        <a id="L704"></a>}
        <a id="L705"></a>return;
    <a id="L706"></a>}
    <a id="L707"></a>if invalidLength(offset, t.length, len(bytes)) {
        <a id="L708"></a>err = SyntaxError{&#34;data truncated&#34;};
        <a id="L709"></a>return;
    <a id="L710"></a>}
    <a id="L711"></a>innerBytes := bytes[offset : offset+t.length];

    <a id="L713"></a><span class="comment">// We deal with the structures defined in this package first.</span>
    <a id="L714"></a>switch fieldType {
    <a id="L715"></a>case objectIdentifierType:
        <a id="L716"></a>newSlice, err1 := parseObjectIdentifier(innerBytes);
        <a id="L717"></a>sliceValue := v.(*reflect.SliceValue);
        <a id="L718"></a>sliceValue.Set(reflect.MakeSlice(sliceValue.Type().(*reflect.SliceType), len(newSlice), len(newSlice)));
        <a id="L719"></a>if err1 == nil {
            <a id="L720"></a>reflect.ArrayCopy(sliceValue, reflect.NewValue(newSlice).(reflect.ArrayOrSliceValue))
        <a id="L721"></a>}
        <a id="L722"></a>offset += t.length;
        <a id="L723"></a>err = err1;
        <a id="L724"></a>return;
    <a id="L725"></a>case bitStringType:
        <a id="L726"></a>structValue := v.(*reflect.StructValue);
        <a id="L727"></a>bs, err1 := parseBitString(innerBytes);
        <a id="L728"></a>offset += t.length;
        <a id="L729"></a>if err1 == nil {
            <a id="L730"></a>structValue.Set(reflect.NewValue(bs).(*reflect.StructValue))
        <a id="L731"></a>}
        <a id="L732"></a>err = err1;
        <a id="L733"></a>return;
    <a id="L734"></a>case timeType:
        <a id="L735"></a>structValue := v.(*reflect.StructValue);
        <a id="L736"></a>time, err1 := parseUTCTime(innerBytes);
        <a id="L737"></a>offset += t.length;
        <a id="L738"></a>if err1 == nil {
            <a id="L739"></a>structValue.Set(reflect.NewValue(time).(*reflect.StructValue))
        <a id="L740"></a>}
        <a id="L741"></a>err = err1;
        <a id="L742"></a>return;
    <a id="L743"></a>}
    <a id="L744"></a>switch val := v.(type) {
    <a id="L745"></a>case *reflect.BoolValue:
        <a id="L746"></a>parsedBool, err1 := parseBool(innerBytes);
        <a id="L747"></a>offset += t.length;
        <a id="L748"></a>if err1 == nil {
            <a id="L749"></a>val.Set(parsedBool)
        <a id="L750"></a>}
        <a id="L751"></a>err = err1;
        <a id="L752"></a>return;
    <a id="L753"></a>case *reflect.IntValue:
        <a id="L754"></a>parsedInt, err1 := parseInt(innerBytes);
        <a id="L755"></a>offset += t.length;
        <a id="L756"></a>if err1 == nil {
            <a id="L757"></a>val.Set(parsedInt)
        <a id="L758"></a>}
        <a id="L759"></a>err = err1;
        <a id="L760"></a>return;
    <a id="L761"></a>case *reflect.Int64Value:
        <a id="L762"></a>parsedInt, err1 := parseInt64(innerBytes);
        <a id="L763"></a>offset += t.length;
        <a id="L764"></a>if err1 == nil {
            <a id="L765"></a>val.Set(parsedInt)
        <a id="L766"></a>}
        <a id="L767"></a>err = err1;
        <a id="L768"></a>return;
    <a id="L769"></a>case *reflect.StructValue:
        <a id="L770"></a>structType := fieldType.(*reflect.StructType);
        <a id="L771"></a>innerOffset := 0;
        <a id="L772"></a>for i := 0; i &lt; structType.NumField(); i++ {
            <a id="L773"></a>field := structType.Field(i);
            <a id="L774"></a>innerOffset, err = parseField(val.Field(i), innerBytes, innerOffset, parseFieldParameters(field.Tag));
            <a id="L775"></a>if err != nil {
                <a id="L776"></a>return
            <a id="L777"></a>}
        <a id="L778"></a>}
        <a id="L779"></a>offset += t.length;
        <a id="L780"></a><span class="comment">// We allow extra bytes at the end of the SEQUENCE because</span>
        <a id="L781"></a><span class="comment">// adding elements to the end has been used in X.509 as the</span>
        <a id="L782"></a><span class="comment">// version numbers have increased.</span>
        <a id="L783"></a>return;
    <a id="L784"></a>case *reflect.SliceValue:
        <a id="L785"></a>sliceType := fieldType.(*reflect.SliceType);
        <a id="L786"></a>if _, ok := sliceType.Elem().(*reflect.Uint8Type); ok {
            <a id="L787"></a>val.Set(reflect.MakeSlice(sliceType, len(innerBytes), len(innerBytes)));
            <a id="L788"></a>reflect.ArrayCopy(val, reflect.NewValue(innerBytes).(reflect.ArrayOrSliceValue));
            <a id="L789"></a>return;
        <a id="L790"></a>}
        <a id="L791"></a>newSlice, err1 := parseSequenceOf(innerBytes, sliceType, sliceType.Elem());
        <a id="L792"></a>offset += t.length;
        <a id="L793"></a>if err1 == nil {
            <a id="L794"></a>val.Set(newSlice)
        <a id="L795"></a>}
        <a id="L796"></a>err = err1;
        <a id="L797"></a>return;
    <a id="L798"></a>case *reflect.StringValue:
        <a id="L799"></a>var v string;
        <a id="L800"></a>switch universalTag {
        <a id="L801"></a>case tagPrintableString:
            <a id="L802"></a>v, err = parsePrintableString(innerBytes)
        <a id="L803"></a>case tagIA5String:
            <a id="L804"></a>v, err = parseIA5String(innerBytes)
        <a id="L805"></a>default:
            <a id="L806"></a>err = SyntaxError{fmt.Sprintf(&#34;internal error: unknown string type %d&#34;, universalTag)}
        <a id="L807"></a>}
        <a id="L808"></a>if err == nil {
            <a id="L809"></a>val.Set(v)
        <a id="L810"></a>}
        <a id="L811"></a>return;
    <a id="L812"></a>}
    <a id="L813"></a>err = StructuralError{&#34;unknown Go type&#34;};
    <a id="L814"></a>return;
<a id="L815"></a>}

<a id="L817"></a><span class="comment">// setDefaultValue is used to install a default value, from a tag string, into</span>
<a id="L818"></a><span class="comment">// a Value. It is successful is the field was optional, even if a default value</span>
<a id="L819"></a><span class="comment">// wasn&#39;t provided or it failed to install it into the Value.</span>
<a id="L820"></a>func setDefaultValue(v reflect.Value, params fieldParameters) (ok bool) {
    <a id="L821"></a>if !params.optional {
        <a id="L822"></a>return
    <a id="L823"></a>}
    <a id="L824"></a>ok = true;
    <a id="L825"></a>if params.defaultValue == nil {
        <a id="L826"></a>return
    <a id="L827"></a>}
    <a id="L828"></a>switch val := v.(type) {
    <a id="L829"></a>case *reflect.IntValue:
        <a id="L830"></a>val.Set(int(*params.defaultValue))
    <a id="L831"></a>case *reflect.Int64Value:
        <a id="L832"></a>val.Set(int64(*params.defaultValue))
    <a id="L833"></a>}
    <a id="L834"></a>return;
<a id="L835"></a>}

<a id="L837"></a><span class="comment">// Unmarshal parses the DER-encoded ASN.1 data structure b</span>
<a id="L838"></a><span class="comment">// and uses the reflect package to fill in an arbitrary value pointed at by val.</span>
<a id="L839"></a><span class="comment">// Because Unmarshal uses the reflect package, the structs</span>
<a id="L840"></a><span class="comment">// being written to must use upper case field names.</span>
<a id="L841"></a><span class="comment">//</span>
<a id="L842"></a><span class="comment">// An ASN.1 INTEGER can be written to an int or int64.</span>
<a id="L843"></a><span class="comment">// If the encoded value does not fit in the Go type,</span>
<a id="L844"></a><span class="comment">// Unmarshal returns a parse error.</span>
<a id="L845"></a><span class="comment">//</span>
<a id="L846"></a><span class="comment">// An ASN.1 BIT STRING can be written to a BitString.</span>
<a id="L847"></a><span class="comment">//</span>
<a id="L848"></a><span class="comment">// An ASN.1 OCTET STRING can be written to a []byte.</span>
<a id="L849"></a><span class="comment">//</span>
<a id="L850"></a><span class="comment">// An ASN.1 OBJECT IDENTIFIER can be written to an</span>
<a id="L851"></a><span class="comment">// ObjectIdentifier.</span>
<a id="L852"></a><span class="comment">//</span>
<a id="L853"></a><span class="comment">// An ASN.1 PrintableString or IA5String can be written to a string.</span>
<a id="L854"></a><span class="comment">//</span>
<a id="L855"></a><span class="comment">// Any of the above ASN.1 values can be written to an interface{}.</span>
<a id="L856"></a><span class="comment">// The value stored in the interface has the corresponding Go type.</span>
<a id="L857"></a><span class="comment">// For integers, that type is int64.</span>
<a id="L858"></a><span class="comment">//</span>
<a id="L859"></a><span class="comment">// An ASN.1 SEQUENCE OF x or SET OF x can be written</span>
<a id="L860"></a><span class="comment">// to a slice if an x can be written to the slice&#39;s element type.</span>
<a id="L861"></a><span class="comment">//</span>
<a id="L862"></a><span class="comment">// An ASN.1 SEQUENCE or SET can be written to a struct</span>
<a id="L863"></a><span class="comment">// if each of the elements in the sequence can be</span>
<a id="L864"></a><span class="comment">// written to the corresponding element in the struct.</span>
<a id="L865"></a><span class="comment">//</span>
<a id="L866"></a><span class="comment">// The following tags on struct fields have special meaning to Unmarshal:</span>
<a id="L867"></a><span class="comment">//</span>
<a id="L868"></a><span class="comment">//	optional		marks the field as ASN.1 OPTIONAL</span>
<a id="L869"></a><span class="comment">//	[explicit] tag:x	specifies the ASN.1 tag number; implies ASN.1 CONTEXT SPECIFIC</span>
<a id="L870"></a><span class="comment">//	default:x		sets the default value for optional integer fields</span>
<a id="L871"></a><span class="comment">//</span>
<a id="L872"></a><span class="comment">// Other ASN.1 types are not supported; if it encounters them,</span>
<a id="L873"></a><span class="comment">// Unmarshal returns a parse error.</span>
<a id="L874"></a>func Unmarshal(val interface{}, b []byte) os.Error {
    <a id="L875"></a>v := reflect.NewValue(val).(*reflect.PtrValue).Elem();
    <a id="L876"></a>_, err := parseField(v, b, 0, fieldParameters{});
    <a id="L877"></a>return err;
<a id="L878"></a>}
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
