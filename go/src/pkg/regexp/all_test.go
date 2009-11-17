<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/regexp/all_test.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/regexp/all_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package regexp

<a id="L7"></a>import (
    <a id="L8"></a>&#34;os&#34;;
    <a id="L9"></a>&#34;strings&#34;;
    <a id="L10"></a>&#34;testing&#34;;
<a id="L11"></a>)

<a id="L13"></a>var good_re = []string{
    <a id="L14"></a>``,
    <a id="L15"></a>`.`,
    <a id="L16"></a>`^.$`,
    <a id="L17"></a>`a`,
    <a id="L18"></a>`a*`,
    <a id="L19"></a>`a+`,
    <a id="L20"></a>`a?`,
    <a id="L21"></a>`a|b`,
    <a id="L22"></a>`a*|b*`,
    <a id="L23"></a>`(a*|b)(c*|d)`,
    <a id="L24"></a>`[a-z]`,
    <a id="L25"></a>`[a-abc-c\-\]\[]`,
    <a id="L26"></a>`[a-z]+`,
    <a id="L27"></a>`[]`,
    <a id="L28"></a>`[abc]`,
    <a id="L29"></a>`[^1234]`,
    <a id="L30"></a>`[^\n]`,
<a id="L31"></a>}

<a id="L33"></a><span class="comment">// TODO: nice to do this with a map</span>
<a id="L34"></a>type stringError struct {
    <a id="L35"></a>re  string;
    <a id="L36"></a>err os.Error;
<a id="L37"></a>}

