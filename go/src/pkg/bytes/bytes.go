<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/bytes/bytes.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/bytes/bytes.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The bytes package implements functions for the manipulation of byte slices.</span>
<a id="L6"></a><span class="comment">// Analagous to the facilities of the strings package.</span>
<a id="L7"></a>package bytes

<a id="L9"></a>import (
    <a id="L10"></a>&#34;unicode&#34;;
    <a id="L11"></a>&#34;utf8&#34;;
<a id="L12"></a>)

<a id="L14"></a><span class="comment">// Compare returns an integer comparing the two byte arrays lexicographically.</span>
<a id="L15"></a><span class="comment">// The result will be 0 if a==b, -1 if a &lt; b, and +1 if a &gt; b</span>
<a id="L16"></a>func Compare(a, b []byte) int {
    <a id="L17"></a>for i := 0; i &lt; len(a) &amp;&amp; i &lt; len(b); i++ {
        <a id="L18"></a>switch {
        <a id="L19"></a>case a[i] &gt; b[i]:
            <a id="L20"></a>return 1
        <a id="L21"></a>case a[i] &lt; b[i]:
            <a id="L22"></a>return -1
        <a id="L23"></a>}
    <a id="L24"></a>}
    <a id="L25"></a>switch {
    <a id="L26"></a>case len(a) &lt; len(b):
        <a id="L27"></a>return -1
    <a id="L28"></a>case len(a) &gt; len(b):
        <a id="L29"></a>return 1
    <a id="L30"></a>}
    <a id="L31"></a>return 0;
<a id="L32"></a>}

<a id="L34"></a><span class="comment">// Equal returns a boolean reporting whether a == b.</span>
<a id="L35"></a>func Equal(a, b []byte) bool {
    <a id="L36"></a>if len(a) != len(b) {
        <a id="L37"></a>return false
    <a id="L38"></a>}
    <a id="L39"></a>for i := 0; i &lt; len(a); i++ {
        <a id="L40"></a>if a[i] != b[i] {
            <a id="L41"></a>return false
        <a id="L42"></a>}
    <a id="L43"></a>}
    <a id="L44"></a>return true;
<a id="L45"></a>}

<a id="L47"></a><span class="comment">// Copy copies bytes from src to dst,</span>
<a id="L48"></a><span class="comment">// stopping when either all of src has been copied</span>
<a id="L49"></a><span class="comment">// or all of dst has been filled.</span>
<a id="L50"></a><span class="comment">// It returns the number of bytes copied.</span>
<a id="L51"></a>func Copy(dst, src []byte) int {
    <a id="L52"></a>if len(src) &gt; len(dst) {
        <a id="L53"></a>src = src[0:len(dst)]
    <a id="L54"></a>}
    <a id="L55"></a>for i, x := range src {
        <a id="L56"></a>dst[i] = x
    <a id="L57"></a>}
    <a id="L58"></a>return len(src);
<a id="L59"></a>}

<a id="L61"></a><span class="comment">// explode splits s into an array of UTF-8 sequences, one per Unicode character (still arrays of bytes),</span>
<a id="L62"></a><span class="comment">// up to a maximum of n byte arrays. Invalid UTF-8 sequences are chopped into individual bytes.</span>
<a id="L63"></a>func explode(s []byte, n int) [][]byte {
    <a id="L64"></a>if n &lt;= 0 {
        <a id="L65"></a>n = len(s)
    <a id="L66"></a>}
    <a id="L67"></a>a := make([][]byte, n);
    <a id="L68"></a>var size int;
    <a id="L69"></a>na := 0;
    <a id="L70"></a>for len(s) &gt; 0 {
        <a id="L71"></a>if na+1 &gt;= n {
            <a id="L72"></a>a[na] = s;
            <a id="L73"></a>na++;
            <a id="L74"></a>break;
        <a id="L75"></a>}
        <a id="L76"></a>_, size = utf8.DecodeRune(s);
        <a id="L77"></a>a[na] = s[0:size];
        <a id="L78"></a>s = s[size:len(s)];
        <a id="L79"></a>na++;
    <a id="L80"></a>}
    <a id="L81"></a>return a[0:na];
<a id="L82"></a>}

<a id="L84"></a><span class="comment">// Count counts the number of non-overlapping instances of sep in s.</span>
<a id="L85"></a>func Count(s, sep []byte) int {
    <a id="L86"></a>if len(sep) == 0 {
        <a id="L87"></a>return utf8.RuneCount(s) + 1
    <a id="L88"></a>}
    <a id="L89"></a>c := sep[0];
    <a id="L90"></a>n := 0;
    <a id="L91"></a>for i := 0; i+len(sep) &lt;= len(s); i++ {
        <a id="L92"></a>if s[i] == c &amp;&amp; (len(sep) == 1 || Equal(s[i:i+len(sep)], sep)) {
            <a id="L93"></a>n++;
            <a id="L94"></a>i += len(sep) - 1;
        <a id="L95"></a>}
    <a id="L96"></a>}
    <a id="L97"></a>return n;
<a id="L98"></a>}

