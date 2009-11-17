<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/gob/decoder.go</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/gob/decoder.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package gob

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;io&#34;;
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;sync&#34;;
<a id="L12"></a>)

<a id="L14"></a><span class="comment">// A Decoder manages the receipt of type and data information read from the</span>
<a id="L15"></a><span class="comment">// remote side of a connection.</span>
<a id="L16"></a>type Decoder struct {
    <a id="L17"></a>mutex      sync.Mutex;           <span class="comment">// each item must be received atomically</span>
    <a id="L18"></a>r          io.Reader;            <span class="comment">// source of the data</span>
    <a id="L19"></a>seen       map[typeId]*wireType; <span class="comment">// which types we&#39;ve already seen described</span>
    <a id="L20"></a>state      *decodeState;         <span class="comment">// reads data from in-memory buffer</span>
    <a id="L21"></a>countState *decodeState;         <span class="comment">// reads counts from wire</span>
    <a id="L22"></a>buf        []byte;
    <a id="L23"></a>oneByte    []byte;
<a id="L24"></a>}

<a id="L26"></a><span class="comment">// NewDecoder returns a new decoder that reads from the io.Reader.</span>
<a id="L27"></a>func NewDecoder(r io.Reader) *Decoder {
    <a id="L28"></a>dec := new(Decoder);
    <a id="L29"></a>dec.r = r;
    <a id="L30"></a>dec.seen = make(map[typeId]*wireType);
    <a id="L31"></a>dec.state = newDecodeState(nil); <span class="comment">// buffer set in Decode(); rest is unimportant</span>
    <a id="L32"></a>dec.oneByte = make([]byte, 1);

    <a id="L34"></a>return dec;
<a id="L35"></a>}

<a id="L37"></a>func (dec *Decoder) recvType(id typeId) {
    <a id="L38"></a><span class="comment">// Have we already seen this type?  That&#39;s an error</span>
    <a id="L39"></a>if _, alreadySeen := dec.seen[id]; alreadySeen {
        <a id="L40"></a>dec.state.err = os.ErrorString(&#34;gob: duplicate type received&#34;);
        <a id="L41"></a>return;
    <a id="L42"></a>}

    <a id="L44"></a><span class="comment">// Type:</span>
    <a id="L45"></a>wire := new(wireType);
    <a id="L46"></a>decode(dec.state.b, tWireType, wire);
    <a id="L47"></a><span class="comment">// Remember we&#39;ve seen this type.</span>
    <a id="L48"></a>dec.seen[id] = wire;
<a id="L49"></a>}

<a id="L51"></a><span class="comment">// Decode reads the next value from the connection and stores</span>
<a id="L52"></a><span class="comment">// it in the data represented by the empty interface value.</span>
<a id="L53"></a><span class="comment">// The value underlying e must be the correct type for the next</span>
<a id="L54"></a><span class="comment">// data item received.</span>
<a id="L55"></a>func (dec *Decoder) Decode(e interface{}) os.Error {
    <a id="L56"></a><span class="comment">// Make sure we&#39;re single-threaded through here.</span>
    <a id="L57"></a>dec.mutex.Lock();
    <a id="L58"></a>defer dec.mutex.Unlock();

    <a id="L60"></a>dec.state.err = nil;
    <a id="L61"></a>for {
        <a id="L62"></a><span class="comment">// Read a count.</span>
        <a id="L63"></a>var nbytes uint64;
        <a id="L64"></a>nbytes, dec.state.err = decodeUintReader(dec.r, dec.oneByte);
        <a id="L65"></a>if dec.state.err != nil {
            <a id="L66"></a>break
        <a id="L67"></a>}
        <a id="L68"></a><span class="comment">// Allocate the buffer.</span>
        <a id="L69"></a>if nbytes &gt; uint64(len(dec.buf)) {
            <a id="L70"></a>dec.buf = make([]byte, nbytes+1000)
        <a id="L71"></a>}
        <a id="L72"></a>dec.state.b = bytes.NewBuffer(dec.buf[0:nbytes]);

        <a id="L74"></a><span class="comment">// Read the data</span>
        <a id="L75"></a>_, dec.state.err = io.ReadFull(dec.r, dec.buf[0:nbytes]);
        <a id="L76"></a>if dec.state.err != nil {
            <a id="L77"></a>if dec.state.err == os.EOF {
                <a id="L78"></a>dec.state.err = io.ErrUnexpectedEOF
            <a id="L79"></a>}
            <a id="L80"></a>break;
        <a id="L81"></a>}

        <a id="L83"></a><span class="comment">// Receive a type id.</span>
        <a id="L84"></a>id := typeId(decodeInt(dec.state));
        <a id="L85"></a>if dec.state.err != nil {
            <a id="L86"></a>break
        <a id="L87"></a>}

        <a id="L89"></a><span class="comment">// Is it a new type?</span>
        <a id="L90"></a>if id &lt; 0 { <span class="comment">// 0 is the error state, handled above</span>
            <a id="L91"></a><span class="comment">// If the id is negative, we have a type.</span>
            <a id="L92"></a>dec.recvType(-id);
            <a id="L93"></a>if dec.state.err != nil {
                <a id="L94"></a>break
            <a id="L95"></a>}
            <a id="L96"></a>continue;
        <a id="L97"></a>}

        <a id="L99"></a><span class="comment">// No, it&#39;s a value.</span>
        <a id="L100"></a>dec.state.err = decode(dec.state.b, id, e);
        <a id="L101"></a>break;
    <a id="L102"></a>}
    <a id="L103"></a>return dec.state.err;
<a id="L104"></a>}
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
