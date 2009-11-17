<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5c/peep.c</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5c/peep.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5c/peep.c
// http://code.google.com/p/inferno-os/source/browse/utils/5c/peep.c
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


#include &#34;gc.h&#34;

int xtramodes(Reg*, Adr*);

void
peep(void)
{
	Reg *r, *r1, *r2;
	Prog *p, *p1;
	int t;
/*
 * complete R structure
 */
	t = 0;
	for(r=firstr; r!=R; r=r1) {
		r1 = r-&gt;link;
		if(r1 == R)
			break;
		p = r-&gt;prog-&gt;link;
		while(p != r1-&gt;prog)
		switch(p-&gt;as) {
		default:
			r2 = rega();
			r-&gt;link = r2;
			r2-&gt;link = r1;

			r2-&gt;prog = p;
			r2-&gt;p1 = r;
			r-&gt;s1 = r2;
			r2-&gt;s1 = r1;
			r1-&gt;p1 = r2;

			r = r2;
			t++;

		case ADATA:
		case AGLOBL:
		case ANAME:
		case ASIGNAME:
			p = p-&gt;link;
		}
	}

loop1:
	t = 0;
	for(r=firstr; r!=R; r=r-&gt;link) {
		p = r-&gt;prog;
		if(p-&gt;as == ASLL || p-&gt;as == ASRL || p-&gt;as == ASRA) {
			/*
			 * elide shift into D_SHIFT operand of subsequent instruction
			 */
			if(shiftprop(r)) {
				excise(r);
				t++;
			}
		}
		if(p-&gt;as == AMOVW || p-&gt;as == AMOVF || p-&gt;as == AMOVD)
		if(regtyp(&amp;p-&gt;to)) {
			if(p-&gt;from.type == D_CONST)
				constprop(&amp;p-&gt;from, &amp;p-&gt;to, r-&gt;s1);
			else if(regtyp(&amp;p-&gt;from))
			if(p-&gt;from.type == p-&gt;to.type) {
				if(copyprop(r)) {
					excise(r);
					t++;
				} else
				if(subprop(r) &amp;&amp; copyprop(r)) {
					excise(r);
					t++;
				}
			}
		}
	}
	if(t)
		goto loop1;
	/*
	 * look for MOVB x,R; MOVB R,R
	 */
	for(r=firstr; r!=R; r=r-&gt;link) {
		p = r-&gt;prog;
		switch(p-&gt;as) {
		default:
			continue;
		case AEOR:
			/*
			 * EOR -1,x,y =&gt; MVN x,y
			 */
			if(p-&gt;from.type == D_CONST &amp;&amp; p-&gt;from.offset == -1) {
				p-&gt;as = AMVN;
				p-&gt;from.type = D_REG;
				if(p-&gt;reg != NREG)
					p-&gt;from.reg = p-&gt;reg;
				else
					p-&gt;from.reg = p-&gt;to.reg;
				p-&gt;reg = NREG;
			}
			continue;
		case AMOVH:
		case AMOVHU:
		case AMOVB:
		case AMOVBU:
			if(p-&gt;to.type != D_REG)
				continue;
			break;
		}
		r1 = r-&gt;link;
		if(r1 == R)
			continue;
		p1 = r1-&gt;prog;
		if(p1-&gt;as != p-&gt;as)
			continue;
		if(p1-&gt;from.type != D_REG || p1-&gt;from.reg != p-&gt;to.reg)
			continue;
		if(p1-&gt;to.type != D_REG || p1-&gt;to.reg != p-&gt;to.reg)
			continue;
		excise(r1);
	}

	for(r=firstr; r!=R; r=r-&gt;link) {
		p = r-&gt;prog;
		switch(p-&gt;as) {
		case AMOVW:
		case AMOVB:
		case AMOVBU:
			if(p-&gt;from.type == D_OREG &amp;&amp; p-&gt;from.offset == 0)
				xtramodes(r, &amp;p-&gt;from);
			else if(p-&gt;to.type == D_OREG &amp;&amp; p-&gt;to.offset == 0)
				xtramodes(r, &amp;p-&gt;to);
			else
				continue;
			break;
		case ACMP:
			/*
			 * elide CMP $0,x if calculation of x can set condition codes
			 */
			if(p-&gt;from.type != D_CONST || p-&gt;from.offset != 0)
				continue;
			r2 = r-&gt;s1;
			if(r2 == R)
				continue;
			t = r2-&gt;prog-&gt;as;
			switch(t) {
			default:
				continue;
			case ABEQ:
			case ABNE:
			case ABMI:
			case ABPL:
				break;
			case ABGE:
				t = ABPL;
				break;
			case ABLT:
				t = ABMI;
				break;
			case ABHI:
				t = ABNE;
				break;
			case ABLS:
				t = ABEQ;
				break;
			}
			r1 = r;
			do
				r1 = uniqp(r1);
			while (r1 != R &amp;&amp; r1-&gt;prog-&gt;as == ANOP);
			if(r1 == R)
				continue;
			p1 = r1-&gt;prog;
			if(p1-&gt;to.type != D_REG)
				continue;
			if(p1-&gt;to.reg != p-&gt;reg)
			if(!(p1-&gt;as == AMOVW &amp;&amp; p1-&gt;from.type == D_REG &amp;&amp; p1-&gt;from.reg == p-&gt;reg))
				continue;
			switch(p1-&gt;as) {
			default:
				continue;
			case AMOVW:
				if(p1-&gt;from.type != D_REG)
					continue;
			case AAND:
			case AEOR:
			case AORR:
			case ABIC:
			case AMVN:
			case ASUB:
			case ARSB:
			case AADD:
			case AADC:
			case ASBC:
			case ARSC:
				break;
			}
			p1-&gt;scond |= C_SBIT;
			r2-&gt;prog-&gt;as = t;
			excise(r);
			continue;
		}
	}

	predicate();
}

