<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/elf/file_test.go</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/debug/elf/file_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package elf

<a id="L7"></a>import (
    <a id="L8"></a>&#34;debug/dwarf&#34;;
    <a id="L9"></a>&#34;encoding/binary&#34;;
    <a id="L10"></a>&#34;reflect&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)

<a id="L14"></a>type fileTest struct {
    <a id="L15"></a>file     string;
    <a id="L16"></a>hdr      FileHeader;
    <a id="L17"></a>sections []SectionHeader;
<a id="L18"></a>}

<a id="L20"></a>var fileTests = []fileTest{
    <a id="L21"></a>fileTest{
        <a id="L22"></a>&#34;testdata/gcc-386-freebsd-exec&#34;,
        <a id="L23"></a>FileHeader{ELFCLASS32, ELFDATA2LSB, EV_CURRENT, ELFOSABI_FREEBSD, 0, binary.LittleEndian, ET_EXEC, EM_386},
        <a id="L24"></a>[]SectionHeader{
            <a id="L25"></a>SectionHeader{&#34;&#34;, SHT_NULL, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
            <a id="L26"></a>SectionHeader{&#34;.interp&#34;, SHT_PROGBITS, SHF_ALLOC, 0x80480d4, 0xd4, 0x15, 0x0, 0x0, 0x1, 0x0},
            <a id="L27"></a>SectionHeader{&#34;.hash&#34;, SHT_HASH, SHF_ALLOC, 0x80480ec, 0xec, 0x90, 0x3, 0x0, 0x4, 0x4},
            <a id="L28"></a>SectionHeader{&#34;.dynsym&#34;, SHT_DYNSYM, SHF_ALLOC, 0x804817c, 0x17c, 0x110, 0x4, 0x1, 0x4, 0x10},
            <a id="L29"></a>SectionHeader{&#34;.dynstr&#34;, SHT_STRTAB, SHF_ALLOC, 0x804828c, 0x28c, 0xbb, 0x0, 0x0, 0x1, 0x0},
            <a id="L30"></a>SectionHeader{&#34;.rel.plt&#34;, SHT_REL, SHF_ALLOC, 0x8048348, 0x348, 0x20, 0x3, 0x7, 0x4, 0x8},
            <a id="L31"></a>SectionHeader{&#34;.init&#34;, SHT_PROGBITS, SHF_ALLOC + SHF_EXECINSTR, 0x8048368, 0x368, 0x11, 0x0, 0x0, 0x4, 0x0},
            <a id="L32"></a>SectionHeader{&#34;.plt&#34;, SHT_PROGBITS, SHF_ALLOC + SHF_EXECINSTR, 0x804837c, 0x37c, 0x50, 0x0, 0x0, 0x4, 0x4},
            <a id="L33"></a>SectionHeader{&#34;.text&#34;, SHT_PROGBITS, SHF_ALLOC + SHF_EXECINSTR, 0x80483cc, 0x3cc, 0x180, 0x0, 0x0, 0x4, 0x0},
            <a id="L34"></a>SectionHeader{&#34;.fini&#34;, SHT_PROGBITS, SHF_ALLOC + SHF_EXECINSTR, 0x804854c, 0x54c, 0xc, 0x0, 0x0, 0x4, 0x0},
            <a id="L35"></a>SectionHeader{&#34;.rodata&#34;, SHT_PROGBITS, SHF_ALLOC, 0x8048558, 0x558, 0xa3, 0x0, 0x0, 0x1, 0x0},
            <a id="L36"></a>SectionHeader{&#34;.data&#34;, SHT_PROGBITS, SHF_WRITE + SHF_ALLOC, 0x80495fc, 0x5fc, 0xc, 0x0, 0x0, 0x4, 0x0},
            <a id="L37"></a>SectionHeader{&#34;.eh_frame&#34;, SHT_PROGBITS, SHF_ALLOC, 0x8049608, 0x608, 0x4, 0x0, 0x0, 0x4, 0x0},
            <a id="L38"></a>SectionHeader{&#34;.dynamic&#34;, SHT_DYNAMIC, SHF_WRITE + SHF_ALLOC, 0x804960c, 0x60c, 0x98, 0x4, 0x0, 0x4, 0x8},
            <a id="L39"></a>SectionHeader{&#34;.ctors&#34;, SHT_PROGBITS, SHF_WRITE + SHF_ALLOC, 0x80496a4, 0x6a4, 0x8, 0x0, 0x0, 0x4, 0x0},
            <a id="L40"></a>SectionHeader{&#34;.dtors&#34;, SHT_PROGBITS, SHF_WRITE + SHF_ALLOC, 0x80496ac, 0x6ac, 0x8, 0x0, 0x0, 0x4, 0x0},
            <a id="L41"></a>SectionHeader{&#34;.jcr&#34;, SHT_PROGBITS, SHF_WRITE + SHF_ALLOC, 0x80496b4, 0x6b4, 0x4, 0x0, 0x0, 0x4, 0x0},
            <a id="L42"></a>SectionHeader{&#34;.got&#34;, SHT_PROGBITS, SHF_WRITE + SHF_ALLOC, 0x80496b8, 0x6b8, 0x1c, 0x0, 0x0, 0x4, 0x4},
            <a id="L43"></a>SectionHeader{&#34;.bss&#34;, SHT_NOBITS, SHF_WRITE + SHF_ALLOC, 0x80496d4, 0x6d4, 0x20, 0x0, 0x0, 0x4, 0x0},
            <a id="L44"></a>SectionHeader{&#34;.comment&#34;, SHT_PROGBITS, 0x0, 0x0, 0x6d4, 0x12d, 0x0, 0x0, 0x1, 0x0},
            <a id="L45"></a>SectionHeader{&#34;.debug_aranges&#34;, SHT_PROGBITS, 0x0, 0x0, 0x801, 0x20, 0x0, 0x0, 0x1, 0x0},
            <a id="L46"></a>SectionHeader{&#34;.debug_pubnames&#34;, SHT_PROGBITS, 0x0, 0x0, 0x821, 0x1b, 0x0, 0x0, 0x1, 0x0},
            <a id="L47"></a>SectionHeader{&#34;.debug_info&#34;, SHT_PROGBITS, 0x0, 0x0, 0x83c, 0x11d, 0x0, 0x0, 0x1, 0x0},
            <a id="L48"></a>SectionHeader{&#34;.debug_abbrev&#34;, SHT_PROGBITS, 0x0, 0x0, 0x959, 0x41, 0x0, 0x0, 0x1, 0x0},
            <a id="L49"></a>SectionHeader{&#34;.debug_line&#34;, SHT_PROGBITS, 0x0, 0x0, 0x99a, 0x35, 0x0, 0x0, 0x1, 0x0},
            <a id="L50"></a>SectionHeader{&#34;.debug_frame&#34;, SHT_PROGBITS, 0x0, 0x0, 0x9d0, 0x30, 0x0, 0x0, 0x4, 0x0},
            <a id="L51"></a>SectionHeader{&#34;.debug_str&#34;, SHT_PROGBITS, 0x0, 0x0, 0xa00, 0xd, 0x0, 0x0, 0x1, 0x0},
            <a id="L52"></a>SectionHeader{&#34;.shstrtab&#34;, SHT_STRTAB, 0x0, 0x0, 0xa0d, 0xf8, 0x0, 0x0, 0x1, 0x0},
            <a id="L53"></a>SectionHeader{&#34;.symtab&#34;, SHT_SYMTAB, 0x0, 0x0, 0xfb8, 0x4b0, 0x1d, 0x38, 0x4, 0x10},
            <a id="L54"></a>SectionHeader{&#34;.strtab&#34;, SHT_STRTAB, 0x0, 0x0, 0x1468, 0x206, 0x0, 0x0, 0x1, 0x0},
        <a id="L55"></a>},
    <a id="L56"></a>},
    <a id="L57"></a>fileTest{
        <a id="L58"></a>&#34;testdata/gcc-amd64-linux-exec&#34;,
        <a id="L59"></a>FileHeader{ELFCLASS64, ELFDATA2LSB, EV_CURRENT, ELFOSABI_NONE, 0, binary.LittleEndian, ET_EXEC, EM_X86_64},
        <a id="L60"></a>[]SectionHeader{
            <a id="L61"></a>SectionHeader{&#34;&#34;, SHT_NULL, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
            <a id="L62"></a>SectionHeader{&#34;.interp&#34;, SHT_PROGBITS, SHF_ALLOC, 0x400200, 0x200, 0x1c, 0x0, 0x0, 0x1, 0x0},
            <a id="L63"></a>SectionHeader{&#34;.note.ABI-tag&#34;, SHT_NOTE, SHF_ALLOC, 0x40021c, 0x21c, 0x20, 0x0, 0x0, 0x4, 0x0},
            <a id="L64"></a>SectionHeader{&#34;.hash&#34;, SHT_HASH, SHF_ALLOC, 0x400240, 0x240, 0x24, 0x5, 0x0, 0x8, 0x4},
            <a id="L65"></a>SectionHeader{&#34;.gnu.hash&#34;, SHT_LOOS + 268435446, SHF_ALLOC, 0x400268, 0x268, 0x1c, 0x5, 0x0, 0x8, 0x0},
            <a id="L66"></a>SectionHeader{&#34;.dynsym&#34;, SHT_DYNSYM, SHF_ALLOC, 0x400288, 0x288, 0x60, 0x6, 0x1, 0x8, 0x18},
            <a id="L67"></a>SectionHeader{&#34;.dynstr&#34;, SHT_STRTAB, SHF_ALLOC, 0x4002e8, 0x2e8, 0x3d, 0x0, 0x0, 0x1, 0x0},
            <a id="L68"></a>SectionHeader{&#34;.gnu.version&#34;, SHT_HIOS, SHF_ALLOC, 0x400326, 0x326, 0x8, 0x5, 0x0, 0x2, 0x2},
            <a id="L69"></a>SectionHeader{&#34;.gnu.version_r&#34;, SHT_LOOS + 268435454, SHF_ALLOC, 0x400330, 0x330, 0x20, 0x6, 0x1, 0x8, 0x0},
            <a id="L70"></a>SectionHeader{&#34;.rela.dyn&#34;, SHT_RELA, SHF_ALLOC, 0x400350, 0x350, 0x18, 0x5, 0x0, 0x8, 0x18},
            <a id="L71"></a>SectionHeader{&#34;.rela.plt&#34;, SHT_RELA, SHF_ALLOC, 0x400368, 0x368, 0x30, 0x5, 0xc, 0x8, 0x18},
            <a id="L72"></a>SectionHeader{&#34;.init&#34;, SHT_PROGBITS, SHF_ALLOC + SHF_EXECINSTR, 0x400398, 0x398, 0x18, 0x0, 0x0, 0x4, 0x0},
            <a id="L73"></a>SectionHeader{&#34;.plt&#34;, SHT_PROGBITS, SHF_ALLOC + SHF_EXECINSTR, 0x4003b0, 0x3b0, 0x30, 0x0, 0x0, 0x4, 0x10},
            <a id="L74"></a>SectionHeader{&#34;.text&#34;, SHT_PROGBITS, SHF_ALLOC + SHF_EXECINSTR, 0x4003e0, 0x3e0, 0x1b4, 0x0, 0x0, 0x10, 0x0},
            <a id="L75"></a>SectionHeader{&#34;.fini&#34;, SHT_PROGBITS, SHF_ALLOC + SHF_EXECINSTR, 0x400594, 0x594, 0xe, 0x0, 0x0, 0x4, 0x0},
            <a id="L76"></a>SectionHeader{&#34;.rodata&#34;, SHT_PROGBITS, SHF_ALLOC, 0x4005a4, 0x5a4, 0x11, 0x0, 0x0, 0x4, 0x0},
            <a id="L77"></a>SectionHeader{&#34;.eh_frame_hdr&#34;, SHT_PROGBITS, SHF_ALLOC, 0x4005b8, 0x5b8, 0x24, 0x0, 0x0, 0x4, 0x0},
            <a id="L78"></a>SectionHeader{&#34;.eh_frame&#34;, SHT_PROGBITS, SHF_ALLOC, 0x4005e0, 0x5e0, 0xa4, 0x0, 0x0, 0x8, 0x0},
            <a id="L79"></a>SectionHeader{&#34;.ctors&#34;, SHT_PROGBITS, SHF_WRITE + SHF_ALLOC, 0x600688, 0x688, 0x10, 0x0, 0x0, 0x8, 0x0},
            <a id="L80"></a>SectionHeader{&#34;.dtors&#34;, SHT_PROGBITS, SHF_WRITE + SHF_ALLOC, 0x600698, 0x698, 0x10, 0x0, 0x0, 0x8, 0x0},
            <a id="L81"></a>SectionHeader{&#34;.jcr&#34;, SHT_PROGBITS, SHF_WRITE + SHF_ALLOC, 0x6006a8, 0x6a8, 0x8, 0x0, 0x0, 0x8, 0x0},
            <a id="L82"></a>SectionHeader{&#34;.dynamic&#34;, SHT_DYNAMIC, SHF_WRITE + SHF_ALLOC, 0x6006b0, 0x6b0, 0x1a0, 0x6, 0x0, 0x8, 0x10},
            <a id="L83"></a>SectionHeader{&#34;.got&#34;, SHT_PROGBITS, SHF_WRITE + SHF_ALLOC, 0x600850, 0x850, 0x8, 0x0, 0x0, 0x8, 0x8},
            <a id="L84"></a>SectionHeader{&#34;.got.plt&#34;, SHT_PROGBITS, SHF_WRITE + SHF_ALLOC, 0x600858, 0x858, 0x28, 0x0, 0x0, 0x8, 0x8},
            <a id="L85"></a>SectionHeader{&#34;.data&#34;, SHT_PROGBITS, SHF_WRITE + SHF_ALLOC, 0x600880, 0x880, 0x18, 0x0, 0x0, 0x8, 0x0},
            <a id="L86"></a>SectionHeader{&#34;.bss&#34;, SHT_NOBITS, SHF_WRITE + SHF_ALLOC, 0x600898, 0x898, 0x8, 0x0, 0x0, 0x4, 0x0},
            <a id="L87"></a>SectionHeader{&#34;.comment&#34;, SHT_PROGBITS, 0x0, 0x0, 0x898, 0x126, 0x0, 0x0, 0x1, 0x0},
            <a id="L88"></a>SectionHeader{&#34;.debug_aranges&#34;, SHT_PROGBITS, 0x0, 0x0, 0x9c0, 0x90, 0x0, 0x0, 0x10, 0x0},
            <a id="L89"></a>SectionHeader{&#34;.debug_pubnames&#34;, SHT_PROGBITS, 0x0, 0x0, 0xa50, 0x25, 0x0, 0x0, 0x1, 0x0},
            <a id="L90"></a>SectionHeader{&#34;.debug_info&#34;, SHT_PROGBITS, 0x0, 0x0, 0xa75, 0x1a7, 0x0, 0x0, 0x1, 0x0},
            <a id="L91"></a>SectionHeader{&#34;.debug_abbrev&#34;, SHT_PROGBITS, 0x0, 0x0, 0xc1c, 0x6f, 0x0, 0x0, 0x1, 0x0},
            <a id="L92"></a>SectionHeader{&#34;.debug_line&#34;, SHT_PROGBITS, 0x0, 0x0, 0xc8b, 0x13f, 0x0, 0x0, 0x1, 0x0},
            <a id="L93"></a>SectionHeader{&#34;.debug_str&#34;, SHT_PROGBITS, SHF_MERGE + SHF_STRINGS, 0x0, 0xdca, 0xb1, 0x0, 0x0, 0x1, 0x1},
            <a id="L94"></a>SectionHeader{&#34;.debug_ranges&#34;, SHT_PROGBITS, 0x0, 0x0, 0xe80, 0x90, 0x0, 0x0, 0x10, 0x0},
            <a id="L95"></a>SectionHeader{&#34;.shstrtab&#34;, SHT_STRTAB, 0x0, 0x0, 0xf10, 0x149, 0x0, 0x0, 0x1, 0x0},
            <a id="L96"></a>SectionHeader{&#34;.symtab&#34;, SHT_SYMTAB, 0x0, 0x0, 0x19a0, 0x6f0, 0x24, 0x39, 0x8, 0x18},
            <a id="L97"></a>SectionHeader{&#34;.strtab&#34;, SHT_STRTAB, 0x0, 0x0, 0x2090, 0x1fc, 0x0, 0x0, 0x1, 0x0},
        <a id="L98"></a>},
    <a id="L99"></a>},
<a id="L100"></a>}

<a id="L102"></a>func TestOpen(t *testing.T) {
    <a id="L103"></a>for i := range fileTests {
        <a id="L104"></a>tt := &amp;fileTests[i];

        <a id="L106"></a>f, err := Open(tt.file);
        <a id="L107"></a>if err != nil {
            <a id="L108"></a>t.Error(err);
            <a id="L109"></a>continue;
        <a id="L110"></a>}
        <a id="L111"></a>if !reflect.DeepEqual(f.FileHeader, tt.hdr) {
            <a id="L112"></a>t.Errorf(&#34;open %s:\n\thave %#v\n\twant %#v\n&#34;, tt.file, f.FileHeader, tt.hdr);
            <a id="L113"></a>continue;
        <a id="L114"></a>}
        <a id="L115"></a>for i, s := range f.Sections {
            <a id="L116"></a>if i &gt;= len(tt.sections) {
                <a id="L117"></a>break
            <a id="L118"></a>}
            <a id="L119"></a>sh := &amp;tt.sections[i];
            <a id="L120"></a>if !reflect.DeepEqual(&amp;s.SectionHeader, sh) {
                <a id="L121"></a>t.Errorf(&#34;open %s, section %d:\n\thave %#v\n\twant %#v\n&#34;, tt.file, i, &amp;s.SectionHeader, sh)
            <a id="L122"></a>}
        <a id="L123"></a>}
        <a id="L124"></a>tn := len(tt.sections);
        <a id="L125"></a>fn := len(f.Sections);
        <a id="L126"></a>if tn != fn {
            <a id="L127"></a>t.Errorf(&#34;open %s: len(Sections) = %d, want %d&#34;, tt.file, fn, tn)
        <a id="L128"></a>}
    <a id="L129"></a>}
<a id="L130"></a>}

<a id="L132"></a>type relocationTest struct {
    <a id="L133"></a>file       string;
    <a id="L134"></a>firstEntry *dwarf.Entry;
<a id="L135"></a>}

<a id="L137"></a>var relocationTests = []relocationTest{
    <a id="L138"></a>relocationTest{
        <a id="L139"></a>&#34;testdata/go-relocation-test-gcc441-x86-64.o&#34;,
        <a id="L140"></a>&amp;dwarf.Entry{Offset: 0xb, Tag: dwarf.TagCompileUnit, Children: true, Field: []dwarf.Field{dwarf.Field{Attr: dwarf.AttrProducer, Val: &#34;GNU C 4.4.1&#34;}, dwarf.Field{Attr: dwarf.AttrLanguage, Val: int64(1)}, dwarf.Field{Attr: dwarf.AttrName, Val: &#34;go-relocation-test.c&#34;}, dwarf.Field{Attr: dwarf.AttrCompDir, Val: &#34;/tmp&#34;}, dwarf.Field{Attr: dwarf.AttrLowpc, Val: uint64(0x0)}, dwarf.Field{Attr: dwarf.AttrHighpc, Val: uint64(0x6)}, dwarf.Field{Attr: dwarf.AttrStmtList, Val: int64(0)}}},
    <a id="L141"></a>},
    <a id="L142"></a>relocationTest{
        <a id="L143"></a>&#34;testdata/go-relocation-test-gcc441-x86.o&#34;,
        <a id="L144"></a>&amp;dwarf.Entry{Offset: 0xb, Tag: dwarf.TagCompileUnit, Children: true, Field: []dwarf.Field{dwarf.Field{Attr: dwarf.AttrProducer, Val: &#34;GNU C 4.4.1&#34;}, dwarf.Field{Attr: dwarf.AttrLanguage, Val: int64(1)}, dwarf.Field{Attr: dwarf.AttrName, Val: &#34;t.c&#34;}, dwarf.Field{Attr: dwarf.AttrCompDir, Val: &#34;/tmp&#34;}, dwarf.Field{Attr: dwarf.AttrLowpc, Val: uint64(0x0)}, dwarf.Field{Attr: dwarf.AttrHighpc, Val: uint64(0x5)}, dwarf.Field{Attr: dwarf.AttrStmtList, Val: int64(0)}}},
    <a id="L145"></a>},
    <a id="L146"></a>relocationTest{
        <a id="L147"></a>&#34;testdata/go-relocation-test-gcc424-x86-64.o&#34;,
        <a id="L148"></a>&amp;dwarf.Entry{Offset: 0xb, Tag: dwarf.TagCompileUnit, Children: true, Field: []dwarf.Field{dwarf.Field{Attr: dwarf.AttrProducer, Val: &#34;GNU C 4.2.4 (Ubuntu 4.2.4-1ubuntu4)&#34;}, dwarf.Field{Attr: dwarf.AttrLanguage, Val: int64(1)}, dwarf.Field{Attr: dwarf.AttrName, Val: &#34;go-relocation-test-gcc424.c&#34;}, dwarf.Field{Attr: dwarf.AttrCompDir, Val: &#34;/tmp&#34;}, dwarf.Field{Attr: dwarf.AttrLowpc, Val: uint64(0x0)}, dwarf.Field{Attr: dwarf.AttrHighpc, Val: uint64(0x6)}, dwarf.Field{Attr: dwarf.AttrStmtList, Val: int64(0)}}},
    <a id="L149"></a>},
<a id="L150"></a>}

<a id="L152"></a>func TestDWARFRelocations(t *testing.T) {
    <a id="L153"></a>for i, test := range relocationTests {
        <a id="L154"></a>f, err := Open(test.file);
        <a id="L155"></a>if err != nil {
            <a id="L156"></a>t.Error(err);
            <a id="L157"></a>continue;
        <a id="L158"></a>}
        <a id="L159"></a>dwarf, err := f.DWARF();
        <a id="L160"></a>if err != nil {
            <a id="L161"></a>t.Error(err);
            <a id="L162"></a>continue;
        <a id="L163"></a>}
        <a id="L164"></a>reader := dwarf.Reader();
        <a id="L165"></a><span class="comment">// Checking only the first entry is sufficient since it has</span>
        <a id="L166"></a><span class="comment">// many different strings. If the relocation had failed, all</span>
        <a id="L167"></a><span class="comment">// the string offsets would be zero and all the strings would</span>
        <a id="L168"></a><span class="comment">// end up being the same.</span>
        <a id="L169"></a>firstEntry, err := reader.Next();
        <a id="L170"></a>if err != nil {
            <a id="L171"></a>t.Error(err);
            <a id="L172"></a>continue;
        <a id="L173"></a>}

        <a id="L175"></a>if !reflect.DeepEqual(test.firstEntry, firstEntry) {
            <a id="L176"></a>t.Errorf(&#34;#%d: mismatch: got:%#v want:%#v&#34;, i, firstEntry, test.firstEntry);
            <a id="L177"></a>continue;
        <a id="L178"></a>}
    <a id="L179"></a>}
<a id="L180"></a>}
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
