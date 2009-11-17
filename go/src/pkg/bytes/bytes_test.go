<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/bytes/bytes_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/bytes/bytes_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package bytes_test

<a id="L7"></a>import (
    <a id="L8"></a>. &#34;bytes&#34;;
    <a id="L9"></a>&#34;strings&#34;;
    <a id="L10"></a>&#34;testing&#34;;
    <a id="L11"></a>&#34;unicode&#34;;
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

<a id="L26"></a>func arrayOfString(a [][]byte) []string {
    <a id="L27"></a>result := make([]string, len(a));
    <a id="L28"></a>for j := 0; j &lt; len(a); j++ {
        <a id="L29"></a>result[j] = string(a[j])
    <a id="L30"></a>}
    <a id="L31"></a>return result;
<a id="L32"></a>}

<a id="L34"></a><span class="comment">// For ease of reading, the test cases use strings that are converted to byte</span>
<a id="L35"></a><span class="comment">// arrays before invoking the functions.</span>

<a id="L37"></a>var abcd = &#34;abcd&#34;
<a id="L38"></a>var faces = &#34;☺☻☹&#34;
<a id="L39"></a>var commas = &#34;1,2,3,4&#34;
<a id="L40"></a>var dots = &#34;1....2....3....4&#34;

<a id="L42"></a>type CompareTest struct {
    <a id="L43"></a>a   string;
    <a id="L44"></a>b   string;
    <a id="L45"></a>cmp int;
<a id="L46"></a>}

<a id="L48"></a>var comparetests = []CompareTest{
    <a id="L49"></a>CompareTest{&#34;&#34;, &#34;&#34;, 0},
    <a id="L50"></a>CompareTest{&#34;a&#34;, &#34;&#34;, 1},
    <a id="L51"></a>CompareTest{&#34;&#34;, &#34;a&#34;, -1},
    <a id="L52"></a>CompareTest{&#34;abc&#34;, &#34;abc&#34;, 0},
    <a id="L53"></a>CompareTest{&#34;ab&#34;, &#34;abc&#34;, -1},
    <a id="L54"></a>CompareTest{&#34;abc&#34;, &#34;ab&#34;, 1},
    <a id="L55"></a>CompareTest{&#34;x&#34;, &#34;ab&#34;, 1},
    <a id="L56"></a>CompareTest{&#34;ab&#34;, &#34;x&#34;, -1},
    <a id="L57"></a>CompareTest{&#34;x&#34;, &#34;a&#34;, 1},
    <a id="L58"></a>CompareTest{&#34;b&#34;, &#34;x&#34;, -1},
<a id="L59"></a>}

<a id="L61"></a>func TestCompare(t *testing.T) {
    <a id="L62"></a>for i := 0; i &lt; len(comparetests); i++ {
        <a id="L63"></a>tt := comparetests[i];
        <a id="L64"></a>a := strings.Bytes(tt.a);
        <a id="L65"></a>b := strings.Bytes(tt.b);
        <a id="L66"></a>cmp := Compare(a, b);
        <a id="L67"></a>eql := Equal(a, b);
        <a id="L68"></a>if cmp != tt.cmp {
            <a id="L69"></a>t.Errorf(`Compare(%q, %q) = %v`, tt.a, tt.b, cmp)
        <a id="L70"></a>}
        <a id="L71"></a>if eql != (tt.cmp == 0) {
            <a id="L72"></a>t.Errorf(`Equal(%q, %q) = %v`, tt.a, tt.b, eql)
        <a id="L73"></a>}
    <a id="L74"></a>}
<a id="L75"></a>}


<a id="L78"></a>type ExplodeTest struct {
    <a id="L79"></a>s   string;
    <a id="L80"></a>n   int;
    <a id="L81"></a>a   []string;
<a id="L82"></a>}

<a id="L84"></a>var explodetests = []ExplodeTest{
    <a id="L85"></a>ExplodeTest{abcd, 0, []string{&#34;a&#34;, &#34;b&#34;, &#34;c&#34;, &#34;d&#34;}},
    <a id="L86"></a>ExplodeTest{faces, 0, []string{&#34;☺&#34;, &#34;☻&#34;, &#34;☹&#34;}},
    <a id="L87"></a>ExplodeTest{abcd, 2, []string{&#34;a&#34;, &#34;bcd&#34;}},
<a id="L88"></a>}

<a id="L90"></a>func TestExplode(t *testing.T) {
    <a id="L91"></a>for _, tt := range (explodetests) {
        <a id="L92"></a>a := Split(strings.Bytes(tt.s), nil, tt.n);
        <a id="L93"></a>result := arrayOfString(a);
        <a id="L94"></a>if !eq(result, tt.a) {
            <a id="L95"></a>t.Errorf(`Explode(&#34;%s&#34;, %d) = %v; want %v`, tt.s, tt.n, result, tt.a);
            <a id="L96"></a>continue;
        <a id="L97"></a>}
        <a id="L98"></a>s := Join(a, []byte{});
        <a id="L99"></a>if string(s) != tt.s {
            <a id="L100"></a>t.Errorf(`Join(Explode(&#34;%s&#34;, %d), &#34;&#34;) = &#34;%s&#34;`, tt.s, tt.n, s)
        <a id="L101"></a>}
    <a id="L102"></a>}
<a id="L103"></a>}


<a id="L106"></a>type SplitTest struct {
    <a id="L107"></a>s   string;
    <a id="L108"></a>sep string;
    <a id="L109"></a>n   int;
    <a id="L110"></a>a   []string;
<a id="L111"></a>}

<a id="L113"></a>var splittests = []SplitTest{
    <a id="L114"></a>SplitTest{abcd, &#34;a&#34;, 0, []string{&#34;&#34;, &#34;bcd&#34;}},
    <a id="L115"></a>SplitTest{abcd, &#34;z&#34;, 0, []string{&#34;abcd&#34;}},
    <a id="L116"></a>SplitTest{abcd, &#34;&#34;, 0, []string{&#34;a&#34;, &#34;b&#34;, &#34;c&#34;, &#34;d&#34;}},
    <a id="L117"></a>SplitTest{commas, &#34;,&#34;, 0, []string{&#34;1&#34;, &#34;2&#34;, &#34;3&#34;, &#34;4&#34;}},
    <a id="L118"></a>SplitTest{dots, &#34;...&#34;, 0, []string{&#34;1&#34;, &#34;.2&#34;, &#34;.3&#34;, &#34;.4&#34;}},
    <a id="L119"></a>SplitTest{faces, &#34;☹&#34;, 0, []string{&#34;☺☻&#34;, &#34;&#34;}},
    <a id="L120"></a>SplitTest{faces, &#34;~&#34;, 0, []string{faces}},
    <a id="L121"></a>SplitTest{faces, &#34;&#34;, 0, []string{&#34;☺&#34;, &#34;☻&#34;, &#34;☹&#34;}},
    <a id="L122"></a>SplitTest{&#34;1 2 3 4&#34;, &#34; &#34;, 3, []string{&#34;1&#34;, &#34;2&#34;, &#34;3 4&#34;}},
    <a id="L123"></a>SplitTest{&#34;1 2 3&#34;, &#34; &#34;, 3, []string{&#34;1&#34;, &#34;2&#34;, &#34;3&#34;}},
    <a id="L124"></a>SplitTest{&#34;1 2&#34;, &#34; &#34;, 3, []string{&#34;1&#34;, &#34;2&#34;}},
    <a id="L125"></a>SplitTest{&#34;123&#34;, &#34;&#34;, 2, []string{&#34;1&#34;, &#34;23&#34;}},
    <a id="L126"></a>SplitTest{&#34;123&#34;, &#34;&#34;, 17, []string{&#34;1&#34;, &#34;2&#34;, &#34;3&#34;}},
<a id="L127"></a>}

<a id="L129"></a>func TestSplit(t *testing.T) {
    <a id="L130"></a>for _, tt := range splittests {
        <a id="L131"></a>a := Split(strings.Bytes(tt.s), strings.Bytes(tt.sep), tt.n);
        <a id="L132"></a>result := arrayOfString(a);
        <a id="L133"></a>if !eq(result, tt.a) {
            <a id="L134"></a>t.Errorf(`Split(%q, %q, %d) = %v; want %v`, tt.s, tt.sep, tt.n, result, tt.a);
            <a id="L135"></a>continue;
        <a id="L136"></a>}
        <a id="L137"></a>s := Join(a, strings.Bytes(tt.sep));
        <a id="L138"></a>if string(s) != tt.s {
            <a id="L139"></a>t.Errorf(`Join(Split(%q, %q, %d), %q) = %q`, tt.s, tt.sep, tt.n, tt.sep, s)
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
        <a id="L162"></a>a := SplitAfter(strings.Bytes(tt.s), strings.Bytes(tt.sep), tt.n);
        <a id="L163"></a>result := arrayOfString(a);
        <a id="L164"></a>if !eq(result, tt.a) {
            <a id="L165"></a>t.Errorf(`Split(%q, %q, %d) = %v; want %v`, tt.s, tt.sep, tt.n, result, tt.a);
            <a id="L166"></a>continue;
        <a id="L167"></a>}
        <a id="L168"></a>s := Join(a, nil);
        <a id="L169"></a>if string(s) != tt.s {
            <a id="L170"></a>t.Errorf(`Join(Split(%q, %q, %d), %q) = %q`, tt.s, tt.sep, tt.n, tt.sep, s)
        <a id="L171"></a>}
    <a id="L172"></a>}
<a id="L173"></a>}

<a id="L175"></a>type CopyTest struct {
    <a id="L176"></a>a   string;
    <a id="L177"></a>b   string;
    <a id="L178"></a>n   int;
    <a id="L179"></a>res string;
<a id="L180"></a>}

<a id="L182"></a>var copytests = []CopyTest{
    <a id="L183"></a>CopyTest{&#34;&#34;, &#34;&#34;, 0, &#34;&#34;},
    <a id="L184"></a>CopyTest{&#34;a&#34;, &#34;&#34;, 0, &#34;a&#34;},
    <a id="L185"></a>CopyTest{&#34;a&#34;, &#34;a&#34;, 1, &#34;a&#34;},
    <a id="L186"></a>CopyTest{&#34;a&#34;, &#34;b&#34;, 1, &#34;b&#34;},
    <a id="L187"></a>CopyTest{&#34;xyz&#34;, &#34;abc&#34;, 3, &#34;abc&#34;},
    <a id="L188"></a>CopyTest{&#34;wxyz&#34;, &#34;abc&#34;, 3, &#34;abcz&#34;},
    <a id="L189"></a>CopyTest{&#34;xyz&#34;, &#34;abcd&#34;, 3, &#34;abc&#34;},
<a id="L190"></a>}

<a id="L192"></a>func TestCopy(t *testing.T) {
    <a id="L193"></a>for i := 0; i &lt; len(copytests); i++ {
        <a id="L194"></a>tt := copytests[i];
        <a id="L195"></a>dst := strings.Bytes(tt.a);
        <a id="L196"></a>n := Copy(dst, strings.Bytes(tt.b));
        <a id="L197"></a>result := string(dst);
        <a id="L198"></a>if result != tt.res || n != tt.n {
            <a id="L199"></a>t.Errorf(`Copy(%q, %q) = %d, %q; want %d, %q`, tt.a, tt.b, n, result, tt.n, tt.res);
            <a id="L200"></a>continue;
        <a id="L201"></a>}
    <a id="L202"></a>}
<a id="L203"></a>}

<a id="L205"></a><span class="comment">// Test case for any function which accepts and returns a byte array.</span>
<a id="L206"></a><span class="comment">// For ease of creation, we write the byte arrays as strings.</span>
<a id="L207"></a>type StringTest struct {
    <a id="L208"></a>in, out string;
<a id="L209"></a>}

<a id="L211"></a>var upperTests = []StringTest{
    <a id="L212"></a>StringTest{&#34;&#34;, &#34;&#34;},
    <a id="L213"></a>StringTest{&#34;abc&#34;, &#34;ABC&#34;},
    <a id="L214"></a>StringTest{&#34;AbC123&#34;, &#34;ABC123&#34;},
    <a id="L215"></a>StringTest{&#34;azAZ09_&#34;, &#34;AZAZ09_&#34;},
    <a id="L216"></a>StringTest{&#34;\u0250\u0250\u0250\u0250\u0250&#34;, &#34;\u2C6F\u2C6F\u2C6F\u2C6F\u2C6F&#34;}, <span class="comment">// grows one byte per char</span>
<a id="L217"></a>}

<a id="L219"></a>var lowerTests = []StringTest{
    <a id="L220"></a>StringTest{&#34;&#34;, &#34;&#34;},
    <a id="L221"></a>StringTest{&#34;abc&#34;, &#34;abc&#34;},
    <a id="L222"></a>StringTest{&#34;AbC123&#34;, &#34;abc123&#34;},
    <a id="L223"></a>StringTest{&#34;azAZ09_&#34;, &#34;azaz09_&#34;},
    <a id="L224"></a>StringTest{&#34;\u2C6D\u2C6D\u2C6D\u2C6D\u2C6D&#34;, &#34;\u0251\u0251\u0251\u0251\u0251&#34;}, <span class="comment">// shrinks one byte per char</span>
<a id="L225"></a>}

<a id="L227"></a>const space = &#34;\t\v\r\f\n\u0085\u00a0\u2000\u3000&#34;

<a id="L229"></a>var trimSpaceTests = []StringTest{
    <a id="L230"></a>StringTest{&#34;&#34;, &#34;&#34;},
    <a id="L231"></a>StringTest{&#34;abc&#34;, &#34;abc&#34;},
    <a id="L232"></a>StringTest{space + &#34;abc&#34; + space, &#34;abc&#34;},
    <a id="L233"></a>StringTest{&#34; &#34;, &#34;&#34;},
    <a id="L234"></a>StringTest{&#34; \t\r\n \t\t\r\r\n\n &#34;, &#34;&#34;},
    <a id="L235"></a>StringTest{&#34; \t\r\n x\t\t\r\r\n\n &#34;, &#34;x&#34;},
    <a id="L236"></a>StringTest{&#34; \u2000\t\r\n x\t\t\r\r\ny\n \u3000&#34;, &#34;x\t\t\r\r\ny&#34;},
    <a id="L237"></a>StringTest{&#34;1 \t\r\n2&#34;, &#34;1 \t\r\n2&#34;},
    <a id="L238"></a>StringTest{&#34; x\x80&#34;, &#34;x\x80&#34;}, <span class="comment">// invalid UTF-8 on end</span>
    <a id="L239"></a>StringTest{&#34; x\xc0&#34;, &#34;x\xc0&#34;}, <span class="comment">// invalid UTF-8 on end</span>
<a id="L240"></a>}

<a id="L242"></a><span class="comment">// Bytes returns a new slice containing the bytes in s.</span>
<a id="L243"></a><span class="comment">// Borrowed from strings to avoid dependency.</span>
<a id="L244"></a>func Bytes(s string) []byte {
    <a id="L245"></a>b := make([]byte, len(s));
    <a id="L246"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L247"></a>b[i] = s[i]
    <a id="L248"></a>}
    <a id="L249"></a>return b;
<a id="L250"></a>}

<a id="L252"></a><span class="comment">// Execute f on each test case.  funcName should be the name of f; it&#39;s used</span>
<a id="L253"></a><span class="comment">// in failure reports.</span>
<a id="L254"></a>func runStringTests(t *testing.T, f func([]byte) []byte, funcName string, testCases []StringTest) {
    <a id="L255"></a>for _, tc := range testCases {
        <a id="L256"></a>actual := string(f(Bytes(tc.in)));
        <a id="L257"></a>if actual != tc.out {
            <a id="L258"></a>t.Errorf(&#34;%s(%q) = %q; want %q&#34;, funcName, tc.in, actual, tc.out)
        <a id="L259"></a>}
    <a id="L260"></a>}
<a id="L261"></a>}

<a id="L263"></a>func tenRunes(rune int) string {
    <a id="L264"></a>r := make([]int, 10);
    <a id="L265"></a>for i := range r {
        <a id="L266"></a>r[i] = rune
    <a id="L267"></a>}
    <a id="L268"></a>return string(r);
<a id="L269"></a>}

<a id="L271"></a>func TestMap(t *testing.T) {
    <a id="L272"></a><span class="comment">// Run a couple of awful growth/shrinkage tests</span>
    <a id="L273"></a>a := tenRunes(&#39;a&#39;);
    <a id="L274"></a><span class="comment">// 1.  Grow.  This triggers two reallocations in Map.</span>
    <a id="L275"></a>maxRune := func(rune int) int { return unicode.MaxRune };
    <a id="L276"></a>m := Map(maxRune, Bytes(a));
    <a id="L277"></a>expect := tenRunes(unicode.MaxRune);
    <a id="L278"></a>if string(m) != expect {
        <a id="L279"></a>t.Errorf(&#34;growing: expected %q got %q&#34;, expect, m)
    <a id="L280"></a>}
    <a id="L281"></a><span class="comment">// 2. Shrink</span>
    <a id="L282"></a>minRune := func(rune int) int { return &#39;a&#39; };
    <a id="L283"></a>m = Map(minRune, Bytes(tenRunes(unicode.MaxRune)));
    <a id="L284"></a>expect = a;
    <a id="L285"></a>if string(m) != expect {
        <a id="L286"></a>t.Errorf(&#34;shrinking: expected %q got %q&#34;, expect, m)
    <a id="L287"></a>}
<a id="L288"></a>}

<a id="L290"></a>func TestToUpper(t *testing.T) { runStringTests(t, ToUpper, &#34;ToUpper&#34;, upperTests) }

<a id="L292"></a>func TestToLower(t *testing.T) { runStringTests(t, ToLower, &#34;ToLower&#34;, lowerTests) }

<a id="L294"></a>func TestTrimSpace(t *testing.T) { runStringTests(t, TrimSpace, &#34;TrimSpace&#34;, trimSpaceTests) }

<a id="L296"></a>type AddTest struct {
    <a id="L297"></a>s, t string;
    <a id="L298"></a>cap  int;
<a id="L299"></a>}

<a id="L301"></a>var addtests = []AddTest{
    <a id="L302"></a>AddTest{&#34;&#34;, &#34;&#34;, 0},
    <a id="L303"></a>AddTest{&#34;a&#34;, &#34;&#34;, 1},
    <a id="L304"></a>AddTest{&#34;a&#34;, &#34;b&#34;, 1},
    <a id="L305"></a>AddTest{&#34;abc&#34;, &#34;def&#34;, 100},
<a id="L306"></a>}

<a id="L308"></a>func TestAdd(t *testing.T) {
    <a id="L309"></a>for _, test := range addtests {
        <a id="L310"></a>b := make([]byte, len(test.s), test.cap);
        <a id="L311"></a>for i := 0; i &lt; len(test.s); i++ {
            <a id="L312"></a>b[i] = test.s[i]
        <a id="L313"></a>}
        <a id="L314"></a>b = Add(b, strings.Bytes(test.t));
        <a id="L315"></a>if string(b) != test.s+test.t {
            <a id="L316"></a>t.Errorf(&#34;Add(%q,%q) = %q&#34;, test.s, test.t, string(b))
        <a id="L317"></a>}
    <a id="L318"></a>}
<a id="L319"></a>}

<a id="L321"></a>func TestAddByte(t *testing.T) {
    <a id="L322"></a>const N = 2e5;
    <a id="L323"></a>b := make([]byte, 0);
    <a id="L324"></a>for i := 0; i &lt; N; i++ {
        <a id="L325"></a>b = AddByte(b, byte(i))
    <a id="L326"></a>}
    <a id="L327"></a>if len(b) != N {
        <a id="L328"></a>t.Errorf(&#34;AddByte: too small; expected %d got %d&#34;, N, len(b))
    <a id="L329"></a>}
    <a id="L330"></a>for i, c := range b {
        <a id="L331"></a>if c != byte(i) {
            <a id="L332"></a>t.Fatalf(&#34;AddByte: b[%d] should be %d is %d&#34;, i, c, byte(i))
        <a id="L333"></a>}
    <a id="L334"></a>}
<a id="L335"></a>}
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
