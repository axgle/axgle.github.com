<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/libmach/macho.h</title>

  <link rel="stylesheet" type="text/css" href="../../doc/style.css">
  <script type="text/javascript" src="../../doc/godocs.js"></script>

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
        <a href="../../index.html"><img src="../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../doc/install.html">Install Go</a></li>
    <li><a href="../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../pkg/index.html">Package documentation</a></li>
    <li><a href="../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/libmach/macho.h</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 *	Definitions needed for  accessing MACH object headers.
 */

typedef struct {
	uint32	magic;		/* mach magic number identifier */
	uint32	cputype;	/* cpu specifier */
	uint32	cpusubtype;	/* machine specifier */
	uint32	filetype;	/* type of file */
	uint32	ncmds;		/* number of load commands */
	uint32	sizeofcmds;	/* the size of all the load commands */
	uint32	flags;		/* flags */
	uint32	reserved;	/* reserved */
} Machhdr;

typedef struct {
	uint32	type;	/* type of load command */
	uint32	size;	/* total size in bytes */
} MachCmd;

typedef struct  {
	MachCmd	cmd;
	char		segname[16];	/* segment name */
	uint32	vmaddr;		/* memory address of this segment */
	uint32	vmsize;		/* memory size of this segment */
	uint32	fileoff;	/* file offset of this segment */
	uint32	filesize;	/* amount to map from the file */
	uint32	maxprot;	/* maximum VM protection */
	uint32	initprot;	/* initial VM protection */
	uint32	nsects;		/* number of sections in segment */
	uint32	flags;		/* flags */
} MachSeg32; /* for 32-bit architectures */

typedef struct  {
	MachCmd	cmd;
	char		segname[16];	/* segment name */
	uvlong	vmaddr;		/* memory address of this segment */
	uvlong	vmsize;		/* memory size of this segment */
	uvlong	fileoff;	/* file offset of this segment */
	uvlong	filesize;	/* amount to map from the file */
	uint32	maxprot;	/* maximum VM protection */
	uint32	initprot;	/* initial VM protection */
	uint32	nsects;		/* number of sections in segment */
	uint32	flags;		/* flags */
} MachSeg64; /* for 64-bit architectures */

typedef struct  {
	MachCmd	cmd;
	uint32	fileoff;	/* file offset of this segment */
	uint32	filesize;	/* amount to map from the file */
} MachSymSeg;

typedef struct  {
	char		sectname[16];	/* name of this section */
	char		segname[16];	/* segment this section goes in */
	uint32	addr;		/* memory address of this section */
	uint32	size;		/* size in bytes of this section */
	uint32	offset;		/* file offset of this section */
	uint32	align;		/* section alignment (power of 2) */
	uint32	reloff;		/* file offset of relocation entries */
	uint32	nreloc;		/* number of relocation entries */
	uint32	flags;		/* flags (section type and attributes)*/
	uint32	reserved1;	/* reserved (for offset or index) */
	uint32	reserved2;	/* reserved (for count or sizeof) */
} MachSect32; /* for 32-bit architectures */

typedef struct  {
	char		sectname[16];	/* name of this section */
	char		segname[16];	/* segment this section goes in */
	uvlong	addr;		/* memory address of this section */
	uvlong	size;		/* size in bytes of this section */
	uint32	offset;		/* file offset of this section */
	uint32	align;		/* section alignment (power of 2) */
	uint32	reloff;		/* file offset of relocation entries */
	uint32	nreloc;		/* number of relocation entries */
	uint32	flags;		/* flags (section type and attributes)*/
	uint32	reserved1;	/* reserved (for offset or index) */
	uint32	reserved2;	/* reserved (for count or sizeof) */
	uint32	reserved3;	/* reserved */
} MachSect64; /* for 64-bit architectures */

enum {
	MACH_CPU_TYPE_X86_64 = (1&lt;&lt;24)|7,
	MACH_CPU_TYPE_X86 = 7,
	MACH_CPU_SUBTYPE_X86 = 3,
	MACH_EXECUTABLE_TYPE = 2,
	MACH_SEGMENT_32 = 1,	/* 32-bit mapped segment */
	MACH_SEGMENT_64 = 0x19,	/* 64-bit mapped segment */
	MACH_SYMSEG = 3,	/* obsolete gdb symtab, reused by go */
	MACH_UNIXTHREAD = 0x5,	/* thread (for stack) */
};


#define	MACH64_MAG		((0xcf&lt;&lt;24) | (0xfa&lt;&lt;16) | (0xed&lt;&lt;8) | 0xfe)
#define	MACH32_MAG		((0xce&lt;&lt;24) | (0xfa&lt;&lt;16) | (0xed&lt;&lt;8) | 0xfe)
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