<a id="L100"></a><span class="comment">// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.</span>
<a id="L101"></a>func Index(s, sep []byte) int {
    <a id="L102"></a>n := len(sep);
    <a id="L103"></a>if n == 0 {
        <a id="L104"></a>return 0
    <a id="L105"></a>}
    <a id="L106"></a>c := sep[0];
    <a id="L107"></a>for i := 0; i+n &lt;= len(s); i++ {
        <a id="L108"></a>if s[i] == c &amp;&amp; (n == 1 || Equal(s[i:i+n], sep)) {
            <a id="L109"></a>return i
        <a id="L110"></a>}
    <a id="L111"></a>}
    <a id="L112"></a>return -1;
<a id="L113"></a>}

<a id="L115"></a><span class="comment">// LastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.</span>
<a id="L116"></a>func LastIndex(s, sep []byte) int {
    <a id="L117"></a>n := len(sep);
    <a id="L118"></a>if n == 0 {
        <a id="L119"></a>return len(s)
    <a id="L120"></a>}
    <a id="L121"></a>c := sep[0];
    <a id="L122"></a>for i := len(s) - n; i &gt;= 0; i-- {
        <a id="L123"></a>if s[i] == c &amp;&amp; (n == 1 || Equal(s[i:i+n], sep)) {
            <a id="L124"></a>return i
        <a id="L125"></a>}
    <a id="L126"></a>}
    <a id="L127"></a>return -1;
<a id="L128"></a>}

<a id="L130"></a><span class="comment">// Generic split: splits after each instance of sep,</span>
<a id="L131"></a><span class="comment">// including sepSave bytes of sep in the subarrays.</span>
<a id="L132"></a>func genSplit(s, sep []byte, sepSave, n int) [][]byte {
    <a id="L133"></a>if len(sep) == 0 {
        <a id="L134"></a>return explode(s, n)
    <a id="L135"></a>}
    <a id="L136"></a>if n &lt;= 0 {
        <a id="L137"></a>n = Count(s, sep) + 1
    <a id="L138"></a>}
    <a id="L139"></a>c := sep[0];
    <a id="L140"></a>start := 0;
    <a id="L141"></a>a := make([][]byte, n);
    <a id="L142"></a>na := 0;
    <a id="L143"></a>for i := 0; i+len(sep) &lt;= len(s) &amp;&amp; na+1 &lt; n; i++ {
        <a id="L144"></a>if s[i] == c &amp;&amp; (len(sep) == 1 || Equal(s[i:i+len(sep)], sep)) {
            <a id="L145"></a>a[na] = s[start : i+sepSave];
            <a id="L146"></a>na++;
            <a id="L147"></a>start = i + len(sep);
            <a id="L148"></a>i += len(sep) - 1;
        <a id="L149"></a>}
    <a id="L150"></a>}
    <a id="L151"></a>a[na] = s[start:len(s)];
    <a id="L152"></a>return a[0 : na+1];
<a id="L153"></a>}

<a id="L155"></a><span class="comment">// Split splits the array s around each instance of sep, returning an array of subarrays of s.</span>
<a id="L156"></a><span class="comment">// If sep is empty, Split splits s after each UTF-8 sequence.</span>
<a id="L157"></a><span class="comment">// If n &gt; 0, Split splits s into at most n subarrays; the last subarray will contain an unsplit remainder.</span>
<a id="L158"></a>func Split(s, sep []byte, n int) [][]byte { return genSplit(s, sep, 0, n) }

<a id="L160"></a><span class="comment">// SplitAfter splits the array s after each instance of sep, returning an array of subarrays of s.</span>
<a id="L161"></a><span class="comment">// If sep is empty, SplitAfter splits s after each UTF-8 sequence.</span>
<a id="L162"></a><span class="comment">// If n &gt; 0, SplitAfter splits s into at most n subarrays; the last subarray will contain an</span>
<a id="L163"></a><span class="comment">// unsplit remainder.</span>
<a id="L164"></a>func SplitAfter(s, sep []byte, n int) [][]byte {
    <a id="L165"></a>return genSplit(s, sep, len(sep), n)
<a id="L166"></a>}

