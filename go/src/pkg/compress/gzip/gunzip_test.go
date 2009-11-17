<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/compress/gzip/gunzip_test.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/compress/gzip/gunzip_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package gzip

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;io&#34;;
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)

<a id="L14"></a>type gzipTest struct {
    <a id="L15"></a>name string;
    <a id="L16"></a>desc string;
    <a id="L17"></a>raw  string;
    <a id="L18"></a>gzip []byte;
    <a id="L19"></a>err  os.Error;
<a id="L20"></a>}

<a id="L22"></a>var gzipTests = []gzipTest{
    <a id="L23"></a>gzipTest{ <span class="comment">// has 1 empty fixed-huffman block</span>
        <a id="L24"></a>&#34;empty.txt&#34;,
        <a id="L25"></a>&#34;empty.txt&#34;,
        <a id="L26"></a>&#34;&#34;,
        <a id="L27"></a>[]byte{
            <a id="L28"></a>0x1f, 0x8b, 0x08, 0x08, 0xf7, 0x5e, 0x14, 0x4a,
            <a id="L29"></a>0x00, 0x03, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
            <a id="L30"></a>0x74, 0x78, 0x74, 0x00, 0x03, 0x00, 0x00, 0x00,
            <a id="L31"></a>0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        <a id="L32"></a>},
        <a id="L33"></a>nil,
    <a id="L34"></a>},
    <a id="L35"></a>gzipTest{ <span class="comment">// has 1 non-empty fixed huffman block</span>
        <a id="L36"></a>&#34;hello.txt&#34;,
        <a id="L37"></a>&#34;hello.txt&#34;,
        <a id="L38"></a>&#34;hello world\n&#34;,
        <a id="L39"></a>[]byte{
            <a id="L40"></a>0x1f, 0x8b, 0x08, 0x08, 0xc8, 0x58, 0x13, 0x4a,
            <a id="L41"></a>0x00, 0x03, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
            <a id="L42"></a>0x74, 0x78, 0x74, 0x00, 0xcb, 0x48, 0xcd, 0xc9,
            <a id="L43"></a>0xc9, 0x57, 0x28, 0xcf, 0x2f, 0xca, 0x49, 0xe1,
            <a id="L44"></a>0x02, 0x00, 0x2d, 0x3b, 0x08, 0xaf, 0x0c, 0x00,
            <a id="L45"></a>0x00, 0x00,
        <a id="L46"></a>},
        <a id="L47"></a>nil,
    <a id="L48"></a>},
    <a id="L49"></a>gzipTest{ <span class="comment">// concatenation</span>
        <a id="L50"></a>&#34;hello.txt&#34;,
        <a id="L51"></a>&#34;hello.txt x2&#34;,
        <a id="L52"></a>&#34;hello world\n&#34;
            <a id="L53"></a>&#34;hello world\n&#34;,
        <a id="L54"></a>[]byte{
            <a id="L55"></a>0x1f, 0x8b, 0x08, 0x08, 0xc8, 0x58, 0x13, 0x4a,
            <a id="L56"></a>0x00, 0x03, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
            <a id="L57"></a>0x74, 0x78, 0x74, 0x00, 0xcb, 0x48, 0xcd, 0xc9,
            <a id="L58"></a>0xc9, 0x57, 0x28, 0xcf, 0x2f, 0xca, 0x49, 0xe1,
            <a id="L59"></a>0x02, 0x00, 0x2d, 0x3b, 0x08, 0xaf, 0x0c, 0x00,
            <a id="L60"></a>0x00, 0x00,
            <a id="L61"></a>0x1f, 0x8b, 0x08, 0x08, 0xc8, 0x58, 0x13, 0x4a,
            <a id="L62"></a>0x00, 0x03, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
            <a id="L63"></a>0x74, 0x78, 0x74, 0x00, 0xcb, 0x48, 0xcd, 0xc9,
            <a id="L64"></a>0xc9, 0x57, 0x28, 0xcf, 0x2f, 0xca, 0x49, 0xe1,
            <a id="L65"></a>0x02, 0x00, 0x2d, 0x3b, 0x08, 0xaf, 0x0c, 0x00,
            <a id="L66"></a>0x00, 0x00,
        <a id="L67"></a>},
        <a id="L68"></a>nil,
    <a id="L69"></a>},
    <a id="L70"></a>gzipTest{ <span class="comment">// has a fixed huffman block with some length-distance pairs</span>
        <a id="L71"></a>&#34;shesells.txt&#34;,
        <a id="L72"></a>&#34;shesells.txt&#34;,
        <a id="L73"></a>&#34;she sells seashells by the seashore\n&#34;,
        <a id="L74"></a>[]byte{
            <a id="L75"></a>0x1f, 0x8b, 0x08, 0x08, 0x72, 0x66, 0x8b, 0x4a,
            <a id="L76"></a>0x00, 0x03, 0x73, 0x68, 0x65, 0x73, 0x65, 0x6c,
            <a id="L77"></a>0x6c, 0x73, 0x2e, 0x74, 0x78, 0x74, 0x00, 0x2b,
            <a id="L78"></a>0xce, 0x48, 0x55, 0x28, 0x4e, 0xcd, 0xc9, 0x29,
            <a id="L79"></a>0x06, 0x92, 0x89, 0xc5, 0x19, 0x60, 0x56, 0x52,
            <a id="L80"></a>0xa5, 0x42, 0x09, 0x58, 0x18, 0x28, 0x90, 0x5f,
            <a id="L81"></a>0x94, 0xca, 0x05, 0x00, 0x76, 0xb0, 0x3b, 0xeb,
            <a id="L82"></a>0x24, 0x00, 0x00, 0x00,
        <a id="L83"></a>},
        <a id="L84"></a>nil,
    <a id="L85"></a>},
    <a id="L86"></a>gzipTest{ <span class="comment">// has dynamic huffman blocks</span>
        <a id="L87"></a>&#34;gettysburg&#34;,
        <a id="L88"></a>&#34;gettysburg&#34;,
        <a id="L89"></a>&#34;  Four score and seven years ago our fathers brought forth on\n&#34;
            <a id="L90"></a>&#34;this continent, a new nation, conceived in Liberty, and dedicated\n&#34;
            <a id="L91"></a>&#34;to the proposition that all men are created equal.\n&#34;
            <a id="L92"></a>&#34;  Now we are engaged in a great Civil War, testing whether that\n&#34;
            <a id="L93"></a>&#34;nation, or any nation so conceived and so dedicated, can long\n&#34;
            <a id="L94"></a>&#34;endure.\n&#34;
            <a id="L95"></a>&#34;  We are met on a great battle-field of that war.\n&#34;
            <a id="L96"></a>&#34;  We have come to dedicate a portion of that field, as a final\n&#34;
            <a id="L97"></a>&#34;resting place for those who here gave their lives that that\n&#34;
            <a id="L98"></a>&#34;nation might live.  It is altogether fitting and proper that\n&#34;
            <a id="L99"></a>&#34;we should do this.\n&#34;
            <a id="L100"></a>&#34;  But, in a larger sense, we can not dedicate — we can not\n&#34;
            <a id="L101"></a>&#34;consecrate — we can not hallow — this ground.\n&#34;
            <a id="L102"></a>&#34;  The brave men, living and dead, who struggled here, have\n&#34;
            <a id="L103"></a>&#34;consecrated it, far above our poor power to add or detract.\n&#34;
            <a id="L104"></a>&#34;The world will little note, nor long remember what we say here,\n&#34;
            <a id="L105"></a>&#34;but it can never forget what they did here.\n&#34;
            <a id="L106"></a>&#34;  It is for us the living, rather, to be dedicated here to the\n&#34;
            <a id="L107"></a>&#34;unfinished work which they who fought here have thus far so\n&#34;
            <a id="L108"></a>&#34;nobly advanced.  It is rather for us to be here dedicated to\n&#34;
            <a id="L109"></a>&#34;the great task remaining before us — that from these honored\n&#34;
            <a id="L110"></a>&#34;dead we take increased devotion to that cause for which they\n&#34;
            <a id="L111"></a>&#34;gave the last full measure of devotion —\n&#34;
            <a id="L112"></a>&#34;  that we here highly resolve that these dead shall not have\n&#34;
            <a id="L113"></a>&#34;died in vain — that this nation, under God, shall have a new\n&#34;
            <a id="L114"></a>&#34;birth of freedom — and that government of the people, by the\n&#34;
            <a id="L115"></a>&#34;people, for the people, shall not perish from this earth.\n&#34;
            <a id="L116"></a>&#34;\n&#34;
            <a id="L117"></a>&#34;Abraham Lincoln, November 19, 1863, Gettysburg, Pennsylvania\n&#34;,
        <a id="L118"></a>[]byte{
            <a id="L119"></a>0x1f, 0x8b, 0x08, 0x08, 0xd1, 0x12, 0x2b, 0x4a,
            <a id="L120"></a>0x00, 0x03, 0x67, 0x65, 0x74, 0x74, 0x79, 0x73,
            <a id="L121"></a>0x62, 0x75, 0x72, 0x67, 0x00, 0x65, 0x54, 0xcd,
            <a id="L122"></a>0x6e, 0xd4, 0x30, 0x10, 0xbe, 0xfb, 0x29, 0xe6,
            <a id="L123"></a>0x01, 0x42, 0xa5, 0x0a, 0x09, 0xc1, 0x11, 0x90,
            <a id="L124"></a>0x40, 0x48, 0xa8, 0xe2, 0x80, 0xd4, 0xf3, 0x24,
            <a id="L125"></a>0x9e, 0x24, 0x56, 0xbd, 0x9e, 0xc5, 0x76, 0x76,
            <a id="L126"></a>0x95, 0x1b, 0x0f, 0xc1, 0x13, 0xf2, 0x24, 0x7c,
            <a id="L127"></a>0x63, 0x77, 0x9b, 0x4a, 0x5c, 0xaa, 0x6e, 0x6c,
            <a id="L128"></a>0xcf, 0x7c, 0x7f, 0x33, 0x44, 0x5f, 0x74, 0xcb,
            <a id="L129"></a>0x54, 0x26, 0xcd, 0x42, 0x9c, 0x3c, 0x15, 0xb9,
            <a id="L130"></a>0x48, 0xa2, 0x5d, 0x38, 0x17, 0xe2, 0x45, 0xc9,
            <a id="L131"></a>0x4e, 0x67, 0xae, 0xab, 0xe0, 0xf7, 0x98, 0x75,
            <a id="L132"></a>0x5b, 0xd6, 0x4a, 0xb3, 0xe6, 0xba, 0x92, 0x26,
            <a id="L133"></a>0x57, 0xd7, 0x50, 0x68, 0xd2, 0x54, 0x43, 0x92,
            <a id="L134"></a>0x54, 0x07, 0x62, 0x4a, 0x72, 0xa5, 0xc4, 0x35,
            <a id="L135"></a>0x68, 0x1a, 0xec, 0x60, 0x92, 0x70, 0x11, 0x4f,
            <a id="L136"></a>0x21, 0xd1, 0xf7, 0x30, 0x4a, 0xae, 0xfb, 0xd0,
            <a id="L137"></a>0x9a, 0x78, 0xf1, 0x61, 0xe2, 0x2a, 0xde, 0x55,
            <a id="L138"></a>0x25, 0xd4, 0xa6, 0x73, 0xd6, 0xb3, 0x96, 0x60,
            <a id="L139"></a>0xef, 0xf0, 0x9b, 0x2b, 0x71, 0x8c, 0x74, 0x02,
            <a id="L140"></a>0x10, 0x06, 0xac, 0x29, 0x8b, 0xdd, 0x25, 0xf9,
            <a id="L141"></a>0xb5, 0x71, 0xbc, 0x73, 0x44, 0x0f, 0x7a, 0xa5,
            <a id="L142"></a>0xab, 0xb4, 0x33, 0x49, 0x0b, 0x2f, 0xbd, 0x03,
            <a id="L143"></a>0xd3, 0x62, 0x17, 0xe9, 0x73, 0xb8, 0x84, 0x48,
            <a id="L144"></a>0x8f, 0x9c, 0x07, 0xaa, 0x52, 0x00, 0x6d, 0xa1,
            <a id="L145"></a>0xeb, 0x2a, 0xc6, 0xa0, 0x95, 0x76, 0x37, 0x78,
            <a id="L146"></a>0x9a, 0x81, 0x65, 0x7f, 0x46, 0x4b, 0x45, 0x5f,
            <a id="L147"></a>0xe1, 0x6d, 0x42, 0xe8, 0x01, 0x13, 0x5c, 0x38,
            <a id="L148"></a>0x51, 0xd4, 0xb4, 0x38, 0x49, 0x7e, 0xcb, 0x62,
            <a id="L149"></a>0x28, 0x1e, 0x3b, 0x82, 0x93, 0x54, 0x48, 0xf1,
            <a id="L150"></a>0xd2, 0x7d, 0xe4, 0x5a, 0xa3, 0xbc, 0x99, 0x83,
            <a id="L151"></a>0x44, 0x4f, 0x3a, 0x77, 0x36, 0x57, 0xce, 0xcf,
            <a id="L152"></a>0x2f, 0x56, 0xbe, 0x80, 0x90, 0x9e, 0x84, 0xea,
            <a id="L153"></a>0x51, 0x1f, 0x8f, 0xcf, 0x90, 0xd4, 0x60, 0xdc,
            <a id="L154"></a>0x5e, 0xb4, 0xf7, 0x10, 0x0b, 0x26, 0xe0, 0xff,
            <a id="L155"></a>0xc4, 0xd1, 0xe5, 0x67, 0x2e, 0xe7, 0xc8, 0x93,
            <a id="L156"></a>0x98, 0x05, 0xb8, 0xa8, 0x45, 0xc0, 0x4d, 0x09,
            <a id="L157"></a>0xdc, 0x84, 0x16, 0x2b, 0x0d, 0x9a, 0x21, 0x53,
            <a id="L158"></a>0x04, 0x8b, 0xd2, 0x0b, 0xbd, 0xa2, 0x4c, 0xa7,
            <a id="L159"></a>0x60, 0xee, 0xd9, 0xe1, 0x1d, 0xd1, 0xb7, 0x4a,
            <a id="L160"></a>0x30, 0x8f, 0x63, 0xd5, 0xa5, 0x8b, 0x33, 0x87,
            <a id="L161"></a>0xda, 0x1a, 0x18, 0x79, 0xf3, 0xe3, 0xa6, 0x17,
            <a id="L162"></a>0x94, 0x2e, 0xab, 0x6e, 0xa0, 0xe3, 0xcd, 0xac,
            <a id="L163"></a>0x50, 0x8c, 0xca, 0xa7, 0x0d, 0x76, 0x37, 0xd1,
            <a id="L164"></a>0x23, 0xe7, 0x05, 0x57, 0x8b, 0xa4, 0x22, 0x83,
            <a id="L165"></a>0xd9, 0x62, 0x52, 0x25, 0xad, 0x07, 0xbb, 0xbf,
            <a id="L166"></a>0xbf, 0xff, 0xbc, 0xfa, 0xee, 0x20, 0x73, 0x91,
            <a id="L167"></a>0x29, 0xff, 0x7f, 0x02, 0x71, 0x62, 0x84, 0xb5,
            <a id="L168"></a>0xf6, 0xb5, 0x25, 0x6b, 0x41, 0xde, 0x92, 0xb7,
            <a id="L169"></a>0x76, 0x3f, 0x91, 0x91, 0x31, 0x1b, 0x41, 0x84,
            <a id="L170"></a>0x62, 0x30, 0x0a, 0x37, 0xa4, 0x5e, 0x18, 0x3a,
            <a id="L171"></a>0x99, 0x08, 0xa5, 0xe6, 0x6d, 0x59, 0x22, 0xec,
            <a id="L172"></a>0x33, 0x39, 0x86, 0x26, 0xf5, 0xab, 0x66, 0xc8,
            <a id="L173"></a>0x08, 0x20, 0xcf, 0x0c, 0xd7, 0x47, 0x45, 0x21,
            <a id="L174"></a>0x0b, 0xf6, 0x59, 0xd5, 0xfe, 0x5c, 0x8d, 0xaa,
            <a id="L175"></a>0x12, 0x7b, 0x6f, 0xa1, 0xf0, 0x52, 0x33, 0x4f,
            <a id="L176"></a>0xf5, 0xce, 0x59, 0xd3, 0xab, 0x66, 0x10, 0xbf,
            <a id="L177"></a>0x06, 0xc4, 0x31, 0x06, 0x73, 0xd6, 0x80, 0xa2,
            <a id="L178"></a>0x78, 0xc2, 0x45, 0xcb, 0x03, 0x65, 0x39, 0xc9,
            <a id="L179"></a>0x09, 0xd1, 0x06, 0x04, 0x33, 0x1a, 0x5a, 0xf1,
            <a id="L180"></a>0xde, 0x01, 0xb8, 0x71, 0x83, 0xc4, 0xb5, 0xb3,
            <a id="L181"></a>0xc3, 0x54, 0x65, 0x33, 0x0d, 0x5a, 0xf7, 0x9b,
            <a id="L182"></a>0x90, 0x7c, 0x27, 0x1f, 0x3a, 0x58, 0xa3, 0xd8,
            <a id="L183"></a>0xfd, 0x30, 0x5f, 0xb7, 0xd2, 0x66, 0xa2, 0x93,
            <a id="L184"></a>0x1c, 0x28, 0xb7, 0xe9, 0x1b, 0x0c, 0xe1, 0x28,
            <a id="L185"></a>0x47, 0x26, 0xbb, 0xe9, 0x7d, 0x7e, 0xdc, 0x96,
            <a id="L186"></a>0x10, 0x92, 0x50, 0x56, 0x7c, 0x06, 0xe2, 0x27,
            <a id="L187"></a>0xb4, 0x08, 0xd3, 0xda, 0x7b, 0x98, 0x34, 0x73,
            <a id="L188"></a>0x9f, 0xdb, 0xf6, 0x62, 0xed, 0x31, 0x41, 0x13,
            <a id="L189"></a>0xd3, 0xa2, 0xa8, 0x4b, 0x3a, 0xc6, 0x1d, 0xe4,
            <a id="L190"></a>0x2f, 0x8c, 0xf8, 0xfb, 0x97, 0x64, 0xf4, 0xb6,
            <a id="L191"></a>0x2f, 0x80, 0x5a, 0xf3, 0x56, 0xe0, 0x40, 0x50,
            <a id="L192"></a>0xd5, 0x19, 0xd0, 0x1e, 0xfc, 0xca, 0xe5, 0xc9,
            <a id="L193"></a>0xd4, 0x60, 0x00, 0x81, 0x2e, 0xa3, 0xcc, 0xb6,
            <a id="L194"></a>0x52, 0xf0, 0xb4, 0xdb, 0x69, 0x99, 0xce, 0x7a,
            <a id="L195"></a>0x32, 0x4c, 0x08, 0xed, 0xaa, 0x10, 0x10, 0xe3,
            <a id="L196"></a>0x6f, 0xee, 0x99, 0x68, 0x95, 0x9f, 0x04, 0x71,
            <a id="L197"></a>0xb2, 0x49, 0x2f, 0x62, 0xa6, 0x5e, 0xb4, 0xef,
            <a id="L198"></a>0x02, 0xed, 0x4f, 0x27, 0xde, 0x4a, 0x0f, 0xfd,
            <a id="L199"></a>0xc1, 0xcc, 0xdd, 0x02, 0x8f, 0x08, 0x16, 0x54,
            <a id="L200"></a>0xdf, 0xda, 0xca, 0xe0, 0x82, 0xf1, 0xb4, 0x31,
            <a id="L201"></a>0x7a, 0xa9, 0x81, 0xfe, 0x90, 0xb7, 0x3e, 0xdb,
            <a id="L202"></a>0xd3, 0x35, 0xc0, 0x20, 0x80, 0x33, 0x46, 0x4a,
            <a id="L203"></a>0x63, 0xab, 0xd1, 0x0d, 0x29, 0xd2, 0xe2, 0x84,
            <a id="L204"></a>0xb8, 0xdb, 0xfa, 0xe9, 0x89, 0x44, 0x86, 0x7c,
            <a id="L205"></a>0xe8, 0x0b, 0xe6, 0x02, 0x6a, 0x07, 0x9b, 0x96,
            <a id="L206"></a>0xd0, 0xdb, 0x2e, 0x41, 0x4c, 0xa1, 0xd5, 0x57,
            <a id="L207"></a>0x45, 0x14, 0xfb, 0xe3, 0xa6, 0x72, 0x5b, 0x87,
            <a id="L208"></a>0x6e, 0x0c, 0x6d, 0x5b, 0xce, 0xe0, 0x2f, 0xe2,
            <a id="L209"></a>0x21, 0x81, 0x95, 0xb0, 0xe8, 0xb6, 0x32, 0x0b,
            <a id="L210"></a>0xb2, 0x98, 0x13, 0x52, 0x5d, 0xfb, 0xec, 0x63,
            <a id="L211"></a>0x17, 0x8a, 0x9e, 0x23, 0x22, 0x36, 0xee, 0xcd,
            <a id="L212"></a>0xda, 0xdb, 0xcf, 0x3e, 0xf1, 0xc7, 0xf1, 0x01,
            <a id="L213"></a>0x12, 0x93, 0x0a, 0xeb, 0x6f, 0xf2, 0x02, 0x15,
            <a id="L214"></a>0x96, 0x77, 0x5d, 0xef, 0x9c, 0xfb, 0x88, 0x91,
            <a id="L215"></a>0x59, 0xf9, 0x84, 0xdd, 0x9b, 0x26, 0x8d, 0x80,
            <a id="L216"></a>0xf9, 0x80, 0x66, 0x2d, 0xac, 0xf7, 0x1f, 0x06,
            <a id="L217"></a>0xba, 0x7f, 0xff, 0xee, 0xed, 0x40, 0x5f, 0xa5,
            <a id="L218"></a>0xd6, 0xbd, 0x8c, 0x5b, 0x46, 0xd2, 0x7e, 0x48,
            <a id="L219"></a>0x4a, 0x65, 0x8f, 0x08, 0x42, 0x60, 0xf7, 0x0f,
            <a id="L220"></a>0xb9, 0x16, 0x0b, 0x0c, 0x1a, 0x06, 0x00, 0x00,
        <a id="L221"></a>},
        <a id="L222"></a>nil,
    <a id="L223"></a>},
    <a id="L224"></a>gzipTest{ <span class="comment">// has 1 non-empty fixed huffman block then garbage</span>
        <a id="L225"></a>&#34;hello.txt&#34;,
        <a id="L226"></a>&#34;hello.txt + garbage&#34;,
        <a id="L227"></a>&#34;hello world\n&#34;,
        <a id="L228"></a>[]byte{
            <a id="L229"></a>0x1f, 0x8b, 0x08, 0x08, 0xc8, 0x58, 0x13, 0x4a,
            <a id="L230"></a>0x00, 0x03, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
            <a id="L231"></a>0x74, 0x78, 0x74, 0x00, 0xcb, 0x48, 0xcd, 0xc9,
            <a id="L232"></a>0xc9, 0x57, 0x28, 0xcf, 0x2f, 0xca, 0x49, 0xe1,
            <a id="L233"></a>0x02, 0x00, 0x2d, 0x3b, 0x08, 0xaf, 0x0c, 0x00,
            <a id="L234"></a>0x00, 0x00, &#39;g&#39;, &#39;a&#39;, &#39;r&#39;, &#39;b&#39;, &#39;a&#39;, &#39;g&#39;, &#39;e&#39;, &#39;!&#39;, &#39;!&#39;, &#39;!&#39;,
        <a id="L235"></a>},
        <a id="L236"></a>HeaderError,
    <a id="L237"></a>},
    <a id="L238"></a>gzipTest{ <span class="comment">// has 1 non-empty fixed huffman block not enough header</span>
        <a id="L239"></a>&#34;hello.txt&#34;,
        <a id="L240"></a>&#34;hello.txt + garbage&#34;,
        <a id="L241"></a>&#34;hello world\n&#34;,
        <a id="L242"></a>[]byte{
            <a id="L243"></a>0x1f, 0x8b, 0x08, 0x08, 0xc8, 0x58, 0x13, 0x4a,
            <a id="L244"></a>0x00, 0x03, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
            <a id="L245"></a>0x74, 0x78, 0x74, 0x00, 0xcb, 0x48, 0xcd, 0xc9,
            <a id="L246"></a>0xc9, 0x57, 0x28, 0xcf, 0x2f, 0xca, 0x49, 0xe1,
            <a id="L247"></a>0x02, 0x00, 0x2d, 0x3b, 0x08, 0xaf, 0x0c, 0x00,
            <a id="L248"></a>0x00, 0x00, gzipID1,
        <a id="L249"></a>},
        <a id="L250"></a>io.ErrUnexpectedEOF,
    <a id="L251"></a>},
    <a id="L252"></a>gzipTest{ <span class="comment">// has 1 non-empty fixed huffman block but corrupt checksum</span>
        <a id="L253"></a>&#34;hello.txt&#34;,
        <a id="L254"></a>&#34;hello.txt + corrupt checksum&#34;,
        <a id="L255"></a>&#34;hello world\n&#34;,
        <a id="L256"></a>[]byte{
            <a id="L257"></a>0x1f, 0x8b, 0x08, 0x08, 0xc8, 0x58, 0x13, 0x4a,
            <a id="L258"></a>0x00, 0x03, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
            <a id="L259"></a>0x74, 0x78, 0x74, 0x00, 0xcb, 0x48, 0xcd, 0xc9,
            <a id="L260"></a>0xc9, 0x57, 0x28, 0xcf, 0x2f, 0xca, 0x49, 0xe1,
            <a id="L261"></a>0x02, 0x00, 0xff, 0xff, 0xff, 0xff, 0x0c, 0x00,
            <a id="L262"></a>0x00, 0x00,
        <a id="L263"></a>},
        <a id="L264"></a>ChecksumError,
    <a id="L265"></a>},
    <a id="L266"></a>gzipTest{ <span class="comment">// has 1 non-empty fixed huffman block but corrupt size</span>
        <a id="L267"></a>&#34;hello.txt&#34;,
        <a id="L268"></a>&#34;hello.txt + corrupt size&#34;,
        <a id="L269"></a>&#34;hello world\n&#34;,
        <a id="L270"></a>[]byte{
            <a id="L271"></a>0x1f, 0x8b, 0x08, 0x08, 0xc8, 0x58, 0x13, 0x4a,
            <a id="L272"></a>0x00, 0x03, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
            <a id="L273"></a>0x74, 0x78, 0x74, 0x00, 0xcb, 0x48, 0xcd, 0xc9,
            <a id="L274"></a>0xc9, 0x57, 0x28, 0xcf, 0x2f, 0xca, 0x49, 0xe1,
            <a id="L275"></a>0x02, 0x00, 0x2d, 0x3b, 0x08, 0xaf, 0xff, 0x00,
            <a id="L276"></a>0x00, 0x00,
        <a id="L277"></a>},
        <a id="L278"></a>ChecksumError,
    <a id="L279"></a>},
<a id="L280"></a>}

