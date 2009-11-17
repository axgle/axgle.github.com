<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/compress/flate/flate_test.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/compress/flate/flate_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This test tests some internals of the flate package.</span>
<a id="L6"></a><span class="comment">// The tests in package compress/gzip serve as the</span>
<a id="L7"></a><span class="comment">// end-to-end test of the inflater.</span>

<a id="L9"></a>package flate

<a id="L11"></a>import (
    <a id="L12"></a>&#34;bytes&#34;;
    <a id="L13"></a>&#34;reflect&#34;;
    <a id="L14"></a>&#34;testing&#34;;
<a id="L15"></a>)

<a id="L17"></a><span class="comment">// The Huffman code lengths used by the fixed-format Huffman blocks.</span>
<a id="L18"></a>var fixedHuffmanBits = [...]int{
    <a id="L19"></a><span class="comment">// 0-143 length 8</span>
    <a id="L20"></a>8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
    <a id="L21"></a>8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
    <a id="L22"></a>8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
    <a id="L23"></a>8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
    <a id="L24"></a>8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
    <a id="L25"></a>8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
    <a id="L26"></a>8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
    <a id="L27"></a>8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
    <a id="L28"></a>8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,

    <a id="L30"></a><span class="comment">// 144-255 length 9</span>
    <a id="L31"></a>9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
    <a id="L32"></a>9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
    <a id="L33"></a>9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
    <a id="L34"></a>9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
    <a id="L35"></a>9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
    <a id="L36"></a>9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
    <a id="L37"></a>9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,

    <a id="L39"></a><span class="comment">// 256-279 length 7</span>
    <a id="L40"></a>7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
    <a id="L41"></a>7, 7, 7, 7, 7, 7, 7, 7,

    <a id="L43"></a><span class="comment">// 280-287 length 8</span>
    <a id="L44"></a>8, 8, 8, 8, 8, 8, 8, 8,
<a id="L45"></a>}

<a id="L47"></a>type InitDecoderTest struct {
    <a id="L48"></a>in  []int;
    <a id="L49"></a>out huffmanDecoder;
    <a id="L50"></a>ok  bool;
<a id="L51"></a>}

<a id="L53"></a>var initDecoderTests = []*InitDecoderTest{
    <a id="L54"></a><span class="comment">// Example from Connell 1973,</span>
    <a id="L55"></a>&amp;InitDecoderTest{
        <a id="L56"></a>[]int{3, 5, 2, 4, 3, 5, 5, 4, 4, 3, 4, 5},
        <a id="L57"></a>huffmanDecoder{
            <a id="L58"></a>2, 5,
            <a id="L59"></a>[maxCodeLen + 1]int{2: 0, 4, 13, 31},
            <a id="L60"></a>[maxCodeLen + 1]int{2: 0, 1, 6, 20},
            <a id="L61"></a><span class="comment">// Paper used different code assignment:</span>
            <a id="L62"></a><span class="comment">// 2, 9, 4, 0, 10, 8, 3, 7, 1, 5, 11, 6</span>
            <a id="L63"></a><span class="comment">// Reordered here so that codes of same length</span>
            <a id="L64"></a><span class="comment">// are assigned to increasing numbers.</span>
            <a id="L65"></a>[]int{2, 0, 4, 9, 3, 7, 8, 10, 1, 5, 6, 11},
        <a id="L66"></a>},
        <a id="L67"></a>true,
    <a id="L68"></a>},

    <a id="L70"></a><span class="comment">// Example from RFC 1951 section 3.2.2</span>
    <a id="L71"></a>&amp;InitDecoderTest{
        <a id="L72"></a>[]int{2, 1, 3, 3},
        <a id="L73"></a>huffmanDecoder{
            <a id="L74"></a>1, 3,
            <a id="L75"></a>[maxCodeLen + 1]int{1: 0, 2, 7},
            <a id="L76"></a>[maxCodeLen + 1]int{1: 0, 1, 4},
            <a id="L77"></a>[]int{1, 0, 2, 3},
        <a id="L78"></a>},
        <a id="L79"></a>true,
    <a id="L80"></a>},

    <a id="L82"></a><span class="comment">// Second example from RFC 1951 section 3.2.2</span>
    <a id="L83"></a>&amp;InitDecoderTest{
        <a id="L84"></a>[]int{3, 3, 3, 3, 3, 2, 4, 4},
        <a id="L85"></a>huffmanDecoder{
            <a id="L86"></a>2, 4,
            <a id="L87"></a>[maxCodeLen + 1]int{2: 0, 6, 15},
            <a id="L88"></a>[maxCodeLen + 1]int{2: 0, 1, 8},
            <a id="L89"></a>[]int{5, 0, 1, 2, 3, 4, 6, 7},
        <a id="L90"></a>},
        <a id="L91"></a>true,
    <a id="L92"></a>},

    <a id="L94"></a><span class="comment">// Static Huffman codes (RFC 1951 section 3.2.6)</span>
    <a id="L95"></a>&amp;InitDecoderTest{
        <a id="L96"></a>&amp;fixedHuffmanBits,
        <a id="L97"></a>fixedHuffmanDecoder,
        <a id="L98"></a>true,
    <a id="L99"></a>},

    <a id="L101"></a><span class="comment">// Illegal input.</span>
    <a id="L102"></a>&amp;InitDecoderTest{
        <a id="L103"></a>[]int{},
        <a id="L104"></a>huffmanDecoder{},
        <a id="L105"></a>false,
    <a id="L106"></a>},

    <a id="L108"></a><span class="comment">// Illegal input.</span>
    <a id="L109"></a>&amp;InitDecoderTest{
        <a id="L110"></a>[]int{0, 0, 0, 0, 0, 0, 0},
        <a id="L111"></a>huffmanDecoder{},
        <a id="L112"></a>false,
    <a id="L113"></a>},
<a id="L114"></a>}

<a id="L116"></a>func TestInitDecoder(t *testing.T) {
    <a id="L117"></a>for i, tt := range initDecoderTests {
        <a id="L118"></a>var h huffmanDecoder;
        <a id="L119"></a>if h.init(tt.in) != tt.ok {
            <a id="L120"></a>t.Errorf(&#34;test %d: init = %v&#34;, i, !tt.ok);
            <a id="L121"></a>continue;
        <a id="L122"></a>}
        <a id="L123"></a>if !reflect.DeepEqual(&amp;h, &amp;tt.out) {
            <a id="L124"></a>t.Errorf(&#34;test %d:\nhave %v\nwant %v&#34;, i, h, tt.out)
        <a id="L125"></a>}
    <a id="L126"></a>}
<a id="L127"></a>}

<a id="L129"></a>func TestUncompressedSource(t *testing.T) {
    <a id="L130"></a>decoder := NewInflater(bytes.NewBuffer([]byte{0x01, 0x01, 0x00, 0xfe, 0xff, 0x11}));
    <a id="L131"></a>output := make([]byte, 1);
    <a id="L132"></a>n, error := decoder.Read(output);
    <a id="L133"></a>if n != 1 || error != nil {
        <a id="L134"></a>t.Fatalf(&#34;decoder.Read() = %d, %v, want 1, nil&#34;, n, error)
    <a id="L135"></a>}
    <a id="L136"></a>if output[0] != 0x11 {
        <a id="L137"></a>t.Errorf(&#34;output[0] = %x, want 0x11&#34;, output[0])
    <a id="L138"></a>}
<a id="L139"></a>}
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
