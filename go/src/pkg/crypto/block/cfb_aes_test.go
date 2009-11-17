<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/cfb_aes_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/cfb_aes_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// CFB AES test vectors.</span>

<a id="L7"></a><span class="comment">// See U.S. National Institute of Standards and Technology (NIST)</span>
<a id="L8"></a><span class="comment">// Special Publication 800-38A, ``Recommendation for Block Cipher</span>
<a id="L9"></a><span class="comment">// Modes of Operation,&#39;&#39; 2001 Edition, pp. 29-52.</span>

<a id="L11"></a>package block

<a id="L13"></a>import (
    <a id="L14"></a>&#34;bytes&#34;;
    <a id="L15"></a>&#34;crypto/aes&#34;;
    <a id="L16"></a>&#34;io&#34;;
    <a id="L17"></a>&#34;testing&#34;;
<a id="L18"></a>)

<a id="L20"></a>type cfbTest struct {
    <a id="L21"></a>name string;
    <a id="L22"></a>s    int;
    <a id="L23"></a>key  []byte;
    <a id="L24"></a>iv   []byte;
    <a id="L25"></a>in   []byte;
    <a id="L26"></a>out  []byte;
<a id="L27"></a>}

<a id="L29"></a>var cfbAESTests = []cfbTest{
    <a id="L30"></a>cfbTest{
        <a id="L31"></a>&#34;CFB1-AES128&#34;,
        <a id="L32"></a>1,
        <a id="L33"></a>commonKey128,
        <a id="L34"></a>commonIV,
        <a id="L35"></a>[]byte{
            <a id="L36"></a>0&lt;&lt;7 | 1&lt;&lt;6 | 1&lt;&lt;5 | 0&lt;&lt;4 | 1&lt;&lt;3 | 0&lt;&lt;2 | 1&lt;&lt;1,
            <a id="L37"></a>1&lt;&lt;7 | 1&lt;&lt;6 | 0&lt;&lt;5 | 0&lt;&lt;4 | 0&lt;&lt;3 | 0&lt;&lt;2 | 0&lt;&lt;1,
        <a id="L38"></a>},
        <a id="L39"></a>[]byte{
            <a id="L40"></a>0&lt;&lt;7 | 1&lt;&lt;6 | 1&lt;&lt;5 | 0&lt;&lt;4 | 1&lt;&lt;3 | 0&lt;&lt;2 | 0&lt;&lt;1,
            <a id="L41"></a>1&lt;&lt;7 | 0&lt;&lt;6 | 1&lt;&lt;5 | 1&lt;&lt;4 | 0&lt;&lt;3 | 0&lt;&lt;2 | 1&lt;&lt;1,
        <a id="L42"></a>},
    <a id="L43"></a>},
    <a id="L44"></a>cfbTest{
        <a id="L45"></a>&#34;CFB1-AES192&#34;,
        <a id="L46"></a>1,
        <a id="L47"></a>commonKey192,
        <a id="L48"></a>commonIV,
        <a id="L49"></a>[]byte{
            <a id="L50"></a>0&lt;&lt;7 | 1&lt;&lt;6 | 1&lt;&lt;5 | 0&lt;&lt;4 | 1&lt;&lt;3 | 0&lt;&lt;2 | 1&lt;&lt;1,
            <a id="L51"></a>1&lt;&lt;7 | 1&lt;&lt;6 | 0&lt;&lt;5 | 0&lt;&lt;4 | 0&lt;&lt;3 | 0&lt;&lt;2 | 0&lt;&lt;1,
        <a id="L52"></a>},
        <a id="L53"></a>[]byte{
            <a id="L54"></a>1&lt;&lt;7 | 0&lt;&lt;6 | 0&lt;&lt;5 | 1&lt;&lt;4 | 0&lt;&lt;3 | 0&lt;&lt;2 | 1&lt;&lt;1,
            <a id="L55"></a>0&lt;&lt;7 | 1&lt;&lt;6 | 0&lt;&lt;5 | 1&lt;&lt;4 | 1&lt;&lt;3 | 0&lt;&lt;2 | 0&lt;&lt;1,
        <a id="L56"></a>},
    <a id="L57"></a>},
    <a id="L58"></a>cfbTest{
        <a id="L59"></a>&#34;CFB1-AES256&#34;,
        <a id="L60"></a>1,
        <a id="L61"></a>commonKey256,
        <a id="L62"></a>commonIV,
        <a id="L63"></a>[]byte{
            <a id="L64"></a>0&lt;&lt;7 | 1&lt;&lt;6 | 1&lt;&lt;5 | 0&lt;&lt;4 | 1&lt;&lt;3 | 0&lt;&lt;2 | 1&lt;&lt;1,
            <a id="L65"></a>1&lt;&lt;7 | 1&lt;&lt;6 | 0&lt;&lt;5 | 0&lt;&lt;4 | 0&lt;&lt;3 | 0&lt;&lt;2 | 0&lt;&lt;1,
        <a id="L66"></a>},
        <a id="L67"></a>[]byte{
            <a id="L68"></a>1&lt;&lt;7 | 0&lt;&lt;6 | 0&lt;&lt;5 | 1&lt;&lt;4 | 0&lt;&lt;3 | 0&lt;&lt;2 | 0&lt;&lt;1,
            <a id="L69"></a>0&lt;&lt;7 | 0&lt;&lt;6 | 1&lt;&lt;5 | 0&lt;&lt;4 | 1&lt;&lt;3 | 0&lt;&lt;2 | 0&lt;&lt;1,
        <a id="L70"></a>},
    <a id="L71"></a>},

    <a id="L73"></a>cfbTest{
        <a id="L74"></a>&#34;CFB8-AES128&#34;,
        <a id="L75"></a>8,
        <a id="L76"></a>commonKey128,
        <a id="L77"></a>commonIV,
        <a id="L78"></a>[]byte{
            <a id="L79"></a>0x6b,
            <a id="L80"></a>0xc1,
            <a id="L81"></a>0xbe,
            <a id="L82"></a>0xe2,
            <a id="L83"></a>0x2e,
            <a id="L84"></a>0x40,
            <a id="L85"></a>0x9f,
            <a id="L86"></a>0x96,
            <a id="L87"></a>0xe9,
            <a id="L88"></a>0x3d,
            <a id="L89"></a>0x7e,
            <a id="L90"></a>0x11,
            <a id="L91"></a>0x73,
            <a id="L92"></a>0x93,
            <a id="L93"></a>0x17,
            <a id="L94"></a>0x2a,
            <a id="L95"></a>0xae,
            <a id="L96"></a>0x2d,
        <a id="L97"></a>},
        <a id="L98"></a>[]byte{
            <a id="L99"></a>0x3b,
            <a id="L100"></a>0x79,
            <a id="L101"></a>0x42,
            <a id="L102"></a>0x4c,
            <a id="L103"></a>0x9c,
            <a id="L104"></a>0x0d,
            <a id="L105"></a>0xd4,
            <a id="L106"></a>0x36,
            <a id="L107"></a>0xba,
            <a id="L108"></a>0xce,
            <a id="L109"></a>0x9e,
            <a id="L110"></a>0x0e,
            <a id="L111"></a>0xd4,
            <a id="L112"></a>0x58,
            <a id="L113"></a>0x6a,
            <a id="L114"></a>0x4f,
            <a id="L115"></a>0x32,
            <a id="L116"></a>0xb9,
        <a id="L117"></a>},
    <a id="L118"></a>},

    <a id="L120"></a>cfbTest{
        <a id="L121"></a>&#34;CFB8-AES192&#34;,
        <a id="L122"></a>8,
        <a id="L123"></a>commonKey192,
        <a id="L124"></a>commonIV,
        <a id="L125"></a>[]byte{
            <a id="L126"></a>0x6b,
            <a id="L127"></a>0xc1,
            <a id="L128"></a>0xbe,
            <a id="L129"></a>0xe2,
            <a id="L130"></a>0x2e,
            <a id="L131"></a>0x40,
            <a id="L132"></a>0x9f,
            <a id="L133"></a>0x96,
            <a id="L134"></a>0xe9,
            <a id="L135"></a>0x3d,
            <a id="L136"></a>0x7e,
            <a id="L137"></a>0x11,
            <a id="L138"></a>0x73,
            <a id="L139"></a>0x93,
            <a id="L140"></a>0x17,
            <a id="L141"></a>0x2a,
            <a id="L142"></a>0xae,
            <a id="L143"></a>0x2d,
        <a id="L144"></a>},
        <a id="L145"></a>[]byte{
            <a id="L146"></a>0xcd,
            <a id="L147"></a>0xa2,
            <a id="L148"></a>0x52,
            <a id="L149"></a>0x1e,
            <a id="L150"></a>0xf0,
            <a id="L151"></a>0xa9,
            <a id="L152"></a>0x05,
            <a id="L153"></a>0xca,
            <a id="L154"></a>0x44,
            <a id="L155"></a>0xcd,
            <a id="L156"></a>0x05,
            <a id="L157"></a>0x7c,
            <a id="L158"></a>0xbf,
            <a id="L159"></a>0x0d,
            <a id="L160"></a>0x47,
            <a id="L161"></a>0xa0,
            <a id="L162"></a>0x67,
            <a id="L163"></a>0x8a,
        <a id="L164"></a>},
    <a id="L165"></a>},

    <a id="L167"></a>cfbTest{
        <a id="L168"></a>&#34;CFB8-AES256&#34;,
        <a id="L169"></a>8,
        <a id="L170"></a>commonKey256,
        <a id="L171"></a>commonIV,
        <a id="L172"></a>[]byte{
            <a id="L173"></a>0x6b,
            <a id="L174"></a>0xc1,
            <a id="L175"></a>0xbe,
            <a id="L176"></a>0xe2,
            <a id="L177"></a>0x2e,
            <a id="L178"></a>0x40,
            <a id="L179"></a>0x9f,
            <a id="L180"></a>0x96,
            <a id="L181"></a>0xe9,
            <a id="L182"></a>0x3d,
            <a id="L183"></a>0x7e,
            <a id="L184"></a>0x11,
            <a id="L185"></a>0x73,
            <a id="L186"></a>0x93,
            <a id="L187"></a>0x17,
            <a id="L188"></a>0x2a,
            <a id="L189"></a>0xae,
            <a id="L190"></a>0x2d,
        <a id="L191"></a>},
        <a id="L192"></a>[]byte{
            <a id="L193"></a>0xdc,
            <a id="L194"></a>0x1f,
            <a id="L195"></a>0x1a,
            <a id="L196"></a>0x85,
            <a id="L197"></a>0x20,
            <a id="L198"></a>0xa6,
            <a id="L199"></a>0x4d,
            <a id="L200"></a>0xb5,
            <a id="L201"></a>0x5f,
            <a id="L202"></a>0xcc,
            <a id="L203"></a>0x8a,
            <a id="L204"></a>0xc5,
            <a id="L205"></a>0x54,
            <a id="L206"></a>0x84,
            <a id="L207"></a>0x4e,
            <a id="L208"></a>0x88,
            <a id="L209"></a>0x97,
            <a id="L210"></a>0x00,
        <a id="L211"></a>},
    <a id="L212"></a>},

    <a id="L214"></a>cfbTest{
        <a id="L215"></a>&#34;CFB128-AES128&#34;,
        <a id="L216"></a>128,
        <a id="L217"></a>commonKey128,
        <a id="L218"></a>commonIV,
        <a id="L219"></a>[]byte{
            <a id="L220"></a>0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
            <a id="L221"></a>0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
            <a id="L222"></a>0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11, 0xe5, 0xfb, 0xc1, 0x19, 0x1a, 0x0a, 0x52, 0xef,
            <a id="L223"></a>0xf6, 0x9f, 0x24, 0x45, 0xdf, 0x4f, 0x9b, 0x17, 0xad, 0x2b, 0x41, 0x7b, 0xe6, 0x6c, 0x37, 0x10,
        <a id="L224"></a>},
        <a id="L225"></a>[]byte{
            <a id="L226"></a>0x3b, 0x3f, 0xd9, 0x2e, 0xb7, 0x2d, 0xad, 0x20, 0x33, 0x34, 0x49, 0xf8, 0xe8, 0x3c, 0xfb, 0x4a,
            <a id="L227"></a>0xc8, 0xa6, 0x45, 0x37, 0xa0, 0xb3, 0xa9, 0x3f, 0xcd, 0xe3, 0xcd, 0xad, 0x9f, 0x1c, 0xe5, 0x8b,
            <a id="L228"></a>0x26, 0x75, 0x1f, 0x67, 0xa3, 0xcb, 0xb1, 0x40, 0xb1, 0x80, 0x8c, 0xf1, 0x87, 0xa4, 0xf4, 0xdf,
            <a id="L229"></a>0xc0, 0x4b, 0x05, 0x35, 0x7c, 0x5d, 0x1c, 0x0e, 0xea, 0xc4, 0xc6, 0x6f, 0x9f, 0xf7, 0xf2, 0xe6,
        <a id="L230"></a>},
    <a id="L231"></a>},

    <a id="L233"></a>cfbTest{
        <a id="L234"></a>&#34;CFB128-AES192&#34;,
        <a id="L235"></a>128,
        <a id="L236"></a>commonKey192,
        <a id="L237"></a>commonIV,
        <a id="L238"></a>[]byte{
            <a id="L239"></a>0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
            <a id="L240"></a>0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
            <a id="L241"></a>0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11, 0xe5, 0xfb, 0xc1, 0x19, 0x1a, 0x0a, 0x52, 0xef,
            <a id="L242"></a>0xf6, 0x9f, 0x24, 0x45, 0xdf, 0x4f, 0x9b, 0x17, 0xad, 0x2b, 0x41, 0x7b, 0xe6, 0x6c, 0x37, 0x10,
        <a id="L243"></a>},
        <a id="L244"></a>[]byte{
            <a id="L245"></a>0xcd, 0xc8, 0x0d, 0x6f, 0xdd, 0xf1, 0x8c, 0xab, 0x34, 0xc2, 0x59, 0x09, 0xc9, 0x9a, 0x41, 0x74,
            <a id="L246"></a>0x67, 0xce, 0x7f, 0x7f, 0x81, 0x17, 0x36, 0x21, 0x96, 0x1a, 0x2b, 0x70, 0x17, 0x1d, 0x3d, 0x7a,
            <a id="L247"></a>0x2e, 0x1e, 0x8a, 0x1d, 0xd5, 0x9b, 0x88, 0xb1, 0xc8, 0xe6, 0x0f, 0xed, 0x1e, 0xfa, 0xc4, 0xc9,
            <a id="L248"></a>0xc0, 0x5f, 0x9f, 0x9c, 0xa9, 0x83, 0x4f, 0xa0, 0x42, 0xae, 0x8f, 0xba, 0x58, 0x4b, 0x09, 0xff,
        <a id="L249"></a>},
    <a id="L250"></a>},

    <a id="L252"></a>cfbTest{
        <a id="L253"></a>&#34;CFB128-AES256&#34;,
        <a id="L254"></a>128,
        <a id="L255"></a>commonKey256,
        <a id="L256"></a>commonIV,
        <a id="L257"></a>[]byte{
            <a id="L258"></a>0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
            <a id="L259"></a>0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
            <a id="L260"></a>0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11, 0xe5, 0xfb, 0xc1, 0x19, 0x1a, 0x0a, 0x52, 0xef,
            <a id="L261"></a>0xf6, 0x9f, 0x24, 0x45, 0xdf, 0x4f, 0x9b, 0x17, 0xad, 0x2b, 0x41, 0x7b, 0xe6, 0x6c, 0x37, 0x10,
        <a id="L262"></a>},
        <a id="L263"></a>[]byte{
            <a id="L264"></a>0xdc, 0x7e, 0x84, 0xbf, 0xda, 0x79, 0x16, 0x4b, 0x7e, 0xcd, 0x84, 0x86, 0x98, 0x5d, 0x38, 0x60,
            <a id="L265"></a>0x39, 0xff, 0xed, 0x14, 0x3b, 0x28, 0xb1, 0xc8, 0x32, 0x11, 0x3c, 0x63, 0x31, 0xe5, 0x40, 0x7b,
            <a id="L266"></a>0xdf, 0x10, 0x13, 0x24, 0x15, 0xe5, 0x4b, 0x92, 0xa1, 0x3e, 0xd0, 0xa8, 0x26, 0x7a, 0xe2, 0xf9,
            <a id="L267"></a>0x75, 0xa3, 0x85, 0x74, 0x1a, 0xb9, 0xce, 0xf8, 0x20, 0x31, 0x62, 0x3d, 0x55, 0xb1, 0xe4, 0x71,
        <a id="L268"></a>},
    <a id="L269"></a>},
<a id="L270"></a>}