<a id="L282"></a>func TestInflater(t *testing.T) {
    <a id="L283"></a>b := new(bytes.Buffer);
    <a id="L284"></a>for _, tt := range gzipTests {
        <a id="L285"></a>in := bytes.NewBuffer(tt.gzip);
        <a id="L286"></a>gzip, err := NewInflater(in);
        <a id="L287"></a>if err != nil {
            <a id="L288"></a>t.Errorf(&#34;%s: NewInflater: %s&#34;, tt.name, err);
            <a id="L289"></a>continue;
        <a id="L290"></a>}
        <a id="L291"></a>defer gzip.Close();
        <a id="L292"></a>if tt.name != gzip.Name {
            <a id="L293"></a>t.Errorf(&#34;%s: got name %s&#34;, tt.name, gzip.Name)
        <a id="L294"></a>}
        <a id="L295"></a>b.Reset();
        <a id="L296"></a>n, err := io.Copy(b, gzip);
        <a id="L297"></a>if err != tt.err {
            <a id="L298"></a>t.Errorf(&#34;%s: io.Copy: %v want %v&#34;, tt.name, err, tt.err)
        <a id="L299"></a>}
        <a id="L300"></a>s := b.String();
        <a id="L301"></a>if s != tt.raw {
            <a id="L302"></a>t.Errorf(&#34;%s: got %d-byte %q want %d-byte %q&#34;, tt.name, n, s, len(tt.raw), tt.raw)
        <a id="L303"></a>}
    <a id="L304"></a>}
<a id="L305"></a>}
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
