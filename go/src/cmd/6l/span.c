<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/6l/span.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/6l/span.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/6l/span.c
// http://code.google.com/p/inferno-os/source/browse/utils/6l/span.c
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

static int	rexflag;
static int	asmode;

void
span(void)
{
	Prog *p, *q;
	int32 v;
	vlong c, idat;
	int m, n, again;

	xdefine(&#34;etext&#34;, STEXT, 0L);
	idat = INITDAT;
	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT)
			curtext = p;
		n = 0;
		if(p-&gt;to.type == D_BRANCH)
			if(p-&gt;pcond == P)
				p-&gt;pcond = p;
		if((q = p-&gt;pcond) != P)
			if(q-&gt;back != 2)
				n = 1;
		p-&gt;back = n;
		if(p-&gt;as == AADJSP) {
			p-&gt;to.type = D_SP;
			v = -p-&gt;from.offset;
			p-&gt;from.offset = v;
			p-&gt;as = p-&gt;mode != 64? AADDL: AADDQ;
			if(v &lt; 0) {
				p-&gt;as = p-&gt;mode != 64? ASUBL: ASUBQ;
				v = -v;
				p-&gt;from.offset = v;
			}
			if(v == 0)
				p-&gt;as = ANOP;
		}
	}
	n = 0;

start:
	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f span\n&#34;, cputime());
	Bflush(&amp;bso);
	c = INITTEXT;
	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT)
			curtext = p;
		if(p-&gt;to.type == D_BRANCH)
			if(p-&gt;back)
				p-&gt;pc = c;
		asmins(p);
		p-&gt;pc = c;
		m = andptr-and;
		p-&gt;mark = m;
		c += m;
	}

loop:
	n++;
	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f span %d\n&#34;, cputime(), n);
	Bflush(&amp;bso);
	if(n &gt; 50) {
		print(&#34;span must be looping\n&#34;);
		errorexit();
	}
	again = 0;
	c = INITTEXT;
	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT)
			curtext = p;
		if(p-&gt;to.type == D_BRANCH || p-&gt;back &amp; 0100) {
			if(p-&gt;back)
				p-&gt;pc = c;
			asmins(p);
			m = andptr-and;
			if(m != p-&gt;mark) {
				p-&gt;mark = m;
				again++;
			}
		}
		p-&gt;pc = c;
		c += p-&gt;mark;
	}
	if(again) {
		textsize = c;
		goto loop;
	}
	if(INITRND) {
		INITDAT = rnd(c, INITRND);
		if(INITDAT != idat) {
			idat = INITDAT;
			goto start;
		}
	}
	xdefine(&#34;etext&#34;, STEXT, c);
	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;etext = %llux\n&#34;, c);
	Bflush(&amp;bso);
	for(p = textp; p != P; p = p-&gt;pcond)
		p-&gt;from.sym-&gt;value = p-&gt;pc;
	textsize = c - INITTEXT;
}

void
xdefine(char *p, int t, vlong v)
{
	Sym *s;

	s = lookup(p, 0);
	if(s-&gt;type == 0 || s-&gt;type == SXREF) {
		s-&gt;type = t;
		s-&gt;value = v;
	}
	if(s-&gt;type == STEXT &amp;&amp; s-&gt;value == 0)
		s-&gt;value = v;
}

