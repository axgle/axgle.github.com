<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/libmach/executable.c</title>

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
  <h1 id="generatedHeader">Text file src/libmach/executable.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno libmach/executable.c
// http://code.google.com/p/inferno-os/source/browse/utils/libmach/executable.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.
//	Power PC support Copyright © 1995-2004 C H Forsyth (forsyth@terzarima.net).
//	Portions Copyright © 1997-1999 Vita Nuova Limited.
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com).
//	Revisions Copyright © 2000-2004 Lucent Technologies Inc. and others.
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the &#34;Software&#34;), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

#include	&lt;u.h&gt;
#include	&lt;libc.h&gt;
#include	&lt;bio.h&gt;
#include	&lt;bootexec.h&gt;
#include	&lt;mach.h&gt;
#include	&#34;elf.h&#34;
#include	&#34;macho.h&#34;

/*
 *	All a.out header types.  The dummy entry allows canonical
 *	processing of the union as a sequence of int32s
 */

typedef struct {
	union{
		/*struct { */
			Exec exechdr;		/* a.out.h */
		/*	uvlong hdr[1];*/
		/*};*/
		Ehdr32 elfhdr32;			/* elf.h */
		Ehdr64 elfhdr64;			/* elf.h */
		struct mipsexec mips;	/* bootexec.h */
		struct mips4kexec mipsk4;	/* bootexec.h */
		struct sparcexec sparc;	/* bootexec.h */
		struct nextexec next;	/* bootexec.h */
		Machhdr machhdr;	/* macho.h */
	} e;
	int32 dummy;			/* padding to ensure extra int32 */
} ExecHdr;

static	int	nextboot(int, Fhdr*, ExecHdr*);
static	int	sparcboot(int, Fhdr*, ExecHdr*);
static	int	mipsboot(int, Fhdr*, ExecHdr*);
static	int	mips4kboot(int, Fhdr*, ExecHdr*);
static	int	common(int, Fhdr*, ExecHdr*);
static	int	commonllp64(int, Fhdr*, ExecHdr*);
static	int	adotout(int, Fhdr*, ExecHdr*);
static	int	elfdotout(int, Fhdr*, ExecHdr*);
static	int	machdotout(int, Fhdr*, ExecHdr*);
static	int	armdotout(int, Fhdr*, ExecHdr*);
static	void	setsym(Fhdr*, int32, int32, int32, vlong);
static	void	setdata(Fhdr*, uvlong, int32, vlong, int32);
static	void	settext(Fhdr*, uvlong, uvlong, int32, vlong);
static	void	hswal(void*, int, uint32(*)(uint32));
static	uvlong	_round(uvlong, uint32);

/*
 *	definition of per-executable file type structures
 */

typedef struct Exectable{
	int32	magic;			/* big-endian magic number of file */
	char	*name;			/* executable identifier */
	char	*dlmname;		/* dynamically loadable module identifier */
	uchar	type;			/* Internal code */
	uchar	_magic;			/* _MAGIC() magic */
	Mach	*mach;			/* Per-machine data */
	int32	hsize;			/* header size */
	uint32	(*swal)(uint32);		/* beswal or leswal */
	int	(*hparse)(int, Fhdr*, ExecHdr*);
} ExecTable;

extern	Mach	mmips;
extern	Mach	mmips2le;
extern	Mach	mmips2be;
extern	Mach	msparc;
extern	Mach	msparc64;
extern	Mach	m68020;
extern	Mach	mi386;
extern	Mach	mamd64;
extern	Mach	marm;
extern	Mach	mpower;
extern	Mach	mpower64;
extern	Mach	malpha;

/* BUG: FIX THESE WHEN NEEDED */
Mach	mmips;
Mach	mmips2le;
Mach	mmips2be;
Mach	msparc;
Mach	msparc64;
Mach	m68020;
Mach	mpower;
Mach	mpower64;
Mach	malpha;

