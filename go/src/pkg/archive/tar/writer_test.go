<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/archive/tar/writer_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/archive/tar/writer_test.go</h1>

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
    <a id="L9"></a>&#34;fmt&#34;;
    <a id="L10"></a>&#34;io&#34;;
    <a id="L11"></a>&#34;testing&#34;;
    <a id="L12"></a>&#34;testing/iotest&#34;;
<a id="L13"></a>)

<a id="L15"></a>type writerTestEntry struct {
    <a id="L16"></a>header   *Header;
    <a id="L17"></a>contents string;
<a id="L18"></a>}

<a id="L20"></a>type writerTest struct {
    <a id="L21"></a>file    string; <span class="comment">// filename of expected output</span>
    <a id="L22"></a>entries []*writerTestEntry;
<a id="L23"></a>}

<a id="L25"></a>var writerTests = []*writerTest{
    <a id="L26"></a>&amp;writerTest{
        <a id="L27"></a>file: &#34;testdata/writer.tar&#34;,
        <a id="L28"></a>entries: []*writerTestEntry{
            <a id="L29"></a>&amp;writerTestEntry{
                <a id="L30"></a>header: &amp;Header{
                    <a id="L31"></a>Name: &#34;small.txt&#34;,
                    <a id="L32"></a>Mode: 0640,
                    <a id="L33"></a>Uid: 73025,
                    <a id="L34"></a>Gid: 5000,
                    <a id="L35"></a>Size: 5,
                    <a id="L36"></a>Mtime: 1246508266,
                    <a id="L37"></a>Typeflag: &#39;0&#39;,
                    <a id="L38"></a>Uname: &#34;dsymonds&#34;,
                    <a id="L39"></a>Gname: &#34;eng&#34;,
                <a id="L40"></a>},
                <a id="L41"></a>contents: &#34;Kilts&#34;,
            <a id="L42"></a>},
            <a id="L43"></a>&amp;writerTestEntry{
                <a id="L44"></a>header: &amp;Header{
                    <a id="L45"></a>Name: &#34;small2.txt&#34;,
                    <a id="L46"></a>Mode: 0640,
                    <a id="L47"></a>Uid: 73025,
                    <a id="L48"></a>Gid: 5000,
                    <a id="L49"></a>Size: 11,
                    <a id="L50"></a>Mtime: 1245217492,
                    <a id="L51"></a>Typeflag: &#39;0&#39;,
                    <a id="L52"></a>Uname: &#34;dsymonds&#34;,
                    <a id="L53"></a>Gname: &#34;eng&#34;,
                <a id="L54"></a>},
                <a id="L55"></a>contents: &#34;Google.com\n&#34;,
            <a id="L56"></a>},
        <a id="L57"></a>},
    <a id="L58"></a>},
    <a id="L59"></a><span class="comment">// The truncated test file was produced using these commands:</span>
    <a id="L60"></a><span class="comment">//   dd if=/dev/zero bs=1048576 count=16384 &gt; /tmp/16gig.txt</span>
    <a id="L61"></a><span class="comment">//   tar -b 1 -c -f- /tmp/16gig.txt | dd bs=512 count=8 &gt; writer-big.tar</span>
    <a id="L62"></a>&amp;writerTest{
        <a id="L63"></a>file: &#34;testdata/writer-big.tar&#34;,
        <a id="L64"></a>entries: []*writerTestEntry{
            <a id="L65"></a>&amp;writerTestEntry{
                <a id="L66"></a>header: &amp;Header{
                    <a id="L67"></a>Name: &#34;tmp/16gig.txt&#34;,
                    <a id="L68"></a>Mode: 0640,
                    <a id="L69"></a>Uid: 73025,
                    <a id="L70"></a>Gid: 5000,
                    <a id="L71"></a>Size: 16 &lt;&lt; 30,
                    <a id="L72"></a>Mtime: 1254699560,
                    <a id="L73"></a>Typeflag: &#39;0&#39;,
                    <a id="L74"></a>Uname: &#34;dsymonds&#34;,
                    <a id="L75"></a>Gname: &#34;eng&#34;,
                <a id="L76"></a>},
                <a id="L77"></a><span class="comment">// no contents</span>
            <a id="L78"></a>},
        <a id="L79"></a>},
    <a id="L80"></a>},
<a id="L81"></a>}

