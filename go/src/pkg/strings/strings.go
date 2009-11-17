<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strings/strings.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/strings/strings.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// A package of simple functions to manipulate strings.</span>
<a id="L6"></a>package strings

<a id="L8"></a>import (
    <a id="L9"></a>&#34;unicode&#34;;
    <a id="L10"></a>&#34;utf8&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">// explode splits s into an array of UTF-8 sequences, one per Unicode character (still strings) up to a maximum of n (n &lt;= 0 means no limit).</span>
<a id="L14"></a><span class="comment">// Invalid UTF-8 sequences become correct encodings of U+FFF8.</span>
<a id="L15"></a>func explode(s string, n int) []string {
    <a id="L16"></a>if n &lt;= 0 {
        <a id="L17"></a>n = len(s)
    <a id="L18"></a>}
    <a id="L19"></a>a := make([]string, n);
    <a id="L20"></a>var size, rune int;
    <a id="L21"></a>na := 0;
    <a id="L22"></a>for len(s) &gt; 0 {
        <a id="L23"></a>if na+1 &gt;= n {
            <a id="L24"></a>a[na] = s;
            <a id="L25"></a>na++;
            <a id="L26"></a>break;
        <a id="L27"></a>}
        <a id="L28"></a>rune, size = utf8.DecodeRuneInString(s);
        <a id="L29"></a>s = s[size:len(s)];
        <a id="L30"></a>a[na] = string(rune);
        <a id="L31"></a>na++;
    <a id="L32"></a>}
    <a id="L33"></a>return a[0:na];
<a id="L34"></a>}

<a id="L36"></a><span class="comment">// Count counts the number of non-overlapping instances of sep in s.</span>
<a id="L37"></a>func Count(s, sep string) int {
    <a id="L38"></a>if sep == &#34;&#34; {
        <a id="L39"></a>return utf8.RuneCountInString(s) + 1
    <a id="L40"></a>}
    <a id="L41"></a>c := sep[0];
    <a id="L42"></a>n := 0;
    <a id="L43"></a>for i := 0; i+len(sep) &lt;= len(s); i++ {
        <a id="L44"></a>if s[i] == c &amp;&amp; (len(sep) == 1 || s[i:i+len(sep)] == sep) {
            <a id="L45"></a>n++;
            <a id="L46"></a>i += len(sep) - 1;
        <a id="L47"></a>}
    <a id="L48"></a>}
    <a id="L49"></a>return n;
<a id="L50"></a>}

<a id="L52"></a><span class="comment">// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.</span>
<a id="L53"></a>func Index(s, sep string) int {
    <a id="L54"></a>n := len(sep);
    <a id="L55"></a>if n == 0 {
        <a id="L56"></a>return 0
    <a id="L57"></a>}
    <a id="L58"></a>c := sep[0];
    <a id="L59"></a>for i := 0; i+n &lt;= len(s); i++ {
        <a id="L60"></a>if s[i] == c &amp;&amp; (n == 1 || s[i:i+n] == sep) {
            <a id="L61"></a>return i
        <a id="L62"></a>}
    <a id="L63"></a>}
    <a id="L64"></a>return -1;
<a id="L65"></a>}

<a id="L67"></a><span class="comment">// LastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.</span>
<a id="L68"></a>func LastIndex(s, sep string) int {
    <a id="L69"></a>n := len(sep);
    <a id="L70"></a>if n == 0 {
        <a id="L71"></a>return len(s)
    <a id="L72"></a>}
    <a id="L73"></a>c := sep[0];
    <a id="L74"></a>for i := len(s) - n; i &gt;= 0; i-- {
        <a id="L75"></a>if s[i] == c &amp;&amp; (n == 1 || s[i:i+n] == sep) {
            <a id="L76"></a>return i
        <a id="L77"></a>}
    <a id="L78"></a>}
    <a id="L79"></a>return -1;
<a id="L80"></a>}

