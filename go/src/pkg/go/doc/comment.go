<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/doc/comment.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/go/doc/comment.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Godoc comment extraction and comment -&gt; HTML formatting.</span>

<a id="L7"></a>package doc

<a id="L9"></a>import (
    <a id="L10"></a>&#34;go/ast&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;strings&#34;;
    <a id="L13"></a>&#34;template&#34;; <span class="comment">// for htmlEscape</span>
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// Comment extraction</span>

<a id="L18"></a><span class="comment">// CommentText returns the text of comment,</span>
<a id="L19"></a><span class="comment">// with the comment markers - //, /*, and */ - removed.</span>
<a id="L20"></a>func CommentText(comment *ast.CommentGroup) string {
    <a id="L21"></a>if comment == nil {
        <a id="L22"></a>return &#34;&#34;
    <a id="L23"></a>}
    <a id="L24"></a>comments := make([]string, len(comment.List));
    <a id="L25"></a>for i, c := range comment.List {
        <a id="L26"></a>comments[i] = string(c.Text)
    <a id="L27"></a>}

    <a id="L29"></a>lines := make([]string, 0, 20);
    <a id="L30"></a>for _, c := range comments {
        <a id="L31"></a><span class="comment">// Remove comment markers.</span>
        <a id="L32"></a><span class="comment">// The parser has given us exactly the comment text.</span>
        <a id="L33"></a>switch n := len(c); {
        <a id="L34"></a>case n &gt;= 4 &amp;&amp; c[0:2] == &#34;/*&#34; &amp;&amp; c[n-2:n] == &#34;*/&#34;:
            <a id="L35"></a>c = c[2 : n-2]
        <a id="L36"></a>case n &gt;= 2 &amp;&amp; c[0:2] == &#34;//&#34;:
            <a id="L37"></a>c = c[2:n];
            <a id="L38"></a><span class="comment">// Remove leading space after //, if there is one.</span>
            <a id="L39"></a>if len(c) &gt; 0 &amp;&amp; c[0] == &#39; &#39; {
                <a id="L40"></a>c = c[1:len(c)]
            <a id="L41"></a>}
        <a id="L42"></a>}

        <a id="L44"></a><span class="comment">// Split on newlines.</span>
        <a id="L45"></a>cl := strings.Split(c, &#34;\n&#34;, 0);

        <a id="L47"></a><span class="comment">// Walk lines, stripping trailing white space and adding to list.</span>
        <a id="L48"></a>for _, l := range cl {
            <a id="L49"></a><span class="comment">// Strip trailing white space</span>
            <a id="L50"></a>m := len(l);
            <a id="L51"></a>for m &gt; 0 &amp;&amp; (l[m-1] == &#39; &#39; || l[m-1] == &#39;\n&#39; || l[m-1] == &#39;\t&#39; || l[m-1] == &#39;\r&#39;) {
                <a id="L52"></a>m--
            <a id="L53"></a>}
            <a id="L54"></a>l = l[0:m];

            <a id="L56"></a><span class="comment">// Add to list.</span>
            <a id="L57"></a>n := len(lines);
            <a id="L58"></a>if n+1 &gt;= cap(lines) {
                <a id="L59"></a>newlines := make([]string, n, 2*cap(lines));
                <a id="L60"></a>for k := range newlines {
                    <a id="L61"></a>newlines[k] = lines[k]
                <a id="L62"></a>}
                <a id="L63"></a>lines = newlines;
            <a id="L64"></a>}
            <a id="L65"></a>lines = lines[0 : n+1];
            <a id="L66"></a>lines[n] = l;
        <a id="L67"></a>}
    <a id="L68"></a>}

    <a id="L70"></a><span class="comment">// Remove leading blank lines; convert runs of</span>
    <a id="L71"></a><span class="comment">// interior blank lines to a single blank line.</span>
    <a id="L72"></a>n := 0;
    <a id="L73"></a>for _, line := range lines {
        <a id="L74"></a>if line != &#34;&#34; || n &gt; 0 &amp;&amp; lines[n-1] != &#34;&#34; {
            <a id="L75"></a>lines[n] = line;
            <a id="L76"></a>n++;
        <a id="L77"></a>}
    <a id="L78"></a>}
    <a id="L79"></a>lines = lines[0:n];

    <a id="L81"></a><span class="comment">// Add final &#34;&#34; entry to get trailing newline from Join.</span>
    <a id="L82"></a><span class="comment">// The original loop always leaves room for one more.</span>
    <a id="L83"></a>if n &gt; 0 &amp;&amp; lines[n-1] != &#34;&#34; {
        <a id="L84"></a>lines = lines[0 : n+1];
        <a id="L85"></a>lines[n] = &#34;&#34;;
    <a id="L86"></a>}

    <a id="L88"></a>return strings.Join(lines, &#34;\n&#34;);
<a id="L89"></a>}

