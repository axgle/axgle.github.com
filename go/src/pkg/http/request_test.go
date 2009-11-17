<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/http/request_test.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/http/request_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package http

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a>type stringMultimap map[string][]string

<a id="L14"></a>type parseTest struct {
    <a id="L15"></a>query string;
    <a id="L16"></a>out   stringMultimap;
<a id="L17"></a>}

<a id="L19"></a>var parseTests = []parseTest{
    <a id="L20"></a>parseTest{
        <a id="L21"></a>query: &#34;a=1&amp;b=2&#34;,
        <a id="L22"></a>out: stringMultimap{&#34;a&#34;: []string{&#34;1&#34;}, &#34;b&#34;: []string{&#34;2&#34;}},
    <a id="L23"></a>},
    <a id="L24"></a>parseTest{
        <a id="L25"></a>query: &#34;a=1&amp;a=2&amp;a=banana&#34;,
        <a id="L26"></a>out: stringMultimap{&#34;a&#34;: []string{&#34;1&#34;, &#34;2&#34;, &#34;banana&#34;}},
    <a id="L27"></a>},
    <a id="L28"></a>parseTest{
        <a id="L29"></a>query: &#34;ascii=%3Ckey%3A+0x90%3E&#34;,
        <a id="L30"></a>out: stringMultimap{&#34;ascii&#34;: []string{&#34;&lt;key: 0x90&gt;&#34;}},
    <a id="L31"></a>},
<a id="L32"></a>}

<a id="L34"></a>func TestParseForm(t *testing.T) {
    <a id="L35"></a>for i, test := range parseTests {
        <a id="L36"></a>form, err := parseForm(test.query);
        <a id="L37"></a>if err != nil {
            <a id="L38"></a>t.Errorf(&#34;test %d: Unexpected error: %v&#34;, i, err);
            <a id="L39"></a>continue;
        <a id="L40"></a>}
        <a id="L41"></a>if len(form) != len(test.out) {
            <a id="L42"></a>t.Errorf(&#34;test %d: len(form) = %d, want %d&#34;, i, len(form), len(test.out))
        <a id="L43"></a>}
        <a id="L44"></a>for k, evs := range test.out {
            <a id="L45"></a>vs, ok := form[k];
            <a id="L46"></a>if !ok {
                <a id="L47"></a>t.Errorf(&#34;test %d: Missing key %q&#34;, i, k);
                <a id="L48"></a>continue;
            <a id="L49"></a>}
            <a id="L50"></a>if len(vs) != len(evs) {
                <a id="L51"></a>t.Errorf(&#34;test %d: len(form[%q]) = %d, want %d&#34;, i, k, len(vs), len(evs));
                <a id="L52"></a>continue;
            <a id="L53"></a>}
            <a id="L54"></a>for j, ev := range evs {
                <a id="L55"></a>if v := vs[j]; v != ev {
                    <a id="L56"></a>t.Errorf(&#34;test %d: form[%q][%d] = %q, want %q&#34;, i, k, j, v, ev)
                <a id="L57"></a>}
            <a id="L58"></a>}
        <a id="L59"></a>}
    <a id="L60"></a>}
<a id="L61"></a>}

<a id="L63"></a>func TestQuery(t *testing.T) {
    <a id="L64"></a>req := &amp;Request{Method: &#34;GET&#34;};
    <a id="L65"></a>req.URL, _ = ParseURL(&#34;http://www.google.com/search?q=foo&amp;q=bar&#34;);
    <a id="L66"></a>if q := req.FormValue(&#34;q&#34;); q != &#34;foo&#34; {
        <a id="L67"></a>t.Errorf(`req.FormValue(&#34;q&#34;) = %q, want &#34;foo&#34;`, q)
    <a id="L68"></a>}
<a id="L69"></a>}

<a id="L71"></a>type stringMap map[string]string
<a id="L72"></a>type parseContentTypeTest struct {
    <a id="L73"></a>contentType stringMap;
    <a id="L74"></a>error       bool;
<a id="L75"></a>}

<a id="L77"></a>var parseContentTypeTests = []parseContentTypeTest{
    <a id="L78"></a>parseContentTypeTest{contentType: stringMap{&#34;Content-Type&#34;: &#34;text/plain&#34;}},
    <a id="L79"></a>parseContentTypeTest{contentType: stringMap{&#34;Content-Type&#34;: &#34;&#34;}},
    <a id="L80"></a>parseContentTypeTest{contentType: stringMap{&#34;Content-Type&#34;: &#34;text/plain; boundary=&#34;}},
    <a id="L81"></a>parseContentTypeTest{
        <a id="L82"></a>contentType: stringMap{&#34;Content-Type&#34;: &#34;application/unknown&#34;},
        <a id="L83"></a>error: true,
    <a id="L84"></a>},
<a id="L85"></a>}

<a id="L87"></a>func TestPostContentTypeParsing(t *testing.T) {
    <a id="L88"></a>for i, test := range parseContentTypeTests {
        <a id="L89"></a>req := &amp;Request{
            <a id="L90"></a>Method: &#34;POST&#34;,
            <a id="L91"></a>Header: test.contentType,
            <a id="L92"></a>Body: bytes.NewBufferString(&#34;body&#34;),
        <a id="L93"></a>};
        <a id="L94"></a>err := req.ParseForm();
        <a id="L95"></a>if !test.error &amp;&amp; err != nil {
            <a id="L96"></a>t.Errorf(&#34;test %d: Unexpected error: %v&#34;, i, err)
        <a id="L97"></a>}
        <a id="L98"></a>if test.error &amp;&amp; err == nil {
            <a id="L99"></a>t.Errorf(&#34;test %d should have returned error&#34;, i)
        <a id="L100"></a>}
    <a id="L101"></a>}
<a id="L102"></a>}

<a id="L104"></a>func TestRedirect(t *testing.T) {
    <a id="L105"></a>const (
        <a id="L106"></a>start = &#34;http://codesearch.google.com/&#34;;
        <a id="L107"></a>end   = &#34;http://www.google.com/codesearch&#34;;
    <a id="L108"></a>)
    <a id="L109"></a>r, url, err := Get(start);
    <a id="L110"></a>if err != nil {
        <a id="L111"></a>t.Fatal(err)
    <a id="L112"></a>}
    <a id="L113"></a>r.Body.Close();
    <a id="L114"></a>if r.StatusCode != 200 || url != end {
        <a id="L115"></a>t.Fatalf(&#34;Get(%s) got status %d at %s, want 200 at %s&#34;, start, r.StatusCode, url, end)
    <a id="L116"></a>}
<a id="L117"></a>}
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