<a id="L82"></a><span class="comment">// Generic split: splits after each instance of sep,</span>
<a id="L83"></a><span class="comment">// including sepSave bytes of sep in the subarrays.</span>
<a id="L84"></a>func genSplit(s, sep string, sepSave, n int) []string {
    <a id="L85"></a>if sep == &#34;&#34; {
        <a id="L86"></a>return explode(s, n)
    <a id="L87"></a>}
    <a id="L88"></a>if n &lt;= 0 {
        <a id="L89"></a>n = Count(s, sep) + 1
    <a id="L90"></a>}
    <a id="L91"></a>c := sep[0];
    <a id="L92"></a>start := 0;
    <a id="L93"></a>a := make([]string, n);
    <a id="L94"></a>na := 0;
    <a id="L95"></a>for i := 0; i+len(sep) &lt;= len(s) &amp;&amp; na+1 &lt; n; i++ {
        <a id="L96"></a>if s[i] == c &amp;&amp; (len(sep) == 1 || s[i:i+len(sep)] == sep) {
            <a id="L97"></a>a[na] = s[start : i+sepSave];
            <a id="L98"></a>na++;
            <a id="L99"></a>start = i + len(sep);
            <a id="L100"></a>i += len(sep) - 1;
        <a id="L101"></a>}
    <a id="L102"></a>}
    <a id="L103"></a>a[na] = s[start:len(s)];
    <a id="L104"></a>return a[0 : na+1];
<a id="L105"></a>}

<a id="L107"></a><span class="comment">// Split splits the string s around each instance of sep, returning an array of substrings of s.</span>
<a id="L108"></a><span class="comment">// If sep is empty, Split splits s after each UTF-8 sequence.</span>
<a id="L109"></a><span class="comment">// If n &gt; 0, split Splits s into at most n substrings; the last substring will be the unsplit remainder.</span>
<a id="L110"></a>func Split(s, sep string, n int) []string { return genSplit(s, sep, 0, n) }

<a id="L112"></a><span class="comment">// SplitAfter splits the string s after each instance of sep, returning an array of substrings of s.</span>
<a id="L113"></a><span class="comment">// If sep is empty, SplitAfter splits s after each UTF-8 sequence.</span>
<a id="L114"></a><span class="comment">// If n &gt; 0, SplitAfter splits s into at most n substrings; the last substring will be the unsplit remainder.</span>
<a id="L115"></a>func SplitAfter(s, sep string, n int) []string {
    <a id="L116"></a>return genSplit(s, sep, len(sep), n)
<a id="L117"></a>}

<a id="L119"></a><span class="comment">// Join concatenates the elements of a to create a single string.   The separator string</span>
<a id="L120"></a><span class="comment">// sep is placed between elements in the resulting string.</span>
<a id="L121"></a>func Join(a []string, sep string) string {
    <a id="L122"></a>if len(a) == 0 {
        <a id="L123"></a>return &#34;&#34;
    <a id="L124"></a>}
    <a id="L125"></a>if len(a) == 1 {
        <a id="L126"></a>return a[0]
    <a id="L127"></a>}
    <a id="L128"></a>n := len(sep) * (len(a) - 1);
    <a id="L129"></a>for i := 0; i &lt; len(a); i++ {
        <a id="L130"></a>n += len(a[i])
    <a id="L131"></a>}

    <a id="L133"></a>b := make([]byte, n);
    <a id="L134"></a>bp := 0;
    <a id="L135"></a>for i := 0; i &lt; len(a); i++ {
        <a id="L136"></a>s := a[i];
        <a id="L137"></a>for j := 0; j &lt; len(s); j++ {
            <a id="L138"></a>b[bp] = s[j];
            <a id="L139"></a>bp++;
        <a id="L140"></a>}
        <a id="L141"></a>if i+1 &lt; len(a) {
            <a id="L142"></a>s = sep;
            <a id="L143"></a>for j := 0; j &lt; len(s); j++ {
                <a id="L144"></a>b[bp] = s[j];
                <a id="L145"></a>bp++;
            <a id="L146"></a>}
        <a id="L147"></a>}
    <a id="L148"></a>}
    <a id="L149"></a>return string(b);
<a id="L150"></a>}

<a id="L152"></a><span class="comment">// HasPrefix tests whether the string s begins with prefix.</span>
<a id="L153"></a>func HasPrefix(s, prefix string) bool {
    <a id="L154"></a>return len(s) &gt;= len(prefix) &amp;&amp; s[0:len(prefix)] == prefix
<a id="L155"></a>}

<a id="L157"></a><span class="comment">// HasSuffix tests whether the string s ends with suffix.</span>
<a id="L158"></a>func HasSuffix(s, suffix string) bool {
    <a id="L159"></a>return len(s) &gt;= len(suffix) &amp;&amp; s[len(s)-len(suffix):len(s)] == suffix
<a id="L160"></a>}

