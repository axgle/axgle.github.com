<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/tls/record_read_test.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/tls/record_read_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package tls

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;testing&#34;;
    <a id="L10"></a>&#34;testing/iotest&#34;;
<a id="L11"></a>)

<a id="L13"></a>func matchRecord(r1, r2 *record) bool {
    <a id="L14"></a>if (r1 == nil) != (r2 == nil) {
        <a id="L15"></a>return false
    <a id="L16"></a>}
    <a id="L17"></a>if r1 == nil {
        <a id="L18"></a>return true
    <a id="L19"></a>}
    <a id="L20"></a>return r1.contentType == r2.contentType &amp;&amp;
        <a id="L21"></a>r1.major == r2.major &amp;&amp;
        <a id="L22"></a>r1.minor == r2.minor &amp;&amp;
        <a id="L23"></a>bytes.Compare(r1.payload, r2.payload) == 0;
<a id="L24"></a>}

<a id="L26"></a>type recordReaderTest struct {
    <a id="L27"></a>in  []byte;
    <a id="L28"></a>out []*record;
<a id="L29"></a>}

<a id="L31"></a>var recordReaderTests = []recordReaderTest{
    <a id="L32"></a>recordReaderTest{nil, nil},
    <a id="L33"></a>recordReaderTest{fromHex(&#34;01&#34;), nil},
    <a id="L34"></a>recordReaderTest{fromHex(&#34;0102&#34;), nil},
    <a id="L35"></a>recordReaderTest{fromHex(&#34;010203&#34;), nil},
    <a id="L36"></a>recordReaderTest{fromHex(&#34;01020300&#34;), nil},
    <a id="L37"></a>recordReaderTest{fromHex(&#34;0102030000&#34;), []*record{&amp;record{1, 2, 3, nil}}},
    <a id="L38"></a>recordReaderTest{fromHex(&#34;01020300000102030000&#34;), []*record{&amp;record{1, 2, 3, nil}, &amp;record{1, 2, 3, nil}}},
    <a id="L39"></a>recordReaderTest{fromHex(&#34;0102030001fe0102030002feff&#34;), []*record{&amp;record{1, 2, 3, []byte{0xfe}}, &amp;record{1, 2, 3, []byte{0xfe, 0xff}}}},
    <a id="L40"></a>recordReaderTest{fromHex(&#34;010203000001020300&#34;), []*record{&amp;record{1, 2, 3, nil}}},
<a id="L41"></a>}

<a id="L43"></a>func TestRecordReader(t *testing.T) {
    <a id="L44"></a>for i, test := range recordReaderTests {
        <a id="L45"></a>buf := bytes.NewBuffer(test.in);
        <a id="L46"></a>c := make(chan *record);
        <a id="L47"></a>go recordReader(c, buf);
        <a id="L48"></a>matchRecordReaderOutput(t, i, test, c);

        <a id="L50"></a>buf = bytes.NewBuffer(test.in);
        <a id="L51"></a>buf2 := iotest.OneByteReader(buf);
        <a id="L52"></a>c = make(chan *record);
        <a id="L53"></a>go recordReader(c, buf2);
        <a id="L54"></a>matchRecordReaderOutput(t, i*2, test, c);
    <a id="L55"></a>}
<a id="L56"></a>}

<a id="L58"></a>func matchRecordReaderOutput(t *testing.T, i int, test recordReaderTest, c &lt;-chan *record) {
    <a id="L59"></a>for j, r1 := range test.out {
        <a id="L60"></a>r2 := &lt;-c;
        <a id="L61"></a>if r2 == nil {
            <a id="L62"></a>t.Errorf(&#34;#%d truncated after %d values&#34;, i, j);
            <a id="L63"></a>break;
        <a id="L64"></a>}
        <a id="L65"></a>if !matchRecord(r1, r2) {
            <a id="L66"></a>t.Errorf(&#34;#%d (%d) got:%#v want:%#v&#34;, i, j, r2, r1)
        <a id="L67"></a>}
    <a id="L68"></a>}
    <a id="L69"></a>&lt;-c;
    <a id="L70"></a>if !closed(c) {
        <a id="L71"></a>t.Errorf(&#34;#%d: channel didn&#39;t close&#34;, i)
    <a id="L72"></a>}
<a id="L73"></a>}
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
