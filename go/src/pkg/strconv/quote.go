<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strconv/quote.go</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/strconv/quote.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package strconv

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>&#34;strings&#34;;
    <a id="L11"></a>&#34;unicode&#34;;
    <a id="L12"></a>&#34;utf8&#34;;
<a id="L13"></a>)

<a id="L15"></a>const lowerhex = &#34;0123456789abcdef&#34;

<a id="L17"></a><span class="comment">// Quote returns a double-quoted Go string literal</span>
<a id="L18"></a><span class="comment">// representing s.  The returned string s uses Go escape</span>
<a id="L19"></a><span class="comment">// sequences (\t, \n, \xFF, \u0100) for control characters</span>
<a id="L20"></a><span class="comment">// and non-ASCII characters.</span>
<a id="L21"></a>func Quote(s string) string {
    <a id="L22"></a>var buf bytes.Buffer;
    <a id="L23"></a>buf.WriteByte(&#39;&#34;&#39;);
    <a id="L24"></a>for ; len(s) &gt; 0; s = s[1:len(s)] {
        <a id="L25"></a>switch c := s[0]; {
        <a id="L26"></a>case c == &#39;&#34;&#39;:
            <a id="L27"></a>buf.WriteString(`\&#34;`)
        <a id="L28"></a>case c == &#39;\\&#39;:
            <a id="L29"></a>buf.WriteString(`\\`)
        <a id="L30"></a>case &#39; &#39; &lt;= c &amp;&amp; c &lt;= &#39;~&#39;:
            <a id="L31"></a>buf.WriteString(string(c))
        <a id="L32"></a>case c == &#39;\a&#39;:
            <a id="L33"></a>buf.WriteString(`\a`)
        <a id="L34"></a>case c == &#39;\b&#39;:
            <a id="L35"></a>buf.WriteString(`\b`)
        <a id="L36"></a>case c == &#39;\f&#39;:
            <a id="L37"></a>buf.WriteString(`\f`)
        <a id="L38"></a>case c == &#39;\n&#39;:
            <a id="L39"></a>buf.WriteString(`\n`)
        <a id="L40"></a>case c == &#39;\r&#39;:
            <a id="L41"></a>buf.WriteString(`\r`)
        <a id="L42"></a>case c == &#39;\t&#39;:
            <a id="L43"></a>buf.WriteString(`\t`)
        <a id="L44"></a>case c == &#39;\v&#39;:
            <a id="L45"></a>buf.WriteString(`\v`)

        <a id="L47"></a>case c &gt;= utf8.RuneSelf &amp;&amp; utf8.FullRuneInString(s):
            <a id="L48"></a>r, size := utf8.DecodeRuneInString(s);
            <a id="L49"></a>if r == utf8.RuneError &amp;&amp; size == 1 {
                <a id="L50"></a>goto EscX
            <a id="L51"></a>}
            <a id="L52"></a>s = s[size-1 : len(s)]; <span class="comment">// next iteration will slice off 1 more</span>
            <a id="L53"></a>if r &lt; 0x10000 {
                <a id="L54"></a>buf.WriteString(`\u`);
                <a id="L55"></a>for j := uint(0); j &lt; 4; j++ {
                    <a id="L56"></a>buf.WriteByte(lowerhex[(r&gt;&gt;(12-4*j))&amp;0xF])
                <a id="L57"></a>}
            <a id="L58"></a>} else {
                <a id="L59"></a>buf.WriteString(`\U`);
                <a id="L60"></a>for j := uint(0); j &lt; 8; j++ {
                    <a id="L61"></a>buf.WriteByte(lowerhex[(r&gt;&gt;(28-4*j))&amp;0xF])
                <a id="L62"></a>}
            <a id="L63"></a>}

        <a id="L65"></a>default:
        <a id="L66"></a>EscX:
            <a id="L67"></a>buf.WriteString(`\x`);
            <a id="L68"></a>buf.WriteByte(lowerhex[c&gt;&gt;4]);
            <a id="L69"></a>buf.WriteByte(lowerhex[c&amp;0xF]);
        <a id="L70"></a>}
    <a id="L71"></a>}
    <a id="L72"></a>buf.WriteByte(&#39;&#34;&#39;);
    <a id="L73"></a>return buf.String();
<a id="L74"></a>}

<a id="L76"></a><span class="comment">// CanBackquote returns whether the string s would be</span>
<a id="L77"></a><span class="comment">// a valid Go string literal if enclosed in backquotes.</span>
<a id="L78"></a>func CanBackquote(s string) bool {
    <a id="L79"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L80"></a>if (s[i] &lt; &#39; &#39; &amp;&amp; s[i] != &#39;\t&#39;) || s[i] == &#39;`&#39; {
            <a id="L81"></a>return false
        <a id="L82"></a>}
    <a id="L83"></a>}
    <a id="L84"></a>return true;
<a id="L85"></a>}