<a id="L272"></a>func TestCFB_AES(t *testing.T) {
    <a id="L273"></a>for _, tt := range cfbAESTests {
        <a id="L274"></a>test := tt.name;

        <a id="L276"></a>if tt.s == 1 {
            <a id="L277"></a><span class="comment">// 1-bit CFB not implemented</span>
            <a id="L278"></a>continue
        <a id="L279"></a>}

        <a id="L281"></a>c, err := aes.NewCipher(tt.key);
        <a id="L282"></a>if err != nil {
            <a id="L283"></a>t.Errorf(&#34;%s: NewCipher(%d bytes) = %s&#34;, test, len(tt.key), err);
            <a id="L284"></a>continue;
        <a id="L285"></a>}

        <a id="L287"></a>var crypt bytes.Buffer;
        <a id="L288"></a>w := NewCFBEncrypter(c, tt.s, tt.iv, &amp;crypt);
        <a id="L289"></a>var r io.Reader = bytes.NewBuffer(tt.in);
        <a id="L290"></a>n, err := io.Copy(w, r);
        <a id="L291"></a>if n != int64(len(tt.in)) || err != nil {
            <a id="L292"></a>t.Errorf(&#34;%s: CFBEncrypter io.Copy = %d, %v want %d, nil&#34;, test, n, err, len(tt.in))
        <a id="L293"></a>} else if d := crypt.Bytes(); !same(tt.out, d) {
            <a id="L294"></a>t.Errorf(&#34;%s: CFBEncrypter\nhave %x\nwant %x&#34;, test, d, tt.out)
        <a id="L295"></a>}

        <a id="L297"></a>var plain bytes.Buffer;
        <a id="L298"></a>r = NewCFBDecrypter(c, tt.s, tt.iv, bytes.NewBuffer(tt.out));
        <a id="L299"></a>w = &amp;plain;
        <a id="L300"></a>n, err = io.Copy(w, r);
        <a id="L301"></a>if n != int64(len(tt.out)) || err != nil {
            <a id="L302"></a>t.Errorf(&#34;%s: CFBDecrypter io.Copy = %d, %v want %d, nil&#34;, test, n, err, len(tt.out))
        <a id="L303"></a>} else if d := plain.Bytes(); !same(tt.in, d) {
            <a id="L304"></a>t.Errorf(&#34;%s: CFBDecrypter\nhave %x\nwant %x&#34;, test, d, tt.in)
        <a id="L305"></a>}

        <a id="L307"></a>if t.Failed() {
            <a id="L308"></a>break
        <a id="L309"></a>}
    <a id="L310"></a>}
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
