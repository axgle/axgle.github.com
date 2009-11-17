<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/libmach/5db.c</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/libmach/5db.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno libmach/5db.c
// http://code.google.com/p/inferno-os/source/browse/utils/libmach/5db.c
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

#include &lt;u.h&gt;
#include &lt;libc.h&gt;
#include &lt;bio.h&gt;
#include &#34;ureg_arm.h&#34;
#include &lt;mach.h&gt;

static int debug = 0;

#define	BITS(a, b)	((1&lt;&lt;(b+1))-(1&lt;&lt;a))

#define LSR(v, s)	((ulong)(v) &gt;&gt; (s))
#define ASR(v, s)	((long)(v) &gt;&gt; (s))
#define ROR(v, s)	(LSR((v), (s)) | (((v) &amp; ((1 &lt;&lt; (s))-1)) &lt;&lt; (32 - (s))))



typedef struct	Instr	Instr;
struct	Instr
{
	Map	*map;
	ulong	w;
	ulong	addr;
	uchar	op;			/* super opcode */

	uchar	cond;			/* bits 28-31 */
	uchar	store;			/* bit 20 */

	uchar	rd;			/* bits 12-15 */
	uchar	rn;			/* bits 16-19 */
	uchar	rs;			/* bits 0-11 (shifter operand) */

	long	imm;			/* rotated imm */
	char*	curr;			/* fill point in buffer */
	char*	end;			/* end of buffer */
	char*	err;			/* error message */
};

typedef struct Opcode Opcode;
struct Opcode
{
	char*	o;
	void	(*fmt)(Opcode*, Instr*);
	ulong	(*foll)(Map*, Rgetter, Instr*, ulong);
	char*	a;
};

static	void	format(char*, Instr*, char*);
static	char	FRAMENAME[] = &#34;.frame&#34;;

/*
 * Arm-specific debugger interface
 */

extern	char	*armexcep(Map*, Rgetter);
static	int	armfoll(Map*, uvlong, Rgetter, uvlong*);
static	int	arminst(Map*, uvlong, char, char*, int);
static	int	armdas(Map*, uvlong, char*, int);
static	int	arminstlen(Map*, uvlong);

/*
 *	Debugger interface
 */
Machdata armmach =
{
	{0, 0, 0, 0xD},		/* break point */
	4,			/* break point size */

	leswab,			/* short to local byte order */
	leswal,			/* long to local byte order */
	leswav,			/* long to local byte order */
	risctrace,		/* C traceback */
	riscframe,		/* Frame finder */
	armexcep,			/* print exception */
	0,			/* breakpoint fixup */
	0,			/* single precision float printer */
	0,			/* double precision float printer */
	armfoll,		/* following addresses */
	arminst,		/* print instruction */
	armdas,			/* dissembler */
	arminstlen,		/* instruction size */
};