<a id="L87"></a>func unhex(b byte) (v int, ok bool) {
    <a id="L88"></a>c := int(b);
    <a id="L89"></a>switch {
    <a id="L90"></a>case &#39;0&#39; &lt;= c &amp;&amp; c &lt;= &#39;9&#39;:
        <a id="L91"></a>return c - &#39;0&#39;, true
    <a id="L92"></a>case &#39;a&#39; &lt;= c &amp;&amp; c &lt;= &#39;f&#39;:
        <a id="L93"></a>return c - &#39;a&#39; + 10, true
    <a id="L94"></a>case &#39;A&#39; &lt;= c &amp;&amp; c &lt;= &#39;F&#39;:
        <a id="L95"></a>return c - &#39;A&#39; + 10, true
    <a id="L96"></a>}
    <a id="L97"></a>return;
<a id="L98"></a>}

<a id="L100"></a><span class="comment">// UnquoteChar decodes the first character or byte in the escaped string</span>
<a id="L101"></a><span class="comment">// or character literal represented by the string s.</span>
<a id="L102"></a><span class="comment">// It returns four values:</span>
<a id="L103"></a><span class="comment">// 1) value, the decoded Unicode code point or byte value;</span>
<a id="L104"></a><span class="comment">// 2) multibyte, a boolean indicating whether the decoded character</span>
<a id="L105"></a><span class="comment">//    requires a multibyte UTF-8 representation;</span>
<a id="L106"></a><span class="comment">// 3) tail, the remainder of the string after the character; and</span>
<a id="L107"></a><span class="comment">// 4) an error that will be nil if the character is syntactically valid.</span>
<a id="L108"></a><span class="comment">// The second argument, quote, specifies the type of literal being parsed</span>
<a id="L109"></a><span class="comment">// and therefore which escaped quote character is permitted.</span>
<a id="L110"></a><span class="comment">// If set to a single quote, it permits the sequence \&#39; and disallows unescaped &#39;.</span>
<a id="L111"></a><span class="comment">// If set to a double quote, it permits \&#34; and disallows unescaped &#34;.</span>
<a id="L112"></a><span class="comment">// If set to zero, it does not permit either escape and allows both quote characters to appear unescaped.</span>
<a id="L113"></a>func UnquoteChar(s string, quote byte) (value int, multibyte bool, tail string, err os.Error) {
    <a id="L114"></a><span class="comment">// easy cases</span>
    <a id="L115"></a>switch c := s[0]; {
    <a id="L116"></a>case c == quote &amp;&amp; (quote == &#39;\&#39;&#39; || quote == &#39;&#34;&#39;):
        <a id="L117"></a>err = os.EINVAL;
        <a id="L118"></a>return;
    <a id="L119"></a>case c &gt;= utf8.RuneSelf:
        <a id="L120"></a>r, size := utf8.DecodeRuneInString(s);
        <a id="L121"></a>return r, true, s[size:len(s)], nil;
    <a id="L122"></a>case c != &#39;\\&#39;:
        <a id="L123"></a>return int(s[0]), false, s[1:len(s)], nil
    <a id="L124"></a>}

    <a id="L126"></a><span class="comment">// hard case: c is backslash</span>
    <a id="L127"></a>if len(s) &lt;= 1 {
        <a id="L128"></a>err = os.EINVAL;
        <a id="L129"></a>return;
    <a id="L130"></a>}
    <a id="L131"></a>c := s[1];
    <a id="L132"></a>s = s[2:len(s)];

    <a id="L134"></a>switch c {
    <a id="L135"></a>case &#39;a&#39;:
        <a id="L136"></a>value = &#39;\a&#39;
    <a id="L137"></a>case &#39;b&#39;:
        <a id="L138"></a>value = &#39;\b&#39;
    <a id="L139"></a>case &#39;f&#39;:
        <a id="L140"></a>value = &#39;\f&#39;
    <a id="L141"></a>case &#39;n&#39;:
        <a id="L142"></a>value = &#39;\n&#39;
    <a id="L143"></a>case &#39;r&#39;:
        <a id="L144"></a>value = &#39;\r&#39;
    <a id="L145"></a>case &#39;t&#39;:
        <a id="L146"></a>value = &#39;\t&#39;
    <a id="L147"></a>case &#39;v&#39;:
        <a id="L148"></a>value = &#39;\v&#39;
    <a id="L149"></a>case &#39;x&#39;, &#39;u&#39;, &#39;U&#39;:
        <a id="L150"></a>n := 0;
        <a id="L151"></a>switch c {
        <a id="L152"></a>case &#39;x&#39;:
            <a id="L153"></a>n = 2
        <a id="L154"></a>case &#39;u&#39;:
            <a id="L155"></a>n = 4
        <a id="L156"></a>case &#39;U&#39;:
            <a id="L157"></a>n = 8
        <a id="L158"></a>}
        <a id="L159"></a>v := 0;
        <a id="L160"></a>if len(s) &lt; n {
            <a id="L161"></a>err = os.EINVAL;
            <a id="L162"></a>return;
        <a id="L163"></a>}
        <a id="L164"></a>for j := 0; j &lt; n; j++ {
            <a id="L165"></a>x, ok := unhex(s[j]);
            <a id="L166"></a>if !ok {
                <a id="L167"></a>err = os.EINVAL;
                <a id="L168"></a>return;
            <a id="L169"></a>}
            <a id="L170"></a>v = v&lt;&lt;4 | x;
        <a id="L171"></a>}
        <a id="L172"></a>s = s[n:len(s)];
        <a id="L173"></a>if c == &#39;x&#39; {
            <a id="L174"></a><span class="comment">// single-byte string, possibly not UTF-8</span>
            <a id="L175"></a>value = v;
            <a id="L176"></a>break;
        <a id="L177"></a>}
        <a id="L178"></a>if v &gt; unicode.MaxRune {
            <a id="L179"></a>err = os.EINVAL;
            <a id="L180"></a>return;
        <a id="L181"></a>}
        <a id="L182"></a>value = v;
        <a id="L183"></a>multibyte = true;
    <a id="L184"></a>case &#39;0&#39;, &#39;1&#39;, &#39;2&#39;, &#39;3&#39;, &#39;4&#39;, &#39;5&#39;, &#39;6&#39;, &#39;7&#39;:
        <a id="L185"></a>v := int(c) - &#39;0&#39;;
        <a id="L186"></a>if len(s) &lt; 2 {
            <a id="L187"></a>err = os.EINVAL;
            <a id="L188"></a>return;
        <a id="L189"></a>}
        <a id="L190"></a>for j := 0; j &lt; 2; j++ { <span class="comment">// one digit already; two more</span>
            <a id="L191"></a>x := int(s[j]) - &#39;0&#39;;
            <a id="L192"></a>if x &lt; 0 || x &gt; 7 {
                <a id="L193"></a>return
            <a id="L194"></a>}
            <a id="L195"></a>v = (v &lt;&lt; 3) | x;
        <a id="L196"></a>}
        <a id="L197"></a>s = s[2:len(s)];
        <a id="L198"></a>if v &gt; 255 {
            <a id="L199"></a>err = os.EINVAL;
            <a id="L200"></a>return;
        <a id="L201"></a>}
        <a id="L202"></a>value = v;
    <a id="L203"></a>case &#39;\\&#39;:
        <a id="L204"></a>value = &#39;\\&#39;
    <a id="L205"></a>case &#39;\&#39;&#39;, &#39;&#34;&#39;:
        <a id="L206"></a>if c != quote {
            <a id="L207"></a>err = os.EINVAL;
            <a id="L208"></a>return;
        <a id="L209"></a>}
        <a id="L210"></a>value = int(c);
    <a id="L211"></a>default:
        <a id="L212"></a>err = os.EINVAL;
        <a id="L213"></a>return;
    <a id="L214"></a>}
    <a id="L215"></a>tail = s;
    <a id="L216"></a>return;
<a id="L217"></a>}

