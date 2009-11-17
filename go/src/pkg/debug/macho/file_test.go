<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/macho/file_test.go</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/debug/macho/file_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package macho

<a id="L7"></a>import (
    <a id="L8"></a>&#34;reflect&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a>type fileTest struct {
    <a id="L13"></a>file     string;
    <a id="L14"></a>hdr      FileHeader;
    <a id="L15"></a>segments []*SegmentHeader;
    <a id="L16"></a>sections []*SectionHeader;
<a id="L17"></a>}

<a id="L19"></a>var fileTests = []fileTest{
    <a id="L20"></a>fileTest{
        <a id="L21"></a>&#34;testdata/gcc-386-darwin-exec&#34;,
        <a id="L22"></a>FileHeader{0xfeedface, Cpu386, 0x3, 0x2, 0xc, 0x3c0, 0x85},
        <a id="L23"></a>[]*SegmentHeader{
            <a id="L24"></a>&amp;SegmentHeader{LoadCmdSegment, 0x38, &#34;__PAGEZERO&#34;, 0x0, 0x1000, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
            <a id="L25"></a>&amp;SegmentHeader{LoadCmdSegment, 0xc0, &#34;__TEXT&#34;, 0x1000, 0x1000, 0x0, 0x1000, 0x7, 0x5, 0x2, 0x0},
            <a id="L26"></a>&amp;SegmentHeader{LoadCmdSegment, 0xc0, &#34;__DATA&#34;, 0x2000, 0x1000, 0x1000, 0x1000, 0x7, 0x3, 0x2, 0x0},
            <a id="L27"></a>&amp;SegmentHeader{LoadCmdSegment, 0x7c, &#34;__IMPORT&#34;, 0x3000, 0x1000, 0x2000, 0x1000, 0x7, 0x7, 0x1, 0x0},
            <a id="L28"></a>&amp;SegmentHeader{LoadCmdSegment, 0x38, &#34;__LINKEDIT&#34;, 0x4000, 0x1000, 0x3000, 0x12c, 0x7, 0x1, 0x0, 0x0},
            <a id="L29"></a>nil,
            <a id="L30"></a>nil,
            <a id="L31"></a>nil,
            <a id="L32"></a>nil,
            <a id="L33"></a>nil,
            <a id="L34"></a>nil,
            <a id="L35"></a>nil,
        <a id="L36"></a>},
        <a id="L37"></a>[]*SectionHeader{
            <a id="L38"></a>&amp;SectionHeader{&#34;__text&#34;, &#34;__TEXT&#34;, 0x1f68, 0x88, 0xf68, 0x2, 0x0, 0x0, 0x80000400},
            <a id="L39"></a>&amp;SectionHeader{&#34;__cstring&#34;, &#34;__TEXT&#34;, 0x1ff0, 0xd, 0xff0, 0x0, 0x0, 0x0, 0x2},
            <a id="L40"></a>&amp;SectionHeader{&#34;__data&#34;, &#34;__DATA&#34;, 0x2000, 0x14, 0x1000, 0x2, 0x0, 0x0, 0x0},
            <a id="L41"></a>&amp;SectionHeader{&#34;__dyld&#34;, &#34;__DATA&#34;, 0x2014, 0x1c, 0x1014, 0x2, 0x0, 0x0, 0x0},
            <a id="L42"></a>&amp;SectionHeader{&#34;__jump_table&#34;, &#34;__IMPORT&#34;, 0x3000, 0xa, 0x2000, 0x6, 0x0, 0x0, 0x4000008},
        <a id="L43"></a>},
    <a id="L44"></a>},
    <a id="L45"></a>fileTest{
        <a id="L46"></a>&#34;testdata/gcc-amd64-darwin-exec&#34;,
        <a id="L47"></a>FileHeader{0xfeedfacf, CpuAmd64, 0x80000003, 0x2, 0xb, 0x568, 0x85},
        <a id="L48"></a>[]*SegmentHeader{
            <a id="L49"></a>&amp;SegmentHeader{LoadCmdSegment64, 0x48, &#34;__PAGEZERO&#34;, 0x0, 0x100000000, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
            <a id="L50"></a>&amp;SegmentHeader{LoadCmdSegment64, 0x1d8, &#34;__TEXT&#34;, 0x100000000, 0x1000, 0x0, 0x1000, 0x7, 0x5, 0x5, 0x0},
            <a id="L51"></a>&amp;SegmentHeader{LoadCmdSegment64, 0x138, &#34;__DATA&#34;, 0x100001000, 0x1000, 0x1000, 0x1000, 0x7, 0x3, 0x3, 0x0},
            <a id="L52"></a>&amp;SegmentHeader{LoadCmdSegment64, 0x48, &#34;__LINKEDIT&#34;, 0x100002000, 0x1000, 0x2000, 0x140, 0x7, 0x1, 0x0, 0x0},
            <a id="L53"></a>nil,
            <a id="L54"></a>nil,
            <a id="L55"></a>nil,
            <a id="L56"></a>nil,
            <a id="L57"></a>nil,
            <a id="L58"></a>nil,
            <a id="L59"></a>nil,
        <a id="L60"></a>},
        <a id="L61"></a>[]*SectionHeader{
            <a id="L62"></a>&amp;SectionHeader{&#34;__text&#34;, &#34;__TEXT&#34;, 0x100000f14, 0x6d, 0xf14, 0x2, 0x0, 0x0, 0x80000400},
            <a id="L63"></a>&amp;SectionHeader{&#34;__symbol_stub1&#34;, &#34;__TEXT&#34;, 0x100000f81, 0xc, 0xf81, 0x0, 0x0, 0x0, 0x80000408},
            <a id="L64"></a>&amp;SectionHeader{&#34;__stub_helper&#34;, &#34;__TEXT&#34;, 0x100000f90, 0x18, 0xf90, 0x2, 0x0, 0x0, 0x0},
            <a id="L65"></a>&amp;SectionHeader{&#34;__cstring&#34;, &#34;__TEXT&#34;, 0x100000fa8, 0xd, 0xfa8, 0x0, 0x0, 0x0, 0x2},
            <a id="L66"></a>&amp;SectionHeader{&#34;__eh_frame&#34;, &#34;__TEXT&#34;, 0x100000fb8, 0x48, 0xfb8, 0x3, 0x0, 0x0, 0x6000000b},
            <a id="L67"></a>&amp;SectionHeader{&#34;__data&#34;, &#34;__DATA&#34;, 0x100001000, 0x1c, 0x1000, 0x3, 0x0, 0x0, 0x0},
            <a id="L68"></a>&amp;SectionHeader{&#34;__dyld&#34;, &#34;__DATA&#34;, 0x100001020, 0x38, 0x1020, 0x3, 0x0, 0x0, 0x0},
            <a id="L69"></a>&amp;SectionHeader{&#34;__la_symbol_ptr&#34;, &#34;__DATA&#34;, 0x100001058, 0x10, 0x1058, 0x2, 0x0, 0x0, 0x7},
        <a id="L70"></a>},
    <a id="L71"></a>},
    <a id="L72"></a>fileTest{
        <a id="L73"></a>&#34;testdata/gcc-amd64-darwin-exec-debug&#34;,
        <a id="L74"></a>FileHeader{0xfeedfacf, CpuAmd64, 0x80000003, 0xa, 0x4, 0x5a0, 0},
        <a id="L75"></a>[]*SegmentHeader{
            <a id="L76"></a>nil,
            <a id="L77"></a>&amp;SegmentHeader{LoadCmdSegment64, 0x1d8, &#34;__TEXT&#34;, 0x100000000, 0x1000, 0x0, 0x0, 0x7, 0x5, 0x5, 0x0},
            <a id="L78"></a>&amp;SegmentHeader{LoadCmdSegment64, 0x138, &#34;__DATA&#34;, 0x100001000, 0x1000, 0x0, 0x0, 0x7, 0x3, 0x3, 0x0},
            <a id="L79"></a>&amp;SegmentHeader{LoadCmdSegment64, 0x278, &#34;__DWARF&#34;, 0x100002000, 0x1000, 0x1000, 0x1bc, 0x7, 0x3, 0x7, 0x0},
        <a id="L80"></a>},
        <a id="L81"></a>[]*SectionHeader{
            <a id="L82"></a>&amp;SectionHeader{&#34;__text&#34;, &#34;__TEXT&#34;, 0x100000f14, 0x0, 0x0, 0x2, 0x0, 0x0, 0x80000400},
            <a id="L83"></a>&amp;SectionHeader{&#34;__symbol_stub1&#34;, &#34;__TEXT&#34;, 0x100000f81, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80000408},
            <a id="L84"></a>&amp;SectionHeader{&#34;__stub_helper&#34;, &#34;__TEXT&#34;, 0x100000f90, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0},
            <a id="L85"></a>&amp;SectionHeader{&#34;__cstring&#34;, &#34;__TEXT&#34;, 0x100000fa8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2},
            <a id="L86"></a>&amp;SectionHeader{&#34;__eh_frame&#34;, &#34;__TEXT&#34;, 0x100000fb8, 0x0, 0x0, 0x3, 0x0, 0x0, 0x6000000b},
            <a id="L87"></a>&amp;SectionHeader{&#34;__data&#34;, &#34;__DATA&#34;, 0x100001000, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0},
            <a id="L88"></a>&amp;SectionHeader{&#34;__dyld&#34;, &#34;__DATA&#34;, 0x100001020, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0},
            <a id="L89"></a>&amp;SectionHeader{&#34;__la_symbol_ptr&#34;, &#34;__DATA&#34;, 0x100001058, 0x0, 0x0, 0x2, 0x0, 0x0, 0x7},
            <a id="L90"></a>&amp;SectionHeader{&#34;__debug_abbrev&#34;, &#34;__DWARF&#34;, 0x100002000, 0x36, 0x1000, 0x0, 0x0, 0x0, 0x0},
            <a id="L91"></a>&amp;SectionHeader{&#34;__debug_aranges&#34;, &#34;__DWARF&#34;, 0x100002036, 0x30, 0x1036, 0x0, 0x0, 0x0, 0x0},
            <a id="L92"></a>&amp;SectionHeader{&#34;__debug_frame&#34;, &#34;__DWARF&#34;, 0x100002066, 0x40, 0x1066, 0x0, 0x0, 0x0, 0x0},
            <a id="L93"></a>&amp;SectionHeader{&#34;__debug_info&#34;, &#34;__DWARF&#34;, 0x1000020a6, 0x54, 0x10a6, 0x0, 0x0, 0x0, 0x0},
            <a id="L94"></a>&amp;SectionHeader{&#34;__debug_line&#34;, &#34;__DWARF&#34;, 0x1000020fa, 0x47, 0x10fa, 0x0, 0x0, 0x0, 0x0},
            <a id="L95"></a>&amp;SectionHeader{&#34;__debug_pubnames&#34;, &#34;__DWARF&#34;, 0x100002141, 0x1b, 0x1141, 0x0, 0x0, 0x0, 0x0},
            <a id="L96"></a>&amp;SectionHeader{&#34;__debug_str&#34;, &#34;__DWARF&#34;, 0x10000215c, 0x60, 0x115c, 0x0, 0x0, 0x0, 0x0},
        <a id="L97"></a>},
    <a id="L98"></a>},
<a id="L99"></a>}

<a id="L101"></a>func TestOpen(t *testing.T) {
    <a id="L102"></a>for i := range fileTests {
        <a id="L103"></a>tt := &amp;fileTests[i];

        <a id="L105"></a>f, err := Open(tt.file);
        <a id="L106"></a>if err != nil {
            <a id="L107"></a>t.Error(err);
            <a id="L108"></a>continue;
        <a id="L109"></a>}
        <a id="L110"></a>if !reflect.DeepEqual(f.FileHeader, tt.hdr) {
            <a id="L111"></a>t.Errorf(&#34;open %s:\n\thave %#v\n\twant %#v\n&#34;, tt.file, f.FileHeader, tt.hdr);
            <a id="L112"></a>continue;
        <a id="L113"></a>}
        <a id="L114"></a>for i, l := range f.Loads {
            <a id="L115"></a>if i &gt;= len(tt.segments) {
                <a id="L116"></a>break
            <a id="L117"></a>}
            <a id="L118"></a>sh := tt.segments[i];
            <a id="L119"></a>s, ok := l.(*Segment);
            <a id="L120"></a>if sh == nil {
                <a id="L121"></a>if ok {
                    <a id="L122"></a>t.Errorf(&#34;open %s, section %d: skipping %#v\n&#34;, tt.file, i, &amp;s.SegmentHeader)
                <a id="L123"></a>}
                <a id="L124"></a>continue;
            <a id="L125"></a>}
            <a id="L126"></a>if !ok {
                <a id="L127"></a>t.Errorf(&#34;open %s, section %d: not *Segment\n&#34;, tt.file, i);
                <a id="L128"></a>continue;
            <a id="L129"></a>}
            <a id="L130"></a>have := &amp;s.SegmentHeader;
            <a id="L131"></a>want := sh;
            <a id="L132"></a>if !reflect.DeepEqual(have, want) {
                <a id="L133"></a>t.Errorf(&#34;open %s, segment %d:\n\thave %#v\n\twant %#v\n&#34;, tt.file, i, have, want)
            <a id="L134"></a>}
        <a id="L135"></a>}
        <a id="L136"></a>tn := len(tt.segments);
        <a id="L137"></a>fn := len(f.Loads);
        <a id="L138"></a>if tn != fn {
            <a id="L139"></a>t.Errorf(&#34;open %s: len(Loads) = %d, want %d&#34;, tt.file, fn, tn)
        <a id="L140"></a>}

        <a id="L142"></a>for i, sh := range f.Sections {
            <a id="L143"></a>if i &gt;= len(tt.sections) {
                <a id="L144"></a>break
            <a id="L145"></a>}
            <a id="L146"></a>have := &amp;sh.SectionHeader;
            <a id="L147"></a>want := tt.sections[i];
            <a id="L148"></a>if !reflect.DeepEqual(have, want) {
                <a id="L149"></a>t.Errorf(&#34;open %s, section %d:\n\thave %#v\n\twant %#v\n&#34;, tt.file, i, have, want)
            <a id="L150"></a>}
        <a id="L151"></a>}
        <a id="L152"></a>tn = len(tt.sections);
        <a id="L153"></a>fn = len(f.Sections);
        <a id="L154"></a>if tn != fn {
            <a id="L155"></a>t.Errorf(&#34;open %s: len(Sections) = %d, want %d&#34;, tt.file, fn, tn)
        <a id="L156"></a>}

    <a id="L158"></a>}
<a id="L159"></a>}
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
