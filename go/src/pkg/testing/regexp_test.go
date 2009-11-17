<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/testing/regexp_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/testing/regexp_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package testing

<a id="L7"></a>import (
    <a id="L8"></a>&#34;strings&#34;;
<a id="L9"></a>)

<a id="L11"></a>var good_re = []string{
    <a id="L12"></a>``,
    <a id="L13"></a>`.`,
    <a id="L14"></a>`^.$`,
    <a id="L15"></a>`a`,
    <a id="L16"></a>`a*`,
    <a id="L17"></a>`a+`,
    <a id="L18"></a>`a?`,
    <a id="L19"></a>`a|b`,
    <a id="L20"></a>`a*|b*`,
    <a id="L21"></a>`(a*|b)(c*|d)`,
    <a id="L22"></a>`[a-z]`,
    <a id="L23"></a>`[a-abc-c\-\]\[]`,
    <a id="L24"></a>`[a-z]+`,
    <a id="L25"></a>`[]`,
    <a id="L26"></a>`[abc]`,
    <a id="L27"></a>`[^1234]`,
    <a id="L28"></a>`[^\n]`,
<a id="L29"></a>}

<a id="L31"></a><span class="comment">// TODO: nice to do this with a map</span>
<a id="L32"></a>type stringError struct {
    <a id="L33"></a>re  string;
    <a id="L34"></a>err string;
<a id="L35"></a>}

