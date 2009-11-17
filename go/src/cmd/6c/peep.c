<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/6c/peep.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/6c/peep.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/6c/peep.c
// http://code.google.com/p/inferno-os/source/browse/utils/6c/peep.c
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

static int
needc(Prog *p)
{
	while(p != P) {
		switch(p-&gt;as) {
		case AADCL:
		case AADCQ:
		case ASBBL:
		case ASBBQ:
		case ARCRL:
		case ARCRQ:
			return 1;
		case AADDL:
		case AADDQ:
		case ASUBL:
		case ASUBQ:
		case AJMP:
		case ARET:
		case ACALL:
			return 0;
		default:
			if(p-&gt;to.type == D_BRANCH)
				return 0;
		}
		p = p-&gt;link;
	}
	return 0;
}

static Reg*
rnops(Reg *r)
{
	Prog *p;
	Reg *r1;

	if(r != R)
	for(;;){
		p = r-&gt;prog;
		if(p-&gt;as != ANOP || p-&gt;from.type != D_NONE || p-&gt;to.type != D_NONE)
			break;
		r1 = uniqs(r);
		if(r1 == R)
			break;
		r = r1;
	}
	return r;
}

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

	pc = 0;	/* speculating it won&#39;t kill */

loop1:

	t = 0;
	for(r=firstr; r!=R; r=r-&gt;link) {
		p = r-&gt;prog;
		switch(p-&gt;as) {
		case AMOVL:
		case AMOVQ:
		case AMOVSS:
		case AMOVSD:
			if(regtyp(&amp;p-&gt;to))
			if(regtyp(&amp;p-&gt;from)) {
				if(copyprop(r)) {
					excise(r);
					t++;
				} else
				if(subprop(r) &amp;&amp; copyprop(r)) {
					excise(r);
					t++;
				}
			}
			break;

		case AMOVBLZX:
		case AMOVWLZX:
		case AMOVBLSX:
		case AMOVWLSX:
			if(regtyp(&amp;p-&gt;to)) {
				r1 = rnops(uniqs(r));
				if(r1 != R) {
					p1 = r1-&gt;prog;
					if(p-&gt;as == p1-&gt;as &amp;&amp; p-&gt;to.type == p1-&gt;from.type){
						p1-&gt;as = AMOVL;
						t++;
					}
				}
			}
			break;

		case AMOVBQSX:
		case AMOVBQZX:
		case AMOVWQSX:
		case AMOVWQZX:
		case AMOVLQSX:
		case AMOVLQZX:
			if(regtyp(&amp;p-&gt;to)) {
				r1 = rnops(uniqs(r));
				if(r1 != R) {
					p1 = r1-&gt;prog;
					if(p-&gt;as == p1-&gt;as &amp;&amp; p-&gt;to.type == p1-&gt;from.type){
						p1-&gt;as = AMOVQ;
						t++;
					}
				}
			}
			break;

		case AADDL:
		case AADDQ:
		case AADDW:
			if(p-&gt;from.type != D_CONST || needc(p-&gt;link))
				break;
			if(p-&gt;from.offset == -1){
				if(p-&gt;as == AADDQ)
					p-&gt;as = ADECQ;
				else if(p-&gt;as == AADDL)
					p-&gt;as = ADECL;
				else
					p-&gt;as = ADECW;
				p-&gt;from = zprog.from;
			}
			else if(p-&gt;from.offset == 1){
				if(p-&gt;as == AADDQ)
					p-&gt;as = AINCQ;
				else if(p-&gt;as == AADDL)
					p-&gt;as = AINCL;
				else
					p-&gt;as = AINCW;
				p-&gt;from = zprog.from;
			}
			break;

		case ASUBL:
		case ASUBQ:
		case ASUBW:
			if(p-&gt;from.type != D_CONST || needc(p-&gt;link))
				break;
			if(p-&gt;from.offset == -1) {
				if(p-&gt;as == ASUBQ)
					p-&gt;as = AINCQ;
				else if(p-&gt;as == ASUBL)
					p-&gt;as = AINCL;
				else
					p-&gt;as = AINCW;
				p-&gt;from = zprog.from;
			}
			else if(p-&gt;from.offset == 1){
				if(p-&gt;as == ASUBQ)
					p-&gt;as = ADECQ;
				else if(p-&gt;as == ASUBL)
					p-&gt;as = ADECL;
				else
					p-&gt;as = ADECW;
				p-&gt;from = zprog.from;
			}
			break;
		}
	}
	if(t)
		goto loop1;
}

