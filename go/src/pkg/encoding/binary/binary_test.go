<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/encoding/binary/binary_test.go</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/encoding/binary/binary_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package binary

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;math&#34;;
    <a id="L10"></a>&#34;reflect&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)

<a id="L14"></a>type Struct struct {
    <a id="L15"></a>Int8    int8;
    <a id="L16"></a>Int16   int16;
    <a id="L17"></a>Int32   int32;
    <a id="L18"></a>Int64   int64;
    <a id="L19"></a>Uint8   uint8;
    <a id="L20"></a>Uint16  uint16;
    <a id="L21"></a>Uint32  uint32;
    <a id="L22"></a>Uint64  uint64;
    <a id="L23"></a>Float64 float64;
    <a id="L24"></a>Array   [4]uint8;
<a id="L25"></a>}

<a id="L27"></a>var s = Struct{
    <a id="L28"></a>0x01,
    <a id="L29"></a>0x0203,
    <a id="L30"></a>0x04050607,
    <a id="L31"></a>0x08090a0b0c0d0e0f,
    <a id="L32"></a>0x10,
    <a id="L33"></a>0x1112,
    <a id="L34"></a>0x13141516,
    <a id="L35"></a>0x1718191a1b1c1d1e,
    <a id="L36"></a>math.Float64frombits(0x1f20212223242526),
    <a id="L37"></a>[4]uint8{0x27, 0x28, 0x29, 0x2a},
<a id="L38"></a>}

<a id="L40"></a>var big = []byte{
    <a id="L41"></a>1,
    <a id="L42"></a>2, 3,
    <a id="L43"></a>4, 5, 6, 7,
    <a id="L44"></a>8, 9, 10, 11, 12, 13, 14, 15,
    <a id="L45"></a>16,
    <a id="L46"></a>17, 18,
    <a id="L47"></a>19, 20, 21, 22,
    <a id="L48"></a>23, 24, 25, 26, 27, 28, 29, 30,
    <a id="L49"></a>31, 32, 33, 34, 35, 36, 37, 38,
    <a id="L50"></a>39, 40, 41, 42,
<a id="L51"></a>}

<a id="L53"></a>var little = []byte{
    <a id="L54"></a>1,
    <a id="L55"></a>3, 2,
    <a id="L56"></a>7, 6, 5, 4,
    <a id="L57"></a>15, 14, 13, 12, 11, 10, 9, 8,
    <a id="L58"></a>16,
    <a id="L59"></a>18, 17,
    <a id="L60"></a>22, 21, 20, 19,
    <a id="L61"></a>30, 29, 28, 27, 26, 25, 24, 23,
    <a id="L62"></a>38, 37, 36, 35, 34, 33, 32, 31,
    <a id="L63"></a>39, 40, 41, 42,
<a id="L64"></a>}

<a id="L66"></a>func TestRead(t *testing.T) {
    <a id="L67"></a>var sl, sb Struct;

    <a id="L69"></a>err := Read(bytes.NewBuffer(big), BigEndian, &amp;sb);
    <a id="L70"></a>if err != nil {
        <a id="L71"></a>t.Errorf(&#34;Read big-endian: %v&#34;, err);
        <a id="L72"></a>goto little;
    <a id="L73"></a>}
    <a id="L74"></a>if !reflect.DeepEqual(sb, s) {
        <a id="L75"></a>t.Errorf(&#34;Read big-endian:\n\thave %+v\n\twant %+v&#34;, sb, s)
    <a id="L76"></a>}

<a id="L78"></a>little:
    <a id="L79"></a>err = Read(bytes.NewBuffer(little), LittleEndian, &amp;sl);
    <a id="L80"></a>if err != nil {
        <a id="L81"></a>t.Errorf(&#34;Read little-endian: %v&#34;, err)
    <a id="L82"></a>}
    <a id="L83"></a>if !reflect.DeepEqual(sl, s) {
        <a id="L84"></a>t.Errorf(&#34;Read big-endian:\n\thave %+v\n\twant %+v&#34;, sl, s)
    <a id="L85"></a>}
<a id="L86"></a>}
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
