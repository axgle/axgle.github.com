<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/elf/file.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/debug/elf/file.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Package elf implements access to ELF object files.</span>
<a id="L6"></a>package elf

<a id="L8"></a>import (
    <a id="L9"></a>&#34;bytes&#34;;
    <a id="L10"></a>&#34;debug/dwarf&#34;;
    <a id="L11"></a>&#34;encoding/binary&#34;;
    <a id="L12"></a>&#34;fmt&#34;;
    <a id="L13"></a>&#34;io&#34;;
    <a id="L14"></a>&#34;os&#34;;
<a id="L15"></a>)

<a id="L17"></a><span class="comment">// TODO: error reporting detail</span>

<a id="L19"></a><span class="comment">/*</span>
<a id="L20"></a><span class="comment"> * Internal ELF representation</span>
<a id="L21"></a><span class="comment"> */</span>

<a id="L23"></a><span class="comment">// A FileHeader represents an ELF file header.</span>
<a id="L24"></a>type FileHeader struct {
    <a id="L25"></a>Class      Class;
    <a id="L26"></a>Data       Data;
    <a id="L27"></a>Version    Version;
    <a id="L28"></a>OSABI      OSABI;
    <a id="L29"></a>ABIVersion uint8;
    <a id="L30"></a>ByteOrder  binary.ByteOrder;
    <a id="L31"></a>Type       Type;
    <a id="L32"></a>Machine    Machine;
<a id="L33"></a>}

<a id="L35"></a><span class="comment">// A File represents an open ELF file.</span>
<a id="L36"></a>type File struct {
    <a id="L37"></a>FileHeader;
    <a id="L38"></a>Sections []*Section;
    <a id="L39"></a>Progs    []*Prog;
    <a id="L40"></a>closer   io.Closer;
<a id="L41"></a>}

<a id="L43"></a><span class="comment">// A SectionHeader represents a single ELF section header.</span>
<a id="L44"></a>type SectionHeader struct {
    <a id="L45"></a>Name      string;
    <a id="L46"></a>Type      SectionType;
    <a id="L47"></a>Flags     SectionFlag;
    <a id="L48"></a>Addr      uint64;
    <a id="L49"></a>Offset    uint64;
    <a id="L50"></a>Size      uint64;
    <a id="L51"></a>Link      uint32;
    <a id="L52"></a>Info      uint32;
    <a id="L53"></a>Addralign uint64;
    <a id="L54"></a>Entsize   uint64;
<a id="L55"></a>}

<a id="L57"></a><span class="comment">// A Section represents a single section in an ELF file.</span>
<a id="L58"></a>type Section struct {
    <a id="L59"></a>SectionHeader;

    <a id="L61"></a><span class="comment">// Embed ReaderAt for ReadAt method.</span>
    <a id="L62"></a><span class="comment">// Do not embed SectionReader directly</span>
    <a id="L63"></a><span class="comment">// to avoid having Read and Seek.</span>
    <a id="L64"></a><span class="comment">// If a client wants Read and Seek it must use</span>
    <a id="L65"></a><span class="comment">// Open() to avoid fighting over the seek offset</span>
    <a id="L66"></a><span class="comment">// with other clients.</span>
    <a id="L67"></a>io.ReaderAt;
    <a id="L68"></a>sr  *io.SectionReader;
<a id="L69"></a>}

<a id="L71"></a><span class="comment">// Data reads and returns the contents of the ELF section.</span>
<a id="L72"></a>func (s *Section) Data() ([]byte, os.Error) {
    <a id="L73"></a>dat := make([]byte, s.sr.Size());
    <a id="L74"></a>n, err := s.sr.ReadAt(dat, 0);
    <a id="L75"></a>return dat[0:n], err;
<a id="L76"></a>}

