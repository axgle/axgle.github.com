<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/6l/pass.c</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/6l/pass.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/6l/pass.c
// http://code.google.com/p/inferno-os/source/browse/utils/6l/pass.c
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

// see ../../runtime/proc.c:/StackGuard
enum
{
	StackSmall = 128,
	StackBig = 4096,
};

void
dodata(void)
{
	int i;
	Sym *s;
	Prog *p;
	int32 t, u;

	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f dodata\n&#34;, cputime());
	Bflush(&amp;bso);
	for(p = datap; p != P; p = p-&gt;link) {
		curtext = p;	// for diag messages
		s = p-&gt;from.sym;
		if(p-&gt;as == ADYNT || p-&gt;as == AINIT)
			s-&gt;value = dtype;
		if(s-&gt;type == SBSS)
			s-&gt;type = SDATA;
		if(s-&gt;type != SDATA)
			diag(&#34;initialize non-data (%d): %s\n%P&#34;,
				s-&gt;type, s-&gt;name, p);
		t = p-&gt;from.offset + p-&gt;width;
		if(t &gt; s-&gt;value)
			diag(&#34;initialize bounds (%lld): %s\n%P&#34;,
				s-&gt;value, s-&gt;name, p);
	}
	/* allocate small guys */
	datsize = 0;
	for(i=0; i&lt;NHASH; i++)
	for(s = hash[i]; s != S; s = s-&gt;link) {
		if(!s-&gt;reachable)
			continue;
		if(s-&gt;type != SDATA)
		if(s-&gt;type != SBSS)
			continue;
		t = s-&gt;value;
		if(t == 0 &amp;&amp; s-&gt;name[0] != &#39;.&#39;) {
			diag(&#34;%s: no size&#34;, s-&gt;name);
			t = 1;
		}
		t = rnd(t, 4);
		s-&gt;value = t;
		if(t &gt; MINSIZ)
			continue;
		if(t &gt;= 8)
			datsize = rnd(datsize, 8);
		s-&gt;size = t;
		s-&gt;value = datsize;
		datsize += t;
		s-&gt;type = SDATA1;
	}

	/* allocate the rest of the data */
	for(i=0; i&lt;NHASH; i++)
	for(s = hash[i]; s != S; s = s-&gt;link) {
		if(!s-&gt;reachable)
			continue;
		if(s-&gt;type != SDATA) {
			if(s-&gt;type == SDATA1)
				s-&gt;type = SDATA;
			continue;
		}
		t = s-&gt;value;
		if(t &gt;= 8)
			datsize = rnd(datsize, 8);
		s-&gt;size = t;
		s-&gt;value = datsize;
		datsize += t;
	}
	if(datsize)
		datsize = rnd(datsize, 8);

	if(debug[&#39;j&#39;]) {
		/*
		 * pad data with bss that fits up to next
		 * 8k boundary, then push data to 8k
		 */
		u = rnd(datsize, 8192);
		u -= datsize;
		for(i=0; i&lt;NHASH; i++)
		for(s = hash[i]; s != S; s = s-&gt;link) {
			if(!s-&gt;reachable)
				continue;
			if(s-&gt;type != SBSS)
				continue;
			t = s-&gt;value;
			if(t &gt; u)
				continue;
			u -= t;
			s-&gt;size = t;
			s-&gt;value = datsize;
			s-&gt;type = SDATA;
			datsize += t;
		}
		datsize += u;
	}
}

void
dobss(void)
{
	int i;
	Sym *s;
	int32 t;

	if(dynptrsize &gt; 0) {
		/* dynamic pointer section between data and bss */
		datsize = rnd(datsize, 8);
	}

	/* now the bss */
	bsssize = 0;
	for(i=0; i&lt;NHASH; i++)
	for(s = hash[i]; s != S; s = s-&gt;link) {
		if(!s-&gt;reachable)
			continue;
		if(s-&gt;type != SBSS)
			continue;
		t = s-&gt;value;
		s-&gt;size = t;
		if(t &gt;= 8)
			bsssize = rnd(bsssize, 8);
		s-&gt;value = bsssize + dynptrsize + datsize;
		bsssize += t;
	}

	xdefine(&#34;data&#34;, SBSS, 0);
	xdefine(&#34;edata&#34;, SBSS, datsize);
	xdefine(&#34;end&#34;, SBSS, dynptrsize + bsssize + datsize);
}

Prog*
brchain(Prog *p)
{
	int i;

	for(i=0; i&lt;20; i++) {
		if(p == P || p-&gt;as != AJMP)
			return p;
		p = p-&gt;pcond;
	}
	return P;
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
	lastp-&gt;link = P;
	firstp = firstp-&gt;link;
}

void
xfol(Prog *p)
{
	Prog *q;
	int i;
	enum as a;

loop:
	if(p == P)
		return;
	if(p-&gt;as == ATEXT)
		curtext = p;
	if(!curtext-&gt;from.sym-&gt;reachable) {
		p = p-&gt;pcond;
		goto loop;
	}
	if(p-&gt;as == AJMP)
	if((q = p-&gt;pcond) != P &amp;&amp; q-&gt;as != ATEXT) {
		p-&gt;mark = 1;
		p = q;
		if(p-&gt;mark == 0)
			goto loop;
	}
	if(p-&gt;mark) {
		/* copy up to 4 instructions to avoid branch */
		for(i=0,q=p; i&lt;4; i++,q=q-&gt;link) {
			if(q == P)
				break;
			if(q == lastp)
				break;
			a = q-&gt;as;
			if(a == ANOP) {
				i--;
				continue;
			}
			switch(a) {
			case AJMP:
			case ARET:
			case AIRETL:
			case AIRETQ:
			case AIRETW:
			case ARETFL:
			case ARETFQ:
			case ARETFW:

			case APUSHL:
			case APUSHFL:
			case APUSHQ:
			case APUSHFQ:
			case APUSHW:
			case APUSHFW:
			case APOPL:
			case APOPFL:
			case APOPQ:
			case APOPFQ:
			case APOPW:
			case APOPFW:
				goto brk;
			}
			if(q-&gt;pcond == P || q-&gt;pcond-&gt;mark)
				continue;
			if(a == ACALL || a == ALOOP)
				continue;
			for(;;) {
				if(p-&gt;as == ANOP) {
					p = p-&gt;link;
					continue;
				}
				q = copyp(p);
				p = p-&gt;link;
				q-&gt;mark = 1;
				lastp-&gt;link = q;
				lastp = q;
				if(q-&gt;as != a || q-&gt;pcond == P || q-&gt;pcond-&gt;mark)
					continue;

				q-&gt;as = relinv(q-&gt;as);
				p = q-&gt;pcond;
				q-&gt;pcond = q-&gt;link;
				q-&gt;link = p;
				xfol(q-&gt;link);
				p = q-&gt;link;
				if(p-&gt;mark)
					return;
				goto loop;
			}
		} /* */
	brk:;
		q = prg();
		q-&gt;as = AJMP;
		q-&gt;line = p-&gt;line;
		q-&gt;to.type = D_BRANCH;
		q-&gt;to.offset = p-&gt;pc;
		q-&gt;pcond = p;
		p = q;
	}
	p-&gt;mark = 1;
	lastp-&gt;link = p;
	lastp = p;
	a = p-&gt;as;
	if(a == AJMP || a == ARET || a == AIRETL || a == AIRETQ || a == AIRETW ||
	   a == ARETFL || a == ARETFQ || a == ARETFW)
		return;
	if(p-&gt;pcond != P)
	if(a != ACALL) {
		q = brchain(p-&gt;link);
		if(q != P &amp;&amp; q-&gt;mark)
		if(a != ALOOP) {
			p-&gt;as = relinv(a);
			p-&gt;link = p-&gt;pcond;
			p-&gt;pcond = q;
		}
		xfol(p-&gt;link);
		q = brchain(p-&gt;pcond);
		if(q-&gt;mark) {
			p-&gt;pcond = q;
			return;
		}
		p = q;
		goto loop;
	}
	p = p-&gt;link;
	goto loop;
}

Prog*
byteq(int v)
{
	Prog *p;

	p = prg();
	p-&gt;as = ABYTE;
	p-&gt;from.type = D_CONST;
	p-&gt;from.offset = v&amp;0xff;
	return p;
}

int
relinv(int a)
{

	switch(a) {
	case AJEQ:	return AJNE;
	case AJNE:	return AJEQ;
	case AJLE:	return AJGT;
	case AJLS:	return AJHI;
	case AJLT:	return AJGE;
	case AJMI:	return AJPL;
	case AJGE:	return AJLT;
	case AJPL:	return AJMI;
	case AJGT:	return AJLE;
	case AJHI:	return AJLS;
	case AJCS:	return AJCC;
	case AJCC:	return AJCS;
	case AJPS:	return AJPC;
	case AJPC:	return AJPS;
	case AJOS:	return AJOC;
	case AJOC:	return AJOS;
	}
	diag(&#34;unknown relation: %s in %s&#34;, anames[a], TNAME);
	return a;
}

void
doinit(void)
{
	Sym *s;
	Prog *p;
	int x;

	for(p = datap; p != P; p = p-&gt;link) {
		x = p-&gt;to.type;
		if(x != D_EXTERN &amp;&amp; x != D_STATIC)
			continue;
		s = p-&gt;to.sym;
		if(s-&gt;type == 0 || s-&gt;type == SXREF)
			diag(&#34;undefined %s initializer of %s&#34;,
				s-&gt;name, p-&gt;from.sym-&gt;name);
		p-&gt;to.offset += s-&gt;value;
		p-&gt;to.type = D_CONST;
		if(s-&gt;type == SDATA || s-&gt;type == SBSS)
			p-&gt;to.offset += INITDAT;
	}
}

void
patch(void)
{
	int32 c;
	Prog *p, *q;
	Sym *s;
	int32 vexit;

	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f mkfwd\n&#34;, cputime());
	Bflush(&amp;bso);
	mkfwd();
	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f patch\n&#34;, cputime());
	Bflush(&amp;bso);

	s = lookup(&#34;exit&#34;, 0);
	vexit = s-&gt;value;
	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT)
			curtext = p;
		if(p-&gt;as == ACALL || (p-&gt;as == AJMP &amp;&amp; p-&gt;to.type != D_BRANCH)) {
			s = p-&gt;to.sym;
			if(s) {
				if(debug[&#39;c&#39;])
					Bprint(&amp;bso, &#34;%s calls %s\n&#34;, TNAME, s-&gt;name);
				switch(s-&gt;type) {
				default:
					/* diag prints TNAME first */
					diag(&#34;undefined: %s&#34;, s-&gt;name);
					s-&gt;type = STEXT;
					s-&gt;value = vexit;
					continue;	// avoid more error messages
				case STEXT:
					p-&gt;to.offset = s-&gt;value;
					break;
				case SUNDEF:
					p-&gt;pcond = UP;
					p-&gt;to.offset = 0;
					break;
				}
				p-&gt;to.type = D_BRANCH;
			}
		}
		if(p-&gt;to.type != D_BRANCH || p-&gt;pcond == UP)
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
			diag(&#34;branch out of range in %s\n%P [%s]&#34;,
				TNAME, p, p-&gt;to.sym ? p-&gt;to.sym-&gt;name : &#34;&lt;nil&gt;&#34;);
			p-&gt;to.type = D_NONE;
		}
		p-&gt;pcond = q;
	}

	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT)
			curtext = p;
		p-&gt;mark = 0;	/* initialization for follow */
		if(p-&gt;pcond != P &amp;&amp; p-&gt;pcond != UP) {
			p-&gt;pcond = brloop(p-&gt;pcond);
			if(p-&gt;pcond != P)
			if(p-&gt;to.type == D_BRANCH)
				p-&gt;to.offset = p-&gt;pcond-&gt;pc;
		}
	}
}

#define	LOG	5
void
mkfwd(void)
{
	Prog *p;
	int i;
	int32 dwn[LOG], cnt[LOG];
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
	int c;
	Prog *q;

	c = 0;
	for(q = p; q != P; q = q-&gt;pcond) {
		if(q-&gt;as != AJMP)
			break;
		c++;
		if(c &gt;= 5000)
			return P;
	}
	return q;
}

static char*
morename[] =
{
	&#34;runtime·morestack00&#34;,
	&#34;runtime·morestack10&#34;,
	&#34;runtime·morestack01&#34;,
	&#34;runtime·morestack11&#34;,

	&#34;runtime·morestack8&#34;,
	&#34;runtime·morestack16&#34;,
	&#34;runtime·morestack24&#34;,
	&#34;runtime·morestack32&#34;,
	&#34;runtime·morestack40&#34;,
	&#34;runtime·morestack48&#34;,
};
Prog*	pmorestack[nelem(morename)];
Sym*	symmorestack[nelem(morename)];

void
dostkoff(void)
{
	Prog *p, *q, *q1;
	int32 autoffset, deltasp;
	int a, f, curframe, curbecome, maxbecome, pcsize;
	uint32 moreconst1, moreconst2, i;

	for(i=0; i&lt;nelem(morename); i++) {
		symmorestack[i] = lookup(morename[i], 0);
		pmorestack[i] = P;
	}

	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT) {
			for(i=0; i&lt;nelem(morename); i++) {
				if(p-&gt;from.sym == symmorestack[i]) {
					pmorestack[i] = p;
					break;
				}
			}
		}
	}

	for(i=0; i&lt;nelem(morename); i++) {
		if(pmorestack[i] == P)
			diag(&#34;morestack trampoline not defined&#34;);
	}

	curframe = 0;
	curbecome = 0;
	maxbecome = 0;
	curtext = 0;
	for(p = firstp; p != P; p = p-&gt;link) {

		/* find out how much arg space is used in this TEXT */
		if(p-&gt;to.type == (D_INDIR+D_SP))
			if(p-&gt;to.offset &gt; curframe)
				curframe = p-&gt;to.offset;

		switch(p-&gt;as) {
		case ATEXT:
			if(curtext &amp;&amp; curtext-&gt;from.sym) {
				curtext-&gt;from.sym-&gt;frame = curframe;
				curtext-&gt;from.sym-&gt;become = curbecome;
				if(curbecome &gt; maxbecome)
					maxbecome = curbecome;
			}
			curframe = 0;
			curbecome = 0;

			curtext = p;
			break;

		case ARET:
			/* special form of RET is BECOME */
			if(p-&gt;from.type == D_CONST)
				if(p-&gt;from.offset &gt; curbecome)
					curbecome = p-&gt;from.offset;
			break;
		}
	}
	if(curtext &amp;&amp; curtext-&gt;from.sym) {
		curtext-&gt;from.sym-&gt;frame = curframe;
		curtext-&gt;from.sym-&gt;become = curbecome;
		if(curbecome &gt; maxbecome)
			maxbecome = curbecome;
	}

	if(debug[&#39;b&#39;])
		print(&#34;max become = %d\n&#34;, maxbecome);
	xdefine(&#34;ALEFbecome&#34;, STEXT, maxbecome);

	curtext = 0;
	for(p = firstp; p != P; p = p-&gt;link) {
		switch(p-&gt;as) {
		case ATEXT:
			curtext = p;
			break;
		case ACALL:
			if(curtext != P &amp;&amp; curtext-&gt;from.sym != S &amp;&amp; curtext-&gt;to.offset &gt;= 0) {
				f = maxbecome - curtext-&gt;from.sym-&gt;frame;
				if(f &lt;= 0)
					break;
				/* calling a become or calling a variable */
				if(p-&gt;to.sym == S || p-&gt;to.sym-&gt;become) {
					curtext-&gt;to.offset += f;
					if(debug[&#39;b&#39;]) {
						curp = p;
						print(&#34;%D calling %D increase %d\n&#34;,
							&amp;curtext-&gt;from, &amp;p-&gt;to, f);
					}
				}
			}
			break;
		}
	}

	autoffset = 0;
	deltasp = 0;
	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT) {
			curtext = p;
			parsetextconst(p-&gt;to.offset);
			autoffset = textstksiz;
			if(autoffset &lt; 0)
				autoffset = 0;

			q = P;
			q1 = P;
			if((p-&gt;from.scale &amp; NOSPLIT) &amp;&amp; autoffset &gt;= StackSmall)
				diag(&#34;nosplit func likely to overflow stack&#34;);

			if(!(p-&gt;from.scale &amp; NOSPLIT)) {
				if(debug[&#39;K&#39;]) {
					// 6l -K means check not only for stack
					// overflow but stack underflow.
					// On underflow, INT 3 (breakpoint).
					// Underflow itself is rare but this also
					// catches out-of-sync stack guard info

					p = appendp(p);
					p-&gt;as = ACMPQ;
					p-&gt;from.type = D_INDIR+D_R15;
					p-&gt;from.offset = 8;
					p-&gt;to.type = D_SP;

					p = appendp(p);
					p-&gt;as = AJHI;
					p-&gt;to.type = D_BRANCH;
					p-&gt;to.offset = 4;
					q1 = p;

					p = appendp(p);
					p-&gt;as = AINT;
					p-&gt;from.type = D_CONST;
					p-&gt;from.offset = 3;
				}

				if(autoffset &lt; StackBig) {  // do we need to call morestack?
					if(autoffset &lt;= StackSmall) {
						// small stack
						p = appendp(p);
						p-&gt;as = ACMPQ;
						p-&gt;from.type = D_SP;
						p-&gt;to.type = D_INDIR+D_R15;
						if(q1) {
							q1-&gt;pcond = p;
							q1 = P;
						}
					} else {
						// large stack
						p = appendp(p);
						p-&gt;as = ALEAQ;
						p-&gt;from.type = D_INDIR+D_SP;
						p-&gt;from.offset = -(autoffset-StackSmall);
						p-&gt;to.type = D_AX;
						if(q1) {
							q1-&gt;pcond = p;
							q1 = P;
						}

						p = appendp(p);
						p-&gt;as = ACMPQ;
						p-&gt;from.type = D_AX;
						p-&gt;to.type = D_INDIR+D_R15;
					}

					// common
					p = appendp(p);
					p-&gt;as = AJHI;
					p-&gt;to.type = D_BRANCH;
					p-&gt;to.offset = 4;
					q = p;
				}

				/* 160 comes from 3 calls (3*8) 4 safes (4*8) and 104 guard */
				moreconst1 = 0;
				if(autoffset+160 &gt; 4096)
					moreconst1 = (autoffset+160) &amp; ~7LL;
				moreconst2 = textarg;

				// 4 varieties varieties (const1==0 cross const2==0)
				// and 6 subvarieties of (const1==0 and const2!=0)
				p = appendp(p);
				if(moreconst1 == 0 &amp;&amp; moreconst2 == 0) {
					p-&gt;as = ACALL;
					p-&gt;to.type = D_BRANCH;
					p-&gt;pcond = pmorestack[0];
					p-&gt;to.sym = symmorestack[0];
					if(q1) {
						q1-&gt;pcond = p;
						q1 = P;
					}
				} else
				if(moreconst1 != 0 &amp;&amp; moreconst2 == 0) {
					p-&gt;as = AMOVL;
					p-&gt;from.type = D_CONST;
					p-&gt;from.offset = moreconst1;
					p-&gt;to.type = D_AX;
					if(q1) {
						q1-&gt;pcond = p;
						q1 = P;
					}

					p = appendp(p);
					p-&gt;as = ACALL;
					p-&gt;to.type = D_BRANCH;
					p-&gt;pcond = pmorestack[1];
					p-&gt;to.sym = symmorestack[1];
				} else
				if(moreconst1 == 0 &amp;&amp; moreconst2 &lt;= 48 &amp;&amp; moreconst2%8 == 0) {
					i = moreconst2/8 + 3;
					p-&gt;as = ACALL;
					p-&gt;to.type = D_BRANCH;
					p-&gt;pcond = pmorestack[i];
					p-&gt;to.sym = symmorestack[i];
					if(q1) {
						q1-&gt;pcond = p;
						q1 = P;
					}
				} else
				if(moreconst1 == 0 &amp;&amp; moreconst2 != 0) {
					p-&gt;as = AMOVL;
					p-&gt;from.type = D_CONST;
					p-&gt;from.offset = moreconst2;
					p-&gt;to.type = D_AX;
					if(q1) {
						q1-&gt;pcond = p;
						q1 = P;
					}

					p = appendp(p);
					p-&gt;as = ACALL;
					p-&gt;to.type = D_BRANCH;
					p-&gt;pcond = pmorestack[2];
					p-&gt;to.sym = symmorestack[2];
				} else {
					p-&gt;as = AMOVQ;
					p-&gt;from.type = D_CONST;
					p-&gt;from.offset = (uint64)moreconst2 &lt;&lt; 32;
					p-&gt;from.offset |= moreconst1;
					p-&gt;to.type = D_AX;
					if(q1) {
						q1-&gt;pcond = p;
						q1 = P;
					}

					p = appendp(p);
					p-&gt;as = ACALL;
					p-&gt;to.type = D_BRANCH;
					p-&gt;pcond = pmorestack[3];
					p-&gt;to.sym = symmorestack[3];
				}
			}

			if(q != P)
				q-&gt;pcond = p-&gt;link;

			if(autoffset) {
				p = appendp(p);
				p-&gt;as = AADJSP;
				p-&gt;from.type = D_CONST;
				p-&gt;from.offset = autoffset;
				if(q != P)
					q-&gt;pcond = p;
			}
			deltasp = autoffset;

			if(debug[&#39;K&#39;] &gt; 1 &amp;&amp; autoffset) {
				// 6l -KK means double-check for stack overflow
				// even after calling morestack and even if the
				// function is marked as nosplit.
				p = appendp(p);
				p-&gt;as = AMOVQ;
				p-&gt;from.type = D_INDIR+D_R15;
				p-&gt;from.offset = 0;
				p-&gt;to.type = D_BX;

				p = appendp(p);
				p-&gt;as = ASUBQ;
				p-&gt;from.type = D_CONST;
				p-&gt;from.offset = StackSmall+32;
				p-&gt;to.type = D_BX;

				p = appendp(p);
				p-&gt;as = ACMPQ;
				p-&gt;from.type = D_SP;
				p-&gt;to.type = D_BX;

				p = appendp(p);
				p-&gt;as = AJHI;
				p-&gt;to.type = D_BRANCH;
				q1 = p;

				p = appendp(p);
				p-&gt;as = AINT;
				p-&gt;from.type = D_CONST;
				p-&gt;from.offset = 3;

				p = appendp(p);
				p-&gt;as = ANOP;
				q1-&gt;pcond = p;
				q1 = P;
			}
		}
		pcsize = p-&gt;mode/8;
		a = p-&gt;from.type;
		if(a == D_AUTO)
			p-&gt;from.offset += deltasp;
		if(a == D_PARAM)
			p-&gt;from.offset += deltasp + pcsize;
		a = p-&gt;to.type;
		if(a == D_AUTO)
			p-&gt;to.offset += deltasp;
		if(a == D_PARAM)
			p-&gt;to.offset += deltasp + pcsize;

		switch(p-&gt;as) {
		default:
			continue;
		case APUSHL:
		case APUSHFL:
			deltasp += 4;
			continue;
		case APUSHQ:
		case APUSHFQ:
			deltasp += 8;
			continue;
		case APUSHW:
		case APUSHFW:
			deltasp += 2;
			continue;
		case APOPL:
		case APOPFL:
			deltasp -= 4;
			continue;
		case APOPQ:
		case APOPFQ:
			deltasp -= 8;
			continue;
		case APOPW:
		case APOPFW:
			deltasp -= 2;
			continue;
		case ARET:
			break;
		}

		if(autoffset != deltasp)
			diag(&#34;unbalanced PUSH/POP&#34;);
		if(p-&gt;from.type == D_CONST)
			goto become;

		if(autoffset) {
			p-&gt;as = AADJSP;
			p-&gt;from.type = D_CONST;
			p-&gt;from.offset = -autoffset;

			p = appendp(p);
			p-&gt;as = ARET;
		}
		continue;

	become:
		q = p;
		p = appendp(p);
		p-&gt;as = AJMP;
		p-&gt;to = q-&gt;to;
		p-&gt;pcond = q-&gt;pcond;

		q-&gt;as = AADJSP;
		q-&gt;from = zprg.from;
		q-&gt;from.type = D_CONST;
		q-&gt;from.offset = -autoffset;
		q-&gt;to = zprg.to;
		continue;
	}
}

vlong
atolwhex(char *s)
{
	vlong n;
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

void
import(void)
{
	int i;
	Sym *s;

	for(i = 0; i &lt; NHASH; i++)
		for(s = hash[i]; s != S; s = s-&gt;link)
			if(s-&gt;sig != 0 &amp;&amp; s-&gt;type == SXREF &amp;&amp; (nimports == 0 || s-&gt;subtype == SIMPORT)){
				if(s-&gt;value != 0)
					diag(&#34;value != 0 on SXREF&#34;);
				undefsym(s);
				Bprint(&amp;bso, &#34;IMPORT: %s sig=%lux v=%lld\n&#34;, s-&gt;name, s-&gt;sig, s-&gt;value);
				if(debug[&#39;S&#39;])
					s-&gt;sig = 0;
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
	if(edatap == P)
		datap = p;
	else
		edatap-&gt;link = p;
	edatap = p;
	p-&gt;as = ADATA;
	p-&gt;width = w;
	p-&gt;from.scale = w;
	p-&gt;from.type = t;
	p-&gt;from.sym = s;
	p-&gt;from.offset = o;
	p-&gt;to.type = D_CONST;
	p-&gt;dlink = s-&gt;data;
	s-&gt;data = p;
	return p;
}

Prog*
newtext(Prog *p, Sym *s)
{
	if(p == P) {
		p = prg();
		p-&gt;as = ATEXT;
		p-&gt;from.sym = s;
	}
	s-&gt;type = STEXT;
	s-&gt;text = p;
	s-&gt;value = pc;
	lastp-&gt;link = p;
	lastp = p;
	p-&gt;pc = pc++;
	if(textp == P)
		textp = p;
	else
		etextp-&gt;pcond = p;
	etextp = p;
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
			if(s-&gt;sig != 0 &amp;&amp; s-&gt;type != SXREF &amp;&amp;
			   s-&gt;type != SUNDEF &amp;&amp;
			   (nexports == 0 || s-&gt;subtype == SEXPORT))
				n++;
	esyms = mal(n*sizeof(Sym*));
	ne = n;
	n = 0;
	for(i = 0; i &lt; NHASH; i++)
		for(s = hash[i]; s != S; s = s-&gt;link)
			if(s-&gt;sig != 0 &amp;&amp; s-&gt;type != SXREF &amp;&amp;
			   s-&gt;type != SUNDEF &amp;&amp;
			   (nexports == 0 || s-&gt;subtype == SEXPORT))
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
		if(debug[&#39;S&#39;])
			s-&gt;sig = 0;
		/* Bprint(&amp;bso, &#34;EXPORT: %s sig=%lux t=%d\n&#34;, s-&gt;name, s-&gt;sig, s-&gt;type); */

		/* signature */
		p = newdata(et, off, sizeof(int32), D_EXTERN);
		off += sizeof(int32);
		p-&gt;to.offset = s-&gt;sig;

		/* address */
		p = newdata(et, off, sizeof(int32), D_EXTERN);
		off += sizeof(int32);
		p-&gt;to.type = D_ADDR;
		p-&gt;to.index = D_EXTERN;
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
				memmove(p-&gt;to.scon, buf, NSNAME);
				nb = 0;
			}
			if(*t++ == 0)
				break;
		}

		/* name */
		p = newdata(et, off, sizeof(int32), D_EXTERN);
		off += sizeof(int32);
		p-&gt;to.type = D_ADDR;
		p-&gt;to.index = D_STATIC;
		p-&gt;to.sym = str;
		p-&gt;to.offset = sv-n;
	}

	if(nb &gt; 0){
		p = newdata(str, sv-nb, nb, D_STATIC);
		p-&gt;to.type = D_SCONST;
		memmove(p-&gt;to.scon, buf, nb);
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