<a id="L39"></a>var bad_re = []stringError{
    <a id="L40"></a>stringError{`*`, ErrBareClosure},
    <a id="L41"></a>stringError{`(abc`, ErrUnmatchedLpar},
    <a id="L42"></a>stringError{`abc)`, ErrUnmatchedRpar},
    <a id="L43"></a>stringError{`x[a-z`, ErrUnmatchedLbkt},
    <a id="L44"></a>stringError{`abc]`, ErrUnmatchedRbkt},
    <a id="L45"></a>stringError{`[z-a]`, ErrBadRange},
    <a id="L46"></a>stringError{`abc\`, ErrExtraneousBackslash},
    <a id="L47"></a>stringError{`a**`, ErrBadClosure},
    <a id="L48"></a>stringError{`a*+`, ErrBadClosure},
    <a id="L49"></a>stringError{`a??`, ErrBadClosure},
    <a id="L50"></a>stringError{`*`, ErrBareClosure},
    <a id="L51"></a>stringError{`\x`, ErrBadBackslash},
<a id="L52"></a>}

<a id="L54"></a>type vec []int

<a id="L56"></a>type tester struct {
    <a id="L57"></a>re    string;
    <a id="L58"></a>text  string;
    <a id="L59"></a>match vec;
<a id="L60"></a>}

<a id="L62"></a>var matches = []tester{
    <a id="L63"></a>tester{``, &#34;&#34;, vec{0, 0}},
    <a id="L64"></a>tester{`a`, &#34;a&#34;, vec{0, 1}},
    <a id="L65"></a>tester{`x`, &#34;y&#34;, vec{}},
    <a id="L66"></a>tester{`b`, &#34;abc&#34;, vec{1, 2}},
    <a id="L67"></a>tester{`.`, &#34;a&#34;, vec{0, 1}},
    <a id="L68"></a>tester{`.*`, &#34;abcdef&#34;, vec{0, 6}},
    <a id="L69"></a>tester{`^abcd$`, &#34;abcd&#34;, vec{0, 4}},
    <a id="L70"></a>tester{`^bcd&#39;`, &#34;abcdef&#34;, vec{}},
    <a id="L71"></a>tester{`^abcd$`, &#34;abcde&#34;, vec{}},
    <a id="L72"></a>tester{`a+`, &#34;baaab&#34;, vec{1, 4}},
    <a id="L73"></a>tester{`a*`, &#34;baaab&#34;, vec{0, 0}},
    <a id="L74"></a>tester{`[a-z]+`, &#34;abcd&#34;, vec{0, 4}},
    <a id="L75"></a>tester{`[^a-z]+`, &#34;ab1234cd&#34;, vec{2, 6}},
    <a id="L76"></a>tester{`[a\-\]z]+`, &#34;az]-bcz&#34;, vec{0, 4}},
    <a id="L77"></a>tester{`[^\n]+`, &#34;abcd\n&#34;, vec{0, 4}},
    <a id="L78"></a>tester{`[日本語]+`, &#34;日本語日本語&#34;, vec{0, 18}},
    <a id="L79"></a>tester{`()`, &#34;&#34;, vec{0, 0, 0, 0}},
    <a id="L80"></a>tester{`(a)`, &#34;a&#34;, vec{0, 1, 0, 1}},
    <a id="L81"></a>tester{`(.)(.)`, &#34;日a&#34;, vec{0, 4, 0, 3, 3, 4}},
    <a id="L82"></a>tester{`(.*)`, &#34;&#34;, vec{0, 0, 0, 0}},
    <a id="L83"></a>tester{`(.*)`, &#34;abcd&#34;, vec{0, 4, 0, 4}},
    <a id="L84"></a>tester{`(..)(..)`, &#34;abcd&#34;, vec{0, 4, 0, 2, 2, 4}},
    <a id="L85"></a>tester{`(([^xyz]*)(d))`, &#34;abcd&#34;, vec{0, 4, 0, 4, 0, 3, 3, 4}},
    <a id="L86"></a>tester{`((a|b|c)*(d))`, &#34;abcd&#34;, vec{0, 4, 0, 4, 2, 3, 3, 4}},
    <a id="L87"></a>tester{`(((a|b|c)*)(d))`, &#34;abcd&#34;, vec{0, 4, 0, 4, 0, 3, 2, 3, 3, 4}},
    <a id="L88"></a>tester{`a*(|(b))c*`, &#34;aacc&#34;, vec{0, 4, 2, 2, -1, -1}},
<a id="L89"></a>}

<a id="L91"></a>func compileTest(t *testing.T, expr string, error os.Error) *Regexp {
    <a id="L92"></a>re, err := Compile(expr);
    <a id="L93"></a>if err != error {
        <a id="L94"></a>t.Error(&#34;compiling `&#34;, expr, &#34;`; unexpected error: &#34;, err.String())
    <a id="L95"></a>}
    <a id="L96"></a>return re;
<a id="L97"></a>}

<a id="L99"></a>func printVec(t *testing.T, m []int) {
    <a id="L100"></a>l := len(m);
    <a id="L101"></a>if l == 0 {
        <a id="L102"></a>t.Log(&#34;\t&lt;no match&gt;&#34;)
    <a id="L103"></a>} else {
        <a id="L104"></a>for i := 0; i &lt; l; i = i + 2 {
            <a id="L105"></a>t.Log(&#34;\t&#34;, m[i], &#34;,&#34;, m[i+1])
        <a id="L106"></a>}
    <a id="L107"></a>}
<a id="L108"></a>}

<a id="L110"></a>func printStrings(t *testing.T, m []string) {
    <a id="L111"></a>l := len(m);
    <a id="L112"></a>if l == 0 {
        <a id="L113"></a>t.Log(&#34;\t&lt;no match&gt;&#34;)
    <a id="L114"></a>} else {
        <a id="L115"></a>for i := 0; i &lt; l; i = i + 2 {
            <a id="L116"></a>t.Logf(&#34;\t%q&#34;, m[i])
        <a id="L117"></a>}
    <a id="L118"></a>}
<a id="L119"></a>}

<a id="L121"></a>func printBytes(t *testing.T, b [][]byte) {
    <a id="L122"></a>l := len(b);
    <a id="L123"></a>if l == 0 {
        <a id="L124"></a>t.Log(&#34;\t&lt;no match&gt;&#34;)
    <a id="L125"></a>} else {
        <a id="L126"></a>for i := 0; i &lt; l; i = i + 2 {
            <a id="L127"></a>t.Logf(&#34;\t%q&#34;, b[i])
        <a id="L128"></a>}
    <a id="L129"></a>}
<a id="L130"></a>}

<a id="L132"></a>func equal(m1, m2 []int) bool {
    <a id="L133"></a>l := len(m1);
    <a id="L134"></a>if l != len(m2) {
        <a id="L135"></a>return false
    <a id="L136"></a>}
    <a id="L137"></a>for i := 0; i &lt; l; i++ {
        <a id="L138"></a>if m1[i] != m2[i] {
            <a id="L139"></a>return false
        <a id="L140"></a>}
    <a id="L141"></a>}
    <a id="L142"></a>return true;
<a id="L143"></a>}

<a id="L145"></a>func equalStrings(m1, m2 []string) bool {
    <a id="L146"></a>l := len(m1);
    <a id="L147"></a>if l != len(m2) {
        <a id="L148"></a>return false
    <a id="L149"></a>}
    <a id="L150"></a>for i := 0; i &lt; l; i++ {
        <a id="L151"></a>if m1[i] != m2[i] {
            <a id="L152"></a>return false
        <a id="L153"></a>}
    <a id="L154"></a>}
    <a id="L155"></a>return true;
<a id="L156"></a>}

<a id="L158"></a>func equalBytes(m1 [][]byte, m2 []string) bool {
    <a id="L159"></a>l := len(m1);
    <a id="L160"></a>if l != len(m2) {
        <a id="L161"></a>return false
    <a id="L162"></a>}
    <a id="L163"></a>for i := 0; i &lt; l; i++ {
        <a id="L164"></a>if string(m1[i]) != m2[i] {
            <a id="L165"></a>return false
        <a id="L166"></a>}
    <a id="L167"></a>}
    <a id="L168"></a>return true;
<a id="L169"></a>}

<a id="L171"></a>func executeTest(t *testing.T, expr string, str string, match []int) {
    <a id="L172"></a>re := compileTest(t, expr, nil);
    <a id="L173"></a>if re == nil {
        <a id="L174"></a>return
    <a id="L175"></a>}
    <a id="L176"></a>m := re.ExecuteString(str);
    <a id="L177"></a>if !equal(m, match) {
        <a id="L178"></a>t.Error(&#34;ExecuteString failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;);
        <a id="L179"></a>printVec(t, m);
        <a id="L180"></a>t.Log(&#34;should be:&#34;);
        <a id="L181"></a>printVec(t, match);
    <a id="L182"></a>}
    <a id="L183"></a><span class="comment">// now try bytes</span>
    <a id="L184"></a>m = re.Execute(strings.Bytes(str));
    <a id="L185"></a>if !equal(m, match) {
        <a id="L186"></a>t.Error(&#34;Execute failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;);
        <a id="L187"></a>printVec(t, m);
        <a id="L188"></a>t.Log(&#34;should be:&#34;);
        <a id="L189"></a>printVec(t, match);
    <a id="L190"></a>}
<a id="L191"></a>}

<a id="L193"></a>func TestGoodCompile(t *testing.T) {
    <a id="L194"></a>for i := 0; i &lt; len(good_re); i++ {
        <a id="L195"></a>compileTest(t, good_re[i], nil)
    <a id="L196"></a>}
<a id="L197"></a>}

<a id="L199"></a>func TestBadCompile(t *testing.T) {
    <a id="L200"></a>for i := 0; i &lt; len(bad_re); i++ {
        <a id="L201"></a>compileTest(t, bad_re[i].re, bad_re[i].err)
    <a id="L202"></a>}
<a id="L203"></a>}

<a id="L205"></a>func TestExecute(t *testing.T) {
    <a id="L206"></a>for i := 0; i &lt; len(matches); i++ {
        <a id="L207"></a>test := &amp;matches[i];
        <a id="L208"></a>executeTest(t, test.re, test.text, test.match);
    <a id="L209"></a>}
<a id="L210"></a>}

<a id="L212"></a>func matchTest(t *testing.T, expr string, str string, match []int) {
    <a id="L213"></a>re := compileTest(t, expr, nil);
    <a id="L214"></a>if re == nil {
        <a id="L215"></a>return
    <a id="L216"></a>}
    <a id="L217"></a>m := re.MatchString(str);
    <a id="L218"></a>if m != (len(match) &gt; 0) {
        <a id="L219"></a>t.Error(&#34;MatchString failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;, m, &#34;should be&#34;, len(match) &gt; 0)
    <a id="L220"></a>}
    <a id="L221"></a><span class="comment">// now try bytes</span>
    <a id="L222"></a>m = re.Match(strings.Bytes(str));
    <a id="L223"></a>if m != (len(match) &gt; 0) {
        <a id="L224"></a>t.Error(&#34;Match failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;, m, &#34;should be&#34;, len(match) &gt; 0)
    <a id="L225"></a>}
<a id="L226"></a>}

<a id="L228"></a>func TestMatch(t *testing.T) {
    <a id="L229"></a>for i := 0; i &lt; len(matches); i++ {
        <a id="L230"></a>test := &amp;matches[i];
        <a id="L231"></a>matchTest(t, test.re, test.text, test.match);
    <a id="L232"></a>}
<a id="L233"></a>}

<a id="L235"></a>func matchStringsTest(t *testing.T, expr string, str string, match []int) {
    <a id="L236"></a>re := compileTest(t, expr, nil);
    <a id="L237"></a>if re == nil {
        <a id="L238"></a>return
    <a id="L239"></a>}
    <a id="L240"></a>strs := make([]string, len(match)/2);
    <a id="L241"></a>for i := 0; i &lt; len(match); i++ {
        <a id="L242"></a>strs[i/2] = str[match[i]:match[i+1]]
    <a id="L243"></a>}
    <a id="L244"></a>m := re.MatchStrings(str);
    <a id="L245"></a>if !equalStrings(m, strs) {
        <a id="L246"></a>t.Error(&#34;MatchStrings failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;);
        <a id="L247"></a>printStrings(t, m);
        <a id="L248"></a>t.Log(&#34;should be:&#34;);
        <a id="L249"></a>printStrings(t, strs);
    <a id="L250"></a>}
    <a id="L251"></a><span class="comment">// now try bytes</span>
    <a id="L252"></a>s := re.MatchSlices(strings.Bytes(str));
    <a id="L253"></a>if !equalBytes(s, strs) {
        <a id="L254"></a>t.Error(&#34;MatchSlices failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;);
        <a id="L255"></a>printBytes(t, s);
        <a id="L256"></a>t.Log(&#34;should be:&#34;);
        <a id="L257"></a>printStrings(t, strs);
    <a id="L258"></a>}
<a id="L259"></a>}

<a id="L261"></a>func TestMatchStrings(t *testing.T) {
    <a id="L262"></a>for i := 0; i &lt; len(matches); i++ {
        <a id="L263"></a>test := &amp;matches[i];
        <a id="L264"></a>matchTest(t, test.re, test.text, test.match);
    <a id="L265"></a>}
<a id="L266"></a>}

<a id="L268"></a>func matchFunctionTest(t *testing.T, expr string, str string, match []int) {
    <a id="L269"></a>m, err := MatchString(expr, str);
    <a id="L270"></a>if err == nil {
        <a id="L271"></a>return
    <a id="L272"></a>}
    <a id="L273"></a>if m != (len(match) &gt; 0) {
        <a id="L274"></a>t.Error(&#34;function Match failure on `&#34;, expr, &#34;` matching `&#34;, str, &#34;`:&#34;, m, &#34;should be&#34;, len(match) &gt; 0)
    <a id="L275"></a>}
<a id="L276"></a>}

<a id="L278"></a>func TestMatchFunction(t *testing.T) {
    <a id="L279"></a>for i := 0; i &lt; len(matches); i++ {
        <a id="L280"></a>test := &amp;matches[i];
        <a id="L281"></a>matchFunctionTest(t, test.re, test.text, test.match);
    <a id="L282"></a>}
<a id="L283"></a>}

<a id="L285"></a>type ReplaceTest struct {
    <a id="L286"></a>pattern, replacement, input, output string;
<a id="L287"></a>}

<a id="L289"></a>var replaceTests = []ReplaceTest{
    <a id="L290"></a><span class="comment">// Test empty input and/or replacement, with pattern that matches the empty string.</span>
    <a id="L291"></a>ReplaceTest{&#34;&#34;, &#34;&#34;, &#34;&#34;, &#34;&#34;},
    <a id="L292"></a>ReplaceTest{&#34;&#34;, &#34;x&#34;, &#34;&#34;, &#34;x&#34;},
    <a id="L293"></a>ReplaceTest{&#34;&#34;, &#34;&#34;, &#34;abc&#34;, &#34;abc&#34;},
    <a id="L294"></a>ReplaceTest{&#34;&#34;, &#34;x&#34;, &#34;abc&#34;, &#34;xaxbxcx&#34;},

    <a id="L296"></a><span class="comment">// Test empty input and/or replacement, with pattern that does not match the empty string.</span>
    <a id="L297"></a>ReplaceTest{&#34;b&#34;, &#34;&#34;, &#34;&#34;, &#34;&#34;},
    <a id="L298"></a>ReplaceTest{&#34;b&#34;, &#34;x&#34;, &#34;&#34;, &#34;&#34;},
    <a id="L299"></a>ReplaceTest{&#34;b&#34;, &#34;&#34;, &#34;abc&#34;, &#34;ac&#34;},
    <a id="L300"></a>ReplaceTest{&#34;b&#34;, &#34;x&#34;, &#34;abc&#34;, &#34;axc&#34;},
    <a id="L301"></a>ReplaceTest{&#34;y&#34;, &#34;&#34;, &#34;&#34;, &#34;&#34;},
    <a id="L302"></a>ReplaceTest{&#34;y&#34;, &#34;x&#34;, &#34;&#34;, &#34;&#34;},
    <a id="L303"></a>ReplaceTest{&#34;y&#34;, &#34;&#34;, &#34;abc&#34;, &#34;abc&#34;},
    <a id="L304"></a>ReplaceTest{&#34;y&#34;, &#34;x&#34;, &#34;abc&#34;, &#34;abc&#34;},

    <a id="L306"></a><span class="comment">// Multibyte characters -- verify that we don&#39;t try to match in the middle</span>
    <a id="L307"></a><span class="comment">// of a character.</span>
    <a id="L308"></a>ReplaceTest{&#34;[a-c]*&#34;, &#34;x&#34;, &#34;\u65e5&#34;, &#34;x\u65e5x&#34;},
    <a id="L309"></a>ReplaceTest{&#34;[^\u65e5]&#34;, &#34;x&#34;, &#34;abc\u65e5def&#34;, &#34;xxx\u65e5xxx&#34;},

    <a id="L311"></a><span class="comment">// Start and end of a string.</span>
    <a id="L312"></a>ReplaceTest{&#34;^[a-c]*&#34;, &#34;x&#34;, &#34;abcdabc&#34;, &#34;xdabc&#34;},
    <a id="L313"></a>ReplaceTest{&#34;[a-c]*$&#34;, &#34;x&#34;, &#34;abcdabc&#34;, &#34;abcdx&#34;},
    <a id="L314"></a>ReplaceTest{&#34;^[a-c]*$&#34;, &#34;x&#34;, &#34;abcdabc&#34;, &#34;abcdabc&#34;},
    <a id="L315"></a>ReplaceTest{&#34;^[a-c]*&#34;, &#34;x&#34;, &#34;abc&#34;, &#34;x&#34;},
    <a id="L316"></a>ReplaceTest{&#34;[a-c]*$&#34;, &#34;x&#34;, &#34;abc&#34;, &#34;x&#34;},
    <a id="L317"></a>ReplaceTest{&#34;^[a-c]*$&#34;, &#34;x&#34;, &#34;abc&#34;, &#34;x&#34;},
    <a id="L318"></a>ReplaceTest{&#34;^[a-c]*&#34;, &#34;x&#34;, &#34;dabce&#34;, &#34;xdabce&#34;},
    <a id="L319"></a>ReplaceTest{&#34;[a-c]*$&#34;, &#34;x&#34;, &#34;dabce&#34;, &#34;dabcex&#34;},
    <a id="L320"></a>ReplaceTest{&#34;^[a-c]*$&#34;, &#34;x&#34;, &#34;dabce&#34;, &#34;dabce&#34;},
    <a id="L321"></a>ReplaceTest{&#34;^[a-c]*&#34;, &#34;x&#34;, &#34;&#34;, &#34;x&#34;},
    <a id="L322"></a>ReplaceTest{&#34;[a-c]*$&#34;, &#34;x&#34;, &#34;&#34;, &#34;x&#34;},
    <a id="L323"></a>ReplaceTest{&#34;^[a-c]*$&#34;, &#34;x&#34;, &#34;&#34;, &#34;x&#34;},

    <a id="L325"></a>ReplaceTest{&#34;^[a-c]+&#34;, &#34;x&#34;, &#34;abcdabc&#34;, &#34;xdabc&#34;},
    <a id="L326"></a>ReplaceTest{&#34;[a-c]+$&#34;, &#34;x&#34;, &#34;abcdabc&#34;, &#34;abcdx&#34;},
    <a id="L327"></a>ReplaceTest{&#34;^[a-c]+$&#34;, &#34;x&#34;, &#34;abcdabc&#34;, &#34;abcdabc&#34;},
    <a id="L328"></a>ReplaceTest{&#34;^[a-c]+&#34;, &#34;x&#34;, &#34;abc&#34;, &#34;x&#34;},
    <a id="L329"></a>ReplaceTest{&#34;[a-c]+$&#34;, &#34;x&#34;, &#34;abc&#34;, &#34;x&#34;},
    <a id="L330"></a>ReplaceTest{&#34;^[a-c]+$&#34;, &#34;x&#34;, &#34;abc&#34;, &#34;x&#34;},
    <a id="L331"></a>ReplaceTest{&#34;^[a-c]+&#34;, &#34;x&#34;, &#34;dabce&#34;, &#34;dabce&#34;},
    <a id="L332"></a>ReplaceTest{&#34;[a-c]+$&#34;, &#34;x&#34;, &#34;dabce&#34;, &#34;dabce&#34;},
    <a id="L333"></a>ReplaceTest{&#34;^[a-c]+$&#34;, &#34;x&#34;, &#34;dabce&#34;, &#34;dabce&#34;},
    <a id="L334"></a>ReplaceTest{&#34;^[a-c]+&#34;, &#34;x&#34;, &#34;&#34;, &#34;&#34;},
    <a id="L335"></a>ReplaceTest{&#34;[a-c]+$&#34;, &#34;x&#34;, &#34;&#34;, &#34;&#34;},
    <a id="L336"></a>ReplaceTest{&#34;^[a-c]+$&#34;, &#34;x&#34;, &#34;&#34;, &#34;&#34;},

    <a id="L338"></a><span class="comment">// Other cases.</span>
    <a id="L339"></a>ReplaceTest{&#34;abc&#34;, &#34;def&#34;, &#34;abcdefg&#34;, &#34;defdefg&#34;},
    <a id="L340"></a>ReplaceTest{&#34;bc&#34;, &#34;BC&#34;, &#34;abcbcdcdedef&#34;, &#34;aBCBCdcdedef&#34;},
    <a id="L341"></a>ReplaceTest{&#34;abc&#34;, &#34;&#34;, &#34;abcdabc&#34;, &#34;d&#34;},
    <a id="L342"></a>ReplaceTest{&#34;x&#34;, &#34;xXx&#34;, &#34;xxxXxxx&#34;, &#34;xXxxXxxXxXxXxxXxxXx&#34;},
    <a id="L343"></a>ReplaceTest{&#34;abc&#34;, &#34;d&#34;, &#34;&#34;, &#34;&#34;},
    <a id="L344"></a>ReplaceTest{&#34;abc&#34;, &#34;d&#34;, &#34;abc&#34;, &#34;d&#34;},
    <a id="L345"></a>ReplaceTest{&#34;.+&#34;, &#34;x&#34;, &#34;abc&#34;, &#34;x&#34;},
    <a id="L346"></a>ReplaceTest{&#34;[a-c]*&#34;, &#34;x&#34;, &#34;def&#34;, &#34;xdxexfx&#34;},
    <a id="L347"></a>ReplaceTest{&#34;[a-c]+&#34;, &#34;x&#34;, &#34;abcbcdcdedef&#34;, &#34;xdxdedef&#34;},
    <a id="L348"></a>ReplaceTest{&#34;[a-c]*&#34;, &#34;x&#34;, &#34;abcbcdcdedef&#34;, &#34;xdxdxexdxexfx&#34;},
<a id="L349"></a>}

<a id="L351"></a>func TestReplaceAll(t *testing.T) {
    <a id="L352"></a>for _, tc := range replaceTests {
        <a id="L353"></a>re, err := Compile(tc.pattern);
        <a id="L354"></a>if err != nil {
            <a id="L355"></a>t.Errorf(&#34;Unexpected error compiling %q: %v&#34;, tc.pattern, err);
            <a id="L356"></a>continue;
        <a id="L357"></a>}
        <a id="L358"></a>actual := re.ReplaceAllString(tc.input, tc.replacement);
        <a id="L359"></a>if actual != tc.output {
            <a id="L360"></a>t.Errorf(&#34;%q.Replace(%q,%q) = %q; want %q&#34;,
                <a id="L361"></a>tc.pattern, tc.input, tc.replacement, actual, tc.output)
        <a id="L362"></a>}
        <a id="L363"></a><span class="comment">// now try bytes</span>
        <a id="L364"></a>actual = string(re.ReplaceAll(strings.Bytes(tc.input), strings.Bytes(tc.replacement)));
        <a id="L365"></a>if actual != tc.output {
            <a id="L366"></a>t.Errorf(&#34;%q.Replace(%q,%q) = %q; want %q&#34;,
                <a id="L367"></a>tc.pattern, tc.input, tc.replacement, actual, tc.output)
        <a id="L368"></a>}
    <a id="L369"></a>}
<a id="L370"></a>}

<a id="L372"></a>type QuoteMetaTest struct {
    <a id="L373"></a>pattern, output string;
<a id="L374"></a>}

<a id="L376"></a>var quoteMetaTests = []QuoteMetaTest{
    <a id="L377"></a>QuoteMetaTest{``, ``},
    <a id="L378"></a>QuoteMetaTest{`foo`, `foo`},
    <a id="L379"></a>QuoteMetaTest{`!@#$%^&amp;*()_+-=[{]}\|,&lt;.&gt;/?~`, `!@#\$%\^&amp;\*\(\)_\+-=\[{\]}\\\|,&lt;\.&gt;/\?~`},
<a id="L380"></a>}

<a id="L382"></a>func TestQuoteMeta(t *testing.T) {
    <a id="L383"></a>for _, tc := range quoteMetaTests {
        <a id="L384"></a><span class="comment">// Verify that QuoteMeta returns the expected string.</span>
        <a id="L385"></a>quoted := QuoteMeta(tc.pattern);
        <a id="L386"></a>if quoted != tc.output {
            <a id="L387"></a>t.Errorf(&#34;QuoteMeta(`%s`) = `%s`; want `%s`&#34;,
                <a id="L388"></a>tc.pattern, quoted, tc.output);
            <a id="L389"></a>continue;
        <a id="L390"></a>}

        <a id="L392"></a><span class="comment">// Verify that the quoted string is in fact treated as expected</span>
        <a id="L393"></a><span class="comment">// by Compile -- i.e. that it matches the original, unquoted string.</span>
        <a id="L394"></a>if tc.pattern != &#34;&#34; {
            <a id="L395"></a>re, err := Compile(quoted);
            <a id="L396"></a>if err != nil {
                <a id="L397"></a>t.Errorf(&#34;Unexpected error compiling QuoteMeta(`%s`): %v&#34;, tc.pattern, err);
                <a id="L398"></a>continue;
            <a id="L399"></a>}
            <a id="L400"></a>src := &#34;abc&#34; + tc.pattern + &#34;def&#34;;
            <a id="L401"></a>repl := &#34;xyz&#34;;
            <a id="L402"></a>replaced := re.ReplaceAllString(src, repl);
            <a id="L403"></a>expected := &#34;abcxyzdef&#34;;
            <a id="L404"></a>if replaced != expected {
                <a id="L405"></a>t.Errorf(&#34;QuoteMeta(`%s`).Replace(`%s`,`%s`) = `%s`; want `%s`&#34;,
                    <a id="L406"></a>tc.pattern, src, repl, replaced, expected)
            <a id="L407"></a>}
        <a id="L408"></a>}
    <a id="L409"></a>}
<a id="L410"></a>}

<a id="L412"></a>type matchCase struct {
    <a id="L413"></a>matchfunc string;
    <a id="L414"></a>input     string;
    <a id="L415"></a>n         int;
    <a id="L416"></a>regexp    string;
    <a id="L417"></a>expected  []string;
<a id="L418"></a>}

<a id="L420"></a>var matchCases = []matchCase{
    <a id="L421"></a>matchCase{&#34;match&#34;, &#34; aa b&#34;, 0, &#34;[^ ]+&#34;, []string{&#34;aa&#34;, &#34;b&#34;}},
    <a id="L422"></a>matchCase{&#34;match&#34;, &#34; aa b&#34;, 0, &#34;[^ ]*&#34;, []string{&#34;&#34;, &#34;aa&#34;, &#34;b&#34;}},
    <a id="L423"></a>matchCase{&#34;match&#34;, &#34;a b c&#34;, 0, &#34;[^ ]*&#34;, []string{&#34;a&#34;, &#34;b&#34;, &#34;c&#34;}},
    <a id="L424"></a>matchCase{&#34;match&#34;, &#34;a:a: a:&#34;, 0, &#34;^.:&#34;, []string{&#34;a:&#34;}},
    <a id="L425"></a>matchCase{&#34;match&#34;, &#34;&#34;, 0, &#34;[^ ]*&#34;, []string{&#34;&#34;}},
    <a id="L426"></a>matchCase{&#34;match&#34;, &#34;&#34;, 0, &#34;&#34;, []string{&#34;&#34;}},
    <a id="L427"></a>matchCase{&#34;match&#34;, &#34;a&#34;, 0, &#34;&#34;, []string{&#34;&#34;, &#34;&#34;}},
    <a id="L428"></a>matchCase{&#34;match&#34;, &#34;ab&#34;, 0, &#34;^&#34;, []string{&#34;&#34;}},
    <a id="L429"></a>matchCase{&#34;match&#34;, &#34;ab&#34;, 0, &#34;$&#34;, []string{&#34;&#34;}},
    <a id="L430"></a>matchCase{&#34;match&#34;, &#34;ab&#34;, 0, &#34;X*&#34;, []string{&#34;&#34;, &#34;&#34;, &#34;&#34;}},
    <a id="L431"></a>matchCase{&#34;match&#34;, &#34;aX&#34;, 0, &#34;X*&#34;, []string{&#34;&#34;, &#34;X&#34;}},
    <a id="L432"></a>matchCase{&#34;match&#34;, &#34;XabX&#34;, 0, &#34;X*&#34;, []string{&#34;X&#34;, &#34;&#34;, &#34;X&#34;}},

    <a id="L434"></a>matchCase{&#34;matchit&#34;, &#34;&#34;, 0, &#34;.&#34;, []string{}},
    <a id="L435"></a>matchCase{&#34;matchit&#34;, &#34;abc&#34;, 2, &#34;.&#34;, []string{&#34;a&#34;, &#34;b&#34;}},
    <a id="L436"></a>matchCase{&#34;matchit&#34;, &#34;abc&#34;, 0, &#34;.&#34;, []string{&#34;a&#34;, &#34;b&#34;, &#34;c&#34;}},
<a id="L437"></a>}

<a id="L439"></a>func printStringSlice(t *testing.T, s []string) {
    <a id="L440"></a>l := len(s);
    <a id="L441"></a>if l == 0 {
        <a id="L442"></a>t.Log(&#34;\t&lt;empty&gt;&#34;)
    <a id="L443"></a>} else {
        <a id="L444"></a>for i := 0; i &lt; l; i++ {
            <a id="L445"></a>t.Logf(&#34;\t%q&#34;, s[i])
        <a id="L446"></a>}
    <a id="L447"></a>}
<a id="L448"></a>}

<a id="L450"></a>func TestAllMatches(t *testing.T) {
    <a id="L451"></a>ch := make(chan matchCase);
    <a id="L452"></a>go func() {
        <a id="L453"></a>for _, c := range matchCases {
            <a id="L454"></a>ch &lt;- c;
            <a id="L455"></a>stringCase := matchCase{
                <a id="L456"></a>&#34;string&#34; + c.matchfunc,
                <a id="L457"></a>c.input,
                <a id="L458"></a>c.n,
                <a id="L459"></a>c.regexp,
                <a id="L460"></a>c.expected,
            <a id="L461"></a>};
            <a id="L462"></a>ch &lt;- stringCase;
        <a id="L463"></a>}
        <a id="L464"></a>close(ch);
    <a id="L465"></a>}();

    <a id="L467"></a>for c := range ch {
        <a id="L468"></a>var result []string;
        <a id="L469"></a>re, _ := Compile(c.regexp);

        <a id="L471"></a>switch c.matchfunc {
        <a id="L472"></a>case &#34;matchit&#34;:
            <a id="L473"></a>result = make([]string, len(c.input)+1);
            <a id="L474"></a>i := 0;
            <a id="L475"></a>b := strings.Bytes(c.input);
            <a id="L476"></a>for match := range re.AllMatchesIter(b, c.n) {
                <a id="L477"></a>result[i] = string(match);
                <a id="L478"></a>i++;
            <a id="L479"></a>}
            <a id="L480"></a>result = result[0:i];
        <a id="L481"></a>case &#34;stringmatchit&#34;:
            <a id="L482"></a>result = make([]string, len(c.input)+1);
            <a id="L483"></a>i := 0;
            <a id="L484"></a>for match := range re.AllMatchesStringIter(c.input, c.n) {
                <a id="L485"></a>result[i] = match;
                <a id="L486"></a>i++;
            <a id="L487"></a>}
            <a id="L488"></a>result = result[0:i];
        <a id="L489"></a>case &#34;match&#34;:
            <a id="L490"></a>result = make([]string, len(c.input)+1);
            <a id="L491"></a>b := strings.Bytes(c.input);
            <a id="L492"></a>i := 0;
            <a id="L493"></a>for _, match := range re.AllMatches(b, c.n) {
                <a id="L494"></a>result[i] = string(match);
                <a id="L495"></a>i++;
            <a id="L496"></a>}
            <a id="L497"></a>result = result[0:i];
        <a id="L498"></a>case &#34;stringmatch&#34;:
            <a id="L499"></a>result = re.AllMatchesString(c.input, c.n)
        <a id="L500"></a>}

        <a id="L502"></a>if !equalStrings(result, c.expected) {
            <a id="L503"></a>t.Errorf(&#34;testing &#39;%s&#39;.%s(&#39;%s&#39;, %d), expected: &#34;,
                <a id="L504"></a>c.regexp, c.matchfunc, c.input, c.n);
            <a id="L505"></a>printStringSlice(t, c.expected);
            <a id="L506"></a>t.Log(&#34;got: &#34;);
            <a id="L507"></a>printStringSlice(t, result);
            <a id="L508"></a>t.Log(&#34;\n&#34;);
        <a id="L509"></a>}
    <a id="L510"></a>}
<a id="L511"></a>}
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