<a id="L78"></a><span class="comment">// Open returns a new ReadSeeker reading the ELF section.</span>
<a id="L79"></a>func (s *Section) Open() io.ReadSeeker { return io.NewSectionReader(s.sr, 0, 1&lt;&lt;63-1) }

<a id="L81"></a><span class="comment">// A ProgHeader represents a single ELF program header.</span>
<a id="L82"></a>type ProgHeader struct {
    <a id="L83"></a>Type   ProgType;
    <a id="L84"></a>Flags  ProgFlag;
    <a id="L85"></a>Vaddr  uint64;
    <a id="L86"></a>Paddr  uint64;
    <a id="L87"></a>Filesz uint64;
    <a id="L88"></a>Memsz  uint64;
    <a id="L89"></a>Align  uint64;
<a id="L90"></a>}

<a id="L92"></a><span class="comment">// A Prog represents a single ELF program header in an ELF binary.</span>
<a id="L93"></a>type Prog struct {
    <a id="L94"></a>ProgHeader;

    <a id="L96"></a><span class="comment">// Embed ReaderAt for ReadAt method.</span>
    <a id="L97"></a><span class="comment">// Do not embed SectionReader directly</span>
    <a id="L98"></a><span class="comment">// to avoid having Read and Seek.</span>
    <a id="L99"></a><span class="comment">// If a client wants Read and Seek it must use</span>
    <a id="L100"></a><span class="comment">// Open() to avoid fighting over the seek offset</span>
    <a id="L101"></a><span class="comment">// with other clients.</span>
    <a id="L102"></a>io.ReaderAt;
    <a id="L103"></a>sr  *io.SectionReader;
<a id="L104"></a>}

<a id="L106"></a><span class="comment">// Open returns a new ReadSeeker reading the ELF program body.</span>
<a id="L107"></a>func (p *Prog) Open() io.ReadSeeker { return io.NewSectionReader(p.sr, 0, 1&lt;&lt;63-1) }

<a id="L109"></a><span class="comment">// A Symbol represents an entry in an ELF symbol table section.</span>
<a id="L110"></a>type Symbol struct {
    <a id="L111"></a>Name        uint32;
    <a id="L112"></a>Info, Other byte;
    <a id="L113"></a>Section     uint32;
    <a id="L114"></a>Value, Size uint64;
<a id="L115"></a>}

<a id="L117"></a><span class="comment">/*</span>
<a id="L118"></a><span class="comment"> * ELF reader</span>
<a id="L119"></a><span class="comment"> */</span>

<a id="L121"></a>type FormatError struct {
    <a id="L122"></a>off int64;
    <a id="L123"></a>msg string;
    <a id="L124"></a>val interface{};
<a id="L125"></a>}

