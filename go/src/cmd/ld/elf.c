<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/ld/elf.c</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/ld/elf.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include	&#34;l.h&#34;
#include	&#34;lib.h&#34;
#include	&#34;../ld/elf.h&#34;

/*
 * We use the 64-bit data structures on both 32- and 64-bit machines
 * in order to write the code just once.  The 64-bit data structure is
 * written in the 32-bit format on the 32-bit machines.
 */
#define	NSECT	32

static	int	elf64;
static	ElfEhdr	hdr;
static	ElfPhdr	*phdr[NSECT];
static	ElfShdr	*shdr[NSECT];

/*
 Initialize the global variable that describes the ELF header. It will be updated as
 we write section and prog headers.
 */
void
elfinit(void)
{
	switch(thechar) {
	// 64-bit architectures
	case &#39;6&#39;:
		elf64 = 1;
		hdr.phoff = ELF64HDRSIZE;	/* Must be be ELF64HDRSIZE: first PHdr must follow ELF header */
		hdr.shoff = ELF64HDRSIZE;	/* Will move as we add PHeaders */
		hdr.ehsize = ELF64HDRSIZE;	/* Must be ELF64HDRSIZE */
		hdr.phentsize = ELF64PHDRSIZE;	/* Must be ELF64PHDRSIZE */
		hdr.shentsize = ELF64SHDRSIZE;	/* Must be ELF64SHDRSIZE */
		break;

	// 32-bit architectures
	default:
		hdr.phoff = ELF32HDRSIZE;	/* Must be be ELF32HDRSIZE: first PHdr must follow ELF header */
		hdr.shoff = ELF32HDRSIZE;	/* Will move as we add PHeaders */
		hdr.ehsize = ELF32HDRSIZE;	/* Must be ELF32HDRSIZE */
		hdr.phentsize = ELF32PHDRSIZE;	/* Must be ELF32PHDRSIZE */
		hdr.shentsize = ELF32SHDRSIZE;	/* Must be ELF32SHDRSIZE */
	}
}

void
elf64phdr(ElfPhdr *e)
{
	LPUT(e-&gt;type);
	LPUT(e-&gt;flags);
	VPUT(e-&gt;off);
	VPUT(e-&gt;vaddr);
	VPUT(e-&gt;paddr);
	VPUT(e-&gt;filesz);
	VPUT(e-&gt;memsz);
	VPUT(e-&gt;align);
}

void
elf32phdr(ElfPhdr *e)
{
	LPUT(e-&gt;type);
	LPUT(e-&gt;off);
	LPUT(e-&gt;vaddr);
	LPUT(e-&gt;paddr);
	LPUT(e-&gt;filesz);
	LPUT(e-&gt;memsz);
	LPUT(e-&gt;flags);
	LPUT(e-&gt;align);
}

void
elf64shdr(ElfShdr *e)
{
	LPUT(e-&gt;name);
	LPUT(e-&gt;type);
	VPUT(e-&gt;flags);
	VPUT(e-&gt;addr);
	VPUT(e-&gt;off);
	VPUT(e-&gt;size);
	LPUT(e-&gt;link);
	LPUT(e-&gt;info);
	VPUT(e-&gt;addralign);
	VPUT(e-&gt;entsize);
}

void
elf32shdr(ElfShdr *e)
{
	LPUT(e-&gt;name);
	LPUT(e-&gt;type);
	LPUT(e-&gt;flags);
	LPUT(e-&gt;addr);
	LPUT(e-&gt;off);
	LPUT(e-&gt;size);
	LPUT(e-&gt;link);
	LPUT(e-&gt;info);
	LPUT(e-&gt;addralign);
	LPUT(e-&gt;entsize);
}

uint32
elfwriteshdrs(void)
{
	int i;

	if (elf64) {
		for (i = 0; i &lt; hdr.shnum; i++)
			elf64shdr(shdr[i]);
		return hdr.shnum * ELF64SHDRSIZE;
	}
	for (i = 0; i &lt; hdr.shnum; i++)
		elf32shdr(shdr[i]);
	return hdr.shnum * ELF32SHDRSIZE;
}

