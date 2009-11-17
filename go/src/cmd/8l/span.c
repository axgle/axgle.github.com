<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8l/span.c</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/8l/span.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/8l/span.c
// http://code.google.com/p/inferno-os/source/browse/utils/8l/span.c
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

void
span(void)
{
	Prog *p, *q;
	int32 v, c, idat;
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
			p-&gt;as = AADDL;
			if(v &lt; 0) {
				p-&gt;as = ASUBL;
				v = -v;
				p-&gt;from.offset = v;
			}
			if(v == 0)
				p-&gt;as = ANOP;
		}
	}

	n = 0;
start:
	do{
		again = 0;
		if(debug[&#39;v&#39;])
			Bprint(&amp;bso, &#34;%5.2f span %d\n&#34;, cputime(), n);
		Bflush(&amp;bso);
		if(n &gt; 500) {
			// TODO(rsc): figure out why nacl takes so long to converge.
			print(&#34;span must be looping - %d\n&#34;, textsize);
			errorexit();
		}
		c = INITTEXT;
		for(p = firstp; p != P; p = p-&gt;link) {
			if(p-&gt;as == ATEXT) {
				curtext = p;
				if(HEADTYPE == 8)
					c = (c+31)&amp;~31;
			}
			if(p-&gt;to.type == D_BRANCH)
				if(p-&gt;back)
					p-&gt;pc = c;
			if(n == 0 || HEADTYPE == 8 || p-&gt;to.type == D_BRANCH) {
				if(HEADTYPE == 8)
					p-&gt;pc = c;
				asmins(p);
				m = andptr-and;
				if(p-&gt;mark != m)
					again = 1;
				p-&gt;mark = m;
			}
			if(HEADTYPE == 8) {
				c = p-&gt;pc + p-&gt;mark;
			} else {
				p-&gt;pc = c;
				c += p-&gt;mark;
			}
		}
		textsize = c;
		n++;
	}while(again);

	if(INITRND) {
		INITDAT = rnd(c, INITRND);
		if(INITDAT != idat) {
			idat = INITDAT;
			goto start;
		}
	}
	xdefine(&#34;etext&#34;, STEXT, c);
	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;etext = %lux\n&#34;, c);
	Bflush(&amp;bso);
	for(p = textp; p != P; p = p-&gt;pcond)
		p-&gt;from.sym-&gt;value = p-&gt;pc;
	textsize = c - INITTEXT;
}

void
xdefine(char *p, int t, int32 v)
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
putsymb(char *s, int t, int32 v, int ver, Sym *go)
{
	int i, f;
	vlong gv;

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
	gv = 0;
	if(go) {
		if(!go-&gt;reachable)
			sysfatal(&#34;unreachable type %s&#34;, go-&gt;name);
		gv = go-&gt;value+INITDAT;
	}
	lput(gv);

	symsize += 4 + 1 + i+1 + 4;

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
			Bprint(&amp;bso, &#34;%c %.8lux %s&lt;%d&gt; %s (%.8llux)\n&#34;, t, v, s, ver, go ? go-&gt;name : &#34;&#34;, gv);
		else
			Bprint(&amp;bso, &#34;%c %.8lux %s\n&#34;, t, v, s, go ? go-&gt;name : &#34;&#34;, gv);
	}
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
		putsymb(s-&gt;name, &#39;T&#39;, s-&gt;value, s-&gt;version, 0);

	for(h=0; h&lt;NHASH; h++)
		for(s=hash[h]; s!=S; s=s-&gt;link)
			switch(s-&gt;type) {
			case SCONST:
				if(!s-&gt;reachable)
					continue;
				putsymb(s-&gt;name, &#39;D&#39;, s-&gt;value, s-&gt;version, s-&gt;gotype);
				continue;

			case SDATA:
				if(!s-&gt;reachable)
					continue;
				putsymb(s-&gt;name, &#39;D&#39;, s-&gt;value+INITDAT, s-&gt;version, s-&gt;gotype);
				continue;

			case SMACHO:
				if(!s-&gt;reachable)
					continue;
				putsymb(s-&gt;name, &#39;D&#39;, s-&gt;value+INITDAT+datsize+bsssize, s-&gt;version, s-&gt;gotype);
				continue;

			case SBSS:
				if(!s-&gt;reachable)
					continue;
				putsymb(s-&gt;name, &#39;B&#39;, s-&gt;value+INITDAT, s-&gt;version, s-&gt;gotype);
				continue;

			case SFILE:
				putsymb(s-&gt;name, &#39;f&#39;, s-&gt;value, s-&gt;version, 0);
				continue;
			}

	for(p=textp; p!=P; p=p-&gt;pcond) {
		s = p-&gt;from.sym;
		if(s-&gt;type != STEXT)
			continue;

		/* filenames first */
		for(a=p-&gt;to.autom; a; a=a-&gt;link)
			if(a-&gt;type == D_FILE)
				putsymb(a-&gt;asym-&gt;name, &#39;z&#39;, a-&gt;aoffset, 0, 0);
			else
			if(a-&gt;type == D_FILE1)
				putsymb(a-&gt;asym-&gt;name, &#39;Z&#39;, a-&gt;aoffset, 0, 0);

		if(!s-&gt;reachable)
			continue;

		putsymb(s-&gt;name, &#39;T&#39;, s-&gt;value, s-&gt;version, s-&gt;gotype);

		/* frame, auto and param after */
		putsymb(&#34;.frame&#34;, &#39;m&#39;, p-&gt;to.offset+4, 0, 0);

		for(a=p-&gt;to.autom; a; a=a-&gt;link)
			if(a-&gt;type == D_AUTO)
				putsymb(a-&gt;asym-&gt;name, &#39;a&#39;, -a-&gt;aoffset, 0, a-&gt;gotype);
			else
			if(a-&gt;type == D_PARAM)
				putsymb(a-&gt;asym-&gt;name, &#39;p&#39;, a-&gt;aoffset, 0, a-&gt;gotype);
	}
	if(debug[&#39;v&#39;] || debug[&#39;n&#39;])
		Bprint(&amp;bso, &#34;symsize = %lud\n&#34;, symsize);
	Bflush(&amp;bso);
}

void
asmlc(void)
{
	int32 oldpc, oldlc;
	Prog *p;
	int32 v, s;

	oldpc = INITTEXT;
	oldlc = 0;
	for(p = firstp; p != P; p = p-&gt;link) {
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

int
prefixof(Adr *a)
{
	switch(a-&gt;type) {
	case D_INDIR+D_CS:
		return 0x2e;
	case D_INDIR+D_DS:
		return 0x3e;
	case D_INDIR+D_ES:
		return 0x26;
	case D_INDIR+D_FS:
		return 0x64;
	case D_INDIR+D_GS:
		return 0x65;
	}
	return 0;
}

int
oclass(Adr *a)
{
	int32 v;

	if((a-&gt;type &gt;= D_INDIR &amp;&amp; a-&gt;type &lt; 2*D_INDIR) || a-&gt;index != D_NONE) {
		if(a-&gt;index != D_NONE &amp;&amp; a-&gt;scale == 0) {
			if(a-&gt;type == D_ADDR) {
				switch(a-&gt;index) {
				case D_EXTERN:
				case D_STATIC:
					return Yi32;
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

	case D_CL:
	case D_DL:
	case D_BL:
	case D_AH:
	case D_CH:
	case D_DH:
	case D_BH:
		return Yrb;

	case D_CX:
		return Ycx;

	case D_DX:
	case D_BX:
		return Yrx;

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
	case D_CONST2:
	case D_ADDR:
		if(a-&gt;sym == S) {
			v = a-&gt;offset;
			if(v == 0)
				return Yi0;
			if(v == 1)
				return Yi1;
			if(v &gt;= -128 &amp;&amp; v &lt;= 127)
				return Yi8;
		}
		return Yi32;

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

	case D_AX:
	case D_CX:
	case D_DX:
	case D_BX:
	case D_BP:
	case D_SI:
	case D_DI:
		i = reg[a-&gt;index] &lt;&lt; 3;
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

int32
symaddr(Sym *s)
{
	Adr a;

	a.type = D_ADDR;
	a.index = D_EXTERN;
	a.offset = 0;
	a.sym = s;
	return vaddr(&amp;a);
}

int32
vaddr(Adr *a)
{
	int t;
	int32 v;
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
					sysfatal(&#34;unreachable symbol in vaddr - %s&#34;, s-&gt;name);
				v += s-&gt;value;
				break;
			case SMACHO:
				if(!s-&gt;reachable)
					sysfatal(&#34;unreachable symbol in vaddr - %s&#34;, s-&gt;name);
				v += INITDAT + datsize + s-&gt;value;
				break;
			default:
				if(!s-&gt;reachable)
					sysfatal(&#34;unreachable symbol in vaddr - %s&#34;, s-&gt;name);
				v += INITDAT + s-&gt;value;
			}
		}
	}
	return v;
}

void
asmand(Adr *a, int r)
{
	int32 v;
	int t;
	Adr aa;

	v = a-&gt;offset;
	t = a-&gt;type;
	if(a-&gt;index != D_NONE) {
		if(t &gt;= D_INDIR &amp;&amp; t &lt; 2*D_INDIR) {
			t -= D_INDIR;
			if(t == D_NONE) {
				*andptr++ = (0 &lt;&lt; 6) | (4 &lt;&lt; 0) | (r &lt;&lt; 3);
				asmidx(a, t);
				put4(v);
				return;
			}
			if(v == 0 &amp;&amp; t != D_BP) {
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
		asmand(&amp;aa, r);
		return;
	}
	if(t &gt;= D_AL &amp;&amp; t &lt;= D_F0+7) {
		if(v)
			goto bad;
		*andptr++ = (3 &lt;&lt; 6) | (reg[t] &lt;&lt; 0) | (r &lt;&lt; 3);
		return;
	}
	if(t &gt;= D_INDIR &amp;&amp; t &lt; 2*D_INDIR) {
		t -= D_INDIR;
		if(t == D_NONE || (D_CS &lt;= t &amp;&amp; t &lt;= D_GS)) {
			*andptr++ = (0 &lt;&lt; 6) | (5 &lt;&lt; 0) | (r &lt;&lt; 3);
			put4(v);
			return;
		}
		if(t == D_SP) {
			if(v == 0) {
				*andptr++ = (0 &lt;&lt; 6) | (4 &lt;&lt; 0) | (r &lt;&lt; 3);
				asmidx(a, D_SP);
				return;
			}
			if(v &gt;= -128 &amp;&amp; v &lt; 128) {
				*andptr++ = (1 &lt;&lt; 6) | (4 &lt;&lt; 0) | (r &lt;&lt; 3);
				asmidx(a, D_SP);
				*andptr++ = v;
				return;
			}
			*andptr++ = (2 &lt;&lt; 6) | (4 &lt;&lt; 0) | (r &lt;&lt; 3);
			asmidx(a, D_SP);
			put4(v);
			return;
		}
		if(t &gt;= D_AX &amp;&amp; t &lt;= D_DI) {
			if(v == 0 &amp;&amp; t != D_BP) {
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
	asmand(&amp;aa, r);
	return;
bad:
	diag(&#34;asmand: bad address %D&#34;, a);
	return;
}

#define	E	0xff
uchar	ymovtab[] =
{
/* push */
	APUSHL,	Ycs,	Ynone,	0,	0x0e,E,0,0,
	APUSHL,	Yss,	Ynone,	0,	0x16,E,0,0,
	APUSHL,	Yds,	Ynone,	0,	0x1e,E,0,0,
	APUSHL,	Yes,	Ynone,	0,	0x06,E,0,0,
	APUSHL,	Yfs,	Ynone,	0,	0x0f,0xa0,E,0,
	APUSHL,	Ygs,	Ynone,	0,	0x0f,0xa8,E,0,

	APUSHW,	Ycs,	Ynone,	0,	Pe,0x0e,E,0,
	APUSHW,	Yss,	Ynone,	0,	Pe,0x16,E,0,
	APUSHW,	Yds,	Ynone,	0,	Pe,0x1e,E,0,
	APUSHW,	Yes,	Ynone,	0,	Pe,0x06,E,0,
	APUSHW,	Yfs,	Ynone,	0,	Pe,0x0f,0xa0,E,
	APUSHW,	Ygs,	Ynone,	0,	Pe,0x0f,0xa8,E,

/* pop */
	APOPL,	Ynone,	Yds,	0,	0x1f,E,0,0,
	APOPL,	Ynone,	Yes,	0,	0x07,E,0,0,
	APOPL,	Ynone,	Yss,	0,	0x17,E,0,0,
	APOPL,	Ynone,	Yfs,	0,	0x0f,0xa1,E,0,
	APOPL,	Ynone,	Ygs,	0,	0x0f,0xa9,E,0,

	APOPW,	Ynone,	Yds,	0,	Pe,0x1f,E,0,
	APOPW,	Ynone,	Yes,	0,	Pe,0x07,E,0,
	APOPW,	Ynone,	Yss,	0,	Pe,0x17,E,0,
	APOPW,	Ynone,	Yfs,	0,	Pe,0x0f,0xa1,E,
	APOPW,	Ynone,	Ygs,	0,	Pe,0x0f,0xa9,E,

/* mov seg */
	AMOVW,	Yes,	Yml,	1,	0x8c,0,0,0,
	AMOVW,	Ycs,	Yml,	1,	0x8c,1,0,0,
	AMOVW,	Yss,	Yml,	1,	0x8c,2,0,0,
	AMOVW,	Yds,	Yml,	1,	0x8c,3,0,0,
	AMOVW,	Yfs,	Yml,	1,	0x8c,4,0,0,
	AMOVW,	Ygs,	Yml,	1,	0x8c,5,0,0,

	AMOVW,	Yml,	Yes,	2,	0x8e,0,0,0,
	AMOVW,	Yml,	Ycs,	2,	0x8e,1,0,0,
	AMOVW,	Yml,	Yss,	2,	0x8e,2,0,0,
	AMOVW,	Yml,	Yds,	2,	0x8e,3,0,0,
	AMOVW,	Yml,	Yfs,	2,	0x8e,4,0,0,
	AMOVW,	Yml,	Ygs,	2,	0x8e,5,0,0,

/* mov cr */
	AMOVL,	Ycr0,	Yml,	3,	0x0f,0x20,0,0,
	AMOVL,	Ycr2,	Yml,	3,	0x0f,0x20,2,0,
	AMOVL,	Ycr3,	Yml,	3,	0x0f,0x20,3,0,
	AMOVL,	Ycr4,	Yml,	3,	0x0f,0x20,4,0,

	AMOVL,	Yml,	Ycr0,	4,	0x0f,0x22,0,0,
	AMOVL,	Yml,	Ycr2,	4,	0x0f,0x22,2,0,
	AMOVL,	Yml,	Ycr3,	4,	0x0f,0x22,3,0,
	AMOVL,	Yml,	Ycr4,	4,	0x0f,0x22,4,0,

/* mov dr */
	AMOVL,	Ydr0,	Yml,	3,	0x0f,0x21,0,0,
	AMOVL,	Ydr6,	Yml,	3,	0x0f,0x21,6,0,
	AMOVL,	Ydr7,	Yml,	3,	0x0f,0x21,7,0,

	AMOVL,	Yml,	Ydr0,	4,	0x0f,0x23,0,0,
	AMOVL,	Yml,	Ydr6,	4,	0x0f,0x23,6,0,
	AMOVL,	Yml,	Ydr7,	4,	0x0f,0x23,7,0,

/* mov tr */
	AMOVL,	Ytr6,	Yml,	3,	0x0f,0x24,6,0,
	AMOVL,	Ytr7,	Yml,	3,	0x0f,0x24,7,0,

	AMOVL,	Yml,	Ytr6,	4,	0x0f,0x26,6,E,
	AMOVL,	Yml,	Ytr7,	4,	0x0f,0x26,7,E,

/* lgdt, sgdt, lidt, sidt */
	AMOVL,	Ym,	Ygdtr,	4,	0x0f,0x01,2,0,
	AMOVL,	Ygdtr,	Ym,	3,	0x0f,0x01,0,0,
	AMOVL,	Ym,	Yidtr,	4,	0x0f,0x01,3,0,
	AMOVL,	Yidtr,	Ym,	3,	0x0f,0x01,1,0,

/* lldt, sldt */
	AMOVW,	Yml,	Yldtr,	4,	0x0f,0x00,2,0,
	AMOVW,	Yldtr,	Yml,	3,	0x0f,0x00,0,0,

/* lmsw, smsw */
	AMOVW,	Yml,	Ymsw,	4,	0x0f,0x01,6,0,
	AMOVW,	Ymsw,	Yml,	3,	0x0f,0x01,4,0,

/* ltr, str */
	AMOVW,	Yml,	Ytask,	4,	0x0f,0x00,3,0,
	AMOVW,	Ytask,	Yml,	3,	0x0f,0x00,1,0,

/* load full pointer */
	AMOVL,	Yml,	Ycol,	5,	0,0,0,0,
	AMOVW,	Yml,	Ycol,	5,	Pe,0,0,0,

/* double shift */
	ASHLL,	Ycol,	Yml,	6,	0xa4,0xa5,0,0,
	ASHRL,	Ycol,	Yml,	6,	0xac,0xad,0,0,

/* extra imul */
	AIMULW,	Yml,	Yrl,	7,	Pq,0xaf,0,0,
	AIMULL,	Yml,	Yrl,	7,	Pm,0xaf,0,0,
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

	if(p-&gt;from.type == from) {
		p-&gt;from.type = to;
		p-&gt;ft = 0;
	}
	if(p-&gt;to.type == from) {
		p-&gt;to.type = to;
		p-&gt;tt = 0;
	}

	if(p-&gt;from.index == from) {
		p-&gt;from.index = to;
		p-&gt;ft = 0;
	}
	if(p-&gt;to.index == from) {
		p-&gt;to.index = to;
		p-&gt;tt = 0;
	}

	from += D_INDIR;
	if(p-&gt;from.type == from) {
		p-&gt;from.type = to+D_INDIR;
		p-&gt;ft = 0;
	}
	if(p-&gt;to.type == from) {
		p-&gt;to.type = to+D_INDIR;
		p-&gt;tt = 0;
	}

	if(debug[&#39;Q&#39;])
		print(&#34;%P\n&#34;, p);
}

// nacl RET:
//	POPL BX
//	ANDL BX, $~31
//	JMP BX
uchar naclret[] = { 0x5b, 0x83, 0xe3, ~31, 0xff, 0xe3 };

// nacl JMP BX:
//	ANDL BX, $~31
//	JMP BX
uchar nacljmpbx[] = { 0x83, 0xe3, ~31, 0xff, 0xe3 };

// nacl CALL BX:
//	ANDL BX, $~31
//	CALL BX
uchar naclcallbx[] = { 0x83, 0xe3, ~31, 0xff, 0xd3 };

void
doasm(Prog *p)
{
	Optab *o;
	Prog *q, pp;
	uchar *t;
	int z, op, ft, tt;
	int32 v, pre;

	pre = prefixof(&amp;p-&gt;from);
	if(pre)
		*andptr++ = pre;
	pre = prefixof(&amp;p-&gt;to);
	if(pre)
		*andptr++ = pre;

	if(p-&gt;ft == 0)
		p-&gt;ft = oclass(&amp;p-&gt;from);
	if(p-&gt;tt == 0)
		p-&gt;tt = oclass(&amp;p-&gt;to);

	ft = p-&gt;ft * Ymax;
	tt = p-&gt;tt * Ymax;
	o = &amp;optab[p-&gt;as];
	t = o-&gt;ytab;
	if(t == 0) {
		diag(&#34;asmins: noproto %P&#34;, p);
		return;
	}
	for(z=0; *t; z+=t[3],t+=4)
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

	case Pm:	/* opcode escape */
		*andptr++ = Pm;
		break;

	case Pe:	/* 16 bit escape */
		*andptr++ = Pe;
		break;

	case Pb:	/* botch */
		break;
	}
	v = vaddr(&amp;p-&gt;from);
	op = o-&gt;op[z];
	switch(t[2]) {
	default:
		diag(&#34;asmins: unknown z %d %P&#34;, t[2], p);
		return;

	case Zpseudo:
		break;

	case Zlit:
		if(HEADTYPE == 8 &amp;&amp; p-&gt;as == ARET) {
			// native client return.
			for(z=0; z&lt;sizeof(naclret); z++)
				*andptr++ = naclret[z];
			break;
		}
		for(; op = o-&gt;op[z]; z++)
			*andptr++ = op;
		break;

	case Zm_r:
		*andptr++ = op;
		asmand(&amp;p-&gt;from, reg[p-&gt;to.type]);
		break;

	case Zaut_r:
		*andptr++ = 0x8d;	/* leal */
		if(p-&gt;from.type != D_ADDR)
			diag(&#34;asmins: Zaut sb type ADDR&#34;);
		p-&gt;from.type = p-&gt;from.index;
		p-&gt;from.index = D_NONE;
		p-&gt;ft = 0;
		asmand(&amp;p-&gt;from, reg[p-&gt;to.type]);
		p-&gt;from.index = p-&gt;from.type;
		p-&gt;from.type = D_ADDR;
		p-&gt;ft = 0;
		break;

	case Zm_o:
		*andptr++ = op;
		asmand(&amp;p-&gt;from, o-&gt;op[z+1]);
		break;

	case Zr_m:
		*andptr++ = op;
		asmand(&amp;p-&gt;to, reg[p-&gt;from.type]);
		break;

	case Zo_m:
		if(HEADTYPE == 8) {
			Adr a;

			switch(p-&gt;as) {
			case AJMP:
				if(p-&gt;to.type &lt; D_AX || p-&gt;to.type &gt; D_DI)
					diag(&#34;indirect jmp must use register in native client&#34;);
				// ANDL $~31, REG
				*andptr++ = 0x83;
				asmand(&amp;p-&gt;to, 04);
				*andptr++ = ~31;
				// JMP REG
				*andptr++ = 0xFF;
				asmand(&amp;p-&gt;to, 04);
				return;

			case ACALL:
				a = p-&gt;to;
				// native client indirect call
				if(a.type &lt; D_AX || a.type &gt; D_DI) {
					// MOVL target into BX
					*andptr++ = 0x8b;
					asmand(&amp;p-&gt;to, reg[D_BX]);
					memset(&amp;a, 0, sizeof a);
					a.type = D_BX;
				}
				// ANDL $~31, REG
				*andptr++ = 0x83;
				asmand(&amp;a, 04);
				*andptr++ = ~31;
				// CALL REG
				*andptr++ = 0xFF;
				asmand(&amp;a, 02);
				return;
			}
		}
		*andptr++ = op;
		asmand(&amp;p-&gt;to, o-&gt;op[z+1]);
		break;

	case Zm_ibo:
		v = vaddr(&amp;p-&gt;to);
		*andptr++ = op;
		asmand(&amp;p-&gt;from, o-&gt;op[z+1]);
		*andptr++ = v;
		break;

	case Zibo_m:
		*andptr++ = op;
		asmand(&amp;p-&gt;to, o-&gt;op[z+1]);
		*andptr++ = v;
		break;

	case Z_ib:
		v = vaddr(&amp;p-&gt;to);
	case Zib_:
		if(HEADTYPE == 8 &amp;&amp; p-&gt;as == AINT &amp;&amp; v == 3) {
			// native client disallows all INT instructions.
			// translate INT $3 to HLT.
			*andptr++ = 0xf4;
			break;
		}
		*andptr++ = op;
		*andptr++ = v;
		break;

	case Zib_rp:
		*andptr++ = op + reg[p-&gt;to.type];
		*andptr++ = v;
		break;

	case Zil_rp:
		*andptr++ = op + reg[p-&gt;to.type];
		if(o-&gt;prefix == Pe) {
			*andptr++ = v;
			*andptr++ = v&gt;&gt;8;
		}
		else
			put4(v);
		break;

	case Zib_rr:
		*andptr++ = op;
		asmand(&amp;p-&gt;to, reg[p-&gt;to.type]);
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
		asmand(&amp;p-&gt;from, o-&gt;op[z+1]);
		if(o-&gt;prefix == Pe) {
			*andptr++ = v;
			*andptr++ = v&gt;&gt;8;
		}
		else
			put4(v);
		break;

	case Zilo_m:
		*andptr++ = op;
		asmand(&amp;p-&gt;to, o-&gt;op[z+1]);
		if(o-&gt;prefix == Pe) {
			*andptr++ = v;
			*andptr++ = v&gt;&gt;8;
		}
		else
			put4(v);
		break;

	case Zil_rr:
		*andptr++ = op;
		asmand(&amp;p-&gt;to, reg[p-&gt;to.type]);
		if(o-&gt;prefix == Pe) {
			*andptr++ = v;
			*andptr++ = v&gt;&gt;8;
		}
		else
			put4(v);
		break;

	case Z_rp:
		*andptr++ = op + reg[p-&gt;to.type];
		break;

	case Zrp_:
		*andptr++ = op + reg[p-&gt;from.type];
		break;

	case Zclr:
		*andptr++ = op;
		asmand(&amp;p-&gt;to, reg[p-&gt;to.type]);
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

	case Zcallcon:
		v = p-&gt;to.offset - p-&gt;pc - 5;
		*andptr++ = op;
		*andptr++ = v;
		*andptr++ = v&gt;&gt;8;
		*andptr++ = v&gt;&gt;16;
		*andptr++ = v&gt;&gt;24;
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

	case Zjmpcon:
		v = p-&gt;to.offset - p-&gt;pc - 5;
		*andptr++ = o-&gt;op[z+1];
		*andptr++ = v;
		*andptr++ = v&gt;&gt;8;
		*andptr++ = v&gt;&gt;16;
		*andptr++ = v&gt;&gt;24;
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
			}
		}
		break;

	case Zmov:
		goto domov;
	}
	return;

domov:
	for(t=ymovtab; *t; t+=8)
		if(p-&gt;as == t[0])
		if(ycover[ft+t[1]])
		if(ycover[tt+t[2]])
			goto mfound;
bad:
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
			asmand(&amp;p-&gt;from, reg[D_BX]);
			subreg(&amp;pp, z, D_BX);
			doasm(&amp;pp);
			*andptr++ = 0x87;			/* xchg lhs,bx */
			asmand(&amp;p-&gt;from, reg[D_BX]);
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
			asmand(&amp;p-&gt;to, reg[D_BX]);
			subreg(&amp;pp, z, D_BX);
			doasm(&amp;pp);
			*andptr++ = 0x87;			/* xchg rhs,bx */
			asmand(&amp;p-&gt;to, reg[D_BX]);
		} else {
			*andptr++ = 0x90 + reg[z];		/* xchg rsh,ax */
			subreg(&amp;pp, z, D_AX);
			doasm(&amp;pp);
			*andptr++ = 0x90 + reg[z];		/* xchg rsh,ax */
		}
		return;
	}
	diag(&#34;doasm: notfound t2=%lux from=%lux to=%lux %P&#34;, t[2], p-&gt;from.type, p-&gt;to.type, p);
	return;

mfound:
	switch(t[3]) {
	default:
		diag(&#34;asmins: unknown mov %d %P&#34;, t[3], p);
		break;

	case 0:	/* lit */
		for(z=4; t[z]!=E; z++)
			*andptr++ = t[z];
		break;

	case 1:	/* r,m */
		*andptr++ = t[4];
		asmand(&amp;p-&gt;to, t[5]);
		break;

	case 2:	/* m,r */
		*andptr++ = t[4];
		asmand(&amp;p-&gt;from, t[5]);
		break;

	case 3:	/* r,m - 2op */
		*andptr++ = t[4];
		*andptr++ = t[5];
		asmand(&amp;p-&gt;to, t[6]);
		break;

	case 4:	/* m,r - 2op */
		*andptr++ = t[4];
		*andptr++ = t[5];
		asmand(&amp;p-&gt;from, t[6]);
		break;

	case 5:	/* load full pointer, trash heap */
		if(t[4])
			*andptr++ = t[4];
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
		asmand(&amp;p-&gt;from, reg[p-&gt;to.type]);
		break;

	case 6:	/* double shift */
		z = p-&gt;from.type;
		switch(z) {
		default:
			goto bad;
		case D_CONST:
			*andptr++ = 0x0f;
			*andptr++ = t[4];
			asmand(&amp;p-&gt;to, reg[p-&gt;from.index]);
			*andptr++ = p-&gt;from.offset;
			break;
		case D_CL:
		case D_CX:
			*andptr++ = 0x0f;
			*andptr++ = t[5];
			asmand(&amp;p-&gt;to, reg[p-&gt;from.index]);
			break;
		}
		break;

	case 7: /* imul rm,r */
		if(t[4] == Pq) {
			*andptr++ = Pe;
			*andptr++ = Pm;
		} else
			*andptr++ = t[4];
		*andptr++ = t[5];
		asmand(&amp;p-&gt;from, reg[p-&gt;to.type]);
		break;
	}
}

void
asmins(Prog *p)
{
	if(HEADTYPE == 8) {
		ulong npc;
		static Prog *prefix;

		// native client
		// - pad indirect jump targets (aka ATEXT) to 32-byte boundary
		// - instructions cannot cross 32-byte boundary
		// - end of call (return address) must be on 32-byte boundary
		if(p-&gt;as == ATEXT)
			p-&gt;pc += 31 &amp; -p-&gt;pc;
		if(p-&gt;as == ACALL) {
			// must end on 32-byte boundary.
			// doasm to find out how long the CALL encoding is.
			andptr = and;
			doasm(p);
			npc = p-&gt;pc + (andptr - and);
			p-&gt;pc += 31 &amp; -npc;
		}
		if(p-&gt;as == AREP || p-&gt;as == AREPN) {
			// save prefix for next instruction,
			// so that inserted NOPs do not split (e.g.) REP / MOVSL sequence.
			prefix = p;
			andptr = and;
			return;
		}
		andptr = and;
		if(prefix)
			doasm(prefix);
		doasm(p);
		npc = p-&gt;pc + (andptr - and);
		if(andptr &gt; and &amp;&amp; (p-&gt;pc&amp;~31) != ((npc-1)&amp;~31)) {
			// crossed 32-byte boundary; pad to boundary and try again
			p-&gt;pc += 31 &amp; -p-&gt;pc;
			andptr = and;
			if(prefix)
				doasm(prefix);
			doasm(p);
		}
		prefix = nil;
	} else {
		andptr = and;
		doasm(p);
	}
	if(andptr &gt; and+sizeof and) {
		print(&#34;and[] is too short - %d byte instruction\n&#34;, andptr - and);
		errorexit();
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
	lput(0);
	t = 0;
	lput(imports);
	t += 4;
	for(i = 0; i &lt; NHASH; i++)
		for(s = hash[i]; s != S; s = s-&gt;link)
			if(s-&gt;type == SUNDEF){
				lput(s-&gt;sig);
				t += 4;
				t += sput(s-&gt;name);
			}

	la = 0;
	r = &amp;rels;
	n = r-&gt;n;
	m = r-&gt;m;
	a = r-&gt;a;
	lput(n);
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
			wput(ra);
			t += 2;
		}
		else{
			lput(ra);
			t += 4;
		}
		la = *a++;
	}

	cflush();
	seek(cout, off, 0);
	lput(t);

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