<a id="L91"></a><span class="comment">// Split bytes into lines.</span>
<a id="L92"></a>func split(text []byte) [][]byte {
    <a id="L93"></a><span class="comment">// count lines</span>
    <a id="L94"></a>n := 0;
    <a id="L95"></a>last := 0;
    <a id="L96"></a>for i, c := range text {
        <a id="L97"></a>if c == &#39;\n&#39; {
            <a id="L98"></a>last = i + 1;
            <a id="L99"></a>n++;
        <a id="L100"></a>}
    <a id="L101"></a>}
    <a id="L102"></a>if last &lt; len(text) {
        <a id="L103"></a>n++
    <a id="L104"></a>}

    <a id="L106"></a><span class="comment">// split</span>
    <a id="L107"></a>out := make([][]byte, n);
    <a id="L108"></a>last = 0;
    <a id="L109"></a>n = 0;
    <a id="L110"></a>for i, c := range text {
        <a id="L111"></a>if c == &#39;\n&#39; {
            <a id="L112"></a>out[n] = text[last : i+1];
            <a id="L113"></a>last = i + 1;
            <a id="L114"></a>n++;
        <a id="L115"></a>}
    <a id="L116"></a>}
    <a id="L117"></a>if last &lt; len(text) {
        <a id="L118"></a>out[n] = text[last:len(text)]
    <a id="L119"></a>}

    <a id="L121"></a>return out;
<a id="L122"></a>}


