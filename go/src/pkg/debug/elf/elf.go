<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/elf/elf.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/debug/elf/elf.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">/*</span>
<a id="L2"></a><span class="comment"> * ELF constants and data structures</span>
<a id="L3"></a><span class="comment"> *</span>
<a id="L4"></a><span class="comment"> * Derived from:</span>
<a id="L5"></a><span class="comment"> * $FreeBSD: src/sys/sys/elf32.h,v 1.8.14.1 2005/12/30 22:13:58 marcel Exp $</span>
<a id="L6"></a><span class="comment"> * $FreeBSD: src/sys/sys/elf64.h,v 1.10.14.1 2005/12/30 22:13:58 marcel Exp $</span>
<a id="L7"></a><span class="comment"> * $FreeBSD: src/sys/sys/elf_common.h,v 1.15.8.1 2005/12/30 22:13:58 marcel Exp $</span>
<a id="L8"></a><span class="comment"> * $FreeBSD: src/sys/alpha/include/elf.h,v 1.14 2003/09/25 01:10:22 peter Exp $</span>
<a id="L9"></a><span class="comment"> * $FreeBSD: src/sys/amd64/include/elf.h,v 1.18 2004/08/03 08:21:48 dfr Exp $</span>
<a id="L10"></a><span class="comment"> * $FreeBSD: src/sys/arm/include/elf.h,v 1.5.2.1 2006/06/30 21:42:52 cognet Exp $</span>
<a id="L11"></a><span class="comment"> * $FreeBSD: src/sys/i386/include/elf.h,v 1.16 2004/08/02 19:12:17 dfr Exp $</span>
<a id="L12"></a><span class="comment"> * $FreeBSD: src/sys/powerpc/include/elf.h,v 1.7 2004/11/02 09:47:01 ssouhlal Exp $</span>
<a id="L13"></a><span class="comment"> * $FreeBSD: src/sys/sparc64/include/elf.h,v 1.12 2003/09/25 01:10:26 peter Exp $</span>
<a id="L14"></a><span class="comment"> *</span>
<a id="L15"></a><span class="comment"> * Copyright (c) 1996-1998 John D. Polstra.  All rights reserved.</span>
<a id="L16"></a><span class="comment"> * Copyright (c) 2001 David E. O&#39;Brien</span>
<a id="L17"></a><span class="comment"> * Portions Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L18"></a><span class="comment"> *</span>
<a id="L19"></a><span class="comment"> * Redistribution and use in source and binary forms, with or without</span>
<a id="L20"></a><span class="comment"> * modification, are permitted provided that the following conditions</span>
<a id="L21"></a><span class="comment"> * are met:</span>
<a id="L22"></a><span class="comment"> * 1. Redistributions of source code must retain the above copyright</span>
<a id="L23"></a><span class="comment"> *    notice, this list of conditions and the following disclaimer.</span>
<a id="L24"></a><span class="comment"> * 2. Redistributions in binary form must reproduce the above copyright</span>
<a id="L25"></a><span class="comment"> *    notice, this list of conditions and the following disclaimer in the</span>
<a id="L26"></a><span class="comment"> *    documentation and/or other materials provided with the distribution.</span>
<a id="L27"></a><span class="comment"> *</span>
<a id="L28"></a><span class="comment"> * THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS&#39;&#39; AND</span>
<a id="L29"></a><span class="comment"> * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE</span>
<a id="L30"></a><span class="comment"> * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE</span>
<a id="L31"></a><span class="comment"> * ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE</span>
<a id="L32"></a><span class="comment"> * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL</span>
<a id="L33"></a><span class="comment"> * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS</span>
<a id="L34"></a><span class="comment"> * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)</span>
<a id="L35"></a><span class="comment"> * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT</span>
<a id="L36"></a><span class="comment"> * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY</span>
<a id="L37"></a><span class="comment"> * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF</span>
<a id="L38"></a><span class="comment"> * SUCH DAMAGE.</span>
<a id="L39"></a><span class="comment"> */</span>

<a id="L41"></a>package elf

<a id="L43"></a>import &#34;strconv&#34;

<a id="L45"></a><span class="comment">/*</span>
<a id="L46"></a><span class="comment"> * Constants</span>
<a id="L47"></a><span class="comment"> */</span>

<a id="L49"></a><span class="comment">// Indexes into the Header.Ident array.</span>
<a id="L50"></a>const (
    <a id="L51"></a>EI_CLASS      = 4;  <span class="comment">/* Class of machine. */</span>
    <a id="L52"></a>EI_DATA       = 5;  <span class="comment">/* Data format. */</span>
    <a id="L53"></a>EI_VERSION    = 6;  <span class="comment">/* ELF format version. */</span>
    <a id="L54"></a>EI_OSABI      = 7;  <span class="comment">/* Operating system / ABI identification */</span>
    <a id="L55"></a>EI_ABIVERSION = 8;  <span class="comment">/* ABI version */</span>
    <a id="L56"></a>EI_PAD        = 9;  <span class="comment">/* Start of padding (per SVR4 ABI). */</span>
    <a id="L57"></a>EI_NIDENT     = 16; <span class="comment">/* Size of e_ident array. */</span>
<a id="L58"></a>)

<a id="L60"></a><span class="comment">// Initial magic number for ELF files.</span>
<a id="L61"></a>const ELFMAG = &#34;\177ELF&#34;

<a id="L63"></a><span class="comment">// Version is found in Header.Ident[EI_VERSION] and Header.Version.</span>
<a id="L64"></a>type Version byte

<a id="L66"></a>const (
    <a id="L67"></a>EV_NONE    Version = 0;
    <a id="L68"></a>EV_CURRENT Version = 1;
<a id="L69"></a>)

