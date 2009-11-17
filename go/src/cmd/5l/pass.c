<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5l/pass.c</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5l/pass.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5l/pass.c
// http://code.google.com/p/inferno-os/source/browse/utils/5l/pass.c
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
dodata(void)
{
	int i, t;
	Sym *s;
	Prog *p;
	int32 orig, v;

	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f dodata\n&#34;, cputime());
	Bflush(&amp;bso);
	for(p = datap; p != P; p = p-&gt;link) {
		s = p-&gt;from.sym;
		if(p-&gt;as == ADYNT || p-&gt;as == AINIT)
			s-&gt;value = dtype;
		if(s-&gt;type == SBSS)
			s-&gt;type = SDATA;
		if(s-&gt;type != SDATA)
			diag(&#34;initialize non-data (%d): %s\n%P&#34;,
				s-&gt;type, s-&gt;name, p);
		v = p-&gt;from.offset + p-&gt;reg;
		if(v &gt; s-&gt;value)
			diag(&#34;initialize bounds (%ld/%ld): %s\n%P&#34;,
				v, s-&gt;value, s-&gt;name, p);
		if((s-&gt;type == SBSS || s-&gt;type == SDATA) &amp;&amp; (p-&gt;to.type == D_CONST || p-&gt;to.type == D_OCONST) &amp;&amp; (p-&gt;to.name == D_EXTERN || p-&gt;to.name == D_STATIC)){
			s = p-&gt;to.sym;
			if(s != S &amp;&amp; (s-&gt;type == STEXT || s-&gt;type == SLEAF || s-&gt;type == SCONST || s-&gt;type == SXREF))
				s-&gt;fnptr = 1;
		}
	}

	if(debug[&#39;t&#39;]) {
		/*
		 * pull out string constants
		 */
		for(p = datap; p != P; p = p-&gt;link) {
			s = p-&gt;from.sym;
			if(p-&gt;to.type == D_SCONST)
				s-&gt;type = SSTRING;
		}
	}

	/*
	 * pass 1
	 *	assign &#39;small&#39; variables to data segment
	 *	(rational is that data segment is more easily
	 *	 addressed through offset on R12)
	 */
	orig = 0;
	for(i=0; i&lt;NHASH; i++)
	for(s = hash[i]; s != S; s = s-&gt;link) {
		t = s-&gt;type;
		if(t != SDATA &amp;&amp; t != SBSS)
			continue;
		v = s-&gt;value;
		if(v == 0) {
			diag(&#34;%s: no size&#34;, s-&gt;name);
			v = 1;
		}
		while(v &amp; 3)
			v++;
		s-&gt;value = v;
		if(v &gt; MINSIZ)
			continue;
		s-&gt;value = orig;
		orig += v;
		s-&gt;type = SDATA1;
	}

	/*
	 * pass 2
	 *	assign large &#39;data&#39; variables to data segment
	 */
	for(i=0; i&lt;NHASH; i++)
	for(s = hash[i]; s != S; s = s-&gt;link) {
		t = s-&gt;type;
		if(t != SDATA) {
			if(t == SDATA1)
				s-&gt;type = SDATA;
			continue;
		}
		v = s-&gt;value;
		s-&gt;value = orig;
		orig += v;
	}

	while(orig &amp; 7)
		orig++;
	datsize = orig;

	/*
	 * pass 3
	 *	everything else to bss segment
	 */
	for(i=0; i&lt;NHASH; i++)
	for(s = hash[i]; s != S; s = s-&gt;link) {
		if(s-&gt;type != SBSS)
			continue;
		v = s-&gt;value;
		s-&gt;value = orig;
		orig += v;
	}
	while(orig &amp; 7)
		orig++;
	bsssize = orig-datsize;

	xdefine(&#34;setR12&#34;, SDATA, 0L+BIG);
	xdefine(&#34;bdata&#34;, SDATA, 0L);
	xdefine(&#34;data&#34;, SBSS, 0);
	xdefine(&#34;edata&#34;, SDATA, datsize);
	xdefine(&#34;end&#34;, SBSS, datsize+bsssize);
	xdefine(&#34;etext&#34;, STEXT, 0L);
}

void
undef(void)
{
	int i;
	Sym *s;

	for(i=0; i&lt;NHASH; i++)
	for(s = hash[i]; s != S; s = s-&gt;link)
		if(s-&gt;type == SXREF)
			diag(&#34;%s: not defined&#34;, s-&gt;name);
}

Prog*
brchain(Prog *p)
{
	int i;

	for(i=0; i&lt;20; i++) {
		if(p == P || p-&gt;as != AB)
			return p;
		p = p-&gt;cond;
	}
	return P;
}

int
relinv(int a)
{
	switch(a) {
	case ABEQ:	return ABNE;
	case ABNE:	return ABEQ;
	case ABCS:	return ABCC;
	case ABHS:	return ABLO;
	case ABCC:	return ABCS;
	case ABLO:	return ABHS;
	case ABMI:	return ABPL;
	case ABPL:	return ABMI;
	case ABVS:	return ABVC;
	case ABVC:	return ABVS;
	case ABHI:	return ABLS;
	case ABLS:	return ABHI;
	case ABGE:	return ABLT;
	case ABLT:	return ABGE;
	case ABGT:	return ABLE;
	case ABLE:	return ABGT;
	}
	diag(&#34;unknown relation: %s&#34;, anames[a]);
	return a;
}

void
follow(void)
{
	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f follow\n&#34;, cputime());
	Bflush(&amp;bso);

	firstp = prg();
	lastp = firstp;
	xfol(textp);

	firstp = firstp-&gt;link;
	lastp-&gt;link = P;
}

void
xfol(Prog *p)
{
	Prog *q, *r;
	int a, i;

loop:
	if(p == P)
		return;
	setarch(p);
	a = p-&gt;as;
	if(a == ATEXT)
		curtext = p;
	if(!curtext-&gt;from.sym-&gt;reachable) {
		p = p-&gt;cond;
		goto loop;
	}
	if(a == AB) {
		q = p-&gt;cond;
		if(q != P) {
			p-&gt;mark |= FOLL;
			p = q;
			if(!(p-&gt;mark &amp; FOLL))
				goto loop;
		}
	}
	if(p-&gt;mark &amp; FOLL) {
		for(i=0,q=p; i&lt;4; i++,q=q-&gt;link) {
			if(q == lastp)
				break;
			a = q-&gt;as;
			if(a == ANOP) {
				i--;
				continue;
			}
			if(a == AB || (a == ARET &amp;&amp; q-&gt;scond == 14) || a == ARFE)
				goto copy;
			if(!q-&gt;cond || (q-&gt;cond-&gt;mark&amp;FOLL))
				continue;
			if(a != ABEQ &amp;&amp; a != ABNE)
				continue;
		copy:
			for(;;) {
				r = prg();
				*r = *p;
				if(!(r-&gt;mark&amp;FOLL))
					print(&#34;cant happen 1\n&#34;);
				r-&gt;mark |= FOLL;
				if(p != q) {
					p = p-&gt;link;
					lastp-&gt;link = r;
					lastp = r;
					continue;
				}
				lastp-&gt;link = r;
				lastp = r;
				if(a == AB || (a == ARET &amp;&amp; q-&gt;scond == 14) || a == ARFE)
					return;
				r-&gt;as = ABNE;
				if(a == ABNE)
					r-&gt;as = ABEQ;
				r-&gt;cond = p-&gt;link;
				r-&gt;link = p-&gt;cond;
				if(!(r-&gt;link-&gt;mark&amp;FOLL))
					xfol(r-&gt;link);
				if(!(r-&gt;cond-&gt;mark&amp;FOLL))
					print(&#34;cant happen 2\n&#34;);
				return;
			}
		}
		a = AB;
		q = prg();
		q-&gt;as = a;
		q-&gt;line = p-&gt;line;
		q-&gt;to.type = D_BRANCH;
		q-&gt;to.offset = p-&gt;pc;
		q-&gt;cond = p;
		p = q;
	}
	p-&gt;mark |= FOLL;
	lastp-&gt;link = p;
	lastp = p;
	if(a == AB || (a == ARET &amp;&amp; p-&gt;scond == 14) || a == ARFE){
		return;
	}
	if(p-&gt;cond != P)
	if(a != ABL &amp;&amp; a != ABX &amp;&amp; p-&gt;link != P) {
		q = brchain(p-&gt;link);
		if(a != ATEXT &amp;&amp; a != ABCASE)
		if(q != P &amp;&amp; (q-&gt;mark&amp;FOLL)) {
			p-&gt;as = relinv(a);
			p-&gt;link = p-&gt;cond;
			p-&gt;cond = q;
		}
		xfol(p-&gt;link);
		q = brchain(p-&gt;cond);
		if(q == P)
			q = p-&gt;cond;
		if(q-&gt;mark&amp;FOLL) {
			p-&gt;cond = q;
			return;
		}
		p = q;
		goto loop;
	}
	p = p-&gt;link;
	goto loop;
}

void
patch(void)
{
	int32 c, vexit;
	Prog *p, *q;
	Sym *s, *s1;
	int a;

	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f patch\n&#34;, cputime());
	Bflush(&amp;bso);
	mkfwd();
	s = lookup(&#34;exit&#34;, 0);
	vexit = s-&gt;value;
	for(p = firstp; p != P; p = p-&gt;link) {
		setarch(p);
		a = p-&gt;as;
		if(a == ATEXT)
			curtext = p;
		if(seenthumb &amp;&amp; a == ABL){
			// if((s = p-&gt;to.sym) != S &amp;&amp; (s1 = curtext-&gt;from.sym) != S)
			//	print(&#34;%s calls %s\n&#34;, s1-&gt;name, s-&gt;name);
			 if((s = p-&gt;to.sym) != S &amp;&amp; (s1 = curtext-&gt;from.sym) != S &amp;&amp; s-&gt;thumb != s1-&gt;thumb)
				s-&gt;foreign = 1;
		}
		if((a == ABL || a == ABX || a == AB || a == ARET) &amp;&amp;
		   p-&gt;to.type != D_BRANCH &amp;&amp; p-&gt;to.sym != S) {
			s = p-&gt;to.sym;
			switch(s-&gt;type) {
			default:
				diag(&#34;undefined: %s&#34;, s-&gt;name);
				s-&gt;type = STEXT;
				s-&gt;value = vexit;
				continue;	// avoid more error messages
			case STEXT:
				p-&gt;to.offset = s-&gt;value;
				p-&gt;to.type = D_BRANCH;
				break;
			case SUNDEF:
				if(p-&gt;as != ABL)
					diag(&#34;help: SUNDEF in AB || ARET&#34;);
				p-&gt;to.offset = 0;
				p-&gt;to.type = D_BRANCH;
				p-&gt;cond = UP;
				break;
			}
		}
		if(p-&gt;to.type != D_BRANCH || p-&gt;cond == UP)
			continue;
		c = p-&gt;to.offset;
		for(q = firstp; q != P;) {
			if(q-&gt;forwd != P)
			if(c &gt;= q-&gt;forwd-&gt;pc) {
				q = q-&gt;forwd;
				continue;
			}
			if(c == q-&gt;pc)
				break;
			q = q-&gt;link;
		}
		if(q == P) {
			diag(&#34;branch out of range %ld\n%P&#34;, c, p);
			p-&gt;to.type = D_NONE;
		}
		p-&gt;cond = q;
	}

	for(p = firstp; p != P; p = p-&gt;link) {
		setarch(p);
		a = p-&gt;as;
		if(p-&gt;as == ATEXT)
			curtext = p;
		if(seenthumb &amp;&amp; a == ABL) {
#ifdef CALLEEBX
			if(0)
				{}
#else
			if((s = p-&gt;to.sym) != S &amp;&amp; (s-&gt;foreign || s-&gt;fnptr))
				p-&gt;as = ABX;
#endif
			else if(p-&gt;to.type == D_OREG)
				p-&gt;as = ABX;
		}
		if(p-&gt;cond != P &amp;&amp; p-&gt;cond != UP) {
			p-&gt;cond = brloop(p-&gt;cond);
			if(p-&gt;cond != P)
			if(p-&gt;to.type == D_BRANCH)
				p-&gt;to.offset = p-&gt;cond-&gt;pc;
		}
	}
}

#define	LOG	5
void
mkfwd(void)
{
	Prog *p;
	int32 dwn[LOG], cnt[LOG], i;
	Prog *lst[LOG];

	for(i=0; i&lt;LOG; i++) {
		if(i == 0)
			cnt[i] = 1; else
			cnt[i] = LOG * cnt[i-1];
		dwn[i] = 1;
		lst[i] = P;
	}
	i = 0;
	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT)
			curtext = p;
		i--;
		if(i &lt; 0)
			i = LOG-1;
		p-&gt;forwd = P;
		dwn[i]--;
		if(dwn[i] &lt;= 0) {
			dwn[i] = cnt[i];
			if(lst[i] != P)
				lst[i]-&gt;forwd = p;
			lst[i] = p;
		}
	}
}

Prog*
brloop(Prog *p)
{
	Prog *q;
	int c;

	for(c=0; p!=P;) {
		if(p-&gt;as != AB)
			return p;
		q = p-&gt;cond;
		if(q &lt;= p) {
			c++;
			if(q == p || c &gt; 5000)
				break;
		}
		p = q;
	}
	return P;
}

int32
atolwhex(char *s)
{
	int32 n;
	int f;

	n = 0;
	f = 0;
	while(*s == &#39; &#39; || *s == &#39;\t&#39;)
		s++;
	if(*s == &#39;-&#39; || *s == &#39;+&#39;) {
		if(*s++ == &#39;-&#39;)
			f = 1;
		while(*s == &#39; &#39; || *s == &#39;\t&#39;)
			s++;
	}
	if(s[0]==&#39;0&#39; &amp;&amp; s[1]){
		if(s[1]==&#39;x&#39; || s[1]==&#39;X&#39;){
			s += 2;
			for(;;){
				if(*s &gt;= &#39;0&#39; &amp;&amp; *s &lt;= &#39;9&#39;)
					n = n*16 + *s++ - &#39;0&#39;;
				else if(*s &gt;= &#39;a&#39; &amp;&amp; *s &lt;= &#39;f&#39;)
					n = n*16 + *s++ - &#39;a&#39; + 10;
				else if(*s &gt;= &#39;A&#39; &amp;&amp; *s &lt;= &#39;F&#39;)
					n = n*16 + *s++ - &#39;A&#39; + 10;
				else
					break;
			}
		} else
			while(*s &gt;= &#39;0&#39; &amp;&amp; *s &lt;= &#39;7&#39;)
				n = n*8 + *s++ - &#39;0&#39;;
	} else
		while(*s &gt;= &#39;0&#39; &amp;&amp; *s &lt;= &#39;9&#39;)
			n = n*10 + *s++ - &#39;0&#39;;
	if(f)
		n = -n;
	return n;
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

#define Reachable(n)	if((s = lookup(n, 0)) != nil) s-&gt;used++

static void
rused(Adr *a)
{
	Sym *s = a-&gt;sym;

	if(s == S)
		return;
	if(a-&gt;type == D_OREG || a-&gt;type == D_OCONST || a-&gt;type == D_CONST){
		if(a-&gt;name == D_EXTERN || a-&gt;name == D_STATIC){
			if(s-&gt;used == 0)
				s-&gt;used = 1;
		}
	}
	else if(a-&gt;type == D_BRANCH){
		if(s-&gt;used == 0)
			s-&gt;used = 1;
	}
}

void
reachable()
{
	Prog *p, *prev, *prevt, *nextt, *q;
	Sym *s, *s0;
	int i, todo;
	char *a;

	Reachable(&#34;_div&#34;);
	Reachable(&#34;_divu&#34;);
	Reachable(&#34;_mod&#34;);
	Reachable(&#34;_modu&#34;);
	a = INITENTRY;
	if(*a &gt;= &#39;0&#39; &amp;&amp; *a &lt;= &#39;9&#39;)
		return;
	s = lookup(a, 0);
	if(s == nil)
		return;
	if(s-&gt;type == 0){
		s-&gt;used = 1;	// to stop asm complaining
		for(p = firstp; p != P &amp;&amp; p-&gt;as != ATEXT; p = p-&gt;link)
			;
		if(p == nil)
			return;
		s = p-&gt;from.sym;
	}
	s-&gt;used = 1;
	do{
		todo = 0;
		for(p = firstp; p != P; p = p-&gt;link){
			if(p-&gt;as == ATEXT &amp;&amp; (s0 = p-&gt;from.sym)-&gt;used == 1){
				todo = 1;
				for(q = p-&gt;link; q != P &amp;&amp; q-&gt;as != ATEXT; q = q-&gt;link){
					rused(&amp;q-&gt;from);
					rused(&amp;q-&gt;to);
				}
				s0-&gt;used = 2;
			}
		}
		for(p = datap; p != P; p = p-&gt;link){
			if((s0 = p-&gt;from.sym)-&gt;used == 1){
				todo = 1;
				for(q = p; q != P; q = q-&gt;link){	// data can be scattered
					if(q-&gt;from.sym == s0)
						rused(&amp;q-&gt;to);
				}
				s0-&gt;used = 2;
			}
		}
	}while(todo);
	prev = nil;
	prevt = nextt = nil;
	for(p = firstp; p != P; ){
		if(p-&gt;as == ATEXT){
			prevt = nextt;
			nextt = p;
		}
		if(p-&gt;as == ATEXT &amp;&amp; (s0 = p-&gt;from.sym)-&gt;used == 0){
			s0-&gt;type = SREMOVED;
			for(q = p-&gt;link; q != P &amp;&amp; q-&gt;as != ATEXT; q = q-&gt;link)
				;
			if(q != p-&gt;cond)
				diag(&#34;bad ptr in reachable()&#34;);
			if(prev == nil)
				firstp = q;
			else
				prev-&gt;link = q;
			if(q == nil)
				lastp = prev;
			if(prevt == nil)
				textp = q;
			else
				prevt-&gt;cond = q;
			if(q == nil)
				etextp = prevt;
			nextt = prevt;
			if(debug[&#39;V&#39;])
				print(&#34;%s unused\n&#34;, s0-&gt;name);
			p = q;
		}
		else{
			prev = p;
			p = p-&gt;link;
		}
	}
	prevt = nil;
	for(p = datap; p != nil; ){
		if((s0 = p-&gt;from.sym)-&gt;used == 0){
			s0-&gt;type = SREMOVED;
			prev = prevt;
			for(q = p; q != nil; q = q-&gt;link){
				if(q-&gt;from.sym == s0){
					if(prev == nil)
						datap = q-&gt;link;
					else
						prev-&gt;link = q-&gt;link;
				}
				else
					prev = q;
			}
			if(debug[&#39;V&#39;])
				print(&#34;%s unused (data)\n&#34;, s0-&gt;name);
			p = prevt-&gt;link;
		}
		else{
			prevt = p;
			p = p-&gt;link;
		}
	}
	for(i=0; i&lt;NHASH; i++){
		for(s = hash[i]; s != S; s = s-&gt;link){
			if(s-&gt;used == 0)
				s-&gt;type = SREMOVED;
		}
	}
}

static void
fused(Adr *a, Prog *p, Prog *ct)
{
	Sym *s = a-&gt;sym;
	Use *u;

	if(s == S)
		return;
	if(a-&gt;type == D_OREG || a-&gt;type == D_OCONST || a-&gt;type == D_CONST){
		if(a-&gt;name == D_EXTERN || a-&gt;name == D_STATIC){
			u = malloc(sizeof(Use));
			u-&gt;p = p;
			u-&gt;ct = ct;
			u-&gt;link = s-&gt;use;
			s-&gt;use = u;
		}
	}
	else if(a-&gt;type == D_BRANCH){
		u = malloc(sizeof(Use));
		u-&gt;p = p;
		u-&gt;ct = ct;
		u-&gt;link = s-&gt;use;
		s-&gt;use = u;
	}
}

static int
ckfpuse(Prog *p, Prog *ct, Sym *fp, Sym *r)
{
	int reg;

	USED(fp);
	USED(ct);
	if(p-&gt;from.sym == r &amp;&amp; p-&gt;as == AMOVW &amp;&amp; (p-&gt;from.type == D_CONST || p-&gt;from.type == D_OREG) &amp;&amp; p-&gt;reg == NREG &amp;&amp; p-&gt;to.type == D_REG){
		reg = p-&gt;to.reg;
		for(p = p-&gt;link; p != P &amp;&amp; p-&gt;as != ATEXT; p = p-&gt;link){
			if((p-&gt;as == ABL || p-&gt;as == ABX) &amp;&amp; p-&gt;to.type == D_OREG &amp;&amp; p-&gt;to.reg == reg)
				return 1;
			if(!debug[&#39;F&#39;] &amp;&amp; (isbranch(p) || p-&gt;as == ARET)){
				// print(&#34;%s: branch %P in %s\n&#34;, fp-&gt;name, p, ct-&gt;from.sym-&gt;name);
				return 0;
			}
			if((p-&gt;from.type == D_REG || p-&gt;from.type == D_OREG) &amp;&amp; p-&gt;from.reg == reg){
				if(!debug[&#39;F&#39;] &amp;&amp; p-&gt;to.type != D_REG){
					// print(&#34;%s: store %P in %s\n&#34;, fp-&gt;name, p, ct-&gt;from.sym-&gt;name);
					return 0;
				}
				reg = p-&gt;to.reg;
			}
		}
	}
	// print(&#34;%s: no MOVW O(R), R\n&#34;, fp-&gt;name);
	return debug[&#39;F&#39;];
}

static void
setfpuse(Prog *p, Sym *fp, Sym *r)
{
	int reg;

	if(p-&gt;from.sym == r &amp;&amp; p-&gt;as == AMOVW &amp;&amp; (p-&gt;from.type == D_CONST || p-&gt;from.type == D_OREG) &amp;&amp; p-&gt;reg == NREG &amp;&amp; p-&gt;to.type == D_REG){
		reg = p-&gt;to.reg;
		for(p = p-&gt;link; p != P &amp;&amp; p-&gt;as != ATEXT; p = p-&gt;link){
			if((p-&gt;as == ABL || p-&gt;as == ABX) &amp;&amp; p-&gt;to.type == D_OREG &amp;&amp; p-&gt;to.reg == reg){
				fp-&gt;fnptr = 0;
				p-&gt;as = ABL;	// safe to do so
// print(&#34;simplified %s call\n&#34;, fp-&gt;name);
				break;
			}
			if(!debug[&#39;F&#39;] &amp;&amp; (isbranch(p) || p-&gt;as == ARET))
				diag(&#34;bad setfpuse call&#34;);
			if((p-&gt;from.type == D_REG || p-&gt;from.type == D_OREG) &amp;&amp; p-&gt;from.reg == reg){
				if(!debug[&#39;F&#39;] &amp;&amp; p-&gt;to.type != D_REG)
					diag(&#34;bad setfpuse call&#34;);
				reg = p-&gt;to.reg;
			}
		}
	}
}

static int
cksymuse(Sym *s, int t)
{
	Prog *p;

	for(p = datap; p != P; p = p-&gt;link){
		if(p-&gt;from.sym == s &amp;&amp; p-&gt;to.sym != nil &amp;&amp; strcmp(p-&gt;to.sym-&gt;name, &#34;.string&#34;) != 0 &amp;&amp; p-&gt;to.sym-&gt;thumb != t){
			// print(&#34;%s %s %d %d &#34;, p-&gt;from.sym-&gt;name, p-&gt;to.sym-&gt;name, p-&gt;to.sym-&gt;thumb, t);
			return 0;
		}
	}
	return 1;
}

/* check the use of s at the given point */
static int
ckuse(Sym *s, Sym *s0, Use *u)
{
	Sym *s1;

	s1 = u-&gt;p-&gt;from.sym;
// print(&#34;ckuse %s %s %s\n&#34;, s-&gt;name, s0-&gt;name, s1 ? s1-&gt;name : &#34;nil&#34;);
	if(u-&gt;ct == nil){	/* in data area */
		if(s0 == s &amp;&amp; !cksymuse(s1, s0-&gt;thumb)){
			// print(&#34;%s: cksymuse fails\n&#34;, s0-&gt;name);
			return 0;
		}
		for(u = s1-&gt;use; u != U; u = u-&gt;link)
			if(!ckuse(s1, s0, u))
				return 0;
	}
	else{		/* in text area */
		if(u-&gt;ct-&gt;from.sym-&gt;thumb != s0-&gt;thumb){
			// print(&#34;%s(%d): foreign call %s(%d)\n&#34;, s0-&gt;name, s0-&gt;thumb, u-&gt;ct-&gt;from.sym-&gt;name, u-&gt;ct-&gt;from.sym-&gt;thumb);
			return 0;
		}
		return ckfpuse(u-&gt;p, u-&gt;ct, s0, s);
	}
	return 1;
}

static void
setuse(Sym *s, Sym *s0, Use *u)
{
	Sym *s1;

	s1 = u-&gt;p-&gt;from.sym;
	if(u-&gt;ct == nil){	/* in data area */
		for(u = s1-&gt;use; u != U; u = u-&gt;link)
			setuse(s1, s0, u);
	}
	else{		/* in text area */
		setfpuse(u-&gt;p, s0, s);
	}
}

/* detect BX O(R) which can be done as BL O(R) */
void
fnptrs()
{
	int i;
	Sym *s;
	Prog *p;
	Use *u;

	for(i=0; i&lt;NHASH; i++){
		for(s = hash[i]; s != S; s = s-&gt;link){
			if(s-&gt;fnptr &amp;&amp; (s-&gt;type == STEXT || s-&gt;type == SLEAF || s-&gt;type == SCONST)){
				// print(&#34;%s : fnptr %d %d\n&#34;, s-&gt;name, s-&gt;thumb, s-&gt;foreign);
			}
		}
	}
	/* record use of syms */
	for(p = firstp; p != P; p = p-&gt;link){
		if(p-&gt;as == ATEXT)
			curtext = p;
		else{
			fused(&amp;p-&gt;from, p, curtext);
			fused(&amp;p-&gt;to, p, curtext);
		}
	}
	for(p = datap; p != P; p = p-&gt;link)
		fused(&amp;p-&gt;to, p, nil);

	/* now look for fn ptrs */
	for(i=0; i&lt;NHASH; i++){
		for(s = hash[i]; s != S; s = s-&gt;link){
			if(s-&gt;fnptr &amp;&amp; (s-&gt;type == STEXT || s-&gt;type == SLEAF || s-&gt;type == SCONST)){
				for(u = s-&gt;use; u != U; u = u-&gt;link){
					if(!ckuse(s, s, u))
						break;
				}
				if(u == U){		// can simplify
					for(u = s-&gt;use; u != U; u = u-&gt;link)
						setuse(s, s, u);
				}
			}
		}
	}

	/*  now free Use structures */
}

void
import(void)
{
	int i;
	Sym *s;

	for(i = 0; i &lt; NHASH; i++)
		for(s = hash[i]; s != S; s = s-&gt;link)
			if(s-&gt;sig != 0 &amp;&amp; s-&gt;type == SXREF &amp;&amp; (nimports == 0 || s-&gt;subtype == SIMPORT)){
				undefsym(s);
				Bprint(&amp;bso, &#34;IMPORT: %s sig=%lux v=%ld\n&#34;, s-&gt;name, s-&gt;sig, s-&gt;value);
			}
}

void
ckoff(Sym *s, int32 v)
{
	if(v &lt; 0 || v &gt;= 1&lt;&lt;Roffset)
		diag(&#34;relocation offset %ld for %s out of range&#34;, v, s-&gt;name);
}

Prog*
newdata(Sym *s, int o, int w, int t)
{
	Prog *p;

	p = prg();
	p-&gt;link = datap;
	datap = p;
	p-&gt;as = ADATA;
	p-&gt;reg = w;
	p-&gt;from.type = D_OREG;
	p-&gt;from.name = t;
	p-&gt;from.sym = s;
	p-&gt;from.offset = o;
	p-&gt;to.type = D_CONST;
	p-&gt;to.name = D_NONE;
	s-&gt;data = p;
	return p;
}

void
export(void)
{
	int i, j, n, off, nb, sv, ne;
	Sym *s, *et, *str, **esyms;
	Prog *p;
	char buf[NSNAME], *t;

	n = 0;
	for(i = 0; i &lt; NHASH; i++)
		for(s = hash[i]; s != S; s = s-&gt;link)
			if(s-&gt;sig != 0 &amp;&amp; s-&gt;type != SXREF &amp;&amp; s-&gt;type != SUNDEF &amp;&amp; (nexports == 0 || s-&gt;subtype == SEXPORT))
				n++;
	esyms = malloc(n*sizeof(Sym*));
	ne = n;
	n = 0;
	for(i = 0; i &lt; NHASH; i++)
		for(s = hash[i]; s != S; s = s-&gt;link)
			if(s-&gt;sig != 0 &amp;&amp; s-&gt;type != SXREF &amp;&amp; s-&gt;type != SUNDEF &amp;&amp; (nexports == 0 || s-&gt;subtype == SEXPORT))
				esyms[n++] = s;
	for(i = 0; i &lt; ne-1; i++)
		for(j = i+1; j &lt; ne; j++)
			if(strcmp(esyms[i]-&gt;name, esyms[j]-&gt;name) &gt; 0){
				s = esyms[i];
				esyms[i] = esyms[j];
				esyms[j] = s;
			}

	nb = 0;
	off = 0;
	et = lookup(EXPTAB, 0);
	if(et-&gt;type != 0 &amp;&amp; et-&gt;type != SXREF)
		diag(&#34;%s already defined&#34;, EXPTAB);
	et-&gt;type = SDATA;
	str = lookup(&#34;.string&#34;, 0);
	if(str-&gt;type == 0)
		str-&gt;type = SDATA;
	sv = str-&gt;value;
	for(i = 0; i &lt; ne; i++){
		s = esyms[i];
		Bprint(&amp;bso, &#34;EXPORT: %s sig=%lux t=%d\n&#34;, s-&gt;name, s-&gt;sig, s-&gt;type);

		/* signature */
		p = newdata(et, off, sizeof(int32), D_EXTERN);
		off += sizeof(int32);
		p-&gt;to.offset = s-&gt;sig;

		/* address */
		p = newdata(et, off, sizeof(int32), D_EXTERN);
		off += sizeof(int32);
		p-&gt;to.name = D_EXTERN;
		p-&gt;to.sym = s;

		/* string */
		t = s-&gt;name;
		n = strlen(t)+1;
		for(;;){
			buf[nb++] = *t;
			sv++;
			if(nb &gt;= NSNAME){
				p = newdata(str, sv-NSNAME, NSNAME, D_STATIC);
				p-&gt;to.type = D_SCONST;
				p-&gt;to.sval = malloc(NSNAME);
				memmove(p-&gt;to.sval, buf, NSNAME);
				nb = 0;
			}
			if(*t++ == 0)
				break;
		}

		/* name */
		p = newdata(et, off, sizeof(int32), D_EXTERN);
		off += sizeof(int32);
		p-&gt;to.name = D_STATIC;
		p-&gt;to.sym = str;
		p-&gt;to.offset = sv-n;
	}

	if(nb &gt; 0){
		p = newdata(str, sv-nb, nb, D_STATIC);
		p-&gt;to.type = D_SCONST;
		p-&gt;to.sval = malloc(NSNAME);
		memmove(p-&gt;to.sval, buf, nb);
	}

	for(i = 0; i &lt; 3; i++){
		newdata(et, off, sizeof(int32), D_EXTERN);
		off += sizeof(int32);
	}
	et-&gt;value = off;
	if(sv == 0)
		sv = 1;
	str-&gt;value = sv;
	exports = ne;
	free(esyms);
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