uint32
elfwritephdrs(void)
{
	int i;

	if (elf64) {
		for (i = 0; i &lt; hdr.phnum; i++)
			elf64phdr(phdr[i]);
		return hdr.phnum * ELF64PHDRSIZE;
	}
	for (i = 0; i &lt; hdr.phnum; i++)
		elf32phdr(phdr[i]);
	return hdr.phnum * ELF32PHDRSIZE;
}

ElfPhdr*
newElfPhdr(void)
{
	ElfPhdr *e;

	e = malloc(sizeof *e);
	memset(e, 0, sizeof *e);
	if (hdr.phnum &gt;= NSECT)
		diag(&#34;too many phdrs&#34;);
	else
		phdr[hdr.phnum++] = e;
	if (elf64)
		hdr.shoff += ELF64PHDRSIZE;
	else
		hdr.shoff += ELF32PHDRSIZE;
	return e;
}

ElfShdr*
newElfShstrtab(vlong name)
{
	hdr.shstrndx = hdr.shnum;
	return newElfShdr(name);
}

ElfShdr*
newElfShdr(vlong name)
{
	ElfShdr *e;

	e = malloc(sizeof *e);
	memset(e, 0, sizeof *e);
	e-&gt;name = name;
	if (hdr.shnum &gt;= NSECT) {
		diag(&#34;too many shdrs&#34;);
	} else {
		shdr[hdr.shnum++] = e;
	}
	return e;
}

ElfEhdr*
getElfEhdr(void)
{
	return &amp;hdr;
}

uint32
elf64writehdr(void)
{
	int i;

	for (i = 0; i &lt; EI_NIDENT; i++)
		cput(hdr.ident[i]);
	WPUT(hdr.type);
	WPUT(hdr.machine);
	LPUT(hdr.version);
	VPUT(hdr.entry);
	VPUT(hdr.phoff);
	VPUT(hdr.shoff);
	LPUT(hdr.flags);
	WPUT(hdr.ehsize);
	WPUT(hdr.phentsize);
	WPUT(hdr.phnum);
	WPUT(hdr.shentsize);
	WPUT(hdr.shnum);
	WPUT(hdr.shstrndx);
	return ELF64HDRSIZE;
}

uint32
elf32writehdr(void)
{
	int i;

	for (i = 0; i &lt; EI_NIDENT; i++)
		cput(hdr.ident[i]);
	WPUT(hdr.type);
	WPUT(hdr.machine);
	LPUT(hdr.version);
	LPUT(hdr.entry);
	LPUT(hdr.phoff);
	LPUT(hdr.shoff);
	LPUT(hdr.flags);
	WPUT(hdr.ehsize);
	WPUT(hdr.phentsize);
	WPUT(hdr.phnum);
	WPUT(hdr.shentsize);
	WPUT(hdr.shnum);
	WPUT(hdr.shstrndx);
	return ELF32HDRSIZE;
}

uint32
elfwritehdr(void)
{
	if(elf64)
		return elf64writehdr();
	return elf32writehdr();
}

/* Taken directly from the definition document for ELF64 */
uint32
elfhash(uchar *name)
{
	uint32 h = 0, g;
	while (*name) {
		h = (h &lt;&lt; 4) + *name++;
		if (g = h &amp; 0xf0000000)
			h ^= g &gt;&gt; 24;
		h &amp;= 0x0fffffff;
	}
	return h;
}

void
elfwritedynent(Sym *s, int tag, uint64 val)
{
	if(elf64) {
		adduint64(s, tag);
		adduint64(s, val);
	} else {
		adduint32(s, tag);
		adduint32(s, val);
	}
}

void
elfwritedynentsym(Sym *s, int tag, Sym *t)
{
	if(elf64)
		adduint64(s, tag);
	else
		adduint32(s, tag);
	addaddr(s, t);
}

void
elfwritedynentsymsize(Sym *s, int tag, Sym *t)
{
	if(elf64)
		adduint64(s, tag);
	else
		adduint32(s, tag);
	addsize(s, t);
}
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
