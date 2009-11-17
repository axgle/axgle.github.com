<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/parse.go</title>

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
	<li>Thu Nov 12 16:00:00 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/net/parse.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Simple file i/o and string manipulation, to avoid</span>
<a id="L6"></a><span class="comment">// depending on strconv and bufio and strings.</span>

<a id="L8"></a>package net

<a id="L10"></a>import (
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;os&#34;;
<a id="L13"></a>)

<a id="L15"></a>type file struct {
    <a id="L16"></a>file *os.File;
    <a id="L17"></a>data []byte;
<a id="L18"></a>}

<a id="L20"></a>func (f *file) close() { f.file.Close() }

<a id="L22"></a>func (f *file) getLineFromData() (s string, ok bool) {
    <a id="L23"></a>data := f.data;
    <a id="L24"></a>for i := 0; i &lt; len(data); i++ {
        <a id="L25"></a>if data[i] == &#39;\n&#39; {
            <a id="L26"></a>s = string(data[0:i]);
            <a id="L27"></a>ok = true;
            <a id="L28"></a><span class="comment">// move data</span>
            <a id="L29"></a>i++;
            <a id="L30"></a>n := len(data) - i;
            <a id="L31"></a>for j := 0; j &lt; n; j++ {
                <a id="L32"></a>data[j] = data[i+j]
            <a id="L33"></a>}
            <a id="L34"></a>f.data = data[0:n];
            <a id="L35"></a>return;
        <a id="L36"></a>}
    <a id="L37"></a>}
    <a id="L38"></a>return;
<a id="L39"></a>}

<a id="L41"></a>func (f *file) readLine() (s string, ok bool) {
    <a id="L42"></a>if s, ok = f.getLineFromData(); ok {
        <a id="L43"></a>return
    <a id="L44"></a>}
    <a id="L45"></a>if len(f.data) &lt; cap(f.data) {
        <a id="L46"></a>ln := len(f.data);
        <a id="L47"></a>n, _ := io.ReadFull(f.file, f.data[ln:cap(f.data)]);
        <a id="L48"></a>if n &gt;= 0 {
            <a id="L49"></a>f.data = f.data[0 : ln+n]
        <a id="L50"></a>}
    <a id="L51"></a>}
    <a id="L52"></a>s, ok = f.getLineFromData();
    <a id="L53"></a>return;
<a id="L54"></a>}

<a id="L56"></a>func open(name string) (*file, os.Error) {
    <a id="L57"></a>fd, err := os.Open(name, os.O_RDONLY, 0);
    <a id="L58"></a>if err != nil {
        <a id="L59"></a>return nil, err
    <a id="L60"></a>}
    <a id="L61"></a>return &amp;file{fd, make([]byte, 1024)[0:0]}, nil;
<a id="L62"></a>}

<a id="L64"></a>func byteIndex(s string, c byte) int {
    <a id="L65"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L66"></a>if s[i] == c {
            <a id="L67"></a>return i
        <a id="L68"></a>}
    <a id="L69"></a>}
    <a id="L70"></a>return -1;
<a id="L71"></a>}

<a id="L73"></a><span class="comment">// Count occurrences in s of any bytes in t.</span>
<a id="L74"></a>func countAnyByte(s string, t string) int {
    <a id="L75"></a>n := 0;
    <a id="L76"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L77"></a>if byteIndex(t, s[i]) &gt;= 0 {
            <a id="L78"></a>n++
        <a id="L79"></a>}
    <a id="L80"></a>}
    <a id="L81"></a>return n;
<a id="L82"></a>}

<a id="L84"></a><span class="comment">// Split s at any bytes in t.</span>
<a id="L85"></a>func splitAtBytes(s string, t string) []string {
    <a id="L86"></a>a := make([]string, 1+countAnyByte(s, t));
    <a id="L87"></a>n := 0;
    <a id="L88"></a>last := 0;
    <a id="L89"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L90"></a>if byteIndex(t, s[i]) &gt;= 0 {
            <a id="L91"></a>if last &lt; i {
                <a id="L92"></a>a[n] = string(s[last:i]);
                <a id="L93"></a>n++;
            <a id="L94"></a>}
            <a id="L95"></a>last = i + 1;
        <a id="L96"></a>}
    <a id="L97"></a>}
    <a id="L98"></a>if last &lt; len(s) {
        <a id="L99"></a>a[n] = string(s[last:len(s)]);
        <a id="L100"></a>n++;
    <a id="L101"></a>}
    <a id="L102"></a>return a[0:n];
<a id="L103"></a>}

