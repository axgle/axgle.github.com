<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/macho/file.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/debug/macho/file.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Package macho implements access to Mach-O object files, as defined by</span>
<a id="L6"></a><span class="comment">// http://developer.apple.com/mac/library/documentation/DeveloperTools/Conceptual/MachORuntime/Reference/reference.html.</span>
<a id="L7"></a>package macho

<a id="L9"></a><span class="comment">// High level access to low level data structures.</span>

<a id="L11"></a>import (
    <a id="L12"></a>&#34;bytes&#34;;
    <a id="L13"></a>&#34;debug/dwarf&#34;;
    <a id="L14"></a>&#34;encoding/binary&#34;;
    <a id="L15"></a>&#34;fmt&#34;;
    <a id="L16"></a>&#34;io&#34;;
    <a id="L17"></a>&#34;os&#34;;
<a id="L18"></a>)

<a id="L20"></a><span class="comment">// A File represents an open Mach-O file.</span>
<a id="L21"></a>type File struct {
    <a id="L22"></a>FileHeader;
    <a id="L23"></a>ByteOrder binary.ByteOrder;
    <a id="L24"></a>Loads     []Load;
    <a id="L25"></a>Sections  []*Section;

    <a id="L27"></a>closer io.Closer;
<a id="L28"></a>}

<a id="L30"></a><span class="comment">// A Load represents any Mach-O load command.</span>
<a id="L31"></a>type Load interface {
    <a id="L32"></a>Raw() []byte;
<a id="L33"></a>}

<a id="L35"></a><span class="comment">// A LoadBytes is the uninterpreted bytes of a Mach-O load command.</span>
<a id="L36"></a>type LoadBytes []byte

<a id="L38"></a>func (b LoadBytes) Raw() []byte { return b }

<a id="L40"></a><span class="comment">// A SegmentHeader is the header for a Mach-O 32-bit or 64-bit load segment command.</span>
<a id="L41"></a>type SegmentHeader struct {
    <a id="L42"></a>Cmd     LoadCmd;
    <a id="L43"></a>Len     uint32;
    <a id="L44"></a>Name    string;
    <a id="L45"></a>Addr    uint64;
    <a id="L46"></a>Memsz   uint64;
    <a id="L47"></a>Offset  uint64;
    <a id="L48"></a>Filesz  uint64;
    <a id="L49"></a>Maxprot uint32;
    <a id="L50"></a>Prot    uint32;
    <a id="L51"></a>Nsect   uint32;
    <a id="L52"></a>Flag    uint32;
<a id="L53"></a>}

<a id="L55"></a><span class="comment">// A Segment represents a Mach-O 32-bit or 64-bit load segment command.</span>
<a id="L56"></a>type Segment struct {
    <a id="L57"></a>LoadBytes;
    <a id="L58"></a>SegmentHeader;

    <a id="L60"></a><span class="comment">// Embed ReaderAt for ReadAt method.</span>
    <a id="L61"></a><span class="comment">// Do not embed SectionReader directly</span>
    <a id="L62"></a><span class="comment">// to avoid having Read and Seek.</span>
    <a id="L63"></a><span class="comment">// If a client wants Read and Seek it must use</span>
    <a id="L64"></a><span class="comment">// Open() to avoid fighting over the seek offset</span>
    <a id="L65"></a><span class="comment">// with other clients.</span>
    <a id="L66"></a>io.ReaderAt;
    <a id="L67"></a>sr  *io.SectionReader;
<a id="L68"></a>}

<a id="L70"></a><span class="comment">// Data reads and returns the contents of the segment.</span>
<a id="L71"></a>func (s *Segment) Data() ([]byte, os.Error) {
    <a id="L72"></a>dat := make([]byte, s.sr.Size());
    <a id="L73"></a>n, err := s.sr.ReadAt(dat, 0);
    <a id="L74"></a>return dat[0:n], err;
<a id="L75"></a>}

<a id="L77"></a><span class="comment">// Open returns a new ReadSeeker reading the segment.</span>
<a id="L78"></a>func (s *Segment) Open() io.ReadSeeker { return io.NewSectionReader(s.sr, 0, 1&lt;&lt;63-1) }

