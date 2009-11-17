<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5l/asm.c</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5l/asm.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5l/asm.c
// http://code.google.com/p/inferno-os/source/browse/utils/5l/asm.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
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

#include	&#34;l.h&#34;
#include	&#34;../ld/lib.h&#34;

int32	OFFSET;

static Prog *PP;

int32
entryvalue(void)
{
	char *a;
	Sym *s;

	a = INITENTRY;
	if(*a &gt;= &#39;0&#39; &amp;&amp; *a &lt;= &#39;9&#39;)
		return atolwhex(a);
	s = lookup(a, 0);
	if(s-&gt;type == 0)
		return INITTEXT;
	switch(s-&gt;type) {
	case STEXT:
	case SLEAF:
		break;
	case SDATA:
		if(dlm)
			return s-&gt;value+INITDAT;
	default:
		diag(&#34;entry not text: %s&#34;, s-&gt;name);
	}
	return s-&gt;value;
}

void
asmb(void)
{
	Prog *p;
	int32 t, etext;
	int np;
	vlong va, fo, w, symo;
	int strtabsize;
	vlong symdatva = 0x99LL&lt;&lt;24;
	Optab *o;

	strtabsize = 0;
	symo = 0;

	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f asm\n&#34;, cputime());
	Bflush(&amp;bso);
	OFFSET = HEADR;
	seek(cout, OFFSET, 0);
	pc = INITTEXT;
	for(p = firstp; p != P; p = p-&gt;link) {
		setarch(p);
		if(p-&gt;as == ATEXT) {
			curtext = p;
			autosize = p-&gt;to.offset + 4;
		}
		if(p-&gt;pc != pc) {
			diag(&#34;phase error %lux sb %lux&#34;,
				p-&gt;pc, pc);
			if(!debug[&#39;a&#39;])
				prasm(curp);
			pc = p-&gt;pc;
		}
		curp = p;
		o = oplook(p);	/* could probably avoid this call */
		if(thumb)
			thumbasmout(p, o);
		else
			asmout(p, o);
		pc += o-&gt;size;
	}
	while(pc-INITTEXT &lt; textsize) {
		cput(0);
		pc++;
	}

	if(debug[&#39;a&#39;])
		Bprint(&amp;bso, &#34;\n&#34;);
	Bflush(&amp;bso);
	cflush();

	/* output strings in text segment */
	etext = INITTEXT + textsize;
	for(t = pc; t &lt; etext; t += sizeof(buf)-100) {
		if(etext-t &gt; sizeof(buf)-100)
			datblk(t, sizeof(buf)-100, 1);
		else
			datblk(t, etext-t, 1);
	}

	/* output section header strings */
	curtext = P;
	switch(HEADTYPE) {
	case 0:
	case 1:
	case 2:
	case 5:
		OFFSET = HEADR+textsize;
		seek(cout, OFFSET, 0);
		break;
	case 3:
		OFFSET = rnd(HEADR+textsize, 4096);
		seek(cout, OFFSET, 0);
		break;
	case 6:
		seek(cout, rnd(HEADR+textsize, INITRND)+datsize, 0);
		strtabsize = linuxstrtable();
		cflush();
		t = rnd(HEADR+textsize, INITRND);
		seek(cout, t, 0);
		break;
	}
	if(dlm){
		char buf[8];

		write(cout, buf, INITDAT-textsize);
		textsize = INITDAT;
	}
	for(t = 0; t &lt; datsize; t += sizeof(buf)-100) {
		if(datsize-t &gt; sizeof(buf)-100)
			datblk(t, sizeof(buf)-100, 0);
		else
			datblk(t, datsize-t, 0);
	}
	cflush();

	/* output symbol table */
	symsize = 0;
	lcsize = 0;
	if(!debug[&#39;s&#39;]) {
		if(debug[&#39;v&#39;])
			Bprint(&amp;bso, &#34;%5.2f sym\n&#34;, cputime());
		Bflush(&amp;bso);
		switch(HEADTYPE) {
		case 0:
		case 1:
		case 4:
		case 5:
			debug[&#39;s&#39;] = 1;
			break;
		case 2:
			OFFSET = HEADR+textsize+datsize;
			seek(cout, OFFSET, 0);
			break;
		case 3:
			OFFSET += rnd(datsize, 4096);
			seek(cout, OFFSET, 0);
			break;
		case 6:
			symo = rnd(HEADR+textsize, INITRND)+datsize+strtabsize;
			symo = rnd(symo, INITRND);
			seek(cout, symo + 8, 0);
			break;
		}
		if(!debug[&#39;s&#39;])
			asmsym();
		if(debug[&#39;v&#39;])
			Bprint(&amp;bso, &#34;%5.2f pc\n&#34;, cputime());
		Bflush(&amp;bso);
		if(!debug[&#39;s&#39;])
			asmlc();
		if(!debug[&#39;s&#39;])
			asmthumbmap();
		if(dlm)
			asmdyn();
		cflush();
	}
	else if(dlm){
		seek(cout, HEADR+textsize+datsize, 0);
		asmdyn();
		cflush();
	}

	curtext = P;
	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f header\n&#34;, cputime());
	Bflush(&amp;bso);
	OFFSET = 0;
	seek(cout, OFFSET, 0);
	switch(HEADTYPE) {
	case 0:	/* no header */
		break;
	case 1:	/* aif for risc os */
		lputl(0xe1a00000);		/* NOP - decompress code */
		lputl(0xe1a00000);		/* NOP - relocation code */
		lputl(0xeb000000 + 12);		/* BL - zero init code */
		lputl(0xeb000000 +
			(entryvalue()
			 - INITTEXT
			 + HEADR
			 - 12
			 - 8) / 4);		/* BL - entry code */

		lputl(0xef000011);		/* SWI - exit code */
		lputl(textsize+HEADR);		/* text size */
		lputl(datsize);			/* data size */
		lputl(0);			/* sym size */

		lputl(bsssize);			/* bss size */
		lputl(0);			/* sym type */
		lputl(INITTEXT-HEADR);		/* text addr */
		lputl(0);			/* workspace - ignored */

		lputl(32);			/* addr mode / data addr flag */
		lputl(0);			/* data addr */
		for(t=0; t&lt;2; t++)
			lputl(0);		/* reserved */

		for(t=0; t&lt;15; t++)
			lputl(0xe1a00000);	/* NOP - zero init code */
		lputl(0xe1a0f00e);		/* B (R14) - zero init return */
		break;
	case 2:	/* plan 9 */
		if(dlm)
			lput(0x80000000|0x647);	/* magic */
		else
			lput(0x647);			/* magic */
		lput(textsize);			/* sizes */
		lput(datsize);
		lput(bsssize);
		lput(symsize);			/* nsyms */
		lput(entryvalue());		/* va of entry */
		lput(0L);
		lput(lcsize);
		break;
	case 3:	/* boot for NetBSD */
		lput((143&lt;&lt;16)|0413);		/* magic */
		lputl(rnd(HEADR+textsize, 4096));
		lputl(rnd(datsize, 4096));
		lputl(bsssize);
		lputl(symsize);			/* nsyms */
		lputl(entryvalue());		/* va of entry */
		lputl(0L);
		lputl(0L);
		break;
	case 4: /* boot for IXP1200 */
		break;
	case 5: /* boot for ipaq */
		lputl(0xe3300000);		/* nop */
		lputl(0xe3300000);		/* nop */
		lputl(0xe3300000);		/* nop */
		lputl(0xe3300000);		/* nop */
		break;
	case 6:
		/* elf arm */
		strnput(&#34;\177ELF&#34;, 4);		/* e_ident */
		cput(1);			/* class = 32 bit */
		cput(1);			/* data = LSB */
		cput(1);			/* version = CURRENT */
		strnput(&#34;&#34;, 9);

		wputl(2);			/* type = EXEC */
		wputl(40);			/* machine = ARM */
		lputl(1L);			/* version = CURRENT */
		lputl(entryvalue());		/* entry vaddr */
		lputl(52L);			/* offset to first phdr */
		np = 3;
		if(!debug[&#39;s&#39;])
			np++;
		lputl(52L+32*np);		/* offset to first shdr */
		lputl(0L);			/* processor specific flags */
		wputl(52);			/* Ehdr size */
		wputl(32);			/* Phdr size */
		wputl(np);			/* # of Phdrs */
		wputl(40);			/* Shdr size */
		if (!debug[&#39;s&#39;])
			wputl(7);			/* # of Shdrs */
		else
			wputl(5);			/* # of Shdrs */
		wputl(4);			/* Shdr with strings */

		fo = 0;
		va = INITTEXT &amp; ~((vlong)INITRND - 1);
		w = HEADR+textsize;

		linuxphdr(1,			/* text - type = PT_LOAD */
			1L+4L,			/* text - flags = PF_X+PF_R */
			0,			/* file offset */
			va,			/* vaddr */
			va,			/* paddr */
			w,			/* file size */
			w,			/* memory size */
			INITRND);		/* alignment */

		fo = rnd(fo+w, INITRND);
		va = rnd(va+w, INITRND);
		w = datsize;

		linuxphdr(1,			/* data - type = PT_LOAD */
			2L+4L,			/* data - flags = PF_W+PF_R */
			fo,			/* file offset */
			va,			/* vaddr */
			va,			/* paddr */
			w,			/* file size */
			w+bsssize,		/* memory size */
			INITRND);		/* alignment */

		if(!debug[&#39;s&#39;]) {
			linuxphdr(1,			/* data - type = PT_LOAD */
				2L+4L,			/* data - flags = PF_W+PF_R */
				symo,		/* file offset */
				symdatva,			/* vaddr */
				symdatva,			/* paddr */
				8+symsize+lcsize,			/* file size */
				8+symsize+lcsize,		/* memory size */
				INITRND);		/* alignment */
		}

		linuxphdr(0x6474e551,		/* gok - type = gok */
			1L+2L+4L,		/* gok - flags = PF_X+PF_W+PF_R */
			0,			/* file offset */
			0,			/* vaddr */
			0,			/* paddr */
			0,			/* file size */
			0,			/* memory size */
			8);			/* alignment */

		linuxshdr(nil,			/* name */
			0,			/* type */
			0,			/* flags */
			0,			/* addr */
			0,			/* off */
			0,			/* size */
			0,			/* link */
			0,			/* info */
			0,			/* align */
			0);			/* entsize */

		stroffset = 1;  /* 0 means no name, so start at 1 */
		fo = HEADR;
		va = (INITTEXT &amp; ~((vlong)INITRND - 1)) + HEADR;
		w = textsize;

		linuxshdr(&#34;.text&#34;,		/* name */
			1,			/* type */
			6,			/* flags */
			va,			/* addr */
			fo,			/* off */
			w,			/* size */
			0,			/* link */
			0,			/* info */
			8,			/* align */
			0);			/* entsize */

		fo = rnd(fo+w, INITRND);
		va = rnd(va+w, INITRND);
		w = datsize;

		linuxshdr(&#34;.data&#34;,		/* name */
			1,			/* type */
			3,			/* flags */
			va,			/* addr */
			fo,			/* off */
			w,			/* size */
			0,			/* link */
			0,			/* info */
			8,			/* align */
			0);			/* entsize */

		fo += w;
		va += w;
		w = bsssize;

		linuxshdr(&#34;.bss&#34;,		/* name */
			8,			/* type */
			3,			/* flags */
			va,			/* addr */
			fo,			/* off */
			w,			/* size */
			0,			/* link */
			0,			/* info */
			8,			/* align */
			0);			/* entsize */

		w = strtabsize;

		linuxshdr(&#34;.shstrtab&#34;,		/* name */
			3,			/* type */
			0,			/* flags */
			0,			/* addr */
			fo,			/* off */
			w,			/* size */
			0,			/* link */
			0,			/* info */
			1,			/* align */
			0);			/* entsize */

		if (debug[&#39;s&#39;])
			break;

		fo = symo+8;
		w = symsize;

		linuxshdr(&#34;.gosymtab&#34;,		/* name */
			1,			/* type 1 = SHT_PROGBITS */
			0,			/* flags */
			0,			/* addr */
			fo,			/* off */
			w,			/* size */
			0,			/* link */
			0,			/* info */
			1,			/* align */
			24);			/* entsize */

		fo += w;
		w = lcsize;

		linuxshdr(&#34;.gopclntab&#34;,		/* name */
			1,			/* type 1 = SHT_PROGBITS*/
			0,			/* flags */
			0,			/* addr */
			fo,			/* off */
			w,			/* size */
			0,			/* link */
			0,			/* info */
			1,			/* align */
			24);			/* entsize */
		break;
	}
	cflush();
	if(debug[&#39;c&#39;]){
		print(&#34;textsize=%ld\n&#34;, textsize);
		print(&#34;datsize=%ld\n&#34;, datsize);
		print(&#34;bsssize=%ld\n&#34;, bsssize);
		print(&#34;symsize=%ld\n&#34;, symsize);
		print(&#34;lcsize=%ld\n&#34;, lcsize);
		print(&#34;total=%ld\n&#34;, textsize+datsize+bsssize+symsize+lcsize);
	}
}

void
strnput(char *s, int n)
{
	for(; *s; s++){
		cput(*s);
		n--;
	}
	for(; n &gt; 0; n--)
		cput(0);
}

void
cput(int c)
{
	cbp[0] = c;
	cbp++;
	cbc--;
	if(cbc &lt;= 0)
		cflush();
}

/*
void
cput(int32 c)
{
	*cbp++ = c;
	if(--cbc &lt;= 0)
		cflush();
}
*/

void
wput(int32 l)
{

	cbp[0] = l&gt;&gt;8;
	cbp[1] = l;
	cbp += 2;
	cbc -= 2;
	if(cbc &lt;= 0)
		cflush();
}

void
wputl(ushort w)
{
	cput(w);
	cput(w&gt;&gt;8);
}


void
hput(int32 l)
{

	cbp[0] = l&gt;&gt;8;
	cbp[1] = l;
	cbp += 2;
	cbc -= 2;
	if(cbc &lt;= 0)
		cflush();
}

void
lput(int32 l)
{

	cbp[0] = l&gt;&gt;24;
	cbp[1] = l&gt;&gt;16;
	cbp[2] = l&gt;&gt;8;
	cbp[3] = l;
	cbp += 4;
	cbc -= 4;
	if(cbc &lt;= 0)
		cflush();
}

void
lputl(int32 l)
{

	cbp[3] = l&gt;&gt;24;
	cbp[2] = l&gt;&gt;16;
	cbp[1] = l&gt;&gt;8;
	cbp[0] = l;
	cbp += 4;
	cbc -= 4;
	if(cbc &lt;= 0)
		cflush();
}

void
cflush(void)
{
	int n;

	/* no bug if cbc &lt; 0 since obuf(cbuf) followed by ibuf in buf! */
	n = sizeof(buf.cbuf) - cbc;
	if(n)
		write(cout, buf.cbuf, n);
	cbp = buf.cbuf;
	cbc = sizeof(buf.cbuf);
}

void
nopstat(char *f, Count *c)
{
	if(c-&gt;outof)
	Bprint(&amp;bso, &#34;%s delay %ld/%ld (%.2f)\n&#34;, f,
		c-&gt;outof - c-&gt;count, c-&gt;outof,
		(double)(c-&gt;outof - c-&gt;count)/c-&gt;outof);
}

void
asmsym(void)
{
	Prog *p;
	Auto *a;
	Sym *s;
	int h;

	s = lookup(&#34;etext&#34;, 0);
	if(s-&gt;type == STEXT)
		putsymb(s-&gt;name, &#39;T&#39;, s-&gt;value, s-&gt;version);

	for(h=0; h&lt;NHASH; h++)
		for(s=hash[h]; s!=S; s=s-&gt;link)
			switch(s-&gt;type) {
			case SCONST:
				putsymb(s-&gt;name, &#39;D&#39;, s-&gt;value, s-&gt;version);
				continue;

			case SDATA:
				putsymb(s-&gt;name, &#39;D&#39;, s-&gt;value+INITDAT, s-&gt;version);
				continue;

			case SBSS:
				putsymb(s-&gt;name, &#39;B&#39;, s-&gt;value+INITDAT, s-&gt;version);
				continue;

			case SSTRING:
				putsymb(s-&gt;name, &#39;T&#39;, s-&gt;value, s-&gt;version);
				continue;

			case SFILE:
				putsymb(s-&gt;name, &#39;f&#39;, s-&gt;value, s-&gt;version);
				continue;
			}

	for(p=textp; p!=P; p=p-&gt;cond) {
		s = p-&gt;from.sym;
		if(s-&gt;type != STEXT &amp;&amp; s-&gt;type != SLEAF)
			continue;

		/* filenames first */
		for(a=p-&gt;to.autom; a; a=a-&gt;link)
			if(a-&gt;type == D_FILE)
				putsymb(a-&gt;asym-&gt;name, &#39;z&#39;, a-&gt;aoffset, 0);
			else
			if(a-&gt;type == D_FILE1)
				putsymb(a-&gt;asym-&gt;name, &#39;Z&#39;, a-&gt;aoffset, 0);

		if(!s-&gt;reachable)
			continue;

		if(s-&gt;type == STEXT)
			putsymb(s-&gt;name, &#39;T&#39;, s-&gt;value, s-&gt;version);
		else
			putsymb(s-&gt;name, &#39;L&#39;, s-&gt;value, s-&gt;version);

		/* frame, auto and param after */
		putsymb(&#34;.frame&#34;, &#39;m&#39;, p-&gt;to.offset+4, 0);
		for(a=p-&gt;to.autom; a; a=a-&gt;link)
			if(a-&gt;type == D_AUTO)
				putsymb(a-&gt;asym-&gt;name, &#39;a&#39;, -a-&gt;aoffset, 0);
			else
			if(a-&gt;type == D_PARAM)
				putsymb(a-&gt;asym-&gt;name, &#39;p&#39;, a-&gt;aoffset, 0);
	}
	if(debug[&#39;v&#39;] || debug[&#39;n&#39;])
		Bprint(&amp;bso, &#34;symsize = %lud\n&#34;, symsize);
	Bflush(&amp;bso);
}

void
putsymb(char *s, int t, int32 v, int ver)
{
	int i, f;

	if(t == &#39;f&#39;)
		s++;
	lput(v);
	if(ver)
		t += &#39;a&#39; - &#39;A&#39;;
	cput(t+0x80);			/* 0x80 is variable length */

	if(t == &#39;Z&#39; || t == &#39;z&#39;) {
		cput(s[0]);
		for(i=1; s[i] != 0 || s[i+1] != 0; i += 2) {
			cput(s[i]);
			cput(s[i+1]);
		}
		cput(0);
		cput(0);
		i++;
	}
	else {
		for(i=0; s[i]; i++)
			cput(s[i]);
		cput(0);
	}
	// TODO(rsc): handle go parameter
	lput(0);

	symsize += 4 + 1 + i + 1 + 4;

	if(debug[&#39;n&#39;]) {
		if(t == &#39;z&#39; || t == &#39;Z&#39;) {
			Bprint(&amp;bso, &#34;%c %.8lux &#34;, t, v);
			for(i=1; s[i] != 0 || s[i+1] != 0; i+=2) {
				f = ((s[i]&amp;0xff) &lt;&lt; 8) | (s[i+1]&amp;0xff);
				Bprint(&amp;bso, &#34;/%x&#34;, f);
			}
			Bprint(&amp;bso, &#34;\n&#34;);
			return;
		}
		if(ver)
			Bprint(&amp;bso, &#34;%c %.8lux %s&lt;%d&gt;\n&#34;, t, v, s, ver);
		else
			Bprint(&amp;bso, &#34;%c %.8lux %s\n&#34;, t, v, s);
	}
}

#define	MINLC	4
void
asmlc(void)
{
	int32 oldpc, oldlc;
	Prog *p;
	int32 v, s;

	oldpc = INITTEXT;
	oldlc = 0;
	for(p = firstp; p != P; p = p-&gt;link) {
		setarch(p);
		if(p-&gt;line == oldlc || p-&gt;as == ATEXT || p-&gt;as == ANOP) {
			if(p-&gt;as == ATEXT)
				curtext = p;
			if(debug[&#39;L&#39;])
				Bprint(&amp;bso, &#34;%6lux %P\n&#34;,
					p-&gt;pc, p);
			continue;
		}
		if(debug[&#39;L&#39;])
			Bprint(&amp;bso, &#34;\t\t%6ld&#34;, lcsize);
		v = (p-&gt;pc - oldpc) / MINLC;
		while(v) {
			s = 127;
			if(v &lt; 127)
				s = v;
			cput(s+128);	/* 129-255 +pc */
			if(debug[&#39;L&#39;])
				Bprint(&amp;bso, &#34; pc+%ld*%d(%ld)&#34;, s, MINLC, s+128);
			v -= s;
			lcsize++;
		}
		s = p-&gt;line - oldlc;
		oldlc = p-&gt;line;
		oldpc = p-&gt;pc + MINLC;
		if(s &gt; 64 || s &lt; -64) {
			cput(0);	/* 0 vv +lc */
			cput(s&gt;&gt;24);
			cput(s&gt;&gt;16);
			cput(s&gt;&gt;8);
			cput(s);
			if(debug[&#39;L&#39;]) {
				if(s &gt; 0)
					Bprint(&amp;bso, &#34; lc+%ld(%d,%ld)\n&#34;,
						s, 0, s);
				else
					Bprint(&amp;bso, &#34; lc%ld(%d,%ld)\n&#34;,
						s, 0, s);
				Bprint(&amp;bso, &#34;%6lux %P\n&#34;,
					p-&gt;pc, p);
			}
			lcsize += 5;
			continue;
		}
		if(s &gt; 0) {
			cput(0+s);	/* 1-64 +lc */
			if(debug[&#39;L&#39;]) {
				Bprint(&amp;bso, &#34; lc+%ld(%ld)\n&#34;, s, 0+s);
				Bprint(&amp;bso, &#34;%6lux %P\n&#34;,
					p-&gt;pc, p);
			}
		} else {
			cput(64-s);	/* 65-128 -lc */
			if(debug[&#39;L&#39;]) {
				Bprint(&amp;bso, &#34; lc%ld(%ld)\n&#34;, s, 64-s);
				Bprint(&amp;bso, &#34;%6lux %P\n&#34;,
					p-&gt;pc, p);
			}
		}
		lcsize++;
	}
	while(lcsize &amp; 1) {
		s = 129;
		cput(s);
		lcsize++;
	}
	if(debug[&#39;v&#39;] || debug[&#39;L&#39;])
		Bprint(&amp;bso, &#34;lcsize = %ld\n&#34;, lcsize);
	Bflush(&amp;bso);
}

static void
outt(int32 f, int32 l)
{
	if(debug[&#39;L&#39;])
		Bprint(&amp;bso, &#34;tmap: %lux-%lux\n&#34;, f, l);
	lput(f);
	lput(l);
}

void
asmthumbmap(void)
{
	int32 pc, lastt;
	Prog *p;

	if(!seenthumb)
		return;
	pc = 0;
	lastt = -1;
	for(p = firstp; p != P; p = p-&gt;link){
		pc = p-&gt;pc - INITTEXT;
		if(p-&gt;as == ATEXT){
			setarch(p);
			if(thumb){
				if(p-&gt;from.sym-&gt;foreign){	// 8 bytes of ARM first
					if(lastt &gt;= 0){
						outt(lastt, pc-1);
						lastt = -1;
					}
					pc += 8;
				}
				if(lastt &lt; 0)
					lastt = pc;
			}
			else{
				if(p-&gt;from.sym-&gt;foreign){	// 4 bytes of THUMB first
					if(lastt &lt; 0)
						lastt = pc;
					pc += 4;
				}
				if(lastt &gt;= 0){
					outt(lastt, pc-1);
					lastt = -1;
				}
			}
		}
	}
	if(lastt &gt;= 0)
		outt(lastt, pc+1);
}

void
datblk(int32 s, int32 n, int str)
{
	Sym *v;
	Prog *p;
	char *cast;
	int32 a, l, fl, j, d;
	int i, c;

	memset(buf.dbuf, 0, n+100);
	for(p = datap; p != P; p = p-&gt;link) {
		if(str != (p-&gt;from.sym-&gt;type == SSTRING))
			continue;
		curp = p;
		a = p-&gt;from.sym-&gt;value + p-&gt;from.offset;
		l = a - s;
		c = p-&gt;reg;
		i = 0;
		if(l &lt; 0) {
			if(l+c &lt;= 0)
				continue;
			while(l &lt; 0) {
				l++;
				i++;
			}
		}
		if(l &gt;= n)
			continue;
		if(p-&gt;as != AINIT &amp;&amp; p-&gt;as != ADYNT) {
			for(j=l+(c-i)-1; j&gt;=l; j--)
				if(buf.dbuf[j]) {
					print(&#34;%P\n&#34;, p);
					diag(&#34;multiple initialization&#34;);
					break;
				}
		}
		switch(p-&gt;to.type) {
		default:
			diag(&#34;unknown mode in initialization%P&#34;, p);
			break;

		case D_FCONST:
			switch(c) {
			default:
			case 4:
				fl = ieeedtof(p-&gt;to.ieee);
				cast = (char*)&amp;fl;
				for(; i&lt;c; i++) {
					buf.dbuf[l] = cast[fnuxi4[i]];
					l++;
				}
				break;
			case 8:
				cast = (char*)p-&gt;to.ieee;
				for(; i&lt;c; i++) {
					buf.dbuf[l] = cast[fnuxi8[i]];
					l++;
				}
				break;
			}
			break;

		case D_SCONST:
			for(; i&lt;c; i++) {
				buf.dbuf[l] = p-&gt;to.sval[i];
				l++;
			}
			break;

		case D_CONST:
			d = p-&gt;to.offset;
			v = p-&gt;to.sym;
			if(v) {
				switch(v-&gt;type) {
				case SUNDEF:
					ckoff(v, d);
					d += v-&gt;value;
					break;
				case STEXT:
				case SLEAF:
					d += v-&gt;value;
#ifdef CALLEEBX
					d += fnpinc(v);
#else
					if(v-&gt;thumb)
						d++;		// T bit
#endif
					break;
				case SSTRING:
					d += v-&gt;value;
					break;
				case SDATA:
				case SBSS:
					d += v-&gt;value + INITDAT;
					break;
				}
				if(dlm)
					dynreloc(v, a+INITDAT, 1);
			}
			cast = (char*)&amp;d;
			switch(c) {
			default:
				diag(&#34;bad nuxi %d %d%P&#34;, c, i, curp);
				break;
			case 1:
				for(; i&lt;c; i++) {
					buf.dbuf[l] = cast[inuxi1[i]];
					l++;
				}
				break;
			case 2:
				for(; i&lt;c; i++) {
					buf.dbuf[l] = cast[inuxi2[i]];
					l++;
				}
				break;
			case 4:
				for(; i&lt;c; i++) {
					buf.dbuf[l] = cast[inuxi4[i]];
					l++;
				}
				break;
			}
			break;

		case D_SBIG:
			if(debug[&#39;a&#39;] &amp;&amp; i == 0) {
				Bprint(&amp;bso, &#34;\t%P\n&#34;, curp);
			}
			for(; i&lt;c; i++) {
				buf.dbuf[l] = p-&gt;to.sbig[i];
				l++;
			}
			break;
		}
	}
	write(cout, buf.dbuf, n);
}

void
asmout(Prog *p, Optab *o)
{
	int32 o1, o2, o3, o4, o5, o6, v;
	int r, rf, rt, rt2;
	Sym *s;

PP = p;
	o1 = 0;
	o2 = 0;
	o3 = 0;
	o4 = 0;
	o5 = 0;
	o6 = 0;
	armsize += o-&gt;size;
if(debug[&#39;P&#39;]) print(&#34;%ulx: %P	type %d\n&#34;, (uint32)(p-&gt;pc), p, o-&gt;type);
	switch(o-&gt;type) {
	default:
		diag(&#34;unknown asm %d&#34;, o-&gt;type);
		prasm(p);
		break;

	case 0:		/* pseudo ops */
if(debug[&#39;G&#39;]) print(&#34;%ulx: %s: arm %d %d %d %d\n&#34;, (uint32)(p-&gt;pc), p-&gt;from.sym-&gt;name, p-&gt;from.sym-&gt;thumb, p-&gt;from.sym-&gt;foreign, p-&gt;from.sym-&gt;fnptr, p-&gt;from.sym-&gt;used);
		break;

	case 1:		/* op R,[R],R */
		o1 = oprrr(p-&gt;as, p-&gt;scond);
		rf = p-&gt;from.reg;
		rt = p-&gt;to.reg;
		r = p-&gt;reg;
		if(p-&gt;to.type == D_NONE)
			rt = 0;
		if(p-&gt;as == AMOVW || p-&gt;as == AMVN)
			r = 0;
		else if(r == NREG)
			r = rt;
		o1 |= rf | (r&lt;&lt;16) | (rt&lt;&lt;12);
		break;

	case 2:		/* movbu $I,[R],R */
		aclass(&amp;p-&gt;from);
		o1 = oprrr(p-&gt;as, p-&gt;scond);
		o1 |= immrot(instoffset);
		rt = p-&gt;to.reg;
		r = p-&gt;reg;
		if(p-&gt;to.type == D_NONE)
			rt = 0;
		if(p-&gt;as == AMOVW || p-&gt;as == AMVN)
			r = 0;
		else if(r == NREG)
			r = rt;
		o1 |= (r&lt;&lt;16) | (rt&lt;&lt;12);
		break;

	case 3:		/* add R&lt;&lt;[IR],[R],R */
	mov:
		aclass(&amp;p-&gt;from);
		o1 = oprrr(p-&gt;as, p-&gt;scond);
		o1 |= p-&gt;from.offset;
		rt = p-&gt;to.reg;
		r = p-&gt;reg;
		if(p-&gt;to.type == D_NONE)
			rt = 0;
		if(p-&gt;as == AMOVW || p-&gt;as == AMVN)
			r = 0;
		else if(r == NREG)
			r = rt;
		o1 |= (r&lt;&lt;16) | (rt&lt;&lt;12);
		break;

	case 4:		/* add $I,[R],R */
		aclass(&amp;p-&gt;from);
		o1 = oprrr(AADD, p-&gt;scond);
		o1 |= immrot(instoffset);
		r = p-&gt;from.reg;
		if(r == NREG)
			r = o-&gt;param;
		o1 |= r &lt;&lt; 16;
		o1 |= p-&gt;to.reg &lt;&lt; 12;
		break;

	case 5:		/* bra s */
		v = -8;
		if(p-&gt;cond == UP) {
			s = p-&gt;to.sym;
			if(s-&gt;type != SUNDEF)
				diag(&#34;bad branch sym type&#34;);
			v = (uint32)s-&gt;value &gt;&gt; (Roffset-2);
			dynreloc(s, p-&gt;pc, 0);
		}
		else if(p-&gt;cond != P)
			v = (p-&gt;cond-&gt;pc - pc) - 8;
#ifdef CALLEEBX
		if(p-&gt;as == ABL)
			v += fninc(p-&gt;to.sym);
#endif
		o1 = opbra(p-&gt;as, p-&gt;scond);
		o1 |= (v &gt;&gt; 2) &amp; 0xffffff;
		break;

	case 6:		/* b ,O(R) -&gt; add $O,R,PC */
		aclass(&amp;p-&gt;to);
		o1 = oprrr(AADD, p-&gt;scond);
		o1 |= immrot(instoffset);
		o1 |= p-&gt;to.reg &lt;&lt; 16;
		o1 |= REGPC &lt;&lt; 12;
		break;

	case 7:		/* bl ,O(R) -&gt; mov PC,link; add $O,R,PC */
		aclass(&amp;p-&gt;to);
		o1 = oprrr(AADD, p-&gt;scond);
		o1 |= immrot(0);
		o1 |= REGPC &lt;&lt; 16;
		o1 |= REGLINK &lt;&lt; 12;

		o2 = oprrr(AADD, p-&gt;scond);
		o2 |= immrot(instoffset);
		o2 |= p-&gt;to.reg &lt;&lt; 16;
		o2 |= REGPC &lt;&lt; 12;
		break;

	case 8:		/* sll $c,[R],R -&gt; mov (R&lt;&lt;$c),R */
		aclass(&amp;p-&gt;from);
		o1 = oprrr(p-&gt;as, p-&gt;scond);
		r = p-&gt;reg;
		if(r == NREG)
			r = p-&gt;to.reg;
		o1 |= r;
		o1 |= (instoffset&amp;31) &lt;&lt; 7;
		o1 |= p-&gt;to.reg &lt;&lt; 12;
		break;

	case 9:		/* sll R,[R],R -&gt; mov (R&lt;&lt;R),R */
		o1 = oprrr(p-&gt;as, p-&gt;scond);
		r = p-&gt;reg;
		if(r == NREG)
			r = p-&gt;to.reg;
		o1 |= r;
		o1 |= (p-&gt;from.reg &lt;&lt; 8) | (1&lt;&lt;4);
		o1 |= p-&gt;to.reg &lt;&lt; 12;
		break;

	case 10:	/* swi [$con] */
		o1 = oprrr(p-&gt;as, p-&gt;scond);
		if(p-&gt;to.type != D_NONE) {
			aclass(&amp;p-&gt;to);
			o1 |= instoffset &amp; 0xffffff;
		}
		break;

	case 11:	/* word */
		switch(aclass(&amp;p-&gt;to)) {
		case C_LCON:
			if(!dlm)
				break;
			if(p-&gt;to.name != D_EXTERN &amp;&amp; p-&gt;to.name != D_STATIC)
				break;
		case C_ADDR:
			if(p-&gt;to.sym-&gt;type == SUNDEF)
				ckoff(p-&gt;to.sym, p-&gt;to.offset);
			dynreloc(p-&gt;to.sym, p-&gt;pc, 1);
		}
		o1 = instoffset;
		break;

	case 12:	/* movw $lcon, reg */
		o1 = omvl(p, &amp;p-&gt;from, p-&gt;to.reg);
		break;

	case 13:	/* op $lcon, [R], R */
		o1 = omvl(p, &amp;p-&gt;from, REGTMP);
		if(!o1)
			break;
		o2 = oprrr(p-&gt;as, p-&gt;scond);
		o2 |= REGTMP;
		r = p-&gt;reg;
		if(p-&gt;as == AMOVW || p-&gt;as == AMVN)
			r = 0;
		else if(r == NREG)
			r = p-&gt;to.reg;
		o2 |= r &lt;&lt; 16;
		if(p-&gt;to.type != D_NONE)
			o2 |= p-&gt;to.reg &lt;&lt; 12;
		break;

	case 14:	/* movb/movbu/movh/movhu R,R */
		o1 = oprrr(ASLL, p-&gt;scond);

		if(p-&gt;as == AMOVBU || p-&gt;as == AMOVHU)
			o2 = oprrr(ASRL, p-&gt;scond);
		else
			o2 = oprrr(ASRA, p-&gt;scond);

		r = p-&gt;to.reg;
		o1 |= (p-&gt;from.reg)|(r&lt;&lt;12);
		o2 |= (r)|(r&lt;&lt;12);
		if(p-&gt;as == AMOVB || p-&gt;as == AMOVBU) {
			o1 |= (24&lt;&lt;7);
			o2 |= (24&lt;&lt;7);
		} else {
			o1 |= (16&lt;&lt;7);
			o2 |= (16&lt;&lt;7);
		}
		break;

	case 15:	/* mul r,[r,]r */
		o1 = oprrr(p-&gt;as, p-&gt;scond);
		rf = p-&gt;from.reg;
		rt = p-&gt;to.reg;
		r = p-&gt;reg;
		if(r == NREG)
			r = rt;
		if(rt == r) {
			r = rf;
			rf = rt;
		}
		if(0)
		if(rt == r || rf == REGPC || r == REGPC || rt == REGPC) {
			diag(&#34;bad registers in MUL&#34;);
			prasm(p);
		}
		o1 |= (rf&lt;&lt;8) | r | (rt&lt;&lt;16);
		break;


	case 16:	/* div r,[r,]r */
		o1 = 0xf &lt;&lt; 28;
		o2 = 0;
		break;

	case 17:
		o1 = oprrr(p-&gt;as, p-&gt;scond);
		rf = p-&gt;from.reg;
		rt = p-&gt;to.reg;
		rt2 = p-&gt;to.offset;
		r = p-&gt;reg;
		o1 |= (rf&lt;&lt;8) | r | (rt&lt;&lt;16) | (rt2&lt;&lt;12);
		break;

	case 20:	/* mov/movb/movbu R,O(R) */
		aclass(&amp;p-&gt;to);
		r = p-&gt;to.reg;
		if(r == NREG)
			r = o-&gt;param;
		o1 = osr(p-&gt;as, p-&gt;from.reg, instoffset, r, p-&gt;scond);
		break;

	case 21:	/* mov/movbu O(R),R -&gt; lr */
		aclass(&amp;p-&gt;from);
		r = p-&gt;from.reg;
		if(r == NREG)
			r = o-&gt;param;
		o1 = olr(instoffset, r, p-&gt;to.reg, p-&gt;scond);
		if(p-&gt;as != AMOVW)
			o1 |= 1&lt;&lt;22;
		break;

	case 22:	/* movb/movh/movhu O(R),R -&gt; lr,shl,shr */
		aclass(&amp;p-&gt;from);
		r = p-&gt;from.reg;
		if(r == NREG)
			r = o-&gt;param;
		o1 = olr(instoffset, r, p-&gt;to.reg, p-&gt;scond);

		o2 = oprrr(ASLL, p-&gt;scond);
		o3 = oprrr(ASRA, p-&gt;scond);
		r = p-&gt;to.reg;
		if(p-&gt;as == AMOVB) {
			o2 |= (24&lt;&lt;7)|(r)|(r&lt;&lt;12);
			o3 |= (24&lt;&lt;7)|(r)|(r&lt;&lt;12);
		} else {
			o2 |= (16&lt;&lt;7)|(r)|(r&lt;&lt;12);
			if(p-&gt;as == AMOVHU)
				o3 = oprrr(ASRL, p-&gt;scond);
			o3 |= (16&lt;&lt;7)|(r)|(r&lt;&lt;12);
		}
		break;

	case 23:	/* movh/movhu R,O(R) -&gt; sb,sb */
		aclass(&amp;p-&gt;to);
		r = p-&gt;to.reg;
		if(r == NREG)
			r = o-&gt;param;
		o1 = osr(AMOVH, p-&gt;from.reg, instoffset, r, p-&gt;scond);

		o2 = oprrr(ASRL, p-&gt;scond);
		o2 |= (8&lt;&lt;7)|(p-&gt;from.reg)|(REGTMP&lt;&lt;12);

		o3 = osr(AMOVH, REGTMP, instoffset+1, r, p-&gt;scond);
		break;

	case 30:	/* mov/movb/movbu R,L(R) */
		o1 = omvl(p, &amp;p-&gt;to, REGTMP);
		if(!o1)
			break;
		r = p-&gt;to.reg;
		if(r == NREG)
			r = o-&gt;param;
		o2 = osrr(p-&gt;from.reg, REGTMP,r, p-&gt;scond);
		if(p-&gt;as != AMOVW)
			o2 |= 1&lt;&lt;22;
		break;

	case 31:	/* mov/movbu L(R),R -&gt; lr[b] */
	case 32:	/* movh/movb L(R),R -&gt; lr[b] */
		o1 = omvl(p, &amp;p-&gt;from, REGTMP);
		if(!o1)
			break;
		r = p-&gt;from.reg;
		if(r == NREG)
			r = o-&gt;param;
		o2 = olrr(REGTMP,r, p-&gt;to.reg, p-&gt;scond);
		if(p-&gt;as == AMOVBU || p-&gt;as == AMOVB)
			o2 |= 1&lt;&lt;22;
		if(o-&gt;type == 31)
			break;

		o3 = oprrr(ASLL, p-&gt;scond);

		if(p-&gt;as == AMOVBU || p-&gt;as == AMOVHU)
			o4 = oprrr(ASRL, p-&gt;scond);
		else
			o4 = oprrr(ASRA, p-&gt;scond);

		r = p-&gt;to.reg;
		o3 |= (r)|(r&lt;&lt;12);
		o4 |= (r)|(r&lt;&lt;12);
		if(p-&gt;as == AMOVB || p-&gt;as == AMOVBU) {
			o3 |= (24&lt;&lt;7);
			o4 |= (24&lt;&lt;7);
		} else {
			o3 |= (16&lt;&lt;7);
			o4 |= (16&lt;&lt;7);
		}
		break;

	case 33:	/* movh/movhu R,L(R) -&gt; sb, sb */
		o1 = omvl(p, &amp;p-&gt;to, REGTMP);
		if(!o1)
			break;
		r = p-&gt;to.reg;
		if(r == NREG)
			r = o-&gt;param;
		o2 = osrr(p-&gt;from.reg, REGTMP, r, p-&gt;scond);
		o2 |= (1&lt;&lt;22) ;

		o3 = oprrr(ASRL, p-&gt;scond);
		o3 |= (8&lt;&lt;7)|(p-&gt;from.reg)|(p-&gt;from.reg&lt;&lt;12);
		o3 |= (1&lt;&lt;6);	/* ROR 8 */

		o4 = oprrr(AADD, p-&gt;scond);
		o4 |= (REGTMP &lt;&lt; 12) | (REGTMP &lt;&lt; 16);
		o4 |= immrot(1);

		o5 = osrr(p-&gt;from.reg, REGTMP,r,p-&gt;scond);
		o5 |= (1&lt;&lt;22);

		o6 = oprrr(ASRL, p-&gt;scond);
		o6 |= (24&lt;&lt;7)|(p-&gt;from.reg)|(p-&gt;from.reg&lt;&lt;12);
		o6 |= (1&lt;&lt;6);	/* ROL 8 */

		break;

	case 34:	/* mov $lacon,R */
		o1 = omvl(p, &amp;p-&gt;from, REGTMP);
		if(!o1)
			break;

		o2 = oprrr(AADD, p-&gt;scond);
		o2 |= REGTMP;
		r = p-&gt;from.reg;
		if(r == NREG)
			r = o-&gt;param;
		o2 |= r &lt;&lt; 16;
		if(p-&gt;to.type != D_NONE)
			o2 |= p-&gt;to.reg &lt;&lt; 12;
		break;

	case 35:	/* mov PSR,R */
		o1 = (2&lt;&lt;23) | (0xf&lt;&lt;16) | (0&lt;&lt;0);
		o1 |= (p-&gt;scond &amp; C_SCOND) &lt;&lt; 28;
		o1 |= (p-&gt;from.reg &amp; 1) &lt;&lt; 22;
		o1 |= p-&gt;to.reg &lt;&lt; 12;
		break;

	case 36:	/* mov R,PSR */
		o1 = (2&lt;&lt;23) | (0x29f&lt;&lt;12) | (0&lt;&lt;4);
		if(p-&gt;scond &amp; C_FBIT)
			o1 ^= 0x010 &lt;&lt; 12;
		o1 |= (p-&gt;scond &amp; C_SCOND) &lt;&lt; 28;
		o1 |= (p-&gt;to.reg &amp; 1) &lt;&lt; 22;
		o1 |= p-&gt;from.reg &lt;&lt; 0;
		break;

	case 37:	/* mov $con,PSR */
		aclass(&amp;p-&gt;from);
		o1 = (2&lt;&lt;23) | (0x29f&lt;&lt;12) | (0&lt;&lt;4);
		if(p-&gt;scond &amp; C_FBIT)
			o1 ^= 0x010 &lt;&lt; 12;
		o1 |= (p-&gt;scond &amp; C_SCOND) &lt;&lt; 28;
		o1 |= immrot(instoffset);
		o1 |= (p-&gt;to.reg &amp; 1) &lt;&lt; 22;
		o1 |= p-&gt;from.reg &lt;&lt; 0;
		break;

	case 38:	/* movm $con,oreg -&gt; stm */
		o1 = (0x4 &lt;&lt; 25);
		o1 |= p-&gt;from.offset &amp; 0xffff;
		o1 |= p-&gt;to.reg &lt;&lt; 16;
		aclass(&amp;p-&gt;to);
		goto movm;

	case 39:	/* movm oreg,$con -&gt; ldm */
		o1 = (0x4 &lt;&lt; 25) | (1 &lt;&lt; 20);
		o1 |= p-&gt;to.offset &amp; 0xffff;
		o1 |= p-&gt;from.reg &lt;&lt; 16;
		aclass(&amp;p-&gt;from);
	movm:
		if(instoffset != 0)
			diag(&#34;offset must be zero in MOVM&#34;);
		o1 |= (p-&gt;scond &amp; C_SCOND) &lt;&lt; 28;
		if(p-&gt;scond &amp; C_PBIT)
			o1 |= 1 &lt;&lt; 24;
		if(p-&gt;scond &amp; C_UBIT)
			o1 |= 1 &lt;&lt; 23;
		if(p-&gt;scond &amp; C_SBIT)
			o1 |= 1 &lt;&lt; 22;
		if(p-&gt;scond &amp; C_WBIT)
			o1 |= 1 &lt;&lt; 21;
		break;

	case 40:	/* swp oreg,reg,reg */
		aclass(&amp;p-&gt;from);
		if(instoffset != 0)
			diag(&#34;offset must be zero in SWP&#34;);
		o1 = (0x2&lt;&lt;23) | (0x9&lt;&lt;4);
		if(p-&gt;as != ASWPW)
			o1 |= 1 &lt;&lt; 22;
		o1 |= p-&gt;from.reg &lt;&lt; 16;
		o1 |= p-&gt;reg &lt;&lt; 0;
		o1 |= p-&gt;to.reg &lt;&lt; 12;
		o1 |= (p-&gt;scond &amp; C_SCOND) &lt;&lt; 28;
		break;

	case 41:	/* rfe -&gt; movm.s.w.u 0(r13),[r15] */
		o1 = 0xe8fd8000;
		break;

	case 50:	/* floating point store */
		v = regoff(&amp;p-&gt;to);
		r = p-&gt;to.reg;
		if(r == NREG)
			r = o-&gt;param;
		o1 = ofsr(p-&gt;as, p-&gt;from.reg, v, r, p-&gt;scond, p);
		break;

	case 51:	/* floating point load */
		v = regoff(&amp;p-&gt;from);
		r = p-&gt;from.reg;
		if(r == NREG)
			r = o-&gt;param;
		o1 = ofsr(p-&gt;as, p-&gt;to.reg, v, r, p-&gt;scond, p) | (1&lt;&lt;20);
		break;

	case 52:	/* floating point store, int32 offset UGLY */
		o1 = omvl(p, &amp;p-&gt;to, REGTMP);
		if(!o1)
			break;
		r = p-&gt;to.reg;
		if(r == NREG)
			r = o-&gt;param;
		o2 = oprrr(AADD, p-&gt;scond) | (REGTMP &lt;&lt; 12) | (REGTMP &lt;&lt; 16) | r;
		o3 = ofsr(p-&gt;as, p-&gt;from.reg, 0, REGTMP, p-&gt;scond, p);
		break;

	case 53:	/* floating point load, int32 offset UGLY */
		o1 = omvl(p, &amp;p-&gt;from, REGTMP);
		if(!o1)
			break;
		r = p-&gt;from.reg;
		if(r == NREG)
			r = o-&gt;param;
		o2 = oprrr(AADD, p-&gt;scond) | (REGTMP &lt;&lt; 12) | (REGTMP &lt;&lt; 16) | r;
		o3 = ofsr(p-&gt;as, p-&gt;to.reg, 0, REGTMP, p-&gt;scond, p) | (1&lt;&lt;20);
		break;

	case 54:	/* floating point arith */
		o1 = oprrr(p-&gt;as, p-&gt;scond);
		if(p-&gt;from.type == D_FCONST) {
			rf = chipfloat(p-&gt;from.ieee);
			if(rf &lt; 0){
				diag(&#34;invalid floating-point immediate\n%P&#34;, p);
				rf = 0;
			}
			rf |= (1&lt;&lt;3);
		} else
			rf = p-&gt;from.reg;
		rt = p-&gt;to.reg;
		r = p-&gt;reg;
		if(p-&gt;to.type == D_NONE)
			rt = 0;	/* CMP[FD] */
		else if(o1 &amp; (1&lt;&lt;15))
			r = 0;	/* monadic */
		else if(r == NREG)
			r = rt;
		o1 |= rf | (r&lt;&lt;16) | (rt&lt;&lt;12);
		break;

	case 55:	/* floating point fix and float */
		o1 = oprrr(p-&gt;as, p-&gt;scond);
		rf = p-&gt;from.reg;
		rt = p-&gt;to.reg;
		if(p-&gt;to.type == D_NONE){
			rt = 0;
			diag(&#34;to.type==D_NONE (asm/fp)&#34;);
		}
		if(p-&gt;from.type == D_REG)
			o1 |= (rf&lt;&lt;12) | (rt&lt;&lt;16);
		else
			o1 |= rf | (rt&lt;&lt;12);
		break;

	case 56:	/* move to FP[CS]R */
		o1 = ((p-&gt;scond &amp; C_SCOND) &lt;&lt; 28) | (0xe &lt;&lt; 24) | (1&lt;&lt;8) | (1&lt;&lt;4);
		o1 |= ((p-&gt;to.reg+1)&lt;&lt;21) | (p-&gt;from.reg &lt;&lt; 12);
		break;

	case 57:	/* move from FP[CS]R */
		o1 = ((p-&gt;scond &amp; C_SCOND) &lt;&lt; 28) | (0xe &lt;&lt; 24) | (1&lt;&lt;8) | (1&lt;&lt;4);
		o1 |= ((p-&gt;from.reg+1)&lt;&lt;21) | (p-&gt;to.reg&lt;&lt;12) | (1&lt;&lt;20);
		break;
	case 58:	/* movbu R,R */
		o1 = oprrr(AAND, p-&gt;scond);
		o1 |= immrot(0xff);
		rt = p-&gt;to.reg;
		r = p-&gt;from.reg;
		if(p-&gt;to.type == D_NONE)
			rt = 0;
		if(r == NREG)
			r = rt;
		o1 |= (r&lt;&lt;16) | (rt&lt;&lt;12);
		break;

	case 59:	/* movw/bu R&lt;&lt;I(R),R -&gt; ldr indexed */
		if(p-&gt;from.reg == NREG) {
			if(p-&gt;as != AMOVW)
				diag(&#34;byte MOV from shifter operand&#34;);
			goto mov;
		}
		if(p-&gt;from.offset&amp;(1&lt;&lt;4))
			diag(&#34;bad shift in LDR&#34;);
		o1 = olrr(p-&gt;from.offset, p-&gt;from.reg, p-&gt;to.reg, p-&gt;scond);
		if(p-&gt;as == AMOVBU)
			o1 |= 1&lt;&lt;22;
		break;

	case 60:	/* movb R(R),R -&gt; ldrsb indexed */
		if(p-&gt;from.reg == NREG) {
			diag(&#34;byte MOV from shifter operand&#34;);
			goto mov;
		}
		if(p-&gt;from.offset&amp;(~0xf))
			diag(&#34;bad shift in LDRSB&#34;);
		o1 = olhrr(p-&gt;from.offset, p-&gt;from.reg, p-&gt;to.reg, p-&gt;scond);
		o1 ^= (1&lt;&lt;5)|(1&lt;&lt;6);
		break;

	case 61:	/* movw/b/bu R,R&lt;&lt;[IR](R) -&gt; str indexed */
		if(p-&gt;to.reg == NREG)
			diag(&#34;MOV to shifter operand&#34;);
		o1 = osrr(p-&gt;from.reg, p-&gt;to.offset, p-&gt;to.reg, p-&gt;scond);
		if(p-&gt;as == AMOVB || p-&gt;as == AMOVBU)
			o1 |= 1&lt;&lt;22;
		break;

	case 62:	/* case R -&gt; movw	R&lt;&lt;2(PC),PC */
		o1 = olrr(p-&gt;from.reg, REGPC, REGPC, p-&gt;scond);
		o1 |= 2&lt;&lt;7;
		break;

	case 63:	/* bcase */
		if(p-&gt;cond != P) {
			o1 = p-&gt;cond-&gt;pc;
			if(dlm)
				dynreloc(S, p-&gt;pc, 1);
		}
		break;

	/* reloc ops */
	case 64:	/* mov/movb/movbu R,addr */
		o1 = omvl(p, &amp;p-&gt;to, REGTMP);
		if(!o1)
			break;
		o2 = osr(p-&gt;as, p-&gt;from.reg, 0, REGTMP, p-&gt;scond);
		break;

	case 65:	/* mov/movbu addr,R */
	case 66:	/* movh/movhu/movb addr,R */
		o1 = omvl(p, &amp;p-&gt;from, REGTMP);
		if(!o1)
			break;
		o2 = olr(0, REGTMP, p-&gt;to.reg, p-&gt;scond);
		if(p-&gt;as == AMOVBU || p-&gt;as == AMOVB)
			o2 |= 1&lt;&lt;22;
		if(o-&gt;type == 65)
			break;

		o3 = oprrr(ASLL, p-&gt;scond);

		if(p-&gt;as == AMOVBU || p-&gt;as == AMOVHU)
			o4 = oprrr(ASRL, p-&gt;scond);
		else
			o4 = oprrr(ASRA, p-&gt;scond);

		r = p-&gt;to.reg;
		o3 |= (r)|(r&lt;&lt;12);
		o4 |= (r)|(r&lt;&lt;12);
		if(p-&gt;as == AMOVB || p-&gt;as == AMOVBU) {
			o3 |= (24&lt;&lt;7);
			o4 |= (24&lt;&lt;7);
		} else {
			o3 |= (16&lt;&lt;7);
			o4 |= (16&lt;&lt;7);
		}
		break;

	case 67:	/* movh/movhu R,addr -&gt; sb, sb */
		o1 = omvl(p, &amp;p-&gt;to, REGTMP);
		if(!o1)
			break;
		o2 = osr(p-&gt;as, p-&gt;from.reg, 0, REGTMP, p-&gt;scond);

		o3 = oprrr(ASRL, p-&gt;scond);
		o3 |= (8&lt;&lt;7)|(p-&gt;from.reg)|(p-&gt;from.reg&lt;&lt;12);
		o3 |= (1&lt;&lt;6);	/* ROR 8 */

		o4 = oprrr(AADD, p-&gt;scond);
		o4 |= (REGTMP &lt;&lt; 12) | (REGTMP &lt;&lt; 16);
		o4 |= immrot(1);

		o5 = osr(p-&gt;as, p-&gt;from.reg, 0, REGTMP, p-&gt;scond);

		o6 = oprrr(ASRL, p-&gt;scond);
		o6 |= (24&lt;&lt;7)|(p-&gt;from.reg)|(p-&gt;from.reg&lt;&lt;12);
		o6 |= (1&lt;&lt;6);	/* ROL 8 */
		break;

	case 68:	/* floating point store -&gt; ADDR */
		o1 = omvl(p, &amp;p-&gt;to, REGTMP);
		if(!o1)
			break;
		o2 = ofsr(p-&gt;as, p-&gt;from.reg, 0, REGTMP, p-&gt;scond, p);
		break;

	case 69:	/* floating point load &lt;- ADDR */
		o1 = omvl(p, &amp;p-&gt;from, REGTMP);
		if(!o1)
			break;
		o2 = ofsr(p-&gt;as, p-&gt;to.reg, 0, REGTMP, p-&gt;scond, p) | (1&lt;&lt;20);
		break;

	/* ArmV4 ops: */
	case 70:	/* movh/movhu R,O(R) -&gt; strh */
		aclass(&amp;p-&gt;to);
		r = p-&gt;to.reg;
		if(r == NREG)
			r = o-&gt;param;
		o1 = oshr(p-&gt;from.reg, instoffset, r, p-&gt;scond);
		break;
	case 71:	/* movb/movh/movhu O(R),R -&gt; ldrsb/ldrsh/ldrh */
		aclass(&amp;p-&gt;from);
		r = p-&gt;from.reg;
		if(r == NREG)
			r = o-&gt;param;
		o1 = olhr(instoffset, r, p-&gt;to.reg, p-&gt;scond);
		if(p-&gt;as == AMOVB)
			o1 ^= (1&lt;&lt;5)|(1&lt;&lt;6);
		else if(p-&gt;as == AMOVH)
			o1 ^= (1&lt;&lt;6);
		break;
	case 72:	/* movh/movhu R,L(R) -&gt; strh */
		o1 = omvl(p, &amp;p-&gt;to, REGTMP);
		if(!o1)
			break;
		r = p-&gt;to.reg;
		if(r == NREG)
			r = o-&gt;param;
		o2 = oshrr(p-&gt;from.reg, REGTMP,r, p-&gt;scond);
		break;
	case 73:	/* movb/movh/movhu L(R),R -&gt; ldrsb/ldrsh/ldrh */
		o1 = omvl(p, &amp;p-&gt;from, REGTMP);
		if(!o1)
			break;
		r = p-&gt;from.reg;
		if(r == NREG)
			r = o-&gt;param;
		o2 = olhrr(REGTMP, r, p-&gt;to.reg, p-&gt;scond);
		if(p-&gt;as == AMOVB)
			o2 ^= (1&lt;&lt;5)|(1&lt;&lt;6);
		else if(p-&gt;as == AMOVH)
			o2 ^= (1&lt;&lt;6);
		break;
	case 74:	/* bx $I */
#ifdef CALLEEBX
		diag(&#34;bx $i case (arm)&#34;);
#endif
		if(!seenthumb)
			diag(&#34;ABX $I and seenthumb==0&#34;);
		v = p-&gt;cond-&gt;pc;
		if(p-&gt;to.sym-&gt;thumb)
			v |= 1;	// T bit
		o1 = olr(8, REGPC, REGTMP, p-&gt;scond&amp;C_SCOND);	// mov 8(PC), Rtmp
		o2 = 	oprrr(AADD, p-&gt;scond) | immrot(8) | (REGPC&lt;&lt;16) | (REGLINK&lt;&lt;12);	// add 8,PC, LR
		o3 = ((p-&gt;scond&amp;C_SCOND)&lt;&lt;28) | (0x12fff&lt;&lt;8) | (1&lt;&lt;4) | REGTMP;		// bx Rtmp
		o4 = opbra(AB, 14);	// B over o6
		o5 = v;
		break;
	case 75:	/* bx O(R) */
		aclass(&amp;p-&gt;to);
		if(instoffset != 0)
			diag(&#34;non-zero offset in ABX&#34;);
/*
		o1 = 	oprrr(AADD, p-&gt;scond) | immrot(0) | (REGPC&lt;&lt;16) | (REGLINK&lt;&lt;12);	// mov PC, LR
		o2 = ((p-&gt;scond&amp;C_SCOND)&lt;&lt;28) | (0x12fff&lt;&lt;8) | (1&lt;&lt;4) | p-&gt;to.reg;		// BX R
*/
		// p-&gt;to.reg may be REGLINK
		o1 = oprrr(AADD, p-&gt;scond);
		o1 |= immrot(instoffset);
		o1 |= p-&gt;to.reg &lt;&lt; 16;
		o1 |= REGTMP &lt;&lt; 12;
		o2 = 	oprrr(AADD, p-&gt;scond) | immrot(0) | (REGPC&lt;&lt;16) | (REGLINK&lt;&lt;12);	// mov PC, LR
		o3 = ((p-&gt;scond&amp;C_SCOND)&lt;&lt;28) | (0x12fff&lt;&lt;8) | (1&lt;&lt;4) | REGTMP;		// BX Rtmp
		break;
	case 76:	/* bx O(R) when returning from fn*/
		if(!seenthumb)
			diag(&#34;ABXRET and seenthumb==0&#34;);
		aclass(&amp;p-&gt;to);
// print(&#34;ARM BXRET %d(R%d)\n&#34;, instoffset, p-&gt;to.reg);
		if(instoffset != 0)
			diag(&#34;non-zero offset in ABXRET&#34;);
		// o1 = olr(instoffset, p-&gt;to.reg, REGTMP, p-&gt;scond);	// mov O(R), Rtmp
		o1 = ((p-&gt;scond&amp;C_SCOND)&lt;&lt;28) | (0x12fff&lt;&lt;8) | (1&lt;&lt;4) | p-&gt;to.reg;		// BX R
		break;
	case 77:	/* ldrex oreg,reg */
		aclass(&amp;p-&gt;from);
		if(instoffset != 0)
			diag(&#34;offset must be zero in LDREX&#34;);
		o1 = (0x19&lt;&lt;20) | (0xf9f);
		o1 |= p-&gt;from.reg &lt;&lt; 16;
		o1 |= p-&gt;to.reg &lt;&lt; 12;
		o1 |= (p-&gt;scond &amp; C_SCOND) &lt;&lt; 28;
		break;
	case 78:	/* strex reg,oreg,reg */
		aclass(&amp;p-&gt;from);
		if(instoffset != 0)
			diag(&#34;offset must be zero in STREX&#34;);
		o1 = (0x3&lt;&lt;23) | (0xf9&lt;&lt;4);
		o1 |= p-&gt;from.reg &lt;&lt; 16;
		o1 |= p-&gt;reg &lt;&lt; 0;
		o1 |= p-&gt;to.reg &lt;&lt; 12;
		o1 |= (p-&gt;scond &amp; C_SCOND) &lt;&lt; 28;
		break;
	}

	v = p-&gt;pc;
	switch(o-&gt;size) {
	default:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34; %.8lux:\t\t%P\n&#34;, v, p);
		break;
	case 4:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34; %.8lux: %.8lux\t%P\n&#34;, v, o1, p);
		lputl(o1);
		break;
	case 8:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34; %.8lux: %.8lux %.8lux%P\n&#34;, v, o1, o2, p);
		lputl(o1);
		lputl(o2);
		break;
	case 12:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34; %.8lux: %.8lux %.8lux %.8lux%P\n&#34;, v, o1, o2, o3, p);
		lputl(o1);
		lputl(o2);
		lputl(o3);
		break;
	case 16:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34; %.8lux: %.8lux %.8lux %.8lux %.8lux%P\n&#34;,
				v, o1, o2, o3, o4, p);
		lputl(o1);
		lputl(o2);
		lputl(o3);
		lputl(o4);
		break;
	case 20:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34; %.8lux: %.8lux %.8lux %.8lux %.8lux %.8lux%P\n&#34;,
				v, o1, o2, o3, o4, o5, p);
		lputl(o1);
		lputl(o2);
		lputl(o3);
		lputl(o4);
		lputl(o5);
		break;
	case 24:
		if(debug[&#39;a&#39;])
			Bprint(&amp;bso, &#34; %.8lux: %.8lux %.8lux %.8lux %.8lux %.8lux %.8lux%P\n&#34;,
				v, o1, o2, o3, o4, o5, o6, p);
		lputl(o1);
		lputl(o2);
		lputl(o3);
		lputl(o4);
		lputl(o5);
		lputl(o6);
		break;
	}
}

int32
oprrr(int a, int sc)
{
	int32 o;

	o = (sc &amp; C_SCOND) &lt;&lt; 28;
	if(sc &amp; C_SBIT)
		o |= 1 &lt;&lt; 20;
	if(sc &amp; (C_PBIT|C_WBIT))
		diag(&#34;.P/.W on dp instruction&#34;);
	switch(a) {
	case AMULU:
	case AMUL:	return o | (0x0&lt;&lt;21) | (0x9&lt;&lt;4);
	case AMULA:	return o | (0x1&lt;&lt;21) | (0x9&lt;&lt;4);
	case AMULLU:	return o | (0x4&lt;&lt;21) | (0x9&lt;&lt;4);
	case AMULL:	return o | (0x6&lt;&lt;21) | (0x9&lt;&lt;4);
	case AMULALU:	return o | (0x5&lt;&lt;21) | (0x9&lt;&lt;4);
	case AMULAL:	return o | (0x7&lt;&lt;21) | (0x9&lt;&lt;4);
	case AAND:	return o | (0x0&lt;&lt;21);
	case AEOR:	return o | (0x1&lt;&lt;21);
	case ASUB:	return o | (0x2&lt;&lt;21);
	case ARSB:	return o | (0x3&lt;&lt;21);
	case AADD:	return o | (0x4&lt;&lt;21);
	case AADC:	return o | (0x5&lt;&lt;21);
	case ASBC:	return o | (0x6&lt;&lt;21);
	case ARSC:	return o | (0x7&lt;&lt;21);
	case ATST:	return o | (0x8&lt;&lt;21) | (1&lt;&lt;20);
	case ATEQ:	return o | (0x9&lt;&lt;21) | (1&lt;&lt;20);
	case ACMP:	return o | (0xa&lt;&lt;21) | (1&lt;&lt;20);
	case ACMN:	return o | (0xb&lt;&lt;21) | (1&lt;&lt;20);
	case AORR:	return o | (0xc&lt;&lt;21);
	case AMOVW:	return o | (0xd&lt;&lt;21);
	case ABIC:	return o | (0xe&lt;&lt;21);
	case AMVN:	return o | (0xf&lt;&lt;21);
	case ASLL:	return o | (0xd&lt;&lt;21) | (0&lt;&lt;5);
	case ASRL:	return o | (0xd&lt;&lt;21) | (1&lt;&lt;5);
	case ASRA:	return o | (0xd&lt;&lt;21) | (2&lt;&lt;5);
	case ASWI:	return o | (0xf&lt;&lt;24);

	case AADDD:	return o | (0xe&lt;&lt;24) | (0x0&lt;&lt;20) | (1&lt;&lt;8) | (1&lt;&lt;7);
	case AADDF:	return o | (0xe&lt;&lt;24) | (0x0&lt;&lt;20) | (1&lt;&lt;8);
	case AMULD:	return o | (0xe&lt;&lt;24) | (0x1&lt;&lt;20) | (1&lt;&lt;8) | (1&lt;&lt;7);
	case AMULF:	return o | (0xe&lt;&lt;24) | (0x1&lt;&lt;20) | (1&lt;&lt;8);
	case ASUBD:	return o | (0xe&lt;&lt;24) | (0x2&lt;&lt;20) | (1&lt;&lt;8) | (1&lt;&lt;7);
	case ASUBF:	return o | (0xe&lt;&lt;24) | (0x2&lt;&lt;20) | (1&lt;&lt;8);
	case ADIVD:	return o | (0xe&lt;&lt;24) | (0x4&lt;&lt;20) | (1&lt;&lt;8) | (1&lt;&lt;7);
	case ADIVF:	return o | (0xe&lt;&lt;24) | (0x4&lt;&lt;20) | (1&lt;&lt;8);
	case ACMPD:
	case ACMPF:	return o | (0xe&lt;&lt;24) | (0x9&lt;&lt;20) | (0xF&lt;&lt;12) | (1&lt;&lt;8) | (1&lt;&lt;4);	/* arguably, ACMPF should expand to RNDF, CMPD */

	case AMOVF:
	case AMOVDF:	return o | (0xe&lt;&lt;24) | (0x0&lt;&lt;20) | (1&lt;&lt;15) | (1&lt;&lt;8);
	case AMOVD:
	case AMOVFD:	return o | (0xe&lt;&lt;24) | (0x0&lt;&lt;20) | (1&lt;&lt;15) | (1&lt;&lt;8) | (1&lt;&lt;7);

	case AMOVWF:	return o | (0xe&lt;&lt;24) | (0&lt;&lt;20) | (1&lt;&lt;8) | (1&lt;&lt;4);
	case AMOVWD:	return o | (0xe&lt;&lt;24) | (0&lt;&lt;20) | (1&lt;&lt;8) | (1&lt;&lt;4) | (1&lt;&lt;7);
	case AMOVFW:	return o | (0xe&lt;&lt;24) | (1&lt;&lt;20) | (1&lt;&lt;8) | (1&lt;&lt;4);
	case AMOVDW:	return o | (0xe&lt;&lt;24) | (1&lt;&lt;20) | (1&lt;&lt;8) | (1&lt;&lt;4) | (1&lt;&lt;7);
	}
	diag(&#34;bad rrr %d&#34;, a);
	prasm(curp);
	return 0;
}

int32
opbra(int a, int sc)
{

	if(sc &amp; (C_SBIT|C_PBIT|C_WBIT))
		diag(&#34;.S/.P/.W on bra instruction&#34;);
	sc &amp;= C_SCOND;
	if(a == ABL)
		return (sc&lt;&lt;28)|(0x5&lt;&lt;25)|(0x1&lt;&lt;24);
	if(sc != 0xe)
		diag(&#34;.COND on bcond instruction&#34;);
	switch(a) {
	case ABEQ:	return (0x0&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABNE:	return (0x1&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABCS:	return (0x2&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABHS:	return (0x2&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABCC:	return (0x3&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABLO:	return (0x3&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABMI:	return (0x4&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABPL:	return (0x5&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABVS:	return (0x6&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABVC:	return (0x7&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABHI:	return (0x8&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABLS:	return (0x9&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABGE:	return (0xa&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABLT:	return (0xb&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABGT:	return (0xc&lt;&lt;28)|(0x5&lt;&lt;25);
	case ABLE:	return (0xd&lt;&lt;28)|(0x5&lt;&lt;25);
	case AB:	return (0xe&lt;&lt;28)|(0x5&lt;&lt;25);
	}
	diag(&#34;bad bra %A&#34;, a);
	prasm(curp);
	return 0;
}

int32
olr(int32 v, int b, int r, int sc)
{
	int32 o;

	if(sc &amp; C_SBIT)
		diag(&#34;.S on LDR/STR instruction&#34;);
	o = (sc &amp; C_SCOND) &lt;&lt; 28;
	if(!(sc &amp; C_PBIT))
		o |= 1 &lt;&lt; 24;
	if(!(sc &amp; C_UBIT))
		o |= 1 &lt;&lt; 23;
	if(sc &amp; C_WBIT)
		o |= 1 &lt;&lt; 21;
	o |= (0x1&lt;&lt;26) | (1&lt;&lt;20);
	if(v &lt; 0) {
		v = -v;
		o ^= 1 &lt;&lt; 23;
	}
	if(v &gt;= (1&lt;&lt;12))
		diag(&#34;literal span too large: %d (R%d)\n%P&#34;, v, b, PP);
	o |= v;
	o |= b &lt;&lt; 16;
	o |= r &lt;&lt; 12;
	return o;
}

int32
olhr(int32 v, int b, int r, int sc)
{
	int32 o;

	if(sc &amp; C_SBIT)
		diag(&#34;.S on LDRH/STRH instruction&#34;);
	o = (sc &amp; C_SCOND) &lt;&lt; 28;
	if(!(sc &amp; C_PBIT))
		o |= 1 &lt;&lt; 24;
	if(sc &amp; C_WBIT)
		o |= 1 &lt;&lt; 21;
	o |= (1&lt;&lt;23) | (1&lt;&lt;20)|(0xb&lt;&lt;4);
	if(v &lt; 0) {
		v = -v;
		o ^= 1 &lt;&lt; 23;
	}
	if(v &gt;= (1&lt;&lt;8))
		diag(&#34;literal span too large: %d (R%d)\n%P&#34;, v, b, PP);
	o |= (v&amp;0xf)|((v&gt;&gt;4)&lt;&lt;8)|(1&lt;&lt;22);
	o |= b &lt;&lt; 16;
	o |= r &lt;&lt; 12;
	return o;
}

int32
osr(int a, int r, int32 v, int b, int sc)
{
	int32 o;

	o = olr(v, b, r, sc) ^ (1&lt;&lt;20);
	if(a != AMOVW)
		o |= 1&lt;&lt;22;
	return o;
}

int32
oshr(int r, int32 v, int b, int sc)
{
	int32 o;

	o = olhr(v, b, r, sc) ^ (1&lt;&lt;20);
	return o;
}


int32
osrr(int r, int i, int b, int sc)
{

	return olr(i, b, r, sc) ^ ((1&lt;&lt;25) | (1&lt;&lt;20));
}

int32
oshrr(int r, int i, int b, int sc)
{
	return olhr(i, b, r, sc) ^ ((1&lt;&lt;22) | (1&lt;&lt;20));
}

int32
olrr(int i, int b, int r, int sc)
{

	return olr(i, b, r, sc) ^ (1&lt;&lt;25);
}

int32
olhrr(int i, int b, int r, int sc)
{
	return olhr(i, b, r, sc) ^ (1&lt;&lt;22);
}

int32
ofsr(int a, int r, int32 v, int b, int sc, Prog *p)
{
	int32 o;

	if(sc &amp; C_SBIT)
		diag(&#34;.S on FLDR/FSTR instruction&#34;);
	o = (sc &amp; C_SCOND) &lt;&lt; 28;
	if(!(sc &amp; C_PBIT))
		o |= 1 &lt;&lt; 24;
	if(sc &amp; C_WBIT)
		o |= 1 &lt;&lt; 21;
	o |= (6&lt;&lt;25) | (1&lt;&lt;24) | (1&lt;&lt;23);
	if(v &lt; 0) {
		v = -v;
		o ^= 1 &lt;&lt; 23;
	}
	if(v &amp; 3)
		diag(&#34;odd offset for floating point op: %d\n%P&#34;, v, p);
	else if(v &gt;= (1&lt;&lt;10))
		diag(&#34;literal span too large: %d\n%P&#34;, v, p);
	o |= (v&gt;&gt;2) &amp; 0xFF;
	o |= b &lt;&lt; 16;
	o |= r &lt;&lt; 12;
	o |= 1 &lt;&lt; 8;

	switch(a) {
	default:
		diag(&#34;bad fst %A&#34;, a);
	case AMOVD:
		o |= 1&lt;&lt;15;
	case AMOVF:
		break;
	}
	return o;
}

int32
omvl(Prog *p, Adr *a, int dr)
{
	int32 v, o1;
	if(!p-&gt;cond) {
		aclass(a);
		v = immrot(~instoffset);
		if(v == 0) {
			diag(&#34;missing literal&#34;);
			prasm(p);
			return 0;
		}
		o1 = oprrr(AMVN, p-&gt;scond&amp;C_SCOND);
		o1 |= v;
		o1 |= dr &lt;&lt; 12;
	} else {
		v = p-&gt;cond-&gt;pc - p-&gt;pc - 8;
		o1 = olr(v, REGPC, dr, p-&gt;scond&amp;C_SCOND);
	}
	return o1;
}

static Ieee chipfloats[] = {
	{0x00000000, 0x00000000}, /* 0 */
	{0x00000000, 0x3ff00000}, /* 1 */
	{0x00000000, 0x40000000}, /* 2 */
	{0x00000000, 0x40080000}, /* 3 */
	{0x00000000, 0x40100000}, /* 4 */
	{0x00000000, 0x40140000}, /* 5 */
	{0x00000000, 0x3fe00000}, /* .5 */
	{0x00000000, 0x40240000}, /* 10 */
};

int
chipfloat(Ieee *e)
{
	Ieee *p;
	int n;

	for(n = sizeof(chipfloats)/sizeof(chipfloats[0]); --n &gt;= 0;){
		p = &amp;chipfloats[n];
		if(p-&gt;l == e-&gt;l &amp;&amp; p-&gt;h == e-&gt;h)
			return n;
	}
	return -1;
}

uint32
linuxheadr(void)
{
	uint32 a;

	a = 64;		/* a.out header */

	a += 56;	/* page zero seg */
	a += 56;	/* text seg */
	a += 56;	/* stack seg */

	a += 64;	/* nil sect */
	a += 64;	/* .text sect */
	a += 64;	/* .data seg */
	a += 64;	/* .bss sect */
	a += 64;	/* .shstrtab sect - strings for headers */
	if (!debug[&#39;s&#39;]) {
		a += 56;	/* symdat seg */
		a += 64;	/* .gosymtab sect */
		a += 64;	/* .gopclntab sect */
	}

	return a;
}

void
linuxphdr(int type, int flags, vlong foff,
	vlong vaddr, vlong paddr,
	vlong filesize, vlong memsize, vlong align)
{

	lputl(type);			/* text - type = PT_LOAD */
	lputl(foff);			/* file offset */
	lputl(vaddr);			/* vaddr */
	lputl(paddr);			/* paddr */
	lputl(filesize);		/* file size */
	lputl(memsize);		/* memory size */
	lputl(flags);			/* text - flags = PF_X+PF_R */
	lputl(align);			/* alignment */
}

void
linuxshdr(char *name, uint32 type, vlong flags, vlong addr, vlong off,
	vlong size, uint32 link, uint32 info, vlong align, vlong entsize)
{
	lputl(stroffset);
	lputl(type);
	lputl(flags);
	lputl(addr);
	lputl(off);
	lputl(size);
	lputl(link);
	lputl(info);
	lputl(align);
	lputl(entsize);

	if(name != nil)
		stroffset += strlen(name)+1;
}

int
putstrtab(char* name)
{
	int w;

	w = strlen(name)+1;
	strnput(name, w);
	return w;
}

int
linuxstrtable(void)
{
	int size;

	size = 0;
	size += putstrtab(&#34;&#34;);
	size += putstrtab(&#34;.text&#34;);
	size += putstrtab(&#34;.data&#34;);
	size += putstrtab(&#34;.bss&#34;);
	size += putstrtab(&#34;.shstrtab&#34;);
	if (!debug[&#39;s&#39;]) {
		size += putstrtab(&#34;.gosymtab&#34;);
		size += putstrtab(&#34;.gopclntab&#34;);
	}
	return size;
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