<a id="L105"></a>func getFields(s string) []string { return splitAtBytes(s, &#34; \r\t\n&#34;) }

<a id="L107"></a><span class="comment">// Bigger than we need, not too big to worry about overflow</span>
<a id="L108"></a>const big = 0xFFFFFF

<a id="L110"></a><span class="comment">// Decimal to integer starting at &amp;s[i0].</span>
<a id="L111"></a><span class="comment">// Returns number, new offset, success.</span>
<a id="L112"></a>func dtoi(s string, i0 int) (n int, i int, ok bool) {
    <a id="L113"></a>n = 0;
    <a id="L114"></a>for i = i0; i &lt; len(s) &amp;&amp; &#39;0&#39; &lt;= s[i] &amp;&amp; s[i] &lt;= &#39;9&#39;; i++ {
        <a id="L115"></a>n = n*10 + int(s[i]-&#39;0&#39;);
        <a id="L116"></a>if n &gt;= big {
            <a id="L117"></a>return 0, i, false
        <a id="L118"></a>}
    <a id="L119"></a>}
    <a id="L120"></a>if i == i0 {
        <a id="L121"></a>return 0, i, false
    <a id="L122"></a>}
    <a id="L123"></a>return n, i, true;
<a id="L124"></a>}

<a id="L126"></a><span class="comment">// Hexadecimal to integer starting at &amp;s[i0].</span>
<a id="L127"></a><span class="comment">// Returns number, new offset, success.</span>
<a id="L128"></a>func xtoi(s string, i0 int) (n int, i int, ok bool) {
    <a id="L129"></a>n = 0;
    <a id="L130"></a>for i = i0; i &lt; len(s); i++ {
        <a id="L131"></a>if &#39;0&#39; &lt;= s[i] &amp;&amp; s[i] &lt;= &#39;9&#39; {
            <a id="L132"></a>n *= 16;
            <a id="L133"></a>n += int(s[i] - &#39;0&#39;);
        <a id="L134"></a>} else if &#39;a&#39; &lt;= s[i] &amp;&amp; s[i] &lt;= &#39;f&#39; {
            <a id="L135"></a>n *= 16;
            <a id="L136"></a>n += int(s[i]-&#39;a&#39;) + 10;
        <a id="L137"></a>} else if &#39;A&#39; &lt;= s[i] &amp;&amp; s[i] &lt;= &#39;F&#39; {
            <a id="L138"></a>n *= 16;
            <a id="L139"></a>n += int(s[i]-&#39;A&#39;) + 10;
        <a id="L140"></a>} else {
            <a id="L141"></a>break
        <a id="L142"></a>}
        <a id="L143"></a>if n &gt;= big {
            <a id="L144"></a>return 0, i, false
        <a id="L145"></a>}
    <a id="L146"></a>}
    <a id="L147"></a>if i == i0 {
        <a id="L148"></a>return 0, i, false
    <a id="L149"></a>}
    <a id="L150"></a>return n, i, true;
<a id="L151"></a>}

<a id="L153"></a><span class="comment">// Integer to decimal.</span>
<a id="L154"></a>func itoa(i int) string {
    <a id="L155"></a>var buf [30]byte;
    <a id="L156"></a>n := len(buf);
    <a id="L157"></a>neg := false;
    <a id="L158"></a>if i &lt; 0 {
        <a id="L159"></a>i = -i;
        <a id="L160"></a>neg = true;
    <a id="L161"></a>}
    <a id="L162"></a>ui := uint(i);
    <a id="L163"></a>for ui &gt; 0 || n == len(buf) {
        <a id="L164"></a>n--;
        <a id="L165"></a>buf[n] = byte(&#39;0&#39; + ui%10);
        <a id="L166"></a>ui /= 10;
    <a id="L167"></a>}
    <a id="L168"></a>if neg {
        <a id="L169"></a>n--;
        <a id="L170"></a>buf[n] = &#39;-&#39;;
    <a id="L171"></a>}
    <a id="L172"></a>return string(buf[n:len(buf)]);
<a id="L173"></a>}

<a id="L175"></a><span class="comment">// Number of occurrences of b in s.</span>
<a id="L176"></a>func count(s string, b byte) int {
    <a id="L177"></a>n := 0;
    <a id="L178"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L179"></a>if s[i] == b {
            <a id="L180"></a>n++
        <a id="L181"></a>}
    <a id="L182"></a>}
    <a id="L183"></a>return n;
<a id="L184"></a>}

<a id="L186"></a><span class="comment">// Index of rightmost occurrence of b in s.</span>
<a id="L187"></a>func last(s string, b byte) int {
    <a id="L188"></a>i := len(s);
    <a id="L189"></a>for i--; i &gt;= 0; i-- {
        <a id="L190"></a>if s[i] == b {
            <a id="L191"></a>break
        <a id="L192"></a>}
    <a id="L193"></a>}
    <a id="L194"></a>return i;
<a id="L195"></a>}
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