<a id="L219"></a><span class="comment">// Unquote interprets s as a single-quoted, double-quoted,</span>
<a id="L220"></a><span class="comment">// or backquoted Go string literal, returning the string value</span>
<a id="L221"></a><span class="comment">// that s quotes.  (If s is single-quoted, it would be a Go</span>
<a id="L222"></a><span class="comment">// character literal; Unquote returns the corresponding</span>
<a id="L223"></a><span class="comment">// one-character string.)</span>
<a id="L224"></a>func Unquote(s string) (t string, err os.Error) {
    <a id="L225"></a>n := len(s);
    <a id="L226"></a>if n &lt; 2 {
        <a id="L227"></a>return &#34;&#34;, os.EINVAL
    <a id="L228"></a>}
    <a id="L229"></a>quote := s[0];
    <a id="L230"></a>if quote != s[n-1] {
        <a id="L231"></a>return &#34;&#34;, os.EINVAL
    <a id="L232"></a>}
    <a id="L233"></a>s = s[1 : n-1];

    <a id="L235"></a>if quote == &#39;`&#39; {
        <a id="L236"></a>if strings.Index(s, &#34;`&#34;) &gt;= 0 {
            <a id="L237"></a>return &#34;&#34;, os.EINVAL
        <a id="L238"></a>}
        <a id="L239"></a>return s, nil;
    <a id="L240"></a>}
    <a id="L241"></a>if quote != &#39;&#34;&#39; &amp;&amp; quote != &#39;\&#39;&#39; {
        <a id="L242"></a>return &#34;&#34;, err
    <a id="L243"></a>}

    <a id="L245"></a>var buf bytes.Buffer;
    <a id="L246"></a>for len(s) &gt; 0 {
        <a id="L247"></a>c, multibyte, ss, err := UnquoteChar(s, quote);
        <a id="L248"></a>if err != nil {
            <a id="L249"></a>return &#34;&#34;, err
        <a id="L250"></a>}
        <a id="L251"></a>s = ss;
        <a id="L252"></a>if c &lt; utf8.RuneSelf || !multibyte {
            <a id="L253"></a>buf.WriteByte(byte(c))
        <a id="L254"></a>} else {
            <a id="L255"></a>buf.WriteString(string(c))
        <a id="L256"></a>}
        <a id="L257"></a>if quote == &#39;\&#39;&#39; &amp;&amp; len(s) != 0 {
            <a id="L258"></a><span class="comment">// single-quoted must be single character</span>
            <a id="L259"></a>return &#34;&#34;, os.EINVAL
        <a id="L260"></a>}
    <a id="L261"></a>}
    <a id="L262"></a>return buf.String(), nil;
<a id="L263"></a>}
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