<a id="L168"></a><span class="comment">// Join concatenates the elements of a to create a single byte array.   The separator</span>
<a id="L169"></a><span class="comment">// sep is placed between elements in the resulting array.</span>
<a id="L170"></a>func Join(a [][]byte, sep []byte) []byte {
    <a id="L171"></a>if len(a) == 0 {
        <a id="L172"></a>return []byte{}
    <a id="L173"></a>}
    <a id="L174"></a>if len(a) == 1 {
        <a id="L175"></a>return a[0]
    <a id="L176"></a>}
    <a id="L177"></a>n := len(sep) * (len(a) - 1);
    <a id="L178"></a>for i := 0; i &lt; len(a); i++ {
        <a id="L179"></a>n += len(a[i])
    <a id="L180"></a>}

    <a id="L182"></a>b := make([]byte, n);
    <a id="L183"></a>bp := 0;
    <a id="L184"></a>for i := 0; i &lt; len(a); i++ {
        <a id="L185"></a>s := a[i];
        <a id="L186"></a>for j := 0; j &lt; len(s); j++ {
            <a id="L187"></a>b[bp] = s[j];
            <a id="L188"></a>bp++;
        <a id="L189"></a>}
        <a id="L190"></a>if i+1 &lt; len(a) {
            <a id="L191"></a>s = sep;
            <a id="L192"></a>for j := 0; j &lt; len(s); j++ {
                <a id="L193"></a>b[bp] = s[j];
                <a id="L194"></a>bp++;
            <a id="L195"></a>}
        <a id="L196"></a>}
    <a id="L197"></a>}
    <a id="L198"></a>return b;
<a id="L199"></a>}

<a id="L201"></a><span class="comment">// HasPrefix tests whether the byte array s begins with prefix.</span>
<a id="L202"></a>func HasPrefix(s, prefix []byte) bool {
    <a id="L203"></a>return len(s) &gt;= len(prefix) &amp;&amp; Equal(s[0:len(prefix)], prefix)
<a id="L204"></a>}

<a id="L206"></a><span class="comment">// HasSuffix tests whether the byte array s ends with suffix.</span>
<a id="L207"></a>func HasSuffix(s, suffix []byte) bool {
    <a id="L208"></a>return len(s) &gt;= len(suffix) &amp;&amp; Equal(s[len(s)-len(suffix):len(s)], suffix)
<a id="L209"></a>}

<a id="L211"></a><span class="comment">// Map returns a copy of the byte array s with all its characters modified</span>
<a id="L212"></a><span class="comment">// according to the mapping function.</span>
<a id="L213"></a>func Map(mapping func(rune int) int, s []byte) []byte {
    <a id="L214"></a><span class="comment">// In the worst case, the array can grow when mapped, making</span>
    <a id="L215"></a><span class="comment">// things unpleasant.  But it&#39;s so rare we barge in assuming it&#39;s</span>
    <a id="L216"></a><span class="comment">// fine.  It could also shrink but that falls out naturally.</span>
    <a id="L217"></a>maxbytes := len(s); <span class="comment">// length of b</span>
    <a id="L218"></a>nbytes := 0;        <span class="comment">// number of bytes encoded in b</span>
    <a id="L219"></a>b := make([]byte, maxbytes);
    <a id="L220"></a>for i := 0; i &lt; len(s); {
        <a id="L221"></a>wid := 1;
        <a id="L222"></a>rune := int(s[i]);
        <a id="L223"></a>if rune &lt; utf8.RuneSelf {
            <a id="L224"></a>rune = mapping(rune)
        <a id="L225"></a>} else {
            <a id="L226"></a>rune, wid = utf8.DecodeRune(s[i:len(s)])
        <a id="L227"></a>}
        <a id="L228"></a>rune = mapping(rune);
        <a id="L229"></a>if nbytes+utf8.RuneLen(rune) &gt; maxbytes {
            <a id="L230"></a><span class="comment">// Grow the buffer.</span>
            <a id="L231"></a>maxbytes = maxbytes*2 + utf8.UTFMax;
            <a id="L232"></a>nb := make([]byte, maxbytes);
            <a id="L233"></a>for i, c := range b[0:nbytes] {
                <a id="L234"></a>nb[i] = c
            <a id="L235"></a>}
            <a id="L236"></a>b = nb;
        <a id="L237"></a>}
        <a id="L238"></a>nbytes += utf8.EncodeRune(rune, b[nbytes:maxbytes]);
        <a id="L239"></a>i += wid;
    <a id="L240"></a>}
    <a id="L241"></a>return b[0:nbytes];
<a id="L242"></a>}

<a id="L244"></a><span class="comment">// ToUpper returns a copy of the byte array s with all Unicode letters mapped to their upper case.</span>
<a id="L245"></a>func ToUpper(s []byte) []byte { return Map(unicode.ToUpper, s) }

<a id="L247"></a><span class="comment">// ToUpper returns a copy of the byte array s with all Unicode letters mapped to their lower case.</span>
<a id="L248"></a>func ToLower(s []byte) []byte { return Map(unicode.ToLower, s) }