<a id="L37"></a>var bad_re = []stringError{
    <a id="L38"></a>stringError{`*`, ErrBareClosure},
    <a id="L39"></a>stringError{`(abc`, ErrUnmatchedLpar},
    <a id="L40"></a>stringError{`abc)`, ErrUnmatchedRpar},
    <a id="L41"></a>stringError{`x[a-z`, ErrUnmatchedLbkt},
    <a id="L42"></a>stringError{`abc]`, ErrUnmatchedRbkt},
    <a id="L43"></a>stringError{`[z-a]`, ErrBadRange},
    <a id="L44"></a>stringError{`abc\`, ErrExtraneousBackslash},
    <a id="L45"></a>stringError{`a**`, ErrBadClosure},
    <a id="L46"></a>stringError{`a*+`, ErrBadClosure},
    <a id="L47"></a>stringError{`a??`, ErrBadClosure},
    <a id="L48"></a>stringError{`*`, ErrBareClosure},
    <a id="L49"></a>stringError{`\x`, ErrBadBackslash},
<a id="L50"></a>}

<a id="L52"></a>type vec []int

<a id="L54"></a>type tester struct {
    <a id="L55"></a>re    string;
    <a id="L56"></a>text  string;
    <a id="L57"></a>match vec;
<a id="L58"></a>}

<a id="L60"></a>var matches = []tester{
    <a id="L61"></a>tester{``, &#34;&#34;, vec{0, 0}},
    <a id="L62"></a>tester{`a`, &#34;a&#34;, vec{0, 1}},
    <a id="L63"></a>tester{`x`, &#34;y&#34;, vec{}},
    <a id="L64"></a>tester{`b`, &#34;abc&#34;, vec{1, 2}},
    <a id="L65"></a>tester{`.`, &#34;a&#34;, vec{0, 1}},
    <a id="L66"></a>tester{`.*`, &#34;abcdef&#34;, vec{0, 6}},
    <a id="L67"></a>tester{`^abcd$`, &#34;abcd&#34;, vec{0, 4}},
    <a id="L68"></a>tester{`^bcd&#39;`, &#34;abcdef&#34;, vec{}},
    <a id="L69"></a>tester{`^abcd$`, &#34;abcde&#34;, vec{}},
    <a id="L70"></a>tester{`a+`, &#34;baaab&#34;, vec{1, 4}},
    <a id="L71"></a>tester{`a*`, &#34;baaab&#34;, vec{0, 0}},
    <a id="L72"></a>tester{`[a-z]+`, &#34;abcd&#34;, vec{0, 4}},
    <a id="L73"></a>tester{`[^a-z]+`, &#34;ab1234cd&#34;, vec{2, 6}},
    <a id="L74"></a>tester{`[a\-\]z]+`, &#34;az]-bcz&#34;, vec{0, 4}},
    <a id="L75"></a>tester{`[^\n]+`, &#34;abcd\n&#34;, vec{0, 4}},
    <a id="L76"></a>tester{`[日本語]+`, &#34;日本語日本語&#34;, vec{0, 18}},
    <a id="L77"></a>tester{`()`, &#34;&#34;, vec{0, 0, 0, 0}},
    <a id="L78"></a>tester{`(a)`, &#34;a&#34;, vec{0, 1, 0, 1}},
    <a id="L79"></a>tester{`(.)(.)`, &#34;日a&#34;, vec{0, 4, 0, 3, 3, 4}},
    <a id="L80"></a>tester{`(.*)`, &#34;&#34;, vec{0, 0, 0, 0}},
    <a id="L81"></a>tester{`(.*)`, &#34;abcd&#34;, vec{0, 4, 0, 4}},
    <a id="L82"></a>tester{`(..)(..)`, &#34;abcd&#34;, vec{0, 4, 0, 2, 2, 4}},
    <a id="L83"></a>tester{`(([^xyz]*)(d))`, &#34;abcd&#34;, vec{0, 4, 0, 4, 0, 3, 3, 4}},
    <a id="L84"></a>tester{`((a|b|c)*(d))`, &#34;abcd&#34;, vec{0, 4, 0, 4, 2, 3, 3, 4}},
    <a id="L85"></a>tester{`(((a|b|c)*)(d))`, &#34;abcd&#34;, vec{0, 4, 0, 4, 0, 3, 2, 3, 3, 4}},
    <a id="L86"></a>tester{`a*(|(b))c*`, &#34;aacc&#34;, vec{0, 4, 2, 2, -1, -1}},
<a id="L87"></a>}

<a id="L89"></a>func compileTest(t *T, expr string, error string) *Regexp {
    <a id="L90"></a>re, err := CompileRegexp(expr);
    <a id="L91"></a>if err != error {
        <a id="L92"></a>t.Error(&#34;compiling `&#34;, expr, &#34;`; unexpected error: &#34;, err)
    <a id="L93"></a>}
    <a id="L94"></a>return re;
<a id="L95"></a>}

<a id="L97"></a>func printVec(t *T, m []int) {
    <a id="L98"></a>l := len(m);
    <a id="L99"></a>if l == 0 {
        <a id="L100"></a>t.Log(&#34;\t&lt;no match&gt;&#34;)
    <a id="L101"></a>} else {
        <a id="L102"></a>for i := 0; i &lt; l; i = i + 2 {
            <a id="L103"></a>t.Log(&#34;\t&#34;, m[i], &#34;,&#34;, m[i+1])
        <a id="L104"></a>}
    <a id="L105"></a>}
<a id="L106"></a>}

<a id="L108"></a>func printStrings(t *T, m []string) {
    <a id="L109"></a>l := len(m);
    <a id="L110"></a>if l == 0 {
        <a id="L111"></a>t.Log(&#34;\t&lt;no match&gt;&#34;)
    <a id="L112"></a>} else {
        <a id="L113"></a>for i := 0; i &lt; l; i = i + 2 {
            <a id="L114"></a>t.Logf(&#34;\t%q&#34;, m[i])
        <a id="L115"></a>}
    <a id="L116"></a>}
<a id="L117"></a>}

<a id="L119"></a>func printBytes(t *T, b [][]byte) {
    <a id="L120"></a>l := len(b);
    <a id="L121"></a>if l == 0 {
        <a id="L122"></a>t.Log(&#34;\t&lt;no match&gt;&#34;)
    <a id="L123"></a>} else {
        <a id="L124"></a>for i := 0; i &lt; l; i = i + 2 {
            <a id="L125"></a>t.Logf(&#34;\t%q&#34;, b[i])
        <a id="L126"></a>}
    <a id="L127"></a>}
<a id="L128"></a>}

<a id="L130"></a>func equal(m1, m2 []int) bool {
    <a id="L131"></a>l := len(m1);
    <a id="L132"></a>if l != len(m2) {
        <a id="L133"></a>return false
    <a id="L134"></a>}
    <a id="L135"></a>for i := 0; i &lt; l; i++ {
        <a id="L136"></a>if m1[i] != m2[i] {
            <a id="L137"></a>return false
        <a id="L138"></a>}
    <a id="L139"></a>}
    <a id="L140"></a>return true;
<a id="L141"></a>}

<a id="L143"></a>func equalStrings(m1, m2 []string) bool {
    <a id="L144"></a>l := len(m1);
    <a id="L145"></a>if l != len(m2) {
        <a id="L146"></a>return false
    <a id="L147"></a>}
    <a id="L148"></a>for i := 0; i &lt; l; i++ {
        <a id="L149"></a>if m1[i] != m2[i] {
            <a id="L150"></a>return false
        <a id="L151"></a>}
    <a id="L152"></a>}
    <a id="L153"></a>return true;
<a id="L154"></a>}

<a id="L156"></a>func equalBytes(m1 [][]byte, m2 []string) bool {
    <a id="L157"></a>l := len(m1);
    <a id="L158"></a>if l != len(m2) {
        <a id="L159"></a>return false
    <a id="L160"></a>}
    <a id="L161"></a>for i := 0; i &lt; l; i++ {
        <a id="L162"></a>if string(m1[i]) != m2[i] {
            <a id="L163"></a>return false
        <a id="L164"></a>}
    <a id="L165"></a>}
    <a id="L166"></a>return true;
<a id="L167"></a>}

<a id="L169"></a>func executeTest(t *T, expr string, str string, match []int) {
    <a id="L170"></a>re := compileTest(t, expr, &#34;&#34;);
    <a id="L171"></a>if re == nil {
        <a id="L172"></a>return
    <a id="L173"></a>}
    <a id="L174"></a>m := re.ExecuteString(str);
    <a id="L175"></a>if !equal(m, match) {
        <a id="L176"></a>t.Error(&#34;ExecuteString failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;);
        <a id="L177"></a>printVec(t, m);
        <a id="L178"></a>t.Log(&#34;should be:&#34;);
        <a id="L179"></a>printVec(t, match);
    <a id="L180"></a>}
    <a id="L181"></a><span class="comment">// now try bytes</span>
    <a id="L182"></a>m = re.Execute(strings.Bytes(str));
    <a id="L183"></a>if !equal(m, match) {
        <a id="L184"></a>t.Error(&#34;Execute failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;);
        <a id="L185"></a>printVec(t, m);
        <a id="L186"></a>t.Log(&#34;should be:&#34;);
        <a id="L187"></a>printVec(t, match);
    <a id="L188"></a>}
<a id="L189"></a>}

<a id="L191"></a>func TestGoodCompile(t *T) {
    <a id="L192"></a>for i := 0; i &lt; len(good_re); i++ {
        <a id="L193"></a>compileTest(t, good_re[i], &#34;&#34;)
    <a id="L194"></a>}
<a id="L195"></a>}

<a id="L197"></a>func TestBadCompile(t *T) {
    <a id="L198"></a>for i := 0; i &lt; len(bad_re); i++ {
        <a id="L199"></a>compileTest(t, bad_re[i].re, bad_re[i].err)
    <a id="L200"></a>}
<a id="L201"></a>}

<a id="L203"></a>func TestExecute(t *T) {
    <a id="L204"></a>for i := 0; i &lt; len(matches); i++ {
        <a id="L205"></a>test := &amp;matches[i];
        <a id="L206"></a>executeTest(t, test.re, test.text, test.match);
    <a id="L207"></a>}
<a id="L208"></a>}

<a id="L210"></a>func matchTest(t *T, expr string, str string, match []int) {
    <a id="L211"></a>re := compileTest(t, expr, &#34;&#34;);
    <a id="L212"></a>if re == nil {
        <a id="L213"></a>return
    <a id="L214"></a>}
    <a id="L215"></a>m := re.MatchString(str);
    <a id="L216"></a>if m != (len(match) &gt; 0) {
        <a id="L217"></a>t.Error(&#34;MatchString failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;, m, &#34;should be&#34;, len(match) &gt; 0)
    <a id="L218"></a>}
    <a id="L219"></a><span class="comment">// now try bytes</span>
    <a id="L220"></a>m = re.Match(strings.Bytes(str));
    <a id="L221"></a>if m != (len(match) &gt; 0) {
        <a id="L222"></a>t.Error(&#34;Match failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;, m, &#34;should be&#34;, len(match) &gt; 0)
    <a id="L223"></a>}
<a id="L224"></a>}

<a id="L226"></a>func TestMatch(t *T) {
    <a id="L227"></a>for i := 0; i &lt; len(matches); i++ {
        <a id="L228"></a>test := &amp;matches[i];
        <a id="L229"></a>matchTest(t, test.re, test.text, test.match);
    <a id="L230"></a>}
<a id="L231"></a>}

<a id="L233"></a>func matchStringsTest(t *T, expr string, str string, match []int) {
    <a id="L234"></a>re := compileTest(t, expr, &#34;&#34;);
    <a id="L235"></a>if re == nil {
        <a id="L236"></a>return
    <a id="L237"></a>}
    <a id="L238"></a>strs := make([]string, len(match)/2);
    <a id="L239"></a>for i := 0; i &lt; len(match); i++ {
        <a id="L240"></a>strs[i/2] = str[match[i]:match[i+1]]
    <a id="L241"></a>}
    <a id="L242"></a>m := re.MatchStrings(str);
    <a id="L243"></a>if !equalStrings(m, strs) {
        <a id="L244"></a>t.Error(&#34;MatchStrings failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;);
        <a id="L245"></a>printStrings(t, m);
        <a id="L246"></a>t.Log(&#34;should be:&#34;);
        <a id="L247"></a>printStrings(t, strs);
    <a id="L248"></a>}
    <a id="L249"></a><span class="comment">// now try bytes</span>
    <a id="L250"></a>s := re.MatchSlices(strings.Bytes(str));
    <a id="L251"></a>if !equalBytes(s, strs) {
        <a id="L252"></a>t.Error(&#34;MatchSlices failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;);
        <a id="L253"></a>printBytes(t, s);
        <a id="L254"></a>t.Log(&#34;should be:&#34;);
        <a id="L255"></a>printStrings(t, strs);
    <a id="L256"></a>}
<a id="L257"></a>}

<a id="L259"></a>func TestMatchStrings(t *T) {
    <a id="L260"></a>for i := 0; i &lt; len(matches); i++ {
        <a id="L261"></a>test := &amp;matches[i];
        <a id="L262"></a>matchTest(t, test.re, test.text, test.match);
    <a id="L263"></a>}
<a id="L264"></a>}

<a id="L266"></a>func matchFunctionTest(t *T, expr string, str string, match []int) {
    <a id="L267"></a>m, err := MatchString(expr, str);
    <a id="L268"></a>if err == &#34;&#34; {
        <a id="L269"></a>return
    <a id="L270"></a>}
    <a id="L271"></a>if m != (len(match) &gt; 0) {
        <a id="L272"></a>t.Error(&#34;function Match failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;, m, &#34;should be&#34;, len(match) &gt; 0)
    <a id="L273"></a>}
<a id="L274"></a>}

<a id="L276"></a>func TestMatchFunction(t *T) {
    <a id="L277"></a>for i := 0; i &lt; len(matches); i++ {
        <a id="L278"></a>test := &amp;matches[i];
        <a id="L279"></a>matchFunctionTest(t, test.re, test.text, test.match);
    <a id="L280"></a>}
<a id="L281"></a>}
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
