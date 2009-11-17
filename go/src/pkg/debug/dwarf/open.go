<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/dwarf/open.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/debug/dwarf/open.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package provides access to DWARF debugging information</span>
<a id="L6"></a><span class="comment">// loaded from executable files, as defined in the DWARF 2.0 Standard</span>
<a id="L7"></a><span class="comment">// at http://dwarfstd.org/dwarf-2.0.0.pdf.</span>
<a id="L8"></a>package dwarf

<a id="L10"></a>import (
    <a id="L11"></a>&#34;encoding/binary&#34;;
    <a id="L12"></a>&#34;os&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">// Data represents the DWARF debugging information</span>
<a id="L16"></a><span class="comment">// loaded from an executable file (for example, an ELF or Mach-O executable).</span>
<a id="L17"></a>type Data struct {
    <a id="L18"></a><span class="comment">// raw data</span>
    <a id="L19"></a>abbrev   []byte;
    <a id="L20"></a>aranges  []byte;
    <a id="L21"></a>frame    []byte;
    <a id="L22"></a>info     []byte;
    <a id="L23"></a>line     []byte;
    <a id="L24"></a>pubnames []byte;
    <a id="L25"></a>ranges   []byte;
    <a id="L26"></a>str      []byte;

    <a id="L28"></a><span class="comment">// parsed data</span>
    <a id="L29"></a>abbrevCache map[uint32]abbrevTable;
    <a id="L30"></a>addrsize    int;
    <a id="L31"></a>order       binary.ByteOrder;
    <a id="L32"></a>typeCache   map[Offset]Type;
    <a id="L33"></a>unit        []unit;
<a id="L34"></a>}

<a id="L36"></a><span class="comment">// New returns a new Data object initialized from the given parameters.</span>
<a id="L37"></a><span class="comment">// Clients should typically use [TODO(rsc): method to be named later] instead of calling</span>
<a id="L38"></a><span class="comment">// New directly.</span>
<a id="L39"></a><span class="comment">//</span>
<a id="L40"></a><span class="comment">// The []byte arguments are the data from the corresponding debug section</span>
<a id="L41"></a><span class="comment">// in the object file; for example, for an ELF object, abbrev is the contents of</span>
<a id="L42"></a><span class="comment">// the &#34;.debug_abbrev&#34; section.</span>
<a id="L43"></a>func New(abbrev, aranges, frame, info, line, pubnames, ranges, str []byte) (*Data, os.Error) {
    <a id="L44"></a>d := &amp;Data{
        <a id="L45"></a>abbrev: abbrev,
        <a id="L46"></a>aranges: aranges,
        <a id="L47"></a>frame: frame,
        <a id="L48"></a>info: info,
        <a id="L49"></a>line: line,
        <a id="L50"></a>pubnames: pubnames,
        <a id="L51"></a>ranges: ranges,
        <a id="L52"></a>str: str,
        <a id="L53"></a>abbrevCache: make(map[uint32]abbrevTable),
        <a id="L54"></a>typeCache: make(map[Offset]Type),
    <a id="L55"></a>};

    <a id="L57"></a><span class="comment">// Sniff .debug_info to figure out byte order.</span>
    <a id="L58"></a><span class="comment">// bytes 4:6 are the version, a tiny 16-bit number (1, 2, 3).</span>
    <a id="L59"></a>if len(d.info) &lt; 6 {
        <a id="L60"></a>return nil, DecodeError{&#34;info&#34;, Offset(len(d.info)), &#34;too short&#34;}
    <a id="L61"></a>}
    <a id="L62"></a>x, y := d.info[4], d.info[5];
    <a id="L63"></a>switch {
    <a id="L64"></a>case x == 0 &amp;&amp; y == 0:
        <a id="L65"></a>return nil, DecodeError{&#34;info&#34;, 4, &#34;unsupported version 0&#34;}
    <a id="L66"></a>case x == 0:
        <a id="L67"></a>d.order = binary.BigEndian
    <a id="L68"></a>case y == 0:
        <a id="L69"></a>d.order = binary.LittleEndian
    <a id="L70"></a>default:
        <a id="L71"></a>return nil, DecodeError{&#34;info&#34;, 4, &#34;cannot determine byte order&#34;}
    <a id="L72"></a>}

    <a id="L74"></a>u, err := d.parseUnits();
    <a id="L75"></a>if err != nil {
        <a id="L76"></a>return nil, err
    <a id="L77"></a>}
    <a id="L78"></a>d.unit = u;
    <a id="L79"></a>return d, nil;
<a id="L80"></a>}
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
