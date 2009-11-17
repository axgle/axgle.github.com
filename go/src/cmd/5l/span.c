<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5l/span.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/5l/span.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5l/span.c
// http://code.google.com/p/inferno-os/source/browse/utils/5l/span.c
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

static struct {
	uint32	start;
	uint32	size;
	uint32	extra;
} pool;

int	checkpool(Prog*, int);
int 	flushpool(Prog*, int, int);

int
isbranch(Prog *p)
{
	int as = p-&gt;as;
	return (as &gt;= ABEQ &amp;&amp; as &lt;= ABLE) || as == AB || as == ABL || as == ABX;
}

static int
ispad(Prog *p)
{
	if(p-&gt;as != AMOVW)
		return 0;
	if(p-&gt;from.type != D_REG || p-&gt;from.reg != REGSB)
		return 0;
	if(p-&gt;to.type != D_REG || p-&gt;to.reg != REGSB)
		return 0;
	return 1;
}

int
fninc(Sym *s)
{
	if(thumb){
		if(s-&gt;thumb){
			if(s-&gt;foreign)
				return 8;
			else
				return 0;
		}
		else{
			if(s-&gt;foreign)
				return 0;
			else
				diag(&#34;T A !foreign in fninc&#34;);
		}
	}
	else{
		if(s-&gt;thumb){
			if(s-&gt;foreign)
				return 0;
			else
				diag(&#34;A T !foreign in fninc&#34;);
		}
		else{
			if(s-&gt;foreign)
				return 4;
			else
				return 0;
		}
	}
	return 0;
}

int
fnpinc(Sym *s)
{
	if(!s-&gt;fnptr){	// a simplified case BX O(R) -&gt; BL O(R)
		if(!debug[&#39;f&#39;])
			diag(&#34;fnptr == 0 in fnpinc&#34;);
		if(s-&gt;foreign)
			diag(&#34;bad usage in fnpinc %s %d %d %d&#34;, s-&gt;name, s-&gt;used, s-&gt;foreign, s-&gt;thumb);
		return 0;
	}
	/* 0, 1, 2, 3 squared */
	if(s-&gt;thumb)
		return s-&gt;foreign ? 9 : 1;
	else
		return s-&gt;foreign ? 4 : 0;
}

static Prog *
pad(Prog *p, int pc)
{
	Prog *q;

	q = prg();
	q-&gt;as = AMOVW;
	q-&gt;line = p-&gt;line;
	q-&gt;from.type = D_REG;
	q-&gt;from.reg = REGSB;
	q-&gt;to.type = D_REG;
	q-&gt;to.reg = REGSB;
	q-&gt;pc = pc;
	q-&gt;link = p-&gt;link;
	return q;
}

static int
scan(Prog *op, Prog *p, int c)
{
	Prog *q;

	for(q = op-&gt;link; q != p; q = q-&gt;link){
		q-&gt;pc = c;
		c += oplook(q)-&gt;size;
		nocache(q);
	}
	return c;
}

/* size of a case statement including jump table */
static int32
casesz(Prog *p)
{
	int jt = 0;
	int32 n = 0;
	Optab *o;

	for( ; p != P; p = p-&gt;link){
		if(p-&gt;as == ABCASE)
			jt = 1;
		else if(jt)
			break;
		o = oplook(p);
		n += o-&gt;size;
	}
	return n;
}

void
span(void)
{
	Prog *p, *op;
	Sym *setext, *s;
	Optab *o;
	int m, bflag, i;
	int32 c, otxt, v;
	int lastthumb = -1;

	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f span\n&#34;, cputime());
	Bflush(&amp;bso);

	bflag = 0;
	c = INITTEXT;
	op = nil;
	otxt = c;
	for(p = firstp; p != P; op = p, p = p-&gt;link) {
		setarch(p);
		p-&gt;pc = c;
		o = oplook(p);
		m = o-&gt;size;
		// must check literal pool here in case p generates many instructions
		if(blitrl){
			if(thumb &amp;&amp; isbranch(p))
				pool.extra += brextra(p);
			if(checkpool(op, p-&gt;as == ACASE ? casesz(p) : m))
				c = p-&gt;pc = scan(op, p, c);
		}
		if(m == 0) {
			if(p-&gt;as == ATEXT) {
				if(blitrl &amp;&amp; lastthumb != -1 &amp;&amp; lastthumb != thumb){	// flush literal pool
					if(flushpool(op, 0, 1))
						c = p-&gt;pc = scan(op, p, c);
				}
				lastthumb = thumb;
				curtext = p;
				autosize = p-&gt;to.offset + 4;
				if(p-&gt;from.sym != S)
					p-&gt;from.sym-&gt;value = c;
				/* need passes to resolve branches */
				if(c-otxt &gt;= 1L&lt;&lt;17)
					bflag = 1;
				otxt = c;
				if(thumb &amp;&amp; blitrl)
					pool.extra += brextra(p);
				continue;
			}
			diag(&#34;zero-width instruction\n%P&#34;, p);
			continue;
		}
		switch(o-&gt;flag &amp; (LFROM|LTO|LPOOL)) {
		case LFROM:
			addpool(p, &amp;p-&gt;from);
			break;
		case LTO:
			addpool(p, &amp;p-&gt;to);
			break;
		case LPOOL:
			if ((p-&gt;scond&amp;C_SCOND) == 14)
				flushpool(p, 0, 0);
			break;
		}
		if(p-&gt;as==AMOVW &amp;&amp; p-&gt;to.type==D_REG &amp;&amp; p-&gt;to.reg==REGPC &amp;&amp; (p-&gt;scond&amp;C_SCOND) == 14)
			flushpool(p, 0, 0);
		c += m;
		if(blitrl &amp;&amp; p-&gt;link == P){
			if(thumb &amp;&amp; isbranch(p))
				pool.extra += brextra(p);
			checkpool(p, 0);
		}
	}

	/*
	 * if any procedure is large enough to
	 * generate a large SBRA branch, then
	 * generate extra passes putting branches
	 * around jmps to fix. this is rare.
	 */
	while(bflag) {
		if(debug[&#39;v&#39;])
			Bprint(&amp;bso, &#34;%5.2f span1\n&#34;, cputime());
		bflag = 0;
		c = INITTEXT;
		for(p = firstp; p != P; p = p-&gt;link) {
			setarch(p);
			p-&gt;pc = c;
			if(thumb &amp;&amp; isbranch(p))
				nocache(p);
			o = oplook(p);
/* very larg branches
			if(o-&gt;type == 6 &amp;&amp; p-&gt;cond) {
				otxt = p-&gt;cond-&gt;pc - c;
				if(otxt &lt; 0)
					otxt = -otxt;
				if(otxt &gt;= (1L&lt;&lt;17) - 10) {
					q = prg();
					q-&gt;link = p-&gt;link;
					p-&gt;link = q;
					q-&gt;as = AB;
					q-&gt;to.type = D_BRANCH;
					q-&gt;cond = p-&gt;cond;
					p-&gt;cond = q;
					q = prg();
					q-&gt;link = p-&gt;link;
					p-&gt;link = q;
					q-&gt;as = AB;
					q-&gt;to.type = D_BRANCH;
					q-&gt;cond = q-&gt;link-&gt;link;
					bflag = 1;
				}
			}
 */
			m = o-&gt;size;
			if(m == 0) {
				if(p-&gt;as == ATEXT) {
					curtext = p;
					autosize = p-&gt;to.offset + 4;
					if(p-&gt;from.sym != S)
						p-&gt;from.sym-&gt;value = c;
					continue;
				}
				diag(&#34;zero-width instruction\n%P&#34;, p);
				continue;
			}
			c += m;
		}
	}

	if(seenthumb){		// branch resolution
		int passes = 0;
		int lastc = 0;
		int again;
		Prog *oop;

	loop:
		passes++;
		if(passes &gt; 100){
			diag(&#34;span looping !&#34;);
			errorexit();
		}
		c = INITTEXT;
		oop = op = nil;
		again = 0;
		for(p = firstp; p != P; oop = op, op = p, p = p-&gt;link){
			setarch(p);
			if(p-&gt;pc != c)
				again = 1;
			p-&gt;pc = c;
			if(thumb &amp;&amp; isbranch(p))
				nocache(p);
			o = oplook(p);
			m = o-&gt;size;
			if(passes == 1 &amp;&amp; thumb &amp;&amp; isbranch(p)){	// start conservative so unneeded alignment is not added
				if(p-&gt;as == ABL)
					m = 4;
				else
					m = 2;
				p-&gt;align = 0;
			}
			if(p-&gt;align){
				if((p-&gt;align == 4 &amp;&amp; (c&amp;3)) || (p-&gt;align == 2 &amp;&amp; !(c&amp;3))){
					if(ispad(op)){
						oop-&gt;link = p;
						op = oop;
						c -= 2;
						p-&gt;pc = c;
					}
					else{
						op-&gt;link = pad(op, c);
						op = op-&gt;link;
						c += 2;
						p-&gt;pc = c;
					}
					again = 1;
				}
			}
			if(m == 0) {
				if(p-&gt;as == ATEXT) {
					curtext = p;
					autosize = p-&gt;to.offset + 4;
					if(p-&gt;from.sym != S)
						p-&gt;from.sym-&gt;value = c;
					continue;
				}
			}
			c += m;
		}
		if(c != lastc || again){
			lastc = c;
			goto loop;
		}
	}

	if(0 &amp;&amp; seenthumb){		// rm redundant padding - obsolete
		int d;

		op = nil;
		d = 0;
		for(p = firstp; p != P; op = p, p = p-&gt;link){
			p-&gt;pc -= d;
			if(p-&gt;as == ATEXT){
				if(p-&gt;from.sym != S)
					p-&gt;from.sym-&gt;value -= d;
// if(p-&gt;from.sym != S) print(&#34;%s %ux %d %d %d\n&#34;, p-&gt;from.sym-&gt;name ? p-&gt;from.sym-&gt;name : &#34;?&#34;, p-&gt;from.sym-&gt;value, p-&gt;from.sym-&gt;thumb, p-&gt;from.sym-&gt;foreign, p-&gt;from.sym-&gt;fnptr);
			}
			if(ispad(p) &amp;&amp; p-&gt;link != P &amp;&amp; ispad(p-&gt;link)){
				op-&gt;link = p-&gt;link-&gt;link;
				d += 4;
				p = op;
			}
		}
		// print(&#34;%d bytes removed (padding)\n&#34;, d);
		c -= d;
	}

	if(debug[&#39;t&#39;]) {
		/*
		 * add strings to text segment
		 */
		c = rnd(c, 8);
		for(i=0; i&lt;NHASH; i++)
		for(s = hash[i]; s != S; s = s-&gt;link) {
			if(s-&gt;type != SSTRING)
				continue;
			v = s-&gt;value;
			while(v &amp; 3)
				v++;
			s-&gt;value = c;
			c += v;
		}
	}

	c = rnd(c, 8);

	setext = lookup(&#34;etext&#34;, 0);
	if(setext != S) {
		setext-&gt;value = c;
		textsize = c - INITTEXT;
	}
	if(INITRND)
		INITDAT = rnd(c, INITRND);
	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;tsize = %lux\n&#34;, textsize);
	Bflush(&amp;bso);
}

/*
 * when the first reference to the literal pool threatens
 * to go out of range of a 12-bit PC-relative offset,
 * drop the pool now, and branch round it.
 * this happens only in extended basic blocks that exceed 4k.
 */
int
checkpool(Prog *p, int sz)
{
	if(thumb){
		if(pool.size &gt;= 0x3fc || (p-&gt;pc+sz+pool.extra+2+2)+(pool.size-4)-pool.start-4 &gt;= 0x3fc)
			return flushpool(p, 1, 0);
		else if(p-&gt;link == P)
			return flushpool(p, 2, 0);
		return 0;
	}
	if(pool.size &gt;= 0xffc || immaddr((p-&gt;pc+sz+4)+4+pool.size - pool.start+8) == 0)
		return flushpool(p, 1, 0);
	else if(p-&gt;link == P)
		return flushpool(p, 2, 0);
	return 0;
}

int
flushpool(Prog *p, int skip, int force)
{
	Prog *q;

	if(blitrl) {
		if(skip){
			if(0 &amp;&amp; skip==1)print(&#34;note: flush literal pool at %lux: len=%lud ref=%lux\n&#34;, p-&gt;pc+4, pool.size, pool.start);
			q = prg();
			q-&gt;as = AB;
			q-&gt;to.type = D_BRANCH;
			q-&gt;cond = p-&gt;link;
			q-&gt;link = blitrl;
			blitrl = q;
		}
		else if(!force &amp;&amp; (p-&gt;pc+pool.size-pool.start &lt; (thumb ? 0x3fc+4-pool.extra : 2048)))
			return 0;
		elitrl-&gt;link = p-&gt;link;
		p-&gt;link = blitrl;
		blitrl = 0;	/* BUG: should refer back to values until out-of-range */
		elitrl = 0;
		pool.size = 0;
		pool.start = 0;
		pool.extra = 0;
		return 1;
	}
	return 0;
}

void
addpool(Prog *p, Adr *a)
{
	Prog *q, t;
	int c;

	if(thumb)
		c = thumbaclass(a, p);
	else
		c = aclass(a);

	t = zprg;
	t.as = AWORD;

	switch(c) {
	default:
		t.to = *a;
		break;

	case	C_SROREG:
	case C_LOREG:
	case C_ROREG:
	case C_FOREG:
	case C_SOREG:
	case C_HOREG:
	case C_GOREG:
	case C_FAUTO:
	case C_SAUTO:
	case C_LAUTO:
	case C_LACON:
	case C_GACON:
		t.to.type = D_CONST;
		t.to.offset = instoffset;
		break;
	}

	for(q = blitrl; q != P; q = q-&gt;link)	/* could hash on t.t0.offset */
		if(memcmp(&amp;q-&gt;to, &amp;t.to, sizeof(t.to)) == 0) {
			p-&gt;cond = q;
			return;
		}

	q = prg();
	*q = t;
	q-&gt;pc = pool.size;

	if(blitrl == P) {
		blitrl = q;
		pool.start = p-&gt;pc;
		q-&gt;align = 4;
	} else
		elitrl-&gt;link = q;
	elitrl = q;
	pool.size += 4;

	p-&gt;cond = q;
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
}

int32
regoff(Adr *a)
{

	instoffset = 0;
	aclass(a);
	return instoffset;
}

int32
immrot(uint32 v)
{
	int i;

	for(i=0; i&lt;16; i++) {
		if((v &amp; ~0xff) == 0)
			return (i&lt;&lt;8) | v | (1&lt;&lt;25);
		v = (v&lt;&lt;2) | (v&gt;&gt;30);
	}
	return 0;
}

int32
immaddr(int32 v)
{
	if(v &gt;= 0 &amp;&amp; v &lt;= 0xfff)
		return (v &amp; 0xfff) |
			(1&lt;&lt;24) |	/* pre indexing */
			(1&lt;&lt;23);	/* pre indexing, up */
	if(v &gt;= -0xfff &amp;&amp; v &lt; 0)
		return (-v &amp; 0xfff) |
			(1&lt;&lt;24);	/* pre indexing */
	return 0;
}

int
immfloat(int32 v)
{
	return (v &amp; 0xC03) == 0;	/* offset will fit in floating-point load/store */
}

int
immhalf(int32 v)
{
	if(v &gt;= 0 &amp;&amp; v &lt;= 0xff)
		return v|
			(1&lt;&lt;24)|	/* pre indexing */
			(1&lt;&lt;23);	/* pre indexing, up */
	if(v &gt;= -0xff &amp;&amp; v &lt; 0)
		return (-v &amp; 0xff)|
			(1&lt;&lt;24);	/* pre indexing */
	return 0;
}

int
aclass(Adr *a)
{
	Sym *s;
	int t;

	switch(a-&gt;type) {
	case D_NONE:
		return C_NONE;

	case D_REG:
		return C_REG;

	case D_REGREG:
		return C_REGREG;

	case D_SHIFT:
		return C_SHIFT;

	case D_FREG:
		return C_FREG;

	case D_FPCR:
		return C_FCR;

	case D_OREG:
		switch(a-&gt;name) {
		case D_EXTERN:
		case D_STATIC:
			if(a-&gt;sym == 0 || a-&gt;sym-&gt;name == 0) {
				print(&#34;null sym external\n&#34;);
				print(&#34;%D\n&#34;, a);
				return C_GOK;
			}
			s = a-&gt;sym;
			t = s-&gt;type;
			if(t == 0 || t == SXREF) {
				diag(&#34;undefined external: %s in %s&#34;,
					s-&gt;name, TNAME);
				s-&gt;type = SDATA;
			}
			if(dlm) {
				switch(t) {
				default:
					instoffset = s-&gt;value + a-&gt;offset + INITDAT;
					break;
				case SUNDEF:
				case STEXT:
				case SCONST:
				case SLEAF:
				case SSTRING:
					instoffset = s-&gt;value + a-&gt;offset;
					break;
				}
				return C_ADDR;
			}
			instoffset = s-&gt;value + a-&gt;offset - BIG;
			t = immaddr(instoffset);
			if(t) {
				if(immhalf(instoffset))
					return immfloat(t) ? C_HFEXT : C_HEXT;
				if(immfloat(t))
					return C_FEXT;
				return C_SEXT;
			}
			return C_LEXT;
		case D_AUTO:
			instoffset = autosize + a-&gt;offset;
			t = immaddr(instoffset);
			if(t){
				if(immhalf(instoffset))
					return immfloat(t) ? C_HFAUTO : C_HAUTO;
				if(immfloat(t))
					return C_FAUTO;
				return C_SAUTO;
			}
			return C_LAUTO;

		case D_PARAM:
			instoffset = autosize + a-&gt;offset + 4L;
			t = immaddr(instoffset);
			if(t){
				if(immhalf(instoffset))
					return immfloat(t) ? C_HFAUTO : C_HAUTO;
				if(immfloat(t))
					return C_FAUTO;
				return C_SAUTO;
			}
			return C_LAUTO;
		case D_NONE:
			instoffset = a-&gt;offset;
			t = immaddr(instoffset);
			if(t) {
				if(immhalf(instoffset))		 /* n.b. that it will also satisfy immrot */
					return immfloat(t) ? C_HFOREG : C_HOREG;
				if(immfloat(t))
					return C_FOREG; /* n.b. that it will also satisfy immrot */
				t = immrot(instoffset);
				if(t)
					return C_SROREG;
				if(immhalf(instoffset))
					return C_HOREG;
				return C_SOREG;
			}
			t = immrot(instoffset);
			if(t)
				return C_ROREG;
			return C_LOREG;
		}
		return C_GOK;

	case D_PSR:
		return C_PSR;

	case D_OCONST:
		switch(a-&gt;name) {
		case D_EXTERN:
		case D_STATIC:
			s = a-&gt;sym;
			t = s-&gt;type;
			if(t == 0 || t == SXREF) {
				diag(&#34;undefined external: %s in %s&#34;,
					s-&gt;name, TNAME);
				s-&gt;type = SDATA;
			}
			instoffset = s-&gt;value + a-&gt;offset + INITDAT;
			if(s-&gt;type == STEXT || s-&gt;type == SLEAF || s-&gt;type == SUNDEF) {
				instoffset = s-&gt;value + a-&gt;offset;
#ifdef CALLEEBX
				instoffset += fnpinc(s);
#else
				if(s-&gt;thumb)
					instoffset++;	// T bit
#endif
				return C_LCON;
			}
			return C_LCON;
		}
		return C_GOK;

	case D_FCONST:
		return C_FCON;

	case D_CONST:
	case D_CONST2:
		switch(a-&gt;name) {

		case D_NONE:
			instoffset = a-&gt;offset;
			if(a-&gt;reg != NREG)
				goto aconsize;

			t = immrot(instoffset);
			if(t)
				return C_RCON;
			t = immrot(~instoffset);
			if(t)
				return C_NCON;
			return C_LCON;

		case D_EXTERN:
		case D_STATIC:
			s = a-&gt;sym;
			if(s == S)
				break;
			t = s-&gt;type;
			switch(t) {
			case 0:
			case SXREF:
				diag(&#34;undefined external: %s in %s&#34;,
					s-&gt;name, TNAME);
				s-&gt;type = SDATA;
				break;
			case SUNDEF:
			case STEXT:
			case SSTRING:
			case SCONST:
			case SLEAF:
				instoffset = s-&gt;value + a-&gt;offset;
#ifdef CALLEEBX
				instoffset += fnpinc(s);
#else
				if(s-&gt;thumb)
					instoffset++;	// T bit
#endif
				return C_LCON;
			}
			if(!dlm) {
				instoffset = s-&gt;value + a-&gt;offset - BIG;
				t = immrot(instoffset);
				if(t &amp;&amp; instoffset != 0)
					return C_RECON;
			}
			instoffset = s-&gt;value + a-&gt;offset + INITDAT;
			return C_LCON;

		case D_AUTO:
			instoffset = autosize + a-&gt;offset;
			goto aconsize;

		case D_PARAM:
			instoffset = autosize + a-&gt;offset + 4L;
		aconsize:
			t = immrot(instoffset);
			if(t)
				return C_RACON;
			return C_LACON;
		}
		return C_GOK;

	case D_BRANCH:
		return C_SBRA;
	}
	return C_GOK;
}

Optab*
oplook(Prog *p)
{
	int a1, a2, a3, r;
	char *c1, *c3;
	Optab *o, *e;
	Optab *otab;
	Oprang *orange;

	if(thumb){
		otab = thumboptab;
		orange = thumboprange;
	}
	else{
		otab = optab;
		orange = oprange;
	}
	a1 = p-&gt;optab;
	if(a1)
		return otab+(a1-1);
	a1 = p-&gt;from.class;
	if(a1 == 0) {
		if(thumb)
			a1 = thumbaclass(&amp;p-&gt;from, p) + 1;
		else
			a1 = aclass(&amp;p-&gt;from) + 1;
		p-&gt;from.class = a1;
	}
	a1--;
	a3 = p-&gt;to.class;
	if(a3 == 0) {
		if(thumb)
			a3 = thumbaclass(&amp;p-&gt;to, p) + 1;
		else
			a3 = aclass(&amp;p-&gt;to) + 1;
		p-&gt;to.class = a3;
	}
	a3--;
	a2 = C_NONE;
	if(p-&gt;reg != NREG)
		a2 = C_REG;
	r = p-&gt;as;
	o = orange[r].start;
	if(o == 0) {
		a1 = opcross[repop[r]][a1][a2][a3];
		if(a1) {
			p-&gt;optab = a1+1;
			return otab+a1;
		}
		o = orange[r].stop; /* just generate an error */
	}
	if(debug[&#39;O&#39;]) {
		print(&#34;oplook %A %O %O %O\n&#34;,
			(int)p-&gt;as, a1, a2, a3);
		print(&#34;		%d %d\n&#34;, p-&gt;from.type, p-&gt;to.type);
	}
	e = orange[r].stop;
	c1 = xcmp[a1];
	c3 = xcmp[a3];
	for(; o&lt;e; o++)
		if(o-&gt;a2 == a2)
		if(c1[o-&gt;a1])
		if(c3[o-&gt;a3]) {
			p-&gt;optab = (o-otab)+1;
			return o;
		}
	diag(&#34;illegal combination %A %O %O %O, %d %d&#34;,
		p-&gt;as, a1, a2, a3, p-&gt;from.type, p-&gt;to.type);
	prasm(p);
	if(o == 0)
		o = otab;
	return o;
}

int
cmp(int a, int b)
{

	if(a == b)
		return 1;
	switch(a) {
	case C_LCON:
		if(b == C_RCON || b == C_NCON)
			return 1;
		break;
	case C_LACON:
		if(b == C_RACON)
			return 1;
		break;
	case C_LECON:
		if(b == C_RECON)
			return 1;
		break;

	case C_HFEXT:
		return b == C_HEXT || b == C_FEXT;
	case C_FEXT:
	case C_HEXT:
		return b == C_HFEXT;
	case C_SEXT:
		return cmp(C_HFEXT, b);
	case C_LEXT:
		return cmp(C_SEXT, b);

	case C_HFAUTO:
		return b == C_HAUTO || b == C_FAUTO;
	case C_FAUTO:
	case C_HAUTO:
		return b == C_HFAUTO;
	case C_SAUTO:
		return cmp(C_HFAUTO, b);
	case C_LAUTO:
		return cmp(C_SAUTO, b);

	case C_HFOREG:
		return b == C_HOREG || b == C_FOREG;
	case C_FOREG:
	case C_HOREG:
		return b == C_HFOREG;
	case C_SROREG:
		return cmp(C_SOREG, b) || cmp(C_ROREG, b);
	case C_SOREG:
	case C_ROREG:
		return b == C_SROREG || cmp(C_HFOREG, b);
	case C_LOREG:
		return cmp(C_SROREG, b);

	case C_LBRA:
		if(b == C_SBRA)
			return 1;
		break;
	case C_GBRA:
		if(b == C_SBRA || b == C_LBRA)
			return 1;

	case C_HREG:
		return cmp(C_SP, b) || cmp(C_PC, b);

	}
	return 0;
}

int
ocmp(const void *a1, const void *a2)
{
	Optab *p1, *p2;
	int n;

	p1 = (Optab*)a1;
	p2 = (Optab*)a2;
	n = p1-&gt;as - p2-&gt;as;
	if(n)
		return n;
	n = (p2-&gt;flag&amp;V4) - (p1-&gt;flag&amp;V4);	/* architecture version */
	if(n)
		return n;
	n = p1-&gt;a1 - p2-&gt;a1;
	if(n)
		return n;
	n = p1-&gt;a2 - p2-&gt;a2;
	if(n)
		return n;
	n = p1-&gt;a3 - p2-&gt;a3;
	if(n)
		return n;
	return 0;
}

void
buildop(void)
{
	int i, n, r;

	armv4 = !debug[&#39;h&#39;];
	for(i=0; i&lt;C_GOK; i++)
		for(n=0; n&lt;C_GOK; n++)
			xcmp[i][n] = cmp(n, i);
	for(n=0; optab[n].as != AXXX; n++)
		if((optab[n].flag &amp; V4) &amp;&amp; !armv4) {
			optab[n].as = AXXX;
			break;
		}
	qsort(optab, n, sizeof(optab[0]), ocmp);
	for(i=0; i&lt;n; i++) {
		r = optab[i].as;
		oprange[r].start = optab+i;
		while(optab[i].as == r)
			i++;
		oprange[r].stop = optab+i;
		i--;

		switch(r)
		{
		default:
			diag(&#34;unknown op in build: %A&#34;, r);
			errorexit();
		case AADD:
			oprange[AAND] = oprange[r];
			oprange[AEOR] = oprange[r];
			oprange[ASUB] = oprange[r];
			oprange[ARSB] = oprange[r];
			oprange[AADC] = oprange[r];
			oprange[ASBC] = oprange[r];
			oprange[ARSC] = oprange[r];
			oprange[AORR] = oprange[r];
			oprange[ABIC] = oprange[r];
			break;
		case ACMP:
			oprange[ATST] = oprange[r];
			oprange[ATEQ] = oprange[r];
			oprange[ACMN] = oprange[r];
			break;
		case AMVN:
			break;
		case ABEQ:
			oprange[ABNE] = oprange[r];
			oprange[ABCS] = oprange[r];
			oprange[ABHS] = oprange[r];
			oprange[ABCC] = oprange[r];
			oprange[ABLO] = oprange[r];
			oprange[ABMI] = oprange[r];
			oprange[ABPL] = oprange[r];
			oprange[ABVS] = oprange[r];
			oprange[ABVC] = oprange[r];
			oprange[ABHI] = oprange[r];
			oprange[ABLS] = oprange[r];
			oprange[ABGE] = oprange[r];
			oprange[ABLT] = oprange[r];
			oprange[ABGT] = oprange[r];
			oprange[ABLE] = oprange[r];
			break;
		case ASLL:
			oprange[ASRL] = oprange[r];
			oprange[ASRA] = oprange[r];
			break;
		case AMUL:
			oprange[AMULU] = oprange[r];
			break;
		case ADIV:
			oprange[AMOD] = oprange[r];
			oprange[AMODU] = oprange[r];
			oprange[ADIVU] = oprange[r];
			break;
		case AMOVW:
		case AMOVB:
		case AMOVBU:
		case AMOVH:
		case AMOVHU:
			break;
		case ASWPW:
			oprange[ASWPBU] = oprange[r];
			break;
		case AB:
		case ABL:
		case ABX:
		case ABXRET:
		case ASWI:
		case AWORD:
		case AMOVM:
		case ARFE:
		case ATEXT:
		case ACASE:
		case ABCASE:
			break;
		case AADDF:
			oprange[AADDD] = oprange[r];
			oprange[ASUBF] = oprange[r];
			oprange[ASUBD] = oprange[r];
			oprange[AMULF] = oprange[r];
			oprange[AMULD] = oprange[r];
			oprange[ADIVF] = oprange[r];
			oprange[ADIVD] = oprange[r];
			oprange[AMOVFD] = oprange[r];
			oprange[AMOVDF] = oprange[r];
			break;

		case ACMPF:
			oprange[ACMPD] = oprange[r];
			break;

		case AMOVF:
			oprange[AMOVD] = oprange[r];
			break;

		case AMOVFW:
			oprange[AMOVWF] = oprange[r];
			oprange[AMOVWD] = oprange[r];
			oprange[AMOVDW] = oprange[r];
			break;

		case AMULL:
			oprange[AMULA] = oprange[r];
			oprange[AMULAL] = oprange[r];
			oprange[AMULLU] = oprange[r];
			oprange[AMULALU] = oprange[r];
			break;
		case ALDREX:
		case ASTREX:
			break;
		}
	}
}

/*
void
buildrep(int x, int as)
{
	Opcross *p;
	Optab *e, *s, *o;
	int a1, a2, a3, n;

	if(C_NONE != 0 || C_REG != 1 || C_GOK &gt;= 32 || x &gt;= nelem(opcross)) {
		diag(&#34;assumptions fail in buildrep&#34;);
		errorexit();
	}
	repop[as] = x;
	p = (opcross + x);
	s = oprange[as].start;
	e = oprange[as].stop;
	for(o=e-1; o&gt;=s; o--) {
		n = o-optab;
		for(a2=0; a2&lt;2; a2++) {
			if(a2) {
				if(o-&gt;a2 == C_NONE)
					continue;
			} else
				if(o-&gt;a2 != C_NONE)
					continue;
			for(a1=0; a1&lt;32; a1++) {
				if(!xcmp[a1][o-&gt;a1])
					continue;
				for(a3=0; a3&lt;32; a3++)
					if(xcmp[a3][o-&gt;a3])
						(*p)[a1][a2][a3] = n;
			}
		}
	}
	oprange[as].start = 0;
}
*/

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
	r-&gt;m = nm = malloc(r-&gt;t*sizeof(uchar));
	r-&gt;a = na = malloc(r-&gt;t*sizeof(uint32));
	memmove(nm, m, t*sizeof(uchar));
	memmove(na, a, t*sizeof(uint32));
	free(m);
	free(a);
}

void
dynreloc(Sym *s, int32 v, int abs)
{
	int i, k, n;
	uchar *m;
	uint32 *a;
	Reloc *r;

	if(v&amp;3)
		diag(&#34;bad relocation address&#34;);
	v &gt;&gt;= 2;
	if(s != S &amp;&amp; s-&gt;type == SUNDEF)
		k = abs ? ABSU : RELU;
	else
		k = abs ? ABSD : RELD;
	/* Bprint(&amp;bso, &#34;R %s a=%ld(%lx) %d\n&#34;, s-&gt;name, a, a, k); */
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
	return  s-p+1;
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
