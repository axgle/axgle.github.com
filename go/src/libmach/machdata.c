<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/libmach/machdata.c</title>

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
  <h1 id="generatedHeader">Text file src/libmach/machdata.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno libmach/machdata.c
// http://code.google.com/p/inferno-os/source/browse/utils/libmach/machdata.c
//
// 	Copyright © 1994-1999 Lucent Technologies Inc.
// 	Power PC support Copyright © 1995-2004 C H Forsyth (forsyth@terzarima.net).
// 	Portions Copyright © 1997-1999 Vita Nuova Limited.
// 	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com).
// 	Revisions Copyright © 2000-2004 Lucent Technologies Inc. and others.
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

/*
 * Debugger utilities shared by at least two architectures
 */

#include &lt;u.h&gt;
#include &lt;libc.h&gt;
#include &lt;bio.h&gt;
#include &lt;mach.h&gt;

#define STARTSYM	&#34;_main&#34;
#define PROFSYM		&#34;_mainp&#34;
#define	FRAMENAME	&#34;.frame&#34;

extern	Machdata	mipsmach;

int	asstype = AMIPS;		/* disassembler type */
Machdata *machdata;		/* machine-dependent functions */

int
localaddr(Map *map, char *fn, char *var, uvlong *r, Rgetter rget)
{
	Symbol s;
	uvlong fp, pc, sp, link;

	if (!lookup(fn, 0, &amp;s)) {
		werrstr(&#34;function not found&#34;);
		return -1;
	}
	pc = rget(map, mach-&gt;pc);
	sp = rget(map, mach-&gt;sp);
	if(mach-&gt;link)
		link = rget(map, mach-&gt;link);
	else
		link = 0;
	fp = machdata-&gt;findframe(map, s.value, pc, sp, link);
	if (fp == 0) {
		werrstr(&#34;stack frame not found&#34;);
		return -1;
	}

	if (!var || !var[0]) {
		*r = fp;
		return 1;
	}

	if (findlocal(&amp;s, var, &amp;s) == 0) {
		werrstr(&#34;local variable not found&#34;);
		return -1;
	}

	switch (s.class) {
	case CAUTO:
		*r = fp - s.value;
		break;
	case CPARAM:		/* assume address size is stack width */
		*r = fp + s.value + mach-&gt;szaddr;
		break;
	default:
		werrstr(&#34;local variable not found: %d&#34;, s.class);
		return -1;
	}
	return 1;
}

/*
 * Print value v as s.name[+offset] if possible, or just v.
 */
int
symoff(char *buf, int n, uvlong v, int space)
{
	Symbol s;
	int r;
	int32 delta;

	r = delta = 0;		/* to shut compiler up */
	if (v) {
		r = findsym(v, space, &amp;s);
		if (r)
			delta = v-s.value;
		if (delta &lt; 0)
			delta = -delta;
	}
	if (v == 0 || r == 0)
		return snprint(buf, n, &#34;%llux&#34;, v);
	if (s.type != &#39;t&#39; &amp;&amp; s.type != &#39;T&#39; &amp;&amp; delta &gt;= 4096)
		return snprint(buf, n, &#34;%llux&#34;, v);
	else if (delta)
		return snprint(buf, n, &#34;%s+%#lux&#34;, s.name, delta);
	else
		return snprint(buf, n, &#34;%s&#34;, s.name);
}

/*
 *	Format floating point registers
 *
 *	Register codes in format field:
 *	&#39;X&#39; - print as 32-bit hexadecimal value
 *	&#39;F&#39; - 64-bit double register when modif == &#39;F&#39;; else 32-bit single reg
 *	&#39;f&#39; - 32-bit ieee float
 *	&#39;8&#39; - big endian 80-bit ieee extended float
 *	&#39;3&#39; - little endian 80-bit ieee extended float with hole in bytes 8&amp;9
 */
int
fpformat(Map *map, Reglist *rp, char *buf, int n, int modif)
{
	char reg[12];
	uint32 r;

	switch(rp-&gt;rformat)
	{
	case &#39;X&#39;:
		if (get4(map, rp-&gt;roffs, &amp;r) &lt; 0)
			return -1;
		snprint(buf, n, &#34;%lux&#34;, r);
		break;
	case &#39;F&#39;:	/* first reg of double reg pair */
		if (modif == &#39;F&#39;)
		if ((rp-&gt;rformat==&#39;F&#39;) || (((rp+1)-&gt;rflags&amp;RFLT) &amp;&amp; (rp+1)-&gt;rformat == &#39;f&#39;)) {
			if (get1(map, rp-&gt;roffs, (uchar *)reg, 8) &lt; 0)
				return -1;
			machdata-&gt;dftos(buf, n, reg);
			if (rp-&gt;rformat == &#39;F&#39;)
				return 1;
			return 2;
		}
			/* treat it like &#39;f&#39; */
		if (get1(map, rp-&gt;roffs, (uchar *)reg, 4) &lt; 0)
			return -1;
		machdata-&gt;sftos(buf, n, reg);
		break;
	case &#39;f&#39;:	/* 32 bit float */
		if (get1(map, rp-&gt;roffs, (uchar *)reg, 4) &lt; 0)
			return -1;
		machdata-&gt;sftos(buf, n, reg);
		break;
	case &#39;3&#39;:	/* little endian ieee 80 with hole in bytes 8&amp;9 */
		if (get1(map, rp-&gt;roffs, (uchar *)reg, 10) &lt; 0)
			return -1;
		memmove(reg+10, reg+8, 2);	/* open hole */
		memset(reg+8, 0, 2);		/* fill it */
		leieee80ftos(buf, n, reg);
		break;
	case &#39;8&#39;:	/* big-endian ieee 80 */
		if (get1(map, rp-&gt;roffs, (uchar *)reg, 10) &lt; 0)
			return -1;
		beieee80ftos(buf, n, reg);
		break;
	default:	/* unknown */
		break;
	}
	return 1;
}

char *
_hexify(char *buf, uint32 p, int zeros)
{
	uint32 d;

	d = p/16;
	if(d)
		buf = _hexify(buf, d, zeros-1);
	else
		while(zeros--)
			*buf++ = &#39;0&#39;;
	*buf++ = &#34;0123456789abcdef&#34;[p&amp;0x0f];
	return buf;
}

/*
 * These routines assume that if the number is representable
 * in IEEE floating point, it will be representable in the native
 * double format.  Naive but workable, probably.
 */
int
ieeedftos(char *buf, int n, uint32 h, uint32 l)
{
	double fr;
	int exp;

	if (n &lt;= 0)
		return 0;


	if(h &amp; (1L&lt;&lt;31)){
		*buf++ = &#39;-&#39;;
		h &amp;= ~(1L&lt;&lt;31);
	}else
		*buf++ = &#39; &#39;;
	n--;
	if(l == 0 &amp;&amp; h == 0)
		return snprint(buf, n, &#34;0.&#34;);
	exp = (h&gt;&gt;20) &amp; ((1L&lt;&lt;11)-1L);
	if(exp == 0)
		return snprint(buf, n, &#34;DeN(%.8lux%.8lux)&#34;, h, l);
	if(exp == ((1L&lt;&lt;11)-1L)){
		if(l==0 &amp;&amp; (h&amp;((1L&lt;&lt;20)-1L)) == 0)
			return snprint(buf, n, &#34;Inf&#34;);
		else
			return snprint(buf, n, &#34;NaN(%.8lux%.8lux)&#34;, h&amp;((1L&lt;&lt;20)-1L), l);
	}
	exp -= (1L&lt;&lt;10) - 2L;
	fr = l &amp; ((1L&lt;&lt;16)-1L);
	fr /= 1L&lt;&lt;16;
	fr += (l&gt;&gt;16) &amp; ((1L&lt;&lt;16)-1L);
	fr /= 1L&lt;&lt;16;
	fr += (h &amp; (1L&lt;&lt;20)-1L) | (1L&lt;&lt;20);
	fr /= 1L&lt;&lt;21;
	fr = ldexp(fr, exp);
	return snprint(buf, n, &#34;%.18g&#34;, fr);
}

int
ieeesftos(char *buf, int n, uint32 h)
{
	double fr;
	int exp;

	if (n &lt;= 0)
		return 0;

	if(h &amp; (1L&lt;&lt;31)){
		*buf++ = &#39;-&#39;;
		h &amp;= ~(1L&lt;&lt;31);
	}else
		*buf++ = &#39; &#39;;
	n--;
	if(h == 0)
		return snprint(buf, n, &#34;0.&#34;);
	exp = (h&gt;&gt;23) &amp; ((1L&lt;&lt;8)-1L);
	if(exp == 0)
		return snprint(buf, n, &#34;DeN(%.8lux)&#34;, h);
	if(exp == ((1L&lt;&lt;8)-1L)){
		if((h&amp;((1L&lt;&lt;23)-1L)) == 0)
			return snprint(buf, n, &#34;Inf&#34;);
		else
			return snprint(buf, n, &#34;NaN(%.8lux)&#34;, h&amp;((1L&lt;&lt;23)-1L));
	}
	exp -= (1L&lt;&lt;7) - 2L;
	fr = (h &amp; ((1L&lt;&lt;23)-1L)) | (1L&lt;&lt;23);
	fr /= 1L&lt;&lt;24;
	fr = ldexp(fr, exp);
	return snprint(buf, n, &#34;%.9g&#34;, fr);
}

int
beieeesftos(char *buf, int n, void *s)
{
	return ieeesftos(buf, n, beswal(*(uint32*)s));
}

int
beieeedftos(char *buf, int n, void *s)
{
	return ieeedftos(buf, n, beswal(*(uint32*)s), beswal(((uint32*)(s))[1]));
}

int
leieeesftos(char *buf, int n, void *s)
{
	return ieeesftos(buf, n, leswal(*(uint32*)s));
}

int
leieeedftos(char *buf, int n, void *s)
{
	return ieeedftos(buf, n, leswal(((uint32*)(s))[1]), leswal(*(uint32*)s));
}

/* packed in 12 bytes, with s[2]==s[3]==0; mantissa starts at s[4]*/
int
beieee80ftos(char *buf, int n, void *s)
{
	uchar *reg = (uchar*)s;
	int i;
	uint32 x;
	uchar ieee[8+8];	/* room for slop */
	uchar *p, *q;

	memset(ieee, 0, sizeof(ieee));
	/* sign */
	if(reg[0] &amp; 0x80)
		ieee[0] |= 0x80;

	/* exponent */
	x = ((reg[0]&amp;0x7F)&lt;&lt;8) | reg[1];
	if(x == 0)		/* number is ±0 */
		goto done;
	if(x == 0x7FFF){
		if(memcmp(reg+4, ieee+1, 8) == 0){ /* infinity */
			x = 2047;
		}else{				/* NaN */
			x = 2047;
			ieee[7] = 0x1;		/* make sure */
		}
		ieee[0] |= x&gt;&gt;4;
		ieee[1] |= (x&amp;0xF)&lt;&lt;4;
		goto done;
	}
	x -= 0x3FFF;		/* exponent bias */
	x += 1023;
	if(x &gt;= (1&lt;&lt;11) || ((reg[4]&amp;0x80)==0 &amp;&amp; x!=0))
		return snprint(buf, n, &#34;not in range&#34;);
	ieee[0] |= x&gt;&gt;4;
	ieee[1] |= (x&amp;0xF)&lt;&lt;4;

	/* mantissa */
	p = reg+4;
	q = ieee+1;
	for(i=0; i&lt;56; i+=8, p++, q++){	/* move one byte */
		x = (p[0]&amp;0x7F) &lt;&lt; 1;
		if(p[1] &amp; 0x80)
			x |= 1;
		q[0] |= x&gt;&gt;4;
		q[1] |= (x&amp;0xF)&lt;&lt;4;
	}
    done:
	return beieeedftos(buf, n, (void*)ieee);
}

int
leieee80ftos(char *buf, int n, void *s)
{
	int i;
	char *cp;
	char b[12];

	cp = (char*) s;
	for(i=0; i&lt;12; i++)
		b[11-i] = *cp++;
	return beieee80ftos(buf, n, b);
}

int
cisctrace(Map *map, uvlong pc, uvlong sp, uvlong link, Tracer trace)
{
	Symbol s;
	int found, i;
	uvlong opc, moved;

	USED(link);
	i = 0;
	opc = 0;
	while(pc &amp;&amp; opc != pc) {
		moved = pc2sp(pc);
		if (moved == ~0)
			break;
		found = findsym(pc, CTEXT, &amp;s);
		if (!found)
			break;
		if(strcmp(STARTSYM, s.name) == 0 || strcmp(PROFSYM, s.name) == 0)
			break;

		sp += moved;
		opc = pc;
		if (geta(map, sp, &amp;pc) &lt; 0)
			break;
		(*trace)(map, pc, sp, &amp;s);
		sp += mach-&gt;szaddr;	/*assumes address size = stack width*/
		if(++i &gt; 40)
			break;
	}
	return i;
}

int
risctrace(Map *map, uvlong pc, uvlong sp, uvlong link, Tracer trace)
{
	int i;
	Symbol s, f;
	uvlong oldpc;

	i = 0;
	while(findsym(pc, CTEXT, &amp;s)) {
		if(strcmp(STARTSYM, s.name) == 0 || strcmp(PROFSYM, s.name) == 0)
			break;

		if(pc == s.value)	/* at first instruction */
			f.value = 0;
		else if(findlocal(&amp;s, FRAMENAME, &amp;f) == 0)
			break;

		oldpc = pc;
		if(s.type == &#39;L&#39; || s.type == &#39;l&#39; || pc &lt;= s.value+mach-&gt;pcquant)
			pc = link;
		else
			if (geta(map, sp, &amp;pc) &lt; 0)
				break;

		if(pc == 0 || (pc == oldpc &amp;&amp; f.value == 0))
			break;

		sp += f.value;
		(*trace)(map, pc-8, sp, &amp;s);

		if(++i &gt; 40)
			break;
	}
	return i;
}

uvlong
ciscframe(Map *map, uvlong addr, uvlong pc, uvlong sp, uvlong link)
{
	Symbol s;
	uvlong moved;

	USED(link);
	for(;;) {
		moved = pc2sp(pc);
		if (moved  == ~0)
			break;
		sp += moved;
		findsym(pc, CTEXT, &amp;s);
		if (addr == s.value)
			return sp;
		if (geta(map, sp, &amp;pc) &lt; 0)
			break;
		sp += mach-&gt;szaddr;	/*assumes sizeof(addr) = stack width*/
	}
	return 0;
}

uvlong
riscframe(Map *map, uvlong addr, uvlong pc, uvlong sp, uvlong link)
{
	Symbol s, f;

	while (findsym(pc, CTEXT, &amp;s)) {
		if(strcmp(STARTSYM, s.name) == 0 || strcmp(PROFSYM, s.name) == 0)
			break;

		if(pc == s.value)	/* at first instruction */
			f.value = 0;
		else
		if(findlocal(&amp;s, FRAMENAME, &amp;f) == 0)
			break;

		sp += f.value;
		if (s.value == addr)
			return sp;

		if (s.type == &#39;L&#39; || s.type == &#39;l&#39; || pc-s.value &lt;= mach-&gt;szaddr*2)
			pc = link;
		else
		if (geta(map, sp-f.value, &amp;pc) &lt; 0)
			break;
	}
	return 0;
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