<a id="L125"></a>var (
    <a id="L126"></a>ldquo = strings.Bytes(&#34;&amp;ldquo;&#34;);
    <a id="L127"></a>rdquo = strings.Bytes(&#34;&amp;rdquo;&#34;);
<a id="L128"></a>)

<a id="L130"></a><span class="comment">// Escape comment text for HTML.</span>
<a id="L131"></a><span class="comment">// Also, turn `` into &amp;ldquo; and &#39;&#39; into &amp;rdquo;.</span>
<a id="L132"></a>func commentEscape(w io.Writer, s []byte) {
    <a id="L133"></a>last := 0;
    <a id="L134"></a>for i := 0; i &lt; len(s)-1; i++ {
        <a id="L135"></a>if s[i] == s[i+1] &amp;&amp; (s[i] == &#39;`&#39; || s[i] == &#39;\&#39;&#39;) {
            <a id="L136"></a>template.HTMLEscape(w, s[last:i]);
            <a id="L137"></a>last = i + 2;
            <a id="L138"></a>switch s[i] {
            <a id="L139"></a>case &#39;`&#39;:
                <a id="L140"></a>w.Write(ldquo)
            <a id="L141"></a>case &#39;\&#39;&#39;:
                <a id="L142"></a>w.Write(rdquo)
            <a id="L143"></a>}
            <a id="L144"></a>i++; <span class="comment">// loop will add one more</span>
        <a id="L145"></a>}
    <a id="L146"></a>}
    <a id="L147"></a>template.HTMLEscape(w, s[last:len(s)]);
<a id="L148"></a>}


<a id="L151"></a>var (
    <a id="L152"></a>html_p      = strings.Bytes(&#34;&lt;p&gt;\n&#34;);
    <a id="L153"></a>html_endp   = strings.Bytes(&#34;&lt;/p&gt;\n&#34;);
    <a id="L154"></a>html_pre    = strings.Bytes(&#34;&lt;pre&gt;&#34;);
    <a id="L155"></a>html_endpre = strings.Bytes(&#34;&lt;/pre&gt;\n&#34;);
<a id="L156"></a>)


<a id="L159"></a>func indentLen(s []byte) int {
    <a id="L160"></a>i := 0;
    <a id="L161"></a>for i &lt; len(s) &amp;&amp; (s[i] == &#39; &#39; || s[i] == &#39;\t&#39;) {
        <a id="L162"></a>i++
    <a id="L163"></a>}
    <a id="L164"></a>return i;
<a id="L165"></a>}


<a id="L168"></a>func isBlank(s []byte) bool { return len(s) == 0 || (len(s) == 1 &amp;&amp; s[0] == &#39;\n&#39;) }


<a id="L171"></a>func commonPrefix(a, b []byte) []byte {
    <a id="L172"></a>i := 0;
    <a id="L173"></a>for i &lt; len(a) &amp;&amp; i &lt; len(b) &amp;&amp; a[i] == b[i] {
        <a id="L174"></a>i++
    <a id="L175"></a>}
    <a id="L176"></a>return a[0:i];
<a id="L177"></a>}


<a id="L180"></a>func unindent(block [][]byte) {
    <a id="L181"></a>if len(block) == 0 {
        <a id="L182"></a>return
    <a id="L183"></a>}

    <a id="L185"></a><span class="comment">// compute maximum common white prefix</span>
    <a id="L186"></a>prefix := block[0][0:indentLen(block[0])];
    <a id="L187"></a>for _, line := range block {
        <a id="L188"></a>if !isBlank(line) {
            <a id="L189"></a>prefix = commonPrefix(prefix, line[0:indentLen(line)])
        <a id="L190"></a>}
    <a id="L191"></a>}
    <a id="L192"></a>n := len(prefix);

    <a id="L194"></a><span class="comment">// remove</span>
    <a id="L195"></a>for i, line := range block {
        <a id="L196"></a>if !isBlank(line) {
            <a id="L197"></a>block[i] = line[n:len(line)]
        <a id="L198"></a>}
    <a id="L199"></a>}
<a id="L200"></a>}


<a id="L203"></a><span class="comment">// Convert comment text to formatted HTML.</span>
<a id="L204"></a><span class="comment">// The comment was prepared by DocReader,</span>
<a id="L205"></a><span class="comment">// so it is known not to have leading, trailing blank lines</span>
<a id="L206"></a><span class="comment">// nor to have trailing spaces at the end of lines.</span>
<a id="L207"></a><span class="comment">// The comment markers have already been removed.</span>
<a id="L208"></a><span class="comment">//</span>
<a id="L209"></a><span class="comment">// Turn each run of multiple \n into &lt;/p&gt;&lt;p&gt;</span>
<a id="L210"></a><span class="comment">// Turn each run of indented lines into &lt;pre&gt; without indent.</span>
<a id="L211"></a><span class="comment">//</span>
<a id="L212"></a><span class="comment">// TODO(rsc): I&#39;d like to pass in an array of variable names []string</span>
<a id="L213"></a><span class="comment">// and then italicize those strings when they appear as words.</span>
<a id="L214"></a>func ToHTML(w io.Writer, s []byte) {
    <a id="L215"></a>inpara := false;

    <a id="L217"></a>close := func() {
        <a id="L218"></a>if inpara {
            <a id="L219"></a>w.Write(html_endp);
            <a id="L220"></a>inpara = false;
        <a id="L221"></a>}
    <a id="L222"></a>};
    <a id="L223"></a>open := func() {
        <a id="L224"></a>if !inpara {
            <a id="L225"></a>w.Write(html_p);
            <a id="L226"></a>inpara = true;
        <a id="L227"></a>}
    <a id="L228"></a>};

    <a id="L230"></a>lines := split(s);
    <a id="L231"></a>unindent(lines);
    <a id="L232"></a>for i := 0; i &lt; len(lines); {
        <a id="L233"></a>line := lines[i];
        <a id="L234"></a>if isBlank(line) {
            <a id="L235"></a><span class="comment">// close paragraph</span>
            <a id="L236"></a>close();
            <a id="L237"></a>i++;
            <a id="L238"></a>continue;
        <a id="L239"></a>}
        <a id="L240"></a>if indentLen(line) &gt; 0 {
            <a id="L241"></a><span class="comment">// close paragraph</span>
            <a id="L242"></a>close();

            <a id="L244"></a><span class="comment">// count indented or blank lines</span>
            <a id="L245"></a>j := i + 1;
            <a id="L246"></a>for j &lt; len(lines) &amp;&amp; (isBlank(lines[j]) || indentLen(lines[j]) &gt; 0) {
                <a id="L247"></a>j++
            <a id="L248"></a>}
            <a id="L249"></a><span class="comment">// but not trailing blank lines</span>
            <a id="L250"></a>for j &gt; i &amp;&amp; isBlank(lines[j-1]) {
                <a id="L251"></a>j--
            <a id="L252"></a>}
            <a id="L253"></a>block := lines[i:j];
            <a id="L254"></a>i = j;

            <a id="L256"></a>unindent(block);

            <a id="L258"></a><span class="comment">// put those lines in a pre block.</span>
            <a id="L259"></a><span class="comment">// they don&#39;t get the nice text formatting,</span>
            <a id="L260"></a><span class="comment">// just html escaping</span>
            <a id="L261"></a>w.Write(html_pre);
            <a id="L262"></a>for _, line := range block {
                <a id="L263"></a>template.HTMLEscape(w, line)
            <a id="L264"></a>}
            <a id="L265"></a>w.Write(html_endpre);
            <a id="L266"></a>continue;
        <a id="L267"></a>}
        <a id="L268"></a><span class="comment">// open paragraph</span>
        <a id="L269"></a>open();
        <a id="L270"></a>commentEscape(w, lines[i]);
        <a id="L271"></a>i++;
    <a id="L272"></a>}
    <a id="L273"></a>close();
<a id="L274"></a>}
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
