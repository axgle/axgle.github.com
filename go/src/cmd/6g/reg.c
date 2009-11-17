<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/6g/reg.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/6g/reg.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Derived from Inferno utils/6c/reg.c
// http://code.google.com/p/inferno-os/source/browse/utils/6c/reg.c
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

#include &#34;gg.h&#34;
#undef	EXTERN
#define	EXTERN
#include &#34;opt.h&#34;

#define	P2R(p)	(Reg*)(p-&gt;reg)

static	int	first	= 1;

Reg*
rega(void)
{
	Reg *r;

	r = freer;
	if(r == R) {
		r = mal(sizeof(*r));
	} else
		freer = r-&gt;link;

	*r = zreg;
	return r;
}

int
rcmp(const void *a1, const void *a2)
{
	Rgn *p1, *p2;
	int c1, c2;

	p1 = (Rgn*)a1;
	p2 = (Rgn*)a2;
	c1 = p2-&gt;cost;
	c2 = p1-&gt;cost;
	if(c1 -= c2)
		return c1;
	return p2-&gt;varno - p1-&gt;varno;
}

void
setoutvar(void)
{
	Type *t;
	Node *n;
	Addr a;
	Iter save;
	Bits bit;
	int z;

	t = structfirst(&amp;save, getoutarg(curfn-&gt;type));
	while(t != T) {
		n = nodarg(t, 1);
		a = zprog.from;
		naddr(n, &amp;a, 0);
		bit = mkvar(R, &amp;a);
		for(z=0; z&lt;BITS; z++)
			ovar.b[z] |= bit.b[z];
		t = structnext(&amp;save);
	}
//if(bany(b))
//print(&#34;ovars = %Q\n&#34;, &amp;ovar);
}