void
excise(Reg *r)
{
	Prog *p;

	p = r-&gt;prog;
	p-&gt;as = ANOP;
	p-&gt;from = zprog.from;
	p-&gt;to = zprog.to;
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
	int t;

	t = a-&gt;type;
	if(t &gt;= D_AX &amp;&amp; t &lt;= D_R15)
		return 1;
	if(t &gt;= D_X0 &amp;&amp; t &lt;= D_X0+15)
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
		case ACALL:
			return 0;

		case AIMULL:
		case AIMULQ:
		case AIMULW:
			if(p-&gt;to.type != D_NONE)
				break;

		case ADIVB:
		case ADIVL:
		case ADIVQ:
		case ADIVW:
		case AIDIVB:
		case AIDIVL:
		case AIDIVQ:
		case AIDIVW:
		case AIMULB:
		case AMULB:
		case AMULL:
		case AMULQ:
		case AMULW:

		case AROLB:
		case AROLL:
		case AROLQ:
		case AROLW:
		case ARORB:
		case ARORL:
		case ARORQ:
		case ARORW:
		case ASALB:
		case ASALL:
		case ASALQ:
		case ASALW:
		case ASARB:
		case ASARL:
		case ASARQ:
		case ASARW:
		case ASHLB:
		case ASHLL:
		case ASHLQ:
		case ASHLW:
		case ASHRB:
		case ASHRL:
		case ASHRQ:
		case ASHRW:

		case AREP:
		case AREPN:

		case ACWD:
		case ACDQ:
		case ACQO:

		case ASTOSB:
		case ASTOSL:
		case ASTOSQ:
		case AMOVSB:
		case AMOVSL:
		case AMOVSQ:
			return 0;

		case AMOVL:
		case AMOVQ:
			if(p-&gt;to.type == v1-&gt;type)
				goto gotit;
			break;
		}
		if(copyau(&amp;p-&gt;from, v2) ||
		   copyau(&amp;p-&gt;to, v2))
			break;
		if(copysub(&amp;p-&gt;from, v1, v2, 0) ||
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
		copysub(&amp;p-&gt;to, v1, v2, 1);
		if(debug[&#39;P&#39;])
			print(&#34;%P\n&#34;, r-&gt;prog);
	}
	t = v1-&gt;type;
	v1-&gt;type = v2-&gt;type;
	v2-&gt;type = t;
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
				print(&#34;; %D rar; return 0\n&#34;, v2);
			return 0;

		case 3:	/* set */
			if(debug[&#39;P&#39;])
				print(&#34;; %D set; return 1\n&#34;, v2);
			return 1;

		case 1:	/* used, substitute */
		case 4:	/* use and set */
			if(f) {
				if(!debug[&#39;P&#39;])
					return 0;
				if(t == 4)
					print(&#34;; %D used+set and f=%d; return 0\n&#34;, v2, f);
				else
					print(&#34;; %D used and f=%d; return 0\n&#34;, v2, f);
				return 0;
			}
			if(copyu(p, v2, v1)) {
				if(debug[&#39;P&#39;])
					print(&#34;; sub fail; return 0\n&#34;);
				return 0;
			}
			if(debug[&#39;P&#39;])
				print(&#34;; sub %D/%D&#34;, v2, v1);
			if(t == 4) {
				if(debug[&#39;P&#39;])
					print(&#34;; %D used+set; return 1\n&#34;, v2);
				return 1;
			}
			break;
		}
		if(!f) {
			t = copyu(p, v1, A);
			if(!f &amp;&amp; (t == 2 || t == 3 || t == 4)) {
				f = 1;
				if(debug[&#39;P&#39;])
					print(&#34;; %D set and !f; f=%d&#34;, v1, f);
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
			print(&#34;unknown op %A\n&#34;, p-&gt;as);
		/* SBBL; ADCL; FLD1; SAHF */
		return 2;


	case ANEGB:
	case ANEGW:
	case ANEGL:
	case ANEGQ:
	case ANOTB:
	case ANOTW:
	case ANOTL:
	case ANOTQ:
		if(copyas(&amp;p-&gt;to, v))
			return 2;
		break;

	case ALEAL:	/* lhs addr, rhs store */
	case ALEAQ:
		if(copyas(&amp;p-&gt;from, v))
			return 2;


	case ANOP:	/* rhs store */
	case AMOVL:
	case AMOVQ:
	case AMOVBLSX:
	case AMOVBLZX:
	case AMOVBQSX:
	case AMOVBQZX:
	case AMOVLQSX:
	case AMOVLQZX:
	case AMOVWLSX:
	case AMOVWLZX:
	case AMOVWQSX:
	case AMOVWQZX:

	case AMOVSS:
	case AMOVSD:
	case ACVTSD2SL:
	case ACVTSD2SQ:
	case ACVTSD2SS:
	case ACVTSL2SD:
	case ACVTSL2SS:
	case ACVTSQ2SD:
	case ACVTSQ2SS:
	case ACVTSS2SD:
	case ACVTSS2SL:
	case ACVTSS2SQ:
	case ACVTTSD2SL:
	case ACVTTSD2SQ:
	case ACVTTSS2SL:
	case ACVTTSS2SQ:
		if(copyas(&amp;p-&gt;to, v)) {
			if(s != A)
				return copysub(&amp;p-&gt;from, v, s, 1);
			if(copyau(&amp;p-&gt;from, v))
				return 4;
			return 3;
		}
		goto caseread;

	case AROLB:
	case AROLL:
	case AROLQ:
	case AROLW:
	case ARORB:
	case ARORL:
	case ARORQ:
	case ARORW:
	case ASALB:
	case ASALL:
	case ASALQ:
	case ASALW:
	case ASARB:
	case ASARL:
	case ASARQ:
	case ASARW:
	case ASHLB:
	case ASHLL:
	case ASHLQ:
	case ASHLW:
	case ASHRB:
	case ASHRL:
	case ASHRQ:
	case ASHRW:
		if(copyas(&amp;p-&gt;to, v))
			return 2;
		if(copyas(&amp;p-&gt;from, v))
			if(p-&gt;from.type == D_CX)
				return 2;
		goto caseread;

	case AADDB:	/* rhs rar */
	case AADDL:
	case AADDQ:
	case AADDW:
	case AANDB:
	case AANDL:
	case AANDQ:
	case AANDW:
	case ADECL:
	case ADECQ:
	case ADECW:
	case AINCL:
	case AINCQ:
	case AINCW:
	case ASUBB:
	case ASUBL:
	case ASUBQ:
	case ASUBW:
	case AORB:
	case AORL:
	case AORQ:
	case AORW:
	case AXORB:
	case AXORL:
	case AXORQ:
	case AXORW:
	case AMOVB:
	case AMOVW:

	case AADDSD:
	case AADDSS:
	case ACMPSD:
	case ACMPSS:
	case ADIVSD:
	case ADIVSS:
	case AMAXSD:
	case AMAXSS:
	case AMINSD:
	case AMINSS:
	case AMULSD:
	case AMULSS:
	case ARCPSS:
	case ARSQRTSS:
	case ASQRTSD:
	case ASQRTSS:
	case ASUBSD:
	case ASUBSS:
	case AXORPD:
		if(copyas(&amp;p-&gt;to, v))
			return 2;
		goto caseread;

	case ACMPL:	/* read only */
	case ACMPW:
	case ACMPB:
	case ACMPQ:

	case ACOMISD:
	case ACOMISS:
	case AUCOMISD:
	case AUCOMISS:
	caseread:
		if(s != A) {
			if(copysub(&amp;p-&gt;from, v, s, 1))
				return 1;
			return copysub(&amp;p-&gt;to, v, s, 1);
		}
		if(copyau(&amp;p-&gt;from, v))
			return 1;
		if(copyau(&amp;p-&gt;to, v))
			return 1;
		break;

	case AJGE:	/* no reference */
	case AJNE:
	case AJLE:
	case AJEQ:
	case AJHI:
	case AJLS:
	case AJMI:
	case AJPL:
	case AJGT:
	case AJLT:
	case AJCC:
	case AJCS:

	case AADJSP:
	case AWAIT:
	case ACLD:
		break;

	case AIMULL:
	case AIMULQ:
	case AIMULW:
		if(p-&gt;to.type != D_NONE) {
			if(copyas(&amp;p-&gt;to, v))
				return 2;
			goto caseread;
		}

	case ADIVB:
	case ADIVL:
	case ADIVQ:
	case ADIVW:
	case AIDIVB:
	case AIDIVL:
	case AIDIVQ:
	case AIDIVW:
	case AIMULB:
	case AMULB:
	case AMULL:
	case AMULQ:
	case AMULW:

	case ACWD:
	case ACDQ:
	case ACQO:
		if(v-&gt;type == D_AX || v-&gt;type == D_DX)
			return 2;
		goto caseread;

	case AREP:
	case AREPN:
		if(v-&gt;type == D_CX)
			return 2;
		goto caseread;

	case AMOVSB:
	case AMOVSL:
	case AMOVSQ:
		if(v-&gt;type == D_DI || v-&gt;type == D_SI)
			return 2;
		goto caseread;

	case ASTOSB:
	case ASTOSL:
	case ASTOSQ:
		if(v-&gt;type == D_AX || v-&gt;type == D_DI)
			return 2;
		goto caseread;

	case AJMP:	/* funny */
		if(s != A) {
			if(copysub(&amp;p-&gt;to, v, s, 1))
				return 1;
			return 0;
		}
		if(copyau(&amp;p-&gt;to, v))
			return 1;
		return 0;

	case ARET:	/* funny */
		if(v-&gt;type == REGRET || v-&gt;type == FREGRET)
			return 2;
		if(s != A)
			return 1;
		return 3;

	case ACALL:	/* funny */
		if(REGEXT &amp;&amp; v-&gt;type &lt;= REGEXT &amp;&amp; v-&gt;type &gt; exregoffset)
			return 2;
		if(REGARG &gt;= 0 &amp;&amp; v-&gt;type == REGARG)
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
		if(REGARG &gt;= 0 &amp;&amp; v-&gt;type == REGARG)
			return 3;
		return 0;
	}
	return 0;
}

/*
 * direct reference,
 * could be set/use depending on
 * semantics
 */
int
copyas(Adr *a, Adr *v)
{
	if(a-&gt;type != v-&gt;type)
		return 0;
	if(regtyp(v))
		return 1;
	if(v-&gt;type == D_AUTO || v-&gt;type == D_PARAM)
		if(v-&gt;offset == a-&gt;offset)
			return 1;
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
	if(regtyp(v)) {
		if(a-&gt;type-D_INDIR == v-&gt;type)
			return 1;
		if(a-&gt;index == v-&gt;type)
			return 1;
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
	int t;

	if(copyas(a, v)) {
		t = s-&gt;type;
		if(t &gt;= D_AX &amp;&amp; t &lt;= D_R15 || t &gt;= D_X0 &amp;&amp; t &lt;= D_X0+15) {
			if(f)
				a-&gt;type = t;
		}
		return 0;
	}
	if(regtyp(v)) {
		t = v-&gt;type;
		if(a-&gt;type == t+D_INDIR) {
			if((s-&gt;type == D_BP || s-&gt;type == D_R13) &amp;&amp; a-&gt;index != D_NONE)
				return 1;	/* can&#39;t use BP-base with index */
			if(f)
				a-&gt;type = s-&gt;type+D_INDIR;
//			return 0;
		}
		if(a-&gt;index == t) {
			if(f)
				a-&gt;index = s-&gt;type;
			return 0;
		}
		return 0;
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
