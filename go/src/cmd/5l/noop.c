<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5l/noop.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/5l/noop.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5l/noop.c
// http://code.google.com/p/inferno-os/source/browse/utils/5l/noop.c
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

// see ../../runtime/proc.c:/StackGuard
enum
{
	StackBig = 4096,
};

static	Sym*	sym_div;
static	Sym*	sym_divu;
static	Sym*	sym_mod;
static	Sym*	sym_modu;

static void setdiv(int);

static Prog *
movrr(Prog *q, int rs, int rd, Prog *p)
{
	if(q == nil)
		q = prg();
	q-&gt;as = AMOVW;
	q-&gt;line = p-&gt;line;
	q-&gt;from.type = D_REG;
	q-&gt;from.reg = rs;
	q-&gt;to.type = D_REG;
	q-&gt;to.reg = rd;
	q-&gt;link = p-&gt;link;
	return q;
}

static Prog *
fnret(Prog *q, int rs, int foreign, Prog *p)
{
	q = movrr(q, rs, REGPC, p);
	if(foreign){	// BX rs
		q-&gt;as = ABXRET;
		q-&gt;from.type = D_NONE;
		q-&gt;from.reg = NREG;
		q-&gt;to.reg = rs;
	}
	return q;
}

static Prog *
aword(int32 w, Prog *p)
{
	Prog *q;

	q = prg();
	q-&gt;as = AWORD;
	q-&gt;line = p-&gt;line;
	q-&gt;from.type = D_NONE;
	q-&gt;reg = NREG;
	q-&gt;to.type = D_CONST;
	q-&gt;to.offset = w;
	q-&gt;link = p-&gt;link;
	p-&gt;link = q;
	return q;
}

static Prog *
adword(int32 w1, int32 w2, Prog *p)
{
	Prog *q;

	q = prg();
	q-&gt;as = ADWORD;
	q-&gt;line = p-&gt;line;
	q-&gt;from.type = D_CONST;
	q-&gt;from.offset = w1;
	q-&gt;reg = NREG;
	q-&gt;to.type = D_CONST;
	q-&gt;to.offset = w2;
	q-&gt;link = p-&gt;link;
	p-&gt;link = q;
	return q;
}

void
noops(void)
{
	Prog *p, *q, *q1, *q2;
	int o, curframe, curbecome, maxbecome, foreign;
	Prog *pmorestack;
	Sym *symmorestack;

	/*
	 * find leaf subroutines
	 * become sizes
	 * frame sizes
	 * strip NOPs
	 * expand RET
	 * expand BECOME pseudo
	 */

	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f noops\n&#34;, cputime());
	Bflush(&amp;bso);

	pmorestack = P;
	symmorestack = lookup(&#34;runtime·morestack&#34;, 0);

	if(symmorestack-&gt;type == STEXT)
	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT) {
			if(p-&gt;from.sym == symmorestack) {
				pmorestack = p;
				p-&gt;reg |= NOSPLIT;
				break;
			}
		}
	}
	// TODO(kaib): make lack of morestack an error