void
regopt(Prog *firstp)
{
	Reg *r, *r1;
	Prog *p;
	int i, z, nr;
	uint32 vreg;
	Bits bit;

	if(first) {
		fmtinstall(&#39;Q&#39;, Qconv);
		exregoffset = D_R13;	// R14,R15 are external
		first = 0;
	}

	// count instructions
	nr = 0;
	for(p=firstp; p!=P; p=p-&gt;link)
		nr++;
	// if too big dont bother
	if(nr &gt;= 10000) {
//		print(&#34;********** %S is too big (%d)\n&#34;, curfn-&gt;nname-&gt;sym, nr);
		return;
	}

	r1 = R;
	firstr = R;
	lastr = R;
	nvar = 0;
	regbits = RtoB(D_SP);
	for(z=0; z&lt;BITS; z++) {
		externs.b[z] = 0;
		params.b[z] = 0;
		consts.b[z] = 0;
		addrs.b[z] = 0;
		ovar.b[z] = 0;
	}

	// build list of return variables
	setoutvar();

	/*
	 * pass 1
	 * build aux data structure
	 * allocate pcs
	 * find use and set of variables
	 */
	nr = 0;
	for(p=firstp; p!=P; p=p-&gt;link) {
		switch(p-&gt;as) {
		case ADATA:
		case AGLOBL:
		case ANAME:
		case ASIGNAME:
			continue;
		}
		r = rega();
		nr++;
		if(firstr == R) {
			firstr = r;
			lastr = r;
		} else {
			lastr-&gt;link = r;
			r-&gt;p1 = lastr;
			lastr-&gt;s1 = r;
			lastr = r;
		}
		r-&gt;prog = p;
		p-&gt;reg = r;

		r1 = r-&gt;p1;
		if(r1 != R) {
			switch(r1-&gt;prog-&gt;as) {
			case ARET:
			case AJMP:
			case AIRETL:
			case AIRETQ:
				r-&gt;p1 = R;
				r1-&gt;s1 = R;
			}
		}

		bit = mkvar(r, &amp;p-&gt;from);
		if(bany(&amp;bit))
		switch(p-&gt;as) {
		/*
		 * funny
		 */
		case ALEAL:
		case ALEAQ:
			for(z=0; z&lt;BITS; z++)
				addrs.b[z] |= bit.b[z];
			break;

		/*
		 * left side read
		 */
		default:
			for(z=0; z&lt;BITS; z++)
				r-&gt;use1.b[z] |= bit.b[z];
			break;

		/*
		 * left side read+write
		 */
		case AXCHGB:
		case AXCHGW:
		case AXCHGL:
		case AXCHGQ:
			for(z=0; z&lt;BITS; z++) {
				r-&gt;use1.b[z] |= bit.b[z];
				r-&gt;set.b[z] |= bit.b[z];
			}
			break;
		}

		bit = mkvar(r, &amp;p-&gt;to);
		if(bany(&amp;bit))
		switch(p-&gt;as) {
		default:
			yyerror(&#34;reg: unknown op: %A&#34;, p-&gt;as);
			break;

		/*
		 * right side read
		 */
		case ACMPB:
		case ACMPL:
		case ACMPQ:
		case ACMPW:
		case ACOMISS:
		case ACOMISD:
		case AUCOMISS:
		case AUCOMISD:
			for(z=0; z&lt;BITS; z++)
				r-&gt;use2.b[z] |= bit.b[z];
			break;

		/*
		 * right side write
		 */
		case ANOP:
		case AMOVL:
		case AMOVQ:
		case AMOVB:
		case AMOVW:
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
			for(z=0; z&lt;BITS; z++)
				r-&gt;set.b[z] |= bit.b[z];
			break;

		/*
		 * right side read+write
		 */
		case AINCB:
		case AINCL:
		case AINCQ:
		case AINCW:
		case ADECB:
		case ADECL:
		case ADECQ:
		case ADECW:

		case AADDB:
		case AADDL:
		case AADDQ:
		case AADDW:
		case AANDB:
		case AANDL:
		case AANDQ:
		case AANDW:
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
		case ASALB:
		case ASALL:
		case ASALQ:
		case ASALW:
		case ASARB:
		case ASARL:
		case ASARQ:
		case ASARW:
		case ARCLB:
		case ARCLL:
		case ARCLQ:
		case ARCLW:
		case ARCRB:
		case ARCRL:
		case ARCRQ:
		case ARCRW:
		case AROLB:
		case AROLL:
		case AROLQ:
		case AROLW:
		case ARORB:
		case ARORL:
		case ARORQ:
		case ARORW:
		case ASHLB:
		case ASHLL:
		case ASHLQ:
		case ASHLW:
		case ASHRB:
		case ASHRL:
		case ASHRQ:
		case ASHRW:
		case AIMULL:
		case AIMULQ:
		case AIMULW:
		case ANEGL:
		case ANEGQ:
		case ANOTL:
		case ANOTQ:
		case AADCL:
		case AADCQ:
		case ASBBL:
		case ASBBQ:

		case AXCHGB:
		case AXCHGW:
		case AXCHGL:
		case AXCHGQ:

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
			for(z=0; z&lt;BITS; z++) {
				r-&gt;set.b[z] |= bit.b[z];
				r-&gt;use2.b[z] |= bit.b[z];
			}
			break;

		/*
		 * funny
		 */
		case ACALL:
			for(z=0; z&lt;BITS; z++)
				addrs.b[z] |= bit.b[z];
			break;
		}

		switch(p-&gt;as) {
		case AIMULL:
		case AIMULQ:
		case AIMULW:
			if(p-&gt;to.type != D_NONE)
				break;

		case AIDIVB:
		case AIDIVL:
		case AIDIVQ:
		case AIDIVW:
		case AIMULB:
		case ADIVB:
		case ADIVL:
		case ADIVQ:
		case ADIVW:
		case AMULB:
		case AMULL:
		case AMULQ:
		case AMULW:

		case ACWD:
		case ACDQ:
		case ACQO:
			r-&gt;regu |= RtoB(D_AX) | RtoB(D_DX);
			break;

		case AREP:
		case AREPN:
		case ALOOP:
		case ALOOPEQ:
		case ALOOPNE:
			r-&gt;regu |= RtoB(D_CX);
			break;

		case AMOVSB:
		case AMOVSL:
		case AMOVSQ:
		case AMOVSW:
		case ACMPSB:
		case ACMPSL:
		case ACMPSQ:
		case ACMPSW:
			r-&gt;regu |= RtoB(D_SI) | RtoB(D_DI);
			break;

		case ASTOSB:
		case ASTOSL:
		case ASTOSQ:
		case ASTOSW:
		case ASCASB:
		case ASCASL:
		case ASCASQ:
		case ASCASW:
			r-&gt;regu |= RtoB(D_AX) | RtoB(D_DI);
			break;

		case AINSB:
		case AINSL:
		case AINSW:
		case AOUTSB:
		case AOUTSL:
		case AOUTSW:
			r-&gt;regu |= RtoB(D_DI) | RtoB(D_DX);
			break;
		}
	}
	if(firstr == R)
		return;

	if(debug[&#39;R&#39;] &amp;&amp; debug[&#39;v&#39;])
		dumpit(&#34;pass1&#34;, firstr);

	/*
	 * pass 2
	 * turn branch references to pointers
	 * build back pointers
	 */
	for(r=firstr; r!=R; r=r-&gt;link) {
		p = r-&gt;prog;
		if(p-&gt;to.type == D_BRANCH) {
			if(p-&gt;to.branch == P)
				fatal(&#34;pnil %P&#34;, p);
			r1 = p-&gt;to.branch-&gt;reg;
			if(r1 == R)
				fatal(&#34;rnil %P&#34;, p);
			if(r1 == r) {
				//fatal(&#34;ref to self %P&#34;, p);
				continue;
			}
			r-&gt;s2 = r1;
			r-&gt;p2link = r1-&gt;p2;
			r1-&gt;p2 = r;
		}
	}

	if(debug[&#39;R&#39;] &amp;&amp; debug[&#39;v&#39;])
		dumpit(&#34;pass2&#34;, firstr);

	/*
	 * pass 2.5
	 * find looping structure
	 */
	for(r = firstr; r != R; r = r-&gt;link)
		r-&gt;active = 0;
	change = 0;
	loopit(firstr, nr);

	if(debug[&#39;R&#39;] &amp;&amp; debug[&#39;v&#39;])
		dumpit(&#34;pass2.5&#34;, firstr);

	/*
	 * pass 3
	 * iterate propagating usage
	 * 	back until flow graph is complete
	 */
loop1:
	change = 0;
	for(r = firstr; r != R; r = r-&gt;link)
		r-&gt;active = 0;
	for(r = firstr; r != R; r = r-&gt;link)
		if(r-&gt;prog-&gt;as == ARET)
			prop(r, zbits, zbits);
loop11:
	/* pick up unreachable code */
	i = 0;
	for(r = firstr; r != R; r = r1) {
		r1 = r-&gt;link;
		if(r1 &amp;&amp; r1-&gt;active &amp;&amp; !r-&gt;active) {
			prop(r, zbits, zbits);
			i = 1;
		}
	}
	if(i)
		goto loop11;
	if(change)
		goto loop1;

	if(debug[&#39;R&#39;] &amp;&amp; debug[&#39;v&#39;])
		dumpit(&#34;pass3&#34;, firstr);

	/*
	 * pass 4
	 * iterate propagating register/variable synchrony
	 * 	forward until graph is complete
	 */
loop2:
	change = 0;
	for(r = firstr; r != R; r = r-&gt;link)
		r-&gt;active = 0;
	synch(firstr, zbits);
	if(change)
		goto loop2;

	if(debug[&#39;R&#39;] &amp;&amp; debug[&#39;v&#39;])
		dumpit(&#34;pass4&#34;, firstr);

	/*
	 * pass 5
	 * isolate regions
	 * calculate costs (paint1)
	 */
	r = firstr;
	if(r) {
		for(z=0; z&lt;BITS; z++)
			bit.b[z] = (r-&gt;refahead.b[z] | r-&gt;calahead.b[z]) &amp;
			  ~(externs.b[z] | params.b[z] | addrs.b[z] | consts.b[z]);
		if(bany(&amp;bit) &amp;&amp; !r-&gt;refset) {
			// should never happen - all variables are preset
			if(debug[&#39;w&#39;])
				print(&#34;%L: used and not set: %Q\n&#34;, r-&gt;prog-&gt;lineno, bit);
			r-&gt;refset = 1;
		}
	}
	for(r = firstr; r != R; r = r-&gt;link)
		r-&gt;act = zbits;
	rgp = region;
	nregion = 0;
	for(r = firstr; r != R; r = r-&gt;link) {
		for(z=0; z&lt;BITS; z++)
			bit.b[z] = r-&gt;set.b[z] &amp;
			  ~(r-&gt;refahead.b[z] | r-&gt;calahead.b[z] | addrs.b[z]);
		if(bany(&amp;bit) &amp;&amp; !r-&gt;refset) {
			if(debug[&#39;w&#39;])
				print(&#34;%L: set and not used: %Q\n&#34;, r-&gt;prog-&gt;lineno, bit);
			r-&gt;refset = 1;
			excise(r);
		}
		for(z=0; z&lt;BITS; z++)
			bit.b[z] = LOAD(r) &amp; ~(r-&gt;act.b[z] | addrs.b[z]);
		while(bany(&amp;bit)) {
			i = bnum(bit);
			rgp-&gt;enter = r;
			rgp-&gt;varno = i;
			change = 0;
			paint1(r, i);
			bit.b[i/32] &amp;= ~(1L&lt;&lt;(i%32));
			if(change &lt;= 0)
				continue;
			rgp-&gt;cost = change;
			nregion++;
			if(nregion &gt;= NRGN) {
				if(debug[&#39;R&#39;] &amp;&amp; debug[&#39;v&#39;])
					print(&#34;too many regions\n&#34;);
				goto brk;
			}
			rgp++;
		}
	}
brk:
	qsort(region, nregion, sizeof(region[0]), rcmp);

	/*
	 * pass 6
	 * determine used registers (paint2)
	 * replace code (paint3)
	 */
	rgp = region;
	for(i=0; i&lt;nregion; i++) {
		bit = blsh(rgp-&gt;varno);
		vreg = paint2(rgp-&gt;enter, rgp-&gt;varno);
		vreg = allreg(vreg, rgp);
		if(rgp-&gt;regno != 0)
			paint3(rgp-&gt;enter, rgp-&gt;varno, vreg, rgp-&gt;regno);
		rgp++;
	}

	if(debug[&#39;R&#39;] &amp;&amp; debug[&#39;v&#39;])
		dumpit(&#34;pass6&#34;, firstr);

	/*
	 * pass 7
	 * peep-hole on basic block
	 */
	if(!debug[&#39;R&#39;] || debug[&#39;P&#39;]) {
		peep();
	}

	/*
	 * eliminate nops
	 * free aux structures
	 */
	for(p=firstp; p!=P; p=p-&gt;link) {
		while(p-&gt;link != P &amp;&amp; p-&gt;link-&gt;as == ANOP)
			p-&gt;link = p-&gt;link-&gt;link;
		if(p-&gt;to.type == D_BRANCH)
			while(p-&gt;to.branch != P &amp;&amp; p-&gt;to.branch-&gt;as == ANOP)
				p-&gt;to.branch = p-&gt;to.branch-&gt;link;
	}

	if(r1 != R) {
		r1-&gt;link = freer;
		freer = firstr;
	}

	if(debug[&#39;R&#39;]) {
		if(ostats.ncvtreg ||
		   ostats.nspill ||
		   ostats.nreload ||
		   ostats.ndelmov ||
		   ostats.nvar ||
		   ostats.naddr ||
		   0)
			print(&#34;\nstats\n&#34;);

		if(ostats.ncvtreg)
			print(&#34;	%4ld cvtreg\n&#34;, ostats.ncvtreg);
		if(ostats.nspill)
			print(&#34;	%4ld spill\n&#34;, ostats.nspill);
		if(ostats.nreload)
			print(&#34;	%4ld reload\n&#34;, ostats.nreload);
		if(ostats.ndelmov)
			print(&#34;	%4ld delmov\n&#34;, ostats.ndelmov);
		if(ostats.nvar)
			print(&#34;	%4ld delmov\n&#34;, ostats.nvar);
		if(ostats.naddr)
			print(&#34;	%4ld delmov\n&#34;, ostats.naddr);

		memset(&amp;ostats, 0, sizeof(ostats));
	}
}

/*
 * add mov b,rn
 * just after r
 */
void
addmove(Reg *r, int bn, int rn, int f)
{
	Prog *p, *p1;
	Adr *a;
	Var *v;

	p1 = mal(sizeof(*p1));
	clearp(p1);
	p1-&gt;loc = 9999;

	p = r-&gt;prog;
	p1-&gt;link = p-&gt;link;
	p-&gt;link = p1;
	p1-&gt;lineno = p-&gt;lineno;

	v = var + bn;

	a = &amp;p1-&gt;to;
	a-&gt;sym = v-&gt;sym;
	a-&gt;offset = v-&gt;offset;
	a-&gt;etype = v-&gt;etype;
	a-&gt;type = v-&gt;name;
	a-&gt;gotype = v-&gt;gotype;

	// need to clean this up with wptr and
	// some of the defaults
	p1-&gt;as = AMOVL;
	switch(v-&gt;etype) {
	default:
		fatal(&#34;unknown type\n&#34;);
	case TINT8:
	case TUINT8:
	case TBOOL:
		p1-&gt;as = AMOVB;
		break;
	case TINT16:
	case TUINT16:
		p1-&gt;as = AMOVW;
		break;
	case TINT64:
	case TUINT64:
	case TUINTPTR:
	case TPTR64:
		p1-&gt;as = AMOVQ;
		break;
	case TFLOAT:
	case TFLOAT32:
		p1-&gt;as = AMOVSS;
		break;
	case TFLOAT64:
		p1-&gt;as = AMOVSD;
		break;
	case TINT:
	case TUINT:
	case TINT32:
	case TUINT32:
	case TPTR32:
		break;
	}

	p1-&gt;from.type = rn;
	if(!f) {
		p1-&gt;from = *a;
		*a = zprog.from;
		a-&gt;type = rn;
		if(v-&gt;etype == TUINT8)
			p1-&gt;as = AMOVB;
		if(v-&gt;etype == TUINT16)
			p1-&gt;as = AMOVW;
	}
	if(debug[&#39;R&#39;] &amp;&amp; debug[&#39;v&#39;])
		print(&#34;%P ===add=== %P\n&#34;, p, p1);
	ostats.nspill++;
}

uint32
doregbits(int r)
{
	uint32 b;

	b = 0;
	if(r &gt;= D_INDIR)
		r -= D_INDIR;
	if(r &gt;= D_AX &amp;&amp; r &lt;= D_R15)
		b |= RtoB(r);
	else
	if(r &gt;= D_AL &amp;&amp; r &lt;= D_R15B)
		b |= RtoB(r-D_AL+D_AX);
	else
	if(r &gt;= D_AH &amp;&amp; r &lt;= D_BH)
		b |= RtoB(r-D_AH+D_AX);
	else
	if(r &gt;= D_X0 &amp;&amp; r &lt;= D_X0+15)
		b |= FtoB(r);
	return b;
}

static int
overlap(Var *v, int o2, int w2)
{
	int o1, w1, t1, t2, z;
	Bits bit;

	o1 = v-&gt;offset;
	w1 = v-&gt;width;
	t1 = o1+w1;
	t2 = o2+w2;
	if(!(t1 &gt; o2 &amp;&amp; t2 &gt; o1))
		return 0;

	// set to max extent
	if(o2 &lt; o1)
		v-&gt;offset = o2;
	if(t1 &gt; t2)
		v-&gt;width = t1-v-&gt;offset;
	else
		v-&gt;width = t2-v-&gt;offset;

	// and dont registerize
	bit = blsh(v-var);
	for(z=0; z&lt;BITS; z++)
		addrs.b[z] |= bit.b[z];

	return 1;
}

Bits
mkvar(Reg *r, Adr *a)
{
	Var *v;
	int i, t, n, et, z, w, flag;
	int32 o;
	Bits bit;
	Sym *s;

	/*
	 * mark registers used
	 */
	t = a-&gt;type;
	if(r != R) {
		r-&gt;regu |= doregbits(t);
		r-&gt;regu |= doregbits(a-&gt;index);
	}

	switch(t) {
	default:
		goto none;
	case D_ADDR:
		a-&gt;type = a-&gt;index;
		bit = mkvar(r, a);
		for(z=0; z&lt;BITS; z++)
			addrs.b[z] |= bit.b[z];
		a-&gt;type = t;
		ostats.naddr++;
		goto none;
	case D_EXTERN:
	case D_STATIC:
	case D_PARAM:
	case D_AUTO:
		n = t;
		break;
	}
	s = a-&gt;sym;
	if(s == S)
		goto none;
	if(s-&gt;name[0] == &#39;.&#39;)
		goto none;
	et = a-&gt;etype;
	o = a-&gt;offset;
	w = a-&gt;width;
	v = var;

	flag = 0;
	for(i=0; i&lt;nvar; i++) {
		if(s == v-&gt;sym)
		if(n == v-&gt;name) {
			// if it is the same, use it
			if(v-&gt;etype == et)
			if(v-&gt;width == w)
			if(v-&gt;offset == o)
				goto out;

			// if it overlaps, set max
			// width and dont registerize
			if(overlap(v, o, w))
				flag = 1;
		}
		v++;
	}
	if(flag)
		goto none;

	switch(et) {
	case 0:
	case TFUNC:
	case TARRAY:
		goto none;
	}

	if(nvar &gt;= NVAR) {
		if(debug[&#39;w&#39;] &gt; 1 &amp;&amp; s)
			fatal(&#34;variable not optimized: %D&#34;, a);
		goto none;
	}
	i = nvar;
	nvar++;
	v = &amp;var[i];
	v-&gt;sym = s;
	v-&gt;offset = o;
	v-&gt;name = n;
	v-&gt;gotype = a-&gt;gotype;
	v-&gt;etype = et;
	v-&gt;width = w;
	if(debug[&#39;R&#39;])
		print(&#34;bit=%2d et=%2d w=%d %D\n&#34;, i, et, w, a);
	ostats.nvar++;

out:
	bit = blsh(i);

	// funny punning
	if(v-&gt;etype != et) {
		if(debug[&#39;R&#39;])
			print(&#34;pun et=%d/%d w=%d/%d o=%d/%d %D\n&#34;,
				v-&gt;etype, et,
				v-&gt;width, w,
				v-&gt;offset, o, a);
		for(z=0; z&lt;BITS; z++)
			addrs.b[z] |= bit.b[z];
		goto none;
	}

	if(n == D_EXTERN || n == D_STATIC)
		for(z=0; z&lt;BITS; z++)
			externs.b[z] |= bit.b[z];
	if(n == D_PARAM)
		for(z=0; z&lt;BITS; z++)
			params.b[z] |= bit.b[z];

	return bit;

none:
	return zbits;
}

void
prop(Reg *r, Bits ref, Bits cal)
{
	Reg *r1, *r2;
	int z;

	for(r1 = r; r1 != R; r1 = r1-&gt;p1) {
		for(z=0; z&lt;BITS; z++) {
			ref.b[z] |= r1-&gt;refahead.b[z];
			if(ref.b[z] != r1-&gt;refahead.b[z]) {
				r1-&gt;refahead.b[z] = ref.b[z];
				change++;
			}
			cal.b[z] |= r1-&gt;calahead.b[z];
			if(cal.b[z] != r1-&gt;calahead.b[z]) {
				r1-&gt;calahead.b[z] = cal.b[z];
				change++;
			}
		}
		switch(r1-&gt;prog-&gt;as) {
		case ACALL:
			if(noreturn(r1-&gt;prog))
				break;
			for(z=0; z&lt;BITS; z++) {
				cal.b[z] |= ref.b[z] | externs.b[z];
				ref.b[z] = 0;
			}
			break;

		case ATEXT:
			for(z=0; z&lt;BITS; z++) {
				cal.b[z] = 0;
				ref.b[z] = 0;
			}
			break;

		case ARET:
			for(z=0; z&lt;BITS; z++) {
				cal.b[z] = externs.b[z] | ovar.b[z];
				ref.b[z] = 0;
			}
			break;
		}
		for(z=0; z&lt;BITS; z++) {
			ref.b[z] = (ref.b[z] &amp; ~r1-&gt;set.b[z]) |
				r1-&gt;use1.b[z] | r1-&gt;use2.b[z];
			cal.b[z] &amp;= ~(r1-&gt;set.b[z] | r1-&gt;use1.b[z] | r1-&gt;use2.b[z]);
			r1-&gt;refbehind.b[z] = ref.b[z];
			r1-&gt;calbehind.b[z] = cal.b[z];
		}
		if(r1-&gt;active)
			break;
		r1-&gt;active = 1;
	}
	for(; r != r1; r = r-&gt;p1)
		for(r2 = r-&gt;p2; r2 != R; r2 = r2-&gt;p2link)
			prop(r2, r-&gt;refbehind, r-&gt;calbehind);
}

/*
 * find looping structure
 *
 * 1) find reverse postordering
 * 2) find approximate dominators,
 *	the actual dominators if the flow graph is reducible
 *	otherwise, dominators plus some other non-dominators.
 *	See Matthew S. Hecht and Jeffrey D. Ullman,
 *	&#34;Analysis of a Simple Algorithm for Global Data Flow Problems&#34;,
 *	Conf.  Record of ACM Symp. on Principles of Prog. Langs, Boston, Massachusetts,
 *	Oct. 1-3, 1973, pp.  207-217.
 * 3) find all nodes with a predecessor dominated by the current node.
 *	such a node is a loop head.
 *	recursively, all preds with a greater rpo number are in the loop
 */
int32
postorder(Reg *r, Reg **rpo2r, int32 n)
{
	Reg *r1;

	r-&gt;rpo = 1;
	r1 = r-&gt;s1;
	if(r1 &amp;&amp; !r1-&gt;rpo)
		n = postorder(r1, rpo2r, n);
	r1 = r-&gt;s2;
	if(r1 &amp;&amp; !r1-&gt;rpo)
		n = postorder(r1, rpo2r, n);
	rpo2r[n] = r;
	n++;
	return n;
}

int32
rpolca(int32 *idom, int32 rpo1, int32 rpo2)
{
	int32 t;

	if(rpo1 == -1)
		return rpo2;
	while(rpo1 != rpo2){
		if(rpo1 &gt; rpo2){
			t = rpo2;
			rpo2 = rpo1;
			rpo1 = t;
		}
		while(rpo1 &lt; rpo2){
			t = idom[rpo2];
			if(t &gt;= rpo2)
				fatal(&#34;bad idom&#34;);
			rpo2 = t;
		}
	}
	return rpo1;
}

int
doms(int32 *idom, int32 r, int32 s)
{
	while(s &gt; r)
		s = idom[s];
	return s == r;
}

int
loophead(int32 *idom, Reg *r)
{
	int32 src;

	src = r-&gt;rpo;
	if(r-&gt;p1 != R &amp;&amp; doms(idom, src, r-&gt;p1-&gt;rpo))
		return 1;
	for(r = r-&gt;p2; r != R; r = r-&gt;p2link)
		if(doms(idom, src, r-&gt;rpo))
			return 1;
	return 0;
}

void
loopmark(Reg **rpo2r, int32 head, Reg *r)
{
	if(r-&gt;rpo &lt; head || r-&gt;active == head)
		return;
	r-&gt;active = head;
	r-&gt;loop += LOOP;
	if(r-&gt;p1 != R)
		loopmark(rpo2r, head, r-&gt;p1);
	for(r = r-&gt;p2; r != R; r = r-&gt;p2link)
		loopmark(rpo2r, head, r);
}

void
loopit(Reg *r, int32 nr)
{
	Reg *r1;
	int32 i, d, me;

	if(nr &gt; maxnr) {
		rpo2r = mal(nr * sizeof(Reg*));
		idom = mal(nr * sizeof(int32));
		maxnr = nr;
	}

	d = postorder(r, rpo2r, 0);
	if(d &gt; nr)
		fatal(&#34;too many reg nodes %d %d&#34;, d, nr);
	nr = d;
	for(i = 0; i &lt; nr / 2; i++) {
		r1 = rpo2r[i];
		rpo2r[i] = rpo2r[nr - 1 - i];
		rpo2r[nr - 1 - i] = r1;
	}
	for(i = 0; i &lt; nr; i++)
		rpo2r[i]-&gt;rpo = i;

	idom[0] = 0;
	for(i = 0; i &lt; nr; i++) {
		r1 = rpo2r[i];
		me = r1-&gt;rpo;
		d = -1;
		if(r1-&gt;p1 != R &amp;&amp; r1-&gt;p1-&gt;rpo &lt; me)
			d = r1-&gt;p1-&gt;rpo;
		for(r1 = r1-&gt;p2; r1 != nil; r1 = r1-&gt;p2link)
			if(r1-&gt;rpo &lt; me)
				d = rpolca(idom, d, r1-&gt;rpo);
		idom[i] = d;
	}

	for(i = 0; i &lt; nr; i++) {
		r1 = rpo2r[i];
		r1-&gt;loop++;
		if(r1-&gt;p2 != R &amp;&amp; loophead(idom, r1))
			loopmark(rpo2r, i, r1);
	}
}

void
synch(Reg *r, Bits dif)
{
	Reg *r1;
	int z;

	for(r1 = r; r1 != R; r1 = r1-&gt;s1) {
		for(z=0; z&lt;BITS; z++) {
			dif.b[z] = (dif.b[z] &amp;
				~(~r1-&gt;refbehind.b[z] &amp; r1-&gt;refahead.b[z])) |
					r1-&gt;set.b[z] | r1-&gt;regdiff.b[z];
			if(dif.b[z] != r1-&gt;regdiff.b[z]) {
				r1-&gt;regdiff.b[z] = dif.b[z];
				change++;
			}
		}
		if(r1-&gt;active)
			break;
		r1-&gt;active = 1;
		for(z=0; z&lt;BITS; z++)
			dif.b[z] &amp;= ~(~r1-&gt;calbehind.b[z] &amp; r1-&gt;calahead.b[z]);
		if(r1-&gt;s2 != R)
			synch(r1-&gt;s2, dif);
	}
}

uint32
allreg(uint32 b, Rgn *r)
{
	Var *v;
	int i;

	v = var + r-&gt;varno;
	r-&gt;regno = 0;
	switch(v-&gt;etype) {

	default:
		fatal(&#34;unknown etype %d/%E&#34;, bitno(b), v-&gt;etype);
		break;

	case TINT8:
	case TUINT8:
	case TINT16:
	case TUINT16:
	case TINT32:
	case TUINT32:
	case TINT64:
	case TUINT64:
	case TINT:
	case TUINT:
	case TUINTPTR:
	case TBOOL:
	case TPTR32:
	case TPTR64:
		i = BtoR(~b);
		if(i &amp;&amp; r-&gt;cost &gt; 0) {
			r-&gt;regno = i;
			return RtoB(i);
		}
		break;

	case TFLOAT32:
	case TFLOAT64:
	case TFLOAT:
		i = BtoF(~b);
		if(i &amp;&amp; r-&gt;cost &gt; 0) {
			r-&gt;regno = i;
			return FtoB(i);
		}
		break;
	}
	return 0;
}

void
paint1(Reg *r, int bn)
{
	Reg *r1;
	Prog *p;
	int z;
	uint32 bb;

	z = bn/32;
	bb = 1L&lt;&lt;(bn%32);
	if(r-&gt;act.b[z] &amp; bb)
		return;
	for(;;) {
		if(!(r-&gt;refbehind.b[z] &amp; bb))
			break;
		r1 = r-&gt;p1;
		if(r1 == R)
			break;
		if(!(r1-&gt;refahead.b[z] &amp; bb))
			break;
		if(r1-&gt;act.b[z] &amp; bb)
			break;
		r = r1;
	}

	if(LOAD(r) &amp; ~(r-&gt;set.b[z]&amp;~(r-&gt;use1.b[z]|r-&gt;use2.b[z])) &amp; bb) {
		change -= CLOAD * r-&gt;loop;
	}
	for(;;) {
		r-&gt;act.b[z] |= bb;
		p = r-&gt;prog;

		if(r-&gt;use1.b[z] &amp; bb) {
			change += CREF * r-&gt;loop;
		}

		if((r-&gt;use2.b[z]|r-&gt;set.b[z]) &amp; bb) {
			change += CREF * r-&gt;loop;
		}

		if(STORE(r) &amp; r-&gt;regdiff.b[z] &amp; bb) {
			change -= CLOAD * r-&gt;loop;
		}

		if(r-&gt;refbehind.b[z] &amp; bb)
			for(r1 = r-&gt;p2; r1 != R; r1 = r1-&gt;p2link)
				if(r1-&gt;refahead.b[z] &amp; bb)
					paint1(r1, bn);

		if(!(r-&gt;refahead.b[z] &amp; bb))
			break;
		r1 = r-&gt;s2;
		if(r1 != R)
			if(r1-&gt;refbehind.b[z] &amp; bb)
				paint1(r1, bn);
		r = r-&gt;s1;
		if(r == R)
			break;
		if(r-&gt;act.b[z] &amp; bb)
			break;
		if(!(r-&gt;refbehind.b[z] &amp; bb))
			break;
	}
}

uint32
regset(Reg *r, uint32 bb)
{
	uint32 b, set;
	Adr v;
	int c;

	set = 0;
	v = zprog.from;
	while(b = bb &amp; ~(bb-1)) {
		v.type = b &amp; 0xFFFF? BtoR(b): BtoF(b);
		if(v.type == 0)
			fatal(&#34;zero v.type for %#lux&#34;, b);
		c = copyu(r-&gt;prog, &amp;v, A);
		if(c == 3)
			set |= b;
		bb &amp;= ~b;
	}
	return set;
}

uint32
reguse(Reg *r, uint32 bb)
{
	uint32 b, set;
	Adr v;
	int c;

	set = 0;
	v = zprog.from;
	while(b = bb &amp; ~(bb-1)) {
		v.type = b &amp; 0xFFFF? BtoR(b): BtoF(b);
		c = copyu(r-&gt;prog, &amp;v, A);
		if(c == 1 || c == 2 || c == 4)
			set |= b;
		bb &amp;= ~b;
	}
	return set;
}

uint32
paint2(Reg *r, int bn)
{
	Reg *r1;
	int z;
	uint32 bb, vreg, x;

	z = bn/32;
	bb = 1L &lt;&lt; (bn%32);
	vreg = regbits;
	if(!(r-&gt;act.b[z] &amp; bb))
		return vreg;
	for(;;) {
		if(!(r-&gt;refbehind.b[z] &amp; bb))
			break;
		r1 = r-&gt;p1;
		if(r1 == R)
			break;
		if(!(r1-&gt;refahead.b[z] &amp; bb))
			break;
		if(!(r1-&gt;act.b[z] &amp; bb))
			break;
		r = r1;
	}
	for(;;) {
		r-&gt;act.b[z] &amp;= ~bb;

		vreg |= r-&gt;regu;

		if(r-&gt;refbehind.b[z] &amp; bb)
			for(r1 = r-&gt;p2; r1 != R; r1 = r1-&gt;p2link)
				if(r1-&gt;refahead.b[z] &amp; bb)
					vreg |= paint2(r1, bn);

		if(!(r-&gt;refahead.b[z] &amp; bb))
			break;
		r1 = r-&gt;s2;
		if(r1 != R)
			if(r1-&gt;refbehind.b[z] &amp; bb)
				vreg |= paint2(r1, bn);
		r = r-&gt;s1;
		if(r == R)
			break;
		if(!(r-&gt;act.b[z] &amp; bb))
			break;
		if(!(r-&gt;refbehind.b[z] &amp; bb))
			break;
	}

	bb = vreg;
	for(; r; r=r-&gt;s1) {
		x = r-&gt;regu &amp; ~bb;
		if(x) {
			vreg |= reguse(r, x);
			bb |= regset(r, x);
		}
	}
	return vreg;
}

void
paint3(Reg *r, int bn, int32 rb, int rn)
{
	Reg *r1;
	Prog *p;
	int z;
	uint32 bb;

	z = bn/32;
	bb = 1L &lt;&lt; (bn%32);
	if(r-&gt;act.b[z] &amp; bb)
		return;
	for(;;) {
		if(!(r-&gt;refbehind.b[z] &amp; bb))
			break;
		r1 = r-&gt;p1;
		if(r1 == R)
			break;
		if(!(r1-&gt;refahead.b[z] &amp; bb))
			break;
		if(r1-&gt;act.b[z] &amp; bb)
			break;
		r = r1;
	}

	if(LOAD(r) &amp; ~(r-&gt;set.b[z] &amp; ~(r-&gt;use1.b[z]|r-&gt;use2.b[z])) &amp; bb)
		addmove(r, bn, rn, 0);
	for(;;) {
		r-&gt;act.b[z] |= bb;
		p = r-&gt;prog;

		if(r-&gt;use1.b[z] &amp; bb) {
			if(debug[&#39;R&#39;] &amp;&amp; debug[&#39;v&#39;])
				print(&#34;%P&#34;, p);
			addreg(&amp;p-&gt;from, rn);
			if(debug[&#39;R&#39;] &amp;&amp; debug[&#39;v&#39;])
				print(&#34; ===change== %P\n&#34;, p);
		}
		if((r-&gt;use2.b[z]|r-&gt;set.b[z]) &amp; bb) {
			if(debug[&#39;R&#39;] &amp;&amp; debug[&#39;v&#39;])
				print(&#34;%P&#34;, p);
			addreg(&amp;p-&gt;to, rn);
			if(debug[&#39;R&#39;] &amp;&amp; debug[&#39;v&#39;])
				print(&#34; ===change== %P\n&#34;, p);
		}

		if(STORE(r) &amp; r-&gt;regdiff.b[z] &amp; bb)
			addmove(r, bn, rn, 1);
		r-&gt;regu |= rb;

		if(r-&gt;refbehind.b[z] &amp; bb)
			for(r1 = r-&gt;p2; r1 != R; r1 = r1-&gt;p2link)
				if(r1-&gt;refahead.b[z] &amp; bb)
					paint3(r1, bn, rb, rn);

		if(!(r-&gt;refahead.b[z] &amp; bb))
			break;
		r1 = r-&gt;s2;
		if(r1 != R)
			if(r1-&gt;refbehind.b[z] &amp; bb)
				paint3(r1, bn, rb, rn);
		r = r-&gt;s1;
		if(r == R)
			break;
		if(r-&gt;act.b[z] &amp; bb)
			break;
		if(!(r-&gt;refbehind.b[z] &amp; bb))
			break;
	}
}

void
addreg(Adr *a, int rn)
{

	a-&gt;sym = 0;
	a-&gt;offset = 0;
	a-&gt;type = rn;

	ostats.ncvtreg++;
}

int32
RtoB(int r)
{

	if(r &lt; D_AX || r &gt; D_R15)
		return 0;
	return 1L &lt;&lt; (r-D_AX);
}

int
BtoR(int32 b)
{
	b &amp;= 0x3fffL;		// no R14 or R15
	if(b == 0)
		return 0;
	return bitno(b) + D_AX;
}

/*
 *	bit	reg
 *	16	X5 (FREGMIN)
 *	...
 *	26	X15 (FREGEXT)
 */
int32
FtoB(int f)
{
	if(f &lt; FREGMIN || f &gt; FREGEXT)
		return 0;
	return 1L &lt;&lt; (f - FREGMIN + 16);
}

int
BtoF(int32 b)
{

	b &amp;= 0xFF0000L;
	if(b == 0)
		return 0;
	return bitno(b) - 16 + FREGMIN;
}

void
dumpone(Reg *r)
{
	int z;
	Bits bit;

	print(&#34;%ld:%P&#34;, r-&gt;loop, r-&gt;prog);
	for(z=0; z&lt;BITS; z++)
		bit.b[z] =
			r-&gt;set.b[z] |
			r-&gt;use1.b[z] |
			r-&gt;use2.b[z] |
			r-&gt;refbehind.b[z] |
			r-&gt;refahead.b[z] |
			r-&gt;calbehind.b[z] |
			r-&gt;calahead.b[z] |
			r-&gt;regdiff.b[z] |
			r-&gt;act.b[z] |
				0;
	if(bany(&amp;bit)) {
		print(&#34;\t&#34;);
		if(bany(&amp;r-&gt;set))
			print(&#34; s:%Q&#34;, r-&gt;set);
		if(bany(&amp;r-&gt;use1))
			print(&#34; u1:%Q&#34;, r-&gt;use1);
		if(bany(&amp;r-&gt;use2))
			print(&#34; u2:%Q&#34;, r-&gt;use2);
		if(bany(&amp;r-&gt;refbehind))
			print(&#34; rb:%Q &#34;, r-&gt;refbehind);
		if(bany(&amp;r-&gt;refahead))
			print(&#34; ra:%Q &#34;, r-&gt;refahead);
		if(bany(&amp;r-&gt;calbehind))
			print(&#34;cb:%Q &#34;, r-&gt;calbehind);
		if(bany(&amp;r-&gt;calahead))
			print(&#34; ca:%Q &#34;, r-&gt;calahead);
		if(bany(&amp;r-&gt;regdiff))
			print(&#34; d:%Q &#34;, r-&gt;regdiff);
		if(bany(&amp;r-&gt;act))
			print(&#34; a:%Q &#34;, r-&gt;act);
	}
	print(&#34;\n&#34;);
}

void
dumpit(char *str, Reg *r0)
{
	Reg *r, *r1;

	print(&#34;\n%s\n&#34;, str);
	for(r = r0; r != R; r = r-&gt;link) {
		dumpone(r);
		r1 = r-&gt;p2;
		if(r1 != R) {
			print(&#34;	pred:&#34;);
			for(; r1 != R; r1 = r1-&gt;p2link)
				print(&#34; %.4lud&#34;, r1-&gt;prog-&gt;loc);
			print(&#34;\n&#34;);
		}
//		r1 = r-&gt;s1;
//		if(r1 != R) {
//			print(&#34;	succ:&#34;);
//			for(; r1 != R; r1 = r1-&gt;s1)
//				print(&#34; %.4lud&#34;, r1-&gt;prog-&gt;loc);
//			print(&#34;\n&#34;);
//		}
	}
}

static Sym*	symlist[10];

int
noreturn(Prog *p)
{
	Sym *s;
	int i;

	if(symlist[0] == S) {
		symlist[0] = pkglookup(&#34;throwindex&#34;, &#34;sys&#34;);
		symlist[1] = pkglookup(&#34;throwslice&#34;, &#34;sys&#34;);
		symlist[2] = pkglookup(&#34;throwinit&#34;, &#34;sys&#34;);
		symlist[3] = pkglookup(&#34;panicl&#34;, &#34;sys&#34;);
	}

	s = p-&gt;to.sym;
	if(s == S)
		return 0;
	for(i=0; symlist[i]!=S; i++)
		if(s == symlist[i])
			return 1;
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
