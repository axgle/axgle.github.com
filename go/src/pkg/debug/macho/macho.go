<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/macho/macho.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/debug/macho/macho.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Mach-O header data structures</span>
<a id="L6"></a><span class="comment">// http://developer.apple.com/mac/library/documentation/DeveloperTools/Conceptual/MachORuntime/Reference/reference.html</span>

<a id="L8"></a>package macho

<a id="L10"></a>import &#34;strconv&#34;

<a id="L12"></a><span class="comment">// A FileHeader represents a Mach-O file header.</span>
<a id="L13"></a>type FileHeader struct {
    <a id="L14"></a>Magic  uint32;
    <a id="L15"></a>Cpu    Cpu;
    <a id="L16"></a>SubCpu uint32;
    <a id="L17"></a>Type   Type;
    <a id="L18"></a>Ncmd   uint32;
    <a id="L19"></a>Cmdsz  uint32;
    <a id="L20"></a>Flags  uint32;
<a id="L21"></a>}

<a id="L23"></a>const (
    <a id="L24"></a>fileHeaderSize32 = 7 * 4;
    <a id="L25"></a>fileHeaderSize64 = 8 * 4;
<a id="L26"></a>)

<a id="L28"></a>const (
    <a id="L29"></a>Magic32 uint32 = 0xfeedface;
    <a id="L30"></a>Magic64 uint32 = 0xfeedfacf;
<a id="L31"></a>)

<a id="L33"></a><span class="comment">// A Type is a Mach-O file type, either an object or an executable.</span>
<a id="L34"></a>type Type uint32

<a id="L36"></a>const (
    <a id="L37"></a>TypeObj  Type = 1;
    <a id="L38"></a>TypeExec Type = 2;
<a id="L39"></a>)

<a id="L41"></a><span class="comment">// A Cpu is a Mach-O cpu type.</span>
<a id="L42"></a>type Cpu uint32

<a id="L44"></a>const (
    <a id="L45"></a>Cpu386   Cpu = 7;
    <a id="L46"></a>CpuAmd64 Cpu = Cpu386 + 1&lt;&lt;24;
<a id="L47"></a>)