ExecTable exectab[] =
{
	{ V_MAGIC,			/* Mips v.out */
		&#34;mips plan 9 executable BE&#34;,
		&#34;mips plan 9 dlm BE&#34;,
		FMIPS,
		1,
		&amp;mmips,
		sizeof(Exec),
		beswal,
		adotout },
	{ P_MAGIC,			/* Mips 0.out (r3k le) */
		&#34;mips plan 9 executable LE&#34;,
		&#34;mips plan 9 dlm LE&#34;,
		FMIPSLE,
		1,
		&amp;mmips,
		sizeof(Exec),
		beswal,
		adotout },
	{ M_MAGIC,			/* Mips 4.out */
		&#34;mips 4k plan 9 executable BE&#34;,
		&#34;mips 4k plan 9 dlm BE&#34;,
		FMIPS2BE,
		1,
		&amp;mmips2be,
		sizeof(Exec),
		beswal,
		adotout },
	{ N_MAGIC,			/* Mips 0.out */
		&#34;mips 4k plan 9 executable LE&#34;,
		&#34;mips 4k plan 9 dlm LE&#34;,
		FMIPS2LE,
		1,
		&amp;mmips2le,
		sizeof(Exec),
		beswal,
		adotout },
	{ 0x160&lt;&lt;16,			/* Mips boot image */
		&#34;mips plan 9 boot image&#34;,
		nil,
		FMIPSB,
		0,
		&amp;mmips,
		sizeof(struct mipsexec),
		beswal,
		mipsboot },
	{ (0x160&lt;&lt;16)|3,		/* Mips boot image */
		&#34;mips 4k plan 9 boot image&#34;,
		nil,
		FMIPSB,
		0,
		&amp;mmips2be,
		sizeof(struct mips4kexec),
		beswal,
		mips4kboot },
	{ K_MAGIC,			/* Sparc k.out */
		&#34;sparc plan 9 executable&#34;,
		&#34;sparc plan 9 dlm&#34;,
		FSPARC,
		1,
		&amp;msparc,
		sizeof(Exec),
		beswal,
		adotout },
	{ 0x01030107, 			/* Sparc boot image */
		&#34;sparc plan 9 boot image&#34;,
		nil,
		FSPARCB,
		0,
		&amp;msparc,
		sizeof(struct sparcexec),
		beswal,
		sparcboot },
	{ U_MAGIC,			/* Sparc64 u.out */
		&#34;sparc64 plan 9 executable&#34;,
		&#34;sparc64 plan 9 dlm&#34;,
		FSPARC64,
		1,
		&amp;msparc64,
		sizeof(Exec),
		beswal,
		adotout },
	{ A_MAGIC,			/* 68020 2.out &amp; boot image */
		&#34;68020 plan 9 executable&#34;,
		&#34;68020 plan 9 dlm&#34;,
		F68020,
		1,
		&amp;m68020,
		sizeof(Exec),
		beswal,
		common },
	{ 0xFEEDFACE,			/* Next boot image */
		&#34;next plan 9 boot image&#34;,
		nil,
		FNEXTB,
		0,
		&amp;m68020,
		sizeof(struct nextexec),
		beswal,
		nextboot },
	{ I_MAGIC,			/* I386 8.out &amp; boot image */
		&#34;386 plan 9 executable&#34;,
		&#34;386 plan 9 dlm&#34;,
		FI386,
		1,
		&amp;mi386,
		sizeof(Exec),
		beswal,
		common },
	{ S_MAGIC,			/* amd64 6.out &amp; boot image */
		&#34;amd64 plan 9 executable&#34;,
		&#34;amd64 plan 9 dlm&#34;,
		FAMD64,
		1,
		&amp;mamd64,
		sizeof(Exec)+8,
		nil,
		commonllp64 },
	{ Q_MAGIC,			/* PowerPC q.out &amp; boot image */
		&#34;power plan 9 executable&#34;,
		&#34;power plan 9 dlm&#34;,
		FPOWER,
		1,
		&amp;mpower,
		sizeof(Exec),
		beswal,
		common },
	{ T_MAGIC,			/* power64 9.out &amp; boot image */
		&#34;power64 plan 9 executable&#34;,
		&#34;power64 plan 9 dlm&#34;,
		FPOWER64,
		1,
		&amp;mpower64,
		sizeof(Exec)+8,
		nil,
		commonllp64 },
	{ ELF_MAG,			/* any elf32 or elf64 */
		&#34;elf executable&#34;,
		nil,
		FNONE,
		0,
		&amp;mi386,
		sizeof(Ehdr64),
		nil,
		elfdotout },
	{ MACH64_MAG,			/* 64-bit MACH (apple mac) */
		&#34;mach executable&#34;,
		nil,
		FAMD64,
		0,
		&amp;mamd64,
		sizeof(Machhdr),
		nil,
		machdotout },
	{ MACH32_MAG,			/* 32-bit MACH (apple mac) */
		&#34;mach executable&#34;,
		nil,
		FI386,
		0,
		&amp;mi386,
		sizeof(Machhdr),
		nil,
		machdotout },
	{ E_MAGIC,			/* Arm 5.out and boot image */
		&#34;arm plan 9 executable&#34;,
		&#34;arm plan 9 dlm&#34;,
		FARM,
		1,
		&amp;marm,
		sizeof(Exec),
		beswal,
		common },
	{ (143&lt;&lt;16)|0413,		/* (Free|Net)BSD Arm */
		&#34;arm *bsd executable&#34;,
		nil,
		FARM,
		0,
		&amp;marm,
		sizeof(Exec),
		leswal,
		armdotout },
	{ L_MAGIC,			/* alpha 7.out */
		&#34;alpha plan 9 executable&#34;,
		&#34;alpha plan 9 dlm&#34;,
		FALPHA,
		1,
		&amp;malpha,
		sizeof(Exec),
		beswal,
		common },
	{ 0x0700e0c3,			/* alpha boot image */
		&#34;alpha plan 9 boot image&#34;,
		nil,
		FALPHA,
		0,
		&amp;malpha,
		sizeof(Exec),
		beswal,
		common },
	{ 0 },
};

Mach	*mach = &amp;mi386;			/* Global current machine table */

