<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/archive/tar/reader_test.go</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/archive/tar/reader_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package tar

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;io&#34;;
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;reflect&#34;;
    <a id="L12"></a>&#34;strings&#34;;
    <a id="L13"></a>&#34;testing&#34;;
<a id="L14"></a>)

<a id="L16"></a>type untarTest struct {
    <a id="L17"></a>file    string;
    <a id="L18"></a>headers []*Header;
<a id="L19"></a>}

<a id="L21"></a>var untarTests = []*untarTest{
    <a id="L22"></a>&amp;untarTest{
        <a id="L23"></a>file: &#34;testdata/gnu.tar&#34;,
        <a id="L24"></a>headers: []*Header{
            <a id="L25"></a>&amp;Header{
                <a id="L26"></a>Name: &#34;small.txt&#34;,
                <a id="L27"></a>Mode: 0640,
                <a id="L28"></a>Uid: 73025,
                <a id="L29"></a>Gid: 5000,
                <a id="L30"></a>Size: 5,
                <a id="L31"></a>Mtime: 1244428340,
                <a id="L32"></a>Typeflag: &#39;0&#39;,
                <a id="L33"></a>Uname: &#34;dsymonds&#34;,
                <a id="L34"></a>Gname: &#34;eng&#34;,
            <a id="L35"></a>},
            <a id="L36"></a>&amp;Header{
                <a id="L37"></a>Name: &#34;small2.txt&#34;,
                <a id="L38"></a>Mode: 0640,
                <a id="L39"></a>Uid: 73025,
                <a id="L40"></a>Gid: 5000,
                <a id="L41"></a>Size: 11,
                <a id="L42"></a>Mtime: 1244436044,
                <a id="L43"></a>Typeflag: &#39;0&#39;,
                <a id="L44"></a>Uname: &#34;dsymonds&#34;,
                <a id="L45"></a>Gname: &#34;eng&#34;,
            <a id="L46"></a>},
        <a id="L47"></a>},
    <a id="L48"></a>},
    <a id="L49"></a>&amp;untarTest{
        <a id="L50"></a>file: &#34;testdata/star.tar&#34;,
        <a id="L51"></a>headers: []*Header{
            <a id="L52"></a>&amp;Header{
                <a id="L53"></a>Name: &#34;small.txt&#34;,
                <a id="L54"></a>Mode: 0640,
                <a id="L55"></a>Uid: 73025,
                <a id="L56"></a>Gid: 5000,
                <a id="L57"></a>Size: 5,
                <a id="L58"></a>Mtime: 1244592783,
                <a id="L59"></a>Typeflag: &#39;0&#39;,
                <a id="L60"></a>Uname: &#34;dsymonds&#34;,
                <a id="L61"></a>Gname: &#34;eng&#34;,
                <a id="L62"></a>Atime: 1244592783,
                <a id="L63"></a>Ctime: 1244592783,
            <a id="L64"></a>},
            <a id="L65"></a>&amp;Header{
                <a id="L66"></a>Name: &#34;small2.txt&#34;,
                <a id="L67"></a>Mode: 0640,
                <a id="L68"></a>Uid: 73025,
                <a id="L69"></a>Gid: 5000,
                <a id="L70"></a>Size: 11,
                <a id="L71"></a>Mtime: 1244592783,
                <a id="L72"></a>Typeflag: &#39;0&#39;,
                <a id="L73"></a>Uname: &#34;dsymonds&#34;,
                <a id="L74"></a>Gname: &#34;eng&#34;,
                <a id="L75"></a>Atime: 1244592783,
                <a id="L76"></a>Ctime: 1244592783,
            <a id="L77"></a>},
        <a id="L78"></a>},
    <a id="L79"></a>},
    <a id="L80"></a>&amp;untarTest{
        <a id="L81"></a>file: &#34;testdata/v7.tar&#34;,
        <a id="L82"></a>headers: []*Header{
            <a id="L83"></a>&amp;Header{
                <a id="L84"></a>Name: &#34;small.txt&#34;,
                <a id="L85"></a>Mode: 0444,
                <a id="L86"></a>Uid: 73025,
                <a id="L87"></a>Gid: 5000,
                <a id="L88"></a>Size: 5,
                <a id="L89"></a>Mtime: 1244593104,
                <a id="L90"></a>Typeflag: &#39;\x00&#39;,
            <a id="L91"></a>},
            <a id="L92"></a>&amp;Header{
                <a id="L93"></a>Name: &#34;small2.txt&#34;,
                <a id="L94"></a>Mode: 0444,
                <a id="L95"></a>Uid: 73025,
                <a id="L96"></a>Gid: 5000,
                <a id="L97"></a>Size: 11,
                <a id="L98"></a>Mtime: 1244593104,
                <a id="L99"></a>Typeflag: &#39;\x00&#39;,
            <a id="L100"></a>},
        <a id="L101"></a>},
    <a id="L102"></a>},
<a id="L103"></a>}