<a id="L49"></a>var cpuStrings = []intName{
    <a id="L50"></a>intName{uint32(Cpu386), &#34;Cpu386&#34;},
    <a id="L51"></a>intName{uint32(CpuAmd64), &#34;CpuAmd64&#34;},
<a id="L52"></a>}

<a id="L54"></a>func (i Cpu) String() string   { return stringName(uint32(i), cpuStrings, false) }
<a id="L55"></a>func (i Cpu) GoString() string { return stringName(uint32(i), cpuStrings, true) }

<a id="L57"></a><span class="comment">// A LoadCmd is a Mach-O load command.</span>
<a id="L58"></a>type LoadCmd uint32

<a id="L60"></a>const (
    <a id="L61"></a>LoadCmdSegment    LoadCmd = 1;
    <a id="L62"></a>LoadCmdSegment64  LoadCmd = 25;
    <a id="L63"></a>LoadCmdThread     LoadCmd = 4;
    <a id="L64"></a>LoadCmdUnixThread LoadCmd = 5; <span class="comment">// thread+stack</span>
<a id="L65"></a>)

<a id="L67"></a>var cmdStrings = []intName{
    <a id="L68"></a>intName{uint32(LoadCmdSegment), &#34;LoadCmdSegment&#34;},
    <a id="L69"></a>intName{uint32(LoadCmdSegment64), &#34;LoadCmdSegment64&#34;},
    <a id="L70"></a>intName{uint32(LoadCmdThread), &#34;LoadCmdThread&#34;},
    <a id="L71"></a>intName{uint32(LoadCmdUnixThread), &#34;LoadCmdUnixThread&#34;},
<a id="L72"></a>}

<a id="L74"></a>func (i LoadCmd) String() string   { return stringName(uint32(i), cmdStrings, false) }
<a id="L75"></a>func (i LoadCmd) GoString() string { return stringName(uint32(i), cmdStrings, true) }

<a id="L77"></a><span class="comment">// A Segment64 is a 64-bit Mach-O segment load command.</span>
<a id="L78"></a>type Segment64 struct {
    <a id="L79"></a>Cmd     LoadCmd;
    <a id="L80"></a>Len     uint32;
    <a id="L81"></a>Name    [16]byte;
    <a id="L82"></a>Addr    uint64;
    <a id="L83"></a>Memsz   uint64;
    <a id="L84"></a>Offset  uint64;
    <a id="L85"></a>Filesz  uint64;
    <a id="L86"></a>Maxprot uint32;
    <a id="L87"></a>Prot    uint32;
    <a id="L88"></a>Nsect   uint32;
    <a id="L89"></a>Flag    uint32;
<a id="L90"></a>}

<a id="L92"></a><span class="comment">// A Segment32 is a 32-bit Mach-O segment load command.</span>
<a id="L93"></a>type Segment32 struct {
    <a id="L94"></a>Cmd     LoadCmd;
    <a id="L95"></a>Len     uint32;
    <a id="L96"></a>Name    [16]byte;
    <a id="L97"></a>Addr    uint32;
    <a id="L98"></a>Memsz   uint32;
    <a id="L99"></a>Offset  uint32;
    <a id="L100"></a>Filesz  uint32;
    <a id="L101"></a>Maxprot uint32;
    <a id="L102"></a>Prot    uint32;
    <a id="L103"></a>Nsect   uint32;
    <a id="L104"></a>Flag    uint32;
<a id="L105"></a>}

<a id="L107"></a><span class="comment">// A Section32 is a 32-bit Mach-O section header.</span>
<a id="L108"></a>type Section32 struct {
    <a id="L109"></a>Name     [16]byte;
    <a id="L110"></a>Seg      [16]byte;
    <a id="L111"></a>Addr     uint32;
    <a id="L112"></a>Size     uint32;
    <a id="L113"></a>Offset   uint32;
    <a id="L114"></a>Align    uint32;
    <a id="L115"></a>Reloff   uint32;
    <a id="L116"></a>Nreloc   uint32;
    <a id="L117"></a>Flags    uint32;
    <a id="L118"></a>Reserve1 uint32;
    <a id="L119"></a>Reserve2 uint32;
<a id="L120"></a>}

<a id="L122"></a><span class="comment">// A Section32 is a 64-bit Mach-O section header.</span>
<a id="L123"></a>type Section64 struct {
    <a id="L124"></a>Name     [16]byte;
    <a id="L125"></a>Seg      [16]byte;
    <a id="L126"></a>Addr     uint64;
    <a id="L127"></a>Size     uint64;
    <a id="L128"></a>Offset   uint32;
    <a id="L129"></a>Align    uint32;
    <a id="L130"></a>Reloff   uint32;
    <a id="L131"></a>Nreloc   uint32;
    <a id="L132"></a>Flags    uint32;
    <a id="L133"></a>Reserve1 uint32;
    <a id="L134"></a>Reserve2 uint32;
    <a id="L135"></a>Reserve3 uint32;
<a id="L136"></a>}

<a id="L138"></a><span class="comment">// A Thread is a Mach-O thread state command.</span>
<a id="L139"></a>type Thread struct {
    <a id="L140"></a>Cmd  LoadCmd;
    <a id="L141"></a>Len  uint32;
    <a id="L142"></a>Type uint32;
    <a id="L143"></a>Data []uint32;
<a id="L144"></a>}

<a id="L146"></a><span class="comment">// Regs386 is the Mach-O 386 register structure.</span>
<a id="L147"></a>type Regs386 struct {
    <a id="L148"></a>AX    uint32;
    <a id="L149"></a>BX    uint32;
    <a id="L150"></a>CX    uint32;
    <a id="L151"></a>DX    uint32;
    <a id="L152"></a>DI    uint32;
    <a id="L153"></a>SI    uint32;
    <a id="L154"></a>BP    uint32;
    <a id="L155"></a>SP    uint32;
    <a id="L156"></a>SS    uint32;
    <a id="L157"></a>FLAGS uint32;
    <a id="L158"></a>IP    uint32;
    <a id="L159"></a>CS    uint32;
    <a id="L160"></a>DS    uint32;
    <a id="L161"></a>ES    uint32;
    <a id="L162"></a>FS    uint32;
    <a id="L163"></a>GS    uint32;
<a id="L164"></a>}

<a id="L166"></a><span class="comment">// RegsAMD64 is the Mach-O AMD64 register structure.</span>
<a id="L167"></a>type RegsAMD64 struct {
    <a id="L168"></a>AX    uint64;
    <a id="L169"></a>BX    uint64;
    <a id="L170"></a>CX    uint64;
    <a id="L171"></a>DX    uint64;
    <a id="L172"></a>DI    uint64;
    <a id="L173"></a>SI    uint64;
    <a id="L174"></a>BP    uint64;
    <a id="L175"></a>SP    uint64;
    <a id="L176"></a>R8    uint64;
    <a id="L177"></a>R9    uint64;
    <a id="L178"></a>R10   uint64;
    <a id="L179"></a>R11   uint64;
    <a id="L180"></a>R12   uint64;
    <a id="L181"></a>R13   uint64;
    <a id="L182"></a>R14   uint64;
    <a id="L183"></a>R15   uint64;
    <a id="L184"></a>IP    uint64;
    <a id="L185"></a>FLAGS uint64;
    <a id="L186"></a>CS    uint64;
    <a id="L187"></a>FS    uint64;
    <a id="L188"></a>GS    uint64;
<a id="L189"></a>}

<a id="L191"></a>type intName struct {
    <a id="L192"></a>i   uint32;
    <a id="L193"></a>s   string;
<a id="L194"></a>}

<a id="L196"></a>func stringName(i uint32, names []intName, goSyntax bool) string {
    <a id="L197"></a>for _, n := range names {
        <a id="L198"></a>if n.i == i {
            <a id="L199"></a>if goSyntax {
                <a id="L200"></a>return &#34;macho.&#34; + n.s
            <a id="L201"></a>}
            <a id="L202"></a>return n.s;
        <a id="L203"></a>}
    <a id="L204"></a>}
    <a id="L205"></a>return strconv.Uitoa64(uint64(i));
<a id="L206"></a>}

<a id="L208"></a>func flagName(i uint32, names []intName, goSyntax bool) string {
    <a id="L209"></a>s := &#34;&#34;;
    <a id="L210"></a>for _, n := range names {
        <a id="L211"></a>if n.i&amp;i == n.i {
            <a id="L212"></a>if len(s) &gt; 0 {
                <a id="L213"></a>s += &#34;+&#34;
            <a id="L214"></a>}
            <a id="L215"></a>if goSyntax {
                <a id="L216"></a>s += &#34;macho.&#34;
            <a id="L217"></a>}
            <a id="L218"></a>s += n.s;
            <a id="L219"></a>i -= n.i;
        <a id="L220"></a>}
    <a id="L221"></a>}
    <a id="L222"></a>if len(s) == 0 {
        <a id="L223"></a>return &#34;0x&#34; + strconv.Uitob64(uint64(i), 16)
    <a id="L224"></a>}
    <a id="L225"></a>if i != 0 {
        <a id="L226"></a>s += &#34;+0x&#34; + strconv.Uitob64(uint64(i), 16)
    <a id="L227"></a>}
    <a id="L228"></a>return s;
<a id="L229"></a>}
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