char*
armexcep(Map *map, Rgetter rget)
{
	long c;

	c = (*rget)(map, &#34;TYPE&#34;);
	switch (c&amp;0x1f) {
	case 0x11:
		return &#34;Fiq interrupt&#34;;
	case 0x12:
		return &#34;Mirq interrupt&#34;;
	case 0x13:
		return &#34;SVC/SWI Exception&#34;;
	case 0x17:
		return &#34;Prefetch Abort/Data Abort&#34;;
	case 0x18:
		return &#34;Data Abort&#34;;
	case 0x1b:
		return &#34;Undefined instruction/Breakpoint&#34;;
	case 0x1f:
		return &#34;Sys trap&#34;;
	default:
		return &#34;Undefined trap&#34;;
	}
}

static
char*	cond[16] =
{
	&#34;EQ&#34;,	&#34;NE&#34;,	&#34;CS&#34;,	&#34;CC&#34;,
	&#34;MI&#34;,	&#34;PL&#34;,	&#34;VS&#34;,	&#34;VC&#34;,
	&#34;HI&#34;,	&#34;LS&#34;,	&#34;GE&#34;,	&#34;LT&#34;,
	&#34;GT&#34;,	&#34;LE&#34;,	0,	&#34;NV&#34;
};

static
char*	shtype[4] =
{
	&#34;&lt;&lt;&#34;,	&#34;&gt;&gt;&#34;,	&#34;-&gt;&#34;,	&#34;@&gt;&#34;
};

static
char *hb[4] =
{
	&#34;???&#34;,	&#34;HU&#34;, &#34;B&#34;, &#34;H&#34;
};

static
char*	addsub[2] =
{
	&#34;-&#34;,	&#34;+&#34;,
};

int
armclass(long w)
{
	int op;

	op = (w &gt;&gt; 25) &amp; 0x7;
	switch(op) {
	case 0:	/* data processing r,r,r */
		op = ((w &gt;&gt; 4) &amp; 0xf);
		if(op == 0x9) {
			op = 48+16;		/* mul */
			if(w &amp; (1&lt;&lt;24)) {
				op += 2;
				if(w &amp; (1&lt;&lt;22))
					op++;	/* swap */
				break;
			}
			if(w &amp; (1&lt;&lt;21))
				op++;		/* mla */
			break;
		}
		if ((op &amp; 0x9) == 0x9)		/* ld/st byte/half s/u */
		{
			op = (48+16+4) + ((w &gt;&gt; 22) &amp; 0x1) + ((w &gt;&gt; 19) &amp; 0x2);
			break;
		}
		op = (w &gt;&gt; 21) &amp; 0xf;
		if(w &amp; (1&lt;&lt;4))
			op += 32;
		else
		if(w &amp; (31&lt;&lt;7))
			op += 16;
		break;
	case 1:	/* data processing i,r,r */
		op = (48) + ((w &gt;&gt; 21) &amp; 0xf);
		break;
	case 2:	/* load/store byte/word i(r) */
		op = (48+24) + ((w &gt;&gt; 22) &amp; 0x1) + ((w &gt;&gt; 19) &amp; 0x2);
		break;
	case 3:	/* load/store byte/word (r)(r) */
		op = (48+24+4) + ((w &gt;&gt; 22) &amp; 0x1) + ((w &gt;&gt; 19) &amp; 0x2);
		break;
	case 4:	/* block data transfer (r)(r) */
		op = (48+24+4+4) + ((w &gt;&gt; 20) &amp; 0x1);
		break;
	case 5:	/* branch / branch link */
		op = (48+24+4+4+2) + ((w &gt;&gt; 24) &amp; 0x1);
		break;
	case 7:	/* coprocessor crap */
		op = (48+24+4+4+2+2) + ((w &gt;&gt; 3) &amp; 0x2) + ((w &gt;&gt; 20) &amp; 0x1);
		break;
	default:
		op = (48+24+4+4+2+2+4);
		break;
	}
	return op;
}

static int
decode(Map *map, ulong pc, Instr *i)
{
	uint32 w;

	if(get4(map, pc, &amp;w) &lt; 0) {
		werrstr(&#34;can&#39;t read instruction: %r&#34;);
		return -1;
	}
	i-&gt;w = w;
	i-&gt;addr = pc;
	i-&gt;cond = (w &gt;&gt; 28) &amp; 0xF;
	i-&gt;op = armclass(w);
	i-&gt;map = map;
	return 1;
}

static void
bprint(Instr *i, char *fmt, ...)
{
	va_list arg;

	va_start(arg, fmt);
	i-&gt;curr = vseprint(i-&gt;curr, i-&gt;end, fmt, arg);
	va_end(arg);
}

static int
plocal(Instr *i)
{
	char *reg;
	Symbol s;
	char *fn;
	int class;
	int offset;

	if(!findsym(i-&gt;addr, CTEXT, &amp;s)) {
		if(debug)fprint(2,&#34;fn not found @%lux: %r\n&#34;, i-&gt;addr);
		return 0;
	}
	fn = s.name;
	if (!findlocal(&amp;s, FRAMENAME, &amp;s)) {
		if(debug)fprint(2,&#34;%s.%s not found @%s: %r\n&#34;, fn, FRAMENAME, s.name);
			return 0;
	}
	if(s.value &gt; i-&gt;imm) {
		class = CAUTO;
		offset = s.value-i-&gt;imm;
		reg = &#34;(SP)&#34;;
	} else {
		class = CPARAM;
		offset = i-&gt;imm-s.value-4;
		reg = &#34;(FP)&#34;;
	}
	if(!getauto(&amp;s, offset, class, &amp;s)) {
		if(debug)fprint(2,&#34;%s %s not found @%ux: %r\n&#34;, fn,
			class == CAUTO ? &#34; auto&#34; : &#34;param&#34;, offset);
		return 0;
	}
	bprint(i, &#34;%s%c%d%s&#34;, s.name, class == CPARAM ? &#39;+&#39; : &#39;-&#39;, s.value, reg);
	return 1;
}

/*
 * Print value v as name[+offset]
 */
int
gsymoff(char *buf, int n, long v, int space)
{
	Symbol s;
	int r;
	long delta;

	r = delta = 0;		/* to shut compiler up */
	if (v) {
		r = findsym(v, space, &amp;s);
		if (r)
			delta = v-s.value;
		if (delta &lt; 0)
			delta = -delta;
	}
	if (v == 0 || r == 0 || delta &gt;= 4096)
		return snprint(buf, n, &#34;#%lux&#34;, v);
	if (strcmp(s.name, &#34;.string&#34;) == 0)
		return snprint(buf, n, &#34;#%lux&#34;, v);
	if (!delta)
		return snprint(buf, n, &#34;%s&#34;, s.name);
	if (s.type != &#39;t&#39; &amp;&amp; s.type != &#39;T&#39;)
		return snprint(buf, n, &#34;%s+%lux&#34;, s.name, v-s.value);
	else
		return snprint(buf, n, &#34;#%lux&#34;, v);
}

static void
armdps(Opcode *o, Instr *i)
{
	i-&gt;store = (i-&gt;w &gt;&gt; 20) &amp; 1;
	i-&gt;rn = (i-&gt;w &gt;&gt; 16) &amp; 0xf;
	i-&gt;rd = (i-&gt;w &gt;&gt; 12) &amp; 0xf;
	i-&gt;rs = (i-&gt;w &gt;&gt; 0) &amp; 0xf;
	if(i-&gt;rn == 15 &amp;&amp; i-&gt;rs == 0) {
		if(i-&gt;op == 8) {
			format(&#34;MOVW&#34;, i,&#34;CPSR, R%d&#34;);
			return;
		} else
		if(i-&gt;op == 10) {
			format(&#34;MOVW&#34;, i,&#34;SPSR, R%d&#34;);
			return;
		}
	} else
	if(i-&gt;rn == 9 &amp;&amp; i-&gt;rd == 15) {
		if(i-&gt;op == 9) {
			format(&#34;MOVW&#34;, i, &#34;R%s, CPSR&#34;);
			return;
		} else
		if(i-&gt;op == 11) {
			format(&#34;MOVW&#34;, i, &#34;R%s, SPSR&#34;);
			return;
		}
	}
	format(o-&gt;o, i, o-&gt;a);
}

static void
armdpi(Opcode *o, Instr *i)
{
	ulong v;
	int c;

	v = (i-&gt;w &gt;&gt; 0) &amp; 0xff;
	c = (i-&gt;w &gt;&gt; 8) &amp; 0xf;
	while(c) {
		v = (v&lt;&lt;30) | (v&gt;&gt;2);
		c--;
	}
	i-&gt;imm = v;
	i-&gt;store = (i-&gt;w &gt;&gt; 20) &amp; 1;
	i-&gt;rn = (i-&gt;w &gt;&gt; 16) &amp; 0xf;
	i-&gt;rd = (i-&gt;w &gt;&gt; 12) &amp; 0xf;
	i-&gt;rs = i-&gt;w&amp;0x0f;

		/* RET is encoded as ADD #0,R14,R15 */
	if((i-&gt;w &amp; 0x0fffffff) == 0x028ef000){
		format(&#34;RET%C&#34;, i, &#34;&#34;);
		return;
	}
	if((i-&gt;w &amp; 0x0ff0ffff) == 0x0280f000){
		format(&#34;B%C&#34;, i, &#34;0(R%n)&#34;);
		return;
	}
	format(o-&gt;o, i, o-&gt;a);
}

static void
armsdti(Opcode *o, Instr *i)
{
	ulong v;

	v = i-&gt;w &amp; 0xfff;
	if(!(i-&gt;w &amp; (1&lt;&lt;23)))
		v = -v;
	i-&gt;store = ((i-&gt;w &gt;&gt; 23) &amp; 0x2) | ((i-&gt;w &gt;&gt;21) &amp; 0x1);
	i-&gt;imm = v;
	i-&gt;rn = (i-&gt;w &gt;&gt; 16) &amp; 0xf;
	i-&gt;rd = (i-&gt;w &gt;&gt; 12) &amp; 0xf;
		/* RET is encoded as LW.P x,R13,R15 */
	if ((i-&gt;w &amp; 0x0ffff000) == 0x049df000)
	{
		format(&#34;RET%C%p&#34;, i, &#34;%I&#34;);
		return;
	}
	format(o-&gt;o, i, o-&gt;a);
}

/* arm V4 ld/st halfword, signed byte */
static void
armhwby(Opcode *o, Instr *i)
{
	i-&gt;store = ((i-&gt;w &gt;&gt; 23) &amp; 0x2) | ((i-&gt;w &gt;&gt;21) &amp; 0x1);
	i-&gt;imm = (i-&gt;w &amp; 0xf) | ((i-&gt;w &gt;&gt; 8) &amp; 0xf);
	if (!(i-&gt;w &amp; (1 &lt;&lt; 23)))
		i-&gt;imm = - i-&gt;imm;
	i-&gt;rn = (i-&gt;w &gt;&gt; 16) &amp; 0xf;
	i-&gt;rd = (i-&gt;w &gt;&gt; 12) &amp; 0xf;
	i-&gt;rs = (i-&gt;w &gt;&gt; 0) &amp; 0xf;
	format(o-&gt;o, i, o-&gt;a);
}

static void
armsdts(Opcode *o, Instr *i)
{
	i-&gt;store = ((i-&gt;w &gt;&gt; 23) &amp; 0x2) | ((i-&gt;w &gt;&gt;21) &amp; 0x1);
	i-&gt;rs = (i-&gt;w &gt;&gt; 0) &amp; 0xf;
	i-&gt;rn = (i-&gt;w &gt;&gt; 16) &amp; 0xf;
	i-&gt;rd = (i-&gt;w &gt;&gt; 12) &amp; 0xf;
	format(o-&gt;o, i, o-&gt;a);
}

static void
armbdt(Opcode *o, Instr *i)
{
	i-&gt;store = (i-&gt;w &gt;&gt; 21) &amp; 0x3;		/* S &amp; W bits */
	i-&gt;rn = (i-&gt;w &gt;&gt; 16) &amp; 0xf;
	i-&gt;imm = i-&gt;w &amp; 0xffff;
	if(i-&gt;w == 0xe8fd8000)
		format(&#34;RFE&#34;, i, &#34;&#34;);
	else
		format(o-&gt;o, i, o-&gt;a);
}

/*
static void
armund(Opcode *o, Instr *i)
{
	format(o-&gt;o, i, o-&gt;a);
}

static void
armcdt(Opcode *o, Instr *i)
{
	format(o-&gt;o, i, o-&gt;a);
}
*/

static void
armunk(Opcode *o, Instr *i)
{
	format(o-&gt;o, i, o-&gt;a);
}

static void
armb(Opcode *o, Instr *i)
{
	ulong v;

	v = i-&gt;w &amp; 0xffffff;
	if(v &amp; 0x800000)
		v |= ~0xffffff;
	i-&gt;imm = (v&lt;&lt;2) + i-&gt;addr + 8;
	format(o-&gt;o, i, o-&gt;a);
}

static void
armco(Opcode *o, Instr *i)		/* coprocessor instructions */
{
	int op, p, cp;

	char buf[1024];

	i-&gt;rn = (i-&gt;w &gt;&gt; 16) &amp; 0xf;
	i-&gt;rd = (i-&gt;w &gt;&gt; 12) &amp; 0xf;
	i-&gt;rs = i-&gt;w&amp;0xf;
	cp = (i-&gt;w &gt;&gt; 8) &amp; 0xf;
	p = (i-&gt;w &gt;&gt; 5) &amp; 0x7;
	if(i-&gt;w&amp;(1&lt;&lt;4)) {
		op = (i-&gt;w &gt;&gt; 21) &amp; 0x07;
		snprint(buf, sizeof(buf), &#34;#%x, #%x, R%d, C(%d), C(%d), #%x\n&#34;, cp, op, i-&gt;rd, i-&gt;rn, i-&gt;rs, p);
	} else {
		op = (i-&gt;w &gt;&gt; 20) &amp; 0x0f;
		snprint(buf, sizeof(buf), &#34;#%x, #%x, C(%d), C(%d), C(%d), #%x\n&#34;, cp, op, i-&gt;rd, i-&gt;rn, i-&gt;rs, p);
	}
	format(o-&gt;o, i, buf);
}

static int
armcondpass(Map *map, Rgetter rget, uchar cond)
{
	ulong psr;
	uchar n;
	uchar z;
	uchar c;
	uchar v;

	psr = rget(map, &#34;PSR&#34;);
	n = (psr &gt;&gt; 31) &amp; 1;
	z = (psr &gt;&gt; 30) &amp; 1;
	c = (psr &gt;&gt; 29) &amp; 1;
	v = (psr &gt;&gt; 28) &amp; 1;

	switch(cond) {
		case 0:		return z;
		case 1:		return !z;
		case 2:		return c;
		case 3:		return !c;
		case 4:		return n;
		case 5:		return !n;
		case 6:		return v;
		case 7:		return !v;
		case 8:		return c &amp;&amp; !z;
		case 9:		return !c || z;
		case 10:	return n == v;
		case 11:	return n != v;
		case 12:	return !z &amp;&amp; (n == v);
		case 13:	return z &amp;&amp; (n != v);
		case 14:	return 1;
		case 15:	return 0;
	}
	return 0;
}

static ulong
armshiftval(Map *map, Rgetter rget, Instr *i)
{
	if(i-&gt;w &amp; (1 &lt;&lt; 25)) {				/* immediate */
		ulong imm = i-&gt;w &amp; BITS(0, 7);
		ulong s = (i-&gt;w &amp; BITS(8, 11)) &gt;&gt; 7; /* this contains the *2 */
		return ROR(imm, s);
	} else {
		char buf[8];
		ulong v;
		ulong s = (i-&gt;w &amp; BITS(7,11)) &gt;&gt; 7;

		sprint(buf, &#34;R%ld&#34;, i-&gt;w &amp; 0xf);
		v = rget(map, buf);

		switch((i-&gt;w &amp; BITS(4, 6)) &gt;&gt; 4) {
		case 0:					/* LSLIMM */
			return v &lt;&lt; s;
		case 1:					/* LSLREG */
			sprint(buf, &#34;R%lud&#34;, s &gt;&gt; 1);
			s = rget(map, buf) &amp; 0xFF;
			if(s &gt;= 32) return 0;
			return v &lt;&lt; s;
		case 2:					/* LSRIMM */
			return LSR(v, s);
		case 3:					/* LSRREG */
			sprint(buf, &#34;R%ld&#34;, s &gt;&gt; 1);
			s = rget(map, buf) &amp; 0xFF;
			if(s &gt;= 32) return 0;
			return LSR(v, s);
		case 4:					/* ASRIMM */
			if(s == 0) {
				if((v &amp; (1U&lt;&lt;31)) == 0)
					return 0;
				return 0xFFFFFFFF;
			}
			return ASR(v, s);
		case 5:					/* ASRREG */
			sprint(buf, &#34;R%ld&#34;, s &gt;&gt; 1);
			s = rget(map, buf) &amp; 0xFF;
			if(s &gt;= 32) {
				if((v &amp; (1U&lt;&lt;31)) == 0)
					return 0;
				return 0xFFFFFFFF;
			}
			return ASR(v, s);
		case 6:					/* RORIMM */
			if(s == 0) {
				ulong c = (rget(map, &#34;PSR&#34;) &gt;&gt; 29) &amp; 1;

				return (c &lt;&lt; 31) | LSR(v, 1);
			}
			return ROR(v, s);
		case 7:					/* RORREG */
			sprint(buf, &#34;R%ld&#34;, (s&gt;&gt;1)&amp;0xF);
			s = rget(map, buf);
			if(s == 0 || (s &amp; 0xF) == 0)
				return v;
			return ROR(v, s &amp; 0xF);
		}
	}
	return 0;
}

static int
nbits(ulong v)
{
	int n = 0;
	int i;

	for(i=0; i &lt; 32 ; i++) {
		if(v &amp; 1) ++n;
		v &gt;&gt;= 1;
	}

	return n;
}

static ulong
armmaddr(Map *map, Rgetter rget, Instr *i)
{
	ulong v;
	ulong nb;
	char buf[8];
	ulong rn;

	rn = (i-&gt;w &gt;&gt; 16) &amp; 0xf;
	sprint(buf,&#34;R%ld&#34;, rn);

	v = rget(map, buf);
	nb = nbits(i-&gt;w &amp; ((1 &lt;&lt; 15) - 1));

	switch((i-&gt;w &gt;&gt; 23) &amp; 3) {
		case 0: return (v - (nb*4)) + 4;
		case 1: return v;
		case 2: return v - (nb*4);
		case 3: return v + 4;
	}
	return 0;
}

static ulong
armaddr(Map *map, Rgetter rget, Instr *i)
{
	char buf[8];
	ulong rn;

	sprint(buf, &#34;R%ld&#34;, (i-&gt;w &gt;&gt; 16) &amp; 0xf);
	rn = rget(map, buf);

	if((i-&gt;w &amp; (1&lt;&lt;24)) == 0) {			/* POSTIDX */
		sprint(buf, &#34;R%ld&#34;, rn);
		return rget(map, buf);
	}

	if((i-&gt;w &amp; (1&lt;&lt;25)) == 0) {			/* OFFSET */
		sprint(buf, &#34;R%ld&#34;, rn);
		if(i-&gt;w &amp; (1U&lt;&lt;23))
			return rget(map, buf) + (i-&gt;w &amp; BITS(0,11));
		return rget(map, buf) - (i-&gt;w &amp; BITS(0,11));
	} else {					/* REGOFF */
		ulong index = 0;
		uchar c;
		uchar rm;

		sprint(buf, &#34;R%ld&#34;, i-&gt;w &amp; 0xf);
		rm = rget(map, buf);

		switch((i-&gt;w &amp; BITS(5,6)) &gt;&gt; 5) {
		case 0: index = rm &lt;&lt; ((i-&gt;w &amp; BITS(7,11)) &gt;&gt; 7);	break;
		case 1: index = LSR(rm, ((i-&gt;w &amp; BITS(7,11)) &gt;&gt; 7));	break;
		case 2: index = ASR(rm, ((i-&gt;w &amp; BITS(7,11)) &gt;&gt; 7));	break;
		case 3:
			if((i-&gt;w &amp; BITS(7,11)) == 0) {
				c = (rget(map, &#34;PSR&#34;) &gt;&gt; 29) &amp; 1;
				index = c &lt;&lt; 31 | LSR(rm, 1);
			} else {
				index = ROR(rm, ((i-&gt;w &amp; BITS(7,11)) &gt;&gt; 7));
			}
			break;
		}
		if(i-&gt;w &amp; (1&lt;&lt;23))
			return rn + index;
		return rn - index;
	}
}

static ulong
armfadd(Map *map, Rgetter rget, Instr *i, ulong pc)
{
	char buf[8];
	int r;

	r = (i-&gt;w &gt;&gt; 12) &amp; 0xf;
	if(r != 15 || !armcondpass(map, rget, (i-&gt;w &gt;&gt; 28) &amp; 0xf))
		return pc+4;

	r = (i-&gt;w &gt;&gt; 16) &amp; 0xf;
	sprint(buf, &#34;R%d&#34;, r);

	return rget(map, buf) + armshiftval(map, rget, i);
}

static ulong
armfmovm(Map *map, Rgetter rget, Instr *i, ulong pc)
{
	uint32 v;
	ulong addr;

	v = i-&gt;w &amp; 1&lt;&lt;15;
	if(!v || !armcondpass(map, rget, (i-&gt;w&gt;&gt;28)&amp;0xf))
		return pc+4;

	addr = armmaddr(map, rget, i) + nbits(i-&gt;w &amp; BITS(0,15));
	if(get4(map, addr, &amp;v) &lt; 0) {
		werrstr(&#34;can&#39;t read addr: %r&#34;);
		return -1;
	}
	return v;
}

static ulong
armfbranch(Map *map, Rgetter rget, Instr *i, ulong pc)
{
	if(!armcondpass(map, rget, (i-&gt;w &gt;&gt; 28) &amp; 0xf))
		return pc+4;

	return pc + (((signed long)i-&gt;w &lt;&lt; 8) &gt;&gt; 6) + 8;
}

static ulong
armfmov(Map *map, Rgetter rget, Instr *i, ulong pc)
{
	ulong rd;
	uint32 v;

	rd = (i-&gt;w &gt;&gt; 12) &amp; 0xf;
	if(rd != 15 || !armcondpass(map, rget, (i-&gt;w&gt;&gt;28)&amp;0xf))
		return pc+4;

	 /* LDR */
	/* BUG: Needs LDH/B, too */
	if(((i-&gt;w&gt;&gt;26)&amp;0x3) == 1) {
		if(get4(map, armaddr(map, rget, i), &amp;v) &lt; 0) {
			werrstr(&#34;can&#39;t read instruction: %r&#34;);
			return pc+4;
		}
		return v;
	}

	 /* MOV */
	return armshiftval(map, rget, i);
}

static Opcode opcodes[] =
{
	&#34;AND%C%S&#34;,	armdps, 0,	&#34;R%s,R%n,R%d&#34;,
	&#34;EOR%C%S&#34;,	armdps, 0,	&#34;R%s,R%n,R%d&#34;,
	&#34;SUB%C%S&#34;,	armdps, 0,	&#34;R%s,R%n,R%d&#34;,
	&#34;RSB%C%S&#34;,	armdps, 0,	&#34;R%s,R%n,R%d&#34;,
	&#34;ADD%C%S&#34;,	armdps, armfadd,	&#34;R%s,R%n,R%d&#34;,
	&#34;ADC%C%S&#34;,	armdps, 0,	&#34;R%s,R%n,R%d&#34;,
	&#34;SBC%C%S&#34;,	armdps, 0,	&#34;R%s,R%n,R%d&#34;,
	&#34;RSC%C%S&#34;,	armdps, 0,	&#34;R%s,R%n,R%d&#34;,
	&#34;TST%C%S&#34;,	armdps, 0,	&#34;R%s,R%n&#34;,
	&#34;TEQ%C%S&#34;,	armdps, 0,	&#34;R%s,R%n&#34;,
	&#34;CMP%C%S&#34;,	armdps, 0,	&#34;R%s,R%n&#34;,
	&#34;CMN%C%S&#34;,	armdps, 0,	&#34;R%s,R%n&#34;,
	&#34;ORR%C%S&#34;,	armdps, 0,	&#34;R%s,R%n,R%d&#34;,
	&#34;MOVW%C%S&#34;,	armdps, armfmov,	&#34;R%s,R%d&#34;,
	&#34;BIC%C%S&#34;,	armdps, 0,	&#34;R%s,R%n,R%d&#34;,
	&#34;MVN%C%S&#34;,	armdps, 0,	&#34;R%s,R%d&#34;,

/* 16 */
	&#34;AND%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n,R%d&#34;,
	&#34;EOR%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n,R%d&#34;,
	&#34;SUB%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n,R%d&#34;,
	&#34;RSB%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n,R%d&#34;,
	&#34;ADD%C%S&#34;,	armdps, armfadd,	&#34;(R%s%h%m),R%n,R%d&#34;,
	&#34;ADC%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n,R%d&#34;,
	&#34;SBC%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n,R%d&#34;,
	&#34;RSC%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n,R%d&#34;,
	&#34;TST%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n&#34;,
	&#34;TEQ%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n&#34;,
	&#34;CMP%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n&#34;,
	&#34;CMN%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n&#34;,
	&#34;ORR%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n,R%d&#34;,
	&#34;MOVW%C%S&#34;,	armdps, armfmov,	&#34;(R%s%h%m),R%d&#34;,
	&#34;BIC%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%n,R%d&#34;,
	&#34;MVN%C%S&#34;,	armdps, 0,	&#34;(R%s%h%m),R%d&#34;,

/* 32 */
	&#34;AND%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n,R%d&#34;,
	&#34;EOR%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n,R%d&#34;,
	&#34;SUB%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n,R%d&#34;,
	&#34;RSB%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n,R%d&#34;,
	&#34;ADD%C%S&#34;,	armdps, armfadd,	&#34;(R%s%hR%M),R%n,R%d&#34;,
	&#34;ADC%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n,R%d&#34;,
	&#34;SBC%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n,R%d&#34;,
	&#34;RSC%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n,R%d&#34;,
	&#34;TST%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n&#34;,
	&#34;TEQ%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n&#34;,
	&#34;CMP%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n&#34;,
	&#34;CMN%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n&#34;,
	&#34;ORR%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n,R%d&#34;,
	&#34;MOVW%C%S&#34;,	armdps, armfmov,	&#34;(R%s%hR%M),R%d&#34;,
	&#34;BIC%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%n,R%d&#34;,
	&#34;MVN%C%S&#34;,	armdps, 0,	&#34;(R%s%hR%M),R%d&#34;,

/* 48 */
	&#34;AND%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n,R%d&#34;,
	&#34;EOR%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n,R%d&#34;,
	&#34;SUB%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n,R%d&#34;,
	&#34;RSB%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n,R%d&#34;,
	&#34;ADD%C%S&#34;,	armdpi, armfadd,	&#34;$#%i,R%n,R%d&#34;,
	&#34;ADC%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n,R%d&#34;,
	&#34;SBC%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n,R%d&#34;,
	&#34;RSC%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n,R%d&#34;,
	&#34;TST%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n&#34;,
	&#34;TEQ%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n&#34;,
	&#34;CMP%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n&#34;,
	&#34;CMN%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n&#34;,
	&#34;ORR%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n,R%d&#34;,
	&#34;MOVW%C%S&#34;,	armdpi, armfmov,	&#34;$#%i,R%d&#34;,
	&#34;BIC%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%n,R%d&#34;,
	&#34;MVN%C%S&#34;,	armdpi, 0,	&#34;$#%i,R%d&#34;,

/* 48+16 */
	&#34;MUL%C%S&#34;,	armdpi, 0,	&#34;R%s,R%M,R%n&#34;,
	&#34;MULA%C%S&#34;,	armdpi, 0,	&#34;R%s,R%M,R%n,R%d&#34;,
	&#34;SWPW&#34;,		armdpi, 0,	&#34;R%s,(R%n),R%d&#34;,
	&#34;SWPB&#34;,		armdpi, 0,	&#34;R%s,(R%n),R%d&#34;,

/* 48+16+4 */
	&#34;MOV%u%C%p&#34;,	armhwby, 0,	&#34;R%d,(R%n%UR%M)&#34;,
	&#34;MOV%u%C%p&#34;,	armhwby, 0,	&#34;R%d,%I&#34;,
	&#34;MOV%u%C%p&#34;,	armhwby, armfmov,	&#34;(R%n%UR%M),R%d&#34;,
	&#34;MOV%u%C%p&#34;,	armhwby, armfmov,	&#34;%I,R%d&#34;,

/* 48+24 */
	&#34;MOVW%C%p&#34;,	armsdti, 0,	&#34;R%d,%I&#34;,
	&#34;MOVB%C%p&#34;,	armsdti, 0,	&#34;R%d,%I&#34;,
	&#34;MOVW%C%p&#34;,	armsdti, armfmov,	&#34;%I,R%d&#34;,
	&#34;MOVBU%C%p&#34;,	armsdti, armfmov,	&#34;%I,R%d&#34;,

	&#34;MOVW%C%p&#34;,	armsdts, 0,	&#34;R%d,(R%s%h%m)(R%n)&#34;,
	&#34;MOVB%C%p&#34;,	armsdts, 0,	&#34;R%d,(R%s%h%m)(R%n)&#34;,
	&#34;MOVW%C%p&#34;,	armsdts, armfmov,	&#34;(R%s%h%m)(R%n),R%d&#34;,
	&#34;MOVBU%C%p&#34;,	armsdts, armfmov,	&#34;(R%s%h%m)(R%n),R%d&#34;,

	&#34;MOVM%C%P%a&#34;,	armbdt, armfmovm,		&#34;[%r],(R%n)&#34;,
	&#34;MOVM%C%P%a&#34;,	armbdt, armfmovm,		&#34;(R%n),[%r]&#34;,

	&#34;B%C&#34;,		armb, armfbranch,		&#34;%b&#34;,
	&#34;BL%C&#34;,		armb, armfbranch,		&#34;%b&#34;,

	&#34;CDP%C&#34;,	armco, 0,		&#34;&#34;,
	&#34;CDP%C&#34;,	armco, 0,		&#34;&#34;,
	&#34;MCR%C&#34;,	armco, 0,		&#34;&#34;,
	&#34;MRC%C&#34;,	armco, 0,		&#34;&#34;,

	&#34;UNK&#34;,		armunk, 0,	&#34;&#34;,
};

static void
gaddr(Instr *i)
{
	*i-&gt;curr++ = &#39;$&#39;;
	i-&gt;curr += gsymoff(i-&gt;curr, i-&gt;end-i-&gt;curr, i-&gt;imm, CANY);
}

static	char *mode[] = { 0, &#34;IA&#34;, &#34;DB&#34;, &#34;IB&#34; };
static	char *pw[] = { &#34;P&#34;, &#34;PW&#34;, 0, &#34;W&#34; };
static	char *sw[] = { 0, &#34;W&#34;, &#34;S&#34;, &#34;SW&#34; };

static void
format(char *mnemonic, Instr *i, char *f)
{
	int j, k, m, n;
	int g;
	char *fmt;

	if(mnemonic)
		format(0, i, mnemonic);
	if(f == 0)
		return;
	if(mnemonic)
		if(i-&gt;curr &lt; i-&gt;end)
			*i-&gt;curr++ = &#39;\t&#39;;
	for ( ; *f &amp;&amp; i-&gt;curr &lt; i-&gt;end; f++) {
		if(*f != &#39;%&#39;) {
			*i-&gt;curr++ = *f;
			continue;
		}
		switch (*++f) {

		case &#39;C&#39;:	/* .CONDITION */
			if(cond[i-&gt;cond])
				bprint(i, &#34;.%s&#34;, cond[i-&gt;cond]);
			break;

		case &#39;S&#39;:	/* .STORE */
			if(i-&gt;store)
				bprint(i, &#34;.S&#34;);
			break;

		case &#39;P&#39;:	/* P &amp; U bits for block move */
			n = (i-&gt;w &gt;&gt;23) &amp; 0x3;
			if (mode[n])
				bprint(i, &#34;.%s&#34;, mode[n]);
			break;

		case &#39;p&#39;:	/* P &amp; W bits for single data xfer*/
			if (pw[i-&gt;store])
				bprint(i, &#34;.%s&#34;, pw[i-&gt;store]);
			break;

		case &#39;a&#39;:	/* S &amp; W bits for single data xfer*/
			if (sw[i-&gt;store])
				bprint(i, &#34;.%s&#34;, sw[i-&gt;store]);
			break;

		case &#39;s&#39;:
			bprint(i, &#34;%d&#34;, i-&gt;rs &amp; 0xf);
			break;

		case &#39;M&#39;:
			bprint(i, &#34;%d&#34;, (i-&gt;w&gt;&gt;8) &amp; 0xf);
			break;

		case &#39;m&#39;:
			bprint(i, &#34;%d&#34;, (i-&gt;w&gt;&gt;7) &amp; 0x1f);
			break;

		case &#39;h&#39;:
			bprint(i, shtype[(i-&gt;w&gt;&gt;5) &amp; 0x3]);
			break;

		case &#39;u&#39;:		/* Signed/unsigned Byte/Halfword */
			bprint(i, hb[(i-&gt;w&gt;&gt;5) &amp; 0x3]);
			break;

		case &#39;I&#39;:
			if (i-&gt;rn == 13) {
				if (plocal(i))
					break;
			}
			g = 0;
			fmt = &#34;#%lx(R%d)&#34;;
			if (i-&gt;rn == 15) {
				/* convert load of offset(PC) to a load immediate */
				uint32 x;
				if (get4(i-&gt;map, i-&gt;addr+i-&gt;imm+8, &amp;x) &gt; 0)
				{
					i-&gt;imm = (int32)x;
					g = 1;
					fmt = &#34;&#34;;
				}
			}
			if (mach-&gt;sb)
			{
				if (i-&gt;rd == 11) {
					uint32 nxti;

					if (get4(i-&gt;map, i-&gt;addr+4, &amp;nxti) &gt; 0) {
						if ((nxti &amp; 0x0e0f0fff) == 0x060c000b) {
							i-&gt;imm += mach-&gt;sb;
							g = 1;
							fmt = &#34;-SB&#34;;
						}
					}
				}
				if (i-&gt;rn == 12)
				{
					i-&gt;imm += mach-&gt;sb;
					g = 1;
					fmt = &#34;-SB(SB)&#34;;
				}
			}
			if (g)
			{
				gaddr(i);
				bprint(i, fmt, i-&gt;rn);
			}
			else
				bprint(i, fmt, i-&gt;imm, i-&gt;rn);
			break;
		case &#39;U&#39;:		/* Add/subtract from base */
			bprint(i, addsub[(i-&gt;w &gt;&gt; 23) &amp; 1]);
			break;

		case &#39;n&#39;:
			bprint(i, &#34;%d&#34;, i-&gt;rn);
			break;

		case &#39;d&#39;:
			bprint(i, &#34;%d&#34;, i-&gt;rd);
			break;

		case &#39;i&#39;:
			bprint(i, &#34;%lux&#34;, i-&gt;imm);
			break;

		case &#39;b&#39;:
			i-&gt;curr += symoff(i-&gt;curr, i-&gt;end-i-&gt;curr,
				i-&gt;imm, CTEXT);
			break;

		case &#39;g&#39;:
			i-&gt;curr += gsymoff(i-&gt;curr, i-&gt;end-i-&gt;curr,
				i-&gt;imm, CANY);
			break;

		case &#39;r&#39;:
			n = i-&gt;imm&amp;0xffff;
			j = 0;
			k = 0;
			while(n) {
				m = j;
				while(n&amp;0x1) {
					j++;
					n &gt;&gt;= 1;
				}
				if(j != m) {
					if(k)
						bprint(i, &#34;,&#34;);
					if(j == m+1)
						bprint(i, &#34;R%d&#34;, m);
					else
						bprint(i, &#34;R%d-R%d&#34;, m, j-1);
					k = 1;
				}
				j++;
				n &gt;&gt;= 1;
			}
			break;

		case &#39;\0&#39;:
			*i-&gt;curr++ = &#39;%&#39;;
			return;

		default:
			bprint(i, &#34;%%%c&#34;, *f);
			break;
		}
	}
	*i-&gt;curr = 0;
}

static int
printins(Map *map, ulong pc, char *buf, int n)
{
	Instr i;

	i.curr = buf;
	i.end = buf+n-1;
	if(decode(map, pc, &amp;i) &lt; 0)
		return -1;

	(*opcodes[i.op].fmt)(&amp;opcodes[i.op], &amp;i);
	return 4;
}

static int
arminst(Map *map, uvlong pc, char modifier, char *buf, int n)
{
	USED(modifier);
	return printins(map, pc, buf, n);
}

static int
armdas(Map *map, uvlong pc, char *buf, int n)
{
	Instr i;

	i.curr = buf;
	i.end = buf+n;
	if(decode(map, pc, &amp;i) &lt; 0)
		return -1;
	if(i.end-i.curr &gt; 8)
		i.curr = _hexify(buf, i.w, 7);
	*i.curr = 0;
	return 4;
}

static int
arminstlen(Map *map, uvlong pc)
{
	Instr i;

	if(decode(map, pc, &amp;i) &lt; 0)
		return -1;
	return 4;
}

static int
armfoll(Map *map, uvlong pc, Rgetter rget, uvlong *foll)
{
	ulong d;
	Instr i;

	if(decode(map, pc, &amp;i) &lt; 0)
		return -1;

	if(opcodes[i.op].foll) {
		d = (*opcodes[i.op].foll)(map, rget, &amp;i, pc);
		if(d == -1)
			return -1;
	} else
		d = pc+4;

	foll[0] = d;
	return 1;
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