<a id="L71"></a>var versionStrings = []intName{
    <a id="L72"></a>intName{0, &#34;EV_NONE&#34;},
    <a id="L73"></a>intName{1, &#34;EV_CURRENT&#34;},
<a id="L74"></a>}

<a id="L76"></a>func (i Version) String() string   { return stringName(uint32(i), versionStrings, false) }
<a id="L77"></a>func (i Version) GoString() string { return stringName(uint32(i), versionStrings, true) }

<a id="L79"></a><span class="comment">// Class is found in Header.Ident[EI_CLASS] and Header.Class.</span>
<a id="L80"></a>type Class byte

<a id="L82"></a>const (
    <a id="L83"></a>ELFCLASSNONE Class = 0; <span class="comment">/* Unknown class. */</span>
    <a id="L84"></a>ELFCLASS32   Class = 1; <span class="comment">/* 32-bit architecture. */</span>
    <a id="L85"></a>ELFCLASS64   Class = 2; <span class="comment">/* 64-bit architecture. */</span>
<a id="L86"></a>)

<a id="L88"></a>var classStrings = []intName{
    <a id="L89"></a>intName{0, &#34;ELFCLASSNONE&#34;},
    <a id="L90"></a>intName{1, &#34;ELFCLASS32&#34;},
    <a id="L91"></a>intName{2, &#34;ELFCLASS64&#34;},
<a id="L92"></a>}

<a id="L94"></a>func (i Class) String() string   { return stringName(uint32(i), classStrings, false) }
<a id="L95"></a>func (i Class) GoString() string { return stringName(uint32(i), classStrings, true) }

<a id="L97"></a><span class="comment">// Data is found in Header.Ident[EI_DATA] and Header.Data.</span>
<a id="L98"></a>type Data byte

<a id="L100"></a>const (
    <a id="L101"></a>ELFDATANONE Data = 0; <span class="comment">/* Unknown data format. */</span>
    <a id="L102"></a>ELFDATA2LSB Data = 1; <span class="comment">/* 2&#39;s complement little-endian. */</span>
    <a id="L103"></a>ELFDATA2MSB Data = 2; <span class="comment">/* 2&#39;s complement big-endian. */</span>
<a id="L104"></a>)

<a id="L106"></a>var dataStrings = []intName{
    <a id="L107"></a>intName{0, &#34;ELFDATANONE&#34;},
    <a id="L108"></a>intName{1, &#34;ELFDATA2LSB&#34;},
    <a id="L109"></a>intName{2, &#34;ELFDATA2MSB&#34;},
<a id="L110"></a>}

<a id="L112"></a>func (i Data) String() string   { return stringName(uint32(i), dataStrings, false) }
<a id="L113"></a>func (i Data) GoString() string { return stringName(uint32(i), dataStrings, true) }

<a id="L115"></a><span class="comment">// OSABI is found in Header.Ident[EI_OSABI] and Header.OSABI.</span>
<a id="L116"></a>type OSABI byte

<a id="L118"></a>const (
    <a id="L119"></a>ELFOSABI_NONE       OSABI = 0;   <span class="comment">/* UNIX System V ABI */</span>
    <a id="L120"></a>ELFOSABI_HPUX       OSABI = 1;   <span class="comment">/* HP-UX operating system */</span>
    <a id="L121"></a>ELFOSABI_NETBSD     OSABI = 2;   <span class="comment">/* NetBSD */</span>
    <a id="L122"></a>ELFOSABI_LINUX      OSABI = 3;   <span class="comment">/* GNU/Linux */</span>
    <a id="L123"></a>ELFOSABI_HURD       OSABI = 4;   <span class="comment">/* GNU/Hurd */</span>
    <a id="L124"></a>ELFOSABI_86OPEN     OSABI = 5;   <span class="comment">/* 86Open common IA32 ABI */</span>
    <a id="L125"></a>ELFOSABI_SOLARIS    OSABI = 6;   <span class="comment">/* Solaris */</span>
    <a id="L126"></a>ELFOSABI_AIX        OSABI = 7;   <span class="comment">/* AIX */</span>
    <a id="L127"></a>ELFOSABI_IRIX       OSABI = 8;   <span class="comment">/* IRIX */</span>
    <a id="L128"></a>ELFOSABI_FREEBSD    OSABI = 9;   <span class="comment">/* FreeBSD */</span>
    <a id="L129"></a>ELFOSABI_TRU64      OSABI = 10;  <span class="comment">/* TRU64 UNIX */</span>
    <a id="L130"></a>ELFOSABI_MODESTO    OSABI = 11;  <span class="comment">/* Novell Modesto */</span>
    <a id="L131"></a>ELFOSABI_OPENBSD    OSABI = 12;  <span class="comment">/* OpenBSD */</span>
    <a id="L132"></a>ELFOSABI_OPENVMS    OSABI = 13;  <span class="comment">/* Open VMS */</span>
    <a id="L133"></a>ELFOSABI_NSK        OSABI = 14;  <span class="comment">/* HP Non-Stop Kernel */</span>
    <a id="L134"></a>ELFOSABI_ARM        OSABI = 97;  <span class="comment">/* ARM */</span>
    <a id="L135"></a>ELFOSABI_STANDALONE OSABI = 255; <span class="comment">/* Standalone (embedded) application */</span>
<a id="L136"></a>)

<a id="L138"></a>var osabiStrings = []intName{
    <a id="L139"></a>intName{0, &#34;ELFOSABI_NONE&#34;},
    <a id="L140"></a>intName{1, &#34;ELFOSABI_HPUX&#34;},
    <a id="L141"></a>intName{2, &#34;ELFOSABI_NETBSD&#34;},
    <a id="L142"></a>intName{3, &#34;ELFOSABI_LINUX&#34;},
    <a id="L143"></a>intName{4, &#34;ELFOSABI_HURD&#34;},
    <a id="L144"></a>intName{5, &#34;ELFOSABI_86OPEN&#34;},
    <a id="L145"></a>intName{6, &#34;ELFOSABI_SOLARIS&#34;},
    <a id="L146"></a>intName{7, &#34;ELFOSABI_AIX&#34;},
    <a id="L147"></a>intName{8, &#34;ELFOSABI_IRIX&#34;},
    <a id="L148"></a>intName{9, &#34;ELFOSABI_FREEBSD&#34;},
    <a id="L149"></a>intName{10, &#34;ELFOSABI_TRU64&#34;},
    <a id="L150"></a>intName{11, &#34;ELFOSABI_MODESTO&#34;},
    <a id="L151"></a>intName{12, &#34;ELFOSABI_OPENBSD&#34;},
    <a id="L152"></a>intName{13, &#34;ELFOSABI_OPENVMS&#34;},
    <a id="L153"></a>intName{14, &#34;ELFOSABI_NSK&#34;},
    <a id="L154"></a>intName{97, &#34;ELFOSABI_ARM&#34;},
    <a id="L155"></a>intName{255, &#34;ELFOSABI_STANDALONE&#34;},
<a id="L156"></a>}

<a id="L158"></a>func (i OSABI) String() string   { return stringName(uint32(i), osabiStrings, false) }
<a id="L159"></a>func (i OSABI) GoString() string { return stringName(uint32(i), osabiStrings, true) }

<a id="L161"></a><span class="comment">// Type is found in Header.Type.</span>
<a id="L162"></a>type Type uint16

<a id="L164"></a>const (
    <a id="L165"></a>ET_NONE   Type = 0;      <span class="comment">/* Unknown type. */</span>
    <a id="L166"></a>ET_REL    Type = 1;      <span class="comment">/* Relocatable. */</span>
    <a id="L167"></a>ET_EXEC   Type = 2;      <span class="comment">/* Executable. */</span>
    <a id="L168"></a>ET_DYN    Type = 3;      <span class="comment">/* Shared object. */</span>
    <a id="L169"></a>ET_CORE   Type = 4;      <span class="comment">/* Core file. */</span>
    <a id="L170"></a>ET_LOOS   Type = 0xfe00; <span class="comment">/* First operating system specific. */</span>
    <a id="L171"></a>ET_HIOS   Type = 0xfeff; <span class="comment">/* Last operating system-specific. */</span>
    <a id="L172"></a>ET_LOPROC Type = 0xff00; <span class="comment">/* First processor-specific. */</span>
    <a id="L173"></a>ET_HIPROC Type = 0xffff; <span class="comment">/* Last processor-specific. */</span>
<a id="L174"></a>)

<a id="L176"></a>var typeStrings = []intName{
    <a id="L177"></a>intName{0, &#34;ET_NONE&#34;},
    <a id="L178"></a>intName{1, &#34;ET_REL&#34;},
    <a id="L179"></a>intName{2, &#34;ET_EXEC&#34;},
    <a id="L180"></a>intName{3, &#34;ET_DYN&#34;},
    <a id="L181"></a>intName{4, &#34;ET_CORE&#34;},
    <a id="L182"></a>intName{0xfe00, &#34;ET_LOOS&#34;},
    <a id="L183"></a>intName{0xfeff, &#34;ET_HIOS&#34;},
    <a id="L184"></a>intName{0xff00, &#34;ET_LOPROC&#34;},
    <a id="L185"></a>intName{0xffff, &#34;ET_HIPROC&#34;},
<a id="L186"></a>}

<a id="L188"></a>func (i Type) String() string   { return stringName(uint32(i), typeStrings, false) }
<a id="L189"></a>func (i Type) GoString() string { return stringName(uint32(i), typeStrings, true) }

<a id="L191"></a><span class="comment">// Machine is found in Header.Machine.</span>
<a id="L192"></a>type Machine uint16

<a id="L194"></a>const (
    <a id="L195"></a>EM_NONE        Machine = 0;  <span class="comment">/* Unknown machine. */</span>
    <a id="L196"></a>EM_M32         Machine = 1;  <span class="comment">/* AT&amp;T WE32100. */</span>
    <a id="L197"></a>EM_SPARC       Machine = 2;  <span class="comment">/* Sun SPARC. */</span>
    <a id="L198"></a>EM_386         Machine = 3;  <span class="comment">/* Intel i386. */</span>
    <a id="L199"></a>EM_68K         Machine = 4;  <span class="comment">/* Motorola 68000. */</span>
    <a id="L200"></a>EM_88K         Machine = 5;  <span class="comment">/* Motorola 88000. */</span>
    <a id="L201"></a>EM_860         Machine = 7;  <span class="comment">/* Intel i860. */</span>
    <a id="L202"></a>EM_MIPS        Machine = 8;  <span class="comment">/* MIPS R3000 Big-Endian only. */</span>
    <a id="L203"></a>EM_S370        Machine = 9;  <span class="comment">/* IBM System/370. */</span>
    <a id="L204"></a>EM_MIPS_RS3_LE Machine = 10; <span class="comment">/* MIPS R3000 Little-Endian. */</span>
    <a id="L205"></a>EM_PARISC      Machine = 15; <span class="comment">/* HP PA-RISC. */</span>
    <a id="L206"></a>EM_VPP500      Machine = 17; <span class="comment">/* Fujitsu VPP500. */</span>
    <a id="L207"></a>EM_SPARC32PLUS Machine = 18; <span class="comment">/* SPARC v8plus. */</span>
    <a id="L208"></a>EM_960         Machine = 19; <span class="comment">/* Intel 80960. */</span>
    <a id="L209"></a>EM_PPC         Machine = 20; <span class="comment">/* PowerPC 32-bit. */</span>
    <a id="L210"></a>EM_PPC64       Machine = 21; <span class="comment">/* PowerPC 64-bit. */</span>
    <a id="L211"></a>EM_S390        Machine = 22; <span class="comment">/* IBM System/390. */</span>
    <a id="L212"></a>EM_V800        Machine = 36; <span class="comment">/* NEC V800. */</span>
    <a id="L213"></a>EM_FR20        Machine = 37; <span class="comment">/* Fujitsu FR20. */</span>
    <a id="L214"></a>EM_RH32        Machine = 38; <span class="comment">/* TRW RH-32. */</span>
    <a id="L215"></a>EM_RCE         Machine = 39; <span class="comment">/* Motorola RCE. */</span>
    <a id="L216"></a>EM_ARM         Machine = 40; <span class="comment">/* ARM. */</span>
    <a id="L217"></a>EM_SH          Machine = 42; <span class="comment">/* Hitachi SH. */</span>
    <a id="L218"></a>EM_SPARCV9     Machine = 43; <span class="comment">/* SPARC v9 64-bit. */</span>
    <a id="L219"></a>EM_TRICORE     Machine = 44; <span class="comment">/* Siemens TriCore embedded processor. */</span>
    <a id="L220"></a>EM_ARC         Machine = 45; <span class="comment">/* Argonaut RISC Core. */</span>
    <a id="L221"></a>EM_H8_300      Machine = 46; <span class="comment">/* Hitachi H8/300. */</span>
    <a id="L222"></a>EM_H8_300H     Machine = 47; <span class="comment">/* Hitachi H8/300H. */</span>
    <a id="L223"></a>EM_H8S         Machine = 48; <span class="comment">/* Hitachi H8S. */</span>
    <a id="L224"></a>EM_H8_500      Machine = 49; <span class="comment">/* Hitachi H8/500. */</span>
    <a id="L225"></a>EM_IA_64       Machine = 50; <span class="comment">/* Intel IA-64 Processor. */</span>
    <a id="L226"></a>EM_MIPS_X      Machine = 51; <span class="comment">/* Stanford MIPS-X. */</span>
    <a id="L227"></a>EM_COLDFIRE    Machine = 52; <span class="comment">/* Motorola ColdFire. */</span>
    <a id="L228"></a>EM_68HC12      Machine = 53; <span class="comment">/* Motorola M68HC12. */</span>
    <a id="L229"></a>EM_MMA         Machine = 54; <span class="comment">/* Fujitsu MMA. */</span>
    <a id="L230"></a>EM_PCP         Machine = 55; <span class="comment">/* Siemens PCP. */</span>
    <a id="L231"></a>EM_NCPU        Machine = 56; <span class="comment">/* Sony nCPU. */</span>
    <a id="L232"></a>EM_NDR1        Machine = 57; <span class="comment">/* Denso NDR1 microprocessor. */</span>
    <a id="L233"></a>EM_STARCORE    Machine = 58; <span class="comment">/* Motorola Star*Core processor. */</span>
    <a id="L234"></a>EM_ME16        Machine = 59; <span class="comment">/* Toyota ME16 processor. */</span>
    <a id="L235"></a>EM_ST100       Machine = 60; <span class="comment">/* STMicroelectronics ST100 processor. */</span>
    <a id="L236"></a>EM_TINYJ       Machine = 61; <span class="comment">/* Advanced Logic Corp. TinyJ processor. */</span>
    <a id="L237"></a>EM_X86_64      Machine = 62; <span class="comment">/* Advanced Micro Devices x86-64 */</span>

    <a id="L239"></a><span class="comment">/* Non-standard or deprecated. */</span>
    <a id="L240"></a>EM_486         Machine = 6;      <span class="comment">/* Intel i486. */</span>
    <a id="L241"></a>EM_MIPS_RS4_BE Machine = 10;     <span class="comment">/* MIPS R4000 Big-Endian */</span>
    <a id="L242"></a>EM_ALPHA_STD   Machine = 41;     <span class="comment">/* Digital Alpha (standard value). */</span>
    <a id="L243"></a>EM_ALPHA       Machine = 0x9026; <span class="comment">/* Alpha (written in the absence of an ABI) */</span>
<a id="L244"></a>)

<a id="L246"></a>var machineStrings = []intName{
    <a id="L247"></a>intName{0, &#34;EM_NONE&#34;},
    <a id="L248"></a>intName{1, &#34;EM_M32&#34;},
    <a id="L249"></a>intName{2, &#34;EM_SPARC&#34;},
    <a id="L250"></a>intName{3, &#34;EM_386&#34;},
    <a id="L251"></a>intName{4, &#34;EM_68K&#34;},
    <a id="L252"></a>intName{5, &#34;EM_88K&#34;},
    <a id="L253"></a>intName{7, &#34;EM_860&#34;},
    <a id="L254"></a>intName{8, &#34;EM_MIPS&#34;},
    <a id="L255"></a>intName{9, &#34;EM_S370&#34;},
    <a id="L256"></a>intName{10, &#34;EM_MIPS_RS3_LE&#34;},
    <a id="L257"></a>intName{15, &#34;EM_PARISC&#34;},
    <a id="L258"></a>intName{17, &#34;EM_VPP500&#34;},
    <a id="L259"></a>intName{18, &#34;EM_SPARC32PLUS&#34;},
    <a id="L260"></a>intName{19, &#34;EM_960&#34;},
    <a id="L261"></a>intName{20, &#34;EM_PPC&#34;},
    <a id="L262"></a>intName{21, &#34;EM_PPC64&#34;},
    <a id="L263"></a>intName{22, &#34;EM_S390&#34;},
    <a id="L264"></a>intName{36, &#34;EM_V800&#34;},
    <a id="L265"></a>intName{37, &#34;EM_FR20&#34;},
    <a id="L266"></a>intName{38, &#34;EM_RH32&#34;},
    <a id="L267"></a>intName{39, &#34;EM_RCE&#34;},
    <a id="L268"></a>intName{40, &#34;EM_ARM&#34;},
    <a id="L269"></a>intName{42, &#34;EM_SH&#34;},
    <a id="L270"></a>intName{43, &#34;EM_SPARCV9&#34;},
    <a id="L271"></a>intName{44, &#34;EM_TRICORE&#34;},
    <a id="L272"></a>intName{45, &#34;EM_ARC&#34;},
    <a id="L273"></a>intName{46, &#34;EM_H8_300&#34;},
    <a id="L274"></a>intName{47, &#34;EM_H8_300H&#34;},
    <a id="L275"></a>intName{48, &#34;EM_H8S&#34;},
    <a id="L276"></a>intName{49, &#34;EM_H8_500&#34;},
    <a id="L277"></a>intName{50, &#34;EM_IA_64&#34;},
    <a id="L278"></a>intName{51, &#34;EM_MIPS_X&#34;},
    <a id="L279"></a>intName{52, &#34;EM_COLDFIRE&#34;},
    <a id="L280"></a>intName{53, &#34;EM_68HC12&#34;},
    <a id="L281"></a>intName{54, &#34;EM_MMA&#34;},
    <a id="L282"></a>intName{55, &#34;EM_PCP&#34;},
    <a id="L283"></a>intName{56, &#34;EM_NCPU&#34;},
    <a id="L284"></a>intName{57, &#34;EM_NDR1&#34;},
    <a id="L285"></a>intName{58, &#34;EM_STARCORE&#34;},
    <a id="L286"></a>intName{59, &#34;EM_ME16&#34;},
    <a id="L287"></a>intName{60, &#34;EM_ST100&#34;},
    <a id="L288"></a>intName{61, &#34;EM_TINYJ&#34;},
    <a id="L289"></a>intName{62, &#34;EM_X86_64&#34;},

    <a id="L291"></a><span class="comment">/* Non-standard or deprecated. */</span>
    <a id="L292"></a>intName{6, &#34;EM_486&#34;},
    <a id="L293"></a>intName{10, &#34;EM_MIPS_RS4_BE&#34;},
    <a id="L294"></a>intName{41, &#34;EM_ALPHA_STD&#34;},
    <a id="L295"></a>intName{0x9026, &#34;EM_ALPHA&#34;},
<a id="L296"></a>}

<a id="L298"></a>func (i Machine) String() string   { return stringName(uint32(i), machineStrings, false) }
<a id="L299"></a>func (i Machine) GoString() string { return stringName(uint32(i), machineStrings, true) }

<a id="L301"></a><span class="comment">// Special section indices.</span>
<a id="L302"></a>type SectionIndex int

<a id="L304"></a>const (
    <a id="L305"></a>SHN_UNDEF     SectionIndex = 0;      <span class="comment">/* Undefined, missing, irrelevant. */</span>
    <a id="L306"></a>SHN_LORESERVE SectionIndex = 0xff00; <span class="comment">/* First of reserved range. */</span>
    <a id="L307"></a>SHN_LOPROC    SectionIndex = 0xff00; <span class="comment">/* First processor-specific. */</span>
    <a id="L308"></a>SHN_HIPROC    SectionIndex = 0xff1f; <span class="comment">/* Last processor-specific. */</span>
    <a id="L309"></a>SHN_LOOS      SectionIndex = 0xff20; <span class="comment">/* First operating system-specific. */</span>
    <a id="L310"></a>SHN_HIOS      SectionIndex = 0xff3f; <span class="comment">/* Last operating system-specific. */</span>
    <a id="L311"></a>SHN_ABS       SectionIndex = 0xfff1; <span class="comment">/* Absolute values. */</span>
    <a id="L312"></a>SHN_COMMON    SectionIndex = 0xfff2; <span class="comment">/* Common data. */</span>
    <a id="L313"></a>SHN_XINDEX    SectionIndex = 0xffff; <span class="comment">/* Escape -- index stored elsewhere. */</span>
    <a id="L314"></a>SHN_HIRESERVE SectionIndex = 0xffff; <span class="comment">/* Last of reserved range. */</span>
<a id="L315"></a>)

<a id="L317"></a>var shnStrings = []intName{
    <a id="L318"></a>intName{0, &#34;SHN_UNDEF&#34;},
    <a id="L319"></a>intName{0xff00, &#34;SHN_LOPROC&#34;},
    <a id="L320"></a>intName{0xff20, &#34;SHN_LOOS&#34;},
    <a id="L321"></a>intName{0xfff1, &#34;SHN_ABS&#34;},
    <a id="L322"></a>intName{0xfff2, &#34;SHN_COMMON&#34;},
    <a id="L323"></a>intName{0xffff, &#34;SHN_XINDEX&#34;},
<a id="L324"></a>}

<a id="L326"></a>func (i SectionIndex) String() string   { return stringName(uint32(i), shnStrings, false) }
<a id="L327"></a>func (i SectionIndex) GoString() string { return stringName(uint32(i), shnStrings, true) }

<a id="L329"></a><span class="comment">// Section type.</span>
<a id="L330"></a>type SectionType uint32

<a id="L332"></a>const (
    <a id="L333"></a>SHT_NULL          SectionType = 0;          <span class="comment">/* inactive */</span>
    <a id="L334"></a>SHT_PROGBITS      SectionType = 1;          <span class="comment">/* program defined information */</span>
    <a id="L335"></a>SHT_SYMTAB        SectionType = 2;          <span class="comment">/* symbol table section */</span>
    <a id="L336"></a>SHT_STRTAB        SectionType = 3;          <span class="comment">/* string table section */</span>
    <a id="L337"></a>SHT_RELA          SectionType = 4;          <span class="comment">/* relocation section with addends */</span>
    <a id="L338"></a>SHT_HASH          SectionType = 5;          <span class="comment">/* symbol hash table section */</span>
    <a id="L339"></a>SHT_DYNAMIC       SectionType = 6;          <span class="comment">/* dynamic section */</span>
    <a id="L340"></a>SHT_NOTE          SectionType = 7;          <span class="comment">/* note section */</span>
    <a id="L341"></a>SHT_NOBITS        SectionType = 8;          <span class="comment">/* no space section */</span>
    <a id="L342"></a>SHT_REL           SectionType = 9;          <span class="comment">/* relocation section - no addends */</span>
    <a id="L343"></a>SHT_SHLIB         SectionType = 10;         <span class="comment">/* reserved - purpose unknown */</span>
    <a id="L344"></a>SHT_DYNSYM        SectionType = 11;         <span class="comment">/* dynamic symbol table section */</span>
    <a id="L345"></a>SHT_INIT_ARRAY    SectionType = 14;         <span class="comment">/* Initialization function pointers. */</span>
    <a id="L346"></a>SHT_FINI_ARRAY    SectionType = 15;         <span class="comment">/* Termination function pointers. */</span>
    <a id="L347"></a>SHT_PREINIT_ARRAY SectionType = 16;         <span class="comment">/* Pre-initialization function ptrs. */</span>
    <a id="L348"></a>SHT_GROUP         SectionType = 17;         <span class="comment">/* Section group. */</span>
    <a id="L349"></a>SHT_SYMTAB_SHNDX  SectionType = 18;         <span class="comment">/* Section indexes (see SHN_XINDEX). */</span>
    <a id="L350"></a>SHT_LOOS          SectionType = 0x60000000; <span class="comment">/* First of OS specific semantics */</span>
    <a id="L351"></a>SHT_HIOS          SectionType = 0x6fffffff; <span class="comment">/* Last of OS specific semantics */</span>
    <a id="L352"></a>SHT_LOPROC        SectionType = 0x70000000; <span class="comment">/* reserved range for processor */</span>
    <a id="L353"></a>SHT_HIPROC        SectionType = 0x7fffffff; <span class="comment">/* specific section header types */</span>
    <a id="L354"></a>SHT_LOUSER        SectionType = 0x80000000; <span class="comment">/* reserved range for application */</span>
    <a id="L355"></a>SHT_HIUSER        SectionType = 0xffffffff; <span class="comment">/* specific indexes */</span>
<a id="L356"></a>)

<a id="L358"></a>var shtStrings = []intName{
    <a id="L359"></a>intName{0, &#34;SHT_NULL&#34;},
    <a id="L360"></a>intName{1, &#34;SHT_PROGBITS&#34;},
    <a id="L361"></a>intName{2, &#34;SHT_SYMTAB&#34;},
    <a id="L362"></a>intName{3, &#34;SHT_STRTAB&#34;},
    <a id="L363"></a>intName{4, &#34;SHT_RELA&#34;},
    <a id="L364"></a>intName{5, &#34;SHT_HASH&#34;},
    <a id="L365"></a>intName{6, &#34;SHT_DYNAMIC&#34;},
    <a id="L366"></a>intName{7, &#34;SHT_NOTE&#34;},
    <a id="L367"></a>intName{8, &#34;SHT_NOBITS&#34;},
    <a id="L368"></a>intName{9, &#34;SHT_REL&#34;},
    <a id="L369"></a>intName{10, &#34;SHT_SHLIB&#34;},
    <a id="L370"></a>intName{11, &#34;SHT_DYNSYM&#34;},
    <a id="L371"></a>intName{14, &#34;SHT_INIT_ARRAY&#34;},
    <a id="L372"></a>intName{15, &#34;SHT_FINI_ARRAY&#34;},
    <a id="L373"></a>intName{16, &#34;SHT_PREINIT_ARRAY&#34;},
    <a id="L374"></a>intName{17, &#34;SHT_GROUP&#34;},
    <a id="L375"></a>intName{18, &#34;SHT_SYMTAB_SHNDX&#34;},
    <a id="L376"></a>intName{0x60000000, &#34;SHT_LOOS&#34;},
    <a id="L377"></a>intName{0x6fffffff, &#34;SHT_HIOS&#34;},
    <a id="L378"></a>intName{0x70000000, &#34;SHT_LOPROC&#34;},
    <a id="L379"></a>intName{0x7fffffff, &#34;SHT_HIPROC&#34;},
    <a id="L380"></a>intName{0x80000000, &#34;SHT_LOUSER&#34;},
    <a id="L381"></a>intName{0xffffffff, &#34;SHT_HIUSER&#34;},
<a id="L382"></a>}

<a id="L384"></a>func (i SectionType) String() string   { return stringName(uint32(i), shtStrings, false) }
<a id="L385"></a>func (i SectionType) GoString() string { return stringName(uint32(i), shtStrings, true) }

<a id="L387"></a><span class="comment">// Section flags.</span>
<a id="L388"></a>type SectionFlag uint32

<a id="L390"></a>const (
    <a id="L391"></a>SHF_WRITE            SectionFlag = 0x1;        <span class="comment">/* Section contains writable data. */</span>
    <a id="L392"></a>SHF_ALLOC            SectionFlag = 0x2;        <span class="comment">/* Section occupies memory. */</span>
    <a id="L393"></a>SHF_EXECINSTR        SectionFlag = 0x4;        <span class="comment">/* Section contains instructions. */</span>
    <a id="L394"></a>SHF_MERGE            SectionFlag = 0x10;       <span class="comment">/* Section may be merged. */</span>
    <a id="L395"></a>SHF_STRINGS          SectionFlag = 0x20;       <span class="comment">/* Section contains strings. */</span>
    <a id="L396"></a>SHF_INFO_LINK        SectionFlag = 0x40;       <span class="comment">/* sh_info holds section index. */</span>
    <a id="L397"></a>SHF_LINK_ORDER       SectionFlag = 0x80;       <span class="comment">/* Special ordering requirements. */</span>
    <a id="L398"></a>SHF_OS_NONCONFORMING SectionFlag = 0x100;      <span class="comment">/* OS-specific processing required. */</span>
    <a id="L399"></a>SHF_GROUP            SectionFlag = 0x200;      <span class="comment">/* Member of section group. */</span>
    <a id="L400"></a>SHF_TLS              SectionFlag = 0x400;      <span class="comment">/* Section contains TLS data. */</span>
    <a id="L401"></a>SHF_MASKOS           SectionFlag = 0x0ff00000; <span class="comment">/* OS-specific semantics. */</span>
    <a id="L402"></a>SHF_MASKPROC         SectionFlag = 0xf0000000; <span class="comment">/* Processor-specific semantics. */</span>
<a id="L403"></a>)

<a id="L405"></a>var shfStrings = []intName{
    <a id="L406"></a>intName{0x1, &#34;SHF_WRITE&#34;},
    <a id="L407"></a>intName{0x2, &#34;SHF_ALLOC&#34;},
    <a id="L408"></a>intName{0x4, &#34;SHF_EXECINSTR&#34;},
    <a id="L409"></a>intName{0x10, &#34;SHF_MERGE&#34;},
    <a id="L410"></a>intName{0x20, &#34;SHF_STRINGS&#34;},
    <a id="L411"></a>intName{0x40, &#34;SHF_INFO_LINK&#34;},
    <a id="L412"></a>intName{0x80, &#34;SHF_LINK_ORDER&#34;},
    <a id="L413"></a>intName{0x100, &#34;SHF_OS_NONCONFORMING&#34;},
    <a id="L414"></a>intName{0x200, &#34;SHF_GROUP&#34;},
    <a id="L415"></a>intName{0x400, &#34;SHF_TLS&#34;},
<a id="L416"></a>}

<a id="L418"></a>func (i SectionFlag) String() string   { return flagName(uint32(i), shfStrings, false) }
<a id="L419"></a>func (i SectionFlag) GoString() string { return flagName(uint32(i), shfStrings, true) }

<a id="L421"></a><span class="comment">// Prog.Type</span>
<a id="L422"></a>type ProgType int

<a id="L424"></a>const (
    <a id="L425"></a>PT_NULL    ProgType = 0;          <span class="comment">/* Unused entry. */</span>
    <a id="L426"></a>PT_LOAD    ProgType = 1;          <span class="comment">/* Loadable segment. */</span>
    <a id="L427"></a>PT_DYNAMIC ProgType = 2;          <span class="comment">/* Dynamic linking information segment. */</span>
    <a id="L428"></a>PT_INTERP  ProgType = 3;          <span class="comment">/* Pathname of interpreter. */</span>
    <a id="L429"></a>PT_NOTE    ProgType = 4;          <span class="comment">/* Auxiliary information. */</span>
    <a id="L430"></a>PT_SHLIB   ProgType = 5;          <span class="comment">/* Reserved (not used). */</span>
    <a id="L431"></a>PT_PHDR    ProgType = 6;          <span class="comment">/* Location of program header itself. */</span>
    <a id="L432"></a>PT_TLS     ProgType = 7;          <span class="comment">/* Thread local storage segment */</span>
    <a id="L433"></a>PT_LOOS    ProgType = 0x60000000; <span class="comment">/* First OS-specific. */</span>
    <a id="L434"></a>PT_HIOS    ProgType = 0x6fffffff; <span class="comment">/* Last OS-specific. */</span>
    <a id="L435"></a>PT_LOPROC  ProgType = 0x70000000; <span class="comment">/* First processor-specific type. */</span>
    <a id="L436"></a>PT_HIPROC  ProgType = 0x7fffffff; <span class="comment">/* Last processor-specific type. */</span>
<a id="L437"></a>)

<a id="L439"></a>var ptStrings = []intName{
    <a id="L440"></a>intName{0, &#34;PT_NULL&#34;},
    <a id="L441"></a>intName{1, &#34;PT_LOAD&#34;},
    <a id="L442"></a>intName{2, &#34;PT_DYNAMIC&#34;},
    <a id="L443"></a>intName{3, &#34;PT_INTERP&#34;},
    <a id="L444"></a>intName{4, &#34;PT_NOTE&#34;},
    <a id="L445"></a>intName{5, &#34;PT_SHLIB&#34;},
    <a id="L446"></a>intName{6, &#34;PT_PHDR&#34;},
    <a id="L447"></a>intName{7, &#34;PT_TLS&#34;},
    <a id="L448"></a>intName{0x60000000, &#34;PT_LOOS&#34;},
    <a id="L449"></a>intName{0x6fffffff, &#34;PT_HIOS&#34;},
    <a id="L450"></a>intName{0x70000000, &#34;PT_LOPROC&#34;},
    <a id="L451"></a>intName{0x7fffffff, &#34;PT_HIPROC&#34;},
<a id="L452"></a>}

<a id="L454"></a>func (i ProgType) String() string   { return stringName(uint32(i), ptStrings, false) }
<a id="L455"></a>func (i ProgType) GoString() string { return stringName(uint32(i), ptStrings, true) }

<a id="L457"></a><span class="comment">// Prog.Flag</span>
<a id="L458"></a>type ProgFlag uint32

<a id="L460"></a>const (
    <a id="L461"></a>PF_X        ProgFlag = 0x1;        <span class="comment">/* Executable. */</span>
    <a id="L462"></a>PF_W        ProgFlag = 0x2;        <span class="comment">/* Writable. */</span>
    <a id="L463"></a>PF_R        ProgFlag = 0x4;        <span class="comment">/* Readable. */</span>
    <a id="L464"></a>PF_MASKOS   ProgFlag = 0x0ff00000; <span class="comment">/* Operating system-specific. */</span>
    <a id="L465"></a>PF_MASKPROC ProgFlag = 0xf0000000; <span class="comment">/* Processor-specific. */</span>
<a id="L466"></a>)

<a id="L468"></a>var pfStrings = []intName{
    <a id="L469"></a>intName{0x1, &#34;PF_X&#34;},
    <a id="L470"></a>intName{0x2, &#34;PF_W&#34;},
    <a id="L471"></a>intName{0x4, &#34;PF_R&#34;},
<a id="L472"></a>}

<a id="L474"></a>func (i ProgFlag) String() string   { return flagName(uint32(i), pfStrings, false) }
<a id="L475"></a>func (i ProgFlag) GoString() string { return flagName(uint32(i), pfStrings, true) }

<a id="L477"></a><span class="comment">// Dyn.Tag</span>
<a id="L478"></a>type DynTag int

<a id="L480"></a>const (
    <a id="L481"></a>DT_NULL         DynTag = 0;  <span class="comment">/* Terminating entry. */</span>
    <a id="L482"></a>DT_NEEDED       DynTag = 1;  <span class="comment">/* String table offset of a needed shared library. */</span>
    <a id="L483"></a>DT_PLTRELSZ     DynTag = 2;  <span class="comment">/* Total size in bytes of PLT relocations. */</span>
    <a id="L484"></a>DT_PLTGOT       DynTag = 3;  <span class="comment">/* Processor-dependent address. */</span>
    <a id="L485"></a>DT_HASH         DynTag = 4;  <span class="comment">/* Address of symbol hash table. */</span>
    <a id="L486"></a>DT_STRTAB       DynTag = 5;  <span class="comment">/* Address of string table. */</span>
    <a id="L487"></a>DT_SYMTAB       DynTag = 6;  <span class="comment">/* Address of symbol table. */</span>
    <a id="L488"></a>DT_RELA         DynTag = 7;  <span class="comment">/* Address of ElfNN_Rela relocations. */</span>
    <a id="L489"></a>DT_RELASZ       DynTag = 8;  <span class="comment">/* Total size of ElfNN_Rela relocations. */</span>
    <a id="L490"></a>DT_RELAENT      DynTag = 9;  <span class="comment">/* Size of each ElfNN_Rela relocation entry. */</span>
    <a id="L491"></a>DT_STRSZ        DynTag = 10; <span class="comment">/* Size of string table. */</span>
    <a id="L492"></a>DT_SYMENT       DynTag = 11; <span class="comment">/* Size of each symbol table entry. */</span>
    <a id="L493"></a>DT_INIT         DynTag = 12; <span class="comment">/* Address of initialization function. */</span>
    <a id="L494"></a>DT_FINI         DynTag = 13; <span class="comment">/* Address of finalization function. */</span>
    <a id="L495"></a>DT_SONAME       DynTag = 14; <span class="comment">/* String table offset of shared object name. */</span>
    <a id="L496"></a>DT_RPATH        DynTag = 15; <span class="comment">/* String table offset of library path. [sup] */</span>
    <a id="L497"></a>DT_SYMBOLIC     DynTag = 16; <span class="comment">/* Indicates &#34;symbolic&#34; linking. [sup] */</span>
    <a id="L498"></a>DT_REL          DynTag = 17; <span class="comment">/* Address of ElfNN_Rel relocations. */</span>
    <a id="L499"></a>DT_RELSZ        DynTag = 18; <span class="comment">/* Total size of ElfNN_Rel relocations. */</span>
    <a id="L500"></a>DT_RELENT       DynTag = 19; <span class="comment">/* Size of each ElfNN_Rel relocation. */</span>
    <a id="L501"></a>DT_PLTREL       DynTag = 20; <span class="comment">/* Type of relocation used for PLT. */</span>
    <a id="L502"></a>DT_DEBUG        DynTag = 21; <span class="comment">/* Reserved (not used). */</span>
    <a id="L503"></a>DT_TEXTREL      DynTag = 22; <span class="comment">/* Indicates there may be relocations in non-writable segments. [sup] */</span>
    <a id="L504"></a>DT_JMPREL       DynTag = 23; <span class="comment">/* Address of PLT relocations. */</span>
    <a id="L505"></a>DT_BIND_NOW     DynTag = 24; <span class="comment">/* [sup] */</span>
    <a id="L506"></a>DT_INIT_ARRAY   DynTag = 25; <span class="comment">/* Address of the array of pointers to initialization functions */</span>
    <a id="L507"></a>DT_FINI_ARRAY   DynTag = 26; <span class="comment">/* Address of the array of pointers to termination functions */</span>
    <a id="L508"></a>DT_INIT_ARRAYSZ DynTag = 27; <span class="comment">/* Size in bytes of the array of initialization functions. */</span>
    <a id="L509"></a>DT_FINI_ARRAYSZ DynTag = 28; <span class="comment">/* Size in bytes of the array of terminationfunctions. */</span>
    <a id="L510"></a>DT_RUNPATH      DynTag = 29; <span class="comment">/* String table offset of a null-terminated library search path string. */</span>
    <a id="L511"></a>DT_FLAGS        DynTag = 30; <span class="comment">/* Object specific flag values. */</span>
    <a id="L512"></a>DT_ENCODING     DynTag = 32; <span class="comment">/* Values greater than or equal to DT_ENCODING</span>
    <a id="L513"></a><span class="comment">   and less than DT_LOOS follow the rules for</span>
    <a id="L514"></a><span class="comment">   the interpretation of the d_un union</span>
    <a id="L515"></a><span class="comment">   as follows: even == &#39;d_ptr&#39;, even == &#39;d_val&#39;</span>
    <a id="L516"></a><span class="comment">   or none */</span>
    <a id="L517"></a>DT_PREINIT_ARRAY   DynTag = 32;         <span class="comment">/* Address of the array of pointers to pre-initialization functions. */</span>
    <a id="L518"></a>DT_PREINIT_ARRAYSZ DynTag = 33;         <span class="comment">/* Size in bytes of the array of pre-initialization functions. */</span>
    <a id="L519"></a>DT_LOOS            DynTag = 0x6000000d; <span class="comment">/* First OS-specific */</span>
    <a id="L520"></a>DT_HIOS            DynTag = 0x6ffff000; <span class="comment">/* Last OS-specific */</span>
    <a id="L521"></a>DT_LOPROC          DynTag = 0x70000000; <span class="comment">/* First processor-specific type. */</span>
    <a id="L522"></a>DT_HIPROC          DynTag = 0x7fffffff; <span class="comment">/* Last processor-specific type. */</span>
<a id="L523"></a>)

<a id="L525"></a>var dtStrings = []intName{
    <a id="L526"></a>intName{0, &#34;DT_NULL&#34;},
    <a id="L527"></a>intName{1, &#34;DT_NEEDED&#34;},
    <a id="L528"></a>intName{2, &#34;DT_PLTRELSZ&#34;},
    <a id="L529"></a>intName{3, &#34;DT_PLTGOT&#34;},
    <a id="L530"></a>intName{4, &#34;DT_HASH&#34;},
    <a id="L531"></a>intName{5, &#34;DT_STRTAB&#34;},
    <a id="L532"></a>intName{6, &#34;DT_SYMTAB&#34;},
    <a id="L533"></a>intName{7, &#34;DT_RELA&#34;},
    <a id="L534"></a>intName{8, &#34;DT_RELASZ&#34;},
    <a id="L535"></a>intName{9, &#34;DT_RELAENT&#34;},
    <a id="L536"></a>intName{10, &#34;DT_STRSZ&#34;},
    <a id="L537"></a>intName{11, &#34;DT_SYMENT&#34;},
    <a id="L538"></a>intName{12, &#34;DT_INIT&#34;},
    <a id="L539"></a>intName{13, &#34;DT_FINI&#34;},
    <a id="L540"></a>intName{14, &#34;DT_SONAME&#34;},
    <a id="L541"></a>intName{15, &#34;DT_RPATH&#34;},
    <a id="L542"></a>intName{16, &#34;DT_SYMBOLIC&#34;},
    <a id="L543"></a>intName{17, &#34;DT_REL&#34;},
    <a id="L544"></a>intName{18, &#34;DT_RELSZ&#34;},
    <a id="L545"></a>intName{19, &#34;DT_RELENT&#34;},
    <a id="L546"></a>intName{20, &#34;DT_PLTREL&#34;},
    <a id="L547"></a>intName{21, &#34;DT_DEBUG&#34;},
    <a id="L548"></a>intName{22, &#34;DT_TEXTREL&#34;},
    <a id="L549"></a>intName{23, &#34;DT_JMPREL&#34;},
    <a id="L550"></a>intName{24, &#34;DT_BIND_NOW&#34;},
    <a id="L551"></a>intName{25, &#34;DT_INIT_ARRAY&#34;},
    <a id="L552"></a>intName{26, &#34;DT_FINI_ARRAY&#34;},
    <a id="L553"></a>intName{27, &#34;DT_INIT_ARRAYSZ&#34;},
    <a id="L554"></a>intName{28, &#34;DT_FINI_ARRAYSZ&#34;},
    <a id="L555"></a>intName{29, &#34;DT_RUNPATH&#34;},
    <a id="L556"></a>intName{30, &#34;DT_FLAGS&#34;},
    <a id="L557"></a>intName{32, &#34;DT_ENCODING&#34;},
    <a id="L558"></a>intName{32, &#34;DT_PREINIT_ARRAY&#34;},
    <a id="L559"></a>intName{33, &#34;DT_PREINIT_ARRAYSZ&#34;},
    <a id="L560"></a>intName{0x6000000d, &#34;DT_LOOS&#34;},
    <a id="L561"></a>intName{0x6ffff000, &#34;DT_HIOS&#34;},
    <a id="L562"></a>intName{0x70000000, &#34;DT_LOPROC&#34;},
    <a id="L563"></a>intName{0x7fffffff, &#34;DT_HIPROC&#34;},
<a id="L564"></a>}

<a id="L566"></a>func (i DynTag) String() string   { return stringName(uint32(i), dtStrings, false) }
<a id="L567"></a>func (i DynTag) GoString() string { return stringName(uint32(i), dtStrings, true) }

<a id="L569"></a><span class="comment">// DT_FLAGS values.</span>
<a id="L570"></a>type DynFlag int

<a id="L572"></a>const (
    <a id="L573"></a>DF_ORIGIN DynFlag = 0x0001; <span class="comment">/* Indicates that the object being loaded may</span>
    <a id="L574"></a><span class="comment">   make reference to the</span>
    <a id="L575"></a><span class="comment">   $ORIGIN substitution string */</span>
    <a id="L576"></a>DF_SYMBOLIC DynFlag = 0x0002; <span class="comment">/* Indicates &#34;symbolic&#34; linking. */</span>
    <a id="L577"></a>DF_TEXTREL  DynFlag = 0x0004; <span class="comment">/* Indicates there may be relocations in non-writable segments. */</span>
    <a id="L578"></a>DF_BIND_NOW DynFlag = 0x0008; <span class="comment">/* Indicates that the dynamic linker should</span>
    <a id="L579"></a><span class="comment">   process all relocations for the object</span>
    <a id="L580"></a><span class="comment">   containing this entry before transferring</span>
    <a id="L581"></a><span class="comment">   control to the program. */</span>
    <a id="L582"></a>DF_STATIC_TLS DynFlag = 0x0010; <span class="comment">/* Indicates that the shared object or</span>
    <a id="L583"></a><span class="comment">   executable contains code using a static</span>
    <a id="L584"></a><span class="comment">   thread-local storage scheme. */</span>
<a id="L585"></a>)

<a id="L587"></a>var dflagStrings = []intName{
    <a id="L588"></a>intName{0x0001, &#34;DF_ORIGIN&#34;},
    <a id="L589"></a>intName{0x0002, &#34;DF_SYMBOLIC&#34;},
    <a id="L590"></a>intName{0x0004, &#34;DF_TEXTREL&#34;},
    <a id="L591"></a>intName{0x0008, &#34;DF_BIND_NOW&#34;},
    <a id="L592"></a>intName{0x0010, &#34;DF_STATIC_TLS&#34;},
<a id="L593"></a>}

<a id="L595"></a>func (i DynFlag) String() string   { return flagName(uint32(i), dflagStrings, false) }
<a id="L596"></a>func (i DynFlag) GoString() string { return flagName(uint32(i), dflagStrings, true) }

<a id="L598"></a><span class="comment">// NType values; used in core files.</span>
<a id="L599"></a>type NType int

<a id="L601"></a>const (
    <a id="L602"></a>NT_PRSTATUS NType = 1; <span class="comment">/* Process status. */</span>
    <a id="L603"></a>NT_FPREGSET NType = 2; <span class="comment">/* Floating point registers. */</span>
    <a id="L604"></a>NT_PRPSINFO NType = 3; <span class="comment">/* Process state info. */</span>
<a id="L605"></a>)

<a id="L607"></a>var ntypeStrings = []intName{
    <a id="L608"></a>intName{1, &#34;NT_PRSTATUS&#34;},
    <a id="L609"></a>intName{2, &#34;NT_FPREGSET&#34;},
    <a id="L610"></a>intName{3, &#34;NT_PRPSINFO&#34;},
<a id="L611"></a>}

<a id="L613"></a>func (i NType) String() string   { return stringName(uint32(i), ntypeStrings, false) }
<a id="L614"></a>func (i NType) GoString() string { return stringName(uint32(i), ntypeStrings, true) }

<a id="L616"></a><span class="comment">/* Symbol Binding - ELFNN_ST_BIND - st_info */</span>
<a id="L617"></a>type SymBind int

<a id="L619"></a>const (
    <a id="L620"></a>STB_LOCAL  SymBind = 0;  <span class="comment">/* Local symbol */</span>
    <a id="L621"></a>STB_GLOBAL SymBind = 1;  <span class="comment">/* Global symbol */</span>
    <a id="L622"></a>STB_WEAK   SymBind = 2;  <span class="comment">/* like global - lower precedence */</span>
    <a id="L623"></a>STB_LOOS   SymBind = 10; <span class="comment">/* Reserved range for operating system */</span>
    <a id="L624"></a>STB_HIOS   SymBind = 12; <span class="comment">/*   specific semantics. */</span>
    <a id="L625"></a>STB_LOPROC SymBind = 13; <span class="comment">/* reserved range for processor */</span>
    <a id="L626"></a>STB_HIPROC SymBind = 15; <span class="comment">/*   specific semantics. */</span>
<a id="L627"></a>)

<a id="L629"></a>var stbStrings = []intName{
    <a id="L630"></a>intName{0, &#34;STB_LOCAL&#34;},
    <a id="L631"></a>intName{1, &#34;STB_GLOBAL&#34;},
    <a id="L632"></a>intName{2, &#34;STB_WEAK&#34;},
    <a id="L633"></a>intName{10, &#34;STB_LOOS&#34;},
    <a id="L634"></a>intName{12, &#34;STB_HIOS&#34;},
    <a id="L635"></a>intName{13, &#34;STB_LOPROC&#34;},
    <a id="L636"></a>intName{15, &#34;STB_HIPROC&#34;},
<a id="L637"></a>}

<a id="L639"></a>func (i SymBind) String() string   { return stringName(uint32(i), stbStrings, false) }
<a id="L640"></a>func (i SymBind) GoString() string { return stringName(uint32(i), stbStrings, true) }

<a id="L642"></a><span class="comment">/* Symbol type - ELFNN_ST_TYPE - st_info */</span>
<a id="L643"></a>type SymType int

<a id="L645"></a>const (
    <a id="L646"></a>STT_NOTYPE  SymType = 0;  <span class="comment">/* Unspecified type. */</span>
    <a id="L647"></a>STT_OBJECT  SymType = 1;  <span class="comment">/* Data object. */</span>
    <a id="L648"></a>STT_FUNC    SymType = 2;  <span class="comment">/* Function. */</span>
    <a id="L649"></a>STT_SECTION SymType = 3;  <span class="comment">/* Section. */</span>
    <a id="L650"></a>STT_FILE    SymType = 4;  <span class="comment">/* Source file. */</span>
    <a id="L651"></a>STT_COMMON  SymType = 5;  <span class="comment">/* Uninitialized common block. */</span>
    <a id="L652"></a>STT_TLS     SymType = 6;  <span class="comment">/* TLS object. */</span>
    <a id="L653"></a>STT_LOOS    SymType = 10; <span class="comment">/* Reserved range for operating system */</span>
    <a id="L654"></a>STT_HIOS    SymType = 12; <span class="comment">/*   specific semantics. */</span>
    <a id="L655"></a>STT_LOPROC  SymType = 13; <span class="comment">/* reserved range for processor */</span>
    <a id="L656"></a>STT_HIPROC  SymType = 15; <span class="comment">/*   specific semantics. */</span>
<a id="L657"></a>)

<a id="L659"></a>var sttStrings = []intName{
    <a id="L660"></a>intName{0, &#34;STT_NOTYPE&#34;},
    <a id="L661"></a>intName{1, &#34;STT_OBJECT&#34;},
    <a id="L662"></a>intName{2, &#34;STT_FUNC&#34;},
    <a id="L663"></a>intName{3, &#34;STT_SECTION&#34;},
    <a id="L664"></a>intName{4, &#34;STT_FILE&#34;},
    <a id="L665"></a>intName{5, &#34;STT_COMMON&#34;},
    <a id="L666"></a>intName{6, &#34;STT_TLS&#34;},
    <a id="L667"></a>intName{10, &#34;STT_LOOS&#34;},
    <a id="L668"></a>intName{12, &#34;STT_HIOS&#34;},
    <a id="L669"></a>intName{13, &#34;STT_LOPROC&#34;},
    <a id="L670"></a>intName{15, &#34;STT_HIPROC&#34;},
<a id="L671"></a>}

<a id="L673"></a>func (i SymType) String() string   { return stringName(uint32(i), sttStrings, false) }
<a id="L674"></a>func (i SymType) GoString() string { return stringName(uint32(i), sttStrings, true) }

<a id="L676"></a><span class="comment">/* Symbol visibility - ELFNN_ST_VISIBILITY - st_other */</span>
<a id="L677"></a>type SymVis int

<a id="L679"></a>const (
    <a id="L680"></a>STV_DEFAULT   SymVis = 0x0; <span class="comment">/* Default visibility (see binding). */</span>
    <a id="L681"></a>STV_INTERNAL  SymVis = 0x1; <span class="comment">/* Special meaning in relocatable objects. */</span>
    <a id="L682"></a>STV_HIDDEN    SymVis = 0x2; <span class="comment">/* Not visible. */</span>
    <a id="L683"></a>STV_PROTECTED SymVis = 0x3; <span class="comment">/* Visible but not preemptible. */</span>
<a id="L684"></a>)

<a id="L686"></a>var stvStrings = []intName{
    <a id="L687"></a>intName{0x0, &#34;STV_DEFAULT&#34;},
    <a id="L688"></a>intName{0x1, &#34;STV_INTERNAL&#34;},
    <a id="L689"></a>intName{0x2, &#34;STV_HIDDEN&#34;},
    <a id="L690"></a>intName{0x3, &#34;STV_PROTECTED&#34;},
<a id="L691"></a>}

<a id="L693"></a>func (i SymVis) String() string   { return stringName(uint32(i), stvStrings, false) }
<a id="L694"></a>func (i SymVis) GoString() string { return stringName(uint32(i), stvStrings, true) }

<a id="L696"></a><span class="comment">/*</span>
<a id="L697"></a><span class="comment"> * Relocation types.</span>
<a id="L698"></a><span class="comment"> */</span>

<a id="L700"></a><span class="comment">// Relocation types for x86-64.</span>
<a id="L701"></a>type R_X86_64 int

<a id="L703"></a>const (
    <a id="L704"></a>R_X86_64_NONE     R_X86_64 = 0;  <span class="comment">/* No relocation. */</span>
    <a id="L705"></a>R_X86_64_64       R_X86_64 = 1;  <span class="comment">/* Add 64 bit symbol value. */</span>
    <a id="L706"></a>R_X86_64_PC32     R_X86_64 = 2;  <span class="comment">/* PC-relative 32 bit signed sym value. */</span>
    <a id="L707"></a>R_X86_64_GOT32    R_X86_64 = 3;  <span class="comment">/* PC-relative 32 bit GOT offset. */</span>
    <a id="L708"></a>R_X86_64_PLT32    R_X86_64 = 4;  <span class="comment">/* PC-relative 32 bit PLT offset. */</span>
    <a id="L709"></a>R_X86_64_COPY     R_X86_64 = 5;  <span class="comment">/* Copy data from shared object. */</span>
    <a id="L710"></a>R_X86_64_GLOB_DAT R_X86_64 = 6;  <span class="comment">/* Set GOT entry to data address. */</span>
    <a id="L711"></a>R_X86_64_JMP_SLOT R_X86_64 = 7;  <span class="comment">/* Set GOT entry to code address. */</span>
    <a id="L712"></a>R_X86_64_RELATIVE R_X86_64 = 8;  <span class="comment">/* Add load address of shared object. */</span>
    <a id="L713"></a>R_X86_64_GOTPCREL R_X86_64 = 9;  <span class="comment">/* Add 32 bit signed pcrel offset to GOT. */</span>
    <a id="L714"></a>R_X86_64_32       R_X86_64 = 10; <span class="comment">/* Add 32 bit zero extended symbol value */</span>
    <a id="L715"></a>R_X86_64_32S      R_X86_64 = 11; <span class="comment">/* Add 32 bit sign extended symbol value */</span>
    <a id="L716"></a>R_X86_64_16       R_X86_64 = 12; <span class="comment">/* Add 16 bit zero extended symbol value */</span>
    <a id="L717"></a>R_X86_64_PC16     R_X86_64 = 13; <span class="comment">/* Add 16 bit signed extended pc relative symbol value */</span>
    <a id="L718"></a>R_X86_64_8        R_X86_64 = 14; <span class="comment">/* Add 8 bit zero extended symbol value */</span>
    <a id="L719"></a>R_X86_64_PC8      R_X86_64 = 15; <span class="comment">/* Add 8 bit signed extended pc relative symbol value */</span>
    <a id="L720"></a>R_X86_64_DTPMOD64 R_X86_64 = 16; <span class="comment">/* ID of module containing symbol */</span>
    <a id="L721"></a>R_X86_64_DTPOFF64 R_X86_64 = 17; <span class="comment">/* Offset in TLS block */</span>
    <a id="L722"></a>R_X86_64_TPOFF64  R_X86_64 = 18; <span class="comment">/* Offset in static TLS block */</span>
    <a id="L723"></a>R_X86_64_TLSGD    R_X86_64 = 19; <span class="comment">/* PC relative offset to GD GOT entry */</span>
    <a id="L724"></a>R_X86_64_TLSLD    R_X86_64 = 20; <span class="comment">/* PC relative offset to LD GOT entry */</span>
    <a id="L725"></a>R_X86_64_DTPOFF32 R_X86_64 = 21; <span class="comment">/* Offset in TLS block */</span>
    <a id="L726"></a>R_X86_64_GOTTPOFF R_X86_64 = 22; <span class="comment">/* PC relative offset to IE GOT entry */</span>
    <a id="L727"></a>R_X86_64_TPOFF32  R_X86_64 = 23; <span class="comment">/* Offset in static TLS block */</span>
<a id="L728"></a>)

<a id="L730"></a>var rx86_64Strings = []intName{
    <a id="L731"></a>intName{0, &#34;R_X86_64_NONE&#34;},
    <a id="L732"></a>intName{1, &#34;R_X86_64_64&#34;},
    <a id="L733"></a>intName{2, &#34;R_X86_64_PC32&#34;},
    <a id="L734"></a>intName{3, &#34;R_X86_64_GOT32&#34;},
    <a id="L735"></a>intName{4, &#34;R_X86_64_PLT32&#34;},
    <a id="L736"></a>intName{5, &#34;R_X86_64_COPY&#34;},
    <a id="L737"></a>intName{6, &#34;R_X86_64_GLOB_DAT&#34;},
    <a id="L738"></a>intName{7, &#34;R_X86_64_JMP_SLOT&#34;},
    <a id="L739"></a>intName{8, &#34;R_X86_64_RELATIVE&#34;},
    <a id="L740"></a>intName{9, &#34;R_X86_64_GOTPCREL&#34;},
    <a id="L741"></a>intName{10, &#34;R_X86_64_32&#34;},
    <a id="L742"></a>intName{11, &#34;R_X86_64_32S&#34;},
    <a id="L743"></a>intName{12, &#34;R_X86_64_16&#34;},
    <a id="L744"></a>intName{13, &#34;R_X86_64_PC16&#34;},
    <a id="L745"></a>intName{14, &#34;R_X86_64_8&#34;},
    <a id="L746"></a>intName{15, &#34;R_X86_64_PC8&#34;},
    <a id="L747"></a>intName{16, &#34;R_X86_64_DTPMOD64&#34;},
    <a id="L748"></a>intName{17, &#34;R_X86_64_DTPOFF64&#34;},
    <a id="L749"></a>intName{18, &#34;R_X86_64_TPOFF64&#34;},
    <a id="L750"></a>intName{19, &#34;R_X86_64_TLSGD&#34;},
    <a id="L751"></a>intName{20, &#34;R_X86_64_TLSLD&#34;},
    <a id="L752"></a>intName{21, &#34;R_X86_64_DTPOFF32&#34;},
    <a id="L753"></a>intName{22, &#34;R_X86_64_GOTTPOFF&#34;},
    <a id="L754"></a>intName{23, &#34;R_X86_64_TPOFF32&#34;},
<a id="L755"></a>}

<a id="L757"></a>func (i R_X86_64) String() string   { return stringName(uint32(i), rx86_64Strings, false) }
<a id="L758"></a>func (i R_X86_64) GoString() string { return stringName(uint32(i), rx86_64Strings, true) }

<a id="L760"></a><span class="comment">// Relocation types for Alpha.</span>
<a id="L761"></a>type R_ALPHA int

<a id="L763"></a>const (
    <a id="L764"></a>R_ALPHA_NONE           R_ALPHA = 0;  <span class="comment">/* No reloc */</span>
    <a id="L765"></a>R_ALPHA_REFLONG        R_ALPHA = 1;  <span class="comment">/* Direct 32 bit */</span>
    <a id="L766"></a>R_ALPHA_REFQUAD        R_ALPHA = 2;  <span class="comment">/* Direct 64 bit */</span>
    <a id="L767"></a>R_ALPHA_GPREL32        R_ALPHA = 3;  <span class="comment">/* GP relative 32 bit */</span>
    <a id="L768"></a>R_ALPHA_LITERAL        R_ALPHA = 4;  <span class="comment">/* GP relative 16 bit w/optimization */</span>
    <a id="L769"></a>R_ALPHA_LITUSE         R_ALPHA = 5;  <span class="comment">/* Optimization hint for LITERAL */</span>
    <a id="L770"></a>R_ALPHA_GPDISP         R_ALPHA = 6;  <span class="comment">/* Add displacement to GP */</span>
    <a id="L771"></a>R_ALPHA_BRADDR         R_ALPHA = 7;  <span class="comment">/* PC+4 relative 23 bit shifted */</span>
    <a id="L772"></a>R_ALPHA_HINT           R_ALPHA = 8;  <span class="comment">/* PC+4 relative 16 bit shifted */</span>
    <a id="L773"></a>R_ALPHA_SREL16         R_ALPHA = 9;  <span class="comment">/* PC relative 16 bit */</span>
    <a id="L774"></a>R_ALPHA_SREL32         R_ALPHA = 10; <span class="comment">/* PC relative 32 bit */</span>
    <a id="L775"></a>R_ALPHA_SREL64         R_ALPHA = 11; <span class="comment">/* PC relative 64 bit */</span>
    <a id="L776"></a>R_ALPHA_OP_PUSH        R_ALPHA = 12; <span class="comment">/* OP stack push */</span>
    <a id="L777"></a>R_ALPHA_OP_STORE       R_ALPHA = 13; <span class="comment">/* OP stack pop and store */</span>
    <a id="L778"></a>R_ALPHA_OP_PSUB        R_ALPHA = 14; <span class="comment">/* OP stack subtract */</span>
    <a id="L779"></a>R_ALPHA_OP_PRSHIFT     R_ALPHA = 15; <span class="comment">/* OP stack right shift */</span>
    <a id="L780"></a>R_ALPHA_GPVALUE        R_ALPHA = 16;
    <a id="L781"></a>R_ALPHA_GPRELHIGH      R_ALPHA = 17;
    <a id="L782"></a>R_ALPHA_GPRELLOW       R_ALPHA = 18;
    <a id="L783"></a>R_ALPHA_IMMED_GP_16    R_ALPHA = 19;
    <a id="L784"></a>R_ALPHA_IMMED_GP_HI32  R_ALPHA = 20;
    <a id="L785"></a>R_ALPHA_IMMED_SCN_HI32 R_ALPHA = 21;
    <a id="L786"></a>R_ALPHA_IMMED_BR_HI32  R_ALPHA = 22;
    <a id="L787"></a>R_ALPHA_IMMED_LO32     R_ALPHA = 23;
    <a id="L788"></a>R_ALPHA_COPY           R_ALPHA = 24; <span class="comment">/* Copy symbol at runtime */</span>
    <a id="L789"></a>R_ALPHA_GLOB_DAT       R_ALPHA = 25; <span class="comment">/* Create GOT entry */</span>
    <a id="L790"></a>R_ALPHA_JMP_SLOT       R_ALPHA = 26; <span class="comment">/* Create PLT entry */</span>
    <a id="L791"></a>R_ALPHA_RELATIVE       R_ALPHA = 27; <span class="comment">/* Adjust by program base */</span>
<a id="L792"></a>)

<a id="L794"></a>var ralphaStrings = []intName{
    <a id="L795"></a>intName{0, &#34;R_ALPHA_NONE&#34;},
    <a id="L796"></a>intName{1, &#34;R_ALPHA_REFLONG&#34;},
    <a id="L797"></a>intName{2, &#34;R_ALPHA_REFQUAD&#34;},
    <a id="L798"></a>intName{3, &#34;R_ALPHA_GPREL32&#34;},
    <a id="L799"></a>intName{4, &#34;R_ALPHA_LITERAL&#34;},
    <a id="L800"></a>intName{5, &#34;R_ALPHA_LITUSE&#34;},
    <a id="L801"></a>intName{6, &#34;R_ALPHA_GPDISP&#34;},
    <a id="L802"></a>intName{7, &#34;R_ALPHA_BRADDR&#34;},
    <a id="L803"></a>intName{8, &#34;R_ALPHA_HINT&#34;},
    <a id="L804"></a>intName{9, &#34;R_ALPHA_SREL16&#34;},
    <a id="L805"></a>intName{10, &#34;R_ALPHA_SREL32&#34;},
    <a id="L806"></a>intName{11, &#34;R_ALPHA_SREL64&#34;},
    <a id="L807"></a>intName{12, &#34;R_ALPHA_OP_PUSH&#34;},
    <a id="L808"></a>intName{13, &#34;R_ALPHA_OP_STORE&#34;},
    <a id="L809"></a>intName{14, &#34;R_ALPHA_OP_PSUB&#34;},
    <a id="L810"></a>intName{15, &#34;R_ALPHA_OP_PRSHIFT&#34;},
    <a id="L811"></a>intName{16, &#34;R_ALPHA_GPVALUE&#34;},
    <a id="L812"></a>intName{17, &#34;R_ALPHA_GPRELHIGH&#34;},
    <a id="L813"></a>intName{18, &#34;R_ALPHA_GPRELLOW&#34;},
    <a id="L814"></a>intName{19, &#34;R_ALPHA_IMMED_GP_16&#34;},
    <a id="L815"></a>intName{20, &#34;R_ALPHA_IMMED_GP_HI32&#34;},
    <a id="L816"></a>intName{21, &#34;R_ALPHA_IMMED_SCN_HI32&#34;},
    <a id="L817"></a>intName{22, &#34;R_ALPHA_IMMED_BR_HI32&#34;},
    <a id="L818"></a>intName{23, &#34;R_ALPHA_IMMED_LO32&#34;},
    <a id="L819"></a>intName{24, &#34;R_ALPHA_COPY&#34;},
    <a id="L820"></a>intName{25, &#34;R_ALPHA_GLOB_DAT&#34;},
    <a id="L821"></a>intName{26, &#34;R_ALPHA_JMP_SLOT&#34;},
    <a id="L822"></a>intName{27, &#34;R_ALPHA_RELATIVE&#34;},
<a id="L823"></a>}

<a id="L825"></a>func (i R_ALPHA) String() string   { return stringName(uint32(i), ralphaStrings, false) }
<a id="L826"></a>func (i R_ALPHA) GoString() string { return stringName(uint32(i), ralphaStrings, true) }

<a id="L828"></a><span class="comment">// Relocation types for ARM.</span>
<a id="L829"></a>type R_ARM int

<a id="L831"></a>const (
    <a id="L832"></a>R_ARM_NONE          R_ARM = 0; <span class="comment">/* No relocation. */</span>
    <a id="L833"></a>R_ARM_PC24          R_ARM = 1;
    <a id="L834"></a>R_ARM_ABS32         R_ARM = 2;
    <a id="L835"></a>R_ARM_REL32         R_ARM = 3;
    <a id="L836"></a>R_ARM_PC13          R_ARM = 4;
    <a id="L837"></a>R_ARM_ABS16         R_ARM = 5;
    <a id="L838"></a>R_ARM_ABS12         R_ARM = 6;
    <a id="L839"></a>R_ARM_THM_ABS5      R_ARM = 7;
    <a id="L840"></a>R_ARM_ABS8          R_ARM = 8;
    <a id="L841"></a>R_ARM_SBREL32       R_ARM = 9;
    <a id="L842"></a>R_ARM_THM_PC22      R_ARM = 10;
    <a id="L843"></a>R_ARM_THM_PC8       R_ARM = 11;
    <a id="L844"></a>R_ARM_AMP_VCALL9    R_ARM = 12;
    <a id="L845"></a>R_ARM_SWI24         R_ARM = 13;
    <a id="L846"></a>R_ARM_THM_SWI8      R_ARM = 14;
    <a id="L847"></a>R_ARM_XPC25         R_ARM = 15;
    <a id="L848"></a>R_ARM_THM_XPC22     R_ARM = 16;
    <a id="L849"></a>R_ARM_COPY          R_ARM = 20; <span class="comment">/* Copy data from shared object. */</span>
    <a id="L850"></a>R_ARM_GLOB_DAT      R_ARM = 21; <span class="comment">/* Set GOT entry to data address. */</span>
    <a id="L851"></a>R_ARM_JUMP_SLOT     R_ARM = 22; <span class="comment">/* Set GOT entry to code address. */</span>
    <a id="L852"></a>R_ARM_RELATIVE      R_ARM = 23; <span class="comment">/* Add load address of shared object. */</span>
    <a id="L853"></a>R_ARM_GOTOFF        R_ARM = 24; <span class="comment">/* Add GOT-relative symbol address. */</span>
    <a id="L854"></a>R_ARM_GOTPC         R_ARM = 25; <span class="comment">/* Add PC-relative GOT table address. */</span>
    <a id="L855"></a>R_ARM_GOT32         R_ARM = 26; <span class="comment">/* Add PC-relative GOT offset. */</span>
    <a id="L856"></a>R_ARM_PLT32         R_ARM = 27; <span class="comment">/* Add PC-relative PLT offset. */</span>
    <a id="L857"></a>R_ARM_GNU_VTENTRY   R_ARM = 100;
    <a id="L858"></a>R_ARM_GNU_VTINHERIT R_ARM = 101;
    <a id="L859"></a>R_ARM_RSBREL32      R_ARM = 250;
    <a id="L860"></a>R_ARM_THM_RPC22     R_ARM = 251;
    <a id="L861"></a>R_ARM_RREL32        R_ARM = 252;
    <a id="L862"></a>R_ARM_RABS32        R_ARM = 253;
    <a id="L863"></a>R_ARM_RPC24         R_ARM = 254;
    <a id="L864"></a>R_ARM_RBASE         R_ARM = 255;
<a id="L865"></a>)

<a id="L867"></a>var rarmStrings = []intName{
    <a id="L868"></a>intName{0, &#34;R_ARM_NONE&#34;},
    <a id="L869"></a>intName{1, &#34;R_ARM_PC24&#34;},
    <a id="L870"></a>intName{2, &#34;R_ARM_ABS32&#34;},
    <a id="L871"></a>intName{3, &#34;R_ARM_REL32&#34;},
    <a id="L872"></a>intName{4, &#34;R_ARM_PC13&#34;},
    <a id="L873"></a>intName{5, &#34;R_ARM_ABS16&#34;},
    <a id="L874"></a>intName{6, &#34;R_ARM_ABS12&#34;},
    <a id="L875"></a>intName{7, &#34;R_ARM_THM_ABS5&#34;},
    <a id="L876"></a>intName{8, &#34;R_ARM_ABS8&#34;},
    <a id="L877"></a>intName{9, &#34;R_ARM_SBREL32&#34;},
    <a id="L878"></a>intName{10, &#34;R_ARM_THM_PC22&#34;},
    <a id="L879"></a>intName{11, &#34;R_ARM_THM_PC8&#34;},
    <a id="L880"></a>intName{12, &#34;R_ARM_AMP_VCALL9&#34;},
    <a id="L881"></a>intName{13, &#34;R_ARM_SWI24&#34;},
    <a id="L882"></a>intName{14, &#34;R_ARM_THM_SWI8&#34;},
    <a id="L883"></a>intName{15, &#34;R_ARM_XPC25&#34;},
    <a id="L884"></a>intName{16, &#34;R_ARM_THM_XPC22&#34;},
    <a id="L885"></a>intName{20, &#34;R_ARM_COPY&#34;},
    <a id="L886"></a>intName{21, &#34;R_ARM_GLOB_DAT&#34;},
    <a id="L887"></a>intName{22, &#34;R_ARM_JUMP_SLOT&#34;},
    <a id="L888"></a>intName{23, &#34;R_ARM_RELATIVE&#34;},
    <a id="L889"></a>intName{24, &#34;R_ARM_GOTOFF&#34;},
    <a id="L890"></a>intName{25, &#34;R_ARM_GOTPC&#34;},
    <a id="L891"></a>intName{26, &#34;R_ARM_GOT32&#34;},
    <a id="L892"></a>intName{27, &#34;R_ARM_PLT32&#34;},
    <a id="L893"></a>intName{100, &#34;R_ARM_GNU_VTENTRY&#34;},
    <a id="L894"></a>intName{101, &#34;R_ARM_GNU_VTINHERIT&#34;},
    <a id="L895"></a>intName{250, &#34;R_ARM_RSBREL32&#34;},
    <a id="L896"></a>intName{251, &#34;R_ARM_THM_RPC22&#34;},
    <a id="L897"></a>intName{252, &#34;R_ARM_RREL32&#34;},
    <a id="L898"></a>intName{253, &#34;R_ARM_RABS32&#34;},
    <a id="L899"></a>intName{254, &#34;R_ARM_RPC24&#34;},
    <a id="L900"></a>intName{255, &#34;R_ARM_RBASE&#34;},
<a id="L901"></a>}

<a id="L903"></a>func (i R_ARM) String() string   { return stringName(uint32(i), rarmStrings, false) }
<a id="L904"></a>func (i R_ARM) GoString() string { return stringName(uint32(i), rarmStrings, true) }

<a id="L906"></a><span class="comment">// Relocation types for 386.</span>
<a id="L907"></a>type R_386 int

<a id="L909"></a>const (
    <a id="L910"></a>R_386_NONE         R_386 = 0;  <span class="comment">/* No relocation. */</span>
    <a id="L911"></a>R_386_32           R_386 = 1;  <span class="comment">/* Add symbol value. */</span>
    <a id="L912"></a>R_386_PC32         R_386 = 2;  <span class="comment">/* Add PC-relative symbol value. */</span>
    <a id="L913"></a>R_386_GOT32        R_386 = 3;  <span class="comment">/* Add PC-relative GOT offset. */</span>
    <a id="L914"></a>R_386_PLT32        R_386 = 4;  <span class="comment">/* Add PC-relative PLT offset. */</span>
    <a id="L915"></a>R_386_COPY         R_386 = 5;  <span class="comment">/* Copy data from shared object. */</span>
    <a id="L916"></a>R_386_GLOB_DAT     R_386 = 6;  <span class="comment">/* Set GOT entry to data address. */</span>
    <a id="L917"></a>R_386_JMP_SLOT     R_386 = 7;  <span class="comment">/* Set GOT entry to code address. */</span>
    <a id="L918"></a>R_386_RELATIVE     R_386 = 8;  <span class="comment">/* Add load address of shared object. */</span>
    <a id="L919"></a>R_386_GOTOFF       R_386 = 9;  <span class="comment">/* Add GOT-relative symbol address. */</span>
    <a id="L920"></a>R_386_GOTPC        R_386 = 10; <span class="comment">/* Add PC-relative GOT table address. */</span>
    <a id="L921"></a>R_386_TLS_TPOFF    R_386 = 14; <span class="comment">/* Negative offset in static TLS block */</span>
    <a id="L922"></a>R_386_TLS_IE       R_386 = 15; <span class="comment">/* Absolute address of GOT for -ve static TLS */</span>
    <a id="L923"></a>R_386_TLS_GOTIE    R_386 = 16; <span class="comment">/* GOT entry for negative static TLS block */</span>
    <a id="L924"></a>R_386_TLS_LE       R_386 = 17; <span class="comment">/* Negative offset relative to static TLS */</span>
    <a id="L925"></a>R_386_TLS_GD       R_386 = 18; <span class="comment">/* 32 bit offset to GOT (index,off) pair */</span>
    <a id="L926"></a>R_386_TLS_LDM      R_386 = 19; <span class="comment">/* 32 bit offset to GOT (index,zero) pair */</span>
    <a id="L927"></a>R_386_TLS_GD_32    R_386 = 24; <span class="comment">/* 32 bit offset to GOT (index,off) pair */</span>
    <a id="L928"></a>R_386_TLS_GD_PUSH  R_386 = 25; <span class="comment">/* pushl instruction for Sun ABI GD sequence */</span>
    <a id="L929"></a>R_386_TLS_GD_CALL  R_386 = 26; <span class="comment">/* call instruction for Sun ABI GD sequence */</span>
    <a id="L930"></a>R_386_TLS_GD_POP   R_386 = 27; <span class="comment">/* popl instruction for Sun ABI GD sequence */</span>
    <a id="L931"></a>R_386_TLS_LDM_32   R_386 = 28; <span class="comment">/* 32 bit offset to GOT (index,zero) pair */</span>
    <a id="L932"></a>R_386_TLS_LDM_PUSH R_386 = 29; <span class="comment">/* pushl instruction for Sun ABI LD sequence */</span>
    <a id="L933"></a>R_386_TLS_LDM_CALL R_386 = 30; <span class="comment">/* call instruction for Sun ABI LD sequence */</span>
    <a id="L934"></a>R_386_TLS_LDM_POP  R_386 = 31; <span class="comment">/* popl instruction for Sun ABI LD sequence */</span>
    <a id="L935"></a>R_386_TLS_LDO_32   R_386 = 32; <span class="comment">/* 32 bit offset from start of TLS block */</span>
    <a id="L936"></a>R_386_TLS_IE_32    R_386 = 33; <span class="comment">/* 32 bit offset to GOT static TLS offset entry */</span>
    <a id="L937"></a>R_386_TLS_LE_32    R_386 = 34; <span class="comment">/* 32 bit offset within static TLS block */</span>
    <a id="L938"></a>R_386_TLS_DTPMOD32 R_386 = 35; <span class="comment">/* GOT entry containing TLS index */</span>
    <a id="L939"></a>R_386_TLS_DTPOFF32 R_386 = 36; <span class="comment">/* GOT entry containing TLS offset */</span>
    <a id="L940"></a>R_386_TLS_TPOFF32  R_386 = 37; <span class="comment">/* GOT entry of -ve static TLS offset */</span>
<a id="L941"></a>)

<a id="L943"></a>var r386Strings = []intName{
    <a id="L944"></a>intName{0, &#34;R_386_NONE&#34;},
    <a id="L945"></a>intName{1, &#34;R_386_32&#34;},
    <a id="L946"></a>intName{2, &#34;R_386_PC32&#34;},
    <a id="L947"></a>intName{3, &#34;R_386_GOT32&#34;},
    <a id="L948"></a>intName{4, &#34;R_386_PLT32&#34;},
    <a id="L949"></a>intName{5, &#34;R_386_COPY&#34;},
    <a id="L950"></a>intName{6, &#34;R_386_GLOB_DAT&#34;},
    <a id="L951"></a>intName{7, &#34;R_386_JMP_SLOT&#34;},
    <a id="L952"></a>intName{8, &#34;R_386_RELATIVE&#34;},
    <a id="L953"></a>intName{9, &#34;R_386_GOTOFF&#34;},
    <a id="L954"></a>intName{10, &#34;R_386_GOTPC&#34;},
    <a id="L955"></a>intName{14, &#34;R_386_TLS_TPOFF&#34;},
    <a id="L956"></a>intName{15, &#34;R_386_TLS_IE&#34;},
    <a id="L957"></a>intName{16, &#34;R_386_TLS_GOTIE&#34;},
    <a id="L958"></a>intName{17, &#34;R_386_TLS_LE&#34;},
    <a id="L959"></a>intName{18, &#34;R_386_TLS_GD&#34;},
    <a id="L960"></a>intName{19, &#34;R_386_TLS_LDM&#34;},
    <a id="L961"></a>intName{24, &#34;R_386_TLS_GD_32&#34;},
    <a id="L962"></a>intName{25, &#34;R_386_TLS_GD_PUSH&#34;},
    <a id="L963"></a>intName{26, &#34;R_386_TLS_GD_CALL&#34;},
    <a id="L964"></a>intName{27, &#34;R_386_TLS_GD_POP&#34;},
    <a id="L965"></a>intName{28, &#34;R_386_TLS_LDM_32&#34;},
    <a id="L966"></a>intName{29, &#34;R_386_TLS_LDM_PUSH&#34;},
    <a id="L967"></a>intName{30, &#34;R_386_TLS_LDM_CALL&#34;},
    <a id="L968"></a>intName{31, &#34;R_386_TLS_LDM_POP&#34;},
    <a id="L969"></a>intName{32, &#34;R_386_TLS_LDO_32&#34;},
    <a id="L970"></a>intName{33, &#34;R_386_TLS_IE_32&#34;},
    <a id="L971"></a>intName{34, &#34;R_386_TLS_LE_32&#34;},
    <a id="L972"></a>intName{35, &#34;R_386_TLS_DTPMOD32&#34;},
    <a id="L973"></a>intName{36, &#34;R_386_TLS_DTPOFF32&#34;},
    <a id="L974"></a>intName{37, &#34;R_386_TLS_TPOFF32&#34;},
<a id="L975"></a>}

<a id="L977"></a>func (i R_386) String() string   { return stringName(uint32(i), r386Strings, false) }
<a id="L978"></a>func (i R_386) GoString() string { return stringName(uint32(i), r386Strings, true) }

<a id="L980"></a><span class="comment">// Relocation types for PowerPC.</span>
<a id="L981"></a>type R_PPC int

<a id="L983"></a>const (
    <a id="L984"></a>R_PPC_NONE            R_PPC = 0; <span class="comment">/* No relocation. */</span>
    <a id="L985"></a>R_PPC_ADDR32          R_PPC = 1;
    <a id="L986"></a>R_PPC_ADDR24          R_PPC = 2;
    <a id="L987"></a>R_PPC_ADDR16          R_PPC = 3;
    <a id="L988"></a>R_PPC_ADDR16_LO       R_PPC = 4;
    <a id="L989"></a>R_PPC_ADDR16_HI       R_PPC = 5;
    <a id="L990"></a>R_PPC_ADDR16_HA       R_PPC = 6;
    <a id="L991"></a>R_PPC_ADDR14          R_PPC = 7;
    <a id="L992"></a>R_PPC_ADDR14_BRTAKEN  R_PPC = 8;
    <a id="L993"></a>R_PPC_ADDR14_BRNTAKEN R_PPC = 9;
    <a id="L994"></a>R_PPC_REL24           R_PPC = 10;
    <a id="L995"></a>R_PPC_REL14           R_PPC = 11;
    <a id="L996"></a>R_PPC_REL14_BRTAKEN   R_PPC = 12;
    <a id="L997"></a>R_PPC_REL14_BRNTAKEN  R_PPC = 13;
    <a id="L998"></a>R_PPC_GOT16           R_PPC = 14;
    <a id="L999"></a>R_PPC_GOT16_LO        R_PPC = 15;
    <a id="L1000"></a>R_PPC_GOT16_HI        R_PPC = 16;
    <a id="L1001"></a>R_PPC_GOT16_HA        R_PPC = 17;
    <a id="L1002"></a>R_PPC_PLTREL24        R_PPC = 18;
    <a id="L1003"></a>R_PPC_COPY            R_PPC = 19;
    <a id="L1004"></a>R_PPC_GLOB_DAT        R_PPC = 20;
    <a id="L1005"></a>R_PPC_JMP_SLOT        R_PPC = 21;
    <a id="L1006"></a>R_PPC_RELATIVE        R_PPC = 22;
    <a id="L1007"></a>R_PPC_LOCAL24PC       R_PPC = 23;
    <a id="L1008"></a>R_PPC_UADDR32         R_PPC = 24;
    <a id="L1009"></a>R_PPC_UADDR16         R_PPC = 25;
    <a id="L1010"></a>R_PPC_REL32           R_PPC = 26;
    <a id="L1011"></a>R_PPC_PLT32           R_PPC = 27;
    <a id="L1012"></a>R_PPC_PLTREL32        R_PPC = 28;
    <a id="L1013"></a>R_PPC_PLT16_LO        R_PPC = 29;
    <a id="L1014"></a>R_PPC_PLT16_HI        R_PPC = 30;
    <a id="L1015"></a>R_PPC_PLT16_HA        R_PPC = 31;
    <a id="L1016"></a>R_PPC_SDAREL16        R_PPC = 32;
    <a id="L1017"></a>R_PPC_SECTOFF         R_PPC = 33;
    <a id="L1018"></a>R_PPC_SECTOFF_LO      R_PPC = 34;
    <a id="L1019"></a>R_PPC_SECTOFF_HI      R_PPC = 35;
    <a id="L1020"></a>R_PPC_SECTOFF_HA      R_PPC = 36;
    <a id="L1021"></a>R_PPC_TLS             R_PPC = 67;
    <a id="L1022"></a>R_PPC_DTPMOD32        R_PPC = 68;
    <a id="L1023"></a>R_PPC_TPREL16         R_PPC = 69;
    <a id="L1024"></a>R_PPC_TPREL16_LO      R_PPC = 70;
    <a id="L1025"></a>R_PPC_TPREL16_HI      R_PPC = 71;
    <a id="L1026"></a>R_PPC_TPREL16_HA      R_PPC = 72;
    <a id="L1027"></a>R_PPC_TPREL32         R_PPC = 73;
    <a id="L1028"></a>R_PPC_DTPREL16        R_PPC = 74;
    <a id="L1029"></a>R_PPC_DTPREL16_LO     R_PPC = 75;
    <a id="L1030"></a>R_PPC_DTPREL16_HI     R_PPC = 76;
    <a id="L1031"></a>R_PPC_DTPREL16_HA     R_PPC = 77;
    <a id="L1032"></a>R_PPC_DTPREL32        R_PPC = 78;
    <a id="L1033"></a>R_PPC_GOT_TLSGD16     R_PPC = 79;
    <a id="L1034"></a>R_PPC_GOT_TLSGD16_LO  R_PPC = 80;
    <a id="L1035"></a>R_PPC_GOT_TLSGD16_HI  R_PPC = 81;
    <a id="L1036"></a>R_PPC_GOT_TLSGD16_HA  R_PPC = 82;
    <a id="L1037"></a>R_PPC_GOT_TLSLD16     R_PPC = 83;
    <a id="L1038"></a>R_PPC_GOT_TLSLD16_LO  R_PPC = 84;
    <a id="L1039"></a>R_PPC_GOT_TLSLD16_HI  R_PPC = 85;
    <a id="L1040"></a>R_PPC_GOT_TLSLD16_HA  R_PPC = 86;
    <a id="L1041"></a>R_PPC_GOT_TPREL16     R_PPC = 87;
    <a id="L1042"></a>R_PPC_GOT_TPREL16_LO  R_PPC = 88;
    <a id="L1043"></a>R_PPC_GOT_TPREL16_HI  R_PPC = 89;
    <a id="L1044"></a>R_PPC_GOT_TPREL16_HA  R_PPC = 90;
    <a id="L1045"></a>R_PPC_EMB_NADDR32     R_PPC = 101;
    <a id="L1046"></a>R_PPC_EMB_NADDR16     R_PPC = 102;
    <a id="L1047"></a>R_PPC_EMB_NADDR16_LO  R_PPC = 103;
    <a id="L1048"></a>R_PPC_EMB_NADDR16_HI  R_PPC = 104;
    <a id="L1049"></a>R_PPC_EMB_NADDR16_HA  R_PPC = 105;
    <a id="L1050"></a>R_PPC_EMB_SDAI16      R_PPC = 106;
    <a id="L1051"></a>R_PPC_EMB_SDA2I16     R_PPC = 107;
    <a id="L1052"></a>R_PPC_EMB_SDA2REL     R_PPC = 108;
    <a id="L1053"></a>R_PPC_EMB_SDA21       R_PPC = 109;
    <a id="L1054"></a>R_PPC_EMB_MRKREF      R_PPC = 110;
    <a id="L1055"></a>R_PPC_EMB_RELSEC16    R_PPC = 111;
    <a id="L1056"></a>R_PPC_EMB_RELST_LO    R_PPC = 112;
    <a id="L1057"></a>R_PPC_EMB_RELST_HI    R_PPC = 113;
    <a id="L1058"></a>R_PPC_EMB_RELST_HA    R_PPC = 114;
    <a id="L1059"></a>R_PPC_EMB_BIT_FLD     R_PPC = 115;
    <a id="L1060"></a>R_PPC_EMB_RELSDA      R_PPC = 116;
<a id="L1061"></a>)

<a id="L1063"></a>var rppcStrings = []intName{
    <a id="L1064"></a>intName{0, &#34;R_PPC_NONE&#34;},
    <a id="L1065"></a>intName{1, &#34;R_PPC_ADDR32&#34;},
    <a id="L1066"></a>intName{2, &#34;R_PPC_ADDR24&#34;},
    <a id="L1067"></a>intName{3, &#34;R_PPC_ADDR16&#34;},
    <a id="L1068"></a>intName{4, &#34;R_PPC_ADDR16_LO&#34;},
    <a id="L1069"></a>intName{5, &#34;R_PPC_ADDR16_HI&#34;},
    <a id="L1070"></a>intName{6, &#34;R_PPC_ADDR16_HA&#34;},
    <a id="L1071"></a>intName{7, &#34;R_PPC_ADDR14&#34;},
    <a id="L1072"></a>intName{8, &#34;R_PPC_ADDR14_BRTAKEN&#34;},
    <a id="L1073"></a>intName{9, &#34;R_PPC_ADDR14_BRNTAKEN&#34;},
    <a id="L1074"></a>intName{10, &#34;R_PPC_REL24&#34;},
    <a id="L1075"></a>intName{11, &#34;R_PPC_REL14&#34;},
    <a id="L1076"></a>intName{12, &#34;R_PPC_REL14_BRTAKEN&#34;},
    <a id="L1077"></a>intName{13, &#34;R_PPC_REL14_BRNTAKEN&#34;},
    <a id="L1078"></a>intName{14, &#34;R_PPC_GOT16&#34;},
    <a id="L1079"></a>intName{15, &#34;R_PPC_GOT16_LO&#34;},
    <a id="L1080"></a>intName{16, &#34;R_PPC_GOT16_HI&#34;},
    <a id="L1081"></a>intName{17, &#34;R_PPC_GOT16_HA&#34;},
    <a id="L1082"></a>intName{18, &#34;R_PPC_PLTREL24&#34;},
    <a id="L1083"></a>intName{19, &#34;R_PPC_COPY&#34;},
    <a id="L1084"></a>intName{20, &#34;R_PPC_GLOB_DAT&#34;},
    <a id="L1085"></a>intName{21, &#34;R_PPC_JMP_SLOT&#34;},
    <a id="L1086"></a>intName{22, &#34;R_PPC_RELATIVE&#34;},
    <a id="L1087"></a>intName{23, &#34;R_PPC_LOCAL24PC&#34;},
    <a id="L1088"></a>intName{24, &#34;R_PPC_UADDR32&#34;},
    <a id="L1089"></a>intName{25, &#34;R_PPC_UADDR16&#34;},
    <a id="L1090"></a>intName{26, &#34;R_PPC_REL32&#34;},
    <a id="L1091"></a>intName{27, &#34;R_PPC_PLT32&#34;},
    <a id="L1092"></a>intName{28, &#34;R_PPC_PLTREL32&#34;},
    <a id="L1093"></a>intName{29, &#34;R_PPC_PLT16_LO&#34;},
    <a id="L1094"></a>intName{30, &#34;R_PPC_PLT16_HI&#34;},
    <a id="L1095"></a>intName{31, &#34;R_PPC_PLT16_HA&#34;},
    <a id="L1096"></a>intName{32, &#34;R_PPC_SDAREL16&#34;},
    <a id="L1097"></a>intName{33, &#34;R_PPC_SECTOFF&#34;},
    <a id="L1098"></a>intName{34, &#34;R_PPC_SECTOFF_LO&#34;},
    <a id="L1099"></a>intName{35, &#34;R_PPC_SECTOFF_HI&#34;},
    <a id="L1100"></a>intName{36, &#34;R_PPC_SECTOFF_HA&#34;},

    <a id="L1102"></a>intName{67, &#34;R_PPC_TLS&#34;},
    <a id="L1103"></a>intName{68, &#34;R_PPC_DTPMOD32&#34;},
    <a id="L1104"></a>intName{69, &#34;R_PPC_TPREL16&#34;},
    <a id="L1105"></a>intName{70, &#34;R_PPC_TPREL16_LO&#34;},
    <a id="L1106"></a>intName{71, &#34;R_PPC_TPREL16_HI&#34;},
    <a id="L1107"></a>intName{72, &#34;R_PPC_TPREL16_HA&#34;},
    <a id="L1108"></a>intName{73, &#34;R_PPC_TPREL32&#34;},
    <a id="L1109"></a>intName{74, &#34;R_PPC_DTPREL16&#34;},
    <a id="L1110"></a>intName{75, &#34;R_PPC_DTPREL16_LO&#34;},
    <a id="L1111"></a>intName{76, &#34;R_PPC_DTPREL16_HI&#34;},
    <a id="L1112"></a>intName{77, &#34;R_PPC_DTPREL16_HA&#34;},
    <a id="L1113"></a>intName{78, &#34;R_PPC_DTPREL32&#34;},
    <a id="L1114"></a>intName{79, &#34;R_PPC_GOT_TLSGD16&#34;},
    <a id="L1115"></a>intName{80, &#34;R_PPC_GOT_TLSGD16_LO&#34;},
    <a id="L1116"></a>intName{81, &#34;R_PPC_GOT_TLSGD16_HI&#34;},
    <a id="L1117"></a>intName{82, &#34;R_PPC_GOT_TLSGD16_HA&#34;},
    <a id="L1118"></a>intName{83, &#34;R_PPC_GOT_TLSLD16&#34;},
    <a id="L1119"></a>intName{84, &#34;R_PPC_GOT_TLSLD16_LO&#34;},
    <a id="L1120"></a>intName{85, &#34;R_PPC_GOT_TLSLD16_HI&#34;},
    <a id="L1121"></a>intName{86, &#34;R_PPC_GOT_TLSLD16_HA&#34;},
    <a id="L1122"></a>intName{87, &#34;R_PPC_GOT_TPREL16&#34;},
    <a id="L1123"></a>intName{88, &#34;R_PPC_GOT_TPREL16_LO&#34;},
    <a id="L1124"></a>intName{89, &#34;R_PPC_GOT_TPREL16_HI&#34;},
    <a id="L1125"></a>intName{90, &#34;R_PPC_GOT_TPREL16_HA&#34;},

    <a id="L1127"></a>intName{101, &#34;R_PPC_EMB_NADDR32&#34;},
    <a id="L1128"></a>intName{102, &#34;R_PPC_EMB_NADDR16&#34;},
    <a id="L1129"></a>intName{103, &#34;R_PPC_EMB_NADDR16_LO&#34;},
    <a id="L1130"></a>intName{104, &#34;R_PPC_EMB_NADDR16_HI&#34;},
    <a id="L1131"></a>intName{105, &#34;R_PPC_EMB_NADDR16_HA&#34;},
    <a id="L1132"></a>intName{106, &#34;R_PPC_EMB_SDAI16&#34;},
    <a id="L1133"></a>intName{107, &#34;R_PPC_EMB_SDA2I16&#34;},
    <a id="L1134"></a>intName{108, &#34;R_PPC_EMB_SDA2REL&#34;},
    <a id="L1135"></a>intName{109, &#34;R_PPC_EMB_SDA21&#34;},
    <a id="L1136"></a>intName{110, &#34;R_PPC_EMB_MRKREF&#34;},
    <a id="L1137"></a>intName{111, &#34;R_PPC_EMB_RELSEC16&#34;},
    <a id="L1138"></a>intName{112, &#34;R_PPC_EMB_RELST_LO&#34;},
    <a id="L1139"></a>intName{113, &#34;R_PPC_EMB_RELST_HI&#34;},
    <a id="L1140"></a>intName{114, &#34;R_PPC_EMB_RELST_HA&#34;},
    <a id="L1141"></a>intName{115, &#34;R_PPC_EMB_BIT_FLD&#34;},
    <a id="L1142"></a>intName{116, &#34;R_PPC_EMB_RELSDA&#34;},
<a id="L1143"></a>}

<a id="L1145"></a>func (i R_PPC) String() string   { return stringName(uint32(i), rppcStrings, false) }
<a id="L1146"></a>func (i R_PPC) GoString() string { return stringName(uint32(i), rppcStrings, true) }

<a id="L1148"></a><span class="comment">// Relocation types for SPARC.</span>
<a id="L1149"></a>type R_SPARC int

<a id="L1151"></a>const (
    <a id="L1152"></a>R_SPARC_NONE     R_SPARC = 0;
    <a id="L1153"></a>R_SPARC_8        R_SPARC = 1;
    <a id="L1154"></a>R_SPARC_16       R_SPARC = 2;
    <a id="L1155"></a>R_SPARC_32       R_SPARC = 3;
    <a id="L1156"></a>R_SPARC_DISP8    R_SPARC = 4;
    <a id="L1157"></a>R_SPARC_DISP16   R_SPARC = 5;
    <a id="L1158"></a>R_SPARC_DISP32   R_SPARC = 6;
    <a id="L1159"></a>R_SPARC_WDISP30  R_SPARC = 7;
    <a id="L1160"></a>R_SPARC_WDISP22  R_SPARC = 8;
    <a id="L1161"></a>R_SPARC_HI22     R_SPARC = 9;
    <a id="L1162"></a>R_SPARC_22       R_SPARC = 10;
    <a id="L1163"></a>R_SPARC_13       R_SPARC = 11;
    <a id="L1164"></a>R_SPARC_LO10     R_SPARC = 12;
    <a id="L1165"></a>R_SPARC_GOT10    R_SPARC = 13;
    <a id="L1166"></a>R_SPARC_GOT13    R_SPARC = 14;
    <a id="L1167"></a>R_SPARC_GOT22    R_SPARC = 15;
    <a id="L1168"></a>R_SPARC_PC10     R_SPARC = 16;
    <a id="L1169"></a>R_SPARC_PC22     R_SPARC = 17;
    <a id="L1170"></a>R_SPARC_WPLT30   R_SPARC = 18;
    <a id="L1171"></a>R_SPARC_COPY     R_SPARC = 19;
    <a id="L1172"></a>R_SPARC_GLOB_DAT R_SPARC = 20;
    <a id="L1173"></a>R_SPARC_JMP_SLOT R_SPARC = 21;
    <a id="L1174"></a>R_SPARC_RELATIVE R_SPARC = 22;
    <a id="L1175"></a>R_SPARC_UA32     R_SPARC = 23;
    <a id="L1176"></a>R_SPARC_PLT32    R_SPARC = 24;
    <a id="L1177"></a>R_SPARC_HIPLT22  R_SPARC = 25;
    <a id="L1178"></a>R_SPARC_LOPLT10  R_SPARC = 26;
    <a id="L1179"></a>R_SPARC_PCPLT32  R_SPARC = 27;
    <a id="L1180"></a>R_SPARC_PCPLT22  R_SPARC = 28;
    <a id="L1181"></a>R_SPARC_PCPLT10  R_SPARC = 29;
    <a id="L1182"></a>R_SPARC_10       R_SPARC = 30;
    <a id="L1183"></a>R_SPARC_11       R_SPARC = 31;
    <a id="L1184"></a>R_SPARC_64       R_SPARC = 32;
    <a id="L1185"></a>R_SPARC_OLO10    R_SPARC = 33;
    <a id="L1186"></a>R_SPARC_HH22     R_SPARC = 34;
    <a id="L1187"></a>R_SPARC_HM10     R_SPARC = 35;
    <a id="L1188"></a>R_SPARC_LM22     R_SPARC = 36;
    <a id="L1189"></a>R_SPARC_PC_HH22  R_SPARC = 37;
    <a id="L1190"></a>R_SPARC_PC_HM10  R_SPARC = 38;
    <a id="L1191"></a>R_SPARC_PC_LM22  R_SPARC = 39;
    <a id="L1192"></a>R_SPARC_WDISP16  R_SPARC = 40;
    <a id="L1193"></a>R_SPARC_WDISP19  R_SPARC = 41;
    <a id="L1194"></a>R_SPARC_GLOB_JMP R_SPARC = 42;
    <a id="L1195"></a>R_SPARC_7        R_SPARC = 43;
    <a id="L1196"></a>R_SPARC_5        R_SPARC = 44;
    <a id="L1197"></a>R_SPARC_6        R_SPARC = 45;
    <a id="L1198"></a>R_SPARC_DISP64   R_SPARC = 46;
    <a id="L1199"></a>R_SPARC_PLT64    R_SPARC = 47;
    <a id="L1200"></a>R_SPARC_HIX22    R_SPARC = 48;
    <a id="L1201"></a>R_SPARC_LOX10    R_SPARC = 49;
    <a id="L1202"></a>R_SPARC_H44      R_SPARC = 50;
    <a id="L1203"></a>R_SPARC_M44      R_SPARC = 51;
    <a id="L1204"></a>R_SPARC_L44      R_SPARC = 52;
    <a id="L1205"></a>R_SPARC_REGISTER R_SPARC = 53;
    <a id="L1206"></a>R_SPARC_UA64     R_SPARC = 54;
    <a id="L1207"></a>R_SPARC_UA16     R_SPARC = 55;
<a id="L1208"></a>)

<a id="L1210"></a>var rsparcStrings = []intName{
    <a id="L1211"></a>intName{0, &#34;R_SPARC_NONE&#34;},
    <a id="L1212"></a>intName{1, &#34;R_SPARC_8&#34;},
    <a id="L1213"></a>intName{2, &#34;R_SPARC_16&#34;},
    <a id="L1214"></a>intName{3, &#34;R_SPARC_32&#34;},
    <a id="L1215"></a>intName{4, &#34;R_SPARC_DISP8&#34;},
    <a id="L1216"></a>intName{5, &#34;R_SPARC_DISP16&#34;},
    <a id="L1217"></a>intName{6, &#34;R_SPARC_DISP32&#34;},
    <a id="L1218"></a>intName{7, &#34;R_SPARC_WDISP30&#34;},
    <a id="L1219"></a>intName{8, &#34;R_SPARC_WDISP22&#34;},
    <a id="L1220"></a>intName{9, &#34;R_SPARC_HI22&#34;},
    <a id="L1221"></a>intName{10, &#34;R_SPARC_22&#34;},
    <a id="L1222"></a>intName{11, &#34;R_SPARC_13&#34;},
    <a id="L1223"></a>intName{12, &#34;R_SPARC_LO10&#34;},
    <a id="L1224"></a>intName{13, &#34;R_SPARC_GOT10&#34;},
    <a id="L1225"></a>intName{14, &#34;R_SPARC_GOT13&#34;},
    <a id="L1226"></a>intName{15, &#34;R_SPARC_GOT22&#34;},
    <a id="L1227"></a>intName{16, &#34;R_SPARC_PC10&#34;},
    <a id="L1228"></a>intName{17, &#34;R_SPARC_PC22&#34;},
    <a id="L1229"></a>intName{18, &#34;R_SPARC_WPLT30&#34;},
    <a id="L1230"></a>intName{19, &#34;R_SPARC_COPY&#34;},
    <a id="L1231"></a>intName{20, &#34;R_SPARC_GLOB_DAT&#34;},
    <a id="L1232"></a>intName{21, &#34;R_SPARC_JMP_SLOT&#34;},
    <a id="L1233"></a>intName{22, &#34;R_SPARC_RELATIVE&#34;},
    <a id="L1234"></a>intName{23, &#34;R_SPARC_UA32&#34;},
    <a id="L1235"></a>intName{24, &#34;R_SPARC_PLT32&#34;},
    <a id="L1236"></a>intName{25, &#34;R_SPARC_HIPLT22&#34;},
    <a id="L1237"></a>intName{26, &#34;R_SPARC_LOPLT10&#34;},
    <a id="L1238"></a>intName{27, &#34;R_SPARC_PCPLT32&#34;},
    <a id="L1239"></a>intName{28, &#34;R_SPARC_PCPLT22&#34;},
    <a id="L1240"></a>intName{29, &#34;R_SPARC_PCPLT10&#34;},
    <a id="L1241"></a>intName{30, &#34;R_SPARC_10&#34;},
    <a id="L1242"></a>intName{31, &#34;R_SPARC_11&#34;},
    <a id="L1243"></a>intName{32, &#34;R_SPARC_64&#34;},
    <a id="L1244"></a>intName{33, &#34;R_SPARC_OLO10&#34;},
    <a id="L1245"></a>intName{34, &#34;R_SPARC_HH22&#34;},
    <a id="L1246"></a>intName{35, &#34;R_SPARC_HM10&#34;},
    <a id="L1247"></a>intName{36, &#34;R_SPARC_LM22&#34;},
    <a id="L1248"></a>intName{37, &#34;R_SPARC_PC_HH22&#34;},
    <a id="L1249"></a>intName{38, &#34;R_SPARC_PC_HM10&#34;},
    <a id="L1250"></a>intName{39, &#34;R_SPARC_PC_LM22&#34;},
    <a id="L1251"></a>intName{40, &#34;R_SPARC_WDISP16&#34;},
    <a id="L1252"></a>intName{41, &#34;R_SPARC_WDISP19&#34;},
    <a id="L1253"></a>intName{42, &#34;R_SPARC_GLOB_JMP&#34;},
    <a id="L1254"></a>intName{43, &#34;R_SPARC_7&#34;},
    <a id="L1255"></a>intName{44, &#34;R_SPARC_5&#34;},
    <a id="L1256"></a>intName{45, &#34;R_SPARC_6&#34;},
    <a id="L1257"></a>intName{46, &#34;R_SPARC_DISP64&#34;},
    <a id="L1258"></a>intName{47, &#34;R_SPARC_PLT64&#34;},
    <a id="L1259"></a>intName{48, &#34;R_SPARC_HIX22&#34;},
    <a id="L1260"></a>intName{49, &#34;R_SPARC_LOX10&#34;},
    <a id="L1261"></a>intName{50, &#34;R_SPARC_H44&#34;},
    <a id="L1262"></a>intName{51, &#34;R_SPARC_M44&#34;},
    <a id="L1263"></a>intName{52, &#34;R_SPARC_L44&#34;},
    <a id="L1264"></a>intName{53, &#34;R_SPARC_REGISTER&#34;},
    <a id="L1265"></a>intName{54, &#34;R_SPARC_UA64&#34;},
    <a id="L1266"></a>intName{55, &#34;R_SPARC_UA16&#34;},
<a id="L1267"></a>}

<a id="L1269"></a>func (i R_SPARC) String() string   { return stringName(uint32(i), rsparcStrings, false) }
<a id="L1270"></a>func (i R_SPARC) GoString() string { return stringName(uint32(i), rsparcStrings, true) }

<a id="L1272"></a><span class="comment">/*</span>
<a id="L1273"></a><span class="comment"> * Magic number for the elf trampoline, chosen wisely to be an immediate</span>
<a id="L1274"></a><span class="comment"> * value.</span>
<a id="L1275"></a><span class="comment"> */</span>
<a id="L1276"></a>const ARM_MAGIC_TRAMP_NUMBER = 0x5c000003


<a id="L1279"></a><span class="comment">/*</span>
<a id="L1280"></a><span class="comment"> * ELF32 File header.</span>
<a id="L1281"></a><span class="comment"> */</span>
<a id="L1282"></a>type Header32 struct {
    <a id="L1283"></a>Ident     [EI_NIDENT]byte; <span class="comment">/* File identification. */</span>
    <a id="L1284"></a>Type      uint16;          <span class="comment">/* File type. */</span>
    <a id="L1285"></a>Machine   uint16;          <span class="comment">/* Machine architecture. */</span>
    <a id="L1286"></a>Version   uint32;          <span class="comment">/* ELF format version. */</span>
    <a id="L1287"></a>Entry     uint32;          <span class="comment">/* Entry point. */</span>
    <a id="L1288"></a>Phoff     uint32;          <span class="comment">/* Program header file offset. */</span>
    <a id="L1289"></a>Shoff     uint32;          <span class="comment">/* Section header file offset. */</span>
    <a id="L1290"></a>Flags     uint32;          <span class="comment">/* Architecture-specific flags. */</span>
    <a id="L1291"></a>Ehsize    uint16;          <span class="comment">/* Size of ELF header in bytes. */</span>
    <a id="L1292"></a>Phentsize uint16;          <span class="comment">/* Size of program header entry. */</span>
    <a id="L1293"></a>Phnum     uint16;          <span class="comment">/* Number of program header entries. */</span>
    <a id="L1294"></a>Shentsize uint16;          <span class="comment">/* Size of section header entry. */</span>
    <a id="L1295"></a>Shnum     uint16;          <span class="comment">/* Number of section header entries. */</span>
    <a id="L1296"></a>Shstrndx  uint16;          <span class="comment">/* Section name strings section. */</span>
<a id="L1297"></a>}

<a id="L1299"></a><span class="comment">/*</span>
<a id="L1300"></a><span class="comment"> * ELF32 Section header.</span>
<a id="L1301"></a><span class="comment"> */</span>
<a id="L1302"></a>type Section32 struct {
    <a id="L1303"></a>Name      uint32; <span class="comment">/* Section name (index into the section header string table). */</span>
    <a id="L1304"></a>Type      uint32; <span class="comment">/* Section type. */</span>
    <a id="L1305"></a>Flags     uint32; <span class="comment">/* Section flags. */</span>
    <a id="L1306"></a>Addr      uint32; <span class="comment">/* Address in memory image. */</span>
    <a id="L1307"></a>Off       uint32; <span class="comment">/* Offset in file. */</span>
    <a id="L1308"></a>Size      uint32; <span class="comment">/* Size in bytes. */</span>
    <a id="L1309"></a>Link      uint32; <span class="comment">/* Index of a related section. */</span>
    <a id="L1310"></a>Info      uint32; <span class="comment">/* Depends on section type. */</span>
    <a id="L1311"></a>Addralign uint32; <span class="comment">/* Alignment in bytes. */</span>
    <a id="L1312"></a>Entsize   uint32; <span class="comment">/* Size of each entry in section. */</span>
<a id="L1313"></a>}

<a id="L1315"></a><span class="comment">/*</span>
<a id="L1316"></a><span class="comment"> * ELF32 Program header.</span>
<a id="L1317"></a><span class="comment"> */</span>
<a id="L1318"></a>type Prog32 struct {
    <a id="L1319"></a>Type   uint32; <span class="comment">/* Entry type. */</span>
    <a id="L1320"></a>Off    uint32; <span class="comment">/* File offset of contents. */</span>
    <a id="L1321"></a>Vaddr  uint32; <span class="comment">/* Virtual address in memory image. */</span>
    <a id="L1322"></a>Paddr  uint32; <span class="comment">/* Physical address (not used). */</span>
    <a id="L1323"></a>Filesz uint32; <span class="comment">/* Size of contents in file. */</span>
    <a id="L1324"></a>Memsz  uint32; <span class="comment">/* Size of contents in memory. */</span>
    <a id="L1325"></a>Flags  uint32; <span class="comment">/* Access permission flags. */</span>
    <a id="L1326"></a>Align  uint32; <span class="comment">/* Alignment in memory and file. */</span>
<a id="L1327"></a>}

<a id="L1329"></a><span class="comment">/*</span>
<a id="L1330"></a><span class="comment"> * ELF32 Dynamic structure.  The &#34;.dynamic&#34; section contains an array of them.</span>
<a id="L1331"></a><span class="comment"> */</span>
<a id="L1332"></a>type Dyn32 struct {
    <a id="L1333"></a>Tag int32;  <span class="comment">/* Entry type. */</span>
    <a id="L1334"></a>Val uint32; <span class="comment">/* Integer/Address value. */</span>
<a id="L1335"></a>}

<a id="L1337"></a><span class="comment">/*</span>
<a id="L1338"></a><span class="comment"> * Relocation entries.</span>
<a id="L1339"></a><span class="comment"> */</span>

<a id="L1341"></a><span class="comment">// ELF32 Relocations that don&#39;t need an addend field.</span>
<a id="L1342"></a>type Rel32 struct {
    <a id="L1343"></a>Off  uint32; <span class="comment">/* Location to be relocated. */</span>
    <a id="L1344"></a>Info uint32; <span class="comment">/* Relocation type and symbol index. */</span>
<a id="L1345"></a>}

<a id="L1347"></a><span class="comment">// ELF32 Relocations that need an addend field.</span>
<a id="L1348"></a>type Rela32 struct {
    <a id="L1349"></a>Off    uint32; <span class="comment">/* Location to be relocated. */</span>
    <a id="L1350"></a>Info   uint32; <span class="comment">/* Relocation type and symbol index. */</span>
    <a id="L1351"></a>Addend int32;  <span class="comment">/* Addend. */</span>
<a id="L1352"></a>}

<a id="L1354"></a>func R_SYM32(info uint32) uint32      { return uint32(info &gt;&gt; 8) }
<a id="L1355"></a>func R_TYPE32(info uint32) uint32     { return uint32(info &amp; 0xff) }
<a id="L1356"></a>func R_INFO32(sym, typ uint32) uint32 { return sym&lt;&lt;8 | typ }

<a id="L1358"></a><span class="comment">// ELF32 Symbol.</span>
<a id="L1359"></a>type Sym32 struct {
    <a id="L1360"></a>Name  uint32;
    <a id="L1361"></a>Value uint32;
    <a id="L1362"></a>Size  uint32;
    <a id="L1363"></a>Info  uint8;
    <a id="L1364"></a>Other uint8;
    <a id="L1365"></a>Shndx uint16;
<a id="L1366"></a>}

<a id="L1368"></a>const Sym32Size = 16

<a id="L1370"></a>func ST_BIND(info uint8) SymBind              { return SymBind(info &gt;&gt; 4) }
<a id="L1371"></a>func ST_TYPE(bind SymBind, typ SymType) uint8 { return uint8(bind)&lt;&lt;4 | uint8(typ)&amp;0xf }
<a id="L1372"></a>func ST_VISIBILITY(other uint8) SymVis        { return SymVis(other &amp; 3) }

<a id="L1374"></a><span class="comment">/*</span>
<a id="L1375"></a><span class="comment"> * ELF64</span>
<a id="L1376"></a><span class="comment"> */</span>

<a id="L1378"></a><span class="comment">/*</span>
<a id="L1379"></a><span class="comment"> * ELF64 file header.</span>
<a id="L1380"></a><span class="comment"> */</span>

<a id="L1382"></a>type Header64 struct {
    <a id="L1383"></a>Ident     [EI_NIDENT]byte; <span class="comment">/* File identification. */</span>
    <a id="L1384"></a>Type      uint16;          <span class="comment">/* File type. */</span>
    <a id="L1385"></a>Machine   uint16;          <span class="comment">/* Machine architecture. */</span>
    <a id="L1386"></a>Version   uint32;          <span class="comment">/* ELF format version. */</span>
    <a id="L1387"></a>Entry     uint64;          <span class="comment">/* Entry point. */</span>
    <a id="L1388"></a>Phoff     uint64;          <span class="comment">/* Program header file offset. */</span>
    <a id="L1389"></a>Shoff     uint64;          <span class="comment">/* Section header file offset. */</span>
    <a id="L1390"></a>Flags     uint32;          <span class="comment">/* Architecture-specific flags. */</span>
    <a id="L1391"></a>Ehsize    uint16;          <span class="comment">/* Size of ELF header in bytes. */</span>
    <a id="L1392"></a>Phentsize uint16;          <span class="comment">/* Size of program header entry. */</span>
    <a id="L1393"></a>Phnum     uint16;          <span class="comment">/* Number of program header entries. */</span>
    <a id="L1394"></a>Shentsize uint16;          <span class="comment">/* Size of section header entry. */</span>
    <a id="L1395"></a>Shnum     uint16;          <span class="comment">/* Number of section header entries. */</span>
    <a id="L1396"></a>Shstrndx  uint16;          <span class="comment">/* Section name strings section. */</span>
<a id="L1397"></a>}

<a id="L1399"></a><span class="comment">/*</span>
<a id="L1400"></a><span class="comment"> * ELF64 Section header.</span>
<a id="L1401"></a><span class="comment"> */</span>

<a id="L1403"></a>type Section64 struct {
    <a id="L1404"></a>Name      uint32; <span class="comment">/* Section name (index into the section header string table). */</span>
    <a id="L1405"></a>Type      uint32; <span class="comment">/* Section type. */</span>
    <a id="L1406"></a>Flags     uint64; <span class="comment">/* Section flags. */</span>
    <a id="L1407"></a>Addr      uint64; <span class="comment">/* Address in memory image. */</span>
    <a id="L1408"></a>Off       uint64; <span class="comment">/* Offset in file. */</span>
    <a id="L1409"></a>Size      uint64; <span class="comment">/* Size in bytes. */</span>
    <a id="L1410"></a>Link      uint32; <span class="comment">/* Index of a related section. */</span>
    <a id="L1411"></a>Info      uint32; <span class="comment">/* Depends on section type. */</span>
    <a id="L1412"></a>Addralign uint64; <span class="comment">/* Alignment in bytes. */</span>
    <a id="L1413"></a>Entsize   uint64; <span class="comment">/* Size of each entry in section. */</span>
<a id="L1414"></a>}

<a id="L1416"></a><span class="comment">/*</span>
<a id="L1417"></a><span class="comment"> * ELF64 Program header.</span>
<a id="L1418"></a><span class="comment"> */</span>

<a id="L1420"></a>type Prog64 struct {
    <a id="L1421"></a>Type   uint32; <span class="comment">/* Entry type. */</span>
    <a id="L1422"></a>Flags  uint32; <span class="comment">/* Access permission flags. */</span>
    <a id="L1423"></a>Off    uint64; <span class="comment">/* File offset of contents. */</span>
    <a id="L1424"></a>Vaddr  uint64; <span class="comment">/* Virtual address in memory image. */</span>
    <a id="L1425"></a>Paddr  uint64; <span class="comment">/* Physical address (not used). */</span>
    <a id="L1426"></a>Filesz uint64; <span class="comment">/* Size of contents in file. */</span>
    <a id="L1427"></a>Memsz  uint64; <span class="comment">/* Size of contents in memory. */</span>
    <a id="L1428"></a>Align  uint64; <span class="comment">/* Alignment in memory and file. */</span>
<a id="L1429"></a>}

<a id="L1431"></a><span class="comment">/*</span>
<a id="L1432"></a><span class="comment"> * ELF64 Dynamic structure.  The &#34;.dynamic&#34; section contains an array of them.</span>
<a id="L1433"></a><span class="comment"> */</span>

<a id="L1435"></a>type Dyn64 struct {
    <a id="L1436"></a>Tag int64;  <span class="comment">/* Entry type. */</span>
    <a id="L1437"></a>Val uint64; <span class="comment">/* Integer/address value */</span>
<a id="L1438"></a>}

<a id="L1440"></a><span class="comment">/*</span>
<a id="L1441"></a><span class="comment"> * Relocation entries.</span>
<a id="L1442"></a><span class="comment"> */</span>

<a id="L1444"></a><span class="comment">/* ELF64 relocations that don&#39;t need an addend field. */</span>
<a id="L1445"></a>type Rel64 struct {
    <a id="L1446"></a>Off  uint64; <span class="comment">/* Location to be relocated. */</span>
    <a id="L1447"></a>Info uint64; <span class="comment">/* Relocation type and symbol index. */</span>
<a id="L1448"></a>}

<a id="L1450"></a><span class="comment">/* ELF64 relocations that need an addend field. */</span>
<a id="L1451"></a>type Rela64 struct {
    <a id="L1452"></a>Off    uint64; <span class="comment">/* Location to be relocated. */</span>
    <a id="L1453"></a>Info   uint64; <span class="comment">/* Relocation type and symbol index. */</span>
    <a id="L1454"></a>Addend int64;  <span class="comment">/* Addend. */</span>
<a id="L1455"></a>}

<a id="L1457"></a>func R_SYM64(info uint64) uint32    { return uint32(info &gt;&gt; 32) }
<a id="L1458"></a>func R_TYPE64(info uint64) uint32   { return uint32(info) }
<a id="L1459"></a>func R_INFO(sym, typ uint32) uint64 { return uint64(sym)&lt;&lt;32 | uint64(typ) }


<a id="L1462"></a><span class="comment">/*</span>
<a id="L1463"></a><span class="comment"> * ELF64 symbol table entries.</span>
<a id="L1464"></a><span class="comment"> */</span>
<a id="L1465"></a>type Sym64 struct {
    <a id="L1466"></a>Name  uint32; <span class="comment">/* String table index of name. */</span>
    <a id="L1467"></a>Info  uint8;  <span class="comment">/* Type and binding information. */</span>
    <a id="L1468"></a>Other uint8;  <span class="comment">/* Reserved (not used). */</span>
    <a id="L1469"></a>Shndx uint16; <span class="comment">/* Section index of symbol. */</span>
    <a id="L1470"></a>Value uint64; <span class="comment">/* Symbol value. */</span>
    <a id="L1471"></a>Size  uint64; <span class="comment">/* Size of associated object. */</span>
<a id="L1472"></a>}

<a id="L1474"></a>const Sym64Size = 24

<a id="L1476"></a>type intName struct {
    <a id="L1477"></a>i   uint32;
    <a id="L1478"></a>s   string;
<a id="L1479"></a>}

<a id="L1481"></a>func stringName(i uint32, names []intName, goSyntax bool) string {
    <a id="L1482"></a>for _, n := range names {
        <a id="L1483"></a>if n.i == i {
            <a id="L1484"></a>if goSyntax {
                <a id="L1485"></a>return &#34;elf.&#34; + n.s
            <a id="L1486"></a>}
            <a id="L1487"></a>return n.s;
        <a id="L1488"></a>}
    <a id="L1489"></a>}

    <a id="L1491"></a><span class="comment">// second pass - look for smaller to add with.</span>
    <a id="L1492"></a><span class="comment">// assume sorted already</span>
    <a id="L1493"></a>for j := len(names) - 1; j &gt;= 0; j-- {
        <a id="L1494"></a>n := names[j];
        <a id="L1495"></a>if n.i &lt; i {
            <a id="L1496"></a>s := n.s;
            <a id="L1497"></a>if goSyntax {
                <a id="L1498"></a>s = &#34;elf.&#34; + s
            <a id="L1499"></a>}
            <a id="L1500"></a>return s + &#34;+&#34; + strconv.Uitoa64(uint64(i-n.i));
        <a id="L1501"></a>}
    <a id="L1502"></a>}

    <a id="L1504"></a>return strconv.Uitoa64(uint64(i));
<a id="L1505"></a>}

<a id="L1507"></a>func flagName(i uint32, names []intName, goSyntax bool) string {
    <a id="L1508"></a>s := &#34;&#34;;
    <a id="L1509"></a>for _, n := range names {
        <a id="L1510"></a>if n.i&amp;i == n.i {
            <a id="L1511"></a>if len(s) &gt; 0 {
                <a id="L1512"></a>s += &#34;+&#34;
            <a id="L1513"></a>}
            <a id="L1514"></a>if goSyntax {
                <a id="L1515"></a>s += &#34;elf.&#34;
            <a id="L1516"></a>}
            <a id="L1517"></a>s += n.s;
            <a id="L1518"></a>i -= n.i;
        <a id="L1519"></a>}
    <a id="L1520"></a>}
    <a id="L1521"></a>if len(s) == 0 {
        <a id="L1522"></a>return &#34;0x&#34; + strconv.Uitob64(uint64(i), 16)
    <a id="L1523"></a>}
    <a id="L1524"></a>if i != 0 {
        <a id="L1525"></a>s += &#34;+0x&#34; + strconv.Uitob64(uint64(i), 16)
    <a id="L1526"></a>}
    <a id="L1527"></a>return s;
<a id="L1528"></a>}
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