void
excise(Reg *r)
{
	Prog *p;

	p = r-&gt;prog;
	p-&gt;as = ANOP;
	p-&gt;scond = zprog.scond;
	p-&gt;from = zprog.from;
	p-&gt;to = zprog.to;
	p-&gt;reg = zprog.reg; /**/
}

Reg*
uniqp(Reg *r)
{
	Reg *r1;

	r1 = r-&gt;p1;
	if(r1 == R) {
		r1 = r-&gt;p2;
		if(r1 == R || r1-&gt;p2link != R)
			return R;
	} else
		if(r-&gt;p2 != R)
			return R;
	return r1;
}

Reg*
uniqs(Reg *r)
{
	Reg *r1;

	r1 = r-&gt;s1;
	if(r1 == R) {
		r1 = r-&gt;s2;
		if(r1 == R)
			return R;
	} else
		if(r-&gt;s2 != R)
			return R;
	return r1;
}

int
regtyp(Adr *a)
{

	if(a-&gt;type == D_REG)
		return 1;
	if(a-&gt;type == D_FREG)
		return 1;
	return 0;
}

/*
 * the idea is to substitute
 * one register for another
 * from one MOV to another
 *	MOV	a, R0
 *	ADD	b, R0	/ no use of R1
 *	MOV	R0, R1
 * would be converted to
 *	MOV	a, R1
 *	ADD	b, R1
 *	MOV	R1, R0
 * hopefully, then the former or latter MOV
 * will be eliminated by copy propagation.
 */
