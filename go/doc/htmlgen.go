<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /doc/htmlgen.go</title>

  <link rel="stylesheet" type="text/css" href="style.css">
  <script type="text/javascript" src="godocs.js"></script>

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
        <a href="../index.html"><img src="logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="go_tutorial.html">Tutorial</a></li>
    <li><a href="effective_go.html">Effective Go</a></li>
    <li><a href="go_faq.html">FAQ</a></li>
    <li><a href="go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="go_spec.html">Language Specification</a></li>
    <li><a href="go_mem.html">Memory Model</a></li>
    <li><a href="go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="install.html">Install Go</a></li>
    <li><a href="contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../cmd/index.html">Command documentation</a></li>
    <li><a href="../pkg/index.html">Package documentation</a></li>
    <li><a href="../src/index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /doc/htmlgen.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Process plain text into HTML.</span>
<a id="L6"></a><span class="comment">//	- h2&#39;s are made from lines followed by a line &#34;----\n&#34;</span>
<a id="L7"></a><span class="comment">//	- tab-indented blocks become &lt;pre&gt; blocks</span>
<a id="L8"></a><span class="comment">//	- blank lines become &lt;p&gt; marks</span>
<a id="L9"></a><span class="comment">//	- &#34;quoted strings&#34; become &lt;code&gt;quoted strings&lt;/code&gt;</span>

<a id="L11"></a>package main

<a id="L13"></a>import (
    <a id="L14"></a>&#34;bufio&#34;;
    <a id="L15"></a>&#34;bytes&#34;;
    <a id="L16"></a>&#34;log&#34;;
    <a id="L17"></a>&#34;os&#34;;
    <a id="L18"></a>&#34;strings&#34;;
<a id="L19"></a>)