<a id="L250"></a><span class="comment">// ToTitle returns a copy of the byte array s with all Unicode letters mapped to their title case.</span>
<a id="L251"></a>func ToTitle(s []byte) []byte { return Map(unicode.ToTitle, s) }

<a id="L253"></a><span class="comment">// Trim returns a slice of the string s, with all leading and trailing white space</span>
<a id="L254"></a><span class="comment">// removed, as defined by Unicode.</span>
<a id="L255"></a>func TrimSpace(s []byte) []byte {
    <a id="L256"></a>start, end := 0, len(s);
    <a id="L257"></a>for start &lt; end {
        <a id="L258"></a>wid := 1;
        <a id="L259"></a>rune := int(s[start]);
        <a id="L260"></a>if rune &gt;= utf8.RuneSelf {
            <a id="L261"></a>rune, wid = utf8.DecodeRune(s[start:end])
        <a id="L262"></a>}
        <a id="L263"></a>if !unicode.IsSpace(rune) {
            <a id="L264"></a>break
        <a id="L265"></a>}
        <a id="L266"></a>start += wid;
    <a id="L267"></a>}
    <a id="L268"></a>for start &lt; end {
        <a id="L269"></a>wid := 1;
        <a id="L270"></a>rune := int(s[end-1]);
        <a id="L271"></a>if rune &gt;= utf8.RuneSelf {
            <a id="L272"></a><span class="comment">// Back up carefully looking for beginning of rune. Mustn&#39;t pass start.</span>
            <a id="L273"></a>for wid = 2; start &lt;= end-wid &amp;&amp; !utf8.RuneStart(s[end-wid]); wid++ {
            <a id="L274"></a>}
            <a id="L275"></a>if start &gt; end-wid { <span class="comment">// invalid UTF-8 sequence; stop processing</span>
                <a id="L276"></a>return s[start:end]
            <a id="L277"></a>}
            <a id="L278"></a>rune, wid = utf8.DecodeRune(s[end-wid : end]);
        <a id="L279"></a>}
        <a id="L280"></a>if !unicode.IsSpace(rune) {
            <a id="L281"></a>break
        <a id="L282"></a>}
        <a id="L283"></a>end -= wid;
    <a id="L284"></a>}
    <a id="L285"></a>return s[start:end];
<a id="L286"></a>}

<a id="L288"></a><span class="comment">// How big to make a byte array when growing.</span>
<a id="L289"></a><span class="comment">// Heuristic: Scale by 50% to give n log n time.</span>
<a id="L290"></a>func resize(n int) int {
    <a id="L291"></a>if n &lt; 16 {
        <a id="L292"></a>n = 16
    <a id="L293"></a>}
    <a id="L294"></a>return n + n/2;
<a id="L295"></a>}

<a id="L297"></a><span class="comment">// Add appends the contents of t to the end of s and returns the result.</span>
<a id="L298"></a><span class="comment">// If s has enough capacity, it is extended in place; otherwise a</span>
<a id="L299"></a><span class="comment">// new array is allocated and returned.</span>
<a id="L300"></a>func Add(s, t []byte) []byte {
    <a id="L301"></a>lens := len(s);
    <a id="L302"></a>lent := len(t);
    <a id="L303"></a>if lens+lent &lt;= cap(s) {
        <a id="L304"></a>s = s[0 : lens+lent]
    <a id="L305"></a>} else {
        <a id="L306"></a>news := make([]byte, lens+lent, resize(lens+lent));
        <a id="L307"></a>Copy(news, s);
        <a id="L308"></a>s = news;
    <a id="L309"></a>}
    <a id="L310"></a>Copy(s[lens:lens+lent], t);
    <a id="L311"></a>return s;
<a id="L312"></a>}

<a id="L314"></a><span class="comment">// AddByte appends byte b to the end of s and returns the result.</span>
<a id="L315"></a><span class="comment">// If s has enough capacity, it is extended in place; otherwise a</span>
<a id="L316"></a><span class="comment">// new array is allocated and returned.</span>
<a id="L317"></a>func AddByte(s []byte, t byte) []byte {
    <a id="L318"></a>lens := len(s);
    <a id="L319"></a>if lens+1 &lt;= cap(s) {
        <a id="L320"></a>s = s[0 : lens+1]
    <a id="L321"></a>} else {
        <a id="L322"></a>news := make([]byte, lens+1, resize(lens+1));
        <a id="L323"></a>Copy(news, s);
        <a id="L324"></a>s = news;
    <a id="L325"></a>}
    <a id="L326"></a>s[lens] = t;
    <a id="L327"></a>return s;
<a id="L328"></a>}
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