int
subprop(Reg *r0)
{
	Prog *p;
	Adr *v1, *v2;
	Reg *r;
	int t;

	p = r0-&gt;prog;
	v1 = &amp;p-&gt;from;
	if(!regtyp(v1))
		return 0;
	v2 = &amp;p-&gt;to;
	if(!regtyp(v2))
		return 0;
	for(r=uniqp(r0); r!=R; r=uniqp(r)) {
		if(uniqs(r) == R)
			break;
		p = r-&gt;prog;
		switch(p-&gt;as) {
		case ABL:
			return 0;

		case ACMP:
		case ACMN:
		case AADD:
		case ASUB:
		case ARSB:
		case ASLL:
		case ASRL:
		case ASRA:
		case AORR:
		case AAND:
		case AEOR:
		case AMUL:
		case ADIV:
		case ADIVU:

		case ACMPF:
		case ACMPD:
		case AADDD:
		case AADDF:
		case ASUBD:
		case ASUBF:
		case AMULD:
		case AMULF:
		case ADIVD:
		case ADIVF:
			if(p-&gt;to.type == v1-&gt;type)
			if(p-&gt;to.reg == v1-&gt;reg) {
				if(p-&gt;reg == NREG)
					p-&gt;reg = p-&gt;to.reg;
				goto gotit;
			}
			break;

		case AMOVF:
		case AMOVD:
		case AMOVW:
			if(p-&gt;to.type == v1-&gt;type)
			if(p-&gt;to.reg == v1-&gt;reg)
				goto gotit;
			break;

		case AMOVM:
			t = 1&lt;&lt;v2-&gt;reg;
			if((p-&gt;from.type == D_CONST &amp;&amp; (p-&gt;from.offset&amp;t)) ||
			   (p-&gt;to.type == D_CONST &amp;&amp; (p-&gt;to.offset&amp;t)))
				return 0;
			break;
		}
		if(copyau(&amp;p-&gt;from, v2) ||
		   copyau1(p, v2) ||
		   copyau(&amp;p-&gt;to, v2))
			break;
		if(copysub(&amp;p-&gt;from, v1, v2, 0) ||
		   copysub1(p, v1, v2, 0) ||
		   copysub(&amp;p-&gt;to, v1, v2, 0))
			break;
	}
	return 0;

gotit:
	copysub(&amp;p-&gt;to, v1, v2, 1);
	if(debug[&#39;P&#39;]) {
		print(&#34;gotit: %D-&gt;%D\n%P&#34;, v1, v2, r-&gt;prog);
		if(p-&gt;from.type == v2-&gt;type)
			print(&#34; excise&#34;);
		print(&#34;\n&#34;);
	}
	for(r=uniqs(r); r!=r0; r=uniqs(r)) {
		p = r-&gt;prog;
		copysub(&amp;p-&gt;from, v1, v2, 1);
		copysub1(p, v1, v2, 1);
		copysub(&amp;p-&gt;to, v1, v2, 1);
		if(debug[&#39;P&#39;])
			print(&#34;%P\n&#34;, r-&gt;prog);
	}
	t = v1-&gt;reg;
	v1-&gt;reg = v2-&gt;reg;
	v2-&gt;reg = t;
	if(debug[&#39;P&#39;])
		print(&#34;%P last\n&#34;, r-&gt;prog);
	return 1;
}

/*
 * The idea is to remove redundant copies.
 *	v1-&gt;v2	F=0
 *	(use v2	s/v2/v1/)*
 *	set v1	F=1
 *	use v2	return fail
 *	-----------------
 *	v1-&gt;v2	F=0
 *	(use v2	s/v2/v1/)*
 *	set v1	F=1
 *	set v2	return success
 */
int
copyprop(Reg *r0)
{
	Prog *p;
	Adr *v1, *v2;
	Reg *r;

	p = r0-&gt;prog;
	v1 = &amp;p-&gt;from;
	v2 = &amp;p-&gt;to;
	if(copyas(v1, v2))
		return 1;
	for(r=firstr; r!=R; r=r-&gt;link)
		r-&gt;active = 0;
	return copy1(v1, v2, r0-&gt;s1, 0);
}

int
copy1(Adr *v1, Adr *v2, Reg *r, int f)
{
	int t;
	Prog *p;

	if(r-&gt;active) {
		if(debug[&#39;P&#39;])
			print(&#34;act set; return 1\n&#34;);
		return 1;
	}
	r-&gt;active = 1;
	if(debug[&#39;P&#39;])
		print(&#34;copy %D-&gt;%D f=%d\n&#34;, v1, v2, f);
	for(; r != R; r = r-&gt;s1) {
		p = r-&gt;prog;
		if(debug[&#39;P&#39;])
			print(&#34;%P&#34;, p);
		if(!f &amp;&amp; uniqp(r) == R) {
			f = 1;
			if(debug[&#39;P&#39;])
				print(&#34;; merge; f=%d&#34;, f);
		}
		t = copyu(p, v2, A);
		switch(t) {
		case 2:	/* rar, cant split */
			if(debug[&#39;P&#39;])
				print(&#34;; %Drar; return 0\n&#34;, v2);
			return 0;

		case 3:	/* set */
			if(debug[&#39;P&#39;])
				print(&#34;; %Dset; return 1\n&#34;, v2);
			return 1;

		case 1:	/* used, substitute */
		case 4:	/* use and set */
			if(f) {
				if(!debug[&#39;P&#39;])
					return 0;
				if(t == 4)
					print(&#34;; %Dused+set and f=%d; return 0\n&#34;, v2, f);
				else
					print(&#34;; %Dused and f=%d; return 0\n&#34;, v2, f);
				return 0;
			}
			if(copyu(p, v2, v1)) {
				if(debug[&#39;P&#39;])
					print(&#34;; sub fail; return 0\n&#34;);
				return 0;
			}
			if(debug[&#39;P&#39;])
				print(&#34;; sub%D/%D&#34;, v2, v1);
			if(t == 4) {
				if(debug[&#39;P&#39;])
					print(&#34;; %Dused+set; return 1\n&#34;, v2);
				return 1;
			}
			break;
		}
		if(!f) {
			t = copyu(p, v1, A);
			if(!f &amp;&amp; (t == 2 || t == 3 || t == 4)) {
				f = 1;
				if(debug[&#39;P&#39;])
					print(&#34;; %Dset and !f; f=%d&#34;, v1, f);
			}
		}
		if(debug[&#39;P&#39;])
			print(&#34;\n&#34;);
		if(r-&gt;s2)
			if(!copy1(v1, v2, r-&gt;s2, f))
				return 0;
	}
	return 1;
}

/*
 * The idea is to remove redundant constants.
 *	$c1-&gt;v1
 *	($c1-&gt;v2 s/$c1/v1)*
 *	set v1  return
 * The v1-&gt;v2 should be eliminated by copy propagation.
 */
void
constprop(Adr *c1, Adr *v1, Reg *r)
{
	Prog *p;

	if(debug[&#39;C&#39;])
		print(&#34;constprop %D-&gt;%D\n&#34;, c1, v1);
	for(; r != R; r = r-&gt;s1) {
		p = r-&gt;prog;
		if(debug[&#39;C&#39;])
			print(&#34;%P&#34;, p);
		if(uniqp(r) == R) {
			if(debug[&#39;C&#39;])
				print(&#34;; merge; return\n&#34;);
			return;
		}
		if(p-&gt;as == AMOVW &amp;&amp; copyas(&amp;p-&gt;from, c1)) {
				if(debug[&#39;C&#39;])
					print(&#34;; sub%D/%D&#34;, &amp;p-&gt;from, v1);
				p-&gt;from = *v1;
		} else if(copyu(p, v1, A) &gt; 1) {
			if(debug[&#39;C&#39;])
				print(&#34;; %Dset; return\n&#34;, v1);
			return;
		}
		if(debug[&#39;C&#39;])
			print(&#34;\n&#34;);
		if(r-&gt;s2)
			constprop(c1, v1, r-&gt;s2);
	}
}

/*
 * ASLL x,y,w
 * .. (not use w, not set x y w)
 * AXXX w,a,b (a != w)
 * .. (not use w)
 * (set w)
 * ----------- changed to
 * ..
 * AXXX (x&lt;&lt;y),a,b
 * ..
 */
#define FAIL(msg) { if(debug[&#39;H&#39;]) print(&#34;\t%s; FAILURE\n&#34;, msg); return 0; }
int
shiftprop(Reg *r)
{
	Reg *r1;
	Prog *p, *p1, *p2;
	int n, o;
	Adr a;

	p = r-&gt;prog;
	if(p-&gt;to.type != D_REG)
		FAIL(&#34;BOTCH: result not reg&#34;);
	n = p-&gt;to.reg;
	a = zprog.from;
	if(p-&gt;reg != NREG &amp;&amp; p-&gt;reg != p-&gt;to.reg) {
		a.type = D_REG;
		a.reg = p-&gt;reg;
	}
	if(debug[&#39;H&#39;])
		print(&#34;shiftprop\n%P&#34;, p);
	r1 = r;
	for(;;) {
		/* find first use of shift result; abort if shift operands or result are changed */
		r1 = uniqs(r1);
		if(r1 == R)
			FAIL(&#34;branch&#34;);
		if(uniqp(r1) == R)
			FAIL(&#34;merge&#34;);
		p1 = r1-&gt;prog;
		if(debug[&#39;H&#39;])
			print(&#34;\n%P&#34;, p1);
		switch(copyu(p1, &amp;p-&gt;to, A)) {
		case 0:	/* not used or set */
			if((p-&gt;from.type == D_REG &amp;&amp; copyu(p1, &amp;p-&gt;from, A) &gt; 1) ||
			   (a.type == D_REG &amp;&amp; copyu(p1, &amp;a, A) &gt; 1))
				FAIL(&#34;args modified&#34;);
			continue;
		case 3:	/* set, not used */
			FAIL(&#34;BOTCH: noref&#34;);
		}
		break;
	}
	/* check whether substitution can be done */
	switch(p1-&gt;as) {
	default:
		FAIL(&#34;non-dpi&#34;);
	case AAND:
	case AEOR:
	case AADD:
	case AADC:
	case AORR:
	case ASUB:
	case ARSB:
	case ASBC:
	case ARSC:
		if(p1-&gt;reg == n || (p1-&gt;reg == NREG &amp;&amp; p1-&gt;to.type == D_REG &amp;&amp; p1-&gt;to.reg == n)) {
			if(p1-&gt;from.type != D_REG)
				FAIL(&#34;can&#39;t swap&#34;);
			p1-&gt;reg = p1-&gt;from.reg;
			p1-&gt;from.reg = n;
			switch(p1-&gt;as) {
			case ASUB:
				p1-&gt;as = ARSB;
				break;
			case ARSB:
				p1-&gt;as = ASUB;
				break;
			case ASBC:
				p1-&gt;as = ARSC;
				break;
			case ARSC:
				p1-&gt;as = ASBC;
				break;
			}
			if(debug[&#39;H&#39;])
				print(&#34;\t=&gt;%P&#34;, p1);
		}
	case ABIC:
	case ACMP:
	case ACMN:
		if(p1-&gt;reg == n)
			FAIL(&#34;can&#39;t swap&#34;);
		if(p1-&gt;reg == NREG &amp;&amp; p1-&gt;to.reg == n)
			FAIL(&#34;shift result used twice&#34;);
	case AMVN:
		if(p1-&gt;from.type == D_SHIFT)
			FAIL(&#34;shift result used in shift&#34;);
		if(p1-&gt;from.type != D_REG || p1-&gt;from.reg != n)
			FAIL(&#34;BOTCH: where is it used?&#34;);
		break;
	}
	/* check whether shift result is used subsequently */
	p2 = p1;
	if(p1-&gt;to.reg != n)
	for (;;) {
		r1 = uniqs(r1);
		if(r1 == R)
			FAIL(&#34;inconclusive&#34;);
		p1 = r1-&gt;prog;
		if(debug[&#39;H&#39;])
			print(&#34;\n%P&#34;, p1);
		switch(copyu(p1, &amp;p-&gt;to, A)) {
		case 0:	/* not used or set */
			continue;
		case 3: /* set, not used */
			break;
		default:/* used */
			FAIL(&#34;reused&#34;);
		}
		break;
	}
	/* make the substitution */
	p2-&gt;from.type = D_SHIFT;
	p2-&gt;from.reg = NREG;
	o = p-&gt;reg;
	if(o == NREG)
		o = p-&gt;to.reg;
	switch(p-&gt;from.type){
	case D_CONST:
		o |= (p-&gt;from.offset&amp;0x1f)&lt;&lt;7;
		break;
	case D_REG:
		o |= (1&lt;&lt;4) | (p-&gt;from.reg&lt;&lt;8);
		break;
	}
	switch(p-&gt;as){
	case ASLL:
		o |= 0&lt;&lt;5;
		break;
	case ASRL:
		o |= 1&lt;&lt;5;
		break;
	case ASRA:
		o |= 2&lt;&lt;5;
		break;
	}
	p2-&gt;from.offset = o;
	if(debug[&#39;H&#39;])
		print(&#34;\t=&gt;%P\tSUCCEED\n&#34;, p2);
	return 1;
}

Reg*
findpre(Reg *r, Adr *v)
{
	Reg *r1;

	for(r1=uniqp(r); r1!=R; r=r1,r1=uniqp(r)) {
		if(uniqs(r1) != r)
			return R;
		switch(copyu(r1-&gt;prog, v, A)) {
		case 1: /* used */
		case 2: /* read-alter-rewrite */
			return R;
		case 3: /* set */
		case 4: /* set and used */
			return r1;
		}
	}
	return R;
}

Reg*
findinc(Reg *r, Reg *r2, Adr *v)
{
	Reg *r1;
	Prog *p;


	for(r1=uniqs(r); r1!=R &amp;&amp; r1!=r2; r=r1,r1=uniqs(r)) {
		if(uniqp(r1) != r)
			return R;
		switch(copyu(r1-&gt;prog, v, A)) {
		case 0: /* not touched */
			continue;
		case 4: /* set and used */
			p = r1-&gt;prog;
			if(p-&gt;as == AADD)
			if(p-&gt;from.type == D_CONST)
			if(p-&gt;from.offset &gt; -4096 &amp;&amp; p-&gt;from.offset &lt; 4096)
				return r1;
		default:
			return R;
		}
	}
	return R;
}

int
nochange(Reg *r, Reg *r2, Prog *p)
{
	Adr a[3];
	int i, n;

	if(r == r2)
		return 1;
	n = 0;
	if(p-&gt;reg != NREG &amp;&amp; p-&gt;reg != p-&gt;to.reg) {
		a[n].type = D_REG;
		a[n++].reg = p-&gt;reg;
	}
	switch(p-&gt;from.type) {
	case D_SHIFT:
		a[n].type = D_REG;
		a[n++].reg = p-&gt;from.offset&amp;0xf;
	case D_REG:
		a[n].type = D_REG;
		a[n++].reg = p-&gt;from.reg;
	}
	if(n == 0)
		return 1;
	for(; r!=R &amp;&amp; r!=r2; r=uniqs(r)) {
		p = r-&gt;prog;
		for(i=0; i&lt;n; i++)
			if(copyu(p, &amp;a[i], A) &gt; 1)
				return 0;
	}
	return 1;
}

int
findu1(Reg *r, Adr *v)
{
	for(; r != R; r = r-&gt;s1) {
		if(r-&gt;active)
			return 0;
		r-&gt;active = 1;
		switch(copyu(r-&gt;prog, v, A)) {
		case 1: /* used */
		case 2: /* read-alter-rewrite */
		case 4: /* set and used */
			return 1;
		case 3: /* set */
			return 0;
		}
		if(r-&gt;s2)
			if (findu1(r-&gt;s2, v))
				return 1;
	}
	return 0;
}

int
finduse(Reg *r, Adr *v)
{
	Reg *r1;

	for(r1=firstr; r1!=R; r1=r1-&gt;link)
		r1-&gt;active = 0;
	return findu1(r, v);
}

int
xtramodes(Reg *r, Adr *a)
{
	Reg *r1, *r2, *r3;
	Prog *p, *p1;
	Adr v;

	p = r-&gt;prog;
	if(debug[&#39;h&#39;] &amp;&amp; p-&gt;as == AMOVB &amp;&amp; p-&gt;from.type == D_OREG)	/* byte load */
		return 0;
	v = *a;
	v.type = D_REG;
	r1 = findpre(r, &amp;v);
	if(r1 != R) {
		p1 = r1-&gt;prog;
		if(p1-&gt;to.type == D_REG &amp;&amp; p1-&gt;to.reg == v.reg)
		switch(p1-&gt;as) {
		case AADD:
			if(p1-&gt;from.type == D_REG ||
			   (p1-&gt;from.type == D_SHIFT &amp;&amp; (p1-&gt;from.offset&amp;(1&lt;&lt;4)) == 0 &amp;&amp;
			    (p-&gt;as != AMOVB || (a == &amp;p-&gt;from &amp;&amp; (p1-&gt;from.offset&amp;~0xf) == 0))) ||
			   (p1-&gt;from.type == D_CONST &amp;&amp;
			    p1-&gt;from.offset &gt; -4096 &amp;&amp; p1-&gt;from.offset &lt; 4096))
			if(nochange(uniqs(r1), r, p1)) {
				if(a != &amp;p-&gt;from || v.reg != p-&gt;to.reg)
				if (finduse(r-&gt;s1, &amp;v)) {
					if(p1-&gt;reg == NREG || p1-&gt;reg == v.reg)
						/* pre-indexing */
						p-&gt;scond |= C_WBIT;
					else return 0;
				}
				switch (p1-&gt;from.type) {
				case D_REG:
					/* register offset */
					a-&gt;type = D_SHIFT;
					a-&gt;offset = p1-&gt;from.reg;
					break;
				case D_SHIFT:
					/* scaled register offset */
					a-&gt;type = D_SHIFT;
				case D_CONST:
					/* immediate offset */
					a-&gt;offset = p1-&gt;from.offset;
					break;
				}
				if(p1-&gt;reg != NREG)
					a-&gt;reg = p1-&gt;reg;
				excise(r1);
				return 1;
			}
			break;
		case AMOVW:
			if(p1-&gt;from.type == D_REG)
			if((r2 = findinc(r1, r, &amp;p1-&gt;from)) != R) {
			for(r3=uniqs(r2); r3-&gt;prog-&gt;as==ANOP; r3=uniqs(r3))
				;
			if(r3 == r) {
				/* post-indexing */
				p1 = r2-&gt;prog;
				a-&gt;reg = p1-&gt;to.reg;
				a-&gt;offset = p1-&gt;from.offset;
				p-&gt;scond |= C_PBIT;
				if(!finduse(r, &amp;r1-&gt;prog-&gt;to))
					excise(r1);
				excise(r2);
				return 1;
			}
			}
			break;
		}
	}
	if(a != &amp;p-&gt;from || a-&gt;reg != p-&gt;to.reg)
	if((r1 = findinc(r, R, &amp;v)) != R) {
		/* post-indexing */
		p1 = r1-&gt;prog;
		a-&gt;offset = p1-&gt;from.offset;
		p-&gt;scond |= C_PBIT;
		excise(r1);
		return 1;
	}
	return 0;
}

/*
 * return
 * 1 if v only used (and substitute),
 * 2 if read-alter-rewrite
 * 3 if set
 * 4 if set and used
 * 0 otherwise (not touched)
 */
int
copyu(Prog *p, Adr *v, Adr *s)
{

	switch(p-&gt;as) {

	default:
		if(debug[&#39;P&#39;])
			print(&#34; (?)&#34;);
		return 2;

	case AMOVM:
		if(v-&gt;type != D_REG)
			return 0;
		if(p-&gt;from.type == D_CONST) {	/* read reglist, read/rar */
			if(s != A) {
				if(p-&gt;from.offset&amp;(1&lt;&lt;v-&gt;reg))
					return 1;
				if(copysub(&amp;p-&gt;to, v, s, 1))
					return 1;
				return 0;
			}
			if(copyau(&amp;p-&gt;to, v)) {
				if(p-&gt;scond&amp;C_WBIT)
					return 2;
				return 1;
			}
			if(p-&gt;from.offset&amp;(1&lt;&lt;v-&gt;reg))
				return 1;
		} else {			/* read/rar, write reglist */
			if(s != A) {
				if(p-&gt;to.offset&amp;(1&lt;&lt;v-&gt;reg))
					return 1;
				if(copysub(&amp;p-&gt;from, v, s, 1))
					return 1;
				return 0;
			}
			if(copyau(&amp;p-&gt;from, v)) {
				if(p-&gt;scond&amp;C_WBIT)
					return 2;
				if(p-&gt;to.offset&amp;(1&lt;&lt;v-&gt;reg))
					return 4;
				return 1;
			}
			if(p-&gt;to.offset&amp;(1&lt;&lt;v-&gt;reg))
				return 3;
		}
		return 0;

	case ANOP:	/* read, write */
	case AMOVW:
	case AMOVF:
	case AMOVD:
	case AMOVH:
	case AMOVHU:
	case AMOVB:
	case AMOVBU:
	case AMOVDW:
	case AMOVWD:
	case AMOVFD:
	case AMOVDF:
		if(p-&gt;scond&amp;(C_WBIT|C_PBIT))
		if(v-&gt;type == D_REG) {
			if(p-&gt;from.type == D_OREG || p-&gt;from.type == D_SHIFT) {
				if(p-&gt;from.reg == v-&gt;reg)
					return 2;
			} else {
		  		if(p-&gt;to.reg == v-&gt;reg)
				return 2;
			}
		}
		if(s != A) {
			if(copysub(&amp;p-&gt;from, v, s, 1))
				return 1;
			if(!copyas(&amp;p-&gt;to, v))
				if(copysub(&amp;p-&gt;to, v, s, 1))
					return 1;
			return 0;
		}
		if(copyas(&amp;p-&gt;to, v)) {
			if(copyau(&amp;p-&gt;from, v))
				return 4;
			return 3;
		}
		if(copyau(&amp;p-&gt;from, v))
			return 1;
		if(copyau(&amp;p-&gt;to, v))
			return 1;
		return 0;


	case AADD:	/* read, read, write */
	case ASUB:
	case ARSB:
	case ASLL:
	case ASRL:
	case ASRA:
	case AORR:
	case AAND:
	case AEOR:
	case AMUL:
	case ADIV:
	case ADIVU:
	case AADDF:
	case AADDD:
	case ASUBF:
	case ASUBD:
	case AMULF:
	case AMULD:
	case ADIVF:
	case ADIVD:

	case ACMPF:
	case ACMPD:
	case ACMP:
	case ACMN:
	case ACASE:
		if(s != A) {
			if(copysub(&amp;p-&gt;from, v, s, 1))
				return 1;
			if(copysub1(p, v, s, 1))
				return 1;
			if(!copyas(&amp;p-&gt;to, v))
				if(copysub(&amp;p-&gt;to, v, s, 1))
					return 1;
			return 0;
		}
		if(copyas(&amp;p-&gt;to, v)) {
			if(p-&gt;reg == NREG)
				p-&gt;reg = p-&gt;to.reg;
			if(copyau(&amp;p-&gt;from, v))
				return 4;
			if(copyau1(p, v))
				return 4;
			return 3;
		}
		if(copyau(&amp;p-&gt;from, v))
			return 1;
		if(copyau1(p, v))
			return 1;
		if(copyau(&amp;p-&gt;to, v))
			return 1;
		return 0;

	case ABEQ:	/* read, read */
	case ABNE:
	case ABCS:
	case ABHS:
	case ABCC:
	case ABLO:
	case ABMI:
	case ABPL:
	case ABVS:
	case ABVC:
	case ABHI:
	case ABLS:
	case ABGE:
	case ABLT:
	case ABGT:
	case ABLE:
		if(s != A) {
			if(copysub(&amp;p-&gt;from, v, s, 1))
				return 1;
			return copysub1(p, v, s, 1);
		}
		if(copyau(&amp;p-&gt;from, v))
			return 1;
		if(copyau1(p, v))
			return 1;
		return 0;

	case AB:	/* funny */
		if(s != A) {
			if(copysub(&amp;p-&gt;to, v, s, 1))
				return 1;
			return 0;
		}
		if(copyau(&amp;p-&gt;to, v))
			return 1;
		return 0;

	case ARET:	/* funny */
		if(v-&gt;type == D_REG)
		if(v-&gt;reg == REGRET)
			return 2;
		if(v-&gt;type == D_FREG)
		if(v-&gt;reg == FREGRET)
			return 2;

	case ABL:	/* funny */
		if(v-&gt;type == D_REG) {
			if(v-&gt;reg &lt;= REGEXT &amp;&amp; v-&gt;reg &gt; exregoffset)
				return 2;
			if(v-&gt;reg == REGARG)
				return 2;
		}
		if(v-&gt;type == D_FREG)
			if(v-&gt;reg &lt;= FREGEXT &amp;&amp; v-&gt;reg &gt; exfregoffset)
				return 2;

		if(s != A) {
			if(copysub(&amp;p-&gt;to, v, s, 1))
				return 1;
			return 0;
		}
		if(copyau(&amp;p-&gt;to, v))
			return 4;
		return 3;

	case ATEXT:	/* funny */
		if(v-&gt;type == D_REG)
			if(v-&gt;reg == REGARG)
				return 3;
		return 0;
	}
	return 0;
}

int
a2type(Prog *p)
{

	switch(p-&gt;as) {

	case ACMP:
	case ACMN:

	case AADD:
	case ASUB:
	case ARSB:
	case ASLL:
	case ASRL:
	case ASRA:
	case AORR:
	case AAND:
	case AEOR:
	case AMUL:
	case ADIV:
	case ADIVU:
		return D_REG;

	case ACMPF:
	case ACMPD:

	case AADDF:
	case AADDD:
	case ASUBF:
	case ASUBD:
	case AMULF:
	case AMULD:
	case ADIVF:
	case ADIVD:
		return D_FREG;
	}
	return D_NONE;
}

/*
 * direct reference,
 * could be set/use depending on
 * semantics
 */
int
copyas(Adr *a, Adr *v)
{

	if(regtyp(v)) {
		if(a-&gt;type == v-&gt;type)
		if(a-&gt;reg == v-&gt;reg)
			return 1;
	} else if(v-&gt;type == D_CONST) {		/* for constprop */
		if(a-&gt;type == v-&gt;type)
		if(a-&gt;name == v-&gt;name)
		if(a-&gt;sym == v-&gt;sym)
		if(a-&gt;reg == v-&gt;reg)
		if(a-&gt;offset == v-&gt;offset)
			return 1;
	}
	return 0;
}

/*
 * either direct or indirect
 */
int
copyau(Adr *a, Adr *v)
{

	if(copyas(a, v))
		return 1;
	if(v-&gt;type == D_REG) {
		if(a-&gt;type == D_OREG) {
			if(v-&gt;reg == a-&gt;reg)
				return 1;
		} else if(a-&gt;type == D_SHIFT) {
			if((a-&gt;offset&amp;0xf) == v-&gt;reg)
				return 1;
			if((a-&gt;offset&amp;(1&lt;&lt;4)) &amp;&amp; (a-&gt;offset&gt;&gt;8) == v-&gt;reg)
				return 1;
		}
	}
	return 0;
}

int
copyau1(Prog *p, Adr *v)
{

	if(regtyp(v)) {
		if(a2type(p) == v-&gt;type)
		if(p-&gt;reg == v-&gt;reg) {
			if(a2type(p) != v-&gt;type)
				print(&#34;botch a2type %P\n&#34;, p);
			return 1;
		}
	}
	return 0;
}

/*
 * substitute s for v in a
 * return failure to substitute
 */
int
copysub(Adr *a, Adr *v, Adr *s, int f)
{

	if(f)
	if(copyau(a, v)) {
		if(a-&gt;type == D_SHIFT) {
			if((a-&gt;offset&amp;0xf) == v-&gt;reg)
				a-&gt;offset = (a-&gt;offset&amp;~0xf)|s-&gt;reg;
			if((a-&gt;offset&amp;(1&lt;&lt;4)) &amp;&amp; (a-&gt;offset&gt;&gt;8) == v-&gt;reg)
				a-&gt;offset = (a-&gt;offset&amp;~(0xf&lt;&lt;8))|(s-&gt;reg&lt;&lt;8);
		} else
			a-&gt;reg = s-&gt;reg;
	}
	return 0;
}

int
copysub1(Prog *p1, Adr *v, Adr *s, int f)
{

	if(f)
	if(copyau1(p1, v))
		p1-&gt;reg = s-&gt;reg;
	return 0;
}

struct {
	int opcode;
	int notopcode;
	int scond;
	int notscond;
} predinfo[]  = {
	{ ABEQ,	ABNE,	0x0,	0x1, },
	{ ABNE,	ABEQ,	0x1,	0x0, },
	{ ABCS,	ABCC,	0x2,	0x3, },
	{ ABHS,	ABLO,	0x2,	0x3, },
	{ ABCC,	ABCS,	0x3,	0x2, },
	{ ABLO,	ABHS,	0x3,	0x2, },
	{ ABMI,	ABPL,	0x4,	0x5, },
	{ ABPL,	ABMI,	0x5,	0x4, },
	{ ABVS,	ABVC,	0x6,	0x7, },
	{ ABVC,	ABVS,	0x7,	0x6, },
	{ ABHI,	ABLS,	0x8,	0x9, },
	{ ABLS,	ABHI,	0x9,	0x8, },
	{ ABGE,	ABLT,	0xA,	0xB, },
	{ ABLT,	ABGE,	0xB,	0xA, },
	{ ABGT,	ABLE,	0xC,	0xD, },
	{ ABLE,	ABGT,	0xD,	0xC, },
};

typedef struct {
	Reg *start;
	Reg *last;
	Reg *end;
	int len;
} Joininfo;

enum {
	Join,
	Split,
	End,
	Branch,
	Setcond,
	Toolong
};

enum {
	Falsecond,
	Truecond,
	Delbranch,
	Keepbranch
};

int
isbranch(Prog *p)
{
	return (ABEQ &lt;= p-&gt;as) &amp;&amp; (p-&gt;as &lt;= ABLE);
}

int
predicable(Prog *p)
{
	if (isbranch(p)
		|| p-&gt;as == ANOP
		|| p-&gt;as == AXXX
		|| p-&gt;as == ADATA
		|| p-&gt;as == AGLOBL
		|| p-&gt;as == AGOK
		|| p-&gt;as == AHISTORY
		|| p-&gt;as == ANAME
		|| p-&gt;as == ASIGNAME
		|| p-&gt;as == ATEXT
		|| p-&gt;as == AWORD
		|| p-&gt;as == ADYNT
		|| p-&gt;as == AINIT
		|| p-&gt;as == ABCASE
		|| p-&gt;as == ACASE)
		return 0;
	return 1;
}

/*
 * Depends on an analysis of the encodings performed by 5l.
 * These seem to be all of the opcodes that lead to the &#34;S&#34; bit
 * being set in the instruction encodings.
 *
 * C_SBIT may also have been set explicitly in p-&gt;scond.
 */
int
modifiescpsr(Prog *p)
{
	return (p-&gt;scond&amp;C_SBIT)
		|| p-&gt;as == ATST
		|| p-&gt;as == ATEQ
		|| p-&gt;as == ACMN
		|| p-&gt;as == ACMP
		|| p-&gt;as == AMULU
		|| p-&gt;as == ADIVU
		|| p-&gt;as == AMUL
		|| p-&gt;as == ADIV
		|| p-&gt;as == AMOD
		|| p-&gt;as == AMODU
		|| p-&gt;as == ABL;
}

/*
 * Find the maximal chain of instructions starting with r which could
 * be executed conditionally
 */
int
joinsplit(Reg *r, Joininfo *j)
{
	j-&gt;start = r;
	j-&gt;last = r;
	j-&gt;len = 0;
	do {
		if (r-&gt;p2 &amp;&amp; (r-&gt;p1 || r-&gt;p2-&gt;p2link)) {
			j-&gt;end = r;
			return Join;
		}
		if (r-&gt;s1 &amp;&amp; r-&gt;s2) {
			j-&gt;end = r;
			return Split;
		}
		j-&gt;last = r;
		if (r-&gt;prog-&gt;as != ANOP)
			j-&gt;len++;
		if (!r-&gt;s1 &amp;&amp; !r-&gt;s2) {
			j-&gt;end = r-&gt;link;
			return End;
		}
		if (r-&gt;s2) {
			j-&gt;end = r-&gt;s2;
			return Branch;
		}
		if (modifiescpsr(r-&gt;prog)) {
			j-&gt;end = r-&gt;s1;
			return Setcond;
		}
		r = r-&gt;s1;
	} while (j-&gt;len &lt; 4);
	j-&gt;end = r;
	return Toolong;
}

Reg *
successor(Reg *r)
{
	if (r-&gt;s1)
		return r-&gt;s1;
	else
		return r-&gt;s2;
}

void
applypred(Reg *rstart, Joininfo *j, int cond, int branch)
{
	int pred;
	Reg *r;

	if(j-&gt;len == 0)
		return;
	if (cond == Truecond)
		pred = predinfo[rstart-&gt;prog-&gt;as - ABEQ].scond;
	else
		pred = predinfo[rstart-&gt;prog-&gt;as - ABEQ].notscond;

	for (r = j-&gt;start; ; r = successor(r)) {
		if (r-&gt;prog-&gt;as == AB) {
			if (r != j-&gt;last || branch == Delbranch)
				excise(r);
			else {
			  if (cond == Truecond)
				r-&gt;prog-&gt;as = predinfo[rstart-&gt;prog-&gt;as - ABEQ].opcode;
			  else
				r-&gt;prog-&gt;as = predinfo[rstart-&gt;prog-&gt;as - ABEQ].notopcode;
			}
		}
		else if (predicable(r-&gt;prog))
			r-&gt;prog-&gt;scond = (r-&gt;prog-&gt;scond&amp;~C_SCOND)|pred;
		if (r-&gt;s1 != r-&gt;link) {
			r-&gt;s1 = r-&gt;link;
			r-&gt;link-&gt;p1 = r;
		}
		if (r == j-&gt;last)
			break;
	}
}

void
predicate(void)
{
	Reg *r;
	int t1, t2;
	Joininfo j1, j2;

	for(r=firstr; r!=R; r=r-&gt;link) {
		if (isbranch(r-&gt;prog)) {
			t1 = joinsplit(r-&gt;s1, &amp;j1);
			t2 = joinsplit(r-&gt;s2, &amp;j2);
			if(j1.last-&gt;link != j2.start)
				continue;
			if(j1.end == j2.end)
			if((t1 == Branch &amp;&amp; (t2 == Join || t2 == Setcond)) ||
			   (t2 == Join &amp;&amp; (t1 == Join || t1 == Setcond))) {
				applypred(r, &amp;j1, Falsecond, Delbranch);
				applypred(r, &amp;j2, Truecond, Delbranch);
				excise(r);
				continue;
			}
			if(t1 == End || t1 == Branch) {
				applypred(r, &amp;j1, Falsecond, Keepbranch);
				excise(r);
				continue;
			}
		}
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