<a id="L105"></a>func TestReader(t *testing.T) {
<a id="L106"></a>testLoop:
    <a id="L107"></a>for i, test := range untarTests {
        <a id="L108"></a>f, err := os.Open(test.file, os.O_RDONLY, 0444);
        <a id="L109"></a>if err != nil {
            <a id="L110"></a>t.Errorf(&#34;test %d: Unexpected error: %v&#34;, i, err);
            <a id="L111"></a>continue;
        <a id="L112"></a>}
        <a id="L113"></a>tr := NewReader(f);
        <a id="L114"></a>for j, header := range test.headers {
            <a id="L115"></a>hdr, err := tr.Next();
            <a id="L116"></a>if err != nil || hdr == nil {
                <a id="L117"></a>t.Errorf(&#34;test %d, entry %d: Didn&#39;t get entry: %v&#34;, i, j, err);
                <a id="L118"></a>f.Close();
                <a id="L119"></a>continue testLoop;
            <a id="L120"></a>}
            <a id="L121"></a>if !reflect.DeepEqual(hdr, header) {
                <a id="L122"></a>t.Errorf(&#34;test %d, entry %d: Incorrect header:\nhave %+v\nwant %+v&#34;,
                    <a id="L123"></a>i, j, *hdr, *header)
            <a id="L124"></a>}
        <a id="L125"></a>}
        <a id="L126"></a>hdr, err := tr.Next();
        <a id="L127"></a>if hdr != nil || err != nil {
            <a id="L128"></a>t.Errorf(&#34;test %d: Unexpected entry or error: hdr=%v err=%v&#34;, i, err)
        <a id="L129"></a>}
        <a id="L130"></a>f.Close();
    <a id="L131"></a>}
<a id="L132"></a>}

<a id="L134"></a>func TestPartialRead(t *testing.T) {
    <a id="L135"></a>f, err := os.Open(&#34;testdata/gnu.tar&#34;, os.O_RDONLY, 0444);
    <a id="L136"></a>if err != nil {
        <a id="L137"></a>t.Fatalf(&#34;Unexpected error: %v&#34;, err)
    <a id="L138"></a>}
    <a id="L139"></a>defer f.Close();

    <a id="L141"></a>tr := NewReader(f);

    <a id="L143"></a><span class="comment">// Read the first four bytes; Next() should skip the last byte.</span>
    <a id="L144"></a>hdr, err := tr.Next();
    <a id="L145"></a>if err != nil || hdr == nil {
        <a id="L146"></a>t.Fatalf(&#34;Didn&#39;t get first file: %v&#34;, err)
    <a id="L147"></a>}
    <a id="L148"></a>buf := make([]byte, 4);
    <a id="L149"></a>if _, err := io.ReadFull(tr, buf); err != nil {
        <a id="L150"></a>t.Fatalf(&#34;Unexpected error: %v&#34;, err)
    <a id="L151"></a>}
    <a id="L152"></a>if expected := strings.Bytes(&#34;Kilt&#34;); !bytes.Equal(buf, expected) {
        <a id="L153"></a>t.Errorf(&#34;Contents = %v, want %v&#34;, buf, expected)
    <a id="L154"></a>}

    <a id="L156"></a><span class="comment">// Second file</span>
    <a id="L157"></a>hdr, err = tr.Next();
    <a id="L158"></a>if err != nil || hdr == nil {
        <a id="L159"></a>t.Fatalf(&#34;Didn&#39;t get second file: %v&#34;, err)
    <a id="L160"></a>}
    <a id="L161"></a>buf = make([]byte, 6);
    <a id="L162"></a>if _, err := io.ReadFull(tr, buf); err != nil {
        <a id="L163"></a>t.Fatalf(&#34;Unexpected error: %v&#34;, err)
    <a id="L164"></a>}
    <a id="L165"></a>if expected := strings.Bytes(&#34;Google&#34;); !bytes.Equal(buf, expected) {
        <a id="L166"></a>t.Errorf(&#34;Contents = %v, want %v&#34;, buf, expected)
    <a id="L167"></a>}
<a id="L168"></a>}
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