<a id="L162"></a><span class="comment">// Map returns a copy of the string s with all its characters modified</span>
<a id="L163"></a><span class="comment">// according to the mapping function.</span>
<a id="L164"></a>func Map(mapping func(rune int) int, s string) string {
    <a id="L165"></a><span class="comment">// In the worst case, the string can grow when mapped, making</span>
    <a id="L166"></a><span class="comment">// things unpleasant.  But it&#39;s so rare we barge in assuming it&#39;s</span>
    <a id="L167"></a><span class="comment">// fine.  It could also shrink but that falls out naturally.</span>
    <a id="L168"></a>maxbytes := len(s); <span class="comment">// length of b</span>
    <a id="L169"></a>nbytes := 0;        <span class="comment">// number of bytes encoded in b</span>
    <a id="L170"></a>b := make([]byte, maxbytes);
    <a id="L171"></a>for _, c := range s {
        <a id="L172"></a>rune := mapping(c);
        <a id="L173"></a>wid := 1;
        <a id="L174"></a>if rune &gt;= utf8.RuneSelf {
            <a id="L175"></a>wid = utf8.RuneLen(rune)
        <a id="L176"></a>}
        <a id="L177"></a>if nbytes+wid &gt; maxbytes {
            <a id="L178"></a><span class="comment">// Grow the buffer.</span>
            <a id="L179"></a>maxbytes = maxbytes*2 + utf8.UTFMax;
            <a id="L180"></a>nb := make([]byte, maxbytes);
            <a id="L181"></a>for i, c := range b[0:nbytes] {
                <a id="L182"></a>nb[i] = c
            <a id="L183"></a>}
            <a id="L184"></a>b = nb;
        <a id="L185"></a>}
        <a id="L186"></a>nbytes += utf8.EncodeRune(rune, b[nbytes:maxbytes]);
    <a id="L187"></a>}
    <a id="L188"></a>return string(b[0:nbytes]);
<a id="L189"></a>}

<a id="L191"></a><span class="comment">// ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case.</span>
<a id="L192"></a>func ToUpper(s string) string { return Map(unicode.ToUpper, s) }

<a id="L194"></a><span class="comment">// ToUpper returns a copy of the string s with all Unicode letters mapped to their lower case.</span>
<a id="L195"></a>func ToLower(s string) string { return Map(unicode.ToLower, s) }

<a id="L197"></a><span class="comment">// ToTitle returns a copy of the string s with all Unicode letters mapped to their title case.</span>
<a id="L198"></a>func ToTitle(s string) string { return Map(unicode.ToTitle, s) }

<a id="L200"></a><span class="comment">// Trim returns a slice of the string s, with all leading and trailing white space</span>
<a id="L201"></a><span class="comment">// removed, as defined by Unicode.</span>
<a id="L202"></a>func TrimSpace(s string) string {
    <a id="L203"></a>start, end := 0, len(s);
    <a id="L204"></a>for start &lt; end {
        <a id="L205"></a>wid := 1;
        <a id="L206"></a>rune := int(s[start]);
        <a id="L207"></a>if rune &gt;= utf8.RuneSelf {
            <a id="L208"></a>rune, wid = utf8.DecodeRuneInString(s[start:end])
        <a id="L209"></a>}
        <a id="L210"></a>if !unicode.IsSpace(rune) {
            <a id="L211"></a>break
        <a id="L212"></a>}
        <a id="L213"></a>start += wid;
    <a id="L214"></a>}
    <a id="L215"></a>for start &lt; end {
        <a id="L216"></a>wid := 1;
        <a id="L217"></a>rune := int(s[end-1]);
        <a id="L218"></a>if rune &gt;= utf8.RuneSelf {
            <a id="L219"></a><span class="comment">// Back up carefully looking for beginning of rune. Mustn&#39;t pass start.</span>
            <a id="L220"></a>for wid = 2; start &lt;= end-wid &amp;&amp; !utf8.RuneStart(s[end-wid]); wid++ {
            <a id="L221"></a>}
            <a id="L222"></a>if start &gt; end-wid { <span class="comment">// invalid UTF-8 sequence; stop processing</span>
                <a id="L223"></a>return s[start:end]
            <a id="L224"></a>}
            <a id="L225"></a>rune, wid = utf8.DecodeRuneInString(s[end-wid : end]);
        <a id="L226"></a>}
        <a id="L227"></a>if !unicode.IsSpace(rune) {
            <a id="L228"></a>break
        <a id="L229"></a>}
        <a id="L230"></a>end -= wid;
    <a id="L231"></a>}
    <a id="L232"></a>return s[start:end];
<a id="L233"></a>}

<a id="L235"></a><span class="comment">// Bytes returns a new slice containing the bytes in s.</span>
<a id="L236"></a>func Bytes(s string) []byte {
    <a id="L237"></a>b := make([]byte, len(s));
    <a id="L238"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L239"></a>b[i] = s[i]
    <a id="L240"></a>}
    <a id="L241"></a>return b;
<a id="L242"></a>}
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