<a id="L83"></a><span class="comment">// Render byte array in a two-character hexadecimal string, spaced for easy visual inspection.</span>
<a id="L84"></a>func bytestr(offset int, b []byte) string {
    <a id="L85"></a>const rowLen = 32;
    <a id="L86"></a>s := fmt.Sprintf(&#34;%04x &#34;, offset);
    <a id="L87"></a>for _, ch := range b {
        <a id="L88"></a>switch {
        <a id="L89"></a>case &#39;0&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;9&#39;, &#39;A&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;Z&#39;, &#39;a&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;z&#39;:
            <a id="L90"></a>s += fmt.Sprintf(&#34;  %c&#34;, ch)
        <a id="L91"></a>default:
            <a id="L92"></a>s += fmt.Sprintf(&#34; %02x&#34;, ch)
        <a id="L93"></a>}
    <a id="L94"></a>}
    <a id="L95"></a>return s;
<a id="L96"></a>}

<a id="L98"></a><span class="comment">// Render a pseudo-diff between two blocks of bytes.</span>
<a id="L99"></a>func bytediff(a []byte, b []byte) string {
    <a id="L100"></a>const rowLen = 32;
    <a id="L101"></a>s := fmt.Sprintf(&#34;(%d bytes vs. %d bytes)\n&#34;, len(a), len(b));
    <a id="L102"></a>for offset := 0; len(a)+len(b) &gt; 0; offset += rowLen {
        <a id="L103"></a>na, nb := rowLen, rowLen;
        <a id="L104"></a>if na &gt; len(a) {
            <a id="L105"></a>na = len(a)
        <a id="L106"></a>}
        <a id="L107"></a>if nb &gt; len(b) {
            <a id="L108"></a>nb = len(b)
        <a id="L109"></a>}
        <a id="L110"></a>sa := bytestr(offset, a[0:na]);
        <a id="L111"></a>sb := bytestr(offset, b[0:nb]);
        <a id="L112"></a>if sa != sb {
            <a id="L113"></a>s += fmt.Sprintf(&#34;-%v\n+%v\n&#34;, sa, sb)
        <a id="L114"></a>}
        <a id="L115"></a>a = a[na:len(a)];
        <a id="L116"></a>b = b[nb:len(b)];
    <a id="L117"></a>}
    <a id="L118"></a>return s;
<a id="L119"></a>}

<a id="L121"></a>func TestWriter(t *testing.T) {
<a id="L122"></a>testLoop:
    <a id="L123"></a>for i, test := range writerTests {
        <a id="L124"></a>expected, err := io.ReadFile(test.file);
        <a id="L125"></a>if err != nil {
            <a id="L126"></a>t.Errorf(&#34;test %d: Unexpected error: %v&#34;, i, err);
            <a id="L127"></a>continue;
        <a id="L128"></a>}

        <a id="L130"></a>buf := new(bytes.Buffer);
        <a id="L131"></a>tw := NewWriter(iotest.TruncateWriter(buf, 4&lt;&lt;10)); <span class="comment">// only catch the first 4 KB</span>
        <a id="L132"></a>for j, entry := range test.entries {
            <a id="L133"></a>if err := tw.WriteHeader(entry.header); err != nil {
                <a id="L134"></a>t.Errorf(&#34;test %d, entry %d: Failed writing header: %v&#34;, i, j, err);
                <a id="L135"></a>continue testLoop;
            <a id="L136"></a>}
            <a id="L137"></a>if _, err := io.WriteString(tw, entry.contents); err != nil {
                <a id="L138"></a>t.Errorf(&#34;test %d, entry %d: Failed writing contents: %v&#34;, i, j, err);
                <a id="L139"></a>continue testLoop;
            <a id="L140"></a>}
        <a id="L141"></a>}
        <a id="L142"></a>if err := tw.Close(); err != nil {
            <a id="L143"></a>t.Errorf(&#34;test %d: Failed closing archive: %v&#34;, err);
            <a id="L144"></a>continue testLoop;
        <a id="L145"></a>}

        <a id="L147"></a>actual := buf.Bytes();
        <a id="L148"></a>if !bytes.Equal(expected, actual) {
            <a id="L149"></a>t.Errorf(&#34;test %d: Incorrect result: (-=expected, +=actual)\n%v&#34;,
                <a id="L150"></a>i, bytediff(expected, actual))
        <a id="L151"></a>}
    <a id="L152"></a>}
<a id="L153"></a>}
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