static ExecTable*
couldbe4k(ExecTable *mp)
{
	Dir *d;
	ExecTable *f;

	if((d=dirstat(&#34;/proc/1/regs&#34;)) == nil)
		return mp;
	if(d-&gt;length &lt; 32*8){		/* R3000 */
		free(d);
		return mp;
	}
	free(d);
	for (f = exectab; f-&gt;magic; f++)
		if(f-&gt;magic == M_MAGIC) {
			f-&gt;name = &#34;mips plan 9 executable on mips2 kernel&#34;;
			return f;
		}
	return mp;
}

int
crackhdr(int fd, Fhdr *fp)
{
	ExecTable *mp;
	ExecHdr d;
	int nb, ret;
	uint32 magic;

	fp-&gt;type = FNONE;
	nb = read(fd, (char *)&amp;d.e, sizeof(d.e));
	if (nb &lt;= 0)
		return 0;

	ret = 0;
	magic = beswal(d.e.exechdr.magic);		/* big-endian */
	for (mp = exectab; mp-&gt;magic; mp++) {
		if (nb &lt; mp-&gt;hsize)
			continue;

		/*
		 * The magic number has morphed into something
		 * with fields (the straw was DYN_MAGIC) so now
		 * a flag is needed in Fhdr to distinguish _MAGIC()
		 * magic numbers from foreign magic numbers.
		 *
		 * This code is creaking a bit and if it has to
		 * be modified/extended much more it&#39;s probably
		 * time to step back and redo it all.
		 */
		if(mp-&gt;_magic){
			if(mp-&gt;magic != (magic &amp; ~DYN_MAGIC))
				continue;

			if(mp-&gt;magic == V_MAGIC)
				mp = couldbe4k(mp);

			if ((magic &amp; DYN_MAGIC) &amp;&amp; mp-&gt;dlmname != nil)
				fp-&gt;name = mp-&gt;dlmname;
			else
				fp-&gt;name = mp-&gt;name;
		}
		else{
			if(mp-&gt;magic != magic)
				continue;
			fp-&gt;name = mp-&gt;name;
		}
		fp-&gt;type = mp-&gt;type;
		fp-&gt;hdrsz = mp-&gt;hsize;		/* will be zero on bootables */
		fp-&gt;_magic = mp-&gt;_magic;
		fp-&gt;magic = magic;

		mach = mp-&gt;mach;
		if(mp-&gt;swal != nil)
			hswal(&amp;d, sizeof(d.e)/sizeof(uint32), mp-&gt;swal);
		ret = mp-&gt;hparse(fd, fp, &amp;d);
		seek(fd, mp-&gt;hsize, 0);		/* seek to end of header */
		break;
	}
	if(mp-&gt;magic == 0)
		werrstr(&#34;unknown header type&#34;);
	return ret;
}

/*
 * Convert header to canonical form
 */
static void
hswal(void *v, int n, uint32 (*swap)(uint32))
{
	uint32 *ulp;

	for(ulp = v; n--; ulp++)
		*ulp = (*swap)(*ulp);
}

/*
 *	Crack a normal a.out-type header
 */
static int
adotout(int fd, Fhdr *fp, ExecHdr *hp)
{
	int32 pgsize;

	USED(fd);
	pgsize = mach-&gt;pgsize;
	settext(fp, hp-&gt;e.exechdr.entry, pgsize+sizeof(Exec),
			hp-&gt;e.exechdr.text, sizeof(Exec));
	setdata(fp, _round(pgsize+fp-&gt;txtsz+sizeof(Exec), pgsize),
		hp-&gt;e.exechdr.data, fp-&gt;txtsz+sizeof(Exec), hp-&gt;e.exechdr.bss);
	setsym(fp, hp-&gt;e.exechdr.syms, hp-&gt;e.exechdr.spsz, hp-&gt;e.exechdr.pcsz, fp-&gt;datoff+fp-&gt;datsz);
	return 1;
}

static void
commonboot(Fhdr *fp)
{
	if (!(fp-&gt;entry &amp; mach-&gt;ktmask))
		return;

	switch(fp-&gt;type) {				/* boot image */
	case F68020:
		fp-&gt;type = F68020B;
		fp-&gt;name = &#34;68020 plan 9 boot image&#34;;
		break;
	case FI386:
		fp-&gt;type = FI386B;
		fp-&gt;txtaddr = (u32int)fp-&gt;entry;
		fp-&gt;name = &#34;386 plan 9 boot image&#34;;
		fp-&gt;dataddr = _round(fp-&gt;txtaddr+fp-&gt;txtsz, mach-&gt;pgsize);
		break;
	case FARM:
		fp-&gt;type = FARMB;
		fp-&gt;txtaddr = (u32int)fp-&gt;entry;
		fp-&gt;name = &#34;ARM plan 9 boot image&#34;;
		fp-&gt;dataddr = _round(fp-&gt;txtaddr+fp-&gt;txtsz, mach-&gt;pgsize);
		return;
	case FALPHA:
		fp-&gt;type = FALPHAB;
		fp-&gt;txtaddr = (u32int)fp-&gt;entry;
		fp-&gt;name = &#34;alpha plan 9 boot image&#34;;
		fp-&gt;dataddr = fp-&gt;txtaddr+fp-&gt;txtsz;
		break;
	case FPOWER:
		fp-&gt;type = FPOWERB;
		fp-&gt;txtaddr = (u32int)fp-&gt;entry;
		fp-&gt;name = &#34;power plan 9 boot image&#34;;
		fp-&gt;dataddr = fp-&gt;txtaddr+fp-&gt;txtsz;
		break;
	case FAMD64:
		fp-&gt;type = FAMD64B;
		fp-&gt;txtaddr = fp-&gt;entry;
		fp-&gt;name = &#34;amd64 plan 9 boot image&#34;;
		fp-&gt;dataddr = _round(fp-&gt;txtaddr+fp-&gt;txtsz, mach-&gt;pgsize);
		break;
	default:
		return;
	}
	fp-&gt;hdrsz = 0;			/* header stripped */
}

/*
 *	_MAGIC() style headers and
 *	alpha plan9-style bootable images for axp &#34;headerless&#34; boot
 *
 */
static int
common(int fd, Fhdr *fp, ExecHdr *hp)
{
	adotout(fd, fp, hp);
	if(hp-&gt;e.exechdr.magic &amp; DYN_MAGIC) {
		fp-&gt;txtaddr = 0;
		fp-&gt;dataddr = fp-&gt;txtsz;
		return 1;
	}
	commonboot(fp);
	return 1;
}

static int
commonllp64(int unused, Fhdr *fp, ExecHdr *hp)
{
	int32 pgsize;
	uvlong entry;

	hswal(&amp;hp-&gt;e, sizeof(Exec)/sizeof(int32), beswal);
	if(!(hp-&gt;e.exechdr.magic &amp; HDR_MAGIC))
		return 0;

	/*
	 * There can be more magic here if the
	 * header ever needs more expansion.
	 * For now just catch use of any of the
	 * unused bits.
	 */
	if((hp-&gt;e.exechdr.magic &amp; ~DYN_MAGIC)&gt;&gt;16)
		return 0;
	union {
		char *p;
		uvlong *v;
	} u;
	u.p = (char*)&amp;hp-&gt;e.exechdr;
	entry = beswav(*u.v);

	pgsize = mach-&gt;pgsize;
	settext(fp, entry, pgsize+fp-&gt;hdrsz, hp-&gt;e.exechdr.text, fp-&gt;hdrsz);
	setdata(fp, _round(pgsize+fp-&gt;txtsz+fp-&gt;hdrsz, pgsize),
		hp-&gt;e.exechdr.data, fp-&gt;txtsz+fp-&gt;hdrsz, hp-&gt;e.exechdr.bss);
	setsym(fp, hp-&gt;e.exechdr.syms, hp-&gt;e.exechdr.spsz, hp-&gt;e.exechdr.pcsz, fp-&gt;datoff+fp-&gt;datsz);

	if(hp-&gt;e.exechdr.magic &amp; DYN_MAGIC) {
		fp-&gt;txtaddr = 0;
		fp-&gt;dataddr = fp-&gt;txtsz;
		return 1;
	}
	commonboot(fp);
	return 1;
}

/*
 *	mips bootable image.
 */
static int
mipsboot(int fd, Fhdr *fp, ExecHdr *hp)
{
abort();
#ifdef unused
	USED(fd);
	fp-&gt;type = FMIPSB;
	switch(hp-&gt;e.exechdr.amagic) {
	default:
	case 0407:	/* some kind of mips */
		settext(fp, (u32int)hp-&gt;e.mentry, (u32int)hp-&gt;e.text_start,
			hp-&gt;e.tsize, sizeof(struct mipsexec)+4);
		setdata(fp, (u32int)hp-&gt;e.data_start, hp-&gt;e.dsize,
			fp-&gt;txtoff+hp-&gt;e.tsize, hp-&gt;e.bsize);
		break;
	case 0413:	/* some kind of mips */
		settext(fp, (u32int)hp-&gt;e.mentry, (u32int)hp-&gt;e.text_start,
			hp-&gt;e.tsize, 0);
		setdata(fp, (u32int)hp-&gt;e.data_start, hp-&gt;e.dsize,
			hp-&gt;e.tsize, hp-&gt;e.bsize);
		break;
	}
	setsym(fp, hp-&gt;e.nsyms, 0, hp-&gt;e.pcsize, hp-&gt;e.symptr);
	fp-&gt;hdrsz = 0;			/* header stripped */
#endif
	return 1;
}

/*
 *	mips4k bootable image.
 */
static int
mips4kboot(int fd, Fhdr *fp, ExecHdr *hp)
{
abort();
#ifdef unused
	USED(fd);
	fp-&gt;type = FMIPSB;
	switch(hp-&gt;e.h.amagic) {
	default:
	case 0407:	/* some kind of mips */
		settext(fp, (u32int)hp-&gt;e.h.mentry, (u32int)hp-&gt;e.h.text_start,
			hp-&gt;e.h.tsize, sizeof(struct mips4kexec));
		setdata(fp, (u32int)hp-&gt;e.h.data_start, hp-&gt;e.h.dsize,
			fp-&gt;txtoff+hp-&gt;e.h.tsize, hp-&gt;e.h.bsize);
		break;
	case 0413:	/* some kind of mips */
		settext(fp, (u32int)hp-&gt;e.h.mentry, (u32int)hp-&gt;e.h.text_start,
			hp-&gt;e.h.tsize, 0);
		setdata(fp, (u32int)hp-&gt;e.h.data_start, hp-&gt;e.h.dsize,
			hp-&gt;e.h.tsize, hp-&gt;e.h.bsize);
		break;
	}
	setsym(fp, hp-&gt;e.h.nsyms, 0, hp-&gt;e.h.pcsize, hp-&gt;e.h.symptr);
	fp-&gt;hdrsz = 0;			/* header stripped */
#endif
	return 1;
}

/*
 *	sparc bootable image
 */
static int
sparcboot(int fd, Fhdr *fp, ExecHdr *hp)
{
abort();
#ifdef unused
	USED(fd);
	fp-&gt;type = FSPARCB;
	settext(fp, hp-&gt;e.sentry, hp-&gt;e.sentry, hp-&gt;e.stext,
		sizeof(struct sparcexec));
	setdata(fp, hp-&gt;e.sentry+hp-&gt;e.stext, hp-&gt;e.sdata,
		fp-&gt;txtoff+hp-&gt;e.stext, hp-&gt;e.sbss);
	setsym(fp, hp-&gt;e.ssyms, 0, hp-&gt;e.sdrsize, fp-&gt;datoff+hp-&gt;e.sdata);
	fp-&gt;hdrsz = 0;			/* header stripped */
#endif
	return 1;
}

/*
 *	next bootable image
 */
static int
nextboot(int fd, Fhdr *fp, ExecHdr *hp)
{
abort();
#ifdef unused
	USED(fd);
	fp-&gt;type = FNEXTB;
	settext(fp, hp-&gt;e.textc.vmaddr, hp-&gt;e.textc.vmaddr,
		hp-&gt;e.texts.size, hp-&gt;e.texts.offset);
	setdata(fp, hp-&gt;e.datac.vmaddr, hp-&gt;e.datas.size,
		hp-&gt;e.datas.offset, hp-&gt;e.bsss.size);
	setsym(fp, hp-&gt;e.symc.nsyms, hp-&gt;e.symc.spoff, hp-&gt;e.symc.pcoff,
		hp-&gt;e.symc.symoff);
	fp-&gt;hdrsz = 0;			/* header stripped */
#endif
	return 1;
}

/*
 * Elf32 and Elf64 binaries.
 */
static int
elf64dotout(int fd, Fhdr *fp, ExecHdr *hp)
{

	uvlong (*swav)(uvlong);
	uint32 (*swal)(uint32);
	ushort (*swab)(ushort);
	Ehdr64 *ep;
	Phdr64 *ph;
	Shdr64 *sh;
	int i, it, id, is, phsz, shsz;

	/* bitswap the header according to the DATA format */
	ep = &amp;hp-&gt;e.elfhdr64;
	if(ep-&gt;ident[CLASS] != ELFCLASS64) {
		werrstr(&#34;bad ELF class - not 32 bit or 64 bit&#34;);
		return 0;
	}
	if(ep-&gt;ident[DATA] == ELFDATA2LSB) {
		swab = leswab;
		swal = leswal;
		swav = leswav;
	} else if(ep-&gt;ident[DATA] == ELFDATA2MSB) {
		swab = beswab;
		swal = beswal;
		swav = beswav;
	} else {
		werrstr(&#34;bad ELF encoding - not big or little endian&#34;);
		return 0;
	}

	ep-&gt;type = swab(ep-&gt;type);
	ep-&gt;machine = swab(ep-&gt;machine);
	ep-&gt;version = swal(ep-&gt;version);
	ep-&gt;elfentry = swal(ep-&gt;elfentry);
	ep-&gt;phoff = swav(ep-&gt;phoff);
	ep-&gt;shoff = swav(ep-&gt;shoff);
	ep-&gt;flags = swav(ep-&gt;flags);
	ep-&gt;ehsize = swab(ep-&gt;ehsize);
	ep-&gt;phentsize = swab(ep-&gt;phentsize);
	ep-&gt;phnum = swab(ep-&gt;phnum);
	ep-&gt;shentsize = swab(ep-&gt;shentsize);
	ep-&gt;shnum = swab(ep-&gt;shnum);
	ep-&gt;shstrndx = swab(ep-&gt;shstrndx);
	if(ep-&gt;type != EXEC || ep-&gt;version != CURRENT)
		return 0;

	/* we could definitely support a lot more machines here */
	fp-&gt;magic = ELF_MAG;
	fp-&gt;hdrsz = (ep-&gt;ehsize+ep-&gt;phnum*ep-&gt;phentsize+16)&amp;~15;
	switch(ep-&gt;machine) {
	case AMD64:
		mach = &amp;mamd64;
		fp-&gt;type = FAMD64;
		break;
	default:
		return 0;
	}

	if(ep-&gt;phentsize != sizeof(Phdr64)) {
		werrstr(&#34;bad ELF header size&#34;);
		return 0;
	}
	phsz = sizeof(Phdr64)*ep-&gt;phnum;
	ph = malloc(phsz);
	if(!ph)
		return 0;
	seek(fd, ep-&gt;phoff, 0);
	if(read(fd, ph, phsz) &lt; 0) {
		free(ph);
		return 0;
	}
	hswal(ph, phsz/sizeof(uint32), swal);

	shsz = sizeof(Shdr64)*ep-&gt;shnum;
	sh = malloc(shsz);
	if(sh) {
		seek(fd, ep-&gt;shoff, 0);
		if(read(fd, sh, shsz) &lt; 0) {
			free(sh);
			sh = 0;
		} else
			hswal(sh, shsz/sizeof(uint32), swal);
	}

	/* find text, data and symbols and install them */
	it = id = is = -1;
	for(i = 0; i &lt; ep-&gt;phnum; i++) {
		if(ph[i].type == LOAD
		&amp;&amp; (ph[i].flags &amp; (R|X)) == (R|X) &amp;&amp; it == -1)
			it = i;
		else if(ph[i].type == LOAD
		&amp;&amp; (ph[i].flags &amp; (R|W)) == (R|W) &amp;&amp; id == -1)
			id = i;
		else if(ph[i].type == NOPTYPE &amp;&amp; is == -1)
			is = i;
	}
	if(it == -1 || id == -1) {
		/*
		 * The SPARC64 boot image is something of an ELF hack.
		 * Text+Data+BSS are represented by ph[0].  Symbols
		 * are represented by ph[1]:
		 *
		 *		filesz, memsz, vaddr, paddr, off
		 * ph[0] : txtsz+datsz, txtsz+datsz+bsssz, txtaddr-KZERO, datasize, txtoff
		 * ph[1] : symsz, lcsz, 0, 0, symoff
		 */
		if(ep-&gt;machine == SPARC64 &amp;&amp; ep-&gt;phnum == 2) {
			uint32 txtaddr, txtsz, dataddr, bsssz;

			txtaddr = ph[0].vaddr | 0x80000000;
			txtsz = ph[0].filesz - ph[0].paddr;
			dataddr = txtaddr + txtsz;
			bsssz = ph[0].memsz - ph[0].filesz;
			settext(fp, ep-&gt;elfentry | 0x80000000, txtaddr, txtsz, ph[0].offset);
			setdata(fp, dataddr, ph[0].paddr, ph[0].offset + txtsz, bsssz);
			setsym(fp, ph[1].filesz, 0, ph[1].memsz, ph[1].offset);
			free(ph);
			return 1;
		}

		werrstr(&#34;No TEXT or DATA sections&#34;);
error:
		free(ph);
		free(sh);
		return 0;
	}

	settext(fp, ep-&gt;elfentry, ph[it].vaddr, ph[it].memsz, ph[it].offset);
	setdata(fp, ph[id].vaddr, ph[id].filesz, ph[id].offset, ph[id].memsz - ph[id].filesz);
	if(is != -1)
		setsym(fp, ph[is].filesz, 0, ph[is].memsz, ph[is].offset);
	else if(sh != 0){
		char *buf;
		uvlong symsize = 0;
		uvlong symoff = 0;
		uvlong pclnsz = 0;

		/* load shstrtab names */
		buf = malloc(sh[ep-&gt;shstrndx].size);
		if (buf == 0)
			goto done;
		memset(buf, 0, sizeof buf);
		seek(fd, sh[ep-&gt;shstrndx].offset, 0);
		read(fd, buf, sh[ep-&gt;shstrndx].size);

		for(i = 0; i &lt; ep-&gt;shnum; i++) {
			if (strcmp(&amp;buf[sh[i].name], &#34;.gosymtab&#34;) == 0) {
				symsize = sh[i].size;
				symoff = sh[i].offset;
			}
			if (strcmp(&amp;buf[sh[i].name], &#34;.gopclntab&#34;) == 0) {
				if (sh[i].offset != symoff+symsize) {
					werrstr(&#34;pc line table not contiguous with symbol table&#34;);
					free(buf);
					goto error;
				}
				pclnsz = sh[i].size;
			}
		}
		setsym(fp, symsize, 0, pclnsz, symoff);
		free(buf);
	}
done:
	free(ph);
	free(sh);
	return 1;
}

static int
elfdotout(int fd, Fhdr *fp, ExecHdr *hp)
{

	uint32 (*swal)(uint32);
	ushort (*swab)(ushort);
	Ehdr32 *ep;
	Phdr32 *ph;
	int i, it, id, is, phsz, shsz;
	Shdr32 *sh;

	/* bitswap the header according to the DATA format */
	ep = &amp;hp-&gt;e.elfhdr32;
	if(ep-&gt;ident[CLASS] != ELFCLASS32) {
		return elf64dotout(fd, fp, hp);
	}
	if(ep-&gt;ident[DATA] == ELFDATA2LSB) {
		swab = leswab;
		swal = leswal;
	} else if(ep-&gt;ident[DATA] == ELFDATA2MSB) {
		swab = beswab;
		swal = beswal;
	} else {
		werrstr(&#34;bad ELF encoding - not big or little endian&#34;);
		return 0;
	}

	ep-&gt;type = swab(ep-&gt;type);
	ep-&gt;machine = swab(ep-&gt;machine);
	ep-&gt;version = swal(ep-&gt;version);
	ep-&gt;elfentry = swal(ep-&gt;elfentry);
	ep-&gt;phoff = swal(ep-&gt;phoff);
	ep-&gt;shoff = swal(ep-&gt;shoff);
	ep-&gt;flags = swal(ep-&gt;flags);
	ep-&gt;ehsize = swab(ep-&gt;ehsize);
	ep-&gt;phentsize = swab(ep-&gt;phentsize);
	ep-&gt;phnum = swab(ep-&gt;phnum);
	ep-&gt;shentsize = swab(ep-&gt;shentsize);
	ep-&gt;shnum = swab(ep-&gt;shnum);
	ep-&gt;shstrndx = swab(ep-&gt;shstrndx);
	if(ep-&gt;type != EXEC || ep-&gt;version != CURRENT)
		return 0;

	/* we could definitely support a lot more machines here */
	fp-&gt;magic = ELF_MAG;
	fp-&gt;hdrsz = (ep-&gt;ehsize+ep-&gt;phnum*ep-&gt;phentsize+16)&amp;~15;
	switch(ep-&gt;machine) {
	case I386:
		mach = &amp;mi386;
		fp-&gt;type = FI386;
		break;
	case MIPS:
		mach = &amp;mmips;
		fp-&gt;type = FMIPS;
		break;
	case SPARC64:
		mach = &amp;msparc64;
		fp-&gt;type = FSPARC64;
		break;
	case POWER:
		mach = &amp;mpower;
		fp-&gt;type = FPOWER;
		break;
	case ARM:
		mach = &amp;marm;
		fp-&gt;type = FARM;
		break;
	default:
		return 0;
	}

	if(ep-&gt;phentsize != sizeof(Phdr32)) {
		werrstr(&#34;bad ELF header size&#34;);
		return 0;
	}
	phsz = sizeof(Phdr32)*ep-&gt;phnum;
	ph = malloc(phsz);
	if(!ph)
		return 0;
	seek(fd, ep-&gt;phoff, 0);
	if(read(fd, ph, phsz) &lt; 0) {
		free(ph);
		return 0;
	}
	hswal(ph, phsz/sizeof(uint32), swal);

	shsz = sizeof(Shdr32)*ep-&gt;shnum;
	sh = malloc(shsz);
	if(sh) {
		seek(fd, ep-&gt;shoff, 0);
		if(read(fd, sh, shsz) &lt; 0) {
			free(sh);
			sh = 0;
		} else
			hswal(sh, shsz/sizeof(uint32), swal);
	}

	/* find text, data and symbols and install them */
	it = id = is = -1;
	for(i = 0; i &lt; ep-&gt;phnum; i++) {
		if(ph[i].type == LOAD
		&amp;&amp; (ph[i].flags &amp; (R|X)) == (R|X) &amp;&amp; it == -1)
			it = i;
		else if(ph[i].type == LOAD
		&amp;&amp; (ph[i].flags &amp; (R|W)) == (R|W) &amp;&amp; id == -1)
			id = i;
		else if(ph[i].type == NOPTYPE &amp;&amp; is == -1)
			is = i;
	}
	if(it == -1 || id == -1) {
		/*
		 * The SPARC64 boot image is something of an ELF hack.
		 * Text+Data+BSS are represented by ph[0].  Symbols
		 * are represented by ph[1]:
		 *
		 *		filesz, memsz, vaddr, paddr, off
		 * ph[0] : txtsz+datsz, txtsz+datsz+bsssz, txtaddr-KZERO, datasize, txtoff
		 * ph[1] : symsz, lcsz, 0, 0, symoff
		 */
		if(ep-&gt;machine == SPARC64 &amp;&amp; ep-&gt;phnum == 2) {
			uint32 txtaddr, txtsz, dataddr, bsssz;

			txtaddr = ph[0].vaddr | 0x80000000;
			txtsz = ph[0].filesz - ph[0].paddr;
			dataddr = txtaddr + txtsz;
			bsssz = ph[0].memsz - ph[0].filesz;
			settext(fp, ep-&gt;elfentry | 0x80000000, txtaddr, txtsz, ph[0].offset);
			setdata(fp, dataddr, ph[0].paddr, ph[0].offset + txtsz, bsssz);
			setsym(fp, ph[1].filesz, 0, ph[1].memsz, ph[1].offset);
			free(ph);
			return 1;
		}

		werrstr(&#34;No TEXT or DATA sections&#34;);
error:
		free(sh);
		free(ph);
		return 0;
	}

	settext(fp, ep-&gt;elfentry, ph[it].vaddr, ph[it].memsz, ph[it].offset);
	setdata(fp, ph[id].vaddr, ph[id].filesz, ph[id].offset, ph[id].memsz - ph[id].filesz);
	if(is != -1)
		setsym(fp, ph[is].filesz, 0, ph[is].memsz, ph[is].offset);
	else if(sh != 0){
		char *buf;
		uvlong symsize = 0;
		uvlong symoff = 0;
		uvlong pclnsz = 0;

		/* load shstrtab names */
		buf = malloc(sh[ep-&gt;shstrndx].size);
		if (buf == 0)
			goto done;
		memset(buf, 0, sizeof buf);
		seek(fd, sh[ep-&gt;shstrndx].offset, 0);
		read(fd, buf, sh[ep-&gt;shstrndx].size);

		for(i = 0; i &lt; ep-&gt;shnum; i++) {
			if (strcmp(&amp;buf[sh[i].name], &#34;.gosymtab&#34;) == 0) {
				symsize = sh[i].size;
				symoff = sh[i].offset;
			}
			if (strcmp(&amp;buf[sh[i].name], &#34;.gopclntab&#34;) == 0) {
				if (sh[i].offset != symoff+symsize) {
					werrstr(&#34;pc line table not contiguous with symbol table&#34;);
					free(buf);
					goto error;
				}
				pclnsz = sh[i].size;
			}
		}
		setsym(fp, symsize, 0, pclnsz, symoff);
		free(buf);
	}
done:
	free(sh);
	free(ph);
	return 1;
}

static int
machdotout(int fd, Fhdr *fp, ExecHdr *hp)
{
	uvlong (*swav)(uvlong);
	uint32 (*swal)(uint32);
	ushort (*swab)(ushort);
	Machhdr *mp;
	MachCmd **cmd;
	MachSymSeg *symtab;
	MachSymSeg *pclntab;
	MachSeg64 *seg;
	MachSect64 *sect;
	MachSeg32 *seg32;
	MachSect32 *sect32;
	uvlong textsize, datasize, bsssize;
	uchar *cmdbuf;
	uchar *cmdp;
	int i, hdrsize;
	uint32 textva, textoff, datava, dataoff;

	mp = &amp;hp-&gt;e.machhdr;
	if (leswal(mp-&gt;filetype) != MACH_EXECUTABLE_TYPE) {
		werrstr(&#34;bad MACH executable type %#ux&#34;, leswal(mp-&gt;filetype));
		return 0;
	}

	swab = leswab;
	swal = leswal;
	swav = leswav;

	mp-&gt;magic = swal(mp-&gt;magic);
	mp-&gt;cputype = swal(mp-&gt;cputype);
	mp-&gt;cpusubtype = swal(mp-&gt;cpusubtype);
	mp-&gt;filetype = swal(mp-&gt;filetype);
	mp-&gt;ncmds = swal(mp-&gt;ncmds);
	mp-&gt;sizeofcmds = swal(mp-&gt;sizeofcmds);
	mp-&gt;flags = swal(mp-&gt;flags);
	mp-&gt;reserved = swal(mp-&gt;reserved);
	hdrsize = 0;

	switch(mp-&gt;magic) {
	case 0xFEEDFACE:	// 32-bit mach
		if (mp-&gt;cputype != MACH_CPU_TYPE_X86) {
			werrstr(&#34;bad MACH cpu type - not 386&#34;);
			return 0;
		}
		if (mp-&gt;cpusubtype != MACH_CPU_SUBTYPE_X86) {
			werrstr(&#34;bad MACH cpu subtype - not 386&#34;);
			return 0;
		}
		if (mp-&gt;filetype != MACH_EXECUTABLE_TYPE) {
			werrstr(&#34;bad MACH executable type&#34;);
			return 0;
		}
		mach = &amp;mi386;
		fp-&gt;type = FI386;
		hdrsize = 28;
		break;

	case 0xFEEDFACF:	// 64-bit mach
		if (mp-&gt;cputype != MACH_CPU_TYPE_X86_64) {
			werrstr(&#34;bad MACH cpu type - not amd64&#34;);
			return 0;
		}

		if (mp-&gt;cpusubtype != MACH_CPU_SUBTYPE_X86) {
			werrstr(&#34;bad MACH cpu subtype - not amd64&#34;);
			return 0;
		}
		mach = &amp;mamd64;
		fp-&gt;type = FAMD64;
		hdrsize = 32;
		break;

	default:
		werrstr(&#34;not mach %#ux&#34;, mp-&gt;magic);
		return 0;
	}

	cmdbuf = malloc(mp-&gt;sizeofcmds);
	seek(fd, hdrsize, 0);
	if(read(fd, cmdbuf, mp-&gt;sizeofcmds) != mp-&gt;sizeofcmds) {
		free(cmdbuf);
		return 0;
	}
	cmd = malloc(mp-&gt;ncmds * sizeof(MachCmd*));
	cmdp = cmdbuf;
	textva = 0;
	textoff = 0;
	dataoff = 0;
	datava = 0;
	symtab = 0;
	pclntab = 0;
	textsize = datasize = bsssize = 0;
	for (i = 0; i &lt; mp-&gt;ncmds; i++) {
		MachCmd *c;

		cmd[i] = (MachCmd*)cmdp;
		c = cmd[i];
		c-&gt;type = swal(c-&gt;type);
		c-&gt;size = swal(c-&gt;size);
		switch(c-&gt;type) {
		case MACH_SEGMENT_32:
			if(mp-&gt;magic != 0xFEEDFACE) {
				werrstr(&#34;segment 32 in mach 64&#34;);
				goto bad;
			}
			seg32 = (MachSeg32*)c;
			seg32-&gt;vmaddr = swav(seg32-&gt;vmaddr);
			seg32-&gt;vmsize = swav(seg32-&gt;vmsize);
			seg32-&gt;fileoff = swav(seg32-&gt;fileoff);
			seg32-&gt;filesize = swav(seg32-&gt;filesize);
			seg32-&gt;maxprot = swal(seg32-&gt;maxprot);
			seg32-&gt;initprot = swal(seg32-&gt;initprot);
			seg32-&gt;nsects = swal(seg32-&gt;nsects);
			seg32-&gt;flags = swal(seg32-&gt;flags);
			if (strcmp(seg32-&gt;segname, &#34;__TEXT&#34;) == 0) {
				textva = seg32-&gt;vmaddr;
				textoff = seg32-&gt;fileoff;
				sect32 = (MachSect32*)(cmdp + sizeof(MachSeg32));
				if (strcmp(sect32-&gt;sectname, &#34;__text&#34;) == 0) {
					textsize = swal(sect32-&gt;size);
				} else {
					werrstr(&#34;no text section&#34;);
					goto bad;
				}
			}
			if (strcmp(seg32-&gt;segname, &#34;__DATA&#34;) == 0) {
				datava = seg32-&gt;vmaddr;
				dataoff = seg32-&gt;fileoff;
				sect32 = (MachSect32*)(cmdp + sizeof(MachSeg32));
				if (strcmp(sect32-&gt;sectname, &#34;__data&#34;) == 0) {
					datasize = swal(sect32-&gt;size);
				} else {
					werrstr(&#34;no data section&#34;);
					goto bad;
				}
				sect32++;
				if (strcmp(sect32-&gt;sectname, &#34;__nl_symbol_ptr&#34;) == 0)
					sect32++;
				if (strcmp(sect32-&gt;sectname, &#34;__bss&#34;) == 0) {
					bsssize = swal(sect32-&gt;size);
				} else {
					werrstr(&#34;no bss section&#34;);
					goto bad;
				}
			}
			break;

		case MACH_SEGMENT_64:
			if(mp-&gt;magic != 0xFEEDFACF) {
				werrstr(&#34;segment 32 in mach 64&#34;);
				goto bad;
			}
			seg = (MachSeg64*)c;
			seg-&gt;vmaddr = swav(seg-&gt;vmaddr);
			seg-&gt;vmsize = swav(seg-&gt;vmsize);
			seg-&gt;fileoff = swav(seg-&gt;fileoff);
			seg-&gt;filesize = swav(seg-&gt;filesize);
			seg-&gt;maxprot = swal(seg-&gt;maxprot);
			seg-&gt;initprot = swal(seg-&gt;initprot);
			seg-&gt;nsects = swal(seg-&gt;nsects);
			seg-&gt;flags = swal(seg-&gt;flags);
			if (strcmp(seg-&gt;segname, &#34;__TEXT&#34;) == 0) {
				textva = seg-&gt;vmaddr;
				textoff = seg-&gt;fileoff;
				sect = (MachSect64*)(cmdp + sizeof(MachSeg64));
				if (strcmp(sect-&gt;sectname, &#34;__text&#34;) == 0) {
					textsize = swav(sect-&gt;size);
				} else {
					werrstr(&#34;no text section&#34;);
					goto bad;
				}
			}
			if (strcmp(seg-&gt;segname, &#34;__DATA&#34;) == 0) {
				datava = seg-&gt;vmaddr;
				dataoff = seg-&gt;fileoff;
				sect = (MachSect64*)(cmdp + sizeof(MachSeg64));
				if (strcmp(sect-&gt;sectname, &#34;__data&#34;) == 0) {
					datasize = swav(sect-&gt;size);
				} else {
					werrstr(&#34;no data section&#34;);
					goto bad;
				}
				sect++;
				if (strcmp(sect-&gt;sectname, &#34;__bss&#34;) == 0) {
					bsssize = swav(sect-&gt;size);
				} else {
					werrstr(&#34;no bss section&#34;);
					goto bad;
				}
			}
			break;
		case MACH_UNIXTHREAD:
			break;
		case MACH_SYMSEG:
			if (symtab == 0)
				symtab = (MachSymSeg*)c;
			else if (pclntab == 0)
				pclntab = (MachSymSeg*)c;
			break;
		}
		cmdp += c-&gt;size;
	}
	if (textva == 0 || datava == 0) {
		free(cmd);
		free(cmdbuf);
		return 0;
	}
	/* compute entry by taking address after header - weird - BUG? */
	settext(fp, textva+sizeof(Machhdr) + mp-&gt;sizeofcmds, textva, textsize, textoff);
	setdata(fp, datava, datasize, dataoff, bsssize);
	if(symtab != 0)
		setsym(fp, symtab-&gt;filesize, 0, pclntab? pclntab-&gt;filesize : 0, symtab-&gt;fileoff);
	free(cmd);
	free(cmdbuf);
	return 1;
bad:
	free(cmd);
	free(cmdbuf);
	return 0;
}

/*
 * (Free|Net)BSD ARM header.
 */
static int
armdotout(int fd, Fhdr *fp, ExecHdr *hp)
{
	uvlong kbase;

	USED(fd);
	settext(fp, hp-&gt;e.exechdr.entry, sizeof(Exec), hp-&gt;e.exechdr.text, sizeof(Exec));
	setdata(fp, fp-&gt;txtsz, hp-&gt;e.exechdr.data, fp-&gt;txtsz, hp-&gt;e.exechdr.bss);
	setsym(fp, hp-&gt;e.exechdr.syms, hp-&gt;e.exechdr.spsz, hp-&gt;e.exechdr.pcsz, fp-&gt;datoff+fp-&gt;datsz);

	kbase = 0xF0000000;
	if ((fp-&gt;entry &amp; kbase) == kbase) {		/* Boot image */
		fp-&gt;txtaddr = kbase+sizeof(Exec);
		fp-&gt;name = &#34;ARM *BSD boot image&#34;;
		fp-&gt;hdrsz = 0;		/* header stripped */
		fp-&gt;dataddr = kbase+fp-&gt;txtsz;
	}
	return 1;
}

static void
settext(Fhdr *fp, uvlong e, uvlong a, int32 s, vlong off)
{
	fp-&gt;txtaddr = a;
	fp-&gt;entry = e;
	fp-&gt;txtsz = s;
	fp-&gt;txtoff = off;
}

static void
setdata(Fhdr *fp, uvlong a, int32 s, vlong off, int32 bss)
{
	fp-&gt;dataddr = a;
	fp-&gt;datsz = s;
	fp-&gt;datoff = off;
	fp-&gt;bsssz = bss;
}

static void
setsym(Fhdr *fp, int32 symsz, int32 sppcsz, int32 lnpcsz, vlong symoff)
{
	fp-&gt;symsz = symsz;
	fp-&gt;symoff = symoff;
	fp-&gt;sppcsz = sppcsz;
	fp-&gt;sppcoff = fp-&gt;symoff+fp-&gt;symsz;
	fp-&gt;lnpcsz = lnpcsz;
	fp-&gt;lnpcoff = fp-&gt;sppcoff+fp-&gt;sppcsz;
}


static uvlong
_round(uvlong a, uint32 b)
{
	uvlong w;

	w = (a/b)*b;
	if (a!=w)
		w += b;
	return(w);
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
