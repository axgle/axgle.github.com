<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8l/asm.c</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/8l/asm.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/8l/asm.c
// http://code.google.com/p/inferno-os/source/browse/utils/8l/asm.c
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
#include	&#34;../ld/elf.h&#34;
#include	&#34;../ld/macho.h&#34;

#define	Dbufslop	100

char linuxdynld[] = &#34;/lib/ld-linux.so.2&#34;;
uint32 symdatva = 0x99&lt;&lt;24;

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
wputl(ushort w)
{
	cput(w);
	cput(w&gt;&gt;8);
}

void
wput(ushort w)
{
	cput(w&gt;&gt;8);
	cput(w);
}

void
lput(int32 l)
{
	cput(l&gt;&gt;24);
	cput(l&gt;&gt;16);
	cput(l&gt;&gt;8);
	cput(l);
}

void
lputl(int32 l)
{
	cput(l);
	cput(l&gt;&gt;8);
	cput(l&gt;&gt;16);
	cput(l&gt;&gt;24);
}

void
vputl(uvlong l)
{
	lputl(l &gt;&gt; 32);
	lputl(l);
}

void
strnput(char *s, int n)
{
	for(; *s &amp;&amp; n &gt; 0; s++) {
		cput(*s);
		n--;
	}
	while(n &gt; 0) {
		cput(0);
		n--;
	}
}

vlong
addstring(Sym *s, char *str)
{
	int n, m;
	vlong r;
	Prog *p;

	if(s-&gt;type == 0)
		s-&gt;type = SDATA;
	s-&gt;reachable = 1;
	r = s-&gt;value;
	n = strlen(str)+1;
	while(n &gt; 0) {
		m = n;
		if(m &gt; sizeof(p-&gt;to.scon))
			m = sizeof(p-&gt;to.scon);
		p = newdata(s, s-&gt;value, m, D_EXTERN);
		p-&gt;to.type = D_SCONST;
		memmove(p-&gt;to.scon, str, m);
		s-&gt;value += m;
		str += m;
		n -= m;
	}
	return r;
}

vlong
adduintxx(Sym *s, uint64 v, int wid)
{
	vlong r;
	Prog *p;

	if(s-&gt;type == 0)
		s-&gt;type = SDATA;
	s-&gt;reachable = 1;
	r = s-&gt;value;
	p = newdata(s, s-&gt;value, wid, D_EXTERN);
	s-&gt;value += wid;
	p-&gt;to.type = D_CONST;
	p-&gt;to.offset = v;
	return r;
}

vlong
adduint8(Sym *s, uint8 v)
{
	return adduintxx(s, v, 1);
}

vlong
adduint16(Sym *s, uint16 v)
{
	return adduintxx(s, v, 2);
}

vlong
adduint32(Sym *s, uint32 v)
{
	return adduintxx(s, v, 4);
}

vlong
adduint64(Sym *s, uint64 v)
{
	return adduintxx(s, v, 8);
}

vlong
addaddr(Sym *s, Sym *t)
{
	vlong r;
	Prog *p;
	enum { Ptrsize = 4 };

	if(s-&gt;type == 0)
		s-&gt;type = SDATA;
	s-&gt;reachable = 1;
	r = s-&gt;value;
	p = newdata(s, s-&gt;value, Ptrsize, D_EXTERN);
	s-&gt;value += Ptrsize;
	p-&gt;to.type = D_ADDR;
	p-&gt;to.index = D_EXTERN;
	p-&gt;to.offset = 0;
	p-&gt;to.sym = t;
	return r;
}

vlong
addsize(Sym *s, Sym *t)
{
	vlong r;
	Prog *p;
	enum { Ptrsize = 4 };

	if(s-&gt;type == 0)
		s-&gt;type = SDATA;
	s-&gt;reachable = 1;
	r = s-&gt;value;
	p = newdata(s, s-&gt;value, Ptrsize, D_EXTERN);
	s-&gt;value += Ptrsize;
	p-&gt;to.type = D_SIZE;
	p-&gt;to.index = D_EXTERN;
	p-&gt;to.offset = 0;
	p-&gt;to.sym = t;
	return r;
}

