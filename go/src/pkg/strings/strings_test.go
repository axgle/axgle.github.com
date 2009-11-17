<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strings/strings_test.go</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/strings/strings_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package strings_test

<a id="L7"></a>import (
    <a id="L8"></a>. &#34;strings&#34;;
    <a id="L9"></a>&#34;testing&#34;;
    <a id="L10"></a>&#34;unicode&#34;;
    <a id="L11"></a>&#34;utf8&#34;;
<a id="L12"></a>)

<a id="L14"></a>func eq(a, b []string) bool {
    <a id="L15"></a>if len(a) != len(b) {
        <a id="L16"></a>return false
    <a id="L17"></a>}
    <a id="L18"></a>for i := 0; i &lt; len(a); i++ {
        <a id="L19"></a>if a[i] != b[i] {
            <a id="L20"></a>return false
        <a id="L21"></a>}
    <a id="L22"></a>}
    <a id="L23"></a>return true;
<a id="L24"></a>}

<a id="L26"></a>var abcd = &#34;abcd&#34;
<a id="L27"></a>var faces = &#34;☺☻☹&#34;
<a id="L28"></a>var commas = &#34;1,2,3,4&#34;
<a id="L29"></a>var dots = &#34;1....2....3....4&#34;

<a id="L31"></a>type IndexTest struct {
    <a id="L32"></a>s   string;
    <a id="L33"></a>sep string;
    <a id="L34"></a>out int;
<a id="L35"></a>}