void
putsymb(char *s, int t, vlong v, vlong size, int ver, Sym *go)
{
	int i, f, l;
	vlong gv;

	if(t == &#39;f&#39;)
		s++;
	l = 4;
	if(!debug[&#39;8&#39;]){
		lputb(v&gt;&gt;32);
		l = 8;
	}
	lputb(v);
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
	gv = 0;
	if(go) {
		if(!go-&gt;reachable)
			diag(&#34;unreachable type %s&#34;, go-&gt;name);
		gv = go-&gt;value+INITDAT;
	}
	if(l == 8)
		lputb(gv&gt;&gt;32);
	lputb(gv);
	symsize += l + 1 + i+1 + l;

	if(debug[&#39;n&#39;]) {
		if(t == &#39;z&#39; || t == &#39;Z&#39;) {
			Bprint(&amp;bso, &#34;%c %.8llux &#34;, t, v);
			for(i=1; s[i] != 0 || s[i+1] != 0; i+=2) {
				f = ((s[i]&amp;0xff) &lt;&lt; 8) | (s[i+1]&amp;0xff);
				Bprint(&amp;bso, &#34;/%x&#34;, f);
			}
			Bprint(&amp;bso, &#34;\n&#34;);
			return;
		}
		if(ver)
			Bprint(&amp;bso, &#34;%c %.8llux %s&lt;%d&gt; %s (%.8llux)\n&#34;, t, v, s, ver, go ? go-&gt;name : &#34;&#34;, gv);
		else
			Bprint(&amp;bso, &#34;%c %.8llux %s %s (%.8llux)\n&#34;, t, v, s, go ? go-&gt;name : &#34;&#34;, gv);
	}
}

void
genasmsym(void (*put)(char*, int, vlong, vlong, int, Sym*))
{
	Prog *p;
	Auto *a;
	Sym *s;
	int h;

	s = lookup(&#34;etext&#34;, 0);
	if(s-&gt;type == STEXT)
		put(s-&gt;name, &#39;T&#39;, s-&gt;value, s-&gt;size, s-&gt;version, 0);

	for(h=0; h&lt;NHASH; h++) {
		for(s=hash[h]; s!=S; s=s-&gt;link) {
			switch(s-&gt;type) {
			case SCONST:
				if(!s-&gt;reachable)
					continue;
				put(s-&gt;name, &#39;D&#39;, s-&gt;value, s-&gt;size, s-&gt;version, s-&gt;gotype);
				continue;

			case SDATA:
				if(!s-&gt;reachable)
					continue;
				put(s-&gt;name, &#39;D&#39;, s-&gt;value+INITDAT, s-&gt;size, s-&gt;version, s-&gt;gotype);
				continue;

			case SMACHO:
				if(!s-&gt;reachable)
					continue;
				put(s-&gt;name, &#39;D&#39;, s-&gt;value+INITDAT+datsize+bsssize, s-&gt;size, s-&gt;version, s-&gt;gotype);
				continue;

			case SBSS:
				if(!s-&gt;reachable)
					continue;
				put(s-&gt;name, &#39;B&#39;, s-&gt;value+INITDAT, s-&gt;size, s-&gt;version, s-&gt;gotype);
				continue;

			case SFILE:
				put(s-&gt;name, &#39;f&#39;, s-&gt;value, 0, s-&gt;version, 0);
				continue;
			}
		}
	}

	for(p = textp; p != P; p = p-&gt;pcond) {
		s = p-&gt;from.sym;
		if(s-&gt;type != STEXT)
			continue;

		/* filenames first */
		for(a=p-&gt;to.autom; a; a=a-&gt;link)
			if(a-&gt;type == D_FILE)
				put(a-&gt;asym-&gt;name, &#39;z&#39;, a-&gt;aoffset, 0, 0, 0);
			else
			if(a-&gt;type == D_FILE1)
				put(a-&gt;asym-&gt;name, &#39;Z&#39;, a-&gt;aoffset, 0, 0, 0);

		if(!s-&gt;reachable)
			continue;
		put(s-&gt;name, &#39;T&#39;, s-&gt;value, s-&gt;size, s-&gt;version, s-&gt;gotype);

		/* frame, auto and param after */
		put(&#34;.frame&#34;, &#39;m&#39;, p-&gt;to.offset+8, 0, 0, 0);

		for(a=p-&gt;to.autom; a; a=a-&gt;link)
			if(a-&gt;type == D_AUTO)
				put(a-&gt;asym-&gt;name, &#39;a&#39;, -a-&gt;aoffset, 0, 0, a-&gt;gotype);
			else
			if(a-&gt;type == D_PARAM)
				put(a-&gt;asym-&gt;name, &#39;p&#39;, a-&gt;aoffset, 0, 0, a-&gt;gotype);
	}
	if(debug[&#39;v&#39;] || debug[&#39;n&#39;])
		Bprint(&amp;bso, &#34;symsize = %lud\n&#34;, symsize);
	Bflush(&amp;bso);
}

void
asmsym(void)
{
	genasmsym(putsymb);
}

char *elfstrdat;
int elfstrsize;
int maxelfstr;

int
putelfstr(char *s)
{
	int off, n;

	if(elfstrsize == 0 &amp;&amp; s[0] != 0) {
		// first entry must be empty string
		putelfstr(&#34;&#34;);
	}

	n = strlen(s)+1;
	if(elfstrsize+n &gt; maxelfstr) {
		maxelfstr = 2*(elfstrsize+n+(1&lt;&lt;20));
		elfstrdat = realloc(elfstrdat, maxelfstr);
	}
	off = elfstrsize;
	elfstrsize += n;
	memmove(elfstrdat+off, s, n);
	return off;
}

void
putelfsymb(char *s, int t, vlong addr, vlong size, int ver, Sym *go)
{
	int bind, type, shndx, stroff;
	
	bind = STB_GLOBAL;
	switch(t) {
	default:
		return;
	case &#39;T&#39;:
		type = STT_FUNC;
		shndx = elftextsh + 0;
		break;
	case &#39;D&#39;:
		type = STT_OBJECT;
		shndx = elftextsh + 1;
		break;
	case &#39;B&#39;:
		type = STT_OBJECT;
		shndx = elftextsh + 2;
		break;
	}
	
	stroff = putelfstr(s);
	lputl(stroff);	// string
	cput((bind&lt;&lt;4)|(type&amp;0xF));
	cput(0);
	wputl(shndx);
	vputl(addr);
	vputl(size);
}

void
asmelfsym(void)
{
	genasmsym(putelfsymb);
}

void
asmlc(void)
{
	vlong oldpc;
	Prog *p;
	int32 oldlc, v, s;

	oldpc = INITTEXT;
	oldlc = 0;
	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;line == oldlc || p-&gt;as == ATEXT || p-&gt;as == ANOP) {
			if(p-&gt;as == ATEXT)
				curtext = p;
			if(debug[&#39;O&#39;])
				Bprint(&amp;bso, &#34;%6llux %P\n&#34;,
					p-&gt;pc, p);
			continue;
		}
		if(debug[&#39;O&#39;])
			Bprint(&amp;bso, &#34;\t\t%6ld&#34;, lcsize);
		v = (p-&gt;pc - oldpc) / MINLC;
		while(v) {
			s = 127;
			if(v &lt; 127)
				s = v;
			cput(s+128);	/* 129-255 +pc */
			if(debug[&#39;O&#39;])
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
			if(debug[&#39;O&#39;]) {
				if(s &gt; 0)
					Bprint(&amp;bso, &#34; lc+%ld(%d,%ld)\n&#34;,
						s, 0, s);
				else
					Bprint(&amp;bso, &#34; lc%ld(%d,%ld)\n&#34;,
						s, 0, s);
				Bprint(&amp;bso, &#34;%6llux %P\n&#34;,
					p-&gt;pc, p);
			}
			lcsize += 5;
			continue;
		}
		if(s &gt; 0) {
			cput(0+s);	/* 1-64 +lc */
			if(debug[&#39;O&#39;]) {
				Bprint(&amp;bso, &#34; lc+%ld(%ld)\n&#34;, s, 0+s);
				Bprint(&amp;bso, &#34;%6llux %P\n&#34;,
					p-&gt;pc, p);
			}
		} else {
			cput(64-s);	/* 65-128 -lc */
			if(debug[&#39;O&#39;]) {
				Bprint(&amp;bso, &#34; lc%ld(%ld)\n&#34;, s, 64-s);
				Bprint(&amp;bso, &#34;%6llux %P\n&#34;,
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
	if(debug[&#39;v&#39;] || debug[&#39;O&#39;])
		Bprint(&amp;bso, &#34;lcsize = %ld\n&#34;, lcsize);
	Bflush(&amp;bso);
}

int
oclass(Adr *a)
{
	vlong v;
	int32 l;

	if(a-&gt;type &gt;= D_INDIR || a-&gt;index != D_NONE) {
		if(a-&gt;index != D_NONE &amp;&amp; a-&gt;scale == 0) {
			if(a-&gt;type == D_ADDR) {
				switch(a-&gt;index) {
				case D_EXTERN:
				case D_STATIC:
					return Yi32;	/* TO DO: Yi64 */
				case D_AUTO:
				case D_PARAM:
					return Yiauto;
				}
				return Yxxx;
			}
			return Ycol;
		}
		return Ym;
	}
	switch(a-&gt;type)
	{
	case D_AL:
		return Yal;

	case D_AX:
		return Yax;

/*
	case D_SPB:
*/
	case D_BPB:
	case D_SIB:
	case D_DIB:
	case D_R8B:
	case D_R9B:
	case D_R10B:
	case D_R11B:
	case D_R12B:
	case D_R13B:
	case D_R14B:
	case D_R15B:
		if(asmode != 64)
			return Yxxx;
	case D_DL:
	case D_BL:
	case D_AH:
	case D_CH:
	case D_DH:
	case D_BH:
		return Yrb;

	case D_CL:
		return Ycl;

	case D_CX:
		return Ycx;

	case D_DX:
	case D_BX:
		return Yrx;

	case D_R8:	/* not really Yrl */
	case D_R9:
	case D_R10:
	case D_R11:
	case D_R12:
	case D_R13:
	case D_R14:
	case D_R15:
		if(asmode != 64)
			return Yxxx;
	case D_SP:
	case D_BP:
	case D_SI:
	case D_DI:
		return Yrl;

	case D_F0+0:
		return	Yf0;

	case D_F0+1:
	case D_F0+2:
	case D_F0+3:
	case D_F0+4:
	case D_F0+5:
	case D_F0+6:
	case D_F0+7:
		return	Yrf;

	case D_M0+0:
	case D_M0+1:
	case D_M0+2:
	case D_M0+3:
	case D_M0+4:
	case D_M0+5:
	case D_M0+6:
	case D_M0+7:
		return	Ymr;

	case D_X0+0:
	case D_X0+1:
	case D_X0+2:
	case D_X0+3:
	case D_X0+4:
	case D_X0+5:
	case D_X0+6:
	case D_X0+7:
	case D_X0+8:
	case D_X0+9:
	case D_X0+10:
	case D_X0+11:
	case D_X0+12:
	case D_X0+13:
	case D_X0+14:
	case D_X0+15:
		return	Yxr;

	case D_NONE:
		return Ynone;

	case D_CS:	return	Ycs;
	case D_SS:	return	Yss;
	case D_DS:	return	Yds;
	case D_ES:	return	Yes;
	case D_FS:	return	Yfs;
	case D_GS:	return	Ygs;

	case D_GDTR:	return	Ygdtr;
	case D_IDTR:	return	Yidtr;
	case D_LDTR:	return	Yldtr;
	case D_MSW:	return	Ymsw;
	case D_TASK:	return	Ytask;

	case D_CR+0:	return	Ycr0;
	case D_CR+1:	return	Ycr1;
	case D_CR+2:	return	Ycr2;
	case D_CR+3:	return	Ycr3;
	case D_CR+4:	return	Ycr4;
	case D_CR+5:	return	Ycr5;
	case D_CR+6:	return	Ycr6;
	case D_CR+7:	return	Ycr7;
	case D_CR+8:	return	Ycr8;

	case D_DR+0:	return	Ydr0;
	case D_DR+1:	return	Ydr1;
	case D_DR+2:	return	Ydr2;
	case D_DR+3:	return	Ydr3;
	case D_DR+4:	return	Ydr4;
	case D_DR+5:	return	Ydr5;
	case D_DR+6:	return	Ydr6;
	case D_DR+7:	return	Ydr7;

	case D_TR+0:	return	Ytr0;
	case D_TR+1:	return	Ytr1;
	case D_TR+2:	return	Ytr2;
	case D_TR+3:	return	Ytr3;
	case D_TR+4:	return	Ytr4;
	case D_TR+5:	return	Ytr5;
	case D_TR+6:	return	Ytr6;
	case D_TR+7:	return	Ytr7;

	case D_EXTERN:
	case D_STATIC:
	case D_AUTO:
	case D_PARAM:
		return Ym;

	case D_CONST:
	case D_ADDR:
		if(a-&gt;sym == S) {
			v = a-&gt;offset;
			if(v == 0)
				return Yi0;
			if(v == 1)
				return Yi1;
			if(v &gt;= -128 &amp;&amp; v &lt;= 127)
				return Yi8;
			l = v;
			if((vlong)l == v)
				return Ys32;	/* can sign extend */
			if((v&gt;&gt;32) == 0)
				return Yi32;	/* unsigned */
			return Yi64;
		}
		return Yi32;	/* TO DO: D_ADDR as Yi64 */

	case D_BRANCH:
		return Ybr;
	}
	return Yxxx;
}

void
asmidx(Adr *a, int base)
{
	int i;

	switch(a-&gt;index) {
	default:
		goto bad;

	case D_NONE:
		i = 4 &lt;&lt; 3;
		goto bas;

	case D_R8:
	case D_R9:
	case D_R10:
	case D_R11:
	case D_R12:
	case D_R13:
	case D_R14:
	case D_R15:
		if(asmode != 64)
			goto bad;
	case D_AX:
	case D_CX:
	case D_DX:
	case D_BX:
	case D_BP:
	case D_SI:
	case D_DI:
		i = reg[(int)a-&gt;index] &lt;&lt; 3;
		break;
	}
	switch(a-&gt;scale) {
	default:
		goto bad;
	case 1:
		break;
	case 2:
		i |= (1&lt;&lt;6);
		break;
	case 4:
		i |= (2&lt;&lt;6);
		break;
	case 8:
		i |= (3&lt;&lt;6);
		break;
	}
bas:
	switch(base) {
	default:
		goto bad;
	case D_NONE:	/* must be mod=00 */
		i |= 5;
		break;
	case D_R8:
	case D_R9:
	case D_R10:
	case D_R11:
	case D_R12:
	case D_R13:
	case D_R14:
	case D_R15:
		if(asmode != 64)
			goto bad;
	case D_AX:
	case D_CX:
	case D_DX:
	case D_BX:
	case D_SP:
	case D_BP:
	case D_SI:
	case D_DI:
		i |= reg[base];
		break;
	}
	*andptr++ = i;
	return;
bad:
	diag(&#34;asmidx: bad address %D&#34;, a);
	*andptr++ = 0;
	return;
}

static void
put4(int32 v)
{
	if(dlm &amp;&amp; curp != P &amp;&amp; reloca != nil){
		dynreloc(reloca-&gt;sym, curp-&gt;pc + andptr - &amp;and[0], 1);
		reloca = nil;
	}
	andptr[0] = v;
	andptr[1] = v&gt;&gt;8;
	andptr[2] = v&gt;&gt;16;
	andptr[3] = v&gt;&gt;24;
	andptr += 4;
}

static void
put8(vlong v)
{
	if(dlm &amp;&amp; curp != P &amp;&amp; reloca != nil){
		dynreloc(reloca-&gt;sym, curp-&gt;pc + andptr - &amp;and[0], 1);	/* TO DO */
		reloca = nil;
	}
	andptr[0] = v;
	andptr[1] = v&gt;&gt;8;
	andptr[2] = v&gt;&gt;16;
	andptr[3] = v&gt;&gt;24;
	andptr[4] = v&gt;&gt;32;
	andptr[5] = v&gt;&gt;40;
	andptr[6] = v&gt;&gt;48;
	andptr[7] = v&gt;&gt;56;
	andptr += 8;
}

vlong
symaddr(Sym *s)
{
	Adr a;

	a.type = D_ADDR;
	a.index = D_EXTERN;
	a.offset = 0;
	a.sym = s;
	return vaddr(&amp;a);
}

vlong
vaddr(Adr *a)
{
	int t;
	vlong v;
	Sym *s;

	t = a-&gt;type;
	v = a-&gt;offset;
	if(t == D_ADDR)
		t = a-&gt;index;
	switch(t) {
	case D_STATIC:
	case D_EXTERN:
		s = a-&gt;sym;
		if(s != nil) {
			if(dlm &amp;&amp; curp != P)
				reloca = a;
			switch(s-&gt;type) {
			case SUNDEF:
				ckoff(s, v);
			case STEXT:
			case SCONST:
				if(!s-&gt;reachable)
					diag(&#34;unreachable symbol in vaddr - %s&#34;, s-&gt;name);
				if((uvlong)s-&gt;value &lt; (uvlong)INITTEXT)
					v += INITTEXT;	/* TO DO */
				v += s-&gt;value;
				break;
			case SMACHO:
				if(!s-&gt;reachable)
					sysfatal(&#34;unreachable symbol in vaddr - %s&#34;, s-&gt;name);
				v += INITDAT + datsize + s-&gt;value;
				break;
			default:
				if(!s-&gt;reachable)
					diag(&#34;unreachable symbol in vaddr - %s&#34;, s-&gt;name);
				v += INITDAT + s-&gt;value;
			}
		}
	}
	return v;
}

static void
asmandsz(Adr *a, int r, int rex, int m64)
{
	int32 v;
	int t;
	Adr aa;

	rex &amp;= (0x40 | Rxr);
	v = a-&gt;offset;
	t = a-&gt;type;
	if(a-&gt;index != D_NONE) {
		if(t &gt;= D_INDIR) {
			t -= D_INDIR;
			rexflag |= (regrex[(int)a-&gt;index] &amp; Rxx) | (regrex[t] &amp; Rxb) | rex;
			if(t == D_NONE) {
				*andptr++ = (0 &lt;&lt; 6) | (4 &lt;&lt; 0) | (r &lt;&lt; 3);
				asmidx(a, t);
				put4(v);
				return;
			}
			if(v == 0 &amp;&amp; t != D_BP &amp;&amp; t != D_R13) {
				*andptr++ = (0 &lt;&lt; 6) | (4 &lt;&lt; 0) | (r &lt;&lt; 3);
				asmidx(a, t);
				return;
			}
			if(v &gt;= -128 &amp;&amp; v &lt; 128) {
				*andptr++ = (1 &lt;&lt; 6) | (4 &lt;&lt; 0) | (r &lt;&lt; 3);
				asmidx(a, t);
				*andptr++ = v;
				return;
			}
			*andptr++ = (2 &lt;&lt; 6) | (4 &lt;&lt; 0) | (r &lt;&lt; 3);
			asmidx(a, t);
			put4(v);
			return;
		}
		switch(t) {
		default:
			goto bad;
		case D_STATIC:
		case D_EXTERN:
			aa.type = D_NONE+D_INDIR;
			break;
		case D_AUTO:
		case D_PARAM:
			aa.type = D_SP+D_INDIR;
			break;
		}
		aa.offset = vaddr(a);
		aa.index = a-&gt;index;
		aa.scale = a-&gt;scale;
		asmandsz(&amp;aa, r, rex, m64);
		return;
	}
	if(t &gt;= D_AL &amp;&amp; t &lt;= D_X0+15) {
		if(v)
			goto bad;
		*andptr++ = (3 &lt;&lt; 6) | (reg[t] &lt;&lt; 0) | (r &lt;&lt; 3);
		rexflag |= (regrex[t] &amp; (0x40 | Rxb)) | rex;
		return;
	}
	if(t &gt;= D_INDIR) {
		t -= D_INDIR;
		rexflag |= (regrex[t] &amp; Rxb) | rex;
		if(t == D_NONE) {
			if(asmode != 64){
				*andptr++ = (0 &lt;&lt; 6) | (5 &lt;&lt; 0) | (r &lt;&lt; 3);
				put4(v);
				return;
			}
			/* temporary */
			*andptr++ = (0 &lt;&lt;  6) | (4 &lt;&lt; 0) | (r &lt;&lt; 3);	/* sib present */
			*andptr++ = (0 &lt;&lt; 6) | (4 &lt;&lt; 3) | (5 &lt;&lt; 0);	/* DS:d32 */
			put4(v);
			return;
		}
		if(t == D_SP || t == D_R12) {
			if(v == 0) {
				*andptr++ = (0 &lt;&lt; 6) | (reg[t] &lt;&lt; 0) | (r &lt;&lt; 3);
				asmidx(a, t);
				return;
			}
			if(v &gt;= -128 &amp;&amp; v &lt; 128) {
				*andptr++ = (1 &lt;&lt; 6) | (reg[t] &lt;&lt; 0) | (r &lt;&lt; 3);
				asmidx(a, t);
				*andptr++ = v;
				return;
			}
			*andptr++ = (2 &lt;&lt; 6) | (reg[t] &lt;&lt; 0) | (r &lt;&lt; 3);
			asmidx(a, t);
			put4(v);
			return;
		}
		if(t &gt;= D_AX &amp;&amp; t &lt;= D_R15) {
			if(v == 0 &amp;&amp; t != D_BP &amp;&amp; t != D_R13) {
				*andptr++ = (0 &lt;&lt; 6) | (reg[t] &lt;&lt; 0) | (r &lt;&lt; 3);
				return;
			}
			if(v &gt;= -128 &amp;&amp; v &lt; 128) {
				andptr[0] = (1 &lt;&lt; 6) | (reg[t] &lt;&lt; 0) | (r &lt;&lt; 3);
				andptr[1] = v;
				andptr += 2;
				return;
			}
			*andptr++ = (2 &lt;&lt; 6) | (reg[t] &lt;&lt; 0) | (r &lt;&lt; 3);
			put4(v);
			return;
		}
		goto bad;
	}
	switch(a-&gt;type) {
	default:
		goto bad;
	case D_STATIC:
	case D_EXTERN:
		aa.type = D_NONE+D_INDIR;
		break;
	case D_AUTO:
	case D_PARAM:
		aa.type = D_SP+D_INDIR;
		break;
	}
	aa.index = D_NONE;
	aa.scale = 1;
	aa.offset = vaddr(a);
	asmandsz(&amp;aa, r, rex, m64);
	return;
bad:
	diag(&#34;asmand: bad address %D&#34;, a);
	return;
}

void
asmand(Adr *a, Adr *ra)
{
	asmandsz(a, reg[ra-&gt;type], regrex[ra-&gt;type], 0);
}

void
asmando(Adr *a, int o)
{
	asmandsz(a, o, 0, 0);
}

static void
bytereg(Adr *a, char *t)
{
	if(a-&gt;index == D_NONE &amp;&amp; (a-&gt;type &gt;= D_AX &amp;&amp; a-&gt;type &lt;= D_R15)) {
		a-&gt;type = D_AL + (a-&gt;type-D_AX);
		*t = 0;
	}
}

#define	E	0xff
Movtab	ymovtab[] =
{
/* push */
	{APUSHL,	Ycs,	Ynone,	0,	0x0e,E,0,0},
	{APUSHL,	Yss,	Ynone,	0,	0x16,E,0,0},
	{APUSHL,	Yds,	Ynone,	0,	0x1e,E,0,0},
	{APUSHL,	Yes,	Ynone,	0,	0x06,E,0,0},
	{APUSHL,	Yfs,	Ynone,	0,	0x0f,0xa0,E,0},
	{APUSHL,	Ygs,	Ynone,	0,	0x0f,0xa8,E,0},
	{APUSHQ,	Yfs,	Ynone,	0,	0x0f,0xa0,E,0},
	{APUSHQ,	Ygs,	Ynone,	0,	0x0f,0xa8,E,0},

	{APUSHW,	Ycs,	Ynone,	0,	Pe,0x0e,E,0},
	{APUSHW,	Yss,	Ynone,	0,	Pe,0x16,E,0},
	{APUSHW,	Yds,	Ynone,	0,	Pe,0x1e,E,0},
	{APUSHW,	Yes,	Ynone,	0,	Pe,0x06,E,0},
	{APUSHW,	Yfs,	Ynone,	0,	Pe,0x0f,0xa0,E},
	{APUSHW,	Ygs,	Ynone,	0,	Pe,0x0f,0xa8,E},

/* pop */
	{APOPL,	Ynone,	Yds,	0,	0x1f,E,0,0},
	{APOPL,	Ynone,	Yes,	0,	0x07,E,0,0},
	{APOPL,	Ynone,	Yss,	0,	0x17,E,0,0},
	{APOPL,	Ynone,	Yfs,	0,	0x0f,0xa1,E,0},
	{APOPL,	Ynone,	Ygs,	0,	0x0f,0xa9,E,0},
	{APOPQ,	Ynone,	Yfs,	0,	0x0f,0xa1,E,0},
	{APOPQ,	Ynone,	Ygs,	0,	0x0f,0xa9,E,0},

	{APOPW,	Ynone,	Yds,	0,	Pe,0x1f,E,0},
	{APOPW,	Ynone,	Yes,	0,	Pe,0x07,E,0},
	{APOPW,	Ynone,	Yss,	0,	Pe,0x17,E,0},
	{APOPW,	Ynone,	Yfs,	0,	Pe,0x0f,0xa1,E},
	{APOPW,	Ynone,	Ygs,	0,	Pe,0x0f,0xa9,E},

/* mov seg */
	{AMOVW,	Yes,	Yml,	1,	0x8c,0,0,0},
	{AMOVW,	Ycs,	Yml,	1,	0x8c,1,0,0},
	{AMOVW,	Yss,	Yml,	1,	0x8c,2,0,0},
	{AMOVW,	Yds,	Yml,	1,	0x8c,3,0,0},
	{AMOVW,	Yfs,	Yml,	1,	0x8c,4,0,0},
	{AMOVW,	Ygs,	Yml,	1,	0x8c,5,0,0},

	{AMOVW,	Yml,	Yes,	2,	0x8e,0,0,0},
	{AMOVW,	Yml,	Ycs,	2,	0x8e,1,0,0},
	{AMOVW,	Yml,	Yss,	2,	0x8e,2,0,0},
	{AMOVW,	Yml,	Yds,	2,	0x8e,3,0,0},
	{AMOVW,	Yml,	Yfs,	2,	0x8e,4,0,0},
	{AMOVW,	Yml,	Ygs,	2,	0x8e,5,0,0},

/* mov cr */
	{AMOVL,	Ycr0,	Yml,	3,	0x0f,0x20,0,0},
	{AMOVL,	Ycr2,	Yml,	3,	0x0f,0x20,2,0},
	{AMOVL,	Ycr3,	Yml,	3,	0x0f,0x20,3,0},
	{AMOVL,	Ycr4,	Yml,	3,	0x0f,0x20,4,0},
	{AMOVL,	Ycr8,	Yml,	3,	0x0f,0x20,8,0},
	{AMOVQ,	Ycr0,	Yml,	3,	0x0f,0x20,0,0},
	{AMOVQ,	Ycr2,	Yml,	3,	0x0f,0x20,2,0},
	{AMOVQ,	Ycr3,	Yml,	3,	0x0f,0x20,3,0},
	{AMOVQ,	Ycr4,	Yml,	3,	0x0f,0x20,4,0},
	{AMOVQ,	Ycr8,	Yml,	3,	0x0f,0x20,8,0},

	{AMOVL,	Yml,	Ycr0,	4,	0x0f,0x22,0,0},
	{AMOVL,	Yml,	Ycr2,	4,	0x0f,0x22,2,0},
	{AMOVL,	Yml,	Ycr3,	4,	0x0f,0x22,3,0},
	{AMOVL,	Yml,	Ycr4,	4,	0x0f,0x22,4,0},
	{AMOVL,	Yml,	Ycr8,	4,	0x0f,0x22,8,0},
	{AMOVQ,	Yml,	Ycr0,	4,	0x0f,0x22,0,0},
	{AMOVQ,	Yml,	Ycr2,	4,	0x0f,0x22,2,0},
	{AMOVQ,	Yml,	Ycr3,	4,	0x0f,0x22,3,0},
	{AMOVQ,	Yml,	Ycr4,	4,	0x0f,0x22,4,0},
	{AMOVQ,	Yml,	Ycr8,	4,	0x0f,0x22,8,0},

/* mov dr */
	{AMOVL,	Ydr0,	Yml,	3,	0x0f,0x21,0,0},
	{AMOVL,	Ydr6,	Yml,	3,	0x0f,0x21,6,0},
	{AMOVL,	Ydr7,	Yml,	3,	0x0f,0x21,7,0},
	{AMOVQ,	Ydr0,	Yml,	3,	0x0f,0x21,0,0},
	{AMOVQ,	Ydr6,	Yml,	3,	0x0f,0x21,6,0},
	{AMOVQ,	Ydr7,	Yml,	3,	0x0f,0x21,7,0},

	{AMOVL,	Yml,	Ydr0,	4,	0x0f,0x23,0,0},
	{AMOVL,	Yml,	Ydr6,	4,	0x0f,0x23,6,0},
	{AMOVL,	Yml,	Ydr7,	4,	0x0f,0x23,7,0},
	{AMOVQ,	Yml,	Ydr0,	4,	0x0f,0x23,0,0},
	{AMOVQ,	Yml,	Ydr6,	4,	0x0f,0x23,6,0},
	{AMOVQ,	Yml,	Ydr7,	4,	0x0f,0x23,7,0},

/* mov tr */
	{AMOVL,	Ytr6,	Yml,	3,	0x0f,0x24,6,0},
	{AMOVL,	Ytr7,	Yml,	3,	0x0f,0x24,7,0},

	{AMOVL,	Yml,	Ytr6,	4,	0x0f,0x26,6,E},
	{AMOVL,	Yml,	Ytr7,	4,	0x0f,0x26,7,E},

/* lgdt, sgdt, lidt, sidt */
	{AMOVL,	Ym,	Ygdtr,	4,	0x0f,0x01,2,0},
	{AMOVL,	Ygdtr,	Ym,	3,	0x0f,0x01,0,0},
	{AMOVL,	Ym,	Yidtr,	4,	0x0f,0x01,3,0},
	{AMOVL,	Yidtr,	Ym,	3,	0x0f,0x01,1,0},
	{AMOVQ,	Ym,	Ygdtr,	4,	0x0f,0x01,2,0},
	{AMOVQ,	Ygdtr,	Ym,	3,	0x0f,0x01,0,0},
	{AMOVQ,	Ym,	Yidtr,	4,	0x0f,0x01,3,0},
	{AMOVQ,	Yidtr,	Ym,	3,	0x0f,0x01,1,0},

/* lldt, sldt */
	{AMOVW,	Yml,	Yldtr,	4,	0x0f,0x00,2,0},
	{AMOVW,	Yldtr,	Yml,	3,	0x0f,0x00,0,0},

/* lmsw, smsw */
	{AMOVW,	Yml,	Ymsw,	4,	0x0f,0x01,6,0},
	{AMOVW,	Ymsw,	Yml,	3,	0x0f,0x01,4,0},

/* ltr, str */
	{AMOVW,	Yml,	Ytask,	4,	0x0f,0x00,3,0},
	{AMOVW,	Ytask,	Yml,	3,	0x0f,0x00,1,0},

/* load full pointer */
	{AMOVL,	Yml,	Ycol,	5,	0,0,0,0},
	{AMOVW,	Yml,	Ycol,	5,	Pe,0,0,0},

/* double shift */
	{ASHLL,	Ycol,	Yml,	6,	0xa4,0xa5,0,0},
	{ASHRL,	Ycol,	Yml,	6,	0xac,0xad,0,0},
	{ASHLQ,	Ycol,	Yml,	6,	Pw,0xa4,0xa5,0},
	{ASHRQ,	Ycol,	Yml,	6,	Pw,0xac,0xad,0},
	{ASHLW,	Ycol,	Yml,	6,	Pe,0xa4,0xa5,0},
	{ASHRW,	Ycol,	Yml,	6,	Pe,0xac,0xad,0},
	0
};

int
isax(Adr *a)
{

	switch(a-&gt;type) {
	case D_AX:
	case D_AL:
	case D_AH:
	case D_INDIR+D_AX:
		return 1;
	}
	if(a-&gt;index == D_AX)
		return 1;
	return 0;
}

void
subreg(Prog *p, int from, int to)
{

	if(debug[&#39;Q&#39;])
		print(&#34;\n%P	s/%R/%R/\n&#34;, p, from, to);

	if(p-&gt;from.type == from)
		p-&gt;from.type = to;
	if(p-&gt;to.type == from)
		p-&gt;to.type = to;

	if(p-&gt;from.index == from)
		p-&gt;from.index = to;
	if(p-&gt;to.index == from)
		p-&gt;to.index = to;

	from += D_INDIR;
	if(p-&gt;from.type == from)
		p-&gt;from.type = to+D_INDIR;
	if(p-&gt;to.type == from)
		p-&gt;to.type = to+D_INDIR;

	if(debug[&#39;Q&#39;])
		print(&#34;%P\n&#34;, p);
}

static int
mediaop(Optab *o, int op, int osize, int z)
{
	switch(op){
	case Pm:
	case Pe:
	case Pf2:
	case Pf3:
		if(osize != 1){
			if(op != Pm)
				*andptr++ = op;
			*andptr++ = Pm;
			op = o-&gt;op[++z];
			break;
		}
	default:
		if(andptr == and || andptr[-1] != Pm)
			*andptr++ = Pm;
		break;
	}
	*andptr++ = op;
	return z;
}

void
doasm(Prog *p)
{
	Optab *o;
	Prog *q, pp;
	uchar *t;
	Movtab *mo;
	int z, op, ft, tt, xo, l;
	vlong v;

	o = opindex[p-&gt;as];
	if(o == nil) {
		diag(&#34;asmins: missing op %P&#34;, p);
		return;
	}

	if(p-&gt;ft == 0)
		p-&gt;ft = oclass(&amp;p-&gt;from);
	if(p-&gt;tt == 0)
		p-&gt;tt = oclass(&amp;p-&gt;to);

	ft = p-&gt;ft * Ymax;
	tt = p-&gt;tt * Ymax;

	t = o-&gt;ytab;
	if(t == 0) {
		diag(&#34;asmins: noproto %P&#34;, p);
		return;
	}
	xo = o-&gt;op[0] == 0x0f;
	for(z=0; *t; z+=t[3]+xo,t+=4)
		if(ycover[ft+t[0]])
		if(ycover[tt+t[1]])
			goto found;
	goto domov;

found:
	switch(o-&gt;prefix) {
	case Pq:	/* 16 bit escape and opcode escape */
		*andptr++ = Pe;
		*andptr++ = Pm;
		break;

	case Pf2:	/* xmm opcode escape */
	case Pf3:
		*andptr++ = o-&gt;prefix;
		*andptr++ = Pm;
		break;

	case Pm:	/* opcode escape */
		*andptr++ = Pm;
		break;

	case Pe:	/* 16 bit escape */
		*andptr++ = Pe;
		break;

	case Pw:	/* 64-bit escape */
		if(p-&gt;mode != 64)
			diag(&#34;asmins: illegal 64: %P&#34;, p);
		rexflag |= Pw;
		break;

	case Pb:	/* botch */
		bytereg(&amp;p-&gt;from, &amp;p-&gt;ft);
		bytereg(&amp;p-&gt;to, &amp;p-&gt;tt);
		break;

	case P32:	/* 32 bit but illegal if 64-bit mode */
		if(p-&gt;mode == 64)
			diag(&#34;asmins: illegal in 64-bit mode: %P&#34;, p);
		break;

	case Py:	/* 64-bit only, no prefix */
		if(p-&gt;mode != 64)
			diag(&#34;asmins: illegal in %d-bit mode: %P&#34;, p-&gt;mode, p);
		break;
	}
	v = vaddr(&amp;p-&gt;from);
	op = o-&gt;op[z];
	if(op == 0x0f) {
		*andptr++ = op;
		op = o-&gt;op[++z];
	}
	switch(t[2]) {
	default:
		diag(&#34;asmins: unknown z %d %P&#34;, t[2], p);
		return;

	case Zpseudo:
		break;

	case Zlit:
		for(; op = o-&gt;op[z]; z++)
			*andptr++ = op;
		break;

	case Zmb_r:
		bytereg(&amp;p-&gt;from, &amp;p-&gt;ft);
		/* fall through */
	case Zm_r:
		*andptr++ = op;
		asmand(&amp;p-&gt;from, &amp;p-&gt;to);
		break;

	case Zm_r_xm:
		mediaop(o, op, t[3], z);
		asmand(&amp;p-&gt;from, &amp;p-&gt;to);
		break;

	case Zm_r_xm_nr:
		rexflag = 0;
		mediaop(o, op, t[3], z);
		asmand(&amp;p-&gt;from, &amp;p-&gt;to);
		break;

	case Zm_r_i_xm:
		mediaop(o, op, t[3], z);
		asmand(&amp;p-&gt;from, &amp;p-&gt;to);
		*andptr++ = p-&gt;to.offset;
		break;

	case Zm_r_3d:
		*andptr++ = 0x0f;
		*andptr++ = 0x0f;
		asmand(&amp;p-&gt;from, &amp;p-&gt;to);
		*andptr++ = op;
		break;

	case Zibm_r:
		*andptr++ = op;
		asmand(&amp;p-&gt;from, &amp;p-&gt;to);
		*andptr++ = p-&gt;to.offset;
		break;

	case Zaut_r:
		*andptr++ = 0x8d;	/* leal */
		if(p-&gt;from.type != D_ADDR)
			diag(&#34;asmins: Zaut sb type ADDR&#34;);
		p-&gt;from.type = p-&gt;from.index;
		p-&gt;from.index = D_NONE;
		asmand(&amp;p-&gt;from, &amp;p-&gt;to);
		p-&gt;from.index = p-&gt;from.type;
		p-&gt;from.type = D_ADDR;
		break;

	case Zm_o:
		*andptr++ = op;
		asmando(&amp;p-&gt;from, o-&gt;op[z+1]);
		break;

	case Zr_m:
		*andptr++ = op;
		asmand(&amp;p-&gt;to, &amp;p-&gt;from);
		break;

	case Zr_m_xm:
		mediaop(o, op, t[3], z);
		asmand(&amp;p-&gt;to, &amp;p-&gt;from);
		break;

	case Zr_m_xm_nr:
		rexflag = 0;
		mediaop(o, op, t[3], z);
		asmand(&amp;p-&gt;to, &amp;p-&gt;from);
		break;

	case Zr_m_i_xm:
		mediaop(o, op, t[3], z);
		asmand(&amp;p-&gt;to, &amp;p-&gt;from);
		*andptr++ = p-&gt;from.offset;
		break;

	case Zo_m:
		*andptr++ = op;
		asmando(&amp;p-&gt;to, o-&gt;op[z+1]);
		break;

	case Zo_m64:
		*andptr++ = op;
		asmandsz(&amp;p-&gt;to, o-&gt;op[z+1], 0, 1);
		break;

	case Zm_ibo:
		v = vaddr(&amp;p-&gt;to);
		*andptr++ = op;
		asmando(&amp;p-&gt;from, o-&gt;op[z+1]);
		*andptr++ = v;
		break;

	case Zibo_m:
		*andptr++ = op;
		asmando(&amp;p-&gt;to, o-&gt;op[z+1]);
		*andptr++ = v;
		break;

	case Zibo_m_xm:
		z = mediaop(o, op, t[3], z);
		asmando(&amp;p-&gt;to, o-&gt;op[z+1]);
		*andptr++ = v;
		break;

	case Z_ib:
		v = vaddr(&amp;p-&gt;to);
	case Zib_:
		*andptr++ = op;
		*andptr++ = v;
		break;

	case Zib_rp:
		rexflag |= regrex[p-&gt;to.type] &amp; (Rxb|0x40);
		*andptr++ = op + reg[p-&gt;to.type];
		*andptr++ = v;
		break;

	case Zil_rp:
		rexflag |= regrex[p-&gt;to.type] &amp; Rxb;
		*andptr++ = op + reg[p-&gt;to.type];
		if(o-&gt;prefix == Pe) {
			*andptr++ = v;
			*andptr++ = v&gt;&gt;8;
		}
		else
			put4(v);
		break;

	case Zo_iw:
		*andptr++ = op;
		if(p-&gt;from.type != D_NONE){
			*andptr++ = v;
			*andptr++ = v&gt;&gt;8;
		}
		break;

	case Ziq_rp:
		l = v&gt;&gt;32;
		if(l == 0){
			//p-&gt;mark |= 0100;
			//print(&#34;zero: %llux %P\n&#34;, v, p);
			rexflag &amp;= ~(0x40|Rxw);
			rexflag |= regrex[p-&gt;to.type] &amp; Rxb;
			*andptr++ = 0xb8 + reg[p-&gt;to.type];
			put4(v);
		}else if(l == -1 &amp;&amp; (v&amp;((uvlong)1&lt;&lt;31))!=0){	/* sign extend */
			//p-&gt;mark |= 0100;
			//print(&#34;sign: %llux %P\n&#34;, v, p);
			*andptr ++ = 0xc7;
			asmando(&amp;p-&gt;to, 0);
			put4(v);
		}else{	/* need all 8 */
			//print(&#34;all: %llux %P\n&#34;, v, p);
			rexflag |= regrex[p-&gt;to.type] &amp; Rxb;
			*andptr++ = op + reg[p-&gt;to.type];
			put8(v);
		}
		break;

	case Zib_rr:
		*andptr++ = op;
		asmand(&amp;p-&gt;to, &amp;p-&gt;to);
		*andptr++ = v;
		break;

	case Z_il:
		v = vaddr(&amp;p-&gt;to);
	case Zil_:
		*andptr++ = op;
		if(o-&gt;prefix == Pe) {
			*andptr++ = v;
			*andptr++ = v&gt;&gt;8;
		}
		else
			put4(v);
		break;

	case Zm_ilo:
		v = vaddr(&amp;p-&gt;to);
		*andptr++ = op;
		asmando(&amp;p-&gt;from, o-&gt;op[z+1]);
		if(o-&gt;prefix == Pe) {
			*andptr++ = v;
			*andptr++ = v&gt;&gt;8;
		}
		else
			put4(v);
		break;

	case Zilo_m:
		*andptr++ = op;
		asmando(&amp;p-&gt;to, o-&gt;op[z+1]);
		if(o-&gt;prefix == Pe) {
			*andptr++ = v;
			*andptr++ = v&gt;&gt;8;
		}
		else
			put4(v);
		break;

	case Zil_rr:
		*andptr++ = op;
		asmand(&amp;p-&gt;to, &amp;p-&gt;to);
		if(o-&gt;prefix == Pe) {
			*andptr++ = v;
			*andptr++ = v&gt;&gt;8;
		}
		else
			put4(v);
		break;

	case Z_rp:
		rexflag |= regrex[p-&gt;to.type] &amp; (Rxb|0x40);
		*andptr++ = op + reg[p-&gt;to.type];
		break;

	case Zrp_:
		rexflag |= regrex[p-&gt;from.type] &amp; (Rxb|0x40);
		*andptr++ = op + reg[p-&gt;from.type];
		break;

	case Zclr:
		*andptr++ = op;
		asmand(&amp;p-&gt;to, &amp;p-&gt;to);
		break;

	case Zbr:
		q = p-&gt;pcond;
		if(q) {
			v = q-&gt;pc - p-&gt;pc - 2;
			if(v &gt;= -128 &amp;&amp; v &lt;= 127) {
				*andptr++ = op;
				*andptr++ = v;
			} else {
				v -= 6-2;
				*andptr++ = 0x0f;
				*andptr++ = o-&gt;op[z+1];
				*andptr++ = v;
				*andptr++ = v&gt;&gt;8;
				*andptr++ = v&gt;&gt;16;
				*andptr++ = v&gt;&gt;24;
			}
		}
		break;

	case Zcall:
		q = p-&gt;pcond;
		if(q) {
			v = q-&gt;pc - p-&gt;pc - 5;
			if(dlm &amp;&amp; curp != P &amp;&amp; p-&gt;to.sym-&gt;type == SUNDEF){
				/* v = 0 - p-&gt;pc - 5; */
				v = 0;
				ckoff(p-&gt;to.sym, v);
				v += p-&gt;to.sym-&gt;value;
				dynreloc(p-&gt;to.sym, p-&gt;pc+1, 0);
			}
			*andptr++ = op;
			*andptr++ = v;
			*andptr++ = v&gt;&gt;8;
			*andptr++ = v&gt;&gt;16;
			*andptr++ = v&gt;&gt;24;
		}
		break;

	case Zjmp:
		q = p-&gt;pcond;
		if(q) {
			v = q-&gt;pc - p-&gt;pc - 2;
			if(v &gt;= -128 &amp;&amp; v &lt;= 127) {
				*andptr++ = op;
				*andptr++ = v;
			} else {
				v -= 5-2;
				*andptr++ = o-&gt;op[z+1];
				*andptr++ = v;
				*andptr++ = v&gt;&gt;8;
				*andptr++ = v&gt;&gt;16;
				*andptr++ = v&gt;&gt;24;
			}
		}
		break;

	case Zloop:
		q = p-&gt;pcond;
		if(q) {
			v = q-&gt;pc - p-&gt;pc - 2;
			if(v &lt; -128 &amp;&amp; v &gt; 127)
				diag(&#34;loop too far: %P&#34;, p);
			*andptr++ = op;
			*andptr++ = v;
		}
		break;

	case Zbyte:
		*andptr++ = v;
		if(op &gt; 1) {
			*andptr++ = v&gt;&gt;8;
			if(op &gt; 2) {
				*andptr++ = v&gt;&gt;16;
				*andptr++ = v&gt;&gt;24;
				if(op &gt; 4) {
					*andptr++ = v&gt;&gt;32;
					*andptr++ = v&gt;&gt;40;
					*andptr++ = v&gt;&gt;48;
					*andptr++ = v&gt;&gt;56;
				}
			}
		}
		break;
	}
	return;

domov:
	for(mo=ymovtab; mo-&gt;as; mo++)
		if(p-&gt;as == mo-&gt;as)
		if(ycover[ft+mo-&gt;ft])
		if(ycover[tt+mo-&gt;tt]){
			t = mo-&gt;op;
			goto mfound;
		}
bad:
	if(p-&gt;mode != 64){
		/*
		 * here, the assembly has failed.
		 * if its a byte instruction that has
		 * unaddressable registers, try to
		 * exchange registers and reissue the
		 * instruction with the operands renamed.
		 */
		pp = *p;
		z = p-&gt;from.type;
		if(z &gt;= D_BP &amp;&amp; z &lt;= D_DI) {
			if(isax(&amp;p-&gt;to)) {
				*andptr++ = 0x87;			/* xchg lhs,bx */
				asmando(&amp;p-&gt;from, reg[D_BX]);
				subreg(&amp;pp, z, D_BX);
				doasm(&amp;pp);
				*andptr++ = 0x87;			/* xchg lhs,bx */
				asmando(&amp;p-&gt;from, reg[D_BX]);
			} else {
				*andptr++ = 0x90 + reg[z];		/* xchg lsh,ax */
				subreg(&amp;pp, z, D_AX);
				doasm(&amp;pp);
				*andptr++ = 0x90 + reg[z];		/* xchg lsh,ax */
			}
			return;
		}
		z = p-&gt;to.type;
		if(z &gt;= D_BP &amp;&amp; z &lt;= D_DI) {
			if(isax(&amp;p-&gt;from)) {
				*andptr++ = 0x87;			/* xchg rhs,bx */
				asmando(&amp;p-&gt;to, reg[D_BX]);
				subreg(&amp;pp, z, D_BX);
				doasm(&amp;pp);
				*andptr++ = 0x87;			/* xchg rhs,bx */
				asmando(&amp;p-&gt;to, reg[D_BX]);
			} else {
				*andptr++ = 0x90 + reg[z];		/* xchg rsh,ax */
				subreg(&amp;pp, z, D_AX);
				doasm(&amp;pp);
				*andptr++ = 0x90 + reg[z];		/* xchg rsh,ax */
			}
			return;
		}
	}
	diag(&#34;doasm: notfound from=%ux to=%ux %P&#34;, p-&gt;from.type, p-&gt;to.type, p);
	return;

mfound:
	switch(mo-&gt;code) {
	default:
		diag(&#34;asmins: unknown mov %d %P&#34;, mo-&gt;code, p);
		break;

	case 0:	/* lit */
		for(z=0; t[z]!=E; z++)
			*andptr++ = t[z];
		break;

	case 1:	/* r,m */
		*andptr++ = t[0];
		asmando(&amp;p-&gt;to, t[1]);
		break;

	case 2:	/* m,r */
		*andptr++ = t[0];
		asmando(&amp;p-&gt;from, t[1]);
		break;

	case 3:	/* r,m - 2op */
		*andptr++ = t[0];
		*andptr++ = t[1];
		asmando(&amp;p-&gt;to, t[2]);
		rexflag |= regrex[p-&gt;from.type] &amp; (Rxr|0x40);
		break;

	case 4:	/* m,r - 2op */
		*andptr++ = t[0];
		*andptr++ = t[1];
		asmando(&amp;p-&gt;from, t[2]);
		rexflag |= regrex[p-&gt;to.type] &amp; (Rxr|0x40);
		break;

	case 5:	/* load full pointer, trash heap */
		if(t[0])
			*andptr++ = t[0];
		switch(p-&gt;to.index) {
		default:
			goto bad;
		case D_DS:
			*andptr++ = 0xc5;
			break;
		case D_SS:
			*andptr++ = 0x0f;
			*andptr++ = 0xb2;
			break;
		case D_ES:
			*andptr++ = 0xc4;
			break;
		case D_FS:
			*andptr++ = 0x0f;
			*andptr++ = 0xb4;
			break;
		case D_GS:
			*andptr++ = 0x0f;
			*andptr++ = 0xb5;
			break;
		}
		asmand(&amp;p-&gt;from, &amp;p-&gt;to);
		break;

	case 6:	/* double shift */
		if(t[0] == Pw){
			if(p-&gt;mode != 64)
				diag(&#34;asmins: illegal 64: %P&#34;, p);
			rexflag |= Pw;
			t++;
		}else if(t[0] == Pe){
			*andptr++ = Pe;
			t++;
		}
		z = p-&gt;from.type;
		switch(z) {
		default:
			goto bad;
		case D_CONST:
			*andptr++ = 0x0f;
			*andptr++ = t[0];
			asmandsz(&amp;p-&gt;to, reg[(int)p-&gt;from.index], regrex[(int)p-&gt;from.index], 0);
			*andptr++ = p-&gt;from.offset;
			break;
		case D_CL:
		case D_CX:
			*andptr++ = 0x0f;
			*andptr++ = t[1];
			asmandsz(&amp;p-&gt;to, reg[(int)p-&gt;from.index], regrex[(int)p-&gt;from.index], 0);
			break;
		}
		break;
	}
}

void
asmins(Prog *p)
{
	int n, np, c;

	rexflag = 0;
	andptr = and;
	asmode = p-&gt;mode;
	doasm(p);
	if(rexflag){
		/*
		 * as befits the whole approach of the architecture,
		 * the rex prefix must appear before the first opcode byte
		 * (and thus after any 66/67/f2/f3 prefix bytes, but
		 * before the 0f opcode escape!), or it might be ignored.
		 * note that the handbook often misleadingly shows 66/f2/f3 in `opcode&#39;.
		 */
		if(p-&gt;mode != 64)
			diag(&#34;asmins: illegal in mode %d: %P&#34;, p-&gt;mode, p);
		n = andptr - and;
		for(np = 0; np &lt; n; np++) {
			c = and[np];
			if(c != 0x66 &amp;&amp; c != 0xf2 &amp;&amp; c != 0xf3 &amp;&amp; c != 0x67)
				break;
		}
		memmove(and+np+1, and+np, n-np);
		and[np] = 0x40 | rexflag;
		andptr++;
	}
}

enum{
	ABSD = 0,
	ABSU = 1,
	RELD = 2,
	RELU = 3,
};

int modemap[4] = { 0, 1, -1, 2, };

typedef struct Reloc Reloc;

struct Reloc
{
	int n;
	int t;
	uchar *m;
	uint32 *a;
};

Reloc rels;

static void
grow(Reloc *r)
{
	int t;
	uchar *m, *nm;
	uint32 *a, *na;

	t = r-&gt;t;
	r-&gt;t += 64;
	m = r-&gt;m;
	a = r-&gt;a;
	r-&gt;m = nm = mal(r-&gt;t*sizeof(uchar));
	r-&gt;a = na = mal(r-&gt;t*sizeof(uint32));
	memmove(nm, m, t*sizeof(uchar));
	memmove(na, a, t*sizeof(uint32));
	free(m);
	free(a);
}

void
dynreloc(Sym *s, uint32 v, int abs)
{
	int i, k, n;
	uchar *m;
	uint32 *a;
	Reloc *r;

	if(s-&gt;type == SUNDEF)
		k = abs ? ABSU : RELU;
	else
		k = abs ? ABSD : RELD;
	/* Bprint(&amp;bso, &#34;R %s a=%ld(%lx) %d\n&#34;, s-&gt;name, v, v, k); */
	k = modemap[k];
	r = &amp;rels;
	n = r-&gt;n;
	if(n &gt;= r-&gt;t)
		grow(r);
	m = r-&gt;m;
	a = r-&gt;a;
	for(i = n; i &gt; 0; i--){
		if(v &lt; a[i-1]){	/* happens occasionally for data */
			m[i] = m[i-1];
			a[i] = a[i-1];
		}
		else
			break;
	}
	m[i] = k;
	a[i] = v;
	r-&gt;n++;
}

static int
sput(char *s)
{
	char *p;

	p = s;
	while(*s)
		cput(*s++);
	cput(0);
	return s-p+1;
}

void
asmdyn()
{
	int i, n, t, c;
	Sym *s;
	uint32 la, ra, *a;
	vlong off;
	uchar *m;
	Reloc *r;

	cflush();
	off = seek(cout, 0, 1);
	lputb(0);
	t = 0;
	lputb(imports);
	t += 4;
	for(i = 0; i &lt; NHASH; i++)
		for(s = hash[i]; s != S; s = s-&gt;link)
			if(s-&gt;type == SUNDEF){
				lputb(s-&gt;sig);
				t += 4;
				t += sput(s-&gt;name);
			}

	la = 0;
	r = &amp;rels;
	n = r-&gt;n;
	m = r-&gt;m;
	a = r-&gt;a;
	lputb(n);
	t += 4;
	for(i = 0; i &lt; n; i++){
		ra = *a-la;
		if(*a &lt; la)
			diag(&#34;bad relocation order&#34;);
		if(ra &lt; 256)
			c = 0;
		else if(ra &lt; 65536)
			c = 1;
		else
			c = 2;
		cput((c&lt;&lt;6)|*m++);
		t++;
		if(c == 0){
			cput(ra);
			t++;
		}
		else if(c == 1){
			wputb(ra);
			t += 2;
		}
		else{
			lputb(ra);
			t += 4;
		}
		la = *a++;
	}

	cflush();
	seek(cout, off, 0);
	lputb(t);

	if(debug[&#39;v&#39;]){
		Bprint(&amp;bso, &#34;import table entries = %d\n&#34;, imports);
		Bprint(&amp;bso, &#34;export table entries = %d\n&#34;, exports);
	}
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