//	if(pmorestack == P)
//		diag(&#34;runtime·morestack not defined&#34;);

	curframe = 0;
	curbecome = 0;
	maxbecome = 0;
	curtext = 0;

	q = P;
	for(p = firstp; p != P; p = p-&gt;link) {
		setarch(p);

		/* find out how much arg space is used in this TEXT */
		if(p-&gt;to.type == D_OREG &amp;&amp; p-&gt;to.reg == REGSP)
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

			p-&gt;mark |= LEAF;
			curtext = p;
			break;

		case ARET:
			/* special form of RET is BECOME */
			if(p-&gt;from.type == D_CONST)
				if(p-&gt;from.offset &gt; curbecome)
					curbecome = p-&gt;from.offset;
			break;

		case ADIV:
		case ADIVU:
		case AMOD:
		case AMODU:
			q = p;
			if(prog_div == P)
				initdiv();
			if(curtext != P)
				curtext-&gt;mark &amp;= ~LEAF;
			setdiv(p-&gt;as);
			continue;

		case ANOP:
			q1 = p-&gt;link;
			q-&gt;link = q1;		/* q is non-nop */
			q1-&gt;mark |= p-&gt;mark;
			continue;

		case ABL:
		case ABX:
			if(curtext != P)
				curtext-&gt;mark &amp;= ~LEAF;

		case ABCASE:
		case AB:

		case ABEQ:
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

			q1 = p-&gt;cond;
			if(q1 != P) {
				while(q1-&gt;as == ANOP) {
					q1 = q1-&gt;link;
					p-&gt;cond = q1;
				}
			}
			break;
		}
		q = p;
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
		setarch(p);
		switch(p-&gt;as) {
		case ATEXT:
			curtext = p;
			break;
		case ABL:
		// case ABX:
			if(curtext != P &amp;&amp; curtext-&gt;from.sym != S &amp;&amp; curtext-&gt;to.offset &gt;= 0) {
				o = maxbecome - curtext-&gt;from.sym-&gt;frame;
				if(o &lt;= 0)
					break;
				/* calling a become or calling a variable */
				if(p-&gt;to.sym == S || p-&gt;to.sym-&gt;become) {
					curtext-&gt;to.offset += o;
					if(debug[&#39;b&#39;]) {
						curp = p;
						print(&#34;%D calling %D increase %d\n&#34;,
							&amp;curtext-&gt;from, &amp;p-&gt;to, o);
					}
				}
			}
			break;
		}
	}

	for(p = firstp; p != P; p = p-&gt;link) {
		setarch(p);
		o = p-&gt;as;
		switch(o) {
		case ATEXT:
			curtext = p;
			autosize = p-&gt;to.offset + 4;
			if(autosize &lt;= 4)
			if(curtext-&gt;mark &amp; LEAF) {
				p-&gt;to.offset = -4;
				autosize = 0;
			}

			if(!autosize &amp;&amp; !(curtext-&gt;mark &amp; LEAF)) {
				if(debug[&#39;v&#39;])
					Bprint(&amp;bso, &#34;save suppressed in: %s\n&#34;,
						curtext-&gt;from.sym-&gt;name);
				Bflush(&amp;bso);
				curtext-&gt;mark |= LEAF;
			}
#ifdef CALLEEBX
			if(p-&gt;from.sym-&gt;foreign){
				if(thumb)
					// don&#39;t allow literal pool to seperate these
					p = adword(0xe28f7001, 0xe12fff17, p); // arm add 1, pc, r7 and bx r7
					// p = aword(0xe12fff17, aword(0xe28f7001, p)); // arm add 1, pc, r7 and bx r7
				else
					p = aword(0x4778, p);	// thumb bx pc and 2 bytes padding
			}
#endif
			if(curtext-&gt;mark &amp; LEAF) {
				if(curtext-&gt;from.sym)
					curtext-&gt;from.sym-&gt;type = SLEAF;
				if(!autosize)
					break;
			}

			if(thumb){
				if(!(p-&gt;reg &amp; NOSPLIT))
					diag(&#34;stack splitting not supported in thumb&#34;);
				if(!(curtext-&gt;mark &amp; LEAF)){
					q = movrr(nil, REGLINK, REGTMPT-1, p);
					p-&gt;link = q;
					q1 = prg();
					q1-&gt;as = AMOVW;
					q1-&gt;line = p-&gt;line;
					q1-&gt;from.type = D_REG;
					q1-&gt;from.reg = REGTMPT-1;
					q1-&gt;to.type = D_OREG;
					q1-&gt;to.name = D_NONE;
					q1-&gt;to.reg = REGSP;
					q1-&gt;to.offset = 0;
					q1-&gt;link = q-&gt;link;
					q-&gt;link = q1;
				}
				if(autosize){
					q2 = prg();
					q2-&gt;as = ASUB;
					q2-&gt;line = p-&gt;line;
					q2-&gt;from.type = D_CONST;
					q2-&gt;from.offset = autosize;
					q2-&gt;to.type = D_REG;
					q2-&gt;to.reg = REGSP;
					q2-&gt;link = p-&gt;link;
					p-&gt;link = q2;
				}
				break;
			}

			if(p-&gt;reg &amp; NOSPLIT) {
				q1 = prg();
				q1-&gt;as = AMOVW;
				q1-&gt;scond |= C_WBIT;
				q1-&gt;line = p-&gt;line;
				q1-&gt;from.type = D_REG;
				q1-&gt;from.reg = REGLINK;
				q1-&gt;to.type = D_OREG;
				q1-&gt;to.offset = -autosize;
				q1-&gt;to.reg = REGSP;
				q1-&gt;link = p-&gt;link;
				p-&gt;link = q1;
			} else if (autosize &lt; StackBig) {
				// split stack check for small functions
				// MOVW			g_stackguard(g), R1
				// CMP			R1, $-autosize(SP)
				// MOVW.LO		$autosize, R1
				// MOVW.LO		$args, R2
				// MOVW.LO		R14, R3
				// BL.LO			runtime·morestack(SB) // modifies LR
				// MOVW.W		R14,$-autosize(SP)

				// TODO(kaib): add more trampolines
				// TODO(kaib): put stackguard in register
				// TODO(kaib): add support for -K and underflow detection

				// MOVW			g_stackguard(g), R1
				p = appendp(p);
				p-&gt;as = AMOVW;
				p-&gt;from.type = D_OREG;
				p-&gt;from.reg = REGG;
				p-&gt;to.type = D_REG;
				p-&gt;to.reg = 1;

				// CMP			R1, $-autosize(SP)
				p = appendp(p);
				p-&gt;as = ACMP;
				p-&gt;from.type = D_REG;
				p-&gt;from.reg = 1;
				p-&gt;from.offset = -autosize;
				p-&gt;reg = REGSP;

				// MOVW.LO		$autosize, R1
				p = appendp(p);
				p-&gt;as = AMOVW;
				p-&gt;scond = C_SCOND_LO;
				p-&gt;from.type = D_CONST;
				p-&gt;from.offset = 0;
				p-&gt;to.type = D_REG;
				p-&gt;to.reg = 1;

				// MOVW.LO		$args +4, R2
				// also need to store the extra 4 bytes.
				p = appendp(p);
				p-&gt;as = AMOVW;
				p-&gt;scond = C_SCOND_LO;
				p-&gt;from.type = D_CONST;
				p-&gt;from.offset = (curtext-&gt;to.offset2 &amp; ~7) + 4;
				p-&gt;to.type = D_REG;
				p-&gt;to.reg = 2;

				// MOVW.LO	R14, R3
				p = appendp(p);
				p-&gt;as = AMOVW;
				p-&gt;scond = C_SCOND_LO;
				p-&gt;from.type = D_REG;
				p-&gt;from.reg = REGLINK;
				p-&gt;to.type = D_REG;
				p-&gt;to.reg = 3;

				// BL.LO		runtime·morestack(SB) // modifies LR
				p = appendp(p);
				p-&gt;as = ABL;
				p-&gt;scond = C_SCOND_LO;
 				p-&gt;to.type = D_BRANCH;
				p-&gt;to.sym = symmorestack;
				p-&gt;cond = pmorestack;

				// MOVW.W		R14,$-autosize(SP)
				p = appendp(p);
				p-&gt;as = AMOVW;
 				p-&gt;scond |= C_WBIT;
				p-&gt;from.type = D_REG;
				p-&gt;from.reg = REGLINK;
				p-&gt;to.type = D_OREG;
				p-&gt;to.offset = -autosize;
				p-&gt;to.reg = REGSP;
			} else { // &gt; StackBig
				// MOVW		$autosize, R1
				// MOVW		$args, R2
				// MOVW		R14, R3
				// BL			runtime·morestack(SB) // modifies LR
				// MOVW.W		R14,$-autosize(SP)

				// MOVW		$autosize, R1
				p = appendp(p);
				p-&gt;as = AMOVW;
				p-&gt;from.type = D_CONST;
				p-&gt;from.offset = autosize;
				p-&gt;to.type = D_REG;
				p-&gt;to.reg = 1;

				// MOVW		$args +4, R2
				// also need to store the extra 4 bytes.
				p = appendp(p);
				p-&gt;as = AMOVW;
				p-&gt;from.type = D_CONST;
				p-&gt;from.offset = (curtext-&gt;to.offset2 &amp; ~7) + 4;
				p-&gt;to.type = D_REG;
				p-&gt;to.reg = 2;

				// MOVW	R14, R3
				p = appendp(p);
				p-&gt;as = AMOVW;
				p-&gt;from.type = D_REG;
				p-&gt;from.reg = REGLINK;
				p-&gt;to.type = D_REG;
				p-&gt;to.reg = 3;

				// BL		runtime·morestack(SB) // modifies LR
				p = appendp(p);
				p-&gt;as = ABL;
 				p-&gt;to.type = D_BRANCH;
				p-&gt;to.sym = symmorestack;
				p-&gt;cond = pmorestack;

				// MOVW.W		R14,$-autosize(SP)
				p = appendp(p);
				p-&gt;as = AMOVW;
 				p-&gt;scond |= C_WBIT;
				p-&gt;from.type = D_REG;
				p-&gt;from.reg = REGLINK;
				p-&gt;to.type = D_OREG;
				p-&gt;to.offset = -autosize;
				p-&gt;to.reg = REGSP;
			}
			break;

		case ARET:
			nocache(p);
			foreign = seenthumb &amp;&amp; curtext-&gt;from.sym != S &amp;&amp; (curtext-&gt;from.sym-&gt;foreign || curtext-&gt;from.sym-&gt;fnptr);
// print(&#34;%s %d %d\n&#34;, curtext-&gt;from.sym-&gt;name, curtext-&gt;from.sym-&gt;foreign, curtext-&gt;from.sym-&gt;fnptr);
			if(p-&gt;from.type == D_CONST)
				goto become;
			if(curtext-&gt;mark &amp; LEAF) {
				if(!autosize) {
					if(thumb){
						p = fnret(p, REGLINK, foreign, p);
						break;
					}
// if(foreign) print(&#34;ABXRET 1 %s\n&#34;, curtext-&gt;from.sym-&gt;name);
					p-&gt;as = foreign ? ABXRET : AB;
					p-&gt;from = zprg.from;
					p-&gt;to.type = D_OREG;
					p-&gt;to.offset = 0;
					p-&gt;to.reg = REGLINK;
					break;
				}
			}
			if(thumb){
				if(curtext-&gt;mark &amp; LEAF){
					if(autosize){
						p-&gt;as = AADD;
						p-&gt;from.type = D_CONST;
						p-&gt;from.offset = autosize;
						p-&gt;to.type = D_REG;
						p-&gt;to.reg = REGSP;
						q = nil;
					}
					else
						q = p;
					q = fnret(q, REGLINK, foreign, p);
					if(q != p)
						p-&gt;link = q;
				}
				else{
					p-&gt;as = AMOVW;
					p-&gt;from.type = D_OREG;
					p-&gt;from.name = D_NONE;
					p-&gt;from.reg = REGSP;
					p-&gt;from.offset = 0;
					p-&gt;to.type = D_REG;
					p-&gt;to.reg = REGTMPT-1;
					if(autosize){
						q = prg();
						q-&gt;as = AADD;
						q-&gt;from.type = D_CONST;
						q-&gt;from.offset = autosize;
						q-&gt;to.type = D_REG;
						q-&gt;to.reg = REGSP;
						q-&gt;link = p-&gt;link;
						p-&gt;link = 	q;
					}
					else
						q = p;
					q1 = fnret(nil, REGTMPT-1, foreign, p);
					q1-&gt;link = q-&gt;link;
					q-&gt;link = q1;
				}
				break;
			}
			if(foreign) {
// if(foreign) print(&#34;ABXRET 3 %s\n&#34;, curtext-&gt;from.sym-&gt;name);
#define	R	1
				p-&gt;as = AMOVW;
				p-&gt;from.type = D_OREG;
				p-&gt;from.name = D_NONE;
				p-&gt;from.reg = REGSP;
				p-&gt;from.offset = 0;
				p-&gt;to.type = D_REG;
				p-&gt;to.reg = R;
				q = prg();
				q-&gt;as = AADD;
				q-&gt;scond = p-&gt;scond;
				q-&gt;line = p-&gt;line;
				q-&gt;from.type = D_CONST;
				q-&gt;from.offset = autosize;
				q-&gt;to.type = D_REG;
				q-&gt;to.reg = REGSP;
				q-&gt;link = p-&gt;link;
				p-&gt;link = q;
				q1 = prg();
				q1-&gt;as = ABXRET;
				q1-&gt;scond = p-&gt;scond;
				q1-&gt;line = p-&gt;line;
				q1-&gt;to.type = D_OREG;
				q1-&gt;to.offset = 0;
				q1-&gt;to.reg = R;
				q1-&gt;link = q-&gt;link;
				q-&gt;link = q1;
#undef	R
			}
			else {
				p-&gt;as = AMOVW;
				p-&gt;scond |= C_PBIT;
				p-&gt;from.type = D_OREG;
				p-&gt;from.offset = autosize;
				p-&gt;from.reg = REGSP;
				p-&gt;to.type = D_REG;
				p-&gt;to.reg = REGPC;
			}
			break;

		become:
			if(foreign){
				diag(&#34;foreign become - help&#34;);
				break;
			}
			if(thumb){
				diag(&#34;thumb become - help&#34;);
				break;
			}
			print(&#34;arm become\n&#34;);
			if(curtext-&gt;mark &amp; LEAF) {

				if(!autosize) {
					p-&gt;as = AB;
					p-&gt;from = zprg.from;
					break;
				}
			}
			q = prg();
			q-&gt;scond = p-&gt;scond;
			q-&gt;line = p-&gt;line;
			q-&gt;as = AB;
			q-&gt;from = zprg.from;
			q-&gt;to = p-&gt;to;
			q-&gt;cond = p-&gt;cond;
			q-&gt;link = p-&gt;link;
			p-&gt;link = q;
			if(thumb){
				q1 = prg();
				q1-&gt;line = p-&gt;line;
				q1-&gt;as = AADD;
				q1-&gt;from.type = D_CONST;
				q1-&gt;from.offset = autosize;
				q1-&gt;to.type = D_REG;
				q1-&gt;to.reg = REGSP;
				p-&gt;as = AMOVW;
				p-&gt;line = p-&gt;line;
				p-&gt;from.type = D_OREG;
				p-&gt;from.name = D_NONE;
				p-&gt;from.reg = REGSP;
				p-&gt;from.offset = 0;
				p-&gt;to.type = D_REG;
				p-&gt;to.reg = REGTMPT-1;
				q1-&gt;link = q;
				p-&gt;link = q1;
				q2 = movrr(nil, REGTMPT-1, REGLINK, p);
				q2-&gt;link = q;
				q1-&gt;link = q2;
				break;
			}
			p-&gt;as = AMOVW;
			p-&gt;scond |= C_PBIT;
			p-&gt;from = zprg.from;
			p-&gt;from.type = D_OREG;
			p-&gt;from.offset = autosize;
			p-&gt;from.reg = REGSP;
			p-&gt;to = zprg.to;
			p-&gt;to.type = D_REG;
			p-&gt;to.reg = REGLINK;

			break;

		case ADIV:
		case ADIVU:
		case AMOD:
		case AMODU:
			if(debug[&#39;M&#39;])
				break;
			if(p-&gt;from.type != D_REG)
				break;
			if(p-&gt;to.type != D_REG)
				break;
			q1 = p;

			/* MOV a,4(SP) */
			q = prg();
			q-&gt;link = p-&gt;link;
			p-&gt;link = q;
			p = q;

			p-&gt;as = AMOVW;
			p-&gt;line = q1-&gt;line;
			p-&gt;from.type = D_REG;
			p-&gt;from.reg = q1-&gt;from.reg;
			p-&gt;to.type = D_OREG;
			p-&gt;to.reg = REGSP;
			p-&gt;to.offset = 4;

			/* MOV b,REGTMP */
			q = prg();
			q-&gt;link = p-&gt;link;
			p-&gt;link = q;
			p = q;

			p-&gt;as = AMOVW;
			p-&gt;line = q1-&gt;line;
			p-&gt;from.type = D_REG;
			p-&gt;from.reg = q1-&gt;reg;
			if(q1-&gt;reg == NREG)
				p-&gt;from.reg = q1-&gt;to.reg;
			p-&gt;to.type = D_REG;
			p-&gt;to.reg = prog_div != UP &amp;&amp; prog_div-&gt;from.sym-&gt;thumb ? REGTMPT : REGTMP;
			p-&gt;to.offset = 0;

			/* CALL appropriate */
			q = prg();
			q-&gt;link = p-&gt;link;
			p-&gt;link = q;
			p = q;

#ifdef CALLEEBX
			p-&gt;as = ABL;
#else
			if(prog_div != UP &amp;&amp; prog_div-&gt;from.sym-&gt;thumb)
				p-&gt;as = thumb ? ABL : ABX;
			else
				p-&gt;as = thumb ? ABX : ABL;
#endif
			p-&gt;line = q1-&gt;line;
			p-&gt;to.type = D_BRANCH;
			p-&gt;cond = p;
			switch(o) {
			case ADIV:
				p-&gt;cond = prog_div;
				p-&gt;to.sym = sym_div;
				break;
			case ADIVU:
				p-&gt;cond = prog_divu;
				p-&gt;to.sym = sym_divu;
				break;
			case AMOD:
				p-&gt;cond = prog_mod;
				p-&gt;to.sym = sym_mod;
				break;
			case AMODU:
				p-&gt;cond = prog_modu;
				p-&gt;to.sym = sym_modu;
				break;
			}

			/* MOV REGTMP, b */
			q = prg();
			q-&gt;link = p-&gt;link;
			p-&gt;link = q;
			p = q;

			p-&gt;as = AMOVW;
			p-&gt;line = q1-&gt;line;
			p-&gt;from.type = D_REG;
			p-&gt;from.reg = prog_div != UP &amp;&amp; prog_div-&gt;from.sym-&gt;thumb ? REGTMPT : REGTMP;
			p-&gt;from.offset = 0;
			p-&gt;to.type = D_REG;
			p-&gt;to.reg = q1-&gt;to.reg;

			/* ADD $8,SP */
			q = prg();
			q-&gt;link = p-&gt;link;
			p-&gt;link = q;
			p = q;

			p-&gt;as = AADD;
			p-&gt;from.type = D_CONST;
			p-&gt;from.reg = NREG;
			p-&gt;from.offset = 8;
			p-&gt;reg = NREG;
			p-&gt;to.type = D_REG;
			p-&gt;to.reg = REGSP;

			/* SUB $8,SP */
			q1-&gt;as = ASUB;
			q1-&gt;from.type = D_CONST;
			q1-&gt;from.offset = 8;
			q1-&gt;from.reg = NREG;
			q1-&gt;reg = NREG;
			q1-&gt;to.type = D_REG;
			q1-&gt;to.reg = REGSP;

			break;
		case AMOVW:
			if(thumb){
				Adr *a = &amp;p-&gt;from;

				if(a-&gt;type == D_CONST &amp;&amp; ((a-&gt;name == D_NONE &amp;&amp; a-&gt;reg == REGSP) || a-&gt;name == D_AUTO || a-&gt;name == D_PARAM) &amp;&amp; (a-&gt;offset &amp; 3))
					diag(&#34;SP offset not multiple of 4&#34;);
			}
			break;
		case AMOVB:
		case AMOVBU:
		case AMOVH:
		case AMOVHU:
			if(thumb){
				if(p-&gt;from.type == D_OREG &amp;&amp; (p-&gt;from.name == D_AUTO || p-&gt;from.name == D_PARAM || (p-&gt;from.name == D_CONST &amp;&amp; p-&gt;from.reg == REGSP))){
					q = prg();
					*q = *p;
					if(p-&gt;from.name == D_AUTO)
						q-&gt;from.offset += autosize;
					else if(p-&gt;from.name == D_PARAM)
						q-&gt;from.offset += autosize+4;
					q-&gt;from.name = D_NONE;
					q-&gt;from.reg = REGTMPT;
					p = movrr(p, REGSP, REGTMPT, p);
					q-&gt;link = p-&gt;link;
					p-&gt;link = q;
				}
				if(p-&gt;to.type == D_OREG &amp;&amp; (p-&gt;to.name == D_AUTO || p-&gt;to.name == D_PARAM || (p-&gt;to.name == D_CONST &amp;&amp; p-&gt;to.reg == REGSP))){
					q = prg();
					*q = *p;
					if(p-&gt;to.name == D_AUTO)
						q-&gt;to.offset += autosize;
					else if(p-&gt;to.name == D_PARAM)
						q-&gt;to.offset += autosize+4;
					q-&gt;to.name = D_NONE;
					q-&gt;to.reg = REGTMPT;
					p = movrr(p, REGSP, REGTMPT, p);
					q-&gt;link = p-&gt;link;
					p-&gt;link = q;
					if(q-&gt;to.offset &lt; 0 || q-&gt;to.offset &gt; 255){	// complicated
						p-&gt;to.reg = REGTMPT+1;			// mov sp, r8
						q1 = prg();
						q1-&gt;line = p-&gt;line;
						q1-&gt;as = AMOVW;
						q1-&gt;from.type = D_CONST;
						q1-&gt;from.offset = q-&gt;to.offset;
						q1-&gt;to.type = D_REG;
						q1-&gt;to.reg = REGTMPT;			// mov $o, r7
						p-&gt;link = q1;
						q1-&gt;link = q;
						q1 = prg();
						q1-&gt;line = p-&gt;line;
						q1-&gt;as = AADD;
						q1-&gt;from.type = D_REG;
						q1-&gt;from.reg = REGTMPT+1;
						q1-&gt;to.type = D_REG;
						q1-&gt;to.reg = REGTMPT;			// add r8, r7
						p-&gt;link-&gt;link = q1;
						q1-&gt;link = q;
						q-&gt;to.offset = 0;				// mov* r, 0(r7)
						/* phew */
					}
				}
			}
			break;
		case AMOVM:
			if(thumb){
				if(p-&gt;from.type == D_OREG){
					if(p-&gt;from.offset == 0)
						p-&gt;from.type = D_REG;
					else
						diag(&#34;non-zero AMOVM offset&#34;);
				}
				else if(p-&gt;to.type == D_OREG){
					if(p-&gt;to.offset == 0)
						p-&gt;to.type = D_REG;
					else
						diag(&#34;non-zero AMOVM offset&#34;);
				}
			}
			break;
		case AB:
			if(thumb &amp;&amp; p-&gt;to.type == D_OREG){
				if(p-&gt;to.offset == 0){
					p-&gt;as = AMOVW;
					p-&gt;from.type = D_REG;
					p-&gt;from.reg = p-&gt;to.reg;
					p-&gt;to.type = D_REG;
					p-&gt;to.reg = REGPC;
				}
				else{
					p-&gt;as = AADD;
					p-&gt;from.type = D_CONST;
					p-&gt;from.offset = p-&gt;to.offset;
					p-&gt;reg = p-&gt;to.reg;
					p-&gt;to.type = D_REG;
					p-&gt;to.reg = REGTMPT-1;
					q = prg();
					q-&gt;as = AMOVW;
					q-&gt;line = p-&gt;line;
					q-&gt;from.type = D_REG;
					q-&gt;from.reg = REGTMPT-1;
					q-&gt;to.type = D_REG;
					q-&gt;to.reg = REGPC;
					q-&gt;link = p-&gt;link;
					p-&gt;link = q;
				}
			}
			if(seenthumb &amp;&amp; !thumb &amp;&amp; p-&gt;to.type == D_OREG &amp;&amp; p-&gt;to.reg == REGLINK){
				// print(&#34;warn %s:	b	(R%d)	assuming a return\n&#34;, curtext-&gt;from.sym-&gt;name, p-&gt;to.reg);
				p-&gt;as = ABXRET;
			}
			break;
		case ABL:
		case ABX:
			if(thumb &amp;&amp; p-&gt;to.type == D_OREG){
				if(p-&gt;to.offset == 0){
					p-&gt;as = o;
					p-&gt;from.type = D_NONE;
					p-&gt;to.type = D_REG;
				}
				else{
					p-&gt;as = AADD;
					p-&gt;from.type = D_CONST;
					p-&gt;from.offset = p-&gt;to.offset;
					p-&gt;reg = p-&gt;to.reg;
					p-&gt;to.type = D_REG;
					p-&gt;to.reg = REGTMPT-1;
					q = prg();
					q-&gt;as = o;
					q-&gt;line = p-&gt;line;
					q-&gt;from.type = D_NONE;
					q-&gt;to.type = D_REG;
					q-&gt;to.reg = REGTMPT-1;
					q-&gt;link = p-&gt;link;
					p-&gt;link = q;
				}
			}
			break;
		}
	}
}

static void
sigdiv(char *n)
{
	Sym *s;

	s = lookup(n, 0);
	if(s-&gt;type == STEXT){
		if(s-&gt;sig == 0)
			s-&gt;sig = SIGNINTERN;
	}
	else if(s-&gt;type == 0 || s-&gt;type == SXREF)
		s-&gt;type = SUNDEF;
}

void
divsig(void)
{
	sigdiv(&#34;_div&#34;);
	sigdiv(&#34;_divu&#34;);
	sigdiv(&#34;_mod&#34;);
	sigdiv(&#34;_modu&#34;);
}

static void
sdiv(Sym *s)
{
	if(s-&gt;type == 0 || s-&gt;type == SXREF){
		/* undefsym(s); */
		s-&gt;type = SXREF;
		if(s-&gt;sig == 0)
			s-&gt;sig = SIGNINTERN;
		s-&gt;subtype = SIMPORT;
	}
	else if(s-&gt;type != STEXT)
		diag(&#34;undefined: %s&#34;, s-&gt;name);
}

void
initdiv(void)
{
	Sym *s2, *s3, *s4, *s5;
	Prog *p;

	if(prog_div != P)
		return;
	sym_div = s2 = lookup(&#34;_div&#34;, 0);
	sym_divu = s3 = lookup(&#34;_divu&#34;, 0);
	sym_mod = s4 = lookup(&#34;_mod&#34;, 0);
	sym_modu = s5 = lookup(&#34;_modu&#34;, 0);
	if(dlm) {
		sdiv(s2); if(s2-&gt;type == SXREF) prog_div = UP;
		sdiv(s3); if(s3-&gt;type == SXREF) prog_divu = UP;
		sdiv(s4); if(s4-&gt;type == SXREF) prog_mod = UP;
		sdiv(s5); if(s5-&gt;type == SXREF) prog_modu = UP;
	}
	for(p = firstp; p != P; p = p-&gt;link)
		if(p-&gt;as == ATEXT) {
			if(p-&gt;from.sym == s2)
				prog_div = p;
			if(p-&gt;from.sym == s3)
				prog_divu = p;
			if(p-&gt;from.sym == s4)
				prog_mod = p;
			if(p-&gt;from.sym == s5)
				prog_modu = p;
		}
	if(prog_div == P) {
		diag(&#34;undefined: %s&#34;, s2-&gt;name);
		prog_div = curtext;
	}
	if(prog_divu == P) {
		diag(&#34;undefined: %s&#34;, s3-&gt;name);
		prog_divu = curtext;
	}
	if(prog_mod == P) {
		diag(&#34;undefined: %s&#34;, s4-&gt;name);
		prog_mod = curtext;
	}
	if(prog_modu == P) {
		diag(&#34;undefined: %s&#34;, s5-&gt;name);
		prog_modu = curtext;
	}
}

static void
setdiv(int as)
{
	Prog *p = nil;

	switch(as){
	case ADIV: p = prog_div; break;
	case ADIVU: p = prog_divu; break;
	case AMOD: p = prog_mod; break;
	case AMODU: p = prog_modu; break;
	}
	if(p != UP &amp;&amp; thumb != p-&gt;from.sym-&gt;thumb)
		p-&gt;from.sym-&gt;foreign = 1;
}

void
nocache(Prog *p)
{
	p-&gt;optab = 0;
	p-&gt;from.class = 0;
	p-&gt;to.class = 0;
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