<a id="L37"></a>var indexTests = []IndexTest{
    <a id="L38"></a>IndexTest{&#34;&#34;, &#34;&#34;, 0},
    <a id="L39"></a>IndexTest{&#34;&#34;, &#34;a&#34;, -1},
    <a id="L40"></a>IndexTest{&#34;&#34;, &#34;foo&#34;, -1},
    <a id="L41"></a>IndexTest{&#34;fo&#34;, &#34;foo&#34;, -1},
    <a id="L42"></a>IndexTest{&#34;foo&#34;, &#34;foo&#34;, 0},
    <a id="L43"></a>IndexTest{&#34;oofofoofooo&#34;, &#34;f&#34;, 2},
    <a id="L44"></a>IndexTest{&#34;oofofoofooo&#34;, &#34;foo&#34;, 4},
    <a id="L45"></a>IndexTest{&#34;barfoobarfoo&#34;, &#34;foo&#34;, 3},
    <a id="L46"></a>IndexTest{&#34;foo&#34;, &#34;&#34;, 0},
    <a id="L47"></a>IndexTest{&#34;foo&#34;, &#34;o&#34;, 1},
    <a id="L48"></a>IndexTest{&#34;abcABCabc&#34;, &#34;A&#34;, 3},
<a id="L49"></a>}

<a id="L51"></a>var lastIndexTests = []IndexTest{
    <a id="L52"></a>IndexTest{&#34;&#34;, &#34;&#34;, 0},
    <a id="L53"></a>IndexTest{&#34;&#34;, &#34;a&#34;, -1},
    <a id="L54"></a>IndexTest{&#34;&#34;, &#34;foo&#34;, -1},
    <a id="L55"></a>IndexTest{&#34;fo&#34;, &#34;foo&#34;, -1},
    <a id="L56"></a>IndexTest{&#34;foo&#34;, &#34;foo&#34;, 0},
    <a id="L57"></a>IndexTest{&#34;oofofoofooo&#34;, &#34;f&#34;, 7},
    <a id="L58"></a>IndexTest{&#34;oofofoofooo&#34;, &#34;foo&#34;, 7},
    <a id="L59"></a>IndexTest{&#34;barfoobarfoo&#34;, &#34;foo&#34;, 9},
    <a id="L60"></a>IndexTest{&#34;foo&#34;, &#34;&#34;, 3},
    <a id="L61"></a>IndexTest{&#34;foo&#34;, &#34;o&#34;, 2},
    <a id="L62"></a>IndexTest{&#34;abcABCabc&#34;, &#34;A&#34;, 3},
    <a id="L63"></a>IndexTest{&#34;abcABCabc&#34;, &#34;a&#34;, 6},
<a id="L64"></a>}

<a id="L66"></a><span class="comment">// Execute f on each test case.  funcName should be the name of f; it&#39;s used</span>
<a id="L67"></a><span class="comment">// in failure reports.</span>
<a id="L68"></a>func runIndexTests(t *testing.T, f func(s, sep string) int, funcName string, testCases []IndexTest) {
    <a id="L69"></a>for _, test := range testCases {
        <a id="L70"></a>actual := f(test.s, test.sep);
        <a id="L71"></a>if actual != test.out {
            <a id="L72"></a>t.Errorf(&#34;%s(%q,%q) = %v; want %v&#34;, funcName, test.s, test.sep, actual, test.out)
        <a id="L73"></a>}
    <a id="L74"></a>}
<a id="L75"></a>}

<a id="L77"></a>func TestIndex(t *testing.T) { runIndexTests(t, Index, &#34;Index&#34;, indexTests) }

<a id="L79"></a>func TestLastIndex(t *testing.T) { runIndexTests(t, LastIndex, &#34;LastIndex&#34;, lastIndexTests) }


<a id="L82"></a>type ExplodeTest struct {
    <a id="L83"></a>s   string;
    <a id="L84"></a>n   int;
    <a id="L85"></a>a   []string;
<a id="L86"></a>}

<a id="L88"></a>var explodetests = []ExplodeTest{
    <a id="L89"></a>ExplodeTest{abcd, 4, []string{&#34;a&#34;, &#34;b&#34;, &#34;c&#34;, &#34;d&#34;}},
    <a id="L90"></a>ExplodeTest{faces, 3, []string{&#34;☺&#34;, &#34;☻&#34;, &#34;☹&#34;}},
    <a id="L91"></a>ExplodeTest{abcd, 2, []string{&#34;a&#34;, &#34;bcd&#34;}},
<a id="L92"></a>}

<a id="L94"></a>func TestExplode(t *testing.T) {
    <a id="L95"></a>for _, tt := range explodetests {
        <a id="L96"></a>a := Split(tt.s, &#34;&#34;, tt.n);
        <a id="L97"></a>if !eq(a, tt.a) {
            <a id="L98"></a>t.Errorf(&#34;explode(%q, %d) = %v; want %v&#34;, tt.s, tt.n, a, tt.a);
            <a id="L99"></a>continue;
        <a id="L100"></a>}
        <a id="L101"></a>s := Join(a, &#34;&#34;);
        <a id="L102"></a>if s != tt.s {
            <a id="L103"></a>t.Errorf(`Join(explode(%q, %d), &#34;&#34;) = %q`, tt.s, tt.n, s)
        <a id="L104"></a>}
    <a id="L105"></a>}
<a id="L106"></a>}

<a id="L108"></a>type SplitTest struct {
    <a id="L109"></a>s   string;
    <a id="L110"></a>sep string;
    <a id="L111"></a>n   int;
    <a id="L112"></a>a   []string;
<a id="L113"></a>}

<a id="L115"></a>var splittests = []SplitTest{
    <a id="L116"></a>SplitTest{abcd, &#34;a&#34;, 0, []string{&#34;&#34;, &#34;bcd&#34;}},
    <a id="L117"></a>SplitTest{abcd, &#34;z&#34;, 0, []string{&#34;abcd&#34;}},
    <a id="L118"></a>SplitTest{abcd, &#34;&#34;, 0, []string{&#34;a&#34;, &#34;b&#34;, &#34;c&#34;, &#34;d&#34;}},
    <a id="L119"></a>SplitTest{commas, &#34;,&#34;, 0, []string{&#34;1&#34;, &#34;2&#34;, &#34;3&#34;, &#34;4&#34;}},
    <a id="L120"></a>SplitTest{dots, &#34;...&#34;, 0, []string{&#34;1&#34;, &#34;.2&#34;, &#34;.3&#34;, &#34;.4&#34;}},
    <a id="L121"></a>SplitTest{faces, &#34;☹&#34;, 0, []string{&#34;☺☻&#34;, &#34;&#34;}},
    <a id="L122"></a>SplitTest{faces, &#34;~&#34;, 0, []string{faces}},
    <a id="L123"></a>SplitTest{faces, &#34;&#34;, 0, []string{&#34;☺&#34;, &#34;☻&#34;, &#34;☹&#34;}},
    <a id="L124"></a>SplitTest{&#34;1 2 3 4&#34;, &#34; &#34;, 3, []string{&#34;1&#34;, &#34;2&#34;, &#34;3 4&#34;}},
    <a id="L125"></a>SplitTest{&#34;1 2&#34;, &#34; &#34;, 3, []string{&#34;1&#34;, &#34;2&#34;}},
    <a id="L126"></a>SplitTest{&#34;123&#34;, &#34;&#34;, 2, []string{&#34;1&#34;, &#34;23&#34;}},
    <a id="L127"></a>SplitTest{&#34;123&#34;, &#34;&#34;, 17, []string{&#34;1&#34;, &#34;2&#34;, &#34;3&#34;}},
<a id="L128"></a>}

<a id="L130"></a>func TestSplit(t *testing.T) {
    <a id="L131"></a>for _, tt := range splittests {
        <a id="L132"></a>a := Split(tt.s, tt.sep, tt.n);
        <a id="L133"></a>if !eq(a, tt.a) {
            <a id="L134"></a>t.Errorf(&#34;Split(%q, %q, %d) = %v; want %v&#34;, tt.s, tt.sep, tt.n, a, tt.a);
            <a id="L135"></a>continue;
        <a id="L136"></a>}
        <a id="L137"></a>s := Join(a, tt.sep);
        <a id="L138"></a>if s != tt.s {
            <a id="L139"></a>t.Errorf(&#34;Join(Split(%q, %q, %d), %q) = %q&#34;, tt.s, tt.sep, tt.n, tt.sep, s)
        <a id="L140"></a>}
    <a id="L141"></a>}
<a id="L142"></a>}

<a id="L144"></a>var splitaftertests = []SplitTest{
    <a id="L145"></a>SplitTest{abcd, &#34;a&#34;, 0, []string{&#34;a&#34;, &#34;bcd&#34;}},
    <a id="L146"></a>SplitTest{abcd, &#34;z&#34;, 0, []string{&#34;abcd&#34;}},
    <a id="L147"></a>SplitTest{abcd, &#34;&#34;, 0, []string{&#34;a&#34;, &#34;b&#34;, &#34;c&#34;, &#34;d&#34;}},
    <a id="L148"></a>SplitTest{commas, &#34;,&#34;, 0, []string{&#34;1,&#34;, &#34;2,&#34;, &#34;3,&#34;, &#34;4&#34;}},
    <a id="L149"></a>SplitTest{dots, &#34;...&#34;, 0, []string{&#34;1...&#34;, &#34;.2...&#34;, &#34;.3...&#34;, &#34;.4&#34;}},
    <a id="L150"></a>SplitTest{faces, &#34;☹&#34;, 0, []string{&#34;☺☻☹&#34;, &#34;&#34;}},
    <a id="L151"></a>SplitTest{faces, &#34;~&#34;, 0, []string{faces}},
    <a id="L152"></a>SplitTest{faces, &#34;&#34;, 0, []string{&#34;☺&#34;, &#34;☻&#34;, &#34;☹&#34;}},
    <a id="L153"></a>SplitTest{&#34;1 2 3 4&#34;, &#34; &#34;, 3, []string{&#34;1 &#34;, &#34;2 &#34;, &#34;3 4&#34;}},
    <a id="L154"></a>SplitTest{&#34;1 2 3&#34;, &#34; &#34;, 3, []string{&#34;1 &#34;, &#34;2 &#34;, &#34;3&#34;}},
    <a id="L155"></a>SplitTest{&#34;1 2&#34;, &#34; &#34;, 3, []string{&#34;1 &#34;, &#34;2&#34;}},
    <a id="L156"></a>SplitTest{&#34;123&#34;, &#34;&#34;, 2, []string{&#34;1&#34;, &#34;23&#34;}},
    <a id="L157"></a>SplitTest{&#34;123&#34;, &#34;&#34;, 17, []string{&#34;1&#34;, &#34;2&#34;, &#34;3&#34;}},
<a id="L158"></a>}

<a id="L160"></a>func TestSplitAfter(t *testing.T) {
    <a id="L161"></a>for _, tt := range splitaftertests {
        <a id="L162"></a>a := SplitAfter(tt.s, tt.sep, tt.n);
        <a id="L163"></a>if !eq(a, tt.a) {
            <a id="L164"></a>t.Errorf(`Split(%q, %q, %d) = %v; want %v`, tt.s, tt.sep, tt.n, a, tt.a);
            <a id="L165"></a>continue;
        <a id="L166"></a>}
        <a id="L167"></a>s := Join(a, &#34;&#34;);
        <a id="L168"></a>if s != tt.s {
            <a id="L169"></a>t.Errorf(`Join(Split(%q, %q, %d), %q) = %q`, tt.s, tt.sep, tt.n, tt.sep, s)
        <a id="L170"></a>}
    <a id="L171"></a>}
<a id="L172"></a>}

<a id="L174"></a><span class="comment">// Test case for any function which accepts and returns a single string.</span>
<a id="L175"></a>type StringTest struct {
    <a id="L176"></a>in, out string;
<a id="L177"></a>}

<a id="L179"></a><span class="comment">// Execute f on each test case.  funcName should be the name of f; it&#39;s used</span>
<a id="L180"></a><span class="comment">// in failure reports.</span>
<a id="L181"></a>func runStringTests(t *testing.T, f func(string) string, funcName string, testCases []StringTest) {
    <a id="L182"></a>for _, tc := range testCases {
        <a id="L183"></a>actual := f(tc.in);
        <a id="L184"></a>if actual != tc.out {
            <a id="L185"></a>t.Errorf(&#34;%s(%q) = %q; want %q&#34;, funcName, tc.in, actual, tc.out)
        <a id="L186"></a>}
    <a id="L187"></a>}
<a id="L188"></a>}

<a id="L190"></a>var upperTests = []StringTest{
    <a id="L191"></a>StringTest{&#34;&#34;, &#34;&#34;},
    <a id="L192"></a>StringTest{&#34;abc&#34;, &#34;ABC&#34;},
    <a id="L193"></a>StringTest{&#34;AbC123&#34;, &#34;ABC123&#34;},
    <a id="L194"></a>StringTest{&#34;azAZ09_&#34;, &#34;AZAZ09_&#34;},
    <a id="L195"></a>StringTest{&#34;\u0250\u0250\u0250\u0250\u0250&#34;, &#34;\u2C6F\u2C6F\u2C6F\u2C6F\u2C6F&#34;}, <span class="comment">// grows one byte per char</span>
<a id="L196"></a>}

<a id="L198"></a>var lowerTests = []StringTest{
    <a id="L199"></a>StringTest{&#34;&#34;, &#34;&#34;},
    <a id="L200"></a>StringTest{&#34;abc&#34;, &#34;abc&#34;},
    <a id="L201"></a>StringTest{&#34;AbC123&#34;, &#34;abc123&#34;},
    <a id="L202"></a>StringTest{&#34;azAZ09_&#34;, &#34;azaz09_&#34;},
    <a id="L203"></a>StringTest{&#34;\u2C6D\u2C6D\u2C6D\u2C6D\u2C6D&#34;, &#34;\u0251\u0251\u0251\u0251\u0251&#34;}, <span class="comment">// shrinks one byte per char</span>
<a id="L204"></a>}

<a id="L206"></a>const space = &#34;\t\v\r\f\n\u0085\u00a0\u2000\u3000&#34;

<a id="L208"></a>var trimSpaceTests = []StringTest{
    <a id="L209"></a>StringTest{&#34;&#34;, &#34;&#34;},
    <a id="L210"></a>StringTest{&#34;abc&#34;, &#34;abc&#34;},
    <a id="L211"></a>StringTest{space + &#34;abc&#34; + space, &#34;abc&#34;},
    <a id="L212"></a>StringTest{&#34; &#34;, &#34;&#34;},
    <a id="L213"></a>StringTest{&#34; \t\r\n \t\t\r\r\n\n &#34;, &#34;&#34;},
    <a id="L214"></a>StringTest{&#34; \t\r\n x\t\t\r\r\n\n &#34;, &#34;x&#34;},
    <a id="L215"></a>StringTest{&#34; \u2000\t\r\n x\t\t\r\r\ny\n \u3000&#34;, &#34;x\t\t\r\r\ny&#34;},
    <a id="L216"></a>StringTest{&#34;1 \t\r\n2&#34;, &#34;1 \t\r\n2&#34;},
    <a id="L217"></a>StringTest{&#34; x\x80&#34;, &#34;x\x80&#34;}, <span class="comment">// invalid UTF-8 on end</span>
    <a id="L218"></a>StringTest{&#34; x\xc0&#34;, &#34;x\xc0&#34;}, <span class="comment">// invalid UTF-8 on end</span>
<a id="L219"></a>}

<a id="L221"></a>func tenRunes(rune int) string {
    <a id="L222"></a>r := make([]int, 10);
    <a id="L223"></a>for i := range r {
        <a id="L224"></a>r[i] = rune
    <a id="L225"></a>}
    <a id="L226"></a>return string(r);
<a id="L227"></a>}

<a id="L229"></a>func TestMap(t *testing.T) {
    <a id="L230"></a><span class="comment">// Run a couple of awful growth/shrinkage tests</span>
    <a id="L231"></a>a := tenRunes(&#39;a&#39;);
    <a id="L232"></a><span class="comment">// 1.  Grow.  This triggers two reallocations in Map.</span>
    <a id="L233"></a>maxRune := func(rune int) int { return unicode.MaxRune };
    <a id="L234"></a>m := Map(maxRune, a);
    <a id="L235"></a>expect := tenRunes(unicode.MaxRune);
    <a id="L236"></a>if m != expect {
        <a id="L237"></a>t.Errorf(&#34;growing: expected %q got %q&#34;, expect, m)
    <a id="L238"></a>}
    <a id="L239"></a><span class="comment">// 2. Shrink</span>
    <a id="L240"></a>minRune := func(rune int) int { return &#39;a&#39; };
    <a id="L241"></a>m = Map(minRune, tenRunes(unicode.MaxRune));
    <a id="L242"></a>expect = a;
    <a id="L243"></a>if m != expect {
        <a id="L244"></a>t.Errorf(&#34;shrinking: expected %q got %q&#34;, expect, m)
    <a id="L245"></a>}
<a id="L246"></a>}

<a id="L248"></a>func TestToUpper(t *testing.T) { runStringTests(t, ToUpper, &#34;ToUpper&#34;, upperTests) }

<a id="L250"></a>func TestToLower(t *testing.T) { runStringTests(t, ToLower, &#34;ToLower&#34;, lowerTests) }

<a id="L252"></a>func TestTrimSpace(t *testing.T) { runStringTests(t, TrimSpace, &#34;TrimSpace&#34;, trimSpaceTests) }

<a id="L254"></a>func equal(m string, s1, s2 string, t *testing.T) bool {
    <a id="L255"></a>if s1 == s2 {
        <a id="L256"></a>return true
    <a id="L257"></a>}
    <a id="L258"></a>e1 := Split(s1, &#34;&#34;, 0);
    <a id="L259"></a>e2 := Split(s2, &#34;&#34;, 0);
    <a id="L260"></a>for i, c1 := range e1 {
        <a id="L261"></a>if i &gt; len(e2) {
            <a id="L262"></a>break
        <a id="L263"></a>}
        <a id="L264"></a>r1, _ := utf8.DecodeRuneInString(c1);
        <a id="L265"></a>r2, _ := utf8.DecodeRuneInString(e2[i]);
        <a id="L266"></a>if r1 != r2 {
            <a id="L267"></a>t.Errorf(&#34;%s diff at %d: U+%04X U+%04X&#34;, m, i, r1, r2)
        <a id="L268"></a>}
    <a id="L269"></a>}
    <a id="L270"></a>return false;
<a id="L271"></a>}

<a id="L273"></a>func TestCaseConsistency(t *testing.T) {
    <a id="L274"></a><span class="comment">// Make a string of all the runes.</span>
    <a id="L275"></a>a := make([]int, unicode.MaxRune+1);
    <a id="L276"></a>for i := range a {
        <a id="L277"></a>a[i] = i
    <a id="L278"></a>}
    <a id="L279"></a>s := string(a);
    <a id="L280"></a><span class="comment">// convert the cases.</span>
    <a id="L281"></a>upper := ToUpper(s);
    <a id="L282"></a>lower := ToLower(s);

    <a id="L284"></a><span class="comment">// Consistency checks</span>
    <a id="L285"></a>if n := utf8.RuneCountInString(upper); n != unicode.MaxRune+1 {
        <a id="L286"></a>t.Error(&#34;rune count wrong in upper:&#34;, n)
    <a id="L287"></a>}
    <a id="L288"></a>if n := utf8.RuneCountInString(lower); n != unicode.MaxRune+1 {
        <a id="L289"></a>t.Error(&#34;rune count wrong in lower:&#34;, n)
    <a id="L290"></a>}
    <a id="L291"></a>if !equal(&#34;ToUpper(upper)&#34;, ToUpper(upper), upper, t) {
        <a id="L292"></a>t.Error(&#34;ToUpper(upper) consistency fail&#34;)
    <a id="L293"></a>}
    <a id="L294"></a>if !equal(&#34;ToLower(lower)&#34;, ToLower(lower), lower, t) {
        <a id="L295"></a>t.Error(&#34;ToLower(lower) consistency fail&#34;)
    <a id="L296"></a>}
    <a id="L297"></a><span class="comment">/*</span>
    <a id="L298"></a><span class="comment">	  These fail because of non-one-to-oneness of the data, such as multiple</span>
    <a id="L299"></a><span class="comment">	  upper case &#39;I&#39; mapping to &#39;i&#39;.  We comment them out but keep them for</span>
    <a id="L300"></a><span class="comment">	  interest.</span>
    <a id="L301"></a><span class="comment">	  For instance: CAPITAL LETTER I WITH DOT ABOVE:</span>
    <a id="L302"></a><span class="comment">		unicode.ToUpper(unicode.ToLower(&#39;\u0130&#39;)) != &#39;\u0130&#39;</span>

    <a id="L304"></a><span class="comment">	if !equal(&#34;ToUpper(lower)&#34;, ToUpper(lower), upper, t) {</span>
    <a id="L305"></a><span class="comment">		t.Error(&#34;ToUpper(lower) consistency fail&#34;);</span>
    <a id="L306"></a><span class="comment">	}</span>
    <a id="L307"></a><span class="comment">	if !equal(&#34;ToLower(upper)&#34;, ToLower(upper), lower, t) {</span>
    <a id="L308"></a><span class="comment">		t.Error(&#34;ToLower(upper) consistency fail&#34;);</span>
    <a id="L309"></a><span class="comment">	}</span>
    <a id="L310"></a><span class="comment">*/</span>
<a id="L311"></a>}
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
