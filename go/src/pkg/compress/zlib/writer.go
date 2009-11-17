<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/compress/zlib/writer.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/compress/zlib/writer.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package zlib

<a id="L7"></a>import (
    <a id="L8"></a>&#34;compress/flate&#34;;
    <a id="L9"></a>&#34;hash&#34;;
    <a id="L10"></a>&#34;hash/adler32&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;os&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">// These constants are copied from the flate package, so that code that imports</span>
<a id="L16"></a><span class="comment">// &#34;compress/zlib&#34; does not also have to import &#34;compress/flate&#34;.</span>
<a id="L17"></a>const (
    <a id="L18"></a>NoCompression      = flate.NoCompression;
    <a id="L19"></a>BestSpeed          = flate.BestSpeed;
    <a id="L20"></a>BestCompression    = flate.BestCompression;
    <a id="L21"></a>DefaultCompression = flate.DefaultCompression;
<a id="L22"></a>)

<a id="L24"></a>type writer struct {
    <a id="L25"></a>w        io.Writer;
    <a id="L26"></a>deflater io.WriteCloser;
    <a id="L27"></a>digest   hash.Hash32;
    <a id="L28"></a>err      os.Error;
    <a id="L29"></a>scratch  [4]byte;
<a id="L30"></a>}

<a id="L32"></a><span class="comment">// NewDeflater calls NewDeflaterLevel with the default compression level.</span>
<a id="L33"></a>func NewDeflater(w io.Writer) (io.WriteCloser, os.Error) {
    <a id="L34"></a>return NewDeflaterLevel(w, DefaultCompression)
<a id="L35"></a>}

<a id="L37"></a><span class="comment">// NewDeflater creates a new io.WriteCloser that satisfies writes by compressing data written to w.</span>
<a id="L38"></a><span class="comment">// It is the caller&#39;s responsibility to call Close on the WriteCloser when done.</span>
<a id="L39"></a><span class="comment">// level is the compression level, which can be DefaultCompression, NoCompression,</span>
<a id="L40"></a><span class="comment">// or any integer value between BestSpeed and BestCompression (inclusive).</span>
<a id="L41"></a>func NewDeflaterLevel(w io.Writer, level int) (io.WriteCloser, os.Error) {
    <a id="L42"></a>z := new(writer);
    <a id="L43"></a><span class="comment">// ZLIB has a two-byte header (as documented in RFC 1950).</span>
    <a id="L44"></a><span class="comment">// The first four bits is the CINFO (compression info), which is 7 for the default deflate window size.</span>
    <a id="L45"></a><span class="comment">// The next four bits is the CM (compression method), which is 8 for deflate.</span>
    <a id="L46"></a>z.scratch[0] = 0x78;
    <a id="L47"></a><span class="comment">// The next two bits is the FLEVEL (compression level). The four values are:</span>
    <a id="L48"></a><span class="comment">// 0=fastest, 1=fast, 2=default, 3=best.</span>
    <a id="L49"></a><span class="comment">// The next bit, FDICT, is unused, in this implementation.</span>
    <a id="L50"></a><span class="comment">// The final five FCHECK bits form a mod-31 checksum.</span>
    <a id="L51"></a>switch level {
    <a id="L52"></a>case 0, 1:
        <a id="L53"></a>z.scratch[1] = 0x01
    <a id="L54"></a>case 2, 3, 4, 5:
        <a id="L55"></a>z.scratch[1] = 0x5e
    <a id="L56"></a>case 6, -1:
        <a id="L57"></a>z.scratch[1] = 0x9c
    <a id="L58"></a>case 7, 8, 9:
        <a id="L59"></a>z.scratch[1] = 0xda
    <a id="L60"></a>default:
        <a id="L61"></a>return nil, os.NewError(&#34;level out of range&#34;)
    <a id="L62"></a>}
    <a id="L63"></a>_, err := w.Write(z.scratch[0:2]);
    <a id="L64"></a>if err != nil {
        <a id="L65"></a>return nil, err
    <a id="L66"></a>}
    <a id="L67"></a>z.w = w;
    <a id="L68"></a>z.deflater = flate.NewDeflater(w, level);
    <a id="L69"></a>z.digest = adler32.New();
    <a id="L70"></a>return z, nil;
<a id="L71"></a>}

<a id="L73"></a>func (z *writer) Write(p []byte) (n int, err os.Error) {
    <a id="L74"></a>if z.err != nil {
        <a id="L75"></a>return 0, z.err
    <a id="L76"></a>}
    <a id="L77"></a>if len(p) == 0 {
        <a id="L78"></a>return 0, nil
    <a id="L79"></a>}
    <a id="L80"></a>n, err = z.deflater.Write(p);
    <a id="L81"></a>if err != nil {
        <a id="L82"></a>z.err = err;
        <a id="L83"></a>return;
    <a id="L84"></a>}
    <a id="L85"></a>z.digest.Write(p);
    <a id="L86"></a>return;
<a id="L87"></a>}

<a id="L89"></a><span class="comment">// Calling Close does not close the wrapped io.Writer originally passed to NewDeflater.</span>
<a id="L90"></a>func (z *writer) Close() os.Error {
    <a id="L91"></a>if z.err != nil {
        <a id="L92"></a>return z.err
    <a id="L93"></a>}
    <a id="L94"></a>z.err = z.deflater.Close();
    <a id="L95"></a>if z.err != nil {
        <a id="L96"></a>return z.err
    <a id="L97"></a>}
    <a id="L98"></a>checksum := z.digest.Sum32();
    <a id="L99"></a><span class="comment">// ZLIB (RFC 1950) is big-endian, unlike GZIP (RFC 1952).</span>
    <a id="L100"></a>z.scratch[0] = uint8(checksum &gt;&gt; 24);
    <a id="L101"></a>z.scratch[1] = uint8(checksum &gt;&gt; 16);
    <a id="L102"></a>z.scratch[2] = uint8(checksum &gt;&gt; 8);
    <a id="L103"></a>z.scratch[3] = uint8(checksum &gt;&gt; 0);
    <a id="L104"></a>_, z.err = z.w.Write(z.scratch[0:4]);
    <a id="L105"></a>return z.err;
<a id="L106"></a>}
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
