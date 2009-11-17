<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/ogle/arch.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/ogle/arch.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package ogle

<a id="L7"></a>import (
    <a id="L8"></a>&#34;debug/proc&#34;;
    <a id="L9"></a>&#34;math&#34;;
<a id="L10"></a>)

<a id="L12"></a>type Arch interface {
    <a id="L13"></a><span class="comment">// ToWord converts an array of up to 8 bytes in memory order</span>
    <a id="L14"></a><span class="comment">// to a word.</span>
    <a id="L15"></a>ToWord(data []byte) proc.Word;
    <a id="L16"></a><span class="comment">// FromWord converts a word to an array of up to 8 bytes in</span>
    <a id="L17"></a><span class="comment">// memory order.</span>
    <a id="L18"></a>FromWord(v proc.Word, out []byte);
    <a id="L19"></a><span class="comment">// ToFloat32 converts a word to a float.  The order of this</span>
    <a id="L20"></a><span class="comment">// word will be the order returned by ToWord on the memory</span>
    <a id="L21"></a><span class="comment">// representation of a float, and thus may require reversing.</span>
    <a id="L22"></a>ToFloat32(bits uint32) float32;
    <a id="L23"></a><span class="comment">// FromFloat32 converts a float to a word.  This should return</span>
    <a id="L24"></a><span class="comment">// a word that can be passed to FromWord to get the memory</span>
    <a id="L25"></a><span class="comment">// representation of a float on this architecture.</span>
    <a id="L26"></a>FromFloat32(f float32) uint32;
    <a id="L27"></a><span class="comment">// ToFloat64 is to float64 as ToFloat32 is to float32.</span>
    <a id="L28"></a>ToFloat64(bits uint64) float64;
    <a id="L29"></a><span class="comment">// FromFloat64 is to float64 as FromFloat32 is to float32.</span>
    <a id="L30"></a>FromFloat64(f float64) uint64;

    <a id="L32"></a><span class="comment">// IntSize returns the number of bytes in an &#39;int&#39;.</span>
    <a id="L33"></a>IntSize() int;
    <a id="L34"></a><span class="comment">// PtrSize returns the number of bytes in a &#39;uintptr&#39;.</span>
    <a id="L35"></a>PtrSize() int;
    <a id="L36"></a><span class="comment">// FloatSize returns the number of bytes in a &#39;float&#39;.</span>
    <a id="L37"></a>FloatSize() int;
    <a id="L38"></a><span class="comment">// Align rounds offset up to the appropriate offset for a</span>
    <a id="L39"></a><span class="comment">// basic type with the given width.</span>
    <a id="L40"></a>Align(offset, width int) int;

    <a id="L42"></a><span class="comment">// G returns the current G pointer.</span>
    <a id="L43"></a>G(regs proc.Regs) proc.Word;

    <a id="L45"></a><span class="comment">// ClosureSize returns the number of bytes expected by</span>
    <a id="L46"></a><span class="comment">// ParseClosure.</span>
    <a id="L47"></a>ClosureSize() int;
    <a id="L48"></a><span class="comment">// ParseClosure takes ClosureSize bytes read from a return PC</span>
    <a id="L49"></a><span class="comment">// in a remote process, determines if the code is a closure,</span>
    <a id="L50"></a><span class="comment">// and returns the frame size of the closure if it is.</span>
    <a id="L51"></a>ParseClosure(data []byte) (frame int, ok bool);
<a id="L52"></a>}

<a id="L54"></a>type ArchLSB struct{}

<a id="L56"></a>func (ArchLSB) ToWord(data []byte) proc.Word {
    <a id="L57"></a>var v proc.Word;
    <a id="L58"></a>for i, b := range data {
        <a id="L59"></a>v |= proc.Word(b) &lt;&lt; (uint(i) * 8)
    <a id="L60"></a>}
    <a id="L61"></a>return v;
<a id="L62"></a>}

<a id="L64"></a>func (ArchLSB) FromWord(v proc.Word, out []byte) {
    <a id="L65"></a>for i := range out {
        <a id="L66"></a>out[i] = byte(v);
        <a id="L67"></a>v &gt;&gt;= 8;
    <a id="L68"></a>}
<a id="L69"></a>}

<a id="L71"></a>func (ArchLSB) ToFloat32(bits uint32) float32 {
    <a id="L72"></a><span class="comment">// TODO(austin) Do these definitions depend on my current</span>
    <a id="L73"></a><span class="comment">// architecture?</span>
    <a id="L74"></a>return math.Float32frombits(bits)
<a id="L75"></a>}

<a id="L77"></a>func (ArchLSB) FromFloat32(f float32) uint32 { return math.Float32bits(f) }

<a id="L79"></a>func (ArchLSB) ToFloat64(bits uint64) float64 { return math.Float64frombits(bits) }

<a id="L81"></a>func (ArchLSB) FromFloat64(f float64) uint64 { return math.Float64bits(f) }

<a id="L83"></a>type ArchAlignedMultiple struct{}

<a id="L85"></a>func (ArchAlignedMultiple) Align(offset, width int) int {
    <a id="L86"></a>return ((offset - 1) | (width - 1)) + 1
<a id="L87"></a>}

<a id="L89"></a>type amd64 struct {
    <a id="L90"></a>ArchLSB;
    <a id="L91"></a>ArchAlignedMultiple;
    <a id="L92"></a>gReg int;
<a id="L93"></a>}

<a id="L95"></a>func (a *amd64) IntSize() int { return 4 }

<a id="L97"></a>func (a *amd64) PtrSize() int { return 8 }

<a id="L99"></a>func (a *amd64) FloatSize() int { return 4 }

<a id="L101"></a>func (a *amd64) G(regs proc.Regs) proc.Word {
    <a id="L102"></a><span class="comment">// See src/pkg/runtime/mkasmh</span>
    <a id="L103"></a>if a.gReg == -1 {
        <a id="L104"></a>ns := regs.Names();
        <a id="L105"></a>for i, n := range ns {
            <a id="L106"></a>if n == &#34;r15&#34; {
                <a id="L107"></a>a.gReg = i;
                <a id="L108"></a>break;
            <a id="L109"></a>}
        <a id="L110"></a>}
    <a id="L111"></a>}

    <a id="L113"></a>return regs.Get(a.gReg);
<a id="L114"></a>}

<a id="L116"></a>func (a *amd64) ClosureSize() int { return 8 }

<a id="L118"></a>func (a *amd64) ParseClosure(data []byte) (int, bool) {
    <a id="L119"></a>if data[0] == 0x48 &amp;&amp; data[1] == 0x81 &amp;&amp; data[2] == 0xc4 &amp;&amp; data[7] == 0xc3 {
        <a id="L120"></a>return int(a.ToWord(data[3:7]) + 8), true
    <a id="L121"></a>}
    <a id="L122"></a>return 0, false;
<a id="L123"></a>}

<a id="L125"></a>var Amd64 = &amp;amd64{gReg: -1}
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
