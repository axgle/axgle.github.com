<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/ld/macho.c</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/ld/macho.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Mach-O file writing
// http://developer.apple.com/mac/library/DOCUMENTATION/DeveloperTools/Conceptual/MachORuntime/Reference/reference.html

#include &#34;l.h&#34;
#include &#34;../ld/lib.h&#34;
#include &#34;../ld/macho.h&#34;

static	int	macho64;
static	MachoHdr	hdr;
static	MachoLoad	load[16];
static	MachoSeg	seg[16];
static	MachoDebug	xdebug[16];
static	int	nload, nseg, ndebug, nsect;

void
machoinit(void)
{
	switch(thechar) {
	// 64-bit architectures
	case &#39;6&#39;:
		macho64 = 1;
		break;

	// 32-bit architectures
	default:
		break;
	}
}

MachoHdr*
getMachoHdr(void)
{
	return &amp;hdr;
}

MachoLoad*
newMachoLoad(uint32 type, uint32 ndata)
{
	MachoLoad *l;

	if(nload &gt;= nelem(load)) {
		diag(&#34;too many loads&#34;);
		errorexit();
	}
	l = &amp;load[nload++];
	l-&gt;type = type;
	l-&gt;ndata = ndata;
	l-&gt;data = mal(ndata*4);
	return l;
}

MachoSeg*
newMachoSeg(char *name, int msect)
{
	MachoSeg *s;

	if(nseg &gt;= nelem(seg)) {
		diag(&#34;too many segs&#34;);
		errorexit();
	}
	s = &amp;seg[nseg++];
	s-&gt;name = name;
	s-&gt;msect = msect;
	s-&gt;sect = mal(msect*sizeof s-&gt;sect[0]);
	return s;
}

MachoSect*
newMachoSect(MachoSeg *seg, char *name)
{
	MachoSect *s;

	if(seg-&gt;nsect &gt;= seg-&gt;msect) {
		diag(&#34;too many sects in segment %s&#34;, seg-&gt;name);
		errorexit();
	}
	s = &amp;seg-&gt;sect[seg-&gt;nsect++];
	s-&gt;name = name;
	nsect++;
	return s;
}

MachoDebug*
newMachoDebug(void)
{
	if(ndebug &gt;= nelem(xdebug)) {
		diag(&#34;too many debugs&#34;);
		errorexit();
	}
	return &amp;xdebug[ndebug++];
}


// Generic linking code.

static uchar *linkdata;
static uint32 nlinkdata;
static uint32 mlinkdata;

static uchar *strtab;
static uint32 nstrtab;
static uint32 mstrtab;

static char **dylib;
static int ndylib;

static vlong linkoff;

int
machowrite(void)
{
	vlong o1;
	int loadsize;
	int i, j;
	MachoSeg *s;
	MachoSect *t;
	MachoDebug *d;
	MachoLoad *l;

	o1 = Boffset(&amp;bso);

	loadsize = 4*4*ndebug;
	for(i=0; i&lt;nload; i++)
		loadsize += 4*(load[i].ndata+2);
	if(macho64) {
		loadsize += 18*4*nseg;
		loadsize += 20*4*nsect;
	} else {
		loadsize += 14*4*nseg;
		loadsize += 17*4*nsect;
	}

	if(macho64)
		LPUT(0xfeedfacf);
	else
		LPUT(0xfeedface);
	LPUT(hdr.cpu);
	LPUT(hdr.subcpu);
	LPUT(2);	/* file type - mach executable */
	LPUT(nload+nseg+ndebug);
	LPUT(loadsize);
	LPUT(1);	/* flags - no undefines */
	if(macho64)
		LPUT(0);	/* reserved */

	for(i=0; i&lt;nseg; i++) {
		s = &amp;seg[i];
		if(macho64) {
			LPUT(25);	/* segment 64 */
			LPUT(72+80*s-&gt;nsect);
			strnput(s-&gt;name, 16);
			VPUT(s-&gt;vaddr);
			VPUT(s-&gt;vsize);
			VPUT(s-&gt;fileoffset);
			VPUT(s-&gt;filesize);
			LPUT(s-&gt;prot1);
			LPUT(s-&gt;prot2);
			LPUT(s-&gt;nsect);
			LPUT(s-&gt;flag);
		} else {
			LPUT(1);	/* segment 32 */
			LPUT(56+68*s-&gt;nsect);
			strnput(s-&gt;name, 16);
			LPUT(s-&gt;vaddr);
			LPUT(s-&gt;vsize);
			LPUT(s-&gt;fileoffset);
			LPUT(s-&gt;filesize);
			LPUT(s-&gt;prot1);
			LPUT(s-&gt;prot2);
			LPUT(s-&gt;nsect);
			LPUT(s-&gt;flag);
		}
		for(j=0; j&lt;s-&gt;nsect; j++) {
			t = &amp;s-&gt;sect[j];
			if(macho64) {
				strnput(t-&gt;name, 16);
				strnput(s-&gt;name, 16);
				VPUT(t-&gt;addr);
				VPUT(t-&gt;size);
				LPUT(t-&gt;off);
				LPUT(t-&gt;align);
				LPUT(t-&gt;reloc);
				LPUT(t-&gt;nreloc);
				LPUT(t-&gt;flag);
				LPUT(0);	/* reserved */
				LPUT(0);	/* reserved */
				LPUT(0);	/* reserved */
			} else {
				strnput(t-&gt;name, 16);
				strnput(s-&gt;name, 16);
				LPUT(t-&gt;addr);
				LPUT(t-&gt;size);
				LPUT(t-&gt;off);
				LPUT(t-&gt;align);
				LPUT(t-&gt;reloc);
				LPUT(t-&gt;nreloc);
				LPUT(t-&gt;flag);
				LPUT(0);	/* reserved */
				LPUT(0);	/* reserved */
			}
		}
	}

	for(i=0; i&lt;nload; i++) {
		l = &amp;load[i];
		LPUT(l-&gt;type);
		LPUT(4*(l-&gt;ndata+2));
		for(j=0; j&lt;l-&gt;ndata; j++)
			LPUT(l-&gt;data[j]);
	}

	for(i=0; i&lt;ndebug; i++) {
		d = &amp;xdebug[i];
		LPUT(3);	/* obsolete gdb debug info */
		LPUT(16);	/* size of symseg command */
		LPUT(d-&gt;fileoffset);
		LPUT(d-&gt;filesize);
	}

	return Boffset(&amp;bso) - o1;
}

static void*
grow(uchar **dat, uint32 *ndat, uint32 *mdat, uint32 n)
{
	uchar *p;
	uint32 old;

	if(*ndat+n &gt; *mdat) {
		old = *mdat;
		*mdat = (*ndat+n)*2 + 128;
		*dat = realloc(*dat, *mdat);
		if(*dat == 0) {
			diag(&#34;out of memory&#34;);
			errorexit();
		}
		memset(*dat+old, 0, *mdat-old);
	}
	p = *dat + *ndat;
	*ndat += n;
	return p;
}

static int
needlib(char *name)
{
	char *p;
	Sym *s;

	/* reuse hash code in symbol table */
	p = smprint(&#34;.machoload.%s&#34;, name);
	s = lookup(p, 0);
	if(s-&gt;type == 0) {
		s-&gt;type = 100;	// avoid SDATA, etc.
		return 1;
	}
	return 0;
}

void
domacho(void)
{
	int h, nsym, ptrsize;
	char *p;
	uchar *dat;
	uint32 x;
	Sym *s;

	ptrsize = 4;
	if(macho64)
		ptrsize = 8;

	// empirically, string table must begin with &#34; \x00&#34;.
	if(!debug[&#39;d&#39;])
		*(char*)grow(&amp;strtab, &amp;nstrtab, &amp;mstrtab, 2) = &#39; &#39;;

	nsym = 0;
	for(h=0; h&lt;NHASH; h++) {
		for(s=hash[h]; s!=S; s=s-&gt;link) {
			if(!s-&gt;reachable || (s-&gt;type != SDATA &amp;&amp; s-&gt;type != SBSS) || s-&gt;dynldname == nil)
				continue;
			if(debug[&#39;d&#39;]) {
				diag(&#34;cannot use dynamic loading and -d&#34;);
				errorexit();
			}
			s-&gt;type = SMACHO;
			s-&gt;value = nsym*ptrsize;

			/* symbol table entry - darwin still puts _ prefixes on all C symbols */
			x = nstrtab;
			p = grow(&amp;strtab, &amp;nstrtab, &amp;mstrtab, 1+strlen(s-&gt;dynldname)+1);
			*p++ = &#39;_&#39;;
			strcpy(p, s-&gt;dynldname);

			dat = grow(&amp;linkdata, &amp;nlinkdata, &amp;mlinkdata, 8+ptrsize);
			dat[0] = x;
			dat[1] = x&gt;&gt;8;
			dat[2] = x&gt;&gt;16;
			dat[3] = x&gt;&gt;24;
			dat[4] = 0x01;	// type: N_EXT - external symbol

			if(needlib(s-&gt;dynldlib)) {
				if(ndylib%32 == 0) {
					dylib = realloc(dylib, (ndylib+32)*sizeof dylib[0]);
					if(dylib == nil) {
						diag(&#34;out of memory&#34;);
						errorexit();
					}
				}
				dylib[ndylib++] = s-&gt;dynldlib;
			}
			nsym++;
		}
	}

	/*
	 * list of symbol table indexes.
	 * we don&#39;t take advantage of the opportunity
	 * to order the symbol table differently from
	 * this list, so it is boring: 0 1 2 3 4 ...
	 */
	for(x=0; x&lt;nsym; x++) {
		dat = grow(&amp;linkdata, &amp;nlinkdata, &amp;mlinkdata, 4);
		dat[0] = x;
		dat[1] = x&gt;&gt;8;
		dat[2] = x&gt;&gt;16;
		dat[3] = x&gt;&gt;24;
	}

	dynptrsize = nsym*ptrsize;
}

vlong
domacholink(void)
{
	linkoff = 0;
	if(nlinkdata &gt; 0) {
		linkoff = rnd(HEADR+textsize, INITRND) + rnd(datsize, INITRND);
		seek(cout, linkoff, 0);
		write(cout, linkdata, nlinkdata);
		write(cout, strtab, nstrtab);
	}
	return rnd(nlinkdata+nstrtab, INITRND);
}

void
asmbmacho(vlong symdatva, vlong symo)
{
	vlong v, w;
	vlong va;
	int a, i, ptrsize;
	MachoHdr *mh;
	MachoSect *msect;
	MachoSeg *ms;
	MachoDebug *md;
	MachoLoad *ml;

	/* apple MACH */
	va = INITTEXT - HEADR;
	mh = getMachoHdr();
	switch(thechar){
	default:
		diag(&#34;unknown mach architecture&#34;);
		errorexit();
	case &#39;6&#39;:
		mh-&gt;cpu = MACHO_CPU_AMD64;
		mh-&gt;subcpu = MACHO_SUBCPU_X86;
		ptrsize = 8;
		break;
	case &#39;8&#39;:
		mh-&gt;cpu = MACHO_CPU_386;
		mh-&gt;subcpu = MACHO_SUBCPU_X86;
		ptrsize = 4;
		break;
	}

	/* segment for zero page */
	ms = newMachoSeg(&#34;__PAGEZERO&#34;, 0);
	ms-&gt;vsize = va;

	/* text */
	v = rnd(HEADR+textsize, INITRND);
	ms = newMachoSeg(&#34;__TEXT&#34;, 1);
	ms-&gt;vaddr = va;
	ms-&gt;vsize = v;
	ms-&gt;filesize = v;
	ms-&gt;prot1 = 7;
	ms-&gt;prot2 = 5;

	msect = newMachoSect(ms, &#34;__text&#34;);
	msect-&gt;addr = INITTEXT;
	msect-&gt;size = textsize;
	msect-&gt;off = INITTEXT - va;
	msect-&gt;flag = 0x400;	/* flag - some instructions */

	/* data */
	w = datsize+dynptrsize+bsssize;
	ms = newMachoSeg(&#34;__DATA&#34;, 2+(dynptrsize&gt;0));
	ms-&gt;vaddr = va+v;
	ms-&gt;vsize = w;
	ms-&gt;fileoffset = v;
	ms-&gt;filesize = datsize;
	ms-&gt;prot1 = 7;
	ms-&gt;prot2 = 3;

	msect = newMachoSect(ms, &#34;__data&#34;);
	msect-&gt;addr = va+v;
	msect-&gt;size = datsize;
	msect-&gt;off = v;

	if(dynptrsize &gt; 0) {
		msect = newMachoSect(ms, &#34;__nl_symbol_ptr&#34;);
		msect-&gt;addr = va+v+datsize;
		msect-&gt;size = dynptrsize;
		msect-&gt;align = 2;
		msect-&gt;flag = 6;	/* section with nonlazy symbol pointers */
		/*
		 * The reserved1 field is supposed to be the index of
		 * the first entry in the list of symbol table indexes
		 * in isymtab for the symbols we need.  We only use
		 * pointers, so we need the entire list, so the index
		 * here should be 0, which luckily is what the Mach-O
		 * writing code emits by default for this not really reserved field.
		msect-&gt;reserved1 = 0; - first indirect symbol table entry we need
		 */
	}

	msect = newMachoSect(ms, &#34;__bss&#34;);
	msect-&gt;addr = va+v+datsize+dynptrsize;
	msect-&gt;size = bsssize;
	msect-&gt;flag = 1;	/* flag - zero fill */

	switch(thechar) {
	default:
		diag(&#34;unknown macho architecture&#34;);
		errorexit();
	case &#39;6&#39;:
		ml = newMachoLoad(5, 42+2);	/* unix thread */
		ml-&gt;data[0] = 4;	/* thread type */
		ml-&gt;data[1] = 42;	/* word count */
		ml-&gt;data[2+32] = entryvalue();	/* start pc */
		ml-&gt;data[2+32+1] = entryvalue()&gt;&gt;32;
		break;
	case &#39;8&#39;:
		ml = newMachoLoad(5, 16+2);	/* unix thread */
		ml-&gt;data[0] = 1;	/* thread type */
		ml-&gt;data[1] = 16;	/* word count */
		ml-&gt;data[2+10] = entryvalue();	/* start pc */
		break;
	}

	if(!debug[&#39;d&#39;]) {
		int nsym;

		nsym = dynptrsize/ptrsize;

		ms = newMachoSeg(&#34;__LINKEDIT&#34;, 0);
		ms-&gt;vaddr = va+v+rnd(datsize+dynptrsize+bsssize, INITRND);
		ms-&gt;vsize = nlinkdata+nstrtab;
		ms-&gt;fileoffset = linkoff;
		ms-&gt;filesize = nlinkdata+nstrtab;
		ms-&gt;prot1 = 7;
		ms-&gt;prot2 = 3;

		ml = newMachoLoad(2, 4);	/* LC_SYMTAB */
		ml-&gt;data[0] = linkoff;	/* symoff */
		ml-&gt;data[1] = nsym;	/* nsyms */
		ml-&gt;data[2] = linkoff + nlinkdata;	/* stroff */
		ml-&gt;data[3] = nstrtab;	/* strsize */

		ml = newMachoLoad(11, 18);	/* LC_DYSYMTAB */
		ml-&gt;data[0] = 0;	/* ilocalsym */
		ml-&gt;data[1] = 0;	/* nlocalsym */
		ml-&gt;data[2] = 0;	/* iextdefsym */
		ml-&gt;data[3] = 0;	/* nextdefsym */
		ml-&gt;data[4] = 0;	/* iundefsym */
		ml-&gt;data[5] = nsym;	/* nundefsym */
		ml-&gt;data[6] = 0;	/* tocoffset */
		ml-&gt;data[7] = 0;	/* ntoc */
		ml-&gt;data[8] = 0;	/* modtaboff */
		ml-&gt;data[9] = 0;	/* nmodtab */
		ml-&gt;data[10] = 0;	/* extrefsymoff */
		ml-&gt;data[11] = 0;	/* nextrefsyms */
		ml-&gt;data[12] = linkoff + nlinkdata - nsym*4;	/* indirectsymoff */
		ml-&gt;data[13] = nsym;	/* nindirectsyms */
		ml-&gt;data[14] = 0;	/* extreloff */
		ml-&gt;data[15] = 0;	/* nextrel */
		ml-&gt;data[16] = 0;	/* locreloff */
		ml-&gt;data[17] = 0;	/* nlocrel */

		ml = newMachoLoad(14, 6);	/* LC_LOAD_DYLINKER */
		ml-&gt;data[0] = 12;	/* offset to string */
		strcpy((char*)&amp;ml-&gt;data[1], &#34;/usr/lib/dyld&#34;);

		for(i=0; i&lt;ndylib; i++) {
			ml = newMachoLoad(12, 4+(strlen(dylib[i])+1+7)/8*2);	/* LC_LOAD_DYLIB */
			ml-&gt;data[0] = 24;	/* offset of string from beginning of load */
			ml-&gt;data[1] = 0;	/* time stamp */
			ml-&gt;data[2] = 0;	/* version */
			ml-&gt;data[3] = 0;	/* compatibility version */
			strcpy((char*)&amp;ml-&gt;data[4], dylib[i]);
		}
	}

	if(!debug[&#39;s&#39;]) {
		ms = newMachoSeg(&#34;__SYMDAT&#34;, 1);
		ms-&gt;vaddr = symdatva;
		ms-&gt;vsize = 8+symsize+lcsize;
		ms-&gt;fileoffset = symo;
		ms-&gt;filesize = 8+symsize+lcsize;
		ms-&gt;prot1 = 7;
		ms-&gt;prot2 = 5;

		md = newMachoDebug();
		md-&gt;fileoffset = symo+8;
		md-&gt;filesize = symsize;

		md = newMachoDebug();
		md-&gt;fileoffset = symo+8+symsize;
		md-&gt;filesize = lcsize;
	}

	a = machowrite();
	if(a &gt; MACHORESERVE)
		diag(&#34;MACHORESERVE too small: %d &gt; %d&#34;, a, MACHORESERVE);
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