vlong
datoff(vlong addr)
{
	if(addr &gt;= INITDAT)
		return addr - INITDAT + rnd(HEADR+textsize, INITRND);
	diag(&#34;datoff %#llx&#34;, addr);
	return 0;
}

enum {
	ElfStrEmpty,
	ElfStrInterp,
	ElfStrHash,
	ElfStrGot,
	ElfStrGotPlt,
	ElfStrDynamic,
	ElfStrDynsym,
	ElfStrDynstr,
	ElfStrRel,
	ElfStrText,
	ElfStrData,
	ElfStrBss,
	ElfStrGosymtab,
	ElfStrGopclntab,
	ElfStrShstrtab,
	NElfStr
};

vlong elfstr[NElfStr];

static int
needlib(char *name)
{
	char *p;
	Sym *s;

	/* reuse hash code in symbol table */
	p = smprint(&#34;.dynlib.%s&#34;, name);
	s = lookup(p, 0);
	if(s-&gt;type == 0) {
		s-&gt;type = 100;	// avoid SDATA, etc.
		return 1;
	}
	return 0;
}

void
doelf(void)
{
	Sym *s, *shstrtab, *dynamic, *dynstr, *d;
	int h, nsym, t;

	if(HEADTYPE != 7 &amp;&amp; HEADTYPE != 8)
		return;

	/* predefine strings we need for section headers */
	shstrtab = lookup(&#34;.shstrtab&#34;, 0);
	shstrtab-&gt;reachable = 1;
	elfstr[ElfStrEmpty] = addstring(shstrtab, &#34;&#34;);
	elfstr[ElfStrText] = addstring(shstrtab, &#34;.text&#34;);
	elfstr[ElfStrData] = addstring(shstrtab, &#34;.data&#34;);
	elfstr[ElfStrBss] = addstring(shstrtab, &#34;.bss&#34;);
	if(!debug[&#39;s&#39;]) {
		elfstr[ElfStrGosymtab] = addstring(shstrtab, &#34;.gosymtab&#34;);
		elfstr[ElfStrGopclntab] = addstring(shstrtab, &#34;.gopclntab&#34;);
	}
	elfstr[ElfStrShstrtab] = addstring(shstrtab, &#34;.shstrtab&#34;);

	if(!debug[&#39;d&#39;]) {	/* -d suppresses dynamic loader format */
		elfstr[ElfStrInterp] = addstring(shstrtab, &#34;.interp&#34;);
		elfstr[ElfStrHash] = addstring(shstrtab, &#34;.hash&#34;);
		elfstr[ElfStrGot] = addstring(shstrtab, &#34;.got&#34;);
		elfstr[ElfStrGotPlt] = addstring(shstrtab, &#34;.got.plt&#34;);
		elfstr[ElfStrDynamic] = addstring(shstrtab, &#34;.dynamic&#34;);
		elfstr[ElfStrDynsym] = addstring(shstrtab, &#34;.dynsym&#34;);
		elfstr[ElfStrDynstr] = addstring(shstrtab, &#34;.dynstr&#34;);
		elfstr[ElfStrRel] = addstring(shstrtab, &#34;.rel&#34;);

		/* interpreter string */
		s = lookup(&#34;.interp&#34;, 0);
		s-&gt;reachable = 1;
		s-&gt;type = SDATA;	// TODO: rodata
		addstring(lookup(&#34;.interp&#34;, 0), linuxdynld);

		/*
		 * hash table - empty for now.
		 * only entries that other objects need to find when
		 * linking us need to be in this table.  right now that
		 * is no entries.
		 *
		 * must have at least 1 bucket, though, to avoid
		 * a divide by zero bug in some copies of the
		 * glibc dynamic loader.
		 */
		s = lookup(&#34;.hash&#34;, 0);
		s-&gt;type = SDATA;	// TODO: rodata
		s-&gt;reachable = 1;
		adduint32(s, 1);	// nbucket
		adduint32(s, 1);	// nchain
		adduint32(s, 0);	// bucket[0]
		adduint32(s, 0);	// chain[0]

		/* dynamic symbol table - first entry all zeros */
		s = lookup(&#34;.dynsym&#34;, 0);
		s-&gt;type = SDATA;
		s-&gt;reachable = 1;
		s-&gt;value += ELF32SYMSIZE;

		/* dynamic string table */
		s = lookup(&#34;.dynstr&#34;, 0);
		addstring(s, &#34;&#34;);
		dynstr = s;

		/* relocation table */
		s = lookup(&#34;.rel&#34;, 0);
		s-&gt;reachable = 1;
		s-&gt;type = SDATA;

		/* global offset table */
		s = lookup(&#34;.got&#34;, 0);
		s-&gt;reachable = 1;
		s-&gt;type = SDATA;

		/* got.plt - ??? */
		s = lookup(&#34;.got.plt&#34;, 0);
		s-&gt;reachable = 1;
		s-&gt;type = SDATA;

		/* define dynamic elf table */
		s = lookup(&#34;.dynamic&#34;, 0);
		dynamic = s;

		/*
		 * relocation entries for dynld symbols
		 */
		nsym = 1;	// sym 0 is reserved
		for(h=0; h&lt;NHASH; h++) {
			for(s=hash[h]; s!=S; s=s-&gt;link) {
				if(!s-&gt;reachable || (s-&gt;type != SDATA &amp;&amp; s-&gt;type != SBSS) || s-&gt;dynldname == nil)
					continue;

				d = lookup(&#34;.rel&#34;, 0);
				addaddr(d, s);
				adduint32(d, ELF32_R_INFO(nsym, R_386_32));
				nsym++;

				d = lookup(&#34;.dynsym&#34;, 0);
				adduint32(d, addstring(lookup(&#34;.dynstr&#34;, 0), s-&gt;dynldname));
				adduint32(d, 0);	/* value */
				adduint32(d, 0);	/* size of object */
				t = STB_GLOBAL &lt;&lt; 4;
				t |= STT_OBJECT;	// works for func too, empirically
				adduint8(d, t);
				adduint8(d, 0);	/* reserved */
				adduint16(d, SHN_UNDEF);	/* section where symbol is defined */

				if(needlib(s-&gt;dynldlib))
					elfwritedynent(dynamic, DT_NEEDED, addstring(dynstr, s-&gt;dynldlib));
			}
		}

		/*
		 * .dynamic table
		 */
		s = dynamic;
		elfwritedynentsym(s, DT_HASH, lookup(&#34;.hash&#34;, 0));
		elfwritedynentsym(s, DT_SYMTAB, lookup(&#34;.dynsym&#34;, 0));
		elfwritedynent(s, DT_SYMENT, ELF32SYMSIZE);
		elfwritedynentsym(s, DT_STRTAB, lookup(&#34;.dynstr&#34;, 0));
		elfwritedynentsymsize(s, DT_STRSZ, lookup(&#34;.dynstr&#34;, 0));
		elfwritedynentsym(s, DT_REL, lookup(&#34;.rel&#34;, 0));
		elfwritedynentsymsize(s, DT_RELSZ, lookup(&#34;.rel&#34;, 0));
		elfwritedynent(s, DT_RELENT, ELF32RELSIZE);
		elfwritedynent(s, DT_NULL, 0);
	}
}

void
shsym(Elf64_Shdr *sh, Sym *s)
{
	sh-&gt;addr = symaddr(s);
	sh-&gt;off = datoff(sh-&gt;addr);
	sh-&gt;size = s-&gt;size;
}

void
phsh(Elf64_Phdr *ph, Elf64_Shdr *sh)
{
	ph-&gt;vaddr = sh-&gt;addr;
	ph-&gt;paddr = ph-&gt;vaddr;
	ph-&gt;off = sh-&gt;off;
	ph-&gt;filesz = sh-&gt;size;
	ph-&gt;memsz = sh-&gt;size;
	ph-&gt;align = sh-&gt;addralign;
}

void
asmb(void)
{
	Prog *p;
	int32 v, magic;
	int a, dynsym;
	uint32 va, fo, w, symo, startva, machlink;
	uchar *op1;
	ulong expectpc;
	ElfEhdr *eh;
	ElfPhdr *ph, *pph;
	ElfShdr *sh;

	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f asmb\n&#34;, cputime());
	Bflush(&amp;bso);

	seek(cout, HEADR, 0);
	pc = INITTEXT;
	curp = firstp;
	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT)
			curtext = p;
		curp = p;
		if(HEADTYPE == 8) {
			// native client
			expectpc = p-&gt;pc;
			p-&gt;pc = pc;
			asmins(p);
			if(p-&gt;pc != expectpc) {
				Bflush(&amp;bso);
				diag(&#34;phase error %lux sb %lux in %s&#34;, p-&gt;pc, expectpc, TNAME);
			}
			while(pc &lt; p-&gt;pc) {
				cput(0x90);	// nop
				pc++;
			}
		}
		if(p-&gt;pc != pc) {
			Bflush(&amp;bso);
			if(!debug[&#39;a&#39;])
				print(&#34;%P\n&#34;, curp);
			diag(&#34;phase error %lux sb %lux in %s&#34;, p-&gt;pc, pc, TNAME);
			pc = p-&gt;pc;
		}
		if(HEADTYPE != 8) {
			asmins(p);
			if(pc != p-&gt;pc) {
				Bflush(&amp;bso);
				diag(&#34;asmins changed pc %lux sb %lux in %s&#34;, p-&gt;pc, pc, TNAME);
			}
		}
		if(cbc &lt; sizeof(and))
			cflush();
		a = (andptr - and);

		if(debug[&#39;a&#39;]) {
			Bprint(&amp;bso, pcstr, pc);
			for(op1 = and; op1 &lt; andptr; op1++)
				Bprint(&amp;bso, &#34;%.2ux&#34;, *op1 &amp; 0xff);
			Bprint(&amp;bso, &#34;\t%P\n&#34;, curp);
		}
		if(dlm) {
			if(p-&gt;as == ATEXT)
				reloca = nil;
			else if(reloca != nil)
				diag(&#34;reloc failure: %P&#34;, curp);
		}
		memmove(cbp, and, a);
		cbp += a;
		pc += a;
		cbc -= a;
	}
	if(HEADTYPE == 8) {
		while(pc &lt; INITDAT) {
			cput(0xf4);	// hlt
			pc++;
		}
	}
	cflush();

	switch(HEADTYPE) {
	default:
		diag(&#34;unknown header type %d&#34;, HEADTYPE);
	case 0:
		seek(cout, rnd(HEADR+textsize, 8192), 0);
		break;
	case 1:
		textsize = rnd(HEADR+textsize, 4096)-HEADR;
		seek(cout, textsize+HEADR, 0);
		break;
	case 2:
		seek(cout, HEADR+textsize, 0);
		break;
	case 3:
	case 4:
		seek(cout, HEADR+rnd(textsize, INITRND), 0);
		break;
	case 6:
		v = HEADR+textsize;
		seek(cout, v, 0);
		v = rnd(v, 4096) - v;
		while(v &gt; 0) {
			cput(0);
			v--;
		}
		cflush();
		break;
	case 7:
	case 8:
		v = rnd(HEADR+textsize, INITRND);
		seek(cout, v, 0);
		break;
	}

	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f datblk\n&#34;, cputime());
	Bflush(&amp;bso);

	if(dlm){
		char buf[8];

		write(cout, buf, INITDAT-textsize);
		textsize = INITDAT;
	}

	for(v = 0; v &lt; datsize; v += sizeof(buf)-Dbufslop) {
		if(datsize-v &gt; sizeof(buf)-Dbufslop)
			datblk(v, sizeof(buf)-Dbufslop);
		else
			datblk(v, datsize-v);
	}

	machlink = 0;
	if(HEADTYPE == 6)
		machlink = domacholink();

	symsize = 0;
	spsize = 0;
	lcsize = 0;
	symo = 0;
	if(!debug[&#39;s&#39;]) {
		if(debug[&#39;v&#39;])
			Bprint(&amp;bso, &#34;%5.2f sym\n&#34;, cputime());
		Bflush(&amp;bso);
		switch(HEADTYPE) {
		default:
		case 0:
			seek(cout, rnd(HEADR+textsize, 8192)+datsize, 0);
			break;
		case 1:
			seek(cout, rnd(HEADR+textsize, INITRND)+datsize, 0);
			break;
		case 2:
			seek(cout, HEADR+textsize+datsize, 0);
			break;
		case 3:
		case 4:
			debug[&#39;s&#39;] = 1;
			symo = HEADR+textsize+datsize;
			break;
		case 6:
			symo = rnd(HEADR+textsize, INITRND)+rnd(datsize, INITRND)+machlink;
			break;
		case 7:
		case 8:
			symo = rnd(HEADR+textsize, INITRND)+datsize;
			symo = rnd(symo, INITRND);
			break;
		}
		seek(cout, symo+8, 0);
		if(!debug[&#39;s&#39;])
			asmsym();
		if(debug[&#39;v&#39;])
			Bprint(&amp;bso, &#34;%5.2f sp\n&#34;, cputime());
		Bflush(&amp;bso);
		if(debug[&#39;v&#39;])
			Bprint(&amp;bso, &#34;%5.2f pc\n&#34;, cputime());
		Bflush(&amp;bso);
		if(!debug[&#39;s&#39;])
			asmlc();
		if(dlm)
			asmdyn();
		cflush();
		seek(cout, symo, 0);
		lputl(symsize);
		lputl(lcsize);
		cflush();
	}
	else if(dlm){
		seek(cout, HEADR+textsize+datsize, 0);
		asmdyn();
		cflush();
	}
	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f headr\n&#34;, cputime());
	Bflush(&amp;bso);
	seek(cout, 0L, 0);
	switch(HEADTYPE) {
	default:
	case 0:	/* garbage */
		lput(0x160L&lt;&lt;16);		/* magic and sections */
		lput(0L);			/* time and date */
		lput(rnd(HEADR+textsize, 4096)+datsize);
		lput(symsize);			/* nsyms */
		lput((0x38L&lt;&lt;16)|7L);		/* size of optional hdr and flags */
		lput((0413&lt;&lt;16)|0437L);		/* magic and version */
		lput(rnd(HEADR+textsize, 4096));	/* sizes */
		lput(datsize);
		lput(bsssize);
		lput(entryvalue());		/* va of entry */
		lput(INITTEXT-HEADR);		/* va of base of text */
		lput(INITDAT);			/* va of base of data */
		lput(INITDAT+datsize);		/* va of base of bss */
		lput(~0L);			/* gp reg mask */
		lput(0L);
		lput(0L);
		lput(0L);
		lput(0L);
		lput(~0L);			/* gp value ?? */
		break;
		lputl(0);			/* x */
	case 1:	/* unix coff */
		/*
		 * file header
		 */
		lputl(0x0004014c);		/* 4 sections, magic */
		lputl(0);			/* unix time stamp */
		lputl(0);			/* symbol table */
		lputl(0);			/* nsyms */
		lputl(0x0003001c);		/* flags, sizeof a.out header */
		/*
		 * a.out header
		 */
		lputl(0x10b);			/* magic, version stamp */
		lputl(rnd(textsize, INITRND));	/* text sizes */
		lputl(datsize);			/* data sizes */
		lputl(bsssize);			/* bss sizes */
		lput(entryvalue());		/* va of entry */
		lputl(INITTEXT);		/* text start */
		lputl(INITDAT);			/* data start */
		/*
		 * text section header
		 */
		s8put(&#34;.text&#34;);
		lputl(HEADR);			/* pa */
		lputl(HEADR);			/* va */
		lputl(textsize);		/* text size */
		lputl(HEADR);			/* file offset */
		lputl(0);			/* relocation */
		lputl(0);			/* line numbers */
		lputl(0);			/* relocation, line numbers */
		lputl(0x20);			/* flags text only */
		/*
		 * data section header
		 */
		s8put(&#34;.data&#34;);
		lputl(INITDAT);			/* pa */
		lputl(INITDAT);			/* va */
		lputl(datsize);			/* data size */
		lputl(HEADR+textsize);		/* file offset */
		lputl(0);			/* relocation */
		lputl(0);			/* line numbers */
		lputl(0);			/* relocation, line numbers */
		lputl(0x40);			/* flags data only */
		/*
		 * bss section header
		 */
		s8put(&#34;.bss&#34;);
		lputl(INITDAT+datsize);		/* pa */
		lputl(INITDAT+datsize);		/* va */
		lputl(bsssize);			/* bss size */
		lputl(0);			/* file offset */
		lputl(0);			/* relocation */
		lputl(0);			/* line numbers */
		lputl(0);			/* relocation, line numbers */
		lputl(0x80);			/* flags bss only */
		/*
		 * comment section header
		 */
		s8put(&#34;.comment&#34;);
		lputl(0);			/* pa */
		lputl(0);			/* va */
		lputl(symsize+lcsize);		/* comment size */
		lputl(HEADR+textsize+datsize);	/* file offset */
		lputl(HEADR+textsize+datsize);	/* offset of syms */
		lputl(HEADR+textsize+datsize+symsize);/* offset of line numbers */
		lputl(0);			/* relocation, line numbers */
		lputl(0x200);			/* flags comment only */
		break;
	case 2:	/* plan9 */
		magic = 4*11*11+7;
		if(dlm)
			magic |= 0x80000000;
		lput(magic);		/* magic */
		lput(textsize);			/* sizes */
		lput(datsize);
		lput(bsssize);
		lput(symsize);			/* nsyms */
		lput(entryvalue());		/* va of entry */
		lput(spsize);			/* sp offsets */
		lput(lcsize);			/* line offsets */
		break;
	case 3:
		/* MS-DOS .COM */
		break;
	case 4:
		/* fake MS-DOS .EXE */
		v = rnd(HEADR+textsize, INITRND)+datsize;
		wputl(0x5A4D);			/* &#39;MZ&#39; */
		wputl(v % 512);			/* bytes in last page */
		wputl(rnd(v, 512)/512);		/* total number of pages */
		wputl(0x0000);			/* number of reloc items */
		v = rnd(HEADR-(INITTEXT &amp; 0xFFFF), 16);
		wputl(v/16);			/* size of header */
		wputl(0x0000);			/* minimum allocation */
		wputl(0xFFFF);			/* maximum allocation */
		wputl(0x0000);			/* initial ss value */
		wputl(0x0100);			/* initial sp value */
		wputl(0x0000);			/* complemented checksum */
		v = entryvalue();
		wputl(v);			/* initial ip value (!) */
		wputl(0x0000);			/* initial cs value */
		wputl(0x0000);
		wputl(0x0000);
		wputl(0x003E);			/* reloc table offset */
		wputl(0x0000);			/* overlay number */
		break;

	case 6:
		asmbmacho(symdatva, symo);
		break;

	case 7:
	case 8:
		/* elf 386 */
		if(HEADTYPE == 8)
			debug[&#39;d&#39;] = 1;

		eh = getElfEhdr();
		fo = HEADR;
		startva = INITTEXT - HEADR;
		va = startva + fo;
		w = textsize;

		/* This null SHdr must appear before all others */
		sh = newElfShdr(elfstr[ElfStrEmpty]);

		/* program header info - but not on native client */
		pph = nil;
		if(HEADTYPE != 8) {
			pph = newElfPhdr();
			pph-&gt;type = PT_PHDR;
			pph-&gt;flags = PF_R + PF_X;
			pph-&gt;off = eh-&gt;ehsize;
			pph-&gt;vaddr = INITTEXT - HEADR + pph-&gt;off;
			pph-&gt;paddr = INITTEXT - HEADR + pph-&gt;off;
			pph-&gt;align = INITRND;
		}

		if(!debug[&#39;d&#39;]) {
			/* interpreter */
			sh = newElfShdr(elfstr[ElfStrInterp]);
			sh-&gt;type = SHT_PROGBITS;
			sh-&gt;flags = SHF_ALLOC;
			sh-&gt;addralign = 1;
			shsym(sh, lookup(&#34;.interp&#34;, 0));

			ph = newElfPhdr();
			ph-&gt;type = PT_INTERP;
			ph-&gt;flags = PF_R;
			phsh(ph, sh);
		}

		ph = newElfPhdr();
		ph-&gt;type = PT_LOAD;
		ph-&gt;flags = PF_X+PF_R;
		ph-&gt;vaddr = va;
		ph-&gt;paddr = va;
		ph-&gt;off = fo;
		ph-&gt;filesz = w;
		ph-&gt;memsz = w;
		ph-&gt;align = INITRND;

		fo = rnd(fo+w, INITRND);
		va = rnd(va+w, INITRND);
		w = datsize;

		ph = newElfPhdr();
		ph-&gt;type = PT_LOAD;
		ph-&gt;flags = PF_W+PF_R;
		ph-&gt;off = fo;
		ph-&gt;vaddr = va;
		ph-&gt;paddr = va;
		ph-&gt;filesz = w;
		ph-&gt;memsz = w+bsssize;
		ph-&gt;align = INITRND;

		if(!debug[&#39;s&#39;] &amp;&amp; HEADTYPE != 8) {
			ph = newElfPhdr();
			ph-&gt;type = PT_LOAD;
			ph-&gt;flags = PF_W+PF_R;
			ph-&gt;off = symo;
			ph-&gt;vaddr = symdatva;
			ph-&gt;paddr = symdatva;
			ph-&gt;filesz = 8+symsize+lcsize;
			ph-&gt;memsz = 8+symsize+lcsize;
			ph-&gt;align = INITRND;
		}

		/* Dynamic linking sections */
		if (!debug[&#39;d&#39;]) {	/* -d suppresses dynamic loader format */
			/* S headers for dynamic linking */
			sh = newElfShdr(elfstr[ElfStrGot]);
			sh-&gt;type = SHT_PROGBITS;
			sh-&gt;flags = SHF_ALLOC+SHF_WRITE;
			sh-&gt;entsize = 4;
			sh-&gt;addralign = 4;
			shsym(sh, lookup(&#34;.got&#34;, 0));

			sh = newElfShdr(elfstr[ElfStrGotPlt]);
			sh-&gt;type = SHT_PROGBITS;
			sh-&gt;flags = SHF_ALLOC+SHF_WRITE;
			sh-&gt;entsize = 4;
			sh-&gt;addralign = 4;
			shsym(sh, lookup(&#34;.got.plt&#34;, 0));

			dynsym = eh-&gt;shnum;
			sh = newElfShdr(elfstr[ElfStrDynsym]);
			sh-&gt;type = SHT_DYNSYM;
			sh-&gt;flags = SHF_ALLOC;
			sh-&gt;entsize = ELF32SYMSIZE;
			sh-&gt;addralign = 4;
			sh-&gt;link = dynsym+1;	// dynstr
			// sh-&gt;info = index of first non-local symbol (number of local symbols)
			shsym(sh, lookup(&#34;.dynsym&#34;, 0));

			sh = newElfShdr(elfstr[ElfStrDynstr]);
			sh-&gt;type = SHT_STRTAB;
			sh-&gt;flags = SHF_ALLOC;
			sh-&gt;addralign = 1;
			shsym(sh, lookup(&#34;.dynstr&#34;, 0));

			sh = newElfShdr(elfstr[ElfStrHash]);
			sh-&gt;type = SHT_HASH;
			sh-&gt;flags = SHF_ALLOC;
			sh-&gt;entsize = 4;
			sh-&gt;addralign = 4;
			sh-&gt;link = dynsym;
			shsym(sh, lookup(&#34;.hash&#34;, 0));

			sh = newElfShdr(elfstr[ElfStrRel]);
			sh-&gt;type = SHT_REL;
			sh-&gt;flags = SHF_ALLOC;
			sh-&gt;entsize = ELF32RELSIZE;
			sh-&gt;addralign = 4;
			sh-&gt;link = dynsym;
			shsym(sh, lookup(&#34;.rel&#34;, 0));

			/* sh and PT_DYNAMIC for .dynamic section */
			sh = newElfShdr(elfstr[ElfStrDynamic]);
			sh-&gt;type = SHT_DYNAMIC;
			sh-&gt;flags = SHF_ALLOC+SHF_WRITE;
			sh-&gt;entsize = 8;
			sh-&gt;addralign = 4;
			sh-&gt;link = dynsym+1;	// dynstr
			shsym(sh, lookup(&#34;.dynamic&#34;, 0));
			ph = newElfPhdr();
			ph-&gt;type = PT_DYNAMIC;
			ph-&gt;flags = PF_R + PF_W;
			phsh(ph, sh);

			/*
			 * Thread-local storage segment (really just size).
			 */
			if(tlsoffset != 0) {
				ph = newElfPhdr();
				ph-&gt;type = PT_TLS;
				ph-&gt;flags = PF_R;
				ph-&gt;memsz = -tlsoffset;
				ph-&gt;align = 4;
			}
		}

		ph = newElfPhdr();
		ph-&gt;type = PT_GNU_STACK;
		ph-&gt;flags = PF_W+PF_R;
		ph-&gt;align = 4;

		fo = ELFRESERVE;
		va = startva + fo;
		w = textsize;

		sh = newElfShdr(elfstr[ElfStrText]);
		sh-&gt;type = SHT_PROGBITS;
		sh-&gt;flags = SHF_ALLOC+SHF_EXECINSTR;
		sh-&gt;addr = va;
		sh-&gt;off = fo;
		sh-&gt;size = w;
		sh-&gt;addralign = 4;

		fo = rnd(fo+w, INITRND);
		va = rnd(va+w, INITRND);
		w = datsize;

		sh = newElfShdr(elfstr[ElfStrData]);
		sh-&gt;type = SHT_PROGBITS;
		sh-&gt;flags = SHF_WRITE+SHF_ALLOC;
		sh-&gt;addr = va;
		sh-&gt;off = fo;
		sh-&gt;size = w;
		sh-&gt;addralign = 4;

		fo += w;
		va += w;
		w = bsssize;

		sh = newElfShdr(elfstr[ElfStrBss]);
		sh-&gt;type = SHT_NOBITS;
		sh-&gt;flags = SHF_WRITE+SHF_ALLOC;
		sh-&gt;addr = va;
		sh-&gt;off = fo;
		sh-&gt;size = w;
		sh-&gt;addralign = 4;

		if (!debug[&#39;s&#39;]) {
			fo = symo+8;
			w = symsize;

			sh = newElfShdr(elfstr[ElfStrGosymtab]);
			sh-&gt;type = SHT_PROGBITS;
			sh-&gt;off = fo;
			sh-&gt;size = w;
			sh-&gt;addralign = 1;

			fo += w;
			w = lcsize;

			sh = newElfShdr(elfstr[ElfStrGopclntab]);
			sh-&gt;type = SHT_PROGBITS;
			sh-&gt;off = fo;
			sh-&gt;size = w;
			sh-&gt;addralign = 1;
		}

		sh = newElfShstrtab(elfstr[ElfStrShstrtab]);
		sh-&gt;type = SHT_STRTAB;
		sh-&gt;addralign = 1;
		shsym(sh, lookup(&#34;.shstrtab&#34;, 0));

		/* Main header */
		eh-&gt;ident[EI_MAG0] = &#39;\177&#39;;
		eh-&gt;ident[EI_MAG1] = &#39;E&#39;;
		eh-&gt;ident[EI_MAG2] = &#39;L&#39;;
		eh-&gt;ident[EI_MAG3] = &#39;F&#39;;
		eh-&gt;ident[EI_CLASS] = ELFCLASS32;
		eh-&gt;ident[EI_DATA] = ELFDATA2LSB;
		eh-&gt;ident[EI_VERSION] = EV_CURRENT;
		if(HEADTYPE == 8) {
			eh-&gt;ident[EI_OSABI] = ELFOSABI_NACL;
			eh-&gt;ident[EI_ABIVERSION] = 6;
			eh-&gt;flags = 0x200000;	// aligned mod 32
		}

		eh-&gt;type = ET_EXEC;
		eh-&gt;machine = EM_386;
		eh-&gt;version = EV_CURRENT;
		eh-&gt;entry = entryvalue();

		if(pph != nil) {
			pph-&gt;filesz = eh-&gt;phnum * eh-&gt;phentsize;
			pph-&gt;memsz = pph-&gt;filesz;
		}

		seek(cout, 0, 0);
		a = 0;
		a += elfwritehdr();
		a += elfwritephdrs();
		a += elfwriteshdrs();
		if (a &gt; ELFRESERVE) {
			diag(&#34;ELFRESERVE too small: %d &gt; %d&#34;, a, ELFRESERVE);
		}
		cflush();
		break;
	}
	cflush();
}

void
s8put(char *n)
{
	char name[8];
	int i;

	strncpy(name, n, sizeof(name));
	for(i=0; i&lt;sizeof(name); i++)
		cput(name[i]);
}

void
cflush(void)
{
	int n;

	n = sizeof(buf.cbuf) - cbc;
	if(n)
		write(cout, buf.cbuf, n);
	cbp = buf.cbuf;
	cbc = sizeof(buf.cbuf);
}

void
datblk(int32 s, int32 n)
{
	Prog *p;
	char *cast;
	int32 l, fl, j;
	int i, c;
	Adr *a;

	memset(buf.dbuf, 0, n+Dbufslop);
	for(p = datap; p != P; p = p-&gt;link) {
		a = &amp;p-&gt;from;

		l = a-&gt;sym-&gt;value + a-&gt;offset - s;
		if(l &gt;= n)
			continue;

		c = a-&gt;scale;
		i = 0;
		if(l &lt; 0) {
			if(l+c &lt;= 0)
				continue;
			i = -l;
			l = 0;
		}

		curp = p;
		if(!a-&gt;sym-&gt;reachable)
			diag(&#34;unreachable symbol in datblk - %s&#34;, a-&gt;sym-&gt;name);
		if(a-&gt;sym-&gt;type == SMACHO)
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
		case D_FCONST:
			switch(c) {
			default:
			case 4:
				fl = ieeedtof(&amp;p-&gt;to.ieee);
				cast = (char*)&amp;fl;
				for(; i&lt;c; i++) {
					buf.dbuf[l] = cast[fnuxi4[i]];
					l++;
				}
				break;
			case 8:
				cast = (char*)&amp;p-&gt;to.ieee;
				for(; i&lt;c; i++) {
					buf.dbuf[l] = cast[fnuxi8[i]];
					l++;
				}
				break;
			}
			break;

		case D_SCONST:
			for(; i&lt;c; i++) {
				buf.dbuf[l] = p-&gt;to.scon[i];
				l++;
			}
			break;

		default:
			fl = p-&gt;to.offset;
			if(p-&gt;to.type == D_SIZE)
				fl += p-&gt;to.sym-&gt;size;
			if(p-&gt;to.type == D_ADDR) {
				if(p-&gt;to.index != D_STATIC &amp;&amp; p-&gt;to.index != D_EXTERN)
					diag(&#34;DADDR type%P&#34;, p);
				if(p-&gt;to.sym) {
					if(p-&gt;to.sym-&gt;type == SUNDEF)
						ckoff(p-&gt;to.sym, fl);
					fl += p-&gt;to.sym-&gt;value;
					if(p-&gt;to.sym-&gt;type != STEXT &amp;&amp; p-&gt;to.sym-&gt;type != SUNDEF)
						fl += INITDAT;
					if(dlm)
						dynreloc(p-&gt;to.sym, l+s+INITDAT, 1);
				}
			}
			cast = (char*)&amp;fl;
			switch(c) {
			default:
				diag(&#34;bad nuxi %d %d\n%P&#34;, c, i, curp);
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
		}
	}

	write(cout, buf.dbuf, n);
	if(!debug[&#39;a&#39;])
		return;

	/*
	 * a second pass just to print the asm
	 */
	for(p = datap; p != P; p = p-&gt;link) {
		a = &amp;p-&gt;from;

		l = a-&gt;sym-&gt;value + a-&gt;offset - s;
		if(l &lt; 0 || l &gt;= n)
			continue;

		c = a-&gt;scale;
		i = 0;

		switch(p-&gt;to.type) {
		case D_FCONST:
			switch(c) {
			default:
			case 4:
				fl = ieeedtof(&amp;p-&gt;to.ieee);
				cast = (char*)&amp;fl;
				Bprint(&amp;bso, pcstr, l+s+INITDAT);
				for(j=0; j&lt;c; j++)
					Bprint(&amp;bso, &#34;%.2ux&#34;, cast[fnuxi4[j]] &amp; 0xff);
				Bprint(&amp;bso, &#34;\t%P\n&#34;, curp);
				break;
			case 8:
				cast = (char*)&amp;p-&gt;to.ieee;
				Bprint(&amp;bso, pcstr, l+s+INITDAT);
				for(j=0; j&lt;c; j++)
					Bprint(&amp;bso, &#34;%.2ux&#34;, cast[fnuxi8[j]] &amp; 0xff);
				Bprint(&amp;bso, &#34;\t%P\n&#34;, curp);
				break;
			}
			break;

		case D_SCONST:
			Bprint(&amp;bso, pcstr, l+s+INITDAT);
			for(j=0; j&lt;c; j++)
				Bprint(&amp;bso, &#34;%.2ux&#34;, p-&gt;to.scon[j] &amp; 0xff);
			Bprint(&amp;bso, &#34;\t%P\n&#34;, curp);
			break;

		default:
			fl = p-&gt;to.offset;
			if(p-&gt;to.type == D_SIZE)
				fl += p-&gt;to.sym-&gt;size;
			if(p-&gt;to.type == D_ADDR) {
				if(p-&gt;to.index != D_STATIC &amp;&amp; p-&gt;to.index != D_EXTERN)
					diag(&#34;DADDR type%P&#34;, p);
				if(p-&gt;to.sym) {
					if(p-&gt;to.sym-&gt;type == SUNDEF)
						ckoff(p-&gt;to.sym, fl);
					fl += p-&gt;to.sym-&gt;value;
					if(p-&gt;to.sym-&gt;type != STEXT &amp;&amp; p-&gt;to.sym-&gt;type != SUNDEF)
						fl += INITDAT;
					if(dlm)
						dynreloc(p-&gt;to.sym, l+s+INITDAT, 1);
				}
			}
			cast = (char*)&amp;fl;
			switch(c) {
			default:
				diag(&#34;bad nuxi %d %d\n%P&#34;, c, i, curp);
				break;
			case 1:
				Bprint(&amp;bso, pcstr, l+s+INITDAT);
				for(j=0; j&lt;c; j++)
					Bprint(&amp;bso, &#34;%.2ux&#34;, cast[inuxi1[j]] &amp; 0xff);
				Bprint(&amp;bso, &#34;\t%P\n&#34;, curp);
				break;
			case 2:
				Bprint(&amp;bso, pcstr, l+s+INITDAT);
				for(j=0; j&lt;c; j++)
					Bprint(&amp;bso, &#34;%.2ux&#34;, cast[inuxi2[j]] &amp; 0xff);
				Bprint(&amp;bso, &#34;\t%P\n&#34;, curp);
				break;
			case 4:
				Bprint(&amp;bso, pcstr, l+s+INITDAT);
				for(j=0; j&lt;c; j++)
					Bprint(&amp;bso, &#34;%.2ux&#34;, cast[inuxi4[j]] &amp; 0xff);
				Bprint(&amp;bso, &#34;\t%P\n&#34;, curp);
				break;
			}
			break;
		}
	}
}

int32
rnd(int32 v, int32 r)
{
	int32 c;

	if(r &lt;= 0)
		return v;
	v += r - 1;
	c = v % r;
	if(c &lt; 0)
		c += r;
	v -= c;
	return v;
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