<a id="L80"></a>type SectionHeader struct {
    <a id="L81"></a>Name   string;
    <a id="L82"></a>Seg    string;
    <a id="L83"></a>Addr   uint64;
    <a id="L84"></a>Size   uint64;
    <a id="L85"></a>Offset uint32;
    <a id="L86"></a>Align  uint32;
    <a id="L87"></a>Reloff uint32;
    <a id="L88"></a>Nreloc uint32;
    <a id="L89"></a>Flags  uint32;
<a id="L90"></a>}

<a id="L92"></a>type Section struct {
    <a id="L93"></a>SectionHeader;

    <a id="L95"></a><span class="comment">// Embed ReaderAt for ReadAt method.</span>
    <a id="L96"></a><span class="comment">// Do not embed SectionReader directly</span>
    <a id="L97"></a><span class="comment">// to avoid having Read and Seek.</span>
    <a id="L98"></a><span class="comment">// If a client wants Read and Seek it must use</span>
    <a id="L99"></a><span class="comment">// Open() to avoid fighting over the seek offset</span>
    <a id="L100"></a><span class="comment">// with other clients.</span>
    <a id="L101"></a>io.ReaderAt;
    <a id="L102"></a>sr  *io.SectionReader;
<a id="L103"></a>}

<a id="L105"></a><span class="comment">// Data reads and returns the contents of the Mach-O section.</span>
<a id="L106"></a>func (s *Section) Data() ([]byte, os.Error) {
    <a id="L107"></a>dat := make([]byte, s.sr.Size());
    <a id="L108"></a>n, err := s.sr.ReadAt(dat, 0);
    <a id="L109"></a>return dat[0:n], err;
<a id="L110"></a>}

<a id="L112"></a><span class="comment">// Open returns a new ReadSeeker reading the Mach-O section.</span>
<a id="L113"></a>func (s *Section) Open() io.ReadSeeker { return io.NewSectionReader(s.sr, 0, 1&lt;&lt;63-1) }


<a id="L116"></a><span class="comment">/*</span>
<a id="L117"></a><span class="comment"> * Mach-O reader</span>
<a id="L118"></a><span class="comment"> */</span>

<a id="L120"></a>type FormatError struct {
    <a id="L121"></a>off int64;
    <a id="L122"></a>msg string;
    <a id="L123"></a>val interface{};
<a id="L124"></a>}