<a id="L127"></a>func (e *FormatError) String() string {
    <a id="L128"></a>msg := e.msg;
    <a id="L129"></a>if e.val != nil {
        <a id="L130"></a>msg += fmt.Sprintf(&#34; &#39;%v&#39; &#34;, e.val)
    <a id="L131"></a>}
    <a id="L132"></a>msg += fmt.Sprintf(&#34;in record at byte %#x&#34;, e.off);
    <a id="L133"></a>return msg;
<a id="L134"></a>}

<a id="L136"></a><span class="comment">// Open opens the named file using os.Open and prepares it for use as an ELF binary.</span>
<a id="L137"></a>func Open(name string) (*File, os.Error) {
    <a id="L138"></a>f, err := os.Open(name, os.O_RDONLY, 0);
    <a id="L139"></a>if err != nil {
        <a id="L140"></a>return nil, err
    <a id="L141"></a>}
    <a id="L142"></a>ff, err := NewFile(f);
    <a id="L143"></a>if err != nil {
        <a id="L144"></a>f.Close();
        <a id="L145"></a>return nil, err;
    <a id="L146"></a>}
    <a id="L147"></a>ff.closer = f;
    <a id="L148"></a>return ff, nil;
<a id="L149"></a>}

<a id="L151"></a><span class="comment">// Close closes the File.</span>
<a id="L152"></a><span class="comment">// If the File was created using NewFile directly instead of Open,</span>
<a id="L153"></a><span class="comment">// Close has no effect.</span>
<a id="L154"></a>func (f *File) Close() os.Error {
    <a id="L155"></a>var err os.Error;
    <a id="L156"></a>if f.closer != nil {
        <a id="L157"></a>err = f.closer.Close();
        <a id="L158"></a>f.closer = nil;
    <a id="L159"></a>}
    <a id="L160"></a>return err;
<a id="L161"></a>}

<a id="L163"></a><span class="comment">// NewFile creates a new File for acecssing an ELF binary in an underlying reader.</span>
<a id="L164"></a><span class="comment">// The ELF binary is expected to start at position 0 in the ReaderAt.</span>
<a id="L165"></a>func NewFile(r io.ReaderAt) (*File, os.Error) {
    <a id="L166"></a>sr := io.NewSectionReader(r, 0, 1&lt;&lt;63-1);
    <a id="L167"></a><span class="comment">// Read and decode ELF identifier</span>
    <a id="L168"></a>var ident [16]uint8;
    <a id="L169"></a>if _, err := r.ReadAt(&amp;ident, 0); err != nil {
        <a id="L170"></a>return nil, err
    <a id="L171"></a>}
    <a id="L172"></a>if ident[0] != &#39;\x7f&#39; || ident[1] != &#39;E&#39; || ident[2] != &#39;L&#39; || ident[3] != &#39;F&#39; {
        <a id="L173"></a>return nil, &amp;FormatError{0, &#34;bad magic number&#34;, ident[0:4]}
    <a id="L174"></a>}

    <a id="L176"></a>f := new(File);
    <a id="L177"></a>f.Class = Class(ident[EI_CLASS]);
    <a id="L178"></a>switch f.Class {
    <a id="L179"></a>case ELFCLASS32:
    <a id="L180"></a>case ELFCLASS64:
        <a id="L181"></a><span class="comment">// ok</span>
    <a id="L182"></a>default:
        <a id="L183"></a>return nil, &amp;FormatError{0, &#34;unknown ELF class&#34;, f.Class}
    <a id="L184"></a>}

    <a id="L186"></a>f.Data = Data(ident[EI_DATA]);
    <a id="L187"></a>switch f.Data {
    <a id="L188"></a>case ELFDATA2LSB:
        <a id="L189"></a>f.ByteOrder = binary.LittleEndian
    <a id="L190"></a>case ELFDATA2MSB:
        <a id="L191"></a>f.ByteOrder = binary.BigEndian
    <a id="L192"></a>default:
        <a id="L193"></a>return nil, &amp;FormatError{0, &#34;unknown ELF data encoding&#34;, f.Data}
    <a id="L194"></a>}

    <a id="L196"></a>f.Version = Version(ident[EI_VERSION]);
    <a id="L197"></a>if f.Version != EV_CURRENT {
        <a id="L198"></a>return nil, &amp;FormatError{0, &#34;unknown ELF version&#34;, f.Version}
    <a id="L199"></a>}

    <a id="L201"></a>f.OSABI = OSABI(ident[EI_OSABI]);
    <a id="L202"></a>f.ABIVersion = ident[EI_ABIVERSION];

    <a id="L204"></a><span class="comment">// Read ELF file header</span>
    <a id="L205"></a>var shoff int64;
    <a id="L206"></a>var shentsize, shnum, shstrndx int;
    <a id="L207"></a>shstrndx = -1;
    <a id="L208"></a>switch f.Class {
    <a id="L209"></a>case ELFCLASS32:
        <a id="L210"></a>hdr := new(Header32);
        <a id="L211"></a>sr.Seek(0, 0);
        <a id="L212"></a>if err := binary.Read(sr, f.ByteOrder, hdr); err != nil {
            <a id="L213"></a>return nil, err
        <a id="L214"></a>}
        <a id="L215"></a>f.Type = Type(hdr.Type);
        <a id="L216"></a>f.Machine = Machine(hdr.Machine);
        <a id="L217"></a>if v := Version(hdr.Version); v != f.Version {
            <a id="L218"></a>return nil, &amp;FormatError{0, &#34;mismatched ELF version&#34;, v}
        <a id="L219"></a>}
        <a id="L220"></a>shoff = int64(hdr.Shoff);
        <a id="L221"></a>shentsize = int(hdr.Shentsize);
        <a id="L222"></a>shnum = int(hdr.Shnum);
        <a id="L223"></a>shstrndx = int(hdr.Shstrndx);
    <a id="L224"></a>case ELFCLASS64:
        <a id="L225"></a>hdr := new(Header64);
        <a id="L226"></a>sr.Seek(0, 0);
        <a id="L227"></a>if err := binary.Read(sr, f.ByteOrder, hdr); err != nil {
            <a id="L228"></a>return nil, err
        <a id="L229"></a>}
        <a id="L230"></a>f.Type = Type(hdr.Type);
        <a id="L231"></a>f.Machine = Machine(hdr.Machine);
        <a id="L232"></a>if v := Version(hdr.Version); v != f.Version {
            <a id="L233"></a>return nil, &amp;FormatError{0, &#34;mismatched ELF version&#34;, v}
        <a id="L234"></a>}
        <a id="L235"></a>shoff = int64(hdr.Shoff);
        <a id="L236"></a>shentsize = int(hdr.Shentsize);
        <a id="L237"></a>shnum = int(hdr.Shnum);
        <a id="L238"></a>shstrndx = int(hdr.Shstrndx);
    <a id="L239"></a>}
    <a id="L240"></a>if shstrndx &lt; 0 || shstrndx &gt;= shnum {
        <a id="L241"></a>return nil, &amp;FormatError{0, &#34;invalid ELF shstrndx&#34;, shstrndx}
    <a id="L242"></a>}

    <a id="L244"></a><span class="comment">// Read program headers</span>
    <a id="L245"></a><span class="comment">// TODO</span>

    <a id="L247"></a><span class="comment">// Read section headers</span>
    <a id="L248"></a>f.Sections = make([]*Section, shnum);
    <a id="L249"></a>names := make([]uint32, shnum);
    <a id="L250"></a>for i := 0; i &lt; shnum; i++ {
        <a id="L251"></a>off := shoff + int64(i)*int64(shentsize);
        <a id="L252"></a>sr.Seek(off, 0);
        <a id="L253"></a>s := new(Section);
        <a id="L254"></a>switch f.Class {
        <a id="L255"></a>case ELFCLASS32:
            <a id="L256"></a>sh := new(Section32);
            <a id="L257"></a>if err := binary.Read(sr, f.ByteOrder, sh); err != nil {
                <a id="L258"></a>return nil, err
            <a id="L259"></a>}
            <a id="L260"></a>names[i] = sh.Name;
            <a id="L261"></a>s.SectionHeader = SectionHeader{
                <a id="L262"></a>Type: SectionType(sh.Type),
                <a id="L263"></a>Flags: SectionFlag(sh.Flags),
                <a id="L264"></a>Addr: uint64(sh.Addr),
                <a id="L265"></a>Offset: uint64(sh.Off),
                <a id="L266"></a>Size: uint64(sh.Size),
                <a id="L267"></a>Link: uint32(sh.Link),
                <a id="L268"></a>Info: uint32(sh.Info),
                <a id="L269"></a>Addralign: uint64(sh.Addralign),
                <a id="L270"></a>Entsize: uint64(sh.Entsize),
            <a id="L271"></a>};
        <a id="L272"></a>case ELFCLASS64:
            <a id="L273"></a>sh := new(Section64);
            <a id="L274"></a>if err := binary.Read(sr, f.ByteOrder, sh); err != nil {
                <a id="L275"></a>return nil, err
            <a id="L276"></a>}
            <a id="L277"></a>names[i] = sh.Name;
            <a id="L278"></a>s.SectionHeader = SectionHeader{
                <a id="L279"></a>Type: SectionType(sh.Type),
                <a id="L280"></a>Flags: SectionFlag(sh.Flags),
                <a id="L281"></a>Offset: uint64(sh.Off),
                <a id="L282"></a>Size: uint64(sh.Size),
                <a id="L283"></a>Addr: uint64(sh.Addr),
                <a id="L284"></a>Link: uint32(sh.Link),
                <a id="L285"></a>Info: uint32(sh.Info),
                <a id="L286"></a>Addralign: uint64(sh.Addralign),
                <a id="L287"></a>Entsize: uint64(sh.Entsize),
            <a id="L288"></a>};
        <a id="L289"></a>}
        <a id="L290"></a>s.sr = io.NewSectionReader(r, int64(s.Offset), int64(s.Size));
        <a id="L291"></a>s.ReaderAt = s.sr;
        <a id="L292"></a>f.Sections[i] = s;
    <a id="L293"></a>}

    <a id="L295"></a><span class="comment">// Load section header string table.</span>
    <a id="L296"></a>s := f.Sections[shstrndx];
    <a id="L297"></a>shstrtab := make([]byte, s.Size);
    <a id="L298"></a>if _, err := r.ReadAt(shstrtab, int64(s.Offset)); err != nil {
        <a id="L299"></a>return nil, err
    <a id="L300"></a>}
    <a id="L301"></a>for i, s := range f.Sections {
        <a id="L302"></a>var ok bool;
        <a id="L303"></a>s.Name, ok = getString(shstrtab, int(names[i]));
        <a id="L304"></a>if !ok {
            <a id="L305"></a>return nil, &amp;FormatError{shoff + int64(i*shentsize), &#34;bad section name index&#34;, names[i]}
        <a id="L306"></a>}
    <a id="L307"></a>}

    <a id="L309"></a>return f, nil;
<a id="L310"></a>}

<a id="L312"></a>func (f *File) getSymbols() ([]Symbol, os.Error) {
    <a id="L313"></a>switch f.Class {
    <a id="L314"></a>case ELFCLASS64:
        <a id="L315"></a>return f.getSymbols64()
    <a id="L316"></a>}

    <a id="L318"></a>return nil, os.ErrorString(&#34;not implemented&#34;);
<a id="L319"></a>}

<a id="L321"></a><span class="comment">// GetSymbols returns a slice of Symbols from parsing the symbol table.</span>
<a id="L322"></a>func (f *File) getSymbols64() ([]Symbol, os.Error) {
    <a id="L323"></a>var symtabSection *Section;
    <a id="L324"></a>for _, section := range f.Sections {
        <a id="L325"></a>if section.Type == SHT_SYMTAB {
            <a id="L326"></a>symtabSection = section;
            <a id="L327"></a>break;
        <a id="L328"></a>}
    <a id="L329"></a>}

    <a id="L331"></a>if symtabSection == nil {
        <a id="L332"></a>return nil, os.ErrorString(&#34;no symbol section&#34;)
    <a id="L333"></a>}

    <a id="L335"></a>data, err := symtabSection.Data();
    <a id="L336"></a>if err != nil {
        <a id="L337"></a>return nil, os.ErrorString(&#34;cannot load symbol section&#34;)
    <a id="L338"></a>}
    <a id="L339"></a>symtab := bytes.NewBuffer(data);
    <a id="L340"></a>if symtab.Len()%Sym64Size != 0 {
        <a id="L341"></a>return nil, os.ErrorString(&#34;length of symbol section is not a multiple of Sym64Size&#34;)
    <a id="L342"></a>}

    <a id="L344"></a><span class="comment">// The first entry is all zeros.</span>
    <a id="L345"></a>var skip [Sym64Size]byte;
    <a id="L346"></a>symtab.Read(skip[0:len(skip)]);

    <a id="L348"></a>symbols := make([]Symbol, symtab.Len()/Sym64Size);

    <a id="L350"></a>i := 0;
    <a id="L351"></a>var sym Sym64;
    <a id="L352"></a>for symtab.Len() &gt; 0 {
        <a id="L353"></a>binary.Read(symtab, f.ByteOrder, &amp;sym);
        <a id="L354"></a>symbols[i].Name = sym.Name;
        <a id="L355"></a>symbols[i].Info = sym.Info;
        <a id="L356"></a>symbols[i].Other = sym.Other;
        <a id="L357"></a>symbols[i].Section = uint32(sym.Shndx);
        <a id="L358"></a>symbols[i].Value = sym.Value;
        <a id="L359"></a>symbols[i].Size = sym.Size;
        <a id="L360"></a>i++;
    <a id="L361"></a>}

    <a id="L363"></a>return symbols, nil;
<a id="L364"></a>}

<a id="L366"></a><span class="comment">// getString extracts a string from an ELF string table.</span>
<a id="L367"></a>func getString(section []byte, start int) (string, bool) {
    <a id="L368"></a>if start &lt; 0 || start &gt;= len(section) {
        <a id="L369"></a>return &#34;&#34;, false
    <a id="L370"></a>}

    <a id="L372"></a>for end := start; end &lt; len(section); end++ {
        <a id="L373"></a>if section[end] == 0 {
            <a id="L374"></a>return string(section[start:end]), true
        <a id="L375"></a>}
    <a id="L376"></a>}
    <a id="L377"></a>return &#34;&#34;, false;
<a id="L378"></a>}

<a id="L380"></a><span class="comment">// Section returns a section with the given name, or nil if no such</span>
<a id="L381"></a><span class="comment">// section exists.</span>
<a id="L382"></a>func (f *File) Section(name string) *Section {
    <a id="L383"></a>for _, s := range f.Sections {
        <a id="L384"></a>if s.Name == name {
            <a id="L385"></a>return s
        <a id="L386"></a>}
    <a id="L387"></a>}
    <a id="L388"></a>return nil;
<a id="L389"></a>}

<a id="L391"></a><span class="comment">// applyRelocations applies relocations to dst. rels is a relocations section</span>
<a id="L392"></a><span class="comment">// in RELA format.</span>
<a id="L393"></a>func (f *File) applyRelocations(dst []byte, rels []byte) os.Error {
    <a id="L394"></a>if f.Class == ELFCLASS64 &amp;&amp; f.Machine == EM_X86_64 {
        <a id="L395"></a>return f.applyRelocationsAMD64(dst, rels)
    <a id="L396"></a>}

    <a id="L398"></a>return os.ErrorString(&#34;not implemented&#34;);
<a id="L399"></a>}

<a id="L401"></a>func (f *File) applyRelocationsAMD64(dst []byte, rels []byte) os.Error {
    <a id="L402"></a>if len(rels)%Sym64Size != 0 {
        <a id="L403"></a>return os.ErrorString(&#34;length of relocation section is not a multiple of Sym64Size&#34;)
    <a id="L404"></a>}

    <a id="L406"></a>symbols, err := f.getSymbols();
    <a id="L407"></a>if err != nil {
        <a id="L408"></a>return err
    <a id="L409"></a>}

    <a id="L411"></a>b := bytes.NewBuffer(rels);
    <a id="L412"></a>var rela Rela64;

    <a id="L414"></a>for b.Len() &gt; 0 {
        <a id="L415"></a>binary.Read(b, f.ByteOrder, &amp;rela);
        <a id="L416"></a>symNo := rela.Info &gt;&gt; 32;
        <a id="L417"></a>t := R_X86_64(rela.Info &amp; 0xffff);

        <a id="L419"></a>if symNo &gt;= uint64(len(symbols)) {
            <a id="L420"></a>continue
        <a id="L421"></a>}
        <a id="L422"></a>sym := &amp;symbols[symNo];
        <a id="L423"></a>if SymType(sym.Info&amp;0xf) != STT_SECTION {
            <a id="L424"></a><span class="comment">// We don&#39;t handle non-section relocations for now.</span>
            <a id="L425"></a>continue
        <a id="L426"></a>}

        <a id="L428"></a>switch t {
        <a id="L429"></a>case R_X86_64_64:
            <a id="L430"></a>if rela.Off+8 &gt;= uint64(len(dst)) || rela.Addend &lt; 0 {
                <a id="L431"></a>continue
            <a id="L432"></a>}
            <a id="L433"></a>f.ByteOrder.PutUint64(dst[rela.Off:rela.Off+8], uint64(rela.Addend));
        <a id="L434"></a>case R_X86_64_32:
            <a id="L435"></a>if rela.Off+4 &gt;= uint64(len(dst)) || rela.Addend &lt; 0 {
                <a id="L436"></a>continue
            <a id="L437"></a>}
            <a id="L438"></a>f.ByteOrder.PutUint32(dst[rela.Off:rela.Off+4], uint32(rela.Addend));
        <a id="L439"></a>}
    <a id="L440"></a>}

    <a id="L442"></a>return nil;
<a id="L443"></a>}

<a id="L445"></a>func (f *File) DWARF() (*dwarf.Data, os.Error) {
    <a id="L446"></a><span class="comment">// There are many other DWARF sections, but these</span>
    <a id="L447"></a><span class="comment">// are the required ones, and the debug/dwarf package</span>
    <a id="L448"></a><span class="comment">// does not use the others, so don&#39;t bother loading them.</span>
    <a id="L449"></a>var names = [...]string{&#34;abbrev&#34;, &#34;info&#34;, &#34;str&#34;};
    <a id="L450"></a>var dat [len(names)][]byte;
    <a id="L451"></a>for i, name := range names {
        <a id="L452"></a>name = &#34;.debug_&#34; + name;
        <a id="L453"></a>s := f.Section(name);
        <a id="L454"></a>if s == nil {
            <a id="L455"></a>continue
        <a id="L456"></a>}
        <a id="L457"></a>b, err := s.Data();
        <a id="L458"></a>if err != nil &amp;&amp; uint64(len(b)) &lt; s.Size {
            <a id="L459"></a>return nil, err
        <a id="L460"></a>}
        <a id="L461"></a>dat[i] = b;
    <a id="L462"></a>}

    <a id="L464"></a><span class="comment">// If there&#39;s a relocation table for .debug_info, we have to process it</span>
    <a id="L465"></a><span class="comment">// now otherwise the data in .debug_info is invalid for x86-64 objects.</span>
    <a id="L466"></a>rela := f.Section(&#34;.rela.debug_info&#34;);
    <a id="L467"></a>if rela != nil &amp;&amp; rela.Type == SHT_RELA &amp;&amp; f.Machine == EM_X86_64 {
        <a id="L468"></a>data, err := rela.Data();
        <a id="L469"></a>if err != nil {
            <a id="L470"></a>return nil, err
        <a id="L471"></a>}
        <a id="L472"></a>err = f.applyRelocations(dat[1], data);
        <a id="L473"></a>if err != nil {
            <a id="L474"></a>return nil, err
        <a id="L475"></a>}
    <a id="L476"></a>}

    <a id="L478"></a>abbrev, info, str := dat[0], dat[1], dat[2];
    <a id="L479"></a>return dwarf.New(abbrev, nil, nil, info, nil, nil, nil, str);
<a id="L480"></a>}
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