<a id="L21"></a>var (
    <a id="L22"></a>lines   = make([][]byte, 0, 10000); <span class="comment">// assume big enough</span>
    <a id="L23"></a>linebuf = make([]byte, 10000);      <span class="comment">// assume big enough</span>

    <a id="L25"></a>empty   = strings.Bytes(&#34;&#34;);
    <a id="L26"></a>newline = strings.Bytes(&#34;\n&#34;);
    <a id="L27"></a>tab     = strings.Bytes(&#34;\t&#34;);
    <a id="L28"></a>quote   = strings.Bytes(`&#34;`);

    <a id="L30"></a>sectionMarker = strings.Bytes(&#34;----\n&#34;);
    <a id="L31"></a>preStart      = strings.Bytes(&#34;&lt;pre&gt;&#34;);
    <a id="L32"></a>preEnd        = strings.Bytes(&#34;&lt;/pre&gt;\n&#34;);
    <a id="L33"></a>pp            = strings.Bytes(&#34;&lt;p&gt;\n&#34;);
<a id="L34"></a>)

<a id="L36"></a>func main() {
    <a id="L37"></a>read();
    <a id="L38"></a>headings();
    <a id="L39"></a>paragraphs();
    <a id="L40"></a>coalesce(preStart, foldPre);
    <a id="L41"></a>coalesce(tab, foldTabs);
    <a id="L42"></a>quotes();
    <a id="L43"></a>write();
<a id="L44"></a>}

<a id="L46"></a>func read() {
    <a id="L47"></a>b := bufio.NewReader(os.Stdin);
    <a id="L48"></a>for {
        <a id="L49"></a>line, err := b.ReadBytes(&#39;\n&#39;);
        <a id="L50"></a>if err == os.EOF {
            <a id="L51"></a>break
        <a id="L52"></a>}
        <a id="L53"></a>if err != nil {
            <a id="L54"></a>log.Exit(err)
        <a id="L55"></a>}
        <a id="L56"></a>n := len(lines);
        <a id="L57"></a>lines = lines[0 : n+1];
        <a id="L58"></a>lines[n] = line;
    <a id="L59"></a>}
<a id="L60"></a>}

<a id="L62"></a>func write() {
    <a id="L63"></a>b := bufio.NewWriter(os.Stdout);
    <a id="L64"></a>for _, line := range lines {
        <a id="L65"></a>b.Write(expandTabs(line))
    <a id="L66"></a>}
    <a id="L67"></a>b.Flush();
<a id="L68"></a>}

<a id="L70"></a><span class="comment">// each time prefix is found on a line, call fold and replace</span>
<a id="L71"></a><span class="comment">// line with return value from fold.</span>
<a id="L72"></a>func coalesce(prefix []byte, fold func(i int) (n int, line []byte)) {
    <a id="L73"></a>j := 0; <span class="comment">// output line number; goes up by one each loop</span>
    <a id="L74"></a>for i := 0; i &lt; len(lines); {
        <a id="L75"></a>if bytes.HasPrefix(lines[i], prefix) {
            <a id="L76"></a>nlines, block := fold(i);
            <a id="L77"></a>lines[j] = block;
            <a id="L78"></a>i += nlines;
        <a id="L79"></a>} else {
            <a id="L80"></a>lines[j] = lines[i];
            <a id="L81"></a>i++;
        <a id="L82"></a>}
        <a id="L83"></a>j++;
    <a id="L84"></a>}
    <a id="L85"></a>lines = lines[0:j];
<a id="L86"></a>}

<a id="L88"></a><span class="comment">// return the &lt;pre&gt; block as a single slice</span>
<a id="L89"></a>func foldPre(i int) (n int, line []byte) {
    <a id="L90"></a>buf := new(bytes.Buffer);
    <a id="L91"></a>for i &lt; len(lines) {
        <a id="L92"></a>buf.Write(lines[i]);
        <a id="L93"></a>n++;
        <a id="L94"></a>if bytes.Equal(lines[i], preEnd) {
            <a id="L95"></a>break
        <a id="L96"></a>}
        <a id="L97"></a>i++;
    <a id="L98"></a>}
    <a id="L99"></a>return n, buf.Bytes();
<a id="L100"></a>}

<a id="L102"></a><span class="comment">// return the tab-indented block as a single &lt;pre&gt;-bounded slice</span>
<a id="L103"></a>func foldTabs(i int) (n int, line []byte) {
    <a id="L104"></a>buf := new(bytes.Buffer);
    <a id="L105"></a>buf.WriteString(&#34;&lt;pre&gt;\n&#34;);
    <a id="L106"></a>for i &lt; len(lines) {
        <a id="L107"></a>if !bytes.HasPrefix(lines[i], tab) {
            <a id="L108"></a>break
        <a id="L109"></a>}
        <a id="L110"></a>buf.Write(lines[i]);
        <a id="L111"></a>n++;
        <a id="L112"></a>i++;
    <a id="L113"></a>}
    <a id="L114"></a>buf.WriteString(&#34;&lt;/pre&gt;\n&#34;);
    <a id="L115"></a>return n, buf.Bytes();
<a id="L116"></a>}

<a id="L118"></a>func headings() {
    <a id="L119"></a>b := bufio.NewWriter(os.Stdout);
    <a id="L120"></a>for i, l := range lines {
        <a id="L121"></a>if i &gt; 0 &amp;&amp; bytes.Equal(l, sectionMarker) {
            <a id="L122"></a>lines[i-1] = strings.Bytes(&#34;&lt;h2&gt;&#34; + string(trim(lines[i-1])) + &#34;&lt;/h2&gt;\n&#34;);
            <a id="L123"></a>lines[i] = empty;
        <a id="L124"></a>}
    <a id="L125"></a>}
    <a id="L126"></a>b.Flush();
<a id="L127"></a>}

<a id="L129"></a>func paragraphs() {
    <a id="L130"></a>for i, l := range lines {
        <a id="L131"></a>if bytes.Equal(l, newline) {
            <a id="L132"></a>lines[i] = pp
        <a id="L133"></a>}
    <a id="L134"></a>}
<a id="L135"></a>}

<a id="L137"></a>func quotes() {
    <a id="L138"></a>for i, l := range lines {
        <a id="L139"></a>lines[i] = codeQuotes(l)
    <a id="L140"></a>}
<a id="L141"></a>}

<a id="L143"></a>func codeQuotes(l []byte) []byte {
    <a id="L144"></a>if bytes.HasPrefix(l, preStart) {
        <a id="L145"></a>return l
    <a id="L146"></a>}
    <a id="L147"></a>n := bytes.Index(l, quote);
    <a id="L148"></a>if n &lt; 0 {
        <a id="L149"></a>return l
    <a id="L150"></a>}
    <a id="L151"></a>buf := new(bytes.Buffer);
    <a id="L152"></a>inQuote := false;
    <a id="L153"></a>for _, c := range l {
        <a id="L154"></a>if c == &#39;&#34;&#39; {
            <a id="L155"></a>if inQuote {
                <a id="L156"></a>buf.WriteString(&#34;&lt;/code&gt;&#34;)
            <a id="L157"></a>} else {
                <a id="L158"></a>buf.WriteString(&#34;&lt;code&gt;&#34;)
            <a id="L159"></a>}
            <a id="L160"></a>inQuote = !inQuote;
        <a id="L161"></a>} else {
            <a id="L162"></a>buf.WriteByte(c)
        <a id="L163"></a>}
    <a id="L164"></a>}
    <a id="L165"></a>return buf.Bytes();
<a id="L166"></a>}

<a id="L168"></a><span class="comment">// drop trailing newline</span>
<a id="L169"></a>func trim(l []byte) []byte {
    <a id="L170"></a>n := len(l);
    <a id="L171"></a>if n &gt; 0 &amp;&amp; l[n-1] == &#39;\n&#39; {
        <a id="L172"></a>return l[0 : n-1]
    <a id="L173"></a>}
    <a id="L174"></a>return l;
<a id="L175"></a>}

<a id="L177"></a><span class="comment">// expand tabs to 4 spaces. don&#39;t worry about columns.</span>
<a id="L178"></a>func expandTabs(l []byte) []byte {
    <a id="L179"></a>j := 0; <span class="comment">// position in linebuf.</span>
    <a id="L180"></a>for _, c := range l {
        <a id="L181"></a>if c == &#39;\t&#39; {
            <a id="L182"></a>for k := 0; k &lt; 4; k++ {
                <a id="L183"></a>linebuf[j] = &#39; &#39;;
                <a id="L184"></a>j++;
            <a id="L185"></a>}
        <a id="L186"></a>} else {
            <a id="L187"></a>linebuf[j] = c;
            <a id="L188"></a>j++;
        <a id="L189"></a>}
    <a id="L190"></a>}
    <a id="L191"></a>return linebuf[0:j];
<a id="L192"></a>}
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