<a id="L126"></a>func (e *FormatError) String() string {
    <a id="L127"></a>msg := e.msg;
    <a id="L128"></a>if e.val != nil {
        <a id="L129"></a>msg += fmt.Sprintf(&#34; &#39;%v&#39; &#34;, e.val)
    <a id="L130"></a>}
    <a id="L131"></a>msg += fmt.Sprintf(&#34;in record at byte %#x&#34;, e.off);
    <a id="L132"></a>return msg;
<a id="L133"></a>}

<a id="L135"></a><span class="comment">// Open opens the named file using os.Open and prepares it for use as a Mach-O binary.</span>
<a id="L136"></a>func Open(name string) (*File, os.Error) {
    <a id="L137"></a>f, err := os.Open(name, os.O_RDONLY, 0);
    <a id="L138"></a>if err != nil {
        <a id="L139"></a>return nil, err
    <a id="L140"></a>}
    <a id="L141"></a>ff, err := NewFile(f);
    <a id="L142"></a>if err != nil {
        <a id="L143"></a>f.Close();
        <a id="L144"></a>return nil, err;
    <a id="L145"></a>}
    <a id="L146"></a>ff.closer = f;
    <a id="L147"></a>return ff, nil;
<a id="L148"></a>}

<a id="L150"></a><span class="comment">// Close closes the File.</span>
<a id="L151"></a><span class="comment">// If the File was created using NewFile directly instead of Open,</span>
<a id="L152"></a><span class="comment">// Close has no effect.</span>
<a id="L153"></a>func (f *File) Close() os.Error {
    <a id="L154"></a>var err os.Error;
    <a id="L155"></a>if f.closer != nil {
        <a id="L156"></a>err = f.closer.Close();
        <a id="L157"></a>f.closer = nil;
    <a id="L158"></a>}
    <a id="L159"></a>return err;
<a id="L160"></a>}

<a id="L162"></a><span class="comment">// NewFile creates a new File for acecssing a Mach-O binary in an underlying reader.</span>
<a id="L163"></a><span class="comment">// The Mach-O binary is expected to start at position 0 in the ReaderAt.</span>
<a id="L164"></a>func NewFile(r io.ReaderAt) (*File, os.Error) {
    <a id="L165"></a>f := new(File);
    <a id="L166"></a>sr := io.NewSectionReader(r, 0, 1&lt;&lt;63-1);

    <a id="L168"></a><span class="comment">// Read and decode Mach magic to determine byte order, size.</span>
    <a id="L169"></a><span class="comment">// Magic32 and Magic64 differ only in the bottom bit.</span>
    <a id="L170"></a>var ident [4]uint8;
    <a id="L171"></a>if _, err := r.ReadAt(&amp;ident, 0); err != nil {
        <a id="L172"></a>return nil, err
    <a id="L173"></a>}
    <a id="L174"></a>be := binary.BigEndian.Uint32(&amp;ident);
    <a id="L175"></a>le := binary.LittleEndian.Uint32(&amp;ident);
    <a id="L176"></a>switch Magic32 &amp;^ 1 {
    <a id="L177"></a>case be &amp;^ 1:
        <a id="L178"></a>f.ByteOrder = binary.BigEndian;
        <a id="L179"></a>f.Magic = be;
    <a id="L180"></a>case le &amp;^ 1:
        <a id="L181"></a>f.ByteOrder = binary.LittleEndian;
        <a id="L182"></a>f.Magic = le;
    <a id="L183"></a>}

    <a id="L185"></a><span class="comment">// Read entire file header.</span>
    <a id="L186"></a>if err := binary.Read(sr, f.ByteOrder, &amp;f.FileHeader); err != nil {
        <a id="L187"></a>return nil, err
    <a id="L188"></a>}

    <a id="L190"></a><span class="comment">// Then load commands.</span>
    <a id="L191"></a>offset := int64(fileHeaderSize32);
    <a id="L192"></a>if f.Magic == Magic64 {
        <a id="L193"></a>offset = fileHeaderSize64
    <a id="L194"></a>}
    <a id="L195"></a>dat := make([]byte, f.Cmdsz);
    <a id="L196"></a>if _, err := r.ReadAt(dat, offset); err != nil {
        <a id="L197"></a>return nil, err
    <a id="L198"></a>}
    <a id="L199"></a>f.Loads = make([]Load, f.Ncmd);
    <a id="L200"></a>bo := f.ByteOrder;
    <a id="L201"></a>for i := range f.Loads {
        <a id="L202"></a><span class="comment">// Each load command begins with uint32 command and length.</span>
        <a id="L203"></a>if len(dat) &lt; 8 {
            <a id="L204"></a>return nil, &amp;FormatError{offset, &#34;command block too small&#34;, nil}
        <a id="L205"></a>}
        <a id="L206"></a>cmd, siz := LoadCmd(bo.Uint32(dat[0:4])), bo.Uint32(dat[4:8]);
        <a id="L207"></a>if siz &lt; 8 || siz &gt; uint32(len(dat)) {
            <a id="L208"></a>return nil, &amp;FormatError{offset, &#34;invalid command block size&#34;, nil}
        <a id="L209"></a>}
        <a id="L210"></a>var cmddat []byte;
        <a id="L211"></a>cmddat, dat = dat[0:siz], dat[siz:len(dat)];
        <a id="L212"></a>offset += int64(siz);
        <a id="L213"></a>var s *Segment;
        <a id="L214"></a>switch cmd {
        <a id="L215"></a>default:
            <a id="L216"></a>f.Loads[i] = LoadBytes(cmddat)

        <a id="L218"></a>case LoadCmdSegment:
            <a id="L219"></a>var seg32 Segment32;
            <a id="L220"></a>b := bytes.NewBuffer(cmddat);
            <a id="L221"></a>if err := binary.Read(b, bo, &amp;seg32); err != nil {
                <a id="L222"></a>return nil, err
            <a id="L223"></a>}
            <a id="L224"></a>s = new(Segment);
            <a id="L225"></a>s.LoadBytes = cmddat;
            <a id="L226"></a>s.Cmd = cmd;
            <a id="L227"></a>s.Len = siz;
            <a id="L228"></a>s.Name = cstring(&amp;seg32.Name);
            <a id="L229"></a>s.Addr = uint64(seg32.Addr);
            <a id="L230"></a>s.Memsz = uint64(seg32.Memsz);
            <a id="L231"></a>s.Offset = uint64(seg32.Offset);
            <a id="L232"></a>s.Filesz = uint64(seg32.Filesz);
            <a id="L233"></a>s.Maxprot = seg32.Maxprot;
            <a id="L234"></a>s.Prot = seg32.Prot;
            <a id="L235"></a>s.Nsect = seg32.Nsect;
            <a id="L236"></a>s.Flag = seg32.Flag;
            <a id="L237"></a>f.Loads[i] = s;
            <a id="L238"></a>for i := 0; i &lt; int(s.Nsect); i++ {
                <a id="L239"></a>var sh32 Section32;
                <a id="L240"></a>if err := binary.Read(b, bo, &amp;sh32); err != nil {
                    <a id="L241"></a>return nil, err
                <a id="L242"></a>}
                <a id="L243"></a>sh := new(Section);
                <a id="L244"></a>sh.Name = cstring(&amp;sh32.Name);
                <a id="L245"></a>sh.Seg = cstring(&amp;sh32.Seg);
                <a id="L246"></a>sh.Addr = uint64(sh32.Addr);
                <a id="L247"></a>sh.Size = uint64(sh32.Size);
                <a id="L248"></a>sh.Offset = sh32.Offset;
                <a id="L249"></a>sh.Align = sh32.Align;
                <a id="L250"></a>sh.Reloff = sh32.Reloff;
                <a id="L251"></a>sh.Nreloc = sh32.Nreloc;
                <a id="L252"></a>sh.Flags = sh32.Flags;
                <a id="L253"></a>f.pushSection(sh, r);
            <a id="L254"></a>}

        <a id="L256"></a>case LoadCmdSegment64:
            <a id="L257"></a>var seg64 Segment64;
            <a id="L258"></a>b := bytes.NewBuffer(cmddat);
            <a id="L259"></a>if err := binary.Read(b, bo, &amp;seg64); err != nil {
                <a id="L260"></a>return nil, err
            <a id="L261"></a>}
            <a id="L262"></a>s = new(Segment);
            <a id="L263"></a>s.LoadBytes = cmddat;
            <a id="L264"></a>s.Cmd = cmd;
            <a id="L265"></a>s.Len = siz;
            <a id="L266"></a>s.Name = cstring(&amp;seg64.Name);
            <a id="L267"></a>s.Addr = seg64.Addr;
            <a id="L268"></a>s.Memsz = seg64.Memsz;
            <a id="L269"></a>s.Offset = seg64.Offset;
            <a id="L270"></a>s.Filesz = seg64.Filesz;
            <a id="L271"></a>s.Maxprot = seg64.Maxprot;
            <a id="L272"></a>s.Prot = seg64.Prot;
            <a id="L273"></a>s.Nsect = seg64.Nsect;
            <a id="L274"></a>s.Flag = seg64.Flag;
            <a id="L275"></a>f.Loads[i] = s;
            <a id="L276"></a>for i := 0; i &lt; int(s.Nsect); i++ {
                <a id="L277"></a>var sh64 Section64;
                <a id="L278"></a>if err := binary.Read(b, bo, &amp;sh64); err != nil {
                    <a id="L279"></a>return nil, err
                <a id="L280"></a>}
                <a id="L281"></a>sh := new(Section);
                <a id="L282"></a>sh.Name = cstring(&amp;sh64.Name);
                <a id="L283"></a>sh.Seg = cstring(&amp;sh64.Seg);
                <a id="L284"></a>sh.Addr = sh64.Addr;
                <a id="L285"></a>sh.Size = sh64.Size;
                <a id="L286"></a>sh.Offset = sh64.Offset;
                <a id="L287"></a>sh.Align = sh64.Align;
                <a id="L288"></a>sh.Reloff = sh64.Reloff;
                <a id="L289"></a>sh.Nreloc = sh64.Nreloc;
                <a id="L290"></a>sh.Flags = sh64.Flags;
                <a id="L291"></a>f.pushSection(sh, r);
            <a id="L292"></a>}
        <a id="L293"></a>}
        <a id="L294"></a>if s != nil {
            <a id="L295"></a>s.sr = io.NewSectionReader(r, int64(s.Offset), int64(s.Filesz));
            <a id="L296"></a>s.ReaderAt = s.sr;
        <a id="L297"></a>}
    <a id="L298"></a>}
    <a id="L299"></a>return f, nil;
<a id="L300"></a>}

<a id="L302"></a>func (f *File) pushSection(sh *Section, r io.ReaderAt) {
    <a id="L303"></a>n := len(f.Sections);
    <a id="L304"></a>if n &gt;= cap(f.Sections) {
        <a id="L305"></a>m := (n + 1) * 2;
        <a id="L306"></a>new := make([]*Section, n, m);
        <a id="L307"></a>for i, sh := range f.Sections {
            <a id="L308"></a>new[i] = sh
        <a id="L309"></a>}
        <a id="L310"></a>f.Sections = new;
    <a id="L311"></a>}
    <a id="L312"></a>f.Sections = f.Sections[0 : n+1];
    <a id="L313"></a>f.Sections[n] = sh;
    <a id="L314"></a>sh.sr = io.NewSectionReader(r, int64(sh.Offset), int64(sh.Size));
    <a id="L315"></a>sh.ReaderAt = sh.sr;
<a id="L316"></a>}

<a id="L318"></a>func cstring(b []byte) string {
    <a id="L319"></a>var i int;
    <a id="L320"></a>for i = 0; i &lt; len(b) &amp;&amp; b[i] != 0; i++ {
    <a id="L321"></a>}
    <a id="L322"></a>return string(b[0:i]);
<a id="L323"></a>}

<a id="L325"></a><span class="comment">// Segment returns the first Segment with the given name, or nil if no such segment exists.</span>
<a id="L326"></a>func (f *File) Segment(name string) *Segment {
    <a id="L327"></a>for _, l := range f.Loads {
        <a id="L328"></a>if s, ok := l.(*Segment); ok &amp;&amp; s.Name == name {
            <a id="L329"></a>return s
        <a id="L330"></a>}
    <a id="L331"></a>}
    <a id="L332"></a>return nil;
<a id="L333"></a>}

<a id="L335"></a><span class="comment">// Section returns the first section with the given name, or nil if no such</span>
<a id="L336"></a><span class="comment">// section exists.</span>
<a id="L337"></a>func (f *File) Section(name string) *Section {
    <a id="L338"></a>for _, s := range f.Sections {
        <a id="L339"></a>if s.Name == name {
            <a id="L340"></a>return s
        <a id="L341"></a>}
    <a id="L342"></a>}
    <a id="L343"></a>return nil;
<a id="L344"></a>}

<a id="L346"></a><span class="comment">// DWARF returns the DWARF debug information for the Mach-O file.</span>
<a id="L347"></a>func (f *File) DWARF() (*dwarf.Data, os.Error) {
    <a id="L348"></a><span class="comment">// There are many other DWARF sections, but these</span>
    <a id="L349"></a><span class="comment">// are the required ones, and the debug/dwarf package</span>
    <a id="L350"></a><span class="comment">// does not use the others, so don&#39;t bother loading them.</span>
    <a id="L351"></a>var names = [...]string{&#34;abbrev&#34;, &#34;info&#34;, &#34;str&#34;};
    <a id="L352"></a>var dat [len(names)][]byte;
    <a id="L353"></a>for i, name := range names {
        <a id="L354"></a>name = &#34;__debug_&#34; + name;
        <a id="L355"></a>s := f.Section(name);
        <a id="L356"></a>if s == nil {
            <a id="L357"></a>return nil, os.NewError(&#34;missing Mach-O section &#34; + name)
        <a id="L358"></a>}
        <a id="L359"></a>b, err := s.Data();
        <a id="L360"></a>if err != nil &amp;&amp; uint64(len(b)) &lt; s.Size {
            <a id="L361"></a>return nil, err
        <a id="L362"></a>}
        <a id="L363"></a>dat[i] = b;
    <a id="L364"></a>}

    <a id="L366"></a>abbrev, info, str := dat[0], dat[1], dat[2];
    <a id="L367"></a>return dwarf.New(abbrev, nil, nil, info, nil, nil, nil, str);
<a id="L368"></a>}
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
