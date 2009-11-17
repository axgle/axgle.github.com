<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/compress/flate/reverse_bits.go</title>

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
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/compress/flate/reverse_bits.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package flate

<a id="L7"></a>var reverseByte = [256]byte{
    <a id="L8"></a>0x00, 0x80, 0x40, 0xc0, 0x20, 0xa0, 0x60, 0xe0,
    <a id="L9"></a>0x10, 0x90, 0x50, 0xd0, 0x30, 0xb0, 0x70, 0xf0,
    <a id="L10"></a>0x08, 0x88, 0x48, 0xc8, 0x28, 0xa8, 0x68, 0xe8,
    <a id="L11"></a>0x18, 0x98, 0x58, 0xd8, 0x38, 0xb8, 0x78, 0xf8,
    <a id="L12"></a>0x04, 0x84, 0x44, 0xc4, 0x24, 0xa4, 0x64, 0xe4,
    <a id="L13"></a>0x14, 0x94, 0x54, 0xd4, 0x34, 0xb4, 0x74, 0xf4,
    <a id="L14"></a>0x0c, 0x8c, 0x4c, 0xcc, 0x2c, 0xac, 0x6c, 0xec,
    <a id="L15"></a>0x1c, 0x9c, 0x5c, 0xdc, 0x3c, 0xbc, 0x7c, 0xfc,
    <a id="L16"></a>0x02, 0x82, 0x42, 0xc2, 0x22, 0xa2, 0x62, 0xe2,
    <a id="L17"></a>0x12, 0x92, 0x52, 0xd2, 0x32, 0xb2, 0x72, 0xf2,
    <a id="L18"></a>0x0a, 0x8a, 0x4a, 0xca, 0x2a, 0xaa, 0x6a, 0xea,
    <a id="L19"></a>0x1a, 0x9a, 0x5a, 0xda, 0x3a, 0xba, 0x7a, 0xfa,
    <a id="L20"></a>0x06, 0x86, 0x46, 0xc6, 0x26, 0xa6, 0x66, 0xe6,
    <a id="L21"></a>0x16, 0x96, 0x56, 0xd6, 0x36, 0xb6, 0x76, 0xf6,
    <a id="L22"></a>0x0e, 0x8e, 0x4e, 0xce, 0x2e, 0xae, 0x6e, 0xee,
    <a id="L23"></a>0x1e, 0x9e, 0x5e, 0xde, 0x3e, 0xbe, 0x7e, 0xfe,
    <a id="L24"></a>0x01, 0x81, 0x41, 0xc1, 0x21, 0xa1, 0x61, 0xe1,
    <a id="L25"></a>0x11, 0x91, 0x51, 0xd1, 0x31, 0xb1, 0x71, 0xf1,
    <a id="L26"></a>0x09, 0x89, 0x49, 0xc9, 0x29, 0xa9, 0x69, 0xe9,
    <a id="L27"></a>0x19, 0x99, 0x59, 0xd9, 0x39, 0xb9, 0x79, 0xf9,
    <a id="L28"></a>0x05, 0x85, 0x45, 0xc5, 0x25, 0xa5, 0x65, 0xe5,
    <a id="L29"></a>0x15, 0x95, 0x55, 0xd5, 0x35, 0xb5, 0x75, 0xf5,
    <a id="L30"></a>0x0d, 0x8d, 0x4d, 0xcd, 0x2d, 0xad, 0x6d, 0xed,
    <a id="L31"></a>0x1d, 0x9d, 0x5d, 0xdd, 0x3d, 0xbd, 0x7d, 0xfd,
    <a id="L32"></a>0x03, 0x83, 0x43, 0xc3, 0x23, 0xa3, 0x63, 0xe3,
    <a id="L33"></a>0x13, 0x93, 0x53, 0xd3, 0x33, 0xb3, 0x73, 0xf3,
    <a id="L34"></a>0x0b, 0x8b, 0x4b, 0xcb, 0x2b, 0xab, 0x6b, 0xeb,
    <a id="L35"></a>0x1b, 0x9b, 0x5b, 0xdb, 0x3b, 0xbb, 0x7b, 0xfb,
    <a id="L36"></a>0x07, 0x87, 0x47, 0xc7, 0x27, 0xa7, 0x67, 0xe7,
    <a id="L37"></a>0x17, 0x97, 0x57, 0xd7, 0x37, 0xb7, 0x77, 0xf7,
    <a id="L38"></a>0x0f, 0x8f, 0x4f, 0xcf, 0x2f, 0xaf, 0x6f, 0xef,
    <a id="L39"></a>0x1f, 0x9f, 0x5f, 0xdf, 0x3f, 0xbf, 0x7f, 0xff,
<a id="L40"></a>}

<a id="L42"></a>func reverseUint16(v uint16) uint16 {
    <a id="L43"></a>return uint16(reverseByte[v&gt;&gt;8]) | uint16(reverseByte[v&amp;0xFF])&lt;&lt;8
<a id="L44"></a>}

<a id="L46"></a>func reverseBits(number uint16, bitLength byte) uint16 {
    <a id="L47"></a>return reverseUint16(number &lt;&lt; uint8(16-bitLength))
<a id="L48"></a>}
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
